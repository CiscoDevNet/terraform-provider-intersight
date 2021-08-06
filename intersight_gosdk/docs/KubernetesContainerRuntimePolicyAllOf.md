# KubernetesContainerRuntimePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ContainerRuntimePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ContainerRuntimePolicy"]
**DockerBridgeNetworkCidr** | Pointer to **string** | Bridge IP (--bip) including Prefix (e.g., 172.17.0.5/24) that Docker will use for the default bridge network (docker0). Containers will connect to this if no other network is configured, not used by kubernetes pods because their network is managed by CNI. However this address space must not collide with other CIDRs on your networks, including the cluster&#39;s Service CIDR, Pod Network CIDR and IP Pools. | [optional] 
**DockerHttpProxy** | Pointer to [**NullableKubernetesProxyConfig**](kubernetes.ProxyConfig.md) |  | [optional] 
**DockerHttpsProxy** | Pointer to [**NullableKubernetesProxyConfig**](kubernetes.ProxyConfig.md) |  | [optional] 
**DockerNoProxy** | Pointer to **[]string** |  | [optional] 
**ClusterProfiles** | Pointer to [**[]KubernetesClusterProfileRelationship**](KubernetesClusterProfileRelationship.md) | An array of relationships to kubernetesClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewKubernetesContainerRuntimePolicyAllOf

`func NewKubernetesContainerRuntimePolicyAllOf(classId string, objectType string, ) *KubernetesContainerRuntimePolicyAllOf`

NewKubernetesContainerRuntimePolicyAllOf instantiates a new KubernetesContainerRuntimePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesContainerRuntimePolicyAllOfWithDefaults

`func NewKubernetesContainerRuntimePolicyAllOfWithDefaults() *KubernetesContainerRuntimePolicyAllOf`

NewKubernetesContainerRuntimePolicyAllOfWithDefaults instantiates a new KubernetesContainerRuntimePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesContainerRuntimePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesContainerRuntimePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesContainerRuntimePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesContainerRuntimePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDockerBridgeNetworkCidr

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerBridgeNetworkCidr() string`

GetDockerBridgeNetworkCidr returns the DockerBridgeNetworkCidr field if non-nil, zero value otherwise.

### GetDockerBridgeNetworkCidrOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerBridgeNetworkCidrOk() (*string, bool)`

GetDockerBridgeNetworkCidrOk returns a tuple with the DockerBridgeNetworkCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerBridgeNetworkCidr

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerBridgeNetworkCidr(v string)`

SetDockerBridgeNetworkCidr sets DockerBridgeNetworkCidr field to given value.

### HasDockerBridgeNetworkCidr

`func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerBridgeNetworkCidr() bool`

HasDockerBridgeNetworkCidr returns a boolean if a field has been set.

### GetDockerHttpProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpProxy() KubernetesProxyConfig`

GetDockerHttpProxy returns the DockerHttpProxy field if non-nil, zero value otherwise.

### GetDockerHttpProxyOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpProxyOk() (*KubernetesProxyConfig, bool)`

GetDockerHttpProxyOk returns a tuple with the DockerHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHttpProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpProxy(v KubernetesProxyConfig)`

SetDockerHttpProxy sets DockerHttpProxy field to given value.

### HasDockerHttpProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerHttpProxy() bool`

HasDockerHttpProxy returns a boolean if a field has been set.

### SetDockerHttpProxyNil

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpProxyNil(b bool)`

 SetDockerHttpProxyNil sets the value for DockerHttpProxy to be an explicit nil

### UnsetDockerHttpProxy
`func (o *KubernetesContainerRuntimePolicyAllOf) UnsetDockerHttpProxy()`

UnsetDockerHttpProxy ensures that no value is present for DockerHttpProxy, not even an explicit nil
### GetDockerHttpsProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpsProxy() KubernetesProxyConfig`

GetDockerHttpsProxy returns the DockerHttpsProxy field if non-nil, zero value otherwise.

### GetDockerHttpsProxyOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerHttpsProxyOk() (*KubernetesProxyConfig, bool)`

GetDockerHttpsProxyOk returns a tuple with the DockerHttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHttpsProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpsProxy(v KubernetesProxyConfig)`

SetDockerHttpsProxy sets DockerHttpsProxy field to given value.

### HasDockerHttpsProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerHttpsProxy() bool`

HasDockerHttpsProxy returns a boolean if a field has been set.

### SetDockerHttpsProxyNil

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerHttpsProxyNil(b bool)`

 SetDockerHttpsProxyNil sets the value for DockerHttpsProxy to be an explicit nil

### UnsetDockerHttpsProxy
`func (o *KubernetesContainerRuntimePolicyAllOf) UnsetDockerHttpsProxy()`

UnsetDockerHttpsProxy ensures that no value is present for DockerHttpsProxy, not even an explicit nil
### GetDockerNoProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerNoProxy() []string`

GetDockerNoProxy returns the DockerNoProxy field if non-nil, zero value otherwise.

### GetDockerNoProxyOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetDockerNoProxyOk() (*[]string, bool)`

GetDockerNoProxyOk returns a tuple with the DockerNoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerNoProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerNoProxy(v []string)`

SetDockerNoProxy sets DockerNoProxy field to given value.

### HasDockerNoProxy

`func (o *KubernetesContainerRuntimePolicyAllOf) HasDockerNoProxy() bool`

HasDockerNoProxy returns a boolean if a field has been set.

### SetDockerNoProxyNil

`func (o *KubernetesContainerRuntimePolicyAllOf) SetDockerNoProxyNil(b bool)`

 SetDockerNoProxyNil sets the value for DockerNoProxy to be an explicit nil

### UnsetDockerNoProxy
`func (o *KubernetesContainerRuntimePolicyAllOf) UnsetDockerNoProxy()`

UnsetDockerNoProxy ensures that no value is present for DockerNoProxy, not even an explicit nil
### GetClusterProfiles

`func (o *KubernetesContainerRuntimePolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *KubernetesContainerRuntimePolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *KubernetesContainerRuntimePolicyAllOf) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *KubernetesContainerRuntimePolicyAllOf) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *KubernetesContainerRuntimePolicyAllOf) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *KubernetesContainerRuntimePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesContainerRuntimePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesContainerRuntimePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesContainerRuntimePolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


