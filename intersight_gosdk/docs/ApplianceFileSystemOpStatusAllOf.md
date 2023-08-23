# ApplianceFileSystemOpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileSystemOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileSystemOpStatus"]
**Capacity** | Pointer to **int64** | Capacity of the file system in bytes. | [optional] [readonly] 
**Mountpoint** | Pointer to **string** | Mount point of this file system. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]
**Usage** | Pointer to **float32** | Percentage of the file system capacity currently in use. | [optional] [readonly] 
**NodeOpStatus** | Pointer to [**ApplianceNodeOpStatusRelationship**](ApplianceNodeOpStatusRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileSystemOpStatusAllOf

`func NewApplianceFileSystemOpStatusAllOf(classId string, objectType string, ) *ApplianceFileSystemOpStatusAllOf`

NewApplianceFileSystemOpStatusAllOf instantiates a new ApplianceFileSystemOpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileSystemOpStatusAllOfWithDefaults

`func NewApplianceFileSystemOpStatusAllOfWithDefaults() *ApplianceFileSystemOpStatusAllOf`

NewApplianceFileSystemOpStatusAllOfWithDefaults instantiates a new ApplianceFileSystemOpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileSystemOpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileSystemOpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileSystemOpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileSystemOpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *ApplianceFileSystemOpStatusAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ApplianceFileSystemOpStatusAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ApplianceFileSystemOpStatusAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMountpoint

`func (o *ApplianceFileSystemOpStatusAllOf) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *ApplianceFileSystemOpStatusAllOf) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *ApplianceFileSystemOpStatusAllOf) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceFileSystemOpStatusAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceFileSystemOpStatusAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceFileSystemOpStatusAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetUsage

`func (o *ApplianceFileSystemOpStatusAllOf) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApplianceFileSystemOpStatusAllOf) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ApplianceFileSystemOpStatusAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetNodeOpStatus

`func (o *ApplianceFileSystemOpStatusAllOf) GetNodeOpStatus() ApplianceNodeOpStatusRelationship`

GetNodeOpStatus returns the NodeOpStatus field if non-nil, zero value otherwise.

### GetNodeOpStatusOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool)`

GetNodeOpStatusOk returns a tuple with the NodeOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOpStatus

`func (o *ApplianceFileSystemOpStatusAllOf) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship)`

SetNodeOpStatus sets NodeOpStatus field to given value.

### HasNodeOpStatus

`func (o *ApplianceFileSystemOpStatusAllOf) HasNodeOpStatus() bool`

HasNodeOpStatus returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceFileSystemOpStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceFileSystemOpStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceFileSystemOpStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceFileSystemOpStatusAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


