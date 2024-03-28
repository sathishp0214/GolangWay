package main

import "fmt"

func main() {
	o := HBAutomation{} //particular automation

	HealthCheckers["HBAutomation"] = o
	HealthCheckers["HBAutomation1"] = HBAutomation1{}
	RunAall()
}

var HealthCheckers = map[string]HealthChecker{} //stores all healthcheck interface implemented automations

type HealthChecker interface {
	Name()
	Enabled() bool
	Run()
	Publish()
}

//running all interface implmented healthcheck automations
func RunAall() {
	for _, h := range HealthCheckers {

		if h.Enabled() {
			h.Run()
		}

	}
}

type HBAutomation struct {
	HealthChecker //compostion with healthcheck interface
}

func (h HBAutomation) Run() {
	fmt.Println("HB automation is  running")
}

func (h HBAutomation) Enabled() bool {
	return true
}

type HBAutomation1 struct {
	HealthChecker //compostion with healthcheck interface
}

func (h HBAutomation1) Run() {
	fmt.Println("HB automation1 is  running")
}

func (h HBAutomation1) Enabled() bool {
	return true
}
