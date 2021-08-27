# BulkMoMerger

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

### NewBulkMoMerger

`func NewBulkMoMerger(classId string, objectType string, ) *BulkMoMerger`

NewBulkMoMerger instantiates a new BulkMoMerger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMoMergerWithDefaults

`func NewBulkMoMergerWithDefaults() *BulkMoMerger`

NewBulkMoMergerWithDefaults instantiates a new BulkMoMerger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkMoMerger) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkMoMerger) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkMoMerger) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkMoMerger) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkMoMerger) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkMoMerger) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMergeAction

`func (o *BulkMoMerger) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *BulkMoMerger) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *BulkMoMerger) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *BulkMoMerger) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetResponses

`func (o *BulkMoMerger) GetResponses() []BulkRestResult`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *BulkMoMerger) GetResponsesOk() (*[]BulkRestResult, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *BulkMoMerger) SetResponses(v []BulkRestResult)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *BulkMoMerger) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### SetResponsesNil

`func (o *BulkMoMerger) SetResponsesNil(b bool)`

 SetResponsesNil sets the value for Responses to be an explicit nil

### UnsetResponses
`func (o *BulkMoMerger) UnsetResponses()`

UnsetResponses ensures that no value is present for Responses, not even an explicit nil
### GetSources

`func (o *BulkMoMerger) GetSources() []MoBaseMo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BulkMoMerger) GetSourcesOk() (*[]MoBaseMo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BulkMoMerger) SetSources(v []MoBaseMo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BulkMoMerger) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *BulkMoMerger) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *BulkMoMerger) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetTargetConfig

`func (o *BulkMoMerger) GetTargetConfig() MoBaseMo`

GetTargetConfig returns the TargetConfig field if non-nil, zero value otherwise.

### GetTargetConfigOk

`func (o *BulkMoMerger) GetTargetConfigOk() (*MoBaseMo, bool)`

GetTargetConfigOk returns a tuple with the TargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfig

`func (o *BulkMoMerger) SetTargetConfig(v MoBaseMo)`

SetTargetConfig sets TargetConfig field to given value.

### HasTargetConfig

`func (o *BulkMoMerger) HasTargetConfig() bool`

HasTargetConfig returns a boolean if a field has been set.

### GetTargets

`func (o *BulkMoMerger) GetTargets() []MoBaseMo`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *BulkMoMerger) GetTargetsOk() (*[]MoBaseMo, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *BulkMoMerger) SetTargets(v []MoBaseMo)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *BulkMoMerger) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *BulkMoMerger) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *BulkMoMerger) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetOrganization

`func (o *BulkMoMerger) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkMoMerger) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkMoMerger) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkMoMerger) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


