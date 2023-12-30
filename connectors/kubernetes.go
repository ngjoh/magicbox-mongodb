package connectors

import (
	"encoding/json"
	"time"
)

func KubernetesClusters() ([]Connector, error) {
	type Config struct {
		Kind        string `json:"kind"`
		APIVersion  string `json:"apiVersion"`
		Preferences struct {
		} `json:"preferences"`
		Clusters []struct {
			Name    string `json:"name"`
			Cluster struct {
				Server                   string `json:"server"`
				CertificateAuthorityData string `json:"certificate-authority-data"`
			} `json:"cluster"`
		} `json:"clusters"`
		Users []struct {
			Name string `json:"name"`
			User struct {
				ClientCertificateData string `json:"client-certificate-data"`
				ClientKeyData         string `json:"client-key-data"`
			} `json:"user"`
		} `json:"users"`
		Contexts []struct {
			Name    string `json:"name"`
			Context struct {
				Cluster   string `json:"cluster"`
				User      string `json:"user"`
				Namespace string `json:"namespace"`
			} `json:"context"`
		} `json:"contexts"`
		CurrentContext string `json:"current-context"`
	}

	bytes, err := Execute("kubectl", *&Options{}, "config", "view", "-o=json")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	config := Config{}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	for _, cluster := range config.Contexts {
		connector := Connector{
			Name:        cluster.Name,
			Description: "Current namespace: " + cluster.Context.Namespace,
			IsCurrent:   cluster.Name == config.CurrentContext,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func KubernetesNamespaces() ([]Connector, error) {
	type Namespace struct {
		APIVersion string `json:"apiVersion"`
		Items      []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				CreationTimestamp time.Time `json:"creationTimestamp"`
				Labels            struct {
					KubernetesIoMetadataName          string `json:"kubernetes.io/metadata.name"`
					KustomizeToolkitFluxcdIoName      string `json:"kustomize.toolkit.fluxcd.io/name"`
					KustomizeToolkitFluxcdIoNamespace string `json:"kustomize.toolkit.fluxcd.io/namespace"`
				} `json:"labels"`
				Name            string `json:"name"`
				ResourceVersion string `json:"resourceVersion"`
				UID             string `json:"uid"`
			} `json:"metadata"`
			Spec struct {
				Finalizers []string `json:"finalizers"`
			} `json:"spec"`
			Status struct {
				Phase string `json:"phase"`
			} `json:"status"`
		} `json:"items"`
		Kind     string `json:"kind"`
		Metadata struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"metadata"`
	}

	bytes, err := Execute("kubectl", *&Options{}, "get", "namespaces", "-o=json")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	namespaces := Namespace{}

	err = json.Unmarshal(bytes, &namespaces)
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaces.Items {
		connector := Connector{
			Name:        namespace.Metadata.Name,
			Description: namespace.Status.Phase,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func KubernetesPods() ([]Connector, error) {
	type Pods struct {
		APIVersion string `json:"apiVersion"`
		Items      []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				Annotations struct {
					ChecksumTokenSecret string `json:"checksum/token-secret"`
				} `json:"annotations"`
				CreationTimestamp time.Time `json:"creationTimestamp"`
				GenerateName      string    `json:"generateName"`
				Labels            struct {
					AppKubernetesIoComponent       string `json:"app.kubernetes.io/component"`
					AppKubernetesIoInstance        string `json:"app.kubernetes.io/instance"`
					AppKubernetesIoManagedBy       string `json:"app.kubernetes.io/managed-by"`
					AppKubernetesIoName            string `json:"app.kubernetes.io/name"`
					AppKubernetesIoVersion         string `json:"app.kubernetes.io/version"`
					ControllerRevisionHash         string `json:"controller-revision-hash"`
					HelmShChart                    string `json:"helm.sh/chart"`
					StatefulsetKubernetesIoPodName string `json:"statefulset.kubernetes.io/pod-name"`
				} `json:"labels"`
				Name            string `json:"name"`
				Namespace       string `json:"namespace"`
				OwnerReferences []struct {
					APIVersion         string `json:"apiVersion"`
					BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
					Controller         bool   `json:"controller"`
					Kind               string `json:"kind"`
					Name               string `json:"name"`
					UID                string `json:"uid"`
				} `json:"ownerReferences"`
				ResourceVersion string `json:"resourceVersion"`
				UID             string `json:"uid"`
			} `json:"metadata"`
			Spec struct {
				Affinity struct {
					PodAntiAffinity struct {
						PreferredDuringSchedulingIgnoredDuringExecution []struct {
							PodAffinityTerm struct {
								LabelSelector struct {
									MatchLabels struct {
										AppKubernetesIoComponent string `json:"app.kubernetes.io/component"`
										AppKubernetesIoInstance  string `json:"app.kubernetes.io/instance"`
										AppKubernetesIoName      string `json:"app.kubernetes.io/name"`
									} `json:"matchLabels"`
								} `json:"labelSelector"`
								TopologyKey string `json:"topologyKey"`
							} `json:"podAffinityTerm"`
							Weight int `json:"weight"`
						} `json:"preferredDuringSchedulingIgnoredDuringExecution"`
					} `json:"podAntiAffinity"`
				} `json:"affinity"`
				Containers []struct {
					Env []struct {
						Name      string `json:"name"`
						Value     string `json:"value,omitempty"`
						ValueFrom struct {
							FieldRef struct {
								APIVersion string `json:"apiVersion"`
								FieldPath  string `json:"fieldPath"`
							} `json:"fieldRef"`
						} `json:"valueFrom,omitempty"`
					} `json:"env"`
					Image           string `json:"image"`
					ImagePullPolicy string `json:"imagePullPolicy"`
					LivenessProbe   struct {
						Exec struct {
							Command []string `json:"command"`
						} `json:"exec"`
						FailureThreshold    int `json:"failureThreshold"`
						InitialDelaySeconds int `json:"initialDelaySeconds"`
						PeriodSeconds       int `json:"periodSeconds"`
						SuccessThreshold    int `json:"successThreshold"`
						TimeoutSeconds      int `json:"timeoutSeconds"`
					} `json:"livenessProbe"`
					Name  string `json:"name"`
					Ports []struct {
						ContainerPort int    `json:"containerPort"`
						Name          string `json:"name"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
					ReadinessProbe struct {
						Exec struct {
							Command []string `json:"command"`
						} `json:"exec"`
						FailureThreshold    int `json:"failureThreshold"`
						InitialDelaySeconds int `json:"initialDelaySeconds"`
						PeriodSeconds       int `json:"periodSeconds"`
						SuccessThreshold    int `json:"successThreshold"`
						TimeoutSeconds      int `json:"timeoutSeconds"`
					} `json:"readinessProbe"`
					Resources struct {
					} `json:"resources"`
					SecurityContext struct {
						AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
						Capabilities             struct {
							Drop []string `json:"drop"`
						} `json:"capabilities"`
						Privileged     bool `json:"privileged"`
						RunAsNonRoot   bool `json:"runAsNonRoot"`
						RunAsUser      int  `json:"runAsUser"`
						SeccompProfile struct {
							Type string `json:"type"`
						} `json:"seccompProfile"`
					} `json:"securityContext"`
					TerminationMessagePath   string `json:"terminationMessagePath"`
					TerminationMessagePolicy string `json:"terminationMessagePolicy"`
					VolumeMounts             []struct {
						MountPath string `json:"mountPath"`
						Name      string `json:"name"`
						ReadOnly  bool   `json:"readOnly,omitempty"`
					} `json:"volumeMounts"`
				} `json:"containers"`
				DNSPolicy          string `json:"dnsPolicy"`
				EnableServiceLinks bool   `json:"enableServiceLinks"`
				Hostname           string `json:"hostname"`
				NodeName           string `json:"nodeName"`
				PreemptionPolicy   string `json:"preemptionPolicy"`
				Priority           int    `json:"priority"`
				RestartPolicy      string `json:"restartPolicy"`
				SchedulerName      string `json:"schedulerName"`
				SecurityContext    struct {
					FsGroup int `json:"fsGroup"`
				} `json:"securityContext"`
				ServiceAccount                string `json:"serviceAccount"`
				ServiceAccountName            string `json:"serviceAccountName"`
				Subdomain                     string `json:"subdomain"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				Tolerations                   []struct {
					Effect            string `json:"effect"`
					Key               string `json:"key"`
					Operator          string `json:"operator"`
					TolerationSeconds int    `json:"tolerationSeconds"`
				} `json:"tolerations"`
				Volumes []struct {
					Name                  string `json:"name"`
					PersistentVolumeClaim struct {
						ClaimName string `json:"claimName"`
					} `json:"persistentVolumeClaim,omitempty"`
					Secret struct {
						DefaultMode int    `json:"defaultMode"`
						SecretName  string `json:"secretName"`
					} `json:"secret,omitempty"`
					Projected struct {
						DefaultMode int `json:"defaultMode"`
						Sources     []struct {
							ServiceAccountToken struct {
								ExpirationSeconds int    `json:"expirationSeconds"`
								Path              string `json:"path"`
							} `json:"serviceAccountToken,omitempty"`
							ConfigMap struct {
								Items []struct {
									Key  string `json:"key"`
									Path string `json:"path"`
								} `json:"items"`
								Name string `json:"name"`
							} `json:"configMap,omitempty"`
							DownwardAPI struct {
								Items []struct {
									FieldRef struct {
										APIVersion string `json:"apiVersion"`
										FieldPath  string `json:"fieldPath"`
									} `json:"fieldRef"`
									Path string `json:"path"`
								} `json:"items"`
							} `json:"downwardAPI,omitempty"`
						} `json:"sources"`
					} `json:"projected,omitempty"`
				} `json:"volumes"`
			} `json:"spec"`
			Status struct {
				Conditions []struct {
					LastProbeTime      interface{} `json:"lastProbeTime"`
					LastTransitionTime time.Time   `json:"lastTransitionTime"`
					Status             string      `json:"status"`
					Type               string      `json:"type"`
				} `json:"conditions"`
				ContainerStatuses []struct {
					ContainerID string `json:"containerID"`
					Image       string `json:"image"`
					ImageID     string `json:"imageID"`
					LastState   struct {
					} `json:"lastState"`
					Name         string `json:"name"`
					Ready        bool   `json:"ready"`
					RestartCount int    `json:"restartCount"`
					Started      bool   `json:"started"`
					State        struct {
						Running struct {
							StartedAt time.Time `json:"startedAt"`
						} `json:"running"`
					} `json:"state"`
				} `json:"containerStatuses"`
				HostIP string `json:"hostIP"`
				Phase  string `json:"phase"`
				PodIP  string `json:"podIP"`
				PodIPs []struct {
					IP string `json:"ip"`
				} `json:"podIPs"`
				QosClass  string    `json:"qosClass"`
				StartTime time.Time `json:"startTime"`
			} `json:"status"`
		} `json:"items"`
		Kind     string `json:"kind"`
		Metadata struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"metadata"`
	}
	bytes, err := Execute("kubectl", *&Options{}, "get", "pods", "-o=json")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	pods := Pods{}

	err = json.Unmarshal(bytes, &pods)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		connector := Connector{
			Name:        pod.Metadata.Name,
			Description: pod.Status.Phase,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}
