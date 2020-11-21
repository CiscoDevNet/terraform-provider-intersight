# HyperflexWwxnPrefixRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.WwxnPrefixRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.WwxnPrefixRange"]
**EndAddr** | Pointer to **string** | The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. | [optional] 
**StartAddr** | Pointer to **string** | The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. | [optional] 

## Methods

### NewHyperflexWwxnPrefixRange

`func NewHyperflexWwxnPrefixRange(classId string, objectType string, ) *HyperflexWwxnPrefixRange`

NewHyperflexWwxnPrefixRange instantiates a new HyperflexWwxnPrefixRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexWwxnPrefixRangeWithDefaults

`func NewHyperflexWwxnPrefixRangeWithDefaults() *HyperflexWwxnPrefixRange`

NewHyperflexWwxnPrefixRangeWithDefaults instantiates a new HyperflexWwxnPrefixRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexWwxnPrefixRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexWwxnPrefixRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexWwxnPrefixRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexWwxnPrefixRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexWwxnPrefixRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexWwxnPrefixRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddr

`func (o *HyperflexWwxnPrefixRange) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *HyperflexWwxnPrefixRange) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *HyperflexWwxnPrefixRange) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *HyperflexWwxnPrefixRange) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetStartAddr

`func (o *HyperflexWwxnPrefixRange) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *HyperflexWwxnPrefixRange) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *HyperflexWwxnPrefixRange) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *HyperflexWwxnPrefixRange) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


