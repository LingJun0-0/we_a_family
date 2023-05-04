package utils

const Timestemp = "2006-01-02 15:04:05"

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
	LoginError    ErrorCode = 1002 //账号或密码错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		LoginError:    "账号或密码错误",
	}
)

type StatusCode int

//用户权限 0-7，7为最高
const (
	MemberStatusCode0 StatusCode = 0 //0为黑用户权限
	MemberStatusCode1 StatusCode = 1 //1为默认，仅可查看标签(及标签下的照片)
	MemberStatusCode2 StatusCode = 2 //2为可修改标签的人(不能增和删)
	MemberStatusCode3 StatusCode = 3 //3为标签拥有者（对标签及内容的增删改查）
	MemberStatusCode4 StatusCode = 4 //4为系统管理员（对用户，标签，照片的增删改查）
	MemberStatusCode5 StatusCode = 5 //5为系统拥有着（系统的所有权限）
)
