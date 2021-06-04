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
	"github.com/gin-gonic/gin"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	_ "strings"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		router := gin.Default()
		router.GET("/api/pod/", func(t *gin.Context) {
			//namespaces := t.Param("name")
			//app := t.Param("name1")
			namespaces := t.DefaultQuery("name", "")
			app := t.DefaultQuery("name1", "")
			config1 := "/root/.kube/config"

			if namespaces == "" && app == "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Pods(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}
				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s    %s\n", d.Namespace, d.Name, d.Status.Phase, d.Spec.NodeName, d.Status.PodIP)
					t.String(http.StatusOK, string(s))

				}

			} else if namespaces != "" && app == "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Pods(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}
				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s    %s\n", d.Namespace, d.Name, d.Status.Phase, d.Spec.NodeName, d.Status.PodIP)
					t.String(http.StatusOK, string(s))

				}

			} else if namespaces != "" && app != "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Pods(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}

				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s    %s\n", d.Namespace, d.Name, d.Status.Phase, d.Spec.NodeName, d.Status.PodIP)
					b := []string{s}
					f := fuzzy.Find(app, b)
					for _, i := range f {
						//k := fmt.Sprintf("f", i)
						t.String(http.StatusOK, string(i))
					}

				}
			} else {
				fmt.Printf("please input namespaces")

			}

		})
		router.GET("/api/svc", func(t *gin.Context) {
			namespaces := t.DefaultQuery("name", "")
			app := t.DefaultQuery("name1", "")
			config1 := "/root/.kube/config"

			if namespaces == "" && app == "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Services(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}
				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s    %s   %#v\n", d.Namespace, d.Name, d.Spec.Type, d.Spec.ClusterIP, d.Spec.ExternalIPs, d.Spec.Ports)
					t.String(http.StatusOK, string(s))

				}

			} else if namespaces != "" && app == "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Services(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}
				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s   %s   %#v\n", d.Namespace, d.Name, d.Spec.Type, d.Spec.ClusterIP, d.Spec.ExternalIPs, d.Spec.Ports)
					t.String(http.StatusOK, string(s))

				}

			} else if namespaces != "" && app != "" {
				dash1, err := clientcmd.BuildConfigFromFlags("", config1)
				if err != nil {
					panic(err)
				}
				clientset, err := kubernetes.NewForConfig(dash1)

				if err != nil {
					panic(err)
				}

				allpods := clientset.CoreV1().Services(namespaces)
				list1, err := allpods.List(context.TODO(), metav1.ListOptions{})

				if err != nil {
					panic(err)
				}

				for _, d := range list1.Items {

					s := fmt.Sprintf("   %s    %s    %s    %s   %s   %#v\n", d.Namespace, d.Name, d.Spec.Type, d.Spec.ClusterIP, d.Spec.ExternalIPs, d.Spec.Ports)
					b := []string{s}
					f := fuzzy.Find(app, b)
					for _, i := range f {
						//k := fmt.Sprintf("f", i)
						t.String(http.StatusOK, string(i))
					}

				}
			} else {
				fmt.Printf("please input namespaces")

			}

		})
		router.Run(":8000")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
