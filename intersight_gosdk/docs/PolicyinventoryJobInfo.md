# PolicyinventoryJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionStatus** | Pointer to **string** | Execution status of the inventory job. * &#x60;Scheduled&#x60; - Inventory job is marked as scheduled. * &#x60;Completed&#x60; - Inventory job is marked as completed. * &#x60;Error&#x60; - Inventory job has errored out. | [optional] [readonly] [default to "Scheduled"]
**LastScheduledTime** | Pointer to [**time.Time**](time.Time.md) | Last scheduled time of the inventory job. | [optional] [readonly] 
**PolicyId** | Pointer to **string** | Policy ID for the inventory job. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Policy name for the inventory job. | [optional] [readonly] 

## Methods

### NewPolicyinventoryJobInfo

`func NewPolicyinventoryJobInfo() *PolicyinventoryJobInfo`

NewPolicyinventoryJobInfo instantiates a new PolicyinventoryJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyinventoryJobInfoWithDefaults

`func NewPolicyinventoryJobInfoWithDefaults() *PolicyinventoryJobInfo`

NewPolicyinventoryJobInfoWithDefaults instantiates a new PolicyinventoryJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionStatus

`func (o *PolicyinventoryJobInfo) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *PolicyinventoryJobInfo) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *PolicyinventoryJobInfo) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *PolicyinventoryJobInfo) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetLastScheduledTime

`func (o *PolicyinventoryJobInfo) GetLastScheduledTime() time.Time`

GetLastScheduledTime returns the LastScheduledTime field if non-nil, zero value otherwise.

### GetLastScheduledTimeOk

`func (o *PolicyinventoryJobInfo) GetLastScheduledTimeOk() (*time.Time, bool)`

GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledTime

`func (o *PolicyinventoryJobInfo) SetLastScheduledTime(v time.Time)`

SetLastScheduledTime sets LastScheduledTime field to given value.

### HasLastScheduledTime

`func (o *PolicyinventoryJobInfo) HasLastScheduledTime() bool`

HasLastScheduledTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyinventoryJobInfo) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyinventoryJobInfo) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyinventoryJobInfo) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyinventoryJobInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyinventoryJobInfo) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyinventoryJobInfo) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyinventoryJobInfo) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyinventoryJobInfo) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


