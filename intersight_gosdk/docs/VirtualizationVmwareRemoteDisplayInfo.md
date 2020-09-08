# VirtualizationVmwareRemoteDisplayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteDisplayPassword** | Pointer to **string** | The password used for remote access. It is stored base64 encoded. | [optional] 
**RemoteDisplayVncKey** | Pointer to **string** | The access key for the remote display, potentially a long string. | [optional] 
**RemoteDisplayVncPort** | Pointer to **int64** | Applies only when remoteDisplayvnc is enabled. | [optional] 

## Methods

### NewVirtualizationVmwareRemoteDisplayInfo

`func NewVirtualizationVmwareRemoteDisplayInfo() *VirtualizationVmwareRemoteDisplayInfo`

NewVirtualizationVmwareRemoteDisplayInfo instantiates a new VirtualizationVmwareRemoteDisplayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareRemoteDisplayInfoWithDefaults

`func NewVirtualizationVmwareRemoteDisplayInfoWithDefaults() *VirtualizationVmwareRemoteDisplayInfo`

NewVirtualizationVmwareRemoteDisplayInfoWithDefaults instantiates a new VirtualizationVmwareRemoteDisplayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayPassword() string`

GetRemoteDisplayPassword returns the RemoteDisplayPassword field if non-nil, zero value otherwise.

### GetRemoteDisplayPasswordOk

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayPasswordOk() (*string, bool)`

GetRemoteDisplayPasswordOk returns a tuple with the RemoteDisplayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayPassword(v string)`

SetRemoteDisplayPassword sets RemoteDisplayPassword field to given value.

### HasRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayPassword() bool`

HasRemoteDisplayPassword returns a boolean if a field has been set.

### GetRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncKey() string`

GetRemoteDisplayVncKey returns the RemoteDisplayVncKey field if non-nil, zero value otherwise.

### GetRemoteDisplayVncKeyOk

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncKeyOk() (*string, bool)`

GetRemoteDisplayVncKeyOk returns a tuple with the RemoteDisplayVncKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayVncKey(v string)`

SetRemoteDisplayVncKey sets RemoteDisplayVncKey field to given value.

### HasRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayVncKey() bool`

HasRemoteDisplayVncKey returns a boolean if a field has been set.

### GetRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncPort() int64`

GetRemoteDisplayVncPort returns the RemoteDisplayVncPort field if non-nil, zero value otherwise.

### GetRemoteDisplayVncPortOk

`func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncPortOk() (*int64, bool)`

GetRemoteDisplayVncPortOk returns a tuple with the RemoteDisplayVncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayVncPort(v int64)`

SetRemoteDisplayVncPort sets RemoteDisplayVncPort field to given value.

### HasRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayVncPort() bool`

HasRemoteDisplayVncPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


