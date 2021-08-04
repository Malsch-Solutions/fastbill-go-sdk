package main

import (
	"github.com/malsch-solutions/fastbill-go-sdk/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/session"
	"os"
)

func main() {
	fastbillClient := session.NewSession(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))

	customerClient := customer.NewCustomerClient(fastbillClient)

	customers, err := customerClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)
}
