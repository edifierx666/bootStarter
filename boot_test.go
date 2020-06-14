package bootStarter

import (
	"log"
	"testing"
)

type S1 struct {
	BasicStarter
}

func (i *S1) Name() string {
	return "S1"
}
func (i *S1) Init() {
	log.Println("init")
}
func TestBootApp_Boot(t *testing.T) {
	app := BootApp{}
	RegisterStarter(&S1{})
	app.Boot()
	app.WaitingExit()
}
