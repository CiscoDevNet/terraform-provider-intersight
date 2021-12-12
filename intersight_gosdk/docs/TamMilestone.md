# TamMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.Milestone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.Milestone"]
**Date** | Pointer to **time.Time** | Date when the specified end-of-life milestone advisory is reached. | [optional] 
**Description** | Pointer to **string** | A description of the milestone defined by an end-of-life advisory. | [optional] 
**EndOffset** | Pointer to **int32** | Number of days (exclusive) relative to the milestone date when the milestone is considered to be not in effect. A nagative number indicates number of days ahead of the milestone date. The default is 2147483647 (0x7FFFFFFF) which means the milestone date range&#39;s upper bound is omitted. | [optional] [default to 2147483647]
**LabelHint** | Pointer to **string** | Extra hint to the type of label to be used in display in addition to severity and effective date. How to use it is at UI&#39;s descretion. * &#x60;upcoming&#x60; - This end-of-life (EOL) milestone is upcoming. The label may be changed to more urgent type such as &#39;imminent&#39; as time progress approaching effective date. * &#x60;imminent&#x60; - This end-of-life (EOL) milestone is imminent. There will be no label change before the effective date. * &#x60;past&#x60; - This end-of-life (EOL) milestone has past the effective date. | [optional] [default to "upcoming"]
**MilestoneType** | Pointer to **string** | Milestone type as defined in Cisco end-of-life (EOL) policy (https://www.cisco.com/c/en/us/products/eos-eol-policy.html) when the specified end-of-life milestone advisory is reached. * &#x60;unknown&#x60; - The type of end-of-life milestone is not defined. * &#x60;endOfSoftwareMaintenanceDate&#x60; - The last date that Cisco Engineering may release any final software maintenance releases or bug fixes. After this date, Cisco Engineering may no longer develop, repair, maintain, or test the product software and only critical security updates will be provided on this release train.  * &#x60;lastDateOfSupport&#x60; - The last date to receive service and support for the software. After this date, all support services for the software are unavailable, and the software becomes obsolete. | [optional] [default to "unknown"]
**Name** | Pointer to **string** | A milestone defined by an end-of-life advisory. | [optional] 
**StartOffset** | Pointer to **int32** | Number of days (inclusive) relative to the milestone date when the milestone is considered to be in effect. A nagative number indicates number of days ahead of the milestone date. The default is 0 which means the milestone take effect exactly on the same date as the specified milestone date. A negative value of -2147483648 (0x80000000) indicates the milestone date range&#39;s lower bound is omitted. | [optional] [default to 0]

## Methods

### NewTamMilestone

`func NewTamMilestone(classId string, objectType string, ) *TamMilestone`

NewTamMilestone instantiates a new TamMilestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamMilestoneWithDefaults

`func NewTamMilestoneWithDefaults() *TamMilestone`

NewTamMilestoneWithDefaults instantiates a new TamMilestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamMilestone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamMilestone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamMilestone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamMilestone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamMilestone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamMilestone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDate

`func (o *TamMilestone) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TamMilestone) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TamMilestone) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *TamMilestone) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *TamMilestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TamMilestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TamMilestone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TamMilestone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndOffset

`func (o *TamMilestone) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *TamMilestone) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *TamMilestone) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *TamMilestone) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetLabelHint

`func (o *TamMilestone) GetLabelHint() string`

GetLabelHint returns the LabelHint field if non-nil, zero value otherwise.

### GetLabelHintOk

`func (o *TamMilestone) GetLabelHintOk() (*string, bool)`

GetLabelHintOk returns a tuple with the LabelHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelHint

`func (o *TamMilestone) SetLabelHint(v string)`

SetLabelHint sets LabelHint field to given value.

### HasLabelHint

`func (o *TamMilestone) HasLabelHint() bool`

HasLabelHint returns a boolean if a field has been set.

### GetMilestoneType

`func (o *TamMilestone) GetMilestoneType() string`

GetMilestoneType returns the MilestoneType field if non-nil, zero value otherwise.

### GetMilestoneTypeOk

`func (o *TamMilestone) GetMilestoneTypeOk() (*string, bool)`

GetMilestoneTypeOk returns a tuple with the MilestoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneType

`func (o *TamMilestone) SetMilestoneType(v string)`

SetMilestoneType sets MilestoneType field to given value.

### HasMilestoneType

`func (o *TamMilestone) HasMilestoneType() bool`

HasMilestoneType returns a boolean if a field has been set.

### GetName

`func (o *TamMilestone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamMilestone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamMilestone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamMilestone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartOffset

`func (o *TamMilestone) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *TamMilestone) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *TamMilestone) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *TamMilestone) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


