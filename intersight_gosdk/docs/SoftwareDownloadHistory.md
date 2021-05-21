# SoftwareDownloadHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.DownloadHistory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.DownloadHistory"]
**Name** | Pointer to **string** | The name of software which was downloaded. | [optional] [readonly] 
**Product** | Pointer to **string** | The product type of the downloaded software. | [optional] 
**Timestamp** | Pointer to **time.Time** | The download time of the software image. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of software which was downloaded. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Image** | Pointer to [**FirmwareBaseDistributableRelationship**](firmware.BaseDistributable.Relationship.md) |  | [optional] 

## Methods

### NewSoftwareDownloadHistory

`func NewSoftwareDownloadHistory(classId string, objectType string, ) *SoftwareDownloadHistory`

NewSoftwareDownloadHistory instantiates a new SoftwareDownloadHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareDownloadHistoryWithDefaults

`func NewSoftwareDownloadHistoryWithDefaults() *SoftwareDownloadHistory`

NewSoftwareDownloadHistoryWithDefaults instantiates a new SoftwareDownloadHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareDownloadHistory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareDownloadHistory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareDownloadHistory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareDownloadHistory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareDownloadHistory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareDownloadHistory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *SoftwareDownloadHistory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareDownloadHistory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareDownloadHistory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwareDownloadHistory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProduct

`func (o *SoftwareDownloadHistory) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SoftwareDownloadHistory) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SoftwareDownloadHistory) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SoftwareDownloadHistory) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTimestamp

`func (o *SoftwareDownloadHistory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SoftwareDownloadHistory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SoftwareDownloadHistory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SoftwareDownloadHistory) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwareDownloadHistory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareDownloadHistory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareDownloadHistory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwareDownloadHistory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *SoftwareDownloadHistory) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SoftwareDownloadHistory) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SoftwareDownloadHistory) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SoftwareDownloadHistory) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetImage

`func (o *SoftwareDownloadHistory) GetImage() FirmwareBaseDistributableRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SoftwareDownloadHistory) GetImageOk() (*FirmwareBaseDistributableRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SoftwareDownloadHistory) SetImage(v FirmwareBaseDistributableRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *SoftwareDownloadHistory) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


