# StorageFlexUtilControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerName** | Pointer to **string** | Name of the Flex Util Controller. | [optional] 
**ControllerStatus** | Pointer to **string** | The current status of the controller. | [optional] 
**FfControllerId** | Pointer to **string** | Identifier for the Storage Flex Util Controller. | [optional] 
**InternalState** | Pointer to **string** | The internal state of the controller. | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**FlexUtilPhysicalDrives** | Pointer to [**[]StorageFlexUtilPhysicalDriveRelationship**](storage.FlexUtilPhysicalDrive.Relationship.md) | An array of relationships to storageFlexUtilPhysicalDrive resources. | [optional] [readonly] 
**FlexUtilVirtualDrives** | Pointer to [**[]StorageFlexUtilVirtualDriveRelationship**](storage.FlexUtilVirtualDrive.Relationship.md) | An array of relationships to storageFlexUtilVirtualDrive resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageFlexUtilControllerAllOf

`func NewStorageFlexUtilControllerAllOf() *StorageFlexUtilControllerAllOf`

NewStorageFlexUtilControllerAllOf instantiates a new StorageFlexUtilControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilControllerAllOfWithDefaults

`func NewStorageFlexUtilControllerAllOfWithDefaults() *StorageFlexUtilControllerAllOf`

NewStorageFlexUtilControllerAllOfWithDefaults instantiates a new StorageFlexUtilControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerName

`func (o *StorageFlexUtilControllerAllOf) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *StorageFlexUtilControllerAllOf) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *StorageFlexUtilControllerAllOf) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *StorageFlexUtilControllerAllOf) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageFlexUtilControllerAllOf) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageFlexUtilControllerAllOf) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageFlexUtilControllerAllOf) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageFlexUtilControllerAllOf) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexUtilControllerAllOf) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexUtilControllerAllOf) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexUtilControllerAllOf) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexUtilControllerAllOf) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetInternalState

`func (o *StorageFlexUtilControllerAllOf) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *StorageFlexUtilControllerAllOf) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *StorageFlexUtilControllerAllOf) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *StorageFlexUtilControllerAllOf) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexUtilControllerAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexUtilControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexUtilControllerAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexUtilControllerAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexUtilPhysicalDrives

`func (o *StorageFlexUtilControllerAllOf) GetFlexUtilPhysicalDrives() []StorageFlexUtilPhysicalDriveRelationship`

GetFlexUtilPhysicalDrives returns the FlexUtilPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexUtilPhysicalDrivesOk

`func (o *StorageFlexUtilControllerAllOf) GetFlexUtilPhysicalDrivesOk() (*[]StorageFlexUtilPhysicalDriveRelationship, bool)`

GetFlexUtilPhysicalDrivesOk returns a tuple with the FlexUtilPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexUtilPhysicalDrives

`func (o *StorageFlexUtilControllerAllOf) SetFlexUtilPhysicalDrives(v []StorageFlexUtilPhysicalDriveRelationship)`

SetFlexUtilPhysicalDrives sets FlexUtilPhysicalDrives field to given value.

### HasFlexUtilPhysicalDrives

`func (o *StorageFlexUtilControllerAllOf) HasFlexUtilPhysicalDrives() bool`

HasFlexUtilPhysicalDrives returns a boolean if a field has been set.

### SetFlexUtilPhysicalDrivesNil

`func (o *StorageFlexUtilControllerAllOf) SetFlexUtilPhysicalDrivesNil(b bool)`

 SetFlexUtilPhysicalDrivesNil sets the value for FlexUtilPhysicalDrives to be an explicit nil

### UnsetFlexUtilPhysicalDrives
`func (o *StorageFlexUtilControllerAllOf) UnsetFlexUtilPhysicalDrives()`

UnsetFlexUtilPhysicalDrives ensures that no value is present for FlexUtilPhysicalDrives, not even an explicit nil
### GetFlexUtilVirtualDrives

`func (o *StorageFlexUtilControllerAllOf) GetFlexUtilVirtualDrives() []StorageFlexUtilVirtualDriveRelationship`

GetFlexUtilVirtualDrives returns the FlexUtilVirtualDrives field if non-nil, zero value otherwise.

### GetFlexUtilVirtualDrivesOk

`func (o *StorageFlexUtilControllerAllOf) GetFlexUtilVirtualDrivesOk() (*[]StorageFlexUtilVirtualDriveRelationship, bool)`

GetFlexUtilVirtualDrivesOk returns a tuple with the FlexUtilVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexUtilVirtualDrives

`func (o *StorageFlexUtilControllerAllOf) SetFlexUtilVirtualDrives(v []StorageFlexUtilVirtualDriveRelationship)`

SetFlexUtilVirtualDrives sets FlexUtilVirtualDrives field to given value.

### HasFlexUtilVirtualDrives

`func (o *StorageFlexUtilControllerAllOf) HasFlexUtilVirtualDrives() bool`

HasFlexUtilVirtualDrives returns a boolean if a field has been set.

### SetFlexUtilVirtualDrivesNil

`func (o *StorageFlexUtilControllerAllOf) SetFlexUtilVirtualDrivesNil(b bool)`

 SetFlexUtilVirtualDrivesNil sets the value for FlexUtilVirtualDrives to be an explicit nil

### UnsetFlexUtilVirtualDrives
`func (o *StorageFlexUtilControllerAllOf) UnsetFlexUtilVirtualDrives()`

UnsetFlexUtilVirtualDrives ensures that no value is present for FlexUtilVirtualDrives, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageFlexUtilControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


