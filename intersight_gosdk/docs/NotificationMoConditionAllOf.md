# NotificationMoConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.MoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.MoCondition"]
**OdataFilter** | Pointer to **string** | Odata filter string managed internally. It is built with specific ObjectType properties. | [optional] 

## Methods

### NewNotificationMoConditionAllOf

`func NewNotificationMoConditionAllOf(classId string, objectType string, ) *NotificationMoConditionAllOf`

NewNotificationMoConditionAllOf instantiates a new NotificationMoConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMoConditionAllOfWithDefaults

`func NewNotificationMoConditionAllOfWithDefaults() *NotificationMoConditionAllOf`

NewNotificationMoConditionAllOfWithDefaults instantiates a new NotificationMoConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationMoConditionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationMoConditionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationMoConditionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationMoConditionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationMoConditionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationMoConditionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOdataFilter

`func (o *NotificationMoConditionAllOf) GetOdataFilter() string`

GetOdataFilter returns the OdataFilter field if non-nil, zero value otherwise.

### GetOdataFilterOk

`func (o *NotificationMoConditionAllOf) GetOdataFilterOk() (*string, bool)`

GetOdataFilterOk returns a tuple with the OdataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilter

`func (o *NotificationMoConditionAllOf) SetOdataFilter(v string)`

SetOdataFilter sets OdataFilter field to given value.

### HasOdataFilter

`func (o *NotificationMoConditionAllOf) HasOdataFilter() bool`

HasOdataFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


