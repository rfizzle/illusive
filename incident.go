package illusive

import "encoding/json"

// IncidentDetails returns the details of an incident
func (c *Client) Incident(incidentID string, parameters ...RequestParameter) (*IncidentDetails, error) {
	// Setup params
	validParams := incidentDetailsParams()
	parameters = append(parameters, Param("incident_id", incidentID))
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/api/v2/incidents/incident", &params)

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
	var details IncidentDetails
	err = json.Unmarshal(responseData, &details)
	return &details, err
}
