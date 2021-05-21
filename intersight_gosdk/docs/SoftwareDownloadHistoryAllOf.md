# SoftwareDownloadHistoryAllOf

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

### NewSoftwareDownloadHistoryAllOf

`func NewSoftwareDownloadHistoryAllOf(classId string, objectType string, ) *SoftwareDownloadHistoryAllOf`

NewSoftwareDownloadHistoryAllOf instantiates a new SoftwareDownloadHistoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareDownloadHistoryAllOfWithDefaults

`func NewSoftwareDownloadHistoryAllOfWithDefaults() *SoftwareDownloadHistoryAllOf`

NewSoftwareDownloadHistoryAllOfWithDefaults instantiates a new SoftwareDownloadHistoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareDownloadHistoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareDownloadHistoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareDownloadHistoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareDownloadHistoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareDownloadHistoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareDownloadHistoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *SoftwareDownloadHistoryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareDownloadHistoryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareDownloadHistoryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwareDownloadHistoryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProduct

`func (o *SoftwareDownloadHistoryAllOf) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SoftwareDownloadHistoryAllOf) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SoftwareDownloadHistoryAllOf) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SoftwareDownloadHistoryAllOf) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTimestamp

`func (o *SoftwareDownloadHistoryAllOf) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SoftwareDownloadHistoryAllOf) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SoftwareDownloadHistoryAllOf) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SoftwareDownloadHistoryAllOf) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *SoftwareDownloadHistoryAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareDownloadHistoryAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareDownloadHistoryAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwareDownloadHistoryAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *SoftwareDownloadHistoryAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SoftwareDownloadHistoryAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SoftwareDownloadHistoryAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SoftwareDownloadHistoryAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetImage

`func (o *SoftwareDownloadHistoryAllOf) GetImage() FirmwareBaseDistributableRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SoftwareDownloadHistoryAllOf) GetImageOk() (*FirmwareBaseDistributableRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SoftwareDownloadHistoryAllOf) SetImage(v FirmwareBaseDistributableRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *SoftwareDownloadHistoryAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


