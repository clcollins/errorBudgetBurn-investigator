package common

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
)

// Common checks that might be used by a variety of errorBudgetBurn alerts
// Eg: api- and console-errorBugetBurn alerts might both check the status of
// the default ingress

const (
	defaultIngressNamespace          = "openshift-ingress"
	defaultIngressPodControllerLabel = "ingresscontroller.operator.openshift.io/deployment-ingresscontroller=default"

	templateNamespaceAndName = "{{.Namespace}}/{{.Name}}"
)

// Check DefaultIngress is an aggregator for the various different checks that would
// validate the health of the default ingress
func CheckDefaultIngress(kubeClient *dynamic.DynamicClient, verbose bool) error {

	err := checkDefaultIngressRouterPods(kubeClient, verbose)
	if err != nil {
		return err
	}

	return nil
}

// checkDefaultIngressRouterPods is an aggregator for the various different checks that would
// validate the health of the default ingress router pods themselves
func checkDefaultIngressRouterPods(kubeClient *dynamic.DynamicClient, verbose bool) error {
	if verbose {
		fmt.Println("Checking default ingress router pods")
	}

	var listOpts metav1.ListOptions
	listOpts.LabelSelector = defaultIngressPodControllerLabel

	pods, err := getPods(kubeClient, defaultIngressNamespace, listOpts)
	if err != nil {
		return err
	}

	_, err = podsInRunningPhase(pods)
	if err != nil {
		return err
	}

	return nil
}

// podsInRunningPhase checks to make sure all the pods in the podlist are running (or creating) and
// returns false and an error if they are in any other phase (eg: Pending, Terminating)
func podsInRunningPhase(pods corev1.PodList) (bool, error) {
	for _, pod := range pods.Items {
		switch pod.Status.Phase {
		case "Running":
			return true, nil
		case "Creating":
			return true, nil
		default:
			return false, fmt.Errorf("%s pod/%s is not running: %s", pod.Namespace, pod.Name, pod.Status.Phase)
		}
	}
	return true, nil
}
