# StoragePurePortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PurePort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PurePort"]
**Failover** | Pointer to **string** | Name of the port to which this port has failed over. | [optional] [readonly] 
**Portal** | Pointer to **string** | Ip address of iSCSI portal configured on the port. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Controller** | Pointer to [**StoragePureControllerRelationship**](StoragePureControllerRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePurePortAllOf

`func NewStoragePurePortAllOf(classId string, objectType string, ) *StoragePurePortAllOf`

NewStoragePurePortAllOf instantiates a new StoragePurePortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePurePortAllOfWithDefaults

`func NewStoragePurePortAllOfWithDefaults() *StoragePurePortAllOf`

NewStoragePurePortAllOfWithDefaults instantiates a new StoragePurePortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePurePortAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePurePortAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePurePortAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePurePortAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePurePortAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePurePortAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


