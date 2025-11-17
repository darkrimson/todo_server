package repo

import (
	"database/sql"
	"fmt"
	"strconv"

	"go1f/pkg/db"
	"go1f/pkg/db/model"
)

func GetTask(idStr string) (*model.Task, error) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid task id format: %w", err)
	}

	var idInt64 int64
	query := `
		SELECT * FROM scheduler WHERE id=:id
`
	var task model.Task
	err = db.DB.QueryRow(query, sql.Named("id", id)).Scan(
		&idInt64, &task.Date, &task.Title, &task.Comment, &task.Repeat,
	)
	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}

	task.ID = strconv.FormatInt(idInt64, 10)

	return &task, nil
}
