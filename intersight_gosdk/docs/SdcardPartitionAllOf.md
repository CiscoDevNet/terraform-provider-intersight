# SdcardPartitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | This specifies the type of the partition. Allowed values are OS, Utility. * &#x60;OS&#x60; - This partition contains virtual drives where user can install operating system. * &#x60;Utility&#x60; - This partition contains virtual drives for utilities such as SCU, HUU, Drivers and Diagnostics. | [optional] [default to "OS"]
**VirtualDrives** | Pointer to [**[]SdcardVirtualDrive**](sdcard.VirtualDrive.md) |  | [optional] 

## Methods

### NewSdcardPartitionAllOf

`func NewSdcardPartitionAllOf() *SdcardPartitionAllOf`

NewSdcardPartitionAllOf instantiates a new SdcardPartitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdcardPartitionAllOfWithDefaults

`func NewSdcardPartitionAllOfWithDefaults() *SdcardPartitionAllOf`

NewSdcardPartitionAllOfWithDefaults instantiates a new SdcardPartitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SdcardPartitionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdcardPartitionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdcardPartitionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdcardPartitionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDrives

`func (o *SdcardPartitionAllOf) GetVirtualDrives() []SdcardVirtualDrive`

GetVirtualDrives returns the VirtualDrives field if non-nil, zero value otherwise.

### GetVirtualDrivesOk

`func (o *SdcardPartitionAllOf) GetVirtualDrivesOk() (*[]SdcardVirtualDrive, bool)`

GetVirtualDrivesOk returns a tuple with the VirtualDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrives

`func (o *SdcardPartitionAllOf) SetVirtualDrives(v []SdcardVirtualDrive)`

SetVirtualDrives sets VirtualDrives field to given value.

### HasVirtualDrives

`func (o *SdcardPartitionAllOf) HasVirtualDrives() bool`

HasVirtualDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


