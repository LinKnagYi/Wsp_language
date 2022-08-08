<p align="center"><img src="./logo.png"
        alt="Logo" width="128" height="128" style="max-width: 100%;"></p>
<h1 align="center">WSP</h1>
<p align="center">一门解释型语言</p>
<p align="center">
    <a href="https://github.com/Linkangyis/Wsp_language/blob/LICENSE">
        <img src="https://img.shields.io/github/license/Ice-Hazymoon/MikuTools.svg" alt="MIT License" />
    </a>
</p>

## 安装WSP

```bash
vi /ect/profile
export WSPPATH=WSP所在目录
```
```bash
ln -s WSP所在目录/wsp /usr/bin
```
## 介绍

基于golang开发的解释型语言 使用wsp虚拟机，效率极高，当前版本 v3.1.2-beta.6,有PHP的简单 Python的实用 Golang的效率

## 开发

```bash
wsp ./xxxx.wsp
```

## 语法
自定义函数
```php
function 变量名(参数){
    //代码块
}
```
自定义变量
```php
$xx=xx;
```
循环
```php
xx;xx;xx形式
for(条件){
    //代码块
}
```
```php
while形式
for(条件){
    //代码块
}
```
```php
do_while形式
for{
    //代码块
}(条件)
```
```php
死循环
for{
    //代码块
}
```
判断
```php
if(条件){
    //代码块
}else if(条件){
    //代码块
}else{
    //代码块
}
```
Switch语句
```php
$a = "3";
switch($a){
    case 1:
        print(1);
    case 2:
        print(2);
    case 3:
        print(3);
    default:
        print(4);
}
```
wgo协程
```php
wgo func();//参数有特殊处理，与非协程有很大变化，在运行前就会先将变量解析为静态值
```
## 扩展开发
```golang
package main

import(
  "fmt"
)
func Func_Info()(map[int]string){  //系统核心
    info := make(map[int]string)  //函数列表
    info[0] = "Testb"
    info[1] = "Tests"
    return info
}
func Package_Info()(string){  //系统核心
    info := "Test"   //包名设置
    return info
}


func Testb(Value map[int]string)(string){    //Testb扩展函数 wsp调用 Test.Testb()
    fmt.Println("b")
    return "TRUE"
}

func Tests(Value map[int]string)(string){    //Tests扩展函数 wsp调用 Test.Tests()
    fmt.Println("s")
    return "TRUE"
}
//扩展编译指令
//go build -buildmode=plugin -o test.so Test.go
```
## License

[MIT](https://github.com/Linkangyis/Wsp_language/blob/LICENSE)
