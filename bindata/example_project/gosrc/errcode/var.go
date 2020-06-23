package errcode

var (
	OK                       = add(0, "ok")
	BadRequest               = add(1, "参数错误")
	Permission               = add(2, "没有操作权限")
	InvalidCaptcha           = add(3, "验证码错误")
	Empty                    = add(4, "不能为空")
	Duplicate                = add(5, "重复")
	NotFound                 = add(20, "数据不存在")
	AlreadyExist             = add(21, "数据已经存在")
	IncorrectUserOrPwd       = add(400, "账号或密码不正确") //400-499  登录相关,前端可控制跳转重新登录
	InvalidToken             = add(403, "登录信息失效,请重新登录")
	SQLAffectedEq0           = add(1012, "执行操作未失效")
	AdminLocked              = add(1013, "管理锁定,无法操作")
	InvalidResource          = add(1015, "不正确的资源")
	TooMany                  = add(1020, "数量限制")
	TemporaryNoDataAvailable = add(1021, "暂无可用数据")
	RateLimit                = add(1111, "频率限制")
	Error                    = add(9999, "错误")
)
