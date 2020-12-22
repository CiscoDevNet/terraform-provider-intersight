# ForecastCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "forecast.Catalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "forecast.Catalog"]
**SchedTime** | Pointer to **string** | The time at which the regression model needs to run for all the metrics specified in catalog. | [optional] [readonly] 
**Version** | Pointer to **string** | The catalog version used in forecast configuration service. | [optional] [readonly] 
**Definition** | Pointer to [**[]ForecastDefinitionRelationship**](ForecastDefinitionRelationship.md) | An array of relationships to forecastDefinition resources. | [optional] [readonly] 

## Methods

### NewForecastCatalog

`func NewForecastCatalog(classId string, objectType string, ) *ForecastCatalog`

NewForecastCatalog instantiates a new ForecastCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastCatalogWithDefaults

`func NewForecastCatalogWithDefaults() *ForecastCatalog`

NewForecastCatalogWithDefaults instantiates a new ForecastCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ForecastCatalog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ForecastCatalog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ForecastCatalog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ForecastCatalog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ForecastCatalog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ForecastCatalog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSchedTime

`func (o *ForecastCatalog) GetSchedTime() string`

GetSchedTime returns the SchedTime field if non-nil, zero value otherwise.

### GetSchedTimeOk

`func (o *ForecastCatalog) GetSchedTimeOk() (*string, bool)`

GetSchedTimeOk returns a tuple with the SchedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedTime

`func (o *ForecastCatalog) SetSchedTime(v string)`

SetSchedTime sets SchedTime field to given value.

### HasSchedTime

`func (o *ForecastCatalog) HasSchedTime() bool`

HasSchedTime returns a boolean if a field has been set.

### GetVersion

`func (o *ForecastCatalog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ForecastCatalog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ForecastCatalog) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ForecastCatalog) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefinition

`func (o *ForecastCatalog) GetDefinition() []ForecastDefinitionRelationship`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ForecastCatalog) GetDefinitionOk() (*[]ForecastDefinitionRelationship, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ForecastCatalog) SetDefinition(v []ForecastDefinitionRelationship)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ForecastCatalog) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *ForecastCatalog) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *ForecastCatalog) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


