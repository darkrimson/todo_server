package utils

import (
	"fmt"
	"time"

	"go1f/pkg/consts"

	"go1f/pkg/db/model"
)

func CheckDate(task *model.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(consts.FORMAT_DATE)
	}
	t, err := time.Parse(consts.FORMAT_DATE, task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format, must be YYYYMMDD")
	}

	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("invalid repeat rule %w", err)
		}

		if !AfterNow(now, t) {
			task.Date = next
		}
	} else {
		if !AfterNow(now, t) {
			task.Date = now.Format(consts.FORMAT_DATE)
		}
	}

	return nil
}
