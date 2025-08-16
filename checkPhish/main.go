package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	BaseUrl = "https://developers.checkphish.ai/api"
)

type CPClient struct {
	hc     *http.Client
	apikey string
}

func NewAuthClient(hc *http.Client, apikey string) *CPClient {
	return &CPClient{
		hc:     hc,
		apikey: apikey,
	}
}

func NewHttpClient() (*http.Client, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
	}

	return http.DefaultClient, nil
}

func (sc *CPClient) NewSearchJob(url string) (CPScanResponse, error) {
	scanHeader := fmt.Sprintf(`{"url":"%s"}`, url)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/neo/scan", BaseUrl), nil)

	if err != nil {
		return CPScanResponse{}, err
	}
	req.Header.Add("apikey", fmt.Sprintf("%s", sc.apikey))
	req.Header.Add("urlInfo", scanHeader)
	req.Header.Add("insights", "true")
	resp, err := sc.hc.Do(req)
	if err = CPErrCodes(resp.StatusCode); err != nil {
		return CPScanResponse{}, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CPScanResponse{}, err
	}

	var sr CPScanResponse
	unmarshalErr := json.Unmarshal(bytes, &sr)
	return sr, unmarshalErr
}

func (sc *CPClient) NewStatusJob(id string) (CPStatusResponse, error) {

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/neo/scan/status", BaseUrl), nil)

	if err != nil {
		return CPStatusResponse{}, err
	}
	req.Header.Add("apikey", fmt.Sprintf("%s", sc.apikey))
	req.Header.Add("jobID", id)
	resp, err := sc.hc.Do(req)
	if err = CPErrCodes(resp.StatusCode); err != nil {
		return CPStatusResponse{}, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CPStatusResponse{}, err
	}

	var sr CPStatusResponse
	unmarshalErr := json.Unmarshal(bytes, &sr)

	return sr, unmarshalErr
}

func (sc *CPClient) GetResultJob(id string) (CPDoneResponse, error) {

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/neo/scan/status", BaseUrl), nil)

	if err != nil {
		return CPDoneResponse{}, err
	}
	req.Header.Add("apikey", fmt.Sprintf("%s", sc.apikey))
	req.Header.Add("jobID", id)
	resp, err := sc.hc.Do(req)
	if err = CPErrCodes(resp.StatusCode); err != nil {
		return CPDoneResponse{}, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CPDoneResponse{}, err
	}

	var sr CPDoneResponse
	unmarshalErr := json.Unmarshal(bytes, &sr)

	return sr, unmarshalErr
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "URL", "apikey")
		return
	}
	apikey := os.Args[2]
	nc, _ := NewHttpClient()
	cp := NewAuthClient(nc, apikey)
	s, err := cp.NewSearchJob(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	ss, _ := cp.NewStatusJob(s.Jobid)

	if ss.Status == "PENDING" {
		PollInterval()
		s, _ := cp.NewStatusJob(s.Jobid)
		if err != nil {
			fmt.Println(err)
		}
		s2, _ := cp.GetResultJob(s.JobID)
		fmt.Println(s2)
	} else {
		s2, _ := cp.GetResultJob(ss.JobID)
		fmt.Println(s2)
	}

}

func PollInterval() {
	time.Sleep(60 * time.Second)
}

func CPErrCodes(code int) error {
	switch code {
	case 200:
		return nil
	case 400:
		return fmt.Errorf("400	Request error. See response body for details.")
	case 401:
		return fmt.Errorf("401	Authentication failure, invalid access credentials. Check headers!")
	case 404:
		return fmt.Errorf("404, Requested endpoint does not exist.")
	case 409:
		return fmt.Errorf("409	Invalid operation for this endpoint. See response body for details.")
	case 500:
		return fmt.Errorf("500	Unspecified internal server error. See response body for details")
	case 503:
		return fmt.Errorf("503 Feature is disabled in configuration file.")
	default:
		return fmt.Errorf("status code (%d)", code)
	}
}

type CPScanResponse struct {
	Jobid     string `json:"jobID"`
	Timestamp int64  `json:"timestamp"`
}

type CPStatusResponse struct {
	JobID  string `json:"job_id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type CPDoneResponse struct {
	JobID       string `json:"job_id"`
	Status      string `json:"status"`
	URL         string `json:"url"`
	URLSha256   string `json:"url_sha256"`
	Disposition string `json:"disposition"`
	Brand       string `json:"brand"`
	Error       bool   `json:"error"`
}
