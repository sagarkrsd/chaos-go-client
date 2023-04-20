package chaos

import (
	"encoding/json"
	"fmt"

	"github.com/sagarkrsd/chaos-go-client/pkg/utils"
)

type ListWorkflowRunResponse struct {
	Data ListWorkflowRunData `json:"data"`
}

type ListWorkflowRunData struct {
	ListWorkflowRun ListWorkflowRn `json:"listWorkflowRun"`
}

type ListWorkflowRn struct {
	// Total number of workflow runs
	TotalNoOfWorkflowRuns int `json:"totalNoOfWorkflowRuns"`
	// Defines details of workflow runs
	WorkflowRuns []*WorkflowRun `json:"workflowRuns"`
}

type UserDetails struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type InstallationType string

type InfrastructureType string

type UpdateStatus string

// Defines the details for a infra
type Infra struct {
	// ID of the infra
	InfraID string `json:"infraID"`
	// Name of the infra
	Name string `json:"name"`
	// Description of the infra
	Description *string `json:"description"`
	// Tags of the infra
	Tags []string `json:"tags"`
	// Environment ID for the infra
	EnvironmentID string `json:"environmentID"`
	// Infra Platform Name eg. GKE,AWS, Others
	PlatformName string `json:"platformName"`
	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive"`
	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed"`
	// Boolean value indicating if chaos infrastructure is removed or not
	IsRemoved bool `json:"isRemoved"`
	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt"`
	// Number of schedules created in the infra
	NoOfSchedules *int `json:"noOfSchedules"`
	// Number of workflows run in the infra
	NoOfWorkflows *int `json:"noOfWorkflows"`
	// Token used to verify and retrieve the infra manifest
	Token string `json:"token"`
	// Namespace where the infra is being installed
	InfraNamespace *string `json:"infraNamespace"`
	// Name of service account used by infra
	ServiceAccount *string `json:"serviceAccount"`
	// Scope of the infra : ns or infra
	InfraScope string `json:"infraScope"`
	// Bool value indicating whether infra ns used already exists on infra or not
	InfraNsExists *bool `json:"infraNsExists"`
	// Bool value indicating whether service account used already exists on infra or not
	InfraSaExists *bool `json:"infraSaExists"`
	// InstallationType connector/manifest
	InstallationType InstallationType `json:"installationType"`
	// K8sConnectorID
	K8sConnectorID *string `json:"k8sConnectorID"`
	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp *string `json:"lastWorkflowTimestamp"`
	// Timestamp when the infra got connected
	StartTime string `json:"startTime"`
	// Version of the infra
	Version string `json:"version"`
	// User who created the infra
	CreatedBy *UserDetails `json:"createdBy"`
	// User who has updated the infra
	UpdatedBy *UserDetails `json:"updatedBy"`
	// Last Heartbeat status sent by the infra
	LastHeartbeat *string `json:"lastHeartbeat"`
	// Type of the infrastructure
	InfraType *InfrastructureType `json:"infraType"`
	// update status of infra
	UpdateStatus UpdateStatus `json:"updateStatus"`
}

type WorkflowRunStatus string

// Defines the details of a workflow run
type WorkflowRun struct {
	// Harness identifiers
	Identifiers *Identifiers `json:"identifiers"`
	// ID of the workflow run which is to be queried
	WorkflowRunID string `json:"workflowRunID"`
	// Type of the workflow
	WorkflowType *string `json:"workflowType"`
	// ID of the workflow
	WorkflowID string `json:"workflowID"`
	// Array containing weightage and name of each chaos experiment in the workflow
	Weightages []*Weightages `json:"weightages"`
	// Timestamp at which workflow run was last updated
	UpdatedAt string `json:"updatedAt"`
	// Timestamp at which workflow run was created
	CreatedAt string `json:"createdAt"`
	// Target infra in which the workflow will run
	Infra *Infra `json:"infra"`
	// Name of the workflow
	WorkflowName string `json:"workflowName"`
	// Manifest of the workflow run
	WorkflowManifest string `json:"workflowManifest"`
	// Phase of the workflow run
	Phase WorkflowRunStatus `json:"phase"`
	// Resiliency score of the workflow
	ResiliencyScore *float64 `json:"resiliencyScore"`
	// Number of experiments passed
	ExperimentsPassed *int `json:"experimentsPassed"`
	// Number of experiments failed
	ExperimentsFailed *int `json:"experimentsFailed"`
	// Number of experiments awaited
	ExperimentsAwaited *int `json:"experimentsAwaited"`
	// Number of experiments stopped
	ExperimentsStopped *int `json:"experimentsStopped"`
	// Number of experiments which are not available
	ExperimentsNa *int `json:"experimentsNa"`
	// Total number of experiments
	TotalExperiments *int `json:"totalExperiments"`
	// Stores all the workflow run details related to the nodes of DAG graph and chaos results of the experiments
	ExecutionData string `json:"executionData"`
	// Bool value indicating if the workflow run has removed
	IsRemoved *bool `json:"isRemoved"`
	// User who has updated the workflow
	UpdatedBy *UserDetails `json:"updatedBy"`
	// User who has created the experiment run
	CreatedBy *UserDetails `json:"createdBy"`
	// Notify ID of the experiment run
	NotifyID *string `json:"notifyID"`
}

// Defines the details of the weightages of each chaos experiment in the workflow
type Weightages struct {
	// Name of the experiment
	ExperimentName string `json:"experimentName"`
	// Weightage of the experiment
	Weightage int `json:"weightage"`
}

// ListWorkflowRun lists all the runs of a given Chaos workflow/experiment.
func ListWorkflowRun(url, notifyID string, identifiers Identifiers) (ListWorkflowRn, error) {
	method := "POST"

	listWorkflowRunAPIQuery :=
		fmt.Sprintf("{\"query\":\"query ListWorkflowRun(\\n  $identifiers: IdentifiersRequest!,\\n  $request: ListWorkflowRunRequest!\\n) {\\n  listWorkflowRun(\\n    identifiers: $identifiers,\\n    request: $request\\n  ) {\\n    totalNoOfWorkflowRuns\\n    workflowRuns {\\n      identifiers {\\n          orgIdentifier\\n          projectIdentifier\\n          accountIdentifier\\n      }\\n      workflowRunID\\n      workflowID\\n      weightages {\\n        experimentName\\n        weightage\\n      }\\n      updatedAt\\n      createdAt\\n      infra {\\n        infraID\\n        infraNamespace\\n        infraScope\\n        isActive\\n        isInfraConfirmed\\n      }\\n      workflowName\\n      workflowManifest\\n      phase\\n      resiliencyScore\\n      experimentsPassed\\n      experimentsFailed\\n      experimentsAwaited\\n      experimentsStopped\\n      experimentsNa\\n      totalExperiments\\n      executionData\\n      isRemoved\\n      updatedBy {\\n        userID\\n        username\\n      }\\n      createdBy {\\n        username\\n        userID\\n      }\\n    }\\n  }\\n}\",\"variables\":{\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"},\"request\":{\"notifyIDs\":[\"%s\"]}}}",
			identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier, notifyID)

	listWorkflowRunRes := ListWorkflowRunResponse{}

	response, err := utils.SendRequest(url, method, listWorkflowRunAPIQuery)
	if err != nil {
		return listWorkflowRunRes.Data.ListWorkflowRun, err
	}

	err = json.Unmarshal(response, &listWorkflowRunRes)
	if err != nil {
		return listWorkflowRunRes.Data.ListWorkflowRun, err
	}

	//fmt.Printf("Successfully listed the runs of a given Chaos experiment/workflow, response: %+v", listWorkflowRunRes)

	return listWorkflowRunRes.Data.ListWorkflowRun, nil
}
