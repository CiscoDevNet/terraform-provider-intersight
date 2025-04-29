# NotificationApplianceUpgradeMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.ApplianceUpgradeMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.ApplianceUpgradeMoCondition"]
**OdataFilter** | Pointer to **string** | Odata filter string managed internally. It is built with specific ObjectType properties. | [optional] [readonly] 
**Statuses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationApplianceUpgradeMoCondition

`func NewNotificationApplianceUpgradeMoCondition(classId string, objectType string, ) *NotificationApplianceUpgradeMoCondition`

NewNotificationApplianceUpgradeMoCondition instantiates a new NotificationApplianceUpgradeMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationApplianceUpgradeMoConditionWithDefaults

`func NewNotificationApplianceUpgradeMoConditionWithDefaults() *NotificationApplianceUpgradeMoCondition`

NewNotificationApplianceUpgradeMoConditionWithDefaults instantiates a new NotificationApplianceUpgradeMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationApplianceUpgradeMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationApplianceUpgradeMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationApplianceUpgradeMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationApplianceUpgradeMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationApplianceUpgradeMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationApplianceUpgradeMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOdataFilter

`func (o *NotificationApplianceUpgradeMoCondition) GetOdataFilter() string`

GetOdataFilter returns the OdataFilter field if non-nil, zero value otherwise.

### GetOdataFilterOk

`func (o *NotificationApplianceUpgradeMoCondition) GetOdataFilterOk() (*string, bool)`

GetOdataFilterOk returns a tuple with the OdataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilter

`func (o *NotificationApplianceUpgradeMoCondition) SetOdataFilter(v string)`

SetOdataFilter sets OdataFilter field to given value.

### HasOdataFilter

`func (o *NotificationApplianceUpgradeMoCondition) HasOdataFilter() bool`

HasOdataFilter returns a boolean if a field has been set.

### GetStatuses

`func (o *NotificationApplianceUpgradeMoCondition) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *NotificationApplianceUpgradeMoCondition) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *NotificationApplianceUpgradeMoCondition) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *NotificationApplianceUpgradeMoCondition) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *NotificationApplianceUpgradeMoCondition) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *NotificationApplianceUpgradeMoCondition) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


