package boot_starter

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type IBoot interface {
	Boot()
	WaitingExit()
	Quit()
}

type BootApp struct {
}

func (i *BootApp) Quit() {
	for _, starter := range GetAllStarter() {
		starter.Stop()
	}
}

func (i *BootApp) WaitingExit() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		c := <-sig
		log.Println("接收到退出信号:", c)
		for _, starter := range GetAllStarter() {
			starter.Stop()
		}
	}
}

func (i *BootApp) Boot() {
	if len(GetAllStarter()) == 0 {
		log.Println("没有任何starter")
		return
	}
	i.init()
	i.setup()
	i.start()
}

func (i *BootApp) init() {
	for _, starter := range GetAllStarter() {
		starter.Init()
	}
}

func (i *BootApp) setup() {
	for _, starter := range GetAllStarter() {
		starter.Setup()
	}
}

func (i *BootApp) start() {
	for _, starter := range GetAllStarter() {
		if starter.Isblocking() {
			go starter.Start()
		} else {
			starter.Start()
		}
	}
}
