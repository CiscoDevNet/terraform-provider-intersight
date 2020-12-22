# VirtualizationEsxiVmComputeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmComputeConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmComputeConfiguration"]
**ResourcePool** | Pointer to **string** | ResourcePool where virtual machine is deployed. | [optional] 

## Methods

### NewVirtualizationEsxiVmComputeConfiguration

`func NewVirtualizationEsxiVmComputeConfiguration(classId string, objectType string, ) *VirtualizationEsxiVmComputeConfiguration`

NewVirtualizationEsxiVmComputeConfiguration instantiates a new VirtualizationEsxiVmComputeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmComputeConfigurationWithDefaults

`func NewVirtualizationEsxiVmComputeConfigurationWithDefaults() *VirtualizationEsxiVmComputeConfiguration`

NewVirtualizationEsxiVmComputeConfigurationWithDefaults instantiates a new VirtualizationEsxiVmComputeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmComputeConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmComputeConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmComputeConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmComputeConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmComputeConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmComputeConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetResourcePool

`func (o *VirtualizationEsxiVmComputeConfiguration) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VirtualizationEsxiVmComputeConfiguration) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VirtualizationEsxiVmComputeConfiguration) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VirtualizationEsxiVmComputeConfiguration) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


