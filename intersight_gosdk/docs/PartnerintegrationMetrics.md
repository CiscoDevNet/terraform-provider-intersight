# PartnerintegrationMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.Metrics"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.Metrics"]
**Data** | Pointer to [**NullablePartnerintegrationMetricsModel**](PartnerintegrationMetricsModel.md) |  | [optional] 
**Name** | Pointer to **string** | Placeholder name for the Metrics. | [optional] 
**Inventory** | Pointer to [**NullablePartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationMetrics

`func NewPartnerintegrationMetrics(classId string, objectType string, ) *PartnerintegrationMetrics`

NewPartnerintegrationMetrics instantiates a new PartnerintegrationMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationMetricsWithDefaults

`func NewPartnerintegrationMetricsWithDefaults() *PartnerintegrationMetrics`

NewPartnerintegrationMetricsWithDefaults instantiates a new PartnerintegrationMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationMetrics) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationMetrics) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationMetrics) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationMetrics) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationMetrics) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationMetrics) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetData

`func (o *PartnerintegrationMetrics) GetData() PartnerintegrationMetricsModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PartnerintegrationMetrics) GetDataOk() (*PartnerintegrationMetricsModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PartnerintegrationMetrics) SetData(v PartnerintegrationMetricsModel)`

SetData sets Data field to given value.

### HasData

`func (o *PartnerintegrationMetrics) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PartnerintegrationMetrics) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PartnerintegrationMetrics) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetName

`func (o *PartnerintegrationMetrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationMetrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationMetrics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationMetrics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInventory

`func (o *PartnerintegrationMetrics) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationMetrics) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationMetrics) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationMetrics) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *PartnerintegrationMetrics) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *PartnerintegrationMetrics) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


