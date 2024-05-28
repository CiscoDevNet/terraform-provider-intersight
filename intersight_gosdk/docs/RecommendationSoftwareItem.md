# RecommendationSoftwareItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.SoftwareItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.SoftwareItem"]
**Message** | Pointer to **string** | The user visible message which informs user of the type of software recommendation. | [optional] [readonly] 
**RecommendationType** | Pointer to **string** | The software-recommendation type, for example, HXDP version, HyperV or ESXi version, etc. * &#x60;None&#x60; - The Enum value None represents the default software recommendation value. * &#x60;HXDPVersion&#x60; - The Enum value HXDPVersion represents that the software recommendation is to upgrade the HyperFlex Data Platform build version. * &#x60;NodeRatioLicense&#x60; - The Enum value NodeRatioLicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using 1:2 converged to compute node ratio limits. * &#x60;NodeRatioEvalLicense&#x60; - The Enum value NodeRatioEvalLicense represents that the software recommendation is to upgrade the Hyperflex Data Platform license to Data Center Premier, after the evaluation period ends. * &#x60;DCNoFILicense&#x60; - The Enum value DCNoFILicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using DC-No-FI limits. * &#x60;DcNoFIEvalLicense&#x60; - The Enum value DcNoFIEvalLicense represents that the software recommendation is to upgrade the Hyperflex Data Platform license to Hyperflex Data Center Advantage or Data Center Premier license, after evaluation period ends. * &#x60;LAZExistingStatus&#x60; - The Enum value LAZExistingStatus represents that the software recommendation is to indicate HyperFlex cluster might have LAZ enabled. * &#x60;LAZNewStatus&#x60; - The Enum value LAZNewStatus represents that the software recommendation is to enable LAZ with expansion on the HyperFlex Cluster. * &#x60;EVCStatus&#x60; - The Enum value EVCStatus represents that the software recommendation is to enable Enhanced VMotion on the HypeFlex Cluster. | [optional] [readonly] [default to "None"]
**ClusterExpansion** | Pointer to [**NullableRecommendationClusterExpansionRelationship**](RecommendationClusterExpansionRelationship.md) |  | [optional] 

## Methods

### NewRecommendationSoftwareItem

`func NewRecommendationSoftwareItem(classId string, objectType string, ) *RecommendationSoftwareItem`

NewRecommendationSoftwareItem instantiates a new RecommendationSoftwareItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationSoftwareItemWithDefaults

`func NewRecommendationSoftwareItemWithDefaults() *RecommendationSoftwareItem`

NewRecommendationSoftwareItemWithDefaults instantiates a new RecommendationSoftwareItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationSoftwareItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationSoftwareItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationSoftwareItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationSoftwareItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationSoftwareItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationSoftwareItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *RecommendationSoftwareItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecommendationSoftwareItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecommendationSoftwareItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RecommendationSoftwareItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecommendationType

`func (o *RecommendationSoftwareItem) GetRecommendationType() string`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *RecommendationSoftwareItem) GetRecommendationTypeOk() (*string, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *RecommendationSoftwareItem) SetRecommendationType(v string)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *RecommendationSoftwareItem) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetClusterExpansion

`func (o *RecommendationSoftwareItem) GetClusterExpansion() RecommendationClusterExpansionRelationship`

GetClusterExpansion returns the ClusterExpansion field if non-nil, zero value otherwise.

### GetClusterExpansionOk

`func (o *RecommendationSoftwareItem) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool)`

GetClusterExpansionOk returns a tuple with the ClusterExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExpansion

`func (o *RecommendationSoftwareItem) SetClusterExpansion(v RecommendationClusterExpansionRelationship)`

SetClusterExpansion sets ClusterExpansion field to given value.

### HasClusterExpansion

`func (o *RecommendationSoftwareItem) HasClusterExpansion() bool`

HasClusterExpansion returns a boolean if a field has been set.

### SetClusterExpansionNil

`func (o *RecommendationSoftwareItem) SetClusterExpansionNil(b bool)`

 SetClusterExpansionNil sets the value for ClusterExpansion to be an explicit nil

### UnsetClusterExpansion
`func (o *RecommendationSoftwareItem) UnsetClusterExpansion()`

UnsetClusterExpansion ensures that no value is present for ClusterExpansion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


