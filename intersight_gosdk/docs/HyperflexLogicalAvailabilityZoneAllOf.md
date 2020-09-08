# HyperflexLogicalAvailabilityZoneAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoConfig** | Pointer to **bool** | Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to Fabric Interconnect attached HyperFlex systems with 8 or more converged nodes. | [optional] 

## Methods

### NewHyperflexLogicalAvailabilityZoneAllOf

`func NewHyperflexLogicalAvailabilityZoneAllOf() *HyperflexLogicalAvailabilityZoneAllOf`

NewHyperflexLogicalAvailabilityZoneAllOf instantiates a new HyperflexLogicalAvailabilityZoneAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLogicalAvailabilityZoneAllOfWithDefaults

`func NewHyperflexLogicalAvailabilityZoneAllOfWithDefaults() *HyperflexLogicalAvailabilityZoneAllOf`

NewHyperflexLogicalAvailabilityZoneAllOfWithDefaults instantiates a new HyperflexLogicalAvailabilityZoneAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoConfig

`func (o *HyperflexLogicalAvailabilityZoneAllOf) GetAutoConfig() bool`

GetAutoConfig returns the AutoConfig field if non-nil, zero value otherwise.

### GetAutoConfigOk

`func (o *HyperflexLogicalAvailabilityZoneAllOf) GetAutoConfigOk() (*bool, bool)`

GetAutoConfigOk returns a tuple with the AutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfig

`func (o *HyperflexLogicalAvailabilityZoneAllOf) SetAutoConfig(v bool)`

SetAutoConfig sets AutoConfig field to given value.

### HasAutoConfig

`func (o *HyperflexLogicalAvailabilityZoneAllOf) HasAutoConfig() bool`

HasAutoConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


