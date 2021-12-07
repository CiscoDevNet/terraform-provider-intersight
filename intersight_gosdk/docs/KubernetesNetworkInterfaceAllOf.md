# KubernetesNetworkInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Addresses** | Pointer to **[]string** |  | [optional] 
**Gateway** | Pointer to **string** | Deprecated. This will add a default route as long as the first default route in Routes is not different. If is different, Gateway will be replaced with that default route. If there is no default Route and this is set, then Routes will be updated with the first entry as a default with this default gateway. If there is only one default Route and this gateway becomes empty, then the default routes will all be removed. Do not set if using Ip Pools, as the gateway is configured in the pool. This will be removed in the future. | [optional] 
**IpV4Configs** | Pointer to [**[]KubernetesIpV4Config**](KubernetesIpV4Config.md) |  | [optional] 
**Mtu** | Pointer to **int64** | The MTU to assign to this Network Interface. | [optional] 
**Name** | Pointer to **string** | Name for this network interface. | [optional] 
**Routes** | Pointer to [**[]KubernetesRoute**](KubernetesRoute.md) |  | [optional] 

## Methods

### NewKubernetesNetworkInterfaceAllOf

`func NewKubernetesNetworkInterfaceAllOf(classId string, objectType string, ) *KubernetesNetworkInterfaceAllOf`

NewKubernetesNetworkInterfaceAllOf instantiates a new KubernetesNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNetworkInterfaceAllOfWithDefaults

`func NewKubernetesNetworkInterfaceAllOfWithDefaults() *KubernetesNetworkInterfaceAllOf`

NewKubernetesNetworkInterfaceAllOfWithDefaults instantiates a new KubernetesNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddresses

`func (o *KubernetesNetworkInterfaceAllOf) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *KubernetesNetworkInterfaceAllOf) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *KubernetesNetworkInterfaceAllOf) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *KubernetesNetworkInterfaceAllOf) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *KubernetesNetworkInterfaceAllOf) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *KubernetesNetworkInterfaceAllOf) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetGateway

`func (o *KubernetesNetworkInterfaceAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *KubernetesNetworkInterfaceAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *KubernetesNetworkInterfaceAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *KubernetesNetworkInterfaceAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpV4Configs

`func (o *KubernetesNetworkInterfaceAllOf) GetIpV4Configs() []KubernetesIpV4Config`

GetIpV4Configs returns the IpV4Configs field if non-nil, zero value otherwise.

### GetIpV4ConfigsOk

`func (o *KubernetesNetworkInterfaceAllOf) GetIpV4ConfigsOk() (*[]KubernetesIpV4Config, bool)`

GetIpV4ConfigsOk returns a tuple with the IpV4Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Configs

`func (o *KubernetesNetworkInterfaceAllOf) SetIpV4Configs(v []KubernetesIpV4Config)`

SetIpV4Configs sets IpV4Configs field to given value.

### HasIpV4Configs

`func (o *KubernetesNetworkInterfaceAllOf) HasIpV4Configs() bool`

HasIpV4Configs returns a boolean if a field has been set.

### SetIpV4ConfigsNil

`func (o *KubernetesNetworkInterfaceAllOf) SetIpV4ConfigsNil(b bool)`

 SetIpV4ConfigsNil sets the value for IpV4Configs to be an explicit nil

### UnsetIpV4Configs
`func (o *KubernetesNetworkInterfaceAllOf) UnsetIpV4Configs()`

UnsetIpV4Configs ensures that no value is present for IpV4Configs, not even an explicit nil
### GetMtu

`func (o *KubernetesNetworkInterfaceAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *KubernetesNetworkInterfaceAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *KubernetesNetworkInterfaceAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *KubernetesNetworkInterfaceAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *KubernetesNetworkInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNetworkInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNetworkInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNetworkInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *KubernetesNetworkInterfaceAllOf) GetRoutes() []KubernetesRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *KubernetesNetworkInterfaceAllOf) GetRoutesOk() (*[]KubernetesRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *KubernetesNetworkInterfaceAllOf) SetRoutes(v []KubernetesRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *KubernetesNetworkInterfaceAllOf) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *KubernetesNetworkInterfaceAllOf) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *KubernetesNetworkInterfaceAllOf) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


