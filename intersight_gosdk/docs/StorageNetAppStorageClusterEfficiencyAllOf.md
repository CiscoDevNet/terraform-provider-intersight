# StorageNetAppStorageClusterEfficiencyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppStorageClusterEfficiency"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppStorageClusterEfficiency"]
**LogicalUsed** | Pointer to **int64** | The logical space used for the cluster. | [optional] [readonly] 
**Ratio** | Pointer to **float32** | Data reduction ratio (logical_used / used). | [optional] [readonly] 
**Savings** | Pointer to **int64** | Space saved by storage efficiencies (logical_used - used). | [optional] [readonly] 

## Methods

### NewStorageNetAppStorageClusterEfficiencyAllOf

`func NewStorageNetAppStorageClusterEfficiencyAllOf(classId string, objectType string, ) *StorageNetAppStorageClusterEfficiencyAllOf`

NewStorageNetAppStorageClusterEfficiencyAllOf instantiates a new StorageNetAppStorageClusterEfficiencyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppStorageClusterEfficiencyAllOfWithDefaults

`func NewStorageNetAppStorageClusterEfficiencyAllOfWithDefaults() *StorageNetAppStorageClusterEfficiencyAllOf`

NewStorageNetAppStorageClusterEfficiencyAllOfWithDefaults instantiates a new StorageNetAppStorageClusterEfficiencyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLogicalUsed

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetLogicalUsed() int64`

GetLogicalUsed returns the LogicalUsed field if non-nil, zero value otherwise.

### GetLogicalUsedOk

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetLogicalUsedOk() (*int64, bool)`

GetLogicalUsedOk returns a tuple with the LogicalUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsed

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) SetLogicalUsed(v int64)`

SetLogicalUsed sets LogicalUsed field to given value.

### HasLogicalUsed

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) HasLogicalUsed() bool`

HasLogicalUsed returns a boolean if a field has been set.

### GetRatio

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetSavings

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetSavings() int64`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) GetSavingsOk() (*int64, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) SetSavings(v int64)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *StorageNetAppStorageClusterEfficiencyAllOf) HasSavings() bool`

HasSavings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


