package main

// #cgo pkg-config: gtk+-3.0
// #include "main.h"
import "C"

func main() {
	C.start()
}
