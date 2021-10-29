# AssetDeploymentDeviceAlarmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeploymentDeviceAlarmInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeploymentDeviceAlarmInfo"]
**EnabledAlarms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAssetDeploymentDeviceAlarmInfo

`func NewAssetDeploymentDeviceAlarmInfo(classId string, objectType string, ) *AssetDeploymentDeviceAlarmInfo`

NewAssetDeploymentDeviceAlarmInfo instantiates a new AssetDeploymentDeviceAlarmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentDeviceAlarmInfoWithDefaults

`func NewAssetDeploymentDeviceAlarmInfoWithDefaults() *AssetDeploymentDeviceAlarmInfo`

NewAssetDeploymentDeviceAlarmInfoWithDefaults instantiates a new AssetDeploymentDeviceAlarmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeploymentDeviceAlarmInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeploymentDeviceAlarmInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeploymentDeviceAlarmInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeploymentDeviceAlarmInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeploymentDeviceAlarmInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeploymentDeviceAlarmInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabledAlarms

`func (o *AssetDeploymentDeviceAlarmInfo) GetEnabledAlarms() []string`

GetEnabledAlarms returns the EnabledAlarms field if non-nil, zero value otherwise.

### GetEnabledAlarmsOk

`func (o *AssetDeploymentDeviceAlarmInfo) GetEnabledAlarmsOk() (*[]string, bool)`

GetEnabledAlarmsOk returns a tuple with the EnabledAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlarms

`func (o *AssetDeploymentDeviceAlarmInfo) SetEnabledAlarms(v []string)`

SetEnabledAlarms sets EnabledAlarms field to given value.

### HasEnabledAlarms

`func (o *AssetDeploymentDeviceAlarmInfo) HasEnabledAlarms() bool`

HasEnabledAlarms returns a boolean if a field has been set.

### SetEnabledAlarmsNil

`func (o *AssetDeploymentDeviceAlarmInfo) SetEnabledAlarmsNil(b bool)`

 SetEnabledAlarmsNil sets the value for EnabledAlarms to be an explicit nil

### UnsetEnabledAlarms
`func (o *AssetDeploymentDeviceAlarmInfo) UnsetEnabledAlarms()`

UnsetEnabledAlarms ensures that no value is present for EnabledAlarms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


