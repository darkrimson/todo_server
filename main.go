package main

import (
	"fmt"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	repeat, _ := NextDate(time.Now(), "20240113", "d 7")
	fmt.Println(repeat)
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.Run()
	if err != nil {
		fmt.Errorf("server run error: %v", err)
		return
	}
}

type RepeatRule struct {
	Type      string
	Days      int
	WeekDays  []int
	MonthDays []int
}

func parseRepeatRule(repeat string) (*RepeatRule, error) {
	if repeat == "" {
		return nil, fmt.Errorf("empty repeat")
	}

	parts := strings.Fields(repeat)
	switch parts[0] {
	case "y":
		return &RepeatRule{Type: "y"}, nil

	case "d":
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid len repeat: %s", repeat)
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil || n < 1 || n > 400 {
			return nil, fmt.Errorf("invalid d repeat: %s", repeat)
		}
		return &RepeatRule{Type: "d", Days: n}, nil

	case "w":
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid WeekDays repeat: %s", repeat)
		}
		var weekDays = make([]int, 5)
		dayStrings := strings.Split(parts[1], ",")
		for _, day := range dayStrings {
			n, err := strconv.Atoi(day)
			if err != nil || n < 1 || n > 7 {
				return nil, fmt.Errorf("invalid WeekDays repeat: %s", repeat)
			}
			weekDays = append(weekDays, n)
		}

		if len(weekDays) == 0 {
			return nil, fmt.Errorf("invalid WeekDays repeat: %s", repeat)
		}

		return &RepeatRule{Type: "w", WeekDays: weekDays}, nil

	case "m":
		return &RepeatRule{Type: "m"}, nil
	}

	return nil, fmt.Errorf("invalid format repeat: %s", repeat)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	parseRepeat, err := parseRepeatRule(repeat)
	if err != nil {
		return "", err
	}

	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}

	switch parseRepeat.Type {
	case "y":
		for !date.After(now) {
			date = date.AddDate(1, 0, 0)
		}
	case "d":
		for !date.After(now) {
			date = date.AddDate(0, 0, parseRepeat.Days)
		}
	case "w":
		repeatWDays := make(map[int]bool)
		for _, d := range parseRepeat.WeekDays {
			repeatWDays[d] = true
		}
		for !date.After(now) {
			wd := weekdayToHuman(date.Weekday())

			if repeatWDays[wd] {
				return "", nil
			}

			date = date.AddDate(0, 0, 1)
		}
	default:
		return "", fmt.Errorf("invalid repeat: %s", repeat)
	}

	return date.Format("20060102"), nil
}

func weekdayToHuman(w time.Weekday) int {
	if w == time.Sunday {
		return 7
	}
	return int(w)
}

//parts := strings.Fields(repeat)
//switch parts[0] {
//case "y":
//// годовое повторение
//case "d":
//// повторение через дни
//case "w":
//// дни недели
//case "m":
//// дни месяца
//default:
//return "", fmt.Errorf("unsupported repeat format: %s", repeat)
//}
