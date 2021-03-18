package vdp

import (
	"encoding/json"
	"time"

	"github.com/hodtien/extension-lib/model"
	"github.com/hodtien/extension-lib/global"
	
	"github.com/google/uuid"
)

// NatsCheckUserExists - NatsCheckUserExists
func NatsCheckUserExists(apiKey, username string) map[string]interface{} {
	subj := "vdp_request.authentication.check_user_exists"
	nReq := map[string]interface{}{
		"request_id": "CheckUserExists-" + uuid.New().String(),
		"api_key":    apiKey,
		"username":   username,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check User Exists Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check User Exists Failed: " + err.Error()}
	}

	return resp
}

// NatsCheckEmailExists - NatsCheckEmailExists
func NatsCheckEmailExists(apiKey, email string) map[string]interface{} {
	subj := "vdp_request.authentication.check_email_exists"

	nReq := map[string]interface{}{
		"request_id": "CheckEmailExists-" + uuid.New().String(),
		"api_key":    apiKey,
		"email":      email,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Email Exists Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Email Exists Failed: " + err.Error()}
	}

	return resp
}

// NatsCheckPhoneExists - NatsCheckPhoneExists
func NatsCheckPhoneExists(apiKey, phone string) map[string]interface{} {
	subj := "vdp_request.authentication.check_phone_number_exists"

	nReq := map[string]interface{}{
		"request_id":   "CheckPhoneExist-" + uuid.New().String(),
		"api_key":      apiKey,
		"phone_number": phone,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Phone Number Exists Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Check Phone Number Exists Failed: " + err.Error()}
	}

	return resp
}

// NatsUserRegister - NatsUserRegister
func NatsUserRegister(apiKey string, dataRegister map[string]interface{}) map[string]interface{} {
	subj := "vdp_request.authentication.register"

	nReq := map[string]interface{}{
		"request_id": "Register-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       dataRegister,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Register Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Register Failed: " + err.Error()}
	}

	return resp
}

// NatsUserLogin - NatsUserLogin
func NatsUserLogin(apiKey, username, password string) map[string]interface{} {
	subj := "vdp_request.authentication.login"

	nReq := map[string]interface{}{
		"request_id": "Login-" + uuid.New().String(),
		"api_key":    apiKey,
		"username":   username,
		"password":   password,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Login Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Login Failed: " + err.Error()}
	}

	return resp
}

// NatsGetProfile - NatsGetProfile
func NatsGetProfile(accessToken string) map[string]interface{} {
	subj := "vdp_request.authentication.get_profile"

	nReq := map[string]interface{}{
		"request_id": "GetProfile-" + uuid.New().String(),
		"api_key":    accessToken,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Profile Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get Profile Failed: " + err.Error()}
	}

	return resp
}

// NatsRefreshToken - NatsRefreshToken
func NatsRefreshToken(apiKey, refreshToken string) map[string]interface{} {
	subj := "vdp_request.authentication.refresh_token"

	nReq := map[string]interface{}{
		"request_id":    "RefrehToken-" + uuid.New().String(),
		"api_key":       apiKey,
		"refresh_token": refreshToken,
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Refresh Token Failed: "+ err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Refresh Token Failed: "+ err.Error()}
	}

	return resp
}

// NatsChangePassword - NatsChangePassword
func NatsChangePassword(apiKey string, old, new, confirm string) map[string]interface{} {
	subj := "vdp_request.authentication.update_password"

	nReq := map[string]interface{}{
		"request_id": "ChangePassword-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       model.PasswordModel{
			OldPassword: old,
			NewPassword: new,
			ConfirmNewPassword: confirm,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Change Password Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Change Password Failed: " + err.Error()}
	}

	return resp
}

// NatsRequestResetPasswordByEmail - NatsRequestResetPasswordByEmail
func NatsRequestResetPasswordByEmail(apiKey string, email string) map[string]interface{} {
	subj := "vdp_request.authentication.request_reset_password_email"

	nReq := map[string]interface{}{
		"request_id": "RequestResetPasswordByEmail-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       model.PasswordModel{
			Email: email,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Request Reset Password Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Request Reset Password Failed: " + err.Error()}
	}

	return resp
}

// NatsResetPasswordByEmail - NatsResetPasswordByEmail
func NatsResetPasswordByEmail(apiKey string, email, verifyCode, newPass, confirmPass string) map[string]interface{} {
	subj := "vdp_request.authentication.reset_password_email"

	nReq := map[string]interface{}{
		"request_id": "ResetPasswordByEmail-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       model.PasswordModel{
			Email: email,
			VerifyCode: verifyCode,
			NewPassword: newPass,
			ConfirmNewPassword: confirmPass,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Reset Password Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Reset Password Failed: " + err.Error()}
	}

	return resp
}

// NatsRequestResetPasswordByPhone - NatsRequestResetPasswordByPhone
func NatsRequestResetPasswordByPhone(apiKey string, phone string) map[string]interface{} {
	subj := "vdp_request.authentication.request_reset_password_phone"

	nReq := map[string]interface{}{
		"request_id": "RequestResetPasswordByPhone-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       model.PasswordModel{
			PhoneNumber: phone,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Request Reset Password Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Request Reset Password Failed: " + err.Error()}
	}

	return resp
}

// NatsResetPasswordByPhone - NatsResetPasswordByPhone
func NatsResetPasswordByPhone(apiKey string, phone, verifyCode, newPass, confirmPass string) map[string]interface{} {
	subj := "vdp_request.authentication.reset_password_phone"

	nReq := map[string]interface{}{
		"request_id": "ResetPasswordByPhone-" + uuid.New().String(),
		"api_key":    apiKey,
		"body":       model.PasswordModel{
			PhoneNumber: phone,
			VerifyCode: verifyCode,
			NewPassword: newPass,
			ConfirmNewPassword: confirmPass,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Reset Password Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Reset Password Failed: " + err.Error()}
	}

	return resp
}

// NatsGetUserByListID - NatsGetUserByListID
func NatsGetUserByListID(apiKey string, listID interface{}) map[string]interface{} {
	subj := "vdp_request.app_user_management.get_user_by_list_user_id"

	nReq := map[string]interface{}{
		"request_id": "GetListUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"body": map[string]interface{}{
			"user_id_list": listID,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get List User Failed: " + err.Error()}
	}

	return resp
}

// GetUserByPhoneOrEmail - GetUserByPhoneOrEmail
func GetUserByPhoneOrEmail(apiKey, field, value string) map[string]interface{} {
	subj := "vdp_request.app_user_management.get_user_by_phone_and_email"

	nReq := map[string]interface{}{
		"request_id": "GetListUser-" + uuid.New().String(),
		"api_key":    apiKey,
		"body": map[string]interface{}{
			field: value,
		},
	}

	payload, _ := json.Marshal(nReq)
	msg, err := global.Nc.Request(subj, payload, 15*time.Second)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get User Failed: " + err.Error()}
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return map[string]interface{}{"code": "10", "message": "Get User Failed: " + err.Error()}
	}

	return resp
}
