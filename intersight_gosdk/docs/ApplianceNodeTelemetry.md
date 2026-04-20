# ApplianceNodeTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeTelemetry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeTelemetry"]
**CpuUsage** | Pointer to **string** | CpuUsage indicates the operational status of CPU utilization on the Assist node. Values are Optimal, Warning, or Critical based on configured thresholds. * &#x60;Optimal&#x60; - Resource usage is within normal operating parameters. * &#x60;Warning&#x60; - Resource usage has exceeded warning thresholds and attention may be needed. * &#x60;Critical&#x60; - Resource usage has exceeded critical thresholds and immediate attention is required. | [optional] [readonly] [default to "Optimal"]
**Instance** | Pointer to **string** | The instance identifier from which the metrics were collected. | [optional] [readonly] 
**MemoryUsage** | Pointer to **string** | MemoryUsage indicates the operational status of memory utilization on the Assist node. Values are Optimal, Warning, or Critical based on configured thresholds. * &#x60;Optimal&#x60; - Resource usage is within normal operating parameters. * &#x60;Warning&#x60; - Resource usage has exceeded warning thresholds and attention may be needed. * &#x60;Critical&#x60; - Resource usage has exceeded critical thresholds and immediate attention is required. | [optional] [readonly] [default to "Optimal"]
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceNodeTelemetry

`func NewApplianceNodeTelemetry(classId string, objectType string, ) *ApplianceNodeTelemetry`

NewApplianceNodeTelemetry instantiates a new ApplianceNodeTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeTelemetryWithDefaults

`func NewApplianceNodeTelemetryWithDefaults() *ApplianceNodeTelemetry`

NewApplianceNodeTelemetryWithDefaults instantiates a new ApplianceNodeTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeTelemetry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeTelemetry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeTelemetry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeTelemetry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeTelemetry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeTelemetry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuUsage

`func (o *ApplianceNodeTelemetry) GetCpuUsage() string`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ApplianceNodeTelemetry) GetCpuUsageOk() (*string, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ApplianceNodeTelemetry) SetCpuUsage(v string)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ApplianceNodeTelemetry) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetInstance

`func (o *ApplianceNodeTelemetry) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ApplianceNodeTelemetry) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ApplianceNodeTelemetry) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ApplianceNodeTelemetry) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *ApplianceNodeTelemetry) GetMemoryUsage() string`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ApplianceNodeTelemetry) GetMemoryUsageOk() (*string, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ApplianceNodeTelemetry) SetMemoryUsage(v string)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *ApplianceNodeTelemetry) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceNodeTelemetry) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceNodeTelemetry) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceNodeTelemetry) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceNodeTelemetry) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceNodeTelemetry) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceNodeTelemetry) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


