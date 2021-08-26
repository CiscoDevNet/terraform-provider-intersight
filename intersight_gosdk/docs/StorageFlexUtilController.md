# StorageFlexUtilController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.FlexUtilController"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.FlexUtilController"]
**ControllerName** | Pointer to **string** | Name of the Flex Util Controller. | [optional] 
**ControllerStatus** | Pointer to **string** | The current status of the controller. | [optional] 
**FfControllerId** | Pointer to **string** | Identifier for the Storage Flex Util Controller. | [optional] 
**InternalState** | Pointer to **string** | The internal state of the controller. | [optional] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**FlexUtilPhysicalDrives** | Pointer to [**[]StorageFlexUtilPhysicalDriveRelationship**](StorageFlexUtilPhysicalDriveRelationship.md) | An array of relationships to storageFlexUtilPhysicalDrive resources. | [optional] [readonly] 
**FlexUtilVirtualDrives** | Pointer to [**[]StorageFlexUtilVirtualDriveRelationship**](StorageFlexUtilVirtualDriveRelationship.md) | An array of relationships to storageFlexUtilVirtualDrive resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageFlexUtilController

`func NewStorageFlexUtilController(classId string, objectType string, ) *StorageFlexUtilController`

NewStorageFlexUtilController instantiates a new StorageFlexUtilController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexUtilControllerWithDefaults

`func NewStorageFlexUtilControllerWithDefaults() *StorageFlexUtilController`

NewStorageFlexUtilControllerWithDefaults instantiates a new StorageFlexUtilController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFlexUtilController) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexUtilController) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexUtilController) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFlexUtilController) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexUtilController) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexUtilController) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerName

`func (o *StorageFlexUtilController) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *StorageFlexUtilController) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *StorageFlexUtilController) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *StorageFlexUtilController) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageFlexUtilController) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageFlexUtilController) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageFlexUtilController) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageFlexUtilController) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetFfControllerId

`func (o *StorageFlexUtilController) GetFfControllerId() string`

GetFfControllerId returns the FfControllerId field if non-nil, zero value otherwise.

### GetFfControllerIdOk

`func (o *StorageFlexUtilController) GetFfControllerIdOk() (*string, bool)`

GetFfControllerIdOk returns a tuple with the FfControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfControllerId

`func (o *StorageFlexUtilController) SetFfControllerId(v string)`

SetFfControllerId sets FfControllerId field to given value.

### HasFfControllerId

`func (o *StorageFlexUtilController) HasFfControllerId() bool`

HasFfControllerId returns a boolean if a field has been set.

### GetInternalState

`func (o *StorageFlexUtilController) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *StorageFlexUtilController) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *StorageFlexUtilController) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *StorageFlexUtilController) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetComputeBoard

`func (o *StorageFlexUtilController) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *StorageFlexUtilController) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *StorageFlexUtilController) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *StorageFlexUtilController) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetFlexUtilPhysicalDrives

`func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrives() []StorageFlexUtilPhysicalDriveRelationship`

GetFlexUtilPhysicalDrives returns the FlexUtilPhysicalDrives field if non-nil, zero value otherwise.

### GetFlexUtilPhysicalDrivesOk

`func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrivesOk() (*[]StorageFlexUtilPhysicalDriveRelationship, bool)`

GetFlexUtilPhysicalDrivesOk returns a tuple with the FlexUtilPhysicalDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexUtilPhysicalDrives

`func (o *StorageFlexUtilController) SetFlexUtilPhysicalDrives(v []StorageFlexUtilPhysicalDriveRelationship)`

SetFlexUtilPhysicalDrives sets FlexUtilPhysicalDrives field to given value.

### HasFlexUtilPhysicalDrives

`func (o *StorageFlexUtilController) HasFlexUtilPhysicalDrives() bool`

HasFlexUtilPhysicalDrives returns a boolean if a field has been set.

### SetFlexUtilPhysicalDrivesNil

`func (o *StorageFlexUtilController) SetFlexUtilPhysicalDrivesNil(b bool)`

 SetFlexUtilPhysicalDrivesNil sets the value for FlexUtilPhysicalDrives to be an explicit nil

### UnsetFlexUtilPhysicalDrives
`func (o *StorageFlexUtilController) UnsetFlexUtilPhysicalDrives()`

UnsetFlexUtilPhysicalDrives ensures that no value is present for FlexUtilPhysicalDrives, not even an explicit nil
### GetFlexUtilVirtualDrives

`func (o *StorageFlexUtilController) GetFlexUtilVirtualDrives() []StorageFlexUtilVirtualDriveRelationship`

GetFlexUtilVirtualDrives returns the FlexUtilVirtualDrives field if non-nil, zero value otherwise.

### GetFlexUtilVirtualDrivesOk

`func (o *StorageFlexUtilController) GetFlexUtilVirtualDrivesOk() (*[]StorageFlexUtilVirtualDriveRelationship, bool)`

GetFlexUtilVirtualDrivesOk returns a tuple with the FlexUtilVirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexUtilVirtualDrives

`func (o *StorageFlexUtilController) SetFlexUtilVirtualDrives(v []StorageFlexUtilVirtualDriveRelationship)`

SetFlexUtilVirtualDrives sets FlexUtilVirtualDrives field to given value.

### HasFlexUtilVirtualDrives

`func (o *StorageFlexUtilController) HasFlexUtilVirtualDrives() bool`

HasFlexUtilVirtualDrives returns a boolean if a field has been set.

### SetFlexUtilVirtualDrivesNil

`func (o *StorageFlexUtilController) SetFlexUtilVirtualDrivesNil(b bool)`

 SetFlexUtilVirtualDrivesNil sets the value for FlexUtilVirtualDrives to be an explicit nil

### UnsetFlexUtilVirtualDrives
`func (o *StorageFlexUtilController) UnsetFlexUtilVirtualDrives()`

UnsetFlexUtilVirtualDrives ensures that no value is present for FlexUtilVirtualDrives, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *StorageFlexUtilController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexUtilController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexUtilController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexUtilController) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexUtilController) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexUtilController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexUtilController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexUtilController) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


