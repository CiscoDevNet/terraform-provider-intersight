# VmediaPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | State of the Virtual Media service on the endpoint. | [optional] 
**Encryption** | Pointer to **bool** | If enabled, allows encryption of all Virtual Media communications. | [optional] 
**LowPowerUsb** | Pointer to **bool** | If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. | [optional] 
**Mappings** | Pointer to [**[]VmediaMapping**](vmedia.Mapping.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewVmediaPolicyAllOf

`func NewVmediaPolicyAllOf() *VmediaPolicyAllOf`

NewVmediaPolicyAllOf instantiates a new VmediaPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmediaPolicyAllOfWithDefaults

`func NewVmediaPolicyAllOfWithDefaults() *VmediaPolicyAllOf`

NewVmediaPolicyAllOfWithDefaults instantiates a new VmediaPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VmediaPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VmediaPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VmediaPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VmediaPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryption

`func (o *VmediaPolicyAllOf) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VmediaPolicyAllOf) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VmediaPolicyAllOf) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VmediaPolicyAllOf) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLowPowerUsb

`func (o *VmediaPolicyAllOf) GetLowPowerUsb() bool`

GetLowPowerUsb returns the LowPowerUsb field if non-nil, zero value otherwise.

### GetLowPowerUsbOk

`func (o *VmediaPolicyAllOf) GetLowPowerUsbOk() (*bool, bool)`

GetLowPowerUsbOk returns a tuple with the LowPowerUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPowerUsb

`func (o *VmediaPolicyAllOf) SetLowPowerUsb(v bool)`

SetLowPowerUsb sets LowPowerUsb field to given value.

### HasLowPowerUsb

`func (o *VmediaPolicyAllOf) HasLowPowerUsb() bool`

HasLowPowerUsb returns a boolean if a field has been set.

### GetMappings

`func (o *VmediaPolicyAllOf) GetMappings() []VmediaMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *VmediaPolicyAllOf) GetMappingsOk() (*[]VmediaMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *VmediaPolicyAllOf) SetMappings(v []VmediaMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *VmediaPolicyAllOf) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetOrganization

`func (o *VmediaPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VmediaPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VmediaPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VmediaPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *VmediaPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *VmediaPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *VmediaPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *VmediaPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *VmediaPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *VmediaPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


