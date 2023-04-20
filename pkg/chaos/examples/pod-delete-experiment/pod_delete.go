package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/sagarkrsd/chaos-go-client/pkg/chaos"
)

func main() {

	// An example url value: http://35.188.55.143:32500/gateway/chaos/manager/api/query?accountIdentifier=lgcHPsKuRaue8RTin4Efgw
	//
	// Apart from these values, please export your Harness API Key as env variable "X_API_Key"
	// wherever you are executing this code...
	url := "<replace-with-Harness-Chaos-server-url>"
	identifiers := chaos.Identifiers{
		AccountIdentifier: "<replace-with-your-Harness-account-ID>",
		OrgIdentifier:     "<replace-with-your-Harness-organisation-ID>",
		ProjectIdentifier: "<replace-with-your-Harness-project-ID>",
	}

	// Register a new infra if not already registered
	fmt.Println("Registering a new Chaos infra...")
	registerInfraRes, err := chaos.RegisterNewInfra(url, identifiers)
	if err != nil {
		fmt.Printf("Error registering a new Chaos infra: %+v", err)
		return
	}

	infraID, manifest := registerInfraRes.InfraID, registerInfraRes.Manifest
	if len(infraID) == 0 || len(manifest) == 0 {
		fmt.Printf("Invalid value for infraID/manifest from registerInfra API call, infraID: %s, manifest: %s",
			infraID, manifest)
		return
	} else {
		fmt.Printf("\nSuccessfully registered a new Infra with following details: infraID: %s\n",
			infraID)
	}

	// Store the infra manifest to a file named myInfraManifest.yaml in the current directory ...
	f, err := os.Create("myInfraManifest.yaml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(manifest + "\n")
	if err != nil {
		log.Fatal(err)
	}
	f.Sync()

	// Apply the infra manifest on a Kubernetes cluster as per the current K8s context where this
	// code is executed i.e., kubectl should be configured as a pre-requisite wherever this code is
	// executed to apply infra manifests.
	//
	// NOTE: You can use "--dry-run=client" with the below command in case you just want to see what all
	// resources will be deployed as part of Chaos infra on your cluster.
	cmd := exec.Command("kubectl", "apply", "-f", "myInfraManifest.yaml")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err, stderr.String())
	} else {
		fmt.Println("Applying manifest for registering a new infra...")
		fmt.Printf("\nOutput... %+v\n", out.String())
	}

	var infraDetails chaos.Infra
getInfraDetails:
	for {
		// Check if the infra is active or not before proceeding with the workflow creation...
		infraDetails, err = chaos.GetInfraDetails(url, infraID, identifiers)
		if err != nil {
			fmt.Printf("\nFailed to retrieve Chaos infra details: %+v\n", err)
			return
		}
		if infraDetails.IsActive {
			fmt.Printf("\nSuccessfully retrieved the infra details and the infra is active now: %+v\n", infraDetails)
			break getInfraDetails
		} else {
			fmt.Printf("\nWaiting on the Chaos infra to become active: %+v\n", infraDetails)
		}
		time.Sleep(15 * time.Second)
	}

	// Now create a Chaos workflow for running pod delete experiment on boutique application - cart service's pod
	// in hce namespace...
	fmt.Println("Creating a new Chaos workflow for running pod delete experiment on Boutique app...")
	createWorkflowRes, err := chaos.CreateChaosWorkFlow(url, infraID, identifiers)
	if err != nil {
		fmt.Printf("\nFailed to create Chaos workflow: %+v\n", err)
		return
	} else {
		fmt.Printf("\nChaos workflow successfully created, workflowID: %s\n", createWorkflowRes.WorkflowID)
	}

	fmt.Println("Preparing to run pod delete experiment on Boutique app now...")
	// Run the above experiment now...
	runChaosExperimentRes, err := chaos.RunChaosExperiment(url, createWorkflowRes.WorkflowID, identifiers)
	if err != nil {
		fmt.Printf("\nFailed to run Chaos experiment: %+v\n", err)
		return
	} else {
		fmt.Printf("\nRunning pod delete experiment on Boutique app now, notifyID to view results: %s\n",
			runChaosExperimentRes.NotifyID)
	}

	fmt.Println("Preparing to view the details of the ongoing pod delete experiment...")

viewChaosExpDetails:
	for {
		// Fetch/Observe the details of the above Chaos experiment run
		listWorkflowRunRes, err := chaos.ListWorkflowRun(url, runChaosExperimentRes.NotifyID, identifiers)
		if err != nil {
			fmt.Printf("\nFailed to list the runs of a given Chaos experiment: %+v\n", err)
			return
		} else {
			fmt.Printf("\nDetails of ongoing pod delete experiment... : %+v\n", listWorkflowRunRes)
		}
		// Run this loop for the duration of Chaos experiment run...

		if listWorkflowRunRes.TotalNoOfWorkflowRuns > 0 {
			for _, workflowRun := range listWorkflowRunRes.WorkflowRuns {
				if workflowRun.Phase == "Queued" || workflowRun.Phase == "Running" {
					fmt.Printf("\nChaos injection is in progress... awaiting results for workflow %s: %+v\n",
						workflowRun.WorkflowName, workflowRun.ExecutionData)
					time.Sleep(30 * time.Second)
				} else {
					// NOTE: Since we are already filtering the workflow runs based on notify IDs here, it will always be only
					// one workflow run running at a moment so we can break out of loop once it is completed here...
					fmt.Printf("\nChaos experiment execution results for workflow %s: %+v\n", workflowRun.WorkflowName,
						workflowRun.ExecutionData)
					break viewChaosExpDetails
				}
			}
		} else {
			break viewChaosExpDetails
		}
	}

}
