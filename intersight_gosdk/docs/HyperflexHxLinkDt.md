# HyperflexHxLinkDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxLinkDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxLinkDt"]
**Comments** | Pointer to **string** | A description that provides additional context for usage of the resource. | [optional] [readonly] 
**Href** | Pointer to **string** | The URI of resource. Target URL for making REST call. | [optional] [readonly] 
**Method** | Pointer to **string** | The HTTP method to use when making the REST call to this resource. * &#x60;POST&#x60; - The http verb for creating POST requests to a rest endpoint. * &#x60;GET&#x60; - The http verb for creating GET requests to a rest endpoint. * &#x60;PUT&#x60; - The http verb for creating PUT requests to a rest endpoint. * &#x60;DELETE&#x60; - The http verb for creating DELETE requests to a rest endpoint. | [optional] [readonly] [default to "POST"]
**Rel** | Pointer to **string** | The relationship of the link to this resource. | [optional] [readonly] 

## Methods

### NewHyperflexHxLinkDt

`func NewHyperflexHxLinkDt(classId string, objectType string, ) *HyperflexHxLinkDt`

NewHyperflexHxLinkDt instantiates a new HyperflexHxLinkDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxLinkDtWithDefaults

`func NewHyperflexHxLinkDtWithDefaults() *HyperflexHxLinkDt`

NewHyperflexHxLinkDtWithDefaults instantiates a new HyperflexHxLinkDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxLinkDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxLinkDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxLinkDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxLinkDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxLinkDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxLinkDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComments

`func (o *HyperflexHxLinkDt) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *HyperflexHxLinkDt) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *HyperflexHxLinkDt) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *HyperflexHxLinkDt) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetHref

`func (o *HyperflexHxLinkDt) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HyperflexHxLinkDt) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HyperflexHxLinkDt) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *HyperflexHxLinkDt) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMethod

`func (o *HyperflexHxLinkDt) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HyperflexHxLinkDt) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HyperflexHxLinkDt) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HyperflexHxLinkDt) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRel

`func (o *HyperflexHxLinkDt) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *HyperflexHxLinkDt) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *HyperflexHxLinkDt) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *HyperflexHxLinkDt) HasRel() bool`

HasRel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


