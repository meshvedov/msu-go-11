package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentQuarter(t *testing.T) {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := NewCalendar(parsed)
		actual := calendar.CurrentQuarter()
		if actual != test.quarter {
			t.Error("Month:", test.month,
				"Expected Quarter:", test.quarter,
				"Actual Quarter:", actual)
		}
	}
}

type Calendar struct {
	month   string
	quarter int
}

func NewCalendar(data time.Time) Calendar {
	switch m := data.Month(); m {
	case 1:
		return Calendar{"01", 1}
	case 2:
		return Calendar{"02", 1}
	case 3:
		return Calendar{"03", 1}
	case 4:
		return Calendar{"04", 2}
	case 5:
		return Calendar{"05", 2}
	case 6:
		return Calendar{"06", 2}
	case 7:
		return Calendar{"07", 3}
	case 8:
		return Calendar{"08", 3}
	case 9:
		return Calendar{"09", 3}
	case 10:
		return Calendar{"10", 4}
	case 11:
		return Calendar{"11", 4}
	case 12:
		return Calendar{"12", 4}
	}
	return Calendar{"00", 0}
}

func (cal Calendar) CurrentQuarter() int {
	return cal.quarter
}
