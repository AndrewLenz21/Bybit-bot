package bot_service

import "strings"

type ChannelA struct{}
type ChannelB struct{}

type SignalParser interface {
	ObtainSymbol(line string) string
	ObtainSide(line string) string
	ObtainEntry(line string) string
}

/************Channel 1 parser***********/
func (c ChannelA) ObtainSymbol(line string) string {
	symbol := ""
	//Symbol is between '#' and '/'
	hashIndex := strings.Index(line, "#")
	slashIndex := strings.Index(line, "/")
	//If there is no slash or hash, just return empty string
	if hashIndex != -1 && slashIndex != -1 {
		symbol = line[hashIndex+1:slashIndex] + "USDT" //for bybit api ex: BTCUSDT
	}
	return symbol
}

func (c ChannelA) ObtainSide(line string) string {
	side := ""
	//Side is after parentesis '('
	parenIndex := strings.Index(line, "(")
	if parenIndex != -1 {
		sideSubstring := line[parenIndex+1:]
		switch {
		case strings.Contains(sideSubstring, "Short"):
			side = "Sell"
		case strings.Contains(sideSubstring, "Long"):
			side = "Buy"
		}
	}
	return side
}

func (c ChannelA) ObtainEntry(line string) string {
	// Entry price is after the '-'
	entry := ""
	if strings.Contains(line, "Entry - ") {
		entryParts := strings.Split(line, "Entry - ")
		entry = strings.TrimSpace(entryParts[1])
	}
	return entry
}

/**********Channel 2 Parser**********/
func (c ChannelB) ObtainSymbol(lines []string) string {
	symbol := ""
	lineWithHash := ""
	//Symbol is on line with '#'
	for _, line := range lines {
		if strings.Contains(line, "#") {
			lineWithHash = line
			break // Encontrada la primera línea con '#', detener la búsqueda
		}
	}
	// Symbol is between '#' and '/'
	if lineWithHash != "" {
		hashIndex := strings.Index(lineWithHash, "#")
		//slashIndex := strings.Index(lineWithHash, "/")
		slashIndex := strings.Index(lineWithHash[hashIndex:], "/") + hashIndex // the first '/' after the '#'

		//If there is no slash or hash, just return empty string
		if hashIndex != -1 && slashIndex != -1 {
			symbol = lineWithHash[hashIndex+1:slashIndex] + "USDT" //for bybit api ex: BTCUSDT
		}
	}

	return symbol
}

func (c ChannelB) ObtainSide(lines []string) string {
	side := ""
	lineWithHash := ""

	//Symbol is on line with '#'
	for _, line := range lines {
		if strings.Contains(line, "#") {
			lineWithHash = strings.ToLower(line)
			break // Encontrada la primera línea con '#', detener la búsqueda
		}
	}
	//Side is the first word
	if lineWithHash != "" {
		switch {
		case strings.Contains(lineWithHash, "short"):
			side = "Sell"
		case strings.Contains(lineWithHash, "long"):
			side = "Buy"
		}
	}

	return side
}

func (c ChannelB) ObtainEntry(lines []string) string {
	// Entry price is after the '-'
	entry := ""
	lineWithEntry := ""

	//Symbol is on line with '#'
	for _, line := range lines {
		if strings.Contains(line, "Entry Point - ") {
			lineWithEntry = line
			break // Encontrada la primera línea con '#', detener la búsqueda
		}
	}
	if lineWithEntry != "" {
		entryParts := strings.Split(lineWithEntry, "Entry Point - ")
		entry = strings.TrimSpace(entryParts[1])
	}
	return entry
}
