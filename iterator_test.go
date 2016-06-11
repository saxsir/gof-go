package iterator

import (
	"testing"
)

func check(a, e interface{}, t *testing.T) {
	if a != e {
		t.Errorf("expected %v, actual %v", e, a)
	}
}

func TestIteratorPattern(t *testing.T) {
	bs := NewBookShelf()
	asserts := []string{"A", "B", "C"}
	for _, assert := range asserts {
		bs.Add(NewBook(assert))
	}
	i := 0
	for iterator := bs.Iterator(); iterator.HasNext(); {
		if book := iterator.Next(); book.(*Book).name != asserts[i] {
			t.Errorf("Expect Book.name to %s, but %s", asserts[i], book.(*Book).name)
		}
		i++
	}
}
