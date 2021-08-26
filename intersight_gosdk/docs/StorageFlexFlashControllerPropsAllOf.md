# StorageFlexFlashControllerPropsAllOf

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

### NewStorageFlexFlashControllerPropsAllOf

`func NewStorageFlexFlashControllerPropsAllOf(classId string, objectType string, ) *StorageFlexFlashControllerPropsAllOf`

NewStorageFlexFlashControllerPropsAllOf instantiates a new StorageFlexFlashControllerPropsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashControllerPropsAllOfWithDefaults

`func NewStorageFlexFlashControllerPropsAllOfWithDefaults() *StorageFlexFlashControllerPropsAllOf`

NewStorageFlexFlashControllerPropsAllOfWithDefaults instantiates a new StorageFlexFlashControllerPropsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFlexFlashControllerPropsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashControllerPropsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFlexFlashControllerPropsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashControllerPropsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardsManageable

`func (o *StorageFlexFlashControllerPropsAllOf) GetCardsManageable() string`

GetCardsManageable returns the CardsManageable field if non-nil, zero value otherwise.

### GetCardsManageableOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetCardsManageableOk() (*string, bool)`

GetCardsManageableOk returns a tuple with the CardsManageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsManageable

`func (o *StorageFlexFlashControllerPropsAllOf) SetCardsManageable(v string)`

SetCardsManageable sets CardsManageable field to given value.

### HasCardsManageable

`func (o *StorageFlexFlashControllerPropsAllOf) HasCardsManageable() bool`

HasCardsManageable returns a boolean if a field has been set.

### GetConfiguredMode

`func (o *StorageFlexFlashControllerPropsAllOf) GetConfiguredMode() string`

GetConfiguredMode returns the ConfiguredMode field if non-nil, zero value otherwise.

### GetConfiguredModeOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetConfiguredModeOk() (*string, bool)`

GetConfiguredModeOk returns a tuple with the ConfiguredMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMode

`func (o *StorageFlexFlashControllerPropsAllOf) SetConfiguredMode(v string)`

SetConfiguredMode sets ConfiguredMode field to given value.

### HasConfiguredMode

`func (o *StorageFlexFlashControllerPropsAllOf) HasConfiguredMode() bool`

HasConfiguredMode returns a boolean if a field has been set.

### GetControllerName

`func (o *StorageFlexFlashControllerPropsAllOf) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *StorageFlexFlashControllerPropsAllOf) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *StorageFlexFlashControllerPropsAllOf) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetControllerStatus

`func (o *StorageFlexFlashControllerPropsAllOf) GetControllerStatus() string`

GetControllerStatus returns the ControllerStatus field if non-nil, zero value otherwise.

### GetControllerStatusOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetControllerStatusOk() (*string, bool)`

GetControllerStatusOk returns a tuple with the ControllerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerStatus

`func (o *StorageFlexFlashControllerPropsAllOf) SetControllerStatus(v string)`

SetControllerStatus sets ControllerStatus field to given value.

### HasControllerStatus

`func (o *StorageFlexFlashControllerPropsAllOf) HasControllerStatus() bool`

HasControllerStatus returns a boolean if a field has been set.

### GetFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) GetFwVersion() string`

GetFwVersion returns the FwVersion field if non-nil, zero value otherwise.

### GetFwVersionOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetFwVersionOk() (*string, bool)`

GetFwVersionOk returns a tuple with the FwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) SetFwVersion(v string)`

SetFwVersion sets FwVersion field to given value.

### HasFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) HasFwVersion() bool`

HasFwVersion returns a boolean if a field has been set.

### GetInternalState

`func (o *StorageFlexFlashControllerPropsAllOf) GetInternalState() string`

GetInternalState returns the InternalState field if non-nil, zero value otherwise.

### GetInternalStateOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetInternalStateOk() (*string, bool)`

GetInternalStateOk returns a tuple with the InternalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalState

`func (o *StorageFlexFlashControllerPropsAllOf) SetInternalState(v string)`

SetInternalState sets InternalState field to given value.

### HasInternalState

`func (o *StorageFlexFlashControllerPropsAllOf) HasInternalState() bool`

HasInternalState returns a boolean if a field has been set.

### GetOperatingMode

`func (o *StorageFlexFlashControllerPropsAllOf) GetOperatingMode() string`

GetOperatingMode returns the OperatingMode field if non-nil, zero value otherwise.

### GetOperatingModeOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetOperatingModeOk() (*string, bool)`

GetOperatingModeOk returns a tuple with the OperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMode

`func (o *StorageFlexFlashControllerPropsAllOf) SetOperatingMode(v string)`

SetOperatingMode sets OperatingMode field to given value.

### HasOperatingMode

`func (o *StorageFlexFlashControllerPropsAllOf) HasOperatingMode() bool`

HasOperatingMode returns a boolean if a field has been set.

### GetPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) GetPhysicalDriveCount() string`

GetPhysicalDriveCount returns the PhysicalDriveCount field if non-nil, zero value otherwise.

### GetPhysicalDriveCountOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetPhysicalDriveCountOk() (*string, bool)`

GetPhysicalDriveCountOk returns a tuple with the PhysicalDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) SetPhysicalDriveCount(v string)`

SetPhysicalDriveCount sets PhysicalDriveCount field to given value.

### HasPhysicalDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) HasPhysicalDriveCount() bool`

HasPhysicalDriveCount returns a boolean if a field has been set.

### GetProductName

`func (o *StorageFlexFlashControllerPropsAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *StorageFlexFlashControllerPropsAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *StorageFlexFlashControllerPropsAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStartupFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) GetStartupFwVersion() string`

GetStartupFwVersion returns the StartupFwVersion field if non-nil, zero value otherwise.

### GetStartupFwVersionOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetStartupFwVersionOk() (*string, bool)`

GetStartupFwVersionOk returns a tuple with the StartupFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) SetStartupFwVersion(v string)`

SetStartupFwVersion sets StartupFwVersion field to given value.

### HasStartupFwVersion

`func (o *StorageFlexFlashControllerPropsAllOf) HasStartupFwVersion() bool`

HasStartupFwVersion returns a boolean if a field has been set.

### GetVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) GetVirtualDriveCount() string`

GetVirtualDriveCount returns the VirtualDriveCount field if non-nil, zero value otherwise.

### GetVirtualDriveCountOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetVirtualDriveCountOk() (*string, bool)`

GetVirtualDriveCountOk returns a tuple with the VirtualDriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) SetVirtualDriveCount(v string)`

SetVirtualDriveCount sets VirtualDriveCount field to given value.

### HasVirtualDriveCount

`func (o *StorageFlexFlashControllerPropsAllOf) HasVirtualDriveCount() bool`

HasVirtualDriveCount returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashControllerPropsAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashControllerPropsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashControllerPropsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashControllerPropsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashControllerPropsAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashControllerPropsAllOf) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


