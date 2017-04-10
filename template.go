package goxp

import (
	"net/http"
	"html/template"
	"path"
)

func rander(w http.ResponseWriter,data interface{},path string){
	t,err:=template.ParseFiles(path)
	if err!=nil {
		Error(err)
	}else {
		t.Execute(w,data)
	}
}

//Rander 模板渲染
func (this *Xp)Rander(paths ...string){
	var p string = pathPrefix
	if pathP,ok:=this.config["pathPrefix"];ok{
		p = pathP
	}
	if len(paths)>0{
		rander(this.Rs,this.Data,p+paths[0])
	}else {
		if ts,ok:=this.config["templateSuffix"];ok{
			rander(this.Rs,this.Data,p+this.thisMethod+ts)
		}else {
			rander(this.Rs,this.Data,path.Join(p,this.thisPackage,this.thisMethod)+templateSuffix)
		}
		
	}
	
}