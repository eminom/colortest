

package main

import (
    "fmt"
    "syscall"
    "strconv"
)

type SetColor struct {
    stdHandle uintptr
    setConsoleTextAttr *syscall.LazyProc
}

func NewSetColor()*SetColor{
    var c int = -11
    od := syscall.NewLazyDLL("kernel32.dll")
    stdHandle, _, _ := od.NewProc("GetStdHandle").Call(uintptr(c))
    setConsoleTextAttr := od.NewProc("SetConsoleTextAttribute")
    return &SetColor{
        stdHandle:stdHandle,
        setConsoleTextAttr:setConsoleTextAttr,
    }
}

func (s *SetColor)SetColor(color int){
    s.setConsoleTextAttr.Call(s.stdHandle, uintptr(color)) 
}

func (s *SetColor)Reset(){
    s.setConsoleTextAttr.Call(s.stdHandle, uintptr(7))
}

func main() {
    set := NewSetColor()
    for i:=0;i<16;i++{
        set.SetColor(i)
        fmt.Println(strconv.Itoa(i)+`:hola`)
    }
}