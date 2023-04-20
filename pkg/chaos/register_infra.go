package chaos

import (
	"encoding/json"

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

type RegisterInfraRequest struct {
	Name             string `json:"name"`
	EnvironmentID    string `json:"environmentID"`
	Description      string `json:"description"`
	PlatformName     string `json:"platformName"`
	InfraNamespace   string `json:"infraNamespace"`
	ServiceAccount   string `json:"serviceAccount"`
	InfraScope       string `json:"infraScope"`
	InfraNsExists    bool   `json:"infraNsExists"`
	InfraSaExists    bool   `json:"infraSaExists"`
	InstallationType string `json:"installationType"`
	SkipSsl          bool   `json:"skipSsl"`
}

type RegisterInfraVariables struct {
	Identifiers Identifiers `json:"identifiers"`
	Request     interface{} `json:"request"`
}

// RegisterInfra registers a new Chaos infrastructure.
func RegisterNewInfra(req RegisterInfraRequest, url string, identifiers Identifiers) (RegisterInfra, error) {
	method := "POST"
	registerInfraRes := RegisterInfraResponse{}

	/*registerInfraAPIQuery :=
	fmt.Sprintf("{\"query\":\"mutation($identifiers: IdentifiersRequest!, $request: RegisterInfraRequest!) {\\n  registerInfra(identifiers: $identifiers, request: $request) {\\n    token\\n    infraID\\n    name\\n    manifest\\n  }\\n}\\n\",\"variables\":{\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"},\"request\":{\"name\":\"my-chaos-demo-infra\",\"environmentID\":\"my-chaos-demo-env\",\"description\":\"Chaos Demo Environment\",\"platformName\":\"my-chaos-demo-platform\",\"infraNamespace\":\"hce\",\"serviceAccount\":\"hce\",\"infraScope\":\"cluster\",\"infraNsExists\":false,\"infraSaExists\":false,\"installationType\":\"MANIFEST\",\"skipSsl\":false}}}",
		identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier)
	*/
	variables := RegisterInfraVariables{
		Identifiers: identifiers,
		Request:     req,
	}
	query := map[string]interface{}{
		"query": `
		mutation($identifiers: IdentifiersRequest!, $request: RegisterInfraRequest!) {
			registerInfra(identifiers: $identifiers, request: $request) {
			  token
			  infraID
			  name
			  manifest
			}
		}
		`,
		"variables": variables,
	}

	response, err := utils.SendRequest(url, method, query)
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
