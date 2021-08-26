# RecommendationPhysicalItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.PhysicalItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.PhysicalItem"]
**Capacity** | Pointer to **int64** | Capacity of the physical entity added. | [optional] [readonly] 
**Count** | Pointer to **int64** | Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation. | [optional] [readonly] 
**IsNew** | Pointer to **bool** | If the PhysicalItem is new, this is set to true, else false. | [optional] [readonly] 
**MaxCount** | Pointer to **int64** | Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation. | [optional] [readonly] 
**Model** | Pointer to **string** | Model of the recommended physical device which is externally identifiable. | [optional] [readonly] 
**SourceMoid** | Pointer to **string** | Moid of the managed object which represents the existing physical entity. | [optional] [readonly] 
**Unit** | Pointer to **string** | Unit of the new capacity. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. | [optional] [readonly] [default to "TB"]
**Uuid** | Pointer to **string** | Uuid of the recommended physical device. | [optional] [readonly] 
**CapacityRunway** | Pointer to [**RecommendationCapacityRunwayRelationship**](RecommendationCapacityRunwayRelationship.md) |  | [optional] 
**PhysicalItem** | Pointer to [**[]RecommendationPhysicalItemRelationship**](RecommendationPhysicalItemRelationship.md) | An array of relationships to recommendationPhysicalItem resources. | [optional] [readonly] 

## Methods

### NewRecommendationPhysicalItemAllOf

`func NewRecommendationPhysicalItemAllOf(classId string, objectType string, ) *RecommendationPhysicalItemAllOf`

NewRecommendationPhysicalItemAllOf instantiates a new RecommendationPhysicalItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationPhysicalItemAllOfWithDefaults

`func NewRecommendationPhysicalItemAllOfWithDefaults() *RecommendationPhysicalItemAllOf`

NewRecommendationPhysicalItemAllOfWithDefaults instantiates a new RecommendationPhysicalItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationPhysicalItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationPhysicalItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationPhysicalItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationPhysicalItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationPhysicalItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationPhysicalItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *RecommendationPhysicalItemAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *RecommendationPhysicalItemAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *RecommendationPhysicalItemAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *RecommendationPhysicalItemAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCount

`func (o *RecommendationPhysicalItemAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecommendationPhysicalItemAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecommendationPhysicalItemAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecommendationPhysicalItemAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIsNew

`func (o *RecommendationPhysicalItemAllOf) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *RecommendationPhysicalItemAllOf) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *RecommendationPhysicalItemAllOf) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *RecommendationPhysicalItemAllOf) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetMaxCount

`func (o *RecommendationPhysicalItemAllOf) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *RecommendationPhysicalItemAllOf) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *RecommendationPhysicalItemAllOf) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *RecommendationPhysicalItemAllOf) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetModel

`func (o *RecommendationPhysicalItemAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RecommendationPhysicalItemAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RecommendationPhysicalItemAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RecommendationPhysicalItemAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSourceMoid

`func (o *RecommendationPhysicalItemAllOf) GetSourceMoid() string`

GetSourceMoid returns the SourceMoid field if non-nil, zero value otherwise.

### GetSourceMoidOk

`func (o *RecommendationPhysicalItemAllOf) GetSourceMoidOk() (*string, bool)`

GetSourceMoidOk returns a tuple with the SourceMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMoid

`func (o *RecommendationPhysicalItemAllOf) SetSourceMoid(v string)`

SetSourceMoid sets SourceMoid field to given value.

### HasSourceMoid

`func (o *RecommendationPhysicalItemAllOf) HasSourceMoid() bool`

HasSourceMoid returns a boolean if a field has been set.

### GetUnit

`func (o *RecommendationPhysicalItemAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RecommendationPhysicalItemAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RecommendationPhysicalItemAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RecommendationPhysicalItemAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUuid

`func (o *RecommendationPhysicalItemAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RecommendationPhysicalItemAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RecommendationPhysicalItemAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RecommendationPhysicalItemAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCapacityRunway

`func (o *RecommendationPhysicalItemAllOf) GetCapacityRunway() RecommendationCapacityRunwayRelationship`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *RecommendationPhysicalItemAllOf) GetCapacityRunwayOk() (*RecommendationCapacityRunwayRelationship, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *RecommendationPhysicalItemAllOf) SetCapacityRunway(v RecommendationCapacityRunwayRelationship)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *RecommendationPhysicalItemAllOf) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetPhysicalItem

`func (o *RecommendationPhysicalItemAllOf) GetPhysicalItem() []RecommendationPhysicalItemRelationship`

GetPhysicalItem returns the PhysicalItem field if non-nil, zero value otherwise.

### GetPhysicalItemOk

`func (o *RecommendationPhysicalItemAllOf) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool)`

GetPhysicalItemOk returns a tuple with the PhysicalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalItem

`func (o *RecommendationPhysicalItemAllOf) SetPhysicalItem(v []RecommendationPhysicalItemRelationship)`

SetPhysicalItem sets PhysicalItem field to given value.

### HasPhysicalItem

`func (o *RecommendationPhysicalItemAllOf) HasPhysicalItem() bool`

HasPhysicalItem returns a boolean if a field has been set.

### SetPhysicalItemNil

`func (o *RecommendationPhysicalItemAllOf) SetPhysicalItemNil(b bool)`

 SetPhysicalItemNil sets the value for PhysicalItem to be an explicit nil

### UnsetPhysicalItem
`func (o *RecommendationPhysicalItemAllOf) UnsetPhysicalItem()`

UnsetPhysicalItem ensures that no value is present for PhysicalItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


