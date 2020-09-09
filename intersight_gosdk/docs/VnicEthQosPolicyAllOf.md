# VnicEthQosPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] 
**Mtu** | Pointer to **int64** | The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. | [optional] 
**Priority** | Pointer to **string** | The priortity matching the System QoS specified in the fabric profile. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [default to "Best Effort"]
**RateLimit** | Pointer to **int64** | The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. | [optional] 
**TrustHostCos** | Pointer to **bool** | Enables usage of the Class of Service provided by the operating system. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicEthQosPolicyAllOf

`func NewVnicEthQosPolicyAllOf() *VnicEthQosPolicyAllOf`

NewVnicEthQosPolicyAllOf instantiates a new VnicEthQosPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthQosPolicyAllOfWithDefaults

`func NewVnicEthQosPolicyAllOfWithDefaults() *VnicEthQosPolicyAllOf`

NewVnicEthQosPolicyAllOfWithDefaults instantiates a new VnicEthQosPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCos

`func (o *VnicEthQosPolicyAllOf) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicEthQosPolicyAllOf) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicEthQosPolicyAllOf) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicEthQosPolicyAllOf) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMtu

`func (o *VnicEthQosPolicyAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VnicEthQosPolicyAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VnicEthQosPolicyAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VnicEthQosPolicyAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetPriority

`func (o *VnicEthQosPolicyAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VnicEthQosPolicyAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VnicEthQosPolicyAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VnicEthQosPolicyAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRateLimit

`func (o *VnicEthQosPolicyAllOf) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VnicEthQosPolicyAllOf) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VnicEthQosPolicyAllOf) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VnicEthQosPolicyAllOf) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetTrustHostCos

`func (o *VnicEthQosPolicyAllOf) GetTrustHostCos() bool`

GetTrustHostCos returns the TrustHostCos field if non-nil, zero value otherwise.

### GetTrustHostCosOk

`func (o *VnicEthQosPolicyAllOf) GetTrustHostCosOk() (*bool, bool)`

GetTrustHostCosOk returns a tuple with the TrustHostCos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustHostCos

`func (o *VnicEthQosPolicyAllOf) SetTrustHostCos(v bool)`

SetTrustHostCos sets TrustHostCos field to given value.

### HasTrustHostCos

`func (o *VnicEthQosPolicyAllOf) HasTrustHostCos() bool`

HasTrustHostCos returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicEthQosPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicEthQosPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicEthQosPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicEthQosPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


