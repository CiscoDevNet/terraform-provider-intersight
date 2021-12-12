# FabricFcStorageRoleListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;fabric.FcStorageRole&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]FabricFcStorageRole**](FabricFcStorageRole.md) | The array of &#39;fabric.FcStorageRole&#39; resources matching the request. | [optional] 

## Methods

### NewFabricFcStorageRoleListAllOf

`func NewFabricFcStorageRoleListAllOf() *FabricFcStorageRoleListAllOf`

NewFabricFcStorageRoleListAllOf instantiates a new FabricFcStorageRoleListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcStorageRoleListAllOfWithDefaults

`func NewFabricFcStorageRoleListAllOfWithDefaults() *FabricFcStorageRoleListAllOf`

NewFabricFcStorageRoleListAllOfWithDefaults instantiates a new FabricFcStorageRoleListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FabricFcStorageRoleListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FabricFcStorageRoleListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FabricFcStorageRoleListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FabricFcStorageRoleListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *FabricFcStorageRoleListAllOf) GetResults() []FabricFcStorageRole`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FabricFcStorageRoleListAllOf) GetResultsOk() (*[]FabricFcStorageRole, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FabricFcStorageRoleListAllOf) SetResults(v []FabricFcStorageRole)`

SetResults sets Results field to given value.

### HasResults

`func (o *FabricFcStorageRoleListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *FabricFcStorageRoleListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *FabricFcStorageRoleListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


