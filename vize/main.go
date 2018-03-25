package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("/Users/thebigapple/Desktop/social-promo.mp4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(b))
}
