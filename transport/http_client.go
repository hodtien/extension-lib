package transport

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// MakeHTTPGetRequest - MakeHTTPGetRequest
func MakeHTTPGetRequest(ctx context.Context, URL string, headerKeys, headerValues []string) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// MakeHTTPPostRequest - MakeHTTPPostRequest
func MakeHTTPPostRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// MakeHTTPPutRequest - MakeHTTPPutRequest
func MakeHTTPPutRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("PUT", URL, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// MakeHTTPPatchRequest - MakeHTTPPatchRequest
func MakeHTTPPatchRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("PATCH", URL, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// MakeHTTPDeleteRequest - MakeHTTPDeleteRequest
func MakeHTTPDeleteRequest(ctx context.Context, URL string, headerKeys, headerValues []string) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("DELETE", URL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// MakeHTTPPostFormRequest - MakeHTTPPostFormRequest
func MakeHTTPPostFormRequest(ctx context.Context, URL string, headerKeys, headerValues []string, multiPartWriter *multipart.Writer, bodyBytes bytes.Buffer) ([]byte, error) {
	timeout := time.Duration(20 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	request, err := http.NewRequest("POST", URL, &bodyBytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	for i := 0; i < len(headerKeys); i++ {
		if headerKeys[i] != "" && headerValues[i] != "" {
			request.Header.Set(headerKeys[i], headerValues[i])
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
