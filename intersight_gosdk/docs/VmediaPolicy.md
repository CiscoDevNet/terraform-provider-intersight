# VmediaPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vmedia.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vmedia.Policy"]
**Enabled** | Pointer to **bool** | State of the Virtual Media service on the endpoint. | [optional] [default to true]
**Encryption** | Pointer to **bool** | If enabled, allows encryption of all Virtual Media communications. Please note that this is no longer applicable for servers running versions 4.2 and above. | [optional] [default to true]
**LowPowerUsb** | Pointer to **bool** | If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. | [optional] [default to true]
**Mappings** | Pointer to [**[]VmediaMapping**](VmediaMapping.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewVmediaPolicy

`func NewVmediaPolicy(classId string, objectType string, ) *VmediaPolicy`

NewVmediaPolicy instantiates a new VmediaPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmediaPolicyWithDefaults

`func NewVmediaPolicyWithDefaults() *VmediaPolicy`

NewVmediaPolicyWithDefaults instantiates a new VmediaPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VmediaPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VmediaPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VmediaPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VmediaPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VmediaPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VmediaPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *VmediaPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VmediaPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VmediaPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VmediaPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryption

`func (o *VmediaPolicy) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VmediaPolicy) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VmediaPolicy) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VmediaPolicy) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLowPowerUsb

`func (o *VmediaPolicy) GetLowPowerUsb() bool`

GetLowPowerUsb returns the LowPowerUsb field if non-nil, zero value otherwise.

### GetLowPowerUsbOk

`func (o *VmediaPolicy) GetLowPowerUsbOk() (*bool, bool)`

GetLowPowerUsbOk returns a tuple with the LowPowerUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPowerUsb

`func (o *VmediaPolicy) SetLowPowerUsb(v bool)`

SetLowPowerUsb sets LowPowerUsb field to given value.

### HasLowPowerUsb

`func (o *VmediaPolicy) HasLowPowerUsb() bool`

HasLowPowerUsb returns a boolean if a field has been set.

### GetMappings

`func (o *VmediaPolicy) GetMappings() []VmediaMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *VmediaPolicy) GetMappingsOk() (*[]VmediaMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *VmediaPolicy) SetMappings(v []VmediaMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *VmediaPolicy) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *VmediaPolicy) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *VmediaPolicy) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetOrganization

`func (o *VmediaPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VmediaPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VmediaPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VmediaPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *VmediaPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *VmediaPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *VmediaPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *VmediaPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *VmediaPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *VmediaPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


