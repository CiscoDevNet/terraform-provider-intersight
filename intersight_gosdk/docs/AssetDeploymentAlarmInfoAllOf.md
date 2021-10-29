# AssetDeploymentAlarmInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeploymentAlarmInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeploymentAlarmInfo"]
**EnabledAlarms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAssetDeploymentAlarmInfoAllOf

`func NewAssetDeploymentAlarmInfoAllOf(classId string, objectType string, ) *AssetDeploymentAlarmInfoAllOf`

NewAssetDeploymentAlarmInfoAllOf instantiates a new AssetDeploymentAlarmInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentAlarmInfoAllOfWithDefaults

`func NewAssetDeploymentAlarmInfoAllOfWithDefaults() *AssetDeploymentAlarmInfoAllOf`

NewAssetDeploymentAlarmInfoAllOfWithDefaults instantiates a new AssetDeploymentAlarmInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeploymentAlarmInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeploymentAlarmInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeploymentAlarmInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeploymentAlarmInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeploymentAlarmInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeploymentAlarmInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabledAlarms

`func (o *AssetDeploymentAlarmInfoAllOf) GetEnabledAlarms() []string`

GetEnabledAlarms returns the EnabledAlarms field if non-nil, zero value otherwise.

### GetEnabledAlarmsOk

`func (o *AssetDeploymentAlarmInfoAllOf) GetEnabledAlarmsOk() (*[]string, bool)`

GetEnabledAlarmsOk returns a tuple with the EnabledAlarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlarms

`func (o *AssetDeploymentAlarmInfoAllOf) SetEnabledAlarms(v []string)`

SetEnabledAlarms sets EnabledAlarms field to given value.

### HasEnabledAlarms

`func (o *AssetDeploymentAlarmInfoAllOf) HasEnabledAlarms() bool`

HasEnabledAlarms returns a boolean if a field has been set.

### SetEnabledAlarmsNil

`func (o *AssetDeploymentAlarmInfoAllOf) SetEnabledAlarmsNil(b bool)`

 SetEnabledAlarmsNil sets the value for EnabledAlarms to be an explicit nil

### UnsetEnabledAlarms
`func (o *AssetDeploymentAlarmInfoAllOf) UnsetEnabledAlarms()`

UnsetEnabledAlarms ensures that no value is present for EnabledAlarms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


