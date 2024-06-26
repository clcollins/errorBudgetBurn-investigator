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

package consoleErrorBudgetBurn

import (
	"fmt"

	"github.com/clcollins/errorBudgetBurn-investigator/pkg/common"
	"k8s.io/client-go/dynamic"
)

func Run(kubeClient *dynamic.DynamicClient, verbose bool) error {
	return consoleErrorBudgetBurn(kubeClient, verbose)
}

func consoleErrorBudgetBurn(kubeClient *dynamic.DynamicClient, verbose bool) error {
	if verbose {
		fmt.Println("Investigating console-errorBudgetBurn")
	}

	err := common.CheckDefaultIngress(kubeClient, verbose)
	if err != nil {
		return err
	}

	return nil
}
