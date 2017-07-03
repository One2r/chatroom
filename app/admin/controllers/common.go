package controllers

//AjaxErrReturn ajax error 返回结构
type AjaxErrReturn struct {
	Error bizExcep
}

//AjaxSuccReturn ajax success 返回结构
type AjaxSuccReturn struct {
	Data interface{}
}

type bizExcep struct {
	Code int
	Msg  string
}

// BizException 业务异常输出结构
func BizException(msg string, code int) AjaxErrReturn {
	return AjaxErrReturn{Error: bizExcep{Code: code, Msg: msg}}
}
