package queue

import (
	"container/list"
	
)

type Customer struct {
	ip string `json:"id" bson:"id"` 
	port string `json:"id" bson:"id"`
}


// Define Queue class
type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (this *Queue) Push(elem interface{}) {
	
	this.list.PushFront(elem)
}

// func (this *Queue) Push(elem Customer) {
	
// 	this.list.PushFront(elem)
// }

func (this *Queue) Pop() (interface{}, bool) {
	
	if this.list.Len() == 0 {
		return nil, false	
	}

// func (this *Queue) Pop() (Customer, bool) {
	
// 	if this.list.Len() == 0 {
// 		c := Customer {
// 			ip: "emptyip",
// 			port: "emptyport",
// 	}
// 		return c, false	
// 	}


// 	element := this.list.Back()
// 	this.list.Remove(element)
// 	 c:= Customer {
// 	 		ip: element.value.ip,
// 			port: element.value.port,
// 	 }
// 	c:= element(Customer)
// 	return c, true
// }

func (this *Queue) Size() int {
	
	return this.list.Len()
}

func (this *Queue) IsEmpty() bool {
	
	return (this.list.Len() == 0)
}