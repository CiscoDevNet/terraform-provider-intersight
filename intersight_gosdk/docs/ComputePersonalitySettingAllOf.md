# ComputePersonalitySettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PersonalitySetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PersonalitySetting"]
**AdditionalInformation** | Pointer to **interface{}** | Additional information to be set along with the personality value. This can include information like the hypervisor type, last update time etc.. | [optional] 
**Personality** | Pointer to **string** | The personality value that is to be set for the server. | [optional] 

## Methods

### NewComputePersonalitySettingAllOf

`func NewComputePersonalitySettingAllOf(classId string, objectType string, ) *ComputePersonalitySettingAllOf`

NewComputePersonalitySettingAllOf instantiates a new ComputePersonalitySettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePersonalitySettingAllOfWithDefaults

`func NewComputePersonalitySettingAllOfWithDefaults() *ComputePersonalitySettingAllOf`

NewComputePersonalitySettingAllOfWithDefaults instantiates a new ComputePersonalitySettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePersonalitySettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePersonalitySettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePersonalitySettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePersonalitySettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePersonalitySettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePersonalitySettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdditionalInformation

`func (o *ComputePersonalitySettingAllOf) GetAdditionalInformation() interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *ComputePersonalitySettingAllOf) GetAdditionalInformationOk() (*interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *ComputePersonalitySettingAllOf) SetAdditionalInformation(v interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *ComputePersonalitySettingAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### SetAdditionalInformationNil

`func (o *ComputePersonalitySettingAllOf) SetAdditionalInformationNil(b bool)`

 SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil

### UnsetAdditionalInformation
`func (o *ComputePersonalitySettingAllOf) UnsetAdditionalInformation()`

UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil
### GetPersonality

`func (o *ComputePersonalitySettingAllOf) GetPersonality() string`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *ComputePersonalitySettingAllOf) GetPersonalityOk() (*string, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *ComputePersonalitySettingAllOf) SetPersonality(v string)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *ComputePersonalitySettingAllOf) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


