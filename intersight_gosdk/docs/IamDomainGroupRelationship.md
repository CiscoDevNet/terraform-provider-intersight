# IamDomainGroupRelationship

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
**Name** | Pointer to **string** | The name of the domain-group. | [optional] [readonly] 
**Partition1** | Pointer to **int64** | The partition number domain group related messages are produced for &#39;Partition1&#39; category service topics. | [optional] [readonly] 
**Partition2** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition2&#39; category service topics. | [optional] [readonly] 
**Partition3** | Pointer to **int64** | In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for &#39;Partition3&#39; category service topics. | [optional] [readonly] 
**PartitionKey** | Pointer to **string** | Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as parition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with &#39;H&#39; and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0. | [optional] [readonly] 
**Usage** | Pointer to **int64** | The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamDomainGroupRelationship

`func NewIamDomainGroupRelationship(classId string, objectType string, ) *IamDomainGroupRelationship`

NewIamDomainGroupRelationship instantiates a new IamDomainGroupRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDomainGroupRelationshipWithDefaults

`func NewIamDomainGroupRelationshipWithDefaults() *IamDomainGroupRelationship`

NewIamDomainGroupRelationshipWithDefaults instantiates a new IamDomainGroupRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamDomainGroupRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamDomainGroupRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamDomainGroupRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamDomainGroupRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamDomainGroupRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamDomainGroupRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamDomainGroupRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamDomainGroupRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamDomainGroupRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamDomainGroupRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamDomainGroupRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamDomainGroupRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamDomainGroupRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamDomainGroupRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamDomainGroupRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamDomainGroupRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamDomainGroupRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamDomainGroupRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamDomainGroupRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamDomainGroupRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamDomainGroupRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamDomainGroupRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamDomainGroupRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamDomainGroupRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamDomainGroupRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamDomainGroupRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamDomainGroupRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamDomainGroupRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamDomainGroupRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamDomainGroupRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamDomainGroupRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamDomainGroupRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamDomainGroupRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamDomainGroupRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamDomainGroupRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamDomainGroupRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamDomainGroupRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamDomainGroupRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamDomainGroupRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamDomainGroupRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamDomainGroupRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamDomainGroupRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamDomainGroupRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamDomainGroupRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamDomainGroupRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamDomainGroupRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamDomainGroupRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamDomainGroupRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamDomainGroupRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamDomainGroupRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamDomainGroupRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamDomainGroupRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamDomainGroupRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamDomainGroupRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamDomainGroupRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamDomainGroupRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamDomainGroupRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamDomainGroupRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamDomainGroupRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamDomainGroupRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamDomainGroupRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamDomainGroupRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamDomainGroupRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamDomainGroupRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetName

`func (o *IamDomainGroupRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamDomainGroupRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamDomainGroupRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamDomainGroupRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartition1

`func (o *IamDomainGroupRelationship) GetPartition1() int64`

GetPartition1 returns the Partition1 field if non-nil, zero value otherwise.

### GetPartition1Ok

`func (o *IamDomainGroupRelationship) GetPartition1Ok() (*int64, bool)`

GetPartition1Ok returns a tuple with the Partition1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition1

`func (o *IamDomainGroupRelationship) SetPartition1(v int64)`

SetPartition1 sets Partition1 field to given value.

### HasPartition1

`func (o *IamDomainGroupRelationship) HasPartition1() bool`

HasPartition1 returns a boolean if a field has been set.

### GetPartition2

`func (o *IamDomainGroupRelationship) GetPartition2() int64`

GetPartition2 returns the Partition2 field if non-nil, zero value otherwise.

### GetPartition2Ok

`func (o *IamDomainGroupRelationship) GetPartition2Ok() (*int64, bool)`

GetPartition2Ok returns a tuple with the Partition2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition2

`func (o *IamDomainGroupRelationship) SetPartition2(v int64)`

SetPartition2 sets Partition2 field to given value.

### HasPartition2

`func (o *IamDomainGroupRelationship) HasPartition2() bool`

HasPartition2 returns a boolean if a field has been set.

### GetPartition3

`func (o *IamDomainGroupRelationship) GetPartition3() int64`

GetPartition3 returns the Partition3 field if non-nil, zero value otherwise.

### GetPartition3Ok

`func (o *IamDomainGroupRelationship) GetPartition3Ok() (*int64, bool)`

GetPartition3Ok returns a tuple with the Partition3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition3

`func (o *IamDomainGroupRelationship) SetPartition3(v int64)`

SetPartition3 sets Partition3 field to given value.

### HasPartition3

`func (o *IamDomainGroupRelationship) HasPartition3() bool`

HasPartition3 returns a boolean if a field has been set.

### GetPartitionKey

`func (o *IamDomainGroupRelationship) GetPartitionKey() string`

GetPartitionKey returns the PartitionKey field if non-nil, zero value otherwise.

### GetPartitionKeyOk

`func (o *IamDomainGroupRelationship) GetPartitionKeyOk() (*string, bool)`

GetPartitionKeyOk returns a tuple with the PartitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionKey

`func (o *IamDomainGroupRelationship) SetPartitionKey(v string)`

SetPartitionKey sets PartitionKey field to given value.

### HasPartitionKey

`func (o *IamDomainGroupRelationship) HasPartitionKey() bool`

HasPartitionKey returns a boolean if a field has been set.

### GetUsage

`func (o *IamDomainGroupRelationship) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IamDomainGroupRelationship) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IamDomainGroupRelationship) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IamDomainGroupRelationship) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAccount

`func (o *IamDomainGroupRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamDomainGroupRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamDomainGroupRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamDomainGroupRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


