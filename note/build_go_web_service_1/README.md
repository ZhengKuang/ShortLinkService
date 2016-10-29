
#基于Go语言的短链接服务---实战笔记

##Go搭建web服务：go如何搭建web服务呢？

Go有一个源包：net/http包，在这个包里面有一个listenandServe方法，调用包以后,利用
http.listenandServe()方法可以监听一个端口

http.ListenAndServe(":3000", nil):表示监听3000端口


ListenAndServe代码分析
```sh
func ListenAndServe(addr string, handler Handler) error {
	// 创建一个Server结构体，调用该结构体的ListenAndServer方法然后返回
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```
调用结构体Server的ListenAndServe()方法，并且初始化Server结构体变量serve的addr和handler，比如在http.ListenAndServe(":3000", nil)这个例子里面，addr为":3000"，handler为nil

```sh
type Server struct {
	Addr           string        // 服务器的IP地址和端口信息
	Handler        Handler       // 请求处理函数的路由复用器
	ReadTimeout    time.Duration 
	WriteTimeout   time.Duration
	MaxHeaderBytes int       
	TLSConfig      *tls.Config  
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
	disableKeepAlives int32 
}


func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"  // 如果不指定服务器地址信息，默认以":http"作为地址信息
	}
	ln, err := net.Listen("tcp", addr)    // 这里创建了一个TCP Listener，之后用于接收客户端的连接请求
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})  // 调用Server.Serve()函数并返回
}


```
srv为Server这个结构体的首地址，结合上文，意思是检查server这个结构体变量里面的addr是否为空，为空就附上默认值为":http"也就是":80"，并且建立一个端口监听addr，在这个例子里面，监听的是3000端口，监听变量为ln，并将ln传入tcpKeepAliveListener方法中作为变量，调用server.Serve方法
```sh
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration 
	// 这个循环就是服务器的主循环了，通过传进来的listener接收来自客户端的请求并建立连接，
	// 然后为每一个连接创建routine执行c.serve()，这个c.serve就是具体的服务处理了
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve() // <-这里为每一个建立的连接创建routine之后进行服务
	}
}
```
这段代码的意思是，ln这个变量是监听3000端口的变量，传入Serve方法中，ln会一直Accept新的请求(这个死循环在互联网应用里面也有学到)，并且返回每一次请求对应的变量rw,如果正常的话，会产生请求对应的连接c，然后连接c调用serve方法
```sh
func (c *conn) serve() {
	origConn := c.rwc // copy it before it's set nil on Close or Hijack

	// 这里做了一些延迟释放和TLS相关的处理...
	
	// 前面的部分都可以忽略，这里才是主要的循环
	for {
		w, err := c.readRequest()  // 读取客户端的请求
		// ...
		serverHandler{c.server}.ServeHTTP(w, w.req) //这里对请求进行处理
		if c.hijacked() {
			return
		}
		w.finishRequest()
		if w.closeAfterReply {
			if w.requestBodyLimitHit {
				c.closeWriteAndWait()
			}
			break
		}
		c.setState(c.rwc, StateIdle)
	}
}
```
从这段代码分析可以得出，c（connection）是一个结构体，有一个方法叫做serve方法，在这个方法中
，c会读取客户端请求（GET，POST，PUSH）请求并且调用**serveHandler.ServeHTTP**这个无比重要的方法来进行相应的处理


总之，调用http.ListenAndServe(":3000", nil)，会创建一个监听3000端口的变量，并且调用serveHandler.ServeHTTP这个方法，至于这个方法到底在哪里，我们接下来会分析
```sh
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```
```sh
func ListenAndServe(addr string, handler Handler) error {
	// 创建一个Server结构体，调用该结构体的ListenAndServer方法然后返回
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```
对比这两段代码，我们可以得知，Handler是一个interface，里面有ServeHttp这个方法，所以如果我们需要传入一个Interface的实例变量handler，我们只要复写一个struct，并且在struct里面的变量复写一个方法即可，再&变量取地址，就相当于把方法传了进去，如果以后3000端口还有新的request，我们就会调用复用的方法
```sh
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
type Request struct {
	Method string
	URL *url.URL
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	Header Header
	Body io.ReadCloser
	TransferEncoding []string
	Host string
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form
	Trailer Header
	RemoteAddr string
	RequestURI string
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
	ctx context.Context
}


type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}
```
从ServeHTTP的两个变量ResponseWriter和Request分析可以知道，Request是一个结构体，里面有Method（Get，Put，Post），有对应的URL等等。
RespenseWriter是一个interface，interface的方法需要复写一下才能用。
注意：interface其实和struct是一样的，方法都要复写，调用的时候都是struct用&实例变量的感觉，，interface不需要实例变量，直接&struct就可以了，interface和struct的不同在于struct里面可以储存实例变量的地址。

最后一行示例代码：
```sh
func main{
    http.ListenandServe(":8080",&BaseHandler)

}

type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a => ", req.URL.Query().Get("a"))

	resp.Write([]byte("hello golang"))
}
```

直接调用BaseHandler的方法ServeHTTP


