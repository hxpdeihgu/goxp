## goxp
golang web框架

## 快速入门

一个经典的Goxp例子如下：

``
go get github.com/hxpdeihgu/goxp
``

1.创建main.go
```go
func main(){
	t:=goxp.New()
	t.Add(new(controller.Index))
	t.Run(":8080")
}
```

2.创建一个controller目录
```go
type Index struct {
	
}
func (Index) Index() string {
	return "Hello World"
}
```
```go
func (Index)Abc(this *goxp.Xp) string{
	return "hello world"
}
```

```go
func (Index) Index2(this *goxp.Xp) {
    this.Rs.Write([]byte("hello world"))
}

func (Index) Index2(rw http.ResponseWriter, r *http.Request) {
    rw.Write([]byte("hello world"))
}
```
然后在浏览器访问`http://localhost:8080`, 将会得到一个对应字符串返回

#模板创建
```go
func (Index) Index2(this *goxp.Xp) {
    this.Data["data"] = "Hello World"
    this.Rander()
}
```
创建一个view目录，新建一个index2.html文件
```go
{.data}
```

继承model方法使用

```go
type Test struct {
	goxp.Model
}

//Test 不能传入指针
func (t Test) Abc(this *goxp.Xp) string {
	return t.Md5("abc")
}


````
#session使用
```go
func main(){
	session.SessionInit()//添加seeion控件
	t:=goxp.New()
	t.Add(new(b.Index))
	t.Run(":8080")
}
```
代码实现
```go
func (Index) Test(this *goxp.Xp) {
	sessions:=session.SessionStart(this)
	//sessions.Add("name","hello world")
	fmt.Println(sessions.Get("name"))
}
```


