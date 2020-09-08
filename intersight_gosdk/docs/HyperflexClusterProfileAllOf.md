# HyperflexClusterProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataIpAddress** | Pointer to **string** | The storage data IP address for the HyperFlex cluster. | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the HyperFlex cluster. * &#x60;ESXi&#x60; - ESXi hypervisor as specified by the user. * &#x60;HYPERV&#x60; - Hyperv hypervisor as specified by the user. * &#x60;KVM&#x60; - KVM hypervisor as specified by the user. | [optional] [default to "ESXi"]
**MacAddressPrefix** | Pointer to **string** | The MAC address prefix in the form of 00:25:B5:XX. | [optional] 
**MgmtIpAddress** | Pointer to **string** | The management IP address for the HyperFlex cluster. | [optional] 
**MgmtPlatform** | Pointer to **string** | The management platform for the HyperFlex cluster. * &#x60;FI&#x60; - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * &#x60;EDGE&#x60; - The host servers used in the cluster deployment are standalone severs. | [optional] [default to "FI"]
**Replication** | Pointer to **int64** | The number of copies of each data block written. | [optional] 
**StorageDataVlan** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**WwxnPrefix** | Pointer to **string** | The WWxN prefix in the form of 20:00:00:25:B5:XX. | [optional] 
**AssociatedCluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**AutoSupport** | Pointer to [**HyperflexAutoSupportPolicyRelationship**](hyperflex.AutoSupportPolicy.Relationship.md) |  | [optional] 
**ClusterNetwork** | Pointer to [**HyperflexClusterNetworkPolicyRelationship**](hyperflex.ClusterNetworkPolicy.Relationship.md) |  | [optional] 
**ClusterStorage** | Pointer to [**HyperflexClusterStoragePolicyRelationship**](hyperflex.ClusterStoragePolicy.Relationship.md) |  | [optional] 
**ConfigResult** | Pointer to [**HyperflexConfigResultRelationship**](hyperflex.ConfigResult.Relationship.md) |  | [optional] 
**ExtFcStorage** | Pointer to [**HyperflexExtFcStoragePolicyRelationship**](hyperflex.ExtFcStoragePolicy.Relationship.md) |  | [optional] 
**ExtIscsiStorage** | Pointer to [**HyperflexExtIscsiStoragePolicyRelationship**](hyperflex.ExtIscsiStoragePolicy.Relationship.md) |  | [optional] 
**LocalCredential** | Pointer to [**HyperflexLocalCredentialPolicyRelationship**](hyperflex.LocalCredentialPolicy.Relationship.md) |  | [optional] 
**NodeConfig** | Pointer to [**HyperflexNodeConfigPolicyRelationship**](hyperflex.NodeConfigPolicy.Relationship.md) |  | [optional] 
**NodeProfileConfig** | Pointer to [**[]HyperflexNodeProfileRelationship**](hyperflex.NodeProfile.Relationship.md) | An array of relationships to hyperflexNodeProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ProxySetting** | Pointer to [**HyperflexProxySettingPolicyRelationship**](hyperflex.ProxySettingPolicy.Relationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 
**SoftwareVersion** | Pointer to [**HyperflexSoftwareVersionPolicyRelationship**](hyperflex.SoftwareVersionPolicy.Relationship.md) |  | [optional] 
**SysConfig** | Pointer to [**HyperflexSysConfigPolicyRelationship**](hyperflex.SysConfigPolicy.Relationship.md) |  | [optional] 
**UcsmConfig** | Pointer to [**HyperflexUcsmConfigPolicyRelationship**](hyperflex.UcsmConfigPolicy.Relationship.md) |  | [optional] 
**VcenterConfig** | Pointer to [**HyperflexVcenterConfigPolicyRelationship**](hyperflex.VcenterConfigPolicy.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterProfileAllOf

`func NewHyperflexClusterProfileAllOf() *HyperflexClusterProfileAllOf`

NewHyperflexClusterProfileAllOf instantiates a new HyperflexClusterProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterProfileAllOfWithDefaults

`func NewHyperflexClusterProfileAllOfWithDefaults() *HyperflexClusterProfileAllOf`

NewHyperflexClusterProfileAllOfWithDefaults instantiates a new HyperflexClusterProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataIpAddress

`func (o *HyperflexClusterProfileAllOf) GetDataIpAddress() string`

GetDataIpAddress returns the DataIpAddress field if non-nil, zero value otherwise.

### GetDataIpAddressOk

`func (o *HyperflexClusterProfileAllOf) GetDataIpAddressOk() (*string, bool)`

GetDataIpAddressOk returns a tuple with the DataIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIpAddress

`func (o *HyperflexClusterProfileAllOf) SetDataIpAddress(v string)`

SetDataIpAddress sets DataIpAddress field to given value.

### HasDataIpAddress

`func (o *HyperflexClusterProfileAllOf) HasDataIpAddress() bool`

HasDataIpAddress returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HyperflexClusterProfileAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexClusterProfileAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexClusterProfileAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexClusterProfileAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetMacAddressPrefix

`func (o *HyperflexClusterProfileAllOf) GetMacAddressPrefix() string`

GetMacAddressPrefix returns the MacAddressPrefix field if non-nil, zero value otherwise.

### GetMacAddressPrefixOk

`func (o *HyperflexClusterProfileAllOf) GetMacAddressPrefixOk() (*string, bool)`

GetMacAddressPrefixOk returns a tuple with the MacAddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressPrefix

`func (o *HyperflexClusterProfileAllOf) SetMacAddressPrefix(v string)`

SetMacAddressPrefix sets MacAddressPrefix field to given value.

### HasMacAddressPrefix

`func (o *HyperflexClusterProfileAllOf) HasMacAddressPrefix() bool`

HasMacAddressPrefix returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *HyperflexClusterProfileAllOf) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *HyperflexClusterProfileAllOf) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *HyperflexClusterProfileAllOf) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *HyperflexClusterProfileAllOf) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *HyperflexClusterProfileAllOf) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *HyperflexClusterProfileAllOf) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *HyperflexClusterProfileAllOf) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *HyperflexClusterProfileAllOf) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetReplication

`func (o *HyperflexClusterProfileAllOf) GetReplication() int64`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *HyperflexClusterProfileAllOf) GetReplicationOk() (*int64, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *HyperflexClusterProfileAllOf) SetReplication(v int64)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *HyperflexClusterProfileAllOf) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetStorageDataVlan

`func (o *HyperflexClusterProfileAllOf) GetStorageDataVlan() HyperflexNamedVlan`

GetStorageDataVlan returns the StorageDataVlan field if non-nil, zero value otherwise.

### GetStorageDataVlanOk

`func (o *HyperflexClusterProfileAllOf) GetStorageDataVlanOk() (*HyperflexNamedVlan, bool)`

GetStorageDataVlanOk returns a tuple with the StorageDataVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDataVlan

`func (o *HyperflexClusterProfileAllOf) SetStorageDataVlan(v HyperflexNamedVlan)`

SetStorageDataVlan sets StorageDataVlan field to given value.

### HasStorageDataVlan

`func (o *HyperflexClusterProfileAllOf) HasStorageDataVlan() bool`

HasStorageDataVlan returns a boolean if a field has been set.

### GetWwxnPrefix

`func (o *HyperflexClusterProfileAllOf) GetWwxnPrefix() string`

GetWwxnPrefix returns the WwxnPrefix field if non-nil, zero value otherwise.

### GetWwxnPrefixOk

`func (o *HyperflexClusterProfileAllOf) GetWwxnPrefixOk() (*string, bool)`

GetWwxnPrefixOk returns a tuple with the WwxnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwxnPrefix

`func (o *HyperflexClusterProfileAllOf) SetWwxnPrefix(v string)`

SetWwxnPrefix sets WwxnPrefix field to given value.

### HasWwxnPrefix

`func (o *HyperflexClusterProfileAllOf) HasWwxnPrefix() bool`

HasWwxnPrefix returns a boolean if a field has been set.

### GetAssociatedCluster

`func (o *HyperflexClusterProfileAllOf) GetAssociatedCluster() HyperflexClusterRelationship`

GetAssociatedCluster returns the AssociatedCluster field if non-nil, zero value otherwise.

### GetAssociatedClusterOk

`func (o *HyperflexClusterProfileAllOf) GetAssociatedClusterOk() (*HyperflexClusterRelationship, bool)`

GetAssociatedClusterOk returns a tuple with the AssociatedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCluster

`func (o *HyperflexClusterProfileAllOf) SetAssociatedCluster(v HyperflexClusterRelationship)`

SetAssociatedCluster sets AssociatedCluster field to given value.

### HasAssociatedCluster

`func (o *HyperflexClusterProfileAllOf) HasAssociatedCluster() bool`

HasAssociatedCluster returns a boolean if a field has been set.

### GetAutoSupport

`func (o *HyperflexClusterProfileAllOf) GetAutoSupport() HyperflexAutoSupportPolicyRelationship`

GetAutoSupport returns the AutoSupport field if non-nil, zero value otherwise.

### GetAutoSupportOk

`func (o *HyperflexClusterProfileAllOf) GetAutoSupportOk() (*HyperflexAutoSupportPolicyRelationship, bool)`

GetAutoSupportOk returns a tuple with the AutoSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSupport

`func (o *HyperflexClusterProfileAllOf) SetAutoSupport(v HyperflexAutoSupportPolicyRelationship)`

SetAutoSupport sets AutoSupport field to given value.

### HasAutoSupport

`func (o *HyperflexClusterProfileAllOf) HasAutoSupport() bool`

HasAutoSupport returns a boolean if a field has been set.

### GetClusterNetwork

`func (o *HyperflexClusterProfileAllOf) GetClusterNetwork() HyperflexClusterNetworkPolicyRelationship`

GetClusterNetwork returns the ClusterNetwork field if non-nil, zero value otherwise.

### GetClusterNetworkOk

`func (o *HyperflexClusterProfileAllOf) GetClusterNetworkOk() (*HyperflexClusterNetworkPolicyRelationship, bool)`

GetClusterNetworkOk returns a tuple with the ClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetwork

`func (o *HyperflexClusterProfileAllOf) SetClusterNetwork(v HyperflexClusterNetworkPolicyRelationship)`

SetClusterNetwork sets ClusterNetwork field to given value.

### HasClusterNetwork

`func (o *HyperflexClusterProfileAllOf) HasClusterNetwork() bool`

HasClusterNetwork returns a boolean if a field has been set.

### GetClusterStorage

`func (o *HyperflexClusterProfileAllOf) GetClusterStorage() HyperflexClusterStoragePolicyRelationship`

GetClusterStorage returns the ClusterStorage field if non-nil, zero value otherwise.

### GetClusterStorageOk

`func (o *HyperflexClusterProfileAllOf) GetClusterStorageOk() (*HyperflexClusterStoragePolicyRelationship, bool)`

GetClusterStorageOk returns a tuple with the ClusterStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorage

`func (o *HyperflexClusterProfileAllOf) SetClusterStorage(v HyperflexClusterStoragePolicyRelationship)`

SetClusterStorage sets ClusterStorage field to given value.

### HasClusterStorage

`func (o *HyperflexClusterProfileAllOf) HasClusterStorage() bool`

HasClusterStorage returns a boolean if a field has been set.

### GetConfigResult

`func (o *HyperflexClusterProfileAllOf) GetConfigResult() HyperflexConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *HyperflexClusterProfileAllOf) GetConfigResultOk() (*HyperflexConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *HyperflexClusterProfileAllOf) SetConfigResult(v HyperflexConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *HyperflexClusterProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetExtFcStorage

`func (o *HyperflexClusterProfileAllOf) GetExtFcStorage() HyperflexExtFcStoragePolicyRelationship`

GetExtFcStorage returns the ExtFcStorage field if non-nil, zero value otherwise.

### GetExtFcStorageOk

`func (o *HyperflexClusterProfileAllOf) GetExtFcStorageOk() (*HyperflexExtFcStoragePolicyRelationship, bool)`

GetExtFcStorageOk returns a tuple with the ExtFcStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtFcStorage

`func (o *HyperflexClusterProfileAllOf) SetExtFcStorage(v HyperflexExtFcStoragePolicyRelationship)`

SetExtFcStorage sets ExtFcStorage field to given value.

### HasExtFcStorage

`func (o *HyperflexClusterProfileAllOf) HasExtFcStorage() bool`

HasExtFcStorage returns a boolean if a field has been set.

### GetExtIscsiStorage

`func (o *HyperflexClusterProfileAllOf) GetExtIscsiStorage() HyperflexExtIscsiStoragePolicyRelationship`

GetExtIscsiStorage returns the ExtIscsiStorage field if non-nil, zero value otherwise.

### GetExtIscsiStorageOk

`func (o *HyperflexClusterProfileAllOf) GetExtIscsiStorageOk() (*HyperflexExtIscsiStoragePolicyRelationship, bool)`

GetExtIscsiStorageOk returns a tuple with the ExtIscsiStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIscsiStorage

`func (o *HyperflexClusterProfileAllOf) SetExtIscsiStorage(v HyperflexExtIscsiStoragePolicyRelationship)`

SetExtIscsiStorage sets ExtIscsiStorage field to given value.

### HasExtIscsiStorage

`func (o *HyperflexClusterProfileAllOf) HasExtIscsiStorage() bool`

HasExtIscsiStorage returns a boolean if a field has been set.

### GetLocalCredential

`func (o *HyperflexClusterProfileAllOf) GetLocalCredential() HyperflexLocalCredentialPolicyRelationship`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *HyperflexClusterProfileAllOf) GetLocalCredentialOk() (*HyperflexLocalCredentialPolicyRelationship, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *HyperflexClusterProfileAllOf) SetLocalCredential(v HyperflexLocalCredentialPolicyRelationship)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *HyperflexClusterProfileAllOf) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### GetNodeConfig

`func (o *HyperflexClusterProfileAllOf) GetNodeConfig() HyperflexNodeConfigPolicyRelationship`

GetNodeConfig returns the NodeConfig field if non-nil, zero value otherwise.

### GetNodeConfigOk

`func (o *HyperflexClusterProfileAllOf) GetNodeConfigOk() (*HyperflexNodeConfigPolicyRelationship, bool)`

GetNodeConfigOk returns a tuple with the NodeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfig

`func (o *HyperflexClusterProfileAllOf) SetNodeConfig(v HyperflexNodeConfigPolicyRelationship)`

SetNodeConfig sets NodeConfig field to given value.

### HasNodeConfig

`func (o *HyperflexClusterProfileAllOf) HasNodeConfig() bool`

HasNodeConfig returns a boolean if a field has been set.

### GetNodeProfileConfig

`func (o *HyperflexClusterProfileAllOf) GetNodeProfileConfig() []HyperflexNodeProfileRelationship`

GetNodeProfileConfig returns the NodeProfileConfig field if non-nil, zero value otherwise.

### GetNodeProfileConfigOk

`func (o *HyperflexClusterProfileAllOf) GetNodeProfileConfigOk() (*[]HyperflexNodeProfileRelationship, bool)`

GetNodeProfileConfigOk returns a tuple with the NodeProfileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeProfileConfig

`func (o *HyperflexClusterProfileAllOf) SetNodeProfileConfig(v []HyperflexNodeProfileRelationship)`

SetNodeProfileConfig sets NodeProfileConfig field to given value.

### HasNodeProfileConfig

`func (o *HyperflexClusterProfileAllOf) HasNodeProfileConfig() bool`

HasNodeProfileConfig returns a boolean if a field has been set.

### SetNodeProfileConfigNil

`func (o *HyperflexClusterProfileAllOf) SetNodeProfileConfigNil(b bool)`

 SetNodeProfileConfigNil sets the value for NodeProfileConfig to be an explicit nil

### UnsetNodeProfileConfig
`func (o *HyperflexClusterProfileAllOf) UnsetNodeProfileConfig()`

UnsetNodeProfileConfig ensures that no value is present for NodeProfileConfig, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProxySetting

`func (o *HyperflexClusterProfileAllOf) GetProxySetting() HyperflexProxySettingPolicyRelationship`

GetProxySetting returns the ProxySetting field if non-nil, zero value otherwise.

### GetProxySettingOk

`func (o *HyperflexClusterProfileAllOf) GetProxySettingOk() (*HyperflexProxySettingPolicyRelationship, bool)`

GetProxySettingOk returns a tuple with the ProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxySetting

`func (o *HyperflexClusterProfileAllOf) SetProxySetting(v HyperflexProxySettingPolicyRelationship)`

SetProxySetting sets ProxySetting field to given value.

### HasProxySetting

`func (o *HyperflexClusterProfileAllOf) HasProxySetting() bool`

HasProxySetting returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *HyperflexClusterProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *HyperflexClusterProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *HyperflexClusterProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *HyperflexClusterProfileAllOf) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *HyperflexClusterProfileAllOf) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *HyperflexClusterProfileAllOf) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil
### GetSoftwareVersion

`func (o *HyperflexClusterProfileAllOf) GetSoftwareVersion() HyperflexSoftwareVersionPolicyRelationship`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *HyperflexClusterProfileAllOf) GetSoftwareVersionOk() (*HyperflexSoftwareVersionPolicyRelationship, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *HyperflexClusterProfileAllOf) SetSoftwareVersion(v HyperflexSoftwareVersionPolicyRelationship)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *HyperflexClusterProfileAllOf) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSysConfig

`func (o *HyperflexClusterProfileAllOf) GetSysConfig() HyperflexSysConfigPolicyRelationship`

GetSysConfig returns the SysConfig field if non-nil, zero value otherwise.

### GetSysConfigOk

`func (o *HyperflexClusterProfileAllOf) GetSysConfigOk() (*HyperflexSysConfigPolicyRelationship, bool)`

GetSysConfigOk returns a tuple with the SysConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysConfig

`func (o *HyperflexClusterProfileAllOf) SetSysConfig(v HyperflexSysConfigPolicyRelationship)`

SetSysConfig sets SysConfig field to given value.

### HasSysConfig

`func (o *HyperflexClusterProfileAllOf) HasSysConfig() bool`

HasSysConfig returns a boolean if a field has been set.

### GetUcsmConfig

`func (o *HyperflexClusterProfileAllOf) GetUcsmConfig() HyperflexUcsmConfigPolicyRelationship`

GetUcsmConfig returns the UcsmConfig field if non-nil, zero value otherwise.

### GetUcsmConfigOk

`func (o *HyperflexClusterProfileAllOf) GetUcsmConfigOk() (*HyperflexUcsmConfigPolicyRelationship, bool)`

GetUcsmConfigOk returns a tuple with the UcsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsmConfig

`func (o *HyperflexClusterProfileAllOf) SetUcsmConfig(v HyperflexUcsmConfigPolicyRelationship)`

SetUcsmConfig sets UcsmConfig field to given value.

### HasUcsmConfig

`func (o *HyperflexClusterProfileAllOf) HasUcsmConfig() bool`

HasUcsmConfig returns a boolean if a field has been set.

### GetVcenterConfig

`func (o *HyperflexClusterProfileAllOf) GetVcenterConfig() HyperflexVcenterConfigPolicyRelationship`

GetVcenterConfig returns the VcenterConfig field if non-nil, zero value otherwise.

### GetVcenterConfigOk

`func (o *HyperflexClusterProfileAllOf) GetVcenterConfigOk() (*HyperflexVcenterConfigPolicyRelationship, bool)`

GetVcenterConfigOk returns a tuple with the VcenterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterConfig

`func (o *HyperflexClusterProfileAllOf) SetVcenterConfig(v HyperflexVcenterConfigPolicyRelationship)`

SetVcenterConfig sets VcenterConfig field to given value.

### HasVcenterConfig

`func (o *HyperflexClusterProfileAllOf) HasVcenterConfig() bool`

HasVcenterConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


