# StoragePureArrayAlerts

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
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 

## Methods

### NewStoragePureArrayAlerts

`func NewStoragePureArrayAlerts(classId string, objectType string, ) *StoragePureArrayAlerts`

NewStoragePureArrayAlerts instantiates a new StoragePureArrayAlerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureArrayAlertsWithDefaults

`func NewStoragePureArrayAlertsWithDefaults() *StoragePureArrayAlerts`

NewStoragePureArrayAlertsWithDefaults instantiates a new StoragePureArrayAlerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureArrayAlerts) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureArrayAlerts) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureArrayAlerts) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureArrayAlerts) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureArrayAlerts) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureArrayAlerts) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCause

`func (o *StoragePureArrayAlerts) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *StoragePureArrayAlerts) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *StoragePureArrayAlerts) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *StoragePureArrayAlerts) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetComponent

`func (o *StoragePureArrayAlerts) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *StoragePureArrayAlerts) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *StoragePureArrayAlerts) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *StoragePureArrayAlerts) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StoragePureArrayAlerts) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoragePureArrayAlerts) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoragePureArrayAlerts) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoragePureArrayAlerts) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *StoragePureArrayAlerts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureArrayAlerts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureArrayAlerts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureArrayAlerts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *StoragePureArrayAlerts) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StoragePureArrayAlerts) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StoragePureArrayAlerts) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StoragePureArrayAlerts) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureArrayAlerts) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureArrayAlerts) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureArrayAlerts) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureArrayAlerts) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureArrayAlerts) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureArrayAlerts) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


