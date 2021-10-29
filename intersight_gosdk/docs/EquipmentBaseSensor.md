# EquipmentBaseSensor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppSensor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.NetAppSensor"]
**ControllerName** | Pointer to **string** | The name of the storage array controller that the sensor applies to. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of a specific sensor. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the specified sensor. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the specified sensor. | [optional] [readonly] 
**Units** | Pointer to **string** | The units that correspond to the value of the sensor, if applicable. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of the specified sensor. | [optional] [readonly] 

## Methods

### NewEquipmentBaseSensor

`func NewEquipmentBaseSensor(classId string, objectType string, ) *EquipmentBaseSensor`

NewEquipmentBaseSensor instantiates a new EquipmentBaseSensor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentBaseSensorWithDefaults

`func NewEquipmentBaseSensorWithDefaults() *EquipmentBaseSensor`

NewEquipmentBaseSensorWithDefaults instantiates a new EquipmentBaseSensor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentBaseSensor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentBaseSensor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentBaseSensor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentBaseSensor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentBaseSensor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentBaseSensor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerName

`func (o *EquipmentBaseSensor) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *EquipmentBaseSensor) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *EquipmentBaseSensor) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *EquipmentBaseSensor) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetName

`func (o *EquipmentBaseSensor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentBaseSensor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentBaseSensor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentBaseSensor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *EquipmentBaseSensor) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EquipmentBaseSensor) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EquipmentBaseSensor) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EquipmentBaseSensor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *EquipmentBaseSensor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquipmentBaseSensor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquipmentBaseSensor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EquipmentBaseSensor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnits

`func (o *EquipmentBaseSensor) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *EquipmentBaseSensor) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *EquipmentBaseSensor) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *EquipmentBaseSensor) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetValue

`func (o *EquipmentBaseSensor) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EquipmentBaseSensor) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EquipmentBaseSensor) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EquipmentBaseSensor) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


