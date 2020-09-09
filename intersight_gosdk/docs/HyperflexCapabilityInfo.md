# HyperflexCapabilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityConstraints** | Pointer to [**[]HclConstraint**](hcl.Constraint.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the capability or feature set consisting of a collection of constraint rules and value. | [optional] 
**Value** | Pointer to **string** | Capability Value which is valid only iff all specified constraints match. | [optional] [readonly] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](hyperflex.AppCatalog.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexCapabilityInfo

`func NewHyperflexCapabilityInfo() *HyperflexCapabilityInfo`

NewHyperflexCapabilityInfo instantiates a new HyperflexCapabilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexCapabilityInfoWithDefaults

`func NewHyperflexCapabilityInfoWithDefaults() *HyperflexCapabilityInfo`

NewHyperflexCapabilityInfoWithDefaults instantiates a new HyperflexCapabilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilityConstraints

`func (o *HyperflexCapabilityInfo) GetCapabilityConstraints() []HclConstraint`

GetCapabilityConstraints returns the CapabilityConstraints field if non-nil, zero value otherwise.

### GetCapabilityConstraintsOk

`func (o *HyperflexCapabilityInfo) GetCapabilityConstraintsOk() (*[]HclConstraint, bool)`

GetCapabilityConstraintsOk returns a tuple with the CapabilityConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityConstraints

`func (o *HyperflexCapabilityInfo) SetCapabilityConstraints(v []HclConstraint)`

SetCapabilityConstraints sets CapabilityConstraints field to given value.

### HasCapabilityConstraints

`func (o *HyperflexCapabilityInfo) HasCapabilityConstraints() bool`

HasCapabilityConstraints returns a boolean if a field has been set.

### GetName

`func (o *HyperflexCapabilityInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexCapabilityInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexCapabilityInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexCapabilityInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *HyperflexCapabilityInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HyperflexCapabilityInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HyperflexCapabilityInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HyperflexCapabilityInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HyperflexCapabilityInfo) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexCapabilityInfo) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexCapabilityInfo) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexCapabilityInfo) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


