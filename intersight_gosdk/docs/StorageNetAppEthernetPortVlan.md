# StorageNetAppEthernetPortVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppEthernetPortVlan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppEthernetPortVlan"]
**BasePort** | Pointer to [**NullableStorageNetAppPort**](StorageNetAppPort.md) |  | [optional] 
**Tag** | Pointer to **int64** | The ID tag of the VLAN for this port. | [optional] [readonly] 

## Methods

### NewStorageNetAppEthernetPortVlan

`func NewStorageNetAppEthernetPortVlan(classId string, objectType string, ) *StorageNetAppEthernetPortVlan`

NewStorageNetAppEthernetPortVlan instantiates a new StorageNetAppEthernetPortVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortVlanWithDefaults

`func NewStorageNetAppEthernetPortVlanWithDefaults() *StorageNetAppEthernetPortVlan`

NewStorageNetAppEthernetPortVlanWithDefaults instantiates a new StorageNetAppEthernetPortVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortVlan) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortVlan) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortVlan) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortVlan) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortVlan) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortVlan) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBasePort

`func (o *StorageNetAppEthernetPortVlan) GetBasePort() StorageNetAppPort`

GetBasePort returns the BasePort field if non-nil, zero value otherwise.

### GetBasePortOk

`func (o *StorageNetAppEthernetPortVlan) GetBasePortOk() (*StorageNetAppPort, bool)`

GetBasePortOk returns a tuple with the BasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePort

`func (o *StorageNetAppEthernetPortVlan) SetBasePort(v StorageNetAppPort)`

SetBasePort sets BasePort field to given value.

### HasBasePort

`func (o *StorageNetAppEthernetPortVlan) HasBasePort() bool`

HasBasePort returns a boolean if a field has been set.

### SetBasePortNil

`func (o *StorageNetAppEthernetPortVlan) SetBasePortNil(b bool)`

 SetBasePortNil sets the value for BasePort to be an explicit nil

### UnsetBasePort
`func (o *StorageNetAppEthernetPortVlan) UnsetBasePort()`

UnsetBasePort ensures that no value is present for BasePort, not even an explicit nil
### GetTag

`func (o *StorageNetAppEthernetPortVlan) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *StorageNetAppEthernetPortVlan) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *StorageNetAppEthernetPortVlan) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *StorageNetAppEthernetPortVlan) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


