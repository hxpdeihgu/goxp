package goxp

import (
	"reflect"
	"sync"
)

//Data 存放模板值
type Data map[string]interface{}

//errFunc404处理函数
type ErrFunc func(*Xp)

//panicFunc 错误处理函数
type PanicFunc func(*Xp)

//Trees存储定义的结构体
type trees struct {
	sync.RWMutex
	c map[string]*Controller
}

//配置项目
type Config map[string]string

//Controller 控制器信息
type Controller struct{
	method string
	Tv *reflect.Value
}




