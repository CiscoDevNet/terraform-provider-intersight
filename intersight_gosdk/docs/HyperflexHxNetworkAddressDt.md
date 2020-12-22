# HyperflexHxNetworkAddressDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxNetworkAddressDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxNetworkAddressDt"]
**Address** | Pointer to **string** | The network address as an FQDN or IPv4 address. | [optional] [readonly] 
**Fqdn** | Pointer to **string** | The fully qualified domain name for the network address. | [optional] [readonly] 
**Ip** | Pointer to **string** | The network address as an IPv4 address. | [optional] [readonly] 

## Methods

### NewHyperflexHxNetworkAddressDt

`func NewHyperflexHxNetworkAddressDt(classId string, objectType string, ) *HyperflexHxNetworkAddressDt`

NewHyperflexHxNetworkAddressDt instantiates a new HyperflexHxNetworkAddressDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxNetworkAddressDtWithDefaults

`func NewHyperflexHxNetworkAddressDtWithDefaults() *HyperflexHxNetworkAddressDt`

NewHyperflexHxNetworkAddressDtWithDefaults instantiates a new HyperflexHxNetworkAddressDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxNetworkAddressDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxNetworkAddressDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxNetworkAddressDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxNetworkAddressDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxNetworkAddressDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxNetworkAddressDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *HyperflexHxNetworkAddressDt) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HyperflexHxNetworkAddressDt) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HyperflexHxNetworkAddressDt) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HyperflexHxNetworkAddressDt) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *HyperflexHxNetworkAddressDt) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *HyperflexHxNetworkAddressDt) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *HyperflexHxNetworkAddressDt) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *HyperflexHxNetworkAddressDt) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetIp

`func (o *HyperflexHxNetworkAddressDt) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexHxNetworkAddressDt) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexHxNetworkAddressDt) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexHxNetworkAddressDt) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


