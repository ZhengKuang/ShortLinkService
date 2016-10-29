#中间件的实现方法
```sh
func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()

	muxHandler := mux.NewMuxHandler()
	muxHandler.Handle("/hello/golang/", &BaseHandler{})
	muxHandler.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})

	log.Println("ShortURl server will start at port" + port)
	log.Fatalln(http.ListenAndServe(port, muxHandler))
}

type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a => ", req.URL.Query().Get("a"))

	resp.Write([]byte("hello golang"))
}

```
从这个代码可以看出，ListenAndServe方法传送了一个muxHandler这个struct（interface）进去，里面包含ServeHTTP这个方法，而且
muxHandler重新定义了Handle和HandleFunc的处理方式，并且在main里面对相应的函数进行初始化，BaseHandler的ServeHTTP是不会被
执行的，因为他没有被传进去