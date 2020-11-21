# VirtualizationVmwareRemoteDisplayInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareRemoteDisplayInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareRemoteDisplayInfo"]
**RemoteDisplayPassword** | Pointer to **string** | The password used for remote access. It is stored base64 encoded. | [optional] 
**RemoteDisplayVncKey** | Pointer to **string** | The access key for the remote display, potentially a long string. | [optional] 
**RemoteDisplayVncPort** | Pointer to **int64** | Applies only when remoteDisplayvnc is enabled. | [optional] 

## Methods

### NewVirtualizationVmwareRemoteDisplayInfoAllOf

`func NewVirtualizationVmwareRemoteDisplayInfoAllOf(classId string, objectType string, ) *VirtualizationVmwareRemoteDisplayInfoAllOf`

NewVirtualizationVmwareRemoteDisplayInfoAllOf instantiates a new VirtualizationVmwareRemoteDisplayInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareRemoteDisplayInfoAllOfWithDefaults

`func NewVirtualizationVmwareRemoteDisplayInfoAllOfWithDefaults() *VirtualizationVmwareRemoteDisplayInfoAllOf`

NewVirtualizationVmwareRemoteDisplayInfoAllOfWithDefaults instantiates a new VirtualizationVmwareRemoteDisplayInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayPassword() string`

GetRemoteDisplayPassword returns the RemoteDisplayPassword field if non-nil, zero value otherwise.

### GetRemoteDisplayPasswordOk

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayPasswordOk() (*string, bool)`

GetRemoteDisplayPasswordOk returns a tuple with the RemoteDisplayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) SetRemoteDisplayPassword(v string)`

SetRemoteDisplayPassword sets RemoteDisplayPassword field to given value.

### HasRemoteDisplayPassword

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) HasRemoteDisplayPassword() bool`

HasRemoteDisplayPassword returns a boolean if a field has been set.

### GetRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayVncKey() string`

GetRemoteDisplayVncKey returns the RemoteDisplayVncKey field if non-nil, zero value otherwise.

### GetRemoteDisplayVncKeyOk

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayVncKeyOk() (*string, bool)`

GetRemoteDisplayVncKeyOk returns a tuple with the RemoteDisplayVncKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) SetRemoteDisplayVncKey(v string)`

SetRemoteDisplayVncKey sets RemoteDisplayVncKey field to given value.

### HasRemoteDisplayVncKey

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) HasRemoteDisplayVncKey() bool`

HasRemoteDisplayVncKey returns a boolean if a field has been set.

### GetRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayVncPort() int64`

GetRemoteDisplayVncPort returns the RemoteDisplayVncPort field if non-nil, zero value otherwise.

### GetRemoteDisplayVncPortOk

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) GetRemoteDisplayVncPortOk() (*int64, bool)`

GetRemoteDisplayVncPortOk returns a tuple with the RemoteDisplayVncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) SetRemoteDisplayVncPort(v int64)`

SetRemoteDisplayVncPort sets RemoteDisplayVncPort field to given value.

### HasRemoteDisplayVncPort

`func (o *VirtualizationVmwareRemoteDisplayInfoAllOf) HasRemoteDisplayVncPort() bool`

HasRemoteDisplayVncPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


