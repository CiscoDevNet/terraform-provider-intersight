# NiatelemetryNetworkInfoAllOf

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

### NewNiatelemetryNetworkInfoAllOf

`func NewNiatelemetryNetworkInfoAllOf(classId string, objectType string, ) *NiatelemetryNetworkInfoAllOf`

NewNiatelemetryNetworkInfoAllOf instantiates a new NiatelemetryNetworkInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNetworkInfoAllOfWithDefaults

`func NewNiatelemetryNetworkInfoAllOfWithDefaults() *NiatelemetryNetworkInfoAllOf`

NewNiatelemetryNetworkInfoAllOfWithDefaults instantiates a new NiatelemetryNetworkInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNetworkInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNetworkInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNetworkInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNetworkInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNetworkInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNetworkInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveNode

`func (o *NiatelemetryNetworkInfoAllOf) GetActiveNode() NiatelemetryNode`

GetActiveNode returns the ActiveNode field if non-nil, zero value otherwise.

### GetActiveNodeOk

`func (o *NiatelemetryNetworkInfoAllOf) GetActiveNodeOk() (*NiatelemetryNode, bool)`

GetActiveNodeOk returns a tuple with the ActiveNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNode

`func (o *NiatelemetryNetworkInfoAllOf) SetActiveNode(v NiatelemetryNode)`

SetActiveNode sets ActiveNode field to given value.

### HasActiveNode

`func (o *NiatelemetryNetworkInfoAllOf) HasActiveNode() bool`

HasActiveNode returns a boolean if a field has been set.

### SetActiveNodeNil

`func (o *NiatelemetryNetworkInfoAllOf) SetActiveNodeNil(b bool)`

 SetActiveNodeNil sets the value for ActiveNode to be an explicit nil

### UnsetActiveNode
`func (o *NiatelemetryNetworkInfoAllOf) UnsetActiveNode()`

UnsetActiveNode ensures that no value is present for ActiveNode, not even an explicit nil
### GetHostname

`func (o *NiatelemetryNetworkInfoAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NiatelemetryNetworkInfoAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NiatelemetryNetworkInfoAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NiatelemetryNetworkInfoAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetManagementtIp

`func (o *NiatelemetryNetworkInfoAllOf) GetManagementtIp() string`

GetManagementtIp returns the ManagementtIp field if non-nil, zero value otherwise.

### GetManagementtIpOk

`func (o *NiatelemetryNetworkInfoAllOf) GetManagementtIpOk() (*string, bool)`

GetManagementtIpOk returns a tuple with the ManagementtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementtIp

`func (o *NiatelemetryNetworkInfoAllOf) SetManagementtIp(v string)`

SetManagementtIp sets ManagementtIp field to given value.

### HasManagementtIp

`func (o *NiatelemetryNetworkInfoAllOf) HasManagementtIp() bool`

HasManagementtIp returns a boolean if a field has been set.

### GetOutofbandIp

`func (o *NiatelemetryNetworkInfoAllOf) GetOutofbandIp() string`

GetOutofbandIp returns the OutofbandIp field if non-nil, zero value otherwise.

### GetOutofbandIpOk

`func (o *NiatelemetryNetworkInfoAllOf) GetOutofbandIpOk() (*string, bool)`

GetOutofbandIpOk returns a tuple with the OutofbandIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutofbandIp

`func (o *NiatelemetryNetworkInfoAllOf) SetOutofbandIp(v string)`

SetOutofbandIp sets OutofbandIp field to given value.

### HasOutofbandIp

`func (o *NiatelemetryNetworkInfoAllOf) HasOutofbandIp() bool`

HasOutofbandIp returns a boolean if a field has been set.

### GetStandbyNode

`func (o *NiatelemetryNetworkInfoAllOf) GetStandbyNode() NiatelemetryNode`

GetStandbyNode returns the StandbyNode field if non-nil, zero value otherwise.

### GetStandbyNodeOk

`func (o *NiatelemetryNetworkInfoAllOf) GetStandbyNodeOk() (*NiatelemetryNode, bool)`

GetStandbyNodeOk returns a tuple with the StandbyNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyNode

`func (o *NiatelemetryNetworkInfoAllOf) SetStandbyNode(v NiatelemetryNode)`

SetStandbyNode sets StandbyNode field to given value.

### HasStandbyNode

`func (o *NiatelemetryNetworkInfoAllOf) HasStandbyNode() bool`

HasStandbyNode returns a boolean if a field has been set.

### SetStandbyNodeNil

`func (o *NiatelemetryNetworkInfoAllOf) SetStandbyNodeNil(b bool)`

 SetStandbyNodeNil sets the value for StandbyNode to be an explicit nil

### UnsetStandbyNode
`func (o *NiatelemetryNetworkInfoAllOf) UnsetStandbyNode()`

UnsetStandbyNode ensures that no value is present for StandbyNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


