# StorageFlexFlashControllerProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.FlexFlashControllerProps"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.FlexFlashControllerProps"]
**CardsManageable** | Pointer to **string** | Manageable card on the flex flash controller. | [optional] 
**ConfiguredMode** | Pointer to **string** | Mode configured on the flex flash controller. | [optional] 
**ControllerName** | Pointer to **string** | The current name of the flex flash controller. | [optional] 
**ControllerStatus** | Pointer to **string** | The current status of the flex flash controller. | [optional] 
**FwVersion** | Pointer to **string** | Firmware version of the flex flash controller. | [optional] 
**InternalState** | Pointer to **string** | Internal state of the flex flash controller. | [optional] 
**OperatingMode** | Pointer to **string** | Operating mode of flex flash controller. | [optional] 
**PhysicalDriveCount** | Pointer to **string** | Number of connected physical drives to a specific Flex flash controller. | [optional] 
**ProductName** | Pointer to **string** | Product name of the flex flash controller. | [optional] 
**StartupFwVersion** | Pointer to **string** | Startup firmware version of the Flex flash controller. | [optional] 
**VirtualDriveCount** | Pointer to **string** | Number of virtual drives for a specific Flex flash controller. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](StorageFlexFlashControllerRelationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashControllerProps

`func NewStorageFlexFlashControllerProps(classId string, objectType string, ) *StorageFlexFlashControllerProps`

NewStorageFlexFlashControllerProps instantiates a new StorageFlexFlashControllerProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerPropsWithDefaults

`func NewStorageFlexFlashControllerPropsWithDefaults() *StorageFlexFlashControllerProps`

NewStorageFlexFlashControllerPropsWithDefaults instantiates a new StorageFlexFlashControllerProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFlexFlashControllerProps) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashControllerProps) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashControllerProps) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFlexFlashControllerProps) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashControllerProps) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashControllerProps) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardsManageable

`func (o *StorageFlexFlashControllerProps) GetCardsManageable() string`

GetCardsManageable returns the CardsManageable field if non-nil, zero value otherwise.

### GetCardsManageableOk

`func (o *StorageFlexFlashControllerProps) GetCardsManageableOk() (*string, bool)`

GetCardsManageableOk returns a tuple with the CardsManageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsManageable

`func (o *StorageFlexFlashControllerProps) SetCardsManageable(v string)`

SetCardsManageable sets CardsManageable field to given value.

### HasCardsManageable

`func (o *StorageFlexFlashControllerProps) HasCardsManageable() bool`

HasCardsManageable returns a boolean if a field has been set.

### GetConfiguredMode

`func (o *StorageFlexFlashControllerProps) GetConfiguredMode() string`

GetConfiguredMode returns the ConfiguredMode field if non-nil, zero value otherwise.

### GetConfiguredModeOk

`func (o *StorageFlexFlashControllerProps) GetConfiguredModeOk() (*string, bool)`

GetConfiguredModeOk returns a tuple with the ConfiguredMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMode

`func (o *StorageFlexFlashControllerProps) SetConfiguredMode(v string)`

SetConfiguredMode sets ConfiguredMode field to given value.

### HasConfiguredMode

`func (o *StorageFlexFlashControllerProps) HasConfiguredMode() bool`

HasConfiguredMode returns a boolean if a field has been set.

### GetControllerName

`func (o *StorageFlexFlashControllerProps) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *StorageFlexFlashControllerProps) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *StorageFlexFlashControllerProps) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *StorageFlexFlashControllerProps) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageFlexFlashControllerProps) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageFlexFlashControllerProps) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageFlexFlashControllerProps) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageFlexFlashControllerProps) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetFwVersion

`func (o *StorageFlexFlashControllerProps) GetFwVersion() string`

GetFwVersion returns the FwVersion field if non-nil, zero value otherwise.

### GetFwVersionOk

`func (o *StorageFlexFlashControllerProps) GetFwVersionOk() (*string, bool)`

GetFwVersionOk returns a tuple with the FwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwVersion

`func (o *StorageFlexFlashControllerProps) SetFwVersion(v string)`

SetFwVersion sets FwVersion field to given value.

### HasFwVersion

`func (o *StorageFlexFlashControllerProps) HasFwVersion() bool`

HasFwVersion returns a boolean if a field has been set.

### GetInternalState

`func (o *StorageFlexFlashControllerProps) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *StorageFlexFlashControllerProps) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *StorageFlexFlashControllerProps) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *StorageFlexFlashControllerProps) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetOperatingMode

`func (o *StorageFlexFlashControllerProps) GetOperatingMode() string`

GetOperatingMode returns the OperatingMode field if non-nil, zero value otherwise.

### GetOperatingModeOk

`func (o *StorageFlexFlashControllerProps) GetOperatingModeOk() (*string, bool)`

GetOperatingModeOk returns a tuple with the OperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMode

`func (o *StorageFlexFlashControllerProps) SetOperatingMode(v string)`

SetOperatingMode sets OperatingMode field to given value.

### HasOperatingMode

`func (o *StorageFlexFlashControllerProps) HasOperatingMode() bool`

HasOperatingMode returns a boolean if a field has been set.

### GetPhysicalDriveCount

`func (o *StorageFlexFlashControllerProps) GetPhysicalDriveCount() string`

GetPhysicalDriveCount returns the PhysicalDriveCount field if non-nil, zero value otherwise.

### GetPhysicalDriveCountOk

`func (o *StorageFlexFlashControllerProps) GetPhysicalDriveCountOk() (*string, bool)`

GetPhysicalDriveCountOk returns a tuple with the PhysicalDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDriveCount

`func (o *StorageFlexFlashControllerProps) SetPhysicalDriveCount(v string)`

SetPhysicalDriveCount sets PhysicalDriveCount field to given value.

### HasPhysicalDriveCount

`func (o *StorageFlexFlashControllerProps) HasPhysicalDriveCount() bool`

HasPhysicalDriveCount returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexFlashControllerProps) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexFlashControllerProps) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexFlashControllerProps) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexFlashControllerProps) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStartupFwVersion

`func (o *StorageFlexFlashControllerProps) GetStartupFwVersion() string`

GetStartupFwVersion returns the StartupFwVersion field if non-nil, zero value otherwise.

### GetStartupFwVersionOk

`func (o *StorageFlexFlashControllerProps) GetStartupFwVersionOk() (*string, bool)`

GetStartupFwVersionOk returns a tuple with the StartupFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupFwVersion

`func (o *StorageFlexFlashControllerProps) SetStartupFwVersion(v string)`

SetStartupFwVersion sets StartupFwVersion field to given value.

### HasStartupFwVersion

`func (o *StorageFlexFlashControllerProps) HasStartupFwVersion() bool`

HasStartupFwVersion returns a boolean if a field has been set.

### GetVirtualDriveCount

`func (o *StorageFlexFlashControllerProps) GetVirtualDriveCount() string`

GetVirtualDriveCount returns the VirtualDriveCount field if non-nil, zero value otherwise.

### GetVirtualDriveCountOk

`func (o *StorageFlexFlashControllerProps) GetVirtualDriveCountOk() (*string, bool)`

GetVirtualDriveCountOk returns a tuple with the VirtualDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveCount

`func (o *StorageFlexFlashControllerProps) SetVirtualDriveCount(v string)`

SetVirtualDriveCount sets VirtualDriveCount field to given value.

### HasVirtualDriveCount

`func (o *StorageFlexFlashControllerProps) HasVirtualDriveCount() bool`

HasVirtualDriveCount returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerProps) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashControllerProps) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerProps) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashControllerProps) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerProps) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerProps) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerProps) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerProps) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashControllerProps) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashControllerProps) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashControllerProps) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashControllerProps) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


