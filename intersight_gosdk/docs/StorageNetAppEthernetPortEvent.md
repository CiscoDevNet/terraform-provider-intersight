# StorageNetAppEthernetPortEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppEthernetPortEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppEthernetPortEvent"]
**EthernetPort** | Pointer to [**StorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppEthernetPortEvent

`func NewStorageNetAppEthernetPortEvent(classId string, objectType string, ) *StorageNetAppEthernetPortEvent`

NewStorageNetAppEthernetPortEvent instantiates a new StorageNetAppEthernetPortEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortEventWithDefaults

`func NewStorageNetAppEthernetPortEventWithDefaults() *StorageNetAppEthernetPortEvent`

NewStorageNetAppEthernetPortEventWithDefaults instantiates a new StorageNetAppEthernetPortEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEthernetPort

`func (o *StorageNetAppEthernetPortEvent) GetEthernetPort() StorageNetAppEthernetPortRelationship`

GetEthernetPort returns the EthernetPort field if non-nil, zero value otherwise.

### GetEthernetPortOk

`func (o *StorageNetAppEthernetPortEvent) GetEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetEthernetPortOk returns a tuple with the EthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPort

`func (o *StorageNetAppEthernetPortEvent) SetEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetEthernetPort sets EthernetPort field to given value.

### HasEthernetPort

`func (o *StorageNetAppEthernetPortEvent) HasEthernetPort() bool`

HasEthernetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


