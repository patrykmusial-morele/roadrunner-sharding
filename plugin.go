package sharding

import (
	"fmt"
)

type Plugin struct{}

type Configurer interface {
	UnmarshalKey(name string, out any) error
	Has(name string) bool
}

func (p *Plugin) Name() string {
	return "sharding"
}

func (p *Plugin) Init(cfg Configurer) error {
	fmt.Println("➡️ Plugin Hello World: Inicjalizacja...")
	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	fmt.Println("✨✨✨ HELLO WORLD PLUGIN AKTYWOWANY! ✨✨✨")

	return errCh
}

func (p *Plugin) Stop() error {
	fmt.Println("👋 Plugin Hello World: Zatrzymywanie...")
	return nil
}