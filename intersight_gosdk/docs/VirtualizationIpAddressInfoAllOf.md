# VirtualizationIpAddressInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IpAddressInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IpAddressInfo"]
**GatewayIp** | Pointer to **string** | IP address of the device on network which forwards local traffic to other networks. | [optional] 
**IpAddress** | Pointer to **string** | An IP address is a 32-bit number. It uniquely identifies a host in given network. | [optional] 
**SubnetMask** | Pointer to **string** | A 32 bit number which helps to identify the host and rest of the network. | [optional] 

## Methods

### NewVirtualizationIpAddressInfoAllOf

`func NewVirtualizationIpAddressInfoAllOf(classId string, objectType string, ) *VirtualizationIpAddressInfoAllOf`

NewVirtualizationIpAddressInfoAllOf instantiates a new VirtualizationIpAddressInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIpAddressInfoAllOfWithDefaults

`func NewVirtualizationIpAddressInfoAllOfWithDefaults() *VirtualizationIpAddressInfoAllOf`

NewVirtualizationIpAddressInfoAllOfWithDefaults instantiates a new VirtualizationIpAddressInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIpAddressInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIpAddressInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIpAddressInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIpAddressInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIpAddressInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIpAddressInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGatewayIp

`func (o *VirtualizationIpAddressInfoAllOf) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *VirtualizationIpAddressInfoAllOf) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *VirtualizationIpAddressInfoAllOf) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *VirtualizationIpAddressInfoAllOf) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationIpAddressInfoAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationIpAddressInfoAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationIpAddressInfoAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationIpAddressInfoAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetSubnetMask

`func (o *VirtualizationIpAddressInfoAllOf) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VirtualizationIpAddressInfoAllOf) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VirtualizationIpAddressInfoAllOf) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *VirtualizationIpAddressInfoAllOf) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


