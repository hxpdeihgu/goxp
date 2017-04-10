package goxp

import (
	"net/http"
	"reflect"
	"github.com/pkg/errors"
	"strings"
	"fmt"
	"encoding/json"
)
const (
	argumentNum = iota
	argumentNum1
	argumentNum2
)
type Xp struct {
	Rs http.ResponseWriter
	Rq *http.Request
	thisMethod string
	thisPackage string
	Data
	trees
	Config
	ErrFunc
	PanicFunc
	ErrMsg error
}

//路由设置
func (this *Xp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			Error(err)
			if this.PanicFunc == nil{
				http.NotFound(rw,r)
				return
			}
			this.PanicFunc(this)
		}
	}()
	this.restXp(rw,r)
	c,err := getController(this,r)
	if err==nil{
		method := c.Tv.MethodByName(c.method)
		if method.IsValid() {
			bf:= c.Tv.MethodByName(BF)
			if bf.IsValid() {
				in:=make([]reflect.Value,argumentNum)
				in = append(in,reflect.ValueOf(this))
				this.returnFun(bf,in)
			}
			num:=method.Type().NumIn()
			switch num {
			case argumentNum:
				this.returnFun(method,nil)
			case argumentNum1:
				in:=make([]reflect.Value,argumentNum)
				in = append(in,reflect.ValueOf(this))
				this.returnFun(method,in)
			case argumentNum2:
				in:=make([]reflect.Value,argumentNum)
				in = append(in,reflect.ValueOf(rw))
				in = append(in,reflect.ValueOf(r))
				this.returnFun(method,in)
			default:
				this.ErrMsg = errors.New("Parameter is not in conformity with the requirements")
				if this.ErrFunc == nil{
					http.NotFound(rw,r)
					break
				}
				this.ErrFunc(this)
			}
			af:= c.Tv.MethodByName(AF)
			if af.IsValid() {
				in:=make([]reflect.Value,argumentNum)
				in = append(in,reflect.ValueOf(this))
				this.returnFun(af,in)
			}
		}else {
			if this.ErrFunc == nil {
				http.NotFound(rw,r)
				return
			}
			this.ErrMsg = errors.New("404 error")
			this.ErrFunc(this)
		}
		
	}else {
		if this.ErrFunc == nil{
			http.NotFound(rw,r)
			return
		}
		this.ErrMsg = errors.New("404 error")
		this.ErrFunc(this)
	}
	
}
func (this *Xp) returnFun(method reflect.Value,in []reflect.Value)  {
	v:=method.Call(in)
	if len(v)>0 {
		b,err:=json.Marshal(v[argumentNum].Interface())
		if err==nil {
			this.Rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
			this.Rs.Write(b)
		}
	}
}
//getReflectTypeValue 获取控制器
func getController(this *Xp,r *http.Request) (*Controller,error) {
	this.RLock()
	defer this.RUnlock()
	
	path:=strings.TrimRight(r.URL.Path,separator)
	methodIndex:=strings.LastIndex(path,separator)
	if c,ok:=this.trees.c[path[:methodIndex]];ok {
		this.thisMethod = path[methodIndex+1:]
		this.thisPackage = path[:methodIndex]
		c.method = strings.Title(path[methodIndex+1:])
		return c,nil
	}
	return nil,errors.New("not found is controller")
}

//重新赋值
func (this *Xp) restXp(rw http.ResponseWriter, r *http.Request) {
	this.Rs = rw
	this.Rq = r
	this.Data = make(map[string]interface{})
	this.thisMethod = ""
}

//addType添加类型
func (this *Xp) addType(d interface{}) {
	rv:=reflect.Indirect(reflect.ValueOf(d))
	rt:= rv.Type()
	this.Lock()
	defer this.Unlock()
	structName := getPath(rt.PkgPath(),rt.Name(),this)
	
	if _,ok:=this.c[structName];ok{
		panic(fmt.Sprint(structName," is exist;"))
	}
	
	c:=new(Controller)
	c.Tv = &rv;
	this.c[structName] = c
}

//getPath 组合访问地址
func getPath(pkg string,name string,this *Xp) string {
	if prefix,ok:= this.Config[controller];ok{
		pkg=strings.TrimPrefix(pkg,prefix)
	}else {
		pkg=strings.TrimPrefix(pkg,controller)
	}
	
	name=strings.ToLower(name)
	return (pkg+separator+name)
}