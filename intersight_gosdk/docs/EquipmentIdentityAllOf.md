# EquipmentIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AdminAction** | Pointer to **string** | Updated by UI/API to trigger specific chassis action type. * &#x60;None&#x60; - No operation value for maintenance actions on an equipment. * &#x60;Decommission&#x60; - Decommission the equipment and temporarily remove it from being managed by Intersight. * &#x60;Recommission&#x60; - Recommission the equipment. * &#x60;Reack&#x60; - Reacknowledge the equipment and discover it again. * &#x60;Remove&#x60; - Remove the equipment permanently from Intersight management. * &#x60;Replace&#x60; - Replace the equipment with the other one. | [optional] [default to "None"]
**AdminActionState** | Pointer to **string** | The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**Identifier** | Pointer to **int64** | Numeric Identifier assigned by the management system to the equipment. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | The equipment&#39;s lifecycle status. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DecommissionInProgress&#x60; - Decommission Inprogress Lifecycle state. * &#x60;RecommissionInProgress&#x60; - Recommission Inprogress Lifecycle state. * &#x60;OperationFailed&#x60; - Failed Operation Lifecycle state. * &#x60;ReackInProgress&#x60; - ReackInProgress Lifecycle state. * &#x60;RemoveInProgress&#x60; - RemoveInProgress Lifecycle state. * &#x60;Discovered&#x60; - Discovered Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;Inactive&#x60; - Inactive Lifecycle state. * &#x60;ReplaceInProgress&#x60; - ReplaceInProgress Lifecycle state. | [optional] [readonly] [default to "None"]
**Model** | Pointer to **string** | The vendor provided model name for the equipment. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the equipment. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The manufacturer of the equipment. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIdentityAllOf

`func NewEquipmentIdentityAllOf(classId string, objectType string, ) *EquipmentIdentityAllOf`

NewEquipmentIdentityAllOf instantiates a new EquipmentIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIdentityAllOfWithDefaults

`func NewEquipmentIdentityAllOfWithDefaults() *EquipmentIdentityAllOf`

NewEquipmentIdentityAllOfWithDefaults instantiates a new EquipmentIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIdentityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIdentityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIdentityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIdentityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIdentityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIdentityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *EquipmentIdentityAllOf) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *EquipmentIdentityAllOf) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *EquipmentIdentityAllOf) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *EquipmentIdentityAllOf) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminActionState

`func (o *EquipmentIdentityAllOf) GetAdminActionState() string`

GetAdminActionState returns the AdminActionState field if non-nil, zero value otherwise.

### GetAdminActionStateOk

`func (o *EquipmentIdentityAllOf) GetAdminActionStateOk() (*string, bool)`

GetAdminActionStateOk returns a tuple with the AdminActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionState

`func (o *EquipmentIdentityAllOf) SetAdminActionState(v string)`

SetAdminActionState sets AdminActionState field to given value.

### HasAdminActionState

`func (o *EquipmentIdentityAllOf) HasAdminActionState() bool`

HasAdminActionState returns a boolean if a field has been set.

### GetIdentifier

`func (o *EquipmentIdentityAllOf) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EquipmentIdentityAllOf) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EquipmentIdentityAllOf) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EquipmentIdentityAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIdentityAllOf) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIdentityAllOf) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIdentityAllOf) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIdentityAllOf) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIdentityAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIdentityAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIdentityAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIdentityAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIdentityAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIdentityAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIdentityAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIdentityAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIdentityAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIdentityAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIdentityAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIdentityAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentIdentityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentIdentityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentIdentityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentIdentityAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


