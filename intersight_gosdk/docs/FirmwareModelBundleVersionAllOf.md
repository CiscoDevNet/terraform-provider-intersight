# FirmwareModelBundleVersionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ModelBundleVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ModelBundleVersion"]
**BundleVersion** | Pointer to **string** | The bundle version to which the server will be upgraded. | [optional] 
**ModelFamily** | Pointer to **string** | The server family that will be impacted by this upgrade. * &#x60;UCSC-C220-M5&#x60; - The upgrade on all C220-M5 servers claimed in setup. * &#x60;UCSC-C220-M4&#x60; - The upgrade on all C220-M4 servers claimed in setup. * &#x60;UCSC-C240-M4&#x60; - The upgrade on all C240-M4 servers claimed in setup. * &#x60;UCSC-C460-M4&#x60; - The upgrade on all C460-M4 servers claimed in setup. * &#x60;UCSC-C240-M5&#x60; - The upgrade on all C240-M5 servers claimed in setup. * &#x60;UCSC-C480-M5&#x60; - The upgrade on all C480-M5 servers claimed in setup. * &#x60;UCSB-B200-M5&#x60; - The upgrade on all B200-M5 servers claimed in setup. * &#x60;UCSB-B480-M5&#x60; - The upgrade on all B480-M5 servers claimed in setup. * &#x60;UCSC-C220-M6&#x60; - The upgrade on all C220-M6 servers claimed in setup. * &#x60;UCSC-C240-M6&#x60; - The upgrade on all C240-M6 servers claimed in setup. * &#x60;UCSC-C225-M6&#x60; - The upgrade on all C225-M6 servers claimed in setup. * &#x60;UCSC-C245-M6&#x60; - The upgrade on all C245-M6 servers claimed in setup. * &#x60;UCSB-B200-M6&#x60; - The upgrade on all B200-M6 servers claimed in setup. * &#x60;UCSX-210C-M6&#x60; - The upgrade on all 210C-M6 servers claimed in setup. * &#x60;UCSX-210C-M7&#x60; - The upgrade on all 210C-M7 servers claimed in setup. * &#x60;UCSX-220-M7&#x60; - The upgrade on all C220-M7 servers claimed in setup. * &#x60;UCSX-240-M7&#x60; - The upgrade on all C240-M7 servers claimed in setup. * &#x60;UCSC-C125&#x60; - The upgrade on all C125 servers claimed in setup. | [optional] [default to "UCSC-C220-M5"]

## Methods

### NewFirmwareModelBundleVersionAllOf

`func NewFirmwareModelBundleVersionAllOf(classId string, objectType string, ) *FirmwareModelBundleVersionAllOf`

NewFirmwareModelBundleVersionAllOf instantiates a new FirmwareModelBundleVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareModelBundleVersionAllOfWithDefaults

`func NewFirmwareModelBundleVersionAllOfWithDefaults() *FirmwareModelBundleVersionAllOf`

NewFirmwareModelBundleVersionAllOfWithDefaults instantiates a new FirmwareModelBundleVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareModelBundleVersionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareModelBundleVersionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareModelBundleVersionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareModelBundleVersionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareModelBundleVersionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareModelBundleVersionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBundleVersion

`func (o *FirmwareModelBundleVersionAllOf) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *FirmwareModelBundleVersionAllOf) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *FirmwareModelBundleVersionAllOf) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *FirmwareModelBundleVersionAllOf) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetModelFamily

`func (o *FirmwareModelBundleVersionAllOf) GetModelFamily() string`

GetModelFamily returns the ModelFamily field if non-nil, zero value otherwise.

### GetModelFamilyOk

`func (o *FirmwareModelBundleVersionAllOf) GetModelFamilyOk() (*string, bool)`

GetModelFamilyOk returns a tuple with the ModelFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFamily

`func (o *FirmwareModelBundleVersionAllOf) SetModelFamily(v string)`

SetModelFamily sets ModelFamily field to given value.

### HasModelFamily

`func (o *FirmwareModelBundleVersionAllOf) HasModelFamily() bool`

HasModelFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


