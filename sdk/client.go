package sdk

import (
	"errors"
)

type Client interface {
	GetUserRoleList(request UserRoleRequest) (resp ResponseOne, err error) //获取单个用户权限列表
	GetRoles(request UserRoleRequest) (resp ResponseOne, err error)        //获取单个用户权限列表
}

type Option func(*messageClient)

var (
	InnerServerErr = errors.New("500")
	NotFoundErr    = errors.New("404")
)

var (
	debugEnable = false
)

// Debug 输出sdk中的log
func Debug(enable bool) {
	debugEnable = enable
	DebugEnable = enable
}

func WithHost(host string) Option {
	return func(client *messageClient) {
		client.host = host
	}
}

// New 初始化.
//
// 内部存储了Credentials,应该确保复用,而不是每次新建
//
// Client 内置了刷新credentials功能,不需要考虑credentials的获取问题.
func New(options ...Option) (m Client) {
	c := &messageClient{}
	for _, opt := range options {
		opt(c)
	}
	m = c
	return
}
