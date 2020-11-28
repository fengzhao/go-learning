# builtin简介


## make函数

用于给 map , slice , chan(only)  这些类型的变量申请内存空间。
```go
// $gopath/src/builtin/builtin.go
func make(t Type, size ...IntegerType) Type
```
函数入参是（类型，长度...）
返回的其类型的变量

```go

// make声明切片
// make(元素类型,切片长度,切片容量)
mySlice1 := make([]int, 5, 10)
// 当长度和容量相同时，可以省略容量，make(元素类型,切片长度)
//mySlice1 := make([]int, 5)
fmt.Println(mySlice1)
// [0 0 0 0 0]

// 切片变量的数据类型
fmt.Println("the type of mySlice1 is :", reflect.TypeOf(mySlice1))
// the type of mySlice1 is : []int


// make短声明一个空map
stringMap := make(map[string][]string, 100)
// 打印空map
fmt.Println(stringMap)
// map[]

// 添加键值对，值是切片字面量
stringMap["Red"] = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
var slice3 = []string{"1", "2", "3", "4"}
stringMap["slice3"] = slice3

fmt.Println(stringMap)
// map[Red:[Red Blue Green Yellow Pink] slice3:[1 2 3 4]]

```


 








## new函数
new函数用于申请一块内存地址（即申请一个指针）
```go
// $gopath/src/builtin/builtin.go
func new(Type) *Type
```
函数入参是类型
函数返回的是一个指针（即内存地址，存放的值是这个类型的零值）

表达式` new(T)` 创建一个未命名的 T 类型变量，初始化为 T 类型的零值，并返回其内存地址。
```go
// 声明一个指针
var p = new(int)
// p的值：0xc000072090 （指针是内存地址）
fmt.Println(p)
// 0xc00000a0d8

// p的数据类型：p是指针
fmt.Println("the type of p is :", reflect.TypeOf(p))
// the type of p is : *int

// 输出这个指针指向的变量的值：其类型的零值
fmt.Print(*p)
// 0
```