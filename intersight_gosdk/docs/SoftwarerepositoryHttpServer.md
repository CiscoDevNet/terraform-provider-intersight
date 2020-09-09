# SoftwarerepositoryHttpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**LocationLink** | Pointer to **string** | HTTP/HTTPS link to the image. Accepted formats are HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. | [optional] 
**Password** | Pointer to **string** | Password as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories. | [optional] 
**Username** | Pointer to **string** | Username as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories. | [optional] 

## Methods

### NewSoftwarerepositoryHttpServer

`func NewSoftwarerepositoryHttpServer() *SoftwarerepositoryHttpServer`

NewSoftwarerepositoryHttpServer instantiates a new SoftwarerepositoryHttpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryHttpServerWithDefaults

`func NewSoftwarerepositoryHttpServerWithDefaults() *SoftwarerepositoryHttpServer`

NewSoftwarerepositoryHttpServerWithDefaults instantiates a new SoftwarerepositoryHttpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPasswordSet

`func (o *SoftwarerepositoryHttpServer) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *SoftwarerepositoryHttpServer) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *SoftwarerepositoryHttpServer) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *SoftwarerepositoryHttpServer) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocationLink

`func (o *SoftwarerepositoryHttpServer) GetLocationLink() string`

GetLocationLink returns the LocationLink field if non-nil, zero value otherwise.

### GetLocationLinkOk

`func (o *SoftwarerepositoryHttpServer) GetLocationLinkOk() (*string, bool)`

GetLocationLinkOk returns a tuple with the LocationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLink

`func (o *SoftwarerepositoryHttpServer) SetLocationLink(v string)`

SetLocationLink sets LocationLink field to given value.

### HasLocationLink

`func (o *SoftwarerepositoryHttpServer) HasLocationLink() bool`

HasLocationLink returns a boolean if a field has been set.

### GetPassword

`func (o *SoftwarerepositoryHttpServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SoftwarerepositoryHttpServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SoftwarerepositoryHttpServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SoftwarerepositoryHttpServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *SoftwarerepositoryHttpServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SoftwarerepositoryHttpServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SoftwarerepositoryHttpServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SoftwarerepositoryHttpServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


