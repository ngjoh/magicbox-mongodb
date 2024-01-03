package stores

import (
	"encoding/json"
	"time"

	"github.com/koksmat-com/koksmat/connectors"
)

func PerconaCRDS() ([]Store, error) {
	type CRDS struct {
		APIVersion string `json:"apiVersion"`
		Items      []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				Annotations struct {
					KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
				} `json:"annotations"`
				CreationTimestamp time.Time `json:"creationTimestamp"`
				Finalizers        []string  `json:"finalizers"`
				Generation        int       `json:"generation"`
				Name              string    `json:"name"`
				Namespace         string    `json:"namespace"`
				ResourceVersion   string    `json:"resourceVersion"`
				UID               string    `json:"uid"`
			} `json:"metadata,omitempty"`
			Spec struct {
				AllowUnsafeConfigurations bool `json:"allowUnsafeConfigurations"`
				Backup                    struct {
					Enabled bool   `json:"enabled"`
					Image   string `json:"image"`
					Pitr    struct {
						CompressionLevel int    `json:"compressionLevel"`
						CompressionType  string `json:"compressionType"`
						Enabled          bool   `json:"enabled"`
					} `json:"pitr"`
					ServiceAccountName string `json:"serviceAccountName"`
					Storages           struct {
						AzureBlob struct {
							Azure struct {
								Container         string `json:"container"`
								CredentialsSecret string `json:"credentialsSecret"`
								Prefix            string `json:"prefix"`
							} `json:"azure"`
							Type string `json:"type"`
						} `json:"azure-blob"`
					} `json:"storages"`
				} `json:"backup"`
				CrVersion       string `json:"crVersion"`
				Image           string `json:"image"`
				ImagePullPolicy string `json:"imagePullPolicy"`
				Pmm             struct {
					Enabled    bool   `json:"enabled"`
					Image      string `json:"image"`
					ServerHost string `json:"serverHost"`
				} `json:"pmm"`
				Replsets []struct {
					Affinity struct {
						AntiAffinityTopologyKey string `json:"antiAffinityTopologyKey"`
					} `json:"affinity"`
					Arbiter struct {
						Affinity struct {
							AntiAffinityTopologyKey string `json:"antiAffinityTopologyKey"`
						} `json:"affinity"`
						Enabled bool `json:"enabled"`
						Size    int  `json:"size"`
					} `json:"arbiter"`
					Expose struct {
						Enabled    bool   `json:"enabled"`
						ExposeType string `json:"exposeType"`
					} `json:"expose"`
					Name      string `json:"name"`
					Nonvoting struct {
						Affinity struct {
							AntiAffinityTopologyKey string `json:"antiAffinityTopologyKey"`
						} `json:"affinity"`
						Enabled             bool `json:"enabled"`
						PodDisruptionBudget struct {
							MaxUnavailable int `json:"maxUnavailable"`
						} `json:"podDisruptionBudget"`
						Resources struct {
							Limits struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						Size       int `json:"size"`
						VolumeSpec struct {
							PersistentVolumeClaim struct {
								Resources struct {
									Requests struct {
										Storage string `json:"storage"`
									} `json:"requests"`
								} `json:"resources"`
							} `json:"persistentVolumeClaim"`
						} `json:"volumeSpec"`
					} `json:"nonvoting"`
					PodDisruptionBudget struct {
						MaxUnavailable int `json:"maxUnavailable"`
					} `json:"podDisruptionBudget"`
					Resources struct {
						Limits struct {
							CPU    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"limits"`
						Requests struct {
							CPU    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"requests"`
					} `json:"resources"`
					Size       int `json:"size"`
					VolumeSpec struct {
						PersistentVolumeClaim struct {
							Resources struct {
								Requests struct {
									Storage string `json:"storage"`
								} `json:"requests"`
							} `json:"resources"`
						} `json:"persistentVolumeClaim"`
					} `json:"volumeSpec"`
				} `json:"replsets"`
				Secrets struct {
					EncryptionKey string `json:"encryptionKey"`
					Users         string `json:"users"`
				} `json:"secrets"`
				Sharding struct {
					ConfigsvrReplSet struct {
						Affinity struct {
							AntiAffinityTopologyKey string `json:"antiAffinityTopologyKey"`
						} `json:"affinity"`
						Expose struct {
							Enabled    bool   `json:"enabled"`
							ExposeType string `json:"exposeType"`
						} `json:"expose"`
						PodDisruptionBudget struct {
							MaxUnavailable int `json:"maxUnavailable"`
						} `json:"podDisruptionBudget"`
						Resources struct {
							Limits struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						Size       int `json:"size"`
						VolumeSpec struct {
							PersistentVolumeClaim struct {
								Resources struct {
									Requests struct {
										Storage string `json:"storage"`
									} `json:"requests"`
								} `json:"resources"`
							} `json:"persistentVolumeClaim"`
						} `json:"volumeSpec"`
					} `json:"configsvrReplSet"`
					Enabled bool `json:"enabled"`
					Mongos  struct {
						Affinity struct {
							AntiAffinityTopologyKey string `json:"antiAffinityTopologyKey"`
						} `json:"affinity"`
						Expose struct {
							ExposeType string `json:"exposeType"`
						} `json:"expose"`
						PodDisruptionBudget struct {
							MaxUnavailable int `json:"maxUnavailable"`
						} `json:"podDisruptionBudget"`
						Resources struct {
							Limits struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						Size int `json:"size"`
					} `json:"mongos"`
				} `json:"sharding"`
				UpdateStrategy string `json:"updateStrategy"`
				UpgradeOptions struct {
					Apply                  string `json:"apply"`
					Schedule               string `json:"schedule"`
					SetFCV                 bool   `json:"setFCV"`
					VersionServiceEndpoint string `json:"versionServiceEndpoint"`
				} `json:"upgradeOptions"`
			} `json:"spec"`
			Status struct {
				Conditions []struct {
					LastTransitionTime time.Time `json:"lastTransitionTime"`
					Status             string    `json:"status"`
					Type               string    `json:"type"`
					Reason             string    `json:"reason,omitempty"`
				} `json:"conditions"`
				Host         string `json:"host"`
				MongoImage   string `json:"mongoImage"`
				MongoVersion string `json:"mongoVersion"`
				Mongos       struct {
					Ready  int    `json:"ready"`
					Size   int    `json:"size"`
					Status string `json:"status"`
				} `json:"mongos"`
				ObservedGeneration int `json:"observedGeneration"`
				Ready              int `json:"ready"`
				Replsets           struct {
					Cfg struct {
						Initialized bool   `json:"initialized"`
						Ready       int    `json:"ready"`
						Size        int    `json:"size"`
						Status      string `json:"status"`
					} `json:"cfg"`
					Rs0 struct {
						AddedAsShard bool   `json:"added_as_shard"`
						Initialized  bool   `json:"initialized"`
						Ready        int    `json:"ready"`
						Size         int    `json:"size"`
						Status       string `json:"status"`
					} `json:"rs0"`
				} `json:"replsets"`
				Size  int    `json:"size"`
				State string `json:"state"`
			} `json:"status"`
		} `json:"items"`
		Kind     string `json:"kind"`
		Metadata struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"metadata"`
	}

	bytes, err := connectors.Execute("kubectl", *&connectors.Options{}, "get", "psmdb", "-o", "json")
	if err != nil {
		return nil, err
	}
	stores := make([]Store, 0)
	crds := CRDS{}

	err = json.Unmarshal(bytes, &crds)
	if err != nil {
		return nil, err
	}
	mateContext, err := connectors.GetMateContext()
	if err != nil {
		return nil, err
	}
	for _, crd := range crds.Items {
		connector := Store{
			Name:        crd.Metadata.Name,
			Description: crd.Status.State,
			IsCurrent:   crd.Metadata.Name == mateContext.Current.Mongo,
		}
		stores = append(stores, connector)
	}

	return stores, nil
}

func Databases() ([]Store, error) {
	stores := make([]Store, 0)

	return stores, nil
}
