package proxy_pattern

import "testing"

func TestProxyPattern(T *testing.T) {
	action := "Run"
	proxy := ProxyObject{}
	proxy.ObjDo(action)
}
