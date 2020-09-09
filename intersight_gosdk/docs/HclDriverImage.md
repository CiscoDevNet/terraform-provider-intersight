# HclDriverImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverIsoUrl** | Pointer to **string** | URL of the driver ISO images. | [optional] 
**ManagementType** | Pointer to **string** | Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * &#x60;UCSM&#x60; - The server is managed by UCS Manager. * &#x60;IMC&#x60; - The server is standalone managed by CIMC. | [optional] [default to "UCSM"]
**ServerPid** | Pointer to **string** | Three part ID representing the server model as returned by UCSM/CIMC XML APIs. | [optional] 

## Methods

### NewHclDriverImage

`func NewHclDriverImage() *HclDriverImage`

NewHclDriverImage instantiates a new HclDriverImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclDriverImageWithDefaults

`func NewHclDriverImageWithDefaults() *HclDriverImage`

NewHclDriverImageWithDefaults instantiates a new HclDriverImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


