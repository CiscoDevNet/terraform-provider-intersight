# SoftwarerepositoryHttpServerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.HttpServer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.HttpServer"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats are HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. | [optional] 
**Password** | Pointer to **string** | Password as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories. | [optional] 
**Username** | Pointer to **string** | Username as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories. | [optional] 

## Methods

### NewSoftwarerepositoryHttpServerAllOf

`func NewSoftwarerepositoryHttpServerAllOf(classId string, objectType string, ) *SoftwarerepositoryHttpServerAllOf`

NewSoftwarerepositoryHttpServerAllOf instantiates a new SoftwarerepositoryHttpServerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryHttpServerAllOfWithDefaults

`func NewSoftwarerepositoryHttpServerAllOfWithDefaults() *SoftwarerepositoryHttpServerAllOf`

NewSoftwarerepositoryHttpServerAllOfWithDefaults instantiates a new SoftwarerepositoryHttpServerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryHttpServerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryHttpServerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryHttpServerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryHttpServerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *SoftwarerepositoryHttpServerAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SoftwarerepositoryHttpServerAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SoftwarerepositoryHttpServerAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocationLink

`func (o *SoftwarerepositoryHttpServerAllOf) GetLocationLink() string`

GetLocationLink returns the LocationLink field if non-nil, zero value otherwise.

### GetLocationLinkOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetLocationLinkOk() (*string, bool)`

GetLocationLinkOk returns a tuple with the LocationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLink

`func (o *SoftwarerepositoryHttpServerAllOf) SetLocationLink(v string)`

SetLocationLink sets LocationLink field to given value.

### HasLocationLink

`func (o *SoftwarerepositoryHttpServerAllOf) HasLocationLink() bool`

HasLocationLink returns a boolean if a field has been set.

### GetPassword

`func (o *SoftwarerepositoryHttpServerAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SoftwarerepositoryHttpServerAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SoftwarerepositoryHttpServerAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *SoftwarerepositoryHttpServerAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SoftwarerepositoryHttpServerAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SoftwarerepositoryHttpServerAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SoftwarerepositoryHttpServerAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


