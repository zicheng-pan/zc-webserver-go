package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LogData struct {
	RequestTime string
	IP          string
	STATUS      int
}

type Log struct {
	data []LogData
}

type commonResponse struct {
	BizCode int
	Msg     string
	Data    interface{}
}
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type HandlerMapping struct {
	Mapping map[string]MappingHandler
}

type MappingHandler struct {
	HandleMethod func(context *Context)
	methods      []string
}

func (c *Context) ReadRequetIntoJson(req interface{}) error {
	var body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	fmt.Printf(string(body))
	err = json.Unmarshal(body, req)
	fmt.Printf("aaaabbbbacccc:%v", req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.Writer.WriteHeader(code)
	var data, error = json.Marshal(resp)
	if error != nil {
		return error
	}
	_, error = c.Writer.Write(data)
	return error
}

func (c *Context) OKJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) ErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
	}
}

func IsContain(items []string, item string) bool {

	for _, eachItem := range items {

		if eachItem == item {

			return true
		}
	}
	return false
}
