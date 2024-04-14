package bot_service

var allowedChannelIDs = map[int64]bool{
	1763199802: true, // tester 1
	1813229153: true, // tester 2
	5572573112: true, // tester 3
	// SIGNALS
	1129940546: true, // Channel 1
	1717037581: true, // Channel 2
}

// PARSING MESSAGE
func SignalReceived(sender string, idMessage int, channel int64, timestamp int, message string) {
	if VerifyChatID(channel) {
		symbol, entry, side := ParseSignalParams(channel, message)

		// If the parse is successful, get the position
		if symbol != "" && entry != "" && side != "" {
			order := NewTradingPosition(symbol, side, entry)
			order.ObtainCoinInfo()
			order.ObtainUserConfiguration()
			order.SetLeverage()
			if channel == 5572573112 {
				order.CompareEntryPrice()
			}
			// We got the configuration and the signal, lets define the order amount and targets (TP and SL)
			order.SetOrderParameters()
			order.OpenNewOrder() //THIS ORDER WILL CALL THE BYBIT HANDLER
		}

		// db := test.NewTestRepo(postgres.GetPool())
		// db.ProvaInsert(int(channel), message)
	}
}

// PARSE THE ALLOWED CHANNELS
func VerifyChatID(id int64) bool {
	_, found := allowedChannelIDs[id]
	return found
}
