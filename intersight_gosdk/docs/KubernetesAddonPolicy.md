# KubernetesAddonPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonPolicy"]
**SystemManaged** | Pointer to **bool** | To determine if Addon Policy is automatically managed by the system. | [optional] 
**Addons** | Pointer to [**[]KubernetesAddonRelationship**](KubernetesAddonRelationship.md) | An array of relationships to kubernetesAddon resources. | [optional] 
**ClusterProfiles** | Pointer to [**[]KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) | An array of relationships to kubernetesClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewKubernetesAddonPolicy

`func NewKubernetesAddonPolicy(classId string, objectType string, ) *KubernetesAddonPolicy`

NewKubernetesAddonPolicy instantiates a new KubernetesAddonPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonPolicyWithDefaults

`func NewKubernetesAddonPolicyWithDefaults() *KubernetesAddonPolicy`

NewKubernetesAddonPolicyWithDefaults instantiates a new KubernetesAddonPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSystemManaged

`func (o *KubernetesAddonPolicy) GetSystemManaged() bool`

GetSystemManaged returns the SystemManaged field if non-nil, zero value otherwise.

### GetSystemManagedOk

`func (o *KubernetesAddonPolicy) GetSystemManagedOk() (*bool, bool)`

GetSystemManagedOk returns a tuple with the SystemManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManaged

`func (o *KubernetesAddonPolicy) SetSystemManaged(v bool)`

SetSystemManaged sets SystemManaged field to given value.

### HasSystemManaged

`func (o *KubernetesAddonPolicy) HasSystemManaged() bool`

HasSystemManaged returns a boolean if a field has been set.

### GetAddons

`func (o *KubernetesAddonPolicy) GetAddons() []KubernetesAddonRelationship`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *KubernetesAddonPolicy) GetAddonsOk() (*[]KubernetesAddonRelationship, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *KubernetesAddonPolicy) SetAddons(v []KubernetesAddonRelationship)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *KubernetesAddonPolicy) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### SetAddonsNil

`func (o *KubernetesAddonPolicy) SetAddonsNil(b bool)`

 SetAddonsNil sets the value for Addons to be an explicit nil

### UnsetAddons
`func (o *KubernetesAddonPolicy) UnsetAddons()`

UnsetAddons ensures that no value is present for Addons, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesAddonPolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesAddonPolicy) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesAddonPolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesAddonPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesAddonPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesAddonPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesAddonPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


