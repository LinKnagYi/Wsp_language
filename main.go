package main

import(
  "fmt"
  "os"
  "Wsp/Analysis/Lex"
  "Wsp/Analysis/Ast"
  "Wsp/Module/Vm"
  "Wsp/Module/Ini"
  "Wsp/Module/Opcache"
  "Wsp/Module/Memory"
  "Wsp/Module/GC"
  "Wsp/Module/Const"
  "Wsp/Compile"
  "io/ioutil"
  "strings"
  "github.com/peterh/liner"
)
const Version = "4.6.3"
func RunCode(Code string,FilesStruct vm.FileValue)vm.FileValue{
    Opcode:=Compile(Code)
    if ini.DebugsIf(){
        fmt.Println("\n---------------------------------------------------------")
        fmt.Println("Opcode:")
        fmt.Println("---------------------------------------------------------")
        for i:=0;i<=len(Opcode.Body)-1;i++{
            Codes:=Opcode.Body[i]
            fmt.Println("第",i,"段:  ID  Type    Value    Name    Text    Mov    Line")
            for z:=0;z<=len(Codes)-1;z++{
                fmt.Println("          ",z,Codes[z])
            }
            fmt.Println("---------------------------------------------------------")
        }
    }
    return vm.WspVmConsole(Opcode,FilesStruct)
}
func Compile(Code string)compile.Res_Struct{
    //fmt.Println(lex.Wsp_Lexical(string(Code+"\n ")))
    Code = strings.Replace(Code, "\r", "", -1)
    return compile.Wsp_Compile(ast.Wsp_Ast(lex.Wsp_Lexical(string(Code+"\n "))))
}

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
var Files_s string
func ConsoleCodeCheck(Code string)bool{
    Tlex := lex.Wsp_Lexical(string(Code+"\n "))
    JsNum := 0
    for i:=0;i<=len(Tlex)-1;i++{
        if Tlex[i].Type==75{
            JsNum++
        }else if Tlex[i].Type==76{
            JsNum--
        }
    }
    if JsNum==0{
        return true
    }
    return false
}
func main(){
    if len(os.Args)==1{
        fmt.Println("Wsp Console [版本 "+Version+"]")
        fmt.Println("版权所有 (c) 2022 WspLang.com 保留部分权利。")
        Inis:=ini.ReadIni()
        if Inis["wsp_var_ram"]=="1"{
            vm.VarRamStart()
        }
        vm.VmStart()
        gc.SetGcSize(Inis["wsp_gc_size"])
        go gc.GC_Runtime()
        FilesStruct := vm.FileValue{}
        var TmpCode string
        MapsLexTab:=lex.TabReturn();
        for k,_ := range MapsLexTab{
            vm.TabStruct.Add(k)
        }
        
        for{
            printfText := "" 
            if ConsoleCodeCheck(TmpCode){
                printfText = ">>>"
            }else{
                printfText = "   "
            }
            lineTmp := liner.NewLiner()
            
            lineTmp.SetCompleter(func(line string) (c []string) {
                if line==""{
                    return 
                }
                Maps:=vm.TabStruct.Read()
                for _, n := range Maps {
                    if strings.HasPrefix(n, strings.ToLower(line)) {
                        c = append(c, n)
                    }
                }
                return
            })
            
            Codebyte,_:= lineTmp.Prompt(printfText);
            lineTmp.Close()
            Code := string(Codebyte)
            if ConsoleCodeCheck(TmpCode+Code){
                FilesStruct = RunCode(TmpCode+Code+"\n",FilesStruct)
                TmpCode = ""
            }else{
                TmpCode += Code+"\n"
            }
        }
        os.Exit(0)
    }
    file := ""
    str, _ := os.Getwd()
    if ok,_ := PathExists(str+"/"+os.Args[1]); ok {
        file = str+"/"+os.Args[1]
    }else if ok,_ := PathExists(os.Args[1]); ok {
        file = os.Args[1]
    }else if os.Args[1] == "version"{
        fmt.Println("Version    V"+Version+"\nOpcache    V1.1.0\nVarCache   V1.0.0\nWspGc      V1.1.1")
        os.Exit(0)
    }else if os.Args[1] == "help"{
        if len(os.Args)==2{
            fmt.Println("Wsp 是用来运行Wsp语言代码的解释器\n")
            fmt.Println("使用方法：\n\n        wsp <路径>")
            fmt.Println("        wsp version 查看版本\n")
            fmt.Println("有关该主题的更多信息，请使用“go help OR wsp help ini”。\n")
            os.Exit(0)
        }else if os.Args[2] == "ini"{
            fmt.Println("wsp_func_del   用来禁用函数，用“,”隔开")
            fmt.Println("wsp_debug      用来显示OPCODE数据 默认0关闭 1开启")
            fmt.Println("wsp_cache      用来开启OPCODE缓存 默认0关闭 1开启")
            fmt.Println("wsp_cache_file 用来设置OPCODE缓存存储路径")
            fmt.Println("wsp_var_ram    用来开启VarCache缓存 默认1开启 0关闭")
            fmt.Println("extension      用来载入SO动态库扩展，需使用绝对路径")
        }
    }else{
        fmt.Println("文件或路径不存在")
        fmt.Println(str+"/"+os.Args[1])
        fmt.Println(os.Args[1])
        os.Exit(0)
    }
    consts.WspConst.SetWspFile(file)
    data, _ := ioutil.ReadFile(file)
    vm.WspCodeFileSet(file)
    files := string(data)
    Inis:=ini.ReadIni()
    if Inis["wsp_var_ram"]=="1"{
        vm.VarRamStart()
    }
    vm.VmStart()
    gc.SetGcSize(Inis["wsp_gc_size"])
    go gc.GC_Runtime()
    cache_file:=Inis["wsp_cache_file"]
    var Opcode compile.Res_Struct
    if ok,_ := PathExists(cache_file+"/"+op.Md5(files)); ok  && Inis["wsp_cache"]=="1"{
        TmpS:=op.Opcaches_Read(cache_file+"/"+op.Md5(files))
        center.S_Memory_FromMap(TmpS.FuncList)
        Opcode=TmpS.Opcode
        vm.Wsp_Vm(TmpS.Opcode)
        
    }else{
        Buildse:=Compile(files)
        if Inis["wsp_cache"]=="1"{
            Ops:=op.Opcodes{Buildse,center.R_Memory_FromMap()}
            op.Opcaches_ADD(Ops,cache_file+"/"+op.Md5(files))
        }
        Opcode=Buildse
        vm.Wsp_Vm(Buildse)
    }
    if ini.DebugsIf(){
        fmt.Println("\n---------------------------------------------------------")
        fmt.Println("Opcode:")
        fmt.Println("---------------------------------------------------------")
        for i:=0;i<=len(Opcode.Body)-1;i++{
            Codes:=Opcode.Body[i]
            fmt.Println("第",i,"段:  ID  Type    Value    Name    Text    Mov    Line")
            for z:=0;z<=len(Codes)-1;z++{
                fmt.Println("          ",z,Codes[z])
            }
            fmt.Println("---------------------------------------------------------")
        }
    }
    gc.Gc_Ends()
    vm.VmEnd()
}