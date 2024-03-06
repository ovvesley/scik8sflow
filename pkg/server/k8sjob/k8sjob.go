package k8sjob

import "gopkg.in/yaml.v3"

type K8sJob struct {
	ApiVersion string         `json:"apiVersion"`
	Kind       string         `json:"kind"`
	Metadata   K8sJobMetadata `json:"metadata"`
	Spec       K8sJobSpec     `json:"spec"`
}

type K8sJobMetadata struct {
	Name string `json:"name"`
}

type K8sJobSpec struct {
	Template K8sJobTemplate `json:"template"`
}

type K8sJobTemplate struct {
	Spec K8sJobSpecTemplate `json:"spec"`
}

type K8sJobSpecTemplate struct {
	Containers    []K8sJobContainer `json:"containers"`
	RestartPolicy string            `json:"restartPolicy"`
}

type K8sJobContainer struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

func (k *K8sJob) ToYaml() string {
	workflowStringByte, _ := yaml.Marshal(k)
	workflowStringYaml := string(workflowStringByte)
	return workflowStringYaml
}

// docker run --rm alpine:latest bin/sh -c 'echo ZWNobyAiSGVsbG8gV29ybGQiCnNsZWVwIDUKZWNobyAiSGVsbG8gV29ybGQgQWdhaW4iCnNsZWVwIDUKZWNobyAiSGVsbG8gV29ybGQgT25lIE1vcmUgVGltZSI=| base64 -d| sh'
