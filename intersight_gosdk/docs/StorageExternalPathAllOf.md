# StorageExternalPathAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ExternalPath"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ExternalPath"]
**BlockedPathMonitoring** | Pointer to **int64** | The time (in seconds) until the external parity group is blocked after all paths to the external parity group are disconnected. | [optional] [readonly] 
**ExternalWwn** | Pointer to **string** | WWN of the external storage system. | [optional] [readonly] 
**IoTimeOut** | Pointer to **int64** | The value (in seconds) set for the I/O time over for the external parity group. | [optional] [readonly] 
**PortId** | Pointer to **string** | Port number of external path on the local storage system. | [optional] [readonly] 
**QueueDepth** | Pointer to **int64** | Number of Read-Write commands that can be queued to the external parity group. | [optional] [readonly] 

## Methods

### NewStorageExternalPathAllOf

`func NewStorageExternalPathAllOf(classId string, objectType string, ) *StorageExternalPathAllOf`

NewStorageExternalPathAllOf instantiates a new StorageExternalPathAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageExternalPathAllOfWithDefaults

`func NewStorageExternalPathAllOfWithDefaults() *StorageExternalPathAllOf`

NewStorageExternalPathAllOfWithDefaults instantiates a new StorageExternalPathAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageExternalPathAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageExternalPathAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageExternalPathAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageExternalPathAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageExternalPathAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageExternalPathAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockedPathMonitoring

`func (o *StorageExternalPathAllOf) GetBlockedPathMonitoring() int64`

GetBlockedPathMonitoring returns the BlockedPathMonitoring field if non-nil, zero value otherwise.

### GetBlockedPathMonitoringOk

`func (o *StorageExternalPathAllOf) GetBlockedPathMonitoringOk() (*int64, bool)`

GetBlockedPathMonitoringOk returns a tuple with the BlockedPathMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPathMonitoring

`func (o *StorageExternalPathAllOf) SetBlockedPathMonitoring(v int64)`

SetBlockedPathMonitoring sets BlockedPathMonitoring field to given value.

### HasBlockedPathMonitoring

`func (o *StorageExternalPathAllOf) HasBlockedPathMonitoring() bool`

HasBlockedPathMonitoring returns a boolean if a field has been set.

### GetExternalWwn

`func (o *StorageExternalPathAllOf) GetExternalWwn() string`

GetExternalWwn returns the ExternalWwn field if non-nil, zero value otherwise.

### GetExternalWwnOk

`func (o *StorageExternalPathAllOf) GetExternalWwnOk() (*string, bool)`

GetExternalWwnOk returns a tuple with the ExternalWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWwn

`func (o *StorageExternalPathAllOf) SetExternalWwn(v string)`

SetExternalWwn sets ExternalWwn field to given value.

### HasExternalWwn

`func (o *StorageExternalPathAllOf) HasExternalWwn() bool`

HasExternalWwn returns a boolean if a field has been set.

### GetIoTimeOut

`func (o *StorageExternalPathAllOf) GetIoTimeOut() int64`

GetIoTimeOut returns the IoTimeOut field if non-nil, zero value otherwise.

### GetIoTimeOutOk

`func (o *StorageExternalPathAllOf) GetIoTimeOutOk() (*int64, bool)`

GetIoTimeOutOk returns a tuple with the IoTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoTimeOut

`func (o *StorageExternalPathAllOf) SetIoTimeOut(v int64)`

SetIoTimeOut sets IoTimeOut field to given value.

### HasIoTimeOut

`func (o *StorageExternalPathAllOf) HasIoTimeOut() bool`

HasIoTimeOut returns a boolean if a field has been set.

### GetPortId

`func (o *StorageExternalPathAllOf) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageExternalPathAllOf) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageExternalPathAllOf) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageExternalPathAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetQueueDepth

`func (o *StorageExternalPathAllOf) GetQueueDepth() int64`

GetQueueDepth returns the QueueDepth field if non-nil, zero value otherwise.

### GetQueueDepthOk

`func (o *StorageExternalPathAllOf) GetQueueDepthOk() (*int64, bool)`

GetQueueDepthOk returns a tuple with the QueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDepth

`func (o *StorageExternalPathAllOf) SetQueueDepth(v int64)`

SetQueueDepth sets QueueDepth field to given value.

### HasQueueDepth

`func (o *StorageExternalPathAllOf) HasQueueDepth() bool`

HasQueueDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


