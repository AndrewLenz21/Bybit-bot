CREATE TABLE dbo_trading_bot.USER_ORDERS
(
	id_order SERIAL PRIMARY KEY
	, user_id 	INT NOT NULL
	, pair_id INT NOT NULL    --BYBIT SYMBOL
	, order_id VARCHAR(100)   --BYBIT RESPONSE
	
	, side VARCHAR(5)  -- Sell or Buy
	, order_type VARCHAR(6)  --Limit or Market
	, target_price VARCHAR(100)   --price
	, amount_order VARCHAR(100)   --order amount
	, order_status_id INT   --order status - we need just the new or filled
	--targets
	, take_profit VARCHAR(100)  -- Take Profit in number
	, stop_loss VARCHAR(100)  -- Stop Loss in number
	--we cant put trailing stop loss here
	, category VARCHAR(100) 
	--REG
    , ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , ins_user INT
    , upt_date TIMESTAMP WITH TIME ZONE
    , upt_user INT
	--KEYS
	, FOREIGN KEY (user_id) REFERENCES dbo.USERS(user_id)
	, FOREIGN KEY (pair_id) REFERENCES dbo.TIP_TRADING_PERP_PAIR(pair_id)
	, FOREIGN KEY (order_status_id) REFERENCES dbo.TIP_ORDER_STATUS(order_status_id)
)