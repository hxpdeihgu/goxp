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
	x.config = make(map[string]string)
	return x
}

func Run(this *Xp){
	if addr,ok:=this.config["addr"];ok{
		log.Fatal(http.ListenAndServe(addr,this))
	}
	log.Fatal(http.ListenAndServe(addr,this))
}
func (this *Xp) Run(addr string){
	log.Fatal(http.ListenAndServe(addr,this))
}

func (this *Xp) Redirect(url string){
	http.Redirect(this.Rs,this.Rq,url,301)
}

func (this *Xp) SetConfig(key,value string)  {
	this.Lock()
	defer this.Unlock()
	this.config[key] = value
}

func (this *Xp) GetConfig(key string) string {
	this.RLock()
	defer this.RUnlock()
	if v,ok:=this.config[key];ok {
		return v
	}
	return ""
}
