/*
Copyright © 2024 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clcollins/errorBudgetBurn-investigator/pkg/apiErrorBudgetBurn"
	"github.com/clcollins/errorBudgetBurn-investigator/pkg/consoleErrorBudgetBurn"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var supportedAlerts = map[string]func(*dynamic.DynamicClient, bool) error{
	"console": consoleErrorBudgetBurn.Run,
	"api":     apiErrorBudgetBurn.Run,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "errorBudgetBurn-investigator",
	Short: "A tool for investigating various errorBudgetBurn alerts",
	Long: `
A tool for investigating various errorBudgetBurn alerts received by SREs, including:

* console-errorBudgetBurn
* api-errorBudgetBurn
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := viper.GetBool("verbose")

		kubeClient, err := newKubeDynamicClient()
		if err != nil {
			fmt.Print(err)
			return
		}

		for k := range supportedAlerts {
			if viper.GetBool(k) {
				err := supportedAlerts[k](kubeClient, verbose)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("verbose", false, "Enable verbose output")

	rootCmd.Flags().Bool("console", false, "Investigate console-errorBudgetBurn alerts")
	rootCmd.Flags().Bool("api", false, "Investigate api-errorBudgetBurn alerts")

	rootCmd.MarkFlagsMutuallyExclusive("console", "api")
	rootCmd.MarkFlagsOneRequired("console", "api")

	err := viper.BindPFlags(rootCmd.Flags())
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

// newKubeDynamicClient builds a kube client for the currently logged-in cluster
func newKubeDynamicClient() (*dynamic.DynamicClient, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// kubeConfigLoader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
	// 	clientcmd.NewDefaultClientConfigLoadingRules(),
	// 	&clientcmd.ConfigOverrides{},
	// )

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", home+"/.kube/config")
	if err != nil {
		return nil, err
	}

	kubeClient, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return kubeClient, nil
}
