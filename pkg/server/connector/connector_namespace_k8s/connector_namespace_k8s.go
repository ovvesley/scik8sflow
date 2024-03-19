package connector_namespace_k8s

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type ConnectorNamespaceK8s struct {
	client *http.Client
}

type IConnectorNamespace interface {
	ListNamespaces()
	GetNamespace(namespace string) (ResponseGetNamespace, error)
	CreateNamespace(namespace string) (ResponseCreateNamespace, error)
}

func New() IConnectorNamespace {
	return &ConnectorNamespaceK8s{
		client: newClient(),
	}
}
func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func (c ConnectorNamespaceK8s) ListNamespaces() {
	//TODO implement me
	panic("implement me")
}

type ResponseGetNamespaceMetadata struct {
	Name              string    `json:"name"`
	Uid               string    `json:"uid"`
	ResourceVersion   string    `json:"resourceVersion"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Labels            struct {
		KubernetesIoMetadataName string `json:"kubernetes.io/metadata.name"`
	} `json:"labels"`
	Annotations struct {
		KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
	} `json:"annotations"`
	ManagedFields []struct {
		Manager    string    `json:"manager"`
		Operation  string    `json:"operation"`
		ApiVersion string    `json:"apiVersion"`
		Time       time.Time `json:"time"`
		FieldsType string    `json:"fieldsType"`
		FieldsV1   struct {
			FMetadata struct {
				FAnnotations struct {
					Field1 struct {
					} `json:"."`
					FKubectlKubernetesIoLastAppliedConfiguration struct {
					} `json:"f:kubectl.kubernetes.io/last-applied-configuration"`
				} `json:"f:annotations"`
				FLabels struct {
					Field1 struct {
					} `json:"."`
					FKubernetesIoMetadataName struct {
					} `json:"f:kubernetes.io/metadata.name"`
				} `json:"f:labels"`
			} `json:"f:metadata"`
		} `json:"fieldsV1"`
	} `json:"managedFields"`
}
type ResponseGetNamespace struct {
	Kind       string                       `json:"kind"`
	ApiVersion string                       `json:"apiVersion"`
	Metadata   ResponseGetNamespaceMetadata `json:"metadata"`
	Spec       struct {
		Finalizers []string `json:"finalizers"`
	} `json:"spec"`
	Status struct {
		Phase string `json:"phase"`
	} `json:"status"`
}

func (c *ConnectorNamespaceK8s) GetNamespace(namespace string) (ResponseGetNamespace, error) {
	token := os.Getenv("K8S_API_SERVER_TOKEN")
	host := os.Getenv("K8S_API_SERVER_HOST")

	req, _ := http.NewRequest("GET", "https://"+host+"/api/v1/namespaces/"+namespace, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)

	if err != nil {
		return ResponseGetNamespace{}, err
	}

	defer resp.Body.Close()

	var result ResponseGetNamespace
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return ResponseGetNamespace{}, err
	}

	return result, nil
}

type ResponseCreateNamespace struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata   struct {
		Name              string    `json:"name"`
		Uid               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			KubernetesIoMetadataName string `json:"kubernetes.io/metadata.name"`
		} `json:"labels"`
		ManagedFields []struct {
			Manager    string    `json:"manager"`
			Operation  string    `json:"operation"`
			ApiVersion string    `json:"apiVersion"`
			Time       time.Time `json:"time"`
			FieldsType string    `json:"fieldsType"`
			FieldsV1   struct {
				FMetadata struct {
					FLabels struct {
						Field1 struct {
						} `json:"."`
						FKubernetesIoMetadataName struct {
						} `json:"f:kubernetes.io/metadata.name"`
					} `json:"f:labels"`
				} `json:"f:metadata"`
			} `json:"fieldsV1"`
		} `json:"managedFields"`
	} `json:"metadata"`
	Spec struct {
		Finalizers []string `json:"finalizers"`
	} `json:"spec"`
	Status struct {
		Phase string `json:"phase"`
	} `json:"status"`
}

func (c *ConnectorNamespaceK8s) CreateNamespace(namespace string) (ResponseCreateNamespace, error) {
	token := os.Getenv("K8S_API_SERVER_TOKEN")
	host := os.Getenv("K8S_API_SERVER_HOST")

	body := []byte(`{"apiVersion": "v1", "kind": "Namespace", "metadata": {"name": "` + namespace + `"}}`)

	req, _ := http.NewRequest("POST", "https://"+host+"/api/v1/namespaces", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	if err != nil {
		return ResponseCreateNamespace{}, err
	}

	defer resp.Body.Close()

	var result ResponseCreateNamespace
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return ResponseCreateNamespace{}, err
	}

	return result, nil

}
