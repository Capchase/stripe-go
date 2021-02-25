package stripe

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	APIResource
	Account         *string `json:"account,omitempty"`
	Created         *int64 `json:"created,omitempty"`
	Data            *EventData    `json:"data,omitempty"`
	ID              *string `json:"id,omitempty"`
	Livemode        *bool `json:"livemode,omitempty"`
	PendingWebhooks *int64 `json:"pending_webhooks,omitempty"`
	Request         *EventRequest `json:"request,omitempty"`
	Type            *string `json:"type,omitempty"`
}

// EventRequest contains information on a request that created an event.
type EventRequest struct {
	// ID is the request ID of the request that created an event, if the event
	// was created by a request.
	ID *string `json:"id,omitempty"`

	// IdempotencyKey is the idempotency key of the request that created an
	// event, if the event was created by a request and if an idempotency key
	// was specified for that request.
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
}

// EventData is the unmarshalled object as a map.
type EventData struct {
	// Object is a raw mapping of the API resource contained in the event.
	// Although marked with json:"-", it's still populated independently by
	// a custom UnmarshalJSON implementation.
	Object             map[string]interface{} `json:"-,omitempty"`
	PreviousAttributes map[string]interface{} `json:"previous_attributes,omitempty"`
	Raw                json.RawMessage        `json:"object,omitempty"`
}

// EventParams is the set of parameters that can be used when retrieving events.
// For more details see https://stripe.com/docs/api#retrieve_events.
type EventParams struct {
	Params `form:"*"`
}

// EventList is a list of events as retrieved from a list endpoint.
type EventList struct {
	APIResource
	ListMeta
	Data []*Event `json:"data,omitempty"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams      `form:"*"`
	Created         *int64            `form:"created"`
	CreatedRange    *RangeQueryParams `form:"created"`
	DeliverySuccess *bool             `form:"delivery_success"`
	Type            *string           `form:"type"`
	Types           []*string         `form:"types"`
}

// GetObjectValue returns the value from the e.Data.Object bag based on the keys hierarchy.
func (e *Event) GetObjectValue(keys ...string) string {
	return getValue(e.Data.Object, keys)
}

// GetPreviousValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPreviousValue(keys ...string) string {
	return getValue(e.Data.PreviousAttributes, keys)
}

// UnmarshalJSON handles deserialization of the EventData.
// This custom unmarshaling exists so that we can keep both the map and raw data.
func (e *EventData) UnmarshalJSON(data []byte) error {
	type eventdata EventData
	var ee eventdata
	err := json.Unmarshal(data, &ee)
	if err != nil {
		return err
	}

	*e = EventData(ee)
	return json.Unmarshal(e.Raw, &e.Object)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		key := keys[i]

		sliceNode, ok := node.([]interface{})
		if ok {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				panic(fmt.Sprintf(
					"Cannot access nested slice element with non-integer key: %s",
					key))
			}
			node = sliceNode[intKey]
			continue
		}

		mapNode, ok := node.(map[string]interface{})
		if ok {
			node = mapNode[key]
			continue
		}

		panic(fmt.Sprintf(
			"Cannot descend into non-map non-slice object with key: %s", key))
	}

	if node == nil {
		return ""
	}

	return fmt.Sprintf("%v", node)
}
