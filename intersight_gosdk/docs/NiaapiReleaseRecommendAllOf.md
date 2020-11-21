# NiaapiReleaseRecommendAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Cll** | Pointer to **string** | Current long-lived release. | [optional] 
**Crr** | Pointer to **string** | Customer recommended releases. | [optional] 
**Pid** | Pointer to **string** | Hardware model identificator. | [optional] 
**Ull** | Pointer to **string** | Upcoming long-lived release. | [optional] 

## Methods

### NewNiaapiReleaseRecommendAllOf

`func NewNiaapiReleaseRecommendAllOf(classId string, objectType string, ) *NiaapiReleaseRecommendAllOf`

NewNiaapiReleaseRecommendAllOf instantiates a new NiaapiReleaseRecommendAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiReleaseRecommendAllOfWithDefaults

`func NewNiaapiReleaseRecommendAllOfWithDefaults() *NiaapiReleaseRecommendAllOf`

NewNiaapiReleaseRecommendAllOfWithDefaults instantiates a new NiaapiReleaseRecommendAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiReleaseRecommendAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiReleaseRecommendAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiReleaseRecommendAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiReleaseRecommendAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiReleaseRecommendAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiReleaseRecommendAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCll

`func (o *NiaapiReleaseRecommendAllOf) GetCll() string`

GetCll returns the Cll field if non-nil, zero value otherwise.

### GetCllOk

`func (o *NiaapiReleaseRecommendAllOf) GetCllOk() (*string, bool)`

GetCllOk returns a tuple with the Cll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCll

`func (o *NiaapiReleaseRecommendAllOf) SetCll(v string)`

SetCll sets Cll field to given value.

### HasCll

`func (o *NiaapiReleaseRecommendAllOf) HasCll() bool`

HasCll returns a boolean if a field has been set.

### GetCrr

`func (o *NiaapiReleaseRecommendAllOf) GetCrr() string`

GetCrr returns the Crr field if non-nil, zero value otherwise.

### GetCrrOk

`func (o *NiaapiReleaseRecommendAllOf) GetCrrOk() (*string, bool)`

GetCrrOk returns a tuple with the Crr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrr

`func (o *NiaapiReleaseRecommendAllOf) SetCrr(v string)`

SetCrr sets Crr field to given value.

### HasCrr

`func (o *NiaapiReleaseRecommendAllOf) HasCrr() bool`

HasCrr returns a boolean if a field has been set.

### GetPid

`func (o *NiaapiReleaseRecommendAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *NiaapiReleaseRecommendAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *NiaapiReleaseRecommendAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *NiaapiReleaseRecommendAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetUll

`func (o *NiaapiReleaseRecommendAllOf) GetUll() string`

GetUll returns the Ull field if non-nil, zero value otherwise.

### GetUllOk

`func (o *NiaapiReleaseRecommendAllOf) GetUllOk() (*string, bool)`

GetUllOk returns a tuple with the Ull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUll

`func (o *NiaapiReleaseRecommendAllOf) SetUll(v string)`

SetUll sets Ull field to given value.

### HasUll

`func (o *NiaapiReleaseRecommendAllOf) HasUll() bool`

HasUll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


