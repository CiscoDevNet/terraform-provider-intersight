# StorageNetAppNonDataIpInterfaceEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNonDataIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNonDataIpInterfaceEvent"]
**IpInterface** | Pointer to [**StorageNetAppNonDataIpInterfaceRelationship**](StorageNetAppNonDataIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNonDataIpInterfaceEventAllOf

`func NewStorageNetAppNonDataIpInterfaceEventAllOf(classId string, objectType string, ) *StorageNetAppNonDataIpInterfaceEventAllOf`

NewStorageNetAppNonDataIpInterfaceEventAllOf instantiates a new StorageNetAppNonDataIpInterfaceEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNonDataIpInterfaceEventAllOfWithDefaults

`func NewStorageNetAppNonDataIpInterfaceEventAllOfWithDefaults() *StorageNetAppNonDataIpInterfaceEventAllOf`

NewStorageNetAppNonDataIpInterfaceEventAllOfWithDefaults instantiates a new StorageNetAppNonDataIpInterfaceEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetIpInterface() StorageNetAppNonDataIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetIpInterfaceOk() (*StorageNetAppNonDataIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetIpInterface(v StorageNetAppNonDataIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppNonDataIpInterfaceEventAllOf) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


