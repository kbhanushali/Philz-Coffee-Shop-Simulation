package main
	
import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"log"
	"github.com/kb/new_queue"
)
 
var pQueue = queue.NewQueue(1)
var oQueue = queue.NewQueue(1)

func startServer() {
	
	http.HandleFunc("/", WelcomeMessage)
	http.HandleFunc("/enqueueOrderQueue", EnqueueOrderQueue)
	http.HandleFunc("/enqueuePaymentQueue", EnqueuePaymentQueue)
	http.HandleFunc("/dequeueOrderQueue", DequeueOrderQueue)
	http.HandleFunc("/dequeuePaymentQueue", DequeuePaymentQueue)
    log.Fatal(http.ListenAndServe("localhost:1330", nil))
}



func WelcomeMessage(w http.ResponseWriter, r *http.Request) {

	
	fmt.Fprintln(w, "Welcome to Philz Coffee !")
	
}

func EnqueueOrderQueue(w http.ResponseWriter, r *http.Request) {

	u, err := url.Parse(r.URL.String())
    if err != nil {
        panic(err)
    }
	m, _ := url.ParseQuery(u.RawQuery)
	
	oQueue.Push(m["ip"][0],m["name"][0],m["port"][0])
	
	fmt.Fprintln(w, m["name"][0] +" Joined Order Queue")
	fmt.Println(m["name"][0] +" Joined Order Queue")
	
}

func EnqueuePaymentQueue(w http.ResponseWriter, r *http.Request) {

	u, err := url.Parse(r.URL.String())
    if err != nil {
        panic(err)
    }
	m, _ := url.ParseQuery(u.RawQuery)
	
	pQueue.Push(m["ip"][0],m["name"][0],m["port"][0])
	
	fmt.Fprintln(w, m["name"][0] +" Joined Payment Queue")
	fmt.Println(m["name"][0] +" Joined Payment Queue")
	
}

func DequeueOrderQueue(w http.ResponseWriter, r *http.Request) {
	
	port:= "0000"
	ip:="00.00.00.00"
	name:=""
	customer := make(map[string]string)

	//fmt.Printf("Number of people waiting to Place order : %d\n", oQueue.Size())
	if(!oQueue.IsEmpty()) {
		ip, name, port= oQueue.Pop()
	// fmt.Printf("Person with ip:  %v name: %v port: %v popped out\n",ip, name, port)
			customer["ip"] = ip
			customer["name"] = name
			customer["port"]= port
		
	}
	if customer["name"] != ""{
	fmt.Printf("%v is Placing Order \n",customer["name"])
}
	fmt.Printf("Number of people remaining to Place order : %d\n", oQueue.Size())



	response, err := json.MarshalIndent(customer,""," ")  
		if err != nil{
			panic(err)
		}
		fmt.Fprintf(w, string(response))
	
	
}

func DequeuePaymentQueue(w http.ResponseWriter, r *http.Request) {
	
	port:= "0000"
	ip:="00.00.00.00"
	name:=""
	customer := make(map[string]string)

	//fmt.Printf("Number of people waiting to Place order : %d\n", pQueue.Size())
	if(!pQueue.IsEmpty()) {
		ip, name, port= pQueue.Pop()
	// fmt.Printf("Person with ip:  %v name: %v port: %v popped out\n",ip, name, port)
	// fmt.Printf("Number of people remaining to Place order : %d\n", pQueue.Size())
			customer["ip"] = ip
			customer["name"] = name
			customer["port"]= port
		
	}

	if customer["name"] != ""{
	fmt.Printf("%v is making the Payment\n",customer["name"])
}
	fmt.Printf("Number of people remaining to make payment : %d\n", oQueue.Size())	
	
	response, err := json.MarshalIndent(customer,""," ")  
		if err != nil{
			panic(err)
		}
		fmt.Fprintf(w, string(response))
	
}


func main() {
	
	startServer()
	
}