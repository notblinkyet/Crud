package posgresql

import (
	"context"
	"errors"
)

func (storage *PostgresStorage) Delete(id int) error {
	ctx := context.Background()
	conn := storage.conn
	commandTag, err := conn.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("no row found to delete")
	}
	return nil
}
