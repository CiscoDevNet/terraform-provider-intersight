# StorageNetAppIpInterfaceEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppIpInterfaceEvent"]
**IpInterface** | Pointer to [**StorageNetAppIpInterfaceRelationship**](StorageNetAppIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppIpInterfaceEventAllOf

`func NewStorageNetAppIpInterfaceEventAllOf(classId string, objectType string, ) *StorageNetAppIpInterfaceEventAllOf`

NewStorageNetAppIpInterfaceEventAllOf instantiates a new StorageNetAppIpInterfaceEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppIpInterfaceEventAllOfWithDefaults

`func NewStorageNetAppIpInterfaceEventAllOfWithDefaults() *StorageNetAppIpInterfaceEventAllOf`

NewStorageNetAppIpInterfaceEventAllOfWithDefaults instantiates a new StorageNetAppIpInterfaceEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppIpInterfaceEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppIpInterfaceEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppIpInterfaceEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppIpInterfaceEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppIpInterfaceEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppIpInterfaceEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppIpInterfaceEventAllOf) GetIpInterface() StorageNetAppIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppIpInterfaceEventAllOf) GetIpInterfaceOk() (*StorageNetAppIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppIpInterfaceEventAllOf) SetIpInterface(v StorageNetAppIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppIpInterfaceEventAllOf) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


