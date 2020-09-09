# PolicyConfigContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigState** | Pointer to **string** | Indicates a profile&#39;s configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Validating, Configuring, Failed. | [optional] [readonly] 
**ControlAction** | Pointer to **string** | System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. | [optional] 
**ErrorState** | Pointer to **string** | Indicates a profile&#39;s error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). | [optional] 
**OperState** | Pointer to **string** | Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. | [optional] [readonly] 

## Methods

### NewPolicyConfigContextAllOf

`func NewPolicyConfigContextAllOf() *PolicyConfigContextAllOf`

NewPolicyConfigContextAllOf instantiates a new PolicyConfigContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigContextAllOfWithDefaults

`func NewPolicyConfigContextAllOfWithDefaults() *PolicyConfigContextAllOf`

NewPolicyConfigContextAllOfWithDefaults instantiates a new PolicyConfigContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigState

`func (o *PolicyConfigContextAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyConfigContextAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyConfigContextAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyConfigContextAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetControlAction

`func (o *PolicyConfigContextAllOf) GetControlAction() string`

GetControlAction returns the ControlAction field if non-nil, zero value otherwise.

### GetControlActionOk

`func (o *PolicyConfigContextAllOf) GetControlActionOk() (*string, bool)`

GetControlActionOk returns a tuple with the ControlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlAction

`func (o *PolicyConfigContextAllOf) SetControlAction(v string)`

SetControlAction sets ControlAction field to given value.

### HasControlAction

`func (o *PolicyConfigContextAllOf) HasControlAction() bool`

HasControlAction returns a boolean if a field has been set.

### GetErrorState

`func (o *PolicyConfigContextAllOf) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *PolicyConfigContextAllOf) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *PolicyConfigContextAllOf) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *PolicyConfigContextAllOf) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetOperState

`func (o *PolicyConfigContextAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PolicyConfigContextAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PolicyConfigContextAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PolicyConfigContextAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


