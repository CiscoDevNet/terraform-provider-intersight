# HyperflexStartReduceReSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.StartReduceReSync"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.StartReduceReSync"]
**ClusterMoIds** | Pointer to **[]string** |  | [optional] 
**Operation** | Pointer to **string** | The cleanup operation to perform. * &#x60;NoOp&#x60; - Does not perform any operation when the API is called. * &#x60;StartReduceResync&#x60; - Start the execution of reduce re-sync and stale mirror cleanup for the HyperFlex clusters associated with the account. | [optional] [default to "NoOp"]
**TargetDetails** | Pointer to [**[]HyperflexReSyncClusterMoIds**](HyperflexReSyncClusterMoIds.md) |  | [optional] 

## Methods

### NewHyperflexStartReduceReSync

`func NewHyperflexStartReduceReSync(classId string, objectType string, ) *HyperflexStartReduceReSync`

NewHyperflexStartReduceReSync instantiates a new HyperflexStartReduceReSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexStartReduceReSyncWithDefaults

`func NewHyperflexStartReduceReSyncWithDefaults() *HyperflexStartReduceReSync`

NewHyperflexStartReduceReSyncWithDefaults instantiates a new HyperflexStartReduceReSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexStartReduceReSync) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexStartReduceReSync) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexStartReduceReSync) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexStartReduceReSync) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexStartReduceReSync) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexStartReduceReSync) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterMoIds

`func (o *HyperflexStartReduceReSync) GetClusterMoIds() []string`

GetClusterMoIds returns the ClusterMoIds field if non-nil, zero value otherwise.

### GetClusterMoIdsOk

`func (o *HyperflexStartReduceReSync) GetClusterMoIdsOk() (*[]string, bool)`

GetClusterMoIdsOk returns a tuple with the ClusterMoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMoIds

`func (o *HyperflexStartReduceReSync) SetClusterMoIds(v []string)`

SetClusterMoIds sets ClusterMoIds field to given value.

### HasClusterMoIds

`func (o *HyperflexStartReduceReSync) HasClusterMoIds() bool`

HasClusterMoIds returns a boolean if a field has been set.

### SetClusterMoIdsNil

`func (o *HyperflexStartReduceReSync) SetClusterMoIdsNil(b bool)`

 SetClusterMoIdsNil sets the value for ClusterMoIds to be an explicit nil

### UnsetClusterMoIds
`func (o *HyperflexStartReduceReSync) UnsetClusterMoIds()`

UnsetClusterMoIds ensures that no value is present for ClusterMoIds, not even an explicit nil
### GetOperation

`func (o *HyperflexStartReduceReSync) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HyperflexStartReduceReSync) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HyperflexStartReduceReSync) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HyperflexStartReduceReSync) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTargetDetails

`func (o *HyperflexStartReduceReSync) GetTargetDetails() []HyperflexReSyncClusterMoIds`

GetTargetDetails returns the TargetDetails field if non-nil, zero value otherwise.

### GetTargetDetailsOk

`func (o *HyperflexStartReduceReSync) GetTargetDetailsOk() (*[]HyperflexReSyncClusterMoIds, bool)`

GetTargetDetailsOk returns a tuple with the TargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDetails

`func (o *HyperflexStartReduceReSync) SetTargetDetails(v []HyperflexReSyncClusterMoIds)`

SetTargetDetails sets TargetDetails field to given value.

### HasTargetDetails

`func (o *HyperflexStartReduceReSync) HasTargetDetails() bool`

HasTargetDetails returns a boolean if a field has been set.

### SetTargetDetailsNil

`func (o *HyperflexStartReduceReSync) SetTargetDetailsNil(b bool)`

 SetTargetDetailsNil sets the value for TargetDetails to be an explicit nil

### UnsetTargetDetails
`func (o *HyperflexStartReduceReSync) UnsetTargetDetails()`

UnsetTargetDetails ensures that no value is present for TargetDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


