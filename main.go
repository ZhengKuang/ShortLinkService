package main


import (
	"fmt"
	"net/http"

)

func main(){
	http.ListenAndServe(":8080",&BaseHandler{})
}


type BaseHandler struct{

}

func (handler *BaseHandler) ServeHTTP(resp  http.ResponseWriter, req *http.Request){
	fmt.Println("url path=>",req.URL.Path)
	fmt.Println("url param a=>",req.URL.Query().Get("a"))
	resp.Write([]byte("hello golang"))
}