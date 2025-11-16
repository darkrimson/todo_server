package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/username/go-final-project/pkg/consts"
)

type RepeatRule struct {
	Type      string
	Days      int
	WeekDays  []int
	MonthDays []int
	Months    []int
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
		weekDays := []int{}
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
		if len(parts) < 2 || len(parts) > 3 {
			return nil, fmt.Errorf("invalid m rule: %s", repeat)
		}

		monthDays := []int{}
		for _, day := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(day)
			if err != nil {
				return nil, fmt.Errorf("invalid repeat: %s", err)
			}

			if !((n >= 1 && n <= 31) || n == -1 || n == -2) {
				return nil, fmt.Errorf("invalid m day: %d", n)
			}
			monthDays = append(monthDays, n)
		}
		var months []int
		if len(parts) == 3 {
			for _, m := range strings.Split(parts[2], ",") {
				n, err := strconv.Atoi(m)
				if err != nil || n < 1 || n > 12 {
					return nil, fmt.Errorf("invalid repeat: %s", repeat)
				}
				months = append(months, n)
			}
		}
		return &RepeatRule{Type: "m", Months: months, MonthDays: monthDays}, nil
	}
	return nil, fmt.Errorf("invalid format repeat: %s", repeat)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	parseRepeat, err := parseRepeatRule(repeat)
	if err != nil {
		return "", err
	}

	date, err := time.Parse(consts.FORMAT_DATE, dstart)
	if err != nil {
		return "", err
	}

	switch parseRepeat.Type {
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if date.After(now) {
				break
			}
		}
	case "d":
		for {
			date = date.AddDate(0, 0, parseRepeat.Days)
			if date.After(now) {
				break
			}
		}
	case "w":
		repeatWDays := make(map[int]bool)
		for _, d := range parseRepeat.WeekDays {
			repeatWDays[d] = true
		}
		for {
			date = date.AddDate(0, 0, 1)

			wd := weekdayToHuman(date.Weekday())

			if repeatWDays[wd] && date.After(now) {
				return date.Format(consts.FORMAT_DATE), nil
			}

		}
	case "m":
		for {
			months := parseRepeat.Months
			if len(months) == 0 {
				months = make([]int, 12)
				for i := 1; i <= 12; i++ {
					months[i-1] = i
				}
			}

			candidates := []time.Time{}
			for _, month := range months {
				// Пропускаем месяцы в прошлом относительно originalInputDate
				if int(date.Month()) > month {
					continue
				}

				for _, d := range parseRepeat.MonthDays {
					var day int
					if d > 0 {
						day = d
					} else {
						daysInTargetMonth := daysInMonth(date.Year(), month)
						day = daysInTargetMonth + d + 1
					}
					if day < 1 || day > daysInMonth(date.Year(), month) {
						continue
					}

					// Если год и месяц совпадают с originalInputDate, пропускаем дни <= originalInputDate.Day()
					if month == int(date.Month()) && day <= date.Day() {
						continue
					}

					candidate := time.Date(date.Year(), time.Month(month), day, 0, 0, 0, 0, time.UTC)
					candidates = append(candidates, candidate)
				}
			}

			if len(candidates) > 0 {
				var minCandidate time.Time
				for _, c := range candidates {
					if c.After(now) {
						if minCandidate.IsZero() || c.Before(minCandidate) {
							minCandidate = c
						}
					}
				}
				if !minCandidate.IsZero() {
					date = minCandidate
					break
				}
			}
			date = time.Date(date.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
		}
	default:
		return "", fmt.Errorf("invalid repeat: %s", repeat)
	}

	return date.Format(consts.FORMAT_DATE), nil
}

func weekdayToHuman(w time.Weekday) int {
	if w == time.Sunday {
		return 7
	}
	return int(w)
}

func daysInMonth(year, month int) int {
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
