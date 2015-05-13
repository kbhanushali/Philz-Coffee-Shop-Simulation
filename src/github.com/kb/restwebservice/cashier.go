package main
	
import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"strings"

)

//var customerIP "127.0.0.1"
var coffeeShopIP = "127.0.0.1"

 
 var custUrl, coffeeName, cream, sugar, xHot, size, answer string

func checkCustomer() {

	for
	{	
		var customer map[string]string

		url :="http://"+coffeeShopIP+":1330/dequeuePaymentQueue"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	err = json.Unmarshal(body, &customer)
	if err !=nil{
		panic(err)
	}


	fmt.Println(customer["ip"], customer["port"], customer["name"])

	if customer["port"] != ""{
	custUrl ="http://"+customer["ip"]+customer["port"]+"/"
	
	GetCoffeeName()
	Whatsize()

	RequestPayment()

	TakeReceipt()
	} else{
	fmt.Println("No one in Payment  Queue")
}
	time.Sleep(3*time.Second)
	}
}
	

func GetCoffeeName() {

	fmt.Println("What was your order Sir?")
  
  	time.Sleep(3*time.Second)


	url :=custUrl+"coffee_name"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	coffeeName = strings.TrimSpace(string(body))
	fmt.Println("Okay, you had  a"+coffeeName+"\n")
}	



func Whatsize() {

	fmt.Println("What size ?")

	  	time.Sleep(1*time.Second)


	url :=custUrl+"size"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	size = strings.TrimSpace(string(body))
	if size == "small"{
	fmt.Println("Your total amount is 1 dollars\n")
	}

	if size == "medium"{
	fmt.Println("Your total amount is 3 dollars\n")
	}

	if size == "large"{
	fmt.Println("Your total amount is 5 dollars\n")
	}

	time.Sleep(1*time.Second)


}

func RequestPayment() {


	fmt.Printf("Would you like to have a reciept?\n\n")
	
	url :=custUrl+"requestPayment"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	answer = strings.TrimSpace(string(body))
	fmt.Println("Thank you for your Payment!!")
	time.Sleep(1*time.Second)

}

func TakeReceipt() {

	if answer != "no"{
	fmt.Println("Here is your Reciept.")
	}
	fmt.Println(" Thank you very much. Have a great day!!")

	url :=custUrl+"takeReceipt"
	http.Get(url)
	// 
	// if err !=nil{
	// 	panic(err)
	// }
	// defer res.Body.Close()
}
	

func main() {
	
	checkCustomer()
	
}