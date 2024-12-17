package graphql

import (
	"net/http"
	"time"
)

/***
 * Designed to be more or less what The Guild says HTTP Details should look like in Extensions
 * https://the-guild.dev/graphql/mesh/docs/plugins/http-details-extensions
 */

type RequestDetail struct {
	Headers   http.Header `json:"headers"`
	Timestamp time.Time   `json:"timestamp"`
	URL       string      `json:"url"`
	Method    string      `json:"method"`
}

type ResponseDetail struct {
	Headers    http.Header `json:"headers"`
	Timestamp  time.Time   `json:"timestamp"`
	StatusText string      `json:"statusText"`
	StatusCode int         `json:"status"`
}

type HTTPDetails struct {
	// not sure these are available
	// SourceName   string            `json:"sourceName"`
	// Path         map[string]string `json:"path"`
	Request      RequestDetail  `json:"request"`
	Response     ResponseDetail `json:"response"`
	ResponseTime int64          `json:"responseTime"`
}

func BuildHTTPDetails(req *http.Request, resp *http.Response, requestTime time.Time, responseTime time.Time) HTTPDetails {
	return HTTPDetails{
		Request: RequestDetail{
			Timestamp: requestTime,
			URL:       req.URL.String(),
			Method:    req.Method,
			Headers:   req.Header,
		},
		Response: ResponseDetail{
			Timestamp:  responseTime,
			StatusCode: resp.StatusCode,
			StatusText: resp.Status,
			Headers:    resp.Header,
		},
		ResponseTime: responseTime.Sub(requestTime).Milliseconds(),
	}
}
