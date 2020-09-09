# HyperflexClusterRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**AlarmSummary** | Pointer to [**HyperflexAlarmSummary**](hyperflex.AlarmSummary.md) |  | [optional] 
**CapacityRunway** | Pointer to **int64** | The number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \&quot;Unknown\&quot; for a cluster that is not connected or with not sufficient data. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The name of this HyperFlex cluster. | [optional] [readonly] 
**ClusterType** | Pointer to **int64** | The storage type of this cluster (All Flash or Hybrid). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | The unique identifier for this HyperFlex cluster. | [optional] [readonly] 
**ComputeNodeCount** | Pointer to **int64** | The number of compute nodes that belong to this cluster. | [optional] [readonly] 
**ConvergedNodeCount** | Pointer to **int64** | The number of converged nodes that belong to this cluster. | [optional] [readonly] 
**DeploymentType** | Pointer to **string** | The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as &#39;NA&#39; (not available). * &#x60;NA&#x60; - The deployment type of the HyperFlex cluster is not available. * &#x60;Datacenter&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * &#x60;Stretched Cluster&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * &#x60;Edge&#x60; - The deployment type of a HyperFlex cluster consisting of 2-4 standalone nodes. | [optional] [readonly] [default to "NA"]
**DeviceId** | Pointer to **string** | The unique identifier of the device registration that represents this HyperFlex cluster&#39;s connection to Intersight. | [optional] [readonly] 
**FltAggr** | Pointer to **int64** | The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms. | [optional] [readonly] 
**HxVersion** | Pointer to **string** | The HyperFlex Data Platform version of this cluster. | [optional] [readonly] 
**HxdpBuildVersion** | Pointer to **string** | The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version. | [optional] [readonly] 
**HypervisorType** | Pointer to **string** | The type of hypervisor running on this cluster. * &#x60;ESXi&#x60; - ESXi hypervisor as specified by the user. * &#x60;HYPERV&#x60; - Hyperv hypervisor as specified by the user. * &#x60;KVM&#x60; - KVM hypervisor as specified by the user. | [optional] [readonly] [default to "ESXi"]
**HypervisorVersion** | Pointer to **string** | The version of hypervisor running on this cluster. | [optional] [readonly] 
**Summary** | Pointer to [**HyperflexSummary**](hyperflex.Summary.md) |  | [optional] 
**UtilizationPercentage** | Pointer to **float32** | The storage utilization percentage is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationTrendPercentage** | Pointer to **float32** | The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | The number of virtual machines present on this cluster. | [optional] [readonly] 
**Alarm** | Pointer to [**[]HyperflexAlarmRelationship**](hyperflex.Alarm.Relationship.md) | An array of relationships to hyperflexAlarm resources. | [optional] [readonly] 
**Health** | Pointer to [**HyperflexHealthRelationship**](hyperflex.Health.Relationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]HyperflexNodeRelationship**](hyperflex.Node.Relationship.md) | An array of relationships to hyperflexNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterRelationship

`func NewHyperflexClusterRelationship(classId string, objectType string, ) *HyperflexClusterRelationship`

NewHyperflexClusterRelationship instantiates a new HyperflexClusterRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterRelationshipWithDefaults

`func NewHyperflexClusterRelationshipWithDefaults() *HyperflexClusterRelationship`

NewHyperflexClusterRelationshipWithDefaults instantiates a new HyperflexClusterRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexClusterRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexClusterRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexClusterRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexClusterRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexClusterRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexClusterRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexClusterRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexClusterRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexClusterRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexClusterRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexClusterRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexClusterRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexClusterRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexClusterRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexClusterRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexClusterRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexClusterRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexClusterRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexClusterRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexClusterRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexClusterRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexClusterRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexClusterRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexClusterRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexClusterRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexClusterRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexClusterRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexClusterRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexClusterRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexClusterRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexClusterRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexClusterRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexClusterRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexClusterRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexClusterRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexClusterRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexClusterRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexClusterRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexClusterRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexClusterRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexClusterRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexClusterRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexClusterRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexClusterRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexClusterRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexClusterRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexClusterRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexClusterRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexClusterRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexClusterRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexClusterRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexClusterRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexClusterRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexClusterRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexClusterRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexClusterRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexClusterRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexClusterRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexClusterRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexClusterRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAlarmSummary

`func (o *HyperflexClusterRelationship) GetAlarmSummary() HyperflexAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *HyperflexClusterRelationship) GetAlarmSummaryOk() (*HyperflexAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *HyperflexClusterRelationship) SetAlarmSummary(v HyperflexAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *HyperflexClusterRelationship) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### GetCapacityRunway

`func (o *HyperflexClusterRelationship) GetCapacityRunway() int64`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *HyperflexClusterRelationship) GetCapacityRunwayOk() (*int64, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *HyperflexClusterRelationship) SetCapacityRunway(v int64)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *HyperflexClusterRelationship) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetClusterName

`func (o *HyperflexClusterRelationship) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HyperflexClusterRelationship) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HyperflexClusterRelationship) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HyperflexClusterRelationship) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterType

`func (o *HyperflexClusterRelationship) GetClusterType() int64`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *HyperflexClusterRelationship) GetClusterTypeOk() (*int64, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *HyperflexClusterRelationship) SetClusterType(v int64)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *HyperflexClusterRelationship) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexClusterRelationship) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexClusterRelationship) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexClusterRelationship) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexClusterRelationship) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetComputeNodeCount

`func (o *HyperflexClusterRelationship) GetComputeNodeCount() int64`

GetComputeNodeCount returns the ComputeNodeCount field if non-nil, zero value otherwise.

### GetComputeNodeCountOk

`func (o *HyperflexClusterRelationship) GetComputeNodeCountOk() (*int64, bool)`

GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodeCount

`func (o *HyperflexClusterRelationship) SetComputeNodeCount(v int64)`

SetComputeNodeCount sets ComputeNodeCount field to given value.

### HasComputeNodeCount

`func (o *HyperflexClusterRelationship) HasComputeNodeCount() bool`

HasComputeNodeCount returns a boolean if a field has been set.

### GetConvergedNodeCount

`func (o *HyperflexClusterRelationship) GetConvergedNodeCount() int64`

GetConvergedNodeCount returns the ConvergedNodeCount field if non-nil, zero value otherwise.

### GetConvergedNodeCountOk

`func (o *HyperflexClusterRelationship) GetConvergedNodeCountOk() (*int64, bool)`

GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergedNodeCount

`func (o *HyperflexClusterRelationship) SetConvergedNodeCount(v int64)`

SetConvergedNodeCount sets ConvergedNodeCount field to given value.

### HasConvergedNodeCount

`func (o *HyperflexClusterRelationship) HasConvergedNodeCount() bool`

HasConvergedNodeCount returns a boolean if a field has been set.

### GetDeploymentType

`func (o *HyperflexClusterRelationship) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *HyperflexClusterRelationship) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *HyperflexClusterRelationship) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *HyperflexClusterRelationship) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDeviceId

`func (o *HyperflexClusterRelationship) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HyperflexClusterRelationship) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HyperflexClusterRelationship) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HyperflexClusterRelationship) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetFltAggr

`func (o *HyperflexClusterRelationship) GetFltAggr() int64`

GetFltAggr returns the FltAggr field if non-nil, zero value otherwise.

### GetFltAggrOk

`func (o *HyperflexClusterRelationship) GetFltAggrOk() (*int64, bool)`

GetFltAggrOk returns a tuple with the FltAggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFltAggr

`func (o *HyperflexClusterRelationship) SetFltAggr(v int64)`

SetFltAggr sets FltAggr field to given value.

### HasFltAggr

`func (o *HyperflexClusterRelationship) HasFltAggr() bool`

HasFltAggr returns a boolean if a field has been set.

### GetHxVersion

`func (o *HyperflexClusterRelationship) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *HyperflexClusterRelationship) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *HyperflexClusterRelationship) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *HyperflexClusterRelationship) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHxdpBuildVersion

`func (o *HyperflexClusterRelationship) GetHxdpBuildVersion() string`

GetHxdpBuildVersion returns the HxdpBuildVersion field if non-nil, zero value otherwise.

### GetHxdpBuildVersionOk

`func (o *HyperflexClusterRelationship) GetHxdpBuildVersionOk() (*string, bool)`

GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpBuildVersion

`func (o *HyperflexClusterRelationship) SetHxdpBuildVersion(v string)`

SetHxdpBuildVersion sets HxdpBuildVersion field to given value.

### HasHxdpBuildVersion

`func (o *HyperflexClusterRelationship) HasHxdpBuildVersion() bool`

HasHxdpBuildVersion returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HyperflexClusterRelationship) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexClusterRelationship) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexClusterRelationship) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexClusterRelationship) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexClusterRelationship) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexClusterRelationship) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexClusterRelationship) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexClusterRelationship) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetSummary

`func (o *HyperflexClusterRelationship) GetSummary() HyperflexSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HyperflexClusterRelationship) GetSummaryOk() (*HyperflexSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HyperflexClusterRelationship) SetSummary(v HyperflexSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *HyperflexClusterRelationship) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUtilizationPercentage

`func (o *HyperflexClusterRelationship) GetUtilizationPercentage() float32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *HyperflexClusterRelationship) GetUtilizationPercentageOk() (*float32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *HyperflexClusterRelationship) SetUtilizationPercentage(v float32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.

### HasUtilizationPercentage

`func (o *HyperflexClusterRelationship) HasUtilizationPercentage() bool`

HasUtilizationPercentage returns a boolean if a field has been set.

### GetUtilizationTrendPercentage

`func (o *HyperflexClusterRelationship) GetUtilizationTrendPercentage() float32`

GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field if non-nil, zero value otherwise.

### GetUtilizationTrendPercentageOk

`func (o *HyperflexClusterRelationship) GetUtilizationTrendPercentageOk() (*float32, bool)`

GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTrendPercentage

`func (o *HyperflexClusterRelationship) SetUtilizationTrendPercentage(v float32)`

SetUtilizationTrendPercentage sets UtilizationTrendPercentage field to given value.

### HasUtilizationTrendPercentage

`func (o *HyperflexClusterRelationship) HasUtilizationTrendPercentage() bool`

HasUtilizationTrendPercentage returns a boolean if a field has been set.

### GetVmCount

`func (o *HyperflexClusterRelationship) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *HyperflexClusterRelationship) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *HyperflexClusterRelationship) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *HyperflexClusterRelationship) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetAlarm

`func (o *HyperflexClusterRelationship) GetAlarm() []HyperflexAlarmRelationship`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *HyperflexClusterRelationship) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *HyperflexClusterRelationship) SetAlarm(v []HyperflexAlarmRelationship)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *HyperflexClusterRelationship) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### SetAlarmNil

`func (o *HyperflexClusterRelationship) SetAlarmNil(b bool)`

 SetAlarmNil sets the value for Alarm to be an explicit nil

### UnsetAlarm
`func (o *HyperflexClusterRelationship) UnsetAlarm()`

UnsetAlarm ensures that no value is present for Alarm, not even an explicit nil
### GetHealth

`func (o *HyperflexClusterRelationship) GetHealth() HyperflexHealthRelationship`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HyperflexClusterRelationship) GetHealthOk() (*HyperflexHealthRelationship, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HyperflexClusterRelationship) SetHealth(v HyperflexHealthRelationship)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *HyperflexClusterRelationship) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetNodes

`func (o *HyperflexClusterRelationship) GetNodes() []HyperflexNodeRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HyperflexClusterRelationship) GetNodesOk() (*[]HyperflexNodeRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HyperflexClusterRelationship) SetNodes(v []HyperflexNodeRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HyperflexClusterRelationship) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *HyperflexClusterRelationship) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *HyperflexClusterRelationship) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetRegisteredDevice

`func (o *HyperflexClusterRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexClusterRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexClusterRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexClusterRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


