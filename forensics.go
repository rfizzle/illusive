package illusive

import "encoding/json"

func (c *Client) ForensicsTimeline(incidentID string, parameters ...RequestParameter) (interface{}, error) {
	// Setup params
	validParams := incidentEventParams()
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
	var timeline interface{}
	err = json.Unmarshal(responseData, &timeline)
	return timeline, err
}
