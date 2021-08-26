# KubernetesAciCniProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AciCniProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AciCniProfile"]
**AaepName** | Pointer to **string** | Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy. | [optional] 
**ExtSvcDynSubnetStart** | Pointer to **string** | Start of range of IP subnets for external services with dynamic IP allocation for use by Kubernetes clusters using this ACI CNI policy. | [optional] 
**ExtSvcStaticSubnetStart** | Pointer to **string** | Start of range of IP subnets for external services with static IP allocation for use by Kubernetes clusters using this ACI CNI policy. | [optional] 
**InfraVlanId** | Pointer to **int64** | Value of ACI infrastructuere VLAN ID for the ACI fabric. | [optional] [readonly] 
**L3OutNetworkName** | Pointer to **string** | Name of ACI L3Out network to be used for all Kubernetes clusters using this policy. | [optional] 
**L3OutPolicyName** | Pointer to **string** | Name of ACI L3Out policy to be used for all Kubernetes clusters using this policy. | [optional] 
**L3OutTenant** | Pointer to **string** | Tenant in ACI used by this L3Out and Common VRF. | [optional] 
**NestedVmmDomain** | Pointer to **string** | VMM domain within which Kubernetes clusters using this policy are nested. | [optional] 
**NodeSvcSubnetStart** | Pointer to **string** | Start of range of ACI Node Service IP subnets to use by Kubernetes clusters using this ACI CNI policy This is used for the service graph which is used for ACI PBR based load balancing. | [optional] 
**NodeVlanRangeEnd** | Pointer to **int64** | Ending value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy. | [optional] 
**NodeVlanRangeStart** | Pointer to **int64** | Starting value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy. | [optional] 
**NumberOfKubernetesClusters** | Pointer to **int64** | Number of k8s clusters currently using this ACI CNI profile. | [optional] [readonly] 
**OpflexMulticastAddressRange** | Pointer to **string** | Range of IP Multicast addresses to be used by the Opflex protocol for Kubernetes clusters using this policy. | [optional] 
**PodSubnetStart** | Pointer to **string** | Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile. | [optional] 
**SvcSubnetStart** | Pointer to **string** | Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only. | [optional] [readonly] 
**Vrf** | Pointer to **string** | VRF (Virtual Routing and Forwarding) domain to be used within ACI fabric by all k8s clusters using this policy. | [optional] 
**ClusterAciAllocations** | Pointer to [**[]KubernetesAciCniTenantClusterAllocationRelationship**](KubernetesAciCniTenantClusterAllocationRelationship.md) | An array of relationships to kubernetesAciCniTenantClusterAllocation resources. | [optional] [readonly] 
**ClusterProfiles** | Pointer to [**[]KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) | An array of relationships to kubernetesClusterProfile resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAciCniProfileAllOf

`func NewKubernetesAciCniProfileAllOf(classId string, objectType string, ) *KubernetesAciCniProfileAllOf`

NewKubernetesAciCniProfileAllOf instantiates a new KubernetesAciCniProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAciCniProfileAllOfWithDefaults

`func NewKubernetesAciCniProfileAllOfWithDefaults() *KubernetesAciCniProfileAllOf`

NewKubernetesAciCniProfileAllOfWithDefaults instantiates a new KubernetesAciCniProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAciCniProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAciCniProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAciCniProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAciCniProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAciCniProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAciCniProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAaepName

`func (o *KubernetesAciCniProfileAllOf) GetAaepName() string`

GetAaepName returns the AaepName field if non-nil, zero value otherwise.

### GetAaepNameOk

`func (o *KubernetesAciCniProfileAllOf) GetAaepNameOk() (*string, bool)`

GetAaepNameOk returns a tuple with the AaepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaepName

`func (o *KubernetesAciCniProfileAllOf) SetAaepName(v string)`

SetAaepName sets AaepName field to given value.

### HasAaepName

`func (o *KubernetesAciCniProfileAllOf) HasAaepName() bool`

HasAaepName returns a boolean if a field has been set.

### GetExtSvcDynSubnetStart

`func (o *KubernetesAciCniProfileAllOf) GetExtSvcDynSubnetStart() string`

GetExtSvcDynSubnetStart returns the ExtSvcDynSubnetStart field if non-nil, zero value otherwise.

### GetExtSvcDynSubnetStartOk

`func (o *KubernetesAciCniProfileAllOf) GetExtSvcDynSubnetStartOk() (*string, bool)`

GetExtSvcDynSubnetStartOk returns a tuple with the ExtSvcDynSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtSvcDynSubnetStart

`func (o *KubernetesAciCniProfileAllOf) SetExtSvcDynSubnetStart(v string)`

SetExtSvcDynSubnetStart sets ExtSvcDynSubnetStart field to given value.

### HasExtSvcDynSubnetStart

`func (o *KubernetesAciCniProfileAllOf) HasExtSvcDynSubnetStart() bool`

HasExtSvcDynSubnetStart returns a boolean if a field has been set.

### GetExtSvcStaticSubnetStart

`func (o *KubernetesAciCniProfileAllOf) GetExtSvcStaticSubnetStart() string`

GetExtSvcStaticSubnetStart returns the ExtSvcStaticSubnetStart field if non-nil, zero value otherwise.

### GetExtSvcStaticSubnetStartOk

`func (o *KubernetesAciCniProfileAllOf) GetExtSvcStaticSubnetStartOk() (*string, bool)`

GetExtSvcStaticSubnetStartOk returns a tuple with the ExtSvcStaticSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtSvcStaticSubnetStart

`func (o *KubernetesAciCniProfileAllOf) SetExtSvcStaticSubnetStart(v string)`

SetExtSvcStaticSubnetStart sets ExtSvcStaticSubnetStart field to given value.

### HasExtSvcStaticSubnetStart

`func (o *KubernetesAciCniProfileAllOf) HasExtSvcStaticSubnetStart() bool`

HasExtSvcStaticSubnetStart returns a boolean if a field has been set.

### GetInfraVlanId

`func (o *KubernetesAciCniProfileAllOf) GetInfraVlanId() int64`

GetInfraVlanId returns the InfraVlanId field if non-nil, zero value otherwise.

### GetInfraVlanIdOk

`func (o *KubernetesAciCniProfileAllOf) GetInfraVlanIdOk() (*int64, bool)`

GetInfraVlanIdOk returns a tuple with the InfraVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraVlanId

`func (o *KubernetesAciCniProfileAllOf) SetInfraVlanId(v int64)`

SetInfraVlanId sets InfraVlanId field to given value.

### HasInfraVlanId

`func (o *KubernetesAciCniProfileAllOf) HasInfraVlanId() bool`

HasInfraVlanId returns a boolean if a field has been set.

### GetL3OutNetworkName

`func (o *KubernetesAciCniProfileAllOf) GetL3OutNetworkName() string`

GetL3OutNetworkName returns the L3OutNetworkName field if non-nil, zero value otherwise.

### GetL3OutNetworkNameOk

`func (o *KubernetesAciCniProfileAllOf) GetL3OutNetworkNameOk() (*string, bool)`

GetL3OutNetworkNameOk returns a tuple with the L3OutNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3OutNetworkName

`func (o *KubernetesAciCniProfileAllOf) SetL3OutNetworkName(v string)`

SetL3OutNetworkName sets L3OutNetworkName field to given value.

### HasL3OutNetworkName

`func (o *KubernetesAciCniProfileAllOf) HasL3OutNetworkName() bool`

HasL3OutNetworkName returns a boolean if a field has been set.

### GetL3OutPolicyName

`func (o *KubernetesAciCniProfileAllOf) GetL3OutPolicyName() string`

GetL3OutPolicyName returns the L3OutPolicyName field if non-nil, zero value otherwise.

### GetL3OutPolicyNameOk

`func (o *KubernetesAciCniProfileAllOf) GetL3OutPolicyNameOk() (*string, bool)`

GetL3OutPolicyNameOk returns a tuple with the L3OutPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3OutPolicyName

`func (o *KubernetesAciCniProfileAllOf) SetL3OutPolicyName(v string)`

SetL3OutPolicyName sets L3OutPolicyName field to given value.

### HasL3OutPolicyName

`func (o *KubernetesAciCniProfileAllOf) HasL3OutPolicyName() bool`

HasL3OutPolicyName returns a boolean if a field has been set.

### GetL3OutTenant

`func (o *KubernetesAciCniProfileAllOf) GetL3OutTenant() string`

GetL3OutTenant returns the L3OutTenant field if non-nil, zero value otherwise.

### GetL3OutTenantOk

`func (o *KubernetesAciCniProfileAllOf) GetL3OutTenantOk() (*string, bool)`

GetL3OutTenantOk returns a tuple with the L3OutTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3OutTenant

`func (o *KubernetesAciCniProfileAllOf) SetL3OutTenant(v string)`

SetL3OutTenant sets L3OutTenant field to given value.

### HasL3OutTenant

`func (o *KubernetesAciCniProfileAllOf) HasL3OutTenant() bool`

HasL3OutTenant returns a boolean if a field has been set.

### GetNestedVmmDomain

`func (o *KubernetesAciCniProfileAllOf) GetNestedVmmDomain() string`

GetNestedVmmDomain returns the NestedVmmDomain field if non-nil, zero value otherwise.

### GetNestedVmmDomainOk

`func (o *KubernetesAciCniProfileAllOf) GetNestedVmmDomainOk() (*string, bool)`

GetNestedVmmDomainOk returns a tuple with the NestedVmmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVmmDomain

`func (o *KubernetesAciCniProfileAllOf) SetNestedVmmDomain(v string)`

SetNestedVmmDomain sets NestedVmmDomain field to given value.

### HasNestedVmmDomain

`func (o *KubernetesAciCniProfileAllOf) HasNestedVmmDomain() bool`

HasNestedVmmDomain returns a boolean if a field has been set.

### GetNodeSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) GetNodeSvcSubnetStart() string`

GetNodeSvcSubnetStart returns the NodeSvcSubnetStart field if non-nil, zero value otherwise.

### GetNodeSvcSubnetStartOk

`func (o *KubernetesAciCniProfileAllOf) GetNodeSvcSubnetStartOk() (*string, bool)`

GetNodeSvcSubnetStartOk returns a tuple with the NodeSvcSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) SetNodeSvcSubnetStart(v string)`

SetNodeSvcSubnetStart sets NodeSvcSubnetStart field to given value.

### HasNodeSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) HasNodeSvcSubnetStart() bool`

HasNodeSvcSubnetStart returns a boolean if a field has been set.

### GetNodeVlanRangeEnd

`func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeEnd() int64`

GetNodeVlanRangeEnd returns the NodeVlanRangeEnd field if non-nil, zero value otherwise.

### GetNodeVlanRangeEndOk

`func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeEndOk() (*int64, bool)`

GetNodeVlanRangeEndOk returns a tuple with the NodeVlanRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeVlanRangeEnd

`func (o *KubernetesAciCniProfileAllOf) SetNodeVlanRangeEnd(v int64)`

SetNodeVlanRangeEnd sets NodeVlanRangeEnd field to given value.

### HasNodeVlanRangeEnd

`func (o *KubernetesAciCniProfileAllOf) HasNodeVlanRangeEnd() bool`

HasNodeVlanRangeEnd returns a boolean if a field has been set.

### GetNodeVlanRangeStart

`func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeStart() int64`

GetNodeVlanRangeStart returns the NodeVlanRangeStart field if non-nil, zero value otherwise.

### GetNodeVlanRangeStartOk

`func (o *KubernetesAciCniProfileAllOf) GetNodeVlanRangeStartOk() (*int64, bool)`

GetNodeVlanRangeStartOk returns a tuple with the NodeVlanRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeVlanRangeStart

`func (o *KubernetesAciCniProfileAllOf) SetNodeVlanRangeStart(v int64)`

SetNodeVlanRangeStart sets NodeVlanRangeStart field to given value.

### HasNodeVlanRangeStart

`func (o *KubernetesAciCniProfileAllOf) HasNodeVlanRangeStart() bool`

HasNodeVlanRangeStart returns a boolean if a field has been set.

### GetNumberOfKubernetesClusters

`func (o *KubernetesAciCniProfileAllOf) GetNumberOfKubernetesClusters() int64`

GetNumberOfKubernetesClusters returns the NumberOfKubernetesClusters field if non-nil, zero value otherwise.

### GetNumberOfKubernetesClustersOk

`func (o *KubernetesAciCniProfileAllOf) GetNumberOfKubernetesClustersOk() (*int64, bool)`

GetNumberOfKubernetesClustersOk returns a tuple with the NumberOfKubernetesClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfKubernetesClusters

`func (o *KubernetesAciCniProfileAllOf) SetNumberOfKubernetesClusters(v int64)`

SetNumberOfKubernetesClusters sets NumberOfKubernetesClusters field to given value.

### HasNumberOfKubernetesClusters

`func (o *KubernetesAciCniProfileAllOf) HasNumberOfKubernetesClusters() bool`

HasNumberOfKubernetesClusters returns a boolean if a field has been set.

### GetOpflexMulticastAddressRange

`func (o *KubernetesAciCniProfileAllOf) GetOpflexMulticastAddressRange() string`

GetOpflexMulticastAddressRange returns the OpflexMulticastAddressRange field if non-nil, zero value otherwise.

### GetOpflexMulticastAddressRangeOk

`func (o *KubernetesAciCniProfileAllOf) GetOpflexMulticastAddressRangeOk() (*string, bool)`

GetOpflexMulticastAddressRangeOk returns a tuple with the OpflexMulticastAddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpflexMulticastAddressRange

`func (o *KubernetesAciCniProfileAllOf) SetOpflexMulticastAddressRange(v string)`

SetOpflexMulticastAddressRange sets OpflexMulticastAddressRange field to given value.

### HasOpflexMulticastAddressRange

`func (o *KubernetesAciCniProfileAllOf) HasOpflexMulticastAddressRange() bool`

HasOpflexMulticastAddressRange returns a boolean if a field has been set.

### GetPodSubnetStart

`func (o *KubernetesAciCniProfileAllOf) GetPodSubnetStart() string`

GetPodSubnetStart returns the PodSubnetStart field if non-nil, zero value otherwise.

### GetPodSubnetStartOk

`func (o *KubernetesAciCniProfileAllOf) GetPodSubnetStartOk() (*string, bool)`

GetPodSubnetStartOk returns a tuple with the PodSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSubnetStart

`func (o *KubernetesAciCniProfileAllOf) SetPodSubnetStart(v string)`

SetPodSubnetStart sets PodSubnetStart field to given value.

### HasPodSubnetStart

`func (o *KubernetesAciCniProfileAllOf) HasPodSubnetStart() bool`

HasPodSubnetStart returns a boolean if a field has been set.

### GetSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) GetSvcSubnetStart() string`

GetSvcSubnetStart returns the SvcSubnetStart field if non-nil, zero value otherwise.

### GetSvcSubnetStartOk

`func (o *KubernetesAciCniProfileAllOf) GetSvcSubnetStartOk() (*string, bool)`

GetSvcSubnetStartOk returns a tuple with the SvcSubnetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) SetSvcSubnetStart(v string)`

SetSvcSubnetStart sets SvcSubnetStart field to given value.

### HasSvcSubnetStart

`func (o *KubernetesAciCniProfileAllOf) HasSvcSubnetStart() bool`

HasSvcSubnetStart returns a boolean if a field has been set.

### GetVrf

`func (o *KubernetesAciCniProfileAllOf) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *KubernetesAciCniProfileAllOf) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *KubernetesAciCniProfileAllOf) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *KubernetesAciCniProfileAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetClusterAciAllocations

`func (o *KubernetesAciCniProfileAllOf) GetClusterAciAllocations() []KubernetesAciCniTenantClusterAllocationRelationship`

GetClusterAciAllocations returns the ClusterAciAllocations field if non-nil, zero value otherwise.

### GetClusterAciAllocationsOk

`func (o *KubernetesAciCniProfileAllOf) GetClusterAciAllocationsOk() (*[]KubernetesAciCniTenantClusterAllocationRelationship, bool)`

GetClusterAciAllocationsOk returns a tuple with the ClusterAciAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAciAllocations

`func (o *KubernetesAciCniProfileAllOf) SetClusterAciAllocations(v []KubernetesAciCniTenantClusterAllocationRelationship)`

SetClusterAciAllocations sets ClusterAciAllocations field to given value.

### HasClusterAciAllocations

`func (o *KubernetesAciCniProfileAllOf) HasClusterAciAllocations() bool`

HasClusterAciAllocations returns a boolean if a field has been set.

### SetClusterAciAllocationsNil

`func (o *KubernetesAciCniProfileAllOf) SetClusterAciAllocationsNil(b bool)`

 SetClusterAciAllocationsNil sets the value for ClusterAciAllocations to be an explicit nil

### UnsetClusterAciAllocations
`func (o *KubernetesAciCniProfileAllOf) UnsetClusterAciAllocations()`

UnsetClusterAciAllocations ensures that no value is present for ClusterAciAllocations, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesAciCniProfileAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesAciCniProfileAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesAciCniProfileAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesAciCniProfileAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesAciCniProfileAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesAciCniProfileAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesAciCniProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAciCniProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAciCniProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAciCniProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *KubernetesAciCniProfileAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesAciCniProfileAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesAciCniProfileAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesAciCniProfileAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


