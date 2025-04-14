# WorkflowMigrationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MigrationHistory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MigrationHistory"]
**Destination** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**MigrationActionInstance** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Source** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** | The time when the migration was performed. | [optional] 

## Methods

### NewWorkflowMigrationHistory

`func NewWorkflowMigrationHistory(classId string, objectType string, ) *WorkflowMigrationHistory`

NewWorkflowMigrationHistory instantiates a new WorkflowMigrationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMigrationHistoryWithDefaults

`func NewWorkflowMigrationHistoryWithDefaults() *WorkflowMigrationHistory`

NewWorkflowMigrationHistoryWithDefaults instantiates a new WorkflowMigrationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMigrationHistory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMigrationHistory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMigrationHistory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMigrationHistory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMigrationHistory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMigrationHistory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestination

`func (o *WorkflowMigrationHistory) GetDestination() MoMoRef`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *WorkflowMigrationHistory) GetDestinationOk() (*MoMoRef, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *WorkflowMigrationHistory) SetDestination(v MoMoRef)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *WorkflowMigrationHistory) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetMigrationActionInstance

`func (o *WorkflowMigrationHistory) GetMigrationActionInstance() MoMoRef`

GetMigrationActionInstance returns the MigrationActionInstance field if non-nil, zero value otherwise.

### GetMigrationActionInstanceOk

`func (o *WorkflowMigrationHistory) GetMigrationActionInstanceOk() (*MoMoRef, bool)`

GetMigrationActionInstanceOk returns a tuple with the MigrationActionInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationActionInstance

`func (o *WorkflowMigrationHistory) SetMigrationActionInstance(v MoMoRef)`

SetMigrationActionInstance sets MigrationActionInstance field to given value.

### HasMigrationActionInstance

`func (o *WorkflowMigrationHistory) HasMigrationActionInstance() bool`

HasMigrationActionInstance returns a boolean if a field has been set.

### GetSource

`func (o *WorkflowMigrationHistory) GetSource() MoMoRef`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WorkflowMigrationHistory) GetSourceOk() (*MoMoRef, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WorkflowMigrationHistory) SetSource(v MoMoRef)`

SetSource sets Source field to given value.

### HasSource

`func (o *WorkflowMigrationHistory) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimestamp

`func (o *WorkflowMigrationHistory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WorkflowMigrationHistory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WorkflowMigrationHistory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WorkflowMigrationHistory) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


