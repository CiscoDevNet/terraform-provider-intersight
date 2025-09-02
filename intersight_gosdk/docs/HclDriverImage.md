# HclDriverImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.DriverImage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.DriverImage"]
**DriverIsoUrl** | Pointer to **string** | URL of the driver ISO images. | [optional] 
**ManagementType** | Pointer to **string** | Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * &#x60;UCSM&#x60; - The server is managed by UCS Manager. * &#x60;IMC&#x60; - The server is standalone managed by CIMC. * &#x60;IMM&#x60; - The server is standalone managed by Intersight. | [optional] [default to "UCSM"]
**ServerPid** | Pointer to **string** | Three part ID representing the server model as returned by UCSM/CIMC XML APIs. | [optional] 

## Methods

### NewHclDriverImage

`func NewHclDriverImage(classId string, objectType string, ) *HclDriverImage`

NewHclDriverImage instantiates a new HclDriverImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclDriverImageWithDefaults

`func NewHclDriverImageWithDefaults() *HclDriverImage`

NewHclDriverImageWithDefaults instantiates a new HclDriverImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclDriverImage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclDriverImage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclDriverImage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclDriverImage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclDriverImage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclDriverImage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriverIsoUrl

`func (o *HclDriverImage) GetDriverIsoUrl() string`

GetDriverIsoUrl returns the DriverIsoUrl field if non-nil, zero value otherwise.

### GetDriverIsoUrlOk

`func (o *HclDriverImage) GetDriverIsoUrlOk() (*string, bool)`

GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverIsoUrl

`func (o *HclDriverImage) SetDriverIsoUrl(v string)`

SetDriverIsoUrl sets DriverIsoUrl field to given value.

### HasDriverIsoUrl

`func (o *HclDriverImage) HasDriverIsoUrl() bool`

HasDriverIsoUrl returns a boolean if a field has been set.

### GetManagementType

`func (o *HclDriverImage) GetManagementType() string`

GetManagementType returns the ManagementType field if non-nil, zero value otherwise.

### GetManagementTypeOk

`func (o *HclDriverImage) GetManagementTypeOk() (*string, bool)`

GetManagementTypeOk returns a tuple with the ManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementType

`func (o *HclDriverImage) SetManagementType(v string)`

SetManagementType sets ManagementType field to given value.

### HasManagementType

`func (o *HclDriverImage) HasManagementType() bool`

HasManagementType returns a boolean if a field has been set.

### GetServerPid

`func (o *HclDriverImage) GetServerPid() string`

GetServerPid returns the ServerPid field if non-nil, zero value otherwise.

### GetServerPidOk

`func (o *HclDriverImage) GetServerPidOk() (*string, bool)`

GetServerPidOk returns a tuple with the ServerPid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPid

`func (o *HclDriverImage) SetServerPid(v string)`

SetServerPid sets ServerPid field to given value.

### HasServerPid

`func (o *HclDriverImage) HasServerPid() bool`

HasServerPid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


