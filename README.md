# Philz-Coffee-Shop-Simulation
This is a Simulation of a Coffee shop where we depict how customers make and order, followed by a payment and then leave the system. 
This is written in Go lang and the communication between customer process, barista process, cashier process and the Queue happens
using restful web services.

To test the system, follow the following steps.

1. Go to src/github.com/kb/restwebservice folder
2. Run 'go run coffee_shop.go' 
3. Run 'go run barista.go'
4. Run 'go run cashier.go'
5. Run 'go run customer.go krishna mocha medium yes low yes 1341 localhost'

The following interaction takes place:

1. Customer is Created by the line no. 5 and it joins the Order Queue
2. The Barista who is continously checking the Order Queue gets the ip and port of the customer by dequeuing it from the Order Queue.
3. Barista starts making rest calls to customer.( Taking order and serving it.) 
4. The Customer then joins the Payment Queue.
5. The Cashier who is continously checking the Payment Queue gets the ip and port of the customer by dequeuing it from the Payment Queue.
6. The Cashier starts making rest calls to customer.( Requesting payment and giving a Receipt.) 
7. Customer process exists. (Leaves the Coffee Shop)
