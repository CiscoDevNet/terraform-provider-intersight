# VirtualizationVmwareProactiveHaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareProactiveHa"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareProactiveHa"]
**LastAcknowledgedAlarmTime** | Pointer to **time.Time** | Last recognized alarm time for a proactive HA alarm instance in a vCenter. | [optional] [readonly] 
**LastSentAlarmTime** | Pointer to **time.Time** | Time at which the last alarm was sent from cloud to the device connector. | [optional] [readonly] 
**AlarmDefinitions** | Pointer to [**[]CondAlarmDefinitionRelationship**](CondAlarmDefinitionRelationship.md) | An array of relationships to condAlarmDefinition resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareProactiveHaAllOf

`func NewVirtualizationVmwareProactiveHaAllOf(classId string, objectType string, ) *VirtualizationVmwareProactiveHaAllOf`

NewVirtualizationVmwareProactiveHaAllOf instantiates a new VirtualizationVmwareProactiveHaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareProactiveHaAllOfWithDefaults

`func NewVirtualizationVmwareProactiveHaAllOfWithDefaults() *VirtualizationVmwareProactiveHaAllOf`

NewVirtualizationVmwareProactiveHaAllOfWithDefaults instantiates a new VirtualizationVmwareProactiveHaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareProactiveHaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareProactiveHaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareProactiveHaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareProactiveHaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastAcknowledgedAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) GetLastAcknowledgedAlarmTime() time.Time`

GetLastAcknowledgedAlarmTime returns the LastAcknowledgedAlarmTime field if non-nil, zero value otherwise.

### GetLastAcknowledgedAlarmTimeOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetLastAcknowledgedAlarmTimeOk() (*time.Time, bool)`

GetLastAcknowledgedAlarmTimeOk returns a tuple with the LastAcknowledgedAlarmTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAcknowledgedAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) SetLastAcknowledgedAlarmTime(v time.Time)`

SetLastAcknowledgedAlarmTime sets LastAcknowledgedAlarmTime field to given value.

### HasLastAcknowledgedAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) HasLastAcknowledgedAlarmTime() bool`

HasLastAcknowledgedAlarmTime returns a boolean if a field has been set.

### GetLastSentAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) GetLastSentAlarmTime() time.Time`

GetLastSentAlarmTime returns the LastSentAlarmTime field if non-nil, zero value otherwise.

### GetLastSentAlarmTimeOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetLastSentAlarmTimeOk() (*time.Time, bool)`

GetLastSentAlarmTimeOk returns a tuple with the LastSentAlarmTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) SetLastSentAlarmTime(v time.Time)`

SetLastSentAlarmTime sets LastSentAlarmTime field to given value.

### HasLastSentAlarmTime

`func (o *VirtualizationVmwareProactiveHaAllOf) HasLastSentAlarmTime() bool`

HasLastSentAlarmTime returns a boolean if a field has been set.

### GetAlarmDefinitions

`func (o *VirtualizationVmwareProactiveHaAllOf) GetAlarmDefinitions() []CondAlarmDefinitionRelationship`

GetAlarmDefinitions returns the AlarmDefinitions field if non-nil, zero value otherwise.

### GetAlarmDefinitionsOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetAlarmDefinitionsOk() (*[]CondAlarmDefinitionRelationship, bool)`

GetAlarmDefinitionsOk returns a tuple with the AlarmDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmDefinitions

`func (o *VirtualizationVmwareProactiveHaAllOf) SetAlarmDefinitions(v []CondAlarmDefinitionRelationship)`

SetAlarmDefinitions sets AlarmDefinitions field to given value.

### HasAlarmDefinitions

`func (o *VirtualizationVmwareProactiveHaAllOf) HasAlarmDefinitions() bool`

HasAlarmDefinitions returns a boolean if a field has been set.

### SetAlarmDefinitionsNil

`func (o *VirtualizationVmwareProactiveHaAllOf) SetAlarmDefinitionsNil(b bool)`

 SetAlarmDefinitionsNil sets the value for AlarmDefinitions to be an explicit nil

### UnsetAlarmDefinitions
`func (o *VirtualizationVmwareProactiveHaAllOf) UnsetAlarmDefinitions()`

UnsetAlarmDefinitions ensures that no value is present for AlarmDefinitions, not even an explicit nil
### GetRegisteredDevice

`func (o *VirtualizationVmwareProactiveHaAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVmwareProactiveHaAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVmwareProactiveHaAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVmwareProactiveHaAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


