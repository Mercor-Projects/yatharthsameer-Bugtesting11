#!/bin/bash
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

source "$(dirname "${BASH_SOURCE}")/../../hack/lib/init.sh"
trap os::test::junit::reconcile_output EXIT

# Cleanup cluster resources created by this test
(
  set +e
  oc delete all,templates --all
  exit 0
) &>/dev/null


os::test::junit::declare_suite_start "cmd/describe"
# This test validates non-duplicate errors when describing an existing resource without a defined describer
os::cmd::expect_success 'oc create -f - << __EOF__
{
  "apiVersion": "v1",
  "involvedObject": {
    "apiVersion": "v1",
    "kind": "Pod",
    "name": "test-pod",
    "namespace": "cmd-describer"
  },
  "kind": "Event",
  "message": "test message",
  "metadata": {
    "name": "test-event"
  }
}
__EOF__
'
os::cmd::try_until_success 'eventnum=$(oc get events | wc -l) && [[ $eventnum -gt 0 ]]'
# resources without describers get a default
os::cmd::expect_success_and_text 'oc describe events' 'Namespace:\s+cmd-describer'

# TemplateInstance
os::cmd::expect_success 'oc create -f test/extended/testdata/templates/templateinstance_objectkinds.yaml'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Name:\s+templateinstance'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Namespace:\s+cmd-describer'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Type:\s+Ready'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Status:\s+True'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Secret:\s+cmd-describer/secret'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Deployment:\s+cmd-describer/deployment'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Route:\s+cmd-describer/route'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'Route:\s+cmd-describer/newroute'
os::cmd::expect_success_and_text 'oc describe templateinstances templateinstance' 'NAME:\s+8 bytes'

echo "describer: ok"
os::test::junit::declare_suite_end
