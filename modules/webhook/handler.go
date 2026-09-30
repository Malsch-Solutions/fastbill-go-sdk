package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
)

// MaxEventSize is the largest notification body ParseEvent reads.
const MaxEventSize = 1 << 20

// userAgent is the User-Agent FastBill sends notifications with.
const userAgent = "FastBill"

// ParseEvent checks that r is a FastBill notification (POST, JSON, FastBill
// user agent, at most MaxEventSize bytes) and decodes it. The caller should
// answer with 400 Bad Request if it returns an error.
//
// FastBill documents no signature for notifications, so these checks do
// not prove the sender; use a secret, hard to guess endpoint URL.
func ParseEvent(r *http.Request) (Event, error) {
	var event Event
	if r.Method != http.MethodPost {
		return event, fmt.Errorf("webhook: method %s, want POST", r.Method)
	}
	contentType := r.Header.Get("Content-Type")
	if mediaType, _, err := mime.ParseMediaType(contentType); err != nil || mediaType != "application/json" {
		return event, fmt.Errorf("webhook: Content-Type %q, want application/json", contentType)
	}
	if ua := r.UserAgent(); ua != userAgent {
		return event, fmt.Errorf("webhook: User-Agent %q, want %s", ua, userAgent)
	}
	if r.Body == nil {
		return event, fmt.Errorf("webhook: empty body")
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, MaxEventSize+1))
	if err != nil {
		return event, fmt.Errorf("webhook: read body: %w", err)
	}
	if len(body) > MaxEventSize {
		return event, fmt.Errorf("webhook: body larger than %d bytes", MaxEventSize)
	}
	if err := json.Unmarshal(body, &event); err != nil {
		return event, fmt.Errorf("webhook: decode event: %w", err)
	}
	return event, nil
}
