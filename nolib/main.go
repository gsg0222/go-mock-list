package main

import (
	"fmt"

	"github.com/gsg0222/go-mock-list/interfaces"
)

func main(){
	var f interfaces.Foo
	// MockFoo1はBazzを実装していないがFooとして利用可能
	f = &MockFoo1{}
	result1, err := f.Bar(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result1)
	// Bazzは実装していないのでpanicになる
	// result2, err := f.Bazz("string")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result2)

	// MockFoo2もBarを実装していないがFooとして利用可能
	// Bazzメソッドで実際に呼び出されるメソッドをここで定義している
	f = &MockFoo2{
		FakeBaz: func(s string)(string, error){
			return s + s, nil
		},
	}

	// Barは実装していないのでpanicになる
	// result3, err := f.Bar(1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(result3)
	result4, err := f.Baz("string")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result4)
}

// FooをembeddedしているのでFooとして利用可能
// ただし実装していないメソッドを実行するとpanicになる
type MockFoo1 struct{
	interfaces.Foo
}

func (f *MockFoo1) Bar(i int)(int, error){
	if i < 0 {
		return 0, fmt.Errorf("error : should not minus but %d", i)
	}
	return i*2, nil
}

// FooをembeddedしているのでFooとして利用可能
// ただし実装していないメソッドを実行するとpanicになる
// 実際にテストで使いやすいようにリテラルでメソッドを実装できるようにしてある
type MockFoo2 struct{
	interfaces.Foo
	FakeBar func (int) (int, error)
	FakeBaz func (string) (string, error)
}

func (f *MockFoo2) Baz(s string) (string, error){
	return f.FakeBaz(s)
}
