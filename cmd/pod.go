/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var svc string

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pod called service")
		svc1, err := clientcmd.BuildConfigFromFlags("", svc)
		if err != nil {
			panic(err)
		}
		clientset, err := kubernetes.NewForConfig(svc1)

		if err != nil {
			panic(err)
		}
		Serviceclient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
		list, err := Serviceclient.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, d := range list.Items {

			fmt.Printf(" * %s %s %d %v+ %v+ %d)\n", d.Name, d.Namespace, d.Spec.Ports, d.Spec.ClusterIPs, d.TypeMeta, d.Spec.ClusterIP)
		}
	},
}

func init() {
	rootCmd.AddCommand(podCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	podCmd.Flags().StringVarP(&svc, "svc", "s", "", "testsvc")
}
