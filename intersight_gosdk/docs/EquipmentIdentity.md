# EquipmentIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AdminAction** | Pointer to **string** | Updated by UI/API to trigger specific action type. * &#x60;None&#x60; - No operation value for maintenance actions on an equipment. * &#x60;Decommission&#x60; - Decommission the equipment and temporarily remove it from being managed by Intersight. * &#x60;Recommission&#x60; - Recommission the equipment. * &#x60;Reack&#x60; - Reacknowledge the equipment and discover it again. * &#x60;Remove&#x60; - Remove the equipment permanently from Intersight management. * &#x60;Replace&#x60; - Replace the equipment with the other one. | [optional] [default to "None"]
**AdminActionState** | Pointer to **string** | The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**Identifier** | Pointer to **int64** | Numeric Identifier assigned by the management system to the equipment. Identifier can only be changed if it has been PATCHED with the AdminAction property set to &#39;Recommission&#39;. | [optional] 
**Lifecycle** | Pointer to **string** | The equipment&#39;s lifecycle status. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DecommissionInProgress&#x60; - Decommission Inprogress Lifecycle state. * &#x60;RecommissionInProgress&#x60; - Recommission Inprogress Lifecycle state. * &#x60;OperationFailed&#x60; - Failed Operation Lifecycle state. * &#x60;ReackInProgress&#x60; - ReackInProgress Lifecycle state. * &#x60;RemoveInProgress&#x60; - RemoveInProgress Lifecycle state. * &#x60;Discovered&#x60; - Discovered Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;SecureEraseInProgress&#x60; - Secure Erase is in progress on given physical entity. * &#x60;ScrubInProgress&#x60; - Scrub is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;Inactive&#x60; - Inactive Lifecycle state. * &#x60;ReplaceInProgress&#x60; - ReplaceInProgress Lifecycle state. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. * &#x60;ReplacementPendingUserAction&#x60; - A new blade server is detected, and discovery is pending cleanup of old blade originally discovered in the new blade&#39;s location. * &#x60;Removed&#x60; - The blade server has been removed from its discovered slot, and not detected anywhere else within the domain. * &#x60;Moved&#x60; - The blade server has been moved from its discovered location to a new location within the domain. * &#x60;Replaced&#x60; - The blade server has been removed from its discovered location and another blade has been inserted. in that location. * &#x60;MovedAndReplaced&#x60; - The blade server has been moved from its discovered location to a new location within the domain and another blade has been inserted into the previously discovered location. * &#x60;DomainRmaPendingUserAction&#x60; - Domain RMA detected due to the presence of an old pair of FIs with mismatched serial numbers within the same account. User to either initiate the &#39;Replace Domain workflow&#39; or unclaim the old domain. | [optional] [readonly] [default to "None"]
**LifecycleModTime** | Pointer to **time.Time** | The time when the last life cycle state change happened. | [optional] [readonly] 
**Model** | Pointer to **string** | The vendor provided model name for the equipment. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the equipment for unique identification. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the equipment. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The manufacturer of the equipment. | [optional] [readonly] 
**CustomPermissionResources** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentIdentity

`func NewEquipmentIdentity(classId string, objectType string, ) *EquipmentIdentity`

NewEquipmentIdentity instantiates a new EquipmentIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIdentityWithDefaults

`func NewEquipmentIdentityWithDefaults() *EquipmentIdentity`

NewEquipmentIdentityWithDefaults instantiates a new EquipmentIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *EquipmentIdentity) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *EquipmentIdentity) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *EquipmentIdentity) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *EquipmentIdentity) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminActionState

`func (o *EquipmentIdentity) GetAdminActionState() string`

GetAdminActionState returns the AdminActionState field if non-nil, zero value otherwise.

### GetAdminActionStateOk

`func (o *EquipmentIdentity) GetAdminActionStateOk() (*string, bool)`

GetAdminActionStateOk returns a tuple with the AdminActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionState

`func (o *EquipmentIdentity) SetAdminActionState(v string)`

SetAdminActionState sets AdminActionState field to given value.

### HasAdminActionState

`func (o *EquipmentIdentity) HasAdminActionState() bool`

HasAdminActionState returns a boolean if a field has been set.

### GetIdentifier

`func (o *EquipmentIdentity) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EquipmentIdentity) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EquipmentIdentity) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EquipmentIdentity) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIdentity) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIdentity) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIdentity) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIdentity) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLifecycleModTime

`func (o *EquipmentIdentity) GetLifecycleModTime() time.Time`

GetLifecycleModTime returns the LifecycleModTime field if non-nil, zero value otherwise.

### GetLifecycleModTimeOk

`func (o *EquipmentIdentity) GetLifecycleModTimeOk() (*time.Time, bool)`

GetLifecycleModTimeOk returns a tuple with the LifecycleModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleModTime

`func (o *EquipmentIdentity) SetLifecycleModTime(v time.Time)`

SetLifecycleModTime sets LifecycleModTime field to given value.

### HasLifecycleModTime

`func (o *EquipmentIdentity) HasLifecycleModTime() bool`

HasLifecycleModTime returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIdentity) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIdentity) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIdentity) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIdentity) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *EquipmentIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIdentity) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIdentity) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIdentity) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIdentity) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIdentity) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIdentity) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIdentity) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIdentity) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCustomPermissionResources

`func (o *EquipmentIdentity) GetCustomPermissionResources() []MoBaseMoRelationship`

GetCustomPermissionResources returns the CustomPermissionResources field if non-nil, zero value otherwise.

### GetCustomPermissionResourcesOk

`func (o *EquipmentIdentity) GetCustomPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPermissionResources

`func (o *EquipmentIdentity) SetCustomPermissionResources(v []MoBaseMoRelationship)`

SetCustomPermissionResources sets CustomPermissionResources field to given value.

### HasCustomPermissionResources

`func (o *EquipmentIdentity) HasCustomPermissionResources() bool`

HasCustomPermissionResources returns a boolean if a field has been set.

### SetCustomPermissionResourcesNil

`func (o *EquipmentIdentity) SetCustomPermissionResourcesNil(b bool)`

 SetCustomPermissionResourcesNil sets the value for CustomPermissionResources to be an explicit nil

### UnsetCustomPermissionResources
`func (o *EquipmentIdentity) UnsetCustomPermissionResources()`

UnsetCustomPermissionResources ensures that no value is present for CustomPermissionResources, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentIdentity) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentIdentity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentIdentity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentIdentity) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentIdentity) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentIdentity) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


