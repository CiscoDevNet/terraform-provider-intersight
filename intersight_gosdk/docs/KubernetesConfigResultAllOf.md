# KubernetesConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ConfigResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ConfigResult"]
**Profile** | Pointer to [**KubernetesNodeProfileRelationship**](KubernetesNodeProfileRelationship.md) |  | [optional] 
**ResultEntries** | Pointer to [**[]KubernetesConfigResultEntryRelationship**](KubernetesConfigResultEntryRelationship.md) | An array of relationships to kubernetesConfigResultEntry resources. | [optional] 

## Methods

### NewKubernetesConfigResultAllOf

`func NewKubernetesConfigResultAllOf(classId string, objectType string, ) *KubernetesConfigResultAllOf`

NewKubernetesConfigResultAllOf instantiates a new KubernetesConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigResultAllOfWithDefaults

`func NewKubernetesConfigResultAllOfWithDefaults() *KubernetesConfigResultAllOf`

NewKubernetesConfigResultAllOfWithDefaults instantiates a new KubernetesConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesConfigResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesConfigResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesConfigResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesConfigResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesConfigResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesConfigResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfile

`func (o *KubernetesConfigResultAllOf) GetProfile() KubernetesNodeProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *KubernetesConfigResultAllOf) GetProfileOk() (*KubernetesNodeProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *KubernetesConfigResultAllOf) SetProfile(v KubernetesNodeProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *KubernetesConfigResultAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResultEntries

`func (o *KubernetesConfigResultAllOf) GetResultEntries() []KubernetesConfigResultEntryRelationship`

GetResultEntries returns the ResultEntries field if non-nil, zero value otherwise.

### GetResultEntriesOk

`func (o *KubernetesConfigResultAllOf) GetResultEntriesOk() (*[]KubernetesConfigResultEntryRelationship, bool)`

GetResultEntriesOk returns a tuple with the ResultEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEntries

`func (o *KubernetesConfigResultAllOf) SetResultEntries(v []KubernetesConfigResultEntryRelationship)`

SetResultEntries sets ResultEntries field to given value.

### HasResultEntries

`func (o *KubernetesConfigResultAllOf) HasResultEntries() bool`

HasResultEntries returns a boolean if a field has been set.

### SetResultEntriesNil

`func (o *KubernetesConfigResultAllOf) SetResultEntriesNil(b bool)`

 SetResultEntriesNil sets the value for ResultEntries to be an explicit nil

### UnsetResultEntries
`func (o *KubernetesConfigResultAllOf) UnsetResultEntries()`

UnsetResultEntries ensures that no value is present for ResultEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


