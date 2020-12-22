# VirtualizationCloudInitConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.CloudInitConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.CloudInitConfig"]
**ConfigType** | Pointer to **string** | Virtual machine cloud init configuration type. * &#x60;&#x60; - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected. * &#x60;NoCloudSource&#x60; - Allows the user to provide user-data to the instance without running a network service. * &#x60;CloudConfigDrive&#x60; - Allows the user to provide user-data and network-data from cloud. | [optional] [default to ""]
**NetworkData** | Pointer to **string** | Network configuration data for a virtual machine. | [optional] 
**NetworkDataBase64Encoded** | Pointer to **bool** | Set to true, if the cloud init network data is in base64 format. | [optional] 
**UserData** | Pointer to **string** | User configuration data for a virtual machine such as adding user, group etc. | [optional] 
**UserDataBase64Encoded** | Pointer to **bool** | Set to true, if the cloud init user data is in base64 format. | [optional] 

## Methods

### NewVirtualizationCloudInitConfigAllOf

`func NewVirtualizationCloudInitConfigAllOf(classId string, objectType string, ) *VirtualizationCloudInitConfigAllOf`

NewVirtualizationCloudInitConfigAllOf instantiates a new VirtualizationCloudInitConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCloudInitConfigAllOfWithDefaults

`func NewVirtualizationCloudInitConfigAllOfWithDefaults() *VirtualizationCloudInitConfigAllOf`

NewVirtualizationCloudInitConfigAllOfWithDefaults instantiates a new VirtualizationCloudInitConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCloudInitConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCloudInitConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCloudInitConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCloudInitConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCloudInitConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCloudInitConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigType

`func (o *VirtualizationCloudInitConfigAllOf) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *VirtualizationCloudInitConfigAllOf) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *VirtualizationCloudInitConfigAllOf) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.

### HasConfigType

`func (o *VirtualizationCloudInitConfigAllOf) HasConfigType() bool`

HasConfigType returns a boolean if a field has been set.

### GetNetworkData

`func (o *VirtualizationCloudInitConfigAllOf) GetNetworkData() string`

GetNetworkData returns the NetworkData field if non-nil, zero value otherwise.

### GetNetworkDataOk

`func (o *VirtualizationCloudInitConfigAllOf) GetNetworkDataOk() (*string, bool)`

GetNetworkDataOk returns a tuple with the NetworkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkData

`func (o *VirtualizationCloudInitConfigAllOf) SetNetworkData(v string)`

SetNetworkData sets NetworkData field to given value.

### HasNetworkData

`func (o *VirtualizationCloudInitConfigAllOf) HasNetworkData() bool`

HasNetworkData returns a boolean if a field has been set.

### GetNetworkDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) GetNetworkDataBase64Encoded() bool`

GetNetworkDataBase64Encoded returns the NetworkDataBase64Encoded field if non-nil, zero value otherwise.

### GetNetworkDataBase64EncodedOk

`func (o *VirtualizationCloudInitConfigAllOf) GetNetworkDataBase64EncodedOk() (*bool, bool)`

GetNetworkDataBase64EncodedOk returns a tuple with the NetworkDataBase64Encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) SetNetworkDataBase64Encoded(v bool)`

SetNetworkDataBase64Encoded sets NetworkDataBase64Encoded field to given value.

### HasNetworkDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) HasNetworkDataBase64Encoded() bool`

HasNetworkDataBase64Encoded returns a boolean if a field has been set.

### GetUserData

`func (o *VirtualizationCloudInitConfigAllOf) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VirtualizationCloudInitConfigAllOf) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VirtualizationCloudInitConfigAllOf) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *VirtualizationCloudInitConfigAllOf) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetUserDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) GetUserDataBase64Encoded() bool`

GetUserDataBase64Encoded returns the UserDataBase64Encoded field if non-nil, zero value otherwise.

### GetUserDataBase64EncodedOk

`func (o *VirtualizationCloudInitConfigAllOf) GetUserDataBase64EncodedOk() (*bool, bool)`

GetUserDataBase64EncodedOk returns a tuple with the UserDataBase64Encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) SetUserDataBase64Encoded(v bool)`

SetUserDataBase64Encoded sets UserDataBase64Encoded field to given value.

### HasUserDataBase64Encoded

`func (o *VirtualizationCloudInitConfigAllOf) HasUserDataBase64Encoded() bool`

HasUserDataBase64Encoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


