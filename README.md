运行main文件可以启动http server 端口号为8080

	httpServer.Route([]string{"POST", "GET"}, "/", handler)
	httpServer.Route([]string{"POST", "GET"}, "/healthz", healthz)
	通过Route方法来注册路由，其中请求/ 路径 handler方法实现了
	1.接收客户端 request，并将 request 中带的 header 写入 response header
    2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
    3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
    
    路由/healthz实现了 healthcheck的功能
    访问 localhost:8080/healthz 时，返回200
    
##利用nsenter来查看容器IP配置
    1.利用命令lsns -t net来查看网络namespace
    
    ubuntu@k8snode:~$ lsns -t net
            NS TYPE NPROCS   PID USER      NETNSID NSFS                      COMMAND
    4026531992 net       6  1003 ubuntu unassigned /run/docker/netns/default /lib/systemd/systemd --user
   
    2.进入namespace运行命令查看ip的命令 nsenter - t <pid> -n ip addr
    
    sudo nsenter -t 1003 -n ip addr
    1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
        link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
        inet 127.0.0.1/8 scope host lo
           valid_lft forever preferred_lft forever
        inet6 ::1/128 scope host
           valid_lft forever preferred_lft forever
    2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
        link/ether 08:00:27:12:d2:e4 brd ff:ff:ff:ff:ff:ff
        inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic enp0s3
           valid_lft 85040sec preferred_lft 85040sec
        inet6 fe80::a00:27ff:fe12:d2e4/64 scope link
           valid_lft forever preferred_lft forever
    3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
        link/ether 08:00:27:03:ea:bb brd ff:ff:ff:ff:ff:ff
        inet 192.168.34.2/24 brd 192.168.34.255 scope global enp0s8
           valid_lft forever preferred_lft forever
        inet6 fe80::a00:27ff:fe03:eabb/64 scope link
           valid_lft forever preferred_lft forever
    4: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
        link/ether 02:42:9c:6d:4a:63 brd ff:ff:ff:ff:ff:ff
        inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
           valid_lft forever preferred_lft forever
        inet6 fe80::42:9cff:fe6d:4a63/64 scope link
           valid_lft forever preferred_lft forever
    12: vethc5eaeaf@if11: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
        link/ether 16:d7:b7:b2:04:d7 brd ff:ff:ff:ff:ff:ff link-netnsid 0
        inet6 fe80::14d7:b7ff:feb2:4d7/64 scope link
           valid_lft forever preferred_lft forever

##将docker build 的容器上传到docker hub上
    1、首先进行docker login命令登录到docker hub中
    2. 给自己通过命令docker build .后的容器打标签
    docker tag 8b<容器id> 007mouse/gp-httpserver:1.0 <自己的容器标签>
    3. 通过docker push 进行上传到docker hub中，连接地址为：
    https://registry.hub.docker.com/r/007mouse/go-httpserver
    