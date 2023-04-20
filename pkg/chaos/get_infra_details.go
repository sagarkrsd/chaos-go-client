package chaos

import (
	"encoding/json"

	"github.com/sagarkrsd/chaos-go-client/pkg/utils"
)

type GetInfraDetailsResponse struct {
	Data GetInfraDetailsData `json:"data"`
}

type GetInfraDetailsData struct {
	GetInfraDetails Infra `json:"getInfraDetails"`
}

type GetInfraDetailsVariables struct {
	Identifiers Identifiers `json:"identifiers"`
	InfraID     interface{} `json:"infraID"`
}

// GetInfraDetails gets the details of a given Chaos infrastructure.
func GetInfraDetails(url, infraID string, identifiers Identifiers) (Infra, error) {
	method := "POST"

	/*
		GetInfraDetailsAPIQuery :=
			fmt.Sprintf("{\"query\":\"query GetInfraDetails(\\n  $infraID: String!,\\n  $identifiers: IdentifiersRequest!\\n) {\\n  getInfraDetails(\\n    infraID: $infraID,\\n    identifiers: $identifiers\\n  ) {\\n    infraID\\n    name\\n    description\\n    tags\\n    environmentID\\n    platformName\\n    isActive\\n    isInfraConfirmed\\n    isRemoved\\n    updatedAt\\n    createdAt\\n    noOfSchedules\\n    noOfWorkflows\\n    token\\n    infraNamespace\\n    serviceAccount\\n    infraScope\\n    infraNsExists\\n    infraSaExists\\n    installationType\\n    k8sConnectorID\\n    lastWorkflowTimestamp\\n    startTime\\n    version\\n    createdBy {\\n      userID\\n      username\\n      email\\n    }\\n    updatedBy {\\n      userID\\n      username\\n      email\\n    }\\n  }\\n}\",\"variables\":{\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"},\"infraID\":\"%s\"}}",
				identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier, infraID)
	*/
	GetInfraDetailsRes := GetInfraDetailsResponse{}

	variables := GetInfraDetailsVariables{
		Identifiers: identifiers,
		InfraID:     infraID,
	}
	query := map[string]interface{}{
		"query": `
		query GetInfraDetails(
			$infraID: String!,
			$identifiers: IdentifiersRequest!
		  ) {
			getInfraDetails(
			  infraID: $infraID,
			  identifiers: $identifiers
			) {
			  infraID
			  name
			  description
			  tags
			  environmentID
			  platformName
			  isActive
			  isInfraConfirmed
			  isRemoved
			  updatedAt
			  createdAt
			  noOfSchedules
			  noOfWorkflows
			  token
			  infraNamespace
			  serviceAccount
			  infraScope
			  infraNsExists
			  infraSaExists
			  installationType
			  k8sConnectorID
			  lastWorkflowTimestamp
			  startTime
			  version
			  createdBy {
				userID
				username
				email
			  }
			  updatedBy {
				userID
				username
				email
			  }
			}
		}
		`,
		"variables": variables,
	}

	response, err := utils.SendRequest(url, method, query)
	if err != nil {
		return GetInfraDetailsRes.Data.GetInfraDetails, err
	}

	err = json.Unmarshal(response, &GetInfraDetailsRes)
	if err != nil {
		return GetInfraDetailsRes.Data.GetInfraDetails, err
	}

	//fmt.Printf("Got the infra details, response: %+v", GetInfraDetailsRes)

	return GetInfraDetailsRes.Data.GetInfraDetails, nil
}
