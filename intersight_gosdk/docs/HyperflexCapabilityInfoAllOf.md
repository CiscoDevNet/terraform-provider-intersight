# HyperflexCapabilityInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.CapabilityInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.CapabilityInfo"]
**CapabilityConstraints** | Pointer to [**[]HclConstraint**](HclConstraint.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the capability or feature set consisting of a collection of constraint rules and value. | [optional] 
**Value** | Pointer to **string** | Capability Value which is valid only iff all specified constraints match. | [optional] [readonly] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHyperflexCapabilityInfoAllOf

`func NewHyperflexCapabilityInfoAllOf(classId string, objectType string, ) *HyperflexCapabilityInfoAllOf`

NewHyperflexCapabilityInfoAllOf instantiates a new HyperflexCapabilityInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexCapabilityInfoAllOfWithDefaults

`func NewHyperflexCapabilityInfoAllOfWithDefaults() *HyperflexCapabilityInfoAllOf`

NewHyperflexCapabilityInfoAllOfWithDefaults instantiates a new HyperflexCapabilityInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexCapabilityInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexCapabilityInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexCapabilityInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexCapabilityInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexCapabilityInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexCapabilityInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapabilityConstraints

`func (o *HyperflexCapabilityInfoAllOf) GetCapabilityConstraints() []HclConstraint`

GetCapabilityConstraints returns the CapabilityConstraints field if non-nil, zero value otherwise.

### GetCapabilityConstraintsOk

`func (o *HyperflexCapabilityInfoAllOf) GetCapabilityConstraintsOk() (*[]HclConstraint, bool)`

GetCapabilityConstraintsOk returns a tuple with the CapabilityConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityConstraints

`func (o *HyperflexCapabilityInfoAllOf) SetCapabilityConstraints(v []HclConstraint)`

SetCapabilityConstraints sets CapabilityConstraints field to given value.

### HasCapabilityConstraints

`func (o *HyperflexCapabilityInfoAllOf) HasCapabilityConstraints() bool`

HasCapabilityConstraints returns a boolean if a field has been set.

### SetCapabilityConstraintsNil

`func (o *HyperflexCapabilityInfoAllOf) SetCapabilityConstraintsNil(b bool)`

 SetCapabilityConstraintsNil sets the value for CapabilityConstraints to be an explicit nil

### UnsetCapabilityConstraints
`func (o *HyperflexCapabilityInfoAllOf) UnsetCapabilityConstraints()`

UnsetCapabilityConstraints ensures that no value is present for CapabilityConstraints, not even an explicit nil
### GetName

`func (o *HyperflexCapabilityInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexCapabilityInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexCapabilityInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexCapabilityInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *HyperflexCapabilityInfoAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HyperflexCapabilityInfoAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HyperflexCapabilityInfoAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HyperflexCapabilityInfoAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HyperflexCapabilityInfoAllOf) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexCapabilityInfoAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexCapabilityInfoAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexCapabilityInfoAllOf) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


