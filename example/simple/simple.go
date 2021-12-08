package main

import "github.com/heart-dance-x/tower"

func main() {
	ss := tower.NewBootStrap(nil)
	ss.Listen()
}
