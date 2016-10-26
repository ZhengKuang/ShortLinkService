package main


import (
	"fmt"
	"net/http"
	"flag"
	"log"


)

var(
	port string
)


func main(){
	flag.StringVar(&port,"port",":8080","port to listen")
	flag.Parse()
	log.Println("ShortURL server will start at prot "+port)
	log.Fatalln(http.ListenAndServe(port,&BaseHandler{}))
}


type BaseHandler struct{

}

func (handler *BaseHandler) ServeHTTP(resp  http.ResponseWriter, req *http.Request){
	fmt.Println("url path=>",req.URL.Path)
	fmt.Println("url param a=>",req.URL.Query().Get("a"))
	resp.Write([]byte("hello golang"))
}