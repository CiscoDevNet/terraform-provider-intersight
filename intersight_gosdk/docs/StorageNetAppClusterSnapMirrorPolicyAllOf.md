# StorageNetAppClusterSnapMirrorPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppClusterSnapMirrorPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppClusterSnapMirrorPolicy"]
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppClusterSnapMirrorPolicyAllOf

`func NewStorageNetAppClusterSnapMirrorPolicyAllOf(classId string, objectType string, ) *StorageNetAppClusterSnapMirrorPolicyAllOf`

NewStorageNetAppClusterSnapMirrorPolicyAllOf instantiates a new StorageNetAppClusterSnapMirrorPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppClusterSnapMirrorPolicyAllOfWithDefaults

`func NewStorageNetAppClusterSnapMirrorPolicyAllOfWithDefaults() *StorageNetAppClusterSnapMirrorPolicyAllOf`

NewStorageNetAppClusterSnapMirrorPolicyAllOfWithDefaults instantiates a new StorageNetAppClusterSnapMirrorPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArray

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppClusterSnapMirrorPolicyAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


