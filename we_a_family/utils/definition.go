package utils

const Timestamp = "2006-01-02 15:04:05"
const GodUserName = 541513140112
const TagOffset = 1 << 20
const PictureOffset = 10 << 20

type ErrorCode int

const (
	SettingsError      ErrorCode = 1001 //系统错误
	MemberDoesNotExist ErrorCode = 1002 //账号不存在
	RegisterError      ErrorCode = 1003 //注册失败
	ChangeError        ErrorCode = 1004 //更新失败
	DeleteError        ErrorCode = 1005 //删除失败
	RegisterAgainError ErrorCode = 1006 //账号已存在
	DeletedMember      ErrorCode = 1007 //账号为黑用户，请联系管理员
	PwdNotRight        ErrorCode = 1008 //账号或密码不对
	ListError          ErrorCode = 1009 //获取列表失败

)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:      "系统错误",
		MemberDoesNotExist: "账号不存在请注册",
		RegisterError:      "注册失败",
		ChangeError:        "更新失败",
		DeleteError:        "删除失败",
		RegisterAgainError: "账号已存在请登录",
		DeletedMember:      "账号为黑用户，请联系管理员",
		PwdNotRight:        "密码不对",
		ListError:          "获取列表失败",
	}
)

type StatusCode int

// 权限 0-5
const (
	// Black 不可看不可写
	// Reader 可看不可写(tag, picture)
	// Writer 不可看可写(tag, picture)
	// Owner 可看可写(self.tag, self.picture)
	// Admin 可看可写(Reader, Writer, Owner, tag, picture)

	Black  StatusCode = 0
	Reader StatusCode = 1
	Writer StatusCode = 2
	Owner  StatusCode = 3
	Admin  StatusCode = 4
)

const (
	Member     = "member"
	Picture    = "picture"
	Tag        = "tag"
	Permission = "permission"
)
