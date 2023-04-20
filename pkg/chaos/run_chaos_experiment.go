package chaos

import (
	"encoding/json"

	"github.com/sagarkrsd/chaos-go-client/pkg/utils"
)

type RunChaosExperimentResponse struct {
	Data RunChaosExperimentData `json:"data"`
}

type RunChaosExperimentData struct {
	RunChaosExp RunChaosExp `json:"runChaosExperiment"`
}

type RunChaosExp struct {
	NotifyID string `json:"notifyID"`
}

type RunChaosExperimentVariables struct {
	Identifiers Identifiers `json:"identifiers"`
	WorkflowID  interface{} `json:"workflowID"`
}

// RunChaosExperiment executes a Chaos experiment on a given Chaos infra.
func RunChaosExperiment(url, workflowID string, identifiers Identifiers) (RunChaosExp, error) {
	method := "POST"

	/*
		runChaosExperimentAPIQuery :=
			fmt.Sprintf("{\"query\":\"mutation RunChaosExperiment(\\n  $workflowID: String!,\\n  $identifiers: IdentifiersRequest!\\n) {\\n  runChaosExperiment(\\n    workflowID: $workflowID,\\n    identifiers: $identifiers\\n  ) {\\n    notifyID\\n  }\\n}\",\"variables\":{\"workflowID\":\"%s\",\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"}}}",
				workflowID, identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier)
	*/

	runChaosExperimentRes := RunChaosExperimentResponse{}

	variables := RunChaosExperimentVariables{
		Identifiers: identifiers,
		WorkflowID:  workflowID,
	}
	query := map[string]interface{}{
		"query": `
		mutation RunChaosExperiment(
			$workflowID: String!,
			$identifiers: IdentifiersRequest!
		  ) {
			runChaosExperiment(
			  workflowID: $workflowID,
			  identifiers: $identifiers
			) {
			  notifyID
			}
		}
		`,
		"variables": variables,
	}

	response, err := utils.SendRequest(url, method, query)
	if err != nil {
		return runChaosExperimentRes.Data.RunChaosExp, err
	}

	err = json.Unmarshal(response, &runChaosExperimentRes)
	if err != nil {
		return runChaosExperimentRes.Data.RunChaosExp, err
	}

	//fmt.Printf("Successfully ran a given Chaos experiment, response: %+v", runChaosExperimentRes)

	return runChaosExperimentRes.Data.RunChaosExp, nil
}
