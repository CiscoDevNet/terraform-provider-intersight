# HclOperatingSystemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.OperatingSystem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.OperatingSystem"]
**Version** | Pointer to **string** | Version of the Operating System. | [optional] 
**Vendor** | Pointer to [**HclOperatingSystemVendorRelationship**](HclOperatingSystemVendorRelationship.md) |  | [optional] 

## Methods

### NewHclOperatingSystemAllOf

`func NewHclOperatingSystemAllOf(classId string, objectType string, ) *HclOperatingSystemAllOf`

NewHclOperatingSystemAllOf instantiates a new HclOperatingSystemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclOperatingSystemAllOfWithDefaults

`func NewHclOperatingSystemAllOfWithDefaults() *HclOperatingSystemAllOf`

NewHclOperatingSystemAllOfWithDefaults instantiates a new HclOperatingSystemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclOperatingSystemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclOperatingSystemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclOperatingSystemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclOperatingSystemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclOperatingSystemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclOperatingSystemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HclOperatingSystemAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HclOperatingSystemAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HclOperatingSystemAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HclOperatingSystemAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVendor

`func (o *HclOperatingSystemAllOf) GetVendor() HclOperatingSystemVendorRelationship`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HclOperatingSystemAllOf) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HclOperatingSystemAllOf) SetVendor(v HclOperatingSystemVendorRelationship)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HclOperatingSystemAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


