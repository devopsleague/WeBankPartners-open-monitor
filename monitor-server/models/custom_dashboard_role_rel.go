package models

type CustomDashBoardRoleRel struct {
	Id              int    `json:"id" xorm:"'id' pk"`
	RoleId          string `json:"roleId" xorm:"role_id"`                   // 角色ID
	CustomDashboard string `json:"customDashboard" xorm:"custom_dashboard"` // 自定义看板
	Permission      string `json:"permission" xorm:"permission"`            // 权限
}