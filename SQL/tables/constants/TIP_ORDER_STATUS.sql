CREATE TABLE dbo_trading_bot.TIP_ORDER_STATUS
(
	order_status_id SERIAL PRIMARY KEY
    , order_status_desc VARCHAR(50) NOT NULL UNIQUE
	, main_order_status VARCHAR(50)
    , flg_active BOOLEAN DEFAULT TRUE
	
    , ins_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , ins_user INT
    , upt_date TIMESTAMP WITH TIME ZONE
    , upt_user INT
	--KEYS
	, FOREIGN KEY (ins_user) REFERENCES USERS(user_id)
)

SELECT * FROM dbo.TIP_ORDER_STATUS

/*
INSERT INTO dbo.TIP_ORDER_STATUS
(
	order_status_desc
	, main_order_status
	, ins_user
)
VALUES
(
	'Deactivated'
	, 'Closed'
	, 1
)
*/

/*
ORDER STATUS 

--open status

New order has been placed successfully
PartiallyFilled
Untriggered Conditional orders are created

--closed status

Rejected
PartiallyFilledCanceled Only spot has this order status
Filled
Cancelled In derivatives, orders with this status may have an executed qty
Triggered instantaneous state for conditional orders from Untriggered to New
Deactivated UTA: Spot tp/sl order, conditional order, OCO order are cancelled before they are triggered
*/