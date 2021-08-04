package main

import (
	"encoding/json"
	"github.com/malsch-solutions/fastbill-go-sdk/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/session"
	"log"
	"os"
)

func main() {
	fastbillSession := session.NewSession(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))

	customerClient := customer.NewCustomerClient(fastbillSession)

	customers, err := customerClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(json.Marshal(customers))
}
