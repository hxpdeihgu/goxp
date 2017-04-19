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
	t.Add(new(controller.Index)))
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

###列出所有服务器地址，nginx 自动均衡分发请求到各个服务器。 
 ```nginx
upstream frontends {    
    ip_hash;  
    server 192.168.199.1:8088;
    server 192.168.199.2:8089;
}
server {
    listen      80; 
    server_name mydomain.com www.mydomain.com;
    location / {
        proxy_pass_header Server;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Scheme $scheme;
        proxy_pass http://frontends;
    }
     
    #静态资源交由nginx管理
    location /static {
        root        /var/www/mydomain/web;
        expires     1d;
        add_header  Cache-Control public;
        access_log  off;
    }
}`
```
//this host ip 192.168.199.1
```go
func main() {
    ...
    t:=goxp.New()
    t.Add(new(b.Index))
    t.Run(":8080")
}
 
...
//other
//this host ip 192.168.199.2
func main() {
    ...
	t:=goxp.New()
	t.Add(new(b.Index))
	t.Run(":8080")
}
```

