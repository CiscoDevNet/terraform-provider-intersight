# CapabilityDimmsEndpointDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.DimmsEndpointDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.DimmsEndpointDescriptor"]
**Description** | Pointer to **string** | This field is to provide description of the dimm. | [optional] [readonly] 
**Model** | Pointer to **string** | This field is to provide model of the dimm. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field is to provide partNumber of the dimm. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field is to provide pid of the dimm. | [optional] [readonly] 
**SupportedPlatformsPids** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | This field is to provide vendor of the dimm. | [optional] [readonly] 

## Methods

### NewCapabilityDimmsEndpointDescriptor

`func NewCapabilityDimmsEndpointDescriptor(classId string, objectType string, ) *CapabilityDimmsEndpointDescriptor`

NewCapabilityDimmsEndpointDescriptor instantiates a new CapabilityDimmsEndpointDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDimmsEndpointDescriptorWithDefaults

`func NewCapabilityDimmsEndpointDescriptorWithDefaults() *CapabilityDimmsEndpointDescriptor`

NewCapabilityDimmsEndpointDescriptorWithDefaults instantiates a new CapabilityDimmsEndpointDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityDimmsEndpointDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityDimmsEndpointDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityDimmsEndpointDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityDimmsEndpointDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityDimmsEndpointDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityDimmsEndpointDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CapabilityDimmsEndpointDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityDimmsEndpointDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityDimmsEndpointDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityDimmsEndpointDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityDimmsEndpointDescriptor) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityDimmsEndpointDescriptor) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityDimmsEndpointDescriptor) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityDimmsEndpointDescriptor) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPartNumber

`func (o *CapabilityDimmsEndpointDescriptor) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CapabilityDimmsEndpointDescriptor) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CapabilityDimmsEndpointDescriptor) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *CapabilityDimmsEndpointDescriptor) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityDimmsEndpointDescriptor) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityDimmsEndpointDescriptor) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityDimmsEndpointDescriptor) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityDimmsEndpointDescriptor) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSupportedPlatformsPids

`func (o *CapabilityDimmsEndpointDescriptor) GetSupportedPlatformsPids() []string`

GetSupportedPlatformsPids returns the SupportedPlatformsPids field if non-nil, zero value otherwise.

### GetSupportedPlatformsPidsOk

`func (o *CapabilityDimmsEndpointDescriptor) GetSupportedPlatformsPidsOk() (*[]string, bool)`

GetSupportedPlatformsPidsOk returns a tuple with the SupportedPlatformsPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatformsPids

`func (o *CapabilityDimmsEndpointDescriptor) SetSupportedPlatformsPids(v []string)`

SetSupportedPlatformsPids sets SupportedPlatformsPids field to given value.

### HasSupportedPlatformsPids

`func (o *CapabilityDimmsEndpointDescriptor) HasSupportedPlatformsPids() bool`

HasSupportedPlatformsPids returns a boolean if a field has been set.

### SetSupportedPlatformsPidsNil

`func (o *CapabilityDimmsEndpointDescriptor) SetSupportedPlatformsPidsNil(b bool)`

 SetSupportedPlatformsPidsNil sets the value for SupportedPlatformsPids to be an explicit nil

### UnsetSupportedPlatformsPids
`func (o *CapabilityDimmsEndpointDescriptor) UnsetSupportedPlatformsPids()`

UnsetSupportedPlatformsPids ensures that no value is present for SupportedPlatformsPids, not even an explicit nil
### GetVendor

`func (o *CapabilityDimmsEndpointDescriptor) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityDimmsEndpointDescriptor) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityDimmsEndpointDescriptor) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityDimmsEndpointDescriptor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


