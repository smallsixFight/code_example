package pkg2

import (
	"errors"
	"fmt"
)

var GlobalValue = "global_value"

func FuncA() string {
	return "FuncA"
}
func FuncB(a string) string {
	return a + " FuncB"
}

type A struct {
	PublicValue  string
	privateValue string
}

var FuncVar = func(a, b int) (c int, err error) {
	return a + b, fmt.Errorf("%d", a+b)
}

func (A) privateMethodValue() error {
	return nil
}

func (*A) privateMethodPoint() error {
	return nil
}

func (*A) PointMethodA() error {
	panic("point method A")
}

func (*A) PointMethodB() error {
	return errors.New("point method B")
}

func (A) MethodA() error {
	return errors.New("method A")
}
func (A) MethodB() error {
	return errors.New("method B")
}
