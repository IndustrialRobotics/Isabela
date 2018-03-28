package main

import (
	"github.com/goburrow/modbus"
	"time"
	"log"
	"os"
	"fmt"
)

//https://github.com/goburrow/modbus

func main() {

	// Modbus TCP
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.Address = "192.168.137.23"
	handler.SlaveId = 0xff
	handler.IdleTimeout = 25 * time.Second
	handler.Logger = log.New(os.Stdout, "Test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session

	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(15, 2)
	results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 2, 0, 4})
	results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})

	if err != nil {
		fmt.Println("Could not connect to the Modbus device...")
	}
	fmt.Println(results)
}
