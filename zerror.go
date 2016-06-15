package zError

import "math"

type ZError interface {
	ErrCode() int64
	ErrStr() string
	New(string) ZError
}

type er struct {
	errType   int64
	errString string
}

func (this *er) ErrCode() int64 {
	return this.errType
}

func (this *er) ErrStr() string {
	return this.errString
}

func (this *er) New(s string) ZError {
	e := new(er)
	if s == "" {
		e.errString = this.errString
	} else {
		e.errString = s
	}
	e.errType = this.errType
	return e
}

var ErrCodeList map[int64]ZError

var TypeConflict ZError

func SignInzError(t int64, s string) (ZError, string) {
	if _, ok := ErrCodeList[t]; ok {
		return nil, TypeConflict.ErrStr()
	} else {
		e := new(er)
		e.errType = t
		e.errString = s
		ErrCodeList[t] = e
		return e, ""
	}
}

func init() {
	ErrCodeList = make(map[int64]ZError, 10)
	TypeConflict, _ = SignInzError(math.MaxInt64, "Error Code is Signed aleady")
}
