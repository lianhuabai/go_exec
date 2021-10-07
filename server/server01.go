package server

import (
	"fmt"
	"homework/logger"
	"homework/tools"
	"net"
	"net/http"
	"strings"
)

// 获取request ip地址
func getIp(r *http.Request) (string, error) {
	ip_real := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip_real) != nil{
		return ip_real,nil
	}
	xForWardedFor := r.Header.Get("X-Forwarded-For")
	ip_forward := strings.TrimSpace(strings.Split(xForWardedFor, ",")[0])
	if ip_forward != ""{
		return ip_forward,nil
	}
	ip_remote, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		return ip_remote, nil
	}
	return "", nil
}


func Healthz(w http.ResponseWriter, r *http.Request)  {
	ip,_ := getIp(r)
	logger.Infof("客户端请求ip: %s", ip)
	logger.Infof("请求返回响应码 %d", 200)
	fmt.Println(ip)
	// 将环境变量写入response headers
	env,bool := tools.GetEnv("VERSION")
	if bool == false {
		fmt.Println("获取系统环境变量失败")
	} else {
		w.Header().Set("version", env)
	}
	// 获取request headers,写入response heaers
	if len(r.Header) > 0 {
		for k,v := range r.Header{
			w.Header().Set(k,v[0])
		}
	}
	w.WriteHeader(200)
	w.Write([]byte("200"))

}

