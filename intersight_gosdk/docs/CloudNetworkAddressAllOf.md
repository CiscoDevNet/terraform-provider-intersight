# CloudNetworkAddressAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.NetworkAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.NetworkAddress"]
**Ip** | Pointer to **string** | IP (Internet Protocol) address value. | [optional] [readonly] 
**IpAllocation** | Pointer to **string** | IP address allocation type (DYNAMIC | STATIC | IPAM_CALLOUT | PRE_ALLOCATE). * &#x60;Dynamic&#x60; - IP address allocation type is dynamic. * &#x60;Static&#x60; - IP address allocation type is static. * &#x60;IpamCallout&#x60; - IP address is assigned with the results of callout scripts execution. * &#x60;PreAllocate&#x60; - IP address allocation type is PreAllocate . | [optional] [readonly] [default to "Dynamic"]
**IpVersion** | Pointer to **string** | Whether IP address is of type IPv4 or IPv6. * &#x60;IPv4&#x60; - Internet protocol version 4. * &#x60;IPv6&#x60; - Internet protocol version 6. | [optional] [readonly] [default to "IPv4"]

## Methods

### NewCloudNetworkAddressAllOf

`func NewCloudNetworkAddressAllOf(classId string, objectType string, ) *CloudNetworkAddressAllOf`

NewCloudNetworkAddressAllOf instantiates a new CloudNetworkAddressAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkAddressAllOfWithDefaults

`func NewCloudNetworkAddressAllOfWithDefaults() *CloudNetworkAddressAllOf`

NewCloudNetworkAddressAllOfWithDefaults instantiates a new CloudNetworkAddressAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudNetworkAddressAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudNetworkAddressAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudNetworkAddressAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudNetworkAddressAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkAddressAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkAddressAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIp

`func (o *CloudNetworkAddressAllOf) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CloudNetworkAddressAllOf) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CloudNetworkAddressAllOf) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CloudNetworkAddressAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpAllocation

`func (o *CloudNetworkAddressAllOf) GetIpAllocation() string`

GetIpAllocation returns the IpAllocation field if non-nil, zero value otherwise.

### GetIpAllocationOk

`func (o *CloudNetworkAddressAllOf) GetIpAllocationOk() (*string, bool)`

GetIpAllocationOk returns a tuple with the IpAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocation

`func (o *CloudNetworkAddressAllOf) SetIpAllocation(v string)`

SetIpAllocation sets IpAllocation field to given value.

### HasIpAllocation

`func (o *CloudNetworkAddressAllOf) HasIpAllocation() bool`

HasIpAllocation returns a boolean if a field has been set.

### GetIpVersion

`func (o *CloudNetworkAddressAllOf) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *CloudNetworkAddressAllOf) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *CloudNetworkAddressAllOf) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *CloudNetworkAddressAllOf) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


