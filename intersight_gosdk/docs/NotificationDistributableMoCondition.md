# NotificationDistributableMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.DistributableMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.DistributableMoCondition"]
**ImageType** | Pointer to **string** | Image type for which user wants to create notification. The list of valid values for Image type of notification.DistributableMoCondition is software.HyperflexBundleDistributable, software.ApplianceDistributableIntelligence, firmware.Distributable, software.HciBundleDistributable. * &#x60;None&#x60; - Default value for DistributableObjectType enum. * &#x60;software.ApplianceDistributableIntelligence&#x60; - Object type for Appliance Intelligence updates related distributable. * &#x60;firmware.Distributable&#x60; - Object type for firmware related distributable. * &#x60;software.HyperflexBundleDistributable&#x60; - Object type for HyperFlex related distributable. * &#x60;software.ApplianceDistributableConnected&#x60; - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Connected appliance.  * &#x60;software.ApplianceDistributablePrivate&#x60; - Object type for Appliance related distributable. By Choosing this, user shows intent to use the downloaded image for Private appliance.  * &#x60;software.UcsdBundleDistributable&#x60; - Object type for UCSD related distributable. * &#x60;software.HciBundleDistributable&#x60; - Object type for HCI Bundle related distributable. | [optional] [default to "None"]
**OdataFilter** | Pointer to **string** | Odata filter string managed internally. It is built with specific ObjectType properties. | [optional] [readonly] 

## Methods

### NewNotificationDistributableMoCondition

`func NewNotificationDistributableMoCondition(classId string, objectType string, ) *NotificationDistributableMoCondition`

NewNotificationDistributableMoCondition instantiates a new NotificationDistributableMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDistributableMoConditionWithDefaults

`func NewNotificationDistributableMoConditionWithDefaults() *NotificationDistributableMoCondition`

NewNotificationDistributableMoConditionWithDefaults instantiates a new NotificationDistributableMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationDistributableMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationDistributableMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationDistributableMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationDistributableMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationDistributableMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationDistributableMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageType

`func (o *NotificationDistributableMoCondition) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *NotificationDistributableMoCondition) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *NotificationDistributableMoCondition) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *NotificationDistributableMoCondition) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOdataFilter

`func (o *NotificationDistributableMoCondition) GetOdataFilter() string`

GetOdataFilter returns the OdataFilter field if non-nil, zero value otherwise.

### GetOdataFilterOk

`func (o *NotificationDistributableMoCondition) GetOdataFilterOk() (*string, bool)`

GetOdataFilterOk returns a tuple with the OdataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilter

`func (o *NotificationDistributableMoCondition) SetOdataFilter(v string)`

SetOdataFilter sets OdataFilter field to given value.

### HasOdataFilter

`func (o *NotificationDistributableMoCondition) HasOdataFilter() bool`

HasOdataFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


