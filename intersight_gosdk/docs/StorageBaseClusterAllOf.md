# StorageBaseClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.Cluster"]
**StorageCapacity** | Pointer to **int64** | The storage capacity in this cluster. | [optional] [readonly] 
**ParentCluster** | Pointer to [**ComputeBaseClusterRelationship**](ComputeBaseClusterRelationship.md) |  | [optional] 

## Methods

### NewStorageBaseClusterAllOf

`func NewStorageBaseClusterAllOf(classId string, objectType string, ) *StorageBaseClusterAllOf`

NewStorageBaseClusterAllOf instantiates a new StorageBaseClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseClusterAllOfWithDefaults

`func NewStorageBaseClusterAllOfWithDefaults() *StorageBaseClusterAllOf`

NewStorageBaseClusterAllOfWithDefaults instantiates a new StorageBaseClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStorageCapacity

`func (o *StorageBaseClusterAllOf) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *StorageBaseClusterAllOf) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *StorageBaseClusterAllOf) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *StorageBaseClusterAllOf) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### GetParentCluster

`func (o *StorageBaseClusterAllOf) GetParentCluster() ComputeBaseClusterRelationship`

GetParentCluster returns the ParentCluster field if non-nil, zero value otherwise.

### GetParentClusterOk

`func (o *StorageBaseClusterAllOf) GetParentClusterOk() (*ComputeBaseClusterRelationship, bool)`

GetParentClusterOk returns a tuple with the ParentCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCluster

`func (o *StorageBaseClusterAllOf) SetParentCluster(v ComputeBaseClusterRelationship)`

SetParentCluster sets ParentCluster field to given value.

### HasParentCluster

`func (o *StorageBaseClusterAllOf) HasParentCluster() bool`

HasParentCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


