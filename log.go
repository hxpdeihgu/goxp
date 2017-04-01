package goxp

import (
	"os"
	"io"
	"log"
)

func logMsg(msg interface{},prefix string)  {
	logger := log.New(os.Stdout,prefix, log.Lshortfile|log.LstdFlags)
	logger.Println(msg)
}

func Error(msg interface{}){
	logMsg(msg,"Error: ")
}

func Warning(msg interface{}){
	logMsg(msg,"Warning: ")
}

func Notice(msg interface{}){
	logMsg(msg,"Notice: ")
}

func msgs(out io.Writer,msg string,prefix string){
	logger := log.New(out,prefix, log.Lshortfile)
	logger.Print(msg)
}

func Msgs(w io.Writer,msg string,prefix string){
	msgs(w,msg,prefix)
}

func Msg(w io.Writer,msg string){
	msgs(w,msg,"")
}
