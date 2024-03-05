/*
Copyright 2022 The Kubernetes Authors.

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

package v1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

// Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations is an autogenerated conversion function.
func Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(in, out, s)
}

// Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations is an autogenerated conversion function.
func Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in, out, s)
}
