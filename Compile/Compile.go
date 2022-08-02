package compile

import(
  "Wsp/Analysis/Ast"
  "Wsp/Module/Memory"
)

func Wsp_Compile(Codes ast.Ast_Tree)Res_Struct{
    MemoryList:=center.R_Memory_FromMap()
    Code = Codes
    Res:=Res_Struct{}
    Res.Body=Wsp_Compile_l(Codes.BodyAst)
    
    Funcs:=make(map[string]map[int]map[int]Body_Struct_Run)
    for i:=0;i<=len(MemoryList)-1;i++{
        Name:=MemoryList[i]
        Funcs[Name]=Wsp_Compile_l(Codes.FuncAst.FuncList[Name])
    }
    
    Res.Func=Func_Struct{Funcs,Codes.FuncAst.FuncVars}
    
    return Res
}

func Wsp_Compile_l(TCode map[int]ast.BodyAst_Struct)map[int]map[int]Body_Struct_Run{
    Res:=make(map[int]map[int]Body_Struct_Run)
    Len_Line := 0
    Res[Len_Line]=make(map[int]Body_Struct_Run)
    for i:=0;i<=len(TCode)-1;i++{
        switch TCode[i].Type{
            case 50:
                if TCode[i+1].Type==95{
                    tmp:=0
                    for z:=i;z<=len(TCode)-1;z++{
                        if TCode[z].Type!=50&&TCode[z].Type!=95{
                            break
                        }else if TCode[z].Type==50&&TCode[z+1].Type==95{
                            Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                                Type : 301,
                                Abrk : TCode[z].Abrk,
                                Name : "SET_VAR",
                                Text : TCode[z].Text,
                                Movs : "<NIL>",
                                Line : TCode[z].Line,
                            }
                        }else{
                            break
                        }
                        tmp++
                    }
                    i+=tmp-1
                }else{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 302,
                        Abrk : TCode[i].Abrk,
                        Name : "SO_VAR",
                        Text : TCode[i].Text,
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }
            case 7:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 200,
                    Abrk : TCode[i].Abrk,
                    Name : "PRINT",
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 80:
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 2:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 201,
                    Abrk : TCode[i].Abrk,
                    Name : "FOR",
                    Text : TCode[i].Xbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 3:
                for z:=i;z<=len(TCode)-1;z++{
                    if TCode[z].Type==3 || TCode[z].Type==4 || TCode[z].Type==5{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 199+TCode[z].Type,
                            Abrk : TCode[z].Abrk,
                            Name : TCode[z].Name,
                            Text : TCode[z].Xbrk[0],
                            Movs : "<NIL>",
                            Line : TCode[z].Line,
                        }
                    }else{
                        break
                    }
                }
                Len_Line++
                Res[Len_Line]=make(map[int]Body_Struct_Run)
            case 0:
                if len(TCode[i].Sbrk)>0{
                    Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                        Type : 205,
                        Abrk : TCode[i].Abrk,
                        Name : TCode[i].Text,
                        Text : TCode[i].Sbrk[0],
                        Movs : "<NIL>",
                        Line : TCode[i].Line,
                    }
                }else if TCode[i].Text!=""{
                    _,ok:=Code.FuncAst.FuncVars[TCode[i].Text]
                    if string(TCode[i].Text[0])=="\""||ast.IsNum(TCode[i].Text)||ok{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 0,
                            Name : TCode[i].Name,
                            Text : TCode[i].Text,
                            Movs : "<NIL>",
                            Line : TCode[i].Line,
                        }
                    }else if _,ok:=Code.FuncAst.FuncList[TCode[i].Text];ok{
                        Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                            Type : 206,
                            Name : TCode[i].Name,
                            Text : TCode[i].Text,
                            Movs : "<NIL>",
                            Line : TCode[i].Line,
                        }
                    }
                }
            case 11:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 207,
                    Abrk : TCode[i].Abrk,
                    Name : "RETURN",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 8:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 208,
                    Name : "WGO",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 9:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 209,
                    Abrk : TCode[i].Abrk,
                    Name : "ADD",
                    Text : TCode[i].Sbrk[0],
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 6:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 210,
                    Name : "BREAK",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
            case 10:
                Res[Len_Line][len(Res[Len_Line])]=Body_Struct_Run{
                    Type : 211,
                    Name : "CONTINUE",
                    Text : "",
                    Movs : "<NIL>",
                    Line : TCode[i].Line,
                }
        }
    }
    Check(Res)
    return Res
}