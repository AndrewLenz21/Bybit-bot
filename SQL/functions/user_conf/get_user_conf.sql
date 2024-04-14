CREATE OR REPLACE FUNCTION dbo.get_user_trading_conf(_user_id INT)
RETURNS TABLE (
    id_configuration INT,
    user_id INT,
    pair_id INT,
    amount_per_position NUMERIC(6,4),
    max_loss_per_position NUMERIC(6,4),
    take_profit NUMERIC(6,4),
    stop_loss NUMERIC(6,4),
    trailing_stop NUMERIC(6,4),
    default_first_target NUMERIC(6,4),
    amount_per_first_target NUMERIC(6,4),
    default_next_target NUMERIC(6,4),
    amount_per_next_target NUMERIC(6,4),
    long_leverage NUMERIC(4,2),
    short_leverage NUMERIC(4,2)
)
LANGUAGE plpgsql
AS $$
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
    FROM dbo.user_trading_conf uc
    WHERE uc.user_id = _user_id;
END;
$$;

SELECT * FROM dbo.get_user_trading_conf(1); -- Reemplaza 1 con el ID de configuración específico que quieras consultar.

