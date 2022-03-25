# EquipmentIdentitySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.IdentitySummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.IdentitySummary"]
**AdapterSerial** | Pointer to **string** | Serial Identifier of an adapter participating in SWM. | [optional] [readonly] 
**AdminAction** | Pointer to **string** | Updated by UI/API to trigger specific chassis action type. * &#x60;None&#x60; - No operation value for maintenance actions on an equipment. * &#x60;Decommission&#x60; - Decommission the equipment and temporarily remove it from being managed by Intersight. * &#x60;Recommission&#x60; - Recommission the equipment. * &#x60;Reack&#x60; - Reacknowledge the equipment and discover it again. * &#x60;Remove&#x60; - Remove the equipment permanently from Intersight management. * &#x60;Replace&#x60; - Replace the equipment with the other one. | [optional] [readonly] [default to "None"]
**AdminActionState** | Pointer to **string** | The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**ChassisId** | Pointer to **int64** | Chassis Identifier of a blade server. | [optional] [readonly] 
**CurrentChassisId** | Pointer to **int64** | The id of the chassis that the blade is currently located in. | [optional] [readonly] 
**CurrentSlotId** | Pointer to **int64** | The slot number in the chassis that the blade is currently located in. | [optional] [readonly] 
**FirmwareSupportability** | Pointer to **string** | Describes whether the running CIMC version supports Intersight managed mode. * &#x60;Unknown&#x60; - The running firmware version is unknown. * &#x60;Supported&#x60; - The running firmware version is known and supports IMM mode. * &#x60;NotSupported&#x60; - The running firmware version is known and does not support IMM mode. | [optional] [readonly] [default to "Unknown"]
**Identifier** | Pointer to **int64** | Numeric Identifier assigned by the management system to the equipment. | [optional] [readonly] 
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](EquipmentIoCardIdentity.md) |  | [optional] 
**LastDiscoveryTriggered** | Pointer to **string** | Denotes the type of discovery that was most recently triggered on this server. * &#x60;Unknown&#x60; - The last discovery type is unknown. * &#x60;Deep&#x60; - The last discovery triggered is deep. * &#x60;Shallow&#x60; - The last discovery triggered is shallow. | [optional] [readonly] [default to "Unknown"]
**Lifecycle** | Pointer to **string** | The equipment&#39;s lifecycle status. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DecommissionInProgress&#x60; - Decommission Inprogress Lifecycle state. * &#x60;RecommissionInProgress&#x60; - Recommission Inprogress Lifecycle state. * &#x60;OperationFailed&#x60; - Failed Operation Lifecycle state. * &#x60;ReackInProgress&#x60; - ReackInProgress Lifecycle state. * &#x60;RemoveInProgress&#x60; - RemoveInProgress Lifecycle state. * &#x60;Discovered&#x60; - Discovered Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. * &#x60;FirmwareUpgradeInProgress&#x60; - Firmware upgrade is in progress on given physical entity. * &#x60;BladeMigrationInProgress&#x60; - Server slot migration is in progress on given physical entity. * &#x60;Inactive&#x60; - Inactive Lifecycle state. * &#x60;ReplaceInProgress&#x60; - ReplaceInProgress Lifecycle state. * &#x60;SlotMismatch&#x60; - The blade server is detected in a different chassis/slot than it was previously. | [optional] [readonly] [default to "None"]
**Model** | Pointer to **string** | The vendor provided model name for the equipment. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence state of the blade server. * &#x60;Unknown&#x60; - The default presence state. * &#x60;Equipped&#x60; - The server is equipped in the slot. * &#x60;EquippedMismatch&#x60; - The slot is equipped, but there is another server currently inventoried in the slot. * &#x60;Missing&#x60; - The server is not present in the given slot. | [optional] [readonly] [default to "Unknown"]
**Serial** | Pointer to **string** | The serial number of the equipment. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Chassis slot number of a blade server. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of this view MO. | [optional] [readonly] 
**SwitchId** | Pointer to **int64** | Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The manufacturer of the equipment. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentIdentitySummary

`func NewEquipmentIdentitySummary(classId string, objectType string, ) *EquipmentIdentitySummary`

NewEquipmentIdentitySummary instantiates a new EquipmentIdentitySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIdentitySummaryWithDefaults

`func NewEquipmentIdentitySummaryWithDefaults() *EquipmentIdentitySummary`

NewEquipmentIdentitySummaryWithDefaults instantiates a new EquipmentIdentitySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIdentitySummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIdentitySummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIdentitySummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIdentitySummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIdentitySummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIdentitySummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterSerial

`func (o *EquipmentIdentitySummary) GetAdapterSerial() string`

GetAdapterSerial returns the AdapterSerial field if non-nil, zero value otherwise.

### GetAdapterSerialOk

`func (o *EquipmentIdentitySummary) GetAdapterSerialOk() (*string, bool)`

GetAdapterSerialOk returns a tuple with the AdapterSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterSerial

`func (o *EquipmentIdentitySummary) SetAdapterSerial(v string)`

SetAdapterSerial sets AdapterSerial field to given value.

### HasAdapterSerial

`func (o *EquipmentIdentitySummary) HasAdapterSerial() bool`

HasAdapterSerial returns a boolean if a field has been set.

### GetAdminAction

`func (o *EquipmentIdentitySummary) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *EquipmentIdentitySummary) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *EquipmentIdentitySummary) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *EquipmentIdentitySummary) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminActionState

`func (o *EquipmentIdentitySummary) GetAdminActionState() string`

GetAdminActionState returns the AdminActionState field if non-nil, zero value otherwise.

### GetAdminActionStateOk

`func (o *EquipmentIdentitySummary) GetAdminActionStateOk() (*string, bool)`

GetAdminActionStateOk returns a tuple with the AdminActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionState

`func (o *EquipmentIdentitySummary) SetAdminActionState(v string)`

SetAdminActionState sets AdminActionState field to given value.

### HasAdminActionState

`func (o *EquipmentIdentitySummary) HasAdminActionState() bool`

HasAdminActionState returns a boolean if a field has been set.

### GetChassisId

`func (o *EquipmentIdentitySummary) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentIdentitySummary) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentIdentitySummary) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentIdentitySummary) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetCurrentChassisId

`func (o *EquipmentIdentitySummary) GetCurrentChassisId() int64`

GetCurrentChassisId returns the CurrentChassisId field if non-nil, zero value otherwise.

### GetCurrentChassisIdOk

`func (o *EquipmentIdentitySummary) GetCurrentChassisIdOk() (*int64, bool)`

GetCurrentChassisIdOk returns a tuple with the CurrentChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChassisId

`func (o *EquipmentIdentitySummary) SetCurrentChassisId(v int64)`

SetCurrentChassisId sets CurrentChassisId field to given value.

### HasCurrentChassisId

`func (o *EquipmentIdentitySummary) HasCurrentChassisId() bool`

HasCurrentChassisId returns a boolean if a field has been set.

### GetCurrentSlotId

`func (o *EquipmentIdentitySummary) GetCurrentSlotId() int64`

GetCurrentSlotId returns the CurrentSlotId field if non-nil, zero value otherwise.

### GetCurrentSlotIdOk

`func (o *EquipmentIdentitySummary) GetCurrentSlotIdOk() (*int64, bool)`

GetCurrentSlotIdOk returns a tuple with the CurrentSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSlotId

`func (o *EquipmentIdentitySummary) SetCurrentSlotId(v int64)`

SetCurrentSlotId sets CurrentSlotId field to given value.

### HasCurrentSlotId

`func (o *EquipmentIdentitySummary) HasCurrentSlotId() bool`

HasCurrentSlotId returns a boolean if a field has been set.

### GetFirmwareSupportability

`func (o *EquipmentIdentitySummary) GetFirmwareSupportability() string`

GetFirmwareSupportability returns the FirmwareSupportability field if non-nil, zero value otherwise.

### GetFirmwareSupportabilityOk

`func (o *EquipmentIdentitySummary) GetFirmwareSupportabilityOk() (*string, bool)`

GetFirmwareSupportabilityOk returns a tuple with the FirmwareSupportability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareSupportability

`func (o *EquipmentIdentitySummary) SetFirmwareSupportability(v string)`

SetFirmwareSupportability sets FirmwareSupportability field to given value.

### HasFirmwareSupportability

`func (o *EquipmentIdentitySummary) HasFirmwareSupportability() bool`

HasFirmwareSupportability returns a boolean if a field has been set.

### GetIdentifier

`func (o *EquipmentIdentitySummary) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EquipmentIdentitySummary) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EquipmentIdentitySummary) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EquipmentIdentitySummary) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIoCardIdentityList

`func (o *EquipmentIdentitySummary) GetIoCardIdentityList() []EquipmentIoCardIdentity`

GetIoCardIdentityList returns the IoCardIdentityList field if non-nil, zero value otherwise.

### GetIoCardIdentityListOk

`func (o *EquipmentIdentitySummary) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool)`

GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardIdentityList

`func (o *EquipmentIdentitySummary) SetIoCardIdentityList(v []EquipmentIoCardIdentity)`

SetIoCardIdentityList sets IoCardIdentityList field to given value.

### HasIoCardIdentityList

`func (o *EquipmentIdentitySummary) HasIoCardIdentityList() bool`

HasIoCardIdentityList returns a boolean if a field has been set.

### SetIoCardIdentityListNil

`func (o *EquipmentIdentitySummary) SetIoCardIdentityListNil(b bool)`

 SetIoCardIdentityListNil sets the value for IoCardIdentityList to be an explicit nil

### UnsetIoCardIdentityList
`func (o *EquipmentIdentitySummary) UnsetIoCardIdentityList()`

UnsetIoCardIdentityList ensures that no value is present for IoCardIdentityList, not even an explicit nil
### GetLastDiscoveryTriggered

`func (o *EquipmentIdentitySummary) GetLastDiscoveryTriggered() string`

GetLastDiscoveryTriggered returns the LastDiscoveryTriggered field if non-nil, zero value otherwise.

### GetLastDiscoveryTriggeredOk

`func (o *EquipmentIdentitySummary) GetLastDiscoveryTriggeredOk() (*string, bool)`

GetLastDiscoveryTriggeredOk returns a tuple with the LastDiscoveryTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveryTriggered

`func (o *EquipmentIdentitySummary) SetLastDiscoveryTriggered(v string)`

SetLastDiscoveryTriggered sets LastDiscoveryTriggered field to given value.

### HasLastDiscoveryTriggered

`func (o *EquipmentIdentitySummary) HasLastDiscoveryTriggered() bool`

HasLastDiscoveryTriggered returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIdentitySummary) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIdentitySummary) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIdentitySummary) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIdentitySummary) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIdentitySummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIdentitySummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIdentitySummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIdentitySummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentIdentitySummary) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentIdentitySummary) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentIdentitySummary) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentIdentitySummary) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIdentitySummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIdentitySummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIdentitySummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIdentitySummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentIdentitySummary) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentIdentitySummary) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentIdentitySummary) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentIdentitySummary) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *EquipmentIdentitySummary) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *EquipmentIdentitySummary) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *EquipmentIdentitySummary) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *EquipmentIdentitySummary) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentIdentitySummary) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentIdentitySummary) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentIdentitySummary) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentIdentitySummary) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIdentitySummary) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIdentitySummary) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIdentitySummary) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIdentitySummary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentIdentitySummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentIdentitySummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentIdentitySummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentIdentitySummary) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


