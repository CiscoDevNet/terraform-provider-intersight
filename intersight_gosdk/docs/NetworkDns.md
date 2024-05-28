# NetworkDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Dns"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Dns"]
**AdditionalDomains** | Pointer to **[]string** |  | [optional] 
**DefaultDomain** | Pointer to **string** | Default domain configured for VRF. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**VrfName** | Pointer to **string** | Name of the VRF configured for the DNS. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkDns

`func NewNetworkDns(classId string, objectType string, ) *NetworkDns`

NewNetworkDns instantiates a new NetworkDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDnsWithDefaults

`func NewNetworkDnsWithDefaults() *NetworkDns`

NewNetworkDnsWithDefaults instantiates a new NetworkDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkDns) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkDns) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkDns) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkDns) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkDns) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkDns) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalDomains

`func (o *NetworkDns) GetAdditionalDomains() []string`

GetAdditionalDomains returns the AdditionalDomains field if non-nil, zero value otherwise.

### GetAdditionalDomainsOk

`func (o *NetworkDns) GetAdditionalDomainsOk() (*[]string, bool)`

GetAdditionalDomainsOk returns a tuple with the AdditionalDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDomains

`func (o *NetworkDns) SetAdditionalDomains(v []string)`

SetAdditionalDomains sets AdditionalDomains field to given value.

### HasAdditionalDomains

`func (o *NetworkDns) HasAdditionalDomains() bool`

HasAdditionalDomains returns a boolean if a field has been set.

### SetAdditionalDomainsNil

`func (o *NetworkDns) SetAdditionalDomainsNil(b bool)`

 SetAdditionalDomainsNil sets the value for AdditionalDomains to be an explicit nil

### UnsetAdditionalDomains
`func (o *NetworkDns) UnsetAdditionalDomains()`

UnsetAdditionalDomains ensures that no value is present for AdditionalDomains, not even an explicit nil
### GetDefaultDomain

`func (o *NetworkDns) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *NetworkDns) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *NetworkDns) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *NetworkDns) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetNameServers

`func (o *NetworkDns) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *NetworkDns) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *NetworkDns) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *NetworkDns) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *NetworkDns) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *NetworkDns) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetVrfName

`func (o *NetworkDns) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *NetworkDns) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *NetworkDns) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *NetworkDns) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkDns) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkDns) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkDns) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkDns) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkDns) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkDns) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkDns) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkDns) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkDns) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkDns) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkDns) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkDns) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


