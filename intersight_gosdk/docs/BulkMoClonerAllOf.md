# BulkMoClonerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.MoCloner"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.MoCloner"]
**Responses** | Pointer to [**[]BulkRestResult**](BulkRestResult.md) |  | [optional] 
**Sources** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**Targets** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkMoClonerAllOf

`func NewBulkMoClonerAllOf(classId string, objectType string, ) *BulkMoClonerAllOf`

NewBulkMoClonerAllOf instantiates a new BulkMoClonerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoClonerAllOfWithDefaults

`func NewBulkMoClonerAllOfWithDefaults() *BulkMoClonerAllOf`

NewBulkMoClonerAllOfWithDefaults instantiates a new BulkMoClonerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoClonerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoClonerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoClonerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoClonerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoClonerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoClonerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetResponses

`func (o *BulkMoClonerAllOf) GetResponses() []BulkRestResult`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *BulkMoClonerAllOf) GetResponsesOk() (*[]BulkRestResult, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *BulkMoClonerAllOf) SetResponses(v []BulkRestResult)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *BulkMoClonerAllOf) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### SetResponsesNil

`func (o *BulkMoClonerAllOf) SetResponsesNil(b bool)`

 SetResponsesNil sets the value for Responses to be an explicit nil

### UnsetResponses
`func (o *BulkMoClonerAllOf) UnsetResponses()`

UnsetResponses ensures that no value is present for Responses, not even an explicit nil
### GetSources

`func (o *BulkMoClonerAllOf) GetSources() []MoBaseMo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkMoClonerAllOf) GetSourcesOk() (*[]MoBaseMo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkMoClonerAllOf) SetSources(v []MoBaseMo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkMoClonerAllOf) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *BulkMoClonerAllOf) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *BulkMoClonerAllOf) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetTargets

`func (o *BulkMoClonerAllOf) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoClonerAllOf) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoClonerAllOf) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoClonerAllOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoClonerAllOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoClonerAllOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetOrganization

`func (o *BulkMoClonerAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoClonerAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoClonerAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoClonerAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


