# StorageNetAppFcPortEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcPortEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcPortEvent"]
**FcPort** | Pointer to [**StorageNetAppFcPortRelationship**](StorageNetAppFcPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppFcPortEvent

`func NewStorageNetAppFcPortEvent(classId string, objectType string, ) *StorageNetAppFcPortEvent`

NewStorageNetAppFcPortEvent instantiates a new StorageNetAppFcPortEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcPortEventWithDefaults

`func NewStorageNetAppFcPortEventWithDefaults() *StorageNetAppFcPortEvent`

NewStorageNetAppFcPortEventWithDefaults instantiates a new StorageNetAppFcPortEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcPortEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcPortEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcPortEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcPortEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcPortEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcPortEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFcPort

`func (o *StorageNetAppFcPortEvent) GetFcPort() StorageNetAppFcPortRelationship`

GetFcPort returns the FcPort field if non-nil, zero value otherwise.

### GetFcPortOk

`func (o *StorageNetAppFcPortEvent) GetFcPortOk() (*StorageNetAppFcPortRelationship, bool)`

GetFcPortOk returns a tuple with the FcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPort

`func (o *StorageNetAppFcPortEvent) SetFcPort(v StorageNetAppFcPortRelationship)`

SetFcPort sets FcPort field to given value.

### HasFcPort

`func (o *StorageNetAppFcPortEvent) HasFcPort() bool`

HasFcPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


