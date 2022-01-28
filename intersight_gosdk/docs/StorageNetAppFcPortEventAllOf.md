# StorageNetAppFcPortEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcPortEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcPortEvent"]
**FcPort** | Pointer to [**StorageNetAppFcPortRelationship**](StorageNetAppFcPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppFcPortEventAllOf

`func NewStorageNetAppFcPortEventAllOf(classId string, objectType string, ) *StorageNetAppFcPortEventAllOf`

NewStorageNetAppFcPortEventAllOf instantiates a new StorageNetAppFcPortEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcPortEventAllOfWithDefaults

`func NewStorageNetAppFcPortEventAllOfWithDefaults() *StorageNetAppFcPortEventAllOf`

NewStorageNetAppFcPortEventAllOfWithDefaults instantiates a new StorageNetAppFcPortEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcPortEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcPortEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcPortEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcPortEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcPortEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcPortEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFcPort

`func (o *StorageNetAppFcPortEventAllOf) GetFcPort() StorageNetAppFcPortRelationship`

GetFcPort returns the FcPort field if non-nil, zero value otherwise.

### GetFcPortOk

`func (o *StorageNetAppFcPortEventAllOf) GetFcPortOk() (*StorageNetAppFcPortRelationship, bool)`

GetFcPortOk returns a tuple with the FcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPort

`func (o *StorageNetAppFcPortEventAllOf) SetFcPort(v StorageNetAppFcPortRelationship)`

SetFcPort sets FcPort field to given value.

### HasFcPort

`func (o *StorageNetAppFcPortEventAllOf) HasFcPort() bool`

HasFcPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


