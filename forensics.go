package illusive

import "encoding/json"

func (c *Client) ForensicsIncidentSummary(incidentID string, parameters ...RequestParameter) ([]IncidentAnalysisSummary, error) {
	// Setup params
	validParams := forensicsIncidentSummary()
	parameters = append(parameters, Param("incident_id", incidentID))
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/api/v1/forensics/incident_summary", &params)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Conduct request
	responseData, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	// Unmarshal json body into structs and return
	var summary []IncidentAnalysisSummary
	err = json.Unmarshal(responseData, &summary)
	return summary, err
}

func (c *Client) ForensicsTimeline(incidentID string, parameters ...RequestParameter) ([]TimelineItem, error) {
	// Setup params
	validParams := forensicsTimelineParams()
	parameters = append(parameters, Param("incident_id", incidentID))
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/api/v1/forensics/timeline", &params)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Conduct request
	responseData, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	// Unmarshal json body into structs and return
	var timeline []TimelineItem
	err = json.Unmarshal(responseData, &timeline)
	return timeline, err
}
