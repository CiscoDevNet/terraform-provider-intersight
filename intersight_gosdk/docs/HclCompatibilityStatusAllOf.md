# HclCompatibilityStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.CompatibilityStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.CompatibilityStatus"]
**ProfileList** | Pointer to [**[]HclHardwareCompatibilityProfile**](HclHardwareCompatibilityProfile.md) |  | [optional] 
**RequestType** | Pointer to **string** | Type of the request to be served. * &#x60;FillSupportedVersions&#x60; - Responds with the supported firmware and driver versions. The API doesn&#39;t expect firmware and driver versions to be passed in the request and ignores if passed. * &#x60;CheckCompatibility&#x60; - Checks the compatibility for the given firmware and driver versions. This request type expects the firmware and driver versions to filled and the service validates the values and responds back with the error codes. * &#x60;GetRecommendedDrivers&#x60; - Responds with the supported drivers. The API expects firmware version to be filled. The API populates driver ISO url for the given server model. Today the link is same for all servers managed by UCSM whereas it depends on the model for Standalone servers. | [optional] [default to "FillSupportedVersions"]

## Methods

### NewHclCompatibilityStatusAllOf

`func NewHclCompatibilityStatusAllOf(classId string, objectType string, ) *HclCompatibilityStatusAllOf`

NewHclCompatibilityStatusAllOf instantiates a new HclCompatibilityStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclCompatibilityStatusAllOfWithDefaults

`func NewHclCompatibilityStatusAllOfWithDefaults() *HclCompatibilityStatusAllOf`

NewHclCompatibilityStatusAllOfWithDefaults instantiates a new HclCompatibilityStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclCompatibilityStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclCompatibilityStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclCompatibilityStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclCompatibilityStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclCompatibilityStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclCompatibilityStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProfileList

`func (o *HclCompatibilityStatusAllOf) GetProfileList() []HclHardwareCompatibilityProfile`

GetProfileList returns the ProfileList field if non-nil, zero value otherwise.

### GetProfileListOk

`func (o *HclCompatibilityStatusAllOf) GetProfileListOk() (*[]HclHardwareCompatibilityProfile, bool)`

GetProfileListOk returns a tuple with the ProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileList

`func (o *HclCompatibilityStatusAllOf) SetProfileList(v []HclHardwareCompatibilityProfile)`

SetProfileList sets ProfileList field to given value.

### HasProfileList

`func (o *HclCompatibilityStatusAllOf) HasProfileList() bool`

HasProfileList returns a boolean if a field has been set.

### SetProfileListNil

`func (o *HclCompatibilityStatusAllOf) SetProfileListNil(b bool)`

 SetProfileListNil sets the value for ProfileList to be an explicit nil

### UnsetProfileList
`func (o *HclCompatibilityStatusAllOf) UnsetProfileList()`

UnsetProfileList ensures that no value is present for ProfileList, not even an explicit nil
### GetRequestType

`func (o *HclCompatibilityStatusAllOf) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HclCompatibilityStatusAllOf) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HclCompatibilityStatusAllOf) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *HclCompatibilityStatusAllOf) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


