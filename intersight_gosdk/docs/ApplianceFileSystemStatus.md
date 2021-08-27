# ApplianceFileSystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileSystemStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileSystemStatus"]
**Capacity** | Pointer to **int64** | Capacity of the file system in bytes. | [optional] 
**Mountpoint** | Pointer to **string** | Mount point of this file system. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - Operational status of the Intersight Appliance entity is Unknown. * &#x60;Operational&#x60; - Operational status of the Intersight Appliance entity is Operational. * &#x60;Impaired&#x60; - Operational status of the Intersight Appliance entity is Impaired. * &#x60;AttentionNeeded&#x60; - Operational status of the Intersight Appliance entity is AttentionNeeded. | [optional] [readonly] [default to "Unknown"]
**StatusChecks** | Pointer to [**[]ApplianceStatusCheck**](ApplianceStatusCheck.md) |  | [optional] 
**Usage** | Pointer to **float32** | Percentage of the file system capacity currently in use. | [optional] [readonly] 
**NodeStatus** | Pointer to [**ApplianceNodeStatusRelationship**](ApplianceNodeStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileSystemStatus

`func NewApplianceFileSystemStatus(classId string, objectType string, ) *ApplianceFileSystemStatus`

NewApplianceFileSystemStatus instantiates a new ApplianceFileSystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileSystemStatusWithDefaults

`func NewApplianceFileSystemStatusWithDefaults() *ApplianceFileSystemStatus`

NewApplianceFileSystemStatusWithDefaults instantiates a new ApplianceFileSystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileSystemStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileSystemStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileSystemStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileSystemStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileSystemStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileSystemStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *ApplianceFileSystemStatus) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ApplianceFileSystemStatus) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ApplianceFileSystemStatus) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ApplianceFileSystemStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMountpoint

`func (o *ApplianceFileSystemStatus) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *ApplianceFileSystemStatus) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *ApplianceFileSystemStatus) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *ApplianceFileSystemStatus) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceFileSystemStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceFileSystemStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceFileSystemStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceFileSystemStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetStatusChecks

`func (o *ApplianceFileSystemStatus) GetStatusChecks() []ApplianceStatusCheck`

GetStatusChecks returns the StatusChecks field if non-nil, zero value otherwise.

### GetStatusChecksOk

`func (o *ApplianceFileSystemStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool)`

GetStatusChecksOk returns a tuple with the StatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecks

`func (o *ApplianceFileSystemStatus) SetStatusChecks(v []ApplianceStatusCheck)`

SetStatusChecks sets StatusChecks field to given value.

### HasStatusChecks

`func (o *ApplianceFileSystemStatus) HasStatusChecks() bool`

HasStatusChecks returns a boolean if a field has been set.

### SetStatusChecksNil

`func (o *ApplianceFileSystemStatus) SetStatusChecksNil(b bool)`

 SetStatusChecksNil sets the value for StatusChecks to be an explicit nil

### UnsetStatusChecks
`func (o *ApplianceFileSystemStatus) UnsetStatusChecks()`

UnsetStatusChecks ensures that no value is present for StatusChecks, not even an explicit nil
### GetUsage

`func (o *ApplianceFileSystemStatus) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApplianceFileSystemStatus) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApplianceFileSystemStatus) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ApplianceFileSystemStatus) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetNodeStatus

`func (o *ApplianceFileSystemStatus) GetNodeStatus() ApplianceNodeStatusRelationship`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *ApplianceFileSystemStatus) GetNodeStatusOk() (*ApplianceNodeStatusRelationship, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *ApplianceFileSystemStatus) SetNodeStatus(v ApplianceNodeStatusRelationship)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *ApplianceFileSystemStatus) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


