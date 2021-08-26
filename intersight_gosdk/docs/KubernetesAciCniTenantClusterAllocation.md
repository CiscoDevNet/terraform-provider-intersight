# KubernetesAciCniTenantClusterAllocation

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

### NewKubernetesAciCniTenantClusterAllocation

`func NewKubernetesAciCniTenantClusterAllocation(classId string, objectType string, ) *KubernetesAciCniTenantClusterAllocation`

NewKubernetesAciCniTenantClusterAllocation instantiates a new KubernetesAciCniTenantClusterAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAciCniTenantClusterAllocationWithDefaults

`func NewKubernetesAciCniTenantClusterAllocationWithDefaults() *KubernetesAciCniTenantClusterAllocation`

NewKubernetesAciCniTenantClusterAllocationWithDefaults instantiates a new KubernetesAciCniTenantClusterAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAciCniTenantClusterAllocation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAciCniTenantClusterAllocation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAciCniTenantClusterAllocation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAciCniTenantClusterAllocation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) GetNodeSvcIpSubnet() string`

GetNodeSvcIpSubnet returns the NodeSvcIpSubnet field if non-nil, zero value otherwise.

### GetNodeSvcIpSubnetOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetNodeSvcIpSubnetOk() (*string, bool)`

GetNodeSvcIpSubnetOk returns a tuple with the NodeSvcIpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) SetNodeSvcIpSubnet(v string)`

SetNodeSvcIpSubnet sets NodeSvcIpSubnet field to given value.

### HasNodeSvcIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) HasNodeSvcIpSubnet() bool`

HasNodeSvcIpSubnet returns a boolean if a field has been set.

### GetPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) GetPodIpSubnet() string`

GetPodIpSubnet returns the PodIpSubnet field if non-nil, zero value otherwise.

### GetPodIpSubnetOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetPodIpSubnetOk() (*string, bool)`

GetPodIpSubnetOk returns a tuple with the PodIpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) SetPodIpSubnet(v string)`

SetPodIpSubnet sets PodIpSubnet field to given value.

### HasPodIpSubnet

`func (o *KubernetesAciCniTenantClusterAllocation) HasPodIpSubnet() bool`

HasPodIpSubnet returns a boolean if a field has been set.

### GetVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocation) GetVlanEnd() string`

GetVlanEnd returns the VlanEnd field if non-nil, zero value otherwise.

### GetVlanEndOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetVlanEndOk() (*string, bool)`

GetVlanEndOk returns a tuple with the VlanEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocation) SetVlanEnd(v string)`

SetVlanEnd sets VlanEnd field to given value.

### HasVlanEnd

`func (o *KubernetesAciCniTenantClusterAllocation) HasVlanEnd() bool`

HasVlanEnd returns a boolean if a field has been set.

### GetVlanStart

`func (o *KubernetesAciCniTenantClusterAllocation) GetVlanStart() string`

GetVlanStart returns the VlanStart field if non-nil, zero value otherwise.

### GetVlanStartOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetVlanStartOk() (*string, bool)`

GetVlanStartOk returns a tuple with the VlanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanStart

`func (o *KubernetesAciCniTenantClusterAllocation) SetVlanStart(v string)`

SetVlanStart sets VlanStart field to given value.

### HasVlanStart

`func (o *KubernetesAciCniTenantClusterAllocation) HasVlanStart() bool`

HasVlanStart returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAciCniTenantClusterAllocation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAciCniTenantClusterAllocation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAciCniTenantClusterAllocation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAciCniTenantClusterAllocation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


