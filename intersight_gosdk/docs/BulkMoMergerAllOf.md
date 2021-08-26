# BulkMoMergerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.MoMerger"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.MoMerger"]
**MergeAction** | Pointer to **string** | The type of merge action to be applied on the target MOs.  * &#x60;Merge&#x60; - The null properties/relationships of the source MO will be ignored for the target MO. The non-null properties/relationships of the source will override the target MO properties/relationships. * &#x60;Replace&#x60; - Merge action as described in RFC 7386. The null properties/relationships of the source MO will be deleted on the target MO.The non-null properties/relationships of the source will override the target MO properties/relationships.When source object type is different from target, only the properties common to both source and target  will be affected.Other properties on the target will be ignored. | [optional] [default to "Merge"]
**Responses** | Pointer to [**[]BulkRestResult**](BulkRestResult.md) |  | [optional] 
**Sources** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**TargetConfig** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 
**Targets** | Pointer to [**[]MoBaseMo**](MoBaseMo.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkMoMergerAllOf

`func NewBulkMoMergerAllOf(classId string, objectType string, ) *BulkMoMergerAllOf`

NewBulkMoMergerAllOf instantiates a new BulkMoMergerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoMergerAllOfWithDefaults

`func NewBulkMoMergerAllOfWithDefaults() *BulkMoMergerAllOf`

NewBulkMoMergerAllOfWithDefaults instantiates a new BulkMoMergerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoMergerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoMergerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoMergerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoMergerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoMergerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoMergerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMergeAction

`func (o *BulkMoMergerAllOf) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *BulkMoMergerAllOf) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *BulkMoMergerAllOf) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *BulkMoMergerAllOf) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetResponses

`func (o *BulkMoMergerAllOf) GetResponses() []BulkRestResult`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *BulkMoMergerAllOf) GetResponsesOk() (*[]BulkRestResult, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *BulkMoMergerAllOf) SetResponses(v []BulkRestResult)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *BulkMoMergerAllOf) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### SetResponsesNil

`func (o *BulkMoMergerAllOf) SetResponsesNil(b bool)`

 SetResponsesNil sets the value for Responses to be an explicit nil

### UnsetResponses
`func (o *BulkMoMergerAllOf) UnsetResponses()`

UnsetResponses ensures that no value is present for Responses, not even an explicit nil
### GetSources

`func (o *BulkMoMergerAllOf) GetSources() []MoBaseMo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkMoMergerAllOf) GetSourcesOk() (*[]MoBaseMo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkMoMergerAllOf) SetSources(v []MoBaseMo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkMoMergerAllOf) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *BulkMoMergerAllOf) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *BulkMoMergerAllOf) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetTargetConfig

`func (o *BulkMoMergerAllOf) GetTargetConfig() MoBaseMo`

GetTargetConfig returns the TargetConfig field if non-nil, zero value otherwise.

### GetTargetConfigOk

`func (o *BulkMoMergerAllOf) GetTargetConfigOk() (*MoBaseMo, bool)`

GetTargetConfigOk returns a tuple with the TargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfig

`func (o *BulkMoMergerAllOf) SetTargetConfig(v MoBaseMo)`

SetTargetConfig sets TargetConfig field to given value.

### HasTargetConfig

`func (o *BulkMoMergerAllOf) HasTargetConfig() bool`

HasTargetConfig returns a boolean if a field has been set.

### GetTargets

`func (o *BulkMoMergerAllOf) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoMergerAllOf) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoMergerAllOf) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoMergerAllOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoMergerAllOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoMergerAllOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetOrganization

`func (o *BulkMoMergerAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoMergerAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoMergerAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoMergerAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


