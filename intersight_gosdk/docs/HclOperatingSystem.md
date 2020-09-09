# HclOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the Operating System. | [optional] 
**Vendor** | Pointer to [**HclOperatingSystemVendorRelationship**](hcl.OperatingSystemVendor.Relationship.md) |  | [optional] 

## Methods

### NewHclOperatingSystem

`func NewHclOperatingSystem() *HclOperatingSystem`

NewHclOperatingSystem instantiates a new HclOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclOperatingSystemWithDefaults

`func NewHclOperatingSystemWithDefaults() *HclOperatingSystem`

NewHclOperatingSystemWithDefaults instantiates a new HclOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HclOperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HclOperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HclOperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HclOperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVendor

`func (o *HclOperatingSystem) GetVendor() HclOperatingSystemVendorRelationship`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HclOperatingSystem) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HclOperatingSystem) SetVendor(v HclOperatingSystemVendorRelationship)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HclOperatingSystem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


