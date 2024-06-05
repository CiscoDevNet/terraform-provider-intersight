# DnacDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.Device"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.Device"]
**DeviceId** | Pointer to **string** | Unique identity of the network device. | [optional] 
**HostName** | Pointer to **string** | Host name of the network device. | [optional] 
**ManagementIpAddress** | Pointer to **string** | IP address of the network device. | [optional] 

## Methods

### NewDnacDevice

`func NewDnacDevice(classId string, objectType string, ) *DnacDevice`

NewDnacDevice instantiates a new DnacDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacDeviceWithDefaults

`func NewDnacDeviceWithDefaults() *DnacDevice`

NewDnacDeviceWithDefaults instantiates a new DnacDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *DnacDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DnacDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DnacDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DnacDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHostName

`func (o *DnacDevice) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *DnacDevice) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *DnacDevice) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *DnacDevice) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *DnacDevice) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *DnacDevice) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *DnacDevice) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *DnacDevice) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


