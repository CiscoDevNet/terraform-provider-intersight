# VnicCdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.Cdn"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.Cdn"]
**Source** | Pointer to **string** | Source of the CDN. It can either be user specified or be the same as the vNIC name. * &#x60;vnic&#x60; - Source of the CDN is the same as the vNIC name. * &#x60;user&#x60; - Source of the CDN is specified by the user. | [optional] [default to "vnic"]
**Value** | Pointer to **string** | The CDN value entered in case of user defined mode. | [optional] 

## Methods

### NewVnicCdn

`func NewVnicCdn(classId string, objectType string, ) *VnicCdn`

NewVnicCdn instantiates a new VnicCdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicCdnWithDefaults

`func NewVnicCdnWithDefaults() *VnicCdn`

NewVnicCdnWithDefaults instantiates a new VnicCdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicCdn) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicCdn) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicCdn) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicCdn) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicCdn) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicCdn) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSource

`func (o *VnicCdn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VnicCdn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VnicCdn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VnicCdn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetValue

`func (o *VnicCdn) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VnicCdn) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VnicCdn) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VnicCdn) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


