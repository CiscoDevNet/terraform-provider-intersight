# CondHclStatusRelationship

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
**ComponentStatus** | Pointer to **string** | The overall status for the components found in the HCL. This will provide the HCL validation status for all the components. It can be one of the following. \&quot;Validated\&quot; - all the components hardware/software profiles are listed in the HCL. \&quot;Not-Listed\&quot; - one or more components hardware/software profiles are not listed in the HCL \&quot;Incomplete\&quot; - the components are not evaluated as the server&#39;s software/hardware profiles are not listed in the HCL. \&quot;Not-Evaluated\&quot; - The components are not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**HardwareStatus** | Pointer to **string** | The server model, processor and firmware are considered as part of the hardware profile for the server. This will provide the HCL validation status for the hardware profile. For the failure reason see the serverReason property. The status can be one of the following \&quot;Validated\&quot; - The server model, processor and firmware combination is listed in the HCL \&quot;Not-Listed\&quot; - The server model, processor and firmware combination is not listed in the HCL. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**HclFirmwareVersion** | Pointer to **string** | The current CIMC version for the server normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclModel** | Pointer to **string** | The managed object&#39;s model to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclOsVendor** | Pointer to **string** | The OS Vendor for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclOsVersion** | Pointer to **string** | The OS Version for the managed object to validate normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**HclProcessor** | Pointer to **string** | The managed object&#39;s processor to validate if applicable normalized for querying HCL data. It is empty if we are missing this information. | [optional] 
**InvFirmwareVersion** | Pointer to **string** | The current CIMC version for the server as received from inventory. It is empty if we are missing this information. | [optional] 
**InvModel** | Pointer to **string** | The managed object&#39;s model to validate as received from the inventory. It is empty if we are missing this information. | [optional] 
**InvOsVendor** | Pointer to **string** | The OS Vendor for the managed object to validate as received from inventory. It is empty if we are missing this information. | [optional] 
**InvOsVersion** | Pointer to **string** | The OS Version for the managed object to validate as received from inventory. It is empty if we are missing this information. | [optional] 
**InvProcessor** | Pointer to **string** | The managed object&#39;s processor to validate if applicable as received from inventory. It is empty if we are missing this information. | [optional] 
**Reason** | Pointer to **string** | The reason for the HCL status. It will be one of the following \&quot;Missing-Os-Info\&quot; - we are missing os information in the inventory from the device connector \&quot;Incompatible-Components\&quot; - we have 1 or more components with \&quot;Not-Validated\&quot; status \&quot;Compatible\&quot; - all the components have \&quot;Validated\&quot; status. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Missing-Os-Info&#x60; - This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Components&#x60; - This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails. * &#x60;Compatible&#x60; - This means the HclStatus for the sever has passed HCL validation for all of its related components. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the sever has not been evaluated because it is exempted. | [optional] [default to "Missing-Os-Info"]
**ServerReason** | Pointer to **string** | The reason generated by the server&#39;s HCL validation. For HCL the evaluation can be seen in three logical stages 1. Evaluate the server&#39;s hardware status 2. Evaluate the server&#39;s software status 3. Evaluate the server&#39;s components (each component has its own hardware/software evaluation) The evaluation does not proceed to the next stage until the previous stage is evaluated. Therefore there can be only one validation reason. \&quot;Incompatible-Server\&quot; - the server model is not listed in the HCL. \&quot;Incompatible-Processor\&quot; - the server model and processor combination is not listed in HCL. \&quot;Incompatible-Firmware\&quot; - the server model, processor and server firmware is not listed in HCL. \&quot;Missing-Os-Info\&quot; - the os vendor and version is not listed in HCL with the HW profile. \&quot;Incompatible-Os-Info\&quot; - the os vendor and version is not listed in HCL with the HW profile. \&quot;Incompatible-Components\&quot; - there is one or more components with \&quot;Not-Validated\&quot; status \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). \&quot;Compatible\&quot; - the server and all its components are validated. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information. * &#x60;Incompatible-Server&#x60; - The validation failed for this server because the server&#39;s model was not listed in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not listed for the given server model. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has returned a service error or unrecognized result. * &#x60;Not-Evaluated&#x60; - This means the HclStatus for the sever has not been evaluated because it is exempted. * &#x60;Incompatible-Components&#x60; - The validation has failed for this server because one or more components have \&quot;Not-Listed\&quot; status. * &#x60;Compatible&#x60; - The validation has passed for this server&#39;s model, processor, OS vendor and version. | [optional] [default to "Missing-Os-Driver-Info"]
**SoftwareStatus** | Pointer to **string** | The OS vendor and version are considered part of the software profile for the server. This will provide the HCL validation status for the software profile. For the failure reason see the serverReason property. The status can be be one of the following \&quot;Validated\&quot; - The os vendor/version is listed in the HCL for the server model, processor and firmware \&quot;Not-Listed\&quot; - The os vendor/version is not listed in the HCL for the server model, processor and firmware \&quot;Incomplete\&quot; - The inventory is missing os vendor/version and HCL validation was not performed. \&quot;Not-Evaluated\&quot; - The server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Status** | Pointer to **string** | The HCL compatibility status of the managed object. The status can be one of the following \&quot;Incomplete\&quot; - there is no enough information to evaluate against the HCL data \&quot;Validated\&quot; - all components have been evaluated against the HCL and they all have \&quot;Validated\&quot; status \&quot;Not-Listed\&quot; - all components have been evaluated against the HCL and one or more have \&quot;Not-Listed\&quot; status. \&quot;Not-Evaluated\&quot; - server is not evaluated against the HCL because it is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Details** | Pointer to [**[]CondHclStatusDetailRelationship**](cond.HclStatusDetail.Relationship.md) | An array of relationships to condHclStatusDetail resources. | [optional] [readonly] 
**ManagedObject** | Pointer to [**InventoryBaseRelationship**](inventory.Base.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCondHclStatusRelationship

`func NewCondHclStatusRelationship(classId string, objectType string, ) *CondHclStatusRelationship`

NewCondHclStatusRelationship instantiates a new CondHclStatusRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusRelationshipWithDefaults

`func NewCondHclStatusRelationshipWithDefaults() *CondHclStatusRelationship`

NewCondHclStatusRelationshipWithDefaults instantiates a new CondHclStatusRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *CondHclStatusRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *CondHclStatusRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *CondHclStatusRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *CondHclStatusRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *CondHclStatusRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondHclStatusRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondHclStatusRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *CondHclStatusRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CondHclStatusRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CondHclStatusRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CondHclStatusRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *CondHclStatusRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *CondHclStatusRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *CondHclStatusRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *CondHclStatusRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *CondHclStatusRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *CondHclStatusRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *CondHclStatusRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *CondHclStatusRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *CondHclStatusRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *CondHclStatusRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *CondHclStatusRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *CondHclStatusRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *CondHclStatusRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondHclStatusRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondHclStatusRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *CondHclStatusRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *CondHclStatusRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *CondHclStatusRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *CondHclStatusRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *CondHclStatusRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *CondHclStatusRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *CondHclStatusRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *CondHclStatusRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *CondHclStatusRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CondHclStatusRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CondHclStatusRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CondHclStatusRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *CondHclStatusRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *CondHclStatusRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *CondHclStatusRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *CondHclStatusRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *CondHclStatusRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *CondHclStatusRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *CondHclStatusRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *CondHclStatusRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *CondHclStatusRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *CondHclStatusRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *CondHclStatusRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CondHclStatusRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CondHclStatusRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CondHclStatusRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *CondHclStatusRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *CondHclStatusRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *CondHclStatusRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *CondHclStatusRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *CondHclStatusRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *CondHclStatusRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *CondHclStatusRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *CondHclStatusRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *CondHclStatusRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *CondHclStatusRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *CondHclStatusRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *CondHclStatusRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetComponentStatus

`func (o *CondHclStatusRelationship) GetComponentStatus() string`

GetComponentStatus returns the ComponentStatus field if non-nil, zero value otherwise.

### GetComponentStatusOk

`func (o *CondHclStatusRelationship) GetComponentStatusOk() (*string, bool)`

GetComponentStatusOk returns a tuple with the ComponentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentStatus

`func (o *CondHclStatusRelationship) SetComponentStatus(v string)`

SetComponentStatus sets ComponentStatus field to given value.

### HasComponentStatus

`func (o *CondHclStatusRelationship) HasComponentStatus() bool`

HasComponentStatus returns a boolean if a field has been set.

### GetHardwareStatus

`func (o *CondHclStatusRelationship) GetHardwareStatus() string`

GetHardwareStatus returns the HardwareStatus field if non-nil, zero value otherwise.

### GetHardwareStatusOk

`func (o *CondHclStatusRelationship) GetHardwareStatusOk() (*string, bool)`

GetHardwareStatusOk returns a tuple with the HardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareStatus

`func (o *CondHclStatusRelationship) SetHardwareStatus(v string)`

SetHardwareStatus sets HardwareStatus field to given value.

### HasHardwareStatus

`func (o *CondHclStatusRelationship) HasHardwareStatus() bool`

HasHardwareStatus returns a boolean if a field has been set.

### GetHclFirmwareVersion

`func (o *CondHclStatusRelationship) GetHclFirmwareVersion() string`

GetHclFirmwareVersion returns the HclFirmwareVersion field if non-nil, zero value otherwise.

### GetHclFirmwareVersionOk

`func (o *CondHclStatusRelationship) GetHclFirmwareVersionOk() (*string, bool)`

GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclFirmwareVersion

`func (o *CondHclStatusRelationship) SetHclFirmwareVersion(v string)`

SetHclFirmwareVersion sets HclFirmwareVersion field to given value.

### HasHclFirmwareVersion

`func (o *CondHclStatusRelationship) HasHclFirmwareVersion() bool`

HasHclFirmwareVersion returns a boolean if a field has been set.

### GetHclModel

`func (o *CondHclStatusRelationship) GetHclModel() string`

GetHclModel returns the HclModel field if non-nil, zero value otherwise.

### GetHclModelOk

`func (o *CondHclStatusRelationship) GetHclModelOk() (*string, bool)`

GetHclModelOk returns a tuple with the HclModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclModel

`func (o *CondHclStatusRelationship) SetHclModel(v string)`

SetHclModel sets HclModel field to given value.

### HasHclModel

`func (o *CondHclStatusRelationship) HasHclModel() bool`

HasHclModel returns a boolean if a field has been set.

### GetHclOsVendor

`func (o *CondHclStatusRelationship) GetHclOsVendor() string`

GetHclOsVendor returns the HclOsVendor field if non-nil, zero value otherwise.

### GetHclOsVendorOk

`func (o *CondHclStatusRelationship) GetHclOsVendorOk() (*string, bool)`

GetHclOsVendorOk returns a tuple with the HclOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVendor

`func (o *CondHclStatusRelationship) SetHclOsVendor(v string)`

SetHclOsVendor sets HclOsVendor field to given value.

### HasHclOsVendor

`func (o *CondHclStatusRelationship) HasHclOsVendor() bool`

HasHclOsVendor returns a boolean if a field has been set.

### GetHclOsVersion

`func (o *CondHclStatusRelationship) GetHclOsVersion() string`

GetHclOsVersion returns the HclOsVersion field if non-nil, zero value otherwise.

### GetHclOsVersionOk

`func (o *CondHclStatusRelationship) GetHclOsVersionOk() (*string, bool)`

GetHclOsVersionOk returns a tuple with the HclOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclOsVersion

`func (o *CondHclStatusRelationship) SetHclOsVersion(v string)`

SetHclOsVersion sets HclOsVersion field to given value.

### HasHclOsVersion

`func (o *CondHclStatusRelationship) HasHclOsVersion() bool`

HasHclOsVersion returns a boolean if a field has been set.

### GetHclProcessor

`func (o *CondHclStatusRelationship) GetHclProcessor() string`

GetHclProcessor returns the HclProcessor field if non-nil, zero value otherwise.

### GetHclProcessorOk

`func (o *CondHclStatusRelationship) GetHclProcessorOk() (*string, bool)`

GetHclProcessorOk returns a tuple with the HclProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclProcessor

`func (o *CondHclStatusRelationship) SetHclProcessor(v string)`

SetHclProcessor sets HclProcessor field to given value.

### HasHclProcessor

`func (o *CondHclStatusRelationship) HasHclProcessor() bool`

HasHclProcessor returns a boolean if a field has been set.

### GetInvFirmwareVersion

`func (o *CondHclStatusRelationship) GetInvFirmwareVersion() string`

GetInvFirmwareVersion returns the InvFirmwareVersion field if non-nil, zero value otherwise.

### GetInvFirmwareVersionOk

`func (o *CondHclStatusRelationship) GetInvFirmwareVersionOk() (*string, bool)`

GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvFirmwareVersion

`func (o *CondHclStatusRelationship) SetInvFirmwareVersion(v string)`

SetInvFirmwareVersion sets InvFirmwareVersion field to given value.

### HasInvFirmwareVersion

`func (o *CondHclStatusRelationship) HasInvFirmwareVersion() bool`

HasInvFirmwareVersion returns a boolean if a field has been set.

### GetInvModel

`func (o *CondHclStatusRelationship) GetInvModel() string`

GetInvModel returns the InvModel field if non-nil, zero value otherwise.

### GetInvModelOk

`func (o *CondHclStatusRelationship) GetInvModelOk() (*string, bool)`

GetInvModelOk returns a tuple with the InvModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvModel

`func (o *CondHclStatusRelationship) SetInvModel(v string)`

SetInvModel sets InvModel field to given value.

### HasInvModel

`func (o *CondHclStatusRelationship) HasInvModel() bool`

HasInvModel returns a boolean if a field has been set.

### GetInvOsVendor

`func (o *CondHclStatusRelationship) GetInvOsVendor() string`

GetInvOsVendor returns the InvOsVendor field if non-nil, zero value otherwise.

### GetInvOsVendorOk

`func (o *CondHclStatusRelationship) GetInvOsVendorOk() (*string, bool)`

GetInvOsVendorOk returns a tuple with the InvOsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVendor

`func (o *CondHclStatusRelationship) SetInvOsVendor(v string)`

SetInvOsVendor sets InvOsVendor field to given value.

### HasInvOsVendor

`func (o *CondHclStatusRelationship) HasInvOsVendor() bool`

HasInvOsVendor returns a boolean if a field has been set.

### GetInvOsVersion

`func (o *CondHclStatusRelationship) GetInvOsVersion() string`

GetInvOsVersion returns the InvOsVersion field if non-nil, zero value otherwise.

### GetInvOsVersionOk

`func (o *CondHclStatusRelationship) GetInvOsVersionOk() (*string, bool)`

GetInvOsVersionOk returns a tuple with the InvOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvOsVersion

`func (o *CondHclStatusRelationship) SetInvOsVersion(v string)`

SetInvOsVersion sets InvOsVersion field to given value.

### HasInvOsVersion

`func (o *CondHclStatusRelationship) HasInvOsVersion() bool`

HasInvOsVersion returns a boolean if a field has been set.

### GetInvProcessor

`func (o *CondHclStatusRelationship) GetInvProcessor() string`

GetInvProcessor returns the InvProcessor field if non-nil, zero value otherwise.

### GetInvProcessorOk

`func (o *CondHclStatusRelationship) GetInvProcessorOk() (*string, bool)`

GetInvProcessorOk returns a tuple with the InvProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvProcessor

`func (o *CondHclStatusRelationship) SetInvProcessor(v string)`

SetInvProcessor sets InvProcessor field to given value.

### HasInvProcessor

`func (o *CondHclStatusRelationship) HasInvProcessor() bool`

HasInvProcessor returns a boolean if a field has been set.

### GetReason

`func (o *CondHclStatusRelationship) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondHclStatusRelationship) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondHclStatusRelationship) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondHclStatusRelationship) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetServerReason

`func (o *CondHclStatusRelationship) GetServerReason() string`

GetServerReason returns the ServerReason field if non-nil, zero value otherwise.

### GetServerReasonOk

`func (o *CondHclStatusRelationship) GetServerReasonOk() (*string, bool)`

GetServerReasonOk returns a tuple with the ServerReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerReason

`func (o *CondHclStatusRelationship) SetServerReason(v string)`

SetServerReason sets ServerReason field to given value.

### HasServerReason

`func (o *CondHclStatusRelationship) HasServerReason() bool`

HasServerReason returns a boolean if a field has been set.

### GetSoftwareStatus

`func (o *CondHclStatusRelationship) GetSoftwareStatus() string`

GetSoftwareStatus returns the SoftwareStatus field if non-nil, zero value otherwise.

### GetSoftwareStatusOk

`func (o *CondHclStatusRelationship) GetSoftwareStatusOk() (*string, bool)`

GetSoftwareStatusOk returns a tuple with the SoftwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareStatus

`func (o *CondHclStatusRelationship) SetSoftwareStatus(v string)`

SetSoftwareStatus sets SoftwareStatus field to given value.

### HasSoftwareStatus

`func (o *CondHclStatusRelationship) HasSoftwareStatus() bool`

HasSoftwareStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CondHclStatusRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondHclStatusRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondHclStatusRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondHclStatusRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *CondHclStatusRelationship) GetDetails() []CondHclStatusDetailRelationship`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CondHclStatusRelationship) GetDetailsOk() (*[]CondHclStatusDetailRelationship, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CondHclStatusRelationship) SetDetails(v []CondHclStatusDetailRelationship)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CondHclStatusRelationship) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *CondHclStatusRelationship) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *CondHclStatusRelationship) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetManagedObject

`func (o *CondHclStatusRelationship) GetManagedObject() InventoryBaseRelationship`

GetManagedObject returns the ManagedObject field if non-nil, zero value otherwise.

### GetManagedObjectOk

`func (o *CondHclStatusRelationship) GetManagedObjectOk() (*InventoryBaseRelationship, bool)`

GetManagedObjectOk returns a tuple with the ManagedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedObject

`func (o *CondHclStatusRelationship) SetManagedObject(v InventoryBaseRelationship)`

SetManagedObject sets ManagedObject field to given value.

### HasManagedObject

`func (o *CondHclStatusRelationship) HasManagedObject() bool`

HasManagedObject returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondHclStatusRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondHclStatusRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondHclStatusRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondHclStatusRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


