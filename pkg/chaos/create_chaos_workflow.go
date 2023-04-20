package chaos

import (
	"encoding/json"
	"fmt"

	"github.com/sagarkrsd/chaos-go-client/pkg/utils"
)

type createChaosWorkflowResponse struct {
	Data CreateWorkflowData `json:"data"`
}

type CreateWorkflowData struct {
	CreateChaosWorkFlow CreateChaosWorkflow `json:"createChaosWorkFlow"`
}

type CreateChaosWorkflow struct {
	WorkflowID          string   `json:"workflowID"`
	WorkflowName        string   `json:"workflowName"`
	WorkflowDescription string   `json:"workflowDescription"`
	Tags                []string `json:"tags"`
	IsCustomWorkflow    bool     `json:"isCustomWorkflow"`
	CronSyntax          string   `json:"cronSyntax"`
}

type CreateChaosWorkFlowVariables struct {
	Identifiers Identifiers        `json:"identifiers"`
	Request     ListWorkflowRunReq `json:"request"`
}

// CreateChaosWorkFlow creates a new Chaos workflow/experiment.
func CreateChaosWorkFlow(req ChaosWorkFlowRequest, url, infraID string, identifiers Identifiers) (CreateChaosWorkflow, error) {
	method := "POST"

	/*
		createChaosWorkflowAPIQuery :=
			fmt.Sprintf("{\"query\":\"mutation CreateChaosWorkFlow(\\n  $request: ChaosWorkFlowRequest!\\n  $identifiers: IdentifiersRequest!\\n) {\\n  createChaosWorkFlow(request: $request, identifiers: $identifiers) {\\n    workflowID\\n    cronSyntax\\n    workflowName\\n    workflowDescription\\n    isCustomWorkflow\\n    tags\\n  }\\n}\\n\",\"variables\":{\"identifiers\":{\"orgIdentifier\":\"%s\",\"accountIdentifier\":\"%s\",\"projectIdentifier\":\"%s\"},\"request\":{\"workflowManifest\":\"{\\r\\n    \\\"kind\\\": \\\"Workflow\\\",\\r\\n    \\\"apiVersion\\\": \\\"argoproj.io/v1alpha1\\\",\\r\\n    \\\"metadata\\\": {\\r\\n        \\\"name\\\": \\\"my-pod-delete-experiment\\\",\\r\\n        \\\"namespace\\\": \\\"hce\\\",\\r\\n        \\\"creationTimestamp\\\": null,\\r\\n        \\\"labels\\\": {\\r\\n            \\\"infra_id\\\": \\\"e640346b-bff5-40db-a828-68b44e06e2d1\\\",\\r\\n            \\\"revision_id\\\": \\\"f82e8c99-8efd-4289-a324-eadfd9667574\\\",\\r\\n            \\\"workflow_id\\\": \\\"80557c8b-531c-49d2-a39e-9f72c4674c2c\\\",\\r\\n            \\\"workflows.argoproj.io/controller-instanceid\\\": \\\"e640346b-bff5-40db-a828-68b44e06e2d1\\\"\\r\\n        }\\r\\n    },\\r\\n    \\\"spec\\\": {\\r\\n        \\\"templates\\\": [\\r\\n            {\\r\\n                \\\"name\\\": \\\"test-0-7-x-exp\\\",\\r\\n                \\\"inputs\\\": {},\\r\\n                \\\"outputs\\\": {},\\r\\n                \\\"metadata\\\": {},\\r\\n                \\\"steps\\\": [\\r\\n                    [\\r\\n                        {\\r\\n                            \\\"name\\\": \\\"install-chaos-faults\\\",\\r\\n                            \\\"template\\\": \\\"install-chaos-faults\\\",\\r\\n                            \\\"arguments\\\": {}\\r\\n                        }\\r\\n                    ],\\r\\n                    [\\r\\n                        {\\r\\n                            \\\"name\\\": \\\"pod-delete-ji5\\\",\\r\\n                            \\\"template\\\": \\\"pod-delete-ji5\\\",\\r\\n                            \\\"arguments\\\": {}\\r\\n                        }\\r\\n                    ],\\r\\n                    [\\r\\n                        {\\r\\n                            \\\"name\\\": \\\"cleanup-chaos-resources\\\",\\r\\n                            \\\"template\\\": \\\"cleanup-chaos-resources\\\",\\r\\n                            \\\"arguments\\\": {}\\r\\n                        }\\r\\n                    ]\\r\\n                ]\\r\\n            },\\r\\n            {\\r\\n                \\\"name\\\": \\\"install-chaos-faults\\\",\\r\\n                \\\"inputs\\\": {\\r\\n                    \\\"artifacts\\\": [\\r\\n                        {\\r\\n                            \\\"name\\\": \\\"pod-delete-ji5\\\",\\r\\n                            \\\"path\\\": \\\"/tmp/pod-delete-ji5.yaml\\\",\\r\\n                            \\\"raw\\\": {\\r\\n                                \\\"data\\\": \\\"apiVersion: litmuschaos.io/v1alpha1\\\\ndescription:\\\\n  message: |\\\\n    Deletes a pod belonging to a deployment/statefulset/daemonset\\\\nkind: ChaosExperiment\\\\nmetadata:\\\\n  name: pod-delete\\\\n  labels:\\\\n    name: pod-delete\\\\n    app.kubernetes.io/part-of: litmus\\\\n    app.kubernetes.io/component: chaosexperiment\\\\n    app.kubernetes.io/version: ci\\\\nspec:\\\\n  definition:\\\\n    scope: Namespaced\\\\n    permissions:\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - pods\\\\n        verbs:\\\\n          - create\\\\n          - delete\\\\n          - get\\\\n          - list\\\\n          - patch\\\\n          - update\\\\n          - deletecollection\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - events\\\\n        verbs:\\\\n          - create\\\\n          - get\\\\n          - list\\\\n          - patch\\\\n          - update\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - configmaps\\\\n        verbs:\\\\n          - get\\\\n          - list\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - pods/log\\\\n        verbs:\\\\n          - get\\\\n          - list\\\\n          - watch\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - pods/exec\\\\n        verbs:\\\\n          - get\\\\n          - list\\\\n          - create\\\\n      - apiGroups:\\\\n          - apps\\\\n        resources:\\\\n          - deployments\\\\n          - statefulsets\\\\n          - replicasets\\\\n          - daemonsets\\\\n        verbs:\\\\n          - list\\\\n          - get\\\\n      - apiGroups:\\\\n          - apps.openshift.io\\\\n        resources:\\\\n          - deploymentconfigs\\\\n        verbs:\\\\n          - list\\\\n          - get\\\\n      - apiGroups:\\\\n          - \\\\\\\"\\\\\\\"\\\\n        resources:\\\\n          - replicationcontrollers\\\\n        verbs:\\\\n          - get\\\\n          - list\\\\n      - apiGroups:\\\\n          - argoproj.io\\\\n        resources:\\\\n          - rollouts\\\\n        verbs:\\\\n          - list\\\\n          - get\\\\n      - apiGroups:\\\\n          - batch\\\\n        resources:\\\\n          - jobs\\\\n        verbs:\\\\n          - create\\\\n          - list\\\\n          - get\\\\n          - delete\\\\n          - deletecollection\\\\n      - apiGroups:\\\\n          - litmuschaos.io\\\\n        resources:\\\\n          - chaosengines\\\\n          - chaosexperiments\\\\n          - chaosresults\\\\n        verbs:\\\\n          - create\\\\n          - list\\\\n          - get\\\\n          - patch\\\\n          - update\\\\n          - delete\\\\n    image: chaosnative/go-runner:3.0.0-saas\\\\n    imagePullPolicy: Always\\\\n    args:\\\\n      - -c\\\\n      - ./experiments -name pod-delete\\\\n    command:\\\\n      - /bin/bash\\\\n    env:\\\\n      - name: TOTAL_CHAOS_DURATION\\\\n        value: \\\\\\\"15\\\\\\\"\\\\n      - name: RAMP_TIME\\\\n        value: \\\\\\\"\\\\\\\"\\\\n      - name: FORCE\\\\n        value: \\\\\\\"true\\\\\\\"\\\\n      - name: CHAOS_INTERVAL\\\\n        value: \\\\\\\"5\\\\\\\"\\\\n      - name: PODS_AFFECTED_PERC\\\\n        value: \\\\\\\"\\\\\\\"\\\\n      - name: TARGET_CONTAINER\\\\n        value: \\\\\\\"\\\\\\\"\\\\n      - name: TARGET_PODS\\\\n        value: \\\\\\\"\\\\\\\"\\\\n      - name: DEFAULT_HEALTH_CHECK\\\\n        value: \\\\\\\"false\\\\\\\"\\\\n      - name: NODE_LABEL\\\\n        value: \\\\\\\"\\\\\\\"\\\\n      - name: SEQUENCE\\\\n        value: parallel\\\\n    labels:\\\\n      name: pod-delete\\\\n      app.kubernetes.io/part-of: litmus\\\\n      app.kubernetes.io/component: experiment-job\\\\n      app.kubernetes.io/version: ci\\\\n\\\"\\r\\n                            }\\r\\n                        }\\r\\n                    ]\\r\\n                },\\r\\n                \\\"outputs\\\": {},\\r\\n                \\\"metadata\\\": {},\\r\\n                \\\"container\\\": {\\r\\n                    \\\"name\\\": \\\"\\\",\\r\\n                    \\\"image\\\": \\\"chaosnative/k8s:2.11.0\\\",\\r\\n                    \\\"command\\\": [\\r\\n                        \\\"sh\\\",\\r\\n                        \\\"-c\\\"\\r\\n                    ],\\r\\n                    \\\"args\\\": [\\r\\n                        \\\"kubectl apply -f /tmp/ -n {{workflow.parameters.adminModeNamespace}} && sleep 30\\\"\\r\\n                    ],\\r\\n                    \\\"resources\\\": {}\\r\\n                }\\r\\n            },\\r\\n            {\\r\\n                \\\"name\\\": \\\"cleanup-chaos-resources\\\",\\r\\n                \\\"inputs\\\": {},\\r\\n                \\\"outputs\\\": {},\\r\\n                \\\"metadata\\\": {},\\r\\n                \\\"container\\\": {\\r\\n                    \\\"name\\\": \\\"\\\",\\r\\n                    \\\"image\\\": \\\"chaosnative/k8s:2.11.0\\\",\\r\\n                    \\\"command\\\": [\\r\\n                        \\\"sh\\\",\\r\\n                        \\\"-c\\\"\\r\\n                    ],\\r\\n                    \\\"args\\\": [\\r\\n                        \\\"kubectl delete chaosengine -l workflow_run_id={{workflow.uid}} -n {{workflow.parameters.adminModeNamespace}}\\\"\\r\\n                    ],\\r\\n                    \\\"resources\\\": {}\\r\\n                }\\r\\n            },\\r\\n            {\\r\\n                \\\"name\\\": \\\"pod-delete-ji5\\\",\\r\\n                \\\"inputs\\\": {\\r\\n                    \\\"artifacts\\\": [\\r\\n                        {\\r\\n                            \\\"name\\\": \\\"pod-delete-ji5\\\",\\r\\n                            \\\"path\\\": \\\"/tmp/chaosengine-pod-delete-ji5.yaml\\\",\\r\\n                            \\\"raw\\\": {\\r\\n                                \\\"data\\\": \\\"apiVersion: litmuschaos.io\\\\/v1alpha1\\\\r\\\\nkind: ChaosEngine\\\\r\\\\nmetadata:\\\\r\\\\n  namespace: \\\\\\\"{{workflow.parameters.adminModeNamespace}}\\\\\\\"\\\\r\\\\n  generateName: pod-delete-ji5\\\\r\\\\n  labels:\\\\r\\\\n    workflow_run_id: \\\\\\\"{{ workflow.uid }}\\\\\\\"\\\\r\\\\n    workflow_name: my-pod-delete-experiment\\\\r\\\\nspec:\\\\r\\\\n  appinfo:\\\\r\\\\n    appns: hce\\\\r\\\\n    applabel: app=cartservice\\\\r\\\\n    appkind: deployment\\\\r\\\\n  engineState: active\\\\r\\\\n  chaosServiceAccount: litmus-admin\\\\r\\\\n  experiments:\\\\r\\\\n    - name: pod-delete\\\\r\\\\n      spec:\\\\r\\\\n        components:\\\\r\\\\n          env:\\\\r\\\\n            - name: TOTAL_CHAOS_DURATION\\\\r\\\\n              value: \\\\\\\"30\\\\\\\"\\\\r\\\\n            - name: CHAOS_INTERVAL\\\\r\\\\n              value: \\\\\\\"10\\\\\\\"\\\\r\\\\n            - name: FORCE\\\\r\\\\n              value: \\\\\\\"false\\\\\\\"\\\\r\\\\n            - name: PODS_AFFECTED_PERC\\\\r\\\\n              value: \\\\\\\"\\\\\\\"\\\\r\\\\n        probe:\\\\r\\\\n          - name: http-cartservice-probe\\\\r\\\\n            type: httpProbe\\\\r\\\\n            mode: Continuous\\\\r\\\\n            runProperties:\\\\r\\\\n              probeTimeout: 500\\\\r\\\\n              retry: 1\\\\r\\\\n              interval: 1\\\\r\\\\n              stopOnFailure: false\\\\r\\\\n            httpProbe\\\\/inputs:\\\\r\\\\n              url: http:\\\\/\\\\/frontend\\\\/cart\\\\r\\\\n              method:\\\\r\\\\n                get:\\\\r\\\\n                  criteria: ==\\\\r\\\\n                  responseCode: \\\\\\\"200\\\\\\\"\\\\r\\\\n                  responseTimeout: 15\\\\r\\\\n\\\"\\r\\n                            }\\r\\n                        }\\r\\n                    ]\\r\\n                },\\r\\n                \\\"outputs\\\": {},\\r\\n                \\\"metadata\\\": {\\r\\n                    \\\"labels\\\": {\\r\\n                        \\\"weight\\\": \\\"10\\\"\\r\\n                    }\\r\\n                },\\r\\n                \\\"container\\\": {\\r\\n                    \\\"name\\\": \\\"\\\",\\r\\n                    \\\"image\\\": \\\"chaosnative/litmus-checker:2.11.0\\\",\\r\\n                    \\\"args\\\": [\\r\\n                        \\\"-file=/tmp/chaosengine-pod-delete-ji5.yaml\\\",\\r\\n                        \\\"-saveName=/tmp/engine-name\\\"\\r\\n                    ],\\r\\n                    \\\"resources\\\": {}\\r\\n                }\\r\\n            }\\r\\n        ],\\r\\n        \\\"entrypoint\\\": \\\"test-0-7-x-exp\\\",\\r\\n        \\\"arguments\\\": {\\r\\n            \\\"parameters\\\": [\\r\\n                {\\r\\n                    \\\"name\\\": \\\"adminModeNamespace\\\",\\r\\n                    \\\"value\\\": \\\"hce\\\"\\r\\n                }\\r\\n            ]\\r\\n        },\\r\\n        \\\"serviceAccountName\\\": \\\"argo-chaos\\\",\\r\\n        \\\"podGC\\\": {\\r\\n            \\\"strategy\\\": \\\"OnWorkflowCompletion\\\"\\r\\n        },\\r\\n        \\\"securityContext\\\": {\\r\\n            \\\"runAsUser\\\": 1000,\\r\\n            \\\"runAsNonRoot\\\": true\\r\\n        }\\r\\n    },\\r\\n    \\\"status\\\": {\\r\\n        \\\"startedAt\\\": null,\\r\\n        \\\"finishedAt\\\": null\\r\\n    }\\r\\n}\",\"cronSyntax\":\"\",\"workflowName\":\"my-pod-delete-experiment\",\"runExperiment\":false,\"workflowDescription\":\"This is a pod delete experiment\",\"weightages\":[{\"experimentName\":\"pod-delete-ji5\",\"weightage\":10}],\"isCustomWorkflow\":true,\"infraID\":\"%s\",\"infraType\":\"Kubernetes\",\"tags\":[\"test\",\"workflow\",\"gke\"]}}}",
				identifiers.OrgIdentifier, identifiers.AccountIdentifier, identifiers.ProjectIdentifier, infraID)

	*/
	createChaosWorkflowRes := createChaosWorkflowResponse{}

	variables := RegisterInfraVariables{
		Identifiers: identifiers,
		Request:     req,
	}
	query := map[string]interface{}{
		"query": `
		mutation CreateChaosWorkFlow(
			$request: ChaosWorkFlowRequest!
			$identifiers: IdentifiersRequest!
		  ) {
			createChaosWorkFlow(request: $request, identifiers: $identifiers) {
			  workflowID
			  cronSyntax
			  workflowName
			  workflowDescription
			  isCustomWorkflow
			  tags
			}
		}
		`,
		"variables": variables,
	}

	response, err := utils.SendRequest(url, method, query)
	if err != nil {
		return createChaosWorkflowRes.Data.CreateChaosWorkFlow, err
	}

	fmt.Println("response === ", string(response))
	err = json.Unmarshal(response, &createChaosWorkflowRes)
	if err != nil {
		return createChaosWorkflowRes.Data.CreateChaosWorkFlow, err
	}

	//fmt.Printf("Successfully created a new Chaos workflow, response: %+v", createChaosWorkflowRes)

	return createChaosWorkflowRes.Data.CreateChaosWorkFlow, nil
}
