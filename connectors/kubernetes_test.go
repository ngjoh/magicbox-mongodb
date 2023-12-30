package connectors

import (
	"testing"
)

func TestKubernetesClusters(t *testing.T) {

	k, err := KubernetesClusters()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}

func TestKubernetesNamespaces(t *testing.T) {

	k, err := KubernetesNamespaces()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}

func TestKubernetesPods(t *testing.T) {

	k, err := KubernetesPods()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
