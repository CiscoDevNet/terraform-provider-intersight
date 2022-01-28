# StorageNetAppIpInterfaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppIpInterfaceEvent"]
**IpInterface** | Pointer to [**StorageNetAppIpInterfaceRelationship**](StorageNetAppIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppIpInterfaceEvent

`func NewStorageNetAppIpInterfaceEvent(classId string, objectType string, ) *StorageNetAppIpInterfaceEvent`

NewStorageNetAppIpInterfaceEvent instantiates a new StorageNetAppIpInterfaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppIpInterfaceEventWithDefaults

`func NewStorageNetAppIpInterfaceEventWithDefaults() *StorageNetAppIpInterfaceEvent`

NewStorageNetAppIpInterfaceEventWithDefaults instantiates a new StorageNetAppIpInterfaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppIpInterfaceEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppIpInterfaceEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppIpInterfaceEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppIpInterfaceEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppIpInterfaceEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppIpInterfaceEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppIpInterfaceEvent) GetIpInterface() StorageNetAppIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppIpInterfaceEvent) GetIpInterfaceOk() (*StorageNetAppIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppIpInterfaceEvent) SetIpInterface(v StorageNetAppIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppIpInterfaceEvent) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


