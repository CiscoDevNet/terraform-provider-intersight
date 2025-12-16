# EquipmentChassisOperationsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisOperationsDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisOperationsDetails"]
**ConfigState** | Pointer to **string** | The configured state of the settings in the target chassis. The value is any one of Applied, Applying or Failed. Applied - The state denotes that the settings are applied successfully in the target chassis. Applying - The state denotes that the settings are being applied in the target chassis. Failed - The state denotes that the settings could not be applied in the target chassis. | [optional] [readonly] 
**WorkflowType** | Pointer to **string** | The workflow type of the chassis operation workflow. This can be used to distinguish different chassis operations. | [optional] [readonly] 

## Methods

### NewEquipmentChassisOperationsDetails

`func NewEquipmentChassisOperationsDetails(classId string, objectType string, ) *EquipmentChassisOperationsDetails`

NewEquipmentChassisOperationsDetails instantiates a new EquipmentChassisOperationsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisOperationsDetailsWithDefaults

`func NewEquipmentChassisOperationsDetailsWithDefaults() *EquipmentChassisOperationsDetails`

NewEquipmentChassisOperationsDetailsWithDefaults instantiates a new EquipmentChassisOperationsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisOperationsDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisOperationsDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisOperationsDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisOperationsDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisOperationsDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisOperationsDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *EquipmentChassisOperationsDetails) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentChassisOperationsDetails) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentChassisOperationsDetails) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentChassisOperationsDetails) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetWorkflowType

`func (o *EquipmentChassisOperationsDetails) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *EquipmentChassisOperationsDetails) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *EquipmentChassisOperationsDetails) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *EquipmentChassisOperationsDetails) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


