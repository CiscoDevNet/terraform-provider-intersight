# VirtualizationVmwareVirtualMachineGpuAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualMachineGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualMachineGpu"]
**Key** | Pointer to **int64** | The internally assigned key of this virtual GPU device. This entity is not manipulated by users. | [optional] [readonly] 
**VmIdentity** | Pointer to **string** | Identity of the virtual machine where the virtual gpu is created. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareVirtualMachineGpuAllOf

`func NewVirtualizationVmwareVirtualMachineGpuAllOf(classId string, objectType string, ) *VirtualizationVmwareVirtualMachineGpuAllOf`

NewVirtualizationVmwareVirtualMachineGpuAllOf instantiates a new VirtualizationVmwareVirtualMachineGpuAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineGpuAllOfWithDefaults

`func NewVirtualizationVmwareVirtualMachineGpuAllOfWithDefaults() *VirtualizationVmwareVirtualMachineGpuAllOf`

NewVirtualizationVmwareVirtualMachineGpuAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualMachineGpuAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) SetKey(v int64)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVmIdentity

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetVmIdentity() string`

GetVmIdentity returns the VmIdentity field if non-nil, zero value otherwise.

### GetVmIdentityOk

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) GetVmIdentityOk() (*string, bool)`

GetVmIdentityOk returns a tuple with the VmIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIdentity

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) SetVmIdentity(v string)`

SetVmIdentity sets VmIdentity field to given value.

### HasVmIdentity

`func (o *VirtualizationVmwareVirtualMachineGpuAllOf) HasVmIdentity() bool`

HasVmIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


