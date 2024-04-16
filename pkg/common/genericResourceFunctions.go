package common

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// Generic k8s and OpenShift API interaction functions/helpers that can be used by multiple checks

// Resource schema for the dynamic client
var podResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

// getPods takes a *dynamic.DynamicClient, a namespace, and a metav1.ListOptions and returns a corev1.PodList from the cluster
func getPods(kubeClient *dynamic.DynamicClient, namespace string, listOpts metav1.ListOptions) (corev1.PodList, error) {
	var pods corev1.PodList

	resp, err := kubeClient.Resource(podResource).Namespace(namespace).List(context.TODO(), listOpts)
	if err != nil {
		return pods, err
	}

	unstructured := resp.UnstructuredContent()

	err = runtime.DefaultUnstructuredConverter.
		FromUnstructured(unstructured, &pods)

	if err != nil {
		return pods, err
	}

	return pods, nil
}
