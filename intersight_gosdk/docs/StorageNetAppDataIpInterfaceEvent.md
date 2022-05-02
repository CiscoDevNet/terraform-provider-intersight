# StorageNetAppDataIpInterfaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppDataIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppDataIpInterfaceEvent"]
**IpInterface** | Pointer to [**StorageNetAppDataIpInterfaceRelationship**](StorageNetAppDataIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppDataIpInterfaceEvent

`func NewStorageNetAppDataIpInterfaceEvent(classId string, objectType string, ) *StorageNetAppDataIpInterfaceEvent`

NewStorageNetAppDataIpInterfaceEvent instantiates a new StorageNetAppDataIpInterfaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppDataIpInterfaceEventWithDefaults

`func NewStorageNetAppDataIpInterfaceEventWithDefaults() *StorageNetAppDataIpInterfaceEvent`

NewStorageNetAppDataIpInterfaceEventWithDefaults instantiates a new StorageNetAppDataIpInterfaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppDataIpInterfaceEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppDataIpInterfaceEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppDataIpInterfaceEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppDataIpInterfaceEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppDataIpInterfaceEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppDataIpInterfaceEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppDataIpInterfaceEvent) GetIpInterface() StorageNetAppDataIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppDataIpInterfaceEvent) GetIpInterfaceOk() (*StorageNetAppDataIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppDataIpInterfaceEvent) SetIpInterface(v StorageNetAppDataIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppDataIpInterfaceEvent) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


