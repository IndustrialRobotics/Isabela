package main

import "fmt"

//EtherNet_IP- Part 3 - How to Control a GS4 Variable Frequency Drive (VFD) via EIP - YouTube.MP4
func main() {

	bs := []byte{35,18,19,69}
	byteConveter(bs)

}

func byteConveter(bs []byte) {

	fmt.Println(string(bs))
}
