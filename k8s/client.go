/*
Copyright 2022 The Inspur AIStation Group Authors.

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

// Note: the example only works with the code within the same release/branch.
package k8s

import (
	"github.com/cauwulixuan/go-kit/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var Client *kubernetes.Clientset

func Init(kubePath string) {
	GenerateClientSet(kubePath)
}

func GetK8sConfig(kubePath string) (*rest.Config, error) {
	var kubeconfig string
	if kubePath == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			kubeconfig = ""
		}
	}
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()

	// use the current context in kubeconfig
	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func GenerateClientSet(kubePath string) {
	config, err := GetK8sConfig(kubePath)
	if err != nil {
		log.Slogger.Error(err.Error())
	}

	// create the clientset
	Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Slogger.Error(err.Error())
	}
}
