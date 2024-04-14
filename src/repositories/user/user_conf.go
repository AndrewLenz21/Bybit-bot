package user

import (
	"bybitbot/src/config/postgres"
)

func (r *UserRepo) GetUserConfiguration(user int) ([]UserTradingConf, error) {
	sql := postgres.NewQueryService(r.pool)
	rows, err := sql.SelectCall("get_user_trading_conf", user)
	if err != nil {
		return nil, err
	}
	//MAP THE ROWS TO THE ENTITY
	return MapUserTradingConf(rows)
}
