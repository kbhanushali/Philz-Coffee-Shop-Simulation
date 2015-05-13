package main
	
import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"strings"

)

var coffeeShopIP ="127.0.0.1"

 
 var custUrl, coffeeName, cream, sugar, xHot, size string

func checkCustomer() {

	for
	{	
		var customer map[string]string

		url :="http://"+coffeeShopIP+":1330/dequeueOrderQueue"
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
	
	k := GetCoffeeName()
	Whatsize()

	WantCream()
	WantSugar()

	if k != "frappechino"{
		Extra_hot()
	}

	RepeatOrder()

	TakeCoffee()
	}else{
	fmt.Println("No one in Order Queue")
}
	time.Sleep(3*time.Second)
	}
}
	

func GetCoffeeName() string {

	fmt.Println("Hey, How may I help you today?")
  
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
	fmt.Println("Okay, you want "+coffeeName+"\n")

	return coffeeName
}	

func WantCream() {

	fmt.Println("Extra cream?")

	  	time.Sleep(3*time.Second)


	url :=custUrl+"cream"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	cream = strings.TrimSpace(string(body))
	if cream !="no"{
	fmt.Println("Extra cream checked!\n")
	}

}

func WantSugar() {

	fmt.Println("How much sugar?")

	  	time.Sleep(3*time.Second)


	url :=custUrl+"sugar"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	sugar = strings.TrimSpace(string(body))
	fmt.Println("Your coffee's Sweetness will be : "+sugar+"\n")

}

func Extra_hot() {

	fmt.Println("Extra hot?")

	  	time.Sleep(3*time.Second)


	url :=custUrl+"extra_hot"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		panic(err)
	}

	xHot = strings.TrimSpace(string(body))
	if( xHot != "no"){
	fmt.Println("Will make it extra hot.")

}}


func Whatsize() {

	fmt.Println("What size would you like?")

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
	fmt.Println("A "+size+" sized "+coffeeName+"\n")

}

func RepeatOrder() {


	fmt.Printf("Let me repeat your order.\n\nYou ordered a "+size+" "+coffeeName+" ")
	if cream == "yes"{
		fmt.Printf("with cream")}
	if xHot == "yes"{
		fmt.Printf(", Extra hot")}
	if sugar == "yes"{
		fmt.Printf(" and Sweetend.")}
	fmt.Printf("\n\n")
	url :=custUrl+"confirmOrder"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Will get your coffee ready!!")
	time.Sleep(10*time.Second)



}

func TakeCoffee() {

	fmt.Println("Your coffee is ready. Enjoy!")

	url :=custUrl+"takeCoffee"
	res, err :=http.Get(url)
	if err !=nil{
		panic(err)
	}
	defer res.Body.Close()

	time.Sleep(2*time.Second)

}


// func checkCustomer() {

		
// 	var customer map[string]string

// 		url :="http://localhost:1330/dequeueOrderQueue"
// 	res, err :=http.Get(url)
// 	if err !=nil{
// 		panic(err)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err !=nil{
// 		panic(err)
// 	}

// 	err = json.Unmarshal(body, &customer)
// 	if err !=nil{
// 		panic(err)
// 	}

// 	fmt.Println(customer["ip"], customer["port"], customer["name"])
	
// 	time.Sleep(3*time.Second)

// }	
	

func main() {
	
	checkCustomer()
	
}