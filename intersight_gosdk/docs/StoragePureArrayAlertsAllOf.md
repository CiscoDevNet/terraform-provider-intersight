# StoragePureArrayAlertsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureArrayAlerts"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureArrayAlerts"]
**Cause** | Pointer to **string** | Reason for the event generation. | [optional] [readonly] 
**Component** | Pointer to **string** | Component for which the event generation. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date on which the event was generated on FlashArrays. | [optional] [readonly] 
**Name** | Pointer to **string** | ID of the alert related to the event. | [optional] [readonly] 
**Severity** | Pointer to **string** | Type of the severity of the event it could be Critical or Warning. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 

## Methods

### NewStoragePureArrayAlertsAllOf

`func NewStoragePureArrayAlertsAllOf(classId string, objectType string, ) *StoragePureArrayAlertsAllOf`

NewStoragePureArrayAlertsAllOf instantiates a new StoragePureArrayAlertsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureArrayAlertsAllOfWithDefaults

`func NewStoragePureArrayAlertsAllOfWithDefaults() *StoragePureArrayAlertsAllOf`

NewStoragePureArrayAlertsAllOfWithDefaults instantiates a new StoragePureArrayAlertsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureArrayAlertsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureArrayAlertsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureArrayAlertsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureArrayAlertsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureArrayAlertsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureArrayAlertsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCause

`func (o *StoragePureArrayAlertsAllOf) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *StoragePureArrayAlertsAllOf) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *StoragePureArrayAlertsAllOf) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *StoragePureArrayAlertsAllOf) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetComponent

`func (o *StoragePureArrayAlertsAllOf) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *StoragePureArrayAlertsAllOf) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *StoragePureArrayAlertsAllOf) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *StoragePureArrayAlertsAllOf) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StoragePureArrayAlertsAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoragePureArrayAlertsAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoragePureArrayAlertsAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoragePureArrayAlertsAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *StoragePureArrayAlertsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureArrayAlertsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureArrayAlertsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureArrayAlertsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *StoragePureArrayAlertsAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StoragePureArrayAlertsAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StoragePureArrayAlertsAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StoragePureArrayAlertsAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureArrayAlertsAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureArrayAlertsAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureArrayAlertsAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureArrayAlertsAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


