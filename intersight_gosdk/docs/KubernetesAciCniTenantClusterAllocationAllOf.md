# KubernetesAciCniTenantClusterAllocationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AciCniTenantClusterAllocation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AciCniTenantClusterAllocation"]
**NodeSvcIpSubnet** | Pointer to **string** | CIDR allocated for ACI node service IPs in this tenant cluster. | [optional] [readonly] 
**PodIpSubnet** | Pointer to **string** | CIDR allocated for pod IPs in this tenant cluster. | [optional] [readonly] 
**VlanEnd** | Pointer to **string** | End of VLAN range allocated to this tenant cluster. | [optional] [readonly] 
**VlanStart** | Pointer to **string** | Start of VLAN range allocated to this tenant cluster. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAciCniTenantClusterAllocationAllOf

`func NewKubernetesAciCniTenantClusterAllocationAllOf(classId string, objectType string, ) *KubernetesAciCniTenantClusterAllocationAllOf`

NewKubernetesAciCniTenantClusterAllocationAllOf instantiates a new KubernetesAciCniTenantClusterAllocationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAciCniTenantClusterAllocationAllOfWithDefaults

`func NewKubernetesAciCniTenantClusterAllocationAllOfWithDefaults() *KubernetesAciCniTenantClusterAllocationAllOf`

NewKubernetesAciCniTenantClusterAllocationAllOfWithDefaults instantiates a new KubernetesAciCniTenantClusterAllocationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetNodeSvcIpSubnet() string`

GetNodeSvcIpSubnet returns the NodeSvcIpSubnet field if non-nil, zero value otherwise.

### GetNodeSvcIpSubnetOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetNodeSvcIpSubnetOk() (*string, bool)`

GetNodeSvcIpSubnetOk returns a tuple with the NodeSvcIpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetNodeSvcIpSubnet(v string)`

SetNodeSvcIpSubnet sets NodeSvcIpSubnet field to given value.

### HasNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasNodeSvcIpSubnet() bool`

HasNodeSvcIpSubnet returns a boolean if a field has been set.

### GetPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetPodIpSubnet() string`

GetPodIpSubnet returns the PodIpSubnet field if non-nil, zero value otherwise.

### GetPodIpSubnetOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetPodIpSubnetOk() (*string, bool)`

GetPodIpSubnetOk returns a tuple with the PodIpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetPodIpSubnet(v string)`

SetPodIpSubnet sets PodIpSubnet field to given value.

### HasPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasPodIpSubnet() bool`

HasPodIpSubnet returns a boolean if a field has been set.

### GetVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanEnd() string`

GetVlanEnd returns the VlanEnd field if non-nil, zero value otherwise.

### GetVlanEndOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanEndOk() (*string, bool)`

GetVlanEndOk returns a tuple with the VlanEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetVlanEnd(v string)`

SetVlanEnd sets VlanEnd field to given value.

### HasVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasVlanEnd() bool`

HasVlanEnd returns a boolean if a field has been set.

### GetVlanStart

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanStart() string`

GetVlanStart returns the VlanStart field if non-nil, zero value otherwise.

### GetVlanStartOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanStartOk() (*string, bool)`

GetVlanStartOk returns a tuple with the VlanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanStart

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetVlanStart(v string)`

SetVlanStart sets VlanStart field to given value.

### HasVlanStart

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasVlanStart() bool`

HasVlanStart returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


