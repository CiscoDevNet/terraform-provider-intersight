# CertificatemanagementPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "certificatemanagement.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "certificatemanagement.Policy"]
**Certificates** | Pointer to [**[]CertificatemanagementCertificateBase**](CertificatemanagementCertificateBase.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewCertificatemanagementPolicy

`func NewCertificatemanagementPolicy(classId string, objectType string, ) *CertificatemanagementPolicy`

NewCertificatemanagementPolicy instantiates a new CertificatemanagementPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatemanagementPolicyWithDefaults

`func NewCertificatemanagementPolicyWithDefaults() *CertificatemanagementPolicy`

NewCertificatemanagementPolicyWithDefaults instantiates a new CertificatemanagementPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CertificatemanagementPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CertificatemanagementPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CertificatemanagementPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CertificatemanagementPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificatemanagementPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificatemanagementPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertificates

`func (o *CertificatemanagementPolicy) GetCertificates() []CertificatemanagementCertificateBase`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *CertificatemanagementPolicy) GetCertificatesOk() (*[]CertificatemanagementCertificateBase, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *CertificatemanagementPolicy) SetCertificates(v []CertificatemanagementCertificateBase)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *CertificatemanagementPolicy) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *CertificatemanagementPolicy) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *CertificatemanagementPolicy) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetOrganization

`func (o *CertificatemanagementPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificatemanagementPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificatemanagementPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificatemanagementPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *CertificatemanagementPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *CertificatemanagementPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *CertificatemanagementPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *CertificatemanagementPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *CertificatemanagementPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *CertificatemanagementPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


