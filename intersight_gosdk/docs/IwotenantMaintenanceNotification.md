# IwotenantMaintenanceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iwotenant.MaintenanceNotification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iwotenant.MaintenanceNotification"]
**I18nKey** | Pointer to **string** | Any i18n (internationalization) key defined the message content. If the key already exists then the  message content will be picked based on the key, otherwise provided message value will be used. | [optional] 
**IwoId** | Pointer to **string** | The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. | [optional] 
**MaintenanceStartTime** | Pointer to **time.Time** | The date/time from which the actual maintenance operations will be performed for a Customer&#39;s account. | [optional] 
**NtfnMessage** | Pointer to **string** | The notification message content is to display in the UI banner after the Customer&#39;s login to inform about planned maintenance operations on IWO. | [optional] 
**ShowFromTime** | Pointer to **time.Time** | The date/time from which the maintenance banner message will be shown to the Customer after login in to  Intersight UI. | [optional] 
**ShowUntilTime** | Pointer to **time.Time** | The date/time until which the maintenance banner message will be shown to the Customer after login into  Intersight UI. This will also be the time actual maintenance operation is planned for the finish of a  Customer&#39;s account. | [optional] 

## Methods

### NewIwotenantMaintenanceNotification

`func NewIwotenantMaintenanceNotification(classId string, objectType string, ) *IwotenantMaintenanceNotification`

NewIwotenantMaintenanceNotification instantiates a new IwotenantMaintenanceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIwotenantMaintenanceNotificationWithDefaults

`func NewIwotenantMaintenanceNotificationWithDefaults() *IwotenantMaintenanceNotification`

NewIwotenantMaintenanceNotificationWithDefaults instantiates a new IwotenantMaintenanceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IwotenantMaintenanceNotification) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IwotenantMaintenanceNotification) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IwotenantMaintenanceNotification) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IwotenantMaintenanceNotification) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IwotenantMaintenanceNotification) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IwotenantMaintenanceNotification) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetI18nKey

`func (o *IwotenantMaintenanceNotification) GetI18nKey() string`

GetI18nKey returns the I18nKey field if non-nil, zero value otherwise.

### GetI18nKeyOk

`func (o *IwotenantMaintenanceNotification) GetI18nKeyOk() (*string, bool)`

GetI18nKeyOk returns a tuple with the I18nKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18nKey

`func (o *IwotenantMaintenanceNotification) SetI18nKey(v string)`

SetI18nKey sets I18nKey field to given value.

### HasI18nKey

`func (o *IwotenantMaintenanceNotification) HasI18nKey() bool`

HasI18nKey returns a boolean if a field has been set.

### GetIwoId

`func (o *IwotenantMaintenanceNotification) GetIwoId() string`

GetIwoId returns the IwoId field if non-nil, zero value otherwise.

### GetIwoIdOk

`func (o *IwotenantMaintenanceNotification) GetIwoIdOk() (*string, bool)`

GetIwoIdOk returns a tuple with the IwoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoId

`func (o *IwotenantMaintenanceNotification) SetIwoId(v string)`

SetIwoId sets IwoId field to given value.

### HasIwoId

`func (o *IwotenantMaintenanceNotification) HasIwoId() bool`

HasIwoId returns a boolean if a field has been set.

### GetMaintenanceStartTime

`func (o *IwotenantMaintenanceNotification) GetMaintenanceStartTime() time.Time`

GetMaintenanceStartTime returns the MaintenanceStartTime field if non-nil, zero value otherwise.

### GetMaintenanceStartTimeOk

`func (o *IwotenantMaintenanceNotification) GetMaintenanceStartTimeOk() (*time.Time, bool)`

GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceStartTime

`func (o *IwotenantMaintenanceNotification) SetMaintenanceStartTime(v time.Time)`

SetMaintenanceStartTime sets MaintenanceStartTime field to given value.

### HasMaintenanceStartTime

`func (o *IwotenantMaintenanceNotification) HasMaintenanceStartTime() bool`

HasMaintenanceStartTime returns a boolean if a field has been set.

### GetNtfnMessage

`func (o *IwotenantMaintenanceNotification) GetNtfnMessage() string`

GetNtfnMessage returns the NtfnMessage field if non-nil, zero value otherwise.

### GetNtfnMessageOk

`func (o *IwotenantMaintenanceNotification) GetNtfnMessageOk() (*string, bool)`

GetNtfnMessageOk returns a tuple with the NtfnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfnMessage

`func (o *IwotenantMaintenanceNotification) SetNtfnMessage(v string)`

SetNtfnMessage sets NtfnMessage field to given value.

### HasNtfnMessage

`func (o *IwotenantMaintenanceNotification) HasNtfnMessage() bool`

HasNtfnMessage returns a boolean if a field has been set.

### GetShowFromTime

`func (o *IwotenantMaintenanceNotification) GetShowFromTime() time.Time`

GetShowFromTime returns the ShowFromTime field if non-nil, zero value otherwise.

### GetShowFromTimeOk

`func (o *IwotenantMaintenanceNotification) GetShowFromTimeOk() (*time.Time, bool)`

GetShowFromTimeOk returns a tuple with the ShowFromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFromTime

`func (o *IwotenantMaintenanceNotification) SetShowFromTime(v time.Time)`

SetShowFromTime sets ShowFromTime field to given value.

### HasShowFromTime

`func (o *IwotenantMaintenanceNotification) HasShowFromTime() bool`

HasShowFromTime returns a boolean if a field has been set.

### GetShowUntilTime

`func (o *IwotenantMaintenanceNotification) GetShowUntilTime() time.Time`

GetShowUntilTime returns the ShowUntilTime field if non-nil, zero value otherwise.

### GetShowUntilTimeOk

`func (o *IwotenantMaintenanceNotification) GetShowUntilTimeOk() (*time.Time, bool)`

GetShowUntilTimeOk returns a tuple with the ShowUntilTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUntilTime

`func (o *IwotenantMaintenanceNotification) SetShowUntilTime(v time.Time)`

SetShowUntilTime sets ShowUntilTime field to given value.

### HasShowUntilTime

`func (o *IwotenantMaintenanceNotification) HasShowUntilTime() bool`

HasShowUntilTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


