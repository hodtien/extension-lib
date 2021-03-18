package vdp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hodtien/extension-library/transport"
)

// HttpGetAllUsersByOneCondition - HttpGetAllUsersByOneCondition
func HttpGetAllUsersByOneCondition(ctx context.Context, apiKey, fieldKey, fieldValue string) map[string]interface{} {
	url := "/api/permission/v1/authentication/private/user/list"
	url = transport.Domain + url + "?field_key=" + fieldKey + "&field_value=" + fieldValue

	apiKey = "Bearer " + apiKey

	resp, err := transport.MakeHTTPGetRequest(ctx, url, []string{"Authorization"}, []string{apiKey})
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Get Users Failed: " + err.Error()}
	}

	if strings.Contains(string(resp), "Authentication Token is Invalid") {
		return map[string]interface{}{"code": "1", "message": "Authentication Token is Invalid"}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "-1", "message": "Get Users Failed: " + err.Error()}
	}
	return body
}

// HttpGetUserProfileByUserID - HttpGetUserProfileByUserID
func HttpGetUserProfileByUserID(ctx context.Context, apiKey, userID string) map[string]interface{} {
	urlGetUserProfile := "/api/permission/v1/authentication/user/profile/"

	url := transport.Domain + urlGetUserProfile + userID

	apiKey = "Bearer " + apiKey
	resp, err := transport.MakeHTTPGetRequest(ctx, url, []string{"Authorization"}, []string{apiKey})
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get User Profile Failed: " + err.Error()}
	}

	if string(resp) == "APIKey is Invalid" {
		return map[string]interface{}{"code": "1", "message": "Get User Profile Failed: APIKey is Invalid"}
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(resp, &body)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get User Profile Failed: " + err.Error()}
	}

	if body["code"] == "-1" {
		if body["message"] == "Authentication Token is Invalid" {
			body["code"] = "1"
		} else if body["message"] == "Permission Denied" {
			body["code"] = "3"
		} else {
			return map[string]interface{}{"code": "4", "message": "User doesn't exist"}
		}
	}
	return body
}

// HttpUpdateUserProfileByUserID - HttpUpdateUserProfileByUserID
func HttpUpdateUserProfileByUserID(ctx context.Context, apiKey, userID string, bodyUpdate interface{}) map[string]interface{} {
	url := transport.Domain + "/api/permission/v1/authentication/user/update/" + userID

	apiKey = "Bearer " + apiKey

	bodyBytes, _ := json.Marshal(bodyUpdate)

	resp, err := transport.MakeHTTPPatchRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Update User Profile Failed: " + err.Error()}
	}

	bodyResp := make(map[string]interface{})
	err = json.Unmarshal(resp, &bodyResp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Update User Profile Failed: " + err.Error()}
	}
	return bodyResp
}

// HttpGetAllUserByOneCondition - HttpGetAllUserByOneCondition
func HttpGetAllUserByOneCondition(ctx context.Context, apiKey, role string) map[string]interface{} {
	// Initial
	url := transport.Domain + "/api/permission/v1/authentication/private/user/list?field_key=role&field_value=" + role
	apiKey = "Bearer " + apiKey

	// Exc
	resp, err := transport.MakeHTTPGetRequest(ctx, url, []string{"Authorization"}, []string{apiKey})
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}
	}

	bodyResp := make(map[string]interface{})
	err = json.Unmarshal(resp, &bodyResp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}
	}
	return bodyResp
}

// HttpGetAllUser - HttpGetAllUser
func HttpGetAllUser(ctx context.Context, apiKey string) map[string]interface{} {
	// Initial
	url := transport.Domain + "/api/permission/v1/authentication/user/all"
	apiKey = "Bearer " + apiKey

	// Exc
	resp, err := transport.MakeHTTPGetRequest(ctx, url, []string{"Authorization"}, []string{apiKey})
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"code": "10", "message": "Get All User Failed: " + err.Error()}
	}

	bodyResp := make(map[string]interface{})
	err = json.Unmarshal(resp, &bodyResp)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"code": "10", "message": "Get All User Failed: " + err.Error()}
	}

	return bodyResp
}

// HttpGetAllUserByIDList - HttpGetAllUserByIDList
func HttpGetAllUserByIDList(ctx context.Context, apiKey string, bodyBytes []byte) map[string]interface{} {
	// Initial
	url := transport.Domain + "/api/permission/v1/authentication/user/retrieve_many"
	apiKey = "Bearer " + apiKey

	// Exc
	resp, err := transport.MakeHTTPPostRequest(ctx, url, []string{"Authorization"}, []string{apiKey}, bodyBytes)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}
	}

	bodyResp := make(map[string]interface{})
	err = json.Unmarshal(resp, &bodyResp)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}

	}
	return bodyResp
}
