// Call other function from other place such as bot.go, config.go

package main

import (
	"fmt"

	"github.com/jinheehanaaa/simple-discord-ping/bot"
	"github.com/jinheehanaaa/simple-discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	// No need to close because there are no values to receive

	return
}
