package main

import (
	"fmt"
	"net/http"
)

type ZCHttpServer struct {
	handlerMapping *HandlerMapping
	ServerName     string
}

type HttpServerAPI interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

func (s *ZCHttpServer) Route(methods []string, pattern string, handlerFunc func(context *Context)) {

	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	var context = NewContext(writer, request)
	//	handlerFunc(context)
	//})

	s.handlerMapping.Mapping[pattern] = MappingHandler{
		HandleMethod: handlerFunc,
		methods:      methods,
	}

	fmt.Printf("add pattern : %v", pattern)
	fmt.Printf("add mapping : %v", s.handlerMapping.Mapping)

}

func (s *ZCHttpServer) handlerFunc(context *Context) {

	request := context.Request
	method := request.Method
	path := request.URL.Path

	handler, ok := s.handlerMapping.Mapping[path]
	if !ok {
		context.WriteJson(http.StatusNotFound, "request path not found")
		return
	}

	if IsContain(handler.methods, method) {
		s.handlerMapping.Mapping[path].HandleMethod(context)
	} else {
		context.WriteJson(http.StatusForbidden, "request operation is not correct")
	}
}

func (s *ZCHttpServer) LoadServerRequest() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var context = NewContext(writer, request)
		s.handlerFunc(context)
	})
}

func (s *ZCHttpServer) Start(address string) error {
	s.LoadServerRequest()
	return http.ListenAndServe(address, nil)
}
