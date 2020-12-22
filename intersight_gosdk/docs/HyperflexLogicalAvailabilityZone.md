# HyperflexLogicalAvailabilityZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.LogicalAvailabilityZone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.LogicalAvailabilityZone"]
**AutoConfig** | Pointer to **bool** | Enable or disable Logical Availability Zones (LAZ). If enabled, HyperFlex Data Platform automatically selects and groups nodes into different availability zones. For HyperFlex Data Platform versions prior to 3.0 release, this setting does not apply. For HyperFlex Data Platform versions 3.0 or higher, this setting is only applicable to Fabric Interconnect attached HyperFlex systems with 8 or more converged nodes. | [optional] [default to false]

## Methods

### NewHyperflexLogicalAvailabilityZone

`func NewHyperflexLogicalAvailabilityZone(classId string, objectType string, ) *HyperflexLogicalAvailabilityZone`

NewHyperflexLogicalAvailabilityZone instantiates a new HyperflexLogicalAvailabilityZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLogicalAvailabilityZoneWithDefaults

`func NewHyperflexLogicalAvailabilityZoneWithDefaults() *HyperflexLogicalAvailabilityZone`

NewHyperflexLogicalAvailabilityZoneWithDefaults instantiates a new HyperflexLogicalAvailabilityZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexLogicalAvailabilityZone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexLogicalAvailabilityZone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexLogicalAvailabilityZone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexLogicalAvailabilityZone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexLogicalAvailabilityZone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexLogicalAvailabilityZone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoConfig

`func (o *HyperflexLogicalAvailabilityZone) GetAutoConfig() bool`

GetAutoConfig returns the AutoConfig field if non-nil, zero value otherwise.

### GetAutoConfigOk

`func (o *HyperflexLogicalAvailabilityZone) GetAutoConfigOk() (*bool, bool)`

GetAutoConfigOk returns a tuple with the AutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfig

`func (o *HyperflexLogicalAvailabilityZone) SetAutoConfig(v bool)`

SetAutoConfig sets AutoConfig field to given value.

### HasAutoConfig

`func (o *HyperflexLogicalAvailabilityZone) HasAutoConfig() bool`

HasAutoConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


