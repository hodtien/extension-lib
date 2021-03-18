package vdd

import (
	"context"
	"encoding/json"

	"github.com/hodtien/extension-lib/transport"
)

// MatchData - MatchData
func MatchData(ctx context.Context, apiKey, bucketID string, bodyBytes []byte) map[string]interface{} {
	apiKey = "Bearer " + apiKey

	url := transport.Domain + "/api/core/v1/data/all_in_bucket/match/" + bucketID

	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
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

// MatchDataWithPaggingAndLogic - MatchDataWithPaggingAndLogic
func MatchDataWithPaggingAndLogic(ctx context.Context, apiKey, bucketID, page, limit string, bodyBytes []byte, getTotal bool) map[string]interface{} {
	apiKey = "Bearer " + apiKey

	url := transport.Domain + "/api/core/v1/data/all_in_bucket/match_with_pagging/logic/" + bucketID
	if page != "" && limit != "" {
		url = url + "?page=" + page
		url = url + "&limit=" + limit
	}

	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
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
