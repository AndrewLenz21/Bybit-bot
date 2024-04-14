package test

import (
	"bybitbot/src/config/postgres"
)

func (r *TestRepo) ProvaInsert(number int, description string) (string, error) {
	sql := postgres.NewQueryService(r.pool)
	msg, err := sql.InsertCall("insert_test_table_todo_func", number, description)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (r *TestRepo) ProvaSelect(number int) ([]TestTableTodo, error) {
	sql := postgres.NewQueryService(r.pool)
	rows, err := sql.SelectCall("select_from_test_table_todo", number)
	if err != nil {
		return nil, err
	}
	//MAP THE ROWS TO THE ENTITY
	return MapTestTableTodo(rows)
}
