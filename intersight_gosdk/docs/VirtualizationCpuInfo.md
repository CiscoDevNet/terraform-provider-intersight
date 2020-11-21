# VirtualizationCpuInfo

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

### NewVirtualizationCpuInfo

`func NewVirtualizationCpuInfo(classId string, objectType string, ) *VirtualizationCpuInfo`

NewVirtualizationCpuInfo instantiates a new VirtualizationCpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCpuInfoWithDefaults

`func NewVirtualizationCpuInfoWithDefaults() *VirtualizationCpuInfo`

NewVirtualizationCpuInfoWithDefaults instantiates a new VirtualizationCpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCpuInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCpuInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCpuInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCpuInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCpuInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCpuInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCores

`func (o *VirtualizationCpuInfo) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VirtualizationCpuInfo) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VirtualizationCpuInfo) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VirtualizationCpuInfo) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualizationCpuInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationCpuInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationCpuInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationCpuInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSockets

`func (o *VirtualizationCpuInfo) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *VirtualizationCpuInfo) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *VirtualizationCpuInfo) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *VirtualizationCpuInfo) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpeed

`func (o *VirtualizationCpuInfo) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualizationCpuInfo) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualizationCpuInfo) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualizationCpuInfo) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationCpuInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationCpuInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationCpuInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationCpuInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


