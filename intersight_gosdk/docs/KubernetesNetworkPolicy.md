# KubernetesNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NetworkPolicy"]
**CniConfig** | Pointer to [**NullableKubernetesCniConfig**](kubernetes.CniConfig.md) |  | [optional] 
**CniType** | Pointer to **string** | Supported CNI type. Currently we only support Calico. * &#x60;Calico&#x60; - Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin. * &#x60;Aci&#x60; - Cisco ACI Container Network Interface plugin. | [optional] [default to "Calico"]
**PodNetworkCidr** | Pointer to **string** | CIDR block to allocate Pod network IP addresses from. | [optional] 
**ServiceCidr** | Pointer to **string** | CIDR block to allocate cluster service IP addresses from. | [optional] 
**ClusterProfiles** | Pointer to [**[]KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) | An array of relationships to kubernetesClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewKubernetesNetworkPolicy

`func NewKubernetesNetworkPolicy(classId string, objectType string, ) *KubernetesNetworkPolicy`

NewKubernetesNetworkPolicy instantiates a new KubernetesNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNetworkPolicyWithDefaults

`func NewKubernetesNetworkPolicyWithDefaults() *KubernetesNetworkPolicy`

NewKubernetesNetworkPolicyWithDefaults instantiates a new KubernetesNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNetworkPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNetworkPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNetworkPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNetworkPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNetworkPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNetworkPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCniConfig

`func (o *KubernetesNetworkPolicy) GetCniConfig() KubernetesCniConfig`

GetCniConfig returns the CniConfig field if non-nil, zero value otherwise.

### GetCniConfigOk

`func (o *KubernetesNetworkPolicy) GetCniConfigOk() (*KubernetesCniConfig, bool)`

GetCniConfigOk returns a tuple with the CniConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCniConfig

`func (o *KubernetesNetworkPolicy) SetCniConfig(v KubernetesCniConfig)`

SetCniConfig sets CniConfig field to given value.

### HasCniConfig

`func (o *KubernetesNetworkPolicy) HasCniConfig() bool`

HasCniConfig returns a boolean if a field has been set.

### SetCniConfigNil

`func (o *KubernetesNetworkPolicy) SetCniConfigNil(b bool)`

 SetCniConfigNil sets the value for CniConfig to be an explicit nil

### UnsetCniConfig
`func (o *KubernetesNetworkPolicy) UnsetCniConfig()`

UnsetCniConfig ensures that no value is present for CniConfig, not even an explicit nil
### GetCniType

`func (o *KubernetesNetworkPolicy) GetCniType() string`

GetCniType returns the CniType field if non-nil, zero value otherwise.

### GetCniTypeOk

`func (o *KubernetesNetworkPolicy) GetCniTypeOk() (*string, bool)`

GetCniTypeOk returns a tuple with the CniType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCniType

`func (o *KubernetesNetworkPolicy) SetCniType(v string)`

SetCniType sets CniType field to given value.

### HasCniType

`func (o *KubernetesNetworkPolicy) HasCniType() bool`

HasCniType returns a boolean if a field has been set.

### GetPodNetworkCidr

`func (o *KubernetesNetworkPolicy) GetPodNetworkCidr() string`

GetPodNetworkCidr returns the PodNetworkCidr field if non-nil, zero value otherwise.

### GetPodNetworkCidrOk

`func (o *KubernetesNetworkPolicy) GetPodNetworkCidrOk() (*string, bool)`

GetPodNetworkCidrOk returns a tuple with the PodNetworkCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodNetworkCidr

`func (o *KubernetesNetworkPolicy) SetPodNetworkCidr(v string)`

SetPodNetworkCidr sets PodNetworkCidr field to given value.

### HasPodNetworkCidr

`func (o *KubernetesNetworkPolicy) HasPodNetworkCidr() bool`

HasPodNetworkCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *KubernetesNetworkPolicy) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *KubernetesNetworkPolicy) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *KubernetesNetworkPolicy) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *KubernetesNetworkPolicy) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *KubernetesNetworkPolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesNetworkPolicy) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesNetworkPolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesNetworkPolicy) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesNetworkPolicy) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesNetworkPolicy) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesNetworkPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


