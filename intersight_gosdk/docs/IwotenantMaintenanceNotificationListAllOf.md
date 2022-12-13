# IwotenantMaintenanceNotificationListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;iwotenant.MaintenanceNotification&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IwotenantMaintenanceNotification**](IwotenantMaintenanceNotification.md) | The array of &#39;iwotenant.MaintenanceNotification&#39; resources matching the request. | [optional] 

## Methods

### NewIwotenantMaintenanceNotificationListAllOf

`func NewIwotenantMaintenanceNotificationListAllOf() *IwotenantMaintenanceNotificationListAllOf`

NewIwotenantMaintenanceNotificationListAllOf instantiates a new IwotenantMaintenanceNotificationListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIwotenantMaintenanceNotificationListAllOfWithDefaults

`func NewIwotenantMaintenanceNotificationListAllOfWithDefaults() *IwotenantMaintenanceNotificationListAllOf`

NewIwotenantMaintenanceNotificationListAllOfWithDefaults instantiates a new IwotenantMaintenanceNotificationListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IwotenantMaintenanceNotificationListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IwotenantMaintenanceNotificationListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IwotenantMaintenanceNotificationListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IwotenantMaintenanceNotificationListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IwotenantMaintenanceNotificationListAllOf) GetResults() []IwotenantMaintenanceNotification`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IwotenantMaintenanceNotificationListAllOf) GetResultsOk() (*[]IwotenantMaintenanceNotification, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IwotenantMaintenanceNotificationListAllOf) SetResults(v []IwotenantMaintenanceNotification)`

SetResults sets Results field to given value.

### HasResults

`func (o *IwotenantMaintenanceNotificationListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *IwotenantMaintenanceNotificationListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *IwotenantMaintenanceNotificationListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


