# RecommendationSoftwareItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.SoftwareItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.SoftwareItem"]
**Message** | Pointer to **string** | The user visible message which informs user of the type of software recommendation. | [optional] [readonly] 
**RecommendationType** | Pointer to **string** | The software-recommendation type, for example, HXDP version, HyperV or ESXi version, etc. * &#x60;None&#x60; - The Enum value None represents the default software recommendation value. * &#x60;HXDPVersion&#x60; - The Enum value HXDPVersion represents that the software recommendation is to upgrade the HyperFlex Data Platform build version. * &#x60;NodeRatioLicense&#x60; - The Enum value NodeRatioLicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using 1:2 converged to compute node ratio limits. * &#x60;DCNoFILicense&#x60; - The Enum value DCNoFILicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using DC-No-FI limits. * &#x60;LAZExistingStatus&#x60; - The Enum value LAZExistingStatus represents that the software recommendation indicates that the HyperFlex cluster might have LAZ enabled. * &#x60;LAZNewStatus&#x60; - The Enum value LAZNewStatus represents that the software recommendation is to enable LAZ with expansion on the HyperFlex Cluster. * &#x60;EVCStatus&#x60; - The Enum value EVCStatus represents that the software recommendation is to enable Enhanced VMotion on the HypeFlex Cluster. | [optional] [readonly] [default to "None"]
**ClusterExpansion** | Pointer to [**RecommendationClusterExpansionRelationship**](RecommendationClusterExpansionRelationship.md) |  | [optional] 

## Methods

### NewRecommendationSoftwareItemAllOf

`func NewRecommendationSoftwareItemAllOf(classId string, objectType string, ) *RecommendationSoftwareItemAllOf`

NewRecommendationSoftwareItemAllOf instantiates a new RecommendationSoftwareItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationSoftwareItemAllOfWithDefaults

`func NewRecommendationSoftwareItemAllOfWithDefaults() *RecommendationSoftwareItemAllOf`

NewRecommendationSoftwareItemAllOfWithDefaults instantiates a new RecommendationSoftwareItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationSoftwareItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationSoftwareItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationSoftwareItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationSoftwareItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationSoftwareItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationSoftwareItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *RecommendationSoftwareItemAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecommendationSoftwareItemAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecommendationSoftwareItemAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RecommendationSoftwareItemAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecommendationType

`func (o *RecommendationSoftwareItemAllOf) GetRecommendationType() string`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *RecommendationSoftwareItemAllOf) GetRecommendationTypeOk() (*string, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *RecommendationSoftwareItemAllOf) SetRecommendationType(v string)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *RecommendationSoftwareItemAllOf) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetClusterExpansion

`func (o *RecommendationSoftwareItemAllOf) GetClusterExpansion() RecommendationClusterExpansionRelationship`

GetClusterExpansion returns the ClusterExpansion field if non-nil, zero value otherwise.

### GetClusterExpansionOk

`func (o *RecommendationSoftwareItemAllOf) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool)`

GetClusterExpansionOk returns a tuple with the ClusterExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExpansion

`func (o *RecommendationSoftwareItemAllOf) SetClusterExpansion(v RecommendationClusterExpansionRelationship)`

SetClusterExpansion sets ClusterExpansion field to given value.

### HasClusterExpansion

`func (o *RecommendationSoftwareItemAllOf) HasClusterExpansion() bool`

HasClusterExpansion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


