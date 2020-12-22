# VnicEthInterruptSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthInterruptSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthInterruptSettings"]
**CoalescingTime** | Pointer to **int64** | The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field. | [optional] [default to 125]
**CoalescingType** | Pointer to **string** | Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field. * &#x60;MIN&#x60; - The system waits for the time specified in the Coalescing Time field before sending another interrupt event. * &#x60;IDLE&#x60; - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field. | [optional] [default to "MIN"]
**Count** | Pointer to **int64** | The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources. | [optional] [default to 8]
**Mode** | Pointer to **string** | Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option. * &#x60;MSIx&#x60; - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option. * &#x60;MSI&#x60; - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts. * &#x60;INTx&#x60; - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems. | [optional] [default to "MSIx"]

## Methods

### NewVnicEthInterruptSettingsAllOf

`func NewVnicEthInterruptSettingsAllOf(classId string, objectType string, ) *VnicEthInterruptSettingsAllOf`

NewVnicEthInterruptSettingsAllOf instantiates a new VnicEthInterruptSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthInterruptSettingsAllOfWithDefaults

`func NewVnicEthInterruptSettingsAllOfWithDefaults() *VnicEthInterruptSettingsAllOf`

NewVnicEthInterruptSettingsAllOfWithDefaults instantiates a new VnicEthInterruptSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthInterruptSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthInterruptSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthInterruptSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthInterruptSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthInterruptSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthInterruptSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCoalescingTime

`func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTime() int64`

GetCoalescingTime returns the CoalescingTime field if non-nil, zero value otherwise.

### GetCoalescingTimeOk

`func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTimeOk() (*int64, bool)`

GetCoalescingTimeOk returns a tuple with the CoalescingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingTime

`func (o *VnicEthInterruptSettingsAllOf) SetCoalescingTime(v int64)`

SetCoalescingTime sets CoalescingTime field to given value.

### HasCoalescingTime

`func (o *VnicEthInterruptSettingsAllOf) HasCoalescingTime() bool`

HasCoalescingTime returns a boolean if a field has been set.

### GetCoalescingType

`func (o *VnicEthInterruptSettingsAllOf) GetCoalescingType() string`

GetCoalescingType returns the CoalescingType field if non-nil, zero value otherwise.

### GetCoalescingTypeOk

`func (o *VnicEthInterruptSettingsAllOf) GetCoalescingTypeOk() (*string, bool)`

GetCoalescingTypeOk returns a tuple with the CoalescingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingType

`func (o *VnicEthInterruptSettingsAllOf) SetCoalescingType(v string)`

SetCoalescingType sets CoalescingType field to given value.

### HasCoalescingType

`func (o *VnicEthInterruptSettingsAllOf) HasCoalescingType() bool`

HasCoalescingType returns a boolean if a field has been set.

### GetCount

`func (o *VnicEthInterruptSettingsAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicEthInterruptSettingsAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicEthInterruptSettingsAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicEthInterruptSettingsAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMode

`func (o *VnicEthInterruptSettingsAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VnicEthInterruptSettingsAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VnicEthInterruptSettingsAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VnicEthInterruptSettingsAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


