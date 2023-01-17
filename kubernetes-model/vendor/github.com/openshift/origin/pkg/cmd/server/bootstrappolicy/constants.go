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
package bootstrappolicy

// known namespaces
const (
	DefaultOpenShiftSharedResourcesNamespace = "openshift"
	DefaultOpenShiftInfraNamespace           = "openshift-infra"
	DefaultOpenShiftNodeNamespace            = "openshift-node"
)

// users
const (
	DefaultServiceAccountName  = "default"
	BuilderServiceAccountName  = "builder"
	DeployerServiceAccountName = "deployer"

	MasterUnqualifiedUsername     = "openshift-master"
	AggregatorUnqualifiedUsername = "openshift-aggregator"

	MasterUsername      = "system:" + MasterUnqualifiedUsername
	AggregatorUsername  = "system:" + AggregatorUnqualifiedUsername
	SystemAdminUsername = "system:admin"

	// Not granted any API permissions, just an identity for a client certificate for the API proxy to use
	// Should not be changed without considering impact to pods that may be verifying this identity by default
	MasterProxyUnqualifiedUsername = "master-proxy"
	MasterProxyUsername            = "system:" + MasterProxyUnqualifiedUsername

	// Previous versions used this as the username for the master to connect to the kubelet
	// This should remain in the default role bindings for the NodeAdmin role
	LegacyMasterKubeletAdminClientUsername = "system:master"
	MasterKubeletAdminClientUsername       = "system:openshift-node-admin"
)

// groups
const (
	UnauthenticatedUsername = "system:anonymous"

	AuthenticatedGroup      = "system:authenticated"
	AuthenticatedOAuthGroup = "system:authenticated:oauth"
	UnauthenticatedGroup    = "system:unauthenticated"
	ClusterAdminGroup       = "system:cluster-admins"
	ClusterReaderGroup      = "system:cluster-readers"
	MastersGroup            = "system:masters"
	NodesGroup              = "system:nodes"
	NodeAdminsGroup         = "system:node-admins"
	NodeReadersGroup        = "system:node-readers"
)

// Roles
const (
	ClusterAdminRoleName       = "cluster-admin"
	SudoerRoleName             = "sudoer"
	ScopeImpersonationRoleName = "system:scope-impersonation"
	ClusterReaderRoleName      = "cluster-reader"
	StorageAdminRoleName       = "storage-admin"
	ClusterDebuggerRoleName    = "cluster-debugger"
	AdminRoleName              = "admin"
	EditRoleName               = "edit"
	ViewRoleName               = "view"
	AggregatedAdminRoleName    = "system:openshift:aggregate-to-admin"
	AggregatedEditRoleName     = "system:openshift:aggregate-to-edit"
	AggregatedViewRoleName     = "system:openshift:aggregate-to-view"
	SelfProvisionerRoleName    = "self-provisioner"
	BasicUserRoleName          = "basic-user"
	StatusCheckerRoleName      = "cluster-status"
	SelfAccessReviewerRoleName = "self-access-reviewer"

	RegistryAdminRoleName  = "registry-admin"
	RegistryViewerRoleName = "registry-viewer"
	RegistryEditorRoleName = "registry-editor"

	TemplateServiceBrokerClientRoleName = "system:openshift:templateservicebroker-client"

	BuildStrategyDockerRoleName          = "system:build-strategy-docker"
	BuildStrategyCustomRoleName          = "system:build-strategy-custom"
	BuildStrategySourceRoleName          = "system:build-strategy-source"
	BuildStrategyJenkinsPipelineRoleName = "system:build-strategy-jenkinspipeline"

	ImageAuditorRoleName                = "system:image-auditor"
	ImagePullerRoleName                 = "system:image-puller"
	ImagePusherRoleName                 = "system:image-pusher"
	ImageBuilderRoleName                = "system:image-builder"
	ImagePrunerRoleName                 = "system:image-pruner"
	ImageSignerRoleName                 = "system:image-signer"
	DeployerRoleName                    = "system:deployer"
	RouterRoleName                      = "system:router"
	RegistryRoleName                    = "system:registry"
	MasterRoleName                      = "system:master"
	NodeRoleName                        = "system:node"
	NodeProxierRoleName                 = "system:node-proxier"
	SDNReaderRoleName                   = "system:sdn-reader"
	SDNManagerRoleName                  = "system:sdn-manager"
	OAuthTokenDeleterRoleName           = "system:oauth-token-deleter"
	WebHooksRoleName                    = "system:webhook"
	DiscoveryRoleName                   = "system:discovery"
	PersistentVolumeProvisionerRoleName = "system:persistent-volume-provisioner"

	// NodeAdmin has full access to the API provided by the kubelet
	NodeAdminRoleName = "system:node-admin"
	// NodeReader has read access to the metrics and stats provided by the kubelet
	NodeReaderRoleName = "system:node-reader"

	OpenshiftSharedResourceViewRoleName = "shared-resource-viewer"

	NodeBootstrapRoleName    = "system:node-bootstrapper"
	NodeConfigReaderRoleName = "system:node-config-reader"
)

// RoleBindings
const (
	// Legacy roles that must continue to have a plural form
	SelfAccessReviewerRoleBindingName = SelfAccessReviewerRoleName + "s"
	SelfProvisionerRoleBindingName    = SelfProvisionerRoleName + "s"
	DeployerRoleBindingName           = DeployerRoleName + "s"
	ClusterAdminRoleBindingName       = ClusterAdminRoleName + "s"
	ClusterReaderRoleBindingName      = ClusterReaderRoleName + "s"
	BasicUserRoleBindingName          = BasicUserRoleName + "s"
	OAuthTokenDeleterRoleBindingName  = OAuthTokenDeleterRoleName + "s"
	StatusCheckerRoleBindingName      = StatusCheckerRoleName + "-binding"
	ImagePullerRoleBindingName        = ImagePullerRoleName + "s"
	ImageBuilderRoleBindingName       = ImageBuilderRoleName + "s"
	RouterRoleBindingName             = RouterRoleName + "s"
	RegistryRoleBindingName           = RegistryRoleName + "s"
	MasterRoleBindingName             = MasterRoleName + "s"
	NodeRoleBindingName               = NodeRoleName + "s"
	NodeProxierRoleBindingName        = NodeProxierRoleName + "s"
	NodeAdminRoleBindingName          = NodeAdminRoleName + "s"
	NodeReaderRoleBindingName         = NodeReaderRoleName + "s"
	SDNReaderRoleBindingName          = SDNReaderRoleName + "s"
	SDNManagerRoleBindingName         = SDNManagerRoleName + "s"
	WebHooksRoleBindingName           = WebHooksRoleName + "s"
	DiscoveryRoleBindingName          = DiscoveryRoleName + "-binding"
	RegistryAdminRoleBindingName      = RegistryAdminRoleName + "s"
	RegistryViewerRoleBindingName     = RegistryViewerRoleName + "s"
	RegistryEditorRoleBindingName     = RegistryEditorRoleName + "s"

	OpenshiftSharedResourceViewRoleBindingName = OpenshiftSharedResourceViewRoleName + "s"

	// Bindings
	BuildStrategyDockerRoleBindingName          = BuildStrategyDockerRoleName + "-binding"
	BuildStrategyCustomRoleBindingName          = BuildStrategyCustomRoleName + "-binding"
	BuildStrategySourceRoleBindingName          = BuildStrategySourceRoleName + "-binding"
	BuildStrategyJenkinsPipelineRoleBindingName = BuildStrategyJenkinsPipelineRoleName + "-binding"
)
