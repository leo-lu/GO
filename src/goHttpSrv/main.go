package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"

	//"data"
	//flatbuffers "github.com/google/flatbuffers/go"
	//data "goHttpSrv/data"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type friend struct {
	Name string
	Age  int
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	rst, _ := ioutil.ReadAll(r.Body)

	var f friend
	json.Unmarshal(rst, &f)
	fmt.Println("收到客户端请求：", r.Method, f)
	//v := url.Values{"name": {"xiaoma"}, "email": {"xiaoma.xx.com"}}
	v := user{Name: "xiaohua", Email: "xiao.hua.com"}
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("err: ", err)
	}
	w.Write([]byte(b))

	//b := flatbuffers.NewBuilder(0)
	//
	//data.MonsterStart(b)
	//data.MonsterAddPos(b, data.CreateVec2(b, 10, 20))
	//data.MonsterAddHp(b, 80)
	//data.MonsterAddName(b, b.CreateString("myM"))
	//mon := data.MonsterEnd(b)
	//fmt.Println(mon)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
	fmt.Println("end")
}
