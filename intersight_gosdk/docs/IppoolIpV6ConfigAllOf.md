# IppoolIpV6ConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpV6Config"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpV6Config"]
**Gateway** | Pointer to **string** | IP address of the default IPv6 gateway. | [optional] 
**Prefix** | Pointer to **int64** | A prefix length which masks the  IP address and divides the IP address into network address and host address. | [optional] 
**PrimaryDns** | Pointer to **string** | IP Address of the primary Domain Name System (DNS) server. | [optional] 
**SecondaryDns** | Pointer to **string** | IP Address of the secondary Domain Name System (DNS) server. | [optional] 

## Methods

### NewIppoolIpV6ConfigAllOf

`func NewIppoolIpV6ConfigAllOf(classId string, objectType string, ) *IppoolIpV6ConfigAllOf`

NewIppoolIpV6ConfigAllOf instantiates a new IppoolIpV6ConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpV6ConfigAllOfWithDefaults

`func NewIppoolIpV6ConfigAllOfWithDefaults() *IppoolIpV6ConfigAllOf`

NewIppoolIpV6ConfigAllOfWithDefaults instantiates a new IppoolIpV6ConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpV6ConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpV6ConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpV6ConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpV6ConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpV6ConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpV6ConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *IppoolIpV6ConfigAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IppoolIpV6ConfigAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IppoolIpV6ConfigAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IppoolIpV6ConfigAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPrefix

`func (o *IppoolIpV6ConfigAllOf) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IppoolIpV6ConfigAllOf) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IppoolIpV6ConfigAllOf) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IppoolIpV6ConfigAllOf) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *IppoolIpV6ConfigAllOf) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *IppoolIpV6ConfigAllOf) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *IppoolIpV6ConfigAllOf) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *IppoolIpV6ConfigAllOf) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *IppoolIpV6ConfigAllOf) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *IppoolIpV6ConfigAllOf) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *IppoolIpV6ConfigAllOf) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *IppoolIpV6ConfigAllOf) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


