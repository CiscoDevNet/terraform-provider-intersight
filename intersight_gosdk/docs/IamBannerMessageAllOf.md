# IamBannerMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.BannerMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.BannerMessage"]
**BannerContents** | Pointer to **string** | Contents of the banner message. | [optional] 
**BannerDisplay** | Pointer to **bool** | Whether or not to display the banner message. | [optional] 
**BannerTitle** | Pointer to **string** | Title of the banner message. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamBannerMessageAllOf

`func NewIamBannerMessageAllOf(classId string, objectType string, ) *IamBannerMessageAllOf`

NewIamBannerMessageAllOf instantiates a new IamBannerMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamBannerMessageAllOfWithDefaults

`func NewIamBannerMessageAllOfWithDefaults() *IamBannerMessageAllOf`

NewIamBannerMessageAllOfWithDefaults instantiates a new IamBannerMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamBannerMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamBannerMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamBannerMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamBannerMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamBannerMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamBannerMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBannerContents

`func (o *IamBannerMessageAllOf) GetBannerContents() string`

GetBannerContents returns the BannerContents field if non-nil, zero value otherwise.

### GetBannerContentsOk

`func (o *IamBannerMessageAllOf) GetBannerContentsOk() (*string, bool)`

GetBannerContentsOk returns a tuple with the BannerContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerContents

`func (o *IamBannerMessageAllOf) SetBannerContents(v string)`

SetBannerContents sets BannerContents field to given value.

### HasBannerContents

`func (o *IamBannerMessageAllOf) HasBannerContents() bool`

HasBannerContents returns a boolean if a field has been set.

### GetBannerDisplay

`func (o *IamBannerMessageAllOf) GetBannerDisplay() bool`

GetBannerDisplay returns the BannerDisplay field if non-nil, zero value otherwise.

### GetBannerDisplayOk

`func (o *IamBannerMessageAllOf) GetBannerDisplayOk() (*bool, bool)`

GetBannerDisplayOk returns a tuple with the BannerDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerDisplay

`func (o *IamBannerMessageAllOf) SetBannerDisplay(v bool)`

SetBannerDisplay sets BannerDisplay field to given value.

### HasBannerDisplay

`func (o *IamBannerMessageAllOf) HasBannerDisplay() bool`

HasBannerDisplay returns a boolean if a field has been set.

### GetBannerTitle

`func (o *IamBannerMessageAllOf) GetBannerTitle() string`

GetBannerTitle returns the BannerTitle field if non-nil, zero value otherwise.

### GetBannerTitleOk

`func (o *IamBannerMessageAllOf) GetBannerTitleOk() (*string, bool)`

GetBannerTitleOk returns a tuple with the BannerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerTitle

`func (o *IamBannerMessageAllOf) SetBannerTitle(v string)`

SetBannerTitle sets BannerTitle field to given value.

### HasBannerTitle

`func (o *IamBannerMessageAllOf) HasBannerTitle() bool`

HasBannerTitle returns a boolean if a field has been set.

### GetAccount

`func (o *IamBannerMessageAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamBannerMessageAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamBannerMessageAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamBannerMessageAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


