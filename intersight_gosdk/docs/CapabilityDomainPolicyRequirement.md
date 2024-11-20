# CapabilityDomainPolicyRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.DomainPolicyRequirement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.DomainPolicyRequirement"]
**CertificatePolicyConstraints** | Pointer to [**NullableCapabilityCertificatePropertyConstraints**](CapabilityCertificatePropertyConstraints.md) |  | [optional] 
**LdapConstraints** | Pointer to [**NullableCapabilityLdapBasePropertyConstraints**](CapabilityLdapBasePropertyConstraints.md) |  | [optional] 
**MinBundleVersion** | Pointer to **string** | Minimum Bundle version from which it is supported. | [optional] [readonly] 
**MinVersion** | Pointer to **string** | Minimum Version from which policy is supported. | [optional] [readonly] 
**Model** | Pointer to **string** | Type of the platform for which version compatibility is specified. Example - 3GFI, 4GFI, etc. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Policy Name for which version compatibility is specified. Example - snmp.Policy, ldap.Policy. | [optional] [readonly] 

## Methods

### NewCapabilityDomainPolicyRequirement

`func NewCapabilityDomainPolicyRequirement(classId string, objectType string, ) *CapabilityDomainPolicyRequirement`

NewCapabilityDomainPolicyRequirement instantiates a new CapabilityDomainPolicyRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDomainPolicyRequirementWithDefaults

`func NewCapabilityDomainPolicyRequirementWithDefaults() *CapabilityDomainPolicyRequirement`

NewCapabilityDomainPolicyRequirementWithDefaults instantiates a new CapabilityDomainPolicyRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityDomainPolicyRequirement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityDomainPolicyRequirement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityDomainPolicyRequirement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityDomainPolicyRequirement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityDomainPolicyRequirement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityDomainPolicyRequirement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificatePolicyConstraints

`func (o *CapabilityDomainPolicyRequirement) GetCertificatePolicyConstraints() CapabilityCertificatePropertyConstraints`

GetCertificatePolicyConstraints returns the CertificatePolicyConstraints field if non-nil, zero value otherwise.

### GetCertificatePolicyConstraintsOk

`func (o *CapabilityDomainPolicyRequirement) GetCertificatePolicyConstraintsOk() (*CapabilityCertificatePropertyConstraints, bool)`

GetCertificatePolicyConstraintsOk returns a tuple with the CertificatePolicyConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicyConstraints

`func (o *CapabilityDomainPolicyRequirement) SetCertificatePolicyConstraints(v CapabilityCertificatePropertyConstraints)`

SetCertificatePolicyConstraints sets CertificatePolicyConstraints field to given value.

### HasCertificatePolicyConstraints

`func (o *CapabilityDomainPolicyRequirement) HasCertificatePolicyConstraints() bool`

HasCertificatePolicyConstraints returns a boolean if a field has been set.

### SetCertificatePolicyConstraintsNil

`func (o *CapabilityDomainPolicyRequirement) SetCertificatePolicyConstraintsNil(b bool)`

 SetCertificatePolicyConstraintsNil sets the value for CertificatePolicyConstraints to be an explicit nil

### UnsetCertificatePolicyConstraints
`func (o *CapabilityDomainPolicyRequirement) UnsetCertificatePolicyConstraints()`

UnsetCertificatePolicyConstraints ensures that no value is present for CertificatePolicyConstraints, not even an explicit nil
### GetLdapConstraints

`func (o *CapabilityDomainPolicyRequirement) GetLdapConstraints() CapabilityLdapBasePropertyConstraints`

GetLdapConstraints returns the LdapConstraints field if non-nil, zero value otherwise.

### GetLdapConstraintsOk

`func (o *CapabilityDomainPolicyRequirement) GetLdapConstraintsOk() (*CapabilityLdapBasePropertyConstraints, bool)`

GetLdapConstraintsOk returns a tuple with the LdapConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapConstraints

`func (o *CapabilityDomainPolicyRequirement) SetLdapConstraints(v CapabilityLdapBasePropertyConstraints)`

SetLdapConstraints sets LdapConstraints field to given value.

### HasLdapConstraints

`func (o *CapabilityDomainPolicyRequirement) HasLdapConstraints() bool`

HasLdapConstraints returns a boolean if a field has been set.

### SetLdapConstraintsNil

`func (o *CapabilityDomainPolicyRequirement) SetLdapConstraintsNil(b bool)`

 SetLdapConstraintsNil sets the value for LdapConstraints to be an explicit nil

### UnsetLdapConstraints
`func (o *CapabilityDomainPolicyRequirement) UnsetLdapConstraints()`

UnsetLdapConstraints ensures that no value is present for LdapConstraints, not even an explicit nil
### GetMinBundleVersion

`func (o *CapabilityDomainPolicyRequirement) GetMinBundleVersion() string`

GetMinBundleVersion returns the MinBundleVersion field if non-nil, zero value otherwise.

### GetMinBundleVersionOk

`func (o *CapabilityDomainPolicyRequirement) GetMinBundleVersionOk() (*string, bool)`

GetMinBundleVersionOk returns a tuple with the MinBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBundleVersion

`func (o *CapabilityDomainPolicyRequirement) SetMinBundleVersion(v string)`

SetMinBundleVersion sets MinBundleVersion field to given value.

### HasMinBundleVersion

`func (o *CapabilityDomainPolicyRequirement) HasMinBundleVersion() bool`

HasMinBundleVersion returns a boolean if a field has been set.

### GetMinVersion

`func (o *CapabilityDomainPolicyRequirement) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *CapabilityDomainPolicyRequirement) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *CapabilityDomainPolicyRequirement) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *CapabilityDomainPolicyRequirement) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetModel

`func (o *CapabilityDomainPolicyRequirement) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CapabilityDomainPolicyRequirement) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CapabilityDomainPolicyRequirement) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CapabilityDomainPolicyRequirement) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPolicyName

`func (o *CapabilityDomainPolicyRequirement) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CapabilityDomainPolicyRequirement) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *CapabilityDomainPolicyRequirement) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *CapabilityDomainPolicyRequirement) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


