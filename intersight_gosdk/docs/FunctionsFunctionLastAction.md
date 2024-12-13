# FunctionsFunctionLastAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.FunctionLastAction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.FunctionLastAction"]
**Action** | Pointer to **string** | Name of the last action performed. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Build&#x60; - Build an instance of a Function. * &#x60;Deploy&#x60; - Deploy the build Function. * &#x60;Undeploy&#x60; - Undeploy a Function that was previously successfully deployed. * &#x60;Delete&#x60; - Delete a Function that has yet to be deployed or that was recently undeployed. | [optional] [readonly] [default to "None"]
**FailureReason** | Pointer to **string** | Failure reason for the last action performed. | [optional] [readonly] 

## Methods

### NewFunctionsFunctionLastAction

`func NewFunctionsFunctionLastAction(classId string, objectType string, ) *FunctionsFunctionLastAction`

NewFunctionsFunctionLastAction instantiates a new FunctionsFunctionLastAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionLastActionWithDefaults

`func NewFunctionsFunctionLastActionWithDefaults() *FunctionsFunctionLastAction`

NewFunctionsFunctionLastActionWithDefaults instantiates a new FunctionsFunctionLastAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsFunctionLastAction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsFunctionLastAction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsFunctionLastAction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsFunctionLastAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsFunctionLastAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsFunctionLastAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FunctionsFunctionLastAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FunctionsFunctionLastAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FunctionsFunctionLastAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FunctionsFunctionLastAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFailureReason

`func (o *FunctionsFunctionLastAction) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *FunctionsFunctionLastAction) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *FunctionsFunctionLastAction) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *FunctionsFunctionLastAction) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


