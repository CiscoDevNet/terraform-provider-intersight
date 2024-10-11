# InventoryJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.JobInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.JobInfo"]
**ExecutionStatus** | Pointer to **string** | The execution status of the job scheduled. * &#x60;Unknown&#x60; - Operational job status on a managed device is not known. * &#x60;Scheduled&#x60; - Job on a managed device is scheduled. * &#x60;Completed&#x60; - Job on a managed device is completed. * &#x60;Error&#x60; - Job on a managed device is errored out. | [optional] [readonly] [default to "Unknown"]
**JobName** | Pointer to **string** | The name of the job scheduled. | [optional] [readonly] 
**LastProcessedTime** | Pointer to **time.Time** | Timestamp to indicate the last processed time for this job. | [optional] [readonly] 
**LastScheduledTime** | Pointer to **time.Time** | Last timestamp when this inventory was scheduled. | [optional] [readonly] 
**Properties** | Pointer to **[]string** |  | [optional] 
**Regex** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInventoryJobInfo

`func NewInventoryJobInfo(classId string, objectType string, ) *InventoryJobInfo`

NewInventoryJobInfo instantiates a new InventoryJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryJobInfoWithDefaults

`func NewInventoryJobInfoWithDefaults() *InventoryJobInfo`

NewInventoryJobInfoWithDefaults instantiates a new InventoryJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryJobInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryJobInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryJobInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryJobInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryJobInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryJobInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExecutionStatus

`func (o *InventoryJobInfo) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *InventoryJobInfo) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *InventoryJobInfo) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *InventoryJobInfo) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetJobName

`func (o *InventoryJobInfo) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *InventoryJobInfo) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *InventoryJobInfo) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *InventoryJobInfo) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetLastProcessedTime

`func (o *InventoryJobInfo) GetLastProcessedTime() time.Time`

GetLastProcessedTime returns the LastProcessedTime field if non-nil, zero value otherwise.

### GetLastProcessedTimeOk

`func (o *InventoryJobInfo) GetLastProcessedTimeOk() (*time.Time, bool)`

GetLastProcessedTimeOk returns a tuple with the LastProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedTime

`func (o *InventoryJobInfo) SetLastProcessedTime(v time.Time)`

SetLastProcessedTime sets LastProcessedTime field to given value.

### HasLastProcessedTime

`func (o *InventoryJobInfo) HasLastProcessedTime() bool`

HasLastProcessedTime returns a boolean if a field has been set.

### GetLastScheduledTime

`func (o *InventoryJobInfo) GetLastScheduledTime() time.Time`

GetLastScheduledTime returns the LastScheduledTime field if non-nil, zero value otherwise.

### GetLastScheduledTimeOk

`func (o *InventoryJobInfo) GetLastScheduledTimeOk() (*time.Time, bool)`

GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledTime

`func (o *InventoryJobInfo) SetLastScheduledTime(v time.Time)`

SetLastScheduledTime sets LastScheduledTime field to given value.

### HasLastScheduledTime

`func (o *InventoryJobInfo) HasLastScheduledTime() bool`

HasLastScheduledTime returns a boolean if a field has been set.

### GetProperties

`func (o *InventoryJobInfo) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InventoryJobInfo) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InventoryJobInfo) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InventoryJobInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *InventoryJobInfo) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *InventoryJobInfo) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRegex

`func (o *InventoryJobInfo) GetRegex() []string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *InventoryJobInfo) GetRegexOk() (*[]string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *InventoryJobInfo) SetRegex(v []string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *InventoryJobInfo) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *InventoryJobInfo) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *InventoryJobInfo) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


