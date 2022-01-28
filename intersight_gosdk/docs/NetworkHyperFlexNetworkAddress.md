# NetworkHyperFlexNetworkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.HyperFlexNetworkAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.HyperFlexNetworkAddress"]
**Address** | Pointer to **string** | The IP or FQN of the HyperFlex appliance. | [optional] [readonly] 
**Fqdn** | Pointer to **string** | The fully qualified domain name of the HyperFlex appliance. | [optional] [readonly] 
**Ip** | Pointer to **string** | The IP address of the HyperFlex appliance. | [optional] [readonly] 

## Methods

### NewNetworkHyperFlexNetworkAddress

`func NewNetworkHyperFlexNetworkAddress(classId string, objectType string, ) *NetworkHyperFlexNetworkAddress`

NewNetworkHyperFlexNetworkAddress instantiates a new NetworkHyperFlexNetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHyperFlexNetworkAddressWithDefaults

`func NewNetworkHyperFlexNetworkAddressWithDefaults() *NetworkHyperFlexNetworkAddress`

NewNetworkHyperFlexNetworkAddressWithDefaults instantiates a new NetworkHyperFlexNetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkHyperFlexNetworkAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkHyperFlexNetworkAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkHyperFlexNetworkAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkHyperFlexNetworkAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkHyperFlexNetworkAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkHyperFlexNetworkAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *NetworkHyperFlexNetworkAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkHyperFlexNetworkAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkHyperFlexNetworkAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NetworkHyperFlexNetworkAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *NetworkHyperFlexNetworkAddress) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NetworkHyperFlexNetworkAddress) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NetworkHyperFlexNetworkAddress) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NetworkHyperFlexNetworkAddress) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetIp

`func (o *NetworkHyperFlexNetworkAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkHyperFlexNetworkAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkHyperFlexNetworkAddress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NetworkHyperFlexNetworkAddress) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


