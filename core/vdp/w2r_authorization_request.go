package vdp

import (
	"encoding/json"
	"time"

	"github.com/hodtien/extension-lib/transport"
	
	"github.com/google/uuid"
)

// NatsCreateRoleInDomain - NatsCreateRoleInDomain
func NatsCreateRoleInDomain(apiKey string, role interface{}) map[string]interface{} {
	subj := "vdp_request.authorization.create_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "CreateRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Create Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Create Role Failed: " + err.Error()}
	}

	return resp
}

// NatsUpdateRoleInDomain - NatsUpdateRoleInDomain
func NatsUpdateRoleInDomain(apiKey string, role interface{}) map[string]interface{} {
	subj := "vdp_request.authorization.update_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "UpdateRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Update Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Update Role Failed: " + err.Error()}
	}

	return resp
}

// NatsGetDetailRoleInDomain - NatsGetDetailRoleInDomain
func NatsGetDetailRoleInDomain(apiKey string, role interface{}) map[string]interface{} {
	subj := "vdp_request.authorization.get_detail_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetDetailRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
		"role":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Role Failed: " + err.Error()}
	}

	return resp
}

// NatsDeleteRoleInDomain - NatsDeleteRoleInDomain
func NatsDeleteRoleInDomain(apiKey, role string) map[string]interface{} {
	subj := "vdp_request.authorization.delete_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "DeleteRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
		"role":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Role Failed: " + err.Error()}
	}

	return resp
}

// NatsGetAllRoleInDomain - NatsGetAllRoleInDomain
func NatsGetAllRoleInDomain(apiKey string) map[string]interface{} {
	subj := "vdp_request.authorization.get_all_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetAllRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All Role Failed: " + err.Error()}
	}

	return resp
}

// NatsGetAllUserIDByRole - NatsGetAllUserIDByRole
func NatsGetAllUserIDByRole(apiKey, role string, page, limit int) map[string]interface{} {
	subj := "vdp_request.authorization.get_users_for_role_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetUsersForRoleInDomain-" + uuid.New().String(),
		"api_key":    apiKey,
		"role":       role,
	}
	if page >= 0 && limit >= 0 {
		nReq["page"] = page
		nReq["limit"] = limit
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All User By Role Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get All User By Role Failed: " + err.Error()}
	}

	return resp
}

// NatsHasPermissionForUser - NatsHasPermissionForUser
func NatsHasPermissionForUser(apiKey, userID string, permission []string) map[string]interface{} {
	subj := "vdp_request.authorization.has_permission_for_user"

	nReq := map[string]interface{}{
		"request_id": "HasPermissionsForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
		"permission": permission,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Permission Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Permission Failed: " + err.Error()}
	}

	return resp
}

// NatsGetRolesForUserInDomain - NatsGetRolesForUserInDomain
func NatsGetRolesForUserInDomain(apiKey, userID string) map[string]interface{} {
	subj := "vdp_request.authorization.get_roles_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetRolesForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Roles For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Roles For User Failed: " + err.Error()}
	}

	return resp
}

// NatsAddRoleForUserInDomain - NatsAddRoleForUserInDomain
func NatsAddRoleForUserInDomain(apiKey, userID, role string) map[string]interface{} {
	subj := "vdp_request.authorization.add_role_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "AddRoleForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
		"role":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Add Role For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Add Role For User Failed: " + err.Error()}
	}

	return resp
}

// NatsDeleteRoleForUserInDomain - NatsDeleteRoleForUserInDomain
func NatsDeleteRoleForUserInDomain(apiKey, userID, role string) map[string]interface{} {
	subj := "vdp_request.authorization.delete_role_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "DeleteRoleForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
		"role":       role,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Role For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Role For User Failed: " + err.Error()}
	}

	return resp
}

// NatsGetPermissionsForUserInDomain - NatsGetPermissionsForUserInDomain
func NatsGetPermissionsForUserInDomain(apiKey, userID string) map[string]interface{} {
	subj := "vdp_request.authorization.get_permissions_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetPermissionsForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Permissions For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Permissions For User Failed: " + err.Error()}
	}

	return resp
}

// NatsAddPermissionForUserInDomain - NatsAddPermissionForUserInDomain
func NatsAddPermissionForUserInDomain(apiKey, userID string, permission []string) map[string]interface{} {
	subj := "vdp_request.authorization.add_permission_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "AddPermissionForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
		"permission": permission,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Add Permission For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Add Permission For User Failed: " + err.Error()}
	}

	return resp
}

// NatsDeletePermissionForUserInDomain - NatsDeletePermissionForUserInDomain
func NatsDeletePermissionForUserInDomain(apiKey, userID string, permission []string) map[string]interface{} {
	subj := "vdp_request.authorization.delete_permission_for_user_in_domain"

	nReq := map[string]interface{}{
		"request_id": "DeletePermissionForUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"user_id":    userID,
		"permission": permission,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Permission For User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Delete Permission For User Failed: " + err.Error()}
	}

	return resp
}

// NatsGetUsersForPermissionInDomain - NatsGetUsersForPermissionInDomain
func NatsGetUsersForPermissionInDomain(apiKey string, permission []string) map[string]interface{} {
	subj := "vdp_request.authorization.get_users_for_permission_in_domain"

	nReq := map[string]interface{}{
		"request_id": "GetUsersForPermission-" + uuid.New().String(),
		"api_key":    apiKey,
		"permission": permission,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := transport.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Users For Permission Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Users For Permission Failed: " + err.Error()}
	}

	return resp
}