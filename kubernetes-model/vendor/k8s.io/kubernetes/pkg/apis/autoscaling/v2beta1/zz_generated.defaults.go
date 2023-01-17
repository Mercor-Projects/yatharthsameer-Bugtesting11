/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// +build !ignore_autogenerated

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v2beta1

import (
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v2beta1.HorizontalPodAutoscaler{}, func(obj interface{}) {
		SetObjectDefaults_HorizontalPodAutoscaler(obj.(*v2beta1.HorizontalPodAutoscaler))
	})
	scheme.AddTypeDefaultingFunc(&v2beta1.HorizontalPodAutoscalerList{}, func(obj interface{}) {
		SetObjectDefaults_HorizontalPodAutoscalerList(obj.(*v2beta1.HorizontalPodAutoscalerList))
	})
	return nil
}

func SetObjectDefaults_HorizontalPodAutoscaler(in *v2beta1.HorizontalPodAutoscaler) {
	SetDefaults_HorizontalPodAutoscaler(in)
}

func SetObjectDefaults_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_HorizontalPodAutoscaler(a)
	}
}
