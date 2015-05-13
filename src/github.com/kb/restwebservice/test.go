package main

import (
	"github.com/kb/queue"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"

	//"strconv"
	//"os"
	//"log"
)

type Order struct {
	Coffees [10]Coffee

	}

type Coffee map[string]string


type Customer struct {
	ip string `json:"id" bson:"id"` 
	port string `json:"id" bson:"id"`
}

// type Coffee struct {
// 	cofName string
// 	xtraHot bool
// 	xtraMilk bool
// 	xtraSweet bool
// 	custName string
// 	custId int
// 	size string

// }

//payment Queue	
var pQueue *queue.Queue 
//order Queue	
var oQueue *queue.Queue


func main() {
		pQueue = queue.NewQueue()
		oQueue = queue.NewQueue()
		http.HandleFunc("/enqueueOrderQueue", EnqueueOrderQueue)
		// http.HandleFunc("/enqueuePaymentQueue", EnqueuePaymentQueue)
		// http.HandleFunc("/dequeueOrderQueue", DequeueOrderQueue)
		// http.HandleFunc("/dequeuePaymentQueue", DequeuePaymentQueue)
		http.ListenAndServe("localhost:1330", nil)
}

//adds customer to Order Queue
func EnqueueOrderQueue(w http.ResponseWriter , r *http.Request) {
	u, err := url.Parse(r.URL.String())
    if err != nil {
        panic(err)
    }

	m, _ := url.ParseQuery(u.RawQuery)

	c := Customer {
			ip: m["name"][0],
			port: m["port"][0],
	}
	//var o Order

	 // for i:=0; i < len(o.Coffees); i++ {
	 // 		cof := make(map[string]string)
		// cof["cofName"] = "Mocha"
		// cof["xtraHot"] = "no"
		// cof["xtraMilk"]= "no"
		// cof["xtraSweet"] = "yes"
		// cof["custName"] = "A"
		// cof["custId"] = strconv.Itoa(i)
		// cof["size"] = "small"
		// o.Coffees[i] = cof 
		oQueue.Push(c)

	//}

		//var onm = Order{o.Coffees}
	 	fmt.Printf("Number of people waiting to Place order : %d\n", oQueue.Size())
	
	var c1 Customer
	if ele, ok := oQueue.Pop();ok {
 			fmt.Printf("Cutomer: %v was on top of Order Queue\n", ele)
 			//fmt.Printf("Number of people waiting to Place order : %d\n", oQueue.Size())
	 		c1 = ele
 		}

	response, err := json.MarshalIndent(c1, "", "\t")  
		if err != nil{
			panic(err)
		}

		//fmt.Println(string(response))

		fmt.Fprintf(w, string(response))
	  
}


// //adds customer to Payment Queue
// func EnqueuePaymentQueue(w http.ResponseWriter , r *http.Request) { 
// 	var p Payload
// 	for i:=0; i < len(p.order); i++ {
// 		p.order[i].cofName = "Mocha"
// 		p.order[i].xtraHot = false
// 		p.order[i].xtraMilk = false
// 		p.order[i].xtraSweet = true
// 		p.order[i].custName = "A"
// 		p.order[i].id = i
// 		pQueue.Push(p.order[i])
		
// 	}

// }


// //dequeues the order queue
// func DequeueOrderQueue(w http.ResponseWriter , r *http.Request) { 

// 	for i:=0; i < len(p.order); i++ {
// 		if ele, ok := oQueue.Pop(); ok {
// 			fmt.Printf("Cutomer: %v left the Order Queue\n", ele)
// 			fmt.Printf("Number of people waiting to Place order : %d\n", oQueue.Size())
// 		}
		
// 	}
	
// }

// //dequeues the payment queue
// func DequeuePaymentQueue(w http.ResponseWriter , r *http.Request) { 

// 	for i:=0; i < len(p.order); i++ {
// 		if ele, ok := pQueue.Pop(); ok {
// 			fmt.Printf("Cutomer: %v left the Order Queue\n", ele)
// 			fmt.Printf("Number of people waiting to Place order : %d\n", pQueue.Size())
// 		}
		
// 	}
	
// }
