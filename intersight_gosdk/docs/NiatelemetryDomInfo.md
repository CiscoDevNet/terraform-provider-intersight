# NiatelemetryDomInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DomInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DomInfo"]
**Avg** | Pointer to **string** | Returns Average value of the transceiver sensor. | [optional] 
**Dn** | Pointer to **string** | Returns distinguished name of the transceiver. | [optional] 
**Instant** | Pointer to **string** | Returns instant value of the  transceiversensor. | [optional] 
**Max** | Pointer to **string** | Returns Maximum value reported by the transceiver sensor. | [optional] 
**Min** | Pointer to **string** | Returns Minimum value reported by the transceiver sensor. | [optional] 
**Unit** | Pointer to **string** | Returns transceiver sensor&#39;s unit identifier. | [optional] 
**Value** | Pointer to **string** | Returns calibration value (unit) of transceiver sensor. | [optional] 

## Methods

### NewNiatelemetryDomInfo

`func NewNiatelemetryDomInfo(classId string, objectType string, ) *NiatelemetryDomInfo`

NewNiatelemetryDomInfo instantiates a new NiatelemetryDomInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDomInfoWithDefaults

`func NewNiatelemetryDomInfoWithDefaults() *NiatelemetryDomInfo`

NewNiatelemetryDomInfoWithDefaults instantiates a new NiatelemetryDomInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDomInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDomInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDomInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDomInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDomInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDomInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvg

`func (o *NiatelemetryDomInfo) GetAvg() string`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *NiatelemetryDomInfo) GetAvgOk() (*string, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *NiatelemetryDomInfo) SetAvg(v string)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *NiatelemetryDomInfo) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryDomInfo) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryDomInfo) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryDomInfo) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryDomInfo) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetInstant

`func (o *NiatelemetryDomInfo) GetInstant() string`

GetInstant returns the Instant field if non-nil, zero value otherwise.

### GetInstantOk

`func (o *NiatelemetryDomInfo) GetInstantOk() (*string, bool)`

GetInstantOk returns a tuple with the Instant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstant

`func (o *NiatelemetryDomInfo) SetInstant(v string)`

SetInstant sets Instant field to given value.

### HasInstant

`func (o *NiatelemetryDomInfo) HasInstant() bool`

HasInstant returns a boolean if a field has been set.

### GetMax

`func (o *NiatelemetryDomInfo) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NiatelemetryDomInfo) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NiatelemetryDomInfo) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *NiatelemetryDomInfo) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *NiatelemetryDomInfo) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NiatelemetryDomInfo) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NiatelemetryDomInfo) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *NiatelemetryDomInfo) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetUnit

`func (o *NiatelemetryDomInfo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *NiatelemetryDomInfo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *NiatelemetryDomInfo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *NiatelemetryDomInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *NiatelemetryDomInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NiatelemetryDomInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NiatelemetryDomInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NiatelemetryDomInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


