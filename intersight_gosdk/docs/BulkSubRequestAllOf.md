# BulkSubRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "bulk.RestSubRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "bulk.RestSubRequest"]
**Uri** | Pointer to **string** | The URI on which this action is to be performed. | [optional] 
**Verb** | Pointer to **string** | The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value is used to override the top level verb. * &#x60;POST&#x60; - Used to create a REST resource. * &#x60;PATCH&#x60; - Used to update a REST resource. * &#x60;DELETE&#x60; - Used to delete a REST resource. | [optional] [default to "POST"]

## Methods

### NewBulkSubRequestAllOf

`func NewBulkSubRequestAllOf(classId string, objectType string, ) *BulkSubRequestAllOf`

NewBulkSubRequestAllOf instantiates a new BulkSubRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSubRequestAllOfWithDefaults

`func NewBulkSubRequestAllOfWithDefaults() *BulkSubRequestAllOf`

NewBulkSubRequestAllOfWithDefaults instantiates a new BulkSubRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkSubRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkSubRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkSubRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkSubRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkSubRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkSubRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUri

`func (o *BulkSubRequestAllOf) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkSubRequestAllOf) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkSubRequestAllOf) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkSubRequestAllOf) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVerb

`func (o *BulkSubRequestAllOf) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *BulkSubRequestAllOf) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *BulkSubRequestAllOf) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *BulkSubRequestAllOf) HasVerb() bool`

HasVerb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


