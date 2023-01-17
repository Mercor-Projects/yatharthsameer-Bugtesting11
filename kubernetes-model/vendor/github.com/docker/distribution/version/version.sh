#!/bin/sh
#
# Copyright (C) 2015 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


# This bash script outputs the current, desired content of version.go, using
# git describe. For best effect, pipe this to the target file. Generally, this
# only needs to updated for releases. The actual value of will be replaced
# during build time if the makefile is used.

set -e

cat <<EOF
package version

// Package is the overall, canonical project import path under which the
// package was built.
var Package = "$(go list)"

// Version indicates which version of the binary is running. This is set to
// the latest release tag by hand, always suffixed by "+unknown". During
// build, it will be replaced by the actual version. The value here will be
// used if the registry is run after a go get based install.
var Version = "$(git describe --match 'v[0-9]*' --dirty='.m' --always)+unknown"
EOF