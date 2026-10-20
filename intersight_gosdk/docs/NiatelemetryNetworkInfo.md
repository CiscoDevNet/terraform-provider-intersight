# NiatelemetryNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NetworkInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NetworkInfo"]
**ActiveNode** | Pointer to [**NullableNiatelemetryNode**](NiatelemetryNode.md) |  | [optional] 
**Hostname** | Pointer to **string** | Returns hostname of the network. | [optional] 
**ManagementtIp** | Pointer to **string** | Returns management IP of the network. | [optional] 
**OutofbandIp** | Pointer to **string** | Returns out of band IP of the network. | [optional] 
**StandbyNode** | Pointer to [**NullableNiatelemetryNode**](NiatelemetryNode.md) |  | [optional] 

## Methods

### NewNiatelemetryNetworkInfo

`func NewNiatelemetryNetworkInfo(classId string, objectType string, ) *NiatelemetryNetworkInfo`

NewNiatelemetryNetworkInfo instantiates a new NiatelemetryNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNetworkInfoWithDefaults

`func NewNiatelemetryNetworkInfoWithDefaults() *NiatelemetryNetworkInfo`

NewNiatelemetryNetworkInfoWithDefaults instantiates a new NiatelemetryNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNetworkInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNetworkInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNetworkInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNetworkInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNetworkInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNetworkInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveNode

`func (o *NiatelemetryNetworkInfo) GetActiveNode() NiatelemetryNode`

GetActiveNode returns the ActiveNode field if non-nil, zero value otherwise.

### GetActiveNodeOk

`func (o *NiatelemetryNetworkInfo) GetActiveNodeOk() (*NiatelemetryNode, bool)`

GetActiveNodeOk returns a tuple with the ActiveNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNode

`func (o *NiatelemetryNetworkInfo) SetActiveNode(v NiatelemetryNode)`

SetActiveNode sets ActiveNode field to given value.

### HasActiveNode

`func (o *NiatelemetryNetworkInfo) HasActiveNode() bool`

HasActiveNode returns a boolean if a field has been set.

### SetActiveNodeNil

`func (o *NiatelemetryNetworkInfo) SetActiveNodeNil(b bool)`

 SetActiveNodeNil sets the value for ActiveNode to be an explicit nil

### UnsetActiveNode
`func (o *NiatelemetryNetworkInfo) UnsetActiveNode()`

UnsetActiveNode ensures that no value is present for ActiveNode, not even an explicit nil
### GetHostname

`func (o *NiatelemetryNetworkInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NiatelemetryNetworkInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NiatelemetryNetworkInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NiatelemetryNetworkInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetManagementtIp

`func (o *NiatelemetryNetworkInfo) GetManagementtIp() string`

GetManagementtIp returns the ManagementtIp field if non-nil, zero value otherwise.

### GetManagementtIpOk

`func (o *NiatelemetryNetworkInfo) GetManagementtIpOk() (*string, bool)`

GetManagementtIpOk returns a tuple with the ManagementtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementtIp

`func (o *NiatelemetryNetworkInfo) SetManagementtIp(v string)`

SetManagementtIp sets ManagementtIp field to given value.

### HasManagementtIp

`func (o *NiatelemetryNetworkInfo) HasManagementtIp() bool`

HasManagementtIp returns a boolean if a field has been set.

### GetOutofbandIp

`func (o *NiatelemetryNetworkInfo) GetOutofbandIp() string`

GetOutofbandIp returns the OutofbandIp field if non-nil, zero value otherwise.

### GetOutofbandIpOk

`func (o *NiatelemetryNetworkInfo) GetOutofbandIpOk() (*string, bool)`

GetOutofbandIpOk returns a tuple with the OutofbandIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutofbandIp

`func (o *NiatelemetryNetworkInfo) SetOutofbandIp(v string)`

SetOutofbandIp sets OutofbandIp field to given value.

### HasOutofbandIp

`func (o *NiatelemetryNetworkInfo) HasOutofbandIp() bool`

HasOutofbandIp returns a boolean if a field has been set.

### GetStandbyNode

`func (o *NiatelemetryNetworkInfo) GetStandbyNode() NiatelemetryNode`

GetStandbyNode returns the StandbyNode field if non-nil, zero value otherwise.

### GetStandbyNodeOk

`func (o *NiatelemetryNetworkInfo) GetStandbyNodeOk() (*NiatelemetryNode, bool)`

GetStandbyNodeOk returns a tuple with the StandbyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyNode

`func (o *NiatelemetryNetworkInfo) SetStandbyNode(v NiatelemetryNode)`

SetStandbyNode sets StandbyNode field to given value.

### HasStandbyNode

`func (o *NiatelemetryNetworkInfo) HasStandbyNode() bool`

HasStandbyNode returns a boolean if a field has been set.

### SetStandbyNodeNil

`func (o *NiatelemetryNetworkInfo) SetStandbyNodeNil(b bool)`

 SetStandbyNodeNil sets the value for StandbyNode to be an explicit nil

### UnsetStandbyNode
`func (o *NiatelemetryNetworkInfo) UnsetStandbyNode()`

UnsetStandbyNode ensures that no value is present for StandbyNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


