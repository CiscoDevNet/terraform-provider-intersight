# NotificationSecurityAdvisoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.SecurityAdvisoryFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.SecurityAdvisoryFilter"]
**Severity** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationSecurityAdvisoryFilter

`func NewNotificationSecurityAdvisoryFilter(classId string, objectType string, ) *NotificationSecurityAdvisoryFilter`

NewNotificationSecurityAdvisoryFilter instantiates a new NotificationSecurityAdvisoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSecurityAdvisoryFilterWithDefaults

`func NewNotificationSecurityAdvisoryFilterWithDefaults() *NotificationSecurityAdvisoryFilter`

NewNotificationSecurityAdvisoryFilterWithDefaults instantiates a new NotificationSecurityAdvisoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSecurityAdvisoryFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSecurityAdvisoryFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSecurityAdvisoryFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSecurityAdvisoryFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSecurityAdvisoryFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSecurityAdvisoryFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSeverity

`func (o *NotificationSecurityAdvisoryFilter) GetSeverity() []string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationSecurityAdvisoryFilter) GetSeverityOk() (*[]string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationSecurityAdvisoryFilter) SetSeverity(v []string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationSecurityAdvisoryFilter) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *NotificationSecurityAdvisoryFilter) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *NotificationSecurityAdvisoryFilter) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


