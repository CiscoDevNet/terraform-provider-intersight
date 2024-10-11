# ComputeScrubOpConfigutation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.ScrubOpConfigutation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.ScrubOpConfigutation"]
**ScrubTargets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewComputeScrubOpConfigutation

`func NewComputeScrubOpConfigutation(classId string, objectType string, ) *ComputeScrubOpConfigutation`

NewComputeScrubOpConfigutation instantiates a new ComputeScrubOpConfigutation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeScrubOpConfigutationWithDefaults

`func NewComputeScrubOpConfigutationWithDefaults() *ComputeScrubOpConfigutation`

NewComputeScrubOpConfigutationWithDefaults instantiates a new ComputeScrubOpConfigutation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeScrubOpConfigutation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeScrubOpConfigutation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeScrubOpConfigutation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeScrubOpConfigutation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeScrubOpConfigutation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeScrubOpConfigutation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetScrubTargets

`func (o *ComputeScrubOpConfigutation) GetScrubTargets() []string`

GetScrubTargets returns the ScrubTargets field if non-nil, zero value otherwise.

### GetScrubTargetsOk

`func (o *ComputeScrubOpConfigutation) GetScrubTargetsOk() (*[]string, bool)`

GetScrubTargetsOk returns a tuple with the ScrubTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubTargets

`func (o *ComputeScrubOpConfigutation) SetScrubTargets(v []string)`

SetScrubTargets sets ScrubTargets field to given value.

### HasScrubTargets

`func (o *ComputeScrubOpConfigutation) HasScrubTargets() bool`

HasScrubTargets returns a boolean if a field has been set.

### SetScrubTargetsNil

`func (o *ComputeScrubOpConfigutation) SetScrubTargetsNil(b bool)`

 SetScrubTargetsNil sets the value for ScrubTargets to be an explicit nil

### UnsetScrubTargets
`func (o *ComputeScrubOpConfigutation) UnsetScrubTargets()`

UnsetScrubTargets ensures that no value is present for ScrubTargets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


