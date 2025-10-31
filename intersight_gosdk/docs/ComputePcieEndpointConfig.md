# ComputePcieEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**PcieEndpointType** | Pointer to **string** | Type of PCIe endpoint device. * &#x60;GPU&#x60; - The PCIe endpoint type is GPU. * &#x60;Adapter&#x60; - The PCIe endpoint type is Adapter. | [optional] [default to "GPU"]

## Methods

### NewComputePcieEndpointConfig

`func NewComputePcieEndpointConfig(classId string, objectType string, ) *ComputePcieEndpointConfig`

NewComputePcieEndpointConfig instantiates a new ComputePcieEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePcieEndpointConfigWithDefaults

`func NewComputePcieEndpointConfigWithDefaults() *ComputePcieEndpointConfig`

NewComputePcieEndpointConfigWithDefaults instantiates a new ComputePcieEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePcieEndpointConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePcieEndpointConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePcieEndpointConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePcieEndpointConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePcieEndpointConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePcieEndpointConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcieEndpointType

`func (o *ComputePcieEndpointConfig) GetPcieEndpointType() string`

GetPcieEndpointType returns the PcieEndpointType field if non-nil, zero value otherwise.

### GetPcieEndpointTypeOk

`func (o *ComputePcieEndpointConfig) GetPcieEndpointTypeOk() (*string, bool)`

GetPcieEndpointTypeOk returns a tuple with the PcieEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieEndpointType

`func (o *ComputePcieEndpointConfig) SetPcieEndpointType(v string)`

SetPcieEndpointType sets PcieEndpointType field to given value.

### HasPcieEndpointType

`func (o *ComputePcieEndpointConfig) HasPcieEndpointType() bool`

HasPcieEndpointType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


