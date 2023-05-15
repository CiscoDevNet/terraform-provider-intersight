# CapabilityGpuEndpointDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.GpuEndpointDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.GpuEndpointDescriptor"]
**Description** | Pointer to **string** | This field is to provide description of the gpu. | [optional] [readonly] 
**Model** | Pointer to **string** | This field is to provide model of the gpu. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field is to provide partNumber of the gpu. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field is to provide pid of the gpu. | [optional] [readonly] 
**SupportedPlatformsPids** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | This field is to provide vendor of the gpu. | [optional] [readonly] 

## Methods

### NewCapabilityGpuEndpointDescriptor

`func NewCapabilityGpuEndpointDescriptor(classId string, objectType string, ) *CapabilityGpuEndpointDescriptor`

NewCapabilityGpuEndpointDescriptor instantiates a new CapabilityGpuEndpointDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityGpuEndpointDescriptorWithDefaults

`func NewCapabilityGpuEndpointDescriptorWithDefaults() *CapabilityGpuEndpointDescriptor`

NewCapabilityGpuEndpointDescriptorWithDefaults instantiates a new CapabilityGpuEndpointDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityGpuEndpointDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityGpuEndpointDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityGpuEndpointDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityGpuEndpointDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityGpuEndpointDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityGpuEndpointDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityGpuEndpointDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityGpuEndpointDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityGpuEndpointDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityGpuEndpointDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityGpuEndpointDescriptor) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityGpuEndpointDescriptor) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityGpuEndpointDescriptor) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityGpuEndpointDescriptor) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPartNumber

`func (o *CapabilityGpuEndpointDescriptor) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CapabilityGpuEndpointDescriptor) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CapabilityGpuEndpointDescriptor) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *CapabilityGpuEndpointDescriptor) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityGpuEndpointDescriptor) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityGpuEndpointDescriptor) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityGpuEndpointDescriptor) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityGpuEndpointDescriptor) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSupportedPlatformsPids

`func (o *CapabilityGpuEndpointDescriptor) GetSupportedPlatformsPids() []string`

GetSupportedPlatformsPids returns the SupportedPlatformsPids field if non-nil, zero value otherwise.

### GetSupportedPlatformsPidsOk

`func (o *CapabilityGpuEndpointDescriptor) GetSupportedPlatformsPidsOk() (*[]string, bool)`

GetSupportedPlatformsPidsOk returns a tuple with the SupportedPlatformsPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatformsPids

`func (o *CapabilityGpuEndpointDescriptor) SetSupportedPlatformsPids(v []string)`

SetSupportedPlatformsPids sets SupportedPlatformsPids field to given value.

### HasSupportedPlatformsPids

`func (o *CapabilityGpuEndpointDescriptor) HasSupportedPlatformsPids() bool`

HasSupportedPlatformsPids returns a boolean if a field has been set.

### SetSupportedPlatformsPidsNil

`func (o *CapabilityGpuEndpointDescriptor) SetSupportedPlatformsPidsNil(b bool)`

 SetSupportedPlatformsPidsNil sets the value for SupportedPlatformsPids to be an explicit nil

### UnsetSupportedPlatformsPids
`func (o *CapabilityGpuEndpointDescriptor) UnsetSupportedPlatformsPids()`

UnsetSupportedPlatformsPids ensures that no value is present for SupportedPlatformsPids, not even an explicit nil
### GetVendor

`func (o *CapabilityGpuEndpointDescriptor) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityGpuEndpointDescriptor) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityGpuEndpointDescriptor) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityGpuEndpointDescriptor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


