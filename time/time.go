package string

import (
	"time"
)

func MonthList(start, end string) (result []string) {
	startMonth, err := time.Parse("200601", start)
	if err != nil {
		return nil
	}
	endMonth, err := time.Parse("200601", end)
	if err != nil {
		return nil
	}

	for startMonth.Before(endMonth) {
		result = append(result, startMonth.Format("200601"))
		startMonth = startMonth.AddDate(0, 1, 0)
	}
	//if not equal, then means start > end
	if startMonth == endMonth {
		result = append(result, end)
	}
	return result
}
