package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unicode/utf8"
)

/*
	Golang数据类型
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
*/

//变量定义，数值型，字符串，布尔
func VarDefineDemo() {
	//1 var + varName + varType
	var VarInt int = 800
	bb := 32.3
	cc := "你好Golang"
	t := reflect.ValueOf(VarInt)
	fmt.Printf("Value=%v,Type=%T,Type=%v \n", VarInt, VarInt, t.Kind())
	fmt.Printf("Value=%v,Type=%T \n", bb, bb)
	fmt.Printf("Value=%v,Type=%T \n", cc, cc)

	// 定义基本数据类型
	p := true                             // bool
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", p, p)
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)

	//字符串长度
	strlen := utf8.RuneCountInString(cc)
	fmt.Printf("cc字符串长度 %v \n", strlen)
}

//数组
func ArrayDemo() {

	//一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	//定义数组，4种方式
	var IntArr1 [3]int

	//初始化
	var IntArr [3]int = [3]int{1, 2, 3}
	var c = [3]string{"a", "b", "c"}
	var d = [...]string{"a", "b", "c"}
	var e = [...]string{2: "a", "b", 6: "c"}
	var f = [...]float64{4: 0}
	var g = [...]string{
		"你",
		"是",
		"谁",
	}
	g = g //自赋值

	//数组的CRUD
	IntArr[0] = 12

	//打印数组长度和容量
	fmt.Printf("Len=%#v,Cap=%v \n", len(IntArr), cap(IntArr))

	//循环遍历
	for i := 0; i < 3; i++ {
		fmt.Println(IntArr[i])
	}

	for k, val := range IntArr {
		fmt.Println(k, val)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("请输入第%v个值 \n", i)
		//v, _ := fmt.Scanln(&f[i])
	}

	fmt.Printf("Value=%#v,Type=%T \n", IntArr1, IntArr1)
	fmt.Printf("Value=%#v,Type=%T \n", IntArr, IntArr)
	fmt.Printf("Value=%#v,Type=%T \n", c, c)
	fmt.Printf("Value=%#v,Type=%T \n", d, d)
	fmt.Printf("Value=%#v,Type=%T \n", e, e)
	fmt.Printf("Value=%#v,Type=%T \n", f, f)
}

//Slice演示
func SliceDemo() {
	var ll []int64
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 1)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 2, 3, 4, 5, 6)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, []int64{7, 8, 9}...)
	fmt.Printf("%#v\n", ll)

	fmt.Println(ll[0:2])
	fmt.Println(ll[:2])
	fmt.Println(ll[0:])
	fmt.Println(ll[:])

	var sli_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)

	var sli_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	fmt.Println("sli[1] ==", sli[1])
	fmt.Println("sli[:] ==", sli[:])
	fmt.Println("sli[1:] ==", sli[1:])
	fmt.Println("sli[:4] ==", sli[:4])

	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4])
}

func MapDemo() {
	//	1
	var m1 map[string]string = map[string]string{"a": "aaa", "b": "bbb"}
	fmt.Printf("%#v %T \n", m1, m1)
	//	2
	var m2 = map[string]string{"a": "aaa", "b": "bbb"}
	fmt.Printf("%#v %T \n", m2, m2)

	//	3
	m3 := map[string]string{"a": "aaa", "b": "bbb"}
	fmt.Printf("%#v %T \n", m3, m3)

	//	4
	var m4 map[string]string = make(map[string]string)
	m4["aa"] = "aaa"
	fmt.Printf("%#v %T \n", m4, m4)

	//	5
	var m5 map[string]string
	m5 = make(map[string]string)
	m5["aa"] = "aaa"
	fmt.Printf("%#v %T \n", m5, m5)

	//	6
	var m6 map[string]string = map[string]string{}
	m6["aa"] = "aaa"
	fmt.Printf("%#v %T \n", m6, m6)

	res := map[string]interface{}{}
	res["code"] = "200"
	res["msg"] = "请求成功"
	res["data"] = map[string]string{
		"list": "列表数据",
		"ret":  "ok",
	}
	fmt.Printf("%#v %T\n", res, res)
	retjosn, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("%#v %T\n", string(retjosn), retjosn)
	res2 := map[string]interface{}{}
	err = json.Unmarshal([]byte(retjosn), &res2)
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	delete(res2, "code")
	fmt.Println("%#v", res2)

}

/*
控制语句
*/
func LoopDemo() {
	fn1 := func(s int) string {
		switch {
		case s > 0:
			fmt.Println("000")
			fallthrough
		case s > 10:
			fmt.Println("1111")
			return "1111"
			fallthrough
		case s > 100:
			fmt.Println("2222")
			return "2222"
		default:
			fmt.Println("3")
			return "3"
		}
		return " ssssssssssssss"
	}

	//fn1(2)
	//fn1(3)
	fn1(20)

	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	//使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}

	var person2 = []string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(person2), cap(person2), person2)

	fmt.Println("")

	//循环
	for k, v := range person2 {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	fmt.Println("")

	for i := range person2 {
		fmt.Printf("person[%d]: %s\n", i, person2[i])
	}

	fmt.Println("")

	for i := 0; i < len(person2); i++ {
		fmt.Printf("person[%d]: %s\n", i, person2[i])
	}

	fmt.Println("")

	//使用空白符
	for _, name := range person2 {
		fmt.Println("name :", name)
	}

	i := 3
	fmt.Printf("当 i = %d 时：\n", i)

	switch i {
	case 1:
		fmt.Println("输出 i =", 1)
	case 2:
		fmt.Println("输出 i =", 2)
	case 3:
		fmt.Println("输出 i =", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i =", "4 or 5 or 6")
	default:
		fmt.Println("输出 i =", "xxx")
	}
}

func StructDemo() {
	type MyData struct {
		One int
		two string
	}
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) // main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // {"One":1}	// 私有字段 two 被忽略了

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // main.MyData{One:1, two:""}
}
func main() {
	//基础类型
	//VarDefineDemo()
	//数组
	//ArrayDemo()
	//Slice
	//SliceDemo()
	//Loop
	//LoopDemo()
	//Struct
	//StructDemo()
	//Map
	MapDemo()
	fmt.Println("程序执行完成！")
}
