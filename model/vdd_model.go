package model

type (
	// NATSRequest - NATSRequest
	NATSRequest struct {
		ConnectionID string      `json:"conn_id"`
		RequestID    string      `json:"request_id" validate:"required"`
		UserID       string      `json:"user_id,omitempty"`
		DeviceID     string      `json:"device_id,omitempty"`
		APIKey       string      `json:"api_key" validate:"required"`
		BucketID     string      `json:"bucket_id" validate:"required"`
		RecordID     string      `json:"record_id,omitempty"`
		QueryPath    string      `json:"query_path,omitempty"`
		PagingStart  int64       `json:"start,omitempty"`
		PagingStop   int64       `json:"stop,omitempty"`
		Pagging      int         `json:"page,omitempty"`
		Limit        int         `json:"limit,omitempty"`
		FieldSort    string      `json:"field_sort,omitempty"`
		OrderSort    string      `json:"order_sort,omitempty"`
		ListRecordID []string    `json:"list_record_id,omitempty"`
		FieldUpdate  string      `json:"field_update,omitempty"`
		ValueUpdate  string      `json:"value_update,omitempty"`
		Body         interface{} `json:"body,omitempty"`
	}
)
