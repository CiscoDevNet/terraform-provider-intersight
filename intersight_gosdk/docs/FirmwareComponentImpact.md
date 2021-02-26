# FirmwareComponentImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ComponentImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ComponentImpact"]
**Component** | Pointer to **string** | Impact on the component after the upgrade. * &#x60;ALL&#x60; - This represents all the components. * &#x60;ALL,HDD&#x60; - This represents all the components plus the HDDs. * &#x60;Drive-U.2&#x60; - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category). * &#x60;Storage&#x60; - This represents the storage controller components. * &#x60;None&#x60; - This represents none of the components. * &#x60;NXOS&#x60; - This represents NXOS components. * &#x60;IOM&#x60; - This represents IOM components. * &#x60;PSU&#x60; - This represents PSU components. * &#x60;CIMC&#x60; - This represents CIMC components. * &#x60;BIOS&#x60; - This represents BIOS components. * &#x60;PCIE&#x60; - This represents PCIE components. * &#x60;Drive&#x60; - This represents Drive components. * &#x60;DIMM&#x60; - This represents DIMM components. * &#x60;BoardController&#x60; - This represents Board Controller components. * &#x60;StorageController&#x60; - This represents Storage Controller components. * &#x60;HBA&#x60; - This represents HBA components. * &#x60;GPU&#x60; - This represents GPU components. * &#x60;SasExpander&#x60; - This represents SasExpander components. * &#x60;MSwitch&#x60; - This represents mSwitch components. * &#x60;CMC&#x60; - This represents CMC components. | [optional] [default to "ALL"]

## Methods

### NewFirmwareComponentImpact

`func NewFirmwareComponentImpact(classId string, objectType string, ) *FirmwareComponentImpact`

NewFirmwareComponentImpact instantiates a new FirmwareComponentImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareComponentImpactWithDefaults

`func NewFirmwareComponentImpactWithDefaults() *FirmwareComponentImpact`

NewFirmwareComponentImpactWithDefaults instantiates a new FirmwareComponentImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareComponentImpact) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareComponentImpact) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareComponentImpact) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareComponentImpact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareComponentImpact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareComponentImpact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComponent

`func (o *FirmwareComponentImpact) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FirmwareComponentImpact) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FirmwareComponentImpact) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FirmwareComponentImpact) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


