//This is a parser of all the victim and attacker details
//around a player hurt event. Demoinfocs(https://github.com/markus-wa/demoinfocs-golang/)
//has been used as the base parser
// USEAGE go run main.go "path_to_demos" -> outputs data to /data

package main

import (
	"io/ioutil"
	"log"
	"os"

	p "github.com/mwaurawakati/csgo/demoparser"
)

//export startparsing
func startparsing() {
	source_go := os.Args[1]
	files, err := ioutil.ReadDir(source_go)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		p.FrameRateCheck(source_go, f.Name())
	}
}

func main() {

	startparsing()

}
