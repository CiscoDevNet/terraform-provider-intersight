# EquipmentChassisOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ChassisOperationStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ChassisOperationStatus"]
**ConfigState** | Pointer to **string** | The configured state of the settings in the target chassis. The value is any one of Applied, Applying or Failed. Applied - The state denotes that the settings are applied successfully in the target chassis. Applying - The state denotes that the settings are being applied in the target chassis. Failed - The state denotes that the settings could not be applied in the target chassis. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**SlotId** | Pointer to **int64** | The slot id of the device within the chassis on which the chassis operation is performed. | [optional] [readonly] 
**WorkflowId** | Pointer to **string** | The workflow Id of the chassis operations workflow. | [optional] [readonly] 
**WorkflowType** | Pointer to **string** | The workflow type of the chassis operation workflow. This can be used to distinguish different chassis operations. | [optional] [readonly] 

## Methods

### NewEquipmentChassisOperationStatus

`func NewEquipmentChassisOperationStatus(classId string, objectType string, ) *EquipmentChassisOperationStatus`

NewEquipmentChassisOperationStatus instantiates a new EquipmentChassisOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentChassisOperationStatusWithDefaults

`func NewEquipmentChassisOperationStatusWithDefaults() *EquipmentChassisOperationStatus`

NewEquipmentChassisOperationStatusWithDefaults instantiates a new EquipmentChassisOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentChassisOperationStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentChassisOperationStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentChassisOperationStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentChassisOperationStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentChassisOperationStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentChassisOperationStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *EquipmentChassisOperationStatus) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentChassisOperationStatus) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentChassisOperationStatus) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentChassisOperationStatus) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentChassisOperationStatus) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentChassisOperationStatus) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentChassisOperationStatus) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentChassisOperationStatus) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *EquipmentChassisOperationStatus) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *EquipmentChassisOperationStatus) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *EquipmentChassisOperationStatus) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *EquipmentChassisOperationStatus) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *EquipmentChassisOperationStatus) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *EquipmentChassisOperationStatus) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *EquipmentChassisOperationStatus) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *EquipmentChassisOperationStatus) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


