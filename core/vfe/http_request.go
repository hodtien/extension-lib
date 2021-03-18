package vfe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/hodtien/extension-lib/transport"
)

func HTTPUploadFiles(ctx context.Context, apiKey, bucketID string, multiPartWriter *multipart.Writer, byteBuff bytes.Buffer) map[string]interface{} {
	apiKey = "Bearer " + apiKey
	url := transport.Domain + "/api/vnpt-filestorage-engine/v1/minio/upload/" + bucketID

	bodyBytes, err := transport.MakeHTTPPostFormRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, multiPartWriter, byteBuff)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Upload File Failed" + fmt.Sprintf("%v", err)}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Upload File Failed" + fmt.Sprintf("%v", err)}
	}

	return resp
}
