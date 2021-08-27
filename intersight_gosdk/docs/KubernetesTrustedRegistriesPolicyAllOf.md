# KubernetesTrustedRegistriesPolicyAllOf

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

### NewKubernetesTrustedRegistriesPolicyAllOf

`func NewKubernetesTrustedRegistriesPolicyAllOf(classId string, objectType string, ) *KubernetesTrustedRegistriesPolicyAllOf`

NewKubernetesTrustedRegistriesPolicyAllOf instantiates a new KubernetesTrustedRegistriesPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTrustedRegistriesPolicyAllOfWithDefaults

`func NewKubernetesTrustedRegistriesPolicyAllOfWithDefaults() *KubernetesTrustedRegistriesPolicyAllOf`

NewKubernetesTrustedRegistriesPolicyAllOfWithDefaults instantiates a new KubernetesTrustedRegistriesPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetRootCaRegistries() []string`

GetRootCaRegistries returns the RootCaRegistries field if non-nil, zero value otherwise.

### GetRootCaRegistriesOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetRootCaRegistriesOk() (*[]string, bool)`

GetRootCaRegistriesOk returns a tuple with the RootCaRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetRootCaRegistries(v []string)`

SetRootCaRegistries sets RootCaRegistries field to given value.

### HasRootCaRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) HasRootCaRegistries() bool`

HasRootCaRegistries returns a boolean if a field has been set.

### SetRootCaRegistriesNil

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetRootCaRegistriesNil(b bool)`

 SetRootCaRegistriesNil sets the value for RootCaRegistries to be an explicit nil

### UnsetRootCaRegistries
`func (o *KubernetesTrustedRegistriesPolicyAllOf) UnsetRootCaRegistries()`

UnsetRootCaRegistries ensures that no value is present for RootCaRegistries, not even an explicit nil
### GetUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetUnsignedRegistries() []string`

GetUnsignedRegistries returns the UnsignedRegistries field if non-nil, zero value otherwise.

### GetUnsignedRegistriesOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetUnsignedRegistriesOk() (*[]string, bool)`

GetUnsignedRegistriesOk returns a tuple with the UnsignedRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetUnsignedRegistries(v []string)`

SetUnsignedRegistries sets UnsignedRegistries field to given value.

### HasUnsignedRegistries

`func (o *KubernetesTrustedRegistriesPolicyAllOf) HasUnsignedRegistries() bool`

HasUnsignedRegistries returns a boolean if a field has been set.

### SetUnsignedRegistriesNil

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetUnsignedRegistriesNil(b bool)`

 SetUnsignedRegistriesNil sets the value for UnsignedRegistries to be an explicit nil

### UnsetUnsignedRegistries
`func (o *KubernetesTrustedRegistriesPolicyAllOf) UnsetUnsignedRegistries()`

UnsetUnsignedRegistries ensures that no value is present for UnsignedRegistries, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesTrustedRegistriesPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesTrustedRegistriesPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesTrustedRegistriesPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesTrustedRegistriesPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesTrustedRegistriesPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


