package util

type Dao[RT any] interface {
	FindOne(id string) *RT
}

type Getter[T any] interface {
	Get() T
}

type MyDao struct {
}

type ReturnType struct {
	id int
}

func (m *MyDao) FindOne(id string) *ReturnType {
	panic("implement me")
}

type MyStruct[T any] struct {
	Val T
}

// -----------------com ponteiros---------
// / parametrized pointer receiver
func (m *MyStruct[T]) Get() T {
	return m.Val
}

func foo() Getter[string] {
	return &MyStruct[string]{} // ok
}

//----------------------

func NewMyDao() Dao[ReturnType] {
	return &MyDao{}
}
