# RecommendationAbstractItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.PhysicalItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "recommendation.PhysicalItem"]
**Name** | Pointer to **string** | The name of the physical device recommended. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of physical device recommended. * &#x60;Disk&#x60; - The Enum value Disk represents that the item type recommended is a Physical Disk. * &#x60;Node&#x60; - The Enum value Node represents that the item type recommended is a Storage Node. * &#x60;Cluster&#x60; - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster. | [optional] [readonly] [default to "Disk"]

## Methods

### NewRecommendationAbstractItem

`func NewRecommendationAbstractItem(classId string, objectType string, ) *RecommendationAbstractItem`

NewRecommendationAbstractItem instantiates a new RecommendationAbstractItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationAbstractItemWithDefaults

`func NewRecommendationAbstractItemWithDefaults() *RecommendationAbstractItem`

NewRecommendationAbstractItemWithDefaults instantiates a new RecommendationAbstractItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationAbstractItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationAbstractItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationAbstractItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationAbstractItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationAbstractItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationAbstractItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *RecommendationAbstractItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecommendationAbstractItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecommendationAbstractItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecommendationAbstractItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RecommendationAbstractItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecommendationAbstractItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecommendationAbstractItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecommendationAbstractItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


