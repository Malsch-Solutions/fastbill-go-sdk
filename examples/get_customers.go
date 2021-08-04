package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/malsch-solutions/fastbill-go-sdk/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
)

func main() {
	fastbillService := service.NewService(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))

	customerClient := customer.NewCustomerClient(fastbillService)

	customers, err := customerClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}
	customerJSON, _ := json.Marshal(customers)
	log.Println(string(customerJSON))
}
