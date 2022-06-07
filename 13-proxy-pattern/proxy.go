package proxy_pattern

import "fmt"

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can, %s", action)
}

type ProxyObject struct {
	object *Object
}

func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "Run" {
		p.object.ObjDo(action)
	}
}
