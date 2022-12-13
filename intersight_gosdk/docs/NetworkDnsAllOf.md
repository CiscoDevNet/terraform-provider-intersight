# NetworkDnsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Dns"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Dns"]
**AdditionalDomains** | Pointer to **[]string** |  | [optional] 
**DefaultDomain** | Pointer to **string** | Default domain configured for VRF. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**VrfName** | Pointer to **string** | Name of the VRF configured for the DNS. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkDnsAllOf

`func NewNetworkDnsAllOf(classId string, objectType string, ) *NetworkDnsAllOf`

NewNetworkDnsAllOf instantiates a new NetworkDnsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDnsAllOfWithDefaults

`func NewNetworkDnsAllOfWithDefaults() *NetworkDnsAllOf`

NewNetworkDnsAllOfWithDefaults instantiates a new NetworkDnsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkDnsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkDnsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkDnsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkDnsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkDnsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkDnsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalDomains

`func (o *NetworkDnsAllOf) GetAdditionalDomains() []string`

GetAdditionalDomains returns the AdditionalDomains field if non-nil, zero value otherwise.

### GetAdditionalDomainsOk

`func (o *NetworkDnsAllOf) GetAdditionalDomainsOk() (*[]string, bool)`

GetAdditionalDomainsOk returns a tuple with the AdditionalDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDomains

`func (o *NetworkDnsAllOf) SetAdditionalDomains(v []string)`

SetAdditionalDomains sets AdditionalDomains field to given value.

### HasAdditionalDomains

`func (o *NetworkDnsAllOf) HasAdditionalDomains() bool`

HasAdditionalDomains returns a boolean if a field has been set.

### SetAdditionalDomainsNil

`func (o *NetworkDnsAllOf) SetAdditionalDomainsNil(b bool)`

 SetAdditionalDomainsNil sets the value for AdditionalDomains to be an explicit nil

### UnsetAdditionalDomains
`func (o *NetworkDnsAllOf) UnsetAdditionalDomains()`

UnsetAdditionalDomains ensures that no value is present for AdditionalDomains, not even an explicit nil
### GetDefaultDomain

`func (o *NetworkDnsAllOf) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *NetworkDnsAllOf) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *NetworkDnsAllOf) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *NetworkDnsAllOf) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetNameServers

`func (o *NetworkDnsAllOf) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *NetworkDnsAllOf) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *NetworkDnsAllOf) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *NetworkDnsAllOf) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *NetworkDnsAllOf) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *NetworkDnsAllOf) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetVrfName

`func (o *NetworkDnsAllOf) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NetworkDnsAllOf) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NetworkDnsAllOf) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NetworkDnsAllOf) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkDnsAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkDnsAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkDnsAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkDnsAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkDnsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkDnsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkDnsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkDnsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


