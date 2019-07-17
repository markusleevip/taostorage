package kit

const (
	RetStateOk   int8 = 1
	RetStateFail int8 = -1
)

type Ret struct {
	State int8
	Msg   string
	Data  interface{}
}

// Get common Ret
// 获取通用Ret
func GetCommonRet() Ret {
	ret := Ret{}
	ret.State = RetStateFail
	return ret
}
