# ComputeBaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**StorageClusters** | Pointer to [**[]StorageBaseClusterRelationship**](StorageBaseClusterRelationship.md) | An array of relationships to storageBaseCluster resources. | [optional] 

## Methods

### NewComputeBaseCluster

`func NewComputeBaseCluster(classId string, objectType string, ) *ComputeBaseCluster`

NewComputeBaseCluster instantiates a new ComputeBaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeBaseClusterWithDefaults

`func NewComputeBaseClusterWithDefaults() *ComputeBaseCluster`

NewComputeBaseClusterWithDefaults instantiates a new ComputeBaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeBaseCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeBaseCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeBaseCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeBaseCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeBaseCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeBaseCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStorageClusters

`func (o *ComputeBaseCluster) GetStorageClusters() []StorageBaseClusterRelationship`

GetStorageClusters returns the StorageClusters field if non-nil, zero value otherwise.

### GetStorageClustersOk

`func (o *ComputeBaseCluster) GetStorageClustersOk() (*[]StorageBaseClusterRelationship, bool)`

GetStorageClustersOk returns a tuple with the StorageClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClusters

`func (o *ComputeBaseCluster) SetStorageClusters(v []StorageBaseClusterRelationship)`

SetStorageClusters sets StorageClusters field to given value.

### HasStorageClusters

`func (o *ComputeBaseCluster) HasStorageClusters() bool`

HasStorageClusters returns a boolean if a field has been set.

### SetStorageClustersNil

`func (o *ComputeBaseCluster) SetStorageClustersNil(b bool)`

 SetStorageClustersNil sets the value for StorageClusters to be an explicit nil

### UnsetStorageClusters
`func (o *ComputeBaseCluster) UnsetStorageClusters()`

UnsetStorageClusters ensures that no value is present for StorageClusters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


