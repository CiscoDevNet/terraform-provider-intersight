# ComputeRootPcieEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "compute.CpuConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "compute.CpuConfig"]
**RootPcieEndpointType** | Pointer to **string** | Type of root PCIe endpoint device. * &#x60;CPU&#x60; - The root PCIe endpoint type is CPU. | [optional] [default to "CPU"]

## Methods

### NewComputeRootPcieEndpointConfig

`func NewComputeRootPcieEndpointConfig(classId string, objectType string, ) *ComputeRootPcieEndpointConfig`

NewComputeRootPcieEndpointConfig instantiates a new ComputeRootPcieEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRootPcieEndpointConfigWithDefaults

`func NewComputeRootPcieEndpointConfigWithDefaults() *ComputeRootPcieEndpointConfig`

NewComputeRootPcieEndpointConfigWithDefaults instantiates a new ComputeRootPcieEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeRootPcieEndpointConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeRootPcieEndpointConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeRootPcieEndpointConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeRootPcieEndpointConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeRootPcieEndpointConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeRootPcieEndpointConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRootPcieEndpointType

`func (o *ComputeRootPcieEndpointConfig) GetRootPcieEndpointType() string`

GetRootPcieEndpointType returns the RootPcieEndpointType field if non-nil, zero value otherwise.

### GetRootPcieEndpointTypeOk

`func (o *ComputeRootPcieEndpointConfig) GetRootPcieEndpointTypeOk() (*string, bool)`

GetRootPcieEndpointTypeOk returns a tuple with the RootPcieEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPcieEndpointType

`func (o *ComputeRootPcieEndpointConfig) SetRootPcieEndpointType(v string)`

SetRootPcieEndpointType sets RootPcieEndpointType field to given value.

### HasRootPcieEndpointType

`func (o *ComputeRootPcieEndpointConfig) HasRootPcieEndpointType() bool`

HasRootPcieEndpointType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


