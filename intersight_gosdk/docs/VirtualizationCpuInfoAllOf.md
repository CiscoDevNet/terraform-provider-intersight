# VirtualizationCpuInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.CpuInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.CpuInfo"]
**Cores** | Pointer to **int64** | Number of cores per CPU, as reported by the manufacturer. | [optional] 
**Description** | Pointer to **string** | The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz. | [optional] 
**Sockets** | Pointer to **int64** | Number of CPU sockets available. | [optional] 
**Speed** | Pointer to **int64** | Speed of the CPUs in Hertz. For example, 2593749663. | [optional] 
**Vendor** | Pointer to **string** | Manufacturer of the CPU . For example, Intel. | [optional] 

## Methods

### NewVirtualizationCpuInfoAllOf

`func NewVirtualizationCpuInfoAllOf(classId string, objectType string, ) *VirtualizationCpuInfoAllOf`

NewVirtualizationCpuInfoAllOf instantiates a new VirtualizationCpuInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCpuInfoAllOfWithDefaults

`func NewVirtualizationCpuInfoAllOfWithDefaults() *VirtualizationCpuInfoAllOf`

NewVirtualizationCpuInfoAllOfWithDefaults instantiates a new VirtualizationCpuInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCpuInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCpuInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCpuInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCpuInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCpuInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCpuInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCores

`func (o *VirtualizationCpuInfoAllOf) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VirtualizationCpuInfoAllOf) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VirtualizationCpuInfoAllOf) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VirtualizationCpuInfoAllOf) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationCpuInfoAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationCpuInfoAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationCpuInfoAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationCpuInfoAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSockets

`func (o *VirtualizationCpuInfoAllOf) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *VirtualizationCpuInfoAllOf) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *VirtualizationCpuInfoAllOf) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *VirtualizationCpuInfoAllOf) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpeed

`func (o *VirtualizationCpuInfoAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualizationCpuInfoAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualizationCpuInfoAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualizationCpuInfoAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationCpuInfoAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationCpuInfoAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationCpuInfoAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationCpuInfoAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


