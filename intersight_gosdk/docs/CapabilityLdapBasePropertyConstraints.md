# CapabilityLdapBasePropertyConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.LdapBasePropertyConstraints"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.LdapBasePropertyConstraints"]
**MaxLdapGroupsSupported** | Pointer to **int64** | Maximum LDAP groups supported in FI. | [optional] [readonly] 
**MaxLdapProvidersSupported** | Pointer to **int64** | Maximum LDAP providers supported in FI. | [optional] [readonly] 
**RolesSupported** | Pointer to **[]string** |  | [optional] 
**SupportedMaxLengthForDn** | Pointer to **int64** | Maximum length supported for distinguished name. | [optional] [readonly] 
**SupportedMaxLengthForFilter** | Pointer to **int64** | Maximum length supported for ldap search filter. | [optional] [readonly] 

## Methods

### NewCapabilityLdapBasePropertyConstraints

`func NewCapabilityLdapBasePropertyConstraints(classId string, objectType string, ) *CapabilityLdapBasePropertyConstraints`

NewCapabilityLdapBasePropertyConstraints instantiates a new CapabilityLdapBasePropertyConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityLdapBasePropertyConstraintsWithDefaults

`func NewCapabilityLdapBasePropertyConstraintsWithDefaults() *CapabilityLdapBasePropertyConstraints`

NewCapabilityLdapBasePropertyConstraintsWithDefaults instantiates a new CapabilityLdapBasePropertyConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityLdapBasePropertyConstraints) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityLdapBasePropertyConstraints) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityLdapBasePropertyConstraints) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityLdapBasePropertyConstraints) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityLdapBasePropertyConstraints) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityLdapBasePropertyConstraints) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaxLdapGroupsSupported

`func (o *CapabilityLdapBasePropertyConstraints) GetMaxLdapGroupsSupported() int64`

GetMaxLdapGroupsSupported returns the MaxLdapGroupsSupported field if non-nil, zero value otherwise.

### GetMaxLdapGroupsSupportedOk

`func (o *CapabilityLdapBasePropertyConstraints) GetMaxLdapGroupsSupportedOk() (*int64, bool)`

GetMaxLdapGroupsSupportedOk returns a tuple with the MaxLdapGroupsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLdapGroupsSupported

`func (o *CapabilityLdapBasePropertyConstraints) SetMaxLdapGroupsSupported(v int64)`

SetMaxLdapGroupsSupported sets MaxLdapGroupsSupported field to given value.

### HasMaxLdapGroupsSupported

`func (o *CapabilityLdapBasePropertyConstraints) HasMaxLdapGroupsSupported() bool`

HasMaxLdapGroupsSupported returns a boolean if a field has been set.

### GetMaxLdapProvidersSupported

`func (o *CapabilityLdapBasePropertyConstraints) GetMaxLdapProvidersSupported() int64`

GetMaxLdapProvidersSupported returns the MaxLdapProvidersSupported field if non-nil, zero value otherwise.

### GetMaxLdapProvidersSupportedOk

`func (o *CapabilityLdapBasePropertyConstraints) GetMaxLdapProvidersSupportedOk() (*int64, bool)`

GetMaxLdapProvidersSupportedOk returns a tuple with the MaxLdapProvidersSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLdapProvidersSupported

`func (o *CapabilityLdapBasePropertyConstraints) SetMaxLdapProvidersSupported(v int64)`

SetMaxLdapProvidersSupported sets MaxLdapProvidersSupported field to given value.

### HasMaxLdapProvidersSupported

`func (o *CapabilityLdapBasePropertyConstraints) HasMaxLdapProvidersSupported() bool`

HasMaxLdapProvidersSupported returns a boolean if a field has been set.

### GetRolesSupported

`func (o *CapabilityLdapBasePropertyConstraints) GetRolesSupported() []string`

GetRolesSupported returns the RolesSupported field if non-nil, zero value otherwise.

### GetRolesSupportedOk

`func (o *CapabilityLdapBasePropertyConstraints) GetRolesSupportedOk() (*[]string, bool)`

GetRolesSupportedOk returns a tuple with the RolesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesSupported

`func (o *CapabilityLdapBasePropertyConstraints) SetRolesSupported(v []string)`

SetRolesSupported sets RolesSupported field to given value.

### HasRolesSupported

`func (o *CapabilityLdapBasePropertyConstraints) HasRolesSupported() bool`

HasRolesSupported returns a boolean if a field has been set.

### SetRolesSupportedNil

`func (o *CapabilityLdapBasePropertyConstraints) SetRolesSupportedNil(b bool)`

 SetRolesSupportedNil sets the value for RolesSupported to be an explicit nil

### UnsetRolesSupported
`func (o *CapabilityLdapBasePropertyConstraints) UnsetRolesSupported()`

UnsetRolesSupported ensures that no value is present for RolesSupported, not even an explicit nil
### GetSupportedMaxLengthForDn

`func (o *CapabilityLdapBasePropertyConstraints) GetSupportedMaxLengthForDn() int64`

GetSupportedMaxLengthForDn returns the SupportedMaxLengthForDn field if non-nil, zero value otherwise.

### GetSupportedMaxLengthForDnOk

`func (o *CapabilityLdapBasePropertyConstraints) GetSupportedMaxLengthForDnOk() (*int64, bool)`

GetSupportedMaxLengthForDnOk returns a tuple with the SupportedMaxLengthForDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMaxLengthForDn

`func (o *CapabilityLdapBasePropertyConstraints) SetSupportedMaxLengthForDn(v int64)`

SetSupportedMaxLengthForDn sets SupportedMaxLengthForDn field to given value.

### HasSupportedMaxLengthForDn

`func (o *CapabilityLdapBasePropertyConstraints) HasSupportedMaxLengthForDn() bool`

HasSupportedMaxLengthForDn returns a boolean if a field has been set.

### GetSupportedMaxLengthForFilter

`func (o *CapabilityLdapBasePropertyConstraints) GetSupportedMaxLengthForFilter() int64`

GetSupportedMaxLengthForFilter returns the SupportedMaxLengthForFilter field if non-nil, zero value otherwise.

### GetSupportedMaxLengthForFilterOk

`func (o *CapabilityLdapBasePropertyConstraints) GetSupportedMaxLengthForFilterOk() (*int64, bool)`

GetSupportedMaxLengthForFilterOk returns a tuple with the SupportedMaxLengthForFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMaxLengthForFilter

`func (o *CapabilityLdapBasePropertyConstraints) SetSupportedMaxLengthForFilter(v int64)`

SetSupportedMaxLengthForFilter sets SupportedMaxLengthForFilter field to given value.

### HasSupportedMaxLengthForFilter

`func (o *CapabilityLdapBasePropertyConstraints) HasSupportedMaxLengthForFilter() bool`

HasSupportedMaxLengthForFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


