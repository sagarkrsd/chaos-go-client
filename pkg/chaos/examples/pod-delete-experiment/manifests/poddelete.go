package manifests

var PodDeleteWorkflowManifest = `
{
	"kind": "Workflow",
	"apiVersion": "argoproj.io/v1alpha1",
	"metadata": {
		"name": "my-pod-delete-experiment",
		"namespace": "hce",
		"creationTimestamp": null,
		"labels": {
			"infra_id": "e640346b-bff5-40db-a828-68b44e06e2d1",
			"revision_id": "f82e8c99-8efd-4289-a324-eadfd9667574",
			"workflow_id": "80557c8b-531c-49d2-a39e-9f72c4674c2c",
			"workflows.argoproj.io/controller-instanceid": "e640346b-bff5-40db-a828-68b44e06e2d1"
		}
	},
	"spec": {
		"templates": [
			{
				"name": "test-0-7-x-exp",
				"inputs": {},
				"outputs": {},
				"metadata": {},
				"steps": [
					[
						{
							"name": "install-chaos-faults",
							"template": "install-chaos-faults",
							"arguments": {}
						}
					],
					[
						{
							"name": "pod-delete-ji5",
							"template": "pod-delete-ji5",
							"arguments": {}
						}
					],
					[
						{
							"name": "cleanup-chaos-resources",
							"template": "cleanup-chaos-resources",
							"arguments": {}
						}
					]
				]
			},
			{
				"name": "install-chaos-faults",
				"inputs": {
					"artifacts": [
						{
							"name": "pod-delete-ji5",
							"path": "/tmp/pod-delete-ji5.yaml",
							"raw": {
								"data": "apiVersion: litmuschaos.io/v1alpha1\ndescription:\n  message: |\n    Deletes a pod belonging to a deployment/statefulset/daemonset\nkind: ChaosExperiment\nmetadata:\n  name: pod-delete\n  labels:\n    name: pod-delete\n    app.kubernetes.io/part-of: litmus\n    app.kubernetes.io/component: chaosexperiment\n    app.kubernetes.io/version: ci\nspec:\n  definition:\n    scope: Namespaced\n    permissions:\n      - apiGroups:\n          - \"\"\n        resources:\n          - pods\n        verbs:\n          - create\n          - delete\n          - get\n          - list\n          - patch\n          - update\n          - deletecollection\n      - apiGroups:\n          - \"\"\n        resources:\n          - events\n        verbs:\n          - create\n          - get\n          - list\n          - patch\n          - update\n      - apiGroups:\n          - \"\"\n        resources:\n          - configmaps\n        verbs:\n          - get\n          - list\n      - apiGroups:\n          - \"\"\n        resources:\n          - pods/log\n        verbs:\n          - get\n          - list\n          - watch\n      - apiGroups:\n          - \"\"\n        resources:\n          - pods/exec\n        verbs:\n          - get\n          - list\n          - create\n      - apiGroups:\n          - apps\n        resources:\n          - deployments\n          - statefulsets\n          - replicasets\n          - daemonsets\n        verbs:\n          - list\n          - get\n      - apiGroups:\n          - apps.openshift.io\n        resources:\n          - deploymentconfigs\n        verbs:\n          - list\n          - get\n      - apiGroups:\n          - \"\"\n        resources:\n          - replicationcontrollers\n        verbs:\n          - get\n          - list\n      - apiGroups:\n          - argoproj.io\n        resources:\n          - rollouts\n        verbs:\n          - list\n          - get\n      - apiGroups:\n          - batch\n        resources:\n          - jobs\n        verbs:\n          - create\n          - list\n          - get\n          - delete\n          - deletecollection\n      - apiGroups:\n          - litmuschaos.io\n        resources:\n          - chaosengines\n          - chaosexperiments\n          - chaosresults\n        verbs:\n          - create\n          - list\n          - get\n          - patch\n          - update\n          - delete\n    image: chaosnative/go-runner:3.0.0-saas\n    imagePullPolicy: Always\n    args:\n      - -c\n      - ./experiments -name pod-delete\n    command:\n      - /bin/bash\n    env:\n      - name: TOTAL_CHAOS_DURATION\n        value: \"15\"\n      - name: RAMP_TIME\n        value: \"\"\n      - name: FORCE\n        value: \"true\"\n      - name: CHAOS_INTERVAL\n        value: \"5\"\n      - name: PODS_AFFECTED_PERC\n        value: \"\"\n      - name: TARGET_CONTAINER\n        value: \"\"\n      - name: TARGET_PODS\n        value: \"\"\n      - name: DEFAULT_HEALTH_CHECK\n        value: \"false\"\n      - name: NODE_LABEL\n        value: \"\"\n      - name: SEQUENCE\n        value: parallel\n    labels:\n      name: pod-delete\n      app.kubernetes.io/part-of: litmus\n      app.kubernetes.io/component: experiment-job\n      app.kubernetes.io/version: ci\n"
							}
						}
					]
				},
				"outputs": {},
				"metadata": {},
				"container": {
					"name": "",
					"image": "chaosnative/k8s:2.11.0",
					"command": [
						"sh",
						"-c"
					],
					"args": [
						"kubectl apply -f /tmp/ -n {{workflow.parameters.adminModeNamespace}} && sleep 30"
					],
					"resources": {}
				}
			},
			{
				"name": "cleanup-chaos-resources",
				"inputs": {},
				"outputs": {},
				"metadata": {},
				"container": {
					"name": "",
					"image": "chaosnative/k8s:2.11.0",
					"command": [
						"sh",
						"-c"
					],
					"args": [
						"kubectl delete chaosengine -l workflow_run_id={{workflow.uid}} -n {{workflow.parameters.adminModeNamespace}}"
					],
					"resources": {}
				}
			},
			{
				"name": "pod-delete-ji5",
				"inputs": {
					"artifacts": [
						{
							"name": "pod-delete-ji5",
							"path": "/tmp/chaosengine-pod-delete-ji5.yaml",
							"raw": {
								"data": "apiVersion: litmuschaos.io\/v1alpha1\r\nkind: ChaosEngine\r\nmetadata:\r\n  namespace: \"{{workflow.parameters.adminModeNamespace}}\"\r\n  generateName: pod-delete-ji5\r\n  labels:\r\n    workflow_run_id: \"{{ workflow.uid }}\"\r\n    workflow_name: my-pod-delete-experiment\r\nspec:\r\n  appinfo:\r\n    appns: hce\r\n    applabel: app=cartservice\r\n    appkind: deployment\r\n  engineState: active\r\n  chaosServiceAccount: litmus-admin\r\n  experiments:\r\n    - name: pod-delete\r\n      spec:\r\n        components:\r\n          env:\r\n            - name: TOTAL_CHAOS_DURATION\r\n              value: \"30\"\r\n            - name: CHAOS_INTERVAL\r\n              value: \"10\"\r\n            - name: FORCE\r\n              value: \"false\"\r\n            - name: PODS_AFFECTED_PERC\r\n              value: \"\"\r\n        probe:\r\n          - name: http-cartservice-probe\r\n            type: httpProbe\r\n            mode: Continuous\r\n            runProperties:\r\n              probeTimeout: 500\r\n              retry: 1\r\n              interval: 1\r\n              stopOnFailure: false\r\n            httpProbe\/inputs:\r\n              url: http:\/\/frontend\/cart\r\n              method:\r\n                get:\r\n                  criteria: ==\r\n                  responseCode: \"200\"\r\n                  responseTimeout: 15\r\n"
							}
						}
					]
				},
				"outputs": {},
				"metadata": {
					"labels": {
						"weight": "10"
					}
				},
				"container": {
					"name": "",
					"image": "chaosnative/litmus-checker:2.11.0",
					"args": [
						"-file=/tmp/chaosengine-pod-delete-ji5.yaml",
						"-saveName=/tmp/engine-name"
					],
					"resources": {}
				}
			}
		],
		"entrypoint": "test-0-7-x-exp",
		"arguments": {
			"parameters": [
				{
					"name": "adminModeNamespace",
					"value": "hce"
				}
			]
		},
		"serviceAccountName": "argo-chaos",
		"podGC": {
			"strategy": "OnWorkflowCompletion"
		},
		"securityContext": {
			"runAsUser": 1000,
			"runAsNonRoot": true
		}
	},
	"status": {
		"startedAt": null,
		"finishedAt": null
	}
}
`
