# ServicenowIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicenow.Incident"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicenow.Incident"]
**Approval** | Pointer to **string** | The approver property of Incident. | [optional] 
**Category** | Pointer to **string** | Category property for Incident. | [optional] 
**Comments** | Pointer to **string** | Comments property on Incident. | [optional] 
**CreatedBy** | Pointer to **string** | Creator property of Incident. | [optional] 
**CreatedOn** | Pointer to **string** | Incident create date property. | [optional] 
**Description** | Pointer to **string** | Description for Incident. | [optional] 
**DueDate** | Pointer to **string** | Due date property for Incident. | [optional] 
**ExpectedStart** | Pointer to **string** | Expected start date for Incident. | [optional] 
**Impact** | Pointer to **string** | Impact property for Incident. | [optional] 
**IncidentState** | Pointer to **string** | State property of the Incident. | [optional] 
**OpenedBy** | Pointer to **string** | Assigned to value for Incident. | [optional] 
**Priority** | Pointer to **string** | Priority property for Incident. | [optional] 
**Risk** | Pointer to **string** | The risk property for Incident. | [optional] 
**Severity** | Pointer to **string** | Severity property of the Incident. | [optional] 
**ShortDescription** | Pointer to **string** | Short Description for Incident. | [optional] 
**SysId** | Pointer to **string** | System Id property for Incident. | [optional] 
**TaskEffectiveNumber** | Pointer to **string** | Task Effective Number for Incident. | [optional] 
**UpdatedBy** | Pointer to **string** | Last update by on Incident. | [optional] 
**Urgency** | Pointer to **string** | Urgency property of the Incident. | [optional] 
**WorkEnd** | Pointer to **string** | Work end date for Incident. | [optional] 
**WorkStart** | Pointer to **string** | Work start date for Incident. | [optional] 

## Methods

### NewServicenowIncident

`func NewServicenowIncident(classId string, objectType string, ) *ServicenowIncident`

NewServicenowIncident instantiates a new ServicenowIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicenowIncidentWithDefaults

`func NewServicenowIncidentWithDefaults() *ServicenowIncident`

NewServicenowIncidentWithDefaults instantiates a new ServicenowIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicenowIncident) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicenowIncident) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicenowIncident) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicenowIncident) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicenowIncident) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicenowIncident) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApproval

`func (o *ServicenowIncident) GetApproval() string`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *ServicenowIncident) GetApprovalOk() (*string, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *ServicenowIncident) SetApproval(v string)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *ServicenowIncident) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetCategory

`func (o *ServicenowIncident) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServicenowIncident) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServicenowIncident) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ServicenowIncident) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetComments

`func (o *ServicenowIncident) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ServicenowIncident) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ServicenowIncident) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ServicenowIncident) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ServicenowIncident) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServicenowIncident) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServicenowIncident) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServicenowIncident) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ServicenowIncident) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ServicenowIncident) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ServicenowIncident) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ServicenowIncident) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDescription

`func (o *ServicenowIncident) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicenowIncident) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicenowIncident) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicenowIncident) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *ServicenowIncident) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ServicenowIncident) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ServicenowIncident) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ServicenowIncident) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExpectedStart

`func (o *ServicenowIncident) GetExpectedStart() string`

GetExpectedStart returns the ExpectedStart field if non-nil, zero value otherwise.

### GetExpectedStartOk

`func (o *ServicenowIncident) GetExpectedStartOk() (*string, bool)`

GetExpectedStartOk returns a tuple with the ExpectedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStart

`func (o *ServicenowIncident) SetExpectedStart(v string)`

SetExpectedStart sets ExpectedStart field to given value.

### HasExpectedStart

`func (o *ServicenowIncident) HasExpectedStart() bool`

HasExpectedStart returns a boolean if a field has been set.

### GetImpact

`func (o *ServicenowIncident) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ServicenowIncident) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ServicenowIncident) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ServicenowIncident) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetIncidentState

`func (o *ServicenowIncident) GetIncidentState() string`

GetIncidentState returns the IncidentState field if non-nil, zero value otherwise.

### GetIncidentStateOk

`func (o *ServicenowIncident) GetIncidentStateOk() (*string, bool)`

GetIncidentStateOk returns a tuple with the IncidentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentState

`func (o *ServicenowIncident) SetIncidentState(v string)`

SetIncidentState sets IncidentState field to given value.

### HasIncidentState

`func (o *ServicenowIncident) HasIncidentState() bool`

HasIncidentState returns a boolean if a field has been set.

### GetOpenedBy

`func (o *ServicenowIncident) GetOpenedBy() string`

GetOpenedBy returns the OpenedBy field if non-nil, zero value otherwise.

### GetOpenedByOk

`func (o *ServicenowIncident) GetOpenedByOk() (*string, bool)`

GetOpenedByOk returns a tuple with the OpenedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedBy

`func (o *ServicenowIncident) SetOpenedBy(v string)`

SetOpenedBy sets OpenedBy field to given value.

### HasOpenedBy

`func (o *ServicenowIncident) HasOpenedBy() bool`

HasOpenedBy returns a boolean if a field has been set.

### GetPriority

`func (o *ServicenowIncident) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ServicenowIncident) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ServicenowIncident) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ServicenowIncident) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRisk

`func (o *ServicenowIncident) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *ServicenowIncident) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *ServicenowIncident) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *ServicenowIncident) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSeverity

`func (o *ServicenowIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServicenowIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServicenowIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServicenowIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShortDescription

`func (o *ServicenowIncident) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ServicenowIncident) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ServicenowIncident) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ServicenowIncident) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSysId

`func (o *ServicenowIncident) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *ServicenowIncident) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *ServicenowIncident) SetSysId(v string)`

SetSysId sets SysId field to given value.

### HasSysId

`func (o *ServicenowIncident) HasSysId() bool`

HasSysId returns a boolean if a field has been set.

### GetTaskEffectiveNumber

`func (o *ServicenowIncident) GetTaskEffectiveNumber() string`

GetTaskEffectiveNumber returns the TaskEffectiveNumber field if non-nil, zero value otherwise.

### GetTaskEffectiveNumberOk

`func (o *ServicenowIncident) GetTaskEffectiveNumberOk() (*string, bool)`

GetTaskEffectiveNumberOk returns a tuple with the TaskEffectiveNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskEffectiveNumber

`func (o *ServicenowIncident) SetTaskEffectiveNumber(v string)`

SetTaskEffectiveNumber sets TaskEffectiveNumber field to given value.

### HasTaskEffectiveNumber

`func (o *ServicenowIncident) HasTaskEffectiveNumber() bool`

HasTaskEffectiveNumber returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ServicenowIncident) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServicenowIncident) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServicenowIncident) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServicenowIncident) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUrgency

`func (o *ServicenowIncident) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *ServicenowIncident) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *ServicenowIncident) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *ServicenowIncident) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.

### GetWorkEnd

`func (o *ServicenowIncident) GetWorkEnd() string`

GetWorkEnd returns the WorkEnd field if non-nil, zero value otherwise.

### GetWorkEndOk

`func (o *ServicenowIncident) GetWorkEndOk() (*string, bool)`

GetWorkEndOk returns a tuple with the WorkEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkEnd

`func (o *ServicenowIncident) SetWorkEnd(v string)`

SetWorkEnd sets WorkEnd field to given value.

### HasWorkEnd

`func (o *ServicenowIncident) HasWorkEnd() bool`

HasWorkEnd returns a boolean if a field has been set.

### GetWorkStart

`func (o *ServicenowIncident) GetWorkStart() string`

GetWorkStart returns the WorkStart field if non-nil, zero value otherwise.

### GetWorkStartOk

`func (o *ServicenowIncident) GetWorkStartOk() (*string, bool)`

GetWorkStartOk returns a tuple with the WorkStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkStart

`func (o *ServicenowIncident) SetWorkStart(v string)`

SetWorkStart sets WorkStart field to given value.

### HasWorkStart

`func (o *ServicenowIncident) HasWorkStart() bool`

HasWorkStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


