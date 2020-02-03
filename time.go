package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Location())

	then := time.Date(2020, 2, 3, 13, 00, 00, 651387273, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(now.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err == nil {
		fmt.Println(now.In(loc))
	} else {
		fmt.Println("loi")
	}
	t1 := now.AddDate(-1, 2, 3)
	p(t1)
	t2 := time.Now().AddDate(1, 2, 3)
	p(t2)
}
