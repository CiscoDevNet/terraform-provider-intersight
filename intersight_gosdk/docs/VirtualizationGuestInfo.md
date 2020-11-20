# VirtualizationGuestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.GuestInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.GuestInfo"]
**Hostname** | Pointer to **string** | Name provided to the host OS (example, ubuntu6410, test-gateway, etc.). | [optional] 
**IpAddress** | Pointer to **string** | Primary IP address of the guest os. | [optional] 
**Name** | Pointer to **string** | The name of the guest running on this VM. This may not be the same as the hostname. | [optional] 
**OperatingSystem** | Pointer to **string** | The name of the guest OS running on this VM (Cent OS 4/5/6/7). | [optional] 

## Methods

### NewVirtualizationGuestInfo

`func NewVirtualizationGuestInfo(classId string, objectType string, ) *VirtualizationGuestInfo`

NewVirtualizationGuestInfo instantiates a new VirtualizationGuestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationGuestInfoWithDefaults

`func NewVirtualizationGuestInfoWithDefaults() *VirtualizationGuestInfo`

NewVirtualizationGuestInfoWithDefaults instantiates a new VirtualizationGuestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationGuestInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationGuestInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationGuestInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationGuestInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationGuestInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationGuestInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *VirtualizationGuestInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VirtualizationGuestInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VirtualizationGuestInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VirtualizationGuestInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationGuestInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationGuestInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationGuestInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationGuestInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationGuestInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationGuestInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationGuestInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationGuestInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *VirtualizationGuestInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *VirtualizationGuestInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *VirtualizationGuestInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *VirtualizationGuestInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


