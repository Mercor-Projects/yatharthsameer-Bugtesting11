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
package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/kubernetes/pkg/kubectl/scheme"

	"github.com/openshift/origin/pkg/cmd/util/serviceability"
	"github.com/openshift/origin/pkg/oc/cli"

	// install all APIs
	apiinstall "github.com/openshift/origin/pkg/api/install"
	apilegacy "github.com/openshift/origin/pkg/api/legacy"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"))()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiinstall.InstallAll(scheme.Scheme, scheme.GroupFactoryRegistry, scheme.Registry)
	apilegacy.LegacyInstallAll(scheme.Scheme, scheme.Registry)

	basename := filepath.Base(os.Args[0])
	command := cli.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
