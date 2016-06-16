package zError

import (
	"math"
)

//
//type ZError interface {
//	ErrCode() int64
//	ErrStr() string
//	New(string) ZError
//}

type ZError struct {
	errType   int64
	errString string
}

func (this *ZError) ErrCode() int64 {
	return this.errType
}

func (this *ZError) Error() string {
	return this.errString
}

func (this *ZError) New(s string) *ZError {
	if s == "" {
		return this
	} else {
		return &ZError{errType: this.errType, errString: s}
	}
}

var ErrCodeList map[int64]*ZError

var TypeConflict *ZError

func SignInzError(t int64, s string) (*ZError, string) {
	if _, ok := ErrCodeList[t]; ok {
		return nil, TypeConflict.Error()
	} else {
		e := ZError{errType: t, errString: s}
		ErrCodeList[t] = &e
		return &e, ""
	}
}

func init() {
	ErrCodeList = make(map[int64]*ZError, 10)
	TypeConflict, _ = SignInzError(math.MaxInt64, "Error Code is Signed aleady")
}
