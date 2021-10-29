# HyperflexVmProtectionSpaceUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmProtectionSpaceUsage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmProtectionSpaceUsage"]
**SpaceUsage** | Pointer to **int64** | Space usage of the Virtual Machine from StDataServiceManager. | [optional] [readonly] 

## Methods

### NewHyperflexVmProtectionSpaceUsageAllOf

`func NewHyperflexVmProtectionSpaceUsageAllOf(classId string, objectType string, ) *HyperflexVmProtectionSpaceUsageAllOf`

NewHyperflexVmProtectionSpaceUsageAllOf instantiates a new HyperflexVmProtectionSpaceUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmProtectionSpaceUsageAllOfWithDefaults

`func NewHyperflexVmProtectionSpaceUsageAllOfWithDefaults() *HyperflexVmProtectionSpaceUsageAllOf`

NewHyperflexVmProtectionSpaceUsageAllOfWithDefaults instantiates a new HyperflexVmProtectionSpaceUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmProtectionSpaceUsageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmProtectionSpaceUsageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSpaceUsage

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetSpaceUsage() int64`

GetSpaceUsage returns the SpaceUsage field if non-nil, zero value otherwise.

### GetSpaceUsageOk

`func (o *HyperflexVmProtectionSpaceUsageAllOf) GetSpaceUsageOk() (*int64, bool)`

GetSpaceUsageOk returns a tuple with the SpaceUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceUsage

`func (o *HyperflexVmProtectionSpaceUsageAllOf) SetSpaceUsage(v int64)`

SetSpaceUsage sets SpaceUsage field to given value.

### HasSpaceUsage

`func (o *HyperflexVmProtectionSpaceUsageAllOf) HasSpaceUsage() bool`

HasSpaceUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


