# HyperflexReSyncClusterMoIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ReSyncClusterMoIds"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ReSyncClusterMoIds"]
**AccountMoId** | Pointer to **string** | The customer account MoId. | [optional] [readonly] 
**ClusterMoIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHyperflexReSyncClusterMoIds

`func NewHyperflexReSyncClusterMoIds(classId string, objectType string, ) *HyperflexReSyncClusterMoIds`

NewHyperflexReSyncClusterMoIds instantiates a new HyperflexReSyncClusterMoIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexReSyncClusterMoIdsWithDefaults

`func NewHyperflexReSyncClusterMoIdsWithDefaults() *HyperflexReSyncClusterMoIds`

NewHyperflexReSyncClusterMoIdsWithDefaults instantiates a new HyperflexReSyncClusterMoIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexReSyncClusterMoIds) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexReSyncClusterMoIds) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexReSyncClusterMoIds) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexReSyncClusterMoIds) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexReSyncClusterMoIds) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexReSyncClusterMoIds) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountMoId

`func (o *HyperflexReSyncClusterMoIds) GetAccountMoId() string`

GetAccountMoId returns the AccountMoId field if non-nil, zero value otherwise.

### GetAccountMoIdOk

`func (o *HyperflexReSyncClusterMoIds) GetAccountMoIdOk() (*string, bool)`

GetAccountMoIdOk returns a tuple with the AccountMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoId

`func (o *HyperflexReSyncClusterMoIds) SetAccountMoId(v string)`

SetAccountMoId sets AccountMoId field to given value.

### HasAccountMoId

`func (o *HyperflexReSyncClusterMoIds) HasAccountMoId() bool`

HasAccountMoId returns a boolean if a field has been set.

### GetClusterMoIds

`func (o *HyperflexReSyncClusterMoIds) GetClusterMoIds() []string`

GetClusterMoIds returns the ClusterMoIds field if non-nil, zero value otherwise.

### GetClusterMoIdsOk

`func (o *HyperflexReSyncClusterMoIds) GetClusterMoIdsOk() (*[]string, bool)`

GetClusterMoIdsOk returns a tuple with the ClusterMoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMoIds

`func (o *HyperflexReSyncClusterMoIds) SetClusterMoIds(v []string)`

SetClusterMoIds sets ClusterMoIds field to given value.

### HasClusterMoIds

`func (o *HyperflexReSyncClusterMoIds) HasClusterMoIds() bool`

HasClusterMoIds returns a boolean if a field has been set.

### SetClusterMoIdsNil

`func (o *HyperflexReSyncClusterMoIds) SetClusterMoIdsNil(b bool)`

 SetClusterMoIdsNil sets the value for ClusterMoIds to be an explicit nil

### UnsetClusterMoIds
`func (o *HyperflexReSyncClusterMoIds) UnsetClusterMoIds()`

UnsetClusterMoIds ensures that no value is present for ClusterMoIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


