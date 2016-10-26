package main 
import(
	"flag"
	"fmt"
)

var (
	a int
)



func main(){
	flag.IntVar(&a,"age",18,"age value")

	score:=flag.Int("score",80,"score value")
	flag.Parse()
	fmt.Println("a=>",a)
	fmt.Println("score=>",*score)
}