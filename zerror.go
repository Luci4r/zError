package zError

import "math"

type zError interface {
	ErrCode() int64
	ErrStr() string
	New(string) zError
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

func (this *er) New(s string) zError {
	e := new(er)
	if s == "" {
		e.errString = this.errString
	} else {
		e.errString = s
	}
	e.errType = this.errType
	return e
}

var ErrCodeList map[int64]zError

var TypeConflict zError

func SignInzError(t int64, s string) (zError, string) {
	if s, ok := ErrCodeList[t]; ok {
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
	ErrCodeList = make(map[int64]string, 10)
	TypeConflict = SignInzError(math.MaxInt64, "Error Code is Signed aleady")
}
