package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	Name  string
	Email string
}

type friend struct {
	Name string
	Age  int
}

func main() {
	f := friend{Name: "ffff", Age: 10}
	b, err := json.Marshal(f)

	resp, err := http.Post("http://127.0.0.1:9090", "application/json;charset=utf-8", bytes.NewBuffer([]byte(b)))
	if err != nil {
		fmt.Println("err: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	var data user
	err = json.Unmarshal(body, &data)
	if nil != err {
		fmt.Println("err:", err)
	}
	fmt.Println(string(body), data.Name, data.Email)
}
