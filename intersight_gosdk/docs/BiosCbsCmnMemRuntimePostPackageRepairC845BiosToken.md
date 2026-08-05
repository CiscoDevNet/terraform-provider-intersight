# BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnMemRuntimePostPackageRepairC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnMemRuntimePostPackageRepairC845BiosToken"]
**Value** | Pointer to **string** | Enable or Disable DRAM Run Time Post Package Repair. * &#x60;Disable&#x60; - Value -- Disable for configuring CbsCmnMemRuntimePostPackageRepair token. * &#x60;Enable&#x60; - Value -- Enable for configuring CbsCmnMemRuntimePostPackageRepair token. | [optional] [default to "Disable"]

## Methods

### NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosToken

`func NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosToken(classId string, objectType string, ) *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken`

NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosToken instantiates a new BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosTokenWithDefaults

`func NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosTokenWithDefaults() *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken`

NewBiosCbsCmnMemRuntimePostPackageRepairC845BiosTokenWithDefaults instantiates a new BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnMemRuntimePostPackageRepairC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


