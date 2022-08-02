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
  "Wsp/Compile"
  "io/ioutil"
)

func RunCode(Code string){
    vm.Wsp_Vm(Compile(Code))
}
func Compile(Code string)compile.Res_Struct{
    return compile.Wsp_Compile(ast.Wsp_Ast(lex.Wsp_Lexical(string(Code+" "))))
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
func main(){
    if len(os.Args)==1{
        fmt.Println("文件或路径不存在")
        os.Exit(0)
    }
    file := ""
    str, _ := os.Getwd()
    if ok,_ := PathExists(str+"/"+os.Args[1]); ok {
        file = str+"/"+os.Args[1]
    }else if ok,_ := PathExists(os.Args[1]); ok {
        file = os.Args[1]
    }else if os.Args[1] == "version"{
        fmt.Println("Version: V3.0\nOpcache: V1.0")
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
            fmt.Println("wsp_debug      用来显示OPCODE数据 默认0关闭 1为开启")
            fmt.Println("wsp_cache      用来开启OPCODE缓存 默认1开启 0关闭")
            fmt.Println("wsp_cache_file 用来设置OPCODE缓存存储路径")
            fmt.Println("extension      用来载入SO动态库扩展，需使用绝对路径")
        }
    }else{
        fmt.Println("文件或路径不存在")
        fmt.Println(str+"/"+os.Args[1])
        fmt.Println(os.Args[1])
        os.Exit(0)
    }
    
    data, _ := ioutil.ReadFile(file)
    files := string(data)
    Inis:=ini.ReadIni()
    cache_file:=Inis["wsp_cache_file"]
    var Opcode compile.Res_Struct
    if ok,_ := PathExists(cache_file+"/"+op.Md5(files)); ok  && Inis["wsp_cache"]=="1"{
        TmpS:=op.Opcaches_Read(cache_file+"/"+op.Md5(files))
        center.S_Memory_FromMap(TmpS.FuncList)
        Opcode=TmpS.Opcode
        if len(Compile(files).Body)==len(TmpS.Opcode.Body){
            vm.Wsp_Vm(TmpS.Opcode)
        }
        
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
        for i:=0;i<=len(Opcode.Body)-1;i++{
            fmt.Println(i,Opcode.Body[i])
        }
    }
}