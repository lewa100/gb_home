package queue

import (
	"fmt"
	ph "github.com/gb_home/hw3/phoneBook"
)

var (
	QR queue
	lastId = 1
)

type queue struct {
	people []ph.Person
}

// AddInQueue - add
func (q *queue) AddInQueue(name, phone string) {
	tmp := ph.Person{
		Id: lastId,
		Name:  name,
		Phone: phone,
	}
	lastId++
	q.people = append(q.people, tmp)
}

func (q *queue) GetFromQueue() (man ph.Person) {
	if len(q.people) == 0{
		fmt.Println("Queue is Empty")
		return
	}
	m := q.people[0]
	q.people = q.people[1:]
	return m
}