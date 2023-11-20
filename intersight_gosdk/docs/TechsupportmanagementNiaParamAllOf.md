# TechsupportmanagementNiaParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.NiaParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.NiaParam"]
**CollectionClass** | Pointer to **int32** | CollectionClass specifies which app to collect ND techsupport from - infra/NDI/NDO/NDFC. * &#x60;1&#x60; - Collect infra logs for Nexus Dashboard TACASSIST. * &#x60;2&#x60; - Collect logs for Nexus Dashboard Insights app through NDTACASSIST. * &#x60;3&#x60; - Collect logs for Nexus Dashboard Orchestrator app through NDTACASSIST. * &#x60;4&#x60; - Collect logs for Nexus Dashboard Fabric Controller app through NDTACASSIST. * &#x60;5&#x60; - Collect logs for Nexus Data Broker app through NDTACASSIST. * &#x60;6&#x60; - Collect logs for Nexus Data Broker Orchestrator app through NDTACASSIST. | [optional] [default to 1]
**CollectionLevel** | Pointer to **int32** | CollectionLevel controls the size / depth of the tech support files collected. * &#x60;1&#x60; - Use cached data in the returned collection. * &#x60;2&#x60; - Use current data in the returned collection. | [optional] [default to 1]
**Filename** | Pointer to **string** | Filename specifies an individual filename to collect from the endpoint. | [optional] 
**ForceFresh** | Pointer to **bool** | ForceFresh controls whether to return pre-collected files or force the collection of new files. | [optional] 
**Period** | Pointer to **int64** | Number of days for which to collect techsupport. | [optional] [default to 2]
**Pids** | Pointer to **[]string** |  | [optional] 
**SerialNumbers** | Pointer to **[]string** |  | [optional] 
**UpgradeLogs** | Pointer to **bool** | UpgradeLogs controls the inclusion of upgrade logs in tech support bundles. | [optional] 

## Methods

### NewTechsupportmanagementNiaParamAllOf

`func NewTechsupportmanagementNiaParamAllOf(classId string, objectType string, ) *TechsupportmanagementNiaParamAllOf`

NewTechsupportmanagementNiaParamAllOf instantiates a new TechsupportmanagementNiaParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementNiaParamAllOfWithDefaults

`func NewTechsupportmanagementNiaParamAllOfWithDefaults() *TechsupportmanagementNiaParamAllOf`

NewTechsupportmanagementNiaParamAllOfWithDefaults instantiates a new TechsupportmanagementNiaParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementNiaParamAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementNiaParamAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementNiaParamAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementNiaParamAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementNiaParamAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementNiaParamAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectionClass

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionClass() int32`

GetCollectionClass returns the CollectionClass field if non-nil, zero value otherwise.

### GetCollectionClassOk

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionClassOk() (*int32, bool)`

GetCollectionClassOk returns a tuple with the CollectionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionClass

`func (o *TechsupportmanagementNiaParamAllOf) SetCollectionClass(v int32)`

SetCollectionClass sets CollectionClass field to given value.

### HasCollectionClass

`func (o *TechsupportmanagementNiaParamAllOf) HasCollectionClass() bool`

HasCollectionClass returns a boolean if a field has been set.

### GetCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevel() int32`

GetCollectionLevel returns the CollectionLevel field if non-nil, zero value otherwise.

### GetCollectionLevelOk

`func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevelOk() (*int32, bool)`

GetCollectionLevelOk returns a tuple with the CollectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) SetCollectionLevel(v int32)`

SetCollectionLevel sets CollectionLevel field to given value.

### HasCollectionLevel

`func (o *TechsupportmanagementNiaParamAllOf) HasCollectionLevel() bool`

HasCollectionLevel returns a boolean if a field has been set.

### GetFilename

`func (o *TechsupportmanagementNiaParamAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *TechsupportmanagementNiaParamAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *TechsupportmanagementNiaParamAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *TechsupportmanagementNiaParamAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) GetForceFresh() bool`

GetForceFresh returns the ForceFresh field if non-nil, zero value otherwise.

### GetForceFreshOk

`func (o *TechsupportmanagementNiaParamAllOf) GetForceFreshOk() (*bool, bool)`

GetForceFreshOk returns a tuple with the ForceFresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) SetForceFresh(v bool)`

SetForceFresh sets ForceFresh field to given value.

### HasForceFresh

`func (o *TechsupportmanagementNiaParamAllOf) HasForceFresh() bool`

HasForceFresh returns a boolean if a field has been set.

### GetPeriod

`func (o *TechsupportmanagementNiaParamAllOf) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TechsupportmanagementNiaParamAllOf) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TechsupportmanagementNiaParamAllOf) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TechsupportmanagementNiaParamAllOf) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPids

`func (o *TechsupportmanagementNiaParamAllOf) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *TechsupportmanagementNiaParamAllOf) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *TechsupportmanagementNiaParamAllOf) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *TechsupportmanagementNiaParamAllOf) HasPids() bool`

HasPids returns a boolean if a field has been set.

### SetPidsNil

`func (o *TechsupportmanagementNiaParamAllOf) SetPidsNil(b bool)`

 SetPidsNil sets the value for Pids to be an explicit nil

### UnsetPids
`func (o *TechsupportmanagementNiaParamAllOf) UnsetPids()`

UnsetPids ensures that no value is present for Pids, not even an explicit nil
### GetSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *TechsupportmanagementNiaParamAllOf) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### SetSerialNumbersNil

`func (o *TechsupportmanagementNiaParamAllOf) SetSerialNumbersNil(b bool)`

 SetSerialNumbersNil sets the value for SerialNumbers to be an explicit nil

### UnsetSerialNumbers
`func (o *TechsupportmanagementNiaParamAllOf) UnsetSerialNumbers()`

UnsetSerialNumbers ensures that no value is present for SerialNumbers, not even an explicit nil
### GetUpgradeLogs

`func (o *TechsupportmanagementNiaParamAllOf) GetUpgradeLogs() bool`

GetUpgradeLogs returns the UpgradeLogs field if non-nil, zero value otherwise.

### GetUpgradeLogsOk

`func (o *TechsupportmanagementNiaParamAllOf) GetUpgradeLogsOk() (*bool, bool)`

GetUpgradeLogsOk returns a tuple with the UpgradeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLogs

`func (o *TechsupportmanagementNiaParamAllOf) SetUpgradeLogs(v bool)`

SetUpgradeLogs sets UpgradeLogs field to given value.

### HasUpgradeLogs

`func (o *TechsupportmanagementNiaParamAllOf) HasUpgradeLogs() bool`

HasUpgradeLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


