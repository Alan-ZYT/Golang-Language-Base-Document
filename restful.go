package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取UID
	uid := ps.ByName("uid")
	fmt.Fprintln(w, "获得到用户名：", uid)
	//获取请求类型
	method := r.Method
	fmt.Fprintln(w, "获得请求方式：", method)

}

func main() {
	rou := httprouter.New()
	rou.GET("/getuser/:uid", getuser)
	http.Handle("/", rou)
	http.ListenAndServe(":8088", nil)
}
