# NiatelemetryDomThresInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DomThresInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DomThresInfo"]
**Dn** | Pointer to **string** | Returns distinguished name of the transceiver. | [optional] 
**HighAlarm** | Pointer to **string** | Returns highalarm value of the transceiver sensor. | [optional] 
**HighWarning** | Pointer to **string** | Returns highwarning value of the transceiver sensor. | [optional] 
**LowAlarm** | Pointer to **string** | Returns lowalarm value of the transceiver sensor. | [optional] 
**LowWarning** | Pointer to **string** | Returns lowwarning value of the transceiver sensor. | [optional] 
**NumLanes** | Pointer to **string** | Returns numlanes of the transceiver sensors. | [optional] 
**PartNumber** | Pointer to **string** | Returns part number of the transceiver. | [optional] 
**Type** | Pointer to **string** | Returns type of the transceiver sfp or qsfp. | [optional] 
**TypeName** | Pointer to **string** | Returns type name of the transceiver sfp or qsfp. | [optional] 
**Unit** | Pointer to **string** | Returns calibration value (unit) of transceiver sensor. | [optional] 
**VendorName** | Pointer to **string** | Returns vendor name string of the transceiver. | [optional] 
**VendorPn** | Pointer to **string** | Returns vendor part number of the transceiver. | [optional] 
**VendorRev** | Pointer to **string** | Returns vendor revision number of the transceiver. | [optional] 
**VendorSn** | Pointer to **string** | Returns vendor serial number of the transceiver. | [optional] 
**VersionId** | Pointer to **string** | Returns version identifier of the transceiver. | [optional] 

## Methods

### NewNiatelemetryDomThresInfo

`func NewNiatelemetryDomThresInfo(classId string, objectType string, ) *NiatelemetryDomThresInfo`

NewNiatelemetryDomThresInfo instantiates a new NiatelemetryDomThresInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDomThresInfoWithDefaults

`func NewNiatelemetryDomThresInfoWithDefaults() *NiatelemetryDomThresInfo`

NewNiatelemetryDomThresInfoWithDefaults instantiates a new NiatelemetryDomThresInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDomThresInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDomThresInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDomThresInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDomThresInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDomThresInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDomThresInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryDomThresInfo) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryDomThresInfo) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryDomThresInfo) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryDomThresInfo) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHighAlarm

`func (o *NiatelemetryDomThresInfo) GetHighAlarm() string`

GetHighAlarm returns the HighAlarm field if non-nil, zero value otherwise.

### GetHighAlarmOk

`func (o *NiatelemetryDomThresInfo) GetHighAlarmOk() (*string, bool)`

GetHighAlarmOk returns a tuple with the HighAlarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAlarm

`func (o *NiatelemetryDomThresInfo) SetHighAlarm(v string)`

SetHighAlarm sets HighAlarm field to given value.

### HasHighAlarm

`func (o *NiatelemetryDomThresInfo) HasHighAlarm() bool`

HasHighAlarm returns a boolean if a field has been set.

### GetHighWarning

`func (o *NiatelemetryDomThresInfo) GetHighWarning() string`

GetHighWarning returns the HighWarning field if non-nil, zero value otherwise.

### GetHighWarningOk

`func (o *NiatelemetryDomThresInfo) GetHighWarningOk() (*string, bool)`

GetHighWarningOk returns a tuple with the HighWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWarning

`func (o *NiatelemetryDomThresInfo) SetHighWarning(v string)`

SetHighWarning sets HighWarning field to given value.

### HasHighWarning

`func (o *NiatelemetryDomThresInfo) HasHighWarning() bool`

HasHighWarning returns a boolean if a field has been set.

### GetLowAlarm

`func (o *NiatelemetryDomThresInfo) GetLowAlarm() string`

GetLowAlarm returns the LowAlarm field if non-nil, zero value otherwise.

### GetLowAlarmOk

`func (o *NiatelemetryDomThresInfo) GetLowAlarmOk() (*string, bool)`

GetLowAlarmOk returns a tuple with the LowAlarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowAlarm

`func (o *NiatelemetryDomThresInfo) SetLowAlarm(v string)`

SetLowAlarm sets LowAlarm field to given value.

### HasLowAlarm

`func (o *NiatelemetryDomThresInfo) HasLowAlarm() bool`

HasLowAlarm returns a boolean if a field has been set.

### GetLowWarning

`func (o *NiatelemetryDomThresInfo) GetLowWarning() string`

GetLowWarning returns the LowWarning field if non-nil, zero value otherwise.

### GetLowWarningOk

`func (o *NiatelemetryDomThresInfo) GetLowWarningOk() (*string, bool)`

GetLowWarningOk returns a tuple with the LowWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWarning

`func (o *NiatelemetryDomThresInfo) SetLowWarning(v string)`

SetLowWarning sets LowWarning field to given value.

### HasLowWarning

`func (o *NiatelemetryDomThresInfo) HasLowWarning() bool`

HasLowWarning returns a boolean if a field has been set.

### GetNumLanes

`func (o *NiatelemetryDomThresInfo) GetNumLanes() string`

GetNumLanes returns the NumLanes field if non-nil, zero value otherwise.

### GetNumLanesOk

`func (o *NiatelemetryDomThresInfo) GetNumLanesOk() (*string, bool)`

GetNumLanesOk returns a tuple with the NumLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLanes

`func (o *NiatelemetryDomThresInfo) SetNumLanes(v string)`

SetNumLanes sets NumLanes field to given value.

### HasNumLanes

`func (o *NiatelemetryDomThresInfo) HasNumLanes() bool`

HasNumLanes returns a boolean if a field has been set.

### GetPartNumber

`func (o *NiatelemetryDomThresInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *NiatelemetryDomThresInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *NiatelemetryDomThresInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *NiatelemetryDomThresInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryDomThresInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryDomThresInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryDomThresInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryDomThresInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *NiatelemetryDomThresInfo) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *NiatelemetryDomThresInfo) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *NiatelemetryDomThresInfo) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *NiatelemetryDomThresInfo) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetUnit

`func (o *NiatelemetryDomThresInfo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *NiatelemetryDomThresInfo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *NiatelemetryDomThresInfo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *NiatelemetryDomThresInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVendorName

`func (o *NiatelemetryDomThresInfo) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *NiatelemetryDomThresInfo) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *NiatelemetryDomThresInfo) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *NiatelemetryDomThresInfo) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetVendorPn

`func (o *NiatelemetryDomThresInfo) GetVendorPn() string`

GetVendorPn returns the VendorPn field if non-nil, zero value otherwise.

### GetVendorPnOk

`func (o *NiatelemetryDomThresInfo) GetVendorPnOk() (*string, bool)`

GetVendorPnOk returns a tuple with the VendorPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPn

`func (o *NiatelemetryDomThresInfo) SetVendorPn(v string)`

SetVendorPn sets VendorPn field to given value.

### HasVendorPn

`func (o *NiatelemetryDomThresInfo) HasVendorPn() bool`

HasVendorPn returns a boolean if a field has been set.

### GetVendorRev

`func (o *NiatelemetryDomThresInfo) GetVendorRev() string`

GetVendorRev returns the VendorRev field if non-nil, zero value otherwise.

### GetVendorRevOk

`func (o *NiatelemetryDomThresInfo) GetVendorRevOk() (*string, bool)`

GetVendorRevOk returns a tuple with the VendorRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorRev

`func (o *NiatelemetryDomThresInfo) SetVendorRev(v string)`

SetVendorRev sets VendorRev field to given value.

### HasVendorRev

`func (o *NiatelemetryDomThresInfo) HasVendorRev() bool`

HasVendorRev returns a boolean if a field has been set.

### GetVendorSn

`func (o *NiatelemetryDomThresInfo) GetVendorSn() string`

GetVendorSn returns the VendorSn field if non-nil, zero value otherwise.

### GetVendorSnOk

`func (o *NiatelemetryDomThresInfo) GetVendorSnOk() (*string, bool)`

GetVendorSnOk returns a tuple with the VendorSn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSn

`func (o *NiatelemetryDomThresInfo) SetVendorSn(v string)`

SetVendorSn sets VendorSn field to given value.

### HasVendorSn

`func (o *NiatelemetryDomThresInfo) HasVendorSn() bool`

HasVendorSn returns a boolean if a field has been set.

### GetVersionId

`func (o *NiatelemetryDomThresInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *NiatelemetryDomThresInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *NiatelemetryDomThresInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *NiatelemetryDomThresInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


