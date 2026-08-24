package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

var errorMessages = map[int]string{
	http.StatusBadRequest:                   "Malformed input",                                    // 400 // RFC 9110, 15.5.1
	http.StatusUnauthorized:                 "Authentication required. Identity not verified.",    // 401 // RFC 9110, 15.5.2
	http.StatusForbidden:                    "Access denied. Permission mismatch.",                // 403 // RFC 9110, 15.5.4
	http.StatusNotFound:                     "Target resource does not exist.",                    // 404 // RFC 9110, 15.5.5
	http.StatusMethodNotAllowed:             "Operation rejected. Method not permitted.",          // 405 // RFC 9110, 15.5.6
	http.StatusNotAcceptable:                "No compatible representation available.",            // 406 // RFC 9110, 15.5.7
	http.StatusProxyAuthRequired:            "Proxy authentication required. Forwarding halted.",  // 407 // RFC 9110, 15.5.8
	http.StatusRequestTimeout:               "Signal lost. Client did not respond in time.",       // 408 // RFC 9110, 15.5.9
	http.StatusConflict:                     "State conflict detected. Cannot apply request.",     // 409 // RFC 9110, 15.5.10
	http.StatusGone:                         "Resource permanently removed.",                      // 410 // RFC 9110, 15.5.11
	http.StatusLengthRequired:               "Missing payload length. Transmission aborted.",      // 411 // RFC 9110, 15.5.12
	http.StatusRequestEntityTooLarge:        "Payload exceeds permitted size limits.",             // 413 // RFC 9110, 15.5.14
	http.StatusRequestURITooLong:            "Request path overflow. URI too long.",               // 414 // RFC 9110, 15.5.15
	http.StatusUnsupportedMediaType:         "Unsupported data format. Cannot decode payload.",    // 415 // RFC 9110, 15.5.16
	http.StatusRequestedRangeNotSatisfiable: "Requested range outside available bounds.",          // 416 // RFC 9110, 15.5.17
	http.StatusMisdirectedRequest:           "Request routed to incorrect origin.",                // 421 // RFC 9110, 15.5.20
	http.StatusLocked:                       "Resource locked. Concurrent access denied.",         // 423 // RFC 4918, 11.3
	http.StatusFailedDependency:             "Dependency failure. Upstream operation aborted.",    // 424 // RFC 4918, 11.4
	http.StatusTooEarly:                     "Request arrived before state stabilization.",        // 425 // RFC 8470, 5.2.
	http.StatusUpgradeRequired:              "Protocol upgrade required to proceed.",              // 426 // RFC 9110, 15.5.22
	http.StatusTooManyRequests:              "Rate limit exceeded. Slow down.",                    // 429 // RFC 6585, 4
	http.StatusRequestHeaderFieldsTooLarge:  "Header data overflow. Request rejected.",            // 431 // RFC 6585, 5
	http.StatusUnavailableForLegalReasons:   "Access blocked due to legal constraints.",           // 451 // RFC 7725, 3
	http.StatusInternalServerError:          "Internal fault. Execution path corrupted.",          // 500 // RFC 9110, 15.6.1
	http.StatusNotImplemented:               "Not implemented.",                                   // 501 // RFC 9110, 15.6.2
	http.StatusBadGateway:                   "Upstream response invalid or corrupted.",            // 502 // RFC 9110, 15.6.3
	http.StatusServiceUnavailable:           "Service offline. System under load or maintenance.", // 503 // RFC 9110, 15.6.4
	http.StatusGatewayTimeout:               "Upstream did not respond in time.",                  // 504 // RFC 9110, 15.6.5
	http.StatusHTTPVersionNotSupported:      "Protocol version unsupported by this server.",       // 505 // RFC 9110, 15.6.6
}

func Error(w http.ResponseWriter, code int) {}

/*
func Error(w http.ResponseWriter, code int) {
	var (
		errorType    = strconv.Itoa(code)
		errorMessage = "invalid."
	)

	// get message
	msg, oke := errorMessages[code]
	if oke {
		errorMessage = msg
	}

	// render error html
	var errorResp = fmt.Sprintf("%s: %s", errorType, errorMessage)

	renderedError, err := renderErrorHTML(errorType, errorMessage)
	if err == nil {
		errorResp = renderedError
	}

	// Respond
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(errorResp))
}

func renderErrorHTML(errType, errMessage string) (string, error) {
	return renderHTMLTemplate(
		errorHTML,
		struct {
			ErrorType    string
			ErrorMessage string
		}{
			ErrorType:    errType,
			ErrorMessage: errMessage,
		},
	)
}
*/
