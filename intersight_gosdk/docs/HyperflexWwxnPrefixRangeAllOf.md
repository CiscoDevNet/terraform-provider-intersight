# HyperflexWwxnPrefixRangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.WwxnPrefixRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.WwxnPrefixRange"]
**EndAddr** | Pointer to **string** | The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. | [optional] 
**StartAddr** | Pointer to **string** | The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. | [optional] 

## Methods

### NewHyperflexWwxnPrefixRangeAllOf

`func NewHyperflexWwxnPrefixRangeAllOf(classId string, objectType string, ) *HyperflexWwxnPrefixRangeAllOf`

NewHyperflexWwxnPrefixRangeAllOf instantiates a new HyperflexWwxnPrefixRangeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexWwxnPrefixRangeAllOfWithDefaults

`func NewHyperflexWwxnPrefixRangeAllOfWithDefaults() *HyperflexWwxnPrefixRangeAllOf`

NewHyperflexWwxnPrefixRangeAllOfWithDefaults instantiates a new HyperflexWwxnPrefixRangeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexWwxnPrefixRangeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexWwxnPrefixRangeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexWwxnPrefixRangeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexWwxnPrefixRangeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexWwxnPrefixRangeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexWwxnPrefixRangeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *HyperflexWwxnPrefixRangeAllOf) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetStartAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *HyperflexWwxnPrefixRangeAllOf) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *HyperflexWwxnPrefixRangeAllOf) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


