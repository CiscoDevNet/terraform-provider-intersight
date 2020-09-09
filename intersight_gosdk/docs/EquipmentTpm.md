# EquipmentTpm

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

### NewEquipmentTpm

`func NewEquipmentTpm() *EquipmentTpm`

NewEquipmentTpm instantiates a new EquipmentTpm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentTpmWithDefaults

`func NewEquipmentTpmWithDefaults() *EquipmentTpm`

NewEquipmentTpmWithDefaults instantiates a new EquipmentTpm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationStatus

`func (o *EquipmentTpm) GetActivationStatus() string`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *EquipmentTpm) GetActivationStatusOk() (*string, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *EquipmentTpm) SetActivationStatus(v string)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *EquipmentTpm) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetAdminState

`func (o *EquipmentTpm) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EquipmentTpm) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EquipmentTpm) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EquipmentTpm) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetOwnership

`func (o *EquipmentTpm) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *EquipmentTpm) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *EquipmentTpm) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *EquipmentTpm) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentTpm) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentTpm) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentTpm) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentTpm) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetTpmId

`func (o *EquipmentTpm) GetTpmId() int64`

GetTpmId returns the TpmId field if non-nil, zero value otherwise.

### GetTpmIdOk

`func (o *EquipmentTpm) GetTpmIdOk() (*int64, bool)`

GetTpmIdOk returns a tuple with the TpmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmId

`func (o *EquipmentTpm) SetTpmId(v int64)`

SetTpmId sets TpmId field to given value.

### HasTpmId

`func (o *EquipmentTpm) HasTpmId() bool`

HasTpmId returns a boolean if a field has been set.

### GetVersion

`func (o *EquipmentTpm) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EquipmentTpm) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EquipmentTpm) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EquipmentTpm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetComputeBoard

`func (o *EquipmentTpm) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *EquipmentTpm) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *EquipmentTpm) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *EquipmentTpm) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentTpm) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentTpm) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentTpm) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentTpm) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentTpm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentTpm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentTpm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentTpm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


