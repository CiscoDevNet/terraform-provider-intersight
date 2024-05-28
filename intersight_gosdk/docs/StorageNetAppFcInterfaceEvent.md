# StorageNetAppFcInterfaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcInterfaceEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcInterfaceEvent"]
**FcInterface** | Pointer to [**NullableStorageNetAppFcInterfaceRelationship**](StorageNetAppFcInterfaceRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppFcInterfaceEvent

`func NewStorageNetAppFcInterfaceEvent(classId string, objectType string, ) *StorageNetAppFcInterfaceEvent`

NewStorageNetAppFcInterfaceEvent instantiates a new StorageNetAppFcInterfaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcInterfaceEventWithDefaults

`func NewStorageNetAppFcInterfaceEventWithDefaults() *StorageNetAppFcInterfaceEvent`

NewStorageNetAppFcInterfaceEventWithDefaults instantiates a new StorageNetAppFcInterfaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcInterfaceEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcInterfaceEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcInterfaceEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcInterfaceEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcInterfaceEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcInterfaceEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFcInterface

`func (o *StorageNetAppFcInterfaceEvent) GetFcInterface() StorageNetAppFcInterfaceRelationship`

GetFcInterface returns the FcInterface field if non-nil, zero value otherwise.

### GetFcInterfaceOk

`func (o *StorageNetAppFcInterfaceEvent) GetFcInterfaceOk() (*StorageNetAppFcInterfaceRelationship, bool)`

GetFcInterfaceOk returns a tuple with the FcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcInterface

`func (o *StorageNetAppFcInterfaceEvent) SetFcInterface(v StorageNetAppFcInterfaceRelationship)`

SetFcInterface sets FcInterface field to given value.

### HasFcInterface

`func (o *StorageNetAppFcInterfaceEvent) HasFcInterface() bool`

HasFcInterface returns a boolean if a field has been set.

### SetFcInterfaceNil

`func (o *StorageNetAppFcInterfaceEvent) SetFcInterfaceNil(b bool)`

 SetFcInterfaceNil sets the value for FcInterface to be an explicit nil

### UnsetFcInterface
`func (o *StorageNetAppFcInterfaceEvent) UnsetFcInterface()`

UnsetFcInterface ensures that no value is present for FcInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


