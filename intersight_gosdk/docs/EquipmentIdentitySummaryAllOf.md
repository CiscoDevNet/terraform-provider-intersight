# EquipmentIdentitySummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdapterSerial** | Pointer to **string** | Serial Identifier of an adapter participating in SWM. | [optional] [readonly] 
**AdminAction** | Pointer to **string** | Updated by UI/API to trigger specific chassis action type. * &#x60;None&#x60; - No operation value for maintenance actions on an equipment. * &#x60;Decommission&#x60; - Decommission the equipment and temporarily remove it from being managed by Intersight. * &#x60;Recommission&#x60; - Recommission the equipment. * &#x60;Reack&#x60; - Reacknowledge the equipment and discover it again. * &#x60;Remove&#x60; - Remove the equipment permanently from Intersight management. | [optional] [readonly] [default to "None"]
**AdminActionState** | Pointer to **string** | The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**ChassisId** | Pointer to **int64** | Chassis Identifier of a blade server. | [optional] [readonly] 
**DeviceMoId** | Pointer to **string** | FI Device registration Mo ID. | [optional] [readonly] 
**Identifier** | Pointer to **int64** | Numeric Identifier assigned by the management system to the equipment. | [optional] [readonly] 
**IoCardIdentityList** | Pointer to [**[]EquipmentIoCardIdentity**](equipment.IoCardIdentity.md) |  | [optional] 
**Lifecycle** | Pointer to **string** | The equipment&#39;s lifecycle status. * &#x60;None&#x60; - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * &#x60;Active&#x60; - Default Lifecycle State for a physical entity. * &#x60;Decommissioned&#x60; - Decommission Lifecycle state. * &#x60;DecommissionInProgress&#x60; - Decommission Inprogress Lifecycle state. * &#x60;RecommissionInProgress&#x60; - Recommission Inprogress Lifecycle state. * &#x60;OperationFailed&#x60; - Failed Operation Lifecycle state. * &#x60;ReackInProgress&#x60; - ReackInProgress Lifecycle state. * &#x60;RemoveInProgress&#x60; - RemoveInProgress Lifecycle state. * &#x60;Discovered&#x60; - Discovered Lifecycle state. * &#x60;DiscoveryInProgress&#x60; - DiscoveryInProgress Lifecycle state. * &#x60;DiscoveryFailed&#x60; - DiscoveryFailed Lifecycle state. | [optional] [readonly] [default to "None"]
**Model** | Pointer to **string** | The vendor provided model name for the equipment. | [optional] [readonly] 
**PendingDiscovery** | Pointer to **string** | Indicates pending discovery flag. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the equipment. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Chassis slot number of a blade server. | [optional] [readonly] 
**SourceObjectType** | Pointer to **string** | The source object type of this view MO. | [optional] [readonly] 
**SwitchId** | Pointer to **int64** | Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The manufacturer of the equipment. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIdentitySummaryAllOf

`func NewEquipmentIdentitySummaryAllOf() *EquipmentIdentitySummaryAllOf`

NewEquipmentIdentitySummaryAllOf instantiates a new EquipmentIdentitySummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIdentitySummaryAllOfWithDefaults

`func NewEquipmentIdentitySummaryAllOfWithDefaults() *EquipmentIdentitySummaryAllOf`

NewEquipmentIdentitySummaryAllOfWithDefaults instantiates a new EquipmentIdentitySummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapterSerial

`func (o *EquipmentIdentitySummaryAllOf) GetAdapterSerial() string`

GetAdapterSerial returns the AdapterSerial field if non-nil, zero value otherwise.

### GetAdapterSerialOk

`func (o *EquipmentIdentitySummaryAllOf) GetAdapterSerialOk() (*string, bool)`

GetAdapterSerialOk returns a tuple with the AdapterSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterSerial

`func (o *EquipmentIdentitySummaryAllOf) SetAdapterSerial(v string)`

SetAdapterSerial sets AdapterSerial field to given value.

### HasAdapterSerial

`func (o *EquipmentIdentitySummaryAllOf) HasAdapterSerial() bool`

HasAdapterSerial returns a boolean if a field has been set.

### GetAdminAction

`func (o *EquipmentIdentitySummaryAllOf) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *EquipmentIdentitySummaryAllOf) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *EquipmentIdentitySummaryAllOf) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *EquipmentIdentitySummaryAllOf) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminActionState

`func (o *EquipmentIdentitySummaryAllOf) GetAdminActionState() string`

GetAdminActionState returns the AdminActionState field if non-nil, zero value otherwise.

### GetAdminActionStateOk

`func (o *EquipmentIdentitySummaryAllOf) GetAdminActionStateOk() (*string, bool)`

GetAdminActionStateOk returns a tuple with the AdminActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminActionState

`func (o *EquipmentIdentitySummaryAllOf) SetAdminActionState(v string)`

SetAdminActionState sets AdminActionState field to given value.

### HasAdminActionState

`func (o *EquipmentIdentitySummaryAllOf) HasAdminActionState() bool`

HasAdminActionState returns a boolean if a field has been set.

### GetChassisId

`func (o *EquipmentIdentitySummaryAllOf) GetChassisId() int64`

GetChassisId returns the ChassisId field if non-nil, zero value otherwise.

### GetChassisIdOk

`func (o *EquipmentIdentitySummaryAllOf) GetChassisIdOk() (*int64, bool)`

GetChassisIdOk returns a tuple with the ChassisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisId

`func (o *EquipmentIdentitySummaryAllOf) SetChassisId(v int64)`

SetChassisId sets ChassisId field to given value.

### HasChassisId

`func (o *EquipmentIdentitySummaryAllOf) HasChassisId() bool`

HasChassisId returns a boolean if a field has been set.

### GetDeviceMoId

`func (o *EquipmentIdentitySummaryAllOf) GetDeviceMoId() string`

GetDeviceMoId returns the DeviceMoId field if non-nil, zero value otherwise.

### GetDeviceMoIdOk

`func (o *EquipmentIdentitySummaryAllOf) GetDeviceMoIdOk() (*string, bool)`

GetDeviceMoIdOk returns a tuple with the DeviceMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMoId

`func (o *EquipmentIdentitySummaryAllOf) SetDeviceMoId(v string)`

SetDeviceMoId sets DeviceMoId field to given value.

### HasDeviceMoId

`func (o *EquipmentIdentitySummaryAllOf) HasDeviceMoId() bool`

HasDeviceMoId returns a boolean if a field has been set.

### GetIdentifier

`func (o *EquipmentIdentitySummaryAllOf) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *EquipmentIdentitySummaryAllOf) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *EquipmentIdentitySummaryAllOf) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *EquipmentIdentitySummaryAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIoCardIdentityList

`func (o *EquipmentIdentitySummaryAllOf) GetIoCardIdentityList() []EquipmentIoCardIdentity`

GetIoCardIdentityList returns the IoCardIdentityList field if non-nil, zero value otherwise.

### GetIoCardIdentityListOk

`func (o *EquipmentIdentitySummaryAllOf) GetIoCardIdentityListOk() (*[]EquipmentIoCardIdentity, bool)`

GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardIdentityList

`func (o *EquipmentIdentitySummaryAllOf) SetIoCardIdentityList(v []EquipmentIoCardIdentity)`

SetIoCardIdentityList sets IoCardIdentityList field to given value.

### HasIoCardIdentityList

`func (o *EquipmentIdentitySummaryAllOf) HasIoCardIdentityList() bool`

HasIoCardIdentityList returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIdentitySummaryAllOf) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIdentitySummaryAllOf) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIdentitySummaryAllOf) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIdentitySummaryAllOf) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIdentitySummaryAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIdentitySummaryAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIdentitySummaryAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIdentitySummaryAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPendingDiscovery

`func (o *EquipmentIdentitySummaryAllOf) GetPendingDiscovery() string`

GetPendingDiscovery returns the PendingDiscovery field if non-nil, zero value otherwise.

### GetPendingDiscoveryOk

`func (o *EquipmentIdentitySummaryAllOf) GetPendingDiscoveryOk() (*string, bool)`

GetPendingDiscoveryOk returns a tuple with the PendingDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDiscovery

`func (o *EquipmentIdentitySummaryAllOf) SetPendingDiscovery(v string)`

SetPendingDiscovery sets PendingDiscovery field to given value.

### HasPendingDiscovery

`func (o *EquipmentIdentitySummaryAllOf) HasPendingDiscovery() bool`

HasPendingDiscovery returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIdentitySummaryAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIdentitySummaryAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIdentitySummaryAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIdentitySummaryAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentIdentitySummaryAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentIdentitySummaryAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentIdentitySummaryAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentIdentitySummaryAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSourceObjectType

`func (o *EquipmentIdentitySummaryAllOf) GetSourceObjectType() string`

GetSourceObjectType returns the SourceObjectType field if non-nil, zero value otherwise.

### GetSourceObjectTypeOk

`func (o *EquipmentIdentitySummaryAllOf) GetSourceObjectTypeOk() (*string, bool)`

GetSourceObjectTypeOk returns a tuple with the SourceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectType

`func (o *EquipmentIdentitySummaryAllOf) SetSourceObjectType(v string)`

SetSourceObjectType sets SourceObjectType field to given value.

### HasSourceObjectType

`func (o *EquipmentIdentitySummaryAllOf) HasSourceObjectType() bool`

HasSourceObjectType returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentIdentitySummaryAllOf) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentIdentitySummaryAllOf) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentIdentitySummaryAllOf) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentIdentitySummaryAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIdentitySummaryAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIdentitySummaryAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIdentitySummaryAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIdentitySummaryAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentIdentitySummaryAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentIdentitySummaryAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentIdentitySummaryAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentIdentitySummaryAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


