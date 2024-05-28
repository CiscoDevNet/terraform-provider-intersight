# RecommendationPhysicalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.PhysicalItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.PhysicalItem"]
**Capacity** | Pointer to **int64** | Capacity of the physical entity added. | [optional] [readonly] 
**ConfigurationPath** | Pointer to **string** | Configuration path for the physical entity to be used when ordering it through the Cisco Commerce Workspace. | [optional] [readonly] 
**Count** | Pointer to **int64** | Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation. | [optional] [readonly] 
**IsNew** | Pointer to **bool** | If the PhysicalItem is new, this is set to true, else false. | [optional] [readonly] 
**MaxCount** | Pointer to **int64** | Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation. | [optional] [readonly] 
**Model** | Pointer to **string** | Model of the recommended physical device which is externally identifiable. | [optional] [readonly] 
**ParentMoid** | Pointer to **string** | Moid of the managed object which represents the parent physical entity. | [optional] [readonly] 
**SourceMoid** | Pointer to **string** | Moid of the managed object which represents the existing physical entity. | [optional] [readonly] 
**Unit** | Pointer to **string** | Unit of the new capacity. * &#x60;TB&#x60; - The Enum value TB represents that the measurement unit is in terabytes. * &#x60;MB&#x60; - The Enum value MB represents that the measurement unit is in megabytes. * &#x60;GB&#x60; - The Enum value GB represents that the measurement unit is in gigabytes. * &#x60;MHz&#x60; - The Enum value MHz represents that the measurement unit is in megahertz. * &#x60;GHz&#x60; - The Enum value GHz represents that the measurement unit is in gigahertz. * &#x60;Percentage&#x60; - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. | [optional] [readonly] [default to "TB"]
**Uuid** | Pointer to **string** | Uuid of the recommended physical device. | [optional] [readonly] 
**CapacityRunway** | Pointer to [**NullableRecommendationCapacityRunwayRelationship**](RecommendationCapacityRunwayRelationship.md) |  | [optional] 
**ClusterExpansion** | Pointer to [**NullableRecommendationClusterExpansionRelationship**](RecommendationClusterExpansionRelationship.md) |  | [optional] 
**PhysicalItem** | Pointer to [**[]RecommendationPhysicalItemRelationship**](RecommendationPhysicalItemRelationship.md) | An array of relationships to recommendationPhysicalItem resources. | [optional] [readonly] 

## Methods

### NewRecommendationPhysicalItem

`func NewRecommendationPhysicalItem(classId string, objectType string, ) *RecommendationPhysicalItem`

NewRecommendationPhysicalItem instantiates a new RecommendationPhysicalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationPhysicalItemWithDefaults

`func NewRecommendationPhysicalItemWithDefaults() *RecommendationPhysicalItem`

NewRecommendationPhysicalItemWithDefaults instantiates a new RecommendationPhysicalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationPhysicalItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationPhysicalItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationPhysicalItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationPhysicalItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationPhysicalItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationPhysicalItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *RecommendationPhysicalItem) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *RecommendationPhysicalItem) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *RecommendationPhysicalItem) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *RecommendationPhysicalItem) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConfigurationPath

`func (o *RecommendationPhysicalItem) GetConfigurationPath() string`

GetConfigurationPath returns the ConfigurationPath field if non-nil, zero value otherwise.

### GetConfigurationPathOk

`func (o *RecommendationPhysicalItem) GetConfigurationPathOk() (*string, bool)`

GetConfigurationPathOk returns a tuple with the ConfigurationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPath

`func (o *RecommendationPhysicalItem) SetConfigurationPath(v string)`

SetConfigurationPath sets ConfigurationPath field to given value.

### HasConfigurationPath

`func (o *RecommendationPhysicalItem) HasConfigurationPath() bool`

HasConfigurationPath returns a boolean if a field has been set.

### GetCount

`func (o *RecommendationPhysicalItem) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecommendationPhysicalItem) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecommendationPhysicalItem) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecommendationPhysicalItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIsNew

`func (o *RecommendationPhysicalItem) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *RecommendationPhysicalItem) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *RecommendationPhysicalItem) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *RecommendationPhysicalItem) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetMaxCount

`func (o *RecommendationPhysicalItem) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *RecommendationPhysicalItem) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *RecommendationPhysicalItem) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *RecommendationPhysicalItem) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetModel

`func (o *RecommendationPhysicalItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RecommendationPhysicalItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RecommendationPhysicalItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RecommendationPhysicalItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetParentMoid

`func (o *RecommendationPhysicalItem) GetParentMoid() string`

GetParentMoid returns the ParentMoid field if non-nil, zero value otherwise.

### GetParentMoidOk

`func (o *RecommendationPhysicalItem) GetParentMoidOk() (*string, bool)`

GetParentMoidOk returns a tuple with the ParentMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentMoid

`func (o *RecommendationPhysicalItem) SetParentMoid(v string)`

SetParentMoid sets ParentMoid field to given value.

### HasParentMoid

`func (o *RecommendationPhysicalItem) HasParentMoid() bool`

HasParentMoid returns a boolean if a field has been set.

### GetSourceMoid

`func (o *RecommendationPhysicalItem) GetSourceMoid() string`

GetSourceMoid returns the SourceMoid field if non-nil, zero value otherwise.

### GetSourceMoidOk

`func (o *RecommendationPhysicalItem) GetSourceMoidOk() (*string, bool)`

GetSourceMoidOk returns a tuple with the SourceMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMoid

`func (o *RecommendationPhysicalItem) SetSourceMoid(v string)`

SetSourceMoid sets SourceMoid field to given value.

### HasSourceMoid

`func (o *RecommendationPhysicalItem) HasSourceMoid() bool`

HasSourceMoid returns a boolean if a field has been set.

### GetUnit

`func (o *RecommendationPhysicalItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RecommendationPhysicalItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RecommendationPhysicalItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RecommendationPhysicalItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUuid

`func (o *RecommendationPhysicalItem) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RecommendationPhysicalItem) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RecommendationPhysicalItem) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RecommendationPhysicalItem) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCapacityRunway

`func (o *RecommendationPhysicalItem) GetCapacityRunway() RecommendationCapacityRunwayRelationship`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *RecommendationPhysicalItem) GetCapacityRunwayOk() (*RecommendationCapacityRunwayRelationship, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *RecommendationPhysicalItem) SetCapacityRunway(v RecommendationCapacityRunwayRelationship)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *RecommendationPhysicalItem) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### SetCapacityRunwayNil

`func (o *RecommendationPhysicalItem) SetCapacityRunwayNil(b bool)`

 SetCapacityRunwayNil sets the value for CapacityRunway to be an explicit nil

### UnsetCapacityRunway
`func (o *RecommendationPhysicalItem) UnsetCapacityRunway()`

UnsetCapacityRunway ensures that no value is present for CapacityRunway, not even an explicit nil
### GetClusterExpansion

`func (o *RecommendationPhysicalItem) GetClusterExpansion() RecommendationClusterExpansionRelationship`

GetClusterExpansion returns the ClusterExpansion field if non-nil, zero value otherwise.

### GetClusterExpansionOk

`func (o *RecommendationPhysicalItem) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool)`

GetClusterExpansionOk returns a tuple with the ClusterExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExpansion

`func (o *RecommendationPhysicalItem) SetClusterExpansion(v RecommendationClusterExpansionRelationship)`

SetClusterExpansion sets ClusterExpansion field to given value.

### HasClusterExpansion

`func (o *RecommendationPhysicalItem) HasClusterExpansion() bool`

HasClusterExpansion returns a boolean if a field has been set.

### SetClusterExpansionNil

`func (o *RecommendationPhysicalItem) SetClusterExpansionNil(b bool)`

 SetClusterExpansionNil sets the value for ClusterExpansion to be an explicit nil

### UnsetClusterExpansion
`func (o *RecommendationPhysicalItem) UnsetClusterExpansion()`

UnsetClusterExpansion ensures that no value is present for ClusterExpansion, not even an explicit nil
### GetPhysicalItem

`func (o *RecommendationPhysicalItem) GetPhysicalItem() []RecommendationPhysicalItemRelationship`

GetPhysicalItem returns the PhysicalItem field if non-nil, zero value otherwise.

### GetPhysicalItemOk

`func (o *RecommendationPhysicalItem) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool)`

GetPhysicalItemOk returns a tuple with the PhysicalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalItem

`func (o *RecommendationPhysicalItem) SetPhysicalItem(v []RecommendationPhysicalItemRelationship)`

SetPhysicalItem sets PhysicalItem field to given value.

### HasPhysicalItem

`func (o *RecommendationPhysicalItem) HasPhysicalItem() bool`

HasPhysicalItem returns a boolean if a field has been set.

### SetPhysicalItemNil

`func (o *RecommendationPhysicalItem) SetPhysicalItemNil(b bool)`

 SetPhysicalItemNil sets the value for PhysicalItem to be an explicit nil

### UnsetPhysicalItem
`func (o *RecommendationPhysicalItem) UnsetPhysicalItem()`

UnsetPhysicalItem ensures that no value is present for PhysicalItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


