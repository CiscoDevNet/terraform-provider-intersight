# ApplianceFileSystemTelemetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileSystemTelemetry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileSystemTelemetry"]
**DiskOrder** | Pointer to **string** | Symbolic order identifier of the physical disk from /dev/disk/by-order (e.g., disk1, disk2). This provides a stable, human-readable identifier for the disk that persists across reboots, unlike device names which may change. Currently, the relationship between filesystems to disks is one-to-one and this is unlikely to change in the future. | [optional] [readonly] 
**Instance** | Pointer to **string** | The instance identifier from which the metrics were collected. | [optional] [readonly] 
**Mountpoint** | Pointer to **string** | Mount point of this file system. | [optional] [readonly] 
**Usage** | Pointer to **string** | Operational status of filesystem utilization. Values are Optimal, Warning, or Critical based on configured thresholds. * &#x60;Optimal&#x60; - Resource usage is within normal operating parameters. * &#x60;Warning&#x60; - Resource usage has exceeded warning thresholds and attention may be needed. * &#x60;Critical&#x60; - Resource usage has exceeded critical thresholds and immediate attention is required. | [optional] [readonly] [default to "Optimal"]
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileSystemTelemetry

`func NewApplianceFileSystemTelemetry(classId string, objectType string, ) *ApplianceFileSystemTelemetry`

NewApplianceFileSystemTelemetry instantiates a new ApplianceFileSystemTelemetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileSystemTelemetryWithDefaults

`func NewApplianceFileSystemTelemetryWithDefaults() *ApplianceFileSystemTelemetry`

NewApplianceFileSystemTelemetryWithDefaults instantiates a new ApplianceFileSystemTelemetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileSystemTelemetry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileSystemTelemetry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileSystemTelemetry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileSystemTelemetry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileSystemTelemetry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileSystemTelemetry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiskOrder

`func (o *ApplianceFileSystemTelemetry) GetDiskOrder() string`

GetDiskOrder returns the DiskOrder field if non-nil, zero value otherwise.

### GetDiskOrderOk

`func (o *ApplianceFileSystemTelemetry) GetDiskOrderOk() (*string, bool)`

GetDiskOrderOk returns a tuple with the DiskOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskOrder

`func (o *ApplianceFileSystemTelemetry) SetDiskOrder(v string)`

SetDiskOrder sets DiskOrder field to given value.

### HasDiskOrder

`func (o *ApplianceFileSystemTelemetry) HasDiskOrder() bool`

HasDiskOrder returns a boolean if a field has been set.

### GetInstance

`func (o *ApplianceFileSystemTelemetry) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ApplianceFileSystemTelemetry) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ApplianceFileSystemTelemetry) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ApplianceFileSystemTelemetry) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetMountpoint

`func (o *ApplianceFileSystemTelemetry) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *ApplianceFileSystemTelemetry) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *ApplianceFileSystemTelemetry) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *ApplianceFileSystemTelemetry) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetUsage

`func (o *ApplianceFileSystemTelemetry) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApplianceFileSystemTelemetry) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApplianceFileSystemTelemetry) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ApplianceFileSystemTelemetry) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceFileSystemTelemetry) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceFileSystemTelemetry) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceFileSystemTelemetry) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceFileSystemTelemetry) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceFileSystemTelemetry) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceFileSystemTelemetry) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


