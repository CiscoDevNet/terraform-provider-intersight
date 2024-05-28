# EquipmentSensor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Sensor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Sensor"]
**AlarmStatus** | Pointer to **string** | The alarm status of the sensor. * &#x60;absent&#x60; - The sensor is absent on the device. * &#x60;failed&#x60; - The sensor in the device has failed. * &#x60;normal&#x60; - The sensor status is normal. * &#x60;minor&#x60; - The sensor has crossed minor threshold. * &#x60;major&#x60; - The sensor has crossed major threshold. * &#x60;bad-asic&#x60; - The sensor is in bad-asic state. | [optional] [readonly] [default to "absent"]
**Name** | Pointer to **string** | The name or description of the sensor. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSensor

`func NewEquipmentSensor(classId string, objectType string, ) *EquipmentSensor`

NewEquipmentSensor instantiates a new EquipmentSensor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSensorWithDefaults

`func NewEquipmentSensorWithDefaults() *EquipmentSensor`

NewEquipmentSensorWithDefaults instantiates a new EquipmentSensor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSensor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSensor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSensor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSensor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSensor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSensor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmStatus

`func (o *EquipmentSensor) GetAlarmStatus() string`

GetAlarmStatus returns the AlarmStatus field if non-nil, zero value otherwise.

### GetAlarmStatusOk

`func (o *EquipmentSensor) GetAlarmStatusOk() (*string, bool)`

GetAlarmStatusOk returns a tuple with the AlarmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmStatus

`func (o *EquipmentSensor) SetAlarmStatus(v string)`

SetAlarmStatus sets AlarmStatus field to given value.

### HasAlarmStatus

`func (o *EquipmentSensor) HasAlarmStatus() bool`

HasAlarmStatus returns a boolean if a field has been set.

### GetName

`func (o *EquipmentSensor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentSensor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentSensor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentSensor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentSensor) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSensor) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSensor) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSensor) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *EquipmentSensor) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *EquipmentSensor) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSensor) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSensor) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSensor) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSensor) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentSensor) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentSensor) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


