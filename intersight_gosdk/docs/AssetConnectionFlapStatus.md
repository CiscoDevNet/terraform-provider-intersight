# AssetConnectionFlapStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.ConnectionFlapStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.ConnectionFlapStatus"]
**FlapAlertTime** | Pointer to **time.Time** | Time when flapping was reported. | [optional] 
**FlapCount** | Pointer to **int64** | The number of times a device disconnected within a specified time window. | [optional] 
**FlapDetected** | Pointer to **bool** | Indicates if the device is flapping. | [optional] 
**WindowSize** | Pointer to **string** | The time window during which device disconnects are counted. E.g. values are PT30M or PT1H. | [optional] 

## Methods

### NewAssetConnectionFlapStatus

`func NewAssetConnectionFlapStatus(classId string, objectType string, ) *AssetConnectionFlapStatus`

NewAssetConnectionFlapStatus instantiates a new AssetConnectionFlapStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetConnectionFlapStatusWithDefaults

`func NewAssetConnectionFlapStatusWithDefaults() *AssetConnectionFlapStatus`

NewAssetConnectionFlapStatusWithDefaults instantiates a new AssetConnectionFlapStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetConnectionFlapStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetConnectionFlapStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetConnectionFlapStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetConnectionFlapStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetConnectionFlapStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetConnectionFlapStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFlapAlertTime

`func (o *AssetConnectionFlapStatus) GetFlapAlertTime() time.Time`

GetFlapAlertTime returns the FlapAlertTime field if non-nil, zero value otherwise.

### GetFlapAlertTimeOk

`func (o *AssetConnectionFlapStatus) GetFlapAlertTimeOk() (*time.Time, bool)`

GetFlapAlertTimeOk returns a tuple with the FlapAlertTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapAlertTime

`func (o *AssetConnectionFlapStatus) SetFlapAlertTime(v time.Time)`

SetFlapAlertTime sets FlapAlertTime field to given value.

### HasFlapAlertTime

`func (o *AssetConnectionFlapStatus) HasFlapAlertTime() bool`

HasFlapAlertTime returns a boolean if a field has been set.

### GetFlapCount

`func (o *AssetConnectionFlapStatus) GetFlapCount() int64`

GetFlapCount returns the FlapCount field if non-nil, zero value otherwise.

### GetFlapCountOk

`func (o *AssetConnectionFlapStatus) GetFlapCountOk() (*int64, bool)`

GetFlapCountOk returns a tuple with the FlapCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapCount

`func (o *AssetConnectionFlapStatus) SetFlapCount(v int64)`

SetFlapCount sets FlapCount field to given value.

### HasFlapCount

`func (o *AssetConnectionFlapStatus) HasFlapCount() bool`

HasFlapCount returns a boolean if a field has been set.

### GetFlapDetected

`func (o *AssetConnectionFlapStatus) GetFlapDetected() bool`

GetFlapDetected returns the FlapDetected field if non-nil, zero value otherwise.

### GetFlapDetectedOk

`func (o *AssetConnectionFlapStatus) GetFlapDetectedOk() (*bool, bool)`

GetFlapDetectedOk returns a tuple with the FlapDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapDetected

`func (o *AssetConnectionFlapStatus) SetFlapDetected(v bool)`

SetFlapDetected sets FlapDetected field to given value.

### HasFlapDetected

`func (o *AssetConnectionFlapStatus) HasFlapDetected() bool`

HasFlapDetected returns a boolean if a field has been set.

### GetWindowSize

`func (o *AssetConnectionFlapStatus) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *AssetConnectionFlapStatus) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *AssetConnectionFlapStatus) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *AssetConnectionFlapStatus) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


