# VirtualizationEsxiConsole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiConsole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiConsole"]
**StreamId** | Pointer to **string** | The stream ID of the host console session opened. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Host** | Pointer to [**NullableVirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiConsole

`func NewVirtualizationEsxiConsole(classId string, objectType string, ) *VirtualizationEsxiConsole`

NewVirtualizationEsxiConsole instantiates a new VirtualizationEsxiConsole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiConsoleWithDefaults

`func NewVirtualizationEsxiConsoleWithDefaults() *VirtualizationEsxiConsole`

NewVirtualizationEsxiConsoleWithDefaults instantiates a new VirtualizationEsxiConsole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiConsole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiConsole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiConsole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiConsole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiConsole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiConsole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStreamId

`func (o *VirtualizationEsxiConsole) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *VirtualizationEsxiConsole) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *VirtualizationEsxiConsole) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *VirtualizationEsxiConsole) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *VirtualizationEsxiConsole) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *VirtualizationEsxiConsole) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *VirtualizationEsxiConsole) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *VirtualizationEsxiConsole) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### SetDeviceRegistrationNil

`func (o *VirtualizationEsxiConsole) SetDeviceRegistrationNil(b bool)`

 SetDeviceRegistrationNil sets the value for DeviceRegistration to be an explicit nil

### UnsetDeviceRegistration
`func (o *VirtualizationEsxiConsole) UnsetDeviceRegistration()`

UnsetDeviceRegistration ensures that no value is present for DeviceRegistration, not even an explicit nil
### GetHost

`func (o *VirtualizationEsxiConsole) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationEsxiConsole) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationEsxiConsole) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationEsxiConsole) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *VirtualizationEsxiConsole) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *VirtualizationEsxiConsole) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


