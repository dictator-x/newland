package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./fsdemo/demo.txt")
	defer file.Close()
	if err != nil {
		panic("can not open file")
	}

	//r := make([]byte, 128)
	//var all []byte
	//for {
	//	n, err := file.Read(r)
	//	if err == io.EOF {
	//		break
	//	}
	//	all = append(all, r[:n]...)
	//}
	//fmt.Println(string(r))

	//r, _ := ioutil.ReadAll(file)
	//fmt.Println(string(r))

	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, e := reader.ReadString('\n')
		if e == io.EOF {
			fileStr += str
			break
		}
		fileStr += str
	}
	fmt.Println(fileStr)

	write, ew := os.OpenFile("./fsdemo/write.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer write.Close()
	if ew != nil {
		panic("can not open file")
	}
	//write.Write([]byte("lalala\nllala"))
	//write.Sync()

	out := bufio.NewWriter(write)
	out.WriteString("hello")
	out.Flush()
	os.Rename("./fsdemo/write.txt","./fsdemo/write1.txt")


}