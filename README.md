运行main文件可以启动http server 端口号为8080

	httpServer.Route([]string{"POST", "GET"}, "/", handler)
	httpServer.Route([]string{"POST", "GET"}, "/healthz", healthz)
	通过Route方法来注册路由，其中请求/ 路径 handler方法实现了
	1.接收客户端 request，并将 request 中带的 header 写入 response header
    2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
    3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
    
    路由/healthz实现了 healthcheck的功能
    访问 localhost:8080/healthz 时，返回200
    
   