#GO-web编程http.ListenandServe采用默认mux

```sh
func main{
    http.Handle("/hello/golang/",&BaseHandler{})
    http.Handlefunc("/hello/world",func(resp http.ResponseWriter, req *http.Request){
    resp.Write([]byte("hello world"))
    })
    http.ListenandServe(":8080",nil)

}

type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a => ", req.URL.Query().Get("a"))

	resp.Write([]byte("hello golang"))
}
```
把handle变量置为nil,会调用默认的路由，路由就会检索对应的url比如"/hello/golang/(模糊匹配)"或者 "/hello/world"，再用到相应的方法。