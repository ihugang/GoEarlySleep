package main

import "github.com/kardianos/service"

var (
	arg string
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	doIt()
}

func (p *program) Stop(s service.Service) error {
	return nil
}
