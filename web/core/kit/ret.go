package kit

const (
	RetStateOk   string = "ok"
	RetStateFail string = "fail"
)

type Ret struct {
	State string
	Msg   string
	Param map[string]string
	Data  interface{}
}

// Get common Ret
// 获取通用Ret
func GetCommonRet() Ret {
	ret := Ret{}
	ret.State = RetStateFail
	ret.Param = map[string]string{"Title": "Taostorage"}
	return ret
}
