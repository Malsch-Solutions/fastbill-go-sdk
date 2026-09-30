package fastbill

import (
	"context"
	"encoding/json"
	"fmt"
)

// Records lists a `*.get` service without decoding the records: each one is
// returned as FastBill sent it, with every field. key is the list in the
// response, e.g. "INVOICES" for invoice.get. Use it to archive data, or for
// fields the typed structs do not have.
func Records(ctx context.Context, r Requester, service, key string, filter any, page Page) ([]json.RawMessage, error) {
	req := Request{Service: service, Limit: page.Limit, Offset: page.Offset, Filter: filter}
	var response map[string]List[json.RawMessage]
	if err := r.Do(ctx, req, &response); err != nil {
		return nil, err
	}
	return response[key], nil
}

// All calls get with growing offsets until a page comes back shorter than
// pageSize (at most MaxLimit), and returns every record.
//
//	invoices, err := fastbill.All(ctx, fastbill.MaxLimit, func(ctx context.Context, p fastbill.Page) ([]invoice.Invoice, error) {
//		return invoices.Get(ctx, p, nil)
//	})
func All[T any](ctx context.Context, pageSize int, get func(context.Context, Page) ([]T, error)) ([]T, error) {
	if pageSize <= 0 || pageSize > MaxLimit {
		return nil, fmt.Errorf("fastbill: page size %d, want 1 to %d", pageSize, MaxLimit)
	}
	var all []T
	for offset := 0; ; offset += pageSize {
		page, err := get(ctx, Page{Limit: pageSize, Offset: offset})
		if err != nil {
			return all, err
		}
		all = append(all, page...)
		if len(page) < pageSize {
			return all, nil
		}
	}
}
