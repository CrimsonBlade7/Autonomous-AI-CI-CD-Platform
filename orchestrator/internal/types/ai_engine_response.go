package types

import "encoding/json"

// The response from the AI Engine.
// Response should come with a HMAC-Signature-256 header.
// Contains the tests and the summary.
type AIEngineResponse struct {
	Wfid        int
	PullRequest PullRequest

	// Tests are ignored if Done.
	TestCmd  []string
	TestName string
	Tests    []byte

	// Done should always be accompanied by Summary.
	Done    bool
	Summary string

	// Review guidance is returned when a generated test still fails.
	Suggestions   string
	Documentation string
}

func (aier *AIEngineResponse) UnmarshalJSON(data []byte) error {
	type alias AIEngineResponse
	var aux struct {
		*alias
		TestCommand []string `json:"TestCommand"`
	}
	aux.alias = (*alias)(aier)

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aier.TestCmd) == 0 {
		aier.TestCmd = aux.TestCommand
	}
	return nil
}
