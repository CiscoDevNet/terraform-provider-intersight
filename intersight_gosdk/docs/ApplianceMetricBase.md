# ApplianceMetricBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "appliance.MetricsIngestionUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "appliance.MetricsIngestionUtilization"]
**Attributes** | Pointer to [**MoBaseComplexType**](MoBaseComplexType.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | The end of the measurement window. | [optional] [readonly] 
**MetaField** | Pointer to **string** | Uniquely identifies the source of the metric. | [optional] [readonly] 
**Time** | Pointer to **time.Time** | The start of the measurement window. | [optional] [readonly] 

## Methods

### NewApplianceMetricBase

`func NewApplianceMetricBase(classId string, objectType string, ) *ApplianceMetricBase`

NewApplianceMetricBase instantiates a new ApplianceMetricBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceMetricBaseWithDefaults

`func NewApplianceMetricBaseWithDefaults() *ApplianceMetricBase`

NewApplianceMetricBaseWithDefaults instantiates a new ApplianceMetricBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceMetricBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceMetricBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceMetricBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceMetricBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceMetricBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceMetricBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributes

`func (o *ApplianceMetricBase) GetAttributes() MoBaseComplexType`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplianceMetricBase) GetAttributesOk() (*MoBaseComplexType, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplianceMetricBase) SetAttributes(v MoBaseComplexType)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplianceMetricBase) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceMetricBase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceMetricBase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceMetricBase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceMetricBase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMetaField

`func (o *ApplianceMetricBase) GetMetaField() string`

GetMetaField returns the MetaField field if non-nil, zero value otherwise.

### GetMetaFieldOk

`func (o *ApplianceMetricBase) GetMetaFieldOk() (*string, bool)`

GetMetaFieldOk returns a tuple with the MetaField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaField

`func (o *ApplianceMetricBase) SetMetaField(v string)`

SetMetaField sets MetaField field to given value.

### HasMetaField

`func (o *ApplianceMetricBase) HasMetaField() bool`

HasMetaField returns a boolean if a field has been set.

### GetTime

`func (o *ApplianceMetricBase) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApplianceMetricBase) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApplianceMetricBase) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApplianceMetricBase) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


