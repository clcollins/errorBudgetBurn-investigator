// Copyright 2021-2024 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
