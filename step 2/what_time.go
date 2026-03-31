package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	err = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func TimeNow() time.Time {
	return time.Now()
}

// ! Текущий день недели
func currentDayOfTheWeek() string {
	day := TimeNow().Weekday()

	switch day {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	default:
		return "Ошибка"
	}
}

// ! check dayOrNight
func dayOrNight() string {
	now := TimeNow()

	hour := now.Hour()

	if hour >= 10 && hour <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

// ! Проверка сколько до пятницы
func nextFriday() int {
	day := TimeNow().Weekday()
	return (5 - int(day) + 7) % 7
}

// !проверка какой день недели сейчас
func CheckCurrentDayOfTheWeek(answer string) bool {
	ans := strings.ToLower(answer)

	realDay := strings.ToLower(currentDayOfTheWeek())
	return ans == realDay
}

// ! проверка что сейчас день или ночь
func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, err
	}

	realTime := strings.ToLower(dayOrNight())

	userAnswer := strings.ToLower(answer)

	if realTime == userAnswer {
		return true, nil
	}
	return false, nil
}
