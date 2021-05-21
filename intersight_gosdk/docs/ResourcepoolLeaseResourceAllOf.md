# ResourcepoolLeaseResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.LeaseResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.LeaseResource"]
**Feature** | Pointer to **string** | Lease opertion applied for the feature. | [optional] [readonly] 
**Var0Lease** | Pointer to [**ResourcepoolLeaseRelationship**](resourcepool.Lease.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewResourcepoolLeaseResourceAllOf

`func NewResourcepoolLeaseResourceAllOf(classId string, objectType string, ) *ResourcepoolLeaseResourceAllOf`

NewResourcepoolLeaseResourceAllOf instantiates a new ResourcepoolLeaseResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolLeaseResourceAllOfWithDefaults

`func NewResourcepoolLeaseResourceAllOfWithDefaults() *ResourcepoolLeaseResourceAllOf`

NewResourcepoolLeaseResourceAllOfWithDefaults instantiates a new ResourcepoolLeaseResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolLeaseResourceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolLeaseResourceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolLeaseResourceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolLeaseResourceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolLeaseResourceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolLeaseResourceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeature

`func (o *ResourcepoolLeaseResourceAllOf) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ResourcepoolLeaseResourceAllOf) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ResourcepoolLeaseResourceAllOf) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ResourcepoolLeaseResourceAllOf) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetVar0Lease

`func (o *ResourcepoolLeaseResourceAllOf) GetVar0Lease() ResourcepoolLeaseRelationship`

GetVar0Lease returns the Var0Lease field if non-nil, zero value otherwise.

### GetVar0LeaseOk

`func (o *ResourcepoolLeaseResourceAllOf) GetVar0LeaseOk() (*ResourcepoolLeaseRelationship, bool)`

GetVar0LeaseOk returns a tuple with the Var0Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0Lease

`func (o *ResourcepoolLeaseResourceAllOf) SetVar0Lease(v ResourcepoolLeaseRelationship)`

SetVar0Lease sets Var0Lease field to given value.

### HasVar0Lease

`func (o *ResourcepoolLeaseResourceAllOf) HasVar0Lease() bool`

HasVar0Lease returns a boolean if a field has been set.

### GetResource

`func (o *ResourcepoolLeaseResourceAllOf) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourcepoolLeaseResourceAllOf) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourcepoolLeaseResourceAllOf) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourcepoolLeaseResourceAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


