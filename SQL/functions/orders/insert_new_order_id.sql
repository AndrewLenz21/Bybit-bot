CREATE OR REPLACE FUNCTION dbo_trading_bot.insert_new_order_id(
	user_id_in integer,
	pair_string character varying,
	order_id_in character varying,
	side character varying,
	order_type_in character varying,
	target_price character varying,
	amount_order character varying,
	exec_amount_order character varying,
	order_status_string character varying,
	take_profit character varying,
	stop_loss character varying,
	category character varying,
	create_type_in character varying,
	reduce_only_in boolean)
    RETURNS text
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
    id_pair INT;
	id_order_status INT;
	id_order character varying;
BEGIN

	/*--FIND PAIR ID-*/
	SELECT pair_id INTO id_pair 
	FROM dbo_trading_bot.tip_trading_perp_pair 
	WHERE pair_name = TRIM(pair_string);
	--if not exists, insert new coin
	IF id_pair IS NULL THEN
        INSERT INTO dbo_trading_bot.tip_trading_perp_pair (pair_name, flg_active, ins_user)
        VALUES (pair_string, true, user_id_in)
        RETURNING pair_id INTO id_pair;
    END IF;
	
	/*--FIND ORDER STATUS ID-*/
	SELECT order_status_id INTO id_order_status 
	FROM dbo_trading_bot.tip_order_status 
	WHERE order_status_desc = TRIM(order_status_string);
	
	/*-----Verify if order_id exists with the same status-----*/
	SELECT order_id INTO id_order
	FROM dbo_trading_bot.user_orders
	WHERE 
		order_id = order_id_in
		--AND reduce_only = reduce_only_in
		AND order_status_id = id_order_status
		--AND order_type = order_type_in
		--AND create_type = create_type_in
		AND flg_last_status = true;
	--if not exists, update the last status, and insert this new order_status
	IF id_order IS NULL THEN
		--Update the last status
		UPDATE dbo_trading_bot.user_orders
		SET 
			flg_last_status = false
			, upt_user = user_id_in
			, upt_date = CURRENT_TIMESTAMP
		WHERE order_id = order_id_in
		AND flg_last_status = true;
		
		--Insert new record
		INSERT INTO dbo_trading_bot.user_orders (
			user_id
			, pair_id
			, order_id
			, side
			
			, order_type
			, target_price
			, amount_order
			, exec_amount_order
			
			, order_status_id
			, take_profit
			, stop_loss
			, category
			, create_type
			
			, reduce_only
			, ins_user
		)
    	VALUES (
			user_id_in
			, id_pair
			, order_id_in
			, side
			
			, order_type_in
			, target_price
			, amount_order
			, exec_amount_order
			
			, id_order_status
			, take_profit
			, stop_loss
			, category
			, create_type_in
			
			, reduce_only_in
			, user_id_in
		);
		--flg_last_status -> DEFAULT = true
	END IF;
	/*---------------*/
    
	-- Succeded
    RETURN 'Operation succeded';
EXCEPTION
    WHEN OTHERS THEN
    -- Error
    RETURN 'Failed: ' || SQLERRM;
END;
$BODY$;