package bot_service

import (
	"fmt"
	"strings"
)

// Parsing
func ParseSignalParams(channel int64, username string, message string) (symbol string, entry string, side string) {
	// Divide in lines
	lines := strings.Split(message, "\n")
	// Channel 1
	if len(lines) == 9 && channel == 1717037581 {
		fmt.Println("/**/")
		fmt.Println("/************STEP 1*************/")
		symbol = ChannelA{}.ObtainSymbol(lines[0])
		side = ChannelA{}.ObtainSide(lines[0])
		entry = ChannelA{}.ObtainEntry(lines[2])
	}

	// Channel 2
	if len(lines) > 4 && channel == 1129940546 {
		fmt.Println("/**/")
		fmt.Println("/************STEP 1*************/")
		symbol = ChannelB{}.ObtainSymbol(lines)
		side = ChannelB{}.ObtainSide(lines)
		entry = ChannelB{}.ObtainEntry(lines)
	}

	fmt.Printf("RECEIVED:  \nFrom -> %s \nMessage -> %s \n", username, message)
	return symbol, entry, side
}
