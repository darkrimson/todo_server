package repo

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"go1f/pkg/consts"
	"go1f/pkg/db"
	"go1f/pkg/db/model"
)

func GetTasks(limit int, search string) ([]*model.Task, error) {

	if limit <= 0 {
		limit = 1
	}
	if limit > 100 {
		limit = 100
	}

	var rows *sql.Rows
	var err error

	if search != "" {
		if parsedDate, errParse := time.Parse("02.01.2006", search); errParse == nil {
			searchDate := parsedDate.Format(consts.FORMAT_DATE)
			querySearchByDate := `SELECT id, date, title, comment, "repeat" FROM scheduler WHERE date = :date ORDER BY date ASC LIMIT :limit`
			rows, err = db.DB.Query(querySearchByDate,
				sql.Named("date", searchDate),
				sql.Named("limit", limit),
			)
		} else {
			likePattern := "%" + search + "%"
			querySearchByText := `SELECT id, date, title, comment, "repeat" FROM scheduler WHERE title LIKE :pattern OR comment LIKE :pattern ORDER BY date ASC LIMIT :limit`
			rows, err = db.DB.Query(querySearchByText,
				sql.Named("pattern", likePattern),
				sql.Named("limit", limit),
			)
		}
	} else {
		queryLimit := `SELECT id, date, title, comment, "repeat" FROM scheduler ORDER BY date ASC LIMIT :limit`
		rows, err = db.DB.Query(queryLimit, sql.Named("limit", limit))
	}

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var id int64        // ✅ Сканируем id как int64
		var task model.Task // task.ID — string
		// 🔴 Не используем sql.Named в Scan — только *destination
		err := rows.Scan(&id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		// ✅ Преобразуем int64 в string
		task.ID = strconv.FormatInt(id, 10)
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration failed: %w", err)
	}

	if tasks == nil {
		tasks = make([]*model.Task, 0)
	}

	return tasks, nil
}
