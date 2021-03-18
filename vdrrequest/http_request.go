package vdrrequest

import (
	"context"
	"encoding/json"

	"github.com/hodtien/extension-lib/global"
	"github.com/hodtien/extension-lib/httpclient"
)

// CreateRecord - CreateRecord
func CreateRecord(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := global.Domain + "/api/report/v1/secure/record/create/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := httpclient.MakeHTTPPutRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Create VDR Record Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Create VDR Record Failed: " + err.Error()}
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

// DataIncrement - DataIncrement
func DataIncrement(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := global.Domain + "/api/report/v1/secure/record/data/increment/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := httpclient.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Increment Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Increment Failed: " + err.Error()}
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

// DataDecrement - DataDecrement
func DataDecrement(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := global.Domain + "/api/report/v1/secure/record/data/decrement/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := httpclient.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Decrement Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Data Decrement Failed: " + err.Error()}
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

// RetrieveStatistics - RetrieveStatistics
func RetrieveStatistics(ctx context.Context, apiKey, bucketID string, bodyRequest []byte) map[string]interface{} {
	url := global.Domain + "/api/report/v1/secure/statistics/retrieve/" + bucketID
	apiKey = "Bearer " + apiKey

	resp, err := httpclient.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyRequest)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Retrieve Statistics Failed: " + err.Error()}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Retrieve Statistics Failed: " + err.Error()}
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
