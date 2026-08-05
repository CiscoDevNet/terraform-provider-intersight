# BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken"]
**Value** | Pointer to **string** | This item enables a full memory test. Please note that this is a memory content test and is separate and distinct from the MBIST test of Interface and Data Eye.  PMU Mem BIST: this uses PMU firmware to test memory on all channels simultaneously. Failing memory will be repaired using soft or hard PPR depending on the PPR configuration.  Self-Healing Mem BIST: this runs the JEDEC DRAM self healing, if the device and DIMM support the feature. The DRAM will do a hard repair for failing memory.  PMU and Self-Healing Mem BIST: this option runs the PMU Mem BIST then the Self-Healing Mem BIST tests sequentially. * &#x60;Disabled&#x60; - Value -- Disabled for configuring CbsCmnMemHealingBistEnableBitMaskDdr token. * &#x60;PMU Mem BIST&#x60; - Value -- PMU Mem BIST for configuring CbsCmnMemHealingBistEnableBitMaskDdr token. * &#x60;Self-Healing Mem BIST&#x60; - Value -- Self-Healing Mem BIST for configuring CbsCmnMemHealingBistEnableBitMaskDdr token. * &#x60;PMU and Self-Healing Mem BIST&#x60; - Value -- PMU and Self-Healing Mem BIST for configuring CbsCmnMemHealingBistEnableBitMaskDdr token. | [optional] [default to "Disabled"]

## Methods

### NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken

`func NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken(classId string, objectType string, ) *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken`

NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken instantiates a new BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosTokenWithDefaults

`func NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosTokenWithDefaults() *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken`

NewBiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosTokenWithDefaults instantiates a new BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnMemHealingBistEnableBitMaskDdrC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


