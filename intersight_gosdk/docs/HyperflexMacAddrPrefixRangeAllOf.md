# HyperflexMacAddrPrefixRangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.MacAddrPrefixRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.MacAddrPrefixRange"]
**EndAddr** | Pointer to **string** | The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. | [optional] 
**StartAddr** | Pointer to **string** | The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. | [optional] 

## Methods

### NewHyperflexMacAddrPrefixRangeAllOf

`func NewHyperflexMacAddrPrefixRangeAllOf(classId string, objectType string, ) *HyperflexMacAddrPrefixRangeAllOf`

NewHyperflexMacAddrPrefixRangeAllOf instantiates a new HyperflexMacAddrPrefixRangeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexMacAddrPrefixRangeAllOfWithDefaults

`func NewHyperflexMacAddrPrefixRangeAllOfWithDefaults() *HyperflexMacAddrPrefixRangeAllOf`

NewHyperflexMacAddrPrefixRangeAllOfWithDefaults instantiates a new HyperflexMacAddrPrefixRangeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexMacAddrPrefixRangeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexMacAddrPrefixRangeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetStartAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *HyperflexMacAddrPrefixRangeAllOf) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *HyperflexMacAddrPrefixRangeAllOf) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


