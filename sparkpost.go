package sg

import "encoding/json"

// SparkPostService serializes a mail for SparkPost API.
type SparkPostService struct{}

// Authorize implements the Service interface.
func (*SparkPostService) Authorize(key string) string { return key }

// Serialize implements the Service interface.
func (*SparkPostService) Serialize(m *Mail) ([]byte, error) {
	return json.Marshal(&struct {
		Recipients       []H `json:"recipients"`
		SubstitutionData H   `json:"substitution_data,omitempty"`
		Content          H   `json:"content"`
	}{
		Recipients:       []H{{"address": m.To}},
		Content:          H{"template_id": m.TemplateID},
		SubstitutionData: m.Substitutions,
	})
}

// NewSparkPostClient creates a new client with a SparkPost API key.
func NewSparkPostClient(apiKey string) Sender {
	return &Client{
		APIKey:  apiKey,
		APIURL:  "https://api.sparkpost.com/api/v1/transmissions?num_rcpt_errors=3",
		Service: new(SparkPostService),
	}
}
