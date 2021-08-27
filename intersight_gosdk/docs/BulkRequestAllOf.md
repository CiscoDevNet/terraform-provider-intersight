# BulkRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.Request"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.Request"]
**Requests** | Pointer to [**[]BulkSubRequest**](BulkSubRequest.md) |  | [optional] 
**Results** | Pointer to [**[]BulkApiResult**](BulkApiResult.md) |  | [optional] 
**Uri** | Pointer to **string** | The URI on which this bulk action is to be performed. | [optional] 
**Verb** | Pointer to **string** | The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * &#x60;POST&#x60; - Used to create a REST resource. * &#x60;PATCH&#x60; - Used to update a REST resource. * &#x60;DELETE&#x60; - Used to delete a REST resource. | [optional] [default to "POST"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkRequestAllOf

`func NewBulkRequestAllOf(classId string, objectType string, ) *BulkRequestAllOf`

NewBulkRequestAllOf instantiates a new BulkRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRequestAllOfWithDefaults

`func NewBulkRequestAllOfWithDefaults() *BulkRequestAllOf`

NewBulkRequestAllOfWithDefaults instantiates a new BulkRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRequests

`func (o *BulkRequestAllOf) GetRequests() []BulkSubRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BulkRequestAllOf) GetRequestsOk() (*[]BulkSubRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BulkRequestAllOf) SetRequests(v []BulkSubRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BulkRequestAllOf) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *BulkRequestAllOf) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *BulkRequestAllOf) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil
### GetResults

`func (o *BulkRequestAllOf) GetResults() []BulkApiResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkRequestAllOf) GetResultsOk() (*[]BulkApiResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkRequestAllOf) SetResults(v []BulkApiResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkRequestAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *BulkRequestAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BulkRequestAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetUri

`func (o *BulkRequestAllOf) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkRequestAllOf) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkRequestAllOf) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkRequestAllOf) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVerb

`func (o *BulkRequestAllOf) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *BulkRequestAllOf) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *BulkRequestAllOf) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *BulkRequestAllOf) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetOrganization

`func (o *BulkRequestAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkRequestAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


