package main


import (
	"fmt"
	"net/http"
	"flag"
	"log"


	"short_url/lib/mux"


)

var(
	port string
)


func main(){
	flag.StringVar(&port,"port",":8080","port to listen")
	flag.Parse()

	router:=mux.NewMuxHandler()

	router.Handle("/hello/golang/",&BaseHandler{})
	router.HandleFunc("/hello/world",func(resp  http.ResponseWriter, req *http.Request){
		resp.Write([]byte("hello world"))
	})
	log.Println("ShortURL server will start at prot "+port)
	log.Fatalln(http.ListenAndServe(port,router))
}


type BaseHandler struct{

}

func (handler *BaseHandler) ServeHTTP(resp  http.ResponseWriter, req *http.Request){
	fmt.Println("url path=>",req.URL.Path)
	fmt.Println("url param a=>",req.URL.Query().Get("a"))
	resp.Write([]byte("hello golang"))
}