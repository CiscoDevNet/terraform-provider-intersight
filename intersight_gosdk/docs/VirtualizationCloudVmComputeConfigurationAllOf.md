# VirtualizationCloudVmComputeConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmComputeConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmComputeConfiguration"]
**InstanceTypeId** | Pointer to **string** | Instance Type used by this VM. | [optional] 

## Methods

### NewVirtualizationCloudVmComputeConfigurationAllOf

`func NewVirtualizationCloudVmComputeConfigurationAllOf(classId string, objectType string, ) *VirtualizationCloudVmComputeConfigurationAllOf`

NewVirtualizationCloudVmComputeConfigurationAllOf instantiates a new VirtualizationCloudVmComputeConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCloudVmComputeConfigurationAllOfWithDefaults

`func NewVirtualizationCloudVmComputeConfigurationAllOfWithDefaults() *VirtualizationCloudVmComputeConfigurationAllOf`

NewVirtualizationCloudVmComputeConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmComputeConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInstanceTypeId

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *VirtualizationCloudVmComputeConfigurationAllOf) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


