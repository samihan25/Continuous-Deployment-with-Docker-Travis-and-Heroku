package main

import (
	"testing"
	"time"
)

func Test_hello(t *testing.T) {
	msg := hello()
	if msg != "Hello World!!" {
		t.Errorf("Hello function was incorrect, got: %s, want: %s.", msg, "Hello World!!")
	}
}

func Test_ist(t *testing.T) {
	msg := hello()
	if msg != "Hello World!!" {
		t.Errorf("hello function was incorrect, got: %s, want: %s.", msg, "Hello World!!")
	}

	nowTimeTemp, nowDateTemp := ist()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	nowTime := time.Now().In(loc).Format("3:04 PM")
	nowDate := time.Now().In(loc).Format("02-01-2006 Monday")
	if nowTimeTemp != nowTime {
		t.Errorf("ist function was incorrect for time, got: %s, want: %s.", nowTimeTemp, nowTime)
	}
	if nowDateTemp != nowDate {
		t.Errorf("ist function was incorrect for date, got: %s, want: %s.", nowDateTemp, nowDate)
	}
}
