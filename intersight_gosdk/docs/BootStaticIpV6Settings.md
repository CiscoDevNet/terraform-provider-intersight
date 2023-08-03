# BootStaticIpV6Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.StaticIpV6Settings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.StaticIpV6Settings"]
**DnsIp** | Pointer to **string** | IP address of DNS server. | [optional] 
**GatewayIp** | Pointer to **string** | IP address of default gateway. | [optional] 
**Ip** | Pointer to **string** | Ipv6 static Internet Protocol address. | [optional] 
**PrefixLength** | Pointer to **int64** | A prefix length which masks the IP address and divides the IP address into network address and host address. | [optional] [default to 1]

## Methods

### NewBootStaticIpV6Settings

`func NewBootStaticIpV6Settings(classId string, objectType string, ) *BootStaticIpV6Settings`

NewBootStaticIpV6Settings instantiates a new BootStaticIpV6Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootStaticIpV6SettingsWithDefaults

`func NewBootStaticIpV6SettingsWithDefaults() *BootStaticIpV6Settings`

NewBootStaticIpV6SettingsWithDefaults instantiates a new BootStaticIpV6Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootStaticIpV6Settings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootStaticIpV6Settings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootStaticIpV6Settings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootStaticIpV6Settings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootStaticIpV6Settings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootStaticIpV6Settings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDnsIp

`func (o *BootStaticIpV6Settings) GetDnsIp() string`

GetDnsIp returns the DnsIp field if non-nil, zero value otherwise.

### GetDnsIpOk

`func (o *BootStaticIpV6Settings) GetDnsIpOk() (*string, bool)`

GetDnsIpOk returns a tuple with the DnsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIp

`func (o *BootStaticIpV6Settings) SetDnsIp(v string)`

SetDnsIp sets DnsIp field to given value.

### HasDnsIp

`func (o *BootStaticIpV6Settings) HasDnsIp() bool`

HasDnsIp returns a boolean if a field has been set.

### GetGatewayIp

`func (o *BootStaticIpV6Settings) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *BootStaticIpV6Settings) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *BootStaticIpV6Settings) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *BootStaticIpV6Settings) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetIp

`func (o *BootStaticIpV6Settings) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *BootStaticIpV6Settings) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *BootStaticIpV6Settings) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *BootStaticIpV6Settings) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPrefixLength

`func (o *BootStaticIpV6Settings) GetPrefixLength() int64`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *BootStaticIpV6Settings) GetPrefixLengthOk() (*int64, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *BootStaticIpV6Settings) SetPrefixLength(v int64)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *BootStaticIpV6Settings) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


