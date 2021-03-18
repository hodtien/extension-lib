package vdd

import (
	"encoding/json"
	"time"
	
	"github.com/hodtien/extension-lib/transport"
	"github.com/hodtien/extension-lib/model"

	"github.com/google/uuid"
)

// RetrieveData - RetrieveData
func RetrieveData(apiKey, bucketID, recordID, queryPath string) map[string]interface{} {
	subj := "vdd_request.RetrieveData"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.QueryPath = queryPath

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveData FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveData FAILED: " + err.Error()}
	}

	return resp
}

// RetrieveAllDataInBucket - RetrieveAllDataInBucket
func RetrieveAllDataInBucket(apiKey, bucketID, queryPath string) map[string]interface{} {
	subj := "vdd_request.RetrieveAllDataInBucket"
	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.QueryPath = queryPath

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveAllDataInBucket FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveAllDataInBucket FAILED: " + err.Error()}
	}

	return resp
}

// RetrieveAllData - RetrieveAllData
func RetrieveAllData(userID string, start, stop int64) map[string]interface{} {
	subj := "vdd_request.RetrieveAllData"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.UserID = userID
	nReq.PagingStart = start
	nReq.PagingStop = stop

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveAllData FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveAllData FAILED: " + err.Error()}
	}

	if resp["code"] == "1" {
		resp["code"] = "4"
		resp["message"] = "Data Not Found"
	}

	return resp
}

// RetrieveManyDataByListID - RetrieveManyDataByListID
func RetrieveManyDataByListID(apiKey, bucketID string, listRecordID []string) map[string]interface{} {
	subj := "vdd_request.RetrieveManyDataByListID"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.ListRecordID = listRecordID

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveManyDataByListID FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "RetrieveManyDataByListID FAILED: " + err.Error()}
	}

	if resp["code"] == "1" {
		resp["code"] = "4"
		resp["message"] = "Data Not Found"
	}

	return resp
}

// SaveRecord - SaveRecord
func SaveRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.SaveRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "SaveRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "SaveRecord FAILED: " + err.Error()}
	}

	return resp
}

// SaveRecordValidator - SaveRecordValidator
func SaveRecordValidator(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.SaveRecordValidator"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "SaveRecordValidator FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "SaveRecordValidator FAILED: " + err.Error()}
	}

	return resp
}

// UpdateRecordValidator - UpdateRecordValidator
func UpdateRecordValidator(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.UpdateRecordValidator"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRecordValidator FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRecordValidator FAILED: " + err.Error()}
	}

	return resp
}

// DeleteRecord - DeleteRecord
func DeleteRecord(apiKey, bucketID, recordID string) map[string]interface{} {
	subj := "vdd_request.DeleteRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteRecord FAILED: " + err.Error()}
	}

	return resp
}

// DeleteManyRecord - DeleteManyRecord
func DeleteManyRecord(apiKey, bucketID string, listRecordID []string) map[string]interface{} {
	subj := "vdd_request.DeleteManyRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.ListRecordID = listRecordID

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteManyRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "DeleteManyRecord FAILED: " + err.Error()}
	}

	if resp["code"] == "1" {
		resp["code"] = "4"
		resp["message"] = "Data Not Found"
	}

	return resp
}

// UpdateAFieldRecord - UpdateAFieldRecord
func UpdateAFieldRecord(apiKey, bucketID, recordID, queryPath string, body interface{}) map[string]interface{} {
	subj := "vdd_request.UpdateAFieldRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.QueryPath = queryPath
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateAFieldRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateAFieldRecord FAILED: " + err.Error()}
	}

	return resp
}

// UpdateRecord - UpdateRecord
func UpdateRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.UpdateRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateRecord FAILED: " + err.Error()}
	}

	return resp
}

// UpdateSomeFieldRecord - UpdateSomeFieldRecord
func UpdateSomeFieldRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.UpdateSomeFieldRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.RecordID = recordID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateSomeFieldRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateSomeFieldRecord FAILED: " + err.Error()}
	}

	return resp
}

// UpdateManyRecords - UpdateManyRecords
func UpdateManyRecords(apiKey, bucketID string, listRecordID []string, fieldUpdate, valueUpdate string) map[string]interface{} {
	subj := "vdd_request.UpdateManyRecord"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.ListRecordID = listRecordID
	nReq.FieldUpdate = fieldUpdate
	nReq.ValueUpdate = valueUpdate

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateManyRecord FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "UpdateManyRecord FAILED: " + err.Error()}
	}

	return resp
}

// NatsMatchData - NatsMatchData
func NatsMatchData(apiKey, bucketID, fieldSort, orderSort string, body map[string]interface{}) map[string]interface{} {
	subj := "vdd_request.MatchData"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.Body = body
	nReq.FieldSort = fieldSort
	nReq.OrderSort = orderSort

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchData FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchData FAILED: " + err.Error()}
	}

	return resp
}

// NatsMatchDataWithPagging - NatsMatchDataWithPagging
func NatsMatchDataWithPagging(apiKey, bucketID, fieldSort, orderSort string, page, limit int, body map[string]interface{}) map[string]interface{} {
	subj := "vdd_request.MatchDataWithPagging"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.Body = body
	nReq.Pagging = page
	nReq.Limit = limit
	nReq.FieldSort = fieldSort
	nReq.OrderSort = orderSort

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchDataWithPagging FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchDataWithPagging FAILED: " + err.Error()}
	}

	return resp
}

// NatsInnerJoin - NatsInnerJoin
func NatsInnerJoin(apiKey, bucketID string, page, limit int, body interface{}) map[string]interface{} {
	subj := "vdd_request.MatchDataInnerJoin"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.Body = body
	nReq.Pagging = page
	nReq.Limit = limit

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchDataInnerJoin FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "MatchDataInnerJoin FAILED: " + err.Error()}
	}
	return resp
}

// NatsGetTotalInnerJoin - NatsGetTotalInnerJoin
func NatsGetTotalInnerJoin(apiKey, bucketID string, body interface{}) map[string]interface{} {
	subj := "vdd_request.GetTotalInnerJoin"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.BucketID = bucketID
	nReq.Body = body

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetTotalInnerJoin FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "GetTotalInnerJoin FAILED: " + err.Error()}
	}

	return resp
}

// IncreaseGlobalID - IncreaseGlobalID
func IncreaseGlobalID(apiKey, globalID string) map[string]interface{} {
	subj := "vdd_request.global_id_management.increase"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
	nReq.Body = map[string]interface{}{"global_id": globalID}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "IncreaseGlobalID FAILED: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "IncreaseGlobalID FAILED: " + err.Error()}
	}

	return resp
}
