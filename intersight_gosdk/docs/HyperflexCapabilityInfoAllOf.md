# HyperflexCapabilityInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityConstraints** | Pointer to [**[]HclConstraint**](hcl.Constraint.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the capability or feature set consisting of a collection of constraint rules and value. | [optional] 
**Value** | Pointer to **string** | Capability Value which is valid only iff all specified constraints match. | [optional] [readonly] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](hyperflex.AppCatalog.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexCapabilityInfoAllOf

`func NewHyperflexCapabilityInfoAllOf() *HyperflexCapabilityInfoAllOf`

NewHyperflexCapabilityInfoAllOf instantiates a new HyperflexCapabilityInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexCapabilityInfoAllOfWithDefaults

`func NewHyperflexCapabilityInfoAllOfWithDefaults() *HyperflexCapabilityInfoAllOf`

NewHyperflexCapabilityInfoAllOfWithDefaults instantiates a new HyperflexCapabilityInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


