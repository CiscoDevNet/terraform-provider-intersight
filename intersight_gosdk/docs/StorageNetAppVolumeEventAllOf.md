# StorageNetAppVolumeEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppVolumeEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppVolumeEvent"]
**Volume** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppVolumeEventAllOf

`func NewStorageNetAppVolumeEventAllOf(classId string, objectType string, ) *StorageNetAppVolumeEventAllOf`

NewStorageNetAppVolumeEventAllOf instantiates a new StorageNetAppVolumeEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppVolumeEventAllOfWithDefaults

`func NewStorageNetAppVolumeEventAllOfWithDefaults() *StorageNetAppVolumeEventAllOf`

NewStorageNetAppVolumeEventAllOfWithDefaults instantiates a new StorageNetAppVolumeEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppVolumeEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppVolumeEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppVolumeEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppVolumeEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppVolumeEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppVolumeEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVolume

`func (o *StorageNetAppVolumeEventAllOf) GetVolume() StorageNetAppVolumeRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageNetAppVolumeEventAllOf) GetVolumeOk() (*StorageNetAppVolumeRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageNetAppVolumeEventAllOf) SetVolume(v StorageNetAppVolumeRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageNetAppVolumeEventAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


