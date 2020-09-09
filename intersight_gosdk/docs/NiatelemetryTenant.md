# NiatelemetryTenant

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

### NewNiatelemetryTenant

`func NewNiatelemetryTenant() *NiatelemetryTenant`

NewNiatelemetryTenant instantiates a new NiatelemetryTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryTenantWithDefaults

`func NewNiatelemetryTenantWithDefaults() *NiatelemetryTenant`

NewNiatelemetryTenantWithDefaults instantiates a new NiatelemetryTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfdIfPolCount

`func (o *NiatelemetryTenant) GetBfdIfPolCount() int64`

GetBfdIfPolCount returns the BfdIfPolCount field if non-nil, zero value otherwise.

### GetBfdIfPolCountOk

`func (o *NiatelemetryTenant) GetBfdIfPolCountOk() (*int64, bool)`

GetBfdIfPolCountOk returns a tuple with the BfdIfPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfPolCount

`func (o *NiatelemetryTenant) SetBfdIfPolCount(v int64)`

SetBfdIfPolCount sets BfdIfPolCount field to given value.

### HasBfdIfPolCount

`func (o *NiatelemetryTenant) HasBfdIfPolCount() bool`

HasBfdIfPolCount returns a boolean if a field has been set.

### GetBfdIfpCount

`func (o *NiatelemetryTenant) GetBfdIfpCount() int64`

GetBfdIfpCount returns the BfdIfpCount field if non-nil, zero value otherwise.

### GetBfdIfpCountOk

`func (o *NiatelemetryTenant) GetBfdIfpCountOk() (*int64, bool)`

GetBfdIfpCountOk returns a tuple with the BfdIfpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdIfpCount

`func (o *NiatelemetryTenant) SetBfdIfpCount(v int64)`

SetBfdIfpCount sets BfdIfpCount field to given value.

### HasBfdIfpCount

`func (o *NiatelemetryTenant) HasBfdIfpCount() bool`

HasBfdIfpCount returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryTenant) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryTenant) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryTenant) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryTenant) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFhsBdPolCount

`func (o *NiatelemetryTenant) GetFhsBdPolCount() int64`

GetFhsBdPolCount returns the FhsBdPolCount field if non-nil, zero value otherwise.

### GetFhsBdPolCountOk

`func (o *NiatelemetryTenant) GetFhsBdPolCountOk() (*int64, bool)`

GetFhsBdPolCountOk returns a tuple with the FhsBdPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFhsBdPolCount

`func (o *NiatelemetryTenant) SetFhsBdPolCount(v int64)`

SetFhsBdPolCount sets FhsBdPolCount field to given value.

### HasFhsBdPolCount

`func (o *NiatelemetryTenant) HasFhsBdPolCount() bool`

HasFhsBdPolCount returns a boolean if a field has been set.

### GetFvRsBdToFhsCount

`func (o *NiatelemetryTenant) GetFvRsBdToFhsCount() int64`

GetFvRsBdToFhsCount returns the FvRsBdToFhsCount field if non-nil, zero value otherwise.

### GetFvRsBdToFhsCountOk

`func (o *NiatelemetryTenant) GetFvRsBdToFhsCountOk() (*int64, bool)`

GetFvRsBdToFhsCountOk returns a tuple with the FvRsBdToFhsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsBdToFhsCount

`func (o *NiatelemetryTenant) SetFvRsBdToFhsCount(v int64)`

SetFvRsBdToFhsCount sets FvRsBdToFhsCount field to given value.

### HasFvRsBdToFhsCount

`func (o *NiatelemetryTenant) HasFvRsBdToFhsCount() bool`

HasFvRsBdToFhsCount returns a boolean if a field has been set.

### GetFvSiteConnpCount

`func (o *NiatelemetryTenant) GetFvSiteConnpCount() int64`

GetFvSiteConnpCount returns the FvSiteConnpCount field if non-nil, zero value otherwise.

### GetFvSiteConnpCountOk

`func (o *NiatelemetryTenant) GetFvSiteConnpCountOk() (*int64, bool)`

GetFvSiteConnpCountOk returns a tuple with the FvSiteConnpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvSiteConnpCount

`func (o *NiatelemetryTenant) SetFvSiteConnpCount(v int64)`

SetFvSiteConnpCount sets FvSiteConnpCount field to given value.

### HasFvSiteConnpCount

`func (o *NiatelemetryTenant) HasFvSiteConnpCount() bool`

HasFvSiteConnpCount returns a boolean if a field has been set.

### GetL3MulticastCount

`func (o *NiatelemetryTenant) GetL3MulticastCount() int64`

GetL3MulticastCount returns the L3MulticastCount field if non-nil, zero value otherwise.

### GetL3MulticastCountOk

`func (o *NiatelemetryTenant) GetL3MulticastCountOk() (*int64, bool)`

GetL3MulticastCountOk returns a tuple with the L3MulticastCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCount

`func (o *NiatelemetryTenant) SetL3MulticastCount(v int64)`

SetL3MulticastCount sets L3MulticastCount field to given value.

### HasL3MulticastCount

`func (o *NiatelemetryTenant) HasL3MulticastCount() bool`

HasL3MulticastCount returns a boolean if a field has been set.

### GetL3MulticastCtxCount

`func (o *NiatelemetryTenant) GetL3MulticastCtxCount() int64`

GetL3MulticastCtxCount returns the L3MulticastCtxCount field if non-nil, zero value otherwise.

### GetL3MulticastCtxCountOk

`func (o *NiatelemetryTenant) GetL3MulticastCtxCountOk() (*int64, bool)`

GetL3MulticastCtxCountOk returns a tuple with the L3MulticastCtxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastCtxCount

`func (o *NiatelemetryTenant) SetL3MulticastCtxCount(v int64)`

SetL3MulticastCtxCount sets L3MulticastCtxCount field to given value.

### HasL3MulticastCtxCount

`func (o *NiatelemetryTenant) HasL3MulticastCtxCount() bool`

HasL3MulticastCtxCount returns a boolean if a field has been set.

### GetL3MulticastIfCount

`func (o *NiatelemetryTenant) GetL3MulticastIfCount() int64`

GetL3MulticastIfCount returns the L3MulticastIfCount field if non-nil, zero value otherwise.

### GetL3MulticastIfCountOk

`func (o *NiatelemetryTenant) GetL3MulticastIfCountOk() (*int64, bool)`

GetL3MulticastIfCountOk returns a tuple with the L3MulticastIfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3MulticastIfCount

`func (o *NiatelemetryTenant) SetL3MulticastIfCount(v int64)`

SetL3MulticastIfCount sets L3MulticastIfCount field to given value.

### HasL3MulticastIfCount

`func (o *NiatelemetryTenant) HasL3MulticastIfCount() bool`

HasL3MulticastIfCount returns a boolean if a field has been set.

### GetL3outCount

`func (o *NiatelemetryTenant) GetL3outCount() int64`

GetL3outCount returns the L3outCount field if non-nil, zero value otherwise.

### GetL3outCountOk

`func (o *NiatelemetryTenant) GetL3outCountOk() (*int64, bool)`

GetL3outCountOk returns a tuple with the L3outCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3outCount

`func (o *NiatelemetryTenant) SetL3outCount(v int64)`

SetL3outCount sets L3outCount field to given value.

### HasL3outCount

`func (o *NiatelemetryTenant) HasL3outCount() bool`

HasL3outCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryTenant) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryTenant) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryTenant) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryTenant) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSsm

`func (o *NiatelemetryTenant) GetSsm() string`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *NiatelemetryTenant) GetSsmOk() (*string, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *NiatelemetryTenant) SetSsm(v string)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *NiatelemetryTenant) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetTraceRouteEpCount

`func (o *NiatelemetryTenant) GetTraceRouteEpCount() int64`

GetTraceRouteEpCount returns the TraceRouteEpCount field if non-nil, zero value otherwise.

### GetTraceRouteEpCountOk

`func (o *NiatelemetryTenant) GetTraceRouteEpCountOk() (*int64, bool)`

GetTraceRouteEpCountOk returns a tuple with the TraceRouteEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpCount

`func (o *NiatelemetryTenant) SetTraceRouteEpCount(v int64)`

SetTraceRouteEpCount sets TraceRouteEpCount field to given value.

### HasTraceRouteEpCount

`func (o *NiatelemetryTenant) HasTraceRouteEpCount() bool`

HasTraceRouteEpCount returns a boolean if a field has been set.

### GetTraceRouteEpExtCount

`func (o *NiatelemetryTenant) GetTraceRouteEpExtCount() int64`

GetTraceRouteEpExtCount returns the TraceRouteEpExtCount field if non-nil, zero value otherwise.

### GetTraceRouteEpExtCountOk

`func (o *NiatelemetryTenant) GetTraceRouteEpExtCountOk() (*int64, bool)`

GetTraceRouteEpExtCountOk returns a tuple with the TraceRouteEpExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteEpExtCount

`func (o *NiatelemetryTenant) SetTraceRouteEpExtCount(v int64)`

SetTraceRouteEpExtCount sets TraceRouteEpExtCount field to given value.

### HasTraceRouteEpExtCount

`func (o *NiatelemetryTenant) HasTraceRouteEpExtCount() bool`

HasTraceRouteEpExtCount returns a boolean if a field has been set.

### GetTraceRouteExtEpCount

`func (o *NiatelemetryTenant) GetTraceRouteExtEpCount() int64`

GetTraceRouteExtEpCount returns the TraceRouteExtEpCount field if non-nil, zero value otherwise.

### GetTraceRouteExtEpCountOk

`func (o *NiatelemetryTenant) GetTraceRouteExtEpCountOk() (*int64, bool)`

GetTraceRouteExtEpCountOk returns a tuple with the TraceRouteExtEpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtEpCount

`func (o *NiatelemetryTenant) SetTraceRouteExtEpCount(v int64)`

SetTraceRouteExtEpCount sets TraceRouteExtEpCount field to given value.

### HasTraceRouteExtEpCount

`func (o *NiatelemetryTenant) HasTraceRouteExtEpCount() bool`

HasTraceRouteExtEpCount returns a boolean if a field has been set.

### GetTraceRouteExtExtCount

`func (o *NiatelemetryTenant) GetTraceRouteExtExtCount() int64`

GetTraceRouteExtExtCount returns the TraceRouteExtExtCount field if non-nil, zero value otherwise.

### GetTraceRouteExtExtCountOk

`func (o *NiatelemetryTenant) GetTraceRouteExtExtCountOk() (*int64, bool)`

GetTraceRouteExtExtCountOk returns a tuple with the TraceRouteExtExtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRouteExtExtCount

`func (o *NiatelemetryTenant) SetTraceRouteExtExtCount(v int64)`

SetTraceRouteExtExtCount sets TraceRouteExtExtCount field to given value.

### HasTraceRouteExtExtCount

`func (o *NiatelemetryTenant) HasTraceRouteExtExtCount() bool`

HasTraceRouteExtExtCount returns a boolean if a field has been set.

### GetVnsAbsGraphCount

`func (o *NiatelemetryTenant) GetVnsAbsGraphCount() int64`

GetVnsAbsGraphCount returns the VnsAbsGraphCount field if non-nil, zero value otherwise.

### GetVnsAbsGraphCountOk

`func (o *NiatelemetryTenant) GetVnsAbsGraphCountOk() (*int64, bool)`

GetVnsAbsGraphCountOk returns a tuple with the VnsAbsGraphCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsAbsGraphCount

`func (o *NiatelemetryTenant) SetVnsAbsGraphCount(v int64)`

SetVnsAbsGraphCount sets VnsAbsGraphCount field to given value.

### HasVnsAbsGraphCount

`func (o *NiatelemetryTenant) HasVnsAbsGraphCount() bool`

HasVnsAbsGraphCount returns a boolean if a field has been set.

### GetVnsBackupPolCount

`func (o *NiatelemetryTenant) GetVnsBackupPolCount() int64`

GetVnsBackupPolCount returns the VnsBackupPolCount field if non-nil, zero value otherwise.

### GetVnsBackupPolCountOk

`func (o *NiatelemetryTenant) GetVnsBackupPolCountOk() (*int64, bool)`

GetVnsBackupPolCountOk returns a tuple with the VnsBackupPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsBackupPolCount

`func (o *NiatelemetryTenant) SetVnsBackupPolCount(v int64)`

SetVnsBackupPolCount sets VnsBackupPolCount field to given value.

### HasVnsBackupPolCount

`func (o *NiatelemetryTenant) HasVnsBackupPolCount() bool`

HasVnsBackupPolCount returns a boolean if a field has been set.

### GetVnsRedirectDestCount

`func (o *NiatelemetryTenant) GetVnsRedirectDestCount() int64`

GetVnsRedirectDestCount returns the VnsRedirectDestCount field if non-nil, zero value otherwise.

### GetVnsRedirectDestCountOk

`func (o *NiatelemetryTenant) GetVnsRedirectDestCountOk() (*int64, bool)`

GetVnsRedirectDestCountOk returns a tuple with the VnsRedirectDestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsRedirectDestCount

`func (o *NiatelemetryTenant) SetVnsRedirectDestCount(v int64)`

SetVnsRedirectDestCount sets VnsRedirectDestCount field to given value.

### HasVnsRedirectDestCount

`func (o *NiatelemetryTenant) HasVnsRedirectDestCount() bool`

HasVnsRedirectDestCount returns a boolean if a field has been set.

### GetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) GetVnsSvcRedirectPolCount() int64`

GetVnsSvcRedirectPolCount returns the VnsSvcRedirectPolCount field if non-nil, zero value otherwise.

### GetVnsSvcRedirectPolCountOk

`func (o *NiatelemetryTenant) GetVnsSvcRedirectPolCountOk() (*int64, bool)`

GetVnsSvcRedirectPolCountOk returns a tuple with the VnsSvcRedirectPolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) SetVnsSvcRedirectPolCount(v int64)`

SetVnsSvcRedirectPolCount sets VnsSvcRedirectPolCount field to given value.

### HasVnsSvcRedirectPolCount

`func (o *NiatelemetryTenant) HasVnsSvcRedirectPolCount() bool`

HasVnsSvcRedirectPolCount returns a boolean if a field has been set.

### GetVrfCount

`func (o *NiatelemetryTenant) GetVrfCount() int64`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *NiatelemetryTenant) GetVrfCountOk() (*int64, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *NiatelemetryTenant) SetVrfCount(v int64)`

SetVrfCount sets VrfCount field to given value.

### HasVrfCount

`func (o *NiatelemetryTenant) HasVrfCount() bool`

HasVrfCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryTenant) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryTenant) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryTenant) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryTenant) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


