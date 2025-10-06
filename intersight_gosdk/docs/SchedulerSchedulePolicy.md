# SchedulerSchedulePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.SchedulePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.SchedulePolicy"]
**BlockDates** | Pointer to [**[]SchedulerBlockDate**](SchedulerBlockDate.md) |  | [optional] 
**EnableBlockDates** | Pointer to **bool** | Enable or disable block dates. If set to true, the schedule will not run during the block date interval. | [optional] 
**ScheduleParams** | Pointer to [**[]SchedulerBaseScheduleParams**](SchedulerBaseScheduleParams.md) |  | [optional] 
**UsageCount** | Pointer to **int64** | The number of profiles, templates and deployments that are using this policy.  This is used to determine if the policy can be deleted. If the usageCount is greater than 0, the policy cannot be deleted. | [optional] [readonly] 
**AssociatedObjects** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewSchedulerSchedulePolicy

`func NewSchedulerSchedulePolicy(classId string, objectType string, ) *SchedulerSchedulePolicy`

NewSchedulerSchedulePolicy instantiates a new SchedulerSchedulePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerSchedulePolicyWithDefaults

`func NewSchedulerSchedulePolicyWithDefaults() *SchedulerSchedulePolicy`

NewSchedulerSchedulePolicyWithDefaults instantiates a new SchedulerSchedulePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerSchedulePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerSchedulePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerSchedulePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerSchedulePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerSchedulePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerSchedulePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockDates

`func (o *SchedulerSchedulePolicy) GetBlockDates() []SchedulerBlockDate`

GetBlockDates returns the BlockDates field if non-nil, zero value otherwise.

### GetBlockDatesOk

`func (o *SchedulerSchedulePolicy) GetBlockDatesOk() (*[]SchedulerBlockDate, bool)`

GetBlockDatesOk returns a tuple with the BlockDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDates

`func (o *SchedulerSchedulePolicy) SetBlockDates(v []SchedulerBlockDate)`

SetBlockDates sets BlockDates field to given value.

### HasBlockDates

`func (o *SchedulerSchedulePolicy) HasBlockDates() bool`

HasBlockDates returns a boolean if a field has been set.

### SetBlockDatesNil

`func (o *SchedulerSchedulePolicy) SetBlockDatesNil(b bool)`

 SetBlockDatesNil sets the value for BlockDates to be an explicit nil

### UnsetBlockDates
`func (o *SchedulerSchedulePolicy) UnsetBlockDates()`

UnsetBlockDates ensures that no value is present for BlockDates, not even an explicit nil
### GetEnableBlockDates

`func (o *SchedulerSchedulePolicy) GetEnableBlockDates() bool`

GetEnableBlockDates returns the EnableBlockDates field if non-nil, zero value otherwise.

### GetEnableBlockDatesOk

`func (o *SchedulerSchedulePolicy) GetEnableBlockDatesOk() (*bool, bool)`

GetEnableBlockDatesOk returns a tuple with the EnableBlockDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlockDates

`func (o *SchedulerSchedulePolicy) SetEnableBlockDates(v bool)`

SetEnableBlockDates sets EnableBlockDates field to given value.

### HasEnableBlockDates

`func (o *SchedulerSchedulePolicy) HasEnableBlockDates() bool`

HasEnableBlockDates returns a boolean if a field has been set.

### GetScheduleParams

`func (o *SchedulerSchedulePolicy) GetScheduleParams() []SchedulerBaseScheduleParams`

GetScheduleParams returns the ScheduleParams field if non-nil, zero value otherwise.

### GetScheduleParamsOk

`func (o *SchedulerSchedulePolicy) GetScheduleParamsOk() (*[]SchedulerBaseScheduleParams, bool)`

GetScheduleParamsOk returns a tuple with the ScheduleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleParams

`func (o *SchedulerSchedulePolicy) SetScheduleParams(v []SchedulerBaseScheduleParams)`

SetScheduleParams sets ScheduleParams field to given value.

### HasScheduleParams

`func (o *SchedulerSchedulePolicy) HasScheduleParams() bool`

HasScheduleParams returns a boolean if a field has been set.

### SetScheduleParamsNil

`func (o *SchedulerSchedulePolicy) SetScheduleParamsNil(b bool)`

 SetScheduleParamsNil sets the value for ScheduleParams to be an explicit nil

### UnsetScheduleParams
`func (o *SchedulerSchedulePolicy) UnsetScheduleParams()`

UnsetScheduleParams ensures that no value is present for ScheduleParams, not even an explicit nil
### GetUsageCount

`func (o *SchedulerSchedulePolicy) GetUsageCount() int64`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *SchedulerSchedulePolicy) GetUsageCountOk() (*int64, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *SchedulerSchedulePolicy) SetUsageCount(v int64)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *SchedulerSchedulePolicy) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetAssociatedObjects

`func (o *SchedulerSchedulePolicy) GetAssociatedObjects() []MoBaseMoRelationship`

GetAssociatedObjects returns the AssociatedObjects field if non-nil, zero value otherwise.

### GetAssociatedObjectsOk

`func (o *SchedulerSchedulePolicy) GetAssociatedObjectsOk() (*[]MoBaseMoRelationship, bool)`

GetAssociatedObjectsOk returns a tuple with the AssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjects

`func (o *SchedulerSchedulePolicy) SetAssociatedObjects(v []MoBaseMoRelationship)`

SetAssociatedObjects sets AssociatedObjects field to given value.

### HasAssociatedObjects

`func (o *SchedulerSchedulePolicy) HasAssociatedObjects() bool`

HasAssociatedObjects returns a boolean if a field has been set.

### SetAssociatedObjectsNil

`func (o *SchedulerSchedulePolicy) SetAssociatedObjectsNil(b bool)`

 SetAssociatedObjectsNil sets the value for AssociatedObjects to be an explicit nil

### UnsetAssociatedObjects
`func (o *SchedulerSchedulePolicy) UnsetAssociatedObjects()`

UnsetAssociatedObjects ensures that no value is present for AssociatedObjects, not even an explicit nil
### GetOrganization

`func (o *SchedulerSchedulePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SchedulerSchedulePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SchedulerSchedulePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SchedulerSchedulePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *SchedulerSchedulePolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *SchedulerSchedulePolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


