# StorageHitachiInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiInitiator"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiInitiator"]
**Wwpn** | Pointer to **string** | World wide port name, 64 bit address represented in hexa decimal notation. | [optional] [readonly] 

## Methods

### NewStorageHitachiInitiator

`func NewStorageHitachiInitiator(classId string, objectType string, ) *StorageHitachiInitiator`

NewStorageHitachiInitiator instantiates a new StorageHitachiInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiInitiatorWithDefaults

`func NewStorageHitachiInitiatorWithDefaults() *StorageHitachiInitiator`

NewStorageHitachiInitiatorWithDefaults instantiates a new StorageHitachiInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiInitiator) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiInitiator) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiInitiator) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiInitiator) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiInitiator) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiInitiator) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWwpn

`func (o *StorageHitachiInitiator) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *StorageHitachiInitiator) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *StorageHitachiInitiator) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *StorageHitachiInitiator) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


