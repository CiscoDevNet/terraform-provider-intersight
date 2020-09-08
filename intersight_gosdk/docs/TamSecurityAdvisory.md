# TamSecurityAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]TamAction**](tam.Action.md) |  | [optional] 
**AdvisoryId** | Pointer to **string** | Cisco generated identifier for the published security advisory. | [optional] 
**ApiDataSources** | Pointer to [**[]TamApiDataSource**](tam.ApiDataSource.md) |  | [optional] 
**BaseScore** | Pointer to **float32** | CVSS version 3 base score for the security Advisory. | [optional] 
**CveIds** | Pointer to **[]string** |  | [optional] 
**DatePublished** | Pointer to [**time.Time**](time.Time.md) | Date when the security advisory was first published by Cisco. | [optional] 
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date when the security advisory was last updated by Cisco. | [optional] 
**EnvironmentalScore** | Pointer to **float32** | CVSS version 3 environmental score for the security Advisory. | [optional] 
**ExternalUrl** | Pointer to **string** | A link to an external URL describing security Advisory in more details. | [optional] 
**Recommendation** | Pointer to **string** | Recommended action to resolve the security advisory. | [optional] 
**Status** | Pointer to **string** | Cisco assigned status of the published advisory based on whether the investigation is complete or on-going. * &#x60;interim&#x60; - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available. * &#x60;final&#x60; - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability. | [optional] [default to "interim"]
**TemporalScore** | Pointer to **float32** | CVSS version 3 temporal score for the security Advisory. | [optional] 
**Version** | Pointer to **string** | Cisco assigned advisory version after latest revision. | [optional] 
**Workaround** | Pointer to **string** | Workarounds available for the advisory. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewTamSecurityAdvisory

`func NewTamSecurityAdvisory() *TamSecurityAdvisory`

NewTamSecurityAdvisory instantiates a new TamSecurityAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamSecurityAdvisoryWithDefaults

`func NewTamSecurityAdvisoryWithDefaults() *TamSecurityAdvisory`

NewTamSecurityAdvisoryWithDefaults instantiates a new TamSecurityAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TamSecurityAdvisory) GetActions() []TamAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TamSecurityAdvisory) GetActionsOk() (*[]TamAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TamSecurityAdvisory) SetActions(v []TamAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TamSecurityAdvisory) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdvisoryId

`func (o *TamSecurityAdvisory) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *TamSecurityAdvisory) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *TamSecurityAdvisory) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *TamSecurityAdvisory) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetApiDataSources

`func (o *TamSecurityAdvisory) GetApiDataSources() []TamApiDataSource`

GetApiDataSources returns the ApiDataSources field if non-nil, zero value otherwise.

### GetApiDataSourcesOk

`func (o *TamSecurityAdvisory) GetApiDataSourcesOk() (*[]TamApiDataSource, bool)`

GetApiDataSourcesOk returns a tuple with the ApiDataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDataSources

`func (o *TamSecurityAdvisory) SetApiDataSources(v []TamApiDataSource)`

SetApiDataSources sets ApiDataSources field to given value.

### HasApiDataSources

`func (o *TamSecurityAdvisory) HasApiDataSources() bool`

HasApiDataSources returns a boolean if a field has been set.

### GetBaseScore

`func (o *TamSecurityAdvisory) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *TamSecurityAdvisory) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *TamSecurityAdvisory) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *TamSecurityAdvisory) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetCveIds

`func (o *TamSecurityAdvisory) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *TamSecurityAdvisory) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *TamSecurityAdvisory) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *TamSecurityAdvisory) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### GetDatePublished

`func (o *TamSecurityAdvisory) GetDatePublished() time.Time`

GetDatePublished returns the DatePublished field if non-nil, zero value otherwise.

### GetDatePublishedOk

`func (o *TamSecurityAdvisory) GetDatePublishedOk() (*time.Time, bool)`

GetDatePublishedOk returns a tuple with the DatePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublished

`func (o *TamSecurityAdvisory) SetDatePublished(v time.Time)`

SetDatePublished sets DatePublished field to given value.

### HasDatePublished

`func (o *TamSecurityAdvisory) HasDatePublished() bool`

HasDatePublished returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TamSecurityAdvisory) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TamSecurityAdvisory) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TamSecurityAdvisory) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TamSecurityAdvisory) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetEnvironmentalScore

`func (o *TamSecurityAdvisory) GetEnvironmentalScore() float32`

GetEnvironmentalScore returns the EnvironmentalScore field if non-nil, zero value otherwise.

### GetEnvironmentalScoreOk

`func (o *TamSecurityAdvisory) GetEnvironmentalScoreOk() (*float32, bool)`

GetEnvironmentalScoreOk returns a tuple with the EnvironmentalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentalScore

`func (o *TamSecurityAdvisory) SetEnvironmentalScore(v float32)`

SetEnvironmentalScore sets EnvironmentalScore field to given value.

### HasEnvironmentalScore

`func (o *TamSecurityAdvisory) HasEnvironmentalScore() bool`

HasEnvironmentalScore returns a boolean if a field has been set.

### GetExternalUrl

`func (o *TamSecurityAdvisory) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *TamSecurityAdvisory) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *TamSecurityAdvisory) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *TamSecurityAdvisory) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetRecommendation

`func (o *TamSecurityAdvisory) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *TamSecurityAdvisory) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *TamSecurityAdvisory) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *TamSecurityAdvisory) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetStatus

`func (o *TamSecurityAdvisory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TamSecurityAdvisory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TamSecurityAdvisory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TamSecurityAdvisory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemporalScore

`func (o *TamSecurityAdvisory) GetTemporalScore() float32`

GetTemporalScore returns the TemporalScore field if non-nil, zero value otherwise.

### GetTemporalScoreOk

`func (o *TamSecurityAdvisory) GetTemporalScoreOk() (*float32, bool)`

GetTemporalScoreOk returns a tuple with the TemporalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalScore

`func (o *TamSecurityAdvisory) SetTemporalScore(v float32)`

SetTemporalScore sets TemporalScore field to given value.

### HasTemporalScore

`func (o *TamSecurityAdvisory) HasTemporalScore() bool`

HasTemporalScore returns a boolean if a field has been set.

### GetVersion

`func (o *TamSecurityAdvisory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TamSecurityAdvisory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TamSecurityAdvisory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TamSecurityAdvisory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkaround

`func (o *TamSecurityAdvisory) GetWorkaround() string`

GetWorkaround returns the Workaround field if non-nil, zero value otherwise.

### GetWorkaroundOk

`func (o *TamSecurityAdvisory) GetWorkaroundOk() (*string, bool)`

GetWorkaroundOk returns a tuple with the Workaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaround

`func (o *TamSecurityAdvisory) SetWorkaround(v string)`

SetWorkaround sets Workaround field to given value.

### HasWorkaround

`func (o *TamSecurityAdvisory) HasWorkaround() bool`

HasWorkaround returns a boolean if a field has been set.

### GetOrganization

`func (o *TamSecurityAdvisory) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TamSecurityAdvisory) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TamSecurityAdvisory) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TamSecurityAdvisory) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


