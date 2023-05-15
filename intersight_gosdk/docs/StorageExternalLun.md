# StorageExternalLun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ExternalLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ExternalLun"]
**ExternalLun** | Pointer to **int64** | LUN within the ports of the external storage system. | [optional] [readonly] 
**ExternalWwn** | Pointer to **string** | WWN of the external storage system. | [optional] [readonly] 
**PathStatus** | Pointer to **string** | Status of the external path. | [optional] [readonly] 
**PortId** | Pointer to **string** | Port number of the local storage system. | [optional] [readonly] 
**Priority** | Pointer to **int64** | Priority within the external path group. | [optional] [readonly] 

## Methods

### NewStorageExternalLun

`func NewStorageExternalLun(classId string, objectType string, ) *StorageExternalLun`

NewStorageExternalLun instantiates a new StorageExternalLun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageExternalLunWithDefaults

`func NewStorageExternalLunWithDefaults() *StorageExternalLun`

NewStorageExternalLunWithDefaults instantiates a new StorageExternalLun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageExternalLun) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageExternalLun) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageExternalLun) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageExternalLun) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageExternalLun) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageExternalLun) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalLun

`func (o *StorageExternalLun) GetExternalLun() int64`

GetExternalLun returns the ExternalLun field if non-nil, zero value otherwise.

### GetExternalLunOk

`func (o *StorageExternalLun) GetExternalLunOk() (*int64, bool)`

GetExternalLunOk returns a tuple with the ExternalLun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLun

`func (o *StorageExternalLun) SetExternalLun(v int64)`

SetExternalLun sets ExternalLun field to given value.

### HasExternalLun

`func (o *StorageExternalLun) HasExternalLun() bool`

HasExternalLun returns a boolean if a field has been set.

### GetExternalWwn

`func (o *StorageExternalLun) GetExternalWwn() string`

GetExternalWwn returns the ExternalWwn field if non-nil, zero value otherwise.

### GetExternalWwnOk

`func (o *StorageExternalLun) GetExternalWwnOk() (*string, bool)`

GetExternalWwnOk returns a tuple with the ExternalWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWwn

`func (o *StorageExternalLun) SetExternalWwn(v string)`

SetExternalWwn sets ExternalWwn field to given value.

### HasExternalWwn

`func (o *StorageExternalLun) HasExternalWwn() bool`

HasExternalWwn returns a boolean if a field has been set.

### GetPathStatus

`func (o *StorageExternalLun) GetPathStatus() string`

GetPathStatus returns the PathStatus field if non-nil, zero value otherwise.

### GetPathStatusOk

`func (o *StorageExternalLun) GetPathStatusOk() (*string, bool)`

GetPathStatusOk returns a tuple with the PathStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathStatus

`func (o *StorageExternalLun) SetPathStatus(v string)`

SetPathStatus sets PathStatus field to given value.

### HasPathStatus

`func (o *StorageExternalLun) HasPathStatus() bool`

HasPathStatus returns a boolean if a field has been set.

### GetPortId

`func (o *StorageExternalLun) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StorageExternalLun) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StorageExternalLun) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *StorageExternalLun) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPriority

`func (o *StorageExternalLun) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StorageExternalLun) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StorageExternalLun) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StorageExternalLun) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


