# NiatelemetryNexusDashboardMemoryDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusDashboardMemoryDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusDashboardMemoryDetails"]
**DeviceName** | Pointer to **string** | Name of the node in Nexus Dashboard cluster. | [optional] 
**MemoryCapacity** | Pointer to **int64** | Memory capacity of a node in Nexus Dashboard. | [optional] 
**NexusDashboard** | Pointer to [**NiatelemetryNexusDashboardsRelationship**](niatelemetry.NexusDashboards.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusDashboardMemoryDetailsAllOf

`func NewNiatelemetryNexusDashboardMemoryDetailsAllOf(classId string, objectType string, ) *NiatelemetryNexusDashboardMemoryDetailsAllOf`

NewNiatelemetryNexusDashboardMemoryDetailsAllOf instantiates a new NiatelemetryNexusDashboardMemoryDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusDashboardMemoryDetailsAllOfWithDefaults

`func NewNiatelemetryNexusDashboardMemoryDetailsAllOfWithDefaults() *NiatelemetryNexusDashboardMemoryDetailsAllOf`

NewNiatelemetryNexusDashboardMemoryDetailsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardMemoryDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceName

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetMemoryCapacity() int64`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetMemoryCapacityOk() (*int64, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetMemoryCapacity(v int64)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetNexusDashboard

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetNexusDashboard() NiatelemetryNexusDashboardsRelationship`

GetNexusDashboard returns the NexusDashboard field if non-nil, zero value otherwise.

### GetNexusDashboardOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetNexusDashboardOk() (*NiatelemetryNexusDashboardsRelationship, bool)`

GetNexusDashboardOk returns a tuple with the NexusDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusDashboard

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetNexusDashboard(v NiatelemetryNexusDashboardsRelationship)`

SetNexusDashboard sets NexusDashboard field to given value.

### HasNexusDashboard

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) HasNexusDashboard() bool`

HasNexusDashboard returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNexusDashboardMemoryDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


