# MoVersionContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterestedMos** | Pointer to [**[]MoMoRef**](mo.MoRef.md) |  | [optional] 
**RefMo** | Pointer to [**MoMoRef**](mo.MoRef.md) |  | [optional] 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | The time this versioned Managed Object was created. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the Managed Object, e.g. an incrementing number or a hash id. | [optional] [readonly] 
**VersionType** | Pointer to **string** | Specifies type of version. Currently the only supported value is \&quot;Configured\&quot; that is used to keep track of snapshots of policies and profiles that are intended to be configured to target endpoints. * &#x60;Modified&#x60; - Version created every time an object is modified. * &#x60;Configured&#x60; - Version created every time an object is configured to the service profile. * &#x60;Deployed&#x60; - Version created for objects related to a service profile when it is deployed. | [optional] [readonly] [default to "Modified"]

## Methods

### NewMoVersionContextAllOf

`func NewMoVersionContextAllOf() *MoVersionContextAllOf`

NewMoVersionContextAllOf instantiates a new MoVersionContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoVersionContextAllOfWithDefaults

`func NewMoVersionContextAllOfWithDefaults() *MoVersionContextAllOf`

NewMoVersionContextAllOfWithDefaults instantiates a new MoVersionContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterestedMos

`func (o *MoVersionContextAllOf) GetInterestedMos() []MoMoRef`

GetInterestedMos returns the InterestedMos field if non-nil, zero value otherwise.

### GetInterestedMosOk

`func (o *MoVersionContextAllOf) GetInterestedMosOk() (*[]MoMoRef, bool)`

GetInterestedMosOk returns a tuple with the InterestedMos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestedMos

`func (o *MoVersionContextAllOf) SetInterestedMos(v []MoMoRef)`

SetInterestedMos sets InterestedMos field to given value.

### HasInterestedMos

`func (o *MoVersionContextAllOf) HasInterestedMos() bool`

HasInterestedMos returns a boolean if a field has been set.

### GetRefMo

`func (o *MoVersionContextAllOf) GetRefMo() MoMoRef`

GetRefMo returns the RefMo field if non-nil, zero value otherwise.

### GetRefMoOk

`func (o *MoVersionContextAllOf) GetRefMoOk() (*MoMoRef, bool)`

GetRefMoOk returns a tuple with the RefMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefMo

`func (o *MoVersionContextAllOf) SetRefMo(v MoMoRef)`

SetRefMo sets RefMo field to given value.

### HasRefMo

`func (o *MoVersionContextAllOf) HasRefMo() bool`

HasRefMo returns a boolean if a field has been set.

### GetTimestamp

`func (o *MoVersionContextAllOf) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MoVersionContextAllOf) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MoVersionContextAllOf) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MoVersionContextAllOf) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *MoVersionContextAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MoVersionContextAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MoVersionContextAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MoVersionContextAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionType

`func (o *MoVersionContextAllOf) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *MoVersionContextAllOf) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *MoVersionContextAllOf) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *MoVersionContextAllOf) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


