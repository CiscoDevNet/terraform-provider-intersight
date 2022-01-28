# StorageNetAppFcInterfaceEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcInterfaceEvent"]
**FcInterface** | Pointer to [**StorageNetAppFcInterfaceRelationship**](StorageNetAppFcInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppFcInterfaceEventAllOf

`func NewStorageNetAppFcInterfaceEventAllOf(classId string, objectType string, ) *StorageNetAppFcInterfaceEventAllOf`

NewStorageNetAppFcInterfaceEventAllOf instantiates a new StorageNetAppFcInterfaceEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcInterfaceEventAllOfWithDefaults

`func NewStorageNetAppFcInterfaceEventAllOfWithDefaults() *StorageNetAppFcInterfaceEventAllOf`

NewStorageNetAppFcInterfaceEventAllOfWithDefaults instantiates a new StorageNetAppFcInterfaceEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcInterfaceEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcInterfaceEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcInterfaceEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcInterfaceEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcInterfaceEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcInterfaceEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFcInterface

`func (o *StorageNetAppFcInterfaceEventAllOf) GetFcInterface() StorageNetAppFcInterfaceRelationship`

GetFcInterface returns the FcInterface field if non-nil, zero value otherwise.

### GetFcInterfaceOk

`func (o *StorageNetAppFcInterfaceEventAllOf) GetFcInterfaceOk() (*StorageNetAppFcInterfaceRelationship, bool)`

GetFcInterfaceOk returns a tuple with the FcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcInterface

`func (o *StorageNetAppFcInterfaceEventAllOf) SetFcInterface(v StorageNetAppFcInterfaceRelationship)`

SetFcInterface sets FcInterface field to given value.

### HasFcInterface

`func (o *StorageNetAppFcInterfaceEventAllOf) HasFcInterface() bool`

HasFcInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


