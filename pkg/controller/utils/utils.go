/*
Copyright 2018 The Kubernetes Authors.

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

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/clientcmd"

	capiv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	openshiftapiv1 "github.com/openshift/api/config/v1"
	routev1 "github.com/openshift/api/route/v1"
)

// BuildClusterAPIClientFromKubeconfig will return a kubeclient usin the provided kubeconfig
func BuildClusterAPIClientFromKubeconfig(kubeconfigData string) (client.Client, error) {
	config, err := clientcmd.Load([]byte(kubeconfigData))
	if err != nil {
		return nil, err
	}
	kubeConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{})
	cfg, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	scheme, err := capiv1.SchemeBuilder.Build()
	if err != nil {
		return nil, err
	}

	if err := openshiftapiv1.Install(scheme); err != nil {
		return nil, err
	}

	if err := routev1.Install(scheme); err != nil {
		return nil, err
	}

	return client.New(cfg, client.Options{
		Scheme: scheme,
	})
}

// FixupEmptyClusterVersionFields will un-'nil' fields that would fail validation in the ClusterVersion.Status
func FixupEmptyClusterVersionFields(clusterVersionStatus *openshiftapiv1.ClusterVersionStatus) {

	// Fetching clusterVersion object can result in nil clusterVersion.Status.AvailableUpdates,
	// clusterVersion.Status.Conditions, and clusterVersion.Status.History
	// Place an empty list if needed to satisfy the object validation.

	if clusterVersionStatus.AvailableUpdates == nil {
		clusterVersionStatus.AvailableUpdates = []openshiftapiv1.Update{}
	}

	if clusterVersionStatus.Conditions == nil {
		clusterVersionStatus.Conditions = []openshiftapiv1.ClusterOperatorStatusCondition{}
	}

	if clusterVersionStatus.History == nil {
		clusterVersionStatus.History = []openshiftapiv1.UpdateHistory{}
	}
}

// HasFinalizer returns true if the given object has the given finalizer
func HasFinalizer(object metav1.Object, finalizer string) bool {
	for _, f := range object.GetFinalizers() {
		if f == finalizer {
			return true
		}
	}
	return false
}

// AddFinalizer adds a finalizer to the given object
func AddFinalizer(object metav1.Object, finalizer string) {
	finalizers := sets.NewString(object.GetFinalizers()...)
	finalizers.Insert(finalizer)
	object.SetFinalizers(finalizers.List())
}

// DeleteFinalizer removes a finalizer from the given object
func DeleteFinalizer(object metav1.Object, finalizer string) {
	finalizers := sets.NewString(object.GetFinalizers()...)
	finalizers.Delete(finalizer)
	object.SetFinalizers(finalizers.List())
}
