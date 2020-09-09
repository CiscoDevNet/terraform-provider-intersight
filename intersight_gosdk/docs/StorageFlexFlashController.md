# StorageFlexFlashController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerState** | Pointer to **string** | State of the Flex Flash Storage Controller. | [optional] [readonly] 
**FfControllerId** | Pointer to **string** | Identifier for the Flex Flash Storage Controller. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**FlexFlashControllerProps** | Pointer to [**[]StorageFlexFlashControllerPropsRelationship**](storage.FlexFlashControllerProps.Relationship.md) | An array of relationships to storageFlexFlashControllerProps resources. | [optional] [readonly] 
**FlexFlashPhysicalDrives** | Pointer to [**[]StorageFlexFlashPhysicalDriveRelationship**](storage.FlexFlashPhysicalDrive.Relationship.md) | An array of relationships to storageFlexFlashPhysicalDrive resources. | [optional] [readonly] 
**FlexFlashVirtualDrives** | Pointer to [**[]StorageFlexFlashVirtualDriveRelationship**](storage.FlexFlashVirtualDrive.Relationship.md) | An array of relationships to storageFlexFlashVirtualDrive resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](firmware.RunningFirmware.Relationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewStorageFlexFlashController

`func NewStorageFlexFlashController() *StorageFlexFlashController`

NewStorageFlexFlashController instantiates a new StorageFlexFlashController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerWithDefaults

`func NewStorageFlexFlashControllerWithDefaults() *StorageFlexFlashController`

NewStorageFlexFlashControllerWithDefaults instantiates a new StorageFlexFlashController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerState

`func (o *StorageFlexFlashController) GetControllerState() string`

GetControllerState returns the ControllerState field if non-nil, zero value otherwise.

### GetControllerStateOk

`func (o *StorageFlexFlashController) GetControllerStateOk() (*string, bool)`

GetControllerStateOk returns a tuple with the ControllerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerState

`func (o *StorageFlexFlashController) SetControllerState(v string)`

SetControllerState sets ControllerState field to given value.

### HasControllerState

`func (o *StorageFlexFlashController) HasControllerState() bool`

HasControllerState returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexFlashController) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexFlashController) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexFlashController) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexFlashController) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexFlashController) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexFlashController) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexFlashController) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexFlashController) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexFlashControllerProps

`func (o *StorageFlexFlashController) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship`

GetFlexFlashControllerProps returns the FlexFlashControllerProps field if non-nil, zero value otherwise.

### GetFlexFlashControllerPropsOk

`func (o *StorageFlexFlashController) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool)`

GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashControllerProps

`func (o *StorageFlexFlashController) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship)`

SetFlexFlashControllerProps sets FlexFlashControllerProps field to given value.

### HasFlexFlashControllerProps

`func (o *StorageFlexFlashController) HasFlexFlashControllerProps() bool`

HasFlexFlashControllerProps returns a boolean if a field has been set.

### SetFlexFlashControllerPropsNil

`func (o *StorageFlexFlashController) SetFlexFlashControllerPropsNil(b bool)`

 SetFlexFlashControllerPropsNil sets the value for FlexFlashControllerProps to be an explicit nil

### UnsetFlexFlashControllerProps
`func (o *StorageFlexFlashController) UnsetFlexFlashControllerProps()`

UnsetFlexFlashControllerProps ensures that no value is present for FlexFlashControllerProps, not even an explicit nil
### GetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashController) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship`

GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexFlashPhysicalDrivesOk

`func (o *StorageFlexFlashController) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool)`

GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashController) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship)`

SetFlexFlashPhysicalDrives sets FlexFlashPhysicalDrives field to given value.

### HasFlexFlashPhysicalDrives

`func (o *StorageFlexFlashController) HasFlexFlashPhysicalDrives() bool`

HasFlexFlashPhysicalDrives returns a boolean if a field has been set.

### SetFlexFlashPhysicalDrivesNil

`func (o *StorageFlexFlashController) SetFlexFlashPhysicalDrivesNil(b bool)`

 SetFlexFlashPhysicalDrivesNil sets the value for FlexFlashPhysicalDrives to be an explicit nil

### UnsetFlexFlashPhysicalDrives
`func (o *StorageFlexFlashController) UnsetFlexFlashPhysicalDrives()`

UnsetFlexFlashPhysicalDrives ensures that no value is present for FlexFlashPhysicalDrives, not even an explicit nil
### GetFlexFlashVirtualDrives

`func (o *StorageFlexFlashController) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship`

GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field if non-nil, zero value otherwise.

### GetFlexFlashVirtualDrivesOk

`func (o *StorageFlexFlashController) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool)`

GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashVirtualDrives

`func (o *StorageFlexFlashController) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship)`

SetFlexFlashVirtualDrives sets FlexFlashVirtualDrives field to given value.

### HasFlexFlashVirtualDrives

`func (o *StorageFlexFlashController) HasFlexFlashVirtualDrives() bool`

HasFlexFlashVirtualDrives returns a boolean if a field has been set.

### SetFlexFlashVirtualDrivesNil

`func (o *StorageFlexFlashController) SetFlexFlashVirtualDrivesNil(b bool)`

 SetFlexFlashVirtualDrivesNil sets the value for FlexFlashVirtualDrives to be an explicit nil

### UnsetFlexFlashVirtualDrives
`func (o *StorageFlexFlashController) UnsetFlexFlashVirtualDrives()`

UnsetFlexFlashVirtualDrives ensures that no value is present for FlexFlashVirtualDrives, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageFlexFlashController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashController) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashController) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashController) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageFlexFlashController) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageFlexFlashController) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageFlexFlashController) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageFlexFlashController) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageFlexFlashController) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageFlexFlashController) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


