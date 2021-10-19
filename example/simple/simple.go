package main

import "github.com/go-tower/tower"

func main() {
	ss := tower.NewBootStrap(nil)
	ss.Listen()
}
