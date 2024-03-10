package connector

import "time"

type ResponseGetJob struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata   struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		Uid               string    `json:"uid"`
		ResourceVersion   string    `json:"resourceVersion"`
		Generation        int       `json:"generation"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			ControllerUid string `json:"controller-uid"`
			JobName       string `json:"job-name"`
		} `json:"labels"`
		Annotations struct {
			BatchKubernetesIoJobTracking string `json:"batch.kubernetes.io/job-tracking"`
		} `json:"annotations"`
		ManagedFields []struct {
			Manager    string    `json:"manager"`
			Operation  string    `json:"operation"`
			ApiVersion string    `json:"apiVersion"`
			Time       time.Time `json:"time"`
			FieldsType string    `json:"fieldsType"`
			FieldsV1   struct {
				FSpec struct {
					FBackoffLimit struct {
					} `json:"f:backoffLimit"`
					FCompletionMode struct {
					} `json:"f:completionMode"`
					FCompletions struct {
					} `json:"f:completions"`
					FParallelism struct {
					} `json:"f:parallelism"`
					FSuspend struct {
					} `json:"f:suspend"`
					FTemplate struct {
						FSpec struct {
							FContainers struct {
								KNameActivity077 struct {
									Field1 struct {
									} `json:"."`
									FCommand struct {
									} `json:"f:command"`
									FImage struct {
									} `json:"f:image"`
									FImagePullPolicy struct {
									} `json:"f:imagePullPolicy"`
									FName struct {
									} `json:"f:name"`
									FResources struct {
									} `json:"f:resources"`
									FTerminationMessagePath struct {
									} `json:"f:terminationMessagePath"`
									FTerminationMessagePolicy struct {
									} `json:"f:terminationMessagePolicy"`
								} `json:"k:{"name":"activity-077"}"`
							} `json:"f:containers"`
							FDnsPolicy struct {
							} `json:"f:dnsPolicy"`
							FRestartPolicy struct {
							} `json:"f:restartPolicy"`
							FSchedulerName struct {
							} `json:"f:schedulerName"`
							FSecurityContext struct {
							} `json:"f:securityContext"`
							FTerminationGracePeriodSeconds struct {
							} `json:"f:terminationGracePeriodSeconds"`
						} `json:"f:spec"`
					} `json:"f:template"`
				} `json:"f:spec,omitempty"`
				FStatus struct {
					FCompletionTime struct {
					} `json:"f:completionTime"`
					FConditions struct {
					} `json:"f:conditions"`
					FReady struct {
					} `json:"f:ready"`
					FStartTime struct {
					} `json:"f:startTime"`
					FSucceeded struct {
					} `json:"f:succeeded"`
					FUncountedTerminatedPods struct {
					} `json:"f:uncountedTerminatedPods"`
				} `json:"f:status,omitempty"`
			} `json:"fieldsV1"`
			Subresource string `json:"subresource,omitempty"`
		} `json:"managedFields"`
	} `json:"metadata"`
	Spec struct {
		Parallelism  int `json:"parallelism"`
		Completions  int `json:"completions"`
		BackoffLimit int `json:"backoffLimit"`
		Selector     struct {
			MatchLabels struct {
				ControllerUid string `json:"controller-uid"`
			} `json:"matchLabels"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				CreationTimestamp interface{} `json:"creationTimestamp"`
				Labels            struct {
					ControllerUid string `json:"controller-uid"`
					JobName       string `json:"job-name"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Name      string   `json:"name"`
					Image     string   `json:"image"`
					Command   []string `json:"command"`
					Resources struct {
					} `json:"resources"`
					TerminationMessagePath   string `json:"terminationMessagePath"`
					TerminationMessagePolicy string `json:"terminationMessagePolicy"`
					ImagePullPolicy          string `json:"imagePullPolicy"`
				} `json:"containers"`
				RestartPolicy                 string `json:"restartPolicy"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				DnsPolicy                     string `json:"dnsPolicy"`
				SecurityContext               struct {
				} `json:"securityContext"`
				SchedulerName string `json:"schedulerName"`
			} `json:"spec"`
		} `json:"template"`
		CompletionMode string `json:"completionMode"`
		Suspend        bool   `json:"suspend"`
	} `json:"spec"`
	Status struct {
		Conditions []struct {
			Type               string    `json:"type"`
			Status             string    `json:"status"`
			LastProbeTime      time.Time `json:"lastProbeTime"`
			LastTransitionTime time.Time `json:"lastTransitionTime"`
		} `json:"conditions"`
		StartTime               time.Time `json:"startTime"`
		CompletionTime          time.Time `json:"completionTime"`
		Succeeded               int       `json:"succeeded"`
		UncountedTerminatedPods struct {
		} `json:"uncountedTerminatedPods"`
		Ready  int `json:"ready"`
		Active int `json:"active"`
	} `json:"status"`
}
