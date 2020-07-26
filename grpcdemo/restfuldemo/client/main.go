package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	url := "http://localhost:8080/v1/example/echo"
	jsonStr := []byte(`{"value":" world"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
