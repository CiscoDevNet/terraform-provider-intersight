# FabricPcOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | Admin configured state to disable the port channel. * &#x60;Enabled&#x60; - Admin configured Enabled State. * &#x60;Disabled&#x60; - Admin configured Disabled State. | [optional] [default to "Enabled"]
**PcId** | Pointer to **int64** | Port Channel Identifier for the collection of ports. | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 

## Methods

### NewFabricPcOperation

`func NewFabricPcOperation() *FabricPcOperation`

NewFabricPcOperation instantiates a new FabricPcOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPcOperationWithDefaults

`func NewFabricPcOperationWithDefaults() *FabricPcOperation`

NewFabricPcOperationWithDefaults instantiates a new FabricPcOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *FabricPcOperation) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricPcOperation) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricPcOperation) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricPcOperation) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetPcId

`func (o *FabricPcOperation) GetPcId() int64`

GetPcId returns the PcId field if non-nil, zero value otherwise.

### GetPcIdOk

`func (o *FabricPcOperation) GetPcIdOk() (*int64, bool)`

GetPcIdOk returns a tuple with the PcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcId

`func (o *FabricPcOperation) SetPcId(v int64)`

SetPcId sets PcId field to given value.

### HasPcId

`func (o *FabricPcOperation) HasPcId() bool`

HasPcId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricPcOperation) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricPcOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricPcOperation) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricPcOperation) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


