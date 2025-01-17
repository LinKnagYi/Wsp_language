package vm

import(
    "strings"
)

func VmReturn(From TransmitValue)string{
    From.VarValue.ResLock=true
    Lids := From.OpRunId
    Op := From.Opcode[Lids+1]
    return CodeBlockRunSingle(Op,From.VarValue)
}

func VmWgo(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    From.VarValue.Govm=false
    From.VarValue.SetWgoId(Op.Abrk[0].Text)
    return ""
}

func VmBreak(From TransmitValue)string{
    From.VarValue.LockBreakList="<BREAK>"
    return ""
}

func VmContinue(From TransmitValue)string{
    From.VarValue.AllCodeStop = true
    From.VarValue.LockBreakList="<CONTINUE>"
    return ""
}

func VmFuncInit(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Name:=Op.Text
    UserFuncInitManual(Name)
    return Name
}

func VmFuncInitBody(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Name:=Op.Name
    UserFuncInitManual_9C(Name)
    return Name
}

func VmGlobal(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Text := Op.Text
    FatherList := strings.Split(Text,",")
    for i:=0;i<=len(FatherList)-1;i++{
        pathMain:=From.VarValue.FILE+"Main/Main"+FatherList[i]
        pathFunc:=From.VarValue.paths+From.VarValue.FuncName+FatherList[i]
        Var_Pointer(From.VarValue.FuncName+FatherList[i],From.VarValue)
        CopyVmArray(pathMain,pathFunc)
    }
    
    return ""
}

func VmClassLock(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    Text := Op.Text
    ClassLock[Text]=true
    return ""
}

func SetEnv(From TransmitValue)string{
    Lids := From.OpRunId
    Op := From.Opcode[Lids]
    FromIp := Op.Abrk[0].Text
    Port := Op.Abrk[1].Text
    FormId := Op.Abrk[2].Text
    if FromIp!="this"{
        Value := ReadEnv(FormId)
        WebSocketWg.Add(2)
        go SyncVar(FromIp,Port,FormId,Value.FILE)
        go SyncVarSever(FromIp,Port,FormId,Value.FILE)
        *From.VarValue = Value
        return "True"
    }
    Value := ReadEnv(FormId)
    if Port!="0"{
        WebSocketWg.Add(2)
        go SververSocketClient(Port,Value.FILE,FormId)
        go SververSocketClientUser(Port,Value.FILE,FormId)
    }
    *From.VarValue = Value
    return "True"
}

func Free(From TransmitValue)string{
    Value := From.Value
    varName :=Value[1:len(Value)]
    File:= So_Array_Stick(From.VarValue.FuncName+varName,From.VarValue)
    Del_Dir(File+"/")
    Del_Dir(File)
    Del_File(File)
    Del_Files(File)
    return "<TRUE>"
}