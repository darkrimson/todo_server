package api

import (
	"net/http"
	"time"

	"github.com/username/go-final-project/pkg/utils"
)

func NextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.URL.Query().Get("now")
	dateStr := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	if nowStr == "" || dateStr == "" || repeat == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now, err := time.Parse(FORMAT_DATE, nowStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nextDate, err := utils.NextDate(now, dateStr, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
