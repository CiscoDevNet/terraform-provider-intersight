# NotificationFieldNoticeAdvisoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.FieldNoticeAdvisoryFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.FieldNoticeAdvisoryFilter"]
**Severity** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationFieldNoticeAdvisoryFilter

`func NewNotificationFieldNoticeAdvisoryFilter(classId string, objectType string, ) *NotificationFieldNoticeAdvisoryFilter`

NewNotificationFieldNoticeAdvisoryFilter instantiates a new NotificationFieldNoticeAdvisoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationFieldNoticeAdvisoryFilterWithDefaults

`func NewNotificationFieldNoticeAdvisoryFilterWithDefaults() *NotificationFieldNoticeAdvisoryFilter`

NewNotificationFieldNoticeAdvisoryFilterWithDefaults instantiates a new NotificationFieldNoticeAdvisoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationFieldNoticeAdvisoryFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationFieldNoticeAdvisoryFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationFieldNoticeAdvisoryFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationFieldNoticeAdvisoryFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationFieldNoticeAdvisoryFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationFieldNoticeAdvisoryFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSeverity

`func (o *NotificationFieldNoticeAdvisoryFilter) GetSeverity() []string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationFieldNoticeAdvisoryFilter) GetSeverityOk() (*[]string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationFieldNoticeAdvisoryFilter) SetSeverity(v []string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationFieldNoticeAdvisoryFilter) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *NotificationFieldNoticeAdvisoryFilter) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *NotificationFieldNoticeAdvisoryFilter) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


