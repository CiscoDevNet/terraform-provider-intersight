# PolicyConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigContext"]
**ConfigState** | Pointer to **string** | Indicates a profile&#39;s configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. | [optional] [readonly] 
**ControlAction** | Pointer to **string** | System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. | [optional] 
**ErrorState** | Pointer to **string** | Indicates a profile&#39;s error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). | [optional] 
**OperState** | Pointer to **string** | Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. | [optional] [readonly] 

## Methods

### NewPolicyConfigContext

`func NewPolicyConfigContext(classId string, objectType string, ) *PolicyConfigContext`

NewPolicyConfigContext instantiates a new PolicyConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigContextWithDefaults

`func NewPolicyConfigContextWithDefaults() *PolicyConfigContext`

NewPolicyConfigContextWithDefaults instantiates a new PolicyConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigContext) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigContext) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigContext) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *PolicyConfigContext) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyConfigContext) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyConfigContext) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyConfigContext) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetControlAction

`func (o *PolicyConfigContext) GetControlAction() string`

GetControlAction returns the ControlAction field if non-nil, zero value otherwise.

### GetControlActionOk

`func (o *PolicyConfigContext) GetControlActionOk() (*string, bool)`

GetControlActionOk returns a tuple with the ControlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlAction

`func (o *PolicyConfigContext) SetControlAction(v string)`

SetControlAction sets ControlAction field to given value.

### HasControlAction

`func (o *PolicyConfigContext) HasControlAction() bool`

HasControlAction returns a boolean if a field has been set.

### GetErrorState

`func (o *PolicyConfigContext) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *PolicyConfigContext) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *PolicyConfigContext) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *PolicyConfigContext) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetOperState

`func (o *PolicyConfigContext) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PolicyConfigContext) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PolicyConfigContext) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PolicyConfigContext) HasOperState() bool`

HasOperState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


