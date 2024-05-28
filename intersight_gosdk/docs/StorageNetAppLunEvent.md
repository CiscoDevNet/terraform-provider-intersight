# StorageNetAppLunEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLunEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLunEvent"]
**Lun** | Pointer to [**NullableStorageNetAppLunRelationship**](StorageNetAppLunRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLunEvent

`func NewStorageNetAppLunEvent(classId string, objectType string, ) *StorageNetAppLunEvent`

NewStorageNetAppLunEvent instantiates a new StorageNetAppLunEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLunEventWithDefaults

`func NewStorageNetAppLunEventWithDefaults() *StorageNetAppLunEvent`

NewStorageNetAppLunEventWithDefaults instantiates a new StorageNetAppLunEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLunEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLunEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLunEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLunEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLunEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLunEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLun

`func (o *StorageNetAppLunEvent) GetLun() StorageNetAppLunRelationship`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *StorageNetAppLunEvent) GetLunOk() (*StorageNetAppLunRelationship, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *StorageNetAppLunEvent) SetLun(v StorageNetAppLunRelationship)`

SetLun sets Lun field to given value.

### HasLun

`func (o *StorageNetAppLunEvent) HasLun() bool`

HasLun returns a boolean if a field has been set.

### SetLunNil

`func (o *StorageNetAppLunEvent) SetLunNil(b bool)`

 SetLunNil sets the value for Lun to be an explicit nil

### UnsetLun
`func (o *StorageNetAppLunEvent) UnsetLun()`

UnsetLun ensures that no value is present for Lun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


