package bot_service

import "fmt"

var allowedChannelIDs = map[int64]bool{
	1763199802: true, // testing messages channel
	1813229153: true, // channer sender
	5572573112: true, // BTC scalp bot (my bot))
	// SIGNALS
	1129940546: true, // CHANNEL 1
	1717037581: true, // CHANNEL 2
}

/*
This bot will use:
1. Format message for channel 1:
🔥 #JUP/USDT (Long📈, x20) 🔥

Entry - 1.209
Take-Profit:

🥉 1.2337 (40% of profit)
🥈 1.2464 (60% of profit)
🥇 1.2594 (80% of profit)
🚀 1.2726 (100% of profit)

2. Format message for channel 2:
Long/Buy #PIXEL/USDT

# Entry Point - 4685

Targets: 4705 - 4725 - 4740 - 4775
Leverage - 10x
Stop Loss - 4490

BOT: We will parse the price and side, then we will use different configuration for the bot
*/

func SignalReceived(channel int64, username string, idMessage int, message string, timestamp int) {
	if VerifyChatID(channel) {
		symbol, entry, side := ParseSignalParams(channel, username, message)

		// If the parse is successful, get the position
		if symbol != "" && entry != "" && side != "" {
			fmt.Printf("Parsed Items:  \nSymbol -> %s \nEntry -> %s \nSide -> %s\n", symbol, entry, side)
			order := NewTradingPosition(symbol, side, entry)
			order.ObtainCoinInfo()          //Obtain the coin rules from Bybit: MinQty, MaxQty, Ticker for decimals
			order.ObtainUserConfiguration() //Obtain the configuration from database: StopLoss, Targets, Max Loss
			order.SetLeverage()             //Do not use a leverage different from user configuration

			//After SetLeverage we controlled if we already have positions
			if order.flg_approval {
				if channel == 1717037581 { //To resolve decimals problems from Channel 2
					order.CompareEntryPrice()
				}
				// We got the configuration and the signal, lets define the order amount and targets (TP and SL)
				order.SetOrderParameters()
				order.OpenNewOrder() //THIS ORDER WILL CALL THE BYBIT HANDLER
			} else {
				fmt.Printf("Order on -> %s \n Side -> %s \n NOT APPROVED", symbol, side)
				//fmt.Println("POSSITION NOT APPROVED")
			}
		}
	}
}

// PARSE THE ALLOWED CHANNELS
func VerifyChatID(id int64) bool {
	_, found := allowedChannelIDs[id]
	return found
}
