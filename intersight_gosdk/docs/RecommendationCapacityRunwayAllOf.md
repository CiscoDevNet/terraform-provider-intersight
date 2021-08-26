# RecommendationCapacityRunwayAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.CapacityRunway"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.CapacityRunway"]
**AdditionalCapacity** | Pointer to **int64** | Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster. | [optional] [readonly] 
**Period** | Pointer to **int64** | Number of months in future for which recommendation is provided for. | [optional] [readonly] 
**Runway** | Pointer to **int64** | This represents the new runway, that is the number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit after the recommended hardware is added. | [optional] [readonly] 
**TotalCapacity** | Pointer to **int64** | Total capacity of the cluster after the recommended hardware is added. | [optional] [readonly] 
**Unit** | Pointer to **string** | Unit for the new capacity. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. | [optional] [readonly] [default to "TB"]
**ForecastInstance** | Pointer to [**ForecastInstanceRelationship**](ForecastInstanceRelationship.md) |  | [optional] 
**PhysicalItem** | Pointer to [**[]RecommendationPhysicalItemRelationship**](RecommendationPhysicalItemRelationship.md) | An array of relationships to recommendationPhysicalItem resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewRecommendationCapacityRunwayAllOf

`func NewRecommendationCapacityRunwayAllOf(classId string, objectType string, ) *RecommendationCapacityRunwayAllOf`

NewRecommendationCapacityRunwayAllOf instantiates a new RecommendationCapacityRunwayAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationCapacityRunwayAllOfWithDefaults

`func NewRecommendationCapacityRunwayAllOfWithDefaults() *RecommendationCapacityRunwayAllOf`

NewRecommendationCapacityRunwayAllOfWithDefaults instantiates a new RecommendationCapacityRunwayAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationCapacityRunwayAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationCapacityRunwayAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationCapacityRunwayAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationCapacityRunwayAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationCapacityRunwayAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationCapacityRunwayAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalCapacity

`func (o *RecommendationCapacityRunwayAllOf) GetAdditionalCapacity() int64`

GetAdditionalCapacity returns the AdditionalCapacity field if non-nil, zero value otherwise.

### GetAdditionalCapacityOk

`func (o *RecommendationCapacityRunwayAllOf) GetAdditionalCapacityOk() (*int64, bool)`

GetAdditionalCapacityOk returns a tuple with the AdditionalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCapacity

`func (o *RecommendationCapacityRunwayAllOf) SetAdditionalCapacity(v int64)`

SetAdditionalCapacity sets AdditionalCapacity field to given value.

### HasAdditionalCapacity

`func (o *RecommendationCapacityRunwayAllOf) HasAdditionalCapacity() bool`

HasAdditionalCapacity returns a boolean if a field has been set.

### GetPeriod

`func (o *RecommendationCapacityRunwayAllOf) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RecommendationCapacityRunwayAllOf) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RecommendationCapacityRunwayAllOf) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RecommendationCapacityRunwayAllOf) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRunway

`func (o *RecommendationCapacityRunwayAllOf) GetRunway() int64`

GetRunway returns the Runway field if non-nil, zero value otherwise.

### GetRunwayOk

`func (o *RecommendationCapacityRunwayAllOf) GetRunwayOk() (*int64, bool)`

GetRunwayOk returns a tuple with the Runway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunway

`func (o *RecommendationCapacityRunwayAllOf) SetRunway(v int64)`

SetRunway sets Runway field to given value.

### HasRunway

`func (o *RecommendationCapacityRunwayAllOf) HasRunway() bool`

HasRunway returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *RecommendationCapacityRunwayAllOf) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *RecommendationCapacityRunwayAllOf) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *RecommendationCapacityRunwayAllOf) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *RecommendationCapacityRunwayAllOf) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUnit

`func (o *RecommendationCapacityRunwayAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RecommendationCapacityRunwayAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RecommendationCapacityRunwayAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RecommendationCapacityRunwayAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetForecastInstance

`func (o *RecommendationCapacityRunwayAllOf) GetForecastInstance() ForecastInstanceRelationship`

GetForecastInstance returns the ForecastInstance field if non-nil, zero value otherwise.

### GetForecastInstanceOk

`func (o *RecommendationCapacityRunwayAllOf) GetForecastInstanceOk() (*ForecastInstanceRelationship, bool)`

GetForecastInstanceOk returns a tuple with the ForecastInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastInstance

`func (o *RecommendationCapacityRunwayAllOf) SetForecastInstance(v ForecastInstanceRelationship)`

SetForecastInstance sets ForecastInstance field to given value.

### HasForecastInstance

`func (o *RecommendationCapacityRunwayAllOf) HasForecastInstance() bool`

HasForecastInstance returns a boolean if a field has been set.

### GetPhysicalItem

`func (o *RecommendationCapacityRunwayAllOf) GetPhysicalItem() []RecommendationPhysicalItemRelationship`

GetPhysicalItem returns the PhysicalItem field if non-nil, zero value otherwise.

### GetPhysicalItemOk

`func (o *RecommendationCapacityRunwayAllOf) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool)`

GetPhysicalItemOk returns a tuple with the PhysicalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalItem

`func (o *RecommendationCapacityRunwayAllOf) SetPhysicalItem(v []RecommendationPhysicalItemRelationship)`

SetPhysicalItem sets PhysicalItem field to given value.

### HasPhysicalItem

`func (o *RecommendationCapacityRunwayAllOf) HasPhysicalItem() bool`

HasPhysicalItem returns a boolean if a field has been set.

### SetPhysicalItemNil

`func (o *RecommendationCapacityRunwayAllOf) SetPhysicalItemNil(b bool)`

 SetPhysicalItemNil sets the value for PhysicalItem to be an explicit nil

### UnsetPhysicalItem
`func (o *RecommendationCapacityRunwayAllOf) UnsetPhysicalItem()`

UnsetPhysicalItem ensures that no value is present for PhysicalItem, not even an explicit nil
### GetRegisteredDevice

`func (o *RecommendationCapacityRunwayAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *RecommendationCapacityRunwayAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *RecommendationCapacityRunwayAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *RecommendationCapacityRunwayAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


