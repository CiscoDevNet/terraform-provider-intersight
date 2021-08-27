# KubernetesClusterProfileAllOf

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

### NewKubernetesClusterProfileAllOf

`func NewKubernetesClusterProfileAllOf(classId string, objectType string, ) *KubernetesClusterProfileAllOf`

NewKubernetesClusterProfileAllOf instantiates a new KubernetesClusterProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterProfileAllOfWithDefaults

`func NewKubernetesClusterProfileAllOfWithDefaults() *KubernetesClusterProfileAllOf`

NewKubernetesClusterProfileAllOfWithDefaults instantiates a new KubernetesClusterProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionInfo

`func (o *KubernetesClusterProfileAllOf) GetActionInfo() KubernetesActionInfo`

GetActionInfo returns the ActionInfo field if non-nil, zero value otherwise.

### GetActionInfoOk

`func (o *KubernetesClusterProfileAllOf) GetActionInfoOk() (*KubernetesActionInfo, bool)`

GetActionInfoOk returns a tuple with the ActionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionInfo

`func (o *KubernetesClusterProfileAllOf) SetActionInfo(v KubernetesActionInfo)`

SetActionInfo sets ActionInfo field to given value.

### HasActionInfo

`func (o *KubernetesClusterProfileAllOf) HasActionInfo() bool`

HasActionInfo returns a boolean if a field has been set.

### SetActionInfoNil

`func (o *KubernetesClusterProfileAllOf) SetActionInfoNil(b bool)`

 SetActionInfoNil sets the value for ActionInfo to be an explicit nil

### UnsetActionInfo
`func (o *KubernetesClusterProfileAllOf) UnsetActionInfo()`

UnsetActionInfo ensures that no value is present for ActionInfo, not even an explicit nil
### GetCertConfig

`func (o *KubernetesClusterProfileAllOf) GetCertConfig() KubernetesClusterCertificateConfiguration`

GetCertConfig returns the CertConfig field if non-nil, zero value otherwise.

### GetCertConfigOk

`func (o *KubernetesClusterProfileAllOf) GetCertConfigOk() (*KubernetesClusterCertificateConfiguration, bool)`

GetCertConfigOk returns a tuple with the CertConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertConfig

`func (o *KubernetesClusterProfileAllOf) SetCertConfig(v KubernetesClusterCertificateConfiguration)`

SetCertConfig sets CertConfig field to given value.

### HasCertConfig

`func (o *KubernetesClusterProfileAllOf) HasCertConfig() bool`

HasCertConfig returns a boolean if a field has been set.

### SetCertConfigNil

`func (o *KubernetesClusterProfileAllOf) SetCertConfigNil(b bool)`

 SetCertConfigNil sets the value for CertConfig to be an explicit nil

### UnsetCertConfig
`func (o *KubernetesClusterProfileAllOf) UnsetCertConfig()`

UnsetCertConfig ensures that no value is present for CertConfig, not even an explicit nil
### GetEssentialAddons

`func (o *KubernetesClusterProfileAllOf) GetEssentialAddons() []KubernetesEssentialAddon`

GetEssentialAddons returns the EssentialAddons field if non-nil, zero value otherwise.

### GetEssentialAddonsOk

`func (o *KubernetesClusterProfileAllOf) GetEssentialAddonsOk() (*[]KubernetesEssentialAddon, bool)`

GetEssentialAddonsOk returns a tuple with the EssentialAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssentialAddons

`func (o *KubernetesClusterProfileAllOf) SetEssentialAddons(v []KubernetesEssentialAddon)`

SetEssentialAddons sets EssentialAddons field to given value.

### HasEssentialAddons

`func (o *KubernetesClusterProfileAllOf) HasEssentialAddons() bool`

HasEssentialAddons returns a boolean if a field has been set.

### SetEssentialAddonsNil

`func (o *KubernetesClusterProfileAllOf) SetEssentialAddonsNil(b bool)`

 SetEssentialAddonsNil sets the value for EssentialAddons to be an explicit nil

### UnsetEssentialAddons
`func (o *KubernetesClusterProfileAllOf) UnsetEssentialAddons()`

UnsetEssentialAddons ensures that no value is present for EssentialAddons, not even an explicit nil
### GetKubeConfig

`func (o *KubernetesClusterProfileAllOf) GetKubeConfig() KubernetesConfiguration`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesClusterProfileAllOf) GetKubeConfigOk() (*KubernetesConfiguration, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesClusterProfileAllOf) SetKubeConfig(v KubernetesConfiguration)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesClusterProfileAllOf) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *KubernetesClusterProfileAllOf) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *KubernetesClusterProfileAllOf) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetManagedMode

`func (o *KubernetesClusterProfileAllOf) GetManagedMode() string`

GetManagedMode returns the ManagedMode field if non-nil, zero value otherwise.

### GetManagedModeOk

`func (o *KubernetesClusterProfileAllOf) GetManagedModeOk() (*string, bool)`

GetManagedModeOk returns a tuple with the ManagedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMode

`func (o *KubernetesClusterProfileAllOf) SetManagedMode(v string)`

SetManagedMode sets ManagedMode field to given value.

### HasManagedMode

`func (o *KubernetesClusterProfileAllOf) HasManagedMode() bool`

HasManagedMode returns a boolean if a field has been set.

### GetManagementConfig

`func (o *KubernetesClusterProfileAllOf) GetManagementConfig() KubernetesClusterManagementConfig`

GetManagementConfig returns the ManagementConfig field if non-nil, zero value otherwise.

### GetManagementConfigOk

`func (o *KubernetesClusterProfileAllOf) GetManagementConfigOk() (*KubernetesClusterManagementConfig, bool)`

GetManagementConfigOk returns a tuple with the ManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementConfig

`func (o *KubernetesClusterProfileAllOf) SetManagementConfig(v KubernetesClusterManagementConfig)`

SetManagementConfig sets ManagementConfig field to given value.

### HasManagementConfig

`func (o *KubernetesClusterProfileAllOf) HasManagementConfig() bool`

HasManagementConfig returns a boolean if a field has been set.

### SetManagementConfigNil

`func (o *KubernetesClusterProfileAllOf) SetManagementConfigNil(b bool)`

 SetManagementConfigNil sets the value for ManagementConfig to be an explicit nil

### UnsetManagementConfig
`func (o *KubernetesClusterProfileAllOf) UnsetManagementConfig()`

UnsetManagementConfig ensures that no value is present for ManagementConfig, not even an explicit nil
### GetStatus

`func (o *KubernetesClusterProfileAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesClusterProfileAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesClusterProfileAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesClusterProfileAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAciCniProfile

`func (o *KubernetesClusterProfileAllOf) GetAciCniProfile() KubernetesAciCniProfileRelationship`

GetAciCniProfile returns the AciCniProfile field if non-nil, zero value otherwise.

### GetAciCniProfileOk

`func (o *KubernetesClusterProfileAllOf) GetAciCniProfileOk() (*KubernetesAciCniProfileRelationship, bool)`

GetAciCniProfileOk returns a tuple with the AciCniProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCniProfile

`func (o *KubernetesClusterProfileAllOf) SetAciCniProfile(v KubernetesAciCniProfileRelationship)`

SetAciCniProfile sets AciCniProfile field to given value.

### HasAciCniProfile

`func (o *KubernetesClusterProfileAllOf) HasAciCniProfile() bool`

HasAciCniProfile returns a boolean if a field has been set.

### GetAssociatedCluster

`func (o *KubernetesClusterProfileAllOf) GetAssociatedCluster() KubernetesClusterRelationship`

GetAssociatedCluster returns the AssociatedCluster field if non-nil, zero value otherwise.

### GetAssociatedClusterOk

`func (o *KubernetesClusterProfileAllOf) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool)`

GetAssociatedClusterOk returns a tuple with the AssociatedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCluster

`func (o *KubernetesClusterProfileAllOf) SetAssociatedCluster(v KubernetesClusterRelationship)`

SetAssociatedCluster sets AssociatedCluster field to given value.

### HasAssociatedCluster

`func (o *KubernetesClusterProfileAllOf) HasAssociatedCluster() bool`

HasAssociatedCluster returns a boolean if a field has been set.

### GetClusterIpPools

`func (o *KubernetesClusterProfileAllOf) GetClusterIpPools() []IppoolPoolRelationship`

GetClusterIpPools returns the ClusterIpPools field if non-nil, zero value otherwise.

### GetClusterIpPoolsOk

`func (o *KubernetesClusterProfileAllOf) GetClusterIpPoolsOk() (*[]IppoolPoolRelationship, bool)`

GetClusterIpPoolsOk returns a tuple with the ClusterIpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpPools

`func (o *KubernetesClusterProfileAllOf) SetClusterIpPools(v []IppoolPoolRelationship)`

SetClusterIpPools sets ClusterIpPools field to given value.

### HasClusterIpPools

`func (o *KubernetesClusterProfileAllOf) HasClusterIpPools() bool`

HasClusterIpPools returns a boolean if a field has been set.

### SetClusterIpPoolsNil

`func (o *KubernetesClusterProfileAllOf) SetClusterIpPoolsNil(b bool)`

 SetClusterIpPoolsNil sets the value for ClusterIpPools to be an explicit nil

### UnsetClusterIpPools
`func (o *KubernetesClusterProfileAllOf) UnsetClusterIpPools()`

UnsetClusterIpPools ensures that no value is present for ClusterIpPools, not even an explicit nil
### GetContainerRuntimeConfig

`func (o *KubernetesClusterProfileAllOf) GetContainerRuntimeConfig() KubernetesContainerRuntimePolicyRelationship`

GetContainerRuntimeConfig returns the ContainerRuntimeConfig field if non-nil, zero value otherwise.

### GetContainerRuntimeConfigOk

`func (o *KubernetesClusterProfileAllOf) GetContainerRuntimeConfigOk() (*KubernetesContainerRuntimePolicyRelationship, bool)`

GetContainerRuntimeConfigOk returns a tuple with the ContainerRuntimeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRuntimeConfig

`func (o *KubernetesClusterProfileAllOf) SetContainerRuntimeConfig(v KubernetesContainerRuntimePolicyRelationship)`

SetContainerRuntimeConfig sets ContainerRuntimeConfig field to given value.

### HasContainerRuntimeConfig

`func (o *KubernetesClusterProfileAllOf) HasContainerRuntimeConfig() bool`

HasContainerRuntimeConfig returns a boolean if a field has been set.

### GetLoadbalancerIpLeases

`func (o *KubernetesClusterProfileAllOf) GetLoadbalancerIpLeases() []IppoolIpLeaseRelationship`

GetLoadbalancerIpLeases returns the LoadbalancerIpLeases field if non-nil, zero value otherwise.

### GetLoadbalancerIpLeasesOk

`func (o *KubernetesClusterProfileAllOf) GetLoadbalancerIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetLoadbalancerIpLeasesOk returns a tuple with the LoadbalancerIpLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerIpLeases

`func (o *KubernetesClusterProfileAllOf) SetLoadbalancerIpLeases(v []IppoolIpLeaseRelationship)`

SetLoadbalancerIpLeases sets LoadbalancerIpLeases field to given value.

### HasLoadbalancerIpLeases

`func (o *KubernetesClusterProfileAllOf) HasLoadbalancerIpLeases() bool`

HasLoadbalancerIpLeases returns a boolean if a field has been set.

### SetLoadbalancerIpLeasesNil

`func (o *KubernetesClusterProfileAllOf) SetLoadbalancerIpLeasesNil(b bool)`

 SetLoadbalancerIpLeasesNil sets the value for LoadbalancerIpLeases to be an explicit nil

### UnsetLoadbalancerIpLeases
`func (o *KubernetesClusterProfileAllOf) UnsetLoadbalancerIpLeases()`

UnsetLoadbalancerIpLeases ensures that no value is present for LoadbalancerIpLeases, not even an explicit nil
### GetMasterVipLease

`func (o *KubernetesClusterProfileAllOf) GetMasterVipLease() IppoolIpLeaseRelationship`

GetMasterVipLease returns the MasterVipLease field if non-nil, zero value otherwise.

### GetMasterVipLeaseOk

`func (o *KubernetesClusterProfileAllOf) GetMasterVipLeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetMasterVipLeaseOk returns a tuple with the MasterVipLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVipLease

`func (o *KubernetesClusterProfileAllOf) SetMasterVipLease(v IppoolIpLeaseRelationship)`

SetMasterVipLease sets MasterVipLease field to given value.

### HasMasterVipLease

`func (o *KubernetesClusterProfileAllOf) HasMasterVipLease() bool`

HasMasterVipLease returns a boolean if a field has been set.

### GetNetConfig

`func (o *KubernetesClusterProfileAllOf) GetNetConfig() KubernetesNetworkPolicyRelationship`

GetNetConfig returns the NetConfig field if non-nil, zero value otherwise.

### GetNetConfigOk

`func (o *KubernetesClusterProfileAllOf) GetNetConfigOk() (*KubernetesNetworkPolicyRelationship, bool)`

GetNetConfigOk returns a tuple with the NetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetConfig

`func (o *KubernetesClusterProfileAllOf) SetNetConfig(v KubernetesNetworkPolicyRelationship)`

SetNetConfig sets NetConfig field to given value.

### HasNetConfig

`func (o *KubernetesClusterProfileAllOf) HasNetConfig() bool`

HasNetConfig returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesClusterProfileAllOf) GetNodeGroups() []KubernetesNodeGroupProfileRelationship`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesClusterProfileAllOf) GetNodeGroupsOk() (*[]KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesClusterProfileAllOf) SetNodeGroups(v []KubernetesNodeGroupProfileRelationship)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesClusterProfileAllOf) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### SetNodeGroupsNil

`func (o *KubernetesClusterProfileAllOf) SetNodeGroupsNil(b bool)`

 SetNodeGroupsNil sets the value for NodeGroups to be an explicit nil

### UnsetNodeGroups
`func (o *KubernetesClusterProfileAllOf) UnsetNodeGroups()`

UnsetNodeGroups ensures that no value is present for NodeGroups, not even an explicit nil
### GetOrganization

`func (o *KubernetesClusterProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesClusterProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesClusterProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesClusterProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSysConfig

`func (o *KubernetesClusterProfileAllOf) GetSysConfig() KubernetesSysConfigPolicyRelationship`

GetSysConfig returns the SysConfig field if non-nil, zero value otherwise.

### GetSysConfigOk

`func (o *KubernetesClusterProfileAllOf) GetSysConfigOk() (*KubernetesSysConfigPolicyRelationship, bool)`

GetSysConfigOk returns a tuple with the SysConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysConfig

`func (o *KubernetesClusterProfileAllOf) SetSysConfig(v KubernetesSysConfigPolicyRelationship)`

SetSysConfig sets SysConfig field to given value.

### HasSysConfig

`func (o *KubernetesClusterProfileAllOf) HasSysConfig() bool`

HasSysConfig returns a boolean if a field has been set.

### GetTrustedRegistries

`func (o *KubernetesClusterProfileAllOf) GetTrustedRegistries() KubernetesTrustedRegistriesPolicyRelationship`

GetTrustedRegistries returns the TrustedRegistries field if non-nil, zero value otherwise.

### GetTrustedRegistriesOk

`func (o *KubernetesClusterProfileAllOf) GetTrustedRegistriesOk() (*KubernetesTrustedRegistriesPolicyRelationship, bool)`

GetTrustedRegistriesOk returns a tuple with the TrustedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRegistries

`func (o *KubernetesClusterProfileAllOf) SetTrustedRegistries(v KubernetesTrustedRegistriesPolicyRelationship)`

SetTrustedRegistries sets TrustedRegistries field to given value.

### HasTrustedRegistries

`func (o *KubernetesClusterProfileAllOf) HasTrustedRegistries() bool`

HasTrustedRegistries returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *KubernetesClusterProfileAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *KubernetesClusterProfileAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *KubernetesClusterProfileAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *KubernetesClusterProfileAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


