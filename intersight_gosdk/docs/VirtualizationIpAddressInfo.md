# VirtualizationIpAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IpAddressInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IpAddressInfo"]
**GatewayIp** | Pointer to **string** | IP address of the device on network which forwards local traffic to other networks. | [optional] 
**IpAddress** | Pointer to **string** | An IP address is a 32-bit number. It uniquely identifies a host in given network. | [optional] 
**SubnetMask** | Pointer to **string** | A 32 bit number which helps to identify the host and rest of the network. | [optional] 

## Methods

### NewVirtualizationIpAddressInfo

`func NewVirtualizationIpAddressInfo(classId string, objectType string, ) *VirtualizationIpAddressInfo`

NewVirtualizationIpAddressInfo instantiates a new VirtualizationIpAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIpAddressInfoWithDefaults

`func NewVirtualizationIpAddressInfoWithDefaults() *VirtualizationIpAddressInfo`

NewVirtualizationIpAddressInfoWithDefaults instantiates a new VirtualizationIpAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIpAddressInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIpAddressInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIpAddressInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIpAddressInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIpAddressInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIpAddressInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGatewayIp

`func (o *VirtualizationIpAddressInfo) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *VirtualizationIpAddressInfo) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *VirtualizationIpAddressInfo) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *VirtualizationIpAddressInfo) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationIpAddressInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationIpAddressInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationIpAddressInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationIpAddressInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetSubnetMask

`func (o *VirtualizationIpAddressInfo) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *VirtualizationIpAddressInfo) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *VirtualizationIpAddressInfo) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *VirtualizationIpAddressInfo) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


