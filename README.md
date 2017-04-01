# goxp
golang web框架
#使用

1.创建main.go
#
```
func main(){
	t:=goxp.New()
	t.Add(new(controller.Index))
	t.Run(":8080")
}
```

2.创建一个controller目录
#
```
type Index struct {
	
}
func (Index) Index() string {
	return "Hello World"
}
```
