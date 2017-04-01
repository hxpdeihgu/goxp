# goxp
golang web框架
#使用
1.创建main.go
```
func main(){
	t:=goxp.New()
	t.Add(new(controller.Index))
	t.Run(":8080")
}
```

2.创建一个controller目录
```
type Index struct {
	
}
func (Index) Index() string {
	return "Hello World"
}
```
```
func (Index)Abc(this *goxp.Xp) string{
	return "hello world"
}
```

```
func (Index) Index2(this *goxp.Xp) {
    this.Rs.Write([]byte("hello world"))
}

func (Index) Index2(rw http.ResponseWriter, r *http.Request) {
    rw.Write([]byte("hello world"))
}
```

#模板创建
```
func (Index) Index2(this *goxp.Xp) {
    this.Data["data"] = "Hello World"
    this.Rander()
}
```
创建一个view目录，新建一个index2.html文件
```
{.data}
```

#session使用
```
func main(){
	session.SessionInit()//添加seeion控件
	t:=goxp.New()
	t.Add(new(b.Index))
	t.Run(":8080")
}
```
代码实现
```
func (Index) Test(this *goxp.Xp) {
	sessions:=session.SessionStart(this)
	//sessions.Add("name","hello world")
	fmt.Println(sessions.Get("name"))
}
```

