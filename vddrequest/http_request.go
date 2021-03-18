package vddrequest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/hodtien/extension-lib/global"
	"github.com/hodtien/extension-lib/httpclient"
)

// MatchData - MatchData
func MatchData(ctx context.Context, apiKey, bucketID string, bodyBytes []byte) map[string]interface{} {
	apiKey = "Bearer " + apiKey

	url := global.Domain + "/api/core/v1/data/all_in_bucket/match/" + bucketID

	resp, err := httpclient.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}
	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}

// UploadFileMinio - UploadFileMinio
func UploadFileMinio(ctx context.Context, apiKey, bucketID string, multiPartWriter *multipart.Writer, bodyBytes bytes.Buffer) interface{} {
	// apiKey := viper.GetString("api_key.vdp")
	apiKey = "Bearer " + apiKey
	url := global.Domain + "/api/core/v1/minio/upload/" + bucketID

	resp, err := httpclient.MakeHTTPPostFormRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, multiPartWriter, bodyBytes)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}

	var body []interface{}
	err = json.Unmarshal(resp, &body)
	if err != nil {
		fmt.Println(string(resp))
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}

	return body
}

// MatchDataWithPaggingAndLogic - MatchDataWithPaggingAndLogic
func MatchDataWithPaggingAndLogic(ctx context.Context, apiKey, bucketID, page, limit string, bodyBytes []byte, getTotal bool) map[string]interface{} {
	apiKey = "Bearer " + apiKey

	url := global.Domain + "/api/core/v1/data/all_in_bucket/match_with_pagging/logic/" + bucketID
	if page != "" && limit != "" {
		url = url + "?page=" + page
		url = url + "&limit=" + limit
	}

	resp, err := httpclient.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Query Data Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		}
	}

	return body
}
