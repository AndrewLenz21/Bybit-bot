CREATE OR REPLACE FUNCTION dbo_trading_bot.close_position(
	user_id_in integer,
	symbol character varying,
	side_in character varying,
	last_bybit_pnl character varying)
    RETURNS text
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
    id_pair INT;
	id_position_found INT;
BEGIN
	/*--FIND PAIR ID-*/
	SELECT pair_id INTO id_pair 
	FROM dbo_trading_bot.tip_trading_perp_pair 
	WHERE pair_name = TRIM(symbol);
	--if not exists, insert new coin
	IF id_pair IS NULL THEN
        INSERT INTO dbo_trading_bot.tip_trading_perp_pair (pair_name, flg_active, ins_user)
        VALUES (symbol, true, user_id_in)
        RETURNING pair_id INTO id_pair;
    END IF;
	
	
	/*-----Verify if the position exists with flg_active = 1-----*/
	SELECT id_position INTO id_position_found
	FROM dbo_trading_bot.user_positions
	WHERE 
		user_id = user_id_in
		AND pair_id = id_pair
		AND side = side_in
		AND flg_active = true;
	--if not exists, update the last status, and insert this new order_status
	IF id_position_found IS NOT NULL THEN
		--Update position
		UPDATE dbo_trading_bot.user_positions
		SET 
			total_pnl = total_pnl + CAST(last_bybit_pnl AS NUMERIC(8,4))
			, flg_active = false
			, upt_user = user_id_in
			, upt_date = CURRENT_TIMESTAMP
		WHERE 
			id_position = id_position_found;
	END IF
	;
	/*---------------*/
    
	-- Succeded
    RETURN 'Operation succeded';
EXCEPTION
    WHEN OTHERS THEN
    -- Error
    RETURN 'Failed: ' || SQLERRM;
END;
$BODY$;