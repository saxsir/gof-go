package iterator

import ()

// 集合体のインターフェース
type Aggregate interface {
	Iterator() Iterator
}

// 数え上げる人のインターフェース
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 本棚 - Aggregateの実装
type BookShelf struct {
	books []*Book
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: []*Book{},
	}
}
func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}
func (bs *BookShelf) GetBookAt(i int) *Book {
	return bs.books[i]
}
func (bs *BookShelf) Add(b *Book) {
	bs.books = append(bs.books, b)
}

// 本棚の中身を数え上げる人 - Iteratorの実装
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bs *BookShelf) Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
		index:     0,
	}
}
func (it *BookShelfIterator) HasNext() bool {
	if it.index < len(it.bookShelf.books) {
		return true
	}
	return false
}
func (it *BookShelfIterator) Next() interface{} {
	b := it.bookShelf.GetBookAt(it.index)
	it.index++
	return b
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
