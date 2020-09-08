# StoragePurePortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failover** | Pointer to **string** | Name of the port to which this port has failed over. | [optional] [readonly] 
**Portal** | Pointer to **string** | Ip address of iSCSI portal configured on the port. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**Controller** | Pointer to [**StoragePureControllerRelationship**](storage.PureController.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePurePortAllOf

`func NewStoragePurePortAllOf() *StoragePurePortAllOf`

NewStoragePurePortAllOf instantiates a new StoragePurePortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePurePortAllOfWithDefaults

`func NewStoragePurePortAllOfWithDefaults() *StoragePurePortAllOf`

NewStoragePurePortAllOfWithDefaults instantiates a new StoragePurePortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailover

`func (o *StoragePurePortAllOf) GetFailover() string`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *StoragePurePortAllOf) GetFailoverOk() (*string, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *StoragePurePortAllOf) SetFailover(v string)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *StoragePurePortAllOf) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetPortal

`func (o *StoragePurePortAllOf) GetPortal() string`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *StoragePurePortAllOf) GetPortalOk() (*string, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *StoragePurePortAllOf) SetPortal(v string)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *StoragePurePortAllOf) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetArray

`func (o *StoragePurePortAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePurePortAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePurePortAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePurePortAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetController

`func (o *StoragePurePortAllOf) GetController() StoragePureControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StoragePurePortAllOf) GetControllerOk() (*StoragePureControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StoragePurePortAllOf) SetController(v StoragePureControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *StoragePurePortAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePurePortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePurePortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePurePortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePurePortAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


