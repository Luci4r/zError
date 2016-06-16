package zError

import (
	"errors"
	"testing"
)

var ErrTest1 ZError
var ErrTest2 ZError
var ErrTest3 ZError

type errType struct {
	errName *ZError
	errStr  string
	errType int64
}

var ErrList1 = []errType{
	errType{
		errName: &ErrTest1,
		errStr:  "Error test1",
		errType: -1,
	},
	errType{
		errName: &ErrTest2,
		errStr:  "Error test2",
		errType: -2,
	},
	errType{
		errName: &ErrTest3,
		errStr:  "Error test3",
		errType: 1024,
	},
}

func TestSignInzError(t *testing.T) {
	for _, v := range ErrList1 {
		e, es := SignInzError(v.errType, v.errStr)
		if es != "" {
			t.Error("Error sign in error ", es)
			return
		} else {
			*(v.errName) = *e
		}
	}
	if ErrTest1.ErrCode() != -1 || ErrTest1.Error() != "Error test1" {
		t.Error("Error sign in error ")
		return
	}
	if ErrTest2.ErrCode() != -2 || ErrTest2.Error() != "Error test2" {
		t.Error("Error sign in error ")
		return
	}
	if ErrTest3.ErrCode() != 1024 || ErrTest3.Error() != "Error test3" {
		t.Error("Error sign in error ")
		return
	}
	net1 := ErrTest1.New("New Error test1")
	if net1.ErrCode() != ErrTest1.ErrCode() || net1.Error() != "New Error test1" {
		t.Error("Error sign in error ")
		return
	}
}

func TestDuplicateSignInzError(t *testing.T) {
	_, es := SignInzError(4, "4")
	if es != "" {
		t.Error("Error sign in error ", es)
		return
	}
	_, es = SignInzError(4, "4 duplicate")
	if es == "" {
		t.Error("Error check duplicate sign in")
		return
	}
}

var ErrStr = "Error test"

func BenchmarkEr_ZErrorNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrTest1.New(ErrStr)
	}
}

func BenchmarkEr_ErrorNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errors.New(ErrStr)
	}
}
