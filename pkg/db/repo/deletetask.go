package repo

import (
	"database/sql"
	"fmt"
	"strconv"

	"go1f/pkg/db"
)

func DeleteTask(idStr string) error {

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid task id format: %w", err)
	}

	query := `
		DELETE FROM scheduler WHERE id = :id
`

	res, err := db.DB.Exec(query, sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("failed to delete task in DB: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete task in DB: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
