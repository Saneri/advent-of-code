package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
	"strings"
)

type ModuleType int

const (
	Flipflop ModuleType = iota
	Conjunction
	Broadcaster
)

type Module struct {
	Name         string
	Type         ModuleType
	State        bool
	Memory       map[string]Pulse
	Destinations []string
}

type Pulse bool

const (
	Low  Pulse = false
	High Pulse = true
)

type Modules map[string]*Module

type QueueItem struct {
	ModuleName string
	Pulse      Pulse
	source     string
}

func initiateModules(data []string) Modules {
	modules := make(Modules)
	for _, line := range data {
		split := strings.Split(line, " -> ")
		destinations := strings.Split(split[1], ", ")
		name := split[0]
		module := Module{State: false, Memory: make(map[string]Pulse), Destinations: destinations}
		if strings.HasPrefix(name, "%") {
			module.Type = Flipflop
			module.Name = strings.TrimPrefix(name, "%")
		} else if strings.HasPrefix(name, "&") {
			module.Type = Conjunction
			module.Name = strings.TrimPrefix(name, "&")
		} else if name == "broadcaster" {
			module.Type = Broadcaster
			module.Name = name
		} else {
			log.Fatal("Unknown:", name)
		}
		modules[module.Name] = &module
	}

	// init memory maps
	for _, module := range modules {
		for _, destination := range module.Destinations {
			destModule, ok := modules[destination]
			if ok {
				destModule.Memory[module.Name] = Low
			}
		}
	}
	return modules
}

func pushButton(modules Modules) (int, int) {
	queue := []QueueItem{}
	for _, destination := range modules["broadcaster"].Destinations {
		queue = append(queue, QueueItem{destination, Low, "broadcaster"})
	}
	addToQueue := func(moduleNames []string, pulse Pulse, source string) {
		for _, destination := range moduleNames {
			queue = append(queue, QueueItem{destination, pulse, source})
		}
	}

	highs := 0
	lows := 1 // button starts with low pulse
	for len(queue) > 0 {
		nextItem := queue[0]
		queue = queue[1:]
		module, ok := modules[nextItem.ModuleName]
		pulse := nextItem.Pulse
		if pulse == High {
			highs++
		} else {
			lows++
		}
		if !ok {
			continue
		}
		if module.Type == Flipflop {
			if pulse == Low {
				module.State = !module.State
				addToQueue(module.Destinations, Pulse(module.State), module.Name)
			} else {
				continue
			}
		} else if module.Type == Conjunction {
			module.Memory[nextItem.source] = pulse
			pulseToSend := Low
			for _, pulse := range module.Memory {
				if pulse == Low {
					pulseToSend = High
					break
				}
			}

			addToQueue(module.Destinations, pulseToSend, module.Name)
		} else {
			log.Fatal("Unexpected module type:", module.Type)
		}
	}
	return highs, lows
}

func main() {
	data := utils.ReadInput("input.txt")
	modules := initiateModules(data)

	highs, lows := 0, 0
	for i := 0; i < 1000; i++ {
		high, low := pushButton(modules)
		highs += high
		lows += low
	}
	fmt.Println("a:", highs*lows)
}
