# IppoolIpV4Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpV4Config"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpV4Config"]
**Gateway** | Pointer to **string** | IP address of the default IPv4 gateway. | [optional] 
**Netmask** | Pointer to **string** | A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address. | [optional] 
**PrimaryDns** | Pointer to **string** | IP Address of the primary Domain Name System (DNS) server. | [optional] 
**SecondaryDns** | Pointer to **string** | IP Address of the secondary Domain Name System (DNS) server. | [optional] 

## Methods

### NewIppoolIpV4Config

`func NewIppoolIpV4Config(classId string, objectType string, ) *IppoolIpV4Config`

NewIppoolIpV4Config instantiates a new IppoolIpV4Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpV4ConfigWithDefaults

`func NewIppoolIpV4ConfigWithDefaults() *IppoolIpV4Config`

NewIppoolIpV4ConfigWithDefaults instantiates a new IppoolIpV4Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpV4Config) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpV4Config) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpV4Config) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpV4Config) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpV4Config) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpV4Config) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *IppoolIpV4Config) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IppoolIpV4Config) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IppoolIpV4Config) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IppoolIpV4Config) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *IppoolIpV4Config) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IppoolIpV4Config) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IppoolIpV4Config) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IppoolIpV4Config) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *IppoolIpV4Config) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *IppoolIpV4Config) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *IppoolIpV4Config) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *IppoolIpV4Config) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *IppoolIpV4Config) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *IppoolIpV4Config) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *IppoolIpV4Config) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *IppoolIpV4Config) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


