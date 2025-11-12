# CapabilityServerPcieTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerPcieTopology"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerPcieTopology"]
**ConnectedAdapter** | Pointer to [**[]CapabilityAdapterConnectionPoint**](CapabilityAdapterConnectionPoint.md) |  | [optional] 
**ConnectedGpu** | Pointer to [**[]CapabilityGpuConnectionPoint**](CapabilityGpuConnectionPoint.md) |  | [optional] 
**PcieSwitchId** | Pointer to **int64** | Unique identifier for the physical PCIe switch in the server topology. | [optional] 
**XfmId** | Pointer to **int64** | Unique identifier for the XFM (Switch Fabric Module) in the server topology. | [optional] 

## Methods

### NewCapabilityServerPcieTopology

`func NewCapabilityServerPcieTopology(classId string, objectType string, ) *CapabilityServerPcieTopology`

NewCapabilityServerPcieTopology instantiates a new CapabilityServerPcieTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerPcieTopologyWithDefaults

`func NewCapabilityServerPcieTopologyWithDefaults() *CapabilityServerPcieTopology`

NewCapabilityServerPcieTopologyWithDefaults instantiates a new CapabilityServerPcieTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerPcieTopology) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerPcieTopology) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerPcieTopology) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerPcieTopology) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerPcieTopology) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerPcieTopology) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectedAdapter

`func (o *CapabilityServerPcieTopology) GetConnectedAdapter() []CapabilityAdapterConnectionPoint`

GetConnectedAdapter returns the ConnectedAdapter field if non-nil, zero value otherwise.

### GetConnectedAdapterOk

`func (o *CapabilityServerPcieTopology) GetConnectedAdapterOk() (*[]CapabilityAdapterConnectionPoint, bool)`

GetConnectedAdapterOk returns a tuple with the ConnectedAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAdapter

`func (o *CapabilityServerPcieTopology) SetConnectedAdapter(v []CapabilityAdapterConnectionPoint)`

SetConnectedAdapter sets ConnectedAdapter field to given value.

### HasConnectedAdapter

`func (o *CapabilityServerPcieTopology) HasConnectedAdapter() bool`

HasConnectedAdapter returns a boolean if a field has been set.

### SetConnectedAdapterNil

`func (o *CapabilityServerPcieTopology) SetConnectedAdapterNil(b bool)`

 SetConnectedAdapterNil sets the value for ConnectedAdapter to be an explicit nil

### UnsetConnectedAdapter
`func (o *CapabilityServerPcieTopology) UnsetConnectedAdapter()`

UnsetConnectedAdapter ensures that no value is present for ConnectedAdapter, not even an explicit nil
### GetConnectedGpu

`func (o *CapabilityServerPcieTopology) GetConnectedGpu() []CapabilityGpuConnectionPoint`

GetConnectedGpu returns the ConnectedGpu field if non-nil, zero value otherwise.

### GetConnectedGpuOk

`func (o *CapabilityServerPcieTopology) GetConnectedGpuOk() (*[]CapabilityGpuConnectionPoint, bool)`

GetConnectedGpuOk returns a tuple with the ConnectedGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGpu

`func (o *CapabilityServerPcieTopology) SetConnectedGpu(v []CapabilityGpuConnectionPoint)`

SetConnectedGpu sets ConnectedGpu field to given value.

### HasConnectedGpu

`func (o *CapabilityServerPcieTopology) HasConnectedGpu() bool`

HasConnectedGpu returns a boolean if a field has been set.

### SetConnectedGpuNil

`func (o *CapabilityServerPcieTopology) SetConnectedGpuNil(b bool)`

 SetConnectedGpuNil sets the value for ConnectedGpu to be an explicit nil

### UnsetConnectedGpu
`func (o *CapabilityServerPcieTopology) UnsetConnectedGpu()`

UnsetConnectedGpu ensures that no value is present for ConnectedGpu, not even an explicit nil
### GetPcieSwitchId

`func (o *CapabilityServerPcieTopology) GetPcieSwitchId() int64`

GetPcieSwitchId returns the PcieSwitchId field if non-nil, zero value otherwise.

### GetPcieSwitchIdOk

`func (o *CapabilityServerPcieTopology) GetPcieSwitchIdOk() (*int64, bool)`

GetPcieSwitchIdOk returns a tuple with the PcieSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSwitchId

`func (o *CapabilityServerPcieTopology) SetPcieSwitchId(v int64)`

SetPcieSwitchId sets PcieSwitchId field to given value.

### HasPcieSwitchId

`func (o *CapabilityServerPcieTopology) HasPcieSwitchId() bool`

HasPcieSwitchId returns a boolean if a field has been set.

### GetXfmId

`func (o *CapabilityServerPcieTopology) GetXfmId() int64`

GetXfmId returns the XfmId field if non-nil, zero value otherwise.

### GetXfmIdOk

`func (o *CapabilityServerPcieTopology) GetXfmIdOk() (*int64, bool)`

GetXfmIdOk returns a tuple with the XfmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXfmId

`func (o *CapabilityServerPcieTopology) SetXfmId(v int64)`

SetXfmId sets XfmId field to given value.

### HasXfmId

`func (o *CapabilityServerPcieTopology) HasXfmId() bool`

HasXfmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


