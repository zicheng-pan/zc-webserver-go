package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var logdata = new(Log)

func handler(context *Context) {
	/*
		var body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "read the data error: %v \n", err)
			return
		}
		fmt.Fprintf(w, "read body data: %s \n", string(body))
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "read the data error: %v \n", err)
			return
		}

		fmt.Fprintf(w, "read body data: %s \n", string(body))
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	*/

	whead := context.Writer.Header()
	for key, value := range context.Request.Header {
		whead.Add(key, fmt.Sprintf("%v", value))
	}
	whead.Add("OSVERSION", fmt.Sprintf("%v", os.Getenv("VERSION")))
	itemLogData := LogData{
		RequestTime: time.ANSIC,
		IP:          context.Request.Host,
		STATUS:      http.StatusOK,
	}
	logdata.data = append(logdata.data, itemLogData)

	fmt.Println("请求时间:%s, 请求地址:%s, 返回状态码:%d\n", itemLogData.RequestTime, itemLogData.IP, itemLogData.STATUS)
	fmt.Fprintf(context.Writer, "请查看请求头!!\n")
	fmt.Fprintf(context.Writer, "请求日志如下:\n")
	for _, data := range logdata.data {
		fmt.Fprintf(context.Writer, "请求时间:%s, 请求地址:%s, 返回状态码:%d\n", data.RequestTime, data.IP, data.STATUS)
	}

}

type UserSignUPReq struct {
	Username string
	Password string
	Email    string
}

func signUP(context *Context) {
	var req UserSignUPReq

	var err = context.ReadRequetIntoJson(&req)
	if err != nil {
		fmt.Fprintf(context.Writer, "regiester error : %v", err)
		context.ErrorJson(err)
		return
	}
	resp := &commonResponse{
		BizCode: 200,
		Data:    req,
	}
	err = context.OKJson(resp)
	if err != nil {
		fmt.Printf("write http response error, err:%v", err)
	}
}

func healthz(ctx *Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
}

func main() {
	var httpServer = ZCHttpServer{
		ServerName: "go http server",
		handlerMapping: &HandlerMapping{
			Mapping: make(map[string]MappingHandler),
		},
	}
	httpServer.Route([]string{"POST", "GET"}, "/", handler)
	httpServer.Route([]string{"POST", "GET"}, "/healthz", healthz)
	httpServer.Route([]string{"POST"}, "/signUP", signUP)
	httpServer.Start(":8080")

}
