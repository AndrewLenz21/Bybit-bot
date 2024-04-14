CREATE TABLE dbo.USER_POSITIONS 
(
	id_position SERIAL PRIMARY KEY
	, user_id 	INT NOT NULL
	, pair_id INT NOT NULL    --BYBIT SYMBOL
	--POSITION
	, entry_price VARCHAR(100)  --not all coins have the same decimals
	, amount_contract VARCHAR(100)   --Number of coins on position SIZE
	--RISK MANAGEMENT
	, position_value VARCHAR(100)
	, position_balance VARCHAR(100)
	, stop_loss VARCHAR(100)  
	, take_profit VARCHAR(100)
	, trailing_stop VARCHAR(100)   --could be null
	--DESC FROM BYBIT API
	, position_idx INT
	--ACTIVE
	, flg_active BOOLEAN DEFAULT TRUE   --position will be closed
	--REALIZED PNL
	, unrealized_pnl VARCHAR(100)
	, cur_realised_pnl VARCHAR(100)
	, cum_realised_pnl VARCHAR(100)
	, total_pnl DECIMAL(8,4)     --total PnL after closed position (USDT)
	--REG
    , ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , ins_user INT
    , upt_date TIMESTAMP WITH TIME ZONE  --This case we will update records
    , upt_user INT
	--KEYS
	, FOREIGN KEY (user_id) REFERENCES dbo.USERS(user_id)
	, FOREIGN KEY (pair_id) REFERENCES dbo.TIP_TRADING_PERP_PAIR(pair_id)
)
--for positions
/*
recognize the columns that we are going to use

id_position PRIMARY KEY IDENTITY
user_id   probably me
pair_id
entry_price
size  VARCHAR
position_value
position_balance
stop_loss  VARCHAR
take_profit VARCHAR
trailing_stop
position_idx INT

--pnls
unrealised_pnl
cur_realised_pnl
cum_realised_pnl

flg_active BOOLEAN

*/
--LATER: CREATE A TRIGGER TO UPDATE THIS TABLE BASED ON ORDERS


