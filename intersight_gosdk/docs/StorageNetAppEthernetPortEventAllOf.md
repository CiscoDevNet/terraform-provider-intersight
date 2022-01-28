# StorageNetAppEthernetPortEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppEthernetPortEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppEthernetPortEvent"]
**EthernetPort** | Pointer to [**StorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppEthernetPortEventAllOf

`func NewStorageNetAppEthernetPortEventAllOf(classId string, objectType string, ) *StorageNetAppEthernetPortEventAllOf`

NewStorageNetAppEthernetPortEventAllOf instantiates a new StorageNetAppEthernetPortEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortEventAllOfWithDefaults

`func NewStorageNetAppEthernetPortEventAllOfWithDefaults() *StorageNetAppEthernetPortEventAllOf`

NewStorageNetAppEthernetPortEventAllOfWithDefaults instantiates a new StorageNetAppEthernetPortEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEthernetPort

`func (o *StorageNetAppEthernetPortEventAllOf) GetEthernetPort() StorageNetAppEthernetPortRelationship`

GetEthernetPort returns the EthernetPort field if non-nil, zero value otherwise.

### GetEthernetPortOk

`func (o *StorageNetAppEthernetPortEventAllOf) GetEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetEthernetPortOk returns a tuple with the EthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetPort

`func (o *StorageNetAppEthernetPortEventAllOf) SetEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetEthernetPort sets EthernetPort field to given value.

### HasEthernetPort

`func (o *StorageNetAppEthernetPortEventAllOf) HasEthernetPort() bool`

HasEthernetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


