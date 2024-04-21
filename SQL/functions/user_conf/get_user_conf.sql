CREATE OR REPLACE FUNCTION dbo_trading_bot.get_user_trading_conf(
	_id_configuration integer)
    RETURNS TABLE(id_configuration integer, user_id integer, pair_id integer, amount_per_position numeric, max_loss_per_position numeric, take_profit numeric, stop_loss numeric, trailing_stop numeric, default_first_target numeric, amount_per_first_target numeric, default_next_target numeric, amount_per_next_target numeric, long_leverage numeric, short_leverage numeric) 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
    RETURN QUERY
    SELECT 
        uc.id_configuration,
        uc.user_id,
        uc.pair_id,
        uc.amount_per_position,
        uc.max_loss_per_position,
        uc.take_profit,
        uc.stop_loss,
        uc.trailing_stop,
        uc.default_first_target,
        uc.amount_per_first_target,
        uc.default_next_target,
        uc.amount_per_next_target,
        uc.long_leverage,
        uc.short_leverage
    FROM dbo_trading_bot.user_trading_conf uc
    WHERE uc.id_configuration = _id_configuration;
END;
$BODY$;