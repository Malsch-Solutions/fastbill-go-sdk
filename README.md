# fastbill-go-sdk

[![Checks](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml/badge.svg)](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/check.yml)
[![CodeQL](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Malsch-Solutions/fastbill-go-sdk/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/Malsch-Solutions/fastbill-go-sdk/branch/main/graph/badge.svg?token=NYMO09X0BU)](https://codecov.io/gh/Malsch-Solutions/fastbill-go-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/malsch-solutions/fastbill-go-sdk/v2.svg)](https://pkg.go.dev/github.com/malsch-solutions/fastbill-go-sdk/v2)

Go client for the [FastBill API](https://apidocs.fastbill.com/fastbill/en/). No dependencies outside the standard library.

## Requirements

Go 1.25 or newer.

```bash
go get github.com/malsch-solutions/fastbill-go-sdk/v2
```

## Usage

```go
import (
	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/invoice"
)

client := fastbill.NewClient(os.Getenv("FASTBILL_EMAIL"), os.Getenv("FASTBILL_API_KEY"))
invoices := invoice.NewClient(client)

drafts, err := invoices.Get(ctx, fastbill.Page{Limit: 100}, &invoice.Filter{Type: invoice.TypeDraft})
```

Every call takes a `context.Context`. The client has a 60-second timeout. Pass `fastbill.WithHTTPClient` to change it.

**All pages.** FastBill returns at most 100 records per call. `fastbill.All` pages through a listing:

```go
all, err := fastbill.All(ctx, fastbill.MaxLimit, func(ctx context.Context, p fastbill.Page) ([]invoice.Invoice, error) {
	return invoices.Get(ctx, p, nil)
})
```

**Raw records.** `fastbill.Records` returns each record as FastBill sent it, with every field, e.g. to archive it:

```go
records, err := fastbill.Records(ctx, client, "invoice.get", "INVOICES", nil, fastbill.Page{Limit: 100})
```

**Documents.** `client.Download(ctx, inv.DocumentURL)` fetches an invoice PDF or a receipt file.

**Errors.** FastBill's `ERRORS` come back as `*fastbill.APIError`, and HTTP failures as `*fastbill.StatusError`. Services that only report a status return an error unless the status is `success`.

**Rate limits.** FastBill limits API calls per hour by plan. The client does not throttle, so space large jobs out yourself.

### Types

FastBill sends the same field as a JSON string in one response and as a number in another. The structs use:

| Type | For | Accepts |
|---|---|---|
| `fastbill.ID` | `*_ID` fields | string, number |
| `fastbill.Number` | amounts, quantities, percentages | number, numeric string; keeps the digits as sent, so money is never rounded through `float64`. `Float64()` parses it; `NewNumber`/`NewInt` create one |
| `fastbill.Flag` | 0/1 fields (`IS_CANCELED`, …) | `"1"`/`"0"`, `1`/`0`, `true`/`false`; `Bool()` reads it |
| `fastbill.List[T]` | lists in responses | array, PHP object keyed by index, `{}` or `null` when empty |

Dates are strings in FastBill's format (`YYYY-MM-DD`).

## API coverage

| Module | Services |
|---|---|
| `article` | Products |
| `contact` | Contacts |
| `customer` | Customers |
| `document` | Documents (inbox) |
| `estimate` | Estimates |
| `expense` | Expenses (receipts) |
| `invoice` | Invoices, drafts, credit notes |
| `item` | Invoice items |
| `project` | Projects |
| `recurring` | Recurring invoices |
| `revenue` | Revenues |
| `template` | Templates |
| `time` | Work times |
| `webhook` | Webhooks, and `webhook.ParseEvent` for incoming calls |

## Upgrading from v1

v2 is a new module path, `github.com/malsch-solutions/fastbill-go-sdk/v2`. v1 keeps working where it is pinned.

- `service.NewService(email, key)` → `fastbill.NewClient(email, key, opts...)`. `request`, `response`, `parameter` and `service` are gone. Use `fastbill.Request`, `fastbill.Page` and `fastbill.Requester`.
- `invoice.NewInvoiceClient(s)` → `invoice.NewClient(c)`, and the same for every module.
- Every method takes `ctx` first. `Get(&parameter.Parameter{Limit, Offset}, filter)` → `Get(ctx, fastbill.Page{Limit, Offset}, filter)`.
- Methods that returned `(bool, error)` (Delete, Cancel, SendByEmail, …) return `error`.
- Fields are typed as in the table above. v1's `int` amounts and quantities (e.g. expense `SUB_TOTAL`, item `QUANTITY`) cut decimals off. They are `fastbill.Number` now.
- Filter dates are `YYYY-MM-DD` strings. v1 sent `time.Time` as RFC 3339, which FastBill does not understand.
- Webhooks: `NewWebhookRequestHandler(req).ValidateAndGetData()` → `webhook.ParseEvent(r)`.
- Fields from FastBill's current docs were added, e.g. invoice `SUBTYPE`, `STATE`, `PAYMENTS` and `DETAILS_URL`.
