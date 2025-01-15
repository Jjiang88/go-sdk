package sdk

import "testing"

func TestSdk(t *testing.T) {
	//a := New(WithHost("http://127.0.0.1:80"))
	//role, err := a.GetUserRole(UserRoleRequestPage{
	//	UserCode:        "",
	//	RoleCode:        "",
	//	ApplicationCode: "",
	//	PageNo:          1,
	//	PageSize:        10,
	//})
	//t.Log(role, err)
	//if err != nil {
	//	return
	//}
}
func TestGetOne(t *testing.T) {
	a := New(WithHost("http://role.thingple.io"))
	role, err := a.GetUserRoleList(UserRoleRequest{
		UserCode:        "87e9ca438fb54e6eb6632d42dbf8915c",
		ApplicationCode: "0e2bbea41a7a4af6b0b666b7f19e33fa",
	})
	t.Log(role, err)
	if err != nil {
		return
	}
}
