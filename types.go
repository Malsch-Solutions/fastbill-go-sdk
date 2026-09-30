package fastbill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// ID is an identifier. FastBill sends IDs as strings in listings and as
// numbers in some create responses; ID accepts both.
type ID string

// UnmarshalJSON accepts a string, a number or null.
func (id *ID) UnmarshalJSON(b []byte) error {
	s, err := scalar(b)
	if err != nil {
		return fmt.Errorf("fastbill: ID: %w", err)
	}
	*id = ID(s)
	return nil
}

func (id ID) String() string { return string(id) }

// Number is a decimal value (amount, quantity, percentage). FastBill sends
// numbers as JSON numbers or numeric strings. Number keeps the digits as sent,
// so amounts are never rounded through float64 on the way.
type Number string

// NewNumber formats f with as many digits as needed.
func NewNumber(f float64) Number { return Number(strconv.FormatFloat(f, 'f', -1, 64)) }

// NewInt formats i.
func NewInt(i int) Number { return Number(strconv.Itoa(i)) }

// UnmarshalJSON accepts a number, a string or null.
func (n *Number) UnmarshalJSON(b []byte) error {
	s, err := scalar(b)
	if err != nil {
		return fmt.Errorf("fastbill: Number: %w", err)
	}
	*n = Number(s)
	return nil
}

// MarshalJSON writes a JSON number, or a string if the value is not numeric.
func (n Number) MarshalJSON() ([]byte, error) {
	if n == "" {
		return []byte("null"), nil
	}
	if _, err := strconv.ParseFloat(string(n), 64); err == nil && json.Valid([]byte(n)) {
		return []byte(n), nil
	}
	return json.Marshal(string(n))
}

// Float64 parses the value. An empty Number is 0.
func (n Number) Float64() (float64, error) {
	if n == "" {
		return 0, nil
	}
	return strconv.ParseFloat(string(n), 64)
}

func (n Number) String() string { return string(n) }

// scalar returns the text of a JSON string, number, boolean or null.
func scalar(b []byte) (string, error) {
	b = bytes.TrimSpace(b)
	switch {
	case len(b) == 0, string(b) == "null":
		return "", nil
	case b[0] == '"':
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return "", err
		}
		return strings.TrimSpace(s), nil
	case b[0] == '{' || b[0] == '[':
		return "", fmt.Errorf("got %s, want a scalar", b[:1])
	default:
		// number, true or false
		return string(b), nil
	}
}

// List is a JSON list as FastBill encodes it: an array, or an object keyed
// by index (PHP), and {} or null when empty.
type List[T any] []T

// UnmarshalJSON accepts an array, an object or null.
func (l *List[T]) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	switch {
	case len(b) == 0, string(b) == "null":
		*l = nil
		return nil
	case b[0] == '[':
		var items []T
		if err := json.Unmarshal(b, &items); err != nil {
			return err
		}
		*l = items
		return nil
	case b[0] == '{':
		var byKey map[string]T
		if err := json.Unmarshal(b, &byKey); err != nil {
			return err
		}
		keys := make([]string, 0, len(byKey))
		for k := range byKey {
			keys = append(keys, k)
		}
		slices.SortFunc(keys, compareKeys)
		items := make([]T, 0, len(keys))
		for _, k := range keys {
			items = append(items, byKey[k])
		}
		*l = items
		return nil
	case b[0] == '"':
		// A single message instead of a list (seen with ERRORS).
		var item T
		if err := json.Unmarshal(b, &item); err != nil {
			return err
		}
		*l = List[T]{item}
		return nil
	}
	return fmt.Errorf("fastbill: list: unexpected %q", b[:1])
}

// compareKeys orders numeric keys by value, the rest as text.
func compareKeys(a, b string) int {
	na, errA := strconv.Atoi(a)
	nb, errB := strconv.Atoi(b)
	if errA == nil && errB == nil {
		return na - nb
	}
	return strings.Compare(a, b)
}

// Flag is a yes/no field (IS_CANCELED, IS_GROSS, …). FastBill sends "1"/"0",
// and Flag also accepts 1/0 and true/false.
type Flag string

// Flag values.
const (
	Yes Flag = "1"
	No  Flag = "0"
)

// NewFlag returns Yes or No.
func NewFlag(b bool) Flag {
	if b {
		return Yes
	}
	return No
}

// UnmarshalJSON accepts a string, a number, a boolean or null.
func (f *Flag) UnmarshalJSON(b []byte) error {
	s, err := scalar(b)
	if err != nil {
		return fmt.Errorf("fastbill: Flag: %w", err)
	}
	switch s {
	case "true":
		s = string(Yes)
	case "false":
		s = string(No)
	}
	*f = Flag(s)
	return nil
}

// Bool reports whether the flag is set.
func (f Flag) Bool() bool { return f == Yes || strings.EqualFold(string(f), "true") }

// StatusResponse is the answer of services that only report success.
type StatusResponse struct {
	Status string `json:"STATUS"`
}

// Err returns an error unless FastBill reported success.
func (s StatusResponse) Err(service string) error {
	if strings.EqualFold(s.Status, "success") {
		return nil
	}
	return &APIError{Service: service, Messages: []string{"status " + strconv.Quote(s.Status)}}
}
