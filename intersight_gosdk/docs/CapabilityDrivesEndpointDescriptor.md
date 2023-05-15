# CapabilityDrivesEndpointDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.DrivesEndpointDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.DrivesEndpointDescriptor"]
**AliasModel** | Pointer to **string** | This field is to provide alias model of the drive. | [optional] [readonly] 
**Description** | Pointer to **string** | This field is to provide description of the drive. | [optional] [readonly] 
**Model** | Pointer to **string** | This field is to provide model of the drive. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field is to provide partNumber of the drive. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field is to provide pid of the drive. | [optional] [readonly] 
**SupportedPlatformsPids** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** | This field is to provide vendor of the drive. | [optional] [readonly] 

## Methods

### NewCapabilityDrivesEndpointDescriptor

`func NewCapabilityDrivesEndpointDescriptor(classId string, objectType string, ) *CapabilityDrivesEndpointDescriptor`

NewCapabilityDrivesEndpointDescriptor instantiates a new CapabilityDrivesEndpointDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDrivesEndpointDescriptorWithDefaults

`func NewCapabilityDrivesEndpointDescriptorWithDefaults() *CapabilityDrivesEndpointDescriptor`

NewCapabilityDrivesEndpointDescriptorWithDefaults instantiates a new CapabilityDrivesEndpointDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityDrivesEndpointDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityDrivesEndpointDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityDrivesEndpointDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityDrivesEndpointDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityDrivesEndpointDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityDrivesEndpointDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAliasModel

`func (o *CapabilityDrivesEndpointDescriptor) GetAliasModel() string`

GetAliasModel returns the AliasModel field if non-nil, zero value otherwise.

### GetAliasModelOk

`func (o *CapabilityDrivesEndpointDescriptor) GetAliasModelOk() (*string, bool)`

GetAliasModelOk returns a tuple with the AliasModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasModel

`func (o *CapabilityDrivesEndpointDescriptor) SetAliasModel(v string)`

SetAliasModel sets AliasModel field to given value.

### HasAliasModel

`func (o *CapabilityDrivesEndpointDescriptor) HasAliasModel() bool`

HasAliasModel returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityDrivesEndpointDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityDrivesEndpointDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityDrivesEndpointDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityDrivesEndpointDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityDrivesEndpointDescriptor) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityDrivesEndpointDescriptor) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityDrivesEndpointDescriptor) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityDrivesEndpointDescriptor) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPartNumber

`func (o *CapabilityDrivesEndpointDescriptor) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CapabilityDrivesEndpointDescriptor) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CapabilityDrivesEndpointDescriptor) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *CapabilityDrivesEndpointDescriptor) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityDrivesEndpointDescriptor) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityDrivesEndpointDescriptor) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityDrivesEndpointDescriptor) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityDrivesEndpointDescriptor) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetSupportedPlatformsPids

`func (o *CapabilityDrivesEndpointDescriptor) GetSupportedPlatformsPids() []string`

GetSupportedPlatformsPids returns the SupportedPlatformsPids field if non-nil, zero value otherwise.

### GetSupportedPlatformsPidsOk

`func (o *CapabilityDrivesEndpointDescriptor) GetSupportedPlatformsPidsOk() (*[]string, bool)`

GetSupportedPlatformsPidsOk returns a tuple with the SupportedPlatformsPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatformsPids

`func (o *CapabilityDrivesEndpointDescriptor) SetSupportedPlatformsPids(v []string)`

SetSupportedPlatformsPids sets SupportedPlatformsPids field to given value.

### HasSupportedPlatformsPids

`func (o *CapabilityDrivesEndpointDescriptor) HasSupportedPlatformsPids() bool`

HasSupportedPlatformsPids returns a boolean if a field has been set.

### SetSupportedPlatformsPidsNil

`func (o *CapabilityDrivesEndpointDescriptor) SetSupportedPlatformsPidsNil(b bool)`

 SetSupportedPlatformsPidsNil sets the value for SupportedPlatformsPids to be an explicit nil

### UnsetSupportedPlatformsPids
`func (o *CapabilityDrivesEndpointDescriptor) UnsetSupportedPlatformsPids()`

UnsetSupportedPlatformsPids ensures that no value is present for SupportedPlatformsPids, not even an explicit nil
### GetVendor

`func (o *CapabilityDrivesEndpointDescriptor) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CapabilityDrivesEndpointDescriptor) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CapabilityDrivesEndpointDescriptor) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CapabilityDrivesEndpointDescriptor) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


