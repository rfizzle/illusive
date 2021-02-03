package illusive

import (
	"time"
)

type IncidentDetails struct {
	Closed            bool      `json:"closed"`
	DeceptionFamilies []string  `json:"deceptionFamilies"`
	Flagged           bool      `json:"flagged"`
	HasForensics      bool      `json:"hasForensics"`
	IncidentID        int       `json:"incidentId"`
	IncidentSeverity  string    `json:"incidentSeverity"`
	IncidentTimeUTC   time.Time `json:"incidentTimeUTC"`
	IncidentTypes     []string  `json:"incidentTypes"`
	LastSeenUser      *string   `json:"lastSeenUser"`
	PolicyName        *string   `json:"policyName"`
	RiskInsights      struct {
		StepsToCrownJewel  int `json:"stepsToCrownJewel"`
		StepsToDomainAdmin int `json:"stepsToDomainAdmin"`
	} `json:"riskInsights"`
	SourceHostname        *string `json:"sourceHostname"`
	SourceIP              *string `json:"sourceIp"`
	SourceOperatingSystem *string `json:"sourceOperatingSystem"`
	Unread                bool    `json:"unread"`
	UserNotes             *string `json:"userNotes"`
}

type IncidentAnalysisSummary struct {
	AnalysisType string            `json:"analysisType"`
	Template     string            `json:"template"`
	TemplateData map[string]string `json:"templateData"`
}
