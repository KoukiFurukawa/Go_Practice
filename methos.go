package main

import (
	"fmt"
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
