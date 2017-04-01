package goxp

import (
	"os"
	"io"
	"log"
)

func logMsg(prefix string) *log.Logger {
	return log.New(os.Stdout,prefix, log.Lshortfile|log.LstdFlags)
}

func NewError(out ...io.Writer) *log.Logger {
	if len(out)>0{
		return log.New(out[0],"Error: ", log.Lshortfile|log.LstdFlags)
	}
	return log.New(os.Stdout,"Error: ", log.Lshortfile|log.LstdFlags)
}

func NewWarning(out ...io.Writer) *log.Logger {
	if len(out)>0{
		return log.New(out[0],"Warning: ", log.Lshortfile|log.LstdFlags)
	}
	return log.New(os.Stdout,"Warning: ", log.Lshortfile|log.LstdFlags)
}

func NewNotice(out ...io.Writer) *log.Logger {
	if len(out)>0{
		return log.New(out[0],"Notice: ", log.Lshortfile|log.LstdFlags)
	}
	return log.New(os.Stdout,"Notice: ", log.Lshortfile|log.LstdFlags)
}

func Error(msg interface{}){
	l:=logMsg("Error: ")
	l.Println(msg)
}

func Warning(msg interface{}){
	l:=logMsg("Warning: ")
	l.Println(msg)
}

func Notice(msg interface{}){
	l:=logMsg("Notice: ")
	l.Println(msg)
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
