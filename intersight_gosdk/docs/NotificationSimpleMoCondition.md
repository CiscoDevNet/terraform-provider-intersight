# NotificationSimpleMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.SimpleMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.SimpleMoCondition"]
**Filter** | Pointer to [**NullableNotificationSimpleFilter**](NotificationSimpleFilter.md) |  | [optional] 
**OdataFilter** | Pointer to **string** | Odata filter string managed internally. It is built with specific ObjectType properties. | [optional] [readonly] 

## Methods

### NewNotificationSimpleMoCondition

`func NewNotificationSimpleMoCondition(classId string, objectType string, ) *NotificationSimpleMoCondition`

NewNotificationSimpleMoCondition instantiates a new NotificationSimpleMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSimpleMoConditionWithDefaults

`func NewNotificationSimpleMoConditionWithDefaults() *NotificationSimpleMoCondition`

NewNotificationSimpleMoConditionWithDefaults instantiates a new NotificationSimpleMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSimpleMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSimpleMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSimpleMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSimpleMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSimpleMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSimpleMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilter

`func (o *NotificationSimpleMoCondition) GetFilter() NotificationSimpleFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NotificationSimpleMoCondition) GetFilterOk() (*NotificationSimpleFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NotificationSimpleMoCondition) SetFilter(v NotificationSimpleFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *NotificationSimpleMoCondition) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *NotificationSimpleMoCondition) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *NotificationSimpleMoCondition) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOdataFilter

`func (o *NotificationSimpleMoCondition) GetOdataFilter() string`

GetOdataFilter returns the OdataFilter field if non-nil, zero value otherwise.

### GetOdataFilterOk

`func (o *NotificationSimpleMoCondition) GetOdataFilterOk() (*string, bool)`

GetOdataFilterOk returns a tuple with the OdataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilter

`func (o *NotificationSimpleMoCondition) SetOdataFilter(v string)`

SetOdataFilter sets OdataFilter field to given value.

### HasOdataFilter

`func (o *NotificationSimpleMoCondition) HasOdataFilter() bool`

HasOdataFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


