package main

/*
#include "dht22.h"
#include <stdlib.h>       // for free()
#cgo LDFLAGS: -L. -lwiringPi
*/
import "C"

import (
	"errors"
)

type DHT22Data struct {
	temperature, humidity float32
}

var setup bool

func main() {
}

func Setup(pin int) error {
	success := C.setup(C.int(pin))

	if success != 0 {
		return errors.New("Unable to initalize, are you root?")
	}

	setup = true
	return nil
}

func GetData() (DHT22Data, error) {

	if setup == false {
		panic("Successful call to Setup(pin int) required")
	}

	sd := C.get_temp_and_humidity()

	if sd == nil {
		return DHT22Data{}, nil
	}

	return DHT22Data{
		temperature: float32(sd.temp),
		humidity:    float32(sd.humidity),
	}, nil

}
