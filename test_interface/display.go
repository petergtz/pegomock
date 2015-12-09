package test_interface

// Display is some sample interface to be mocked.
type Display interface {
	Flash(_param0 string, _param1 int)
	Show(_param0 string)
	SomeValue() string
	MultipleValues() (string, int, float32)
}
