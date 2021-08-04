# fastbill-go-sdk
[![Checks](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml/badge.svg)](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml)

golang sdk for the api of https://www.fastbill.com/

> :warning: **Work in progress**: This SDK is still in an very early state. It is not recommended for production use.

## Usage
Put the sdk package on your import statement

```golang
import "github.com/malsch-solutions/fastbill-go-sdk"
```
Initialize your fastbill session
```golang
fastbillSession := session.NewSession(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))
```
Create the desired client for example for customers
```golang
	customerClient := customer.NewCustomerClient(fastbillSession)
```
Call the method of your choice
```
	customers, err := customerClient.Get(&parameter.Parameter{
		Limit:  10,
		Offset: 0,
	}, nil)
```
