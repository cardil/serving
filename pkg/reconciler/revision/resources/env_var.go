/*
Copyright 2018 The Knative Authors

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

package resources

import (
	corev1 "k8s.io/api/core/v1"
	"knative.dev/serving/pkg/apis/serving"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

const (
	knativeRevisionEnvVariableKey      = "K_REVISION"
	knativeConfigurationEnvVariableKey = "K_CONFIGURATION"
	knativeServiceEnvVariableKey       = "K_SERVICE"
)

func getKnativeEnvVar(rev *v1alpha1.Revision) []corev1.EnvVar {
	return []corev1.EnvVar{{
		Name:  knativeRevisionEnvVariableKey,
		Value: rev.Name,
	}, {
		Name:  knativeConfigurationEnvVariableKey,
		Value: rev.Labels[serving.ConfigurationLabelKey],
	}, {
		Name:  knativeServiceEnvVariableKey,
		Value: rev.Labels[serving.ServiceLabelKey],
	}}
}
