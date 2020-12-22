# KubernetesAddonPolicyAllOf

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

### NewKubernetesAddonPolicyAllOf

`func NewKubernetesAddonPolicyAllOf(classId string, objectType string, ) *KubernetesAddonPolicyAllOf`

NewKubernetesAddonPolicyAllOf instantiates a new KubernetesAddonPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonPolicyAllOfWithDefaults

`func NewKubernetesAddonPolicyAllOfWithDefaults() *KubernetesAddonPolicyAllOf`

NewKubernetesAddonPolicyAllOfWithDefaults instantiates a new KubernetesAddonPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSystemManaged

`func (o *KubernetesAddonPolicyAllOf) GetSystemManaged() bool`

GetSystemManaged returns the SystemManaged field if non-nil, zero value otherwise.

### GetSystemManagedOk

`func (o *KubernetesAddonPolicyAllOf) GetSystemManagedOk() (*bool, bool)`

GetSystemManagedOk returns a tuple with the SystemManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManaged

`func (o *KubernetesAddonPolicyAllOf) SetSystemManaged(v bool)`

SetSystemManaged sets SystemManaged field to given value.

### HasSystemManaged

`func (o *KubernetesAddonPolicyAllOf) HasSystemManaged() bool`

HasSystemManaged returns a boolean if a field has been set.

### GetAddons

`func (o *KubernetesAddonPolicyAllOf) GetAddons() []KubernetesAddonRelationship`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *KubernetesAddonPolicyAllOf) GetAddonsOk() (*[]KubernetesAddonRelationship, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *KubernetesAddonPolicyAllOf) SetAddons(v []KubernetesAddonRelationship)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *KubernetesAddonPolicyAllOf) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### SetAddonsNil

`func (o *KubernetesAddonPolicyAllOf) SetAddonsNil(b bool)`

 SetAddonsNil sets the value for Addons to be an explicit nil

### UnsetAddons
`func (o *KubernetesAddonPolicyAllOf) UnsetAddons()`

UnsetAddons ensures that no value is present for Addons, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesAddonPolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesAddonPolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesAddonPolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesAddonPolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesAddonPolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesAddonPolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesAddonPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


