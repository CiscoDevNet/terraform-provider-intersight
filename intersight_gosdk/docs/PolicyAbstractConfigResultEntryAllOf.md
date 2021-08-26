# PolicyAbstractConfigResultEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CompletedTime** | Pointer to **string** | The completed time of the task in installer. | [optional] 
**Context** | Pointer to [**NullablePolicyConfigResultContext**](PolicyConfigResultContext.md) |  | [optional] 
**Message** | Pointer to **string** | Localized message based on the locale setting of the user&#39;s context. | [optional] 
**OwnerId** | Pointer to **string** | The identifier of the object that owns the result message. The owner ID is used to correlate a given result entry to a task or entity. For example, a config result entry that describes the result of a workflow task may have the task&#39;s instance ID as the owner. | [optional] 
**State** | Pointer to **string** | Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**Type** | Pointer to **string** | Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config. | [optional] 

## Methods

### NewPolicyAbstractConfigResultEntryAllOf

`func NewPolicyAbstractConfigResultEntryAllOf(classId string, objectType string, ) *PolicyAbstractConfigResultEntryAllOf`

NewPolicyAbstractConfigResultEntryAllOf instantiates a new PolicyAbstractConfigResultEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigResultEntryAllOfWithDefaults

`func NewPolicyAbstractConfigResultEntryAllOfWithDefaults() *PolicyAbstractConfigResultEntryAllOf`

NewPolicyAbstractConfigResultEntryAllOfWithDefaults instantiates a new PolicyAbstractConfigResultEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractConfigResultEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractConfigResultEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractConfigResultEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractConfigResultEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletedTime

`func (o *PolicyAbstractConfigResultEntryAllOf) GetCompletedTime() string`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetCompletedTimeOk() (*string, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *PolicyAbstractConfigResultEntryAllOf) SetCompletedTime(v string)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *PolicyAbstractConfigResultEntryAllOf) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetContext

`func (o *PolicyAbstractConfigResultEntryAllOf) GetContext() PolicyConfigResultContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetContextOk() (*PolicyConfigResultContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PolicyAbstractConfigResultEntryAllOf) SetContext(v PolicyConfigResultContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PolicyAbstractConfigResultEntryAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *PolicyAbstractConfigResultEntryAllOf) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *PolicyAbstractConfigResultEntryAllOf) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetMessage

`func (o *PolicyAbstractConfigResultEntryAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyAbstractConfigResultEntryAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PolicyAbstractConfigResultEntryAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOwnerId

`func (o *PolicyAbstractConfigResultEntryAllOf) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PolicyAbstractConfigResultEntryAllOf) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PolicyAbstractConfigResultEntryAllOf) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetState

`func (o *PolicyAbstractConfigResultEntryAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyAbstractConfigResultEntryAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PolicyAbstractConfigResultEntryAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *PolicyAbstractConfigResultEntryAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyAbstractConfigResultEntryAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyAbstractConfigResultEntryAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyAbstractConfigResultEntryAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


