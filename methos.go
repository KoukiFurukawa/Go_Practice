package main

import (
	"fmt"
	"time"
)

func add(a int, b int) (c int) {
	c = a + b
	fmt.Println(c)
	return
}

type Participation struct {
	age      uint8
	isMinor  bool
	isSenior bool
}

func (p Participation) IsMinor() bool {
	return p.age < 18
}
func (p Participation) IsSenior() bool {
	return p.age >= 60
}

func createUserInfo(user []uint8) []Participation {
	var userCount int = len(user)
	var participants = make([]Participation, userCount)
	for i := 0; i < userCount; i++ {
		participants[i].age = user[i]
		participants[i].isMinor = participants[i].IsMinor()
		participants[i].isSenior = participants[i].IsSenior()
	}
	fmt.Println(participants)
	return participants
}

type Date struct {
	dateTime   time.Time
	dayOfMonth int
	dayOfWeek  int
	hour       int
	minute     int
}

// Next.js から受け取った TimeStamp: string を Date型 にする
// 受け取る値は yyyy-MM-dd hh:mm を想定
// 引数の形をチェックするのは今回パス
func (d *Date) ConvertDateString(dateString string) {
	layout := "2006-01-02 15:04"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("InValid dateString")
		return
	}
	d.dateTime = parsedTime
	d.dayOfMonth = parsedTime.Day()
	d.dayOfWeek = int(parsedTime.Weekday())
	d.hour = parsedTime.Hour()
	d.minute = parsedTime.Minute()
}

func datetime_sample() {
	var d_string string = "2024-07-02 15:00"
	var date Date
	date.ConvertDateString(d_string)
	fmt.Println(date)
}
