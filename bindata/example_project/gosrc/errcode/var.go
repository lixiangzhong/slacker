package errcode

var (
	OK                 = add(0, "ok")
	BadRequest         = add(1, "参数错误")
	Permission         = add(2, "没有操作权限")
	InvalidCaptcha     = add(3, "验证码错误")
	NotFound           = add(20, "数据不存在")
	AlreadyExist       = add(21, "数据已经存在")
	IncorrectUserOrPwd = add(400, "账号或密码不正确") //401-499  登录相关,前端可控制跳转重新登录
	InvalidToken       = add(403, "登录信息失效,请重新登录")
	Error              = add(9999, "错误")
)
