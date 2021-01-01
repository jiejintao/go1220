package main

import (
	"fmt"
	"go1220/util"
)

const (
	//Sunday = iota
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
func main()  {
	//var amap map[string]string
	//amap = make(map[string]string)
	//amap["a"] = "b"
	//var a int64
	//a = -2
	//print(a,'\n')
	//fmt.Print(a,'\n')
	//println(a)
	//fmt.Println(a)
	//var slice []int
	//slice = []int{1,2,3,4,5}
	//identifier := []int{1,2,3,4,5}
	//fmt.Println(slice)
	//fmt.Println(identifier)
	//amap := map[string]string{"a":"b","b":"c"}
	//amap["c"] = "d"
	//fmt.Println(amap)
	//fmt.Println(amap["a"])
	//fmt.Printf("%T",amap)
	//print(Tuesday , "\n")
	//print(Saturday)
	//var a int
	//a = 4
	//b := 4
	//c := float64(b)/float64(a)
	//fmt.Println(c)
	a := 0
	b := 0
	test(&a,&b)
	//println(a)
	//println(b)
	type pepole struct {

	}
	//afsdf := pepole{}
	//fmt.Println(afsdf)
	var wx wxpay
	wx.name = "jjt"
	wx.age = 18
	//wx := wxpay{"撒旦法",18}
	//fmt.Println(wx)
	//fmt.Println(wx.age)
	//wx.age = 19
	//fmt.Println(wx.age)
	//wx.call()
	//var ali alipay
	//ali.name= "2314"
	//ali.age= 22
	//
	//var pay pay
	//pay = new(wxpay)
	//pay.call()
	//pay = new(alipay)
	//pay.call()
	fmt.Println("输入你的用户名 : ")
	username := util.CRead()
	fmt.Println("输入用户名为 : ",username)
	fmt.Println("输入a1 : ")
	a1 := util.CRead()
	fmt.Println("您输入的值为:",a1)
}

func test(a,b *int){
	*a = 1
	*b = 2
}

type pay interface {
	call()
}

type wxpay struct {
	name string
	age int
}

type alipay struct {
	name string
	age int
}

func (w wxpay) call()  {
	fmt.Println("This is wxpay!",w.name,w.age)
}

func (a alipay) call()  {
	fmt.Println("This is alipay!",a.name,a.age)
}

func add(a,b int) int {
	return a+b
}