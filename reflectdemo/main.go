package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type W struct {
	//a byte //1
	//b int64 // 8
	//c byte // 1
	//d int64 // 4
	a byte //1
	b byte // 8
	c int64 // 1
	d int64 // 4
}

func unsafedemo() {
	var w *W
	fmt.Println(w)
	fmt.Println(unsafe.Sizeof(w))
	fmt.Println(unsafe.Sizeof(*w))
	fmt.Println("--")
	fmt.Println(unsafe.Alignof(w.a))
	fmt.Println(unsafe.Offsetof(w.a))
	fmt.Println(unsafe.Alignof(w.b))
	fmt.Println(unsafe.Offsetof(w.b))
	fmt.Println(unsafe.Alignof(w.c))
	fmt.Println(unsafe.Offsetof(w.c))
	//fmt.Println(unsafe.Alignof(w.d))
	//fmt.Println(unsafe.Offsetof(w.d))

	var w2 W
	fmt.Println(unsafe.Alignof(w))
	fmt.Println(unsafe.Alignof(w2))
}

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func aligndemo() {
	fmt.Println("----------unsafe.Pointer---------")
	t := &T{1, 2, 3, "this is a example", true}
	ptr := unsafe.Pointer(t)
	t1 := (*byte)(ptr)
	fmt.Println(*t1)
	t2 := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t2)))
	*t2 = 99
	fmt.Println(t)
	t3 := (*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t3)))
	fmt.Println(*t3)
	*t3 = 123
	fmt.Println(t)
}


func reflectDemo(x interface{}) {
	v := reflect.TypeOf(x)
	v.Name()
	v.Kind()
	fmt.Printf("type: %v, typeName: %v, typekind: %v\n", v, v.Name(), v.Kind())

}

func reflaceValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v)

	n := v.Int() + 12
	fmt.Println(n)

	kind := v.Kind()
	fmt.Println(kind)
	switch kind {
	case reflect.Int64:
		fmt.Println(v.Int()+10)
	case reflect.Float32:
		fmt.Println(v.Float() + 10.1)
	case reflect.String:
		fmt.Println(v.String())
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(101)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("aaaaa")
	}
}

type myInt int
type Person struct {
	Name string
	Age int
}

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score int `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("name: %v, age:%v, score:%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print(){
	fmt.Println("hwllo")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	//f0 := t.Elem().Field(0)
	//fmt.Printf("%#v\n", f0)
	//fmt.Println(f0.Tag.Get("json"))
	//
	//f1, ok := t.FieldByName("Age")
	//if ok {
	//	fmt.Printf("%#v\n", f1)
	//}

	//fmt.Println(v.FieldByName("Name"))

	setInfo := v.MethodByName("SetInfo")
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo")
	var parmas []reflect.Value
	parmas = append(parmas, reflect.ValueOf("rona"))
	parmas = append(parmas, reflect.ValueOf(23))
	parmas = append(parmas, reflect.ValueOf(99))
	setInfo.Call(parmas)
	fmt.Println(info.Call(nil))

}

func main() {
	//unsafedemo()
	//aligndemo()
	a := 10
	b := 23.4
	c := true
	reflectDemo(a)
	reflectDemo(b)
	reflectDemo(c)
	reflectDemo(myInt(3))
	reflectDemo(Person{})
	reflectDemo(&Person{})

	var i = [3]int{1,2,3}
	var j = []int{11,22,44}
	reflectDemo(i)
	reflectDemo(j)

	var aa = 13
	reflaceValue(aa)

	var aaa int64 = 100
	reflectSetValue2(&aaa)
	fmt.Println(aaa)

	PrintStructField(&Student{"messi", 11, 12})
}