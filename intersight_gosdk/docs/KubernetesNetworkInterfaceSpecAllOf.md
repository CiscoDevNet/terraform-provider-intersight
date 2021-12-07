# KubernetesNetworkInterfaceSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NetworkInterfaceSpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NetworkInterfaceSpec"]
**Mtu** | Pointer to **int64** | The MTU for this Network Interface.  If left blank a default value will apply by the Operating System. | [optional] 
**Name** | Pointer to **string** | NetworkInterfaceSpec is the specification for network interfaces - including configuration of IP Pools and VRF to determine IP configuration, the operating system device settings, and virtual adapter network settings. It can be left empty when used with VirtualMachineInfraConfigPolicy - it will be filled out based on the hypervisor platform type and will match the naming and order of interfaces provided by the hypervisor. | [optional] 
**Pools** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**ProviderName** | Pointer to **string** | In other words, to which named network from the Instructure Provider will the port be connected. | [optional] 
**Vrf** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewKubernetesNetworkInterfaceSpecAllOf

`func NewKubernetesNetworkInterfaceSpecAllOf(classId string, objectType string, ) *KubernetesNetworkInterfaceSpecAllOf`

NewKubernetesNetworkInterfaceSpecAllOf instantiates a new KubernetesNetworkInterfaceSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNetworkInterfaceSpecAllOfWithDefaults

`func NewKubernetesNetworkInterfaceSpecAllOfWithDefaults() *KubernetesNetworkInterfaceSpecAllOf`

NewKubernetesNetworkInterfaceSpecAllOfWithDefaults instantiates a new KubernetesNetworkInterfaceSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMtu

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *KubernetesNetworkInterfaceSpecAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNetworkInterfaceSpecAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPools

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetPools() []MoMoRef`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetPoolsOk() (*[]MoMoRef, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetPools(v []MoMoRef)`

SetPools sets Pools field to given value.

### HasPools

`func (o *KubernetesNetworkInterfaceSpecAllOf) HasPools() bool`

HasPools returns a boolean if a field has been set.

### SetPoolsNil

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetPoolsNil(b bool)`

 SetPoolsNil sets the value for Pools to be an explicit nil

### UnsetPools
`func (o *KubernetesNetworkInterfaceSpecAllOf) UnsetPools()`

UnsetPools ensures that no value is present for Pools, not even an explicit nil
### GetProviderName

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *KubernetesNetworkInterfaceSpecAllOf) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetVrf

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetVrf() MoMoRef`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *KubernetesNetworkInterfaceSpecAllOf) GetVrfOk() (*MoMoRef, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *KubernetesNetworkInterfaceSpecAllOf) SetVrf(v MoMoRef)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *KubernetesNetworkInterfaceSpecAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


