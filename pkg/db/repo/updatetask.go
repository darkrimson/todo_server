package repo

import (
	"database/sql"
	"fmt"
	"strconv"

	"go1f/pkg/db"
	"go1f/pkg/db/model"
)

func UpdateTask(task *model.Task) error {
	id, err := strconv.ParseInt(task.ID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid task id in DB")
	}

	query := `
		UPDATE scheduler SET date = :date, 
			title = :title, 
			comment = :comment, 
			"repeat" = :repeat 
		WHERE id = :id
`
	res, err := db.DB.Exec(query, sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("failed to update task in DB: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed tp get affected rows: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("task not found in DB: %d", task.ID)
	}

	return nil
}
