# KubernetesVirtualMachineNodeProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.VirtualMachineNodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.VirtualMachineNodeProfile"]
**IpAddresses** | Pointer to [**[]IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) | An array of relationships to ippoolIpLease resources. | [optional] 
**NodeIp** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationVirtualMachineRelationship**](VirtualizationVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVirtualMachineNodeProfileAllOf

`func NewKubernetesVirtualMachineNodeProfileAllOf(classId string, objectType string, ) *KubernetesVirtualMachineNodeProfileAllOf`

NewKubernetesVirtualMachineNodeProfileAllOf instantiates a new KubernetesVirtualMachineNodeProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineNodeProfileAllOfWithDefaults

`func NewKubernetesVirtualMachineNodeProfileAllOfWithDefaults() *KubernetesVirtualMachineNodeProfileAllOf`

NewKubernetesVirtualMachineNodeProfileAllOfWithDefaults instantiates a new KubernetesVirtualMachineNodeProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddresses

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetIpAddresses() []IppoolIpLeaseRelationship`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetIpAddressesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetIpAddresses(v []IppoolIpLeaseRelationship)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *KubernetesVirtualMachineNodeProfileAllOf) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *KubernetesVirtualMachineNodeProfileAllOf) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetNodeIp

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetNodeIp() IppoolIpLeaseRelationship`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetNodeIpOk() (*IppoolIpLeaseRelationship, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetNodeIp(v IppoolIpLeaseRelationship)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *KubernetesVirtualMachineNodeProfileAllOf) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetVirtualMachine() VirtualizationVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *KubernetesVirtualMachineNodeProfileAllOf) GetVirtualMachineOk() (*VirtualizationVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfileAllOf) SetVirtualMachine(v VirtualizationVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *KubernetesVirtualMachineNodeProfileAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


