# CapabilityServerPcieConnectivityLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerPcieConnectivityLayout"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerPcieConnectivityLayout"]
**CpuId** | Pointer to **int64** | Unique identifier for the server CPU in the physical connection topology. | [optional] 
**PcieTopology** | Pointer to [**[]CapabilityServerPcieTopology**](CapabilityServerPcieTopology.md) |  | [optional] 
**ServerSlots** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCapabilityServerPcieConnectivityLayout

`func NewCapabilityServerPcieConnectivityLayout(classId string, objectType string, ) *CapabilityServerPcieConnectivityLayout`

NewCapabilityServerPcieConnectivityLayout instantiates a new CapabilityServerPcieConnectivityLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerPcieConnectivityLayoutWithDefaults

`func NewCapabilityServerPcieConnectivityLayoutWithDefaults() *CapabilityServerPcieConnectivityLayout`

NewCapabilityServerPcieConnectivityLayoutWithDefaults instantiates a new CapabilityServerPcieConnectivityLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerPcieConnectivityLayout) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerPcieConnectivityLayout) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerPcieConnectivityLayout) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerPcieConnectivityLayout) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerPcieConnectivityLayout) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerPcieConnectivityLayout) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuId

`func (o *CapabilityServerPcieConnectivityLayout) GetCpuId() int64`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *CapabilityServerPcieConnectivityLayout) GetCpuIdOk() (*int64, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *CapabilityServerPcieConnectivityLayout) SetCpuId(v int64)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *CapabilityServerPcieConnectivityLayout) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetPcieTopology

`func (o *CapabilityServerPcieConnectivityLayout) GetPcieTopology() []CapabilityServerPcieTopology`

GetPcieTopology returns the PcieTopology field if non-nil, zero value otherwise.

### GetPcieTopologyOk

`func (o *CapabilityServerPcieConnectivityLayout) GetPcieTopologyOk() (*[]CapabilityServerPcieTopology, bool)`

GetPcieTopologyOk returns a tuple with the PcieTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieTopology

`func (o *CapabilityServerPcieConnectivityLayout) SetPcieTopology(v []CapabilityServerPcieTopology)`

SetPcieTopology sets PcieTopology field to given value.

### HasPcieTopology

`func (o *CapabilityServerPcieConnectivityLayout) HasPcieTopology() bool`

HasPcieTopology returns a boolean if a field has been set.

### SetPcieTopologyNil

`func (o *CapabilityServerPcieConnectivityLayout) SetPcieTopologyNil(b bool)`

 SetPcieTopologyNil sets the value for PcieTopology to be an explicit nil

### UnsetPcieTopology
`func (o *CapabilityServerPcieConnectivityLayout) UnsetPcieTopology()`

UnsetPcieTopology ensures that no value is present for PcieTopology, not even an explicit nil
### GetServerSlots

`func (o *CapabilityServerPcieConnectivityLayout) GetServerSlots() []int64`

GetServerSlots returns the ServerSlots field if non-nil, zero value otherwise.

### GetServerSlotsOk

`func (o *CapabilityServerPcieConnectivityLayout) GetServerSlotsOk() (*[]int64, bool)`

GetServerSlotsOk returns a tuple with the ServerSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSlots

`func (o *CapabilityServerPcieConnectivityLayout) SetServerSlots(v []int64)`

SetServerSlots sets ServerSlots field to given value.

### HasServerSlots

`func (o *CapabilityServerPcieConnectivityLayout) HasServerSlots() bool`

HasServerSlots returns a boolean if a field has been set.

### SetServerSlotsNil

`func (o *CapabilityServerPcieConnectivityLayout) SetServerSlotsNil(b bool)`

 SetServerSlotsNil sets the value for ServerSlots to be an explicit nil

### UnsetServerSlots
`func (o *CapabilityServerPcieConnectivityLayout) UnsetServerSlots()`

UnsetServerSlots ensures that no value is present for ServerSlots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


