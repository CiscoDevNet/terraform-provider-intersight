# KubernetesConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ConfigResult"]
**Profile** | Pointer to [**KubernetesNodeProfileRelationship**](KubernetesNodeProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]KubernetesConfigResultEntryRelationship**](KubernetesConfigResultEntryRelationship.md) | An array of relationships to kubernetesConfigResultEntry resources. | [optional] 

## Methods

### NewKubernetesConfigResult

`func NewKubernetesConfigResult(classId string, objectType string, ) *KubernetesConfigResult`

NewKubernetesConfigResult instantiates a new KubernetesConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigResultWithDefaults

`func NewKubernetesConfigResultWithDefaults() *KubernetesConfigResult`

NewKubernetesConfigResultWithDefaults instantiates a new KubernetesConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesConfigResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesConfigResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesConfigResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesConfigResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesConfigResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesConfigResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfile

`func (o *KubernetesConfigResult) GetProfile() KubernetesNodeProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *KubernetesConfigResult) GetProfileOk() (*KubernetesNodeProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *KubernetesConfigResult) SetProfile(v KubernetesNodeProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *KubernetesConfigResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *KubernetesConfigResult) GetResultEntries() []KubernetesConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *KubernetesConfigResult) GetResultEntriesOk() (*[]KubernetesConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *KubernetesConfigResult) SetResultEntries(v []KubernetesConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *KubernetesConfigResult) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *KubernetesConfigResult) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *KubernetesConfigResult) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


