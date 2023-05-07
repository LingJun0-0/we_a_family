package utils

const Timestemp = "2006-01-02 15:04:05"

type ErrorCode int

const (
	SettingsError      ErrorCode = 1001 //系统错误
	MemberDoesNotExist ErrorCode = 1002 //账号不存在
	RegisterError      ErrorCode = 1003 //注册失败
	ChangeError        ErrorCode = 1004 //更新失败
	RegisterAgainError ErrorCode = 1005 //账号已存在
	DeletedMember      ErrorCode = 1006 //账号为黑用户，请联系管理员
	NameOrPwdNotRight  ErrorCode = 1007 //账号或密码不对
	ListError          ErrorCode = 1008 //获取列表失败

)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:      "系统错误",
		MemberDoesNotExist: "账号不存在请注册",
		RegisterError:      "注册失败",
		ChangeError:        "更新失败",
		RegisterAgainError: "账号已存在请登录",
		DeletedMember:      "账号为黑用户，请联系管理员",
		NameOrPwdNotRight:  "账号或密码不对",
		ListError:          "获取列表失败",
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
