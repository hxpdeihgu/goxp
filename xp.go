package goxp

import (
	"log"
	"net/http"
)

//Add设置类型
func(this *Xp) Add(d interface{}) {
	this.addType(d)
}

func New() *Xp{
	x:=new(Xp)
	x.c = make(map[string]*Controller)
	x.Config = make(map[string]string)
	return x
}

func Run(this *Xp){
	if addr,ok:=this.Config["addr"];ok{
		log.Fatal(http.ListenAndServe(addr,this))
	}
	log.Fatal(http.ListenAndServe(addr,this))
}
func (this *Xp) Run(addr string){
	log.Fatal(http.ListenAndServe(addr,this))
}
