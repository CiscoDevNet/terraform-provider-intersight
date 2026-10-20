# ServicenowChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "servicenow.ChangeRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "servicenow.ChangeRequest"]
**Approval** | Pointer to **string** | The approver of Change Request. | [optional] 
**AssignedTo** | Pointer to **string** | Assigned to value for Change Request. | [optional] 
**AssignedToDisplayValue** | Pointer to **string** | Assigned to display value for Change Request. | [optional] 
**ChangeModel** | Pointer to **string** | Change Model for Change Request. | [optional] 
**ChangeModelDisplayValue** | Pointer to **string** | Change Model display value for Change Request. | [optional] 
**ChangeRequestNumber** | Pointer to **string** | Number for Change Request. | [optional] 
**Comments** | Pointer to **string** | Comments on Change Request. | [optional] 
**CreatedBy** | Pointer to **string** | Creator of Change Request. | [optional] 
**Description** | Pointer to **string** | Description for Change Request. | [optional] 
**DueDate** | Pointer to **string** | Due date for Change Request. | [optional] 
**EndDate** | Pointer to **string** | End date for Change Request. | [optional] 
**Impact** | Pointer to **string** | Impact for Change Request. | [optional] 
**ImpactDisplayValue** | Pointer to **string** | Impact display value for Change Request. | [optional] 
**Priority** | Pointer to **string** | Priority for Change Request. | [optional] 
**PriorityDisplayValue** | Pointer to **string** | Priority display value for Change Request. | [optional] 
**Risk** | Pointer to **string** | The risk for Change Request. | [optional] 
**ShortDescription** | Pointer to **string** | Short Description for Change Request. | [optional] 
**StartDate** | Pointer to **string** | Start date for Change Request. | [optional] 
**SysId** | Pointer to **string** | Sys Id for Change Request. | [optional] 
**Type** | Pointer to **string** | The type for Change Request. | [optional] 
**UpdatedBy** | Pointer to **string** | Last update Change Request. | [optional] 

## Methods

### NewServicenowChangeRequest

`func NewServicenowChangeRequest(classId string, objectType string, ) *ServicenowChangeRequest`

NewServicenowChangeRequest instantiates a new ServicenowChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicenowChangeRequestWithDefaults

`func NewServicenowChangeRequestWithDefaults() *ServicenowChangeRequest`

NewServicenowChangeRequestWithDefaults instantiates a new ServicenowChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServicenowChangeRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServicenowChangeRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServicenowChangeRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServicenowChangeRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServicenowChangeRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServicenowChangeRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApproval

`func (o *ServicenowChangeRequest) GetApproval() string`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *ServicenowChangeRequest) GetApprovalOk() (*string, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *ServicenowChangeRequest) SetApproval(v string)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *ServicenowChangeRequest) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetAssignedTo

`func (o *ServicenowChangeRequest) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *ServicenowChangeRequest) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *ServicenowChangeRequest) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *ServicenowChangeRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedToDisplayValue

`func (o *ServicenowChangeRequest) GetAssignedToDisplayValue() string`

GetAssignedToDisplayValue returns the AssignedToDisplayValue field if non-nil, zero value otherwise.

### GetAssignedToDisplayValueOk

`func (o *ServicenowChangeRequest) GetAssignedToDisplayValueOk() (*string, bool)`

GetAssignedToDisplayValueOk returns a tuple with the AssignedToDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToDisplayValue

`func (o *ServicenowChangeRequest) SetAssignedToDisplayValue(v string)`

SetAssignedToDisplayValue sets AssignedToDisplayValue field to given value.

### HasAssignedToDisplayValue

`func (o *ServicenowChangeRequest) HasAssignedToDisplayValue() bool`

HasAssignedToDisplayValue returns a boolean if a field has been set.

### GetChangeModel

`func (o *ServicenowChangeRequest) GetChangeModel() string`

GetChangeModel returns the ChangeModel field if non-nil, zero value otherwise.

### GetChangeModelOk

`func (o *ServicenowChangeRequest) GetChangeModelOk() (*string, bool)`

GetChangeModelOk returns a tuple with the ChangeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeModel

`func (o *ServicenowChangeRequest) SetChangeModel(v string)`

SetChangeModel sets ChangeModel field to given value.

### HasChangeModel

`func (o *ServicenowChangeRequest) HasChangeModel() bool`

HasChangeModel returns a boolean if a field has been set.

### GetChangeModelDisplayValue

`func (o *ServicenowChangeRequest) GetChangeModelDisplayValue() string`

GetChangeModelDisplayValue returns the ChangeModelDisplayValue field if non-nil, zero value otherwise.

### GetChangeModelDisplayValueOk

`func (o *ServicenowChangeRequest) GetChangeModelDisplayValueOk() (*string, bool)`

GetChangeModelDisplayValueOk returns a tuple with the ChangeModelDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeModelDisplayValue

`func (o *ServicenowChangeRequest) SetChangeModelDisplayValue(v string)`

SetChangeModelDisplayValue sets ChangeModelDisplayValue field to given value.

### HasChangeModelDisplayValue

`func (o *ServicenowChangeRequest) HasChangeModelDisplayValue() bool`

HasChangeModelDisplayValue returns a boolean if a field has been set.

### GetChangeRequestNumber

`func (o *ServicenowChangeRequest) GetChangeRequestNumber() string`

GetChangeRequestNumber returns the ChangeRequestNumber field if non-nil, zero value otherwise.

### GetChangeRequestNumberOk

`func (o *ServicenowChangeRequest) GetChangeRequestNumberOk() (*string, bool)`

GetChangeRequestNumberOk returns a tuple with the ChangeRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestNumber

`func (o *ServicenowChangeRequest) SetChangeRequestNumber(v string)`

SetChangeRequestNumber sets ChangeRequestNumber field to given value.

### HasChangeRequestNumber

`func (o *ServicenowChangeRequest) HasChangeRequestNumber() bool`

HasChangeRequestNumber returns a boolean if a field has been set.

### GetComments

`func (o *ServicenowChangeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ServicenowChangeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ServicenowChangeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ServicenowChangeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ServicenowChangeRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServicenowChangeRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServicenowChangeRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServicenowChangeRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ServicenowChangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicenowChangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicenowChangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicenowChangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *ServicenowChangeRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ServicenowChangeRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ServicenowChangeRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ServicenowChangeRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ServicenowChangeRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ServicenowChangeRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ServicenowChangeRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ServicenowChangeRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetImpact

`func (o *ServicenowChangeRequest) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ServicenowChangeRequest) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ServicenowChangeRequest) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ServicenowChangeRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetImpactDisplayValue

`func (o *ServicenowChangeRequest) GetImpactDisplayValue() string`

GetImpactDisplayValue returns the ImpactDisplayValue field if non-nil, zero value otherwise.

### GetImpactDisplayValueOk

`func (o *ServicenowChangeRequest) GetImpactDisplayValueOk() (*string, bool)`

GetImpactDisplayValueOk returns a tuple with the ImpactDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDisplayValue

`func (o *ServicenowChangeRequest) SetImpactDisplayValue(v string)`

SetImpactDisplayValue sets ImpactDisplayValue field to given value.

### HasImpactDisplayValue

`func (o *ServicenowChangeRequest) HasImpactDisplayValue() bool`

HasImpactDisplayValue returns a boolean if a field has been set.

### GetPriority

`func (o *ServicenowChangeRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ServicenowChangeRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ServicenowChangeRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ServicenowChangeRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityDisplayValue

`func (o *ServicenowChangeRequest) GetPriorityDisplayValue() string`

GetPriorityDisplayValue returns the PriorityDisplayValue field if non-nil, zero value otherwise.

### GetPriorityDisplayValueOk

`func (o *ServicenowChangeRequest) GetPriorityDisplayValueOk() (*string, bool)`

GetPriorityDisplayValueOk returns a tuple with the PriorityDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityDisplayValue

`func (o *ServicenowChangeRequest) SetPriorityDisplayValue(v string)`

SetPriorityDisplayValue sets PriorityDisplayValue field to given value.

### HasPriorityDisplayValue

`func (o *ServicenowChangeRequest) HasPriorityDisplayValue() bool`

HasPriorityDisplayValue returns a boolean if a field has been set.

### GetRisk

`func (o *ServicenowChangeRequest) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *ServicenowChangeRequest) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *ServicenowChangeRequest) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *ServicenowChangeRequest) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetShortDescription

`func (o *ServicenowChangeRequest) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ServicenowChangeRequest) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ServicenowChangeRequest) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ServicenowChangeRequest) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *ServicenowChangeRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ServicenowChangeRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ServicenowChangeRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ServicenowChangeRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSysId

`func (o *ServicenowChangeRequest) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *ServicenowChangeRequest) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *ServicenowChangeRequest) SetSysId(v string)`

SetSysId sets SysId field to given value.

### HasSysId

`func (o *ServicenowChangeRequest) HasSysId() bool`

HasSysId returns a boolean if a field has been set.

### GetType

`func (o *ServicenowChangeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicenowChangeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicenowChangeRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServicenowChangeRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ServicenowChangeRequest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServicenowChangeRequest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServicenowChangeRequest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServicenowChangeRequest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


