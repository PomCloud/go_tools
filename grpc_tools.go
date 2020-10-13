package utils

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

//获取请求客户端的远程地址
//
//通过从metadata中获取远程地址信息
func GetClientIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}

	addSlice := strings.Split(pr.Addr.String(), ":")
	if addSlice[0] == "[" {
		//本机地址
		return "localhost", nil
	}
	return addSlice[0], nil
}
