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
package integration

import (
	"testing"
	"time"

	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
	kapi "k8s.io/kubernetes/pkg/apis/core"

	testutil "github.com/openshift/origin/test/util"
	testserver "github.com/openshift/origin/test/util/server"
)

func TestExtensionAPIServerConfigMap(t *testing.T) {
	masterConfig, clusterAdminKubeConfig, err := testserver.StartTestMaster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer testserver.CleanupMasterEtcd(t, masterConfig)

	clusterAdminKubeClient, err := testutil.GetClusterAdminKubeClient(clusterAdminKubeConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var configmap *kapi.ConfigMap
	err = utilwait.PollImmediate(50*time.Millisecond, 10*time.Second, func() (bool, error) {
		configmap, err = clusterAdminKubeClient.Core().ConfigMaps(metav1.NamespaceSystem).Get("extension-apiserver-authentication", metav1.GetOptions{})
		if err == nil {
			return true, nil
		}
		if kapierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := configmap.Data["client-ca-file"]; !ok {
		t.Fatal("missing client-ca-file")
	}
}
