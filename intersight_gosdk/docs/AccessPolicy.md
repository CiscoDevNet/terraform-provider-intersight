# AccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.Policy"]
**AddressType** | Pointer to [**NullableAccessAddressType**](AccessAddressType.md) |  | [optional] 
**ConfigurationType** | Pointer to [**NullableAccessConfigurationType**](AccessConfigurationType.md) |  | [optional] 
**InbandVlan** | Pointer to **int64** | VLAN to be used for server access over Inband network. When Inband is enabled, only numbers between 4 to 4093 are allowed. | [optional] [default to 0]
**InbandIpPool** | Pointer to [**NullableIppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**InbandVrf** | Pointer to [**NullableVrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**OutOfBandIpPool** | Pointer to [**NullableIppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**OutOfBandVrf** | Pointer to [**NullableVrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewAccessPolicy

`func NewAccessPolicy(classId string, objectType string, ) *AccessPolicy`

NewAccessPolicy instantiates a new AccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyWithDefaults

`func NewAccessPolicyWithDefaults() *AccessPolicy`

NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddressType

`func (o *AccessPolicy) GetAddressType() AccessAddressType`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AccessPolicy) GetAddressTypeOk() (*AccessAddressType, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AccessPolicy) SetAddressType(v AccessAddressType)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *AccessPolicy) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### SetAddressTypeNil

`func (o *AccessPolicy) SetAddressTypeNil(b bool)`

 SetAddressTypeNil sets the value for AddressType to be an explicit nil

### UnsetAddressType
`func (o *AccessPolicy) UnsetAddressType()`

UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
### GetConfigurationType

`func (o *AccessPolicy) GetConfigurationType() AccessConfigurationType`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *AccessPolicy) GetConfigurationTypeOk() (*AccessConfigurationType, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *AccessPolicy) SetConfigurationType(v AccessConfigurationType)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *AccessPolicy) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### SetConfigurationTypeNil

`func (o *AccessPolicy) SetConfigurationTypeNil(b bool)`

 SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil

### UnsetConfigurationType
`func (o *AccessPolicy) UnsetConfigurationType()`

UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
### GetInbandVlan

`func (o *AccessPolicy) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *AccessPolicy) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *AccessPolicy) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *AccessPolicy) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInbandIpPool

`func (o *AccessPolicy) GetInbandIpPool() IppoolPoolRelationship`

GetInbandIpPool returns the InbandIpPool field if non-nil, zero value otherwise.

### GetInbandIpPoolOk

`func (o *AccessPolicy) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInbandIpPoolOk returns a tuple with the InbandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpPool

`func (o *AccessPolicy) SetInbandIpPool(v IppoolPoolRelationship)`

SetInbandIpPool sets InbandIpPool field to given value.

### HasInbandIpPool

`func (o *AccessPolicy) HasInbandIpPool() bool`

HasInbandIpPool returns a boolean if a field has been set.

### SetInbandIpPoolNil

`func (o *AccessPolicy) SetInbandIpPoolNil(b bool)`

 SetInbandIpPoolNil sets the value for InbandIpPool to be an explicit nil

### UnsetInbandIpPool
`func (o *AccessPolicy) UnsetInbandIpPool()`

UnsetInbandIpPool ensures that no value is present for InbandIpPool, not even an explicit nil
### GetInbandVrf

`func (o *AccessPolicy) GetInbandVrf() VrfVrfRelationship`

GetInbandVrf returns the InbandVrf field if non-nil, zero value otherwise.

### GetInbandVrfOk

`func (o *AccessPolicy) GetInbandVrfOk() (*VrfVrfRelationship, bool)`

GetInbandVrfOk returns a tuple with the InbandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVrf

`func (o *AccessPolicy) SetInbandVrf(v VrfVrfRelationship)`

SetInbandVrf sets InbandVrf field to given value.

### HasInbandVrf

`func (o *AccessPolicy) HasInbandVrf() bool`

HasInbandVrf returns a boolean if a field has been set.

### SetInbandVrfNil

`func (o *AccessPolicy) SetInbandVrfNil(b bool)`

 SetInbandVrfNil sets the value for InbandVrf to be an explicit nil

### UnsetInbandVrf
`func (o *AccessPolicy) UnsetInbandVrf()`

UnsetInbandVrf ensures that no value is present for InbandVrf, not even an explicit nil
### GetOrganization

`func (o *AccessPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccessPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccessPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccessPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *AccessPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *AccessPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetOutOfBandIpPool

`func (o *AccessPolicy) GetOutOfBandIpPool() IppoolPoolRelationship`

GetOutOfBandIpPool returns the OutOfBandIpPool field if non-nil, zero value otherwise.

### GetOutOfBandIpPoolOk

`func (o *AccessPolicy) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpPool

`func (o *AccessPolicy) SetOutOfBandIpPool(v IppoolPoolRelationship)`

SetOutOfBandIpPool sets OutOfBandIpPool field to given value.

### HasOutOfBandIpPool

`func (o *AccessPolicy) HasOutOfBandIpPool() bool`

HasOutOfBandIpPool returns a boolean if a field has been set.

### SetOutOfBandIpPoolNil

`func (o *AccessPolicy) SetOutOfBandIpPoolNil(b bool)`

 SetOutOfBandIpPoolNil sets the value for OutOfBandIpPool to be an explicit nil

### UnsetOutOfBandIpPool
`func (o *AccessPolicy) UnsetOutOfBandIpPool()`

UnsetOutOfBandIpPool ensures that no value is present for OutOfBandIpPool, not even an explicit nil
### GetOutOfBandVrf

`func (o *AccessPolicy) GetOutOfBandVrf() VrfVrfRelationship`

GetOutOfBandVrf returns the OutOfBandVrf field if non-nil, zero value otherwise.

### GetOutOfBandVrfOk

`func (o *AccessPolicy) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool)`

GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandVrf

`func (o *AccessPolicy) SetOutOfBandVrf(v VrfVrfRelationship)`

SetOutOfBandVrf sets OutOfBandVrf field to given value.

### HasOutOfBandVrf

`func (o *AccessPolicy) HasOutOfBandVrf() bool`

HasOutOfBandVrf returns a boolean if a field has been set.

### SetOutOfBandVrfNil

`func (o *AccessPolicy) SetOutOfBandVrfNil(b bool)`

 SetOutOfBandVrfNil sets the value for OutOfBandVrf to be an explicit nil

### UnsetOutOfBandVrf
`func (o *AccessPolicy) UnsetOutOfBandVrf()`

UnsetOutOfBandVrf ensures that no value is present for OutOfBandVrf, not even an explicit nil
### GetProfiles

`func (o *AccessPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AccessPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AccessPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AccessPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *AccessPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *AccessPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


