# VirtualizationGuestInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Name provided to the host OS (example, ubuntu6410, test-gateway, etc.). | [optional] 
**IpAddress** | Pointer to **string** | Primary IP address of the guest os. | [optional] 
**Name** | Pointer to **string** | The name of the guest running on this VM. This may not be the same as the hostname. | [optional] 
**OperatingSystem** | Pointer to **string** | The name of the guest OS running on this VM (Cent OS 4/5/6/7). | [optional] 

## Methods

### NewVirtualizationGuestInfoAllOf

`func NewVirtualizationGuestInfoAllOf() *VirtualizationGuestInfoAllOf`

NewVirtualizationGuestInfoAllOf instantiates a new VirtualizationGuestInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationGuestInfoAllOfWithDefaults

`func NewVirtualizationGuestInfoAllOfWithDefaults() *VirtualizationGuestInfoAllOf`

NewVirtualizationGuestInfoAllOfWithDefaults instantiates a new VirtualizationGuestInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VirtualizationGuestInfoAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VirtualizationGuestInfoAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VirtualizationGuestInfoAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VirtualizationGuestInfoAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationGuestInfoAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationGuestInfoAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationGuestInfoAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationGuestInfoAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationGuestInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationGuestInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationGuestInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationGuestInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *VirtualizationGuestInfoAllOf) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *VirtualizationGuestInfoAllOf) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *VirtualizationGuestInfoAllOf) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *VirtualizationGuestInfoAllOf) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


