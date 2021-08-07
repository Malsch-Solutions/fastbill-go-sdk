package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/article"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/recurring"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/document"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/template"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/item"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/estimate"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/invoice"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/contact"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
)

func main() {
	fastbillService := service.NewService(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))

	handleArticle(fastbillService)
	return

	customerClient := customer.NewCustomerClient(fastbillService)
	log.Println("Create Customer")
	c, err := customerClient.Create(&customer.Customer{CustomerType: "consumer", FirstName: "foo", LastName: "bar 2"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update Customer")
	_, err = customerClient.Update(&customer.Customer{CustomerID: strconv.Itoa(c.CustomerID), FirstName: "alter"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Customer")
	list, err := customerClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	customerJSON, _ := json.Marshal(list)
	log.Println(string(customerJSON))

	//handleContacts(fastbillService, strconv.Itoa(c.CustomerID))
	//handleEstimates(fastbillService, strconv.Itoa(c.CustomerID))
	handleRecurring(fastbillService, strconv.Itoa(c.CustomerID))
	return
	i := handleInvoices(fastbillService, strconv.Itoa(c.CustomerID))

	handleItems(fastbillService, i)

	/*
		log.Println("Delete Customer")
		deleted, err := customerClient.Delete(strconv.Itoa(c.CustomerID))
		if err != nil {
			log.Fatal(err)
		}

		log.Println(deleted)
	*/
}

func handleItems(fastbillService service.Service, i int) {
	itemClient := item.NewItemClient(fastbillService)

	log.Println("Get items")
	items, err := itemClient.Get(&parameter.Parameter{Limit: 10}, &item.Filter{InvoiceID: i})
	if err != nil {
		log.Fatal(err)
	}

	itemsJson, _ := json.Marshal(items)
	log.Println(string(itemsJson))

	log.Println("delete item")
	_, err = itemClient.Delete(items[0].InvoiceItemID)
	if err != nil {
		log.Fatal(err)
	}
}

func handleContacts(fastbillService service.Service, customerID string) {
	contactClient := contact.NewContactClient(fastbillService)

	log.Println("Create Contact")
	c, err := contactClient.Create(&contact.Contact{CustomerID: customerID, FirstName: "foo", LastName: "bar"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update Contact")
	_, err = contactClient.Update(&contact.Contact{ContactID: strconv.Itoa(c.ContactID), FirstName: "alter"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Contact")
	list, err := contactClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	customerJSON, _ := json.Marshal(list)
	log.Println(string(customerJSON))
	log.Println("Delete Contact")
	deleted, err := contactClient.Delete(strconv.Itoa(c.ContactID), customerID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(deleted)
}

func handleInvoices(fastbillService service.Service, customerID string) int {
	invoiceClient := invoice.NewInvoiceClient(fastbillService)

	log.Println("Create Invoice")
	inv, err := invoiceClient.Create(&invoice.Request{
		CustomerID: customerID,
		Items: []invoice.Item{
			{
				Description: "My awesome it" +
					"" +
					"em",
				Quantity:   5,
				UnitPrice:  1.5,
				VatPercent: 19,
			},
		}})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update Invoice")
	_, err = invoiceClient.Update(&invoice.Request{
		InvoiceID: strconv.Itoa(inv.InvoiceID),
		IntroText: "Some awesome invoice text",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Invoice")
	invData, err := invoiceClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}
	invJSON, _ := json.Marshal(invData)
	log.Println(string(invJSON))

	/*
		log.Println("Complete Invoice")
		_, err = invoiceClient.Complete(strconv.Itoa(inv.InvoiceID))
		if err != nil {
			log.Fatal(err)
		}

		log.Println("SendByMail Invoice")
		_, err = invoiceClient.SendByEmail(&invoice.SendByMailRequest{
			InvoiceID: strconv.Itoa(inv.InvoiceID),

			Recipient: invoice.SendByMailRecipients{
				To: "dome.malsch@web.de",
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("SetPaid Invoice")
		_, err = invoiceClient.SetPaid(&invoice.SetPaidRequest{
			InvoiceID:     strconv.Itoa(inv.InvoiceID),
			PaidDate:      time.Now(),
			PaymentMethod: "paypal",
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Cancel Invoice")
		_, err = invoiceClient.Cancel(strconv.Itoa(inv.InvoiceID))
		if err != nil {
			log.Fatal(err)
		}
	*/
	return inv.InvoiceID
}

func handleRecurring(fastbillService service.Service, customerID string) {
	recClient := recurring.NewRecurringClient(fastbillService)

	log.Println("Create Invoice")
	inv, err := recClient.Create(&recurring.Request{
		CustomerID: customerID,
		StartDate:  time.Now().Format("2006-01-02"),
		OutputType: "outgoing",
		Frequency:  "monthly",
		Items: []recurring.Item{
			{
				Description: "My awesome it" +
					"" +
					"em",
				Quantity:   5,
				UnitPrice:  1.5,
				VatPercent: 19,
			},
		}})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update Invoice")
	_, err = recClient.Update(&recurring.Request{
		InvoiceID: strconv.Itoa(inv.InvoiceID),
		IntroText: "Some awesome invoice text",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Invoice")
	invData, err := recClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}
	invJSON, _ := json.Marshal(invData)
	log.Println(string(invJSON))
}

func handleTemplates(fastbillService service.Service) {
	estimateClient := template.NewTemplateClient(fastbillService)

	log.Println("Get Template")
	invData, err := estimateClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

	invJSON, _ := json.Marshal(invData)
	log.Println(string(invJSON))
}

func handleEstimates(fastbillService service.Service, customerID string) {
	estimateClient := estimate.NewEstimateClient(fastbillService)

	log.Println("Create Estimate")
	inv, err := estimateClient.Create(&estimate.Request{
		CustomerID: customerID,
		Items: []estimate.Item{
			{
				Description: "My awesome item",
				Quantity:    "5",
				UnitPrice:   "1.5",
				VatPercent:  "19",
			},
		}})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Estimate")
	invData, err := estimateClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	invJSON, _ := json.Marshal(invData)
	log.Println(string(invJSON))

	log.Println("Create Invoice")
	_, err = estimateClient.CreateInvoice(strconv.Itoa(inv.EstimateID))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("SendByMail Invoice")
	_, err = estimateClient.SendByEmail(&estimate.SendByMailRequest{
		EstimateID: strconv.Itoa(inv.EstimateID),
		Recipient: estimate.SendByMailRecipients{
			To: "dome.malsch@web.de",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func handleDocument(fastbillService service.Service) {
	documentClient := document.NewDocumentClient(fastbillService)
	log.Println("Get Documents")

	docs, err := documentClient.Get(&parameter.Parameter{
		Limit: 10,
	}, &document.Filter{FolderID: "1"})

	if err != nil {
		log.Fatal(err)
	}

	invJSON, _ := json.Marshal(docs)
	log.Println(string(invJSON))

	log.Println("create document")

	resp, _ := http.Get("https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf")
	defer resp.Body.Close()
	d, err := documentClient.Create(&document.Document{}, resp.Body, "myFile.pdf")

	if err != nil {
		log.Fatal(err)
	}

	x, _ := json.Marshal(d)
	log.Println(string(x))
}

func handleArticle(fastbillService service.Service) {
	contactClient := article.NewArticleClient(fastbillService)

	log.Println("Create Article")
	c, err := contactClient.Create(&article.Article{ArticleNumber: "10", Title: "artikel", UnitPrice: "10.50", Unit: "Stück"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update Article")
	_, err = contactClient.Update(&article.Article{ArticleID: strconv.Itoa(c.ArticleID), Description: "meine beschreibung"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Article")
	list, err := contactClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	customerJSON, _ := json.Marshal(list)
	log.Println(string(customerJSON))
	log.Println("Delete Article")
	deleted, err := contactClient.Delete(strconv.Itoa(c.ArticleID))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(deleted)
}
