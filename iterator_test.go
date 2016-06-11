package iterator

import (
	"testing"
)

func check(a, e interface{}, t *testing.T) {
	if a != e {
		t.Errorf("expected %v, actual %v", e, a)
	}
}

// NewBookできるかテストする
// ちゃんと指定した名前の本ができていれば良い
func TestNewBook(t *testing.T) {
	e := "example"
	b := NewBook(e)
	check(b.GetName(), e, t)
}
