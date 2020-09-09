# NiatelemetryTenantAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdIfPolCount** | Pointer to **int64** | Bidirectional Forwarding Detection bfdIfPol Model Object count scale. | [optional] 
**BfdIfpCount** | Pointer to **int64** | Bidirectional Forwarding Detection Interface Policy count scale. | [optional] 
**Dn** | Pointer to **string** | Dn for the tenants present. | [optional] 
**FhsBdPolCount** | Pointer to **int64** | First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection. | [optional] 
**FvRsBdToFhsCount** | Pointer to **int64** | First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection. | [optional] 
**FvSiteConnpCount** | Pointer to **int64** | Multi-Site scale for fvSiteConnp Model Object. | [optional] 
**L3MulticastCount** | Pointer to **int64** | Number of layer 3 multicasts. | [optional] 
**L3MulticastCtxCount** | Pointer to **int64** | Number of layer 3 multicast CtxP. | [optional] 
**L3MulticastIfCount** | Pointer to **int64** | Number of layer 3 multicast IfP. | [optional] 
**L3outCount** | Pointer to **int64** | L3 out scale for the tenants present. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. | [optional] 
**Ssm** | Pointer to **string** | SSM property feature usage. | [optional] 
**TraceRouteEpCount** | Pointer to **int64** | Number of ITrace route endpoint per tenant. | [optional] 
**TraceRouteEpExtCount** | Pointer to **int64** | Number of ITrace endpoint external routes per tenant. | [optional] 
**TraceRouteExtEpCount** | Pointer to **int64** | Number of ITrace external endpoint routes per tenant. | [optional] 
**TraceRouteExtExtCount** | Pointer to **int64** | Number of ITrace external routes per tenant. | [optional] 
**VnsAbsGraphCount** | Pointer to **int64** | L4 to L7 Services graph count scale. | [optional] 
**VnsBackupPolCount** | Pointer to **int64** | Policy Based Routing standby Node count scale. Checks the Policy Based Routing backup policy. | [optional] 
**VnsRedirectDestCount** | Pointer to **int64** | Policy Based Routing standby Node count scale. Policy based redirect requires a destination to redirect traffic. | [optional] 
**VnsSvcRedirectPolCount** | Pointer to **int64** | Policy Based Routing and Policy Based Service Insertion count scale. Includes without L4-L7 package. | [optional] 
**VrfCount** | Pointer to **int64** | Vrf scale count per tenant. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryTenantAllOf

`func NewNiatelemetryTenantAllOf() *NiatelemetryTenantAllOf`

NewNiatelemetryTenantAllOf instantiates a new NiatelemetryTenantAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryTenantAllOfWithDefaults

`func NewNiatelemetryTenantAllOfWithDefaults() *NiatelemetryTenantAllOf`

NewNiatelemetryTenantAllOfWithDefaults instantiates a new NiatelemetryTenantAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdIfPolCount

`func (o *NiatelemetryTenantAllOf) GetBfdIfPolCount() int64`

GetBfdIfPolCount returns the BfdIfPolCount field if non-nil, zero value otherwise.

### GetBfdIfPolCountOk

`func (o *NiatelemetryTenantAllOf) GetBfdIfPolCountOk() (*int64, bool)`

GetBfdIfPolCountOk returns a tuple with the BfdIfPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfPolCount

`func (o *NiatelemetryTenantAllOf) SetBfdIfPolCount(v int64)`

SetBfdIfPolCount sets BfdIfPolCount field to given value.

### HasBfdIfPolCount

`func (o *NiatelemetryTenantAllOf) HasBfdIfPolCount() bool`

HasBfdIfPolCount returns a boolean if a field has been set.

### GetBfdIfpCount

`func (o *NiatelemetryTenantAllOf) GetBfdIfpCount() int64`

GetBfdIfpCount returns the BfdIfpCount field if non-nil, zero value otherwise.

### GetBfdIfpCountOk

`func (o *NiatelemetryTenantAllOf) GetBfdIfpCountOk() (*int64, bool)`

GetBfdIfpCountOk returns a tuple with the BfdIfpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfpCount

`func (o *NiatelemetryTenantAllOf) SetBfdIfpCount(v int64)`

SetBfdIfpCount sets BfdIfpCount field to given value.

### HasBfdIfpCount

`func (o *NiatelemetryTenantAllOf) HasBfdIfpCount() bool`

HasBfdIfpCount returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryTenantAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryTenantAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryTenantAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryTenantAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFhsBdPolCount

`func (o *NiatelemetryTenantAllOf) GetFhsBdPolCount() int64`

GetFhsBdPolCount returns the FhsBdPolCount field if non-nil, zero value otherwise.

### GetFhsBdPolCountOk

`func (o *NiatelemetryTenantAllOf) GetFhsBdPolCountOk() (*int64, bool)`

GetFhsBdPolCountOk returns a tuple with the FhsBdPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhsBdPolCount

`func (o *NiatelemetryTenantAllOf) SetFhsBdPolCount(v int64)`

SetFhsBdPolCount sets FhsBdPolCount field to given value.

### HasFhsBdPolCount

`func (o *NiatelemetryTenantAllOf) HasFhsBdPolCount() bool`

HasFhsBdPolCount returns a boolean if a field has been set.

### GetFvRsBdToFhsCount

`func (o *NiatelemetryTenantAllOf) GetFvRsBdToFhsCount() int64`

GetFvRsBdToFhsCount returns the FvRsBdToFhsCount field if non-nil, zero value otherwise.

### GetFvRsBdToFhsCountOk

`func (o *NiatelemetryTenantAllOf) GetFvRsBdToFhsCountOk() (*int64, bool)`

GetFvRsBdToFhsCountOk returns a tuple with the FvRsBdToFhsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsBdToFhsCount

`func (o *NiatelemetryTenantAllOf) SetFvRsBdToFhsCount(v int64)`

SetFvRsBdToFhsCount sets FvRsBdToFhsCount field to given value.

### HasFvRsBdToFhsCount

`func (o *NiatelemetryTenantAllOf) HasFvRsBdToFhsCount() bool`

HasFvRsBdToFhsCount returns a boolean if a field has been set.

### GetFvSiteConnpCount

`func (o *NiatelemetryTenantAllOf) GetFvSiteConnpCount() int64`

GetFvSiteConnpCount returns the FvSiteConnpCount field if non-nil, zero value otherwise.

### GetFvSiteConnpCountOk

`func (o *NiatelemetryTenantAllOf) GetFvSiteConnpCountOk() (*int64, bool)`

GetFvSiteConnpCountOk returns a tuple with the FvSiteConnpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvSiteConnpCount

`func (o *NiatelemetryTenantAllOf) SetFvSiteConnpCount(v int64)`

SetFvSiteConnpCount sets FvSiteConnpCount field to given value.

### HasFvSiteConnpCount

`func (o *NiatelemetryTenantAllOf) HasFvSiteConnpCount() bool`

HasFvSiteConnpCount returns a boolean if a field has been set.

### GetL3MulticastCount

`func (o *NiatelemetryTenantAllOf) GetL3MulticastCount() int64`

GetL3MulticastCount returns the L3MulticastCount field if non-nil, zero value otherwise.

### GetL3MulticastCountOk

`func (o *NiatelemetryTenantAllOf) GetL3MulticastCountOk() (*int64, bool)`

GetL3MulticastCountOk returns a tuple with the L3MulticastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCount

`func (o *NiatelemetryTenantAllOf) SetL3MulticastCount(v int64)`

SetL3MulticastCount sets L3MulticastCount field to given value.

### HasL3MulticastCount

`func (o *NiatelemetryTenantAllOf) HasL3MulticastCount() bool`

HasL3MulticastCount returns a boolean if a field has been set.

### GetL3MulticastCtxCount

`func (o *NiatelemetryTenantAllOf) GetL3MulticastCtxCount() int64`

GetL3MulticastCtxCount returns the L3MulticastCtxCount field if non-nil, zero value otherwise.

### GetL3MulticastCtxCountOk

`func (o *NiatelemetryTenantAllOf) GetL3MulticastCtxCountOk() (*int64, bool)`

GetL3MulticastCtxCountOk returns a tuple with the L3MulticastCtxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCtxCount

`func (o *NiatelemetryTenantAllOf) SetL3MulticastCtxCount(v int64)`

SetL3MulticastCtxCount sets L3MulticastCtxCount field to given value.

### HasL3MulticastCtxCount

`func (o *NiatelemetryTenantAllOf) HasL3MulticastCtxCount() bool`

HasL3MulticastCtxCount returns a boolean if a field has been set.

### GetL3MulticastIfCount

`func (o *NiatelemetryTenantAllOf) GetL3MulticastIfCount() int64`

GetL3MulticastIfCount returns the L3MulticastIfCount field if non-nil, zero value otherwise.

### GetL3MulticastIfCountOk

`func (o *NiatelemetryTenantAllOf) GetL3MulticastIfCountOk() (*int64, bool)`

GetL3MulticastIfCountOk returns a tuple with the L3MulticastIfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastIfCount

`func (o *NiatelemetryTenantAllOf) SetL3MulticastIfCount(v int64)`

SetL3MulticastIfCount sets L3MulticastIfCount field to given value.

### HasL3MulticastIfCount

`func (o *NiatelemetryTenantAllOf) HasL3MulticastIfCount() bool`

HasL3MulticastIfCount returns a boolean if a field has been set.

### GetL3outCount

`func (o *NiatelemetryTenantAllOf) GetL3outCount() int64`

GetL3outCount returns the L3outCount field if non-nil, zero value otherwise.

### GetL3outCountOk

`func (o *NiatelemetryTenantAllOf) GetL3outCountOk() (*int64, bool)`

GetL3outCountOk returns a tuple with the L3outCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3outCount

`func (o *NiatelemetryTenantAllOf) SetL3outCount(v int64)`

SetL3outCount sets L3outCount field to given value.

### HasL3outCount

`func (o *NiatelemetryTenantAllOf) HasL3outCount() bool`

HasL3outCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryTenantAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryTenantAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryTenantAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryTenantAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSsm

`func (o *NiatelemetryTenantAllOf) GetSsm() string`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *NiatelemetryTenantAllOf) GetSsmOk() (*string, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *NiatelemetryTenantAllOf) SetSsm(v string)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *NiatelemetryTenantAllOf) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetTraceRouteEpCount

`func (o *NiatelemetryTenantAllOf) GetTraceRouteEpCount() int64`

GetTraceRouteEpCount returns the TraceRouteEpCount field if non-nil, zero value otherwise.

### GetTraceRouteEpCountOk

`func (o *NiatelemetryTenantAllOf) GetTraceRouteEpCountOk() (*int64, bool)`

GetTraceRouteEpCountOk returns a tuple with the TraceRouteEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpCount

`func (o *NiatelemetryTenantAllOf) SetTraceRouteEpCount(v int64)`

SetTraceRouteEpCount sets TraceRouteEpCount field to given value.

### HasTraceRouteEpCount

`func (o *NiatelemetryTenantAllOf) HasTraceRouteEpCount() bool`

HasTraceRouteEpCount returns a boolean if a field has been set.

### GetTraceRouteEpExtCount

`func (o *NiatelemetryTenantAllOf) GetTraceRouteEpExtCount() int64`

GetTraceRouteEpExtCount returns the TraceRouteEpExtCount field if non-nil, zero value otherwise.

### GetTraceRouteEpExtCountOk

`func (o *NiatelemetryTenantAllOf) GetTraceRouteEpExtCountOk() (*int64, bool)`

GetTraceRouteEpExtCountOk returns a tuple with the TraceRouteEpExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpExtCount

`func (o *NiatelemetryTenantAllOf) SetTraceRouteEpExtCount(v int64)`

SetTraceRouteEpExtCount sets TraceRouteEpExtCount field to given value.

### HasTraceRouteEpExtCount

`func (o *NiatelemetryTenantAllOf) HasTraceRouteEpExtCount() bool`

HasTraceRouteEpExtCount returns a boolean if a field has been set.

### GetTraceRouteExtEpCount

`func (o *NiatelemetryTenantAllOf) GetTraceRouteExtEpCount() int64`

GetTraceRouteExtEpCount returns the TraceRouteExtEpCount field if non-nil, zero value otherwise.

### GetTraceRouteExtEpCountOk

`func (o *NiatelemetryTenantAllOf) GetTraceRouteExtEpCountOk() (*int64, bool)`

GetTraceRouteExtEpCountOk returns a tuple with the TraceRouteExtEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtEpCount

`func (o *NiatelemetryTenantAllOf) SetTraceRouteExtEpCount(v int64)`

SetTraceRouteExtEpCount sets TraceRouteExtEpCount field to given value.

### HasTraceRouteExtEpCount

`func (o *NiatelemetryTenantAllOf) HasTraceRouteExtEpCount() bool`

HasTraceRouteExtEpCount returns a boolean if a field has been set.

### GetTraceRouteExtExtCount

`func (o *NiatelemetryTenantAllOf) GetTraceRouteExtExtCount() int64`

GetTraceRouteExtExtCount returns the TraceRouteExtExtCount field if non-nil, zero value otherwise.

### GetTraceRouteExtExtCountOk

`func (o *NiatelemetryTenantAllOf) GetTraceRouteExtExtCountOk() (*int64, bool)`

GetTraceRouteExtExtCountOk returns a tuple with the TraceRouteExtExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtExtCount

`func (o *NiatelemetryTenantAllOf) SetTraceRouteExtExtCount(v int64)`

SetTraceRouteExtExtCount sets TraceRouteExtExtCount field to given value.

### HasTraceRouteExtExtCount

`func (o *NiatelemetryTenantAllOf) HasTraceRouteExtExtCount() bool`

HasTraceRouteExtExtCount returns a boolean if a field has been set.

### GetVnsAbsGraphCount

`func (o *NiatelemetryTenantAllOf) GetVnsAbsGraphCount() int64`

GetVnsAbsGraphCount returns the VnsAbsGraphCount field if non-nil, zero value otherwise.

### GetVnsAbsGraphCountOk

`func (o *NiatelemetryTenantAllOf) GetVnsAbsGraphCountOk() (*int64, bool)`

GetVnsAbsGraphCountOk returns a tuple with the VnsAbsGraphCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsAbsGraphCount

`func (o *NiatelemetryTenantAllOf) SetVnsAbsGraphCount(v int64)`

SetVnsAbsGraphCount sets VnsAbsGraphCount field to given value.

### HasVnsAbsGraphCount

`func (o *NiatelemetryTenantAllOf) HasVnsAbsGraphCount() bool`

HasVnsAbsGraphCount returns a boolean if a field has been set.

### GetVnsBackupPolCount

`func (o *NiatelemetryTenantAllOf) GetVnsBackupPolCount() int64`

GetVnsBackupPolCount returns the VnsBackupPolCount field if non-nil, zero value otherwise.

### GetVnsBackupPolCountOk

`func (o *NiatelemetryTenantAllOf) GetVnsBackupPolCountOk() (*int64, bool)`

GetVnsBackupPolCountOk returns a tuple with the VnsBackupPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsBackupPolCount

`func (o *NiatelemetryTenantAllOf) SetVnsBackupPolCount(v int64)`

SetVnsBackupPolCount sets VnsBackupPolCount field to given value.

### HasVnsBackupPolCount

`func (o *NiatelemetryTenantAllOf) HasVnsBackupPolCount() bool`

HasVnsBackupPolCount returns a boolean if a field has been set.

### GetVnsRedirectDestCount

`func (o *NiatelemetryTenantAllOf) GetVnsRedirectDestCount() int64`

GetVnsRedirectDestCount returns the VnsRedirectDestCount field if non-nil, zero value otherwise.

### GetVnsRedirectDestCountOk

`func (o *NiatelemetryTenantAllOf) GetVnsRedirectDestCountOk() (*int64, bool)`

GetVnsRedirectDestCountOk returns a tuple with the VnsRedirectDestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsRedirectDestCount

`func (o *NiatelemetryTenantAllOf) SetVnsRedirectDestCount(v int64)`

SetVnsRedirectDestCount sets VnsRedirectDestCount field to given value.

### HasVnsRedirectDestCount

`func (o *NiatelemetryTenantAllOf) HasVnsRedirectDestCount() bool`

HasVnsRedirectDestCount returns a boolean if a field has been set.

### GetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenantAllOf) GetVnsSvcRedirectPolCount() int64`

GetVnsSvcRedirectPolCount returns the VnsSvcRedirectPolCount field if non-nil, zero value otherwise.

### GetVnsSvcRedirectPolCountOk

`func (o *NiatelemetryTenantAllOf) GetVnsSvcRedirectPolCountOk() (*int64, bool)`

GetVnsSvcRedirectPolCountOk returns a tuple with the VnsSvcRedirectPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenantAllOf) SetVnsSvcRedirectPolCount(v int64)`

SetVnsSvcRedirectPolCount sets VnsSvcRedirectPolCount field to given value.

### HasVnsSvcRedirectPolCount

`func (o *NiatelemetryTenantAllOf) HasVnsSvcRedirectPolCount() bool`

HasVnsSvcRedirectPolCount returns a boolean if a field has been set.

### GetVrfCount

`func (o *NiatelemetryTenantAllOf) GetVrfCount() int64`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *NiatelemetryTenantAllOf) GetVrfCountOk() (*int64, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *NiatelemetryTenantAllOf) SetVrfCount(v int64)`

SetVrfCount sets VrfCount field to given value.

### HasVrfCount

`func (o *NiatelemetryTenantAllOf) HasVrfCount() bool`

HasVrfCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryTenantAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryTenantAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryTenantAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryTenantAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


