# StorageSasPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The SAS Address assigned to storage port. | [optional] [readonly] 
**DiskId** | Pointer to **int64** | The unique disk identifier. | [optional] [readonly] 
**EndPointId** | Pointer to **int64** | The end-point Id assigned to storage port. | [optional] [readonly] 
**LinkDescription** | Pointer to **string** | The description for the link. | [optional] [readonly] 
**LinkSpeed** | Pointer to **string** | The link speed negotiated for communication. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StoragePhysicalDisk** | Pointer to [**StoragePhysicalDiskRelationship**](storage.PhysicalDisk.Relationship.md) |  | [optional] 

## Methods

### NewStorageSasPortAllOf

`func NewStorageSasPortAllOf() *StorageSasPortAllOf`

NewStorageSasPortAllOf instantiates a new StorageSasPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSasPortAllOfWithDefaults

`func NewStorageSasPortAllOfWithDefaults() *StorageSasPortAllOf`

NewStorageSasPortAllOfWithDefaults instantiates a new StorageSasPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StorageSasPortAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StorageSasPortAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StorageSasPortAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StorageSasPortAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDiskId

`func (o *StorageSasPortAllOf) GetDiskId() int64`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *StorageSasPortAllOf) GetDiskIdOk() (*int64, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *StorageSasPortAllOf) SetDiskId(v int64)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *StorageSasPortAllOf) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetEndPointId

`func (o *StorageSasPortAllOf) GetEndPointId() int64`

GetEndPointId returns the EndPointId field if non-nil, zero value otherwise.

### GetEndPointIdOk

`func (o *StorageSasPortAllOf) GetEndPointIdOk() (*int64, bool)`

GetEndPointIdOk returns a tuple with the EndPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointId

`func (o *StorageSasPortAllOf) SetEndPointId(v int64)`

SetEndPointId sets EndPointId field to given value.

### HasEndPointId

`func (o *StorageSasPortAllOf) HasEndPointId() bool`

HasEndPointId returns a boolean if a field has been set.

### GetLinkDescription

`func (o *StorageSasPortAllOf) GetLinkDescription() string`

GetLinkDescription returns the LinkDescription field if non-nil, zero value otherwise.

### GetLinkDescriptionOk

`func (o *StorageSasPortAllOf) GetLinkDescriptionOk() (*string, bool)`

GetLinkDescriptionOk returns a tuple with the LinkDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDescription

`func (o *StorageSasPortAllOf) SetLinkDescription(v string)`

SetLinkDescription sets LinkDescription field to given value.

### HasLinkDescription

`func (o *StorageSasPortAllOf) HasLinkDescription() bool`

HasLinkDescription returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *StorageSasPortAllOf) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *StorageSasPortAllOf) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *StorageSasPortAllOf) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *StorageSasPortAllOf) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageSasPortAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageSasPortAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageSasPortAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageSasPortAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageSasPortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageSasPortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageSasPortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageSasPortAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStoragePhysicalDisk

`func (o *StorageSasPortAllOf) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship`

GetStoragePhysicalDisk returns the StoragePhysicalDisk field if non-nil, zero value otherwise.

### GetStoragePhysicalDiskOk

`func (o *StorageSasPortAllOf) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool)`

GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalDisk

`func (o *StorageSasPortAllOf) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship)`

SetStoragePhysicalDisk sets StoragePhysicalDisk field to given value.

### HasStoragePhysicalDisk

`func (o *StorageSasPortAllOf) HasStoragePhysicalDisk() bool`

HasStoragePhysicalDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


