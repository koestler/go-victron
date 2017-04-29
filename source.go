package dataflow

import (
	"github.com/koestler/go-ve-sensor/bmv"
	"time"
	"log"
	"math/rand"
)

func SourceBmvStartDummy(device *Device, registers bmv.Registers) <-chan Value {
	output := make(chan Value)
	go func() {
		for _ = range time.Tick(time.Second) {
			log.Print("SourceBmvStartDummy tik");
			for name, register := range registers {
				output <- Value{
					Name:   name,
					Device: device,
					Value:  rand.Float64(),
					Unit:   register.Unit,
				}
			}
		}
	}()
	return output
}
