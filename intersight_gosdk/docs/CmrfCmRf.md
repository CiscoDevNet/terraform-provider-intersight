# CmrfCmRf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moid** | Pointer to **string** | The Moid of the referenced REST resource. | [optional] [readonly] 
**Selector** | Pointer to **string** | An OData $filter expression which describes the REST resource to be referenced. This field may be set instead of &#39;moid&#39; by clients. 1. If &#39;moid&#39; is set this field is ignored. 1. If &#39;selector&#39; is set and &#39;moid&#39; is empty/absent from the request, Intersight determines the Moid of the resource matching the filter expression and populates it in the Cmrf that is part of the object instance being inserted/updated to fulfill the REST request. An error is returned if the filter matches zero or more than one REST resource. An example filter string is: Serial eq &#39;3AA8B7T11&#39;. | [optional] 

## Methods

### NewCmrfCmRf

`func NewCmrfCmRf() *CmrfCmRf`

NewCmrfCmRf instantiates a new CmrfCmRf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmrfCmRfWithDefaults

`func NewCmrfCmRfWithDefaults() *CmrfCmRf`

NewCmrfCmRfWithDefaults instantiates a new CmrfCmRf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoid

`func (o *CmrfCmRf) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *CmrfCmRf) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *CmrfCmRf) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *CmrfCmRf) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetSelector

`func (o *CmrfCmRf) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CmrfCmRf) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CmrfCmRf) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *CmrfCmRf) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


