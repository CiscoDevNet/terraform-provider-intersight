# StorageNetAppDataIpInterfaceEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppDataIpInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppDataIpInterfaceEvent"]
**IpInterface** | Pointer to [**StorageNetAppDataIpInterfaceRelationship**](StorageNetAppDataIpInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppDataIpInterfaceEventAllOf

`func NewStorageNetAppDataIpInterfaceEventAllOf(classId string, objectType string, ) *StorageNetAppDataIpInterfaceEventAllOf`

NewStorageNetAppDataIpInterfaceEventAllOf instantiates a new StorageNetAppDataIpInterfaceEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppDataIpInterfaceEventAllOfWithDefaults

`func NewStorageNetAppDataIpInterfaceEventAllOfWithDefaults() *StorageNetAppDataIpInterfaceEventAllOf`

NewStorageNetAppDataIpInterfaceEventAllOfWithDefaults instantiates a new StorageNetAppDataIpInterfaceEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppDataIpInterfaceEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppDataIpInterfaceEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpInterface

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetIpInterface() StorageNetAppDataIpInterfaceRelationship`

GetIpInterface returns the IpInterface field if non-nil, zero value otherwise.

### GetIpInterfaceOk

`func (o *StorageNetAppDataIpInterfaceEventAllOf) GetIpInterfaceOk() (*StorageNetAppDataIpInterfaceRelationship, bool)`

GetIpInterfaceOk returns a tuple with the IpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInterface

`func (o *StorageNetAppDataIpInterfaceEventAllOf) SetIpInterface(v StorageNetAppDataIpInterfaceRelationship)`

SetIpInterface sets IpInterface field to given value.

### HasIpInterface

`func (o *StorageNetAppDataIpInterfaceEventAllOf) HasIpInterface() bool`

HasIpInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


