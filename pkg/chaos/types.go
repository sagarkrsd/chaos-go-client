package chaos

// Identifiers are the common identifiers for a given entity
type Identifiers struct {
	AccountIdentifier string `json:"accountIdentifier"`
	OrgIdentifier     string `json:"orgIdentifier"`
	ProjectIdentifier string `json:"projectIdentifier"`
}

// Query is the GraphQL query sent to the GraphQL server...
type GraphQLQuery struct {
	Query     string `json:"query"`
	Variables map[string]interface{}
}

type WorkflowType string

// Defines the details of the weightages of each chaos experiment in the workflow
type WeightagesInput struct {
	// Name of the experiment
	ExperimentName string `json:"experimentName"`
	// Weightage of the experiment
	Weightage int `json:"weightage"`
}

type EventMetadataInput struct {
	FaultName             string   `json:"faultName"`
	ServiceIdentifier     []string `json:"serviceIdentifier"`
	EnvironmentIdentifier []string `json:"environmentIdentifier"`
}

// Defines the details for a chaos workflow
type ChaosWorkFlowRequest struct {
	// ID of the workflow
	WorkflowID *string `json:"workflowID"`
	// Boolean check indicating if the created scenario will be executed or not
	RunExperiment *bool `json:"runExperiment"`
	// Manifest of the workflow
	WorkflowManifest string `json:"workflowManifest"`
	// Type of the workflow
	WorkflowType *WorkflowType `json:"workflowType"`
	// Cron syntax of the workflow schedule
	CronSyntax string `json:"cronSyntax"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Description of the workflow
	WorkflowDescription string `json:"workflowDescription"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*WeightagesInput `json:"weightages"`
	// Bool value indicating whether the workflow is a custom workflow or not
	IsCustomWorkflow bool `json:"isCustomWorkflow"`
	// ID of the target infra in which the workflow will run
	InfraID string `json:"infraID"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Type of the infra
	InfraType *InfrastructureType `json:"infraType"`
}
