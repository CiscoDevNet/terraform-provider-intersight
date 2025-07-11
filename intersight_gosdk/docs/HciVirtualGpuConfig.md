# HciVirtualGpuConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.VirtualGpuConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.VirtualGpuConfig"]
**Assignable** | Pointer to **int64** | Indicates if the virtual GPU is assignable. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device id of the virtual GPU. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | The name of the virtual GPU. | [optional] [readonly] 
**Fraction** | Pointer to **int64** | The fraction of the virtual GPU. | [optional] [readonly] 
**FrameBufferSizeBytes** | Pointer to **int64** | The frame buffer size of the virtual GPU. | [optional] [readonly] 
**GuestDriverVersion** | Pointer to **string** | The guest driver version of the virtual GPU. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Indicates if the virtual GPU is in use. | [optional] [readonly] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**MaxResolution** | Pointer to **string** | The maximum resolution per display heads. | [optional] [readonly] 
**Mode** | Pointer to **string** | The mode of the virtual GPU such as UNUSED, USED_FOR_PASSTHROUGH, USED_FOR_VIRTUAL. | [optional] [readonly] 
**NumaNode** | Pointer to **string** | The NUMA node of the virtual GPU. | [optional] [readonly] 
**NumberOfVirtualDisplayHeads** | Pointer to **int64** | The number of virtual display heads of the virtual GPU. | [optional] [readonly] 
**Sbdf** | Pointer to **string** | The SBDF address of the virtual GPU. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the virtual GPU such as PASS_THROUGH_GRAPHICS, PASS_THROUGH_COMPUTE, VIRTUAL. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor name of the virtual GPU. | [optional] [readonly] 

## Methods

### NewHciVirtualGpuConfig

`func NewHciVirtualGpuConfig(classId string, objectType string, ) *HciVirtualGpuConfig`

NewHciVirtualGpuConfig instantiates a new HciVirtualGpuConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciVirtualGpuConfigWithDefaults

`func NewHciVirtualGpuConfigWithDefaults() *HciVirtualGpuConfig`

NewHciVirtualGpuConfigWithDefaults instantiates a new HciVirtualGpuConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciVirtualGpuConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciVirtualGpuConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciVirtualGpuConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciVirtualGpuConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciVirtualGpuConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciVirtualGpuConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssignable

`func (o *HciVirtualGpuConfig) GetAssignable() int64`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *HciVirtualGpuConfig) GetAssignableOk() (*int64, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *HciVirtualGpuConfig) SetAssignable(v int64)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *HciVirtualGpuConfig) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetDeviceId

`func (o *HciVirtualGpuConfig) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HciVirtualGpuConfig) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HciVirtualGpuConfig) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HciVirtualGpuConfig) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *HciVirtualGpuConfig) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *HciVirtualGpuConfig) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *HciVirtualGpuConfig) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *HciVirtualGpuConfig) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetFraction

`func (o *HciVirtualGpuConfig) GetFraction() int64`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *HciVirtualGpuConfig) GetFractionOk() (*int64, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *HciVirtualGpuConfig) SetFraction(v int64)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *HciVirtualGpuConfig) HasFraction() bool`

HasFraction returns a boolean if a field has been set.

### GetFrameBufferSizeBytes

`func (o *HciVirtualGpuConfig) GetFrameBufferSizeBytes() int64`

GetFrameBufferSizeBytes returns the FrameBufferSizeBytes field if non-nil, zero value otherwise.

### GetFrameBufferSizeBytesOk

`func (o *HciVirtualGpuConfig) GetFrameBufferSizeBytesOk() (*int64, bool)`

GetFrameBufferSizeBytesOk returns a tuple with the FrameBufferSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeBytes

`func (o *HciVirtualGpuConfig) SetFrameBufferSizeBytes(v int64)`

SetFrameBufferSizeBytes sets FrameBufferSizeBytes field to given value.

### HasFrameBufferSizeBytes

`func (o *HciVirtualGpuConfig) HasFrameBufferSizeBytes() bool`

HasFrameBufferSizeBytes returns a boolean if a field has been set.

### GetGuestDriverVersion

`func (o *HciVirtualGpuConfig) GetGuestDriverVersion() string`

GetGuestDriverVersion returns the GuestDriverVersion field if non-nil, zero value otherwise.

### GetGuestDriverVersionOk

`func (o *HciVirtualGpuConfig) GetGuestDriverVersionOk() (*string, bool)`

GetGuestDriverVersionOk returns a tuple with the GuestDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestDriverVersion

`func (o *HciVirtualGpuConfig) SetGuestDriverVersion(v string)`

SetGuestDriverVersion sets GuestDriverVersion field to given value.

### HasGuestDriverVersion

`func (o *HciVirtualGpuConfig) HasGuestDriverVersion() bool`

HasGuestDriverVersion returns a boolean if a field has been set.

### GetInUse

`func (o *HciVirtualGpuConfig) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *HciVirtualGpuConfig) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *HciVirtualGpuConfig) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *HciVirtualGpuConfig) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetLicenses

`func (o *HciVirtualGpuConfig) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *HciVirtualGpuConfig) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *HciVirtualGpuConfig) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *HciVirtualGpuConfig) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### SetLicensesNil

`func (o *HciVirtualGpuConfig) SetLicensesNil(b bool)`

 SetLicensesNil sets the value for Licenses to be an explicit nil

### UnsetLicenses
`func (o *HciVirtualGpuConfig) UnsetLicenses()`

UnsetLicenses ensures that no value is present for Licenses, not even an explicit nil
### GetMaxResolution

`func (o *HciVirtualGpuConfig) GetMaxResolution() string`

GetMaxResolution returns the MaxResolution field if non-nil, zero value otherwise.

### GetMaxResolutionOk

`func (o *HciVirtualGpuConfig) GetMaxResolutionOk() (*string, bool)`

GetMaxResolutionOk returns a tuple with the MaxResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolution

`func (o *HciVirtualGpuConfig) SetMaxResolution(v string)`

SetMaxResolution sets MaxResolution field to given value.

### HasMaxResolution

`func (o *HciVirtualGpuConfig) HasMaxResolution() bool`

HasMaxResolution returns a boolean if a field has been set.

### GetMode

`func (o *HciVirtualGpuConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HciVirtualGpuConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HciVirtualGpuConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HciVirtualGpuConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumaNode

`func (o *HciVirtualGpuConfig) GetNumaNode() string`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *HciVirtualGpuConfig) GetNumaNodeOk() (*string, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *HciVirtualGpuConfig) SetNumaNode(v string)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *HciVirtualGpuConfig) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpuConfig) GetNumberOfVirtualDisplayHeads() int64`

GetNumberOfVirtualDisplayHeads returns the NumberOfVirtualDisplayHeads field if non-nil, zero value otherwise.

### GetNumberOfVirtualDisplayHeadsOk

`func (o *HciVirtualGpuConfig) GetNumberOfVirtualDisplayHeadsOk() (*int64, bool)`

GetNumberOfVirtualDisplayHeadsOk returns a tuple with the NumberOfVirtualDisplayHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpuConfig) SetNumberOfVirtualDisplayHeads(v int64)`

SetNumberOfVirtualDisplayHeads sets NumberOfVirtualDisplayHeads field to given value.

### HasNumberOfVirtualDisplayHeads

`func (o *HciVirtualGpuConfig) HasNumberOfVirtualDisplayHeads() bool`

HasNumberOfVirtualDisplayHeads returns a boolean if a field has been set.

### GetSbdf

`func (o *HciVirtualGpuConfig) GetSbdf() string`

GetSbdf returns the Sbdf field if non-nil, zero value otherwise.

### GetSbdfOk

`func (o *HciVirtualGpuConfig) GetSbdfOk() (*string, bool)`

GetSbdfOk returns a tuple with the Sbdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbdf

`func (o *HciVirtualGpuConfig) SetSbdf(v string)`

SetSbdf sets Sbdf field to given value.

### HasSbdf

`func (o *HciVirtualGpuConfig) HasSbdf() bool`

HasSbdf returns a boolean if a field has been set.

### GetType

`func (o *HciVirtualGpuConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HciVirtualGpuConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HciVirtualGpuConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HciVirtualGpuConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *HciVirtualGpuConfig) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HciVirtualGpuConfig) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HciVirtualGpuConfig) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HciVirtualGpuConfig) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


