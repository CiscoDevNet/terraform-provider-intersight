# NotificationEolAdvisoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.EolAdvisoryFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.EolAdvisoryFilter"]
**MilestoneType** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationEolAdvisoryFilter

`func NewNotificationEolAdvisoryFilter(classId string, objectType string, ) *NotificationEolAdvisoryFilter`

NewNotificationEolAdvisoryFilter instantiates a new NotificationEolAdvisoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEolAdvisoryFilterWithDefaults

`func NewNotificationEolAdvisoryFilterWithDefaults() *NotificationEolAdvisoryFilter`

NewNotificationEolAdvisoryFilterWithDefaults instantiates a new NotificationEolAdvisoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationEolAdvisoryFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationEolAdvisoryFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationEolAdvisoryFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationEolAdvisoryFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationEolAdvisoryFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationEolAdvisoryFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMilestoneType

`func (o *NotificationEolAdvisoryFilter) GetMilestoneType() []string`

GetMilestoneType returns the MilestoneType field if non-nil, zero value otherwise.

### GetMilestoneTypeOk

`func (o *NotificationEolAdvisoryFilter) GetMilestoneTypeOk() (*[]string, bool)`

GetMilestoneTypeOk returns a tuple with the MilestoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneType

`func (o *NotificationEolAdvisoryFilter) SetMilestoneType(v []string)`

SetMilestoneType sets MilestoneType field to given value.

### HasMilestoneType

`func (o *NotificationEolAdvisoryFilter) HasMilestoneType() bool`

HasMilestoneType returns a boolean if a field has been set.

### SetMilestoneTypeNil

`func (o *NotificationEolAdvisoryFilter) SetMilestoneTypeNil(b bool)`

 SetMilestoneTypeNil sets the value for MilestoneType to be an explicit nil

### UnsetMilestoneType
`func (o *NotificationEolAdvisoryFilter) UnsetMilestoneType()`

UnsetMilestoneType ensures that no value is present for MilestoneType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


