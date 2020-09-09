# CapabilityPortGroupAggregationDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationCap** | Pointer to **string** | Aggregation capability for port group. | [optional] 
**Hw40GPortGroupCap** | Pointer to **bool** | Indicates support for 40G port group capability. | [optional] 
**Pgtype** | Pointer to **string** | The type of port group for which this capability is defined. | [optional] 

## Methods

### NewCapabilityPortGroupAggregationDef

`func NewCapabilityPortGroupAggregationDef() *CapabilityPortGroupAggregationDef`

NewCapabilityPortGroupAggregationDef instantiates a new CapabilityPortGroupAggregationDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPortGroupAggregationDefWithDefaults

`func NewCapabilityPortGroupAggregationDefWithDefaults() *CapabilityPortGroupAggregationDef`

NewCapabilityPortGroupAggregationDefWithDefaults instantiates a new CapabilityPortGroupAggregationDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationCap

`func (o *CapabilityPortGroupAggregationDef) GetAggregationCap() string`

GetAggregationCap returns the AggregationCap field if non-nil, zero value otherwise.

### GetAggregationCapOk

`func (o *CapabilityPortGroupAggregationDef) GetAggregationCapOk() (*string, bool)`

GetAggregationCapOk returns a tuple with the AggregationCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationCap

`func (o *CapabilityPortGroupAggregationDef) SetAggregationCap(v string)`

SetAggregationCap sets AggregationCap field to given value.

### HasAggregationCap

`func (o *CapabilityPortGroupAggregationDef) HasAggregationCap() bool`

HasAggregationCap returns a boolean if a field has been set.

### GetHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDef) GetHw40GPortGroupCap() bool`

GetHw40GPortGroupCap returns the Hw40GPortGroupCap field if non-nil, zero value otherwise.

### GetHw40GPortGroupCapOk

`func (o *CapabilityPortGroupAggregationDef) GetHw40GPortGroupCapOk() (*bool, bool)`

GetHw40GPortGroupCapOk returns a tuple with the Hw40GPortGroupCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDef) SetHw40GPortGroupCap(v bool)`

SetHw40GPortGroupCap sets Hw40GPortGroupCap field to given value.

### HasHw40GPortGroupCap

`func (o *CapabilityPortGroupAggregationDef) HasHw40GPortGroupCap() bool`

HasHw40GPortGroupCap returns a boolean if a field has been set.

### GetPgtype

`func (o *CapabilityPortGroupAggregationDef) GetPgtype() string`

GetPgtype returns the Pgtype field if non-nil, zero value otherwise.

### GetPgtypeOk

`func (o *CapabilityPortGroupAggregationDef) GetPgtypeOk() (*string, bool)`

GetPgtypeOk returns a tuple with the Pgtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgtype

`func (o *CapabilityPortGroupAggregationDef) SetPgtype(v string)`

SetPgtype sets Pgtype field to given value.

### HasPgtype

`func (o *CapabilityPortGroupAggregationDef) HasPgtype() bool`

HasPgtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


