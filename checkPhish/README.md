Go to checkphish.ai to get an apikey.

# Usage: 
go build main.go
./main evil.com apikey


# checkPhish

```go
import "."
```

## Index

- [Constants](<#constants>)
- [func CPErrCodes(code int) error](<#func-cperrcodes>)
- [func NewHttpClient() (*http.Client, error)](<#func-newhttpclient>)
- [func PollInterval()](<#func-pollinterval>)
- [type CPClient](<#type-cpclient>)
  - [func NewAuthClient(hc *http.Client, apikey string) *CPClient](<#func-newauthclient>)
  - [func (sc *CPClient) GetResultJob(id string) (CPDoneResponse, error)](<#func-cpclient-getresultjob>)
  - [func (sc *CPClient) NewSearchJob(url string) (CPScanResponse, error)](<#func-cpclient-newsearchjob>)
  - [func (sc *CPClient) NewStatusJob(id string) (CPStatusResponse, error)](<#func-cpclient-newstatusjob>)
- [type CPDoneResponse](<#type-cpdoneresponse>)
- [type CPScanResponse](<#type-cpscanresponse>)
- [type CPStatusResponse](<#type-cpstatusresponse>)


## Constants

```go
const (
    BaseUrl = "https://developers.checkphish.ai/api"
)
```

## func CPErrCodes

```go
func CPErrCodes(code int) error
```

## func NewHttpClient

```go
func NewHttpClient() (*http.Client, error)
```

## func PollInterval

```go
func PollInterval()
```

## type CPClient

```go
type CPClient struct {
    // contains filtered or unexported fields
}
```

### func NewAuthClient

```go
func NewAuthClient(hc *http.Client, apikey string) *CPClient
```

### func \(\*CPClient\) GetResultJob

```go
func (sc *CPClient) GetResultJob(id string) (CPDoneResponse, error)
```

### func \(\*CPClient\) NewSearchJob

```go
func (sc *CPClient) NewSearchJob(url string) (CPScanResponse, error)
```

### func \(\*CPClient\) NewStatusJob

```go
func (sc *CPClient) NewStatusJob(id string) (CPStatusResponse, error)
```

## type CPDoneResponse

```go
type CPDoneResponse struct {
    JobID       string `json:"job_id"`
    Status      string `json:"status"`
    URL         string `json:"url"`
    URLSha256   string `json:"url_sha256"`
    Disposition string `json:"disposition"`
    Brand       string `json:"brand"`
    Error       bool   `json:"error"`
}
```

## type CPScanResponse

```go
type CPScanResponse struct {
    Jobid     string `json:"jobID"`
    Timestamp int64  `json:"timestamp"`
}
```

## type CPStatusResponse

```go
type CPStatusResponse struct {
    JobID  string `json:"job_id"`
    URL    string `json:"url"`
    Status string `json:"status"`
}
```
