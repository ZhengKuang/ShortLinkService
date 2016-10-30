#go语言中interface



interface，众所周知，里面应该有method。如果一个struct里面实现了interface的方法，那么就可以说这个struct的实例a就是这个interface的type。

interface还可以是空的，意味着没有方法，这样一来，所有的变量，int a，string b，folat c都持有这个interace。

如果interface的实例变量为ak，那么a可以被赋值成为a，b，c

```sh
type a interface{
 }
//空的intercface a
// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s

```
http://www.jb51.net/article/56812.htm

#go 语言中的map

make ( map [KeyType] ValueType, initialCapacity )
make ( map [KeyType] ValueType )

例子
```sh
func test1() {
    map1 := make(map[string]string, 5)
    map2 := make(map[string]string)
    map3 := map[string]string{}
    map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
    fmt.Println(map1, map2, map3, map4)
}
```

map可以提供原type到目的type的映射。
