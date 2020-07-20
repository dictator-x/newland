package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dictator-x/newland/pb"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func basicDemo() {
	fmt.Println("Hello")
	fmt.Println("Hello World")
	var a int = 10
	var b int = 3
	fmt.Println(a, b)
	fmt.Printf("a=%v, a=%T\n", a, a)

	var num int8 = 97
	fmt.Println(num)

	var c int8 = 15
	fmt.Printf("num=%v type:%T\n", c, c)
	fmt.Println(unsafe.Sizeof(c))

	var a1 int32 = 10
	var a2 int64 = 32
	fmt.Println(int64(a1) + a2)
	fmt.Println(a1 + int32(a2))

	var fa float32 = 3.12
	fmt.Printf("value: %v--%f, type: %T\n", fa, fa, fa)
	fmt.Println(unsafe.Sizeof(fa))

	var faa float64 = 3.12
	fmt.Printf("value: %v--%.1f, type: %T\n", faa, faa, faa)
	fmt.Println(unsafe.Sizeof(faa))

	var num1 float64 = 3.1
	var num2 float64 = 4.2

	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Printf("%v--%T\n", d1, d1)

	aa := 10
	bb := float64(aa)
	fmt.Println(aa, bb)

	var flag bool = true
	fmt.Printf("%v--%T\n", flag, flag)

	var s1 string = "hello golang"
	fmt.Println(s1)

	var s2 string = `hello
  world
  `
	fmt.Println(s2 + s1)
	var s3 = "123-345-546"
	arr := strings.Split(s3, "-")
	fmt.Println(arr)
	fmt.Println(strings.Join(arr, "*"))

	var ca = 'a'
	fmt.Printf("%v--%T\n", ca, ca)

	var str = "this"
	fmt.Printf("%v--%T\n", str[2], str[2])
	fmt.Println(unsafe.Sizeof(str))
	var str1 = "你好"
	fmt.Println(len(str1))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	for _, r := range str1 {
		fmt.Printf("%v=%c", r, r)
	}
	fmt.Println()

	ss1 := "你好big"
	for _, r := range ss1 {
		fmt.Printf("%v=%c--%T ", r, r, r)
	}
	fmt.Println()

	ss2 := "nihao和"
	runeStr := []rune(ss2)
	runeStr[5] = '号'
	var ccc int32 = '豪'
	fmt.Println(string(runeStr))
	fmt.Println(ccc)
	fmt.Println(333 + '你')

	var vi int = 20
	var vf float64 = 12.456
	var vt bool = true
	var vb byte = 'a'

	vs1 := fmt.Sprintf("%d", vi)
	fmt.Printf("%v--%T\n", vs1, vs1)

	vs2 := fmt.Sprintf("%f", vf)
	fmt.Printf("%v--%T\n", vs2, vs2)

	vs3 := fmt.Sprintf("%t", vt)
	fmt.Printf("%v--%T\n", vs3, vs3)

	vs4 := fmt.Sprintf("%c", vb)
	fmt.Printf("%v--%T\n", vs4, vs4)

	var ci int = 20
	st1 := strconv.FormatInt(int64(ci), 10)
	fmt.Println(st1)

	ssa := "123456"
	fmt.Printf("%v-%T\n", ssa, ssa)
	fmt.Println(strconv.ParseInt(ssa, 10, 64))
	fmt.Println(5.0 / 2)

	bbb := true || false
	fmt.Println(bbb)

	if age := 34; age > 20 {
		fmt.Println("old")
	}

	if num := 9; num > 0 {
		fmt.Println(num, "is negative")
	} else if b := 10; num > b {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	arr3 := [...]string{"aaa", "bbb", "ccc"}
	fmt.Println(arr3)
	fmt.Println(len(arr3))

	var str5 = "你好"
	for k, v := range str5 {
		fmt.Println(k, v)
	}

	switch extname := ".html"; extname {
	case ".html":
		fmt.Println("text/html")
	case ".css":
		fmt.Println("text/css")
	case ".js":
		fmt.Println("text/javascript")
	default:
		fmt.Println("unknow")
	}

	switch n := 5; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("even")
	}

	switch age := 30; {
	case age < 20:
		fmt.Println("young")
	default:
		fmt.Println("old")
	}

label:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			break label
			fmt.Println("aaa")
		}
	}

	var nn = 30
	if nn > 24 {
		fmt.Println("adult")
		goto label2
	}
	fmt.Println("adult")
	fmt.Println("adult")
label2:
	fmt.Println("adult")
	fmt.Println("adult")

	type people struct {
		name string
		age  int
	}

	ppp := &people{"aaa", 123}
	fmt.Println(ppp)
	fmt.Println(ppp.name)

	var arrr = [3][2]string{
		{"aa", "vv"},
		{"aa", "vv"},
		{"aa", "vv"},
	}
	arrr1 := arrr
	arrr1[0][1] = "aaaaaaaaaa"

	fmt.Println(arrr1)
	fmt.Println(arrr)

	arr5 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr5))

	slice1 := arr5[:]
	slice1[3] = 100
	fmt.Println(slice1)
	fmt.Println(arr5)

	slice2 := arr5[:3]
	fmt.Println(slice2)
	fmt.Println(arr5)

	sliceA := make([]int, 4, 8)
	fmt.Println(sliceA)

	var sliceB []int
	fmt.Println(sliceB)

	sliceB = append(sliceB, 12)
	fmt.Println(sliceB)

	appendA := []string{"php", "java"}
	appendB := []string{"nodejs", "go"}

	appendC := append(appendA, appendB...)

	fmt.Println(appendC)

	sliceA = []int{1, 2, 3, 4, 5}
	sliceB = make([]int, 4, 4)
	copy(sliceB, sliceA)
	fmt.Println(sliceB, sliceA)

	var map1 = make(map[string]string, 10)
	map1["username"] = "wrong"
	map1["password"] = "adminAB1"
	fmt.Println(map1)

	delete(map1, "username")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	var userinfo = make([]map[string]string, 3, 3)
	fmt.Println(userinfo[0] == nil)
	userinfo[0] = make(map[string]string)

	type calc func(int, int) int
	var cal calc = add
	fmt.Println(cal(1, 2))
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func fn1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("show: ", err)
		}
	}()
	panic("panic")
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("file fails")
	}
}

type Month int

var months = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// TODO: method
func (m Month) String() string {
	if January <= m && m <= December {
		return months[m-1]
	}
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(m))
	fmt.Println("afdsac")
	return "%!Month(" + string(buf[n:]) + ")"
}

func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

func timeDemo() {
	t := time.Now()
	fmt.Println(t.Year(), int(t.Month()), t.Day())
	fmt.Println(Month(112))
	fmt.Println(t.Format("2006/01/02 15:04:05"))
	fmt.Println(t.Unix())
	fmt.Println(t.UnixNano())
	unixTime := 1594605794
	fmt.Println(time.Unix(int64(unixTime), 0))

	ticker := time.NewTicker(time.Second)
	end := 5
	for t := range ticker.C {
		fmt.Println(t)
		end--
		if end < 0 {
			ticker.Stop()
			break
		}
	}
}

func demoProto() {
	p := &pb.Person{
		Name: "Jack",
		Age:  10,
		From: "China",
	}
	fmt.Println("原始数据:", p)
	dataMarshal, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("proto.Unmarshal.Err: ", err)
		return
	}
	fmt.Println("编码数据:", dataMarshal)
	// 反序列化
	entity := pb.Person{}
	err = proto.Unmarshal(dataMarshal, &entity)
	if err != nil {
		fmt.Println("proto.Unmarshal.Err: ", err)
		return
	}

	fmt.Printf("解码数据: 姓名：%s 年龄：%d 国籍：%s ", entity.GetName(), entity.GetAge(), entity.GetFrom())

}

type student struct {
	name string
}

func (s student) sayHello() {
	fmt.Printf("%p\n", &s)
	fmt.Println(s.name)
}

type User struct {
	Username string
	Password string
	Address1
	Address2
}

type Address1 struct {
	room int
}

type Address2 struct {
	room int
}

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	//basicDemo()
	//fn1()
	//imeDemo()
	// demoProto()

	// s := student{
	//	name: "xxxx",
	// }
	// fmt.Printf("%p\n", &s)
	// s.sayHello()
	// (&s).sayHello()

	// user := User{"aaa", "bbb", Address1{123}, Address2{321}}
	// user.Username = "a"
	// fmt.Printf("%#v\n", user)
	// fmt.Println(user.Address1.room)

	var s1 = Student{
		ID:     123,
		Gender: "male",
		Name:   "mansi",
		Sno:    "s0001",
	}

	jsonByte, _ := json.Marshal(s1)
	fmt.Println(string(jsonByte))
	s2 := new(Student)
	json.Unmarshal(jsonByte, s2)
	fmt.Println(s2)
}
