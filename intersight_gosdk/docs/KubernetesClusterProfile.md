# KubernetesClusterProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ClusterProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ClusterProfile"]
**ActionInfo** | Pointer to [**NullableKubernetesActionInfo**](KubernetesActionInfo.md) |  | [optional] 
**CertConfig** | Pointer to [**NullableKubernetesClusterCertificateConfiguration**](KubernetesClusterCertificateConfiguration.md) |  | [optional] 
**EssentialAddons** | Pointer to [**[]KubernetesEssentialAddon**](KubernetesEssentialAddon.md) |  | [optional] 
**KubeConfig** | Pointer to [**NullableKubernetesConfiguration**](KubernetesConfiguration.md) |  | [optional] 
**ManagedMode** | Pointer to **string** | Management mode for the cluster. In some cases Intersight kubernetes service is not required to provision and manage the management entities and endpoints (for e.g. EKS). In most other cases it will be required to provision and manage these entities and endpoints. * &#x60;Provided&#x60; - Cluster management entities and endpoints are provided by the infrastructure platform. * &#x60;Managed&#x60; - Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service. | [optional] [default to "Provided"]
**ManagementConfig** | Pointer to [**NullableKubernetesClusterManagementConfig**](KubernetesClusterManagementConfig.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the Kubernetes cluster and its nodes. * &#x60;Undeployed&#x60; - The cluster is undeployed. * &#x60;Configuring&#x60; - The cluster is being configured. * &#x60;Deploying&#x60; - The cluster is being deployed. * &#x60;Undeploying&#x60; - The cluster is being undeployed. * &#x60;DeployFailedTerminal&#x60; - The cluster deployment failed terminally and can not be recovered. * &#x60;DeployFailed&#x60; - The cluster deployment failed. * &#x60;Upgrading&#x60; - The cluster is being upgraded. * &#x60;Deleting&#x60; - The cluster is being deleted. * &#x60;DeleteFailed&#x60; - The cluster delete failed. * &#x60;Ready&#x60; - The cluster is ready for use. * &#x60;Active&#x60; - The cluster is being active. * &#x60;Shutdown&#x60; - All the nodes in the cluster are powered off. * &#x60;Terminated&#x60; - The cluster is terminated. * &#x60;Deployed&#x60; - The cluster is deployed. The cluster may not yet be ready for use. * &#x60;UndeployFailed&#x60; - The cluster undeploy action failed. * &#x60;NotReady&#x60; - The cluster is created and some nodes are not ready. | [optional] [default to "Undeployed"]
**AciCniProfile** | Pointer to [**KubernetesAciCniProfileRelationship**](KubernetesAciCniProfileRelationship.md) |  | [optional] 
**AssociatedCluster** | Pointer to [**KubernetesClusterRelationship**](KubernetesClusterRelationship.md) |  | [optional] 
**ClusterIpPools** | Pointer to [**[]IppoolPoolRelationship**](IppoolPoolRelationship.md) | An array of relationships to ippoolPool resources. | [optional] 
**ContainerRuntimeConfig** | Pointer to [**KubernetesContainerRuntimePolicyRelationship**](KubernetesContainerRuntimePolicyRelationship.md) |  | [optional] 
**LoadbalancerIpLeases** | Pointer to [**[]IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) | An array of relationships to ippoolIpLease resources. | [optional] 
**MasterVipLease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**NetConfig** | Pointer to [**KubernetesNetworkPolicyRelationship**](KubernetesNetworkPolicyRelationship.md) |  | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesNodeGroupProfileRelationship**](KubernetesNodeGroupProfileRelationship.md) | An array of relationships to kubernetesNodeGroupProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SysConfig** | Pointer to [**KubernetesSysConfigPolicyRelationship**](KubernetesSysConfigPolicyRelationship.md) |  | [optional] 
**TrustedRegistries** | Pointer to [**KubernetesTrustedRegistriesPolicyRelationship**](KubernetesTrustedRegistriesPolicyRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewKubernetesClusterProfile

`func NewKubernetesClusterProfile(classId string, objectType string, ) *KubernetesClusterProfile`

NewKubernetesClusterProfile instantiates a new KubernetesClusterProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterProfileWithDefaults

`func NewKubernetesClusterProfileWithDefaults() *KubernetesClusterProfile`

NewKubernetesClusterProfileWithDefaults instantiates a new KubernetesClusterProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionInfo

`func (o *KubernetesClusterProfile) GetActionInfo() KubernetesActionInfo`

GetActionInfo returns the ActionInfo field if non-nil, zero value otherwise.

### GetActionInfoOk

`func (o *KubernetesClusterProfile) GetActionInfoOk() (*KubernetesActionInfo, bool)`

GetActionInfoOk returns a tuple with the ActionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionInfo

`func (o *KubernetesClusterProfile) SetActionInfo(v KubernetesActionInfo)`

SetActionInfo sets ActionInfo field to given value.

### HasActionInfo

`func (o *KubernetesClusterProfile) HasActionInfo() bool`

HasActionInfo returns a boolean if a field has been set.

### SetActionInfoNil

`func (o *KubernetesClusterProfile) SetActionInfoNil(b bool)`

 SetActionInfoNil sets the value for ActionInfo to be an explicit nil

### UnsetActionInfo
`func (o *KubernetesClusterProfile) UnsetActionInfo()`

UnsetActionInfo ensures that no value is present for ActionInfo, not even an explicit nil
### GetCertConfig

`func (o *KubernetesClusterProfile) GetCertConfig() KubernetesClusterCertificateConfiguration`

GetCertConfig returns the CertConfig field if non-nil, zero value otherwise.

### GetCertConfigOk

`func (o *KubernetesClusterProfile) GetCertConfigOk() (*KubernetesClusterCertificateConfiguration, bool)`

GetCertConfigOk returns a tuple with the CertConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertConfig

`func (o *KubernetesClusterProfile) SetCertConfig(v KubernetesClusterCertificateConfiguration)`

SetCertConfig sets CertConfig field to given value.

### HasCertConfig

`func (o *KubernetesClusterProfile) HasCertConfig() bool`

HasCertConfig returns a boolean if a field has been set.

### SetCertConfigNil

`func (o *KubernetesClusterProfile) SetCertConfigNil(b bool)`

 SetCertConfigNil sets the value for CertConfig to be an explicit nil

### UnsetCertConfig
`func (o *KubernetesClusterProfile) UnsetCertConfig()`

UnsetCertConfig ensures that no value is present for CertConfig, not even an explicit nil
### GetEssentialAddons

`func (o *KubernetesClusterProfile) GetEssentialAddons() []KubernetesEssentialAddon`

GetEssentialAddons returns the EssentialAddons field if non-nil, zero value otherwise.

### GetEssentialAddonsOk

`func (o *KubernetesClusterProfile) GetEssentialAddonsOk() (*[]KubernetesEssentialAddon, bool)`

GetEssentialAddonsOk returns a tuple with the EssentialAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssentialAddons

`func (o *KubernetesClusterProfile) SetEssentialAddons(v []KubernetesEssentialAddon)`

SetEssentialAddons sets EssentialAddons field to given value.

### HasEssentialAddons

`func (o *KubernetesClusterProfile) HasEssentialAddons() bool`

HasEssentialAddons returns a boolean if a field has been set.

### SetEssentialAddonsNil

`func (o *KubernetesClusterProfile) SetEssentialAddonsNil(b bool)`

 SetEssentialAddonsNil sets the value for EssentialAddons to be an explicit nil

### UnsetEssentialAddons
`func (o *KubernetesClusterProfile) UnsetEssentialAddons()`

UnsetEssentialAddons ensures that no value is present for EssentialAddons, not even an explicit nil
### GetKubeConfig

`func (o *KubernetesClusterProfile) GetKubeConfig() KubernetesConfiguration`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesClusterProfile) GetKubeConfigOk() (*KubernetesConfiguration, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesClusterProfile) SetKubeConfig(v KubernetesConfiguration)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesClusterProfile) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *KubernetesClusterProfile) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *KubernetesClusterProfile) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetManagedMode

`func (o *KubernetesClusterProfile) GetManagedMode() string`

GetManagedMode returns the ManagedMode field if non-nil, zero value otherwise.

### GetManagedModeOk

`func (o *KubernetesClusterProfile) GetManagedModeOk() (*string, bool)`

GetManagedModeOk returns a tuple with the ManagedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMode

`func (o *KubernetesClusterProfile) SetManagedMode(v string)`

SetManagedMode sets ManagedMode field to given value.

### HasManagedMode

`func (o *KubernetesClusterProfile) HasManagedMode() bool`

HasManagedMode returns a boolean if a field has been set.

### GetManagementConfig

`func (o *KubernetesClusterProfile) GetManagementConfig() KubernetesClusterManagementConfig`

GetManagementConfig returns the ManagementConfig field if non-nil, zero value otherwise.

### GetManagementConfigOk

`func (o *KubernetesClusterProfile) GetManagementConfigOk() (*KubernetesClusterManagementConfig, bool)`

GetManagementConfigOk returns a tuple with the ManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementConfig

`func (o *KubernetesClusterProfile) SetManagementConfig(v KubernetesClusterManagementConfig)`

SetManagementConfig sets ManagementConfig field to given value.

### HasManagementConfig

`func (o *KubernetesClusterProfile) HasManagementConfig() bool`

HasManagementConfig returns a boolean if a field has been set.

### SetManagementConfigNil

`func (o *KubernetesClusterProfile) SetManagementConfigNil(b bool)`

 SetManagementConfigNil sets the value for ManagementConfig to be an explicit nil

### UnsetManagementConfig
`func (o *KubernetesClusterProfile) UnsetManagementConfig()`

UnsetManagementConfig ensures that no value is present for ManagementConfig, not even an explicit nil
### GetStatus

`func (o *KubernetesClusterProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesClusterProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesClusterProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesClusterProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAciCniProfile

`func (o *KubernetesClusterProfile) GetAciCniProfile() KubernetesAciCniProfileRelationship`

GetAciCniProfile returns the AciCniProfile field if non-nil, zero value otherwise.

### GetAciCniProfileOk

`func (o *KubernetesClusterProfile) GetAciCniProfileOk() (*KubernetesAciCniProfileRelationship, bool)`

GetAciCniProfileOk returns a tuple with the AciCniProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCniProfile

`func (o *KubernetesClusterProfile) SetAciCniProfile(v KubernetesAciCniProfileRelationship)`

SetAciCniProfile sets AciCniProfile field to given value.

### HasAciCniProfile

`func (o *KubernetesClusterProfile) HasAciCniProfile() bool`

HasAciCniProfile returns a boolean if a field has been set.

### GetAssociatedCluster

`func (o *KubernetesClusterProfile) GetAssociatedCluster() KubernetesClusterRelationship`

GetAssociatedCluster returns the AssociatedCluster field if non-nil, zero value otherwise.

### GetAssociatedClusterOk

`func (o *KubernetesClusterProfile) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool)`

GetAssociatedClusterOk returns a tuple with the AssociatedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCluster

`func (o *KubernetesClusterProfile) SetAssociatedCluster(v KubernetesClusterRelationship)`

SetAssociatedCluster sets AssociatedCluster field to given value.

### HasAssociatedCluster

`func (o *KubernetesClusterProfile) HasAssociatedCluster() bool`

HasAssociatedCluster returns a boolean if a field has been set.

### GetClusterIpPools

`func (o *KubernetesClusterProfile) GetClusterIpPools() []IppoolPoolRelationship`

GetClusterIpPools returns the ClusterIpPools field if non-nil, zero value otherwise.

### GetClusterIpPoolsOk

`func (o *KubernetesClusterProfile) GetClusterIpPoolsOk() (*[]IppoolPoolRelationship, bool)`

GetClusterIpPoolsOk returns a tuple with the ClusterIpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpPools

`func (o *KubernetesClusterProfile) SetClusterIpPools(v []IppoolPoolRelationship)`

SetClusterIpPools sets ClusterIpPools field to given value.

### HasClusterIpPools

`func (o *KubernetesClusterProfile) HasClusterIpPools() bool`

HasClusterIpPools returns a boolean if a field has been set.

### SetClusterIpPoolsNil

`func (o *KubernetesClusterProfile) SetClusterIpPoolsNil(b bool)`

 SetClusterIpPoolsNil sets the value for ClusterIpPools to be an explicit nil

### UnsetClusterIpPools
`func (o *KubernetesClusterProfile) UnsetClusterIpPools()`

UnsetClusterIpPools ensures that no value is present for ClusterIpPools, not even an explicit nil
### GetContainerRuntimeConfig

`func (o *KubernetesClusterProfile) GetContainerRuntimeConfig() KubernetesContainerRuntimePolicyRelationship`

GetContainerRuntimeConfig returns the ContainerRuntimeConfig field if non-nil, zero value otherwise.

### GetContainerRuntimeConfigOk

`func (o *KubernetesClusterProfile) GetContainerRuntimeConfigOk() (*KubernetesContainerRuntimePolicyRelationship, bool)`

GetContainerRuntimeConfigOk returns a tuple with the ContainerRuntimeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeConfig

`func (o *KubernetesClusterProfile) SetContainerRuntimeConfig(v KubernetesContainerRuntimePolicyRelationship)`

SetContainerRuntimeConfig sets ContainerRuntimeConfig field to given value.

### HasContainerRuntimeConfig

`func (o *KubernetesClusterProfile) HasContainerRuntimeConfig() bool`

HasContainerRuntimeConfig returns a boolean if a field has been set.

### GetLoadbalancerIpLeases

`func (o *KubernetesClusterProfile) GetLoadbalancerIpLeases() []IppoolIpLeaseRelationship`

GetLoadbalancerIpLeases returns the LoadbalancerIpLeases field if non-nil, zero value otherwise.

### GetLoadbalancerIpLeasesOk

`func (o *KubernetesClusterProfile) GetLoadbalancerIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetLoadbalancerIpLeasesOk returns a tuple with the LoadbalancerIpLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerIpLeases

`func (o *KubernetesClusterProfile) SetLoadbalancerIpLeases(v []IppoolIpLeaseRelationship)`

SetLoadbalancerIpLeases sets LoadbalancerIpLeases field to given value.

### HasLoadbalancerIpLeases

`func (o *KubernetesClusterProfile) HasLoadbalancerIpLeases() bool`

HasLoadbalancerIpLeases returns a boolean if a field has been set.

### SetLoadbalancerIpLeasesNil

`func (o *KubernetesClusterProfile) SetLoadbalancerIpLeasesNil(b bool)`

 SetLoadbalancerIpLeasesNil sets the value for LoadbalancerIpLeases to be an explicit nil

### UnsetLoadbalancerIpLeases
`func (o *KubernetesClusterProfile) UnsetLoadbalancerIpLeases()`

UnsetLoadbalancerIpLeases ensures that no value is present for LoadbalancerIpLeases, not even an explicit nil
### GetMasterVipLease

`func (o *KubernetesClusterProfile) GetMasterVipLease() IppoolIpLeaseRelationship`

GetMasterVipLease returns the MasterVipLease field if non-nil, zero value otherwise.

### GetMasterVipLeaseOk

`func (o *KubernetesClusterProfile) GetMasterVipLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetMasterVipLeaseOk returns a tuple with the MasterVipLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVipLease

`func (o *KubernetesClusterProfile) SetMasterVipLease(v IppoolIpLeaseRelationship)`

SetMasterVipLease sets MasterVipLease field to given value.

### HasMasterVipLease

`func (o *KubernetesClusterProfile) HasMasterVipLease() bool`

HasMasterVipLease returns a boolean if a field has been set.

### GetNetConfig

`func (o *KubernetesClusterProfile) GetNetConfig() KubernetesNetworkPolicyRelationship`

GetNetConfig returns the NetConfig field if non-nil, zero value otherwise.

### GetNetConfigOk

`func (o *KubernetesClusterProfile) GetNetConfigOk() (*KubernetesNetworkPolicyRelationship, bool)`

GetNetConfigOk returns a tuple with the NetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetConfig

`func (o *KubernetesClusterProfile) SetNetConfig(v KubernetesNetworkPolicyRelationship)`

SetNetConfig sets NetConfig field to given value.

### HasNetConfig

`func (o *KubernetesClusterProfile) HasNetConfig() bool`

HasNetConfig returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesClusterProfile) GetNodeGroups() []KubernetesNodeGroupProfileRelationship`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesClusterProfile) GetNodeGroupsOk() (*[]KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesClusterProfile) SetNodeGroups(v []KubernetesNodeGroupProfileRelationship)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesClusterProfile) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### SetNodeGroupsNil

`func (o *KubernetesClusterProfile) SetNodeGroupsNil(b bool)`

 SetNodeGroupsNil sets the value for NodeGroups to be an explicit nil

### UnsetNodeGroups
`func (o *KubernetesClusterProfile) UnsetNodeGroups()`

UnsetNodeGroups ensures that no value is present for NodeGroups, not even an explicit nil
### GetOrganization

`func (o *KubernetesClusterProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesClusterProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesClusterProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSysConfig

`func (o *KubernetesClusterProfile) GetSysConfig() KubernetesSysConfigPolicyRelationship`

GetSysConfig returns the SysConfig field if non-nil, zero value otherwise.

### GetSysConfigOk

`func (o *KubernetesClusterProfile) GetSysConfigOk() (*KubernetesSysConfigPolicyRelationship, bool)`

GetSysConfigOk returns a tuple with the SysConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysConfig

`func (o *KubernetesClusterProfile) SetSysConfig(v KubernetesSysConfigPolicyRelationship)`

SetSysConfig sets SysConfig field to given value.

### HasSysConfig

`func (o *KubernetesClusterProfile) HasSysConfig() bool`

HasSysConfig returns a boolean if a field has been set.

### GetTrustedRegistries

`func (o *KubernetesClusterProfile) GetTrustedRegistries() KubernetesTrustedRegistriesPolicyRelationship`

GetTrustedRegistries returns the TrustedRegistries field if non-nil, zero value otherwise.

### GetTrustedRegistriesOk

`func (o *KubernetesClusterProfile) GetTrustedRegistriesOk() (*KubernetesTrustedRegistriesPolicyRelationship, bool)`

GetTrustedRegistriesOk returns a tuple with the TrustedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRegistries

`func (o *KubernetesClusterProfile) SetTrustedRegistries(v KubernetesTrustedRegistriesPolicyRelationship)`

SetTrustedRegistries sets TrustedRegistries field to given value.

### HasTrustedRegistries

`func (o *KubernetesClusterProfile) HasTrustedRegistries() bool`

HasTrustedRegistries returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *KubernetesClusterProfile) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *KubernetesClusterProfile) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *KubernetesClusterProfile) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *KubernetesClusterProfile) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


