# NiatelemetryAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Advisory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Advisory"]
**AdvisoryId** | Pointer to **string** | Unique identifier assigned to the fabric advisory. | [optional] [readonly] 
**AdvisoryScore** | Pointer to **int64** | Numeric score associated with the fabric advisory. | [optional] [readonly] 
**AdvisoryString** | Pointer to **string** | Human-readable description of the fabric advisory. | [optional] [readonly] 
**CompensatingControl** | Pointer to **string** | Compensating control status reported for the fabric advisory. * &#x60;available&#x60; - A compensating control is available for deployment. * &#x60;notAvailable&#x60; - No compensating control is available for the advisory. * &#x60;protected&#x60; - The advisory is protected by a deployed compensating control. * &#x60;monitoring&#x60; - The compensating control is monitoring relevant activity. * &#x60;partiallyApplied&#x60; - The compensating control is applied to only some affected nodes. * &#x60;deploymentFailed&#x60; - Deployment of the compensating control did not complete successfully. * &#x60;disabled&#x60; - The compensating control is currently disabled. * &#x60;queued&#x60; - The compensating control is queued for deployment. * &#x60;deploying&#x60; - The compensating control deployment is currently in progress. * &#x60;outOfSync&#x60; - The deployed compensating control is not synchronized with its intended configuration. * &#x60;redeployRequired&#x60; - The compensating control must be redeployed to restore the intended protection. | [optional] [readonly] [default to "available"]
**CveIds** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **float32** | CVSS base score associated with the fabric advisory. | [optional] [readonly] 
**EntityName** | Pointer to **string** | Entity name associated with the fabric advisory. | [optional] [readonly] 
**FabricName** | Pointer to **string** | Fabric name associated with the advisory. | [optional] [readonly] 
**ImpactedNodes** | Pointer to [**[]NiatelemetryImpactedNodes**](NiatelemetryImpactedNodes.md) |  | [optional] 
**NodeNames** | Pointer to **[]string** |  | [optional] 
**ProtectedNodes** | Pointer to **[]string** |  | [optional] 
**Recommendations** | Pointer to [**[]NiatelemetryAdvisoryRecommendation**](NiatelemetryAdvisoryRecommendation.md) |  | [optional] 
**SecurityAdvisoryId** | Pointer to **string** | Security advisory identifier associated with the fabric advisory. | [optional] [readonly] 
**SecurityStatus** | Pointer to **string** | Security status reported for the fabric advisory. * &#x60;active&#x60; - The advisory is active and requires evaluation. * &#x60;protected&#x60; - The advisory is protected by a compensating control. * &#x60;partiallyProtected&#x60; - The advisory is partially protected by a compensating control. * &#x60;monitoring&#x60; - The advisory is being monitored for relevant activity. * &#x60;cleared&#x60; - The advisory has been cleared and no longer requires action. | [optional] [readonly] [default to "active"]
**Severity** | Pointer to **string** | Severity level of the advisory. * &#x60;critical&#x60; - The advisory has a critical severity level. * &#x60;major&#x60; - The advisory has a major severity level. * &#x60;minor&#x60; - The advisory has a minor severity level. * &#x60;warning&#x60; - The advisory has a warning severity level. | [optional] [readonly] [default to "critical"]
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryAdvisory

`func NewNiatelemetryAdvisory(classId string, objectType string, ) *NiatelemetryAdvisory`

NewNiatelemetryAdvisory instantiates a new NiatelemetryAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAdvisoryWithDefaults

`func NewNiatelemetryAdvisoryWithDefaults() *NiatelemetryAdvisory`

NewNiatelemetryAdvisoryWithDefaults instantiates a new NiatelemetryAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryAdvisory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryAdvisory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryAdvisory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryAdvisory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryAdvisory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryAdvisory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvisoryId

`func (o *NiatelemetryAdvisory) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *NiatelemetryAdvisory) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *NiatelemetryAdvisory) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *NiatelemetryAdvisory) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetAdvisoryScore

`func (o *NiatelemetryAdvisory) GetAdvisoryScore() int64`

GetAdvisoryScore returns the AdvisoryScore field if non-nil, zero value otherwise.

### GetAdvisoryScoreOk

`func (o *NiatelemetryAdvisory) GetAdvisoryScoreOk() (*int64, bool)`

GetAdvisoryScoreOk returns a tuple with the AdvisoryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryScore

`func (o *NiatelemetryAdvisory) SetAdvisoryScore(v int64)`

SetAdvisoryScore sets AdvisoryScore field to given value.

### HasAdvisoryScore

`func (o *NiatelemetryAdvisory) HasAdvisoryScore() bool`

HasAdvisoryScore returns a boolean if a field has been set.

### GetAdvisoryString

`func (o *NiatelemetryAdvisory) GetAdvisoryString() string`

GetAdvisoryString returns the AdvisoryString field if non-nil, zero value otherwise.

### GetAdvisoryStringOk

`func (o *NiatelemetryAdvisory) GetAdvisoryStringOk() (*string, bool)`

GetAdvisoryStringOk returns a tuple with the AdvisoryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryString

`func (o *NiatelemetryAdvisory) SetAdvisoryString(v string)`

SetAdvisoryString sets AdvisoryString field to given value.

### HasAdvisoryString

`func (o *NiatelemetryAdvisory) HasAdvisoryString() bool`

HasAdvisoryString returns a boolean if a field has been set.

### GetCompensatingControl

`func (o *NiatelemetryAdvisory) GetCompensatingControl() string`

GetCompensatingControl returns the CompensatingControl field if non-nil, zero value otherwise.

### GetCompensatingControlOk

`func (o *NiatelemetryAdvisory) GetCompensatingControlOk() (*string, bool)`

GetCompensatingControlOk returns a tuple with the CompensatingControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensatingControl

`func (o *NiatelemetryAdvisory) SetCompensatingControl(v string)`

SetCompensatingControl sets CompensatingControl field to given value.

### HasCompensatingControl

`func (o *NiatelemetryAdvisory) HasCompensatingControl() bool`

HasCompensatingControl returns a boolean if a field has been set.

### GetCveIds

`func (o *NiatelemetryAdvisory) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *NiatelemetryAdvisory) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *NiatelemetryAdvisory) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *NiatelemetryAdvisory) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### SetCveIdsNil

`func (o *NiatelemetryAdvisory) SetCveIdsNil(b bool)`

 SetCveIdsNil sets the value for CveIds to be an explicit nil

### UnsetCveIds
`func (o *NiatelemetryAdvisory) UnsetCveIds()`

UnsetCveIds ensures that no value is present for CveIds, not even an explicit nil
### GetCvss

`func (o *NiatelemetryAdvisory) GetCvss() float32`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *NiatelemetryAdvisory) GetCvssOk() (*float32, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *NiatelemetryAdvisory) SetCvss(v float32)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *NiatelemetryAdvisory) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetEntityName

`func (o *NiatelemetryAdvisory) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *NiatelemetryAdvisory) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *NiatelemetryAdvisory) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *NiatelemetryAdvisory) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryAdvisory) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryAdvisory) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryAdvisory) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryAdvisory) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetImpactedNodes

`func (o *NiatelemetryAdvisory) GetImpactedNodes() []NiatelemetryImpactedNodes`

GetImpactedNodes returns the ImpactedNodes field if non-nil, zero value otherwise.

### GetImpactedNodesOk

`func (o *NiatelemetryAdvisory) GetImpactedNodesOk() (*[]NiatelemetryImpactedNodes, bool)`

GetImpactedNodesOk returns a tuple with the ImpactedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedNodes

`func (o *NiatelemetryAdvisory) SetImpactedNodes(v []NiatelemetryImpactedNodes)`

SetImpactedNodes sets ImpactedNodes field to given value.

### HasImpactedNodes

`func (o *NiatelemetryAdvisory) HasImpactedNodes() bool`

HasImpactedNodes returns a boolean if a field has been set.

### SetImpactedNodesNil

`func (o *NiatelemetryAdvisory) SetImpactedNodesNil(b bool)`

 SetImpactedNodesNil sets the value for ImpactedNodes to be an explicit nil

### UnsetImpactedNodes
`func (o *NiatelemetryAdvisory) UnsetImpactedNodes()`

UnsetImpactedNodes ensures that no value is present for ImpactedNodes, not even an explicit nil
### GetNodeNames

`func (o *NiatelemetryAdvisory) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *NiatelemetryAdvisory) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *NiatelemetryAdvisory) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *NiatelemetryAdvisory) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### SetNodeNamesNil

`func (o *NiatelemetryAdvisory) SetNodeNamesNil(b bool)`

 SetNodeNamesNil sets the value for NodeNames to be an explicit nil

### UnsetNodeNames
`func (o *NiatelemetryAdvisory) UnsetNodeNames()`

UnsetNodeNames ensures that no value is present for NodeNames, not even an explicit nil
### GetProtectedNodes

`func (o *NiatelemetryAdvisory) GetProtectedNodes() []string`

GetProtectedNodes returns the ProtectedNodes field if non-nil, zero value otherwise.

### GetProtectedNodesOk

`func (o *NiatelemetryAdvisory) GetProtectedNodesOk() (*[]string, bool)`

GetProtectedNodesOk returns a tuple with the ProtectedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNodes

`func (o *NiatelemetryAdvisory) SetProtectedNodes(v []string)`

SetProtectedNodes sets ProtectedNodes field to given value.

### HasProtectedNodes

`func (o *NiatelemetryAdvisory) HasProtectedNodes() bool`

HasProtectedNodes returns a boolean if a field has been set.

### SetProtectedNodesNil

`func (o *NiatelemetryAdvisory) SetProtectedNodesNil(b bool)`

 SetProtectedNodesNil sets the value for ProtectedNodes to be an explicit nil

### UnsetProtectedNodes
`func (o *NiatelemetryAdvisory) UnsetProtectedNodes()`

UnsetProtectedNodes ensures that no value is present for ProtectedNodes, not even an explicit nil
### GetRecommendations

`func (o *NiatelemetryAdvisory) GetRecommendations() []NiatelemetryAdvisoryRecommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *NiatelemetryAdvisory) GetRecommendationsOk() (*[]NiatelemetryAdvisoryRecommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *NiatelemetryAdvisory) SetRecommendations(v []NiatelemetryAdvisoryRecommendation)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *NiatelemetryAdvisory) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### SetRecommendationsNil

`func (o *NiatelemetryAdvisory) SetRecommendationsNil(b bool)`

 SetRecommendationsNil sets the value for Recommendations to be an explicit nil

### UnsetRecommendations
`func (o *NiatelemetryAdvisory) UnsetRecommendations()`

UnsetRecommendations ensures that no value is present for Recommendations, not even an explicit nil
### GetSecurityAdvisoryId

`func (o *NiatelemetryAdvisory) GetSecurityAdvisoryId() string`

GetSecurityAdvisoryId returns the SecurityAdvisoryId field if non-nil, zero value otherwise.

### GetSecurityAdvisoryIdOk

`func (o *NiatelemetryAdvisory) GetSecurityAdvisoryIdOk() (*string, bool)`

GetSecurityAdvisoryIdOk returns a tuple with the SecurityAdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAdvisoryId

`func (o *NiatelemetryAdvisory) SetSecurityAdvisoryId(v string)`

SetSecurityAdvisoryId sets SecurityAdvisoryId field to given value.

### HasSecurityAdvisoryId

`func (o *NiatelemetryAdvisory) HasSecurityAdvisoryId() bool`

HasSecurityAdvisoryId returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *NiatelemetryAdvisory) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *NiatelemetryAdvisory) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *NiatelemetryAdvisory) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *NiatelemetryAdvisory) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *NiatelemetryAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NiatelemetryAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NiatelemetryAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NiatelemetryAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryAdvisory) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryAdvisory) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryAdvisory) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryAdvisory) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryAdvisory) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryAdvisory) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


