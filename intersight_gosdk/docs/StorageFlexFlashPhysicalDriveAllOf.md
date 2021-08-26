# StorageFlexFlashPhysicalDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.FlexFlashPhysicalDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.FlexFlashPhysicalDrive"]
**CardStatus** | Pointer to **string** | The status of the flex flash physical drive. | [optional] 
**CardType** | Pointer to **string** | The card type of the flex flash physical drive. | [optional] 
**OemId** | Pointer to **string** | The OEM Identifier of the flex flash physical drive. | [optional] 
**PdStatus** | Pointer to **string** | The drive status of the flex flash physical drive. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageFlexFlashController** | Pointer to [**StorageFlexFlashControllerRelationship**](StorageFlexFlashControllerRelationship.md) |  | [optional] 

## Methods

### NewStorageFlexFlashPhysicalDriveAllOf

`func NewStorageFlexFlashPhysicalDriveAllOf(classId string, objectType string, ) *StorageFlexFlashPhysicalDriveAllOf`

NewStorageFlexFlashPhysicalDriveAllOf instantiates a new StorageFlexFlashPhysicalDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFlexFlashPhysicalDriveAllOfWithDefaults

`func NewStorageFlexFlashPhysicalDriveAllOfWithDefaults() *StorageFlexFlashPhysicalDriveAllOf`

NewStorageFlexFlashPhysicalDriveAllOfWithDefaults instantiates a new StorageFlexFlashPhysicalDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardStatus() string`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardStatusOk() (*string, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetCardStatus(v string)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetCardType

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetOemId

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetOemId() string`

GetOemId returns the OemId field if non-nil, zero value otherwise.

### GetOemIdOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetOemIdOk() (*string, bool)`

GetOemIdOk returns a tuple with the OemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOemId

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetOemId(v string)`

SetOemId sets OemId field to given value.

### HasOemId

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasOemId() bool`

HasOemId returns a boolean if a field has been set.

### GetPdStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetPdStatus() string`

GetPdStatus returns the PdStatus field if non-nil, zero value otherwise.

### GetPdStatusOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetPdStatusOk() (*string, bool)`

GetPdStatusOk returns a tuple with the PdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetPdStatus(v string)`

SetPdStatus sets PdStatus field to given value.

### HasPdStatus

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasPdStatus() bool`

HasPdStatus returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageFlexFlashController

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship`

GetStorageFlexFlashController returns the StorageFlexFlashController field if non-nil, zero value otherwise.

### GetStorageFlexFlashControllerOk

`func (o *StorageFlexFlashPhysicalDriveAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool)`

GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFlexFlashController

`func (o *StorageFlexFlashPhysicalDriveAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship)`

SetStorageFlexFlashController sets StorageFlexFlashController field to given value.

### HasStorageFlexFlashController

`func (o *StorageFlexFlashPhysicalDriveAllOf) HasStorageFlexFlashController() bool`

HasStorageFlexFlashController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


