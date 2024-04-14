SELECT 
	u.pair_id
	, tip_c.pair_name
	, u.order_id
	, u.side
	, u.order_type
	, u.target_price
	, u.amount_order
	, u.order_status_id
	, tip_o.order_status_desc
	, u.stop_loss
	, u.category
FROM dbo.user_orders u
LEFT JOIN dbo.tip_order_status tip_o ON tip_o.order_status_id = u.order_status_id
LEFT JOIN dbo.tip_trading_perp_pair tip_c ON tip_c.pair_id = u.pair_id


SELECT * FROM dbo.tip_order_status
SELECT * FROM dbo.tip_trading_perp_pair

SELECT * FROM dbo.user_orders

--FUNCTION
CREATE OR REPLACE FUNCTION dbo.insert_new_order_id(
	pair_string VARCHAR(50)
	, order_id VARCHAR(100)
	, side VARCHAR(5)
	, order_type VARCHAR(6)
	, target_price VARCHAR(100)
	, amount_order VARCHAR(100)
	, order_status_string VARCHAR(50)
	--targets
	, take_profit VARCHAR(100)
	, stop_loss VARCHAR(100)
	
	, category VARCHAR(100)
)
RETURNS text AS $$
DECLARE
    id_pair INT;
	id_order_status INT;
BEGIN

	/*--FIND PAIR ID-*/
	SELECT pair_id INTO id_pair FROM dbo.tip_trading_perp_pair WHERE pair_name = TRIM(pair_string);
	--if not exists, insert new coin
	IF id_pair IS NULL THEN
        INSERT INTO dbo.tip_trading_perp_pair (pair_name, flg_active, ins_user)
        VALUES (pair_string, true, 1)
        RETURNING pair_id INTO id_pair;
    END IF;
	/*---------------*/
	/*--FIND ORDER STATUS ID-*/
	SELECT order_status_id INTO id_order_status FROM dbo.tip_order_status WHERE order_status_desc = TRIM(order_status_string);
	/*---------------*/
    INSERT INTO dbo.user_orders (
		user_id
		, pair_id
		, order_id
		, side
		
		, order_type
		, target_price
		, amount_order
		
		, order_status_id
		, take_profit
		, stop_loss
		, category
		, ins_user
	)
    VALUES (
		1
		, id_pair
		, order_id
		, side
		
		, order_type
		, target_price
		, amount_order
		
		, id_order_status
		, take_profit
		, stop_loss
		, category
		, 1
	);
	-- Succeded
    RETURN 'Operation succeded';
EXCEPTION
    WHEN OTHERS THEN
    -- Error
    RETURN 'Failed: ' || SQLERRM;
END;
$$ LANGUAGE plpgsql;


SELECT * FROM dbo.tip_trading_perp_pair
SELECT * FROM dbo.tip_order_status
 

