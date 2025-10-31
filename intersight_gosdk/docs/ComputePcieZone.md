# ComputePcieZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PcieZone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PcieZone"]
**PcieEndpoints** | Pointer to [**[]ComputePcieEndpointConfig**](ComputePcieEndpointConfig.md) |  | [optional] 
**RootPcieEndpoint** | Pointer to [**NullableComputeRootPcieEndpointConfig**](ComputeRootPcieEndpointConfig.md) |  | [optional] 

## Methods

### NewComputePcieZone

`func NewComputePcieZone(classId string, objectType string, ) *ComputePcieZone`

NewComputePcieZone instantiates a new ComputePcieZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePcieZoneWithDefaults

`func NewComputePcieZoneWithDefaults() *ComputePcieZone`

NewComputePcieZoneWithDefaults instantiates a new ComputePcieZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePcieZone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePcieZone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePcieZone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePcieZone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePcieZone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePcieZone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcieEndpoints

`func (o *ComputePcieZone) GetPcieEndpoints() []ComputePcieEndpointConfig`

GetPcieEndpoints returns the PcieEndpoints field if non-nil, zero value otherwise.

### GetPcieEndpointsOk

`func (o *ComputePcieZone) GetPcieEndpointsOk() (*[]ComputePcieEndpointConfig, bool)`

GetPcieEndpointsOk returns a tuple with the PcieEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieEndpoints

`func (o *ComputePcieZone) SetPcieEndpoints(v []ComputePcieEndpointConfig)`

SetPcieEndpoints sets PcieEndpoints field to given value.

### HasPcieEndpoints

`func (o *ComputePcieZone) HasPcieEndpoints() bool`

HasPcieEndpoints returns a boolean if a field has been set.

### SetPcieEndpointsNil

`func (o *ComputePcieZone) SetPcieEndpointsNil(b bool)`

 SetPcieEndpointsNil sets the value for PcieEndpoints to be an explicit nil

### UnsetPcieEndpoints
`func (o *ComputePcieZone) UnsetPcieEndpoints()`

UnsetPcieEndpoints ensures that no value is present for PcieEndpoints, not even an explicit nil
### GetRootPcieEndpoint

`func (o *ComputePcieZone) GetRootPcieEndpoint() ComputeRootPcieEndpointConfig`

GetRootPcieEndpoint returns the RootPcieEndpoint field if non-nil, zero value otherwise.

### GetRootPcieEndpointOk

`func (o *ComputePcieZone) GetRootPcieEndpointOk() (*ComputeRootPcieEndpointConfig, bool)`

GetRootPcieEndpointOk returns a tuple with the RootPcieEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPcieEndpoint

`func (o *ComputePcieZone) SetRootPcieEndpoint(v ComputeRootPcieEndpointConfig)`

SetRootPcieEndpoint sets RootPcieEndpoint field to given value.

### HasRootPcieEndpoint

`func (o *ComputePcieZone) HasRootPcieEndpoint() bool`

HasRootPcieEndpoint returns a boolean if a field has been set.

### SetRootPcieEndpointNil

`func (o *ComputePcieZone) SetRootPcieEndpointNil(b bool)`

 SetRootPcieEndpointNil sets the value for RootPcieEndpoint to be an explicit nil

### UnsetRootPcieEndpoint
`func (o *ComputePcieZone) UnsetRootPcieEndpoint()`

UnsetRootPcieEndpoint ensures that no value is present for RootPcieEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


