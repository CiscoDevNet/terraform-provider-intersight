# StorageFlexFlashControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.FlexFlashController"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.FlexFlashController"]
**ControllerState** | Pointer to **string** | State of the Flex Flash Storage Controller. | [optional] [readonly] 
**FfControllerId** | Pointer to **string** | Identifier for the Flex Flash Storage Controller. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**FlexFlashControllerProps** | Pointer to [**[]StorageFlexFlashControllerPropsRelationship**](StorageFlexFlashControllerPropsRelationship.md) | An array of relationships to storageFlexFlashControllerProps resources. | [optional] [readonly] 
**FlexFlashPhysicalDrives** | Pointer to [**[]StorageFlexFlashPhysicalDriveRelationship**](StorageFlexFlashPhysicalDriveRelationship.md) | An array of relationships to storageFlexFlashPhysicalDrive resources. | [optional] [readonly] 
**FlexFlashVirtualDrives** | Pointer to [**[]StorageFlexFlashVirtualDriveRelationship**](StorageFlexFlashVirtualDriveRelationship.md) | An array of relationships to storageFlexFlashVirtualDrive resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningFirmware** | Pointer to [**[]FirmwareRunningFirmwareRelationship**](FirmwareRunningFirmwareRelationship.md) | An array of relationships to firmwareRunningFirmware resources. | [optional] [readonly] 

## Methods

### NewStorageFlexFlashControllerAllOf

`func NewStorageFlexFlashControllerAllOf(classId string, objectType string, ) *StorageFlexFlashControllerAllOf`

NewStorageFlexFlashControllerAllOf instantiates a new StorageFlexFlashControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerAllOfWithDefaults

`func NewStorageFlexFlashControllerAllOfWithDefaults() *StorageFlexFlashControllerAllOf`

NewStorageFlexFlashControllerAllOfWithDefaults instantiates a new StorageFlexFlashControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFlexFlashControllerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashControllerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashControllerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFlexFlashControllerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashControllerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashControllerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerState

`func (o *StorageFlexFlashControllerAllOf) GetControllerState() string`

GetControllerState returns the ControllerState field if non-nil, zero value otherwise.

### GetControllerStateOk

`func (o *StorageFlexFlashControllerAllOf) GetControllerStateOk() (*string, bool)`

GetControllerStateOk returns a tuple with the ControllerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerState

`func (o *StorageFlexFlashControllerAllOf) SetControllerState(v string)`

SetControllerState sets ControllerState field to given value.

### HasControllerState

`func (o *StorageFlexFlashControllerAllOf) HasControllerState() bool`

HasControllerState returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexFlashControllerAllOf) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexFlashControllerAllOf) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexFlashControllerAllOf) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexFlashControllerAllOf) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexFlashControllerAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexFlashControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexFlashControllerAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexFlashControllerAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship`

GetFlexFlashControllerProps returns the FlexFlashControllerProps field if non-nil, zero value otherwise.

### GetFlexFlashControllerPropsOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool)`

GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship)`

SetFlexFlashControllerProps sets FlexFlashControllerProps field to given value.

### HasFlexFlashControllerProps

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashControllerProps() bool`

HasFlexFlashControllerProps returns a boolean if a field has been set.

### SetFlexFlashControllerPropsNil

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashControllerPropsNil(b bool)`

 SetFlexFlashControllerPropsNil sets the value for FlexFlashControllerProps to be an explicit nil

### UnsetFlexFlashControllerProps
`func (o *StorageFlexFlashControllerAllOf) UnsetFlexFlashControllerProps()`

UnsetFlexFlashControllerProps ensures that no value is present for FlexFlashControllerProps, not even an explicit nil
### GetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship`

GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexFlashPhysicalDrivesOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool)`

GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship)`

SetFlexFlashPhysicalDrives sets FlexFlashPhysicalDrives field to given value.

### HasFlexFlashPhysicalDrives

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashPhysicalDrives() bool`

HasFlexFlashPhysicalDrives returns a boolean if a field has been set.

### SetFlexFlashPhysicalDrivesNil

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashPhysicalDrivesNil(b bool)`

 SetFlexFlashPhysicalDrivesNil sets the value for FlexFlashPhysicalDrives to be an explicit nil

### UnsetFlexFlashPhysicalDrives
`func (o *StorageFlexFlashControllerAllOf) UnsetFlexFlashPhysicalDrives()`

UnsetFlexFlashPhysicalDrives ensures that no value is present for FlexFlashPhysicalDrives, not even an explicit nil
### GetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship`

GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field if non-nil, zero value otherwise.

### GetFlexFlashVirtualDrivesOk

`func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool)`

GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship)`

SetFlexFlashVirtualDrives sets FlexFlashVirtualDrives field to given value.

### HasFlexFlashVirtualDrives

`func (o *StorageFlexFlashControllerAllOf) HasFlexFlashVirtualDrives() bool`

HasFlexFlashVirtualDrives returns a boolean if a field has been set.

### SetFlexFlashVirtualDrivesNil

`func (o *StorageFlexFlashControllerAllOf) SetFlexFlashVirtualDrivesNil(b bool)`

 SetFlexFlashVirtualDrivesNil sets the value for FlexFlashVirtualDrives to be an explicit nil

### UnsetFlexFlashVirtualDrives
`func (o *StorageFlexFlashControllerAllOf) UnsetFlexFlashVirtualDrives()`

UnsetFlexFlashVirtualDrives ensures that no value is present for FlexFlashVirtualDrives, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetRunningFirmware

`func (o *StorageFlexFlashControllerAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship`

GetRunningFirmware returns the RunningFirmware field if non-nil, zero value otherwise.

### GetRunningFirmwareOk

`func (o *StorageFlexFlashControllerAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool)`

GetRunningFirmwareOk returns a tuple with the RunningFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningFirmware

`func (o *StorageFlexFlashControllerAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship)`

SetRunningFirmware sets RunningFirmware field to given value.

### HasRunningFirmware

`func (o *StorageFlexFlashControllerAllOf) HasRunningFirmware() bool`

HasRunningFirmware returns a boolean if a field has been set.

### SetRunningFirmwareNil

`func (o *StorageFlexFlashControllerAllOf) SetRunningFirmwareNil(b bool)`

 SetRunningFirmwareNil sets the value for RunningFirmware to be an explicit nil

### UnsetRunningFirmware
`func (o *StorageFlexFlashControllerAllOf) UnsetRunningFirmware()`

UnsetRunningFirmware ensures that no value is present for RunningFirmware, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


