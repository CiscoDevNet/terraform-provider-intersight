# CondHclStatusDetailRelationship

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
**HardwareStatus** | Pointer to **string** | The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \&quot;Incompatible-Server-With-Component\&quot; - the server model and component combination is not listed in HCL \&quot;Incompatible-Firmware\&quot; - The server&#39;s firmware is not listed for this component&#39;s hardware profile \&quot;Incompatible-Component\&quot; - the component&#39;s model is not listed in the HCL \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Not-Evaluated\&quot; - the hardware profile was not evaulated for the component because the server&#39;s hw/sw status is not listed or server is exempted. \&quot;Compatible\&quot; - this component&#39;s hardware profile is listed in the HCL. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**HclCimcVersion** | Pointer to **string** | The current CIMC version for the server normalized for querying HCL data. | [optional] 
**HclDriverName** | Pointer to **string** | The current driver name of the component we are validating normalized for querying HCL data. | [optional] 
**HclDriverVersion** | Pointer to **string** | The current driver version of the component we are validating normalized for querying HCL data. | [optional] 
**HclFirmwareVersion** | Pointer to **string** | The current firmware version of the component model normalized for querying HCL data. | [optional] 
**HclModel** | Pointer to **string** | The component model we are trying to validate normalized for querying HCL data. | [optional] 
**InvCimcVersion** | Pointer to **string** | The current CIMC version for the server as received from inventory. | [optional] 
**InvDriverName** | Pointer to **string** | The current driver name of the component we are validating as received from inventory. | [optional] 
**InvDriverVersion** | Pointer to **string** | The current driver version of the component we are validating as received from inventory. | [optional] 
**InvFirmwareVersion** | Pointer to **string** | The current firmware version of the component model as received from inventory. | [optional] 
**InvModel** | Pointer to **string** | The component model we are trying to validate as received from inventory. | [optional] 
**Reason** | Pointer to **string** | The reason for the status. The reason can be one of \&quot;Incompatible-Server-With-Component\&quot; - HCL validation has failed because the server model is not validated with this component \&quot;Incompatible-Processor\&quot; - HCL validation has failed because the processor is not validated with this server \&quot;Incompatible-Os-Info\&quot; - HCL validation has failed because the os vendor and version is not validated with this server \&quot;Incompatible-Component-Model\&quot; - HCL validation has failed because the component model is not validated \&quot;Incompatible-Firmware\&quot; - HCL validation has failed because the component or server firmware version is not validated \&quot;Incompatible-Driver\&quot; - HCL validation has failed because the driver version is not validated \&quot;Incompatible-Firmware-Driver\&quot; - HCL validation has failed because the firmware version and driver version is not validated \&quot;Missing-Os-Driver-Info\&quot; - HCL validation was not performed because we are missing os driver information form the inventory \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Service-Error\&quot; - HCL data service is available but an error occured when making the request or parsing the response \&quot;Unrecognized-Protocol\&quot; - This service does not recognize the reason code in the response from the HCL data service \&quot;Compatible\&quot; - this component&#39;s inventory data has \&quot;Validated\&quot; status with the HCL. \&quot;Not-Evaluated\&quot; - The component is not evaluated against the HCL because the server is exempted. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**SoftwareStatus** | Pointer to **string** | The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \&quot;Incompatible-Firmware\&quot; - the component&#39;s firmware is not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Incompatible-Driver\&quot; - the component&#39;s driver is not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Incompatible-Firmware-Driver\&quot; - the component&#39;s firmware and driver are not listed under the server&#39;s hardware and software profile and the component&#39;s hardware profile \&quot;Service-Unavailable\&quot; - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \&quot;Not-Evaluated\&quot; - the component&#39;s hardware status was not evaluated because the server&#39;s hardware or software profile is not listed or server is exempted. \&quot;Compatible\&quot; - this component&#39;s hardware profile is listed in the HCL. * &#x60;Missing-Os-Driver-Info&#x60; - The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Incompatible-Server-With-Component&#x60; - The validation failed for this component because he server model and component model combination was not found in the HCL. * &#x60;Incompatible-Processor&#x60; - The validation failed because the given processor was not found for the given server PID. * &#x60;Incompatible-Os-Info&#x60; - The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination. * &#x60;Incompatible-Component-Model&#x60; - The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version. * &#x60;Incompatible-Firmware&#x60; - The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model. * &#x60;Incompatible-Driver&#x60; - The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware. * &#x60;Incompatible-Firmware-Driver&#x60; - The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware. * &#x60;Service-Unavailable&#x60; - The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data. * &#x60;Service-Error&#x60; - The validation has failed because the HCL data service has return a service error or unrecognized result. * &#x60;Unrecognized-Protocol&#x60; - The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values. * &#x60;Not-Evaluated&#x60; - The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server&#39;s hardware profile fails to validate with HCL, then the server&#39;s software status will not be evaluated. * &#x60;Compatible&#x60; - The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version. | [optional] [default to "Missing-Os-Driver-Info"]
**Status** | Pointer to **string** | The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \&quot;Unknown\&quot; - we do not have enough information to evaluate against the HCL data \&quot;Validated\&quot; - we have validated this component against the HCL and it has \&quot;Validated\&quot; status \&quot;Not-Validated\&quot; - we have validated this component against the HCL and it has \&quot;Not-Validated\&quot; status. \&quot;Not-Evaluated\&quot; - The component is not evaluated against the HCL because the server is exempted. * &#x60;Incomplete&#x60; - This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information. * &#x60;Not-Found&#x60; - At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component&#39;s hardware or software profile was not found in the HCL. * &#x60;Not-Listed&#x60; - At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server&#39;s hardware or software profile was not listed in the HCL or one of the components&#39; hardware or software profile was not found in the HCL. * &#x60;Validated&#x60; - At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component&#39;s hardware or software profile was found in the HCL. * &#x60;Not-Evaluated&#x60; - At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet. | [optional] [default to "Incomplete"]
**Component** | Pointer to [**InventoryBaseRelationship**](inventory.Base.Relationship.md) |  | [optional] 
**HclStatus** | Pointer to [**CondHclStatusRelationship**](cond.HclStatus.Relationship.md) |  | [optional] 

## Methods

### NewCondHclStatusDetailRelationship

`func NewCondHclStatusDetailRelationship(classId string, objectType string, ) *CondHclStatusDetailRelationship`

NewCondHclStatusDetailRelationship instantiates a new CondHclStatusDetailRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondHclStatusDetailRelationshipWithDefaults

`func NewCondHclStatusDetailRelationshipWithDefaults() *CondHclStatusDetailRelationship`

NewCondHclStatusDetailRelationshipWithDefaults instantiates a new CondHclStatusDetailRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *CondHclStatusDetailRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *CondHclStatusDetailRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *CondHclStatusDetailRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *CondHclStatusDetailRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *CondHclStatusDetailRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondHclStatusDetailRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondHclStatusDetailRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *CondHclStatusDetailRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CondHclStatusDetailRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CondHclStatusDetailRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CondHclStatusDetailRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *CondHclStatusDetailRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *CondHclStatusDetailRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *CondHclStatusDetailRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *CondHclStatusDetailRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *CondHclStatusDetailRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *CondHclStatusDetailRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *CondHclStatusDetailRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *CondHclStatusDetailRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *CondHclStatusDetailRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *CondHclStatusDetailRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *CondHclStatusDetailRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *CondHclStatusDetailRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *CondHclStatusDetailRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondHclStatusDetailRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondHclStatusDetailRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *CondHclStatusDetailRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *CondHclStatusDetailRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *CondHclStatusDetailRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *CondHclStatusDetailRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *CondHclStatusDetailRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *CondHclStatusDetailRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *CondHclStatusDetailRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *CondHclStatusDetailRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *CondHclStatusDetailRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CondHclStatusDetailRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CondHclStatusDetailRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CondHclStatusDetailRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *CondHclStatusDetailRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *CondHclStatusDetailRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *CondHclStatusDetailRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *CondHclStatusDetailRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *CondHclStatusDetailRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *CondHclStatusDetailRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *CondHclStatusDetailRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *CondHclStatusDetailRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *CondHclStatusDetailRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *CondHclStatusDetailRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *CondHclStatusDetailRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CondHclStatusDetailRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CondHclStatusDetailRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CondHclStatusDetailRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *CondHclStatusDetailRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *CondHclStatusDetailRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *CondHclStatusDetailRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *CondHclStatusDetailRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *CondHclStatusDetailRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *CondHclStatusDetailRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *CondHclStatusDetailRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *CondHclStatusDetailRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *CondHclStatusDetailRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *CondHclStatusDetailRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *CondHclStatusDetailRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *CondHclStatusDetailRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetHardwareStatus

`func (o *CondHclStatusDetailRelationship) GetHardwareStatus() string`

GetHardwareStatus returns the HardwareStatus field if non-nil, zero value otherwise.

### GetHardwareStatusOk

`func (o *CondHclStatusDetailRelationship) GetHardwareStatusOk() (*string, bool)`

GetHardwareStatusOk returns a tuple with the HardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareStatus

`func (o *CondHclStatusDetailRelationship) SetHardwareStatus(v string)`

SetHardwareStatus sets HardwareStatus field to given value.

### HasHardwareStatus

`func (o *CondHclStatusDetailRelationship) HasHardwareStatus() bool`

HasHardwareStatus returns a boolean if a field has been set.

### GetHclCimcVersion

`func (o *CondHclStatusDetailRelationship) GetHclCimcVersion() string`

GetHclCimcVersion returns the HclCimcVersion field if non-nil, zero value otherwise.

### GetHclCimcVersionOk

`func (o *CondHclStatusDetailRelationship) GetHclCimcVersionOk() (*string, bool)`

GetHclCimcVersionOk returns a tuple with the HclCimcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclCimcVersion

`func (o *CondHclStatusDetailRelationship) SetHclCimcVersion(v string)`

SetHclCimcVersion sets HclCimcVersion field to given value.

### HasHclCimcVersion

`func (o *CondHclStatusDetailRelationship) HasHclCimcVersion() bool`

HasHclCimcVersion returns a boolean if a field has been set.

### GetHclDriverName

`func (o *CondHclStatusDetailRelationship) GetHclDriverName() string`

GetHclDriverName returns the HclDriverName field if non-nil, zero value otherwise.

### GetHclDriverNameOk

`func (o *CondHclStatusDetailRelationship) GetHclDriverNameOk() (*string, bool)`

GetHclDriverNameOk returns a tuple with the HclDriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclDriverName

`func (o *CondHclStatusDetailRelationship) SetHclDriverName(v string)`

SetHclDriverName sets HclDriverName field to given value.

### HasHclDriverName

`func (o *CondHclStatusDetailRelationship) HasHclDriverName() bool`

HasHclDriverName returns a boolean if a field has been set.

### GetHclDriverVersion

`func (o *CondHclStatusDetailRelationship) GetHclDriverVersion() string`

GetHclDriverVersion returns the HclDriverVersion field if non-nil, zero value otherwise.

### GetHclDriverVersionOk

`func (o *CondHclStatusDetailRelationship) GetHclDriverVersionOk() (*string, bool)`

GetHclDriverVersionOk returns a tuple with the HclDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclDriverVersion

`func (o *CondHclStatusDetailRelationship) SetHclDriverVersion(v string)`

SetHclDriverVersion sets HclDriverVersion field to given value.

### HasHclDriverVersion

`func (o *CondHclStatusDetailRelationship) HasHclDriverVersion() bool`

HasHclDriverVersion returns a boolean if a field has been set.

### GetHclFirmwareVersion

`func (o *CondHclStatusDetailRelationship) GetHclFirmwareVersion() string`

GetHclFirmwareVersion returns the HclFirmwareVersion field if non-nil, zero value otherwise.

### GetHclFirmwareVersionOk

`func (o *CondHclStatusDetailRelationship) GetHclFirmwareVersionOk() (*string, bool)`

GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclFirmwareVersion

`func (o *CondHclStatusDetailRelationship) SetHclFirmwareVersion(v string)`

SetHclFirmwareVersion sets HclFirmwareVersion field to given value.

### HasHclFirmwareVersion

`func (o *CondHclStatusDetailRelationship) HasHclFirmwareVersion() bool`

HasHclFirmwareVersion returns a boolean if a field has been set.

### GetHclModel

`func (o *CondHclStatusDetailRelationship) GetHclModel() string`

GetHclModel returns the HclModel field if non-nil, zero value otherwise.

### GetHclModelOk

`func (o *CondHclStatusDetailRelationship) GetHclModelOk() (*string, bool)`

GetHclModelOk returns a tuple with the HclModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclModel

`func (o *CondHclStatusDetailRelationship) SetHclModel(v string)`

SetHclModel sets HclModel field to given value.

### HasHclModel

`func (o *CondHclStatusDetailRelationship) HasHclModel() bool`

HasHclModel returns a boolean if a field has been set.

### GetInvCimcVersion

`func (o *CondHclStatusDetailRelationship) GetInvCimcVersion() string`

GetInvCimcVersion returns the InvCimcVersion field if non-nil, zero value otherwise.

### GetInvCimcVersionOk

`func (o *CondHclStatusDetailRelationship) GetInvCimcVersionOk() (*string, bool)`

GetInvCimcVersionOk returns a tuple with the InvCimcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvCimcVersion

`func (o *CondHclStatusDetailRelationship) SetInvCimcVersion(v string)`

SetInvCimcVersion sets InvCimcVersion field to given value.

### HasInvCimcVersion

`func (o *CondHclStatusDetailRelationship) HasInvCimcVersion() bool`

HasInvCimcVersion returns a boolean if a field has been set.

### GetInvDriverName

`func (o *CondHclStatusDetailRelationship) GetInvDriverName() string`

GetInvDriverName returns the InvDriverName field if non-nil, zero value otherwise.

### GetInvDriverNameOk

`func (o *CondHclStatusDetailRelationship) GetInvDriverNameOk() (*string, bool)`

GetInvDriverNameOk returns a tuple with the InvDriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvDriverName

`func (o *CondHclStatusDetailRelationship) SetInvDriverName(v string)`

SetInvDriverName sets InvDriverName field to given value.

### HasInvDriverName

`func (o *CondHclStatusDetailRelationship) HasInvDriverName() bool`

HasInvDriverName returns a boolean if a field has been set.

### GetInvDriverVersion

`func (o *CondHclStatusDetailRelationship) GetInvDriverVersion() string`

GetInvDriverVersion returns the InvDriverVersion field if non-nil, zero value otherwise.

### GetInvDriverVersionOk

`func (o *CondHclStatusDetailRelationship) GetInvDriverVersionOk() (*string, bool)`

GetInvDriverVersionOk returns a tuple with the InvDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvDriverVersion

`func (o *CondHclStatusDetailRelationship) SetInvDriverVersion(v string)`

SetInvDriverVersion sets InvDriverVersion field to given value.

### HasInvDriverVersion

`func (o *CondHclStatusDetailRelationship) HasInvDriverVersion() bool`

HasInvDriverVersion returns a boolean if a field has been set.

### GetInvFirmwareVersion

`func (o *CondHclStatusDetailRelationship) GetInvFirmwareVersion() string`

GetInvFirmwareVersion returns the InvFirmwareVersion field if non-nil, zero value otherwise.

### GetInvFirmwareVersionOk

`func (o *CondHclStatusDetailRelationship) GetInvFirmwareVersionOk() (*string, bool)`

GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvFirmwareVersion

`func (o *CondHclStatusDetailRelationship) SetInvFirmwareVersion(v string)`

SetInvFirmwareVersion sets InvFirmwareVersion field to given value.

### HasInvFirmwareVersion

`func (o *CondHclStatusDetailRelationship) HasInvFirmwareVersion() bool`

HasInvFirmwareVersion returns a boolean if a field has been set.

### GetInvModel

`func (o *CondHclStatusDetailRelationship) GetInvModel() string`

GetInvModel returns the InvModel field if non-nil, zero value otherwise.

### GetInvModelOk

`func (o *CondHclStatusDetailRelationship) GetInvModelOk() (*string, bool)`

GetInvModelOk returns a tuple with the InvModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvModel

`func (o *CondHclStatusDetailRelationship) SetInvModel(v string)`

SetInvModel sets InvModel field to given value.

### HasInvModel

`func (o *CondHclStatusDetailRelationship) HasInvModel() bool`

HasInvModel returns a boolean if a field has been set.

### GetReason

`func (o *CondHclStatusDetailRelationship) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CondHclStatusDetailRelationship) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CondHclStatusDetailRelationship) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CondHclStatusDetailRelationship) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSoftwareStatus

`func (o *CondHclStatusDetailRelationship) GetSoftwareStatus() string`

GetSoftwareStatus returns the SoftwareStatus field if non-nil, zero value otherwise.

### GetSoftwareStatusOk

`func (o *CondHclStatusDetailRelationship) GetSoftwareStatusOk() (*string, bool)`

GetSoftwareStatusOk returns a tuple with the SoftwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareStatus

`func (o *CondHclStatusDetailRelationship) SetSoftwareStatus(v string)`

SetSoftwareStatus sets SoftwareStatus field to given value.

### HasSoftwareStatus

`func (o *CondHclStatusDetailRelationship) HasSoftwareStatus() bool`

HasSoftwareStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CondHclStatusDetailRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondHclStatusDetailRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondHclStatusDetailRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondHclStatusDetailRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponent

`func (o *CondHclStatusDetailRelationship) GetComponent() InventoryBaseRelationship`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CondHclStatusDetailRelationship) GetComponentOk() (*InventoryBaseRelationship, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CondHclStatusDetailRelationship) SetComponent(v InventoryBaseRelationship)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CondHclStatusDetailRelationship) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetHclStatus

`func (o *CondHclStatusDetailRelationship) GetHclStatus() CondHclStatusRelationship`

GetHclStatus returns the HclStatus field if non-nil, zero value otherwise.

### GetHclStatusOk

`func (o *CondHclStatusDetailRelationship) GetHclStatusOk() (*CondHclStatusRelationship, bool)`

GetHclStatusOk returns a tuple with the HclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclStatus

`func (o *CondHclStatusDetailRelationship) SetHclStatus(v CondHclStatusRelationship)`

SetHclStatus sets HclStatus field to given value.

### HasHclStatus

`func (o *CondHclStatusDetailRelationship) HasHclStatus() bool`

HasHclStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


