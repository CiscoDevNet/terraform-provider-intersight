# RecommendationPurchaseOrderListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.PurchaseOrderList"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.PurchaseOrderList"]
**ItemList** | Pointer to **string** | The comma seperated list of items for the current recommendation. This can be used to generate an estimate by uploading it to Cisco Commerce Workspace. | [optional] [readonly] 
**ClusterExpansion** | Pointer to [**RecommendationClusterExpansionRelationship**](RecommendationClusterExpansionRelationship.md) |  | [optional] 

## Methods

### NewRecommendationPurchaseOrderListAllOf

`func NewRecommendationPurchaseOrderListAllOf(classId string, objectType string, ) *RecommendationPurchaseOrderListAllOf`

NewRecommendationPurchaseOrderListAllOf instantiates a new RecommendationPurchaseOrderListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationPurchaseOrderListAllOfWithDefaults

`func NewRecommendationPurchaseOrderListAllOfWithDefaults() *RecommendationPurchaseOrderListAllOf`

NewRecommendationPurchaseOrderListAllOfWithDefaults instantiates a new RecommendationPurchaseOrderListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationPurchaseOrderListAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationPurchaseOrderListAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationPurchaseOrderListAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationPurchaseOrderListAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationPurchaseOrderListAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationPurchaseOrderListAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetItemList

`func (o *RecommendationPurchaseOrderListAllOf) GetItemList() string`

GetItemList returns the ItemList field if non-nil, zero value otherwise.

### GetItemListOk

`func (o *RecommendationPurchaseOrderListAllOf) GetItemListOk() (*string, bool)`

GetItemListOk returns a tuple with the ItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemList

`func (o *RecommendationPurchaseOrderListAllOf) SetItemList(v string)`

SetItemList sets ItemList field to given value.

### HasItemList

`func (o *RecommendationPurchaseOrderListAllOf) HasItemList() bool`

HasItemList returns a boolean if a field has been set.

### GetClusterExpansion

`func (o *RecommendationPurchaseOrderListAllOf) GetClusterExpansion() RecommendationClusterExpansionRelationship`

GetClusterExpansion returns the ClusterExpansion field if non-nil, zero value otherwise.

### GetClusterExpansionOk

`func (o *RecommendationPurchaseOrderListAllOf) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool)`

GetClusterExpansionOk returns a tuple with the ClusterExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExpansion

`func (o *RecommendationPurchaseOrderListAllOf) SetClusterExpansion(v RecommendationClusterExpansionRelationship)`

SetClusterExpansion sets ClusterExpansion field to given value.

### HasClusterExpansion

`func (o *RecommendationPurchaseOrderListAllOf) HasClusterExpansion() bool`

HasClusterExpansion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


