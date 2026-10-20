# NiatelemetryNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Network"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Network"]
**FabricName** | Pointer to **string** | Fabric name associated with the network. | [optional] [readonly] 
**GatewayIpv4Address** | Pointer to **string** | IPv4 gateway address configured for the network. | [optional] [readonly] 
**GatewayIpv6Address** | Pointer to **string** | IPv6 gateway address configured for the network. | [optional] [readonly] 
**VlanId** | Pointer to **int64** | VLAN identifier associated with the network. | [optional] [readonly] 
**VrfName** | Pointer to **string** | VRF name associated with the network. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNetwork

`func NewNiatelemetryNetwork(classId string, objectType string, ) *NiatelemetryNetwork`

NewNiatelemetryNetwork instantiates a new NiatelemetryNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNetworkWithDefaults

`func NewNiatelemetryNetworkWithDefaults() *NiatelemetryNetwork`

NewNiatelemetryNetworkWithDefaults instantiates a new NiatelemetryNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFabricName

`func (o *NiatelemetryNetwork) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNetwork) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNetwork) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNetwork) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetGatewayIpv4Address

`func (o *NiatelemetryNetwork) GetGatewayIpv4Address() string`

GetGatewayIpv4Address returns the GatewayIpv4Address field if non-nil, zero value otherwise.

### GetGatewayIpv4AddressOk

`func (o *NiatelemetryNetwork) GetGatewayIpv4AddressOk() (*string, bool)`

GetGatewayIpv4AddressOk returns a tuple with the GatewayIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIpv4Address

`func (o *NiatelemetryNetwork) SetGatewayIpv4Address(v string)`

SetGatewayIpv4Address sets GatewayIpv4Address field to given value.

### HasGatewayIpv4Address

`func (o *NiatelemetryNetwork) HasGatewayIpv4Address() bool`

HasGatewayIpv4Address returns a boolean if a field has been set.

### GetGatewayIpv6Address

`func (o *NiatelemetryNetwork) GetGatewayIpv6Address() string`

GetGatewayIpv6Address returns the GatewayIpv6Address field if non-nil, zero value otherwise.

### GetGatewayIpv6AddressOk

`func (o *NiatelemetryNetwork) GetGatewayIpv6AddressOk() (*string, bool)`

GetGatewayIpv6AddressOk returns a tuple with the GatewayIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIpv6Address

`func (o *NiatelemetryNetwork) SetGatewayIpv6Address(v string)`

SetGatewayIpv6Address sets GatewayIpv6Address field to given value.

### HasGatewayIpv6Address

`func (o *NiatelemetryNetwork) HasGatewayIpv6Address() bool`

HasGatewayIpv6Address returns a boolean if a field has been set.

### GetVlanId

`func (o *NiatelemetryNetwork) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NiatelemetryNetwork) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NiatelemetryNetwork) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NiatelemetryNetwork) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVrfName

`func (o *NiatelemetryNetwork) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NiatelemetryNetwork) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NiatelemetryNetwork) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NiatelemetryNetwork) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryNetwork) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryNetwork) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryNetwork) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryNetwork) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryNetwork) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryNetwork) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


