package sdk

type ResponsePage struct {
	Code      int           `json:"code,omitempty"`
	Success   string        `json:"success,omitempty"`
	Message   string        `json:"message,omitempty"`
	MessageId int           `json:"messageId,omitempty"`
	Data      []interface{} `json:"items,omitempty"`
}
type ResponseOne struct {
	Code      int         `json:"code,omitempty"`
	Success   string      `json:"success,omitempty"`
	Message   string      `json:"message,omitempty"`
	MessageId int         `json:"messageId,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
type CommonPageRequest struct {
	PageNo   int `json:"pageNo,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}
type UserRoleRequest struct {
	UserCode string `json:"userCode,omitempty"`
	//RoleCode        string `json:"roleCode,omitempty"`
	ApplicationCode string `json:"applicationCode,omitempty"`
	ApplicationName string `json:"applicationName,omitempty"`
}
