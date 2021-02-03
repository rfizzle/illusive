package examples

import (
	"encoding/json"
	"fmt"
	"github.com/rfizzle/illusive"
	"log"
	"strings"
)

func main() {
	client, err := illusive.NewClient(
		"https://illusive.example.com",
		"Basic <Token>",
		illusive.ClientDisableTLSValidation,
		illusive.ClientTimeout(60),
	)

	if err != nil {
		log.Fatal(err)
	}

	results, err := client.ForensicsIncidentSummary("1")
	summary := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range results {
		tmpText := v.Template
		for k, v2 := range v.TemplateData {
			tmpText = strings.ReplaceAll(tmpText, k, v2)
		}
		summary = append(summary, tmpText)
	}

	for k, v := range summary {
		fmt.Printf("%s:\n%s\n--------\n", results[k].AnalysisType, v)
	}
}
