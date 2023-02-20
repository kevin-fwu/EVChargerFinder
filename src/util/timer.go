package util

import "time"

type TimerSec interface {
	OnTimer()
}

type Timer10Sec interface {
	OnTimer10Sec()
}

type Timer30Sec interface {
	OnTimer30Sec()
}

type Timer60Sec interface {
	OnTimer60Sec()
}

type Timer300Sec interface {
	OnTimer300Sec()
}

type Timer3600Sec interface {
	OnTimer3600Sec()
}

type Timer86400Sec interface {
	OnTimer86400Sec()
}

var counter10Sec int = 0
var counter30Sec int = 0
var counter60Sec int = 0
var counter300Sec int = 0
var counter3600Sec int = 0
var counter86400Sec int = 0

var timerListSec []TimerSec
var timerList10Sec []Timer10Sec
var timerList30Sec []Timer30Sec
var timerList60Sec []Timer60Sec
var timerList300Sec []Timer300Sec
var timerList3600Sec []Timer3600Sec
var timerList86400Sec []Timer86400Sec

func RegisterTimer(t interface{}) {
	if cast, ok := t.(TimerSec); ok {
		timerListSec = append(timerListSec, cast)
	}

	if cast, ok := t.(Timer10Sec); ok {
		timerList10Sec = append(timerList10Sec, cast)
	}

	if cast, ok := t.(Timer30Sec); ok {
		timerList30Sec = append(timerList30Sec, cast)
	}

	if cast, ok := t.(Timer60Sec); ok {
		timerList60Sec = append(timerList60Sec, cast)
	}

	if cast, ok := t.(Timer300Sec); ok {
		timerList300Sec = append(timerList300Sec, cast)
	}

	if cast, ok := t.(Timer3600Sec); ok {
		timerList3600Sec = append(timerList3600Sec, cast)
	}

	if cast, ok := t.(Timer86400Sec); ok {
		timerList86400Sec = append(timerList86400Sec, cast)
	}

}
func StartTimer() {
	go func() {
	labelStopTimer:
		for {
			select {
			case <-chStop:
				break labelStopTimer
			default:
			}

			onTimer()

			time.Sleep(1 * time.Second)
		}
	}()
}

func onTimer() {
	for _, t := range timerListSec {
		t.OnTimer()
	}
	counter10Sec++

	if counter10Sec >= 10 {
		onTimer10Sec()
		counter10Sec = 0
	}
}

func onTimer10Sec() {
	for _, t := range timerList10Sec {
		t.OnTimer10Sec()
	}
	counter30Sec++

	if counter30Sec >= 3 {
		onTimer30Sec()
		counter30Sec = 0
	}
}

func onTimer30Sec() {
	for _, t := range timerList30Sec {
		t.OnTimer30Sec()
	}
	counter60Sec++

	if counter60Sec >= 2 {
		onTimer60Sec()
		counter60Sec = 0
	}
}

func onTimer60Sec() {
	for _, t := range timerList60Sec {
		t.OnTimer60Sec()
	}
	counter300Sec++

	if counter300Sec >= 5 {
		onTimer300Sec()
		counter300Sec = 0
	}
}

func onTimer300Sec() {
	for _, t := range timerList300Sec {
		t.OnTimer300Sec()
	}
	counter3600Sec++

	if counter3600Sec >= 12 {
		onTimer3600Sec()
		counter3600Sec = 0
	}
}

func onTimer3600Sec() {
	for _, t := range timerList3600Sec {
		t.OnTimer3600Sec()
	}
	counter86400Sec++

	if counter86400Sec >= 24 {
		onTimer86400Sec()
		counter86400Sec = 0
	}
}

func onTimer86400Sec() {
	for _, t := range timerList86400Sec {
		t.OnTimer86400Sec()
	}
}
