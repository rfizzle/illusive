package illusive

import (
	"fmt"
	"net/url"
)

func Param(key, value string) func() (string, string) {
	return func() (string, string) {
		return key, value
	}
}

func setupParams(reqParams []RequestParameter, valid []string) (url.Values, error) {
	params := url.Values{}

	// Set request parameters
	for _, op := range reqParams {
		pKey, pVal := op()
		if !containsString(pKey, valid) {
			return nil, fmt.Errorf("invalid parameter: %s", pKey)
		}
		params.Set(pKey, pVal)
	}

	return params, nil
}

func incidentDetailsParams() []string {
	return []string{
		"incident_id",
	}
}

func forensicsIncidentSummary() []string {
	return []string{
		"incident_id",
	}
}

func forensicsTimelineParams() []string {
	return []string{
		"incident_id",
		"start_date",
		"end_date",
	}
}

func containsString(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
