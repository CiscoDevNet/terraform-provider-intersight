# KubernetesTrustedRegistriesPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.TrustedRegistriesPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.TrustedRegistriesPolicy"]
**RootCaRegistries** | Pointer to **[]string** |  | [optional] 
**UnsignedRegistries** | Pointer to **[]string** |  | [optional] 
**ClusterProfiles** | Pointer to [**[]KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) | An array of relationships to kubernetesClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesTrustedRegistriesPolicy

`func NewKubernetesTrustedRegistriesPolicy(classId string, objectType string, ) *KubernetesTrustedRegistriesPolicy`

NewKubernetesTrustedRegistriesPolicy instantiates a new KubernetesTrustedRegistriesPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTrustedRegistriesPolicyWithDefaults

`func NewKubernetesTrustedRegistriesPolicyWithDefaults() *KubernetesTrustedRegistriesPolicy`

NewKubernetesTrustedRegistriesPolicyWithDefaults instantiates a new KubernetesTrustedRegistriesPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesTrustedRegistriesPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesTrustedRegistriesPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesTrustedRegistriesPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesTrustedRegistriesPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesTrustedRegistriesPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesTrustedRegistriesPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicy) GetRootCaRegistries() []string`

GetRootCaRegistries returns the RootCaRegistries field if non-nil, zero value otherwise.

### GetRootCaRegistriesOk

`func (o *KubernetesTrustedRegistriesPolicy) GetRootCaRegistriesOk() (*[]string, bool)`

GetRootCaRegistriesOk returns a tuple with the RootCaRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicy) SetRootCaRegistries(v []string)`

SetRootCaRegistries sets RootCaRegistries field to given value.

### HasRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicy) HasRootCaRegistries() bool`

HasRootCaRegistries returns a boolean if a field has been set.

### SetRootCaRegistriesNil

`func (o *KubernetesTrustedRegistriesPolicy) SetRootCaRegistriesNil(b bool)`

 SetRootCaRegistriesNil sets the value for RootCaRegistries to be an explicit nil

### UnsetRootCaRegistries
`func (o *KubernetesTrustedRegistriesPolicy) UnsetRootCaRegistries()`

UnsetRootCaRegistries ensures that no value is present for RootCaRegistries, not even an explicit nil
### GetUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicy) GetUnsignedRegistries() []string`

GetUnsignedRegistries returns the UnsignedRegistries field if non-nil, zero value otherwise.

### GetUnsignedRegistriesOk

`func (o *KubernetesTrustedRegistriesPolicy) GetUnsignedRegistriesOk() (*[]string, bool)`

GetUnsignedRegistriesOk returns a tuple with the UnsignedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicy) SetUnsignedRegistries(v []string)`

SetUnsignedRegistries sets UnsignedRegistries field to given value.

### HasUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicy) HasUnsignedRegistries() bool`

HasUnsignedRegistries returns a boolean if a field has been set.

### SetUnsignedRegistriesNil

`func (o *KubernetesTrustedRegistriesPolicy) SetUnsignedRegistriesNil(b bool)`

 SetUnsignedRegistriesNil sets the value for UnsignedRegistries to be an explicit nil

### UnsetUnsignedRegistries
`func (o *KubernetesTrustedRegistriesPolicy) UnsetUnsignedRegistries()`

UnsetUnsignedRegistries ensures that no value is present for UnsignedRegistries, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesTrustedRegistriesPolicy) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesTrustedRegistriesPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesTrustedRegistriesPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesTrustedRegistriesPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesTrustedRegistriesPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesTrustedRegistriesPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesTrustedRegistriesPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


