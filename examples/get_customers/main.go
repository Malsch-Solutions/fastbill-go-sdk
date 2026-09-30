// Command get_customers prints every customer of the FastBill account as JSON.
//
//	FASTBILL_EMAIL=… FASTBILL_API_KEY=… go run ./examples/get_customers
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/customer"
)

func main() {
	ctx := context.Background()
	client := fastbill.NewClient(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))
	customers := customer.NewClient(client)

	all, err := fastbill.All(ctx, fastbill.MaxLimit, func(ctx context.Context, page fastbill.Page) ([]customer.Customer, error) {
		return customers.Get(ctx, page, nil)
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(all, "", "  ")
	os.Stdout.Write(append(out, '\n'))
}
