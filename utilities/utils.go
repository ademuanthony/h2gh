package utilities

import (
	"time"
)

func GetLastDayOfTheMonth(year int, month time.Month) int {
	t := time.Date(year, month, 32, 23, 59, 59, 999, time.UTC)
	return 32 - t.Day()
}

const (
	PackageAmount float64 = 5000
	UnitRebateAmount float64 = 4500
	ReferralBonusAmount = 500
)