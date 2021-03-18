package vcm

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hodtien/extension-library/transport"

	"github.com/google/uuid"
)

// NatsSendEmailDefault - NatsSendEmailDefault
func NatsSendEmailDefault(to []string, subject, body string) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"is_default_email": true,
			"body": map[string]interface{}{
				"to":      to,
				"subject": subject,
				"body":    body,
			},
		})

	emailSubject := "vcm_email_request.send_email_default"
	msg, err := transport.Nc.Request(emailSubject, requestBody, 15*time.Second)
	if err != nil {
		fmt.Println("Request VCM Send Email Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}

// NatsPushIOSNotification - NatsPushIOSNotification
func NatsPushIOSNotification(apiKey, to, title, body string) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"notification": map[string]interface{}{
					"title": title,
					"body":  body,
				},
				"to": to,
			},
		},
	)

	subject := "vcm_ios_request.push_notification"
	msg, err := transport.Nc.Request(subject, requestBody, 15*time.Second)
	if err != nil {
		fmt.Println("Request VCM Push IOS Notification Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}

// NatsPushAndroidNotification - NatsPushAndroidNotification
func NatsPushAndroidNotification(apiKey, to, title, body string, data interface{}) int {
	// initial
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"request_id": uuid.New().String(),
			"api_key":    apiKey,
			"body": map[string]interface{}{
				"notification": map[string]interface{}{
					"title":   title,
					"message": body,
				},
				"data": data,
				"to":   to,
			},
		},
	)

	subject := "vcm_android_request.push_notification"
	msg, err := transport.Nc.Request(subject, requestBody, 15*time.Second)
	if err != nil {
		fmt.Println("Request VCM Push Android Notification Error:", err)
		return -1
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(msg.Data, &result); err != nil {
		return -1
	}
	if result["code"] == nil || result["code"].(string) != "0" {
		return -1
	}
	return 0
}
