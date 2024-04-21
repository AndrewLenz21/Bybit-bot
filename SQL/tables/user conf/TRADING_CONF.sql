CREATE TABLE dbo.USER_TRADING_CONF (
    id_configuration SERIAL PRIMARY KEY
    , user_id  INT NOT NULL
	, pair_id INT NOT NULL
    --TRADING
	--default positions
	, ammount_per_position DECIMAL(6,4) -- Percentual per Open an order
	, max_loss_per_position DECIMAL(6,4) -- Percentual per Loss in an order
	--general targets
	, take_profit DECIMAL(6, 4) -- Take Profit in Percentual
	, stop_loss DECIMAL(6, 4) -- Stop Loss in Percentual
	, trailing_stop DECIMAL(6, 4) -- Dinamic stop loss
	--risk management
    , default_first_target DECIMAL(6, 4) -- Percentual limit First target
	, amount_per_first_target DECIMAL(6, 4)  -- Percentual amount to take on first target
	
	, default_next_target DECIMAL(6, 4) -- It will be more targets
	, amount_per_next_target DECIMAL(6, 4)  -- Percentual amount to take on next target
	--leverage
	, long_leverage DECIMAL(4,2)
	, short_leverage DECIMAL(4,2)
	--REG
    , ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , ins_user INT
    , upt_date TIMESTAMP WITH TIME ZONE
    , upt_user INT
	--KEYS
    , FOREIGN KEY (user_id) REFERENCES dbo.USERS(user_id)
	, FOREIGN KEY (pair_id) REFERENCES dbo.TIP_TRADING_PERP_PAIR(pair_id)
);

SELECT * FROM dbo.user_trading_conf
/*
INSERT INTO dbo.user_trading_conf
(
	user_id
	, pair_id
	, max_loss_per_position
	
	, take_profit
	, stop_loss
	
	, default_first_target
	, amount_per_first_target
	
	, default_next_target
	, amount_per_next_target
	
	, long_leverage
	, short_leverage
	
	, ins_user
)
VALUES
(
	  1
	, 1 --DEFAULT PAIR
	, 0.5   --2% default ammount LOSS per position
	
	, 30.0000    --30% from entry price
	, 8.0000     --8% from entry price
	
	, 3.0000    --3% from entry price
	, 80.0000   --80% of entire position
	
	, 2.5000
	, 50.000
	
	, 3.00   --margin x3 long
	, 3.00   --margin x3 short
	
	, 1        --admin user
)

*/




