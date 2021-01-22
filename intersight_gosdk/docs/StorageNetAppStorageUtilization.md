# StorageNetAppStorageUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppStorageUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppStorageUtilization"]
**LogicalUsed** | Pointer to **int64** | Total logical used capacity, represented in bytes. | [optional] [readonly] 
**Savings** | Pointer to **int64** | Total savings capacity, represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageNetAppStorageUtilization

`func NewStorageNetAppStorageUtilization(classId string, objectType string, ) *StorageNetAppStorageUtilization`

NewStorageNetAppStorageUtilization instantiates a new StorageNetAppStorageUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppStorageUtilizationWithDefaults

`func NewStorageNetAppStorageUtilizationWithDefaults() *StorageNetAppStorageUtilization`

NewStorageNetAppStorageUtilizationWithDefaults instantiates a new StorageNetAppStorageUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppStorageUtilization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppStorageUtilization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppStorageUtilization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppStorageUtilization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppStorageUtilization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppStorageUtilization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLogicalUsed

`func (o *StorageNetAppStorageUtilization) GetLogicalUsed() int64`

GetLogicalUsed returns the LogicalUsed field if non-nil, zero value otherwise.

### GetLogicalUsedOk

`func (o *StorageNetAppStorageUtilization) GetLogicalUsedOk() (*int64, bool)`

GetLogicalUsedOk returns a tuple with the LogicalUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsed

`func (o *StorageNetAppStorageUtilization) SetLogicalUsed(v int64)`

SetLogicalUsed sets LogicalUsed field to given value.

### HasLogicalUsed

`func (o *StorageNetAppStorageUtilization) HasLogicalUsed() bool`

HasLogicalUsed returns a boolean if a field has been set.

### GetSavings

`func (o *StorageNetAppStorageUtilization) GetSavings() int64`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *StorageNetAppStorageUtilization) GetSavingsOk() (*int64, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *StorageNetAppStorageUtilization) SetSavings(v int64)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *StorageNetAppStorageUtilization) HasSavings() bool`

HasSavings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


