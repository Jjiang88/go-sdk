package sdk

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/pkg/errors"
)

const baseUrl = "/api"
const (
	userRoleUrl = "/userRoleList"
	rolesUrl    = "/roles"
	ApiUserRole = "/userRole"
)

const (
	CodeNotAllow int = 401
	CodeNotFound int = 404
	CodeSuccess  int = 200
)

// messageClient 消息服务
type messageClient struct {
	host string // 消息服务主机地址. 如, "http://127.0.0.1"
	//credential string
	//appId      string
	//secret     string
}

func (m messageClient) GetRoles(request UserRoleRequest) (resp ResponseOne, err error) {
	if debugEnable {
		log.Debug("GetRoles: %s", request.UserCode)
	}
	req := make(map[string]string)
	if len(request.ApplicationName) > 0 {
		req["applicationName"] = fmt.Sprintf("%s", request.ApplicationName)
	}
	if len(request.UserCode) > 0 {
		req["userCode"] = fmt.Sprintf("%s", request.UserCode)
	}
	err = NewRpc(m.host).BuildReq(func(rest *RestClient) (int, error) {
		code, e := rest.Path(baseUrl, ApiUserRole, rolesUrl).
			ResponseJson(&resp).Get(req)
		return code, e
	}).Exec()

	if err != nil {
		err = errors.Wrap(err, "get roles fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	return
}

func (m messageClient) GetUserRoleList(request UserRoleRequest) (resp ResponseOne, err error) {
	if debugEnable {
		log.Debug("GetUserRoleList: %s", request.UserCode)
	}
	req := make(map[string]string)
	if len(request.ApplicationCode) > 0 {
		req["applicationCode"] = fmt.Sprintf("%s", request.ApplicationCode)
	}
	if len(request.UserCode) > 0 {
		req["userCode"] = fmt.Sprintf("%s", request.UserCode)
	}
	err = NewRpc(m.host).BuildReq(func(rest *RestClient) (int, error) {
		code, e := rest.Path(baseUrl, ApiUserRole, userRoleUrl).
			ResponseJson(&resp).Get(req)
		return code, e
	}).Exec()

	if err != nil {
		err = errors.Wrap(err, "get userRole fail")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	return
}
