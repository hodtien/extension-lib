package vfe

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/hodtien/extension-lib/global"
	"github.com/hodtien/extension-lib/model"
)

// UploadFiles - UploadFiles
func UploadFiles(userID, deviceID, bucketID string, files []*model.FileStorage) map[string]interface{} {
	subj := "fs_request.file_storage.upload_files"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.UserID = userID
	nReq.DeviceID = deviceID
	nReq.BucketID = bucketID
	nReq.Files = files

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}

// DeleteFiles - DeleteFiles
func DeleteFiles(userID, deviceID, bucketID string, recordIDs []interface{}) map[string]interface{} {
	subj := "fs_request.file_storage.delete_files"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.UserID = userID
	nReq.DeviceID = deviceID
	nReq.BucketID = bucketID
	nReq.RecordIDs = recordIDs

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}

// RetrieveObjsFolder - Retrieve all object in folder
func RetrieveObjsFolder(userID, deviceID, bucketID, folder string) map[string]interface{} {
	subj := "fs_request.file_storage.retrieve_objs_folder"
	nReq := new(model.NATSRequestFileStorage)
	nReq.RequestID = uuid.New().String()
	nReq.UserID = userID
	nReq.DeviceID = deviceID
	nReq.BucketID = bucketID
	nReq.Folder = folder

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}
	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "FAILED"}
	}

	return resp
}
