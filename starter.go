package bootStarter

import (
	"fmt"
	"reflect"
	"runtime"
)

type IStarter interface {
	Name() string
	Init()
	Setup()
	Start()
	Isblocking() bool
	Stop()
}

type BasicStarter struct{}

func (s BasicStarter) Name() string {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Sprintf("设置name,文件:%s,第%d行,%s", file, line, reflect.TypeOf(s)))
}

func (BasicStarter) Init() {
}

func (BasicStarter) Setup() {
}

func (BasicStarter) Start() {
}

func (BasicStarter) Isblocking() bool {
	return false
}

func (BasicStarter) Stop() {
}
