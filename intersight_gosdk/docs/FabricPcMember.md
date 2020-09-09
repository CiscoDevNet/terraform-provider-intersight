# FabricPcMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcId** | Pointer to **int64** | Port Channel Identifier for the collection of ports. | [optional] 
**PortPolicy** | Pointer to [**FabricPortPolicyRelationship**](fabric.PortPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricPcMember

`func NewFabricPcMember() *FabricPcMember`

NewFabricPcMember instantiates a new FabricPcMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPcMemberWithDefaults

`func NewFabricPcMemberWithDefaults() *FabricPcMember`

NewFabricPcMemberWithDefaults instantiates a new FabricPcMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcId

`func (o *FabricPcMember) GetPcId() int64`

GetPcId returns the PcId field if non-nil, zero value otherwise.

### GetPcIdOk

`func (o *FabricPcMember) GetPcIdOk() (*int64, bool)`

GetPcIdOk returns a tuple with the PcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcId

`func (o *FabricPcMember) SetPcId(v int64)`

SetPcId sets PcId field to given value.

### HasPcId

`func (o *FabricPcMember) HasPcId() bool`

HasPcId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricPcMember) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricPcMember) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricPcMember) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricPcMember) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


