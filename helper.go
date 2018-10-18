package main

import (
	"fmt"

	//webpush "workspace/snippets/webpush/github.com/sherclockholmes/webpush-go"
	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		// TODO: Handle failure!
	}
	fmt.Println(privateKey)
	fmt.Println(publicKey)
}
