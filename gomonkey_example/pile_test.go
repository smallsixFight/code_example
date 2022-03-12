package gomonkey_example

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/code_example/gomoney_example/pkg2"
	"reflect"

	//. "github.com/agiledragon/gomonkey/v2"
	//. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// go test -gcflags "all=-l" -v -run TestFuncA

func TestApplyFuncVarSeq(t *testing.T) {
	convey.Convey("Test ApplyFuncVarSeq", t, func() {
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{1, nil}},
			{Values: gomonkey.Params{2, nil}},
			{Values: gomonkey.Params{3, nil}},
		}
		patches := gomonkey.ApplyFuncVarSeq(&pkg2.FuncVar, outputs)
		defer patches.Reset()
		res, err := pkg2.FuncVar(1, 4)
		convey.So(err, convey.ShouldBeEmpty)
		convey.So(res, convey.ShouldEqual, 1)
		res, err = pkg2.FuncVar(1, 4)
		convey.So(res, convey.ShouldEqual, 2)
		res, err = pkg2.FuncVar(1, 4)
		convey.So(res, convey.ShouldEqual, 3)
		// 不能超过 outputs 数组设置的结果个数，否则会 panic
		//res, err = pkg2.FuncVar(1, 4)
		//convey.So(res, convey.ShouldEqual, 5)
	})
}

func TestApplyFuncVarReturn(t *testing.T) {
	convey.Convey("Test ApplyFuncReturn", t, func() {
		patches := gomonkey.ApplyFuncVarReturn(&pkg2.FuncVar, 0, errors.New("func var error"))
		defer patches.Reset()
		res, err := pkg2.FuncVar(1, 2)
		convey.So(err, convey.ShouldBeError, errors.New("func var error"))
		convey.So(res, convey.ShouldEqual, 0)
	})
}

func TestApplyFuncVar(t *testing.T) {
	convey.Convey("Test ApplyFuncVar", t, func() {
		patches := gomonkey.ApplyFuncVar(&pkg2.FuncVar, func(_, _ int) (c int, err error) {
			return 1, errors.New("func var error")
		})
		defer patches.Reset()
		res, err := pkg2.FuncVar(1, 2)
		convey.So(err, convey.ShouldBeError, errors.New("func var error"))
		convey.So(res, convey.ShouldEqual, 1)
	})
}

func TestApplyGlobalVal(t *testing.T) {
	convey.Convey("Test ApplyGlobalVal", t, func() {
		patches := gomonkey.ApplyGlobalVar(&pkg2.GlobalValue, "change global")
		defer patches.Reset()
		convey.So(pkg2.GlobalValue, convey.ShouldEqual, "change global")
	})
}

// go 1.17 不支持 ApplyPrivateMethod()，参考 https://github.com/agiledragon/gomonkey/releases/tag/v2.3.0 的解释
func TestApplyPrivateMethod(t *testing.T) {
	//var a pkg2.A
	//convey.Convey("Test ApplyPrivateMethod", t, func() {
	//	convey.Convey("Test privateMethodValue", func() {
	//		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(a), "privateMethodValue", func(_ pkg2.A) error {
	//			return errors.New("private method value error")
	//		})
	//		defer patches.Reset()
	//		convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("private method value error"))
	//	})
	//	convey.Convey("Test privateMethodPoint", func() {
	//		patches := gomonkey.ApplyPrivateMethod(reflect.TypeOf(&a), "privateMethodPoint", func(_ *pkg2.A) error {
	//			return errors.New("private method point error")
	//		})
	//		defer patches.Reset()
	//		convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("private method point error"))
	//	})
	//})
}

func TestApplyMethodSeq(t *testing.T) {
	var a pkg2.A
	convey.Convey("Test ApplyMethodSeq", t, func() {
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{errors.New("error 1")}},
			{Values: gomonkey.Params{errors.New("error 2")}},
			{Values: gomonkey.Params{errors.New("error 3")}},
		}
		patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(&a), "PointMethodA", outputs)
		defer patches.Reset()
		convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("error 1"))
		convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("error 2"))
		convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("error 3"))
	})
}

func TestApplyMethodReturn(t *testing.T) {
	var a pkg2.A
	var aa *pkg2.A
	convey.Convey("Test ApplyMethodReturn", t, func() {
		convey.Convey("Test value", func() {
			patches := gomonkey.ApplyMethodReturn(&a, "PointMethodA", errors.New("error 1"))
			defer patches.Reset()
			convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("error 1"))
			convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("error 1"))
		})
		convey.Convey("Test point", func() {
			patches := gomonkey.ApplyMethodReturn(aa, "PointMethodA", errors.New("error 1"))
			defer patches.Reset()
			convey.So(aa.PointMethodA(), convey.ShouldBeError, errors.New("error 1"))
			convey.So(aa.PointMethodA(), convey.ShouldBeError, errors.New("error 1"))
		})
	})
}

func TestApplyMethodFunc(t *testing.T) {
	var a pkg2.A
	convey.Convey("Test ApplyMethodFunc", t, func() {
		// 对比 ApplyMethod()，是不用使用传类型接收器
		convey.Convey("Test PointMethod", func() {
			patches := gomonkey.ApplyMethodFunc(reflect.TypeOf(&a), "PointMethodA", func() error {
				return errors.New("PointMethodA error first")
			})
			defer patches.Reset()
			convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("PointMethodA error first"))
		})
		convey.Convey("Test Method", func() {
			patches := gomonkey.ApplyMethodFunc(reflect.TypeOf(a), "MethodA", func() error {
				return errors.New("MethodA error first")
			})
			defer patches.Reset()
			convey.So(a.MethodA(), convey.ShouldBeError, errors.New("MethodA error first"))
		})
	})
}

func TestApplyMethod(t *testing.T) {
	a := pkg2.A{}
	convey.Convey("Test ApplyMethod", t, func() {
		convey.Convey("Test ApplyPointMethod", func() {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(&a), "PointMethodA", func(_ *pkg2.A) error {
				return errors.New("PointMethodA error first")
			})
			defer patches.Reset()
			convey.So(a.PointMethodA(), convey.ShouldBeError, errors.New("PointMethodA error first"))
		})
		convey.Convey("Test Method", func() {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(a), "MethodA", func(_ pkg2.A) error {
				return errors.New("MethodA error first")
			})
			defer patches.Reset()
			convey.So(a.MethodA(), convey.ShouldBeError, errors.New("MethodA error first"))
		})
	})
}

func TestApplyFuncReturn(t *testing.T) {
	convey.Convey("Test ApplyFuncReturn", t, func() {
		patches := gomonkey.ApplyFuncReturn(pkg2.FuncB, "a")
		defer patches.Reset()
		convey.So(pkg2.FuncB(""), convey.ShouldEqual, "a")
		convey.So(pkg2.FuncB(""), convey.ShouldEqual, "a")
	})
}

func TestApplyFuncSeq(t *testing.T) {
	convey.Convey("Test ApplyFuncSeq", t, func() {
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{"a"}},
			{Values: gomonkey.Params{"b"}},
			{Values: gomonkey.Params{"c"}},
		}
		patches := gomonkey.ApplyFuncSeq(pkg2.FuncB, outputs)
		defer patches.Reset()
		convey.So(pkg2.FuncB(""), convey.ShouldEqual, "a")
		convey.So(pkg2.FuncB(""), convey.ShouldEqual, "b")
		convey.So(pkg2.FuncB(""), convey.ShouldEqual, "c")
	})
}

func TestApplyFunc(t *testing.T) {
	convey.Convey("Test ApplyFunc", t, func() {
		convey.Convey("Test FuncA", t, func() {
			patches := gomonkey.ApplyFunc(pkg2.FuncA, func() string {
				return "pile func"
			})
			defer patches.Reset()
			t.Log("pkg2.FuncA() exec:", pkg2.FuncA())
		})
		convey.Convey("Test FuncB", t, func() {
			patches := gomonkey.ApplyFunc(pkg2.FuncB, func(a string) string {
				return "pile func B"
			})
			defer patches.Reset()
			convey.So(pkg2.FuncB("a"), convey.ShouldEqual, "pile func B")
		})
	})
}
