# TamAdvisoryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]TamAction**](tam.Action.md) |  | [optional] 
**AdvisoryDetails** | Pointer to [**TamBaseAdvisoryDetails**](tam.BaseAdvisoryDetails.md) |  | [optional] 
**AdvisoryId** | Pointer to **string** | Cisco generated identifier for the published security advisory. | [optional] 
**ApiDataSources** | Pointer to [**[]TamApiDataSource**](tam.ApiDataSource.md) |  | [optional] 
**DatePublished** | Pointer to [**time.Time**](time.Time.md) | Date when the security advisory was first published by Cisco. | [optional] 
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date when the security advisory was last updated by Cisco. | [optional] 
**ExternalUrl** | Pointer to **string** | A link to an external URL describing security Advisory in more details. | [optional] 
**Recommendation** | Pointer to **string** | Recommended action to resolve the security advisory. | [optional] 
**Type** | Pointer to **string** | The type (field notice, security advisory etc.) of Intersight advisory. * &#x60;securityAdvisory&#x60; - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * &#x60;fieldNotice&#x60; - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). | [optional] [default to "securityAdvisory"]
**Version** | Pointer to **string** | Cisco assigned advisory version after latest revision. | [optional] 
**Workaround** | Pointer to **string** | Workarounds available for the advisory. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryDefinition

`func NewTamAdvisoryDefinition() *TamAdvisoryDefinition`

NewTamAdvisoryDefinition instantiates a new TamAdvisoryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryDefinitionWithDefaults

`func NewTamAdvisoryDefinitionWithDefaults() *TamAdvisoryDefinition`

NewTamAdvisoryDefinitionWithDefaults instantiates a new TamAdvisoryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TamAdvisoryDefinition) GetActions() []TamAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TamAdvisoryDefinition) GetActionsOk() (*[]TamAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TamAdvisoryDefinition) SetActions(v []TamAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TamAdvisoryDefinition) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdvisoryDetails

`func (o *TamAdvisoryDefinition) GetAdvisoryDetails() TamBaseAdvisoryDetails`

GetAdvisoryDetails returns the AdvisoryDetails field if non-nil, zero value otherwise.

### GetAdvisoryDetailsOk

`func (o *TamAdvisoryDefinition) GetAdvisoryDetailsOk() (*TamBaseAdvisoryDetails, bool)`

GetAdvisoryDetailsOk returns a tuple with the AdvisoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryDetails

`func (o *TamAdvisoryDefinition) SetAdvisoryDetails(v TamBaseAdvisoryDetails)`

SetAdvisoryDetails sets AdvisoryDetails field to given value.

### HasAdvisoryDetails

`func (o *TamAdvisoryDefinition) HasAdvisoryDetails() bool`

HasAdvisoryDetails returns a boolean if a field has been set.

### GetAdvisoryId

`func (o *TamAdvisoryDefinition) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *TamAdvisoryDefinition) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *TamAdvisoryDefinition) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *TamAdvisoryDefinition) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetApiDataSources

`func (o *TamAdvisoryDefinition) GetApiDataSources() []TamApiDataSource`

GetApiDataSources returns the ApiDataSources field if non-nil, zero value otherwise.

### GetApiDataSourcesOk

`func (o *TamAdvisoryDefinition) GetApiDataSourcesOk() (*[]TamApiDataSource, bool)`

GetApiDataSourcesOk returns a tuple with the ApiDataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiDataSources

`func (o *TamAdvisoryDefinition) SetApiDataSources(v []TamApiDataSource)`

SetApiDataSources sets ApiDataSources field to given value.

### HasApiDataSources

`func (o *TamAdvisoryDefinition) HasApiDataSources() bool`

HasApiDataSources returns a boolean if a field has been set.

### GetDatePublished

`func (o *TamAdvisoryDefinition) GetDatePublished() time.Time`

GetDatePublished returns the DatePublished field if non-nil, zero value otherwise.

### GetDatePublishedOk

`func (o *TamAdvisoryDefinition) GetDatePublishedOk() (*time.Time, bool)`

GetDatePublishedOk returns a tuple with the DatePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublished

`func (o *TamAdvisoryDefinition) SetDatePublished(v time.Time)`

SetDatePublished sets DatePublished field to given value.

### HasDatePublished

`func (o *TamAdvisoryDefinition) HasDatePublished() bool`

HasDatePublished returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TamAdvisoryDefinition) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TamAdvisoryDefinition) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TamAdvisoryDefinition) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TamAdvisoryDefinition) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetExternalUrl

`func (o *TamAdvisoryDefinition) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *TamAdvisoryDefinition) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *TamAdvisoryDefinition) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *TamAdvisoryDefinition) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetRecommendation

`func (o *TamAdvisoryDefinition) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *TamAdvisoryDefinition) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *TamAdvisoryDefinition) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *TamAdvisoryDefinition) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetType

`func (o *TamAdvisoryDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamAdvisoryDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamAdvisoryDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamAdvisoryDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *TamAdvisoryDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TamAdvisoryDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TamAdvisoryDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TamAdvisoryDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkaround

`func (o *TamAdvisoryDefinition) GetWorkaround() string`

GetWorkaround returns the Workaround field if non-nil, zero value otherwise.

### GetWorkaroundOk

`func (o *TamAdvisoryDefinition) GetWorkaroundOk() (*string, bool)`

GetWorkaroundOk returns a tuple with the Workaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaround

`func (o *TamAdvisoryDefinition) SetWorkaround(v string)`

SetWorkaround sets Workaround field to given value.

### HasWorkaround

`func (o *TamAdvisoryDefinition) HasWorkaround() bool`

HasWorkaround returns a boolean if a field has been set.

### GetOrganization

`func (o *TamAdvisoryDefinition) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TamAdvisoryDefinition) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TamAdvisoryDefinition) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *TamAdvisoryDefinition) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


