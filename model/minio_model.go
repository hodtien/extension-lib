package model

type (
	// NATSRequestFileStorage - NATSRequestFileStorage
	NATSRequestFileStorage struct {
		ConnectionID string         `json:"conn_id"`
		RequestID    string         `json:"request_id" validate:"required"`
		ApiKey       string         `json:"api_key,omitempty"`
		UserID       string         `json:"user_id,omitempty"`
		DeviceID     string         `json:"device_id,omitempty"`
		BucketID     string         `json:"bucket_id" validate:"required"`
		RecordIDs    []interface{}  `json:"record_ids,omitempty"`
		Folder       string         `json:"folder,omitempty"`
		Files        []*FileStorage `json:"body,omitempty"`
	}

	// FileStorage - FileStorage
	FileStorage struct {
		ContentType string `json:"content_type" validate:"required"`
		Source      []byte `json:"src" validate:"required"`
		Size        int64  `json:"size" validate:"required"`
		Path        string `json:"path" validate:"required"`
	}
)
