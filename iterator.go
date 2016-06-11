package iterator

import ()

// 集合体のインターフェース
type Aggregate interface {
	iterator()
}

// 数え上げる人のインターフェース
type Iterator interface {
	hasNext() bool
	next()
}

// 本棚 - Aggregateの実装
type BookShelf struct {
	books []*Book
	last  int
}

func (bs *BookShelf) iterator() *BookShelfIterator {
	return NewBookShelfIterator()
}

// 本棚の中身を数え上げる人 - Iteratorの実装
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator() *BookShelfIterator {
	return &BookShelfIterator{}
}
func (it *BookShelfIterator) hasNext() bool {
	return true
}
func (it *BookShelfIterator) next() *Book {
	return &Book{}
}

// 本
type Book struct {
	name string
}

func NewBook(n string) *Book {
	return &Book{
		name: n,
	}
}
func (b *Book) GetName() string {
	return b.name
}
