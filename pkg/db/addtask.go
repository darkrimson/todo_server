package db

import (
	"database/sql"
	"fmt"

	"go1f/pkg/db/model"
)

func AddTask(task *model.Task) (int64, error) {

	var id int64

	query := `
		INSERT INTO scheduler (date, title, comment, "repeat")
		VALUES (:date, :title, :comment, :repeat)
	`

	res, err := DB.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert task: %w", err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to last insert task: %w", err)
	}

	return id, nil
}
