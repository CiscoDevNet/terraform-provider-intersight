# NotificationSoftwareReleaseMetaMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.SoftwareReleaseMetaMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.SoftwareReleaseMetaMoCondition"]
**ImageTypes** | Pointer to **[]string** |  | [optional] 
**OdataFilter** | Pointer to **string** | Odata filter string managed internally. It is built with specific ObjectType properties. | [optional] [readonly] 

## Methods

### NewNotificationSoftwareReleaseMetaMoCondition

`func NewNotificationSoftwareReleaseMetaMoCondition(classId string, objectType string, ) *NotificationSoftwareReleaseMetaMoCondition`

NewNotificationSoftwareReleaseMetaMoCondition instantiates a new NotificationSoftwareReleaseMetaMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSoftwareReleaseMetaMoConditionWithDefaults

`func NewNotificationSoftwareReleaseMetaMoConditionWithDefaults() *NotificationSoftwareReleaseMetaMoCondition`

NewNotificationSoftwareReleaseMetaMoConditionWithDefaults instantiates a new NotificationSoftwareReleaseMetaMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSoftwareReleaseMetaMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSoftwareReleaseMetaMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageTypes

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetImageTypes() []string`

GetImageTypes returns the ImageTypes field if non-nil, zero value otherwise.

### GetImageTypesOk

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetImageTypesOk() (*[]string, bool)`

GetImageTypesOk returns a tuple with the ImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTypes

`func (o *NotificationSoftwareReleaseMetaMoCondition) SetImageTypes(v []string)`

SetImageTypes sets ImageTypes field to given value.

### HasImageTypes

`func (o *NotificationSoftwareReleaseMetaMoCondition) HasImageTypes() bool`

HasImageTypes returns a boolean if a field has been set.

### SetImageTypesNil

`func (o *NotificationSoftwareReleaseMetaMoCondition) SetImageTypesNil(b bool)`

 SetImageTypesNil sets the value for ImageTypes to be an explicit nil

### UnsetImageTypes
`func (o *NotificationSoftwareReleaseMetaMoCondition) UnsetImageTypes()`

UnsetImageTypes ensures that no value is present for ImageTypes, not even an explicit nil
### GetOdataFilter

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetOdataFilter() string`

GetOdataFilter returns the OdataFilter field if non-nil, zero value otherwise.

### GetOdataFilterOk

`func (o *NotificationSoftwareReleaseMetaMoCondition) GetOdataFilterOk() (*string, bool)`

GetOdataFilterOk returns a tuple with the OdataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilter

`func (o *NotificationSoftwareReleaseMetaMoCondition) SetOdataFilter(v string)`

SetOdataFilter sets OdataFilter field to given value.

### HasOdataFilter

`func (o *NotificationSoftwareReleaseMetaMoCondition) HasOdataFilter() bool`

HasOdataFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


