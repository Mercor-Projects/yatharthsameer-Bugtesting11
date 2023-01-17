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
package in_pod

import (
	"errors"
	"fmt"
	"path/filepath"

	kclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"

	"github.com/openshift/origin/pkg/oc/admin/diagnostics/diagnostics/cluster/network/in_pod/util"
	"github.com/openshift/origin/pkg/oc/admin/diagnostics/diagnostics/types"
)

const (
	CollectNetworkInfoName = "CollectNetworkInfo"
)

// CollectNetworkInfo is a Diagnostic to collect network information in the cluster.
type CollectNetworkInfo struct {
	KubeClient kclientset.Interface
}

// Name is part of the Diagnostic interface and just returns name.
func (d CollectNetworkInfo) Name() string {
	return CollectNetworkInfoName
}

// Description is part of the Diagnostic interface and just returns the diagnostic description.
func (d CollectNetworkInfo) Description() string {
	return "Collect network information in the cluster."
}

func (d CollectNetworkInfo) Requirements() (client bool, host bool) {
	return true, false
}

// CanRun is part of the Diagnostic interface; it determines if the conditions are right to run this diagnostic.
func (d CollectNetworkInfo) CanRun() (bool, error) {
	if d.KubeClient == nil {
		return false, errors.New("must have kube client")
	}
	return true, nil
}

// Check is part of the Diagnostic interface; it runs the actual diagnostic logic
func (d CollectNetworkInfo) Check() types.DiagnosticResult {
	r := types.NewDiagnosticResult(CollectNetworkInfoName)

	nodeName, _, err := util.GetLocalNode(d.KubeClient)
	if err != nil {
		r.Error("DColNet1001", err, fmt.Sprintf("Fetching local node info failed: %s", err))
		return r
	}

	l := util.LogInterface{
		Result: r,
		Logdir: filepath.Join(util.NetworkDiagDefaultLogDir, util.NetworkDiagNodeLogDirPrefix, nodeName),
	}
	l.LogNode(d.KubeClient)
	return r
}
