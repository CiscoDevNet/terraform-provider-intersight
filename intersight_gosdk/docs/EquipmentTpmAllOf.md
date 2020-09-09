# EquipmentTpmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationStatus** | Pointer to **string** | Identifies the activation status of the TPM. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Identifies the admin configured state of the TPM. | [optional] [readonly] 
**Ownership** | Pointer to **string** | Identifies the ownership information of the TPM. | [optional] [readonly] 
**Presence** | Pointer to **string** | Identifies the presence of the trusted platform module. | [optional] [readonly] 
**TpmId** | Pointer to **int64** | Enter  the ID of the trusted platform module. | [optional] [readonly] 
**Version** | Pointer to **string** | Identifies the revision of the Trusted Platform Module. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentTpmAllOf

`func NewEquipmentTpmAllOf() *EquipmentTpmAllOf`

NewEquipmentTpmAllOf instantiates a new EquipmentTpmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentTpmAllOfWithDefaults

`func NewEquipmentTpmAllOfWithDefaults() *EquipmentTpmAllOf`

NewEquipmentTpmAllOfWithDefaults instantiates a new EquipmentTpmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationStatus

`func (o *EquipmentTpmAllOf) GetActivationStatus() string`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *EquipmentTpmAllOf) GetActivationStatusOk() (*string, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *EquipmentTpmAllOf) SetActivationStatus(v string)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *EquipmentTpmAllOf) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetAdminState

`func (o *EquipmentTpmAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EquipmentTpmAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EquipmentTpmAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EquipmentTpmAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetOwnership

`func (o *EquipmentTpmAllOf) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *EquipmentTpmAllOf) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *EquipmentTpmAllOf) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *EquipmentTpmAllOf) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentTpmAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentTpmAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentTpmAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentTpmAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetTpmId

`func (o *EquipmentTpmAllOf) GetTpmId() int64`

GetTpmId returns the TpmId field if non-nil, zero value otherwise.

### GetTpmIdOk

`func (o *EquipmentTpmAllOf) GetTpmIdOk() (*int64, bool)`

GetTpmIdOk returns a tuple with the TpmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmId

`func (o *EquipmentTpmAllOf) SetTpmId(v int64)`

SetTpmId sets TpmId field to given value.

### HasTpmId

`func (o *EquipmentTpmAllOf) HasTpmId() bool`

HasTpmId returns a boolean if a field has been set.

### GetVersion

`func (o *EquipmentTpmAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EquipmentTpmAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EquipmentTpmAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EquipmentTpmAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetComputeBoard

`func (o *EquipmentTpmAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *EquipmentTpmAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *EquipmentTpmAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *EquipmentTpmAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentTpmAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentTpmAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentTpmAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentTpmAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentTpmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentTpmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentTpmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentTpmAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


