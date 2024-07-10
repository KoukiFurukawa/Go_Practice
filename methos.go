package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"
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
	var d_string string = "2024-07-02 5:12"
	var date Date
	date.ConvertDateString(d_string)
	fmt.Println(date)
	var date2 Date2
	date2.ConvertDateString(d_string)
	fmt.Println(date2)
}

type Date2 struct {
	dateTime   time.Time
	dateString string
}

// Next.js から受け取った TimeStamp: string を Date型 にする
// 受け取る値は yyyy-MM-dd hh:mm を想定
// 引数の形をチェックするのは今回パス
func (d *Date2) ConvertDateString(dateString string) {
	layout := "2006-01-02 15:04"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("InValid dateString")
		return
	}

	d.dateTime = parsedTime

	weekDays := [7]string{"日", "月", "火", "水", "木", "金", "土"}
	var result bytes.Buffer
	dateField := map[string]string{
		"Year":    strconv.Itoa(parsedTime.Year()),
		"Month":   strconv.Itoa(int(parsedTime.Month())),
		"Day":     strconv.Itoa(parsedTime.Day()),
		"WeekDay": weekDays[parsedTime.Weekday()],
		"Hour":    strconv.Itoa(parsedTime.Hour()),
		"Minute":  strconv.Itoa(parsedTime.Minute()),
	}
	for key, value := range dateField {
		if key != "WeekDay" && len(value) < 2 {
			dateField[key] = "0" + value
		}
	}
	text, err := template.New("resultDate").Parse("{{.Year}}/{{.Month}}/{{.Day}} ({{.WeekDay}}) {{.Hour}}:{{.Minute}}")
	if err != nil {
		log.Fatal(err)
	}
	if err = text.Execute(&result, dateField); err != nil {
		log.Fatal(err)
	} else {
		d.dateString = result.String()
	}
}
