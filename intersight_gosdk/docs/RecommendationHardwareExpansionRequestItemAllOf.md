# RecommendationHardwareExpansionRequestItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.HardwareExpansionRequestItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.HardwareExpansionRequestItem"]
**ItemType** | Pointer to **string** | Type of hardware item for which the capacity increase is requested by the user. For example, CPU. * &#x60;None&#x60; - The Enum value None represents that no value was set for the hardware type. * &#x60;CPU&#x60; - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU. * &#x60;Memory&#x60; - The Enum value Memory represents that the hardware type requested for expansion is a memory unit. * &#x60;Storage&#x60; - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives. | [optional] [default to "None"]
**MaxValue** | Pointer to **float32** | The maximum value allowed for expansion for the hardware type on the referred registered device. | [optional] 
**MaxValueUnit** | Pointer to **string** | Unit type for the maximum value of the resource. For example, TB, GB, MB. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. * &#x60;GB&#x60; - The Enum value GB represents that the measurement unit is in gigabytes. * &#x60;MHz&#x60; - The Enum value MHz represents that the measurement unit is in megahertz. * &#x60;GHz&#x60; - The Enum value GHz represents that the measurement unit is in gigahertz. * &#x60;Percentage&#x60; - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. | [optional] [default to "TB"]
**UnitType** | Pointer to **string** | Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. * &#x60;GB&#x60; - The Enum value GB represents that the measurement unit is in gigabytes. * &#x60;MHz&#x60; - The Enum value MHz represents that the measurement unit is in megahertz. * &#x60;GHz&#x60; - The Enum value GHz represents that the measurement unit is in gigahertz. * &#x60;Percentage&#x60; - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. | [optional] [default to "TB"]
**Value** | Pointer to **float32** | Value of the expansion request which can be absolute value or percentage increase. | [optional] 
**ExpansionRequest** | Pointer to [**RecommendationHardwareExpansionRequestRelationship**](RecommendationHardwareExpansionRequestRelationship.md) |  | [optional] 

## Methods

### NewRecommendationHardwareExpansionRequestItemAllOf

`func NewRecommendationHardwareExpansionRequestItemAllOf(classId string, objectType string, ) *RecommendationHardwareExpansionRequestItemAllOf`

NewRecommendationHardwareExpansionRequestItemAllOf instantiates a new RecommendationHardwareExpansionRequestItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationHardwareExpansionRequestItemAllOfWithDefaults

`func NewRecommendationHardwareExpansionRequestItemAllOfWithDefaults() *RecommendationHardwareExpansionRequestItemAllOf`

NewRecommendationHardwareExpansionRequestItemAllOfWithDefaults instantiates a new RecommendationHardwareExpansionRequestItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMaxValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMaxValueUnit

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetMaxValueUnit() string`

GetMaxValueUnit returns the MaxValueUnit field if non-nil, zero value otherwise.

### GetMaxValueUnitOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetMaxValueUnitOk() (*string, bool)`

GetMaxValueUnitOk returns a tuple with the MaxValueUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValueUnit

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetMaxValueUnit(v string)`

SetMaxValueUnit sets MaxValueUnit field to given value.

### HasMaxValueUnit

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasMaxValueUnit() bool`

HasMaxValueUnit returns a boolean if a field has been set.

### GetUnitType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.

### HasUnitType

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasUnitType() bool`

HasUnitType returns a boolean if a field has been set.

### GetValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExpansionRequest

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetExpansionRequest() RecommendationHardwareExpansionRequestRelationship`

GetExpansionRequest returns the ExpansionRequest field if non-nil, zero value otherwise.

### GetExpansionRequestOk

`func (o *RecommendationHardwareExpansionRequestItemAllOf) GetExpansionRequestOk() (*RecommendationHardwareExpansionRequestRelationship, bool)`

GetExpansionRequestOk returns a tuple with the ExpansionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansionRequest

`func (o *RecommendationHardwareExpansionRequestItemAllOf) SetExpansionRequest(v RecommendationHardwareExpansionRequestRelationship)`

SetExpansionRequest sets ExpansionRequest field to given value.

### HasExpansionRequest

`func (o *RecommendationHardwareExpansionRequestItemAllOf) HasExpansionRequest() bool`

HasExpansionRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


