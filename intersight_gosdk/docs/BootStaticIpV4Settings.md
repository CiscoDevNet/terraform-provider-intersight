# BootStaticIpV4Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.StaticIpV4Settings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.StaticIpV4Settings"]
**DnsIp** | Pointer to **string** | IP address of DNS server. | [optional] 
**GatewayIp** | Pointer to **string** | IP address of default gateway. | [optional] 
**Ip** | Pointer to **string** | Ipv4 static Internet Protocol address. | [optional] 
**NetworkMask** | Pointer to **string** | Network mask of the IP address. | [optional] 

## Methods

### NewBootStaticIpV4Settings

`func NewBootStaticIpV4Settings(classId string, objectType string, ) *BootStaticIpV4Settings`

NewBootStaticIpV4Settings instantiates a new BootStaticIpV4Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootStaticIpV4SettingsWithDefaults

`func NewBootStaticIpV4SettingsWithDefaults() *BootStaticIpV4Settings`

NewBootStaticIpV4SettingsWithDefaults instantiates a new BootStaticIpV4Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootStaticIpV4Settings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootStaticIpV4Settings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootStaticIpV4Settings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootStaticIpV4Settings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootStaticIpV4Settings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootStaticIpV4Settings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDnsIp

`func (o *BootStaticIpV4Settings) GetDnsIp() string`

GetDnsIp returns the DnsIp field if non-nil, zero value otherwise.

### GetDnsIpOk

`func (o *BootStaticIpV4Settings) GetDnsIpOk() (*string, bool)`

GetDnsIpOk returns a tuple with the DnsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIp

`func (o *BootStaticIpV4Settings) SetDnsIp(v string)`

SetDnsIp sets DnsIp field to given value.

### HasDnsIp

`func (o *BootStaticIpV4Settings) HasDnsIp() bool`

HasDnsIp returns a boolean if a field has been set.

### GetGatewayIp

`func (o *BootStaticIpV4Settings) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *BootStaticIpV4Settings) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *BootStaticIpV4Settings) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *BootStaticIpV4Settings) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetIp

`func (o *BootStaticIpV4Settings) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *BootStaticIpV4Settings) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *BootStaticIpV4Settings) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *BootStaticIpV4Settings) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetworkMask

`func (o *BootStaticIpV4Settings) GetNetworkMask() string`

GetNetworkMask returns the NetworkMask field if non-nil, zero value otherwise.

### GetNetworkMaskOk

`func (o *BootStaticIpV4Settings) GetNetworkMaskOk() (*string, bool)`

GetNetworkMaskOk returns a tuple with the NetworkMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMask

`func (o *BootStaticIpV4Settings) SetNetworkMask(v string)`

SetNetworkMask sets NetworkMask field to given value.

### HasNetworkMask

`func (o *BootStaticIpV4Settings) HasNetworkMask() bool`

HasNetworkMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


