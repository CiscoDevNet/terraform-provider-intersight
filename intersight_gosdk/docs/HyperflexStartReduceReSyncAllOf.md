# HyperflexStartReduceReSyncAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.StartReduceReSync"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.StartReduceReSync"]
**ClusterMoIds** | Pointer to **[]string** |  | [optional] 
**Operation** | Pointer to **string** | The cleanup operation to perform. * &#x60;NoOp&#x60; - Does not perform any operation when the API is called. * &#x60;StartReduceResync&#x60; - Start the execution of reduce re-sync and stale mirror cleanup for the HyperFlex clusters associated with the account. | [optional] [default to "NoOp"]

## Methods

### NewHyperflexStartReduceReSyncAllOf

`func NewHyperflexStartReduceReSyncAllOf(classId string, objectType string, ) *HyperflexStartReduceReSyncAllOf`

NewHyperflexStartReduceReSyncAllOf instantiates a new HyperflexStartReduceReSyncAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStartReduceReSyncAllOfWithDefaults

`func NewHyperflexStartReduceReSyncAllOfWithDefaults() *HyperflexStartReduceReSyncAllOf`

NewHyperflexStartReduceReSyncAllOfWithDefaults instantiates a new HyperflexStartReduceReSyncAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStartReduceReSyncAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStartReduceReSyncAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStartReduceReSyncAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStartReduceReSyncAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStartReduceReSyncAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStartReduceReSyncAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterMoIds

`func (o *HyperflexStartReduceReSyncAllOf) GetClusterMoIds() []string`

GetClusterMoIds returns the ClusterMoIds field if non-nil, zero value otherwise.

### GetClusterMoIdsOk

`func (o *HyperflexStartReduceReSyncAllOf) GetClusterMoIdsOk() (*[]string, bool)`

GetClusterMoIdsOk returns a tuple with the ClusterMoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMoIds

`func (o *HyperflexStartReduceReSyncAllOf) SetClusterMoIds(v []string)`

SetClusterMoIds sets ClusterMoIds field to given value.

### HasClusterMoIds

`func (o *HyperflexStartReduceReSyncAllOf) HasClusterMoIds() bool`

HasClusterMoIds returns a boolean if a field has been set.

### SetClusterMoIdsNil

`func (o *HyperflexStartReduceReSyncAllOf) SetClusterMoIdsNil(b bool)`

 SetClusterMoIdsNil sets the value for ClusterMoIds to be an explicit nil

### UnsetClusterMoIds
`func (o *HyperflexStartReduceReSyncAllOf) UnsetClusterMoIds()`

UnsetClusterMoIds ensures that no value is present for ClusterMoIds, not even an explicit nil
### GetOperation

`func (o *HyperflexStartReduceReSyncAllOf) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HyperflexStartReduceReSyncAllOf) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HyperflexStartReduceReSyncAllOf) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HyperflexStartReduceReSyncAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


