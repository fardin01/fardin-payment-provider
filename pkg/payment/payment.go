package payment

import (
	"encoding/json"
	"fmt"
)

type Payment struct {
	CustomerId int
	Value float64
	Currency map[string]string
	Result bool
	// I think putting the HTTP status code here is bad practice, as it's related to the underlying web server and has
	// nothing to do with payments. But for the sake of this mock server, it's easier to put it here.
	StatusCode int
}

var count = 0

func shouldSucceed() bool {
	// This function returns false every other time it's called, so we can fail 50% of the requests.

	// The current state of this function does not allow this app to be stateless, as the value of count is hold in memory.
	// Therefore we cannot cluster this app. Make sure replicaCount is set to 1 until the app is made stateless.
	count++
	return count%2 != 0
}

func Pay(invoice []byte) Payment {
	// We should not rely solely on Antaeus, and we need a mechanism to make sure Antaeus is not requesting > 1 payment
	// for the same invoice. And of course many other validations that I'm not implementing here.

	var payment Payment
	err := json.Unmarshal(invoice, &payment)

	if err != nil {
		fmt.Printf("Failed to process the request: %v", err)
		payment.Result = false
		payment.StatusCode = 400
		return payment
	}

	fmt.Println("Attempting to pay invoice: ", string(invoice))
	if shouldSucceed(){
		fmt.Printf("Successfully paid invoice for customer with ID: %v and value of %v", payment.CustomerId, payment.Value)
		payment.Result = true
		payment.StatusCode = 200
		return payment
	}
	// Ideally, APIs should not return 500 if we are failing something on purpose, but I could not think of a better
	// status code.
	fmt.Printf("Failed to pay invoice for customer with ID: %v and value of %v", payment.CustomerId, payment.Value)
	payment.Result = false
	payment.StatusCode = 500
	return payment
}
