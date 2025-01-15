package sdk

import (
	"errors"
	"fmt"
)

type Rpc struct {
	host string
	//appId, secret string
	//token         string
	//
	logic func(rest *RestClient) (code int, err error)
}

func NewRpc(host string) *Rpc {
	r := &Rpc{host: host}
	return r
}

func (r *Rpc) BuildReq(handler func(*RestClient) (int, error)) *Rpc {
	r.logic = handler
	return r
}
func (r *Rpc) Exec() (err error) {
	var code int
	if r.logic == nil {
		return
	}
	// 执行logic
	code, err = r.logic(NewRest(r.host))
	if code != CodeSuccess {
		err = errors.New(fmt.Sprintf("%d", code))
		return
	}
	return
}
