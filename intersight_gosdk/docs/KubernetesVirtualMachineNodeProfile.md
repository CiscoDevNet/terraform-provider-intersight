# KubernetesVirtualMachineNodeProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.VirtualMachineNodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.VirtualMachineNodeProfile"]
**IpAddresses** | Pointer to [**[]IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) | An array of relationships to ippoolIpLease resources. | [optional] 
**NodeIp** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationVirtualMachineRelationship**](VirtualizationVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVirtualMachineNodeProfile

`func NewKubernetesVirtualMachineNodeProfile(classId string, objectType string, ) *KubernetesVirtualMachineNodeProfile`

NewKubernetesVirtualMachineNodeProfile instantiates a new KubernetesVirtualMachineNodeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineNodeProfileWithDefaults

`func NewKubernetesVirtualMachineNodeProfileWithDefaults() *KubernetesVirtualMachineNodeProfile`

NewKubernetesVirtualMachineNodeProfileWithDefaults instantiates a new KubernetesVirtualMachineNodeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineNodeProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineNodeProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineNodeProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineNodeProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineNodeProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineNodeProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddresses

`func (o *KubernetesVirtualMachineNodeProfile) GetIpAddresses() []IppoolIpLeaseRelationship`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *KubernetesVirtualMachineNodeProfile) GetIpAddressesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *KubernetesVirtualMachineNodeProfile) SetIpAddresses(v []IppoolIpLeaseRelationship)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *KubernetesVirtualMachineNodeProfile) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *KubernetesVirtualMachineNodeProfile) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *KubernetesVirtualMachineNodeProfile) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetNodeIp

`func (o *KubernetesVirtualMachineNodeProfile) GetNodeIp() IppoolIpLeaseRelationship`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *KubernetesVirtualMachineNodeProfile) GetNodeIpOk() (*IppoolIpLeaseRelationship, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *KubernetesVirtualMachineNodeProfile) SetNodeIp(v IppoolIpLeaseRelationship)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *KubernetesVirtualMachineNodeProfile) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfile) GetVirtualMachine() VirtualizationVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *KubernetesVirtualMachineNodeProfile) GetVirtualMachineOk() (*VirtualizationVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfile) SetVirtualMachine(v VirtualizationVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfile) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


