CREATE OR REPLACE FUNCTION dbo_trading_bot.upsert_position(
	user_id_in integer,
	symbol character varying,
	entry_price_in character varying,
	side_in character varying,
	size_in character varying,
	position_value_in character varying,
	position_balance_in character varying,
	stop_loss_in character varying,
	take_profit_in character varying,
	trailing_stop_in character varying,
	position_idx_in integer,
	unrealized_pnl_in character varying,
	cur_realised_pnl_in character varying,
	cum_realised_pnl_in character varying,
	flg_active_in boolean)
    RETURNS text
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
    id_pair INT;
	id_order_status INT;
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
	IF id_position_found IS NULL THEN
		--Insert new position
		INSERT INTO dbo_trading_bot.user_positions (
			user_id
			, pair_id
			
			, entry_price
			, side
			, amount_contract
			, position_value
			, position_balance
			
			, stop_loss
			, take_profit
			, trailing_stop
			
			, position_idx
			, unrealized_pnl
			, cur_realised_pnl
			, cum_realised_pnl
			, total_pnl
			
			, ins_user
		)
    	VALUES (
			user_id_in
			, id_pair
			
			, entry_price_in
			, side_in
			, size_in
			, position_value_in
			, position_balance_in
			
			, stop_loss_in
			, take_profit_in
			, trailing_stop_in
			
			, position_idx_in
			, unrealized_pnl_in
			, cur_realised_pnl_in
			, cum_realised_pnl_in
			, CAST(cur_realised_pnl_in AS NUMERIC(8,4))
			
			, user_id_in
		);
		--flg_last_status -> DEFAULT = true
	ELSE
		--This position Exists, we are going to UPDATE
		UPDATE dbo_trading_bot.user_positions
		SET 
			entry_price = entry_price_in
			, amount_contract = size_in
			, position_value = position_value_in
			, position_balance = position_balance_in
			, stop_loss = stop_loss_in
			, take_profit = take_profit_in
			, trailing_stop = trailing_stop_in
			, position_idx = position_idx_in
			, unrealized_pnl = unrealized_pnl_in
			, cur_realised_pnl = cur_realised_pnl_in
			, cum_realised_pnl = cum_realised_pnl_in
			, total_pnl = CAST(cur_realised_pnl_in AS NUMERIC(8,4))
			, flg_active = flg_active_in
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