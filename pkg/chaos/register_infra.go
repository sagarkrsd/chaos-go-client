package chaos

import (
	"encoding/json"
	"fmt"

	"github.com/sagarkrsd/chaos-go-client/pkg/utils"
)

type RegisterInfraResponse struct {
	Data RegisterInfraData `json:"data"`
}

type RegisterInfraData struct {
	RegisterInfra RegisterInfra `json:"registerInfra"`
}

type RegisterInfra struct {
	Token    string `json:"token"`
	InfraID  string `json:"infraID"`
	Name     string `json:"name"`
	Manifest string `json:"manifest"`
}

// RegisterInfra registers a new Chaos infrastructure.
func RegisterNewInfra(url string, identifiers Identifiers) (RegisterInfra, error) {
	method := "POST"
	registerInfraAPIQuery :=
		fmt.Sprintf("{\"query\":\"mutation($identifiers: IdentifiersRequest!, $request: RegisterInfraRequest!) {\\n  registerInfra(identifiers: $identifiers, request: $request) {\\n    token\\n    infraID\\n    name\\n    manifest\\n  }\\n}\\n\",\"variables\":{\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"},\"request\":{\"name\":\"my-chaos-demo-infra\",\"environmentID\":\"my-chaos-demo-env\",\"description\":\"Chaos Demo Environment\",\"platformName\":\"my-chaos-demo-platform\",\"infraNamespace\":\"hce\",\"serviceAccount\":\"hce\",\"infraScope\":\"cluster\",\"infraNsExists\":false,\"infraSaExists\":false,\"installationType\":\"MANIFEST\",\"skipSsl\":false}}}",
			identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier)

	registerInfraRes := RegisterInfraResponse{}

	response, err := utils.SendRequest(url, method, registerInfraAPIQuery)
	if err != nil {
		return registerInfraRes.Data.RegisterInfra, err
	}

	err = json.Unmarshal(response, &registerInfraRes)
	if err != nil {
		return registerInfraRes.Data.RegisterInfra, err
	}

	//fmt.Printf("Successfully registered a new Chaos infra, response: %+v", registerInfraRes)

	return registerInfraRes.Data.RegisterInfra, nil
}
