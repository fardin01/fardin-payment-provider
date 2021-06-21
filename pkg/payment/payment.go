package payment

type Payment struct {
	Result bool
	// I think putting the HTTP status code here is bad practice, as it's related to the underlying web server and has
	// nothing to do with subscriptions. But for the sake of this mock server, it's easier to put it here.
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

func Pay() Payment {
	// We should not rely solely on Antaeus, and we need a mechanism to make sure Antaeus is not requesting > 1 payment
	// for the same invoice
	if shouldSucceed(){
		p := Payment{true, 201}
		return p
	}
	// Ideally, APIs should not return 500 if we are failing something on purpose, but I could not think of a better
	// status code.
	p := Payment{false, 500}
	return p
}
