package vdd

import (
	"encoding/json"
	"time"

	"github.com/hodtien/extension-library/model"
	"github.com/hodtien/extension-library/transport"

	"github.com/google/uuid"
)

// NatsRetrieveData - NatsRetrieveData
func NatsRetrieveData(apiKey, bucketID, recordID, queryPath string) map[string]interface{} {
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

// NatsRetrieveAllDataInBucket - NatsRetrieveAllDataInBucket
func NatsRetrieveAllDataInBucket(apiKey, bucketID, queryPath string) map[string]interface{} {
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

// NatsRetrieveAllData - NatsRetrieveAllData
func NatsRetrieveAllData(apiKey string, start, stop int64) map[string]interface{} {
	subj := "vdd_request.RetrieveAllData"

	nReq := new(model.NATSRequest)
	nReq.RequestID = uuid.New().String()
	nReq.APIKey = apiKey
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

// NatsRetrieveManyDataByListID - NatsRetrieveManyDataByListID
func NatsRetrieveManyDataByListID(apiKey, bucketID string, listRecordID []string) map[string]interface{} {
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

// NatsSaveRecord - NatsSaveRecord
func NatsSaveRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
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

// NatsSaveRecordValidator - NatsSaveRecordValidator
func NatsSaveRecordValidator(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
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

// NatsUpdateRecordValidator - NatsUpdateRecordValidator
func NatsUpdateRecordValidator(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
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

// NatsDeleteRecord - NatsDeleteRecord
func NatsDeleteRecord(apiKey, bucketID, recordID string) map[string]interface{} {
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

// NatsDeleteManyRecord - NatsDeleteManyRecord
func NatsDeleteManyRecord(apiKey, bucketID string, listRecordID []string) map[string]interface{} {
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

// NatsUpdateAFieldRecord - NatsUpdateAFieldRecord
func NatsUpdateAFieldRecord(apiKey, bucketID, recordID, queryPath string, body interface{}) map[string]interface{} {
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

// NatsUpdateRecord - NatsUpdateRecord
func NatsUpdateRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
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

// NatsUpdateSomeFieldRecord - NatsUpdateSomeFieldRecord
func NatsUpdateSomeFieldRecord(apiKey, bucketID, recordID string, body interface{}) map[string]interface{} {
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

// NatsUpdateManyRecords - NatsUpdateManyRecords
func NatsUpdateManyRecords(apiKey, bucketID string, listRecordID []string, fieldUpdate, valueUpdate string) map[string]interface{} {
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

// NatsIncreaseGlobalID - NatsIncreaseGlobalID
func NatsIncreaseGlobalID(apiKey, globalID string) map[string]interface{} {
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
