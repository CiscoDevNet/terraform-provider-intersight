# PolicyAbstractConfigResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedTime** | Pointer to **string** | The completed time of the task in installer. | [optional] 
**Context** | Pointer to [**PolicyConfigResultContext**](policy.ConfigResultContext.md) |  | [optional] 
**Message** | Pointer to **string** | Localized message based on the locale setting of the user&#39;s context. | [optional] 
**OwnerId** | Pointer to **string** | The identifier of the object that owns the result message. The owner ID is used to correlate a given result entry to a task or entity. For example, a config result entry that describes the result of a workflow task may have the task&#39;s instance ID as the owner. | [optional] 
**State** | Pointer to **string** | Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**Type** | Pointer to **string** | Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config. | [optional] 

## Methods

### NewPolicyAbstractConfigResultEntry

`func NewPolicyAbstractConfigResultEntry() *PolicyAbstractConfigResultEntry`

NewPolicyAbstractConfigResultEntry instantiates a new PolicyAbstractConfigResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigResultEntryWithDefaults

`func NewPolicyAbstractConfigResultEntryWithDefaults() *PolicyAbstractConfigResultEntry`

NewPolicyAbstractConfigResultEntryWithDefaults instantiates a new PolicyAbstractConfigResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedTime

`func (o *PolicyAbstractConfigResultEntry) GetCompletedTime() string`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *PolicyAbstractConfigResultEntry) GetCompletedTimeOk() (*string, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *PolicyAbstractConfigResultEntry) SetCompletedTime(v string)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *PolicyAbstractConfigResultEntry) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetContext

`func (o *PolicyAbstractConfigResultEntry) GetContext() PolicyConfigResultContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PolicyAbstractConfigResultEntry) GetContextOk() (*PolicyConfigResultContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PolicyAbstractConfigResultEntry) SetContext(v PolicyConfigResultContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PolicyAbstractConfigResultEntry) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessage

`func (o *PolicyAbstractConfigResultEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyAbstractConfigResultEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyAbstractConfigResultEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PolicyAbstractConfigResultEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOwnerId

`func (o *PolicyAbstractConfigResultEntry) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PolicyAbstractConfigResultEntry) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PolicyAbstractConfigResultEntry) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PolicyAbstractConfigResultEntry) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetState

`func (o *PolicyAbstractConfigResultEntry) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyAbstractConfigResultEntry) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyAbstractConfigResultEntry) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PolicyAbstractConfigResultEntry) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *PolicyAbstractConfigResultEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyAbstractConfigResultEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyAbstractConfigResultEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyAbstractConfigResultEntry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


