package interfaces

type Foo interface{
	Bar(int) (int, error)
	Baz(string) (string, error)
}