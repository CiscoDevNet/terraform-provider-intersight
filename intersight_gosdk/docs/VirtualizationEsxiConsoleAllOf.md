# VirtualizationEsxiConsoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiConsole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiConsole"]
**StreamId** | Pointer to **string** | The stream ID of the host console session opened. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](VirtualizationVmwareHostRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiConsoleAllOf

`func NewVirtualizationEsxiConsoleAllOf(classId string, objectType string, ) *VirtualizationEsxiConsoleAllOf`

NewVirtualizationEsxiConsoleAllOf instantiates a new VirtualizationEsxiConsoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiConsoleAllOfWithDefaults

`func NewVirtualizationEsxiConsoleAllOfWithDefaults() *VirtualizationEsxiConsoleAllOf`

NewVirtualizationEsxiConsoleAllOfWithDefaults instantiates a new VirtualizationEsxiConsoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiConsoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiConsoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiConsoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiConsoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiConsoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiConsoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStreamId

`func (o *VirtualizationEsxiConsoleAllOf) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *VirtualizationEsxiConsoleAllOf) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *VirtualizationEsxiConsoleAllOf) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *VirtualizationEsxiConsoleAllOf) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *VirtualizationEsxiConsoleAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *VirtualizationEsxiConsoleAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *VirtualizationEsxiConsoleAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *VirtualizationEsxiConsoleAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationEsxiConsoleAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationEsxiConsoleAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationEsxiConsoleAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationEsxiConsoleAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


