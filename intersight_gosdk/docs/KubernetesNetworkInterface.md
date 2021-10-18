# KubernetesNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Addresses** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** | The Network Gateway for the Network Interface. | [optional] 
**Mtu** | Pointer to **int64** | The MTU to assign to this Network Interface. | [optional] 
**Name** | Pointer to **string** | Name for this network interface. | [optional] 

## Methods

### NewKubernetesNetworkInterface

`func NewKubernetesNetworkInterface(classId string, objectType string, ) *KubernetesNetworkInterface`

NewKubernetesNetworkInterface instantiates a new KubernetesNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNetworkInterfaceWithDefaults

`func NewKubernetesNetworkInterfaceWithDefaults() *KubernetesNetworkInterface`

NewKubernetesNetworkInterfaceWithDefaults instantiates a new KubernetesNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddresses

`func (o *KubernetesNetworkInterface) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *KubernetesNetworkInterface) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *KubernetesNetworkInterface) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *KubernetesNetworkInterface) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *KubernetesNetworkInterface) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *KubernetesNetworkInterface) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetGateway

`func (o *KubernetesNetworkInterface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *KubernetesNetworkInterface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *KubernetesNetworkInterface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *KubernetesNetworkInterface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMtu

`func (o *KubernetesNetworkInterface) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *KubernetesNetworkInterface) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *KubernetesNetworkInterface) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *KubernetesNetworkInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *KubernetesNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


