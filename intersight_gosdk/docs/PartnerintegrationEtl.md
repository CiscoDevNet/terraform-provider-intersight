# PartnerintegrationEtl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.Etl"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.Etl"]
**Data** | Pointer to **interface{}** | Transformation model in yaml format. | [optional] 
**Name** | Pointer to **string** | Placeholder name for the ETL. | [optional] 
**Inventory** | Pointer to [**PartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationEtl

`func NewPartnerintegrationEtl(classId string, objectType string, ) *PartnerintegrationEtl`

NewPartnerintegrationEtl instantiates a new PartnerintegrationEtl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationEtlWithDefaults

`func NewPartnerintegrationEtlWithDefaults() *PartnerintegrationEtl`

NewPartnerintegrationEtlWithDefaults instantiates a new PartnerintegrationEtl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationEtl) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationEtl) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationEtl) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationEtl) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationEtl) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationEtl) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetData

`func (o *PartnerintegrationEtl) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PartnerintegrationEtl) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PartnerintegrationEtl) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PartnerintegrationEtl) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PartnerintegrationEtl) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PartnerintegrationEtl) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetName

`func (o *PartnerintegrationEtl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerintegrationEtl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerintegrationEtl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerintegrationEtl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInventory

`func (o *PartnerintegrationEtl) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationEtl) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationEtl) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationEtl) HasInventory() bool`

HasInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


