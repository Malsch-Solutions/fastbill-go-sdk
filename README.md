# fastbill-go-sdk
[![Checks](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml/badge.svg)](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml)
[![codecov](https://codecov.io/gh/Malsch-Solutions/fastbill-go-sdk/branch/main/graph/badge.svg?token=NYMO09X0BU)](https://codecov.io/gh/Malsch-Solutions/fastbill-go-sdk)

golang sdk for the api of https://www.fastbill.com/

> :warning: **Work in progress**: This SDK is still in an very early state. It is not recommended for production use.

## Usage
Put the sdk package on your import statement

```golang
import "github.com/malsch-solutions/fastbill-go-sdk"
```
Initialize your fastbill service
```golang
fastbillService := service.NewService(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))
```
Create the desired client for example for customers
```golang
customerClient := customer.NewCustomerClient(fastbillService)
```
Call the method of your choice
```
customers, err := customerClient.Get(&parameter.Parameter{
    Limit:  10,
    Offset: 0,
}, nil)
```

## API Coverage

- [ ] Customers
- [ ] Estimates
- [ ] Invoices
- [ ] Items
- [ ] Recurring invoices
- [ ] Revenues
- [ ] Expenses
- [ ] Products
- [ ] Projects
- [ ] Work times
- [ ] Documents
- [ ] Templates
- [ ] Webhooks
