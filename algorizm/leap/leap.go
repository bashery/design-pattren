package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

/*
import "time"

// IsLeapYear returns true if the year is a leap year.
func IsLeapYear(y int) bool {
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	//year.YearDay() returns the days of the year and returns 366 if it's a leap year.
	days := year.YearDay()
	if days > 365 {
		return true
	}
	return false
}
*/
