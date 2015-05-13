package main
	
import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
)

type Prop struct {
	ip string
    name string
    coffee string
    cream string
	sugar string
	extra_hot string
	port string
	size string
}

var CoffeeShopHost= "127.0.0.1"

type Order struct {
	Coffees [10]Coffee

	}

type Coffee map[string]string


var cust = new(Prop)

func StartServer() {
	http.HandleFunc("/coffee_name", GetCoffeeName)
    http.HandleFunc("/cream", WantCream)
	http.HandleFunc("/sugar", WantSugar)
	http.HandleFunc("/extra_hot", WantExtraHot)
	http.HandleFunc("/size", GetSize)
	http.HandleFunc("/confirmOrder", ConfirmOrder)
	http.HandleFunc("/takeCoffee",TakeCoffee)
	http.HandleFunc("/requestPayment",PayBill)
	http.HandleFunc("/takeReceipt",TakeReceipt)

    log.Fatal(http.ListenAndServe(cust.port, nil))
}


func TakeReceipt(w http.ResponseWriter, r *http.Request) {
time.Sleep(1*time.Second)
		fmt.Println("You too. Bye")
	time.Sleep(1*time.Second)
    fmt.Fprintln(w, "User exited")
     time.Sleep(6*time.Second)
     os.Exit(2)
}

func PayBill(w http.ResponseWriter, r *http.Request) {
time.Sleep(1*time.Second)
		fmt.Println(cust.name+" Pays the bill")
		fmt.Println("I would like to have a receipt")
	time.Sleep(1*time.Second)

    fmt.Fprintln(w, "yes")
}


func GetCoffeeName(w http.ResponseWriter, r *http.Request) {
time.Sleep(1*time.Second)
		fmt.Println("A "+cust.coffee)
	time.Sleep(1*time.Second)

    fmt.Fprintln(w, cust.coffee)
}

func GetSize(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1*time.Second)
	  fmt.Println(cust.size)
	  	time.Sleep(1*time.Second)

    fmt.Fprintln(w, cust.size)
}

func WantCream(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1*time.Second)
		  fmt.Println("Cream ? "+cust.cream)
		  	time.Sleep(1*time.Second)

    fmt.Fprintln(w, cust.cream)
}

func WantSugar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1*time.Second)
	fmt.Println(cust.sugar+" sugar please.")
		time.Sleep(1*time.Second)

    fmt.Fprintln(w, cust.sugar)
}

func WantExtraHot(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1*time.Second)
		fmt.Println(cust.extra_hot)
			time.Sleep(1*time.Second)

    fmt.Fprintln(w, cust.extra_hot)
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Yes. Thats Right!\n")
	fmt.Fprintln(w, "Yes. Thats Right!\n")
	time.Sleep(1*time.Second)

}

func TakeCoffee(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1*time.Second)

	fmt.Println("The coffee is nice\n")
	time.Sleep(1*time.Second)
	AddToPaymentQueue()


}

func AddToOrderQueue() {
	resp, err := http.Get("http://"+CoffeeShopHost+":1330/enqueueOrderQueue?ip="+cust.ip+"&name="+cust.name+"&port="+cust.port)
	if err != nil {
	// handle error
	}
	fmt.Println(resp)
	//defer resp.Body.Close()

}

func AddToPaymentQueue() {
	resp, err := http.Get("http://"+CoffeeShopHost+":1330/enqueuePaymentQueue?ip="+cust.ip+"&name="+cust.name+"&port="+cust.port)
	if err != nil {
	// handle error
	}
	fmt.Println(resp)
	//defer resp.Body.Close()

}

func main() {
	cust.name = os.Args[1]
	cust.coffee = os.Args[2]
	cust.size = os.Args[3]
	cust.cream = os.Args[4]
	cust.sugar = os.Args[5] 
	cust.extra_hot = os.Args[6]
	cust.port = ":" + os.Args[7]
	cust.ip = os.Args[8]

	
	fmt.Println(cust.name)
	fmt.Println(cust.coffee)
	fmt.Println(cust.size)
	fmt.Println(cust.cream)
	fmt.Println(cust.sugar)
	fmt.Println(cust.extra_hot)
	fmt.Println(cust.port)
    fmt.Println(cust.ip)

	
	AddToOrderQueue()
	StartServer()
	
}