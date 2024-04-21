CREATE OR REPLACE FUNCTION dbo_trading_bot.get_user_open_positions(
	user_id_in integer,
	symbol character varying,
	side_in character varying)
    RETURNS TABLE(id_position integer, side character varying, pair_id_out integer, amount_contract character varying) 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
DECLARE
    id_pair INT;
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
	
    RETURN QUERY
    SELECT 
        up.id_position,
        up.side,
        up.pair_id as pair_id_out,
        up.amount_contract
    FROM dbo_trading_bot.user_positions up
    WHERE 
		up.pair_id = id_pair
		AND up.side = side_in
		AND up.flg_active = true;
END;
$BODY$;