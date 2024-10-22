package posgresql

import (
	"context"
	"errors"
	"strconv"
	"time"
)

var query string = "UPDATE tasks SET"

func (storage *PostgresStorage) Update(id int, title string, description string, status string) error {

	ctx := context.Background()

	conn := storage.conn

	pred := false

	params := make([]interface{}, 0, 5)

	if len(title) == 0 && len(description) == 0 && len(status) == 0 {
		return errors.New("no data to update")
	}

	count := 1

	if len(title) != 0 {
		query += " title = $" + strconv.Itoa(count)
		count++
		pred = true
		params = append(params, title)
	}

	if len(description) != 0 {
		if pred {
			query += ","
		}
		pred = true
		query += " description = $" + strconv.Itoa(count)
		count++
		params = append(params, description)
	}

	if len(status) != 0 {
		if pred {
			query += ","
		}
		query += " status = $" + strconv.Itoa(count)
		count++
		params = append(params, status)
	}

	query += ", time_last_update = $" + strconv.Itoa(count) + " WHERE id = $" + strconv.Itoa(count+1) + ";"

	params = append(params, time.Now().Format("2006-01-02 15:04:05"), strconv.Itoa(id))

	comandTag, err := conn.Exec(ctx, query,
		params...,
	)

	if err != nil {
		return err
	}

	if comandTag.RowsAffected() != 1 {
		return errors.New("no tasks to update")
	}

	return nil

}
