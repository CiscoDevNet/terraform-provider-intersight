# BiosPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.Policy"]
**AcsControlGpu1state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 1 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu2state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 2 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu3state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 3 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu4state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 4 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu5state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 5 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu6state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 6 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu7state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 7 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlGpu8state** | Pointer to **string** | BIOS Token for setting ACS Control GPU 8 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlSlot11state** | Pointer to **string** | BIOS Token for setting ACS Control Slot 11 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlSlot12state** | Pointer to **string** | BIOS Token for setting ACS Control Slot 12 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlSlot13state** | Pointer to **string** | BIOS Token for setting ACS Control Slot 13 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AcsControlSlot14state** | Pointer to **string** | BIOS Token for setting ACS Control Slot 14 configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AdjacentCacheLinePrefetch** | Pointer to **string** | BIOS Token for setting Adjacent Cache Line Prefetcher configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AdvancedMemTest** | Pointer to **string** | BIOS Token for setting Enhanced Memory Test configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring AdvancedMemTest token. * &#x60;disabled&#x60; - Value - disabled for configuring AdvancedMemTest token. * &#x60;enabled&#x60; - Value - enabled for configuring AdvancedMemTest token. | [optional] [default to "platform-default"]
**AllUsbDevices** | Pointer to **string** | BIOS Token for setting All USB Devices configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Altitude** | Pointer to **string** | BIOS Token for setting Altitude configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;300-m&#x60; - Value - 300-m for configuring Altitude token. * &#x60;900-m&#x60; - Value - 900-m for configuring Altitude token. * &#x60;1500-m&#x60; - Value - 1500-m for configuring Altitude token. * &#x60;3000-m&#x60; - Value - 3000-m for configuring Altitude token. * &#x60;auto&#x60; - Value - auto for configuring Altitude token. | [optional] [default to "platform-default"]
**AspmSupport** | Pointer to **string** | BIOS Token for setting ASPM Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring AspmSupport token. * &#x60;Disabled&#x60; - Value - Disabled for configuring AspmSupport token. * &#x60;Force L0s&#x60; - Value - Force L0s for configuring AspmSupport token. * &#x60;L1 Only&#x60; - Value - L1 Only for configuring AspmSupport token. | [optional] [default to "platform-default"]
**AssertNmiOnPerr** | Pointer to **string** | BIOS Token for setting Assert NMI on PERR configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AssertNmiOnSerr** | Pointer to **string** | BIOS Token for setting Assert NMI on SERR configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AutoCcState** | Pointer to **string** | BIOS Token for setting Autonomous Core C State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**AutonumousCstateEnable** | Pointer to **string** | BIOS Token for setting CPU Autonomous C State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**BaudRate** | Pointer to **string** | BIOS Token for setting Baud Rate configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;9600&#x60; - Value - 9600 for configuring BaudRate token. * &#x60;19200&#x60; - Value - 19200 for configuring BaudRate token. * &#x60;38400&#x60; - Value - 38400 for configuring BaudRate token. * &#x60;57600&#x60; - Value - 57600 for configuring BaudRate token. * &#x60;115200&#x60; - Value - 115200 for configuring BaudRate token. | [optional] [default to "platform-default"]
**BmeDmaMitigation** | Pointer to **string** | BIOS Token for setting BME DMA Mitigation configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**BootOptionNumRetry** | Pointer to **string** | BIOS Token for setting Number of Retries configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;5&#x60; - Value - 5 for configuring BootOptionNumRetry token. * &#x60;13&#x60; - Value - 13 for configuring BootOptionNumRetry token. * &#x60;Infinite&#x60; - Value - Infinite for configuring BootOptionNumRetry token. | [optional] [default to "platform-default"]
**BootOptionReCoolDown** | Pointer to **string** | BIOS Token for setting Cool Down Time  (sec) configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;15&#x60; - Value - 15 for configuring BootOptionReCoolDown token. * &#x60;45&#x60; - Value - 45 for configuring BootOptionReCoolDown token. * &#x60;90&#x60; - Value - 90 for configuring BootOptionReCoolDown token. | [optional] [default to "platform-default"]
**BootOptionRetry** | Pointer to **string** | BIOS Token for setting Boot Option Retry configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**BootPerformanceMode** | Pointer to **string** | BIOS Token for setting Boot Performance Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Max Efficient&#x60; - Value - Max Efficient for configuring BootPerformanceMode token. * &#x60;Max Performance&#x60; - Value - Max Performance for configuring BootPerformanceMode token. * &#x60;Set by Intel NM&#x60; - Value - Set by Intel NM for configuring BootPerformanceMode token. | [optional] [default to "platform-default"]
**BurstAndPostponedRefresh** | Pointer to **string** | BIOS Token for setting Burst and Postponed Refresh configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CbsCmnApbdis** | Pointer to **string** | BIOS Token for setting APBDIS configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;0&#x60; - Value - 0 for configuring CbsCmnApbdis token. * &#x60;1&#x60; - Value - 1 for configuring CbsCmnApbdis token. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnApbdis token. | [optional] [default to "platform-default"]
**CbsCmnCpuCpb** | Pointer to **string** | BIOS Token for setting Core Performance Boost configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuCpb token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuCpb token. | [optional] [default to "platform-default"]
**CbsCmnCpuGenDowncoreCtrl** | Pointer to **string** | BIOS Token for setting Downcore Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;FOUR (2 + 2)&#x60; - Value - FOUR (2 + 2) for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;FOUR (4 + 0)&#x60; - Value - FOUR (4 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;SIX (3 + 3)&#x60; - Value - SIX (3 + 3) for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;THREE (3 + 0)&#x60; - Value - THREE (3 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;TWO (1 + 1)&#x60; - Value - TWO (1 + 1) for configuring CbsCmnCpuGenDowncoreCtrl token. * &#x60;TWO (2 + 0)&#x60; - Value - TWO (2 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token. | [optional] [default to "platform-default"]
**CbsCmnCpuGlobalCstateCtrl** | Pointer to **string** | BIOS Token for setting Global C State Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuGlobalCstateCtrl token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuGlobalCstateCtrl token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnCpuGlobalCstateCtrl token. | [optional] [default to "platform-default"]
**CbsCmnCpuL1streamHwPrefetcher** | Pointer to **string** | BIOS Token for setting L1 Stream HW Prefetcher configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuL1streamHwPrefetcher token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuL1streamHwPrefetcher token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnCpuL1streamHwPrefetcher token. | [optional] [default to "platform-default"]
**CbsCmnCpuL2streamHwPrefetcher** | Pointer to **string** | BIOS Token for setting L2 Stream HW Prefetcher configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuL2streamHwPrefetcher token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuL2streamHwPrefetcher token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnCpuL2streamHwPrefetcher token. | [optional] [default to "platform-default"]
**CbsCmnCpuSmee** | Pointer to **string** | BIOS Token for setting CPU SMEE configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuSmee token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuSmee token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnCpuSmee token. | [optional] [default to "platform-default"]
**CbsCmnCpuStreamingStoresCtrl** | Pointer to **string** | BIOS Token for setting Streaming Stores Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnCpuStreamingStoresCtrl token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnCpuStreamingStoresCtrl token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnCpuStreamingStoresCtrl token. | [optional] [default to "platform-default"]
**CbsCmnDeterminismSlider** | Pointer to **string** | BIOS Token for setting Determinism Slider configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnDeterminismSlider token. * &#x60;Performance&#x60; - Value - Performance for configuring CbsCmnDeterminismSlider token. * &#x60;Power&#x60; - Value - Power for configuring CbsCmnDeterminismSlider token. | [optional] [default to "platform-default"]
**CbsCmnEfficiencyModeEn** | Pointer to **string** | BIOS Token for setting Efficiency Mode Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnEfficiencyModeEn token. * &#x60;Enabled&#x60; - Value - Enabled for configuring CbsCmnEfficiencyModeEn token. | [optional] [default to "platform-default"]
**CbsCmnFixedSocPstate** | Pointer to **string** | BIOS Token for setting Fixed SOC P-State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnFixedSocPstate token. * &#x60;P0&#x60; - Value - P0 for configuring CbsCmnFixedSocPstate token. * &#x60;P1&#x60; - Value - P1 for configuring CbsCmnFixedSocPstate token. * &#x60;P2&#x60; - Value - P2 for configuring CbsCmnFixedSocPstate token. * &#x60;P3&#x60; - Value - P3 for configuring CbsCmnFixedSocPstate token. | [optional] [default to "platform-default"]
**CbsCmnGnbNbIommu** | Pointer to **string** | BIOS Token for setting IOMMU configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnGnbNbIommu token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnGnbNbIommu token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnGnbNbIommu token. | [optional] [default to "platform-default"]
**CbsCmnGnbSmuDfCstates** | Pointer to **string** | BIOS Token for setting DF C-States configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnGnbSmuDfCstates token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnGnbSmuDfCstates token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnGnbSmuDfCstates token. | [optional] [default to "platform-default"]
**CbsCmnGnbSmucppc** | Pointer to **string** | BIOS Token for setting CPPC configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnGnbSmucppc token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnGnbSmucppc token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnGnbSmucppc token. | [optional] [default to "platform-default"]
**CbsCmnMemCtrlBankGroupSwapDdr4** | Pointer to **string** | BIOS Token for setting Bank Group Swap configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token. | [optional] [default to "platform-default"]
**CbsCmnMemMapBankInterleaveDdr4** | Pointer to **string** | BIOS Token for setting Chipset Interleave configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmnMemMapBankInterleaveDdr4 token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCmnMemMapBankInterleaveDdr4 token. | [optional] [default to "platform-default"]
**CbsCmncTdpCtl** | Pointer to **string** | BIOS Token for setting cTDP Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCmncTdpCtl token. * &#x60;Manual&#x60; - Value - Manual for configuring CbsCmncTdpCtl token. | [optional] [default to "platform-default"]
**CbsCpuCcdCtrlSsp** | Pointer to **string** | BIOS Token for setting CCD Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;2 CCDs&#x60; - Value - 2 CCDs for configuring CbsCpuCcdCtrlSsp token. * &#x60;3 CCDs&#x60; - Value - 3 CCDs for configuring CbsCpuCcdCtrlSsp token. * &#x60;4 CCDs&#x60; - Value - 4 CCDs for configuring CbsCpuCcdCtrlSsp token. * &#x60;6 CCDs&#x60; - Value - 6 CCDs for configuring CbsCpuCcdCtrlSsp token. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCpuCcdCtrlSsp token. | [optional] [default to "platform-default"]
**CbsCpuCoreCtrl** | Pointer to **string** | BIOS Token for setting CPU Downcore control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCpuCoreCtrl token. * &#x60;FIVE (5 + 0)&#x60; - Value - FIVE (5 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;FOUR (4 + 0)&#x60; - Value - FOUR (4 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;ONE (1 + 0)&#x60; - Value - ONE (1 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;SEVEN (7 + 0)&#x60; - Value - SEVEN (7 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;SIX (6 + 0)&#x60; - Value - SIX (6 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;THREE (3 + 0)&#x60; - Value - THREE (3 + 0) for configuring CbsCpuCoreCtrl token. * &#x60;TWO (2 + 0)&#x60; - Value - TWO (2 + 0) for configuring CbsCpuCoreCtrl token. | [optional] [default to "platform-default"]
**CbsCpuSmtCtrl** | Pointer to **string** | BIOS Token for setting CPU SMT Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsCpuSmtCtrl token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsCpuSmtCtrl token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsCpuSmtCtrl token. | [optional] [default to "platform-default"]
**CbsDbgCpuSnpMemCover** | Pointer to **string** | BIOS Token for setting SNP Memory Coverage configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsDbgCpuSnpMemCover token. * &#x60;Custom&#x60; - Value - Custom for configuring CbsDbgCpuSnpMemCover token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsDbgCpuSnpMemCover token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsDbgCpuSnpMemCover token. | [optional] [default to "platform-default"]
**CbsDbgCpuSnpMemSizeCover** | Pointer to **string** | BIOS Token for setting SNP Memory Size to Cover in MiB configuration (0 - 1048576 MiB). | [optional] [default to "platform-default"]
**CbsDfCmnAcpiSratL3numa** | Pointer to **string** | BIOS Token for setting ACPI SRAT L3 Cache As NUMA Domain configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsDfCmnAcpiSratL3numa token. * &#x60;disabled&#x60; - Value - disabled for configuring CbsDfCmnAcpiSratL3numa token. * &#x60;enabled&#x60; - Value - enabled for configuring CbsDfCmnAcpiSratL3numa token. | [optional] [default to "platform-default"]
**CbsDfCmnDramNps** | Pointer to **string** | BIOS Token for setting NUMA Nodes per Socket configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsDfCmnDramNps token. * &#x60;NPS0&#x60; - Value - NPS0 for configuring CbsDfCmnDramNps token. * &#x60;NPS1&#x60; - Value - NPS1 for configuring CbsDfCmnDramNps token. * &#x60;NPS2&#x60; - Value - NPS2 for configuring CbsDfCmnDramNps token. * &#x60;NPS4&#x60; - Value - NPS4 for configuring CbsDfCmnDramNps token. | [optional] [default to "platform-default"]
**CbsDfCmnMemIntlv** | Pointer to **string** | BIOS Token for setting AMD Memory Interleaving configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CbsDfCmnMemIntlv token. * &#x60;Channel&#x60; - Value - Channel for configuring CbsDfCmnMemIntlv token. * &#x60;Die&#x60; - Value - Die for configuring CbsDfCmnMemIntlv token. * &#x60;None&#x60; - Value - None for configuring CbsDfCmnMemIntlv token. * &#x60;Socket&#x60; - Value - Socket for configuring CbsDfCmnMemIntlv token. | [optional] [default to "platform-default"]
**CbsDfCmnMemIntlvSize** | Pointer to **string** | BIOS Token for setting AMD Memory Interleaving Size configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;256 Bytes&#x60; - Value - 256 Bytes for configuring CbsDfCmnMemIntlvSize token. * &#x60;512 Bytes&#x60; - Value - 512 Bytes for configuring CbsDfCmnMemIntlvSize token. * &#x60;1 KB&#x60; - Value - 1 KiB for configuring CbsDfCmnMemIntlvSize token. * &#x60;2 KB&#x60; - Value - 2 KiB for configuring CbsDfCmnMemIntlvSize token. * &#x60;4 KB&#x60; - Value - 4 KiB for configuring CbsDfCmnMemIntlvSize token. * &#x60;Auto&#x60; - Value - Auto for configuring CbsDfCmnMemIntlvSize token. | [optional] [default to "platform-default"]
**CbsSevSnpSupport** | Pointer to **string** | BIOS Token for setting SEV-SNP Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CdnEnable** | Pointer to **string** | BIOS Token for setting Consistent Device Naming configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CdnSupport** | Pointer to **string** | BIOS Token for setting CDN Support for LOM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring CdnSupport token. * &#x60;enabled&#x60; - Value - enabled for configuring CdnSupport token. * &#x60;LOMs Only&#x60; - Value - LOMs Only for configuring CdnSupport token. | [optional] [default to "platform-default"]
**ChannelInterLeave** | Pointer to **string** | BIOS Token for setting Channel Interleaving configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1-way&#x60; - Value - 1-way for configuring ChannelInterLeave token. * &#x60;2-way&#x60; - Value - 2-way for configuring ChannelInterLeave token. * &#x60;3-way&#x60; - Value - 3-way for configuring ChannelInterLeave token. * &#x60;4-way&#x60; - Value - 4-way for configuring ChannelInterLeave token. * &#x60;auto&#x60; - Value - auto for configuring ChannelInterLeave token. | [optional] [default to "platform-default"]
**CiscoAdaptiveMemTraining** | Pointer to **string** | BIOS Token for setting Adaptive Memory Training configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CiscoDebugLevel** | Pointer to **string** | BIOS Token for setting BIOS Techlog Level configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Maximum&#x60; - Value - Maximum for configuring CiscoDebugLevel token. * &#x60;Minimum&#x60; - Value - Minimum for configuring CiscoDebugLevel token. * &#x60;Normal&#x60; - Value - Normal for configuring CiscoDebugLevel token. | [optional] [default to "platform-default"]
**CiscoOpromLaunchOptimization** | Pointer to **string** | BIOS Token for setting OptionROM Launch Optimization configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CiscoXgmiMaxSpeed** | Pointer to **string** | BIOS Token for setting Cisco xGMI Max Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CkeLowPolicy** | Pointer to **string** | BIOS Token for setting CKE Low Policy configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;auto&#x60; - Value - auto for configuring CkeLowPolicy token. * &#x60;disabled&#x60; - Value - disabled for configuring CkeLowPolicy token. * &#x60;fast&#x60; - Value - fast for configuring CkeLowPolicy token. * &#x60;slow&#x60; - Value - slow for configuring CkeLowPolicy token. | [optional] [default to "platform-default"]
**ClosedLoopThermThrotl** | Pointer to **string** | BIOS Token for setting Closed Loop Thermal Throttling configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CmciEnable** | Pointer to **string** | BIOS Token for setting Processor CMCI configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ConfigTdp** | Pointer to **string** | BIOS Token for setting Config TDP configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ConfigTdpLevel** | Pointer to **string** | BIOS Token for setting Configurable TDP Level configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Level 1&#x60; - Value - Level 1 for configuring ConfigTdpLevel token. * &#x60;Level 2&#x60; - Value - Level 2 for configuring ConfigTdpLevel token. * &#x60;Normal&#x60; - Value - Normal for configuring ConfigTdpLevel token. | [optional] [default to "platform-default"]
**ConsoleRedirection** | Pointer to **string** | BIOS Token for setting Console Redirection configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;com-0&#x60; - Value - com-0 for configuring ConsoleRedirection token. * &#x60;com-1&#x60; - Value - com-1 for configuring ConsoleRedirection token. * &#x60;disabled&#x60; - Value - disabled for configuring ConsoleRedirection token. * &#x60;enabled&#x60; - Value - enabled for configuring ConsoleRedirection token. * &#x60;serial-port-a&#x60; - Value - serial-port-a for configuring ConsoleRedirection token. | [optional] [default to "platform-default"]
**CoreMultiProcessing** | Pointer to **string** | BIOS Token for setting Core Multi Processing configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1&#x60; - Value - 1 for configuring CoreMultiProcessing token. * &#x60;2&#x60; - Value - 2 for configuring CoreMultiProcessing token. * &#x60;3&#x60; - Value - 3 for configuring CoreMultiProcessing token. * &#x60;4&#x60; - Value - 4 for configuring CoreMultiProcessing token. * &#x60;5&#x60; - Value - 5 for configuring CoreMultiProcessing token. * &#x60;6&#x60; - Value - 6 for configuring CoreMultiProcessing token. * &#x60;7&#x60; - Value - 7 for configuring CoreMultiProcessing token. * &#x60;8&#x60; - Value - 8 for configuring CoreMultiProcessing token. * &#x60;9&#x60; - Value - 9 for configuring CoreMultiProcessing token. * &#x60;10&#x60; - Value - 10 for configuring CoreMultiProcessing token. * &#x60;11&#x60; - Value - 11 for configuring CoreMultiProcessing token. * &#x60;12&#x60; - Value - 12 for configuring CoreMultiProcessing token. * &#x60;13&#x60; - Value - 13 for configuring CoreMultiProcessing token. * &#x60;14&#x60; - Value - 14 for configuring CoreMultiProcessing token. * &#x60;15&#x60; - Value - 15 for configuring CoreMultiProcessing token. * &#x60;16&#x60; - Value - 16 for configuring CoreMultiProcessing token. * &#x60;17&#x60; - Value - 17 for configuring CoreMultiProcessing token. * &#x60;18&#x60; - Value - 18 for configuring CoreMultiProcessing token. * &#x60;19&#x60; - Value - 19 for configuring CoreMultiProcessing token. * &#x60;20&#x60; - Value - 20 for configuring CoreMultiProcessing token. * &#x60;21&#x60; - Value - 21 for configuring CoreMultiProcessing token. * &#x60;22&#x60; - Value - 22 for configuring CoreMultiProcessing token. * &#x60;23&#x60; - Value - 23 for configuring CoreMultiProcessing token. * &#x60;24&#x60; - Value - 24 for configuring CoreMultiProcessing token. * &#x60;25&#x60; - Value - 25 for configuring CoreMultiProcessing token. * &#x60;26&#x60; - Value - 26 for configuring CoreMultiProcessing token. * &#x60;27&#x60; - Value - 27 for configuring CoreMultiProcessing token. * &#x60;28&#x60; - Value - 28 for configuring CoreMultiProcessing token. * &#x60;29&#x60; - Value - 29 for configuring CoreMultiProcessing token. * &#x60;30&#x60; - Value - 30 for configuring CoreMultiProcessing token. * &#x60;31&#x60; - Value - 31 for configuring CoreMultiProcessing token. * &#x60;32&#x60; - Value - 32 for configuring CoreMultiProcessing token. * &#x60;33&#x60; - Value - 33 for configuring CoreMultiProcessing token. * &#x60;34&#x60; - Value - 34 for configuring CoreMultiProcessing token. * &#x60;35&#x60; - Value - 35 for configuring CoreMultiProcessing token. * &#x60;36&#x60; - Value - 36 for configuring CoreMultiProcessing token. * &#x60;37&#x60; - Value - 37 for configuring CoreMultiProcessing token. * &#x60;38&#x60; - Value - 38 for configuring CoreMultiProcessing token. * &#x60;39&#x60; - Value - 39 for configuring CoreMultiProcessing token. * &#x60;40&#x60; - Value - 40 for configuring CoreMultiProcessing token. * &#x60;41&#x60; - Value - 41 for configuring CoreMultiProcessing token. * &#x60;42&#x60; - Value - 42 for configuring CoreMultiProcessing token. * &#x60;43&#x60; - Value - 43 for configuring CoreMultiProcessing token. * &#x60;44&#x60; - Value - 44 for configuring CoreMultiProcessing token. * &#x60;45&#x60; - Value - 45 for configuring CoreMultiProcessing token. * &#x60;46&#x60; - Value - 46 for configuring CoreMultiProcessing token. * &#x60;47&#x60; - Value - 47 for configuring CoreMultiProcessing token. * &#x60;48&#x60; - Value - 48 for configuring CoreMultiProcessing token. * &#x60;all&#x60; - Value - all for configuring CoreMultiProcessing token. | [optional] [default to "platform-default"]
**CpuEnergyPerformance** | Pointer to **string** | BIOS Token for setting Energy Performance configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;balanced-energy&#x60; - Value - balanced-energy for configuring CpuEnergyPerformance token. * &#x60;balanced-performance&#x60; - Value - balanced-performance for configuring CpuEnergyPerformance token. * &#x60;balanced-power&#x60; - Value - balanced-power for configuring CpuEnergyPerformance token. * &#x60;energy-efficient&#x60; - Value - energy-efficient for configuring CpuEnergyPerformance token. * &#x60;performance&#x60; - Value - performance for configuring CpuEnergyPerformance token. * &#x60;power&#x60; - Value - power for configuring CpuEnergyPerformance token. | [optional] [default to "platform-default"]
**CpuFrequencyFloor** | Pointer to **string** | BIOS Token for setting Frequency Floor Override configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**CpuPerformance** | Pointer to **string** | BIOS Token for setting CPU Performance configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;custom&#x60; - Value - custom for configuring CpuPerformance token. * &#x60;enterprise&#x60; - Value - enterprise for configuring CpuPerformance token. * &#x60;high-throughput&#x60; - Value - high-throughput for configuring CpuPerformance token. * &#x60;hpc&#x60; - Value - hpc for configuring CpuPerformance token. | [optional] [default to "platform-default"]
**CpuPowerManagement** | Pointer to **string** | BIOS Token for setting Power Technology configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;custom&#x60; - Value - custom for configuring CpuPowerManagement token. * &#x60;disabled&#x60; - Value - disabled for configuring CpuPowerManagement token. * &#x60;energy-efficient&#x60; - Value - energy-efficient for configuring CpuPowerManagement token. * &#x60;performance&#x60; - Value - performance for configuring CpuPowerManagement token. | [optional] [default to "platform-default"]
**CrQos** | Pointer to **string** | BIOS Token for setting CR QoS configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Disabled&#x60; - Value - Disabled for configuring CrQos token. * &#x60;Mode 0 - Disable the PMem QoS Feature&#x60; - Value - Mode 0 - Disable the PMem QoS Feature for configuring CrQos token. * &#x60;Mode 1 - M2M QoS Enable and CHA QoS Disable&#x60; - Value - Mode 1 - M2M QoS Enable and CHA QoS Disable for configuring CrQos token. * &#x60;Mode 2 - M2M QoS Enable and CHA QoS Enable&#x60; - Value - Mode 2 - M2M QoS Enable and CHA QoS Enable for configuring CrQos token. * &#x60;Recipe 1&#x60; - Value - Recipe 1 for configuring CrQos token. * &#x60;Recipe 2&#x60; - Value - Recipe 2 for configuring CrQos token. * &#x60;Recipe 3&#x60; - Value - Recipe 3 for configuring CrQos token. | [optional] [default to "platform-default"]
**CrfastgoConfig** | Pointer to **string** | BIOS Token for setting CR FastGo Config configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring CrfastgoConfig token. * &#x60;Default&#x60; - Value - Default for configuring CrfastgoConfig token. * &#x60;Disable optimization&#x60; - Value - Disable optimization for configuring CrfastgoConfig token. * &#x60;Enable optimization&#x60; - Value - Enable optimization for configuring CrfastgoConfig token. * &#x60;Option 1&#x60; - Value - Option 1 for configuring CrfastgoConfig token. * &#x60;Option 2&#x60; - Value - Option 2 for configuring CrfastgoConfig token. * &#x60;Option 3&#x60; - Value - Option 3 for configuring CrfastgoConfig token. * &#x60;Option 4&#x60; - Value - Option 4 for configuring CrfastgoConfig token. * &#x60;Option 5&#x60; - Value - Option 5 for configuring CrfastgoConfig token. | [optional] [default to "platform-default"]
**DcpmmFirmwareDowngrade** | Pointer to **string** | BIOS Token for setting DCPMM Firmware Downgrade configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**DemandScrub** | Pointer to **string** | BIOS Token for setting Demand Scrub configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**DirectCacheAccess** | Pointer to **string** | BIOS Token for setting Direct Cache Access Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;auto&#x60; - Value - auto for configuring DirectCacheAccess token. * &#x60;disabled&#x60; - Value - disabled for configuring DirectCacheAccess token. * &#x60;enabled&#x60; - Value - enabled for configuring DirectCacheAccess token. | [optional] [default to "platform-default"]
**DramClockThrottling** | Pointer to **string** | BIOS Token for setting DRAM Clock Throttling configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring DramClockThrottling token. * &#x60;Balanced&#x60; - Value - Balanced for configuring DramClockThrottling token. * &#x60;Energy Efficient&#x60; - Value - Energy Efficient for configuring DramClockThrottling token. * &#x60;Performance&#x60; - Value - Performance for configuring DramClockThrottling token. | [optional] [default to "platform-default"]
**DramRefreshRate** | Pointer to **string** | BIOS Token for setting DRAM Refresh Rate configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1x&#x60; - Value - 1x for configuring DramRefreshRate token. * &#x60;2x&#x60; - Value - 2x for configuring DramRefreshRate token. * &#x60;3x&#x60; - Value - 3x for configuring DramRefreshRate token. * &#x60;4x&#x60; - Value - 4x for configuring DramRefreshRate token. * &#x60;Auto&#x60; - Value - Auto for configuring DramRefreshRate token. | [optional] [default to "platform-default"]
**DramSwThermalThrottling** | Pointer to **string** | BIOS Token for setting DRAM SW Thermal Throttling configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EadrSupport** | Pointer to **string** | BIOS Token for setting eADR Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring EadrSupport token. * &#x60;disabled&#x60; - Value - disabled for configuring EadrSupport token. * &#x60;enabled&#x60; - Value - enabled for configuring EadrSupport token. | [optional] [default to "platform-default"]
**EdpcEn** | Pointer to **string** | BIOS Token for setting IIO eDPC Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Disabled&#x60; - Value - Disabled for configuring EdpcEn token. * &#x60;On Fatal Error&#x60; - Value - On Fatal Error for configuring EdpcEn token. * &#x60;On Fatal and Non-Fatal Errors&#x60; - Value - On Fatal and Non-Fatal Errors for configuring EdpcEn token. | [optional] [default to "platform-default"]
**EnableClockSpreadSpec** | Pointer to **string** | BIOS Token for setting External SSC Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EnableMktme** | Pointer to **string** | BIOS Token for setting Multikey Total Memory Encryption  (MK-TME) configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EnableSgx** | Pointer to **string** | BIOS Token for setting Software Guard Extensions  (SGX) configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EnableTme** | Pointer to **string** | BIOS Token for setting Total Memory Encryption  (TME) configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EnergyEfficientTurbo** | Pointer to **string** | BIOS Token for setting Energy Efficient Turbo configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EngPerfTuning** | Pointer to **string** | BIOS Token for setting Energy Performance Tuning configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;BIOS&#x60; - Value - BIOS for configuring EngPerfTuning token. * &#x60;OS&#x60; - Value - OS for configuring EngPerfTuning token. | [optional] [default to "platform-default"]
**EnhancedIntelSpeedStepTech** | Pointer to **string** | BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EpochUpdate** | Pointer to **string** | BIOS Token for setting Select Owner EPOCH Input Type configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Change to New Random Owner EPOCHs&#x60; - Value - Change to New Random Owner EPOCHs for configuring EpochUpdate token. * &#x60;Manual User Defined Owner EPOCHs&#x60; - Value - Manual User Defined Owner EPOCHs for configuring EpochUpdate token. * &#x60;SGX Owner EPOCH activated&#x60; - Value - SGX Owner EPOCH activated for configuring EpochUpdate token. | [optional] [default to "platform-default"]
**EppEnable** | Pointer to **string** | BIOS Token for setting Processor EPP Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**EppProfile** | Pointer to **string** | BIOS Token for setting EPP Profile configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Balanced Performance&#x60; - Value - Balanced Performance for configuring EppProfile token. * &#x60;Balanced Power&#x60; - Value - Balanced Power for configuring EppProfile token. * &#x60;Performance&#x60; - Value - Performance for configuring EppProfile token. * &#x60;Power&#x60; - Value - Power for configuring EppProfile token. | [optional] [default to "platform-default"]
**ExecuteDisableBit** | Pointer to **string** | BIOS Token for setting Execute Disable Bit configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ExtendedApic** | Pointer to **string** | BIOS Token for setting Local X2 Apic configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring ExtendedApic token. * &#x60;enabled&#x60; - Value - enabled for configuring ExtendedApic token. * &#x60;X2APIC&#x60; - Value - X2APIC for configuring ExtendedApic token. * &#x60;XAPIC&#x60; - Value - XAPIC for configuring ExtendedApic token. | [optional] [default to "platform-default"]
**FlowControl** | Pointer to **string** | BIOS Token for setting Flow Control configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;none&#x60; - Value - none for configuring FlowControl token. * &#x60;rts-cts&#x60; - Value - rts-cts for configuring FlowControl token. | [optional] [default to "platform-default"]
**Frb2enable** | Pointer to **string** | BIOS Token for setting FRB-2 Timer configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**HardwarePrefetch** | Pointer to **string** | BIOS Token for setting Hardware Prefetcher configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**HwpmEnable** | Pointer to **string** | BIOS Token for setting CPU Hardware Power Management configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Disabled&#x60; - Value - Disabled for configuring HwpmEnable token. * &#x60;HWPM Native Mode&#x60; - Value - HWPM Native Mode for configuring HwpmEnable token. * &#x60;HWPM OOB Mode&#x60; - Value - HWPM OOB Mode for configuring HwpmEnable token. * &#x60;NATIVE MODE&#x60; - Value - NATIVE MODE for configuring HwpmEnable token. * &#x60;Native Mode with no Legacy&#x60; - Value - Native Mode with no Legacy for configuring HwpmEnable token. * &#x60;OOB MODE&#x60; - Value - OOB MODE for configuring HwpmEnable token. | [optional] [default to "platform-default"]
**ImcInterleave** | Pointer to **string** | BIOS Token for setting IMC Interleaving configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1-way Interleave&#x60; - Value - 1-way Interleave for configuring ImcInterleave token. * &#x60;2-way Interleave&#x60; - Value - 2-way Interleave for configuring ImcInterleave token. * &#x60;Auto&#x60; - Value - Auto for configuring ImcInterleave token. | [optional] [default to "platform-default"]
**IntelDynamicSpeedSelect** | Pointer to **string** | BIOS Token for setting Intel Dynamic Speed Select configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelHyperThreadingTech** | Pointer to **string** | BIOS Token for setting Intel HyperThreading Tech configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelSpeedSelect** | Pointer to **string** | BIOS Token for setting Intel Speed Select configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Base&#x60; - Value - Base for configuring IntelSpeedSelect token. * &#x60;Config 1&#x60; - Value - Config 1 for configuring IntelSpeedSelect token. * &#x60;Config 2&#x60; - Value - Config 2 for configuring IntelSpeedSelect token. * &#x60;Config 3&#x60; - Value - Config 3 for configuring IntelSpeedSelect token. * &#x60;Config 4&#x60; - Value - Config 4 for configuring IntelSpeedSelect token. | [optional] [default to "platform-default"]
**IntelTurboBoostTech** | Pointer to **string** | BIOS Token for setting Intel Turbo Boost Tech configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVirtualizationTechnology** | Pointer to **string** | BIOS Token for setting Intel (R) VT configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVtForDirectedIo** | Pointer to **string** | BIOS Token for setting Intel VT for Directed IO configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVtdCoherencySupport** | Pointer to **string** | BIOS Token for setting Intel (R) VT-d Coherency Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVtdInterruptRemapping** | Pointer to **string** | BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVtdPassThroughDmaSupport** | Pointer to **string** | BIOS Token for setting Intel (R) VT-d PassThrough DMA Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IntelVtdatsSupport** | Pointer to **string** | BIOS Token for setting Intel VTD ATS Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**IohErrorEnable** | Pointer to **string** | BIOS Token for setting IIO Error Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;No&#x60; - Value - No for configuring IohErrorEnable token. * &#x60;Yes&#x60; - Value - Yes for configuring IohErrorEnable token. | [optional] [default to "platform-default"]
**IohResource** | Pointer to **string** | BIOS Token for setting IOH Resource Allocation configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;IOH0 24k IOH1 40k&#x60; - Value - IOH0 24k IOH1 40k for configuring IohResource token. * &#x60;IOH0 32k IOH1 32k&#x60; - Value - IOH0 32k IOH1 32k for configuring IohResource token. * &#x60;IOH0 40k IOH1 24k&#x60; - Value - IOH0 40k IOH1 24k for configuring IohResource token. * &#x60;IOH0 48k IOH1 16k&#x60; - Value - IOH0 48k IOH1 16k for configuring IohResource token. * &#x60;IOH0 56k IOH1 8k&#x60; - Value - IOH0 56k IOH1 8k for configuring IohResource token. | [optional] [default to "platform-default"]
**IpPrefetch** | Pointer to **string** | BIOS Token for setting DCU IP Prefetcher configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Ipv4http** | Pointer to **string** | BIOS Token for setting IPV4 HTTP Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Ipv4pxe** | Pointer to **string** | BIOS Token for setting IPv4 PXE Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Ipv6http** | Pointer to **string** | BIOS Token for setting IPV6 HTTP Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Ipv6pxe** | Pointer to **string** | BIOS Token for setting IPV6 PXE Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**KtiPrefetch** | Pointer to **string** | BIOS Token for setting KTI Prefetch configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring KtiPrefetch token. * &#x60;disabled&#x60; - Value - disabled for configuring KtiPrefetch token. * &#x60;enabled&#x60; - Value - enabled for configuring KtiPrefetch token. | [optional] [default to "platform-default"]
**LegacyOsRedirection** | Pointer to **string** | BIOS Token for setting Legacy OS Redirection configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**LegacyUsbSupport** | Pointer to **string** | BIOS Token for setting Legacy USB Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;auto&#x60; - Value - auto for configuring LegacyUsbSupport token. * &#x60;disabled&#x60; - Value - disabled for configuring LegacyUsbSupport token. * &#x60;enabled&#x60; - Value - enabled for configuring LegacyUsbSupport token. | [optional] [default to "platform-default"]
**LlcPrefetch** | Pointer to **string** | BIOS Token for setting LLC Prefetch configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**LomPort0state** | Pointer to **string** | BIOS Token for setting LOM Port 0 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring LomPort0state token. * &#x60;enabled&#x60; - Value - enabled for configuring LomPort0state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring LomPort0state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring LomPort0state token. | [optional] [default to "platform-default"]
**LomPort1state** | Pointer to **string** | BIOS Token for setting LOM Port 1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring LomPort1state token. * &#x60;enabled&#x60; - Value - enabled for configuring LomPort1state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring LomPort1state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring LomPort1state token. | [optional] [default to "platform-default"]
**LomPort2state** | Pointer to **string** | BIOS Token for setting LOM Port 2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring LomPort2state token. * &#x60;enabled&#x60; - Value - enabled for configuring LomPort2state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring LomPort2state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring LomPort2state token. | [optional] [default to "platform-default"]
**LomPort3state** | Pointer to **string** | BIOS Token for setting LOM Port 3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring LomPort3state token. * &#x60;enabled&#x60; - Value - enabled for configuring LomPort3state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring LomPort3state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring LomPort3state token. | [optional] [default to "platform-default"]
**LomPortsAllState** | Pointer to **string** | BIOS Token for setting All Onboard LOM Ports configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**LvDdrMode** | Pointer to **string** | BIOS Token for setting Low Voltage DDR Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;auto&#x60; - Value - auto for configuring LvDdrMode token. * &#x60;performance-mode&#x60; - Value - performance-mode for configuring LvDdrMode token. * &#x60;power-saving-mode&#x60; - Value - power-saving-mode for configuring LvDdrMode token. | [optional] [default to "platform-default"]
**MakeDeviceNonBootable** | Pointer to **string** | BIOS Token for setting Make Device Non Bootable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**MemoryBandwidthBoost** | Pointer to **string** | BIOS Token for setting Memory Bandwidth Boost configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**MemoryInterLeave** | Pointer to **string** | BIOS Token for setting Intel Memory Interleaving configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1 Way Node Interleave&#x60; - Value - 1 Way Node Interleave for configuring MemoryInterLeave token. * &#x60;2 Way Node Interleave&#x60; - Value - 2 Way Node Interleave for configuring MemoryInterLeave token. * &#x60;4 Way Node Interleave&#x60; - Value - 4 Way Node Interleave for configuring MemoryInterLeave token. * &#x60;8 Way Node Interleave&#x60; - Value - 8 Way Node Interleave for configuring MemoryInterLeave token. * &#x60;disabled&#x60; - Value - disabled for configuring MemoryInterLeave token. * &#x60;enabled&#x60; - Value - enabled for configuring MemoryInterLeave token. | [optional] [default to "platform-default"]
**MemoryMappedIoAbove4gb** | Pointer to **string** | BIOS Token for setting Memory Mapped IO above 4GiB configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**MemoryRefreshRate** | Pointer to **string** | BIOS Token for setting Memory Refresh Rate configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1x Refresh&#x60; - Value - 1x Refresh for configuring MemoryRefreshRate token. * &#x60;2x Refresh&#x60; - Value - 2x Refresh for configuring MemoryRefreshRate token. | [optional] [default to "platform-default"]
**MemorySizeLimit** | Pointer to **string** | BIOS Token for setting Memory Size Limit in GiB configuration (0 - 65535 GiB). | [optional] [default to "platform-default"]
**MemoryThermalThrottling** | Pointer to **string** | BIOS Token for setting Memory Thermal Throttling Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;CLTT with PECI&#x60; - Value - CLTT with PECI for configuring MemoryThermalThrottling token. * &#x60;Disabled&#x60; - Value - Disabled for configuring MemoryThermalThrottling token. | [optional] [default to "platform-default"]
**MirroringMode** | Pointer to **string** | BIOS Token for setting Mirroring Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;inter-socket&#x60; - Value - inter-socket for configuring MirroringMode token. * &#x60;intra-socket&#x60; - Value - intra-socket for configuring MirroringMode token. | [optional] [default to "platform-default"]
**MmcfgBase** | Pointer to **string** | BIOS Token for setting MMCFG BASE configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1 GB&#x60; - Value - 1 GiB for configuring MmcfgBase token. * &#x60;2 GB&#x60; - Value - 2 GiB for configuring MmcfgBase token. * &#x60;2.5 GB&#x60; - Value - 2.5 GiB for configuring MmcfgBase token. * &#x60;3 GB&#x60; - Value - 3 GiB for configuring MmcfgBase token. * &#x60;Auto&#x60; - Value - Auto for configuring MmcfgBase token. | [optional] [default to "platform-default"]
**NetworkStack** | Pointer to **string** | BIOS Token for setting Network Stack configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**NumaOptimized** | Pointer to **string** | BIOS Token for setting NUMA Optimized configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**NvmdimmPerformConfig** | Pointer to **string** | BIOS Token for setting NVM Performance Setting configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;BW Optimized&#x60; - Value - BW Optimized for configuring NvmdimmPerformConfig token. * &#x60;Balanced Profile&#x60; - Value - Balanced Profile for configuring NvmdimmPerformConfig token. * &#x60;Latency Optimized&#x60; - Value - Latency Optimized for configuring NvmdimmPerformConfig token. | [optional] [default to "platform-default"]
**Onboard10gbitLom** | Pointer to **string** | BIOS Token for setting Onboard 10Gbit LOM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**OnboardGbitLom** | Pointer to **string** | BIOS Token for setting Onboard Gbit LOM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**OnboardScuStorageSupport** | Pointer to **string** | BIOS Token for setting Onboard SCU Storage Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**OnboardScuStorageSwStack** | Pointer to **string** | BIOS Token for setting Onboard SCU Storage SW Stack configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Intel RSTe&#x60; - Value - Intel RSTe for configuring OnboardScuStorageSwStack token. * &#x60;LSI SW RAID&#x60; - Value - LSI SW RAID for configuring OnboardScuStorageSwStack token. | [optional] [default to "platform-default"]
**OperationMode** | Pointer to **string** | BIOS Token for setting Operation Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Test Only&#x60; - Value - Test Only for configuring OperationMode token. * &#x60;Test and Repair&#x60; - Value - Test and Repair for configuring OperationMode token. | [optional] [default to "platform-default"]
**OsBootWatchdogTimer** | Pointer to **string** | BIOS Token for setting OS Boot Watchdog Timer configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**OsBootWatchdogTimerPolicy** | Pointer to **string** | BIOS Token for setting OS Boot Watchdog Timer Policy configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;do-nothing&#x60; - Value - do-nothing for configuring OsBootWatchdogTimerPolicy token. * &#x60;power-off&#x60; - Value - power-off for configuring OsBootWatchdogTimerPolicy token. * &#x60;reset&#x60; - Value - reset for configuring OsBootWatchdogTimerPolicy token. | [optional] [default to "platform-default"]
**OsBootWatchdogTimerTimeout** | Pointer to **string** | BIOS Token for setting OS Boot Watchdog Timer Timeout configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;5-minutes&#x60; - Value - 5-minutes for configuring OsBootWatchdogTimerTimeout token. * &#x60;10-minutes&#x60; - Value - 10-minutes for configuring OsBootWatchdogTimerTimeout token. * &#x60;15-minutes&#x60; - Value - 15-minutes for configuring OsBootWatchdogTimerTimeout token. * &#x60;20-minutes&#x60; - Value - 20-minutes for configuring OsBootWatchdogTimerTimeout token. | [optional] [default to "platform-default"]
**OutOfBandMgmtPort** | Pointer to **string** | BIOS Token for setting Out-of-Band Mgmt Port configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PackageCstateLimit** | Pointer to **string** | BIOS Token for setting Package C State Limit configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PackageCstateLimit token. * &#x60;C0 C1 State&#x60; - Value - C0 C1 State for configuring PackageCstateLimit token. * &#x60;C0/C1&#x60; - Value - C0/C1 for configuring PackageCstateLimit token. * &#x60;C2&#x60; - Value - C2 for configuring PackageCstateLimit token. * &#x60;C6 Non Retention&#x60; - Value - C6 Non Retention for configuring PackageCstateLimit token. * &#x60;C6 Retention&#x60; - Value - C6 Retention for configuring PackageCstateLimit token. * &#x60;No Limit&#x60; - Value - No Limit for configuring PackageCstateLimit token. | [optional] [default to "platform-default"]
**PanicHighWatermark** | Pointer to **string** | BIOS Token for setting Panic and High Watermark configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;High&#x60; - Value - High for configuring PanicHighWatermark token. * &#x60;Low&#x60; - Value - Low for configuring PanicHighWatermark token. | [optional] [default to "platform-default"]
**PartialCacheLineSparing** | Pointer to **string** | BIOS Token for setting Partial Cache Line Sparing configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PartialMirrorModeConfig** | Pointer to **string** | BIOS Token for setting Partial Memory Mirror Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring PartialMirrorModeConfig token. * &#x60;Percentage&#x60; - Value - Percentage for configuring PartialMirrorModeConfig token. * &#x60;Value in GB&#x60; - Value - Value in GiB for configuring PartialMirrorModeConfig token. | [optional] [default to "platform-default"]
**PartialMirrorPercent** | Pointer to **string** | BIOS Token for setting Partial Mirror Percentage configuration (0.00 - 50.00 Percentage). | [optional] [default to "platform-default"]
**PartialMirrorValue1** | Pointer to **string** | BIOS Token for setting Partial Mirror1 Size in GiB configuration (0 - 65535 GiB). | [optional] [default to "platform-default"]
**PartialMirrorValue2** | Pointer to **string** | BIOS Token for setting Partial Mirror2 Size in GiB configuration (0 - 65535 GiB). | [optional] [default to "platform-default"]
**PartialMirrorValue3** | Pointer to **string** | BIOS Token for setting Partial Mirror3 Size in GiB configuration (0 - 65535 GiB). | [optional] [default to "platform-default"]
**PartialMirrorValue4** | Pointer to **string** | BIOS Token for setting Partial Mirror4 Size in GiB configuration (0 - 65535 GiB). | [optional] [default to "platform-default"]
**PatrolScrub** | Pointer to **string** | BIOS Token for setting Patrol Scrub configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring PatrolScrub token. * &#x60;Enable at End of POST&#x60; - Value - Enable at End of POST for configuring PatrolScrub token. * &#x60;enabled&#x60; - Value - enabled for configuring PatrolScrub token. | [optional] [default to "platform-default"]
**PatrolScrubDuration** | Pointer to **string** | BIOS Token for setting Patrol Scrub Interval configuration (5 - 23 Hour). | [optional] [default to "platform-default"]
**PcIeRasSupport** | Pointer to **string** | BIOS Token for setting PCIe RAS Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcIeSsdHotPlugSupport** | Pointer to **string** | BIOS Token for setting NVMe SSD Hot-Plug Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PchUsb30mode** | Pointer to **string** | BIOS Token for setting xHCI Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PciOptionRoMs** | Pointer to **string** | BIOS Token for setting All PCIe Slots OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring PciOptionRoMs token. * &#x60;enabled&#x60; - Value - enabled for configuring PciOptionRoMs token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring PciOptionRoMs token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring PciOptionRoMs token. | [optional] [default to "platform-default"]
**PciRomClp** | Pointer to **string** | BIOS Token for setting PCI ROM CLP configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieAriSupport** | Pointer to **string** | BIOS Token for setting PCIe ARI Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieAriSupport token. * &#x60;disabled&#x60; - Value - disabled for configuring PcieAriSupport token. * &#x60;enabled&#x60; - Value - enabled for configuring PcieAriSupport token. | [optional] [default to "platform-default"]
**PciePllSsc** | Pointer to **string** | BIOS Token for setting PCIe PLL SSC configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PciePllSsc token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PciePllSsc token. * &#x60;ZeroPointFive&#x60; - Value - ZeroPointFive for configuring PciePllSsc token. | [optional] [default to "platform-default"]
**PcieSlotMraid1linkSpeed** | Pointer to **string** | BIOS Token for setting MRAID1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotMraid1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotMraid1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotMraid1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotMraid1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotMraid1linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring PcieSlotMraid1linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotMraid1optionRom** | Pointer to **string** | BIOS Token for setting MRAID1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotMraid2linkSpeed** | Pointer to **string** | BIOS Token for setting MRAID2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotMraid2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotMraid2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotMraid2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotMraid2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotMraid2linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring PcieSlotMraid2linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotMraid2optionRom** | Pointer to **string** | BIOS Token for setting MRAID2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotMstorraidLinkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot MSTOR Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotMstorraidLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotMstorraidLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotMstorraidLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotMstorraidLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotMstorraidLinkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring PcieSlotMstorraidLinkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotMstorraidOptionRom** | Pointer to **string** | BIOS Token for setting PCIe Slot MSTOR RAID OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme1linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme1linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme1optionRom** | Pointer to **string** | BIOS Token for setting NVME 1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme2linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme2linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme2optionRom** | Pointer to **string** | BIOS Token for setting NVME 2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme3linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 3 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme3linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme3linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme3linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme3linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme3linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme3optionRom** | Pointer to **string** | BIOS Token for setting NVME 3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme4linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 4 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme4linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme4linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme4linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme4linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme4linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme4optionRom** | Pointer to **string** | BIOS Token for setting NVME 4 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme5linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 5 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme5linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme5linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme5linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme5linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme5linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme5optionRom** | Pointer to **string** | BIOS Token for setting NVME 5 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PcieSlotNvme6linkSpeed** | Pointer to **string** | BIOS Token for setting NVME 6 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring PcieSlotNvme6linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring PcieSlotNvme6linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring PcieSlotNvme6linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring PcieSlotNvme6linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring PcieSlotNvme6linkSpeed token. | [optional] [default to "platform-default"]
**PcieSlotNvme6optionRom** | Pointer to **string** | BIOS Token for setting NVME 6 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PopSupport** | Pointer to **string** | BIOS Token for setting Power ON Password configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PostErrorPause** | Pointer to **string** | BIOS Token for setting POST Error Pause configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**PostPackageRepair** | Pointer to **string** | BIOS Token for setting Post Package Repair configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Disabled&#x60; - Value - Disabled for configuring PostPackageRepair token. * &#x60;Hard PPR&#x60; - Value - Hard PPR for configuring PostPackageRepair token. | [optional] [default to "platform-default"]
**ProcessorC1e** | Pointer to **string** | BIOS Token for setting Processor C1E configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ProcessorC3report** | Pointer to **string** | BIOS Token for setting Processor C3 Report configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ProcessorC6report** | Pointer to **string** | BIOS Token for setting Processor C6 Report configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**ProcessorCstate** | Pointer to **string** | BIOS Token for setting CPU C State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Psata** | Pointer to **string** | BIOS Token for setting P-SATA Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;AHCI&#x60; - Value - AHCI for configuring Psata token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Psata token. * &#x60;LSI SW RAID&#x60; - Value - LSI SW RAID for configuring Psata token. | [optional] [default to "platform-default"]
**PstateCoordType** | Pointer to **string** | BIOS Token for setting P-STATE Coordination configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;HW ALL&#x60; - Value - HW ALL for configuring PstateCoordType token. * &#x60;SW ALL&#x60; - Value - SW ALL for configuring PstateCoordType token. * &#x60;SW ANY&#x60; - Value - SW ANY for configuring PstateCoordType token. | [optional] [default to "platform-default"]
**PuttyKeyPad** | Pointer to **string** | BIOS Token for setting Putty KeyPad configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;ESCN&#x60; - Value - ESCN for configuring PuttyKeyPad token. * &#x60;LINUX&#x60; - Value - LINUX for configuring PuttyKeyPad token. * &#x60;SCO&#x60; - Value - SCO for configuring PuttyKeyPad token. * &#x60;VT100&#x60; - Value - VT100 for configuring PuttyKeyPad token. * &#x60;VT400&#x60; - Value - VT400 for configuring PuttyKeyPad token. * &#x60;XTERMR6&#x60; - Value - XTERMR6 for configuring PuttyKeyPad token. | [optional] [default to "platform-default"]
**PwrPerfTuning** | Pointer to **string** | BIOS Token for setting Power Performance Tuning configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;bios&#x60; - Value - BIOS for configuring PwrPerfTuning token. * &#x60;os&#x60; - Value - os for configuring PwrPerfTuning token. * &#x60;peci&#x60; - Value - peci for configuring PwrPerfTuning token. | [optional] [default to "platform-default"]
**QpiLinkFrequency** | Pointer to **string** | BIOS Token for setting QPI Link Frequency Select configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;6.4-gt/s&#x60; - Value - 6.4-gt/s for configuring QpiLinkFrequency token. * &#x60;7.2-gt/s&#x60; - Value - 7.2-gt/s for configuring QpiLinkFrequency token. * &#x60;8.0-gt/s&#x60; - Value - 8.0-gt/s for configuring QpiLinkFrequency token. * &#x60;9.6-gt/s&#x60; - Value - 9.6-gt/s for configuring QpiLinkFrequency token. * &#x60;auto&#x60; - Value - auto for configuring QpiLinkFrequency token. | [optional] [default to "platform-default"]
**QpiLinkSpeed** | Pointer to **string** | BIOS Token for setting UPI Link Frequency Select configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;10.4GT/s&#x60; - Value - 10.4GT/s for configuring QpiLinkSpeed token. * &#x60;11.2GT/s&#x60; - Value - 11.2GT/s for configuring QpiLinkSpeed token. * &#x60;9.6GT/s&#x60; - Value - 9.6GT/s for configuring QpiLinkSpeed token. * &#x60;Auto&#x60; - Value - Auto for configuring QpiLinkSpeed token. | [optional] [default to "platform-default"]
**QpiSnoopMode** | Pointer to **string** | BIOS Token for setting QPI Snoop Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;auto&#x60; - Value - auto for configuring QpiSnoopMode token. * &#x60;cluster-on-die&#x60; - Value - cluster-on-die for configuring QpiSnoopMode token. * &#x60;early-snoop&#x60; - Value - early-snoop for configuring QpiSnoopMode token. * &#x60;home-directory-snoop&#x60; - Value - home-directory-snoop for configuring QpiSnoopMode token. * &#x60;home-directory-snoop-with-osb&#x60; - Value - home-directory-snoop-with-osb for configuring QpiSnoopMode token. * &#x60;home-snoop&#x60; - Value - home-snoop for configuring QpiSnoopMode token. | [optional] [default to "platform-default"]
**RankInterLeave** | Pointer to **string** | BIOS Token for setting Rank Interleaving configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1-way&#x60; - Value - 1-way for configuring RankInterLeave token. * &#x60;2-way&#x60; - Value - 2-way for configuring RankInterLeave token. * &#x60;4-way&#x60; - Value - 4-way for configuring RankInterLeave token. * &#x60;8-way&#x60; - Value - 8-way for configuring RankInterLeave token. * &#x60;auto&#x60; - Value - auto for configuring RankInterLeave token. | [optional] [default to "platform-default"]
**RedirectionAfterPost** | Pointer to **string** | BIOS Token for setting Redirection After BIOS POST configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Always Enable&#x60; - Value - Always Enable for configuring RedirectionAfterPost token. * &#x60;Bootloader&#x60; - Value - Bootloader for configuring RedirectionAfterPost token. | [optional] [default to "platform-default"]
**SataModeSelect** | Pointer to **string** | BIOS Token for setting SATA Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;AHCI&#x60; - Value - AHCI for configuring SataModeSelect token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SataModeSelect token. * &#x60;LSI SW RAID&#x60; - Value - LSI SW RAID for configuring SataModeSelect token. | [optional] [default to "platform-default"]
**SelectMemoryRasConfiguration** | Pointer to **string** | BIOS Token for setting Memory RAS Configuration configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;adddc-sparing&#x60; - Value - adddc-sparing for configuring SelectMemoryRasConfiguration token. * &#x60;lockstep&#x60; - Value - lockstep for configuring SelectMemoryRasConfiguration token. * &#x60;maximum-performance&#x60; - Value - maximum-performance for configuring SelectMemoryRasConfiguration token. * &#x60;mirror-mode-1lm&#x60; - Value - mirror-mode-1lm for configuring SelectMemoryRasConfiguration token. * &#x60;mirroring&#x60; - Value - mirroring for configuring SelectMemoryRasConfiguration token. * &#x60;partial-mirror-mode-1lm&#x60; - Value - partial-mirror-mode-1lm for configuring SelectMemoryRasConfiguration token. * &#x60;sparing&#x60; - Value - sparing for configuring SelectMemoryRasConfiguration token. | [optional] [default to "platform-default"]
**SelectPprType** | Pointer to **string** | BIOS Token for setting PPR Type configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SelectPprType token. * &#x60;Hard PPR&#x60; - Value - Hard PPR for configuring SelectPprType token. * &#x60;Soft PPR&#x60; - Value - Soft PPR for configuring SelectPprType token. | [optional] [default to "platform-default"]
**SerialPortAenable** | Pointer to **string** | BIOS Token for setting Serial A Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Sev** | Pointer to **string** | BIOS Token for setting Secured Encrypted Virtualization configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;253 ASIDs&#x60; - Value - 253 ASIDs for configuring Sev token. * &#x60;509 ASIDs&#x60; - Value - 509 ASIDs for configuring Sev token. * &#x60;Auto&#x60; - Value - Auto for configuring Sev token. | [optional] [default to "platform-default"]
**SgxAutoRegistrationAgent** | Pointer to **string** | BIOS Token for setting SGX Auto MP Registration Agent configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SgxEpoch0** | Pointer to **string** | BIOS Token for setting SGX Epoch 0 configuration (0 - ffffffffffffffff Hash byte 7-0). | [optional] [default to "platform-default"]
**SgxEpoch1** | Pointer to **string** | BIOS Token for setting SGX Epoch 1 configuration (0 - ffffffffffffffff Hash byte 7-0). | [optional] [default to "platform-default"]
**SgxFactoryReset** | Pointer to **string** | BIOS Token for setting SGX Factory Reset configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SgxLePubKeyHash0** | Pointer to **string** | BIOS Token for setting SGX PubKey Hash0 configuration (0 - ffffffffffffffff Hash byte 7-0). | [optional] [default to "platform-default"]
**SgxLePubKeyHash1** | Pointer to **string** | BIOS Token for setting SGX PubKey Hash1 configuration (0 - ffffffffffffffff Hash byte 15-8). | [optional] [default to "platform-default"]
**SgxLePubKeyHash2** | Pointer to **string** | BIOS Token for setting SGX PubKey Hash2 configuration (0 - ffffffffffffffff Hash byte 23-16). | [optional] [default to "platform-default"]
**SgxLePubKeyHash3** | Pointer to **string** | BIOS Token for setting SGX PubKey Hash3 configuration (0 - ffffffffffffffff Hash byte 31-24). | [optional] [default to "platform-default"]
**SgxLeWr** | Pointer to **string** | BIOS Token for setting SGX Write Enable configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SgxPackageInfoInBandAccess** | Pointer to **string** | BIOS Token for setting SGX Package Information In-Band Access configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SgxQos** | Pointer to **string** | BIOS Token for setting SGX QoS configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SinglePctlEnable** | Pointer to **string** | BIOS Token for setting Single PCTL configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;No&#x60; - Value - No for configuring SinglePctlEnable token. * &#x60;Yes&#x60; - Value - Yes for configuring SinglePctlEnable token. | [optional] [default to "platform-default"]
**Slot10linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:10 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot10linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot10linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot10linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot10linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot10linkSpeed token. | [optional] [default to "platform-default"]
**Slot10state** | Pointer to **string** | BIOS Token for setting Slot 10 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot10state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot10state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot10state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot10state token. | [optional] [default to "platform-default"]
**Slot11linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:11 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot11linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot11linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot11linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot11linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot11linkSpeed token. | [optional] [default to "platform-default"]
**Slot11state** | Pointer to **string** | BIOS Token for setting Slot 11 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Slot12linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:12 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot12linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot12linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot12linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot12linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot12linkSpeed token. | [optional] [default to "platform-default"]
**Slot12state** | Pointer to **string** | BIOS Token for setting Slot 12 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Slot13state** | Pointer to **string** | BIOS Token for setting Slot 13 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Slot14state** | Pointer to **string** | BIOS Token for setting Slot 14 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Slot1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot1linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot1linkSpeed token. | [optional] [default to "platform-default"]
**Slot1state** | Pointer to **string** | BIOS Token for setting Slot 1 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot1state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot1state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot1state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot1state token. | [optional] [default to "platform-default"]
**Slot2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot2linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot2linkSpeed token. | [optional] [default to "platform-default"]
**Slot2state** | Pointer to **string** | BIOS Token for setting Slot 2 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot2state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot2state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot2state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot2state token. | [optional] [default to "platform-default"]
**Slot3linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 3 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot3linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot3linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot3linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot3linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot3linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot3linkSpeed token. | [optional] [default to "platform-default"]
**Slot3state** | Pointer to **string** | BIOS Token for setting Slot 3 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot3state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot3state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot3state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot3state token. | [optional] [default to "platform-default"]
**Slot4linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 4 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot4linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot4linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot4linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot4linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot4linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot4linkSpeed token. | [optional] [default to "platform-default"]
**Slot4state** | Pointer to **string** | BIOS Token for setting Slot 4 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot4state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot4state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot4state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot4state token. | [optional] [default to "platform-default"]
**Slot5linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 5 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot5linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot5linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot5linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot5linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot5linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot5linkSpeed token. | [optional] [default to "platform-default"]
**Slot5state** | Pointer to **string** | BIOS Token for setting Slot 5 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot5state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot5state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot5state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot5state token. | [optional] [default to "platform-default"]
**Slot6linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 6 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot6linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot6linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot6linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot6linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot6linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot6linkSpeed token. | [optional] [default to "platform-default"]
**Slot6state** | Pointer to **string** | BIOS Token for setting Slot 6 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot6state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot6state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot6state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot6state token. | [optional] [default to "platform-default"]
**Slot7linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 7 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot7linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot7linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot7linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot7linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot7linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot7linkSpeed token. | [optional] [default to "platform-default"]
**Slot7state** | Pointer to **string** | BIOS Token for setting Slot 7 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot7state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot7state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot7state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot7state token. | [optional] [default to "platform-default"]
**Slot8linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 8 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot8linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot8linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot8linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot8linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot8linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot8linkSpeed token. | [optional] [default to "platform-default"]
**Slot8state** | Pointer to **string** | BIOS Token for setting Slot 8 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot8state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot8state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot8state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot8state token. | [optional] [default to "platform-default"]
**Slot9linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot: 9 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Slot9linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring Slot9linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring Slot9linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring Slot9linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring Slot9linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring Slot9linkSpeed token. | [optional] [default to "platform-default"]
**Slot9state** | Pointer to **string** | BIOS Token for setting Slot 9 State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring Slot9state token. * &#x60;enabled&#x60; - Value - enabled for configuring Slot9state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring Slot9state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring Slot9state token. | [optional] [default to "platform-default"]
**SlotFlomLinkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:FLOM Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFlomLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFlomLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFlomLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFlomLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFlomLinkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme10linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 10 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme10linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme10linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme10linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme10linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme10linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme10linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme10optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 10 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme11linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 11 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme11linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme11linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme11linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme11linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme11linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme11linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme11optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 11 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme12linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 12 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme12linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme12linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme12linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme12linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme12linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme12linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme12optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 12 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme13optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 13 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme14optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 14 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme15optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 15 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme16optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 16 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme17optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 17 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme18optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 18 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme19optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 19 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Front NVME 1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme1linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme1linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme1optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme20optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 20 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme21optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 21 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme22optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 22 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme23optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 23 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme24optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 24 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Front NVME 2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme2linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme2linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme2optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme3linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 3 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme3linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme3linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme3linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme3linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme3linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme3linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme3optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme4linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 4 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme4linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme4linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme4linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme4linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme4linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme4linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme4optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 4 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme5linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 5 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme5linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme5linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme5linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme5linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme5linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme5linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme5optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 5 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme6linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 6 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme6linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme6linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme6linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme6linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme6linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme6linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme6optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 6 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme7linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 7 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme7linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme7linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme7linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme7linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme7linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme7linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme7optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 7 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme8linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 8 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme8linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme8linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme8linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme8linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme8linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme8linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme8optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 8 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontNvme9linkSpeed** | Pointer to **string** | BIOS Token for setting Front NVME 9 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontNvme9linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontNvme9linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontNvme9linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontNvme9linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontNvme9linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotFrontNvme9linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontNvme9optionRom** | Pointer to **string** | BIOS Token for setting Front NVME 9 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotFrontSlot5linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Front1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontSlot5linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontSlot5linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontSlot5linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontSlot5linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontSlot5linkSpeed token. | [optional] [default to "platform-default"]
**SlotFrontSlot6linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Front2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotFrontSlot6linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotFrontSlot6linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotFrontSlot6linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotFrontSlot6linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotFrontSlot6linkSpeed token. | [optional] [default to "platform-default"]
**SlotGpu1state** | Pointer to **string** | BIOS Token for setting GPU 1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu2state** | Pointer to **string** | BIOS Token for setting GPU 2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu3state** | Pointer to **string** | BIOS Token for setting GPU 3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu4state** | Pointer to **string** | BIOS Token for setting GPU 4 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu5state** | Pointer to **string** | BIOS Token for setting GPU 5 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu6state** | Pointer to **string** | BIOS Token for setting GPU 6 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu7state** | Pointer to **string** | BIOS Token for setting GPU 7 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotGpu8state** | Pointer to **string** | BIOS Token for setting GPU 8 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotHbaLinkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:HBA Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotHbaLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotHbaLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotHbaLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotHbaLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotHbaLinkSpeed token. | [optional] [default to "platform-default"]
**SlotHbaState** | Pointer to **string** | BIOS Token for setting PCIe Slot:HBA OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotHbaState token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotHbaState token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotHbaState token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotHbaState token. | [optional] [default to "platform-default"]
**SlotLom1link** | Pointer to **string** | BIOS Token for setting PCIe LOM:1 Link configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotLom2link** | Pointer to **string** | BIOS Token for setting PCIe LOM:2 Link configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotMezzState** | Pointer to **string** | BIOS Token for setting Slot Mezz State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotMezzState token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotMezzState token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotMezzState token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotMezzState token. | [optional] [default to "platform-default"]
**SlotMlomLinkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:MLOM Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotMlomLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotMlomLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotMlomLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotMlomLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotMlomLinkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotMlomLinkSpeed token. | [optional] [default to "platform-default"]
**SlotMlomState** | Pointer to **string** | BIOS Token for setting PCIe Slot MLOM OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotMlomState token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotMlomState token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotMlomState token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotMlomState token. | [optional] [default to "platform-default"]
**SlotMraidLinkSpeed** | Pointer to **string** | BIOS Token for setting MRAID Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotMraidLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotMraidLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotMraidLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotMraidLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotMraidLinkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotMraidLinkSpeed token. | [optional] [default to "platform-default"]
**SlotMraidState** | Pointer to **string** | BIOS Token for setting PCIe Slot MRAID OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN10state** | Pointer to **string** | BIOS Token for setting PCIe Slot N10 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN11state** | Pointer to **string** | BIOS Token for setting PCIe Slot N11 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN12state** | Pointer to **string** | BIOS Token for setting PCIe Slot N12 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN13state** | Pointer to **string** | BIOS Token for setting PCIe Slot N13 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN14state** | Pointer to **string** | BIOS Token for setting PCIe Slot N14 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN15state** | Pointer to **string** | BIOS Token for setting PCIe Slot N15 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN16state** | Pointer to **string** | BIOS Token for setting PCIe Slot N16 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN17state** | Pointer to **string** | BIOS Token for setting PCIe Slot N17 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN18state** | Pointer to **string** | BIOS Token for setting PCIe Slot N18 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN19state** | Pointer to **string** | BIOS Token for setting PCIe Slot N19 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN1state** | Pointer to **string** | BIOS Token for setting PCIe Slot N1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotN1state token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotN1state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotN1state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotN1state token. | [optional] [default to "platform-default"]
**SlotN20state** | Pointer to **string** | BIOS Token for setting PCIe Slot N20 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN21state** | Pointer to **string** | BIOS Token for setting PCIe Slot N21 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN22state** | Pointer to **string** | BIOS Token for setting PCIe Slot N22 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN23state** | Pointer to **string** | BIOS Token for setting PCIe Slot N23 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN24state** | Pointer to **string** | BIOS Token for setting PCIe Slot N24 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN2state** | Pointer to **string** | BIOS Token for setting PCIe Slot N2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotN2state token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotN2state token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotN2state token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotN2state token. | [optional] [default to "platform-default"]
**SlotN3state** | Pointer to **string** | BIOS Token for setting PCIe Slot N3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN4state** | Pointer to **string** | BIOS Token for setting PCIe Slot N4 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN5state** | Pointer to **string** | BIOS Token for setting PCIe Slot N5 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN6state** | Pointer to **string** | BIOS Token for setting PCIe Slot N6 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN7state** | Pointer to **string** | BIOS Token for setting PCIe Slot N7 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN8state** | Pointer to **string** | BIOS Token for setting PCIe Slot N8 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotN9state** | Pointer to **string** | BIOS Token for setting PCIe Slot N9 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRaidLinkSpeed** | Pointer to **string** | BIOS Token for setting RAID Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRaidLinkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRaidLinkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRaidLinkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRaidLinkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRaidLinkSpeed token. | [optional] [default to "platform-default"]
**SlotRaidState** | Pointer to **string** | BIOS Token for setting PCIe Slot RAID OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRearNvme1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRearNvme1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRearNvme1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRearNvme1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRearNvme1linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotRearNvme1linkSpeed token. | [optional] [default to "platform-default"]
**SlotRearNvme1state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRearNvme2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRearNvme2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRearNvme2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRearNvme2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRearNvme2linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotRearNvme2linkSpeed token. | [optional] [default to "platform-default"]
**SlotRearNvme2state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme3linkSpeed** | Pointer to **string** | BIOS Token for setting Rear NVME 3 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRearNvme3linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRearNvme3linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRearNvme3linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRearNvme3linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRearNvme3linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotRearNvme3linkSpeed token. | [optional] [default to "platform-default"]
**SlotRearNvme3state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme4linkSpeed** | Pointer to **string** | BIOS Token for setting Rear NVME 4 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRearNvme4linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRearNvme4linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRearNvme4linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRearNvme4linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRearNvme4linkSpeed token. * &#x60;GEN4&#x60; - Value - GEN4 for configuring SlotRearNvme4linkSpeed token. | [optional] [default to "platform-default"]
**SlotRearNvme4state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 4 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme5state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 5 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme6state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 6 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme7state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 7 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRearNvme8state** | Pointer to **string** | BIOS Token for setting PCIe Slot:Rear NVME 8 OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SlotRiser1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser1linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser1slot1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser1slot1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser1slot1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser1slot1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser1slot1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser1slot1linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser1slot2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser1slot2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser1slot2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser1slot2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser1slot2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser1slot2linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser1slot3linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser1slot3linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser1slot3linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser1slot3linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser1slot3linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser1slot3linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser2linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser2slot4linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser2slot4linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser2slot4linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser2slot4linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser2slot4linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser2slot4linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser2slot5linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser2slot5linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser2slot5linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser2slot5linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser2slot5linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser2slot5linkSpeed token. | [optional] [default to "platform-default"]
**SlotRiser2slot6linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotRiser2slot6linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotRiser2slot6linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotRiser2slot6linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotRiser2slot6linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotRiser2slot6linkSpeed token. | [optional] [default to "platform-default"]
**SlotSasState** | Pointer to **string** | BIOS Token for setting PCIe Slot:SAS OptionROM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;disabled&#x60; - Value - disabled for configuring SlotSasState token. * &#x60;enabled&#x60; - Value - enabled for configuring SlotSasState token. * &#x60;Legacy Only&#x60; - Value - Legacy Only for configuring SlotSasState token. * &#x60;UEFI Only&#x60; - Value - UEFI Only for configuring SlotSasState token. | [optional] [default to "platform-default"]
**SlotSsdSlot1linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotSsdSlot1linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotSsdSlot1linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotSsdSlot1linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotSsdSlot1linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotSsdSlot1linkSpeed token. | [optional] [default to "platform-default"]
**SlotSsdSlot2linkSpeed** | Pointer to **string** | BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SlotSsdSlot2linkSpeed token. * &#x60;Disabled&#x60; - Value - Disabled for configuring SlotSsdSlot2linkSpeed token. * &#x60;GEN1&#x60; - Value - GEN1 for configuring SlotSsdSlot2linkSpeed token. * &#x60;GEN2&#x60; - Value - GEN2 for configuring SlotSsdSlot2linkSpeed token. * &#x60;GEN3&#x60; - Value - GEN3 for configuring SlotSsdSlot2linkSpeed token. | [optional] [default to "platform-default"]
**Smee** | Pointer to **string** | BIOS Token for setting SMEE configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SmtMode** | Pointer to **string** | BIOS Token for setting SMT Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring SmtMode token. * &#x60;Off&#x60; - Value - Off for configuring SmtMode token. | [optional] [default to "platform-default"]
**Snc** | Pointer to **string** | BIOS Token for setting Sub Numa Clustering configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Snc token. * &#x60;disabled&#x60; - Value - disabled for configuring Snc token. * &#x60;enabled&#x60; - Value - enabled for configuring Snc token. | [optional] [default to "platform-default"]
**SnoopyModeFor2lm** | Pointer to **string** | BIOS Token for setting Snoopy Mode for 2LM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SnoopyModeForAd** | Pointer to **string** | BIOS Token for setting Snoopy Mode for AD configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SparingMode** | Pointer to **string** | BIOS Token for setting Sparing Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;dimm-sparing&#x60; - Value - dimm-sparing for configuring SparingMode token. * &#x60;rank-sparing&#x60; - Value - rank-sparing for configuring SparingMode token. | [optional] [default to "platform-default"]
**SrIov** | Pointer to **string** | BIOS Token for setting SR-IOV Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**StreamerPrefetch** | Pointer to **string** | BIOS Token for setting DCU Streamer Prefetch configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**SvmMode** | Pointer to **string** | BIOS Token for setting SVM Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**TerminalType** | Pointer to **string** | BIOS Token for setting Terminal Type configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;pc-ansi&#x60; - Value - pc-ansi for configuring TerminalType token. * &#x60;vt100&#x60; - Value - vt100 for configuring TerminalType token. * &#x60;vt100-plus&#x60; - Value - vt100-plus for configuring TerminalType token. * &#x60;vt-utf8&#x60; - Value - vt-utf8 for configuring TerminalType token. | [optional] [default to "platform-default"]
**TpmControl** | Pointer to **string** | BIOS Token for setting Trusted Platform Module State configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**TpmPendingOperation** | Pointer to **string** | BIOS Token for setting TPM Pending Operation configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;None&#x60; - Value - None for configuring TpmPendingOperation token. * &#x60;TpmClear&#x60; - Value - TpmClear for configuring TpmPendingOperation token. | [optional] [default to "platform-default"]
**TpmSupport** | Pointer to **string** | BIOS Token for setting TPM Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**Tsme** | Pointer to **string** | BIOS Token for setting Transparent Secure Memory Encryption configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring Tsme token. * &#x60;disabled&#x60; - Value - disabled for configuring Tsme token. * &#x60;enabled&#x60; - Value - enabled for configuring Tsme token. | [optional] [default to "platform-default"]
**TxtSupport** | Pointer to **string** | BIOS Token for setting Intel Trusted Execution Technology Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UcsmBootOrderRule** | Pointer to **string** | BIOS Token for setting Boot Order Rules configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Loose&#x60; - Value - Loose for configuring UcsmBootOrderRule token. * &#x60;Strict&#x60; - Value - Strict for configuring UcsmBootOrderRule token. | [optional] [default to "platform-default"]
**UfsDisable** | Pointer to **string** | BIOS Token for setting Uncore Frequency Scaling configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UmaBasedClustering** | Pointer to **string** | BIOS Token for setting UMA Based Clustering configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Disable (All2All)&#x60; - Value - Disable (All2All) for configuring UmaBasedClustering token. * &#x60;Hemisphere (2-clusters)&#x60; - Value - Hemisphere (2-clusters) for configuring UmaBasedClustering token. | [optional] [default to "platform-default"]
**UsbEmul6064** | Pointer to **string** | BIOS Token for setting Port 60/64 Emulation configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortFront** | Pointer to **string** | BIOS Token for setting USB Port Front configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortInternal** | Pointer to **string** | BIOS Token for setting USB Port Internal configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortKvm** | Pointer to **string** | BIOS Token for setting USB Port KVM configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortRear** | Pointer to **string** | BIOS Token for setting USB Port Rear configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortSdCard** | Pointer to **string** | BIOS Token for setting USB Port SD Card configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbPortVmedia** | Pointer to **string** | BIOS Token for setting USB Port VMedia configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**UsbXhciSupport** | Pointer to **string** | BIOS Token for setting XHCI Legacy Support configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**VgaPriority** | Pointer to **string** | BIOS Token for setting VGA Priority configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Offboard&#x60; - Value - Offboard for configuring VgaPriority token. * &#x60;Onboard&#x60; - Value - Onboard for configuring VgaPriority token. * &#x60;Onboard VGA Disabled&#x60; - Value - Onboard VGA Disabled for configuring VgaPriority token. | [optional] [default to "platform-default"]
**VmdEnable** | Pointer to **string** | BIOS Token for setting VMD Enablement configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;enabled&#x60; - Enables the BIOS setting. * &#x60;disabled&#x60; - Disables the BIOS setting. | [optional] [default to "platform-default"]
**VolMemoryMode** | Pointer to **string** | BIOS Token for setting Volatile Memory Mode configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;1LM&#x60; - Value - 1LM for configuring VolMemoryMode token. * &#x60;2LM&#x60; - Value - 2LM for configuring VolMemoryMode token. | [optional] [default to "platform-default"]
**WorkLoadConfig** | Pointer to **string** | BIOS Token for setting Workload Configuration configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Balanced&#x60; - Value - Balanced for configuring WorkLoadConfig token. * &#x60;I/O Sensitive&#x60; - Value - I/O Sensitive for configuring WorkLoadConfig token. * &#x60;NUMA&#x60; - Value - NUMA for configuring WorkLoadConfig token. * &#x60;UMA&#x60; - Value - UMA for configuring WorkLoadConfig token. | [optional] [default to "platform-default"]
**XptPrefetch** | Pointer to **string** | BIOS Token for setting XPT Prefetch configuration. * &#x60;platform-default&#x60; - Default value used by the platform for the BIOS setting. * &#x60;Auto&#x60; - Value - Auto for configuring XptPrefetch token. * &#x60;disabled&#x60; - Value - disabled for configuring XptPrefetch token. * &#x60;enabled&#x60; - Value - enabled for configuring XptPrefetch token. | [optional] [default to "platform-default"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewBiosPolicyAllOf

`func NewBiosPolicyAllOf(classId string, objectType string, ) *BiosPolicyAllOf`

NewBiosPolicyAllOf instantiates a new BiosPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosPolicyAllOfWithDefaults

`func NewBiosPolicyAllOfWithDefaults() *BiosPolicyAllOf`

NewBiosPolicyAllOfWithDefaults instantiates a new BiosPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAcsControlGpu1state

`func (o *BiosPolicyAllOf) GetAcsControlGpu1state() string`

GetAcsControlGpu1state returns the AcsControlGpu1state field if non-nil, zero value otherwise.

### GetAcsControlGpu1stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu1stateOk() (*string, bool)`

GetAcsControlGpu1stateOk returns a tuple with the AcsControlGpu1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu1state

`func (o *BiosPolicyAllOf) SetAcsControlGpu1state(v string)`

SetAcsControlGpu1state sets AcsControlGpu1state field to given value.

### HasAcsControlGpu1state

`func (o *BiosPolicyAllOf) HasAcsControlGpu1state() bool`

HasAcsControlGpu1state returns a boolean if a field has been set.

### GetAcsControlGpu2state

`func (o *BiosPolicyAllOf) GetAcsControlGpu2state() string`

GetAcsControlGpu2state returns the AcsControlGpu2state field if non-nil, zero value otherwise.

### GetAcsControlGpu2stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu2stateOk() (*string, bool)`

GetAcsControlGpu2stateOk returns a tuple with the AcsControlGpu2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu2state

`func (o *BiosPolicyAllOf) SetAcsControlGpu2state(v string)`

SetAcsControlGpu2state sets AcsControlGpu2state field to given value.

### HasAcsControlGpu2state

`func (o *BiosPolicyAllOf) HasAcsControlGpu2state() bool`

HasAcsControlGpu2state returns a boolean if a field has been set.

### GetAcsControlGpu3state

`func (o *BiosPolicyAllOf) GetAcsControlGpu3state() string`

GetAcsControlGpu3state returns the AcsControlGpu3state field if non-nil, zero value otherwise.

### GetAcsControlGpu3stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu3stateOk() (*string, bool)`

GetAcsControlGpu3stateOk returns a tuple with the AcsControlGpu3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu3state

`func (o *BiosPolicyAllOf) SetAcsControlGpu3state(v string)`

SetAcsControlGpu3state sets AcsControlGpu3state field to given value.

### HasAcsControlGpu3state

`func (o *BiosPolicyAllOf) HasAcsControlGpu3state() bool`

HasAcsControlGpu3state returns a boolean if a field has been set.

### GetAcsControlGpu4state

`func (o *BiosPolicyAllOf) GetAcsControlGpu4state() string`

GetAcsControlGpu4state returns the AcsControlGpu4state field if non-nil, zero value otherwise.

### GetAcsControlGpu4stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu4stateOk() (*string, bool)`

GetAcsControlGpu4stateOk returns a tuple with the AcsControlGpu4state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu4state

`func (o *BiosPolicyAllOf) SetAcsControlGpu4state(v string)`

SetAcsControlGpu4state sets AcsControlGpu4state field to given value.

### HasAcsControlGpu4state

`func (o *BiosPolicyAllOf) HasAcsControlGpu4state() bool`

HasAcsControlGpu4state returns a boolean if a field has been set.

### GetAcsControlGpu5state

`func (o *BiosPolicyAllOf) GetAcsControlGpu5state() string`

GetAcsControlGpu5state returns the AcsControlGpu5state field if non-nil, zero value otherwise.

### GetAcsControlGpu5stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu5stateOk() (*string, bool)`

GetAcsControlGpu5stateOk returns a tuple with the AcsControlGpu5state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu5state

`func (o *BiosPolicyAllOf) SetAcsControlGpu5state(v string)`

SetAcsControlGpu5state sets AcsControlGpu5state field to given value.

### HasAcsControlGpu5state

`func (o *BiosPolicyAllOf) HasAcsControlGpu5state() bool`

HasAcsControlGpu5state returns a boolean if a field has been set.

### GetAcsControlGpu6state

`func (o *BiosPolicyAllOf) GetAcsControlGpu6state() string`

GetAcsControlGpu6state returns the AcsControlGpu6state field if non-nil, zero value otherwise.

### GetAcsControlGpu6stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu6stateOk() (*string, bool)`

GetAcsControlGpu6stateOk returns a tuple with the AcsControlGpu6state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu6state

`func (o *BiosPolicyAllOf) SetAcsControlGpu6state(v string)`

SetAcsControlGpu6state sets AcsControlGpu6state field to given value.

### HasAcsControlGpu6state

`func (o *BiosPolicyAllOf) HasAcsControlGpu6state() bool`

HasAcsControlGpu6state returns a boolean if a field has been set.

### GetAcsControlGpu7state

`func (o *BiosPolicyAllOf) GetAcsControlGpu7state() string`

GetAcsControlGpu7state returns the AcsControlGpu7state field if non-nil, zero value otherwise.

### GetAcsControlGpu7stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu7stateOk() (*string, bool)`

GetAcsControlGpu7stateOk returns a tuple with the AcsControlGpu7state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu7state

`func (o *BiosPolicyAllOf) SetAcsControlGpu7state(v string)`

SetAcsControlGpu7state sets AcsControlGpu7state field to given value.

### HasAcsControlGpu7state

`func (o *BiosPolicyAllOf) HasAcsControlGpu7state() bool`

HasAcsControlGpu7state returns a boolean if a field has been set.

### GetAcsControlGpu8state

`func (o *BiosPolicyAllOf) GetAcsControlGpu8state() string`

GetAcsControlGpu8state returns the AcsControlGpu8state field if non-nil, zero value otherwise.

### GetAcsControlGpu8stateOk

`func (o *BiosPolicyAllOf) GetAcsControlGpu8stateOk() (*string, bool)`

GetAcsControlGpu8stateOk returns a tuple with the AcsControlGpu8state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlGpu8state

`func (o *BiosPolicyAllOf) SetAcsControlGpu8state(v string)`

SetAcsControlGpu8state sets AcsControlGpu8state field to given value.

### HasAcsControlGpu8state

`func (o *BiosPolicyAllOf) HasAcsControlGpu8state() bool`

HasAcsControlGpu8state returns a boolean if a field has been set.

### GetAcsControlSlot11state

`func (o *BiosPolicyAllOf) GetAcsControlSlot11state() string`

GetAcsControlSlot11state returns the AcsControlSlot11state field if non-nil, zero value otherwise.

### GetAcsControlSlot11stateOk

`func (o *BiosPolicyAllOf) GetAcsControlSlot11stateOk() (*string, bool)`

GetAcsControlSlot11stateOk returns a tuple with the AcsControlSlot11state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlSlot11state

`func (o *BiosPolicyAllOf) SetAcsControlSlot11state(v string)`

SetAcsControlSlot11state sets AcsControlSlot11state field to given value.

### HasAcsControlSlot11state

`func (o *BiosPolicyAllOf) HasAcsControlSlot11state() bool`

HasAcsControlSlot11state returns a boolean if a field has been set.

### GetAcsControlSlot12state

`func (o *BiosPolicyAllOf) GetAcsControlSlot12state() string`

GetAcsControlSlot12state returns the AcsControlSlot12state field if non-nil, zero value otherwise.

### GetAcsControlSlot12stateOk

`func (o *BiosPolicyAllOf) GetAcsControlSlot12stateOk() (*string, bool)`

GetAcsControlSlot12stateOk returns a tuple with the AcsControlSlot12state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlSlot12state

`func (o *BiosPolicyAllOf) SetAcsControlSlot12state(v string)`

SetAcsControlSlot12state sets AcsControlSlot12state field to given value.

### HasAcsControlSlot12state

`func (o *BiosPolicyAllOf) HasAcsControlSlot12state() bool`

HasAcsControlSlot12state returns a boolean if a field has been set.

### GetAcsControlSlot13state

`func (o *BiosPolicyAllOf) GetAcsControlSlot13state() string`

GetAcsControlSlot13state returns the AcsControlSlot13state field if non-nil, zero value otherwise.

### GetAcsControlSlot13stateOk

`func (o *BiosPolicyAllOf) GetAcsControlSlot13stateOk() (*string, bool)`

GetAcsControlSlot13stateOk returns a tuple with the AcsControlSlot13state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlSlot13state

`func (o *BiosPolicyAllOf) SetAcsControlSlot13state(v string)`

SetAcsControlSlot13state sets AcsControlSlot13state field to given value.

### HasAcsControlSlot13state

`func (o *BiosPolicyAllOf) HasAcsControlSlot13state() bool`

HasAcsControlSlot13state returns a boolean if a field has been set.

### GetAcsControlSlot14state

`func (o *BiosPolicyAllOf) GetAcsControlSlot14state() string`

GetAcsControlSlot14state returns the AcsControlSlot14state field if non-nil, zero value otherwise.

### GetAcsControlSlot14stateOk

`func (o *BiosPolicyAllOf) GetAcsControlSlot14stateOk() (*string, bool)`

GetAcsControlSlot14stateOk returns a tuple with the AcsControlSlot14state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsControlSlot14state

`func (o *BiosPolicyAllOf) SetAcsControlSlot14state(v string)`

SetAcsControlSlot14state sets AcsControlSlot14state field to given value.

### HasAcsControlSlot14state

`func (o *BiosPolicyAllOf) HasAcsControlSlot14state() bool`

HasAcsControlSlot14state returns a boolean if a field has been set.

### GetAdjacentCacheLinePrefetch

`func (o *BiosPolicyAllOf) GetAdjacentCacheLinePrefetch() string`

GetAdjacentCacheLinePrefetch returns the AdjacentCacheLinePrefetch field if non-nil, zero value otherwise.

### GetAdjacentCacheLinePrefetchOk

`func (o *BiosPolicyAllOf) GetAdjacentCacheLinePrefetchOk() (*string, bool)`

GetAdjacentCacheLinePrefetchOk returns a tuple with the AdjacentCacheLinePrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentCacheLinePrefetch

`func (o *BiosPolicyAllOf) SetAdjacentCacheLinePrefetch(v string)`

SetAdjacentCacheLinePrefetch sets AdjacentCacheLinePrefetch field to given value.

### HasAdjacentCacheLinePrefetch

`func (o *BiosPolicyAllOf) HasAdjacentCacheLinePrefetch() bool`

HasAdjacentCacheLinePrefetch returns a boolean if a field has been set.

### GetAdvancedMemTest

`func (o *BiosPolicyAllOf) GetAdvancedMemTest() string`

GetAdvancedMemTest returns the AdvancedMemTest field if non-nil, zero value otherwise.

### GetAdvancedMemTestOk

`func (o *BiosPolicyAllOf) GetAdvancedMemTestOk() (*string, bool)`

GetAdvancedMemTestOk returns a tuple with the AdvancedMemTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedMemTest

`func (o *BiosPolicyAllOf) SetAdvancedMemTest(v string)`

SetAdvancedMemTest sets AdvancedMemTest field to given value.

### HasAdvancedMemTest

`func (o *BiosPolicyAllOf) HasAdvancedMemTest() bool`

HasAdvancedMemTest returns a boolean if a field has been set.

### GetAllUsbDevices

`func (o *BiosPolicyAllOf) GetAllUsbDevices() string`

GetAllUsbDevices returns the AllUsbDevices field if non-nil, zero value otherwise.

### GetAllUsbDevicesOk

`func (o *BiosPolicyAllOf) GetAllUsbDevicesOk() (*string, bool)`

GetAllUsbDevicesOk returns a tuple with the AllUsbDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsbDevices

`func (o *BiosPolicyAllOf) SetAllUsbDevices(v string)`

SetAllUsbDevices sets AllUsbDevices field to given value.

### HasAllUsbDevices

`func (o *BiosPolicyAllOf) HasAllUsbDevices() bool`

HasAllUsbDevices returns a boolean if a field has been set.

### GetAltitude

`func (o *BiosPolicyAllOf) GetAltitude() string`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *BiosPolicyAllOf) GetAltitudeOk() (*string, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *BiosPolicyAllOf) SetAltitude(v string)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *BiosPolicyAllOf) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetAspmSupport

`func (o *BiosPolicyAllOf) GetAspmSupport() string`

GetAspmSupport returns the AspmSupport field if non-nil, zero value otherwise.

### GetAspmSupportOk

`func (o *BiosPolicyAllOf) GetAspmSupportOk() (*string, bool)`

GetAspmSupportOk returns a tuple with the AspmSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspmSupport

`func (o *BiosPolicyAllOf) SetAspmSupport(v string)`

SetAspmSupport sets AspmSupport field to given value.

### HasAspmSupport

`func (o *BiosPolicyAllOf) HasAspmSupport() bool`

HasAspmSupport returns a boolean if a field has been set.

### GetAssertNmiOnPerr

`func (o *BiosPolicyAllOf) GetAssertNmiOnPerr() string`

GetAssertNmiOnPerr returns the AssertNmiOnPerr field if non-nil, zero value otherwise.

### GetAssertNmiOnPerrOk

`func (o *BiosPolicyAllOf) GetAssertNmiOnPerrOk() (*string, bool)`

GetAssertNmiOnPerrOk returns a tuple with the AssertNmiOnPerr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertNmiOnPerr

`func (o *BiosPolicyAllOf) SetAssertNmiOnPerr(v string)`

SetAssertNmiOnPerr sets AssertNmiOnPerr field to given value.

### HasAssertNmiOnPerr

`func (o *BiosPolicyAllOf) HasAssertNmiOnPerr() bool`

HasAssertNmiOnPerr returns a boolean if a field has been set.

### GetAssertNmiOnSerr

`func (o *BiosPolicyAllOf) GetAssertNmiOnSerr() string`

GetAssertNmiOnSerr returns the AssertNmiOnSerr field if non-nil, zero value otherwise.

### GetAssertNmiOnSerrOk

`func (o *BiosPolicyAllOf) GetAssertNmiOnSerrOk() (*string, bool)`

GetAssertNmiOnSerrOk returns a tuple with the AssertNmiOnSerr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertNmiOnSerr

`func (o *BiosPolicyAllOf) SetAssertNmiOnSerr(v string)`

SetAssertNmiOnSerr sets AssertNmiOnSerr field to given value.

### HasAssertNmiOnSerr

`func (o *BiosPolicyAllOf) HasAssertNmiOnSerr() bool`

HasAssertNmiOnSerr returns a boolean if a field has been set.

### GetAutoCcState

`func (o *BiosPolicyAllOf) GetAutoCcState() string`

GetAutoCcState returns the AutoCcState field if non-nil, zero value otherwise.

### GetAutoCcStateOk

`func (o *BiosPolicyAllOf) GetAutoCcStateOk() (*string, bool)`

GetAutoCcStateOk returns a tuple with the AutoCcState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCcState

`func (o *BiosPolicyAllOf) SetAutoCcState(v string)`

SetAutoCcState sets AutoCcState field to given value.

### HasAutoCcState

`func (o *BiosPolicyAllOf) HasAutoCcState() bool`

HasAutoCcState returns a boolean if a field has been set.

### GetAutonumousCstateEnable

`func (o *BiosPolicyAllOf) GetAutonumousCstateEnable() string`

GetAutonumousCstateEnable returns the AutonumousCstateEnable field if non-nil, zero value otherwise.

### GetAutonumousCstateEnableOk

`func (o *BiosPolicyAllOf) GetAutonumousCstateEnableOk() (*string, bool)`

GetAutonumousCstateEnableOk returns a tuple with the AutonumousCstateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonumousCstateEnable

`func (o *BiosPolicyAllOf) SetAutonumousCstateEnable(v string)`

SetAutonumousCstateEnable sets AutonumousCstateEnable field to given value.

### HasAutonumousCstateEnable

`func (o *BiosPolicyAllOf) HasAutonumousCstateEnable() bool`

HasAutonumousCstateEnable returns a boolean if a field has been set.

### GetBaudRate

`func (o *BiosPolicyAllOf) GetBaudRate() string`

GetBaudRate returns the BaudRate field if non-nil, zero value otherwise.

### GetBaudRateOk

`func (o *BiosPolicyAllOf) GetBaudRateOk() (*string, bool)`

GetBaudRateOk returns a tuple with the BaudRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaudRate

`func (o *BiosPolicyAllOf) SetBaudRate(v string)`

SetBaudRate sets BaudRate field to given value.

### HasBaudRate

`func (o *BiosPolicyAllOf) HasBaudRate() bool`

HasBaudRate returns a boolean if a field has been set.

### GetBmeDmaMitigation

`func (o *BiosPolicyAllOf) GetBmeDmaMitigation() string`

GetBmeDmaMitigation returns the BmeDmaMitigation field if non-nil, zero value otherwise.

### GetBmeDmaMitigationOk

`func (o *BiosPolicyAllOf) GetBmeDmaMitigationOk() (*string, bool)`

GetBmeDmaMitigationOk returns a tuple with the BmeDmaMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmeDmaMitigation

`func (o *BiosPolicyAllOf) SetBmeDmaMitigation(v string)`

SetBmeDmaMitigation sets BmeDmaMitigation field to given value.

### HasBmeDmaMitigation

`func (o *BiosPolicyAllOf) HasBmeDmaMitigation() bool`

HasBmeDmaMitigation returns a boolean if a field has been set.

### GetBootOptionNumRetry

`func (o *BiosPolicyAllOf) GetBootOptionNumRetry() string`

GetBootOptionNumRetry returns the BootOptionNumRetry field if non-nil, zero value otherwise.

### GetBootOptionNumRetryOk

`func (o *BiosPolicyAllOf) GetBootOptionNumRetryOk() (*string, bool)`

GetBootOptionNumRetryOk returns a tuple with the BootOptionNumRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionNumRetry

`func (o *BiosPolicyAllOf) SetBootOptionNumRetry(v string)`

SetBootOptionNumRetry sets BootOptionNumRetry field to given value.

### HasBootOptionNumRetry

`func (o *BiosPolicyAllOf) HasBootOptionNumRetry() bool`

HasBootOptionNumRetry returns a boolean if a field has been set.

### GetBootOptionReCoolDown

`func (o *BiosPolicyAllOf) GetBootOptionReCoolDown() string`

GetBootOptionReCoolDown returns the BootOptionReCoolDown field if non-nil, zero value otherwise.

### GetBootOptionReCoolDownOk

`func (o *BiosPolicyAllOf) GetBootOptionReCoolDownOk() (*string, bool)`

GetBootOptionReCoolDownOk returns a tuple with the BootOptionReCoolDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionReCoolDown

`func (o *BiosPolicyAllOf) SetBootOptionReCoolDown(v string)`

SetBootOptionReCoolDown sets BootOptionReCoolDown field to given value.

### HasBootOptionReCoolDown

`func (o *BiosPolicyAllOf) HasBootOptionReCoolDown() bool`

HasBootOptionReCoolDown returns a boolean if a field has been set.

### GetBootOptionRetry

`func (o *BiosPolicyAllOf) GetBootOptionRetry() string`

GetBootOptionRetry returns the BootOptionRetry field if non-nil, zero value otherwise.

### GetBootOptionRetryOk

`func (o *BiosPolicyAllOf) GetBootOptionRetryOk() (*string, bool)`

GetBootOptionRetryOk returns a tuple with the BootOptionRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionRetry

`func (o *BiosPolicyAllOf) SetBootOptionRetry(v string)`

SetBootOptionRetry sets BootOptionRetry field to given value.

### HasBootOptionRetry

`func (o *BiosPolicyAllOf) HasBootOptionRetry() bool`

HasBootOptionRetry returns a boolean if a field has been set.

### GetBootPerformanceMode

`func (o *BiosPolicyAllOf) GetBootPerformanceMode() string`

GetBootPerformanceMode returns the BootPerformanceMode field if non-nil, zero value otherwise.

### GetBootPerformanceModeOk

`func (o *BiosPolicyAllOf) GetBootPerformanceModeOk() (*string, bool)`

GetBootPerformanceModeOk returns a tuple with the BootPerformanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPerformanceMode

`func (o *BiosPolicyAllOf) SetBootPerformanceMode(v string)`

SetBootPerformanceMode sets BootPerformanceMode field to given value.

### HasBootPerformanceMode

`func (o *BiosPolicyAllOf) HasBootPerformanceMode() bool`

HasBootPerformanceMode returns a boolean if a field has been set.

### GetBurstAndPostponedRefresh

`func (o *BiosPolicyAllOf) GetBurstAndPostponedRefresh() string`

GetBurstAndPostponedRefresh returns the BurstAndPostponedRefresh field if non-nil, zero value otherwise.

### GetBurstAndPostponedRefreshOk

`func (o *BiosPolicyAllOf) GetBurstAndPostponedRefreshOk() (*string, bool)`

GetBurstAndPostponedRefreshOk returns a tuple with the BurstAndPostponedRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstAndPostponedRefresh

`func (o *BiosPolicyAllOf) SetBurstAndPostponedRefresh(v string)`

SetBurstAndPostponedRefresh sets BurstAndPostponedRefresh field to given value.

### HasBurstAndPostponedRefresh

`func (o *BiosPolicyAllOf) HasBurstAndPostponedRefresh() bool`

HasBurstAndPostponedRefresh returns a boolean if a field has been set.

### GetCbsCmnApbdis

`func (o *BiosPolicyAllOf) GetCbsCmnApbdis() string`

GetCbsCmnApbdis returns the CbsCmnApbdis field if non-nil, zero value otherwise.

### GetCbsCmnApbdisOk

`func (o *BiosPolicyAllOf) GetCbsCmnApbdisOk() (*string, bool)`

GetCbsCmnApbdisOk returns a tuple with the CbsCmnApbdis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnApbdis

`func (o *BiosPolicyAllOf) SetCbsCmnApbdis(v string)`

SetCbsCmnApbdis sets CbsCmnApbdis field to given value.

### HasCbsCmnApbdis

`func (o *BiosPolicyAllOf) HasCbsCmnApbdis() bool`

HasCbsCmnApbdis returns a boolean if a field has been set.

### GetCbsCmnCpuCpb

`func (o *BiosPolicyAllOf) GetCbsCmnCpuCpb() string`

GetCbsCmnCpuCpb returns the CbsCmnCpuCpb field if non-nil, zero value otherwise.

### GetCbsCmnCpuCpbOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuCpbOk() (*string, bool)`

GetCbsCmnCpuCpbOk returns a tuple with the CbsCmnCpuCpb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuCpb

`func (o *BiosPolicyAllOf) SetCbsCmnCpuCpb(v string)`

SetCbsCmnCpuCpb sets CbsCmnCpuCpb field to given value.

### HasCbsCmnCpuCpb

`func (o *BiosPolicyAllOf) HasCbsCmnCpuCpb() bool`

HasCbsCmnCpuCpb returns a boolean if a field has been set.

### GetCbsCmnCpuGenDowncoreCtrl

`func (o *BiosPolicyAllOf) GetCbsCmnCpuGenDowncoreCtrl() string`

GetCbsCmnCpuGenDowncoreCtrl returns the CbsCmnCpuGenDowncoreCtrl field if non-nil, zero value otherwise.

### GetCbsCmnCpuGenDowncoreCtrlOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuGenDowncoreCtrlOk() (*string, bool)`

GetCbsCmnCpuGenDowncoreCtrlOk returns a tuple with the CbsCmnCpuGenDowncoreCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuGenDowncoreCtrl

`func (o *BiosPolicyAllOf) SetCbsCmnCpuGenDowncoreCtrl(v string)`

SetCbsCmnCpuGenDowncoreCtrl sets CbsCmnCpuGenDowncoreCtrl field to given value.

### HasCbsCmnCpuGenDowncoreCtrl

`func (o *BiosPolicyAllOf) HasCbsCmnCpuGenDowncoreCtrl() bool`

HasCbsCmnCpuGenDowncoreCtrl returns a boolean if a field has been set.

### GetCbsCmnCpuGlobalCstateCtrl

`func (o *BiosPolicyAllOf) GetCbsCmnCpuGlobalCstateCtrl() string`

GetCbsCmnCpuGlobalCstateCtrl returns the CbsCmnCpuGlobalCstateCtrl field if non-nil, zero value otherwise.

### GetCbsCmnCpuGlobalCstateCtrlOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuGlobalCstateCtrlOk() (*string, bool)`

GetCbsCmnCpuGlobalCstateCtrlOk returns a tuple with the CbsCmnCpuGlobalCstateCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuGlobalCstateCtrl

`func (o *BiosPolicyAllOf) SetCbsCmnCpuGlobalCstateCtrl(v string)`

SetCbsCmnCpuGlobalCstateCtrl sets CbsCmnCpuGlobalCstateCtrl field to given value.

### HasCbsCmnCpuGlobalCstateCtrl

`func (o *BiosPolicyAllOf) HasCbsCmnCpuGlobalCstateCtrl() bool`

HasCbsCmnCpuGlobalCstateCtrl returns a boolean if a field has been set.

### GetCbsCmnCpuL1streamHwPrefetcher

`func (o *BiosPolicyAllOf) GetCbsCmnCpuL1streamHwPrefetcher() string`

GetCbsCmnCpuL1streamHwPrefetcher returns the CbsCmnCpuL1streamHwPrefetcher field if non-nil, zero value otherwise.

### GetCbsCmnCpuL1streamHwPrefetcherOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuL1streamHwPrefetcherOk() (*string, bool)`

GetCbsCmnCpuL1streamHwPrefetcherOk returns a tuple with the CbsCmnCpuL1streamHwPrefetcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuL1streamHwPrefetcher

`func (o *BiosPolicyAllOf) SetCbsCmnCpuL1streamHwPrefetcher(v string)`

SetCbsCmnCpuL1streamHwPrefetcher sets CbsCmnCpuL1streamHwPrefetcher field to given value.

### HasCbsCmnCpuL1streamHwPrefetcher

`func (o *BiosPolicyAllOf) HasCbsCmnCpuL1streamHwPrefetcher() bool`

HasCbsCmnCpuL1streamHwPrefetcher returns a boolean if a field has been set.

### GetCbsCmnCpuL2streamHwPrefetcher

`func (o *BiosPolicyAllOf) GetCbsCmnCpuL2streamHwPrefetcher() string`

GetCbsCmnCpuL2streamHwPrefetcher returns the CbsCmnCpuL2streamHwPrefetcher field if non-nil, zero value otherwise.

### GetCbsCmnCpuL2streamHwPrefetcherOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuL2streamHwPrefetcherOk() (*string, bool)`

GetCbsCmnCpuL2streamHwPrefetcherOk returns a tuple with the CbsCmnCpuL2streamHwPrefetcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuL2streamHwPrefetcher

`func (o *BiosPolicyAllOf) SetCbsCmnCpuL2streamHwPrefetcher(v string)`

SetCbsCmnCpuL2streamHwPrefetcher sets CbsCmnCpuL2streamHwPrefetcher field to given value.

### HasCbsCmnCpuL2streamHwPrefetcher

`func (o *BiosPolicyAllOf) HasCbsCmnCpuL2streamHwPrefetcher() bool`

HasCbsCmnCpuL2streamHwPrefetcher returns a boolean if a field has been set.

### GetCbsCmnCpuSmee

`func (o *BiosPolicyAllOf) GetCbsCmnCpuSmee() string`

GetCbsCmnCpuSmee returns the CbsCmnCpuSmee field if non-nil, zero value otherwise.

### GetCbsCmnCpuSmeeOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuSmeeOk() (*string, bool)`

GetCbsCmnCpuSmeeOk returns a tuple with the CbsCmnCpuSmee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuSmee

`func (o *BiosPolicyAllOf) SetCbsCmnCpuSmee(v string)`

SetCbsCmnCpuSmee sets CbsCmnCpuSmee field to given value.

### HasCbsCmnCpuSmee

`func (o *BiosPolicyAllOf) HasCbsCmnCpuSmee() bool`

HasCbsCmnCpuSmee returns a boolean if a field has been set.

### GetCbsCmnCpuStreamingStoresCtrl

`func (o *BiosPolicyAllOf) GetCbsCmnCpuStreamingStoresCtrl() string`

GetCbsCmnCpuStreamingStoresCtrl returns the CbsCmnCpuStreamingStoresCtrl field if non-nil, zero value otherwise.

### GetCbsCmnCpuStreamingStoresCtrlOk

`func (o *BiosPolicyAllOf) GetCbsCmnCpuStreamingStoresCtrlOk() (*string, bool)`

GetCbsCmnCpuStreamingStoresCtrlOk returns a tuple with the CbsCmnCpuStreamingStoresCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnCpuStreamingStoresCtrl

`func (o *BiosPolicyAllOf) SetCbsCmnCpuStreamingStoresCtrl(v string)`

SetCbsCmnCpuStreamingStoresCtrl sets CbsCmnCpuStreamingStoresCtrl field to given value.

### HasCbsCmnCpuStreamingStoresCtrl

`func (o *BiosPolicyAllOf) HasCbsCmnCpuStreamingStoresCtrl() bool`

HasCbsCmnCpuStreamingStoresCtrl returns a boolean if a field has been set.

### GetCbsCmnDeterminismSlider

`func (o *BiosPolicyAllOf) GetCbsCmnDeterminismSlider() string`

GetCbsCmnDeterminismSlider returns the CbsCmnDeterminismSlider field if non-nil, zero value otherwise.

### GetCbsCmnDeterminismSliderOk

`func (o *BiosPolicyAllOf) GetCbsCmnDeterminismSliderOk() (*string, bool)`

GetCbsCmnDeterminismSliderOk returns a tuple with the CbsCmnDeterminismSlider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnDeterminismSlider

`func (o *BiosPolicyAllOf) SetCbsCmnDeterminismSlider(v string)`

SetCbsCmnDeterminismSlider sets CbsCmnDeterminismSlider field to given value.

### HasCbsCmnDeterminismSlider

`func (o *BiosPolicyAllOf) HasCbsCmnDeterminismSlider() bool`

HasCbsCmnDeterminismSlider returns a boolean if a field has been set.

### GetCbsCmnEfficiencyModeEn

`func (o *BiosPolicyAllOf) GetCbsCmnEfficiencyModeEn() string`

GetCbsCmnEfficiencyModeEn returns the CbsCmnEfficiencyModeEn field if non-nil, zero value otherwise.

### GetCbsCmnEfficiencyModeEnOk

`func (o *BiosPolicyAllOf) GetCbsCmnEfficiencyModeEnOk() (*string, bool)`

GetCbsCmnEfficiencyModeEnOk returns a tuple with the CbsCmnEfficiencyModeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnEfficiencyModeEn

`func (o *BiosPolicyAllOf) SetCbsCmnEfficiencyModeEn(v string)`

SetCbsCmnEfficiencyModeEn sets CbsCmnEfficiencyModeEn field to given value.

### HasCbsCmnEfficiencyModeEn

`func (o *BiosPolicyAllOf) HasCbsCmnEfficiencyModeEn() bool`

HasCbsCmnEfficiencyModeEn returns a boolean if a field has been set.

### GetCbsCmnFixedSocPstate

`func (o *BiosPolicyAllOf) GetCbsCmnFixedSocPstate() string`

GetCbsCmnFixedSocPstate returns the CbsCmnFixedSocPstate field if non-nil, zero value otherwise.

### GetCbsCmnFixedSocPstateOk

`func (o *BiosPolicyAllOf) GetCbsCmnFixedSocPstateOk() (*string, bool)`

GetCbsCmnFixedSocPstateOk returns a tuple with the CbsCmnFixedSocPstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnFixedSocPstate

`func (o *BiosPolicyAllOf) SetCbsCmnFixedSocPstate(v string)`

SetCbsCmnFixedSocPstate sets CbsCmnFixedSocPstate field to given value.

### HasCbsCmnFixedSocPstate

`func (o *BiosPolicyAllOf) HasCbsCmnFixedSocPstate() bool`

HasCbsCmnFixedSocPstate returns a boolean if a field has been set.

### GetCbsCmnGnbNbIommu

`func (o *BiosPolicyAllOf) GetCbsCmnGnbNbIommu() string`

GetCbsCmnGnbNbIommu returns the CbsCmnGnbNbIommu field if non-nil, zero value otherwise.

### GetCbsCmnGnbNbIommuOk

`func (o *BiosPolicyAllOf) GetCbsCmnGnbNbIommuOk() (*string, bool)`

GetCbsCmnGnbNbIommuOk returns a tuple with the CbsCmnGnbNbIommu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnGnbNbIommu

`func (o *BiosPolicyAllOf) SetCbsCmnGnbNbIommu(v string)`

SetCbsCmnGnbNbIommu sets CbsCmnGnbNbIommu field to given value.

### HasCbsCmnGnbNbIommu

`func (o *BiosPolicyAllOf) HasCbsCmnGnbNbIommu() bool`

HasCbsCmnGnbNbIommu returns a boolean if a field has been set.

### GetCbsCmnGnbSmuDfCstates

`func (o *BiosPolicyAllOf) GetCbsCmnGnbSmuDfCstates() string`

GetCbsCmnGnbSmuDfCstates returns the CbsCmnGnbSmuDfCstates field if non-nil, zero value otherwise.

### GetCbsCmnGnbSmuDfCstatesOk

`func (o *BiosPolicyAllOf) GetCbsCmnGnbSmuDfCstatesOk() (*string, bool)`

GetCbsCmnGnbSmuDfCstatesOk returns a tuple with the CbsCmnGnbSmuDfCstates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnGnbSmuDfCstates

`func (o *BiosPolicyAllOf) SetCbsCmnGnbSmuDfCstates(v string)`

SetCbsCmnGnbSmuDfCstates sets CbsCmnGnbSmuDfCstates field to given value.

### HasCbsCmnGnbSmuDfCstates

`func (o *BiosPolicyAllOf) HasCbsCmnGnbSmuDfCstates() bool`

HasCbsCmnGnbSmuDfCstates returns a boolean if a field has been set.

### GetCbsCmnGnbSmucppc

`func (o *BiosPolicyAllOf) GetCbsCmnGnbSmucppc() string`

GetCbsCmnGnbSmucppc returns the CbsCmnGnbSmucppc field if non-nil, zero value otherwise.

### GetCbsCmnGnbSmucppcOk

`func (o *BiosPolicyAllOf) GetCbsCmnGnbSmucppcOk() (*string, bool)`

GetCbsCmnGnbSmucppcOk returns a tuple with the CbsCmnGnbSmucppc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnGnbSmucppc

`func (o *BiosPolicyAllOf) SetCbsCmnGnbSmucppc(v string)`

SetCbsCmnGnbSmucppc sets CbsCmnGnbSmucppc field to given value.

### HasCbsCmnGnbSmucppc

`func (o *BiosPolicyAllOf) HasCbsCmnGnbSmucppc() bool`

HasCbsCmnGnbSmucppc returns a boolean if a field has been set.

### GetCbsCmnMemCtrlBankGroupSwapDdr4

`func (o *BiosPolicyAllOf) GetCbsCmnMemCtrlBankGroupSwapDdr4() string`

GetCbsCmnMemCtrlBankGroupSwapDdr4 returns the CbsCmnMemCtrlBankGroupSwapDdr4 field if non-nil, zero value otherwise.

### GetCbsCmnMemCtrlBankGroupSwapDdr4Ok

`func (o *BiosPolicyAllOf) GetCbsCmnMemCtrlBankGroupSwapDdr4Ok() (*string, bool)`

GetCbsCmnMemCtrlBankGroupSwapDdr4Ok returns a tuple with the CbsCmnMemCtrlBankGroupSwapDdr4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnMemCtrlBankGroupSwapDdr4

`func (o *BiosPolicyAllOf) SetCbsCmnMemCtrlBankGroupSwapDdr4(v string)`

SetCbsCmnMemCtrlBankGroupSwapDdr4 sets CbsCmnMemCtrlBankGroupSwapDdr4 field to given value.

### HasCbsCmnMemCtrlBankGroupSwapDdr4

`func (o *BiosPolicyAllOf) HasCbsCmnMemCtrlBankGroupSwapDdr4() bool`

HasCbsCmnMemCtrlBankGroupSwapDdr4 returns a boolean if a field has been set.

### GetCbsCmnMemMapBankInterleaveDdr4

`func (o *BiosPolicyAllOf) GetCbsCmnMemMapBankInterleaveDdr4() string`

GetCbsCmnMemMapBankInterleaveDdr4 returns the CbsCmnMemMapBankInterleaveDdr4 field if non-nil, zero value otherwise.

### GetCbsCmnMemMapBankInterleaveDdr4Ok

`func (o *BiosPolicyAllOf) GetCbsCmnMemMapBankInterleaveDdr4Ok() (*string, bool)`

GetCbsCmnMemMapBankInterleaveDdr4Ok returns a tuple with the CbsCmnMemMapBankInterleaveDdr4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmnMemMapBankInterleaveDdr4

`func (o *BiosPolicyAllOf) SetCbsCmnMemMapBankInterleaveDdr4(v string)`

SetCbsCmnMemMapBankInterleaveDdr4 sets CbsCmnMemMapBankInterleaveDdr4 field to given value.

### HasCbsCmnMemMapBankInterleaveDdr4

`func (o *BiosPolicyAllOf) HasCbsCmnMemMapBankInterleaveDdr4() bool`

HasCbsCmnMemMapBankInterleaveDdr4 returns a boolean if a field has been set.

### GetCbsCmncTdpCtl

`func (o *BiosPolicyAllOf) GetCbsCmncTdpCtl() string`

GetCbsCmncTdpCtl returns the CbsCmncTdpCtl field if non-nil, zero value otherwise.

### GetCbsCmncTdpCtlOk

`func (o *BiosPolicyAllOf) GetCbsCmncTdpCtlOk() (*string, bool)`

GetCbsCmncTdpCtlOk returns a tuple with the CbsCmncTdpCtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCmncTdpCtl

`func (o *BiosPolicyAllOf) SetCbsCmncTdpCtl(v string)`

SetCbsCmncTdpCtl sets CbsCmncTdpCtl field to given value.

### HasCbsCmncTdpCtl

`func (o *BiosPolicyAllOf) HasCbsCmncTdpCtl() bool`

HasCbsCmncTdpCtl returns a boolean if a field has been set.

### GetCbsCpuCcdCtrlSsp

`func (o *BiosPolicyAllOf) GetCbsCpuCcdCtrlSsp() string`

GetCbsCpuCcdCtrlSsp returns the CbsCpuCcdCtrlSsp field if non-nil, zero value otherwise.

### GetCbsCpuCcdCtrlSspOk

`func (o *BiosPolicyAllOf) GetCbsCpuCcdCtrlSspOk() (*string, bool)`

GetCbsCpuCcdCtrlSspOk returns a tuple with the CbsCpuCcdCtrlSsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCpuCcdCtrlSsp

`func (o *BiosPolicyAllOf) SetCbsCpuCcdCtrlSsp(v string)`

SetCbsCpuCcdCtrlSsp sets CbsCpuCcdCtrlSsp field to given value.

### HasCbsCpuCcdCtrlSsp

`func (o *BiosPolicyAllOf) HasCbsCpuCcdCtrlSsp() bool`

HasCbsCpuCcdCtrlSsp returns a boolean if a field has been set.

### GetCbsCpuCoreCtrl

`func (o *BiosPolicyAllOf) GetCbsCpuCoreCtrl() string`

GetCbsCpuCoreCtrl returns the CbsCpuCoreCtrl field if non-nil, zero value otherwise.

### GetCbsCpuCoreCtrlOk

`func (o *BiosPolicyAllOf) GetCbsCpuCoreCtrlOk() (*string, bool)`

GetCbsCpuCoreCtrlOk returns a tuple with the CbsCpuCoreCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCpuCoreCtrl

`func (o *BiosPolicyAllOf) SetCbsCpuCoreCtrl(v string)`

SetCbsCpuCoreCtrl sets CbsCpuCoreCtrl field to given value.

### HasCbsCpuCoreCtrl

`func (o *BiosPolicyAllOf) HasCbsCpuCoreCtrl() bool`

HasCbsCpuCoreCtrl returns a boolean if a field has been set.

### GetCbsCpuSmtCtrl

`func (o *BiosPolicyAllOf) GetCbsCpuSmtCtrl() string`

GetCbsCpuSmtCtrl returns the CbsCpuSmtCtrl field if non-nil, zero value otherwise.

### GetCbsCpuSmtCtrlOk

`func (o *BiosPolicyAllOf) GetCbsCpuSmtCtrlOk() (*string, bool)`

GetCbsCpuSmtCtrlOk returns a tuple with the CbsCpuSmtCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsCpuSmtCtrl

`func (o *BiosPolicyAllOf) SetCbsCpuSmtCtrl(v string)`

SetCbsCpuSmtCtrl sets CbsCpuSmtCtrl field to given value.

### HasCbsCpuSmtCtrl

`func (o *BiosPolicyAllOf) HasCbsCpuSmtCtrl() bool`

HasCbsCpuSmtCtrl returns a boolean if a field has been set.

### GetCbsDbgCpuSnpMemCover

`func (o *BiosPolicyAllOf) GetCbsDbgCpuSnpMemCover() string`

GetCbsDbgCpuSnpMemCover returns the CbsDbgCpuSnpMemCover field if non-nil, zero value otherwise.

### GetCbsDbgCpuSnpMemCoverOk

`func (o *BiosPolicyAllOf) GetCbsDbgCpuSnpMemCoverOk() (*string, bool)`

GetCbsDbgCpuSnpMemCoverOk returns a tuple with the CbsDbgCpuSnpMemCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDbgCpuSnpMemCover

`func (o *BiosPolicyAllOf) SetCbsDbgCpuSnpMemCover(v string)`

SetCbsDbgCpuSnpMemCover sets CbsDbgCpuSnpMemCover field to given value.

### HasCbsDbgCpuSnpMemCover

`func (o *BiosPolicyAllOf) HasCbsDbgCpuSnpMemCover() bool`

HasCbsDbgCpuSnpMemCover returns a boolean if a field has been set.

### GetCbsDbgCpuSnpMemSizeCover

`func (o *BiosPolicyAllOf) GetCbsDbgCpuSnpMemSizeCover() string`

GetCbsDbgCpuSnpMemSizeCover returns the CbsDbgCpuSnpMemSizeCover field if non-nil, zero value otherwise.

### GetCbsDbgCpuSnpMemSizeCoverOk

`func (o *BiosPolicyAllOf) GetCbsDbgCpuSnpMemSizeCoverOk() (*string, bool)`

GetCbsDbgCpuSnpMemSizeCoverOk returns a tuple with the CbsDbgCpuSnpMemSizeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDbgCpuSnpMemSizeCover

`func (o *BiosPolicyAllOf) SetCbsDbgCpuSnpMemSizeCover(v string)`

SetCbsDbgCpuSnpMemSizeCover sets CbsDbgCpuSnpMemSizeCover field to given value.

### HasCbsDbgCpuSnpMemSizeCover

`func (o *BiosPolicyAllOf) HasCbsDbgCpuSnpMemSizeCover() bool`

HasCbsDbgCpuSnpMemSizeCover returns a boolean if a field has been set.

### GetCbsDfCmnAcpiSratL3numa

`func (o *BiosPolicyAllOf) GetCbsDfCmnAcpiSratL3numa() string`

GetCbsDfCmnAcpiSratL3numa returns the CbsDfCmnAcpiSratL3numa field if non-nil, zero value otherwise.

### GetCbsDfCmnAcpiSratL3numaOk

`func (o *BiosPolicyAllOf) GetCbsDfCmnAcpiSratL3numaOk() (*string, bool)`

GetCbsDfCmnAcpiSratL3numaOk returns a tuple with the CbsDfCmnAcpiSratL3numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDfCmnAcpiSratL3numa

`func (o *BiosPolicyAllOf) SetCbsDfCmnAcpiSratL3numa(v string)`

SetCbsDfCmnAcpiSratL3numa sets CbsDfCmnAcpiSratL3numa field to given value.

### HasCbsDfCmnAcpiSratL3numa

`func (o *BiosPolicyAllOf) HasCbsDfCmnAcpiSratL3numa() bool`

HasCbsDfCmnAcpiSratL3numa returns a boolean if a field has been set.

### GetCbsDfCmnDramNps

`func (o *BiosPolicyAllOf) GetCbsDfCmnDramNps() string`

GetCbsDfCmnDramNps returns the CbsDfCmnDramNps field if non-nil, zero value otherwise.

### GetCbsDfCmnDramNpsOk

`func (o *BiosPolicyAllOf) GetCbsDfCmnDramNpsOk() (*string, bool)`

GetCbsDfCmnDramNpsOk returns a tuple with the CbsDfCmnDramNps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDfCmnDramNps

`func (o *BiosPolicyAllOf) SetCbsDfCmnDramNps(v string)`

SetCbsDfCmnDramNps sets CbsDfCmnDramNps field to given value.

### HasCbsDfCmnDramNps

`func (o *BiosPolicyAllOf) HasCbsDfCmnDramNps() bool`

HasCbsDfCmnDramNps returns a boolean if a field has been set.

### GetCbsDfCmnMemIntlv

`func (o *BiosPolicyAllOf) GetCbsDfCmnMemIntlv() string`

GetCbsDfCmnMemIntlv returns the CbsDfCmnMemIntlv field if non-nil, zero value otherwise.

### GetCbsDfCmnMemIntlvOk

`func (o *BiosPolicyAllOf) GetCbsDfCmnMemIntlvOk() (*string, bool)`

GetCbsDfCmnMemIntlvOk returns a tuple with the CbsDfCmnMemIntlv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDfCmnMemIntlv

`func (o *BiosPolicyAllOf) SetCbsDfCmnMemIntlv(v string)`

SetCbsDfCmnMemIntlv sets CbsDfCmnMemIntlv field to given value.

### HasCbsDfCmnMemIntlv

`func (o *BiosPolicyAllOf) HasCbsDfCmnMemIntlv() bool`

HasCbsDfCmnMemIntlv returns a boolean if a field has been set.

### GetCbsDfCmnMemIntlvSize

`func (o *BiosPolicyAllOf) GetCbsDfCmnMemIntlvSize() string`

GetCbsDfCmnMemIntlvSize returns the CbsDfCmnMemIntlvSize field if non-nil, zero value otherwise.

### GetCbsDfCmnMemIntlvSizeOk

`func (o *BiosPolicyAllOf) GetCbsDfCmnMemIntlvSizeOk() (*string, bool)`

GetCbsDfCmnMemIntlvSizeOk returns a tuple with the CbsDfCmnMemIntlvSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsDfCmnMemIntlvSize

`func (o *BiosPolicyAllOf) SetCbsDfCmnMemIntlvSize(v string)`

SetCbsDfCmnMemIntlvSize sets CbsDfCmnMemIntlvSize field to given value.

### HasCbsDfCmnMemIntlvSize

`func (o *BiosPolicyAllOf) HasCbsDfCmnMemIntlvSize() bool`

HasCbsDfCmnMemIntlvSize returns a boolean if a field has been set.

### GetCbsSevSnpSupport

`func (o *BiosPolicyAllOf) GetCbsSevSnpSupport() string`

GetCbsSevSnpSupport returns the CbsSevSnpSupport field if non-nil, zero value otherwise.

### GetCbsSevSnpSupportOk

`func (o *BiosPolicyAllOf) GetCbsSevSnpSupportOk() (*string, bool)`

GetCbsSevSnpSupportOk returns a tuple with the CbsSevSnpSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbsSevSnpSupport

`func (o *BiosPolicyAllOf) SetCbsSevSnpSupport(v string)`

SetCbsSevSnpSupport sets CbsSevSnpSupport field to given value.

### HasCbsSevSnpSupport

`func (o *BiosPolicyAllOf) HasCbsSevSnpSupport() bool`

HasCbsSevSnpSupport returns a boolean if a field has been set.

### GetCdnEnable

`func (o *BiosPolicyAllOf) GetCdnEnable() string`

GetCdnEnable returns the CdnEnable field if non-nil, zero value otherwise.

### GetCdnEnableOk

`func (o *BiosPolicyAllOf) GetCdnEnableOk() (*string, bool)`

GetCdnEnableOk returns a tuple with the CdnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnEnable

`func (o *BiosPolicyAllOf) SetCdnEnable(v string)`

SetCdnEnable sets CdnEnable field to given value.

### HasCdnEnable

`func (o *BiosPolicyAllOf) HasCdnEnable() bool`

HasCdnEnable returns a boolean if a field has been set.

### GetCdnSupport

`func (o *BiosPolicyAllOf) GetCdnSupport() string`

GetCdnSupport returns the CdnSupport field if non-nil, zero value otherwise.

### GetCdnSupportOk

`func (o *BiosPolicyAllOf) GetCdnSupportOk() (*string, bool)`

GetCdnSupportOk returns a tuple with the CdnSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnSupport

`func (o *BiosPolicyAllOf) SetCdnSupport(v string)`

SetCdnSupport sets CdnSupport field to given value.

### HasCdnSupport

`func (o *BiosPolicyAllOf) HasCdnSupport() bool`

HasCdnSupport returns a boolean if a field has been set.

### GetChannelInterLeave

`func (o *BiosPolicyAllOf) GetChannelInterLeave() string`

GetChannelInterLeave returns the ChannelInterLeave field if non-nil, zero value otherwise.

### GetChannelInterLeaveOk

`func (o *BiosPolicyAllOf) GetChannelInterLeaveOk() (*string, bool)`

GetChannelInterLeaveOk returns a tuple with the ChannelInterLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelInterLeave

`func (o *BiosPolicyAllOf) SetChannelInterLeave(v string)`

SetChannelInterLeave sets ChannelInterLeave field to given value.

### HasChannelInterLeave

`func (o *BiosPolicyAllOf) HasChannelInterLeave() bool`

HasChannelInterLeave returns a boolean if a field has been set.

### GetCiscoAdaptiveMemTraining

`func (o *BiosPolicyAllOf) GetCiscoAdaptiveMemTraining() string`

GetCiscoAdaptiveMemTraining returns the CiscoAdaptiveMemTraining field if non-nil, zero value otherwise.

### GetCiscoAdaptiveMemTrainingOk

`func (o *BiosPolicyAllOf) GetCiscoAdaptiveMemTrainingOk() (*string, bool)`

GetCiscoAdaptiveMemTrainingOk returns a tuple with the CiscoAdaptiveMemTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoAdaptiveMemTraining

`func (o *BiosPolicyAllOf) SetCiscoAdaptiveMemTraining(v string)`

SetCiscoAdaptiveMemTraining sets CiscoAdaptiveMemTraining field to given value.

### HasCiscoAdaptiveMemTraining

`func (o *BiosPolicyAllOf) HasCiscoAdaptiveMemTraining() bool`

HasCiscoAdaptiveMemTraining returns a boolean if a field has been set.

### GetCiscoDebugLevel

`func (o *BiosPolicyAllOf) GetCiscoDebugLevel() string`

GetCiscoDebugLevel returns the CiscoDebugLevel field if non-nil, zero value otherwise.

### GetCiscoDebugLevelOk

`func (o *BiosPolicyAllOf) GetCiscoDebugLevelOk() (*string, bool)`

GetCiscoDebugLevelOk returns a tuple with the CiscoDebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoDebugLevel

`func (o *BiosPolicyAllOf) SetCiscoDebugLevel(v string)`

SetCiscoDebugLevel sets CiscoDebugLevel field to given value.

### HasCiscoDebugLevel

`func (o *BiosPolicyAllOf) HasCiscoDebugLevel() bool`

HasCiscoDebugLevel returns a boolean if a field has been set.

### GetCiscoOpromLaunchOptimization

`func (o *BiosPolicyAllOf) GetCiscoOpromLaunchOptimization() string`

GetCiscoOpromLaunchOptimization returns the CiscoOpromLaunchOptimization field if non-nil, zero value otherwise.

### GetCiscoOpromLaunchOptimizationOk

`func (o *BiosPolicyAllOf) GetCiscoOpromLaunchOptimizationOk() (*string, bool)`

GetCiscoOpromLaunchOptimizationOk returns a tuple with the CiscoOpromLaunchOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoOpromLaunchOptimization

`func (o *BiosPolicyAllOf) SetCiscoOpromLaunchOptimization(v string)`

SetCiscoOpromLaunchOptimization sets CiscoOpromLaunchOptimization field to given value.

### HasCiscoOpromLaunchOptimization

`func (o *BiosPolicyAllOf) HasCiscoOpromLaunchOptimization() bool`

HasCiscoOpromLaunchOptimization returns a boolean if a field has been set.

### GetCiscoXgmiMaxSpeed

`func (o *BiosPolicyAllOf) GetCiscoXgmiMaxSpeed() string`

GetCiscoXgmiMaxSpeed returns the CiscoXgmiMaxSpeed field if non-nil, zero value otherwise.

### GetCiscoXgmiMaxSpeedOk

`func (o *BiosPolicyAllOf) GetCiscoXgmiMaxSpeedOk() (*string, bool)`

GetCiscoXgmiMaxSpeedOk returns a tuple with the CiscoXgmiMaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoXgmiMaxSpeed

`func (o *BiosPolicyAllOf) SetCiscoXgmiMaxSpeed(v string)`

SetCiscoXgmiMaxSpeed sets CiscoXgmiMaxSpeed field to given value.

### HasCiscoXgmiMaxSpeed

`func (o *BiosPolicyAllOf) HasCiscoXgmiMaxSpeed() bool`

HasCiscoXgmiMaxSpeed returns a boolean if a field has been set.

### GetCkeLowPolicy

`func (o *BiosPolicyAllOf) GetCkeLowPolicy() string`

GetCkeLowPolicy returns the CkeLowPolicy field if non-nil, zero value otherwise.

### GetCkeLowPolicyOk

`func (o *BiosPolicyAllOf) GetCkeLowPolicyOk() (*string, bool)`

GetCkeLowPolicyOk returns a tuple with the CkeLowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCkeLowPolicy

`func (o *BiosPolicyAllOf) SetCkeLowPolicy(v string)`

SetCkeLowPolicy sets CkeLowPolicy field to given value.

### HasCkeLowPolicy

`func (o *BiosPolicyAllOf) HasCkeLowPolicy() bool`

HasCkeLowPolicy returns a boolean if a field has been set.

### GetClosedLoopThermThrotl

`func (o *BiosPolicyAllOf) GetClosedLoopThermThrotl() string`

GetClosedLoopThermThrotl returns the ClosedLoopThermThrotl field if non-nil, zero value otherwise.

### GetClosedLoopThermThrotlOk

`func (o *BiosPolicyAllOf) GetClosedLoopThermThrotlOk() (*string, bool)`

GetClosedLoopThermThrotlOk returns a tuple with the ClosedLoopThermThrotl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopThermThrotl

`func (o *BiosPolicyAllOf) SetClosedLoopThermThrotl(v string)`

SetClosedLoopThermThrotl sets ClosedLoopThermThrotl field to given value.

### HasClosedLoopThermThrotl

`func (o *BiosPolicyAllOf) HasClosedLoopThermThrotl() bool`

HasClosedLoopThermThrotl returns a boolean if a field has been set.

### GetCmciEnable

`func (o *BiosPolicyAllOf) GetCmciEnable() string`

GetCmciEnable returns the CmciEnable field if non-nil, zero value otherwise.

### GetCmciEnableOk

`func (o *BiosPolicyAllOf) GetCmciEnableOk() (*string, bool)`

GetCmciEnableOk returns a tuple with the CmciEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmciEnable

`func (o *BiosPolicyAllOf) SetCmciEnable(v string)`

SetCmciEnable sets CmciEnable field to given value.

### HasCmciEnable

`func (o *BiosPolicyAllOf) HasCmciEnable() bool`

HasCmciEnable returns a boolean if a field has been set.

### GetConfigTdp

`func (o *BiosPolicyAllOf) GetConfigTdp() string`

GetConfigTdp returns the ConfigTdp field if non-nil, zero value otherwise.

### GetConfigTdpOk

`func (o *BiosPolicyAllOf) GetConfigTdpOk() (*string, bool)`

GetConfigTdpOk returns a tuple with the ConfigTdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTdp

`func (o *BiosPolicyAllOf) SetConfigTdp(v string)`

SetConfigTdp sets ConfigTdp field to given value.

### HasConfigTdp

`func (o *BiosPolicyAllOf) HasConfigTdp() bool`

HasConfigTdp returns a boolean if a field has been set.

### GetConfigTdpLevel

`func (o *BiosPolicyAllOf) GetConfigTdpLevel() string`

GetConfigTdpLevel returns the ConfigTdpLevel field if non-nil, zero value otherwise.

### GetConfigTdpLevelOk

`func (o *BiosPolicyAllOf) GetConfigTdpLevelOk() (*string, bool)`

GetConfigTdpLevelOk returns a tuple with the ConfigTdpLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTdpLevel

`func (o *BiosPolicyAllOf) SetConfigTdpLevel(v string)`

SetConfigTdpLevel sets ConfigTdpLevel field to given value.

### HasConfigTdpLevel

`func (o *BiosPolicyAllOf) HasConfigTdpLevel() bool`

HasConfigTdpLevel returns a boolean if a field has been set.

### GetConsoleRedirection

`func (o *BiosPolicyAllOf) GetConsoleRedirection() string`

GetConsoleRedirection returns the ConsoleRedirection field if non-nil, zero value otherwise.

### GetConsoleRedirectionOk

`func (o *BiosPolicyAllOf) GetConsoleRedirectionOk() (*string, bool)`

GetConsoleRedirectionOk returns a tuple with the ConsoleRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleRedirection

`func (o *BiosPolicyAllOf) SetConsoleRedirection(v string)`

SetConsoleRedirection sets ConsoleRedirection field to given value.

### HasConsoleRedirection

`func (o *BiosPolicyAllOf) HasConsoleRedirection() bool`

HasConsoleRedirection returns a boolean if a field has been set.

### GetCoreMultiProcessing

`func (o *BiosPolicyAllOf) GetCoreMultiProcessing() string`

GetCoreMultiProcessing returns the CoreMultiProcessing field if non-nil, zero value otherwise.

### GetCoreMultiProcessingOk

`func (o *BiosPolicyAllOf) GetCoreMultiProcessingOk() (*string, bool)`

GetCoreMultiProcessingOk returns a tuple with the CoreMultiProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreMultiProcessing

`func (o *BiosPolicyAllOf) SetCoreMultiProcessing(v string)`

SetCoreMultiProcessing sets CoreMultiProcessing field to given value.

### HasCoreMultiProcessing

`func (o *BiosPolicyAllOf) HasCoreMultiProcessing() bool`

HasCoreMultiProcessing returns a boolean if a field has been set.

### GetCpuEnergyPerformance

`func (o *BiosPolicyAllOf) GetCpuEnergyPerformance() string`

GetCpuEnergyPerformance returns the CpuEnergyPerformance field if non-nil, zero value otherwise.

### GetCpuEnergyPerformanceOk

`func (o *BiosPolicyAllOf) GetCpuEnergyPerformanceOk() (*string, bool)`

GetCpuEnergyPerformanceOk returns a tuple with the CpuEnergyPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnergyPerformance

`func (o *BiosPolicyAllOf) SetCpuEnergyPerformance(v string)`

SetCpuEnergyPerformance sets CpuEnergyPerformance field to given value.

### HasCpuEnergyPerformance

`func (o *BiosPolicyAllOf) HasCpuEnergyPerformance() bool`

HasCpuEnergyPerformance returns a boolean if a field has been set.

### GetCpuFrequencyFloor

`func (o *BiosPolicyAllOf) GetCpuFrequencyFloor() string`

GetCpuFrequencyFloor returns the CpuFrequencyFloor field if non-nil, zero value otherwise.

### GetCpuFrequencyFloorOk

`func (o *BiosPolicyAllOf) GetCpuFrequencyFloorOk() (*string, bool)`

GetCpuFrequencyFloorOk returns a tuple with the CpuFrequencyFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyFloor

`func (o *BiosPolicyAllOf) SetCpuFrequencyFloor(v string)`

SetCpuFrequencyFloor sets CpuFrequencyFloor field to given value.

### HasCpuFrequencyFloor

`func (o *BiosPolicyAllOf) HasCpuFrequencyFloor() bool`

HasCpuFrequencyFloor returns a boolean if a field has been set.

### GetCpuPerformance

`func (o *BiosPolicyAllOf) GetCpuPerformance() string`

GetCpuPerformance returns the CpuPerformance field if non-nil, zero value otherwise.

### GetCpuPerformanceOk

`func (o *BiosPolicyAllOf) GetCpuPerformanceOk() (*string, bool)`

GetCpuPerformanceOk returns a tuple with the CpuPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPerformance

`func (o *BiosPolicyAllOf) SetCpuPerformance(v string)`

SetCpuPerformance sets CpuPerformance field to given value.

### HasCpuPerformance

`func (o *BiosPolicyAllOf) HasCpuPerformance() bool`

HasCpuPerformance returns a boolean if a field has been set.

### GetCpuPowerManagement

`func (o *BiosPolicyAllOf) GetCpuPowerManagement() string`

GetCpuPowerManagement returns the CpuPowerManagement field if non-nil, zero value otherwise.

### GetCpuPowerManagementOk

`func (o *BiosPolicyAllOf) GetCpuPowerManagementOk() (*string, bool)`

GetCpuPowerManagementOk returns a tuple with the CpuPowerManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPowerManagement

`func (o *BiosPolicyAllOf) SetCpuPowerManagement(v string)`

SetCpuPowerManagement sets CpuPowerManagement field to given value.

### HasCpuPowerManagement

`func (o *BiosPolicyAllOf) HasCpuPowerManagement() bool`

HasCpuPowerManagement returns a boolean if a field has been set.

### GetCrQos

`func (o *BiosPolicyAllOf) GetCrQos() string`

GetCrQos returns the CrQos field if non-nil, zero value otherwise.

### GetCrQosOk

`func (o *BiosPolicyAllOf) GetCrQosOk() (*string, bool)`

GetCrQosOk returns a tuple with the CrQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrQos

`func (o *BiosPolicyAllOf) SetCrQos(v string)`

SetCrQos sets CrQos field to given value.

### HasCrQos

`func (o *BiosPolicyAllOf) HasCrQos() bool`

HasCrQos returns a boolean if a field has been set.

### GetCrfastgoConfig

`func (o *BiosPolicyAllOf) GetCrfastgoConfig() string`

GetCrfastgoConfig returns the CrfastgoConfig field if non-nil, zero value otherwise.

### GetCrfastgoConfigOk

`func (o *BiosPolicyAllOf) GetCrfastgoConfigOk() (*string, bool)`

GetCrfastgoConfigOk returns a tuple with the CrfastgoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrfastgoConfig

`func (o *BiosPolicyAllOf) SetCrfastgoConfig(v string)`

SetCrfastgoConfig sets CrfastgoConfig field to given value.

### HasCrfastgoConfig

`func (o *BiosPolicyAllOf) HasCrfastgoConfig() bool`

HasCrfastgoConfig returns a boolean if a field has been set.

### GetDcpmmFirmwareDowngrade

`func (o *BiosPolicyAllOf) GetDcpmmFirmwareDowngrade() string`

GetDcpmmFirmwareDowngrade returns the DcpmmFirmwareDowngrade field if non-nil, zero value otherwise.

### GetDcpmmFirmwareDowngradeOk

`func (o *BiosPolicyAllOf) GetDcpmmFirmwareDowngradeOk() (*string, bool)`

GetDcpmmFirmwareDowngradeOk returns a tuple with the DcpmmFirmwareDowngrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcpmmFirmwareDowngrade

`func (o *BiosPolicyAllOf) SetDcpmmFirmwareDowngrade(v string)`

SetDcpmmFirmwareDowngrade sets DcpmmFirmwareDowngrade field to given value.

### HasDcpmmFirmwareDowngrade

`func (o *BiosPolicyAllOf) HasDcpmmFirmwareDowngrade() bool`

HasDcpmmFirmwareDowngrade returns a boolean if a field has been set.

### GetDemandScrub

`func (o *BiosPolicyAllOf) GetDemandScrub() string`

GetDemandScrub returns the DemandScrub field if non-nil, zero value otherwise.

### GetDemandScrubOk

`func (o *BiosPolicyAllOf) GetDemandScrubOk() (*string, bool)`

GetDemandScrubOk returns a tuple with the DemandScrub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandScrub

`func (o *BiosPolicyAllOf) SetDemandScrub(v string)`

SetDemandScrub sets DemandScrub field to given value.

### HasDemandScrub

`func (o *BiosPolicyAllOf) HasDemandScrub() bool`

HasDemandScrub returns a boolean if a field has been set.

### GetDirectCacheAccess

`func (o *BiosPolicyAllOf) GetDirectCacheAccess() string`

GetDirectCacheAccess returns the DirectCacheAccess field if non-nil, zero value otherwise.

### GetDirectCacheAccessOk

`func (o *BiosPolicyAllOf) GetDirectCacheAccessOk() (*string, bool)`

GetDirectCacheAccessOk returns a tuple with the DirectCacheAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCacheAccess

`func (o *BiosPolicyAllOf) SetDirectCacheAccess(v string)`

SetDirectCacheAccess sets DirectCacheAccess field to given value.

### HasDirectCacheAccess

`func (o *BiosPolicyAllOf) HasDirectCacheAccess() bool`

HasDirectCacheAccess returns a boolean if a field has been set.

### GetDramClockThrottling

`func (o *BiosPolicyAllOf) GetDramClockThrottling() string`

GetDramClockThrottling returns the DramClockThrottling field if non-nil, zero value otherwise.

### GetDramClockThrottlingOk

`func (o *BiosPolicyAllOf) GetDramClockThrottlingOk() (*string, bool)`

GetDramClockThrottlingOk returns a tuple with the DramClockThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDramClockThrottling

`func (o *BiosPolicyAllOf) SetDramClockThrottling(v string)`

SetDramClockThrottling sets DramClockThrottling field to given value.

### HasDramClockThrottling

`func (o *BiosPolicyAllOf) HasDramClockThrottling() bool`

HasDramClockThrottling returns a boolean if a field has been set.

### GetDramRefreshRate

`func (o *BiosPolicyAllOf) GetDramRefreshRate() string`

GetDramRefreshRate returns the DramRefreshRate field if non-nil, zero value otherwise.

### GetDramRefreshRateOk

`func (o *BiosPolicyAllOf) GetDramRefreshRateOk() (*string, bool)`

GetDramRefreshRateOk returns a tuple with the DramRefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDramRefreshRate

`func (o *BiosPolicyAllOf) SetDramRefreshRate(v string)`

SetDramRefreshRate sets DramRefreshRate field to given value.

### HasDramRefreshRate

`func (o *BiosPolicyAllOf) HasDramRefreshRate() bool`

HasDramRefreshRate returns a boolean if a field has been set.

### GetDramSwThermalThrottling

`func (o *BiosPolicyAllOf) GetDramSwThermalThrottling() string`

GetDramSwThermalThrottling returns the DramSwThermalThrottling field if non-nil, zero value otherwise.

### GetDramSwThermalThrottlingOk

`func (o *BiosPolicyAllOf) GetDramSwThermalThrottlingOk() (*string, bool)`

GetDramSwThermalThrottlingOk returns a tuple with the DramSwThermalThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDramSwThermalThrottling

`func (o *BiosPolicyAllOf) SetDramSwThermalThrottling(v string)`

SetDramSwThermalThrottling sets DramSwThermalThrottling field to given value.

### HasDramSwThermalThrottling

`func (o *BiosPolicyAllOf) HasDramSwThermalThrottling() bool`

HasDramSwThermalThrottling returns a boolean if a field has been set.

### GetEadrSupport

`func (o *BiosPolicyAllOf) GetEadrSupport() string`

GetEadrSupport returns the EadrSupport field if non-nil, zero value otherwise.

### GetEadrSupportOk

`func (o *BiosPolicyAllOf) GetEadrSupportOk() (*string, bool)`

GetEadrSupportOk returns a tuple with the EadrSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEadrSupport

`func (o *BiosPolicyAllOf) SetEadrSupport(v string)`

SetEadrSupport sets EadrSupport field to given value.

### HasEadrSupport

`func (o *BiosPolicyAllOf) HasEadrSupport() bool`

HasEadrSupport returns a boolean if a field has been set.

### GetEdpcEn

`func (o *BiosPolicyAllOf) GetEdpcEn() string`

GetEdpcEn returns the EdpcEn field if non-nil, zero value otherwise.

### GetEdpcEnOk

`func (o *BiosPolicyAllOf) GetEdpcEnOk() (*string, bool)`

GetEdpcEnOk returns a tuple with the EdpcEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdpcEn

`func (o *BiosPolicyAllOf) SetEdpcEn(v string)`

SetEdpcEn sets EdpcEn field to given value.

### HasEdpcEn

`func (o *BiosPolicyAllOf) HasEdpcEn() bool`

HasEdpcEn returns a boolean if a field has been set.

### GetEnableClockSpreadSpec

`func (o *BiosPolicyAllOf) GetEnableClockSpreadSpec() string`

GetEnableClockSpreadSpec returns the EnableClockSpreadSpec field if non-nil, zero value otherwise.

### GetEnableClockSpreadSpecOk

`func (o *BiosPolicyAllOf) GetEnableClockSpreadSpecOk() (*string, bool)`

GetEnableClockSpreadSpecOk returns a tuple with the EnableClockSpreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClockSpreadSpec

`func (o *BiosPolicyAllOf) SetEnableClockSpreadSpec(v string)`

SetEnableClockSpreadSpec sets EnableClockSpreadSpec field to given value.

### HasEnableClockSpreadSpec

`func (o *BiosPolicyAllOf) HasEnableClockSpreadSpec() bool`

HasEnableClockSpreadSpec returns a boolean if a field has been set.

### GetEnableMktme

`func (o *BiosPolicyAllOf) GetEnableMktme() string`

GetEnableMktme returns the EnableMktme field if non-nil, zero value otherwise.

### GetEnableMktmeOk

`func (o *BiosPolicyAllOf) GetEnableMktmeOk() (*string, bool)`

GetEnableMktmeOk returns a tuple with the EnableMktme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMktme

`func (o *BiosPolicyAllOf) SetEnableMktme(v string)`

SetEnableMktme sets EnableMktme field to given value.

### HasEnableMktme

`func (o *BiosPolicyAllOf) HasEnableMktme() bool`

HasEnableMktme returns a boolean if a field has been set.

### GetEnableSgx

`func (o *BiosPolicyAllOf) GetEnableSgx() string`

GetEnableSgx returns the EnableSgx field if non-nil, zero value otherwise.

### GetEnableSgxOk

`func (o *BiosPolicyAllOf) GetEnableSgxOk() (*string, bool)`

GetEnableSgxOk returns a tuple with the EnableSgx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSgx

`func (o *BiosPolicyAllOf) SetEnableSgx(v string)`

SetEnableSgx sets EnableSgx field to given value.

### HasEnableSgx

`func (o *BiosPolicyAllOf) HasEnableSgx() bool`

HasEnableSgx returns a boolean if a field has been set.

### GetEnableTme

`func (o *BiosPolicyAllOf) GetEnableTme() string`

GetEnableTme returns the EnableTme field if non-nil, zero value otherwise.

### GetEnableTmeOk

`func (o *BiosPolicyAllOf) GetEnableTmeOk() (*string, bool)`

GetEnableTmeOk returns a tuple with the EnableTme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTme

`func (o *BiosPolicyAllOf) SetEnableTme(v string)`

SetEnableTme sets EnableTme field to given value.

### HasEnableTme

`func (o *BiosPolicyAllOf) HasEnableTme() bool`

HasEnableTme returns a boolean if a field has been set.

### GetEnergyEfficientTurbo

`func (o *BiosPolicyAllOf) GetEnergyEfficientTurbo() string`

GetEnergyEfficientTurbo returns the EnergyEfficientTurbo field if non-nil, zero value otherwise.

### GetEnergyEfficientTurboOk

`func (o *BiosPolicyAllOf) GetEnergyEfficientTurboOk() (*string, bool)`

GetEnergyEfficientTurboOk returns a tuple with the EnergyEfficientTurbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficientTurbo

`func (o *BiosPolicyAllOf) SetEnergyEfficientTurbo(v string)`

SetEnergyEfficientTurbo sets EnergyEfficientTurbo field to given value.

### HasEnergyEfficientTurbo

`func (o *BiosPolicyAllOf) HasEnergyEfficientTurbo() bool`

HasEnergyEfficientTurbo returns a boolean if a field has been set.

### GetEngPerfTuning

`func (o *BiosPolicyAllOf) GetEngPerfTuning() string`

GetEngPerfTuning returns the EngPerfTuning field if non-nil, zero value otherwise.

### GetEngPerfTuningOk

`func (o *BiosPolicyAllOf) GetEngPerfTuningOk() (*string, bool)`

GetEngPerfTuningOk returns a tuple with the EngPerfTuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngPerfTuning

`func (o *BiosPolicyAllOf) SetEngPerfTuning(v string)`

SetEngPerfTuning sets EngPerfTuning field to given value.

### HasEngPerfTuning

`func (o *BiosPolicyAllOf) HasEngPerfTuning() bool`

HasEngPerfTuning returns a boolean if a field has been set.

### GetEnhancedIntelSpeedStepTech

`func (o *BiosPolicyAllOf) GetEnhancedIntelSpeedStepTech() string`

GetEnhancedIntelSpeedStepTech returns the EnhancedIntelSpeedStepTech field if non-nil, zero value otherwise.

### GetEnhancedIntelSpeedStepTechOk

`func (o *BiosPolicyAllOf) GetEnhancedIntelSpeedStepTechOk() (*string, bool)`

GetEnhancedIntelSpeedStepTechOk returns a tuple with the EnhancedIntelSpeedStepTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedIntelSpeedStepTech

`func (o *BiosPolicyAllOf) SetEnhancedIntelSpeedStepTech(v string)`

SetEnhancedIntelSpeedStepTech sets EnhancedIntelSpeedStepTech field to given value.

### HasEnhancedIntelSpeedStepTech

`func (o *BiosPolicyAllOf) HasEnhancedIntelSpeedStepTech() bool`

HasEnhancedIntelSpeedStepTech returns a boolean if a field has been set.

### GetEpochUpdate

`func (o *BiosPolicyAllOf) GetEpochUpdate() string`

GetEpochUpdate returns the EpochUpdate field if non-nil, zero value otherwise.

### GetEpochUpdateOk

`func (o *BiosPolicyAllOf) GetEpochUpdateOk() (*string, bool)`

GetEpochUpdateOk returns a tuple with the EpochUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochUpdate

`func (o *BiosPolicyAllOf) SetEpochUpdate(v string)`

SetEpochUpdate sets EpochUpdate field to given value.

### HasEpochUpdate

`func (o *BiosPolicyAllOf) HasEpochUpdate() bool`

HasEpochUpdate returns a boolean if a field has been set.

### GetEppEnable

`func (o *BiosPolicyAllOf) GetEppEnable() string`

GetEppEnable returns the EppEnable field if non-nil, zero value otherwise.

### GetEppEnableOk

`func (o *BiosPolicyAllOf) GetEppEnableOk() (*string, bool)`

GetEppEnableOk returns a tuple with the EppEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEppEnable

`func (o *BiosPolicyAllOf) SetEppEnable(v string)`

SetEppEnable sets EppEnable field to given value.

### HasEppEnable

`func (o *BiosPolicyAllOf) HasEppEnable() bool`

HasEppEnable returns a boolean if a field has been set.

### GetEppProfile

`func (o *BiosPolicyAllOf) GetEppProfile() string`

GetEppProfile returns the EppProfile field if non-nil, zero value otherwise.

### GetEppProfileOk

`func (o *BiosPolicyAllOf) GetEppProfileOk() (*string, bool)`

GetEppProfileOk returns a tuple with the EppProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEppProfile

`func (o *BiosPolicyAllOf) SetEppProfile(v string)`

SetEppProfile sets EppProfile field to given value.

### HasEppProfile

`func (o *BiosPolicyAllOf) HasEppProfile() bool`

HasEppProfile returns a boolean if a field has been set.

### GetExecuteDisableBit

`func (o *BiosPolicyAllOf) GetExecuteDisableBit() string`

GetExecuteDisableBit returns the ExecuteDisableBit field if non-nil, zero value otherwise.

### GetExecuteDisableBitOk

`func (o *BiosPolicyAllOf) GetExecuteDisableBitOk() (*string, bool)`

GetExecuteDisableBitOk returns a tuple with the ExecuteDisableBit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteDisableBit

`func (o *BiosPolicyAllOf) SetExecuteDisableBit(v string)`

SetExecuteDisableBit sets ExecuteDisableBit field to given value.

### HasExecuteDisableBit

`func (o *BiosPolicyAllOf) HasExecuteDisableBit() bool`

HasExecuteDisableBit returns a boolean if a field has been set.

### GetExtendedApic

`func (o *BiosPolicyAllOf) GetExtendedApic() string`

GetExtendedApic returns the ExtendedApic field if non-nil, zero value otherwise.

### GetExtendedApicOk

`func (o *BiosPolicyAllOf) GetExtendedApicOk() (*string, bool)`

GetExtendedApicOk returns a tuple with the ExtendedApic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedApic

`func (o *BiosPolicyAllOf) SetExtendedApic(v string)`

SetExtendedApic sets ExtendedApic field to given value.

### HasExtendedApic

`func (o *BiosPolicyAllOf) HasExtendedApic() bool`

HasExtendedApic returns a boolean if a field has been set.

### GetFlowControl

`func (o *BiosPolicyAllOf) GetFlowControl() string`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *BiosPolicyAllOf) GetFlowControlOk() (*string, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *BiosPolicyAllOf) SetFlowControl(v string)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *BiosPolicyAllOf) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetFrb2enable

`func (o *BiosPolicyAllOf) GetFrb2enable() string`

GetFrb2enable returns the Frb2enable field if non-nil, zero value otherwise.

### GetFrb2enableOk

`func (o *BiosPolicyAllOf) GetFrb2enableOk() (*string, bool)`

GetFrb2enableOk returns a tuple with the Frb2enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrb2enable

`func (o *BiosPolicyAllOf) SetFrb2enable(v string)`

SetFrb2enable sets Frb2enable field to given value.

### HasFrb2enable

`func (o *BiosPolicyAllOf) HasFrb2enable() bool`

HasFrb2enable returns a boolean if a field has been set.

### GetHardwarePrefetch

`func (o *BiosPolicyAllOf) GetHardwarePrefetch() string`

GetHardwarePrefetch returns the HardwarePrefetch field if non-nil, zero value otherwise.

### GetHardwarePrefetchOk

`func (o *BiosPolicyAllOf) GetHardwarePrefetchOk() (*string, bool)`

GetHardwarePrefetchOk returns a tuple with the HardwarePrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwarePrefetch

`func (o *BiosPolicyAllOf) SetHardwarePrefetch(v string)`

SetHardwarePrefetch sets HardwarePrefetch field to given value.

### HasHardwarePrefetch

`func (o *BiosPolicyAllOf) HasHardwarePrefetch() bool`

HasHardwarePrefetch returns a boolean if a field has been set.

### GetHwpmEnable

`func (o *BiosPolicyAllOf) GetHwpmEnable() string`

GetHwpmEnable returns the HwpmEnable field if non-nil, zero value otherwise.

### GetHwpmEnableOk

`func (o *BiosPolicyAllOf) GetHwpmEnableOk() (*string, bool)`

GetHwpmEnableOk returns a tuple with the HwpmEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwpmEnable

`func (o *BiosPolicyAllOf) SetHwpmEnable(v string)`

SetHwpmEnable sets HwpmEnable field to given value.

### HasHwpmEnable

`func (o *BiosPolicyAllOf) HasHwpmEnable() bool`

HasHwpmEnable returns a boolean if a field has been set.

### GetImcInterleave

`func (o *BiosPolicyAllOf) GetImcInterleave() string`

GetImcInterleave returns the ImcInterleave field if non-nil, zero value otherwise.

### GetImcInterleaveOk

`func (o *BiosPolicyAllOf) GetImcInterleaveOk() (*string, bool)`

GetImcInterleaveOk returns a tuple with the ImcInterleave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImcInterleave

`func (o *BiosPolicyAllOf) SetImcInterleave(v string)`

SetImcInterleave sets ImcInterleave field to given value.

### HasImcInterleave

`func (o *BiosPolicyAllOf) HasImcInterleave() bool`

HasImcInterleave returns a boolean if a field has been set.

### GetIntelDynamicSpeedSelect

`func (o *BiosPolicyAllOf) GetIntelDynamicSpeedSelect() string`

GetIntelDynamicSpeedSelect returns the IntelDynamicSpeedSelect field if non-nil, zero value otherwise.

### GetIntelDynamicSpeedSelectOk

`func (o *BiosPolicyAllOf) GetIntelDynamicSpeedSelectOk() (*string, bool)`

GetIntelDynamicSpeedSelectOk returns a tuple with the IntelDynamicSpeedSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelDynamicSpeedSelect

`func (o *BiosPolicyAllOf) SetIntelDynamicSpeedSelect(v string)`

SetIntelDynamicSpeedSelect sets IntelDynamicSpeedSelect field to given value.

### HasIntelDynamicSpeedSelect

`func (o *BiosPolicyAllOf) HasIntelDynamicSpeedSelect() bool`

HasIntelDynamicSpeedSelect returns a boolean if a field has been set.

### GetIntelHyperThreadingTech

`func (o *BiosPolicyAllOf) GetIntelHyperThreadingTech() string`

GetIntelHyperThreadingTech returns the IntelHyperThreadingTech field if non-nil, zero value otherwise.

### GetIntelHyperThreadingTechOk

`func (o *BiosPolicyAllOf) GetIntelHyperThreadingTechOk() (*string, bool)`

GetIntelHyperThreadingTechOk returns a tuple with the IntelHyperThreadingTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelHyperThreadingTech

`func (o *BiosPolicyAllOf) SetIntelHyperThreadingTech(v string)`

SetIntelHyperThreadingTech sets IntelHyperThreadingTech field to given value.

### HasIntelHyperThreadingTech

`func (o *BiosPolicyAllOf) HasIntelHyperThreadingTech() bool`

HasIntelHyperThreadingTech returns a boolean if a field has been set.

### GetIntelSpeedSelect

`func (o *BiosPolicyAllOf) GetIntelSpeedSelect() string`

GetIntelSpeedSelect returns the IntelSpeedSelect field if non-nil, zero value otherwise.

### GetIntelSpeedSelectOk

`func (o *BiosPolicyAllOf) GetIntelSpeedSelectOk() (*string, bool)`

GetIntelSpeedSelectOk returns a tuple with the IntelSpeedSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelSpeedSelect

`func (o *BiosPolicyAllOf) SetIntelSpeedSelect(v string)`

SetIntelSpeedSelect sets IntelSpeedSelect field to given value.

### HasIntelSpeedSelect

`func (o *BiosPolicyAllOf) HasIntelSpeedSelect() bool`

HasIntelSpeedSelect returns a boolean if a field has been set.

### GetIntelTurboBoostTech

`func (o *BiosPolicyAllOf) GetIntelTurboBoostTech() string`

GetIntelTurboBoostTech returns the IntelTurboBoostTech field if non-nil, zero value otherwise.

### GetIntelTurboBoostTechOk

`func (o *BiosPolicyAllOf) GetIntelTurboBoostTechOk() (*string, bool)`

GetIntelTurboBoostTechOk returns a tuple with the IntelTurboBoostTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelTurboBoostTech

`func (o *BiosPolicyAllOf) SetIntelTurboBoostTech(v string)`

SetIntelTurboBoostTech sets IntelTurboBoostTech field to given value.

### HasIntelTurboBoostTech

`func (o *BiosPolicyAllOf) HasIntelTurboBoostTech() bool`

HasIntelTurboBoostTech returns a boolean if a field has been set.

### GetIntelVirtualizationTechnology

`func (o *BiosPolicyAllOf) GetIntelVirtualizationTechnology() string`

GetIntelVirtualizationTechnology returns the IntelVirtualizationTechnology field if non-nil, zero value otherwise.

### GetIntelVirtualizationTechnologyOk

`func (o *BiosPolicyAllOf) GetIntelVirtualizationTechnologyOk() (*string, bool)`

GetIntelVirtualizationTechnologyOk returns a tuple with the IntelVirtualizationTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVirtualizationTechnology

`func (o *BiosPolicyAllOf) SetIntelVirtualizationTechnology(v string)`

SetIntelVirtualizationTechnology sets IntelVirtualizationTechnology field to given value.

### HasIntelVirtualizationTechnology

`func (o *BiosPolicyAllOf) HasIntelVirtualizationTechnology() bool`

HasIntelVirtualizationTechnology returns a boolean if a field has been set.

### GetIntelVtForDirectedIo

`func (o *BiosPolicyAllOf) GetIntelVtForDirectedIo() string`

GetIntelVtForDirectedIo returns the IntelVtForDirectedIo field if non-nil, zero value otherwise.

### GetIntelVtForDirectedIoOk

`func (o *BiosPolicyAllOf) GetIntelVtForDirectedIoOk() (*string, bool)`

GetIntelVtForDirectedIoOk returns a tuple with the IntelVtForDirectedIo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVtForDirectedIo

`func (o *BiosPolicyAllOf) SetIntelVtForDirectedIo(v string)`

SetIntelVtForDirectedIo sets IntelVtForDirectedIo field to given value.

### HasIntelVtForDirectedIo

`func (o *BiosPolicyAllOf) HasIntelVtForDirectedIo() bool`

HasIntelVtForDirectedIo returns a boolean if a field has been set.

### GetIntelVtdCoherencySupport

`func (o *BiosPolicyAllOf) GetIntelVtdCoherencySupport() string`

GetIntelVtdCoherencySupport returns the IntelVtdCoherencySupport field if non-nil, zero value otherwise.

### GetIntelVtdCoherencySupportOk

`func (o *BiosPolicyAllOf) GetIntelVtdCoherencySupportOk() (*string, bool)`

GetIntelVtdCoherencySupportOk returns a tuple with the IntelVtdCoherencySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVtdCoherencySupport

`func (o *BiosPolicyAllOf) SetIntelVtdCoherencySupport(v string)`

SetIntelVtdCoherencySupport sets IntelVtdCoherencySupport field to given value.

### HasIntelVtdCoherencySupport

`func (o *BiosPolicyAllOf) HasIntelVtdCoherencySupport() bool`

HasIntelVtdCoherencySupport returns a boolean if a field has been set.

### GetIntelVtdInterruptRemapping

`func (o *BiosPolicyAllOf) GetIntelVtdInterruptRemapping() string`

GetIntelVtdInterruptRemapping returns the IntelVtdInterruptRemapping field if non-nil, zero value otherwise.

### GetIntelVtdInterruptRemappingOk

`func (o *BiosPolicyAllOf) GetIntelVtdInterruptRemappingOk() (*string, bool)`

GetIntelVtdInterruptRemappingOk returns a tuple with the IntelVtdInterruptRemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVtdInterruptRemapping

`func (o *BiosPolicyAllOf) SetIntelVtdInterruptRemapping(v string)`

SetIntelVtdInterruptRemapping sets IntelVtdInterruptRemapping field to given value.

### HasIntelVtdInterruptRemapping

`func (o *BiosPolicyAllOf) HasIntelVtdInterruptRemapping() bool`

HasIntelVtdInterruptRemapping returns a boolean if a field has been set.

### GetIntelVtdPassThroughDmaSupport

`func (o *BiosPolicyAllOf) GetIntelVtdPassThroughDmaSupport() string`

GetIntelVtdPassThroughDmaSupport returns the IntelVtdPassThroughDmaSupport field if non-nil, zero value otherwise.

### GetIntelVtdPassThroughDmaSupportOk

`func (o *BiosPolicyAllOf) GetIntelVtdPassThroughDmaSupportOk() (*string, bool)`

GetIntelVtdPassThroughDmaSupportOk returns a tuple with the IntelVtdPassThroughDmaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVtdPassThroughDmaSupport

`func (o *BiosPolicyAllOf) SetIntelVtdPassThroughDmaSupport(v string)`

SetIntelVtdPassThroughDmaSupport sets IntelVtdPassThroughDmaSupport field to given value.

### HasIntelVtdPassThroughDmaSupport

`func (o *BiosPolicyAllOf) HasIntelVtdPassThroughDmaSupport() bool`

HasIntelVtdPassThroughDmaSupport returns a boolean if a field has been set.

### GetIntelVtdatsSupport

`func (o *BiosPolicyAllOf) GetIntelVtdatsSupport() string`

GetIntelVtdatsSupport returns the IntelVtdatsSupport field if non-nil, zero value otherwise.

### GetIntelVtdatsSupportOk

`func (o *BiosPolicyAllOf) GetIntelVtdatsSupportOk() (*string, bool)`

GetIntelVtdatsSupportOk returns a tuple with the IntelVtdatsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelVtdatsSupport

`func (o *BiosPolicyAllOf) SetIntelVtdatsSupport(v string)`

SetIntelVtdatsSupport sets IntelVtdatsSupport field to given value.

### HasIntelVtdatsSupport

`func (o *BiosPolicyAllOf) HasIntelVtdatsSupport() bool`

HasIntelVtdatsSupport returns a boolean if a field has been set.

### GetIohErrorEnable

`func (o *BiosPolicyAllOf) GetIohErrorEnable() string`

GetIohErrorEnable returns the IohErrorEnable field if non-nil, zero value otherwise.

### GetIohErrorEnableOk

`func (o *BiosPolicyAllOf) GetIohErrorEnableOk() (*string, bool)`

GetIohErrorEnableOk returns a tuple with the IohErrorEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIohErrorEnable

`func (o *BiosPolicyAllOf) SetIohErrorEnable(v string)`

SetIohErrorEnable sets IohErrorEnable field to given value.

### HasIohErrorEnable

`func (o *BiosPolicyAllOf) HasIohErrorEnable() bool`

HasIohErrorEnable returns a boolean if a field has been set.

### GetIohResource

`func (o *BiosPolicyAllOf) GetIohResource() string`

GetIohResource returns the IohResource field if non-nil, zero value otherwise.

### GetIohResourceOk

`func (o *BiosPolicyAllOf) GetIohResourceOk() (*string, bool)`

GetIohResourceOk returns a tuple with the IohResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIohResource

`func (o *BiosPolicyAllOf) SetIohResource(v string)`

SetIohResource sets IohResource field to given value.

### HasIohResource

`func (o *BiosPolicyAllOf) HasIohResource() bool`

HasIohResource returns a boolean if a field has been set.

### GetIpPrefetch

`func (o *BiosPolicyAllOf) GetIpPrefetch() string`

GetIpPrefetch returns the IpPrefetch field if non-nil, zero value otherwise.

### GetIpPrefetchOk

`func (o *BiosPolicyAllOf) GetIpPrefetchOk() (*string, bool)`

GetIpPrefetchOk returns a tuple with the IpPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefetch

`func (o *BiosPolicyAllOf) SetIpPrefetch(v string)`

SetIpPrefetch sets IpPrefetch field to given value.

### HasIpPrefetch

`func (o *BiosPolicyAllOf) HasIpPrefetch() bool`

HasIpPrefetch returns a boolean if a field has been set.

### GetIpv4http

`func (o *BiosPolicyAllOf) GetIpv4http() string`

GetIpv4http returns the Ipv4http field if non-nil, zero value otherwise.

### GetIpv4httpOk

`func (o *BiosPolicyAllOf) GetIpv4httpOk() (*string, bool)`

GetIpv4httpOk returns a tuple with the Ipv4http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4http

`func (o *BiosPolicyAllOf) SetIpv4http(v string)`

SetIpv4http sets Ipv4http field to given value.

### HasIpv4http

`func (o *BiosPolicyAllOf) HasIpv4http() bool`

HasIpv4http returns a boolean if a field has been set.

### GetIpv4pxe

`func (o *BiosPolicyAllOf) GetIpv4pxe() string`

GetIpv4pxe returns the Ipv4pxe field if non-nil, zero value otherwise.

### GetIpv4pxeOk

`func (o *BiosPolicyAllOf) GetIpv4pxeOk() (*string, bool)`

GetIpv4pxeOk returns a tuple with the Ipv4pxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4pxe

`func (o *BiosPolicyAllOf) SetIpv4pxe(v string)`

SetIpv4pxe sets Ipv4pxe field to given value.

### HasIpv4pxe

`func (o *BiosPolicyAllOf) HasIpv4pxe() bool`

HasIpv4pxe returns a boolean if a field has been set.

### GetIpv6http

`func (o *BiosPolicyAllOf) GetIpv6http() string`

GetIpv6http returns the Ipv6http field if non-nil, zero value otherwise.

### GetIpv6httpOk

`func (o *BiosPolicyAllOf) GetIpv6httpOk() (*string, bool)`

GetIpv6httpOk returns a tuple with the Ipv6http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6http

`func (o *BiosPolicyAllOf) SetIpv6http(v string)`

SetIpv6http sets Ipv6http field to given value.

### HasIpv6http

`func (o *BiosPolicyAllOf) HasIpv6http() bool`

HasIpv6http returns a boolean if a field has been set.

### GetIpv6pxe

`func (o *BiosPolicyAllOf) GetIpv6pxe() string`

GetIpv6pxe returns the Ipv6pxe field if non-nil, zero value otherwise.

### GetIpv6pxeOk

`func (o *BiosPolicyAllOf) GetIpv6pxeOk() (*string, bool)`

GetIpv6pxeOk returns a tuple with the Ipv6pxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6pxe

`func (o *BiosPolicyAllOf) SetIpv6pxe(v string)`

SetIpv6pxe sets Ipv6pxe field to given value.

### HasIpv6pxe

`func (o *BiosPolicyAllOf) HasIpv6pxe() bool`

HasIpv6pxe returns a boolean if a field has been set.

### GetKtiPrefetch

`func (o *BiosPolicyAllOf) GetKtiPrefetch() string`

GetKtiPrefetch returns the KtiPrefetch field if non-nil, zero value otherwise.

### GetKtiPrefetchOk

`func (o *BiosPolicyAllOf) GetKtiPrefetchOk() (*string, bool)`

GetKtiPrefetchOk returns a tuple with the KtiPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKtiPrefetch

`func (o *BiosPolicyAllOf) SetKtiPrefetch(v string)`

SetKtiPrefetch sets KtiPrefetch field to given value.

### HasKtiPrefetch

`func (o *BiosPolicyAllOf) HasKtiPrefetch() bool`

HasKtiPrefetch returns a boolean if a field has been set.

### GetLegacyOsRedirection

`func (o *BiosPolicyAllOf) GetLegacyOsRedirection() string`

GetLegacyOsRedirection returns the LegacyOsRedirection field if non-nil, zero value otherwise.

### GetLegacyOsRedirectionOk

`func (o *BiosPolicyAllOf) GetLegacyOsRedirectionOk() (*string, bool)`

GetLegacyOsRedirectionOk returns a tuple with the LegacyOsRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyOsRedirection

`func (o *BiosPolicyAllOf) SetLegacyOsRedirection(v string)`

SetLegacyOsRedirection sets LegacyOsRedirection field to given value.

### HasLegacyOsRedirection

`func (o *BiosPolicyAllOf) HasLegacyOsRedirection() bool`

HasLegacyOsRedirection returns a boolean if a field has been set.

### GetLegacyUsbSupport

`func (o *BiosPolicyAllOf) GetLegacyUsbSupport() string`

GetLegacyUsbSupport returns the LegacyUsbSupport field if non-nil, zero value otherwise.

### GetLegacyUsbSupportOk

`func (o *BiosPolicyAllOf) GetLegacyUsbSupportOk() (*string, bool)`

GetLegacyUsbSupportOk returns a tuple with the LegacyUsbSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyUsbSupport

`func (o *BiosPolicyAllOf) SetLegacyUsbSupport(v string)`

SetLegacyUsbSupport sets LegacyUsbSupport field to given value.

### HasLegacyUsbSupport

`func (o *BiosPolicyAllOf) HasLegacyUsbSupport() bool`

HasLegacyUsbSupport returns a boolean if a field has been set.

### GetLlcPrefetch

`func (o *BiosPolicyAllOf) GetLlcPrefetch() string`

GetLlcPrefetch returns the LlcPrefetch field if non-nil, zero value otherwise.

### GetLlcPrefetchOk

`func (o *BiosPolicyAllOf) GetLlcPrefetchOk() (*string, bool)`

GetLlcPrefetchOk returns a tuple with the LlcPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlcPrefetch

`func (o *BiosPolicyAllOf) SetLlcPrefetch(v string)`

SetLlcPrefetch sets LlcPrefetch field to given value.

### HasLlcPrefetch

`func (o *BiosPolicyAllOf) HasLlcPrefetch() bool`

HasLlcPrefetch returns a boolean if a field has been set.

### GetLomPort0state

`func (o *BiosPolicyAllOf) GetLomPort0state() string`

GetLomPort0state returns the LomPort0state field if non-nil, zero value otherwise.

### GetLomPort0stateOk

`func (o *BiosPolicyAllOf) GetLomPort0stateOk() (*string, bool)`

GetLomPort0stateOk returns a tuple with the LomPort0state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomPort0state

`func (o *BiosPolicyAllOf) SetLomPort0state(v string)`

SetLomPort0state sets LomPort0state field to given value.

### HasLomPort0state

`func (o *BiosPolicyAllOf) HasLomPort0state() bool`

HasLomPort0state returns a boolean if a field has been set.

### GetLomPort1state

`func (o *BiosPolicyAllOf) GetLomPort1state() string`

GetLomPort1state returns the LomPort1state field if non-nil, zero value otherwise.

### GetLomPort1stateOk

`func (o *BiosPolicyAllOf) GetLomPort1stateOk() (*string, bool)`

GetLomPort1stateOk returns a tuple with the LomPort1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomPort1state

`func (o *BiosPolicyAllOf) SetLomPort1state(v string)`

SetLomPort1state sets LomPort1state field to given value.

### HasLomPort1state

`func (o *BiosPolicyAllOf) HasLomPort1state() bool`

HasLomPort1state returns a boolean if a field has been set.

### GetLomPort2state

`func (o *BiosPolicyAllOf) GetLomPort2state() string`

GetLomPort2state returns the LomPort2state field if non-nil, zero value otherwise.

### GetLomPort2stateOk

`func (o *BiosPolicyAllOf) GetLomPort2stateOk() (*string, bool)`

GetLomPort2stateOk returns a tuple with the LomPort2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomPort2state

`func (o *BiosPolicyAllOf) SetLomPort2state(v string)`

SetLomPort2state sets LomPort2state field to given value.

### HasLomPort2state

`func (o *BiosPolicyAllOf) HasLomPort2state() bool`

HasLomPort2state returns a boolean if a field has been set.

### GetLomPort3state

`func (o *BiosPolicyAllOf) GetLomPort3state() string`

GetLomPort3state returns the LomPort3state field if non-nil, zero value otherwise.

### GetLomPort3stateOk

`func (o *BiosPolicyAllOf) GetLomPort3stateOk() (*string, bool)`

GetLomPort3stateOk returns a tuple with the LomPort3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomPort3state

`func (o *BiosPolicyAllOf) SetLomPort3state(v string)`

SetLomPort3state sets LomPort3state field to given value.

### HasLomPort3state

`func (o *BiosPolicyAllOf) HasLomPort3state() bool`

HasLomPort3state returns a boolean if a field has been set.

### GetLomPortsAllState

`func (o *BiosPolicyAllOf) GetLomPortsAllState() string`

GetLomPortsAllState returns the LomPortsAllState field if non-nil, zero value otherwise.

### GetLomPortsAllStateOk

`func (o *BiosPolicyAllOf) GetLomPortsAllStateOk() (*string, bool)`

GetLomPortsAllStateOk returns a tuple with the LomPortsAllState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomPortsAllState

`func (o *BiosPolicyAllOf) SetLomPortsAllState(v string)`

SetLomPortsAllState sets LomPortsAllState field to given value.

### HasLomPortsAllState

`func (o *BiosPolicyAllOf) HasLomPortsAllState() bool`

HasLomPortsAllState returns a boolean if a field has been set.

### GetLvDdrMode

`func (o *BiosPolicyAllOf) GetLvDdrMode() string`

GetLvDdrMode returns the LvDdrMode field if non-nil, zero value otherwise.

### GetLvDdrModeOk

`func (o *BiosPolicyAllOf) GetLvDdrModeOk() (*string, bool)`

GetLvDdrModeOk returns a tuple with the LvDdrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvDdrMode

`func (o *BiosPolicyAllOf) SetLvDdrMode(v string)`

SetLvDdrMode sets LvDdrMode field to given value.

### HasLvDdrMode

`func (o *BiosPolicyAllOf) HasLvDdrMode() bool`

HasLvDdrMode returns a boolean if a field has been set.

### GetMakeDeviceNonBootable

`func (o *BiosPolicyAllOf) GetMakeDeviceNonBootable() string`

GetMakeDeviceNonBootable returns the MakeDeviceNonBootable field if non-nil, zero value otherwise.

### GetMakeDeviceNonBootableOk

`func (o *BiosPolicyAllOf) GetMakeDeviceNonBootableOk() (*string, bool)`

GetMakeDeviceNonBootableOk returns a tuple with the MakeDeviceNonBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeDeviceNonBootable

`func (o *BiosPolicyAllOf) SetMakeDeviceNonBootable(v string)`

SetMakeDeviceNonBootable sets MakeDeviceNonBootable field to given value.

### HasMakeDeviceNonBootable

`func (o *BiosPolicyAllOf) HasMakeDeviceNonBootable() bool`

HasMakeDeviceNonBootable returns a boolean if a field has been set.

### GetMemoryBandwidthBoost

`func (o *BiosPolicyAllOf) GetMemoryBandwidthBoost() string`

GetMemoryBandwidthBoost returns the MemoryBandwidthBoost field if non-nil, zero value otherwise.

### GetMemoryBandwidthBoostOk

`func (o *BiosPolicyAllOf) GetMemoryBandwidthBoostOk() (*string, bool)`

GetMemoryBandwidthBoostOk returns a tuple with the MemoryBandwidthBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBandwidthBoost

`func (o *BiosPolicyAllOf) SetMemoryBandwidthBoost(v string)`

SetMemoryBandwidthBoost sets MemoryBandwidthBoost field to given value.

### HasMemoryBandwidthBoost

`func (o *BiosPolicyAllOf) HasMemoryBandwidthBoost() bool`

HasMemoryBandwidthBoost returns a boolean if a field has been set.

### GetMemoryInterLeave

`func (o *BiosPolicyAllOf) GetMemoryInterLeave() string`

GetMemoryInterLeave returns the MemoryInterLeave field if non-nil, zero value otherwise.

### GetMemoryInterLeaveOk

`func (o *BiosPolicyAllOf) GetMemoryInterLeaveOk() (*string, bool)`

GetMemoryInterLeaveOk returns a tuple with the MemoryInterLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInterLeave

`func (o *BiosPolicyAllOf) SetMemoryInterLeave(v string)`

SetMemoryInterLeave sets MemoryInterLeave field to given value.

### HasMemoryInterLeave

`func (o *BiosPolicyAllOf) HasMemoryInterLeave() bool`

HasMemoryInterLeave returns a boolean if a field has been set.

### GetMemoryMappedIoAbove4gb

`func (o *BiosPolicyAllOf) GetMemoryMappedIoAbove4gb() string`

GetMemoryMappedIoAbove4gb returns the MemoryMappedIoAbove4gb field if non-nil, zero value otherwise.

### GetMemoryMappedIoAbove4gbOk

`func (o *BiosPolicyAllOf) GetMemoryMappedIoAbove4gbOk() (*string, bool)`

GetMemoryMappedIoAbove4gbOk returns a tuple with the MemoryMappedIoAbove4gb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMappedIoAbove4gb

`func (o *BiosPolicyAllOf) SetMemoryMappedIoAbove4gb(v string)`

SetMemoryMappedIoAbove4gb sets MemoryMappedIoAbove4gb field to given value.

### HasMemoryMappedIoAbove4gb

`func (o *BiosPolicyAllOf) HasMemoryMappedIoAbove4gb() bool`

HasMemoryMappedIoAbove4gb returns a boolean if a field has been set.

### GetMemoryRefreshRate

`func (o *BiosPolicyAllOf) GetMemoryRefreshRate() string`

GetMemoryRefreshRate returns the MemoryRefreshRate field if non-nil, zero value otherwise.

### GetMemoryRefreshRateOk

`func (o *BiosPolicyAllOf) GetMemoryRefreshRateOk() (*string, bool)`

GetMemoryRefreshRateOk returns a tuple with the MemoryRefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRefreshRate

`func (o *BiosPolicyAllOf) SetMemoryRefreshRate(v string)`

SetMemoryRefreshRate sets MemoryRefreshRate field to given value.

### HasMemoryRefreshRate

`func (o *BiosPolicyAllOf) HasMemoryRefreshRate() bool`

HasMemoryRefreshRate returns a boolean if a field has been set.

### GetMemorySizeLimit

`func (o *BiosPolicyAllOf) GetMemorySizeLimit() string`

GetMemorySizeLimit returns the MemorySizeLimit field if non-nil, zero value otherwise.

### GetMemorySizeLimitOk

`func (o *BiosPolicyAllOf) GetMemorySizeLimitOk() (*string, bool)`

GetMemorySizeLimitOk returns a tuple with the MemorySizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeLimit

`func (o *BiosPolicyAllOf) SetMemorySizeLimit(v string)`

SetMemorySizeLimit sets MemorySizeLimit field to given value.

### HasMemorySizeLimit

`func (o *BiosPolicyAllOf) HasMemorySizeLimit() bool`

HasMemorySizeLimit returns a boolean if a field has been set.

### GetMemoryThermalThrottling

`func (o *BiosPolicyAllOf) GetMemoryThermalThrottling() string`

GetMemoryThermalThrottling returns the MemoryThermalThrottling field if non-nil, zero value otherwise.

### GetMemoryThermalThrottlingOk

`func (o *BiosPolicyAllOf) GetMemoryThermalThrottlingOk() (*string, bool)`

GetMemoryThermalThrottlingOk returns a tuple with the MemoryThermalThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryThermalThrottling

`func (o *BiosPolicyAllOf) SetMemoryThermalThrottling(v string)`

SetMemoryThermalThrottling sets MemoryThermalThrottling field to given value.

### HasMemoryThermalThrottling

`func (o *BiosPolicyAllOf) HasMemoryThermalThrottling() bool`

HasMemoryThermalThrottling returns a boolean if a field has been set.

### GetMirroringMode

`func (o *BiosPolicyAllOf) GetMirroringMode() string`

GetMirroringMode returns the MirroringMode field if non-nil, zero value otherwise.

### GetMirroringModeOk

`func (o *BiosPolicyAllOf) GetMirroringModeOk() (*string, bool)`

GetMirroringModeOk returns a tuple with the MirroringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroringMode

`func (o *BiosPolicyAllOf) SetMirroringMode(v string)`

SetMirroringMode sets MirroringMode field to given value.

### HasMirroringMode

`func (o *BiosPolicyAllOf) HasMirroringMode() bool`

HasMirroringMode returns a boolean if a field has been set.

### GetMmcfgBase

`func (o *BiosPolicyAllOf) GetMmcfgBase() string`

GetMmcfgBase returns the MmcfgBase field if non-nil, zero value otherwise.

### GetMmcfgBaseOk

`func (o *BiosPolicyAllOf) GetMmcfgBaseOk() (*string, bool)`

GetMmcfgBaseOk returns a tuple with the MmcfgBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmcfgBase

`func (o *BiosPolicyAllOf) SetMmcfgBase(v string)`

SetMmcfgBase sets MmcfgBase field to given value.

### HasMmcfgBase

`func (o *BiosPolicyAllOf) HasMmcfgBase() bool`

HasMmcfgBase returns a boolean if a field has been set.

### GetNetworkStack

`func (o *BiosPolicyAllOf) GetNetworkStack() string`

GetNetworkStack returns the NetworkStack field if non-nil, zero value otherwise.

### GetNetworkStackOk

`func (o *BiosPolicyAllOf) GetNetworkStackOk() (*string, bool)`

GetNetworkStackOk returns a tuple with the NetworkStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStack

`func (o *BiosPolicyAllOf) SetNetworkStack(v string)`

SetNetworkStack sets NetworkStack field to given value.

### HasNetworkStack

`func (o *BiosPolicyAllOf) HasNetworkStack() bool`

HasNetworkStack returns a boolean if a field has been set.

### GetNumaOptimized

`func (o *BiosPolicyAllOf) GetNumaOptimized() string`

GetNumaOptimized returns the NumaOptimized field if non-nil, zero value otherwise.

### GetNumaOptimizedOk

`func (o *BiosPolicyAllOf) GetNumaOptimizedOk() (*string, bool)`

GetNumaOptimizedOk returns a tuple with the NumaOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaOptimized

`func (o *BiosPolicyAllOf) SetNumaOptimized(v string)`

SetNumaOptimized sets NumaOptimized field to given value.

### HasNumaOptimized

`func (o *BiosPolicyAllOf) HasNumaOptimized() bool`

HasNumaOptimized returns a boolean if a field has been set.

### GetNvmdimmPerformConfig

`func (o *BiosPolicyAllOf) GetNvmdimmPerformConfig() string`

GetNvmdimmPerformConfig returns the NvmdimmPerformConfig field if non-nil, zero value otherwise.

### GetNvmdimmPerformConfigOk

`func (o *BiosPolicyAllOf) GetNvmdimmPerformConfigOk() (*string, bool)`

GetNvmdimmPerformConfigOk returns a tuple with the NvmdimmPerformConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmdimmPerformConfig

`func (o *BiosPolicyAllOf) SetNvmdimmPerformConfig(v string)`

SetNvmdimmPerformConfig sets NvmdimmPerformConfig field to given value.

### HasNvmdimmPerformConfig

`func (o *BiosPolicyAllOf) HasNvmdimmPerformConfig() bool`

HasNvmdimmPerformConfig returns a boolean if a field has been set.

### GetOnboard10gbitLom

`func (o *BiosPolicyAllOf) GetOnboard10gbitLom() string`

GetOnboard10gbitLom returns the Onboard10gbitLom field if non-nil, zero value otherwise.

### GetOnboard10gbitLomOk

`func (o *BiosPolicyAllOf) GetOnboard10gbitLomOk() (*string, bool)`

GetOnboard10gbitLomOk returns a tuple with the Onboard10gbitLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboard10gbitLom

`func (o *BiosPolicyAllOf) SetOnboard10gbitLom(v string)`

SetOnboard10gbitLom sets Onboard10gbitLom field to given value.

### HasOnboard10gbitLom

`func (o *BiosPolicyAllOf) HasOnboard10gbitLom() bool`

HasOnboard10gbitLom returns a boolean if a field has been set.

### GetOnboardGbitLom

`func (o *BiosPolicyAllOf) GetOnboardGbitLom() string`

GetOnboardGbitLom returns the OnboardGbitLom field if non-nil, zero value otherwise.

### GetOnboardGbitLomOk

`func (o *BiosPolicyAllOf) GetOnboardGbitLomOk() (*string, bool)`

GetOnboardGbitLomOk returns a tuple with the OnboardGbitLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardGbitLom

`func (o *BiosPolicyAllOf) SetOnboardGbitLom(v string)`

SetOnboardGbitLom sets OnboardGbitLom field to given value.

### HasOnboardGbitLom

`func (o *BiosPolicyAllOf) HasOnboardGbitLom() bool`

HasOnboardGbitLom returns a boolean if a field has been set.

### GetOnboardScuStorageSupport

`func (o *BiosPolicyAllOf) GetOnboardScuStorageSupport() string`

GetOnboardScuStorageSupport returns the OnboardScuStorageSupport field if non-nil, zero value otherwise.

### GetOnboardScuStorageSupportOk

`func (o *BiosPolicyAllOf) GetOnboardScuStorageSupportOk() (*string, bool)`

GetOnboardScuStorageSupportOk returns a tuple with the OnboardScuStorageSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardScuStorageSupport

`func (o *BiosPolicyAllOf) SetOnboardScuStorageSupport(v string)`

SetOnboardScuStorageSupport sets OnboardScuStorageSupport field to given value.

### HasOnboardScuStorageSupport

`func (o *BiosPolicyAllOf) HasOnboardScuStorageSupport() bool`

HasOnboardScuStorageSupport returns a boolean if a field has been set.

### GetOnboardScuStorageSwStack

`func (o *BiosPolicyAllOf) GetOnboardScuStorageSwStack() string`

GetOnboardScuStorageSwStack returns the OnboardScuStorageSwStack field if non-nil, zero value otherwise.

### GetOnboardScuStorageSwStackOk

`func (o *BiosPolicyAllOf) GetOnboardScuStorageSwStackOk() (*string, bool)`

GetOnboardScuStorageSwStackOk returns a tuple with the OnboardScuStorageSwStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardScuStorageSwStack

`func (o *BiosPolicyAllOf) SetOnboardScuStorageSwStack(v string)`

SetOnboardScuStorageSwStack sets OnboardScuStorageSwStack field to given value.

### HasOnboardScuStorageSwStack

`func (o *BiosPolicyAllOf) HasOnboardScuStorageSwStack() bool`

HasOnboardScuStorageSwStack returns a boolean if a field has been set.

### GetOperationMode

`func (o *BiosPolicyAllOf) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *BiosPolicyAllOf) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *BiosPolicyAllOf) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *BiosPolicyAllOf) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### GetOsBootWatchdogTimer

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimer() string`

GetOsBootWatchdogTimer returns the OsBootWatchdogTimer field if non-nil, zero value otherwise.

### GetOsBootWatchdogTimerOk

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimerOk() (*string, bool)`

GetOsBootWatchdogTimerOk returns a tuple with the OsBootWatchdogTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBootWatchdogTimer

`func (o *BiosPolicyAllOf) SetOsBootWatchdogTimer(v string)`

SetOsBootWatchdogTimer sets OsBootWatchdogTimer field to given value.

### HasOsBootWatchdogTimer

`func (o *BiosPolicyAllOf) HasOsBootWatchdogTimer() bool`

HasOsBootWatchdogTimer returns a boolean if a field has been set.

### GetOsBootWatchdogTimerPolicy

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimerPolicy() string`

GetOsBootWatchdogTimerPolicy returns the OsBootWatchdogTimerPolicy field if non-nil, zero value otherwise.

### GetOsBootWatchdogTimerPolicyOk

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimerPolicyOk() (*string, bool)`

GetOsBootWatchdogTimerPolicyOk returns a tuple with the OsBootWatchdogTimerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBootWatchdogTimerPolicy

`func (o *BiosPolicyAllOf) SetOsBootWatchdogTimerPolicy(v string)`

SetOsBootWatchdogTimerPolicy sets OsBootWatchdogTimerPolicy field to given value.

### HasOsBootWatchdogTimerPolicy

`func (o *BiosPolicyAllOf) HasOsBootWatchdogTimerPolicy() bool`

HasOsBootWatchdogTimerPolicy returns a boolean if a field has been set.

### GetOsBootWatchdogTimerTimeout

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimerTimeout() string`

GetOsBootWatchdogTimerTimeout returns the OsBootWatchdogTimerTimeout field if non-nil, zero value otherwise.

### GetOsBootWatchdogTimerTimeoutOk

`func (o *BiosPolicyAllOf) GetOsBootWatchdogTimerTimeoutOk() (*string, bool)`

GetOsBootWatchdogTimerTimeoutOk returns a tuple with the OsBootWatchdogTimerTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBootWatchdogTimerTimeout

`func (o *BiosPolicyAllOf) SetOsBootWatchdogTimerTimeout(v string)`

SetOsBootWatchdogTimerTimeout sets OsBootWatchdogTimerTimeout field to given value.

### HasOsBootWatchdogTimerTimeout

`func (o *BiosPolicyAllOf) HasOsBootWatchdogTimerTimeout() bool`

HasOsBootWatchdogTimerTimeout returns a boolean if a field has been set.

### GetOutOfBandMgmtPort

`func (o *BiosPolicyAllOf) GetOutOfBandMgmtPort() string`

GetOutOfBandMgmtPort returns the OutOfBandMgmtPort field if non-nil, zero value otherwise.

### GetOutOfBandMgmtPortOk

`func (o *BiosPolicyAllOf) GetOutOfBandMgmtPortOk() (*string, bool)`

GetOutOfBandMgmtPortOk returns a tuple with the OutOfBandMgmtPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandMgmtPort

`func (o *BiosPolicyAllOf) SetOutOfBandMgmtPort(v string)`

SetOutOfBandMgmtPort sets OutOfBandMgmtPort field to given value.

### HasOutOfBandMgmtPort

`func (o *BiosPolicyAllOf) HasOutOfBandMgmtPort() bool`

HasOutOfBandMgmtPort returns a boolean if a field has been set.

### GetPackageCstateLimit

`func (o *BiosPolicyAllOf) GetPackageCstateLimit() string`

GetPackageCstateLimit returns the PackageCstateLimit field if non-nil, zero value otherwise.

### GetPackageCstateLimitOk

`func (o *BiosPolicyAllOf) GetPackageCstateLimitOk() (*string, bool)`

GetPackageCstateLimitOk returns a tuple with the PackageCstateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCstateLimit

`func (o *BiosPolicyAllOf) SetPackageCstateLimit(v string)`

SetPackageCstateLimit sets PackageCstateLimit field to given value.

### HasPackageCstateLimit

`func (o *BiosPolicyAllOf) HasPackageCstateLimit() bool`

HasPackageCstateLimit returns a boolean if a field has been set.

### GetPanicHighWatermark

`func (o *BiosPolicyAllOf) GetPanicHighWatermark() string`

GetPanicHighWatermark returns the PanicHighWatermark field if non-nil, zero value otherwise.

### GetPanicHighWatermarkOk

`func (o *BiosPolicyAllOf) GetPanicHighWatermarkOk() (*string, bool)`

GetPanicHighWatermarkOk returns a tuple with the PanicHighWatermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanicHighWatermark

`func (o *BiosPolicyAllOf) SetPanicHighWatermark(v string)`

SetPanicHighWatermark sets PanicHighWatermark field to given value.

### HasPanicHighWatermark

`func (o *BiosPolicyAllOf) HasPanicHighWatermark() bool`

HasPanicHighWatermark returns a boolean if a field has been set.

### GetPartialCacheLineSparing

`func (o *BiosPolicyAllOf) GetPartialCacheLineSparing() string`

GetPartialCacheLineSparing returns the PartialCacheLineSparing field if non-nil, zero value otherwise.

### GetPartialCacheLineSparingOk

`func (o *BiosPolicyAllOf) GetPartialCacheLineSparingOk() (*string, bool)`

GetPartialCacheLineSparingOk returns a tuple with the PartialCacheLineSparing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialCacheLineSparing

`func (o *BiosPolicyAllOf) SetPartialCacheLineSparing(v string)`

SetPartialCacheLineSparing sets PartialCacheLineSparing field to given value.

### HasPartialCacheLineSparing

`func (o *BiosPolicyAllOf) HasPartialCacheLineSparing() bool`

HasPartialCacheLineSparing returns a boolean if a field has been set.

### GetPartialMirrorModeConfig

`func (o *BiosPolicyAllOf) GetPartialMirrorModeConfig() string`

GetPartialMirrorModeConfig returns the PartialMirrorModeConfig field if non-nil, zero value otherwise.

### GetPartialMirrorModeConfigOk

`func (o *BiosPolicyAllOf) GetPartialMirrorModeConfigOk() (*string, bool)`

GetPartialMirrorModeConfigOk returns a tuple with the PartialMirrorModeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorModeConfig

`func (o *BiosPolicyAllOf) SetPartialMirrorModeConfig(v string)`

SetPartialMirrorModeConfig sets PartialMirrorModeConfig field to given value.

### HasPartialMirrorModeConfig

`func (o *BiosPolicyAllOf) HasPartialMirrorModeConfig() bool`

HasPartialMirrorModeConfig returns a boolean if a field has been set.

### GetPartialMirrorPercent

`func (o *BiosPolicyAllOf) GetPartialMirrorPercent() string`

GetPartialMirrorPercent returns the PartialMirrorPercent field if non-nil, zero value otherwise.

### GetPartialMirrorPercentOk

`func (o *BiosPolicyAllOf) GetPartialMirrorPercentOk() (*string, bool)`

GetPartialMirrorPercentOk returns a tuple with the PartialMirrorPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorPercent

`func (o *BiosPolicyAllOf) SetPartialMirrorPercent(v string)`

SetPartialMirrorPercent sets PartialMirrorPercent field to given value.

### HasPartialMirrorPercent

`func (o *BiosPolicyAllOf) HasPartialMirrorPercent() bool`

HasPartialMirrorPercent returns a boolean if a field has been set.

### GetPartialMirrorValue1

`func (o *BiosPolicyAllOf) GetPartialMirrorValue1() string`

GetPartialMirrorValue1 returns the PartialMirrorValue1 field if non-nil, zero value otherwise.

### GetPartialMirrorValue1Ok

`func (o *BiosPolicyAllOf) GetPartialMirrorValue1Ok() (*string, bool)`

GetPartialMirrorValue1Ok returns a tuple with the PartialMirrorValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorValue1

`func (o *BiosPolicyAllOf) SetPartialMirrorValue1(v string)`

SetPartialMirrorValue1 sets PartialMirrorValue1 field to given value.

### HasPartialMirrorValue1

`func (o *BiosPolicyAllOf) HasPartialMirrorValue1() bool`

HasPartialMirrorValue1 returns a boolean if a field has been set.

### GetPartialMirrorValue2

`func (o *BiosPolicyAllOf) GetPartialMirrorValue2() string`

GetPartialMirrorValue2 returns the PartialMirrorValue2 field if non-nil, zero value otherwise.

### GetPartialMirrorValue2Ok

`func (o *BiosPolicyAllOf) GetPartialMirrorValue2Ok() (*string, bool)`

GetPartialMirrorValue2Ok returns a tuple with the PartialMirrorValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorValue2

`func (o *BiosPolicyAllOf) SetPartialMirrorValue2(v string)`

SetPartialMirrorValue2 sets PartialMirrorValue2 field to given value.

### HasPartialMirrorValue2

`func (o *BiosPolicyAllOf) HasPartialMirrorValue2() bool`

HasPartialMirrorValue2 returns a boolean if a field has been set.

### GetPartialMirrorValue3

`func (o *BiosPolicyAllOf) GetPartialMirrorValue3() string`

GetPartialMirrorValue3 returns the PartialMirrorValue3 field if non-nil, zero value otherwise.

### GetPartialMirrorValue3Ok

`func (o *BiosPolicyAllOf) GetPartialMirrorValue3Ok() (*string, bool)`

GetPartialMirrorValue3Ok returns a tuple with the PartialMirrorValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorValue3

`func (o *BiosPolicyAllOf) SetPartialMirrorValue3(v string)`

SetPartialMirrorValue3 sets PartialMirrorValue3 field to given value.

### HasPartialMirrorValue3

`func (o *BiosPolicyAllOf) HasPartialMirrorValue3() bool`

HasPartialMirrorValue3 returns a boolean if a field has been set.

### GetPartialMirrorValue4

`func (o *BiosPolicyAllOf) GetPartialMirrorValue4() string`

GetPartialMirrorValue4 returns the PartialMirrorValue4 field if non-nil, zero value otherwise.

### GetPartialMirrorValue4Ok

`func (o *BiosPolicyAllOf) GetPartialMirrorValue4Ok() (*string, bool)`

GetPartialMirrorValue4Ok returns a tuple with the PartialMirrorValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialMirrorValue4

`func (o *BiosPolicyAllOf) SetPartialMirrorValue4(v string)`

SetPartialMirrorValue4 sets PartialMirrorValue4 field to given value.

### HasPartialMirrorValue4

`func (o *BiosPolicyAllOf) HasPartialMirrorValue4() bool`

HasPartialMirrorValue4 returns a boolean if a field has been set.

### GetPatrolScrub

`func (o *BiosPolicyAllOf) GetPatrolScrub() string`

GetPatrolScrub returns the PatrolScrub field if non-nil, zero value otherwise.

### GetPatrolScrubOk

`func (o *BiosPolicyAllOf) GetPatrolScrubOk() (*string, bool)`

GetPatrolScrubOk returns a tuple with the PatrolScrub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatrolScrub

`func (o *BiosPolicyAllOf) SetPatrolScrub(v string)`

SetPatrolScrub sets PatrolScrub field to given value.

### HasPatrolScrub

`func (o *BiosPolicyAllOf) HasPatrolScrub() bool`

HasPatrolScrub returns a boolean if a field has been set.

### GetPatrolScrubDuration

`func (o *BiosPolicyAllOf) GetPatrolScrubDuration() string`

GetPatrolScrubDuration returns the PatrolScrubDuration field if non-nil, zero value otherwise.

### GetPatrolScrubDurationOk

`func (o *BiosPolicyAllOf) GetPatrolScrubDurationOk() (*string, bool)`

GetPatrolScrubDurationOk returns a tuple with the PatrolScrubDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatrolScrubDuration

`func (o *BiosPolicyAllOf) SetPatrolScrubDuration(v string)`

SetPatrolScrubDuration sets PatrolScrubDuration field to given value.

### HasPatrolScrubDuration

`func (o *BiosPolicyAllOf) HasPatrolScrubDuration() bool`

HasPatrolScrubDuration returns a boolean if a field has been set.

### GetPcIeRasSupport

`func (o *BiosPolicyAllOf) GetPcIeRasSupport() string`

GetPcIeRasSupport returns the PcIeRasSupport field if non-nil, zero value otherwise.

### GetPcIeRasSupportOk

`func (o *BiosPolicyAllOf) GetPcIeRasSupportOk() (*string, bool)`

GetPcIeRasSupportOk returns a tuple with the PcIeRasSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcIeRasSupport

`func (o *BiosPolicyAllOf) SetPcIeRasSupport(v string)`

SetPcIeRasSupport sets PcIeRasSupport field to given value.

### HasPcIeRasSupport

`func (o *BiosPolicyAllOf) HasPcIeRasSupport() bool`

HasPcIeRasSupport returns a boolean if a field has been set.

### GetPcIeSsdHotPlugSupport

`func (o *BiosPolicyAllOf) GetPcIeSsdHotPlugSupport() string`

GetPcIeSsdHotPlugSupport returns the PcIeSsdHotPlugSupport field if non-nil, zero value otherwise.

### GetPcIeSsdHotPlugSupportOk

`func (o *BiosPolicyAllOf) GetPcIeSsdHotPlugSupportOk() (*string, bool)`

GetPcIeSsdHotPlugSupportOk returns a tuple with the PcIeSsdHotPlugSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcIeSsdHotPlugSupport

`func (o *BiosPolicyAllOf) SetPcIeSsdHotPlugSupport(v string)`

SetPcIeSsdHotPlugSupport sets PcIeSsdHotPlugSupport field to given value.

### HasPcIeSsdHotPlugSupport

`func (o *BiosPolicyAllOf) HasPcIeSsdHotPlugSupport() bool`

HasPcIeSsdHotPlugSupport returns a boolean if a field has been set.

### GetPchUsb30mode

`func (o *BiosPolicyAllOf) GetPchUsb30mode() string`

GetPchUsb30mode returns the PchUsb30mode field if non-nil, zero value otherwise.

### GetPchUsb30modeOk

`func (o *BiosPolicyAllOf) GetPchUsb30modeOk() (*string, bool)`

GetPchUsb30modeOk returns a tuple with the PchUsb30mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPchUsb30mode

`func (o *BiosPolicyAllOf) SetPchUsb30mode(v string)`

SetPchUsb30mode sets PchUsb30mode field to given value.

### HasPchUsb30mode

`func (o *BiosPolicyAllOf) HasPchUsb30mode() bool`

HasPchUsb30mode returns a boolean if a field has been set.

### GetPciOptionRoMs

`func (o *BiosPolicyAllOf) GetPciOptionRoMs() string`

GetPciOptionRoMs returns the PciOptionRoMs field if non-nil, zero value otherwise.

### GetPciOptionRoMsOk

`func (o *BiosPolicyAllOf) GetPciOptionRoMsOk() (*string, bool)`

GetPciOptionRoMsOk returns a tuple with the PciOptionRoMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciOptionRoMs

`func (o *BiosPolicyAllOf) SetPciOptionRoMs(v string)`

SetPciOptionRoMs sets PciOptionRoMs field to given value.

### HasPciOptionRoMs

`func (o *BiosPolicyAllOf) HasPciOptionRoMs() bool`

HasPciOptionRoMs returns a boolean if a field has been set.

### GetPciRomClp

`func (o *BiosPolicyAllOf) GetPciRomClp() string`

GetPciRomClp returns the PciRomClp field if non-nil, zero value otherwise.

### GetPciRomClpOk

`func (o *BiosPolicyAllOf) GetPciRomClpOk() (*string, bool)`

GetPciRomClpOk returns a tuple with the PciRomClp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciRomClp

`func (o *BiosPolicyAllOf) SetPciRomClp(v string)`

SetPciRomClp sets PciRomClp field to given value.

### HasPciRomClp

`func (o *BiosPolicyAllOf) HasPciRomClp() bool`

HasPciRomClp returns a boolean if a field has been set.

### GetPcieAriSupport

`func (o *BiosPolicyAllOf) GetPcieAriSupport() string`

GetPcieAriSupport returns the PcieAriSupport field if non-nil, zero value otherwise.

### GetPcieAriSupportOk

`func (o *BiosPolicyAllOf) GetPcieAriSupportOk() (*string, bool)`

GetPcieAriSupportOk returns a tuple with the PcieAriSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieAriSupport

`func (o *BiosPolicyAllOf) SetPcieAriSupport(v string)`

SetPcieAriSupport sets PcieAriSupport field to given value.

### HasPcieAriSupport

`func (o *BiosPolicyAllOf) HasPcieAriSupport() bool`

HasPcieAriSupport returns a boolean if a field has been set.

### GetPciePllSsc

`func (o *BiosPolicyAllOf) GetPciePllSsc() string`

GetPciePllSsc returns the PciePllSsc field if non-nil, zero value otherwise.

### GetPciePllSscOk

`func (o *BiosPolicyAllOf) GetPciePllSscOk() (*string, bool)`

GetPciePllSscOk returns a tuple with the PciePllSsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciePllSsc

`func (o *BiosPolicyAllOf) SetPciePllSsc(v string)`

SetPciePllSsc sets PciePllSsc field to given value.

### HasPciePllSsc

`func (o *BiosPolicyAllOf) HasPciePllSsc() bool`

HasPciePllSsc returns a boolean if a field has been set.

### GetPcieSlotMraid1linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotMraid1linkSpeed() string`

GetPcieSlotMraid1linkSpeed returns the PcieSlotMraid1linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotMraid1linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotMraid1linkSpeedOk() (*string, bool)`

GetPcieSlotMraid1linkSpeedOk returns a tuple with the PcieSlotMraid1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMraid1linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotMraid1linkSpeed(v string)`

SetPcieSlotMraid1linkSpeed sets PcieSlotMraid1linkSpeed field to given value.

### HasPcieSlotMraid1linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotMraid1linkSpeed() bool`

HasPcieSlotMraid1linkSpeed returns a boolean if a field has been set.

### GetPcieSlotMraid1optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotMraid1optionRom() string`

GetPcieSlotMraid1optionRom returns the PcieSlotMraid1optionRom field if non-nil, zero value otherwise.

### GetPcieSlotMraid1optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotMraid1optionRomOk() (*string, bool)`

GetPcieSlotMraid1optionRomOk returns a tuple with the PcieSlotMraid1optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMraid1optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotMraid1optionRom(v string)`

SetPcieSlotMraid1optionRom sets PcieSlotMraid1optionRom field to given value.

### HasPcieSlotMraid1optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotMraid1optionRom() bool`

HasPcieSlotMraid1optionRom returns a boolean if a field has been set.

### GetPcieSlotMraid2linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotMraid2linkSpeed() string`

GetPcieSlotMraid2linkSpeed returns the PcieSlotMraid2linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotMraid2linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotMraid2linkSpeedOk() (*string, bool)`

GetPcieSlotMraid2linkSpeedOk returns a tuple with the PcieSlotMraid2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMraid2linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotMraid2linkSpeed(v string)`

SetPcieSlotMraid2linkSpeed sets PcieSlotMraid2linkSpeed field to given value.

### HasPcieSlotMraid2linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotMraid2linkSpeed() bool`

HasPcieSlotMraid2linkSpeed returns a boolean if a field has been set.

### GetPcieSlotMraid2optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotMraid2optionRom() string`

GetPcieSlotMraid2optionRom returns the PcieSlotMraid2optionRom field if non-nil, zero value otherwise.

### GetPcieSlotMraid2optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotMraid2optionRomOk() (*string, bool)`

GetPcieSlotMraid2optionRomOk returns a tuple with the PcieSlotMraid2optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMraid2optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotMraid2optionRom(v string)`

SetPcieSlotMraid2optionRom sets PcieSlotMraid2optionRom field to given value.

### HasPcieSlotMraid2optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotMraid2optionRom() bool`

HasPcieSlotMraid2optionRom returns a boolean if a field has been set.

### GetPcieSlotMstorraidLinkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotMstorraidLinkSpeed() string`

GetPcieSlotMstorraidLinkSpeed returns the PcieSlotMstorraidLinkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotMstorraidLinkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotMstorraidLinkSpeedOk() (*string, bool)`

GetPcieSlotMstorraidLinkSpeedOk returns a tuple with the PcieSlotMstorraidLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMstorraidLinkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotMstorraidLinkSpeed(v string)`

SetPcieSlotMstorraidLinkSpeed sets PcieSlotMstorraidLinkSpeed field to given value.

### HasPcieSlotMstorraidLinkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotMstorraidLinkSpeed() bool`

HasPcieSlotMstorraidLinkSpeed returns a boolean if a field has been set.

### GetPcieSlotMstorraidOptionRom

`func (o *BiosPolicyAllOf) GetPcieSlotMstorraidOptionRom() string`

GetPcieSlotMstorraidOptionRom returns the PcieSlotMstorraidOptionRom field if non-nil, zero value otherwise.

### GetPcieSlotMstorraidOptionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotMstorraidOptionRomOk() (*string, bool)`

GetPcieSlotMstorraidOptionRomOk returns a tuple with the PcieSlotMstorraidOptionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotMstorraidOptionRom

`func (o *BiosPolicyAllOf) SetPcieSlotMstorraidOptionRom(v string)`

SetPcieSlotMstorraidOptionRom sets PcieSlotMstorraidOptionRom field to given value.

### HasPcieSlotMstorraidOptionRom

`func (o *BiosPolicyAllOf) HasPcieSlotMstorraidOptionRom() bool`

HasPcieSlotMstorraidOptionRom returns a boolean if a field has been set.

### GetPcieSlotNvme1linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme1linkSpeed() string`

GetPcieSlotNvme1linkSpeed returns the PcieSlotNvme1linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme1linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme1linkSpeedOk() (*string, bool)`

GetPcieSlotNvme1linkSpeedOk returns a tuple with the PcieSlotNvme1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme1linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme1linkSpeed(v string)`

SetPcieSlotNvme1linkSpeed sets PcieSlotNvme1linkSpeed field to given value.

### HasPcieSlotNvme1linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme1linkSpeed() bool`

HasPcieSlotNvme1linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme1optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme1optionRom() string`

GetPcieSlotNvme1optionRom returns the PcieSlotNvme1optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme1optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme1optionRomOk() (*string, bool)`

GetPcieSlotNvme1optionRomOk returns a tuple with the PcieSlotNvme1optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme1optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme1optionRom(v string)`

SetPcieSlotNvme1optionRom sets PcieSlotNvme1optionRom field to given value.

### HasPcieSlotNvme1optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme1optionRom() bool`

HasPcieSlotNvme1optionRom returns a boolean if a field has been set.

### GetPcieSlotNvme2linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme2linkSpeed() string`

GetPcieSlotNvme2linkSpeed returns the PcieSlotNvme2linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme2linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme2linkSpeedOk() (*string, bool)`

GetPcieSlotNvme2linkSpeedOk returns a tuple with the PcieSlotNvme2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme2linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme2linkSpeed(v string)`

SetPcieSlotNvme2linkSpeed sets PcieSlotNvme2linkSpeed field to given value.

### HasPcieSlotNvme2linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme2linkSpeed() bool`

HasPcieSlotNvme2linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme2optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme2optionRom() string`

GetPcieSlotNvme2optionRom returns the PcieSlotNvme2optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme2optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme2optionRomOk() (*string, bool)`

GetPcieSlotNvme2optionRomOk returns a tuple with the PcieSlotNvme2optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme2optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme2optionRom(v string)`

SetPcieSlotNvme2optionRom sets PcieSlotNvme2optionRom field to given value.

### HasPcieSlotNvme2optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme2optionRom() bool`

HasPcieSlotNvme2optionRom returns a boolean if a field has been set.

### GetPcieSlotNvme3linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme3linkSpeed() string`

GetPcieSlotNvme3linkSpeed returns the PcieSlotNvme3linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme3linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme3linkSpeedOk() (*string, bool)`

GetPcieSlotNvme3linkSpeedOk returns a tuple with the PcieSlotNvme3linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme3linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme3linkSpeed(v string)`

SetPcieSlotNvme3linkSpeed sets PcieSlotNvme3linkSpeed field to given value.

### HasPcieSlotNvme3linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme3linkSpeed() bool`

HasPcieSlotNvme3linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme3optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme3optionRom() string`

GetPcieSlotNvme3optionRom returns the PcieSlotNvme3optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme3optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme3optionRomOk() (*string, bool)`

GetPcieSlotNvme3optionRomOk returns a tuple with the PcieSlotNvme3optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme3optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme3optionRom(v string)`

SetPcieSlotNvme3optionRom sets PcieSlotNvme3optionRom field to given value.

### HasPcieSlotNvme3optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme3optionRom() bool`

HasPcieSlotNvme3optionRom returns a boolean if a field has been set.

### GetPcieSlotNvme4linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme4linkSpeed() string`

GetPcieSlotNvme4linkSpeed returns the PcieSlotNvme4linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme4linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme4linkSpeedOk() (*string, bool)`

GetPcieSlotNvme4linkSpeedOk returns a tuple with the PcieSlotNvme4linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme4linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme4linkSpeed(v string)`

SetPcieSlotNvme4linkSpeed sets PcieSlotNvme4linkSpeed field to given value.

### HasPcieSlotNvme4linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme4linkSpeed() bool`

HasPcieSlotNvme4linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme4optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme4optionRom() string`

GetPcieSlotNvme4optionRom returns the PcieSlotNvme4optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme4optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme4optionRomOk() (*string, bool)`

GetPcieSlotNvme4optionRomOk returns a tuple with the PcieSlotNvme4optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme4optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme4optionRom(v string)`

SetPcieSlotNvme4optionRom sets PcieSlotNvme4optionRom field to given value.

### HasPcieSlotNvme4optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme4optionRom() bool`

HasPcieSlotNvme4optionRom returns a boolean if a field has been set.

### GetPcieSlotNvme5linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme5linkSpeed() string`

GetPcieSlotNvme5linkSpeed returns the PcieSlotNvme5linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme5linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme5linkSpeedOk() (*string, bool)`

GetPcieSlotNvme5linkSpeedOk returns a tuple with the PcieSlotNvme5linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme5linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme5linkSpeed(v string)`

SetPcieSlotNvme5linkSpeed sets PcieSlotNvme5linkSpeed field to given value.

### HasPcieSlotNvme5linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme5linkSpeed() bool`

HasPcieSlotNvme5linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme5optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme5optionRom() string`

GetPcieSlotNvme5optionRom returns the PcieSlotNvme5optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme5optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme5optionRomOk() (*string, bool)`

GetPcieSlotNvme5optionRomOk returns a tuple with the PcieSlotNvme5optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme5optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme5optionRom(v string)`

SetPcieSlotNvme5optionRom sets PcieSlotNvme5optionRom field to given value.

### HasPcieSlotNvme5optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme5optionRom() bool`

HasPcieSlotNvme5optionRom returns a boolean if a field has been set.

### GetPcieSlotNvme6linkSpeed

`func (o *BiosPolicyAllOf) GetPcieSlotNvme6linkSpeed() string`

GetPcieSlotNvme6linkSpeed returns the PcieSlotNvme6linkSpeed field if non-nil, zero value otherwise.

### GetPcieSlotNvme6linkSpeedOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme6linkSpeedOk() (*string, bool)`

GetPcieSlotNvme6linkSpeedOk returns a tuple with the PcieSlotNvme6linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme6linkSpeed

`func (o *BiosPolicyAllOf) SetPcieSlotNvme6linkSpeed(v string)`

SetPcieSlotNvme6linkSpeed sets PcieSlotNvme6linkSpeed field to given value.

### HasPcieSlotNvme6linkSpeed

`func (o *BiosPolicyAllOf) HasPcieSlotNvme6linkSpeed() bool`

HasPcieSlotNvme6linkSpeed returns a boolean if a field has been set.

### GetPcieSlotNvme6optionRom

`func (o *BiosPolicyAllOf) GetPcieSlotNvme6optionRom() string`

GetPcieSlotNvme6optionRom returns the PcieSlotNvme6optionRom field if non-nil, zero value otherwise.

### GetPcieSlotNvme6optionRomOk

`func (o *BiosPolicyAllOf) GetPcieSlotNvme6optionRomOk() (*string, bool)`

GetPcieSlotNvme6optionRomOk returns a tuple with the PcieSlotNvme6optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSlotNvme6optionRom

`func (o *BiosPolicyAllOf) SetPcieSlotNvme6optionRom(v string)`

SetPcieSlotNvme6optionRom sets PcieSlotNvme6optionRom field to given value.

### HasPcieSlotNvme6optionRom

`func (o *BiosPolicyAllOf) HasPcieSlotNvme6optionRom() bool`

HasPcieSlotNvme6optionRom returns a boolean if a field has been set.

### GetPopSupport

`func (o *BiosPolicyAllOf) GetPopSupport() string`

GetPopSupport returns the PopSupport field if non-nil, zero value otherwise.

### GetPopSupportOk

`func (o *BiosPolicyAllOf) GetPopSupportOk() (*string, bool)`

GetPopSupportOk returns a tuple with the PopSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopSupport

`func (o *BiosPolicyAllOf) SetPopSupport(v string)`

SetPopSupport sets PopSupport field to given value.

### HasPopSupport

`func (o *BiosPolicyAllOf) HasPopSupport() bool`

HasPopSupport returns a boolean if a field has been set.

### GetPostErrorPause

`func (o *BiosPolicyAllOf) GetPostErrorPause() string`

GetPostErrorPause returns the PostErrorPause field if non-nil, zero value otherwise.

### GetPostErrorPauseOk

`func (o *BiosPolicyAllOf) GetPostErrorPauseOk() (*string, bool)`

GetPostErrorPauseOk returns a tuple with the PostErrorPause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostErrorPause

`func (o *BiosPolicyAllOf) SetPostErrorPause(v string)`

SetPostErrorPause sets PostErrorPause field to given value.

### HasPostErrorPause

`func (o *BiosPolicyAllOf) HasPostErrorPause() bool`

HasPostErrorPause returns a boolean if a field has been set.

### GetPostPackageRepair

`func (o *BiosPolicyAllOf) GetPostPackageRepair() string`

GetPostPackageRepair returns the PostPackageRepair field if non-nil, zero value otherwise.

### GetPostPackageRepairOk

`func (o *BiosPolicyAllOf) GetPostPackageRepairOk() (*string, bool)`

GetPostPackageRepairOk returns a tuple with the PostPackageRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPackageRepair

`func (o *BiosPolicyAllOf) SetPostPackageRepair(v string)`

SetPostPackageRepair sets PostPackageRepair field to given value.

### HasPostPackageRepair

`func (o *BiosPolicyAllOf) HasPostPackageRepair() bool`

HasPostPackageRepair returns a boolean if a field has been set.

### GetProcessorC1e

`func (o *BiosPolicyAllOf) GetProcessorC1e() string`

GetProcessorC1e returns the ProcessorC1e field if non-nil, zero value otherwise.

### GetProcessorC1eOk

`func (o *BiosPolicyAllOf) GetProcessorC1eOk() (*string, bool)`

GetProcessorC1eOk returns a tuple with the ProcessorC1e field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorC1e

`func (o *BiosPolicyAllOf) SetProcessorC1e(v string)`

SetProcessorC1e sets ProcessorC1e field to given value.

### HasProcessorC1e

`func (o *BiosPolicyAllOf) HasProcessorC1e() bool`

HasProcessorC1e returns a boolean if a field has been set.

### GetProcessorC3report

`func (o *BiosPolicyAllOf) GetProcessorC3report() string`

GetProcessorC3report returns the ProcessorC3report field if non-nil, zero value otherwise.

### GetProcessorC3reportOk

`func (o *BiosPolicyAllOf) GetProcessorC3reportOk() (*string, bool)`

GetProcessorC3reportOk returns a tuple with the ProcessorC3report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorC3report

`func (o *BiosPolicyAllOf) SetProcessorC3report(v string)`

SetProcessorC3report sets ProcessorC3report field to given value.

### HasProcessorC3report

`func (o *BiosPolicyAllOf) HasProcessorC3report() bool`

HasProcessorC3report returns a boolean if a field has been set.

### GetProcessorC6report

`func (o *BiosPolicyAllOf) GetProcessorC6report() string`

GetProcessorC6report returns the ProcessorC6report field if non-nil, zero value otherwise.

### GetProcessorC6reportOk

`func (o *BiosPolicyAllOf) GetProcessorC6reportOk() (*string, bool)`

GetProcessorC6reportOk returns a tuple with the ProcessorC6report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorC6report

`func (o *BiosPolicyAllOf) SetProcessorC6report(v string)`

SetProcessorC6report sets ProcessorC6report field to given value.

### HasProcessorC6report

`func (o *BiosPolicyAllOf) HasProcessorC6report() bool`

HasProcessorC6report returns a boolean if a field has been set.

### GetProcessorCstate

`func (o *BiosPolicyAllOf) GetProcessorCstate() string`

GetProcessorCstate returns the ProcessorCstate field if non-nil, zero value otherwise.

### GetProcessorCstateOk

`func (o *BiosPolicyAllOf) GetProcessorCstateOk() (*string, bool)`

GetProcessorCstateOk returns a tuple with the ProcessorCstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCstate

`func (o *BiosPolicyAllOf) SetProcessorCstate(v string)`

SetProcessorCstate sets ProcessorCstate field to given value.

### HasProcessorCstate

`func (o *BiosPolicyAllOf) HasProcessorCstate() bool`

HasProcessorCstate returns a boolean if a field has been set.

### GetPsata

`func (o *BiosPolicyAllOf) GetPsata() string`

GetPsata returns the Psata field if non-nil, zero value otherwise.

### GetPsataOk

`func (o *BiosPolicyAllOf) GetPsataOk() (*string, bool)`

GetPsataOk returns a tuple with the Psata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsata

`func (o *BiosPolicyAllOf) SetPsata(v string)`

SetPsata sets Psata field to given value.

### HasPsata

`func (o *BiosPolicyAllOf) HasPsata() bool`

HasPsata returns a boolean if a field has been set.

### GetPstateCoordType

`func (o *BiosPolicyAllOf) GetPstateCoordType() string`

GetPstateCoordType returns the PstateCoordType field if non-nil, zero value otherwise.

### GetPstateCoordTypeOk

`func (o *BiosPolicyAllOf) GetPstateCoordTypeOk() (*string, bool)`

GetPstateCoordTypeOk returns a tuple with the PstateCoordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPstateCoordType

`func (o *BiosPolicyAllOf) SetPstateCoordType(v string)`

SetPstateCoordType sets PstateCoordType field to given value.

### HasPstateCoordType

`func (o *BiosPolicyAllOf) HasPstateCoordType() bool`

HasPstateCoordType returns a boolean if a field has been set.

### GetPuttyKeyPad

`func (o *BiosPolicyAllOf) GetPuttyKeyPad() string`

GetPuttyKeyPad returns the PuttyKeyPad field if non-nil, zero value otherwise.

### GetPuttyKeyPadOk

`func (o *BiosPolicyAllOf) GetPuttyKeyPadOk() (*string, bool)`

GetPuttyKeyPadOk returns a tuple with the PuttyKeyPad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuttyKeyPad

`func (o *BiosPolicyAllOf) SetPuttyKeyPad(v string)`

SetPuttyKeyPad sets PuttyKeyPad field to given value.

### HasPuttyKeyPad

`func (o *BiosPolicyAllOf) HasPuttyKeyPad() bool`

HasPuttyKeyPad returns a boolean if a field has been set.

### GetPwrPerfTuning

`func (o *BiosPolicyAllOf) GetPwrPerfTuning() string`

GetPwrPerfTuning returns the PwrPerfTuning field if non-nil, zero value otherwise.

### GetPwrPerfTuningOk

`func (o *BiosPolicyAllOf) GetPwrPerfTuningOk() (*string, bool)`

GetPwrPerfTuningOk returns a tuple with the PwrPerfTuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwrPerfTuning

`func (o *BiosPolicyAllOf) SetPwrPerfTuning(v string)`

SetPwrPerfTuning sets PwrPerfTuning field to given value.

### HasPwrPerfTuning

`func (o *BiosPolicyAllOf) HasPwrPerfTuning() bool`

HasPwrPerfTuning returns a boolean if a field has been set.

### GetQpiLinkFrequency

`func (o *BiosPolicyAllOf) GetQpiLinkFrequency() string`

GetQpiLinkFrequency returns the QpiLinkFrequency field if non-nil, zero value otherwise.

### GetQpiLinkFrequencyOk

`func (o *BiosPolicyAllOf) GetQpiLinkFrequencyOk() (*string, bool)`

GetQpiLinkFrequencyOk returns a tuple with the QpiLinkFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpiLinkFrequency

`func (o *BiosPolicyAllOf) SetQpiLinkFrequency(v string)`

SetQpiLinkFrequency sets QpiLinkFrequency field to given value.

### HasQpiLinkFrequency

`func (o *BiosPolicyAllOf) HasQpiLinkFrequency() bool`

HasQpiLinkFrequency returns a boolean if a field has been set.

### GetQpiLinkSpeed

`func (o *BiosPolicyAllOf) GetQpiLinkSpeed() string`

GetQpiLinkSpeed returns the QpiLinkSpeed field if non-nil, zero value otherwise.

### GetQpiLinkSpeedOk

`func (o *BiosPolicyAllOf) GetQpiLinkSpeedOk() (*string, bool)`

GetQpiLinkSpeedOk returns a tuple with the QpiLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpiLinkSpeed

`func (o *BiosPolicyAllOf) SetQpiLinkSpeed(v string)`

SetQpiLinkSpeed sets QpiLinkSpeed field to given value.

### HasQpiLinkSpeed

`func (o *BiosPolicyAllOf) HasQpiLinkSpeed() bool`

HasQpiLinkSpeed returns a boolean if a field has been set.

### GetQpiSnoopMode

`func (o *BiosPolicyAllOf) GetQpiSnoopMode() string`

GetQpiSnoopMode returns the QpiSnoopMode field if non-nil, zero value otherwise.

### GetQpiSnoopModeOk

`func (o *BiosPolicyAllOf) GetQpiSnoopModeOk() (*string, bool)`

GetQpiSnoopModeOk returns a tuple with the QpiSnoopMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQpiSnoopMode

`func (o *BiosPolicyAllOf) SetQpiSnoopMode(v string)`

SetQpiSnoopMode sets QpiSnoopMode field to given value.

### HasQpiSnoopMode

`func (o *BiosPolicyAllOf) HasQpiSnoopMode() bool`

HasQpiSnoopMode returns a boolean if a field has been set.

### GetRankInterLeave

`func (o *BiosPolicyAllOf) GetRankInterLeave() string`

GetRankInterLeave returns the RankInterLeave field if non-nil, zero value otherwise.

### GetRankInterLeaveOk

`func (o *BiosPolicyAllOf) GetRankInterLeaveOk() (*string, bool)`

GetRankInterLeaveOk returns a tuple with the RankInterLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankInterLeave

`func (o *BiosPolicyAllOf) SetRankInterLeave(v string)`

SetRankInterLeave sets RankInterLeave field to given value.

### HasRankInterLeave

`func (o *BiosPolicyAllOf) HasRankInterLeave() bool`

HasRankInterLeave returns a boolean if a field has been set.

### GetRedirectionAfterPost

`func (o *BiosPolicyAllOf) GetRedirectionAfterPost() string`

GetRedirectionAfterPost returns the RedirectionAfterPost field if non-nil, zero value otherwise.

### GetRedirectionAfterPostOk

`func (o *BiosPolicyAllOf) GetRedirectionAfterPostOk() (*string, bool)`

GetRedirectionAfterPostOk returns a tuple with the RedirectionAfterPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionAfterPost

`func (o *BiosPolicyAllOf) SetRedirectionAfterPost(v string)`

SetRedirectionAfterPost sets RedirectionAfterPost field to given value.

### HasRedirectionAfterPost

`func (o *BiosPolicyAllOf) HasRedirectionAfterPost() bool`

HasRedirectionAfterPost returns a boolean if a field has been set.

### GetSataModeSelect

`func (o *BiosPolicyAllOf) GetSataModeSelect() string`

GetSataModeSelect returns the SataModeSelect field if non-nil, zero value otherwise.

### GetSataModeSelectOk

`func (o *BiosPolicyAllOf) GetSataModeSelectOk() (*string, bool)`

GetSataModeSelectOk returns a tuple with the SataModeSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSataModeSelect

`func (o *BiosPolicyAllOf) SetSataModeSelect(v string)`

SetSataModeSelect sets SataModeSelect field to given value.

### HasSataModeSelect

`func (o *BiosPolicyAllOf) HasSataModeSelect() bool`

HasSataModeSelect returns a boolean if a field has been set.

### GetSelectMemoryRasConfiguration

`func (o *BiosPolicyAllOf) GetSelectMemoryRasConfiguration() string`

GetSelectMemoryRasConfiguration returns the SelectMemoryRasConfiguration field if non-nil, zero value otherwise.

### GetSelectMemoryRasConfigurationOk

`func (o *BiosPolicyAllOf) GetSelectMemoryRasConfigurationOk() (*string, bool)`

GetSelectMemoryRasConfigurationOk returns a tuple with the SelectMemoryRasConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectMemoryRasConfiguration

`func (o *BiosPolicyAllOf) SetSelectMemoryRasConfiguration(v string)`

SetSelectMemoryRasConfiguration sets SelectMemoryRasConfiguration field to given value.

### HasSelectMemoryRasConfiguration

`func (o *BiosPolicyAllOf) HasSelectMemoryRasConfiguration() bool`

HasSelectMemoryRasConfiguration returns a boolean if a field has been set.

### GetSelectPprType

`func (o *BiosPolicyAllOf) GetSelectPprType() string`

GetSelectPprType returns the SelectPprType field if non-nil, zero value otherwise.

### GetSelectPprTypeOk

`func (o *BiosPolicyAllOf) GetSelectPprTypeOk() (*string, bool)`

GetSelectPprTypeOk returns a tuple with the SelectPprType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPprType

`func (o *BiosPolicyAllOf) SetSelectPprType(v string)`

SetSelectPprType sets SelectPprType field to given value.

### HasSelectPprType

`func (o *BiosPolicyAllOf) HasSelectPprType() bool`

HasSelectPprType returns a boolean if a field has been set.

### GetSerialPortAenable

`func (o *BiosPolicyAllOf) GetSerialPortAenable() string`

GetSerialPortAenable returns the SerialPortAenable field if non-nil, zero value otherwise.

### GetSerialPortAenableOk

`func (o *BiosPolicyAllOf) GetSerialPortAenableOk() (*string, bool)`

GetSerialPortAenableOk returns a tuple with the SerialPortAenable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialPortAenable

`func (o *BiosPolicyAllOf) SetSerialPortAenable(v string)`

SetSerialPortAenable sets SerialPortAenable field to given value.

### HasSerialPortAenable

`func (o *BiosPolicyAllOf) HasSerialPortAenable() bool`

HasSerialPortAenable returns a boolean if a field has been set.

### GetSev

`func (o *BiosPolicyAllOf) GetSev() string`

GetSev returns the Sev field if non-nil, zero value otherwise.

### GetSevOk

`func (o *BiosPolicyAllOf) GetSevOk() (*string, bool)`

GetSevOk returns a tuple with the Sev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSev

`func (o *BiosPolicyAllOf) SetSev(v string)`

SetSev sets Sev field to given value.

### HasSev

`func (o *BiosPolicyAllOf) HasSev() bool`

HasSev returns a boolean if a field has been set.

### GetSgxAutoRegistrationAgent

`func (o *BiosPolicyAllOf) GetSgxAutoRegistrationAgent() string`

GetSgxAutoRegistrationAgent returns the SgxAutoRegistrationAgent field if non-nil, zero value otherwise.

### GetSgxAutoRegistrationAgentOk

`func (o *BiosPolicyAllOf) GetSgxAutoRegistrationAgentOk() (*string, bool)`

GetSgxAutoRegistrationAgentOk returns a tuple with the SgxAutoRegistrationAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxAutoRegistrationAgent

`func (o *BiosPolicyAllOf) SetSgxAutoRegistrationAgent(v string)`

SetSgxAutoRegistrationAgent sets SgxAutoRegistrationAgent field to given value.

### HasSgxAutoRegistrationAgent

`func (o *BiosPolicyAllOf) HasSgxAutoRegistrationAgent() bool`

HasSgxAutoRegistrationAgent returns a boolean if a field has been set.

### GetSgxEpoch0

`func (o *BiosPolicyAllOf) GetSgxEpoch0() string`

GetSgxEpoch0 returns the SgxEpoch0 field if non-nil, zero value otherwise.

### GetSgxEpoch0Ok

`func (o *BiosPolicyAllOf) GetSgxEpoch0Ok() (*string, bool)`

GetSgxEpoch0Ok returns a tuple with the SgxEpoch0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxEpoch0

`func (o *BiosPolicyAllOf) SetSgxEpoch0(v string)`

SetSgxEpoch0 sets SgxEpoch0 field to given value.

### HasSgxEpoch0

`func (o *BiosPolicyAllOf) HasSgxEpoch0() bool`

HasSgxEpoch0 returns a boolean if a field has been set.

### GetSgxEpoch1

`func (o *BiosPolicyAllOf) GetSgxEpoch1() string`

GetSgxEpoch1 returns the SgxEpoch1 field if non-nil, zero value otherwise.

### GetSgxEpoch1Ok

`func (o *BiosPolicyAllOf) GetSgxEpoch1Ok() (*string, bool)`

GetSgxEpoch1Ok returns a tuple with the SgxEpoch1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxEpoch1

`func (o *BiosPolicyAllOf) SetSgxEpoch1(v string)`

SetSgxEpoch1 sets SgxEpoch1 field to given value.

### HasSgxEpoch1

`func (o *BiosPolicyAllOf) HasSgxEpoch1() bool`

HasSgxEpoch1 returns a boolean if a field has been set.

### GetSgxFactoryReset

`func (o *BiosPolicyAllOf) GetSgxFactoryReset() string`

GetSgxFactoryReset returns the SgxFactoryReset field if non-nil, zero value otherwise.

### GetSgxFactoryResetOk

`func (o *BiosPolicyAllOf) GetSgxFactoryResetOk() (*string, bool)`

GetSgxFactoryResetOk returns a tuple with the SgxFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxFactoryReset

`func (o *BiosPolicyAllOf) SetSgxFactoryReset(v string)`

SetSgxFactoryReset sets SgxFactoryReset field to given value.

### HasSgxFactoryReset

`func (o *BiosPolicyAllOf) HasSgxFactoryReset() bool`

HasSgxFactoryReset returns a boolean if a field has been set.

### GetSgxLePubKeyHash0

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash0() string`

GetSgxLePubKeyHash0 returns the SgxLePubKeyHash0 field if non-nil, zero value otherwise.

### GetSgxLePubKeyHash0Ok

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash0Ok() (*string, bool)`

GetSgxLePubKeyHash0Ok returns a tuple with the SgxLePubKeyHash0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxLePubKeyHash0

`func (o *BiosPolicyAllOf) SetSgxLePubKeyHash0(v string)`

SetSgxLePubKeyHash0 sets SgxLePubKeyHash0 field to given value.

### HasSgxLePubKeyHash0

`func (o *BiosPolicyAllOf) HasSgxLePubKeyHash0() bool`

HasSgxLePubKeyHash0 returns a boolean if a field has been set.

### GetSgxLePubKeyHash1

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash1() string`

GetSgxLePubKeyHash1 returns the SgxLePubKeyHash1 field if non-nil, zero value otherwise.

### GetSgxLePubKeyHash1Ok

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash1Ok() (*string, bool)`

GetSgxLePubKeyHash1Ok returns a tuple with the SgxLePubKeyHash1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxLePubKeyHash1

`func (o *BiosPolicyAllOf) SetSgxLePubKeyHash1(v string)`

SetSgxLePubKeyHash1 sets SgxLePubKeyHash1 field to given value.

### HasSgxLePubKeyHash1

`func (o *BiosPolicyAllOf) HasSgxLePubKeyHash1() bool`

HasSgxLePubKeyHash1 returns a boolean if a field has been set.

### GetSgxLePubKeyHash2

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash2() string`

GetSgxLePubKeyHash2 returns the SgxLePubKeyHash2 field if non-nil, zero value otherwise.

### GetSgxLePubKeyHash2Ok

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash2Ok() (*string, bool)`

GetSgxLePubKeyHash2Ok returns a tuple with the SgxLePubKeyHash2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxLePubKeyHash2

`func (o *BiosPolicyAllOf) SetSgxLePubKeyHash2(v string)`

SetSgxLePubKeyHash2 sets SgxLePubKeyHash2 field to given value.

### HasSgxLePubKeyHash2

`func (o *BiosPolicyAllOf) HasSgxLePubKeyHash2() bool`

HasSgxLePubKeyHash2 returns a boolean if a field has been set.

### GetSgxLePubKeyHash3

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash3() string`

GetSgxLePubKeyHash3 returns the SgxLePubKeyHash3 field if non-nil, zero value otherwise.

### GetSgxLePubKeyHash3Ok

`func (o *BiosPolicyAllOf) GetSgxLePubKeyHash3Ok() (*string, bool)`

GetSgxLePubKeyHash3Ok returns a tuple with the SgxLePubKeyHash3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxLePubKeyHash3

`func (o *BiosPolicyAllOf) SetSgxLePubKeyHash3(v string)`

SetSgxLePubKeyHash3 sets SgxLePubKeyHash3 field to given value.

### HasSgxLePubKeyHash3

`func (o *BiosPolicyAllOf) HasSgxLePubKeyHash3() bool`

HasSgxLePubKeyHash3 returns a boolean if a field has been set.

### GetSgxLeWr

`func (o *BiosPolicyAllOf) GetSgxLeWr() string`

GetSgxLeWr returns the SgxLeWr field if non-nil, zero value otherwise.

### GetSgxLeWrOk

`func (o *BiosPolicyAllOf) GetSgxLeWrOk() (*string, bool)`

GetSgxLeWrOk returns a tuple with the SgxLeWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxLeWr

`func (o *BiosPolicyAllOf) SetSgxLeWr(v string)`

SetSgxLeWr sets SgxLeWr field to given value.

### HasSgxLeWr

`func (o *BiosPolicyAllOf) HasSgxLeWr() bool`

HasSgxLeWr returns a boolean if a field has been set.

### GetSgxPackageInfoInBandAccess

`func (o *BiosPolicyAllOf) GetSgxPackageInfoInBandAccess() string`

GetSgxPackageInfoInBandAccess returns the SgxPackageInfoInBandAccess field if non-nil, zero value otherwise.

### GetSgxPackageInfoInBandAccessOk

`func (o *BiosPolicyAllOf) GetSgxPackageInfoInBandAccessOk() (*string, bool)`

GetSgxPackageInfoInBandAccessOk returns a tuple with the SgxPackageInfoInBandAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxPackageInfoInBandAccess

`func (o *BiosPolicyAllOf) SetSgxPackageInfoInBandAccess(v string)`

SetSgxPackageInfoInBandAccess sets SgxPackageInfoInBandAccess field to given value.

### HasSgxPackageInfoInBandAccess

`func (o *BiosPolicyAllOf) HasSgxPackageInfoInBandAccess() bool`

HasSgxPackageInfoInBandAccess returns a boolean if a field has been set.

### GetSgxQos

`func (o *BiosPolicyAllOf) GetSgxQos() string`

GetSgxQos returns the SgxQos field if non-nil, zero value otherwise.

### GetSgxQosOk

`func (o *BiosPolicyAllOf) GetSgxQosOk() (*string, bool)`

GetSgxQosOk returns a tuple with the SgxQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgxQos

`func (o *BiosPolicyAllOf) SetSgxQos(v string)`

SetSgxQos sets SgxQos field to given value.

### HasSgxQos

`func (o *BiosPolicyAllOf) HasSgxQos() bool`

HasSgxQos returns a boolean if a field has been set.

### GetSinglePctlEnable

`func (o *BiosPolicyAllOf) GetSinglePctlEnable() string`

GetSinglePctlEnable returns the SinglePctlEnable field if non-nil, zero value otherwise.

### GetSinglePctlEnableOk

`func (o *BiosPolicyAllOf) GetSinglePctlEnableOk() (*string, bool)`

GetSinglePctlEnableOk returns a tuple with the SinglePctlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinglePctlEnable

`func (o *BiosPolicyAllOf) SetSinglePctlEnable(v string)`

SetSinglePctlEnable sets SinglePctlEnable field to given value.

### HasSinglePctlEnable

`func (o *BiosPolicyAllOf) HasSinglePctlEnable() bool`

HasSinglePctlEnable returns a boolean if a field has been set.

### GetSlot10linkSpeed

`func (o *BiosPolicyAllOf) GetSlot10linkSpeed() string`

GetSlot10linkSpeed returns the Slot10linkSpeed field if non-nil, zero value otherwise.

### GetSlot10linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot10linkSpeedOk() (*string, bool)`

GetSlot10linkSpeedOk returns a tuple with the Slot10linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot10linkSpeed

`func (o *BiosPolicyAllOf) SetSlot10linkSpeed(v string)`

SetSlot10linkSpeed sets Slot10linkSpeed field to given value.

### HasSlot10linkSpeed

`func (o *BiosPolicyAllOf) HasSlot10linkSpeed() bool`

HasSlot10linkSpeed returns a boolean if a field has been set.

### GetSlot10state

`func (o *BiosPolicyAllOf) GetSlot10state() string`

GetSlot10state returns the Slot10state field if non-nil, zero value otherwise.

### GetSlot10stateOk

`func (o *BiosPolicyAllOf) GetSlot10stateOk() (*string, bool)`

GetSlot10stateOk returns a tuple with the Slot10state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot10state

`func (o *BiosPolicyAllOf) SetSlot10state(v string)`

SetSlot10state sets Slot10state field to given value.

### HasSlot10state

`func (o *BiosPolicyAllOf) HasSlot10state() bool`

HasSlot10state returns a boolean if a field has been set.

### GetSlot11linkSpeed

`func (o *BiosPolicyAllOf) GetSlot11linkSpeed() string`

GetSlot11linkSpeed returns the Slot11linkSpeed field if non-nil, zero value otherwise.

### GetSlot11linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot11linkSpeedOk() (*string, bool)`

GetSlot11linkSpeedOk returns a tuple with the Slot11linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot11linkSpeed

`func (o *BiosPolicyAllOf) SetSlot11linkSpeed(v string)`

SetSlot11linkSpeed sets Slot11linkSpeed field to given value.

### HasSlot11linkSpeed

`func (o *BiosPolicyAllOf) HasSlot11linkSpeed() bool`

HasSlot11linkSpeed returns a boolean if a field has been set.

### GetSlot11state

`func (o *BiosPolicyAllOf) GetSlot11state() string`

GetSlot11state returns the Slot11state field if non-nil, zero value otherwise.

### GetSlot11stateOk

`func (o *BiosPolicyAllOf) GetSlot11stateOk() (*string, bool)`

GetSlot11stateOk returns a tuple with the Slot11state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot11state

`func (o *BiosPolicyAllOf) SetSlot11state(v string)`

SetSlot11state sets Slot11state field to given value.

### HasSlot11state

`func (o *BiosPolicyAllOf) HasSlot11state() bool`

HasSlot11state returns a boolean if a field has been set.

### GetSlot12linkSpeed

`func (o *BiosPolicyAllOf) GetSlot12linkSpeed() string`

GetSlot12linkSpeed returns the Slot12linkSpeed field if non-nil, zero value otherwise.

### GetSlot12linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot12linkSpeedOk() (*string, bool)`

GetSlot12linkSpeedOk returns a tuple with the Slot12linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot12linkSpeed

`func (o *BiosPolicyAllOf) SetSlot12linkSpeed(v string)`

SetSlot12linkSpeed sets Slot12linkSpeed field to given value.

### HasSlot12linkSpeed

`func (o *BiosPolicyAllOf) HasSlot12linkSpeed() bool`

HasSlot12linkSpeed returns a boolean if a field has been set.

### GetSlot12state

`func (o *BiosPolicyAllOf) GetSlot12state() string`

GetSlot12state returns the Slot12state field if non-nil, zero value otherwise.

### GetSlot12stateOk

`func (o *BiosPolicyAllOf) GetSlot12stateOk() (*string, bool)`

GetSlot12stateOk returns a tuple with the Slot12state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot12state

`func (o *BiosPolicyAllOf) SetSlot12state(v string)`

SetSlot12state sets Slot12state field to given value.

### HasSlot12state

`func (o *BiosPolicyAllOf) HasSlot12state() bool`

HasSlot12state returns a boolean if a field has been set.

### GetSlot13state

`func (o *BiosPolicyAllOf) GetSlot13state() string`

GetSlot13state returns the Slot13state field if non-nil, zero value otherwise.

### GetSlot13stateOk

`func (o *BiosPolicyAllOf) GetSlot13stateOk() (*string, bool)`

GetSlot13stateOk returns a tuple with the Slot13state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot13state

`func (o *BiosPolicyAllOf) SetSlot13state(v string)`

SetSlot13state sets Slot13state field to given value.

### HasSlot13state

`func (o *BiosPolicyAllOf) HasSlot13state() bool`

HasSlot13state returns a boolean if a field has been set.

### GetSlot14state

`func (o *BiosPolicyAllOf) GetSlot14state() string`

GetSlot14state returns the Slot14state field if non-nil, zero value otherwise.

### GetSlot14stateOk

`func (o *BiosPolicyAllOf) GetSlot14stateOk() (*string, bool)`

GetSlot14stateOk returns a tuple with the Slot14state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot14state

`func (o *BiosPolicyAllOf) SetSlot14state(v string)`

SetSlot14state sets Slot14state field to given value.

### HasSlot14state

`func (o *BiosPolicyAllOf) HasSlot14state() bool`

HasSlot14state returns a boolean if a field has been set.

### GetSlot1linkSpeed

`func (o *BiosPolicyAllOf) GetSlot1linkSpeed() string`

GetSlot1linkSpeed returns the Slot1linkSpeed field if non-nil, zero value otherwise.

### GetSlot1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot1linkSpeedOk() (*string, bool)`

GetSlot1linkSpeedOk returns a tuple with the Slot1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot1linkSpeed

`func (o *BiosPolicyAllOf) SetSlot1linkSpeed(v string)`

SetSlot1linkSpeed sets Slot1linkSpeed field to given value.

### HasSlot1linkSpeed

`func (o *BiosPolicyAllOf) HasSlot1linkSpeed() bool`

HasSlot1linkSpeed returns a boolean if a field has been set.

### GetSlot1state

`func (o *BiosPolicyAllOf) GetSlot1state() string`

GetSlot1state returns the Slot1state field if non-nil, zero value otherwise.

### GetSlot1stateOk

`func (o *BiosPolicyAllOf) GetSlot1stateOk() (*string, bool)`

GetSlot1stateOk returns a tuple with the Slot1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot1state

`func (o *BiosPolicyAllOf) SetSlot1state(v string)`

SetSlot1state sets Slot1state field to given value.

### HasSlot1state

`func (o *BiosPolicyAllOf) HasSlot1state() bool`

HasSlot1state returns a boolean if a field has been set.

### GetSlot2linkSpeed

`func (o *BiosPolicyAllOf) GetSlot2linkSpeed() string`

GetSlot2linkSpeed returns the Slot2linkSpeed field if non-nil, zero value otherwise.

### GetSlot2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot2linkSpeedOk() (*string, bool)`

GetSlot2linkSpeedOk returns a tuple with the Slot2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot2linkSpeed

`func (o *BiosPolicyAllOf) SetSlot2linkSpeed(v string)`

SetSlot2linkSpeed sets Slot2linkSpeed field to given value.

### HasSlot2linkSpeed

`func (o *BiosPolicyAllOf) HasSlot2linkSpeed() bool`

HasSlot2linkSpeed returns a boolean if a field has been set.

### GetSlot2state

`func (o *BiosPolicyAllOf) GetSlot2state() string`

GetSlot2state returns the Slot2state field if non-nil, zero value otherwise.

### GetSlot2stateOk

`func (o *BiosPolicyAllOf) GetSlot2stateOk() (*string, bool)`

GetSlot2stateOk returns a tuple with the Slot2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot2state

`func (o *BiosPolicyAllOf) SetSlot2state(v string)`

SetSlot2state sets Slot2state field to given value.

### HasSlot2state

`func (o *BiosPolicyAllOf) HasSlot2state() bool`

HasSlot2state returns a boolean if a field has been set.

### GetSlot3linkSpeed

`func (o *BiosPolicyAllOf) GetSlot3linkSpeed() string`

GetSlot3linkSpeed returns the Slot3linkSpeed field if non-nil, zero value otherwise.

### GetSlot3linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot3linkSpeedOk() (*string, bool)`

GetSlot3linkSpeedOk returns a tuple with the Slot3linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot3linkSpeed

`func (o *BiosPolicyAllOf) SetSlot3linkSpeed(v string)`

SetSlot3linkSpeed sets Slot3linkSpeed field to given value.

### HasSlot3linkSpeed

`func (o *BiosPolicyAllOf) HasSlot3linkSpeed() bool`

HasSlot3linkSpeed returns a boolean if a field has been set.

### GetSlot3state

`func (o *BiosPolicyAllOf) GetSlot3state() string`

GetSlot3state returns the Slot3state field if non-nil, zero value otherwise.

### GetSlot3stateOk

`func (o *BiosPolicyAllOf) GetSlot3stateOk() (*string, bool)`

GetSlot3stateOk returns a tuple with the Slot3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot3state

`func (o *BiosPolicyAllOf) SetSlot3state(v string)`

SetSlot3state sets Slot3state field to given value.

### HasSlot3state

`func (o *BiosPolicyAllOf) HasSlot3state() bool`

HasSlot3state returns a boolean if a field has been set.

### GetSlot4linkSpeed

`func (o *BiosPolicyAllOf) GetSlot4linkSpeed() string`

GetSlot4linkSpeed returns the Slot4linkSpeed field if non-nil, zero value otherwise.

### GetSlot4linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot4linkSpeedOk() (*string, bool)`

GetSlot4linkSpeedOk returns a tuple with the Slot4linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot4linkSpeed

`func (o *BiosPolicyAllOf) SetSlot4linkSpeed(v string)`

SetSlot4linkSpeed sets Slot4linkSpeed field to given value.

### HasSlot4linkSpeed

`func (o *BiosPolicyAllOf) HasSlot4linkSpeed() bool`

HasSlot4linkSpeed returns a boolean if a field has been set.

### GetSlot4state

`func (o *BiosPolicyAllOf) GetSlot4state() string`

GetSlot4state returns the Slot4state field if non-nil, zero value otherwise.

### GetSlot4stateOk

`func (o *BiosPolicyAllOf) GetSlot4stateOk() (*string, bool)`

GetSlot4stateOk returns a tuple with the Slot4state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot4state

`func (o *BiosPolicyAllOf) SetSlot4state(v string)`

SetSlot4state sets Slot4state field to given value.

### HasSlot4state

`func (o *BiosPolicyAllOf) HasSlot4state() bool`

HasSlot4state returns a boolean if a field has been set.

### GetSlot5linkSpeed

`func (o *BiosPolicyAllOf) GetSlot5linkSpeed() string`

GetSlot5linkSpeed returns the Slot5linkSpeed field if non-nil, zero value otherwise.

### GetSlot5linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot5linkSpeedOk() (*string, bool)`

GetSlot5linkSpeedOk returns a tuple with the Slot5linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot5linkSpeed

`func (o *BiosPolicyAllOf) SetSlot5linkSpeed(v string)`

SetSlot5linkSpeed sets Slot5linkSpeed field to given value.

### HasSlot5linkSpeed

`func (o *BiosPolicyAllOf) HasSlot5linkSpeed() bool`

HasSlot5linkSpeed returns a boolean if a field has been set.

### GetSlot5state

`func (o *BiosPolicyAllOf) GetSlot5state() string`

GetSlot5state returns the Slot5state field if non-nil, zero value otherwise.

### GetSlot5stateOk

`func (o *BiosPolicyAllOf) GetSlot5stateOk() (*string, bool)`

GetSlot5stateOk returns a tuple with the Slot5state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot5state

`func (o *BiosPolicyAllOf) SetSlot5state(v string)`

SetSlot5state sets Slot5state field to given value.

### HasSlot5state

`func (o *BiosPolicyAllOf) HasSlot5state() bool`

HasSlot5state returns a boolean if a field has been set.

### GetSlot6linkSpeed

`func (o *BiosPolicyAllOf) GetSlot6linkSpeed() string`

GetSlot6linkSpeed returns the Slot6linkSpeed field if non-nil, zero value otherwise.

### GetSlot6linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot6linkSpeedOk() (*string, bool)`

GetSlot6linkSpeedOk returns a tuple with the Slot6linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot6linkSpeed

`func (o *BiosPolicyAllOf) SetSlot6linkSpeed(v string)`

SetSlot6linkSpeed sets Slot6linkSpeed field to given value.

### HasSlot6linkSpeed

`func (o *BiosPolicyAllOf) HasSlot6linkSpeed() bool`

HasSlot6linkSpeed returns a boolean if a field has been set.

### GetSlot6state

`func (o *BiosPolicyAllOf) GetSlot6state() string`

GetSlot6state returns the Slot6state field if non-nil, zero value otherwise.

### GetSlot6stateOk

`func (o *BiosPolicyAllOf) GetSlot6stateOk() (*string, bool)`

GetSlot6stateOk returns a tuple with the Slot6state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot6state

`func (o *BiosPolicyAllOf) SetSlot6state(v string)`

SetSlot6state sets Slot6state field to given value.

### HasSlot6state

`func (o *BiosPolicyAllOf) HasSlot6state() bool`

HasSlot6state returns a boolean if a field has been set.

### GetSlot7linkSpeed

`func (o *BiosPolicyAllOf) GetSlot7linkSpeed() string`

GetSlot7linkSpeed returns the Slot7linkSpeed field if non-nil, zero value otherwise.

### GetSlot7linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot7linkSpeedOk() (*string, bool)`

GetSlot7linkSpeedOk returns a tuple with the Slot7linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot7linkSpeed

`func (o *BiosPolicyAllOf) SetSlot7linkSpeed(v string)`

SetSlot7linkSpeed sets Slot7linkSpeed field to given value.

### HasSlot7linkSpeed

`func (o *BiosPolicyAllOf) HasSlot7linkSpeed() bool`

HasSlot7linkSpeed returns a boolean if a field has been set.

### GetSlot7state

`func (o *BiosPolicyAllOf) GetSlot7state() string`

GetSlot7state returns the Slot7state field if non-nil, zero value otherwise.

### GetSlot7stateOk

`func (o *BiosPolicyAllOf) GetSlot7stateOk() (*string, bool)`

GetSlot7stateOk returns a tuple with the Slot7state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot7state

`func (o *BiosPolicyAllOf) SetSlot7state(v string)`

SetSlot7state sets Slot7state field to given value.

### HasSlot7state

`func (o *BiosPolicyAllOf) HasSlot7state() bool`

HasSlot7state returns a boolean if a field has been set.

### GetSlot8linkSpeed

`func (o *BiosPolicyAllOf) GetSlot8linkSpeed() string`

GetSlot8linkSpeed returns the Slot8linkSpeed field if non-nil, zero value otherwise.

### GetSlot8linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot8linkSpeedOk() (*string, bool)`

GetSlot8linkSpeedOk returns a tuple with the Slot8linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot8linkSpeed

`func (o *BiosPolicyAllOf) SetSlot8linkSpeed(v string)`

SetSlot8linkSpeed sets Slot8linkSpeed field to given value.

### HasSlot8linkSpeed

`func (o *BiosPolicyAllOf) HasSlot8linkSpeed() bool`

HasSlot8linkSpeed returns a boolean if a field has been set.

### GetSlot8state

`func (o *BiosPolicyAllOf) GetSlot8state() string`

GetSlot8state returns the Slot8state field if non-nil, zero value otherwise.

### GetSlot8stateOk

`func (o *BiosPolicyAllOf) GetSlot8stateOk() (*string, bool)`

GetSlot8stateOk returns a tuple with the Slot8state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot8state

`func (o *BiosPolicyAllOf) SetSlot8state(v string)`

SetSlot8state sets Slot8state field to given value.

### HasSlot8state

`func (o *BiosPolicyAllOf) HasSlot8state() bool`

HasSlot8state returns a boolean if a field has been set.

### GetSlot9linkSpeed

`func (o *BiosPolicyAllOf) GetSlot9linkSpeed() string`

GetSlot9linkSpeed returns the Slot9linkSpeed field if non-nil, zero value otherwise.

### GetSlot9linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlot9linkSpeedOk() (*string, bool)`

GetSlot9linkSpeedOk returns a tuple with the Slot9linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot9linkSpeed

`func (o *BiosPolicyAllOf) SetSlot9linkSpeed(v string)`

SetSlot9linkSpeed sets Slot9linkSpeed field to given value.

### HasSlot9linkSpeed

`func (o *BiosPolicyAllOf) HasSlot9linkSpeed() bool`

HasSlot9linkSpeed returns a boolean if a field has been set.

### GetSlot9state

`func (o *BiosPolicyAllOf) GetSlot9state() string`

GetSlot9state returns the Slot9state field if non-nil, zero value otherwise.

### GetSlot9stateOk

`func (o *BiosPolicyAllOf) GetSlot9stateOk() (*string, bool)`

GetSlot9stateOk returns a tuple with the Slot9state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot9state

`func (o *BiosPolicyAllOf) SetSlot9state(v string)`

SetSlot9state sets Slot9state field to given value.

### HasSlot9state

`func (o *BiosPolicyAllOf) HasSlot9state() bool`

HasSlot9state returns a boolean if a field has been set.

### GetSlotFlomLinkSpeed

`func (o *BiosPolicyAllOf) GetSlotFlomLinkSpeed() string`

GetSlotFlomLinkSpeed returns the SlotFlomLinkSpeed field if non-nil, zero value otherwise.

### GetSlotFlomLinkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFlomLinkSpeedOk() (*string, bool)`

GetSlotFlomLinkSpeedOk returns a tuple with the SlotFlomLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFlomLinkSpeed

`func (o *BiosPolicyAllOf) SetSlotFlomLinkSpeed(v string)`

SetSlotFlomLinkSpeed sets SlotFlomLinkSpeed field to given value.

### HasSlotFlomLinkSpeed

`func (o *BiosPolicyAllOf) HasSlotFlomLinkSpeed() bool`

HasSlotFlomLinkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme10linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme10linkSpeed() string`

GetSlotFrontNvme10linkSpeed returns the SlotFrontNvme10linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme10linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme10linkSpeedOk() (*string, bool)`

GetSlotFrontNvme10linkSpeedOk returns a tuple with the SlotFrontNvme10linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme10linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme10linkSpeed(v string)`

SetSlotFrontNvme10linkSpeed sets SlotFrontNvme10linkSpeed field to given value.

### HasSlotFrontNvme10linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme10linkSpeed() bool`

HasSlotFrontNvme10linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme10optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme10optionRom() string`

GetSlotFrontNvme10optionRom returns the SlotFrontNvme10optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme10optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme10optionRomOk() (*string, bool)`

GetSlotFrontNvme10optionRomOk returns a tuple with the SlotFrontNvme10optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme10optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme10optionRom(v string)`

SetSlotFrontNvme10optionRom sets SlotFrontNvme10optionRom field to given value.

### HasSlotFrontNvme10optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme10optionRom() bool`

HasSlotFrontNvme10optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme11linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme11linkSpeed() string`

GetSlotFrontNvme11linkSpeed returns the SlotFrontNvme11linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme11linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme11linkSpeedOk() (*string, bool)`

GetSlotFrontNvme11linkSpeedOk returns a tuple with the SlotFrontNvme11linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme11linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme11linkSpeed(v string)`

SetSlotFrontNvme11linkSpeed sets SlotFrontNvme11linkSpeed field to given value.

### HasSlotFrontNvme11linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme11linkSpeed() bool`

HasSlotFrontNvme11linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme11optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme11optionRom() string`

GetSlotFrontNvme11optionRom returns the SlotFrontNvme11optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme11optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme11optionRomOk() (*string, bool)`

GetSlotFrontNvme11optionRomOk returns a tuple with the SlotFrontNvme11optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme11optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme11optionRom(v string)`

SetSlotFrontNvme11optionRom sets SlotFrontNvme11optionRom field to given value.

### HasSlotFrontNvme11optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme11optionRom() bool`

HasSlotFrontNvme11optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme12linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme12linkSpeed() string`

GetSlotFrontNvme12linkSpeed returns the SlotFrontNvme12linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme12linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme12linkSpeedOk() (*string, bool)`

GetSlotFrontNvme12linkSpeedOk returns a tuple with the SlotFrontNvme12linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme12linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme12linkSpeed(v string)`

SetSlotFrontNvme12linkSpeed sets SlotFrontNvme12linkSpeed field to given value.

### HasSlotFrontNvme12linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme12linkSpeed() bool`

HasSlotFrontNvme12linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme12optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme12optionRom() string`

GetSlotFrontNvme12optionRom returns the SlotFrontNvme12optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme12optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme12optionRomOk() (*string, bool)`

GetSlotFrontNvme12optionRomOk returns a tuple with the SlotFrontNvme12optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme12optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme12optionRom(v string)`

SetSlotFrontNvme12optionRom sets SlotFrontNvme12optionRom field to given value.

### HasSlotFrontNvme12optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme12optionRom() bool`

HasSlotFrontNvme12optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme13optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme13optionRom() string`

GetSlotFrontNvme13optionRom returns the SlotFrontNvme13optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme13optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme13optionRomOk() (*string, bool)`

GetSlotFrontNvme13optionRomOk returns a tuple with the SlotFrontNvme13optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme13optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme13optionRom(v string)`

SetSlotFrontNvme13optionRom sets SlotFrontNvme13optionRom field to given value.

### HasSlotFrontNvme13optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme13optionRom() bool`

HasSlotFrontNvme13optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme14optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme14optionRom() string`

GetSlotFrontNvme14optionRom returns the SlotFrontNvme14optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme14optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme14optionRomOk() (*string, bool)`

GetSlotFrontNvme14optionRomOk returns a tuple with the SlotFrontNvme14optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme14optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme14optionRom(v string)`

SetSlotFrontNvme14optionRom sets SlotFrontNvme14optionRom field to given value.

### HasSlotFrontNvme14optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme14optionRom() bool`

HasSlotFrontNvme14optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme15optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme15optionRom() string`

GetSlotFrontNvme15optionRom returns the SlotFrontNvme15optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme15optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme15optionRomOk() (*string, bool)`

GetSlotFrontNvme15optionRomOk returns a tuple with the SlotFrontNvme15optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme15optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme15optionRom(v string)`

SetSlotFrontNvme15optionRom sets SlotFrontNvme15optionRom field to given value.

### HasSlotFrontNvme15optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme15optionRom() bool`

HasSlotFrontNvme15optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme16optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme16optionRom() string`

GetSlotFrontNvme16optionRom returns the SlotFrontNvme16optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme16optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme16optionRomOk() (*string, bool)`

GetSlotFrontNvme16optionRomOk returns a tuple with the SlotFrontNvme16optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme16optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme16optionRom(v string)`

SetSlotFrontNvme16optionRom sets SlotFrontNvme16optionRom field to given value.

### HasSlotFrontNvme16optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme16optionRom() bool`

HasSlotFrontNvme16optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme17optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme17optionRom() string`

GetSlotFrontNvme17optionRom returns the SlotFrontNvme17optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme17optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme17optionRomOk() (*string, bool)`

GetSlotFrontNvme17optionRomOk returns a tuple with the SlotFrontNvme17optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme17optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme17optionRom(v string)`

SetSlotFrontNvme17optionRom sets SlotFrontNvme17optionRom field to given value.

### HasSlotFrontNvme17optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme17optionRom() bool`

HasSlotFrontNvme17optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme18optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme18optionRom() string`

GetSlotFrontNvme18optionRom returns the SlotFrontNvme18optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme18optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme18optionRomOk() (*string, bool)`

GetSlotFrontNvme18optionRomOk returns a tuple with the SlotFrontNvme18optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme18optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme18optionRom(v string)`

SetSlotFrontNvme18optionRom sets SlotFrontNvme18optionRom field to given value.

### HasSlotFrontNvme18optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme18optionRom() bool`

HasSlotFrontNvme18optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme19optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme19optionRom() string`

GetSlotFrontNvme19optionRom returns the SlotFrontNvme19optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme19optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme19optionRomOk() (*string, bool)`

GetSlotFrontNvme19optionRomOk returns a tuple with the SlotFrontNvme19optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme19optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme19optionRom(v string)`

SetSlotFrontNvme19optionRom sets SlotFrontNvme19optionRom field to given value.

### HasSlotFrontNvme19optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme19optionRom() bool`

HasSlotFrontNvme19optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme1linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme1linkSpeed() string`

GetSlotFrontNvme1linkSpeed returns the SlotFrontNvme1linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme1linkSpeedOk() (*string, bool)`

GetSlotFrontNvme1linkSpeedOk returns a tuple with the SlotFrontNvme1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme1linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme1linkSpeed(v string)`

SetSlotFrontNvme1linkSpeed sets SlotFrontNvme1linkSpeed field to given value.

### HasSlotFrontNvme1linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme1linkSpeed() bool`

HasSlotFrontNvme1linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme1optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme1optionRom() string`

GetSlotFrontNvme1optionRom returns the SlotFrontNvme1optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme1optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme1optionRomOk() (*string, bool)`

GetSlotFrontNvme1optionRomOk returns a tuple with the SlotFrontNvme1optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme1optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme1optionRom(v string)`

SetSlotFrontNvme1optionRom sets SlotFrontNvme1optionRom field to given value.

### HasSlotFrontNvme1optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme1optionRom() bool`

HasSlotFrontNvme1optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme20optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme20optionRom() string`

GetSlotFrontNvme20optionRom returns the SlotFrontNvme20optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme20optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme20optionRomOk() (*string, bool)`

GetSlotFrontNvme20optionRomOk returns a tuple with the SlotFrontNvme20optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme20optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme20optionRom(v string)`

SetSlotFrontNvme20optionRom sets SlotFrontNvme20optionRom field to given value.

### HasSlotFrontNvme20optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme20optionRom() bool`

HasSlotFrontNvme20optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme21optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme21optionRom() string`

GetSlotFrontNvme21optionRom returns the SlotFrontNvme21optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme21optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme21optionRomOk() (*string, bool)`

GetSlotFrontNvme21optionRomOk returns a tuple with the SlotFrontNvme21optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme21optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme21optionRom(v string)`

SetSlotFrontNvme21optionRom sets SlotFrontNvme21optionRom field to given value.

### HasSlotFrontNvme21optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme21optionRom() bool`

HasSlotFrontNvme21optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme22optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme22optionRom() string`

GetSlotFrontNvme22optionRom returns the SlotFrontNvme22optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme22optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme22optionRomOk() (*string, bool)`

GetSlotFrontNvme22optionRomOk returns a tuple with the SlotFrontNvme22optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme22optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme22optionRom(v string)`

SetSlotFrontNvme22optionRom sets SlotFrontNvme22optionRom field to given value.

### HasSlotFrontNvme22optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme22optionRom() bool`

HasSlotFrontNvme22optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme23optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme23optionRom() string`

GetSlotFrontNvme23optionRom returns the SlotFrontNvme23optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme23optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme23optionRomOk() (*string, bool)`

GetSlotFrontNvme23optionRomOk returns a tuple with the SlotFrontNvme23optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme23optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme23optionRom(v string)`

SetSlotFrontNvme23optionRom sets SlotFrontNvme23optionRom field to given value.

### HasSlotFrontNvme23optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme23optionRom() bool`

HasSlotFrontNvme23optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme24optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme24optionRom() string`

GetSlotFrontNvme24optionRom returns the SlotFrontNvme24optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme24optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme24optionRomOk() (*string, bool)`

GetSlotFrontNvme24optionRomOk returns a tuple with the SlotFrontNvme24optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme24optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme24optionRom(v string)`

SetSlotFrontNvme24optionRom sets SlotFrontNvme24optionRom field to given value.

### HasSlotFrontNvme24optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme24optionRom() bool`

HasSlotFrontNvme24optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme2linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme2linkSpeed() string`

GetSlotFrontNvme2linkSpeed returns the SlotFrontNvme2linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme2linkSpeedOk() (*string, bool)`

GetSlotFrontNvme2linkSpeedOk returns a tuple with the SlotFrontNvme2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme2linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme2linkSpeed(v string)`

SetSlotFrontNvme2linkSpeed sets SlotFrontNvme2linkSpeed field to given value.

### HasSlotFrontNvme2linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme2linkSpeed() bool`

HasSlotFrontNvme2linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme2optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme2optionRom() string`

GetSlotFrontNvme2optionRom returns the SlotFrontNvme2optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme2optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme2optionRomOk() (*string, bool)`

GetSlotFrontNvme2optionRomOk returns a tuple with the SlotFrontNvme2optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme2optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme2optionRom(v string)`

SetSlotFrontNvme2optionRom sets SlotFrontNvme2optionRom field to given value.

### HasSlotFrontNvme2optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme2optionRom() bool`

HasSlotFrontNvme2optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme3linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme3linkSpeed() string`

GetSlotFrontNvme3linkSpeed returns the SlotFrontNvme3linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme3linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme3linkSpeedOk() (*string, bool)`

GetSlotFrontNvme3linkSpeedOk returns a tuple with the SlotFrontNvme3linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme3linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme3linkSpeed(v string)`

SetSlotFrontNvme3linkSpeed sets SlotFrontNvme3linkSpeed field to given value.

### HasSlotFrontNvme3linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme3linkSpeed() bool`

HasSlotFrontNvme3linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme3optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme3optionRom() string`

GetSlotFrontNvme3optionRom returns the SlotFrontNvme3optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme3optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme3optionRomOk() (*string, bool)`

GetSlotFrontNvme3optionRomOk returns a tuple with the SlotFrontNvme3optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme3optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme3optionRom(v string)`

SetSlotFrontNvme3optionRom sets SlotFrontNvme3optionRom field to given value.

### HasSlotFrontNvme3optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme3optionRom() bool`

HasSlotFrontNvme3optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme4linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme4linkSpeed() string`

GetSlotFrontNvme4linkSpeed returns the SlotFrontNvme4linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme4linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme4linkSpeedOk() (*string, bool)`

GetSlotFrontNvme4linkSpeedOk returns a tuple with the SlotFrontNvme4linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme4linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme4linkSpeed(v string)`

SetSlotFrontNvme4linkSpeed sets SlotFrontNvme4linkSpeed field to given value.

### HasSlotFrontNvme4linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme4linkSpeed() bool`

HasSlotFrontNvme4linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme4optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme4optionRom() string`

GetSlotFrontNvme4optionRom returns the SlotFrontNvme4optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme4optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme4optionRomOk() (*string, bool)`

GetSlotFrontNvme4optionRomOk returns a tuple with the SlotFrontNvme4optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme4optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme4optionRom(v string)`

SetSlotFrontNvme4optionRom sets SlotFrontNvme4optionRom field to given value.

### HasSlotFrontNvme4optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme4optionRom() bool`

HasSlotFrontNvme4optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme5linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme5linkSpeed() string`

GetSlotFrontNvme5linkSpeed returns the SlotFrontNvme5linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme5linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme5linkSpeedOk() (*string, bool)`

GetSlotFrontNvme5linkSpeedOk returns a tuple with the SlotFrontNvme5linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme5linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme5linkSpeed(v string)`

SetSlotFrontNvme5linkSpeed sets SlotFrontNvme5linkSpeed field to given value.

### HasSlotFrontNvme5linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme5linkSpeed() bool`

HasSlotFrontNvme5linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme5optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme5optionRom() string`

GetSlotFrontNvme5optionRom returns the SlotFrontNvme5optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme5optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme5optionRomOk() (*string, bool)`

GetSlotFrontNvme5optionRomOk returns a tuple with the SlotFrontNvme5optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme5optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme5optionRom(v string)`

SetSlotFrontNvme5optionRom sets SlotFrontNvme5optionRom field to given value.

### HasSlotFrontNvme5optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme5optionRom() bool`

HasSlotFrontNvme5optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme6linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme6linkSpeed() string`

GetSlotFrontNvme6linkSpeed returns the SlotFrontNvme6linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme6linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme6linkSpeedOk() (*string, bool)`

GetSlotFrontNvme6linkSpeedOk returns a tuple with the SlotFrontNvme6linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme6linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme6linkSpeed(v string)`

SetSlotFrontNvme6linkSpeed sets SlotFrontNvme6linkSpeed field to given value.

### HasSlotFrontNvme6linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme6linkSpeed() bool`

HasSlotFrontNvme6linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme6optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme6optionRom() string`

GetSlotFrontNvme6optionRom returns the SlotFrontNvme6optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme6optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme6optionRomOk() (*string, bool)`

GetSlotFrontNvme6optionRomOk returns a tuple with the SlotFrontNvme6optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme6optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme6optionRom(v string)`

SetSlotFrontNvme6optionRom sets SlotFrontNvme6optionRom field to given value.

### HasSlotFrontNvme6optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme6optionRom() bool`

HasSlotFrontNvme6optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme7linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme7linkSpeed() string`

GetSlotFrontNvme7linkSpeed returns the SlotFrontNvme7linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme7linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme7linkSpeedOk() (*string, bool)`

GetSlotFrontNvme7linkSpeedOk returns a tuple with the SlotFrontNvme7linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme7linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme7linkSpeed(v string)`

SetSlotFrontNvme7linkSpeed sets SlotFrontNvme7linkSpeed field to given value.

### HasSlotFrontNvme7linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme7linkSpeed() bool`

HasSlotFrontNvme7linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme7optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme7optionRom() string`

GetSlotFrontNvme7optionRom returns the SlotFrontNvme7optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme7optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme7optionRomOk() (*string, bool)`

GetSlotFrontNvme7optionRomOk returns a tuple with the SlotFrontNvme7optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme7optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme7optionRom(v string)`

SetSlotFrontNvme7optionRom sets SlotFrontNvme7optionRom field to given value.

### HasSlotFrontNvme7optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme7optionRom() bool`

HasSlotFrontNvme7optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme8linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme8linkSpeed() string`

GetSlotFrontNvme8linkSpeed returns the SlotFrontNvme8linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme8linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme8linkSpeedOk() (*string, bool)`

GetSlotFrontNvme8linkSpeedOk returns a tuple with the SlotFrontNvme8linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme8linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme8linkSpeed(v string)`

SetSlotFrontNvme8linkSpeed sets SlotFrontNvme8linkSpeed field to given value.

### HasSlotFrontNvme8linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme8linkSpeed() bool`

HasSlotFrontNvme8linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme8optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme8optionRom() string`

GetSlotFrontNvme8optionRom returns the SlotFrontNvme8optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme8optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme8optionRomOk() (*string, bool)`

GetSlotFrontNvme8optionRomOk returns a tuple with the SlotFrontNvme8optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme8optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme8optionRom(v string)`

SetSlotFrontNvme8optionRom sets SlotFrontNvme8optionRom field to given value.

### HasSlotFrontNvme8optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme8optionRom() bool`

HasSlotFrontNvme8optionRom returns a boolean if a field has been set.

### GetSlotFrontNvme9linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontNvme9linkSpeed() string`

GetSlotFrontNvme9linkSpeed returns the SlotFrontNvme9linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontNvme9linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme9linkSpeedOk() (*string, bool)`

GetSlotFrontNvme9linkSpeedOk returns a tuple with the SlotFrontNvme9linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme9linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontNvme9linkSpeed(v string)`

SetSlotFrontNvme9linkSpeed sets SlotFrontNvme9linkSpeed field to given value.

### HasSlotFrontNvme9linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontNvme9linkSpeed() bool`

HasSlotFrontNvme9linkSpeed returns a boolean if a field has been set.

### GetSlotFrontNvme9optionRom

`func (o *BiosPolicyAllOf) GetSlotFrontNvme9optionRom() string`

GetSlotFrontNvme9optionRom returns the SlotFrontNvme9optionRom field if non-nil, zero value otherwise.

### GetSlotFrontNvme9optionRomOk

`func (o *BiosPolicyAllOf) GetSlotFrontNvme9optionRomOk() (*string, bool)`

GetSlotFrontNvme9optionRomOk returns a tuple with the SlotFrontNvme9optionRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontNvme9optionRom

`func (o *BiosPolicyAllOf) SetSlotFrontNvme9optionRom(v string)`

SetSlotFrontNvme9optionRom sets SlotFrontNvme9optionRom field to given value.

### HasSlotFrontNvme9optionRom

`func (o *BiosPolicyAllOf) HasSlotFrontNvme9optionRom() bool`

HasSlotFrontNvme9optionRom returns a boolean if a field has been set.

### GetSlotFrontSlot5linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontSlot5linkSpeed() string`

GetSlotFrontSlot5linkSpeed returns the SlotFrontSlot5linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontSlot5linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontSlot5linkSpeedOk() (*string, bool)`

GetSlotFrontSlot5linkSpeedOk returns a tuple with the SlotFrontSlot5linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontSlot5linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontSlot5linkSpeed(v string)`

SetSlotFrontSlot5linkSpeed sets SlotFrontSlot5linkSpeed field to given value.

### HasSlotFrontSlot5linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontSlot5linkSpeed() bool`

HasSlotFrontSlot5linkSpeed returns a boolean if a field has been set.

### GetSlotFrontSlot6linkSpeed

`func (o *BiosPolicyAllOf) GetSlotFrontSlot6linkSpeed() string`

GetSlotFrontSlot6linkSpeed returns the SlotFrontSlot6linkSpeed field if non-nil, zero value otherwise.

### GetSlotFrontSlot6linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotFrontSlot6linkSpeedOk() (*string, bool)`

GetSlotFrontSlot6linkSpeedOk returns a tuple with the SlotFrontSlot6linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotFrontSlot6linkSpeed

`func (o *BiosPolicyAllOf) SetSlotFrontSlot6linkSpeed(v string)`

SetSlotFrontSlot6linkSpeed sets SlotFrontSlot6linkSpeed field to given value.

### HasSlotFrontSlot6linkSpeed

`func (o *BiosPolicyAllOf) HasSlotFrontSlot6linkSpeed() bool`

HasSlotFrontSlot6linkSpeed returns a boolean if a field has been set.

### GetSlotGpu1state

`func (o *BiosPolicyAllOf) GetSlotGpu1state() string`

GetSlotGpu1state returns the SlotGpu1state field if non-nil, zero value otherwise.

### GetSlotGpu1stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu1stateOk() (*string, bool)`

GetSlotGpu1stateOk returns a tuple with the SlotGpu1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu1state

`func (o *BiosPolicyAllOf) SetSlotGpu1state(v string)`

SetSlotGpu1state sets SlotGpu1state field to given value.

### HasSlotGpu1state

`func (o *BiosPolicyAllOf) HasSlotGpu1state() bool`

HasSlotGpu1state returns a boolean if a field has been set.

### GetSlotGpu2state

`func (o *BiosPolicyAllOf) GetSlotGpu2state() string`

GetSlotGpu2state returns the SlotGpu2state field if non-nil, zero value otherwise.

### GetSlotGpu2stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu2stateOk() (*string, bool)`

GetSlotGpu2stateOk returns a tuple with the SlotGpu2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu2state

`func (o *BiosPolicyAllOf) SetSlotGpu2state(v string)`

SetSlotGpu2state sets SlotGpu2state field to given value.

### HasSlotGpu2state

`func (o *BiosPolicyAllOf) HasSlotGpu2state() bool`

HasSlotGpu2state returns a boolean if a field has been set.

### GetSlotGpu3state

`func (o *BiosPolicyAllOf) GetSlotGpu3state() string`

GetSlotGpu3state returns the SlotGpu3state field if non-nil, zero value otherwise.

### GetSlotGpu3stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu3stateOk() (*string, bool)`

GetSlotGpu3stateOk returns a tuple with the SlotGpu3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu3state

`func (o *BiosPolicyAllOf) SetSlotGpu3state(v string)`

SetSlotGpu3state sets SlotGpu3state field to given value.

### HasSlotGpu3state

`func (o *BiosPolicyAllOf) HasSlotGpu3state() bool`

HasSlotGpu3state returns a boolean if a field has been set.

### GetSlotGpu4state

`func (o *BiosPolicyAllOf) GetSlotGpu4state() string`

GetSlotGpu4state returns the SlotGpu4state field if non-nil, zero value otherwise.

### GetSlotGpu4stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu4stateOk() (*string, bool)`

GetSlotGpu4stateOk returns a tuple with the SlotGpu4state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu4state

`func (o *BiosPolicyAllOf) SetSlotGpu4state(v string)`

SetSlotGpu4state sets SlotGpu4state field to given value.

### HasSlotGpu4state

`func (o *BiosPolicyAllOf) HasSlotGpu4state() bool`

HasSlotGpu4state returns a boolean if a field has been set.

### GetSlotGpu5state

`func (o *BiosPolicyAllOf) GetSlotGpu5state() string`

GetSlotGpu5state returns the SlotGpu5state field if non-nil, zero value otherwise.

### GetSlotGpu5stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu5stateOk() (*string, bool)`

GetSlotGpu5stateOk returns a tuple with the SlotGpu5state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu5state

`func (o *BiosPolicyAllOf) SetSlotGpu5state(v string)`

SetSlotGpu5state sets SlotGpu5state field to given value.

### HasSlotGpu5state

`func (o *BiosPolicyAllOf) HasSlotGpu5state() bool`

HasSlotGpu5state returns a boolean if a field has been set.

### GetSlotGpu6state

`func (o *BiosPolicyAllOf) GetSlotGpu6state() string`

GetSlotGpu6state returns the SlotGpu6state field if non-nil, zero value otherwise.

### GetSlotGpu6stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu6stateOk() (*string, bool)`

GetSlotGpu6stateOk returns a tuple with the SlotGpu6state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu6state

`func (o *BiosPolicyAllOf) SetSlotGpu6state(v string)`

SetSlotGpu6state sets SlotGpu6state field to given value.

### HasSlotGpu6state

`func (o *BiosPolicyAllOf) HasSlotGpu6state() bool`

HasSlotGpu6state returns a boolean if a field has been set.

### GetSlotGpu7state

`func (o *BiosPolicyAllOf) GetSlotGpu7state() string`

GetSlotGpu7state returns the SlotGpu7state field if non-nil, zero value otherwise.

### GetSlotGpu7stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu7stateOk() (*string, bool)`

GetSlotGpu7stateOk returns a tuple with the SlotGpu7state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu7state

`func (o *BiosPolicyAllOf) SetSlotGpu7state(v string)`

SetSlotGpu7state sets SlotGpu7state field to given value.

### HasSlotGpu7state

`func (o *BiosPolicyAllOf) HasSlotGpu7state() bool`

HasSlotGpu7state returns a boolean if a field has been set.

### GetSlotGpu8state

`func (o *BiosPolicyAllOf) GetSlotGpu8state() string`

GetSlotGpu8state returns the SlotGpu8state field if non-nil, zero value otherwise.

### GetSlotGpu8stateOk

`func (o *BiosPolicyAllOf) GetSlotGpu8stateOk() (*string, bool)`

GetSlotGpu8stateOk returns a tuple with the SlotGpu8state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotGpu8state

`func (o *BiosPolicyAllOf) SetSlotGpu8state(v string)`

SetSlotGpu8state sets SlotGpu8state field to given value.

### HasSlotGpu8state

`func (o *BiosPolicyAllOf) HasSlotGpu8state() bool`

HasSlotGpu8state returns a boolean if a field has been set.

### GetSlotHbaLinkSpeed

`func (o *BiosPolicyAllOf) GetSlotHbaLinkSpeed() string`

GetSlotHbaLinkSpeed returns the SlotHbaLinkSpeed field if non-nil, zero value otherwise.

### GetSlotHbaLinkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotHbaLinkSpeedOk() (*string, bool)`

GetSlotHbaLinkSpeedOk returns a tuple with the SlotHbaLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotHbaLinkSpeed

`func (o *BiosPolicyAllOf) SetSlotHbaLinkSpeed(v string)`

SetSlotHbaLinkSpeed sets SlotHbaLinkSpeed field to given value.

### HasSlotHbaLinkSpeed

`func (o *BiosPolicyAllOf) HasSlotHbaLinkSpeed() bool`

HasSlotHbaLinkSpeed returns a boolean if a field has been set.

### GetSlotHbaState

`func (o *BiosPolicyAllOf) GetSlotHbaState() string`

GetSlotHbaState returns the SlotHbaState field if non-nil, zero value otherwise.

### GetSlotHbaStateOk

`func (o *BiosPolicyAllOf) GetSlotHbaStateOk() (*string, bool)`

GetSlotHbaStateOk returns a tuple with the SlotHbaState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotHbaState

`func (o *BiosPolicyAllOf) SetSlotHbaState(v string)`

SetSlotHbaState sets SlotHbaState field to given value.

### HasSlotHbaState

`func (o *BiosPolicyAllOf) HasSlotHbaState() bool`

HasSlotHbaState returns a boolean if a field has been set.

### GetSlotLom1link

`func (o *BiosPolicyAllOf) GetSlotLom1link() string`

GetSlotLom1link returns the SlotLom1link field if non-nil, zero value otherwise.

### GetSlotLom1linkOk

`func (o *BiosPolicyAllOf) GetSlotLom1linkOk() (*string, bool)`

GetSlotLom1linkOk returns a tuple with the SlotLom1link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotLom1link

`func (o *BiosPolicyAllOf) SetSlotLom1link(v string)`

SetSlotLom1link sets SlotLom1link field to given value.

### HasSlotLom1link

`func (o *BiosPolicyAllOf) HasSlotLom1link() bool`

HasSlotLom1link returns a boolean if a field has been set.

### GetSlotLom2link

`func (o *BiosPolicyAllOf) GetSlotLom2link() string`

GetSlotLom2link returns the SlotLom2link field if non-nil, zero value otherwise.

### GetSlotLom2linkOk

`func (o *BiosPolicyAllOf) GetSlotLom2linkOk() (*string, bool)`

GetSlotLom2linkOk returns a tuple with the SlotLom2link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotLom2link

`func (o *BiosPolicyAllOf) SetSlotLom2link(v string)`

SetSlotLom2link sets SlotLom2link field to given value.

### HasSlotLom2link

`func (o *BiosPolicyAllOf) HasSlotLom2link() bool`

HasSlotLom2link returns a boolean if a field has been set.

### GetSlotMezzState

`func (o *BiosPolicyAllOf) GetSlotMezzState() string`

GetSlotMezzState returns the SlotMezzState field if non-nil, zero value otherwise.

### GetSlotMezzStateOk

`func (o *BiosPolicyAllOf) GetSlotMezzStateOk() (*string, bool)`

GetSlotMezzStateOk returns a tuple with the SlotMezzState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotMezzState

`func (o *BiosPolicyAllOf) SetSlotMezzState(v string)`

SetSlotMezzState sets SlotMezzState field to given value.

### HasSlotMezzState

`func (o *BiosPolicyAllOf) HasSlotMezzState() bool`

HasSlotMezzState returns a boolean if a field has been set.

### GetSlotMlomLinkSpeed

`func (o *BiosPolicyAllOf) GetSlotMlomLinkSpeed() string`

GetSlotMlomLinkSpeed returns the SlotMlomLinkSpeed field if non-nil, zero value otherwise.

### GetSlotMlomLinkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotMlomLinkSpeedOk() (*string, bool)`

GetSlotMlomLinkSpeedOk returns a tuple with the SlotMlomLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotMlomLinkSpeed

`func (o *BiosPolicyAllOf) SetSlotMlomLinkSpeed(v string)`

SetSlotMlomLinkSpeed sets SlotMlomLinkSpeed field to given value.

### HasSlotMlomLinkSpeed

`func (o *BiosPolicyAllOf) HasSlotMlomLinkSpeed() bool`

HasSlotMlomLinkSpeed returns a boolean if a field has been set.

### GetSlotMlomState

`func (o *BiosPolicyAllOf) GetSlotMlomState() string`

GetSlotMlomState returns the SlotMlomState field if non-nil, zero value otherwise.

### GetSlotMlomStateOk

`func (o *BiosPolicyAllOf) GetSlotMlomStateOk() (*string, bool)`

GetSlotMlomStateOk returns a tuple with the SlotMlomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotMlomState

`func (o *BiosPolicyAllOf) SetSlotMlomState(v string)`

SetSlotMlomState sets SlotMlomState field to given value.

### HasSlotMlomState

`func (o *BiosPolicyAllOf) HasSlotMlomState() bool`

HasSlotMlomState returns a boolean if a field has been set.

### GetSlotMraidLinkSpeed

`func (o *BiosPolicyAllOf) GetSlotMraidLinkSpeed() string`

GetSlotMraidLinkSpeed returns the SlotMraidLinkSpeed field if non-nil, zero value otherwise.

### GetSlotMraidLinkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotMraidLinkSpeedOk() (*string, bool)`

GetSlotMraidLinkSpeedOk returns a tuple with the SlotMraidLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotMraidLinkSpeed

`func (o *BiosPolicyAllOf) SetSlotMraidLinkSpeed(v string)`

SetSlotMraidLinkSpeed sets SlotMraidLinkSpeed field to given value.

### HasSlotMraidLinkSpeed

`func (o *BiosPolicyAllOf) HasSlotMraidLinkSpeed() bool`

HasSlotMraidLinkSpeed returns a boolean if a field has been set.

### GetSlotMraidState

`func (o *BiosPolicyAllOf) GetSlotMraidState() string`

GetSlotMraidState returns the SlotMraidState field if non-nil, zero value otherwise.

### GetSlotMraidStateOk

`func (o *BiosPolicyAllOf) GetSlotMraidStateOk() (*string, bool)`

GetSlotMraidStateOk returns a tuple with the SlotMraidState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotMraidState

`func (o *BiosPolicyAllOf) SetSlotMraidState(v string)`

SetSlotMraidState sets SlotMraidState field to given value.

### HasSlotMraidState

`func (o *BiosPolicyAllOf) HasSlotMraidState() bool`

HasSlotMraidState returns a boolean if a field has been set.

### GetSlotN10state

`func (o *BiosPolicyAllOf) GetSlotN10state() string`

GetSlotN10state returns the SlotN10state field if non-nil, zero value otherwise.

### GetSlotN10stateOk

`func (o *BiosPolicyAllOf) GetSlotN10stateOk() (*string, bool)`

GetSlotN10stateOk returns a tuple with the SlotN10state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN10state

`func (o *BiosPolicyAllOf) SetSlotN10state(v string)`

SetSlotN10state sets SlotN10state field to given value.

### HasSlotN10state

`func (o *BiosPolicyAllOf) HasSlotN10state() bool`

HasSlotN10state returns a boolean if a field has been set.

### GetSlotN11state

`func (o *BiosPolicyAllOf) GetSlotN11state() string`

GetSlotN11state returns the SlotN11state field if non-nil, zero value otherwise.

### GetSlotN11stateOk

`func (o *BiosPolicyAllOf) GetSlotN11stateOk() (*string, bool)`

GetSlotN11stateOk returns a tuple with the SlotN11state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN11state

`func (o *BiosPolicyAllOf) SetSlotN11state(v string)`

SetSlotN11state sets SlotN11state field to given value.

### HasSlotN11state

`func (o *BiosPolicyAllOf) HasSlotN11state() bool`

HasSlotN11state returns a boolean if a field has been set.

### GetSlotN12state

`func (o *BiosPolicyAllOf) GetSlotN12state() string`

GetSlotN12state returns the SlotN12state field if non-nil, zero value otherwise.

### GetSlotN12stateOk

`func (o *BiosPolicyAllOf) GetSlotN12stateOk() (*string, bool)`

GetSlotN12stateOk returns a tuple with the SlotN12state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN12state

`func (o *BiosPolicyAllOf) SetSlotN12state(v string)`

SetSlotN12state sets SlotN12state field to given value.

### HasSlotN12state

`func (o *BiosPolicyAllOf) HasSlotN12state() bool`

HasSlotN12state returns a boolean if a field has been set.

### GetSlotN13state

`func (o *BiosPolicyAllOf) GetSlotN13state() string`

GetSlotN13state returns the SlotN13state field if non-nil, zero value otherwise.

### GetSlotN13stateOk

`func (o *BiosPolicyAllOf) GetSlotN13stateOk() (*string, bool)`

GetSlotN13stateOk returns a tuple with the SlotN13state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN13state

`func (o *BiosPolicyAllOf) SetSlotN13state(v string)`

SetSlotN13state sets SlotN13state field to given value.

### HasSlotN13state

`func (o *BiosPolicyAllOf) HasSlotN13state() bool`

HasSlotN13state returns a boolean if a field has been set.

### GetSlotN14state

`func (o *BiosPolicyAllOf) GetSlotN14state() string`

GetSlotN14state returns the SlotN14state field if non-nil, zero value otherwise.

### GetSlotN14stateOk

`func (o *BiosPolicyAllOf) GetSlotN14stateOk() (*string, bool)`

GetSlotN14stateOk returns a tuple with the SlotN14state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN14state

`func (o *BiosPolicyAllOf) SetSlotN14state(v string)`

SetSlotN14state sets SlotN14state field to given value.

### HasSlotN14state

`func (o *BiosPolicyAllOf) HasSlotN14state() bool`

HasSlotN14state returns a boolean if a field has been set.

### GetSlotN15state

`func (o *BiosPolicyAllOf) GetSlotN15state() string`

GetSlotN15state returns the SlotN15state field if non-nil, zero value otherwise.

### GetSlotN15stateOk

`func (o *BiosPolicyAllOf) GetSlotN15stateOk() (*string, bool)`

GetSlotN15stateOk returns a tuple with the SlotN15state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN15state

`func (o *BiosPolicyAllOf) SetSlotN15state(v string)`

SetSlotN15state sets SlotN15state field to given value.

### HasSlotN15state

`func (o *BiosPolicyAllOf) HasSlotN15state() bool`

HasSlotN15state returns a boolean if a field has been set.

### GetSlotN16state

`func (o *BiosPolicyAllOf) GetSlotN16state() string`

GetSlotN16state returns the SlotN16state field if non-nil, zero value otherwise.

### GetSlotN16stateOk

`func (o *BiosPolicyAllOf) GetSlotN16stateOk() (*string, bool)`

GetSlotN16stateOk returns a tuple with the SlotN16state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN16state

`func (o *BiosPolicyAllOf) SetSlotN16state(v string)`

SetSlotN16state sets SlotN16state field to given value.

### HasSlotN16state

`func (o *BiosPolicyAllOf) HasSlotN16state() bool`

HasSlotN16state returns a boolean if a field has been set.

### GetSlotN17state

`func (o *BiosPolicyAllOf) GetSlotN17state() string`

GetSlotN17state returns the SlotN17state field if non-nil, zero value otherwise.

### GetSlotN17stateOk

`func (o *BiosPolicyAllOf) GetSlotN17stateOk() (*string, bool)`

GetSlotN17stateOk returns a tuple with the SlotN17state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN17state

`func (o *BiosPolicyAllOf) SetSlotN17state(v string)`

SetSlotN17state sets SlotN17state field to given value.

### HasSlotN17state

`func (o *BiosPolicyAllOf) HasSlotN17state() bool`

HasSlotN17state returns a boolean if a field has been set.

### GetSlotN18state

`func (o *BiosPolicyAllOf) GetSlotN18state() string`

GetSlotN18state returns the SlotN18state field if non-nil, zero value otherwise.

### GetSlotN18stateOk

`func (o *BiosPolicyAllOf) GetSlotN18stateOk() (*string, bool)`

GetSlotN18stateOk returns a tuple with the SlotN18state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN18state

`func (o *BiosPolicyAllOf) SetSlotN18state(v string)`

SetSlotN18state sets SlotN18state field to given value.

### HasSlotN18state

`func (o *BiosPolicyAllOf) HasSlotN18state() bool`

HasSlotN18state returns a boolean if a field has been set.

### GetSlotN19state

`func (o *BiosPolicyAllOf) GetSlotN19state() string`

GetSlotN19state returns the SlotN19state field if non-nil, zero value otherwise.

### GetSlotN19stateOk

`func (o *BiosPolicyAllOf) GetSlotN19stateOk() (*string, bool)`

GetSlotN19stateOk returns a tuple with the SlotN19state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN19state

`func (o *BiosPolicyAllOf) SetSlotN19state(v string)`

SetSlotN19state sets SlotN19state field to given value.

### HasSlotN19state

`func (o *BiosPolicyAllOf) HasSlotN19state() bool`

HasSlotN19state returns a boolean if a field has been set.

### GetSlotN1state

`func (o *BiosPolicyAllOf) GetSlotN1state() string`

GetSlotN1state returns the SlotN1state field if non-nil, zero value otherwise.

### GetSlotN1stateOk

`func (o *BiosPolicyAllOf) GetSlotN1stateOk() (*string, bool)`

GetSlotN1stateOk returns a tuple with the SlotN1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN1state

`func (o *BiosPolicyAllOf) SetSlotN1state(v string)`

SetSlotN1state sets SlotN1state field to given value.

### HasSlotN1state

`func (o *BiosPolicyAllOf) HasSlotN1state() bool`

HasSlotN1state returns a boolean if a field has been set.

### GetSlotN20state

`func (o *BiosPolicyAllOf) GetSlotN20state() string`

GetSlotN20state returns the SlotN20state field if non-nil, zero value otherwise.

### GetSlotN20stateOk

`func (o *BiosPolicyAllOf) GetSlotN20stateOk() (*string, bool)`

GetSlotN20stateOk returns a tuple with the SlotN20state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN20state

`func (o *BiosPolicyAllOf) SetSlotN20state(v string)`

SetSlotN20state sets SlotN20state field to given value.

### HasSlotN20state

`func (o *BiosPolicyAllOf) HasSlotN20state() bool`

HasSlotN20state returns a boolean if a field has been set.

### GetSlotN21state

`func (o *BiosPolicyAllOf) GetSlotN21state() string`

GetSlotN21state returns the SlotN21state field if non-nil, zero value otherwise.

### GetSlotN21stateOk

`func (o *BiosPolicyAllOf) GetSlotN21stateOk() (*string, bool)`

GetSlotN21stateOk returns a tuple with the SlotN21state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN21state

`func (o *BiosPolicyAllOf) SetSlotN21state(v string)`

SetSlotN21state sets SlotN21state field to given value.

### HasSlotN21state

`func (o *BiosPolicyAllOf) HasSlotN21state() bool`

HasSlotN21state returns a boolean if a field has been set.

### GetSlotN22state

`func (o *BiosPolicyAllOf) GetSlotN22state() string`

GetSlotN22state returns the SlotN22state field if non-nil, zero value otherwise.

### GetSlotN22stateOk

`func (o *BiosPolicyAllOf) GetSlotN22stateOk() (*string, bool)`

GetSlotN22stateOk returns a tuple with the SlotN22state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN22state

`func (o *BiosPolicyAllOf) SetSlotN22state(v string)`

SetSlotN22state sets SlotN22state field to given value.

### HasSlotN22state

`func (o *BiosPolicyAllOf) HasSlotN22state() bool`

HasSlotN22state returns a boolean if a field has been set.

### GetSlotN23state

`func (o *BiosPolicyAllOf) GetSlotN23state() string`

GetSlotN23state returns the SlotN23state field if non-nil, zero value otherwise.

### GetSlotN23stateOk

`func (o *BiosPolicyAllOf) GetSlotN23stateOk() (*string, bool)`

GetSlotN23stateOk returns a tuple with the SlotN23state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN23state

`func (o *BiosPolicyAllOf) SetSlotN23state(v string)`

SetSlotN23state sets SlotN23state field to given value.

### HasSlotN23state

`func (o *BiosPolicyAllOf) HasSlotN23state() bool`

HasSlotN23state returns a boolean if a field has been set.

### GetSlotN24state

`func (o *BiosPolicyAllOf) GetSlotN24state() string`

GetSlotN24state returns the SlotN24state field if non-nil, zero value otherwise.

### GetSlotN24stateOk

`func (o *BiosPolicyAllOf) GetSlotN24stateOk() (*string, bool)`

GetSlotN24stateOk returns a tuple with the SlotN24state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN24state

`func (o *BiosPolicyAllOf) SetSlotN24state(v string)`

SetSlotN24state sets SlotN24state field to given value.

### HasSlotN24state

`func (o *BiosPolicyAllOf) HasSlotN24state() bool`

HasSlotN24state returns a boolean if a field has been set.

### GetSlotN2state

`func (o *BiosPolicyAllOf) GetSlotN2state() string`

GetSlotN2state returns the SlotN2state field if non-nil, zero value otherwise.

### GetSlotN2stateOk

`func (o *BiosPolicyAllOf) GetSlotN2stateOk() (*string, bool)`

GetSlotN2stateOk returns a tuple with the SlotN2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN2state

`func (o *BiosPolicyAllOf) SetSlotN2state(v string)`

SetSlotN2state sets SlotN2state field to given value.

### HasSlotN2state

`func (o *BiosPolicyAllOf) HasSlotN2state() bool`

HasSlotN2state returns a boolean if a field has been set.

### GetSlotN3state

`func (o *BiosPolicyAllOf) GetSlotN3state() string`

GetSlotN3state returns the SlotN3state field if non-nil, zero value otherwise.

### GetSlotN3stateOk

`func (o *BiosPolicyAllOf) GetSlotN3stateOk() (*string, bool)`

GetSlotN3stateOk returns a tuple with the SlotN3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN3state

`func (o *BiosPolicyAllOf) SetSlotN3state(v string)`

SetSlotN3state sets SlotN3state field to given value.

### HasSlotN3state

`func (o *BiosPolicyAllOf) HasSlotN3state() bool`

HasSlotN3state returns a boolean if a field has been set.

### GetSlotN4state

`func (o *BiosPolicyAllOf) GetSlotN4state() string`

GetSlotN4state returns the SlotN4state field if non-nil, zero value otherwise.

### GetSlotN4stateOk

`func (o *BiosPolicyAllOf) GetSlotN4stateOk() (*string, bool)`

GetSlotN4stateOk returns a tuple with the SlotN4state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN4state

`func (o *BiosPolicyAllOf) SetSlotN4state(v string)`

SetSlotN4state sets SlotN4state field to given value.

### HasSlotN4state

`func (o *BiosPolicyAllOf) HasSlotN4state() bool`

HasSlotN4state returns a boolean if a field has been set.

### GetSlotN5state

`func (o *BiosPolicyAllOf) GetSlotN5state() string`

GetSlotN5state returns the SlotN5state field if non-nil, zero value otherwise.

### GetSlotN5stateOk

`func (o *BiosPolicyAllOf) GetSlotN5stateOk() (*string, bool)`

GetSlotN5stateOk returns a tuple with the SlotN5state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN5state

`func (o *BiosPolicyAllOf) SetSlotN5state(v string)`

SetSlotN5state sets SlotN5state field to given value.

### HasSlotN5state

`func (o *BiosPolicyAllOf) HasSlotN5state() bool`

HasSlotN5state returns a boolean if a field has been set.

### GetSlotN6state

`func (o *BiosPolicyAllOf) GetSlotN6state() string`

GetSlotN6state returns the SlotN6state field if non-nil, zero value otherwise.

### GetSlotN6stateOk

`func (o *BiosPolicyAllOf) GetSlotN6stateOk() (*string, bool)`

GetSlotN6stateOk returns a tuple with the SlotN6state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN6state

`func (o *BiosPolicyAllOf) SetSlotN6state(v string)`

SetSlotN6state sets SlotN6state field to given value.

### HasSlotN6state

`func (o *BiosPolicyAllOf) HasSlotN6state() bool`

HasSlotN6state returns a boolean if a field has been set.

### GetSlotN7state

`func (o *BiosPolicyAllOf) GetSlotN7state() string`

GetSlotN7state returns the SlotN7state field if non-nil, zero value otherwise.

### GetSlotN7stateOk

`func (o *BiosPolicyAllOf) GetSlotN7stateOk() (*string, bool)`

GetSlotN7stateOk returns a tuple with the SlotN7state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN7state

`func (o *BiosPolicyAllOf) SetSlotN7state(v string)`

SetSlotN7state sets SlotN7state field to given value.

### HasSlotN7state

`func (o *BiosPolicyAllOf) HasSlotN7state() bool`

HasSlotN7state returns a boolean if a field has been set.

### GetSlotN8state

`func (o *BiosPolicyAllOf) GetSlotN8state() string`

GetSlotN8state returns the SlotN8state field if non-nil, zero value otherwise.

### GetSlotN8stateOk

`func (o *BiosPolicyAllOf) GetSlotN8stateOk() (*string, bool)`

GetSlotN8stateOk returns a tuple with the SlotN8state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN8state

`func (o *BiosPolicyAllOf) SetSlotN8state(v string)`

SetSlotN8state sets SlotN8state field to given value.

### HasSlotN8state

`func (o *BiosPolicyAllOf) HasSlotN8state() bool`

HasSlotN8state returns a boolean if a field has been set.

### GetSlotN9state

`func (o *BiosPolicyAllOf) GetSlotN9state() string`

GetSlotN9state returns the SlotN9state field if non-nil, zero value otherwise.

### GetSlotN9stateOk

`func (o *BiosPolicyAllOf) GetSlotN9stateOk() (*string, bool)`

GetSlotN9stateOk returns a tuple with the SlotN9state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotN9state

`func (o *BiosPolicyAllOf) SetSlotN9state(v string)`

SetSlotN9state sets SlotN9state field to given value.

### HasSlotN9state

`func (o *BiosPolicyAllOf) HasSlotN9state() bool`

HasSlotN9state returns a boolean if a field has been set.

### GetSlotRaidLinkSpeed

`func (o *BiosPolicyAllOf) GetSlotRaidLinkSpeed() string`

GetSlotRaidLinkSpeed returns the SlotRaidLinkSpeed field if non-nil, zero value otherwise.

### GetSlotRaidLinkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRaidLinkSpeedOk() (*string, bool)`

GetSlotRaidLinkSpeedOk returns a tuple with the SlotRaidLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRaidLinkSpeed

`func (o *BiosPolicyAllOf) SetSlotRaidLinkSpeed(v string)`

SetSlotRaidLinkSpeed sets SlotRaidLinkSpeed field to given value.

### HasSlotRaidLinkSpeed

`func (o *BiosPolicyAllOf) HasSlotRaidLinkSpeed() bool`

HasSlotRaidLinkSpeed returns a boolean if a field has been set.

### GetSlotRaidState

`func (o *BiosPolicyAllOf) GetSlotRaidState() string`

GetSlotRaidState returns the SlotRaidState field if non-nil, zero value otherwise.

### GetSlotRaidStateOk

`func (o *BiosPolicyAllOf) GetSlotRaidStateOk() (*string, bool)`

GetSlotRaidStateOk returns a tuple with the SlotRaidState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRaidState

`func (o *BiosPolicyAllOf) SetSlotRaidState(v string)`

SetSlotRaidState sets SlotRaidState field to given value.

### HasSlotRaidState

`func (o *BiosPolicyAllOf) HasSlotRaidState() bool`

HasSlotRaidState returns a boolean if a field has been set.

### GetSlotRearNvme1linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRearNvme1linkSpeed() string`

GetSlotRearNvme1linkSpeed returns the SlotRearNvme1linkSpeed field if non-nil, zero value otherwise.

### GetSlotRearNvme1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme1linkSpeedOk() (*string, bool)`

GetSlotRearNvme1linkSpeedOk returns a tuple with the SlotRearNvme1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme1linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRearNvme1linkSpeed(v string)`

SetSlotRearNvme1linkSpeed sets SlotRearNvme1linkSpeed field to given value.

### HasSlotRearNvme1linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRearNvme1linkSpeed() bool`

HasSlotRearNvme1linkSpeed returns a boolean if a field has been set.

### GetSlotRearNvme1state

`func (o *BiosPolicyAllOf) GetSlotRearNvme1state() string`

GetSlotRearNvme1state returns the SlotRearNvme1state field if non-nil, zero value otherwise.

### GetSlotRearNvme1stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme1stateOk() (*string, bool)`

GetSlotRearNvme1stateOk returns a tuple with the SlotRearNvme1state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme1state

`func (o *BiosPolicyAllOf) SetSlotRearNvme1state(v string)`

SetSlotRearNvme1state sets SlotRearNvme1state field to given value.

### HasSlotRearNvme1state

`func (o *BiosPolicyAllOf) HasSlotRearNvme1state() bool`

HasSlotRearNvme1state returns a boolean if a field has been set.

### GetSlotRearNvme2linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRearNvme2linkSpeed() string`

GetSlotRearNvme2linkSpeed returns the SlotRearNvme2linkSpeed field if non-nil, zero value otherwise.

### GetSlotRearNvme2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme2linkSpeedOk() (*string, bool)`

GetSlotRearNvme2linkSpeedOk returns a tuple with the SlotRearNvme2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme2linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRearNvme2linkSpeed(v string)`

SetSlotRearNvme2linkSpeed sets SlotRearNvme2linkSpeed field to given value.

### HasSlotRearNvme2linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRearNvme2linkSpeed() bool`

HasSlotRearNvme2linkSpeed returns a boolean if a field has been set.

### GetSlotRearNvme2state

`func (o *BiosPolicyAllOf) GetSlotRearNvme2state() string`

GetSlotRearNvme2state returns the SlotRearNvme2state field if non-nil, zero value otherwise.

### GetSlotRearNvme2stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme2stateOk() (*string, bool)`

GetSlotRearNvme2stateOk returns a tuple with the SlotRearNvme2state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme2state

`func (o *BiosPolicyAllOf) SetSlotRearNvme2state(v string)`

SetSlotRearNvme2state sets SlotRearNvme2state field to given value.

### HasSlotRearNvme2state

`func (o *BiosPolicyAllOf) HasSlotRearNvme2state() bool`

HasSlotRearNvme2state returns a boolean if a field has been set.

### GetSlotRearNvme3linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRearNvme3linkSpeed() string`

GetSlotRearNvme3linkSpeed returns the SlotRearNvme3linkSpeed field if non-nil, zero value otherwise.

### GetSlotRearNvme3linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme3linkSpeedOk() (*string, bool)`

GetSlotRearNvme3linkSpeedOk returns a tuple with the SlotRearNvme3linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme3linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRearNvme3linkSpeed(v string)`

SetSlotRearNvme3linkSpeed sets SlotRearNvme3linkSpeed field to given value.

### HasSlotRearNvme3linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRearNvme3linkSpeed() bool`

HasSlotRearNvme3linkSpeed returns a boolean if a field has been set.

### GetSlotRearNvme3state

`func (o *BiosPolicyAllOf) GetSlotRearNvme3state() string`

GetSlotRearNvme3state returns the SlotRearNvme3state field if non-nil, zero value otherwise.

### GetSlotRearNvme3stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme3stateOk() (*string, bool)`

GetSlotRearNvme3stateOk returns a tuple with the SlotRearNvme3state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme3state

`func (o *BiosPolicyAllOf) SetSlotRearNvme3state(v string)`

SetSlotRearNvme3state sets SlotRearNvme3state field to given value.

### HasSlotRearNvme3state

`func (o *BiosPolicyAllOf) HasSlotRearNvme3state() bool`

HasSlotRearNvme3state returns a boolean if a field has been set.

### GetSlotRearNvme4linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRearNvme4linkSpeed() string`

GetSlotRearNvme4linkSpeed returns the SlotRearNvme4linkSpeed field if non-nil, zero value otherwise.

### GetSlotRearNvme4linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme4linkSpeedOk() (*string, bool)`

GetSlotRearNvme4linkSpeedOk returns a tuple with the SlotRearNvme4linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme4linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRearNvme4linkSpeed(v string)`

SetSlotRearNvme4linkSpeed sets SlotRearNvme4linkSpeed field to given value.

### HasSlotRearNvme4linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRearNvme4linkSpeed() bool`

HasSlotRearNvme4linkSpeed returns a boolean if a field has been set.

### GetSlotRearNvme4state

`func (o *BiosPolicyAllOf) GetSlotRearNvme4state() string`

GetSlotRearNvme4state returns the SlotRearNvme4state field if non-nil, zero value otherwise.

### GetSlotRearNvme4stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme4stateOk() (*string, bool)`

GetSlotRearNvme4stateOk returns a tuple with the SlotRearNvme4state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme4state

`func (o *BiosPolicyAllOf) SetSlotRearNvme4state(v string)`

SetSlotRearNvme4state sets SlotRearNvme4state field to given value.

### HasSlotRearNvme4state

`func (o *BiosPolicyAllOf) HasSlotRearNvme4state() bool`

HasSlotRearNvme4state returns a boolean if a field has been set.

### GetSlotRearNvme5state

`func (o *BiosPolicyAllOf) GetSlotRearNvme5state() string`

GetSlotRearNvme5state returns the SlotRearNvme5state field if non-nil, zero value otherwise.

### GetSlotRearNvme5stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme5stateOk() (*string, bool)`

GetSlotRearNvme5stateOk returns a tuple with the SlotRearNvme5state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme5state

`func (o *BiosPolicyAllOf) SetSlotRearNvme5state(v string)`

SetSlotRearNvme5state sets SlotRearNvme5state field to given value.

### HasSlotRearNvme5state

`func (o *BiosPolicyAllOf) HasSlotRearNvme5state() bool`

HasSlotRearNvme5state returns a boolean if a field has been set.

### GetSlotRearNvme6state

`func (o *BiosPolicyAllOf) GetSlotRearNvme6state() string`

GetSlotRearNvme6state returns the SlotRearNvme6state field if non-nil, zero value otherwise.

### GetSlotRearNvme6stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme6stateOk() (*string, bool)`

GetSlotRearNvme6stateOk returns a tuple with the SlotRearNvme6state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme6state

`func (o *BiosPolicyAllOf) SetSlotRearNvme6state(v string)`

SetSlotRearNvme6state sets SlotRearNvme6state field to given value.

### HasSlotRearNvme6state

`func (o *BiosPolicyAllOf) HasSlotRearNvme6state() bool`

HasSlotRearNvme6state returns a boolean if a field has been set.

### GetSlotRearNvme7state

`func (o *BiosPolicyAllOf) GetSlotRearNvme7state() string`

GetSlotRearNvme7state returns the SlotRearNvme7state field if non-nil, zero value otherwise.

### GetSlotRearNvme7stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme7stateOk() (*string, bool)`

GetSlotRearNvme7stateOk returns a tuple with the SlotRearNvme7state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme7state

`func (o *BiosPolicyAllOf) SetSlotRearNvme7state(v string)`

SetSlotRearNvme7state sets SlotRearNvme7state field to given value.

### HasSlotRearNvme7state

`func (o *BiosPolicyAllOf) HasSlotRearNvme7state() bool`

HasSlotRearNvme7state returns a boolean if a field has been set.

### GetSlotRearNvme8state

`func (o *BiosPolicyAllOf) GetSlotRearNvme8state() string`

GetSlotRearNvme8state returns the SlotRearNvme8state field if non-nil, zero value otherwise.

### GetSlotRearNvme8stateOk

`func (o *BiosPolicyAllOf) GetSlotRearNvme8stateOk() (*string, bool)`

GetSlotRearNvme8stateOk returns a tuple with the SlotRearNvme8state field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRearNvme8state

`func (o *BiosPolicyAllOf) SetSlotRearNvme8state(v string)`

SetSlotRearNvme8state sets SlotRearNvme8state field to given value.

### HasSlotRearNvme8state

`func (o *BiosPolicyAllOf) HasSlotRearNvme8state() bool`

HasSlotRearNvme8state returns a boolean if a field has been set.

### GetSlotRiser1linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser1linkSpeed() string`

GetSlotRiser1linkSpeed returns the SlotRiser1linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser1linkSpeedOk() (*string, bool)`

GetSlotRiser1linkSpeedOk returns a tuple with the SlotRiser1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser1linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser1linkSpeed(v string)`

SetSlotRiser1linkSpeed sets SlotRiser1linkSpeed field to given value.

### HasSlotRiser1linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser1linkSpeed() bool`

HasSlotRiser1linkSpeed returns a boolean if a field has been set.

### GetSlotRiser1slot1linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser1slot1linkSpeed() string`

GetSlotRiser1slot1linkSpeed returns the SlotRiser1slot1linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser1slot1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser1slot1linkSpeedOk() (*string, bool)`

GetSlotRiser1slot1linkSpeedOk returns a tuple with the SlotRiser1slot1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser1slot1linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser1slot1linkSpeed(v string)`

SetSlotRiser1slot1linkSpeed sets SlotRiser1slot1linkSpeed field to given value.

### HasSlotRiser1slot1linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser1slot1linkSpeed() bool`

HasSlotRiser1slot1linkSpeed returns a boolean if a field has been set.

### GetSlotRiser1slot2linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser1slot2linkSpeed() string`

GetSlotRiser1slot2linkSpeed returns the SlotRiser1slot2linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser1slot2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser1slot2linkSpeedOk() (*string, bool)`

GetSlotRiser1slot2linkSpeedOk returns a tuple with the SlotRiser1slot2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser1slot2linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser1slot2linkSpeed(v string)`

SetSlotRiser1slot2linkSpeed sets SlotRiser1slot2linkSpeed field to given value.

### HasSlotRiser1slot2linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser1slot2linkSpeed() bool`

HasSlotRiser1slot2linkSpeed returns a boolean if a field has been set.

### GetSlotRiser1slot3linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser1slot3linkSpeed() string`

GetSlotRiser1slot3linkSpeed returns the SlotRiser1slot3linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser1slot3linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser1slot3linkSpeedOk() (*string, bool)`

GetSlotRiser1slot3linkSpeedOk returns a tuple with the SlotRiser1slot3linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser1slot3linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser1slot3linkSpeed(v string)`

SetSlotRiser1slot3linkSpeed sets SlotRiser1slot3linkSpeed field to given value.

### HasSlotRiser1slot3linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser1slot3linkSpeed() bool`

HasSlotRiser1slot3linkSpeed returns a boolean if a field has been set.

### GetSlotRiser2linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser2linkSpeed() string`

GetSlotRiser2linkSpeed returns the SlotRiser2linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser2linkSpeedOk() (*string, bool)`

GetSlotRiser2linkSpeedOk returns a tuple with the SlotRiser2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser2linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser2linkSpeed(v string)`

SetSlotRiser2linkSpeed sets SlotRiser2linkSpeed field to given value.

### HasSlotRiser2linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser2linkSpeed() bool`

HasSlotRiser2linkSpeed returns a boolean if a field has been set.

### GetSlotRiser2slot4linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser2slot4linkSpeed() string`

GetSlotRiser2slot4linkSpeed returns the SlotRiser2slot4linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser2slot4linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser2slot4linkSpeedOk() (*string, bool)`

GetSlotRiser2slot4linkSpeedOk returns a tuple with the SlotRiser2slot4linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser2slot4linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser2slot4linkSpeed(v string)`

SetSlotRiser2slot4linkSpeed sets SlotRiser2slot4linkSpeed field to given value.

### HasSlotRiser2slot4linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser2slot4linkSpeed() bool`

HasSlotRiser2slot4linkSpeed returns a boolean if a field has been set.

### GetSlotRiser2slot5linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser2slot5linkSpeed() string`

GetSlotRiser2slot5linkSpeed returns the SlotRiser2slot5linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser2slot5linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser2slot5linkSpeedOk() (*string, bool)`

GetSlotRiser2slot5linkSpeedOk returns a tuple with the SlotRiser2slot5linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser2slot5linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser2slot5linkSpeed(v string)`

SetSlotRiser2slot5linkSpeed sets SlotRiser2slot5linkSpeed field to given value.

### HasSlotRiser2slot5linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser2slot5linkSpeed() bool`

HasSlotRiser2slot5linkSpeed returns a boolean if a field has been set.

### GetSlotRiser2slot6linkSpeed

`func (o *BiosPolicyAllOf) GetSlotRiser2slot6linkSpeed() string`

GetSlotRiser2slot6linkSpeed returns the SlotRiser2slot6linkSpeed field if non-nil, zero value otherwise.

### GetSlotRiser2slot6linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotRiser2slot6linkSpeedOk() (*string, bool)`

GetSlotRiser2slot6linkSpeedOk returns a tuple with the SlotRiser2slot6linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotRiser2slot6linkSpeed

`func (o *BiosPolicyAllOf) SetSlotRiser2slot6linkSpeed(v string)`

SetSlotRiser2slot6linkSpeed sets SlotRiser2slot6linkSpeed field to given value.

### HasSlotRiser2slot6linkSpeed

`func (o *BiosPolicyAllOf) HasSlotRiser2slot6linkSpeed() bool`

HasSlotRiser2slot6linkSpeed returns a boolean if a field has been set.

### GetSlotSasState

`func (o *BiosPolicyAllOf) GetSlotSasState() string`

GetSlotSasState returns the SlotSasState field if non-nil, zero value otherwise.

### GetSlotSasStateOk

`func (o *BiosPolicyAllOf) GetSlotSasStateOk() (*string, bool)`

GetSlotSasStateOk returns a tuple with the SlotSasState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotSasState

`func (o *BiosPolicyAllOf) SetSlotSasState(v string)`

SetSlotSasState sets SlotSasState field to given value.

### HasSlotSasState

`func (o *BiosPolicyAllOf) HasSlotSasState() bool`

HasSlotSasState returns a boolean if a field has been set.

### GetSlotSsdSlot1linkSpeed

`func (o *BiosPolicyAllOf) GetSlotSsdSlot1linkSpeed() string`

GetSlotSsdSlot1linkSpeed returns the SlotSsdSlot1linkSpeed field if non-nil, zero value otherwise.

### GetSlotSsdSlot1linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotSsdSlot1linkSpeedOk() (*string, bool)`

GetSlotSsdSlot1linkSpeedOk returns a tuple with the SlotSsdSlot1linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotSsdSlot1linkSpeed

`func (o *BiosPolicyAllOf) SetSlotSsdSlot1linkSpeed(v string)`

SetSlotSsdSlot1linkSpeed sets SlotSsdSlot1linkSpeed field to given value.

### HasSlotSsdSlot1linkSpeed

`func (o *BiosPolicyAllOf) HasSlotSsdSlot1linkSpeed() bool`

HasSlotSsdSlot1linkSpeed returns a boolean if a field has been set.

### GetSlotSsdSlot2linkSpeed

`func (o *BiosPolicyAllOf) GetSlotSsdSlot2linkSpeed() string`

GetSlotSsdSlot2linkSpeed returns the SlotSsdSlot2linkSpeed field if non-nil, zero value otherwise.

### GetSlotSsdSlot2linkSpeedOk

`func (o *BiosPolicyAllOf) GetSlotSsdSlot2linkSpeedOk() (*string, bool)`

GetSlotSsdSlot2linkSpeedOk returns a tuple with the SlotSsdSlot2linkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotSsdSlot2linkSpeed

`func (o *BiosPolicyAllOf) SetSlotSsdSlot2linkSpeed(v string)`

SetSlotSsdSlot2linkSpeed sets SlotSsdSlot2linkSpeed field to given value.

### HasSlotSsdSlot2linkSpeed

`func (o *BiosPolicyAllOf) HasSlotSsdSlot2linkSpeed() bool`

HasSlotSsdSlot2linkSpeed returns a boolean if a field has been set.

### GetSmee

`func (o *BiosPolicyAllOf) GetSmee() string`

GetSmee returns the Smee field if non-nil, zero value otherwise.

### GetSmeeOk

`func (o *BiosPolicyAllOf) GetSmeeOk() (*string, bool)`

GetSmeeOk returns a tuple with the Smee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmee

`func (o *BiosPolicyAllOf) SetSmee(v string)`

SetSmee sets Smee field to given value.

### HasSmee

`func (o *BiosPolicyAllOf) HasSmee() bool`

HasSmee returns a boolean if a field has been set.

### GetSmtMode

`func (o *BiosPolicyAllOf) GetSmtMode() string`

GetSmtMode returns the SmtMode field if non-nil, zero value otherwise.

### GetSmtModeOk

`func (o *BiosPolicyAllOf) GetSmtModeOk() (*string, bool)`

GetSmtModeOk returns a tuple with the SmtMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtMode

`func (o *BiosPolicyAllOf) SetSmtMode(v string)`

SetSmtMode sets SmtMode field to given value.

### HasSmtMode

`func (o *BiosPolicyAllOf) HasSmtMode() bool`

HasSmtMode returns a boolean if a field has been set.

### GetSnc

`func (o *BiosPolicyAllOf) GetSnc() string`

GetSnc returns the Snc field if non-nil, zero value otherwise.

### GetSncOk

`func (o *BiosPolicyAllOf) GetSncOk() (*string, bool)`

GetSncOk returns a tuple with the Snc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnc

`func (o *BiosPolicyAllOf) SetSnc(v string)`

SetSnc sets Snc field to given value.

### HasSnc

`func (o *BiosPolicyAllOf) HasSnc() bool`

HasSnc returns a boolean if a field has been set.

### GetSnoopyModeFor2lm

`func (o *BiosPolicyAllOf) GetSnoopyModeFor2lm() string`

GetSnoopyModeFor2lm returns the SnoopyModeFor2lm field if non-nil, zero value otherwise.

### GetSnoopyModeFor2lmOk

`func (o *BiosPolicyAllOf) GetSnoopyModeFor2lmOk() (*string, bool)`

GetSnoopyModeFor2lmOk returns a tuple with the SnoopyModeFor2lm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopyModeFor2lm

`func (o *BiosPolicyAllOf) SetSnoopyModeFor2lm(v string)`

SetSnoopyModeFor2lm sets SnoopyModeFor2lm field to given value.

### HasSnoopyModeFor2lm

`func (o *BiosPolicyAllOf) HasSnoopyModeFor2lm() bool`

HasSnoopyModeFor2lm returns a boolean if a field has been set.

### GetSnoopyModeForAd

`func (o *BiosPolicyAllOf) GetSnoopyModeForAd() string`

GetSnoopyModeForAd returns the SnoopyModeForAd field if non-nil, zero value otherwise.

### GetSnoopyModeForAdOk

`func (o *BiosPolicyAllOf) GetSnoopyModeForAdOk() (*string, bool)`

GetSnoopyModeForAdOk returns a tuple with the SnoopyModeForAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopyModeForAd

`func (o *BiosPolicyAllOf) SetSnoopyModeForAd(v string)`

SetSnoopyModeForAd sets SnoopyModeForAd field to given value.

### HasSnoopyModeForAd

`func (o *BiosPolicyAllOf) HasSnoopyModeForAd() bool`

HasSnoopyModeForAd returns a boolean if a field has been set.

### GetSparingMode

`func (o *BiosPolicyAllOf) GetSparingMode() string`

GetSparingMode returns the SparingMode field if non-nil, zero value otherwise.

### GetSparingModeOk

`func (o *BiosPolicyAllOf) GetSparingModeOk() (*string, bool)`

GetSparingModeOk returns a tuple with the SparingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparingMode

`func (o *BiosPolicyAllOf) SetSparingMode(v string)`

SetSparingMode sets SparingMode field to given value.

### HasSparingMode

`func (o *BiosPolicyAllOf) HasSparingMode() bool`

HasSparingMode returns a boolean if a field has been set.

### GetSrIov

`func (o *BiosPolicyAllOf) GetSrIov() string`

GetSrIov returns the SrIov field if non-nil, zero value otherwise.

### GetSrIovOk

`func (o *BiosPolicyAllOf) GetSrIovOk() (*string, bool)`

GetSrIovOk returns a tuple with the SrIov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrIov

`func (o *BiosPolicyAllOf) SetSrIov(v string)`

SetSrIov sets SrIov field to given value.

### HasSrIov

`func (o *BiosPolicyAllOf) HasSrIov() bool`

HasSrIov returns a boolean if a field has been set.

### GetStreamerPrefetch

`func (o *BiosPolicyAllOf) GetStreamerPrefetch() string`

GetStreamerPrefetch returns the StreamerPrefetch field if non-nil, zero value otherwise.

### GetStreamerPrefetchOk

`func (o *BiosPolicyAllOf) GetStreamerPrefetchOk() (*string, bool)`

GetStreamerPrefetchOk returns a tuple with the StreamerPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerPrefetch

`func (o *BiosPolicyAllOf) SetStreamerPrefetch(v string)`

SetStreamerPrefetch sets StreamerPrefetch field to given value.

### HasStreamerPrefetch

`func (o *BiosPolicyAllOf) HasStreamerPrefetch() bool`

HasStreamerPrefetch returns a boolean if a field has been set.

### GetSvmMode

`func (o *BiosPolicyAllOf) GetSvmMode() string`

GetSvmMode returns the SvmMode field if non-nil, zero value otherwise.

### GetSvmModeOk

`func (o *BiosPolicyAllOf) GetSvmModeOk() (*string, bool)`

GetSvmModeOk returns a tuple with the SvmMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmMode

`func (o *BiosPolicyAllOf) SetSvmMode(v string)`

SetSvmMode sets SvmMode field to given value.

### HasSvmMode

`func (o *BiosPolicyAllOf) HasSvmMode() bool`

HasSvmMode returns a boolean if a field has been set.

### GetTerminalType

`func (o *BiosPolicyAllOf) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *BiosPolicyAllOf) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *BiosPolicyAllOf) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.

### HasTerminalType

`func (o *BiosPolicyAllOf) HasTerminalType() bool`

HasTerminalType returns a boolean if a field has been set.

### GetTpmControl

`func (o *BiosPolicyAllOf) GetTpmControl() string`

GetTpmControl returns the TpmControl field if non-nil, zero value otherwise.

### GetTpmControlOk

`func (o *BiosPolicyAllOf) GetTpmControlOk() (*string, bool)`

GetTpmControlOk returns a tuple with the TpmControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmControl

`func (o *BiosPolicyAllOf) SetTpmControl(v string)`

SetTpmControl sets TpmControl field to given value.

### HasTpmControl

`func (o *BiosPolicyAllOf) HasTpmControl() bool`

HasTpmControl returns a boolean if a field has been set.

### GetTpmPendingOperation

`func (o *BiosPolicyAllOf) GetTpmPendingOperation() string`

GetTpmPendingOperation returns the TpmPendingOperation field if non-nil, zero value otherwise.

### GetTpmPendingOperationOk

`func (o *BiosPolicyAllOf) GetTpmPendingOperationOk() (*string, bool)`

GetTpmPendingOperationOk returns a tuple with the TpmPendingOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmPendingOperation

`func (o *BiosPolicyAllOf) SetTpmPendingOperation(v string)`

SetTpmPendingOperation sets TpmPendingOperation field to given value.

### HasTpmPendingOperation

`func (o *BiosPolicyAllOf) HasTpmPendingOperation() bool`

HasTpmPendingOperation returns a boolean if a field has been set.

### GetTpmSupport

`func (o *BiosPolicyAllOf) GetTpmSupport() string`

GetTpmSupport returns the TpmSupport field if non-nil, zero value otherwise.

### GetTpmSupportOk

`func (o *BiosPolicyAllOf) GetTpmSupportOk() (*string, bool)`

GetTpmSupportOk returns a tuple with the TpmSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmSupport

`func (o *BiosPolicyAllOf) SetTpmSupport(v string)`

SetTpmSupport sets TpmSupport field to given value.

### HasTpmSupport

`func (o *BiosPolicyAllOf) HasTpmSupport() bool`

HasTpmSupport returns a boolean if a field has been set.

### GetTsme

`func (o *BiosPolicyAllOf) GetTsme() string`

GetTsme returns the Tsme field if non-nil, zero value otherwise.

### GetTsmeOk

`func (o *BiosPolicyAllOf) GetTsmeOk() (*string, bool)`

GetTsmeOk returns a tuple with the Tsme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsme

`func (o *BiosPolicyAllOf) SetTsme(v string)`

SetTsme sets Tsme field to given value.

### HasTsme

`func (o *BiosPolicyAllOf) HasTsme() bool`

HasTsme returns a boolean if a field has been set.

### GetTxtSupport

`func (o *BiosPolicyAllOf) GetTxtSupport() string`

GetTxtSupport returns the TxtSupport field if non-nil, zero value otherwise.

### GetTxtSupportOk

`func (o *BiosPolicyAllOf) GetTxtSupportOk() (*string, bool)`

GetTxtSupportOk returns a tuple with the TxtSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtSupport

`func (o *BiosPolicyAllOf) SetTxtSupport(v string)`

SetTxtSupport sets TxtSupport field to given value.

### HasTxtSupport

`func (o *BiosPolicyAllOf) HasTxtSupport() bool`

HasTxtSupport returns a boolean if a field has been set.

### GetUcsmBootOrderRule

`func (o *BiosPolicyAllOf) GetUcsmBootOrderRule() string`

GetUcsmBootOrderRule returns the UcsmBootOrderRule field if non-nil, zero value otherwise.

### GetUcsmBootOrderRuleOk

`func (o *BiosPolicyAllOf) GetUcsmBootOrderRuleOk() (*string, bool)`

GetUcsmBootOrderRuleOk returns a tuple with the UcsmBootOrderRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsmBootOrderRule

`func (o *BiosPolicyAllOf) SetUcsmBootOrderRule(v string)`

SetUcsmBootOrderRule sets UcsmBootOrderRule field to given value.

### HasUcsmBootOrderRule

`func (o *BiosPolicyAllOf) HasUcsmBootOrderRule() bool`

HasUcsmBootOrderRule returns a boolean if a field has been set.

### GetUfsDisable

`func (o *BiosPolicyAllOf) GetUfsDisable() string`

GetUfsDisable returns the UfsDisable field if non-nil, zero value otherwise.

### GetUfsDisableOk

`func (o *BiosPolicyAllOf) GetUfsDisableOk() (*string, bool)`

GetUfsDisableOk returns a tuple with the UfsDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUfsDisable

`func (o *BiosPolicyAllOf) SetUfsDisable(v string)`

SetUfsDisable sets UfsDisable field to given value.

### HasUfsDisable

`func (o *BiosPolicyAllOf) HasUfsDisable() bool`

HasUfsDisable returns a boolean if a field has been set.

### GetUmaBasedClustering

`func (o *BiosPolicyAllOf) GetUmaBasedClustering() string`

GetUmaBasedClustering returns the UmaBasedClustering field if non-nil, zero value otherwise.

### GetUmaBasedClusteringOk

`func (o *BiosPolicyAllOf) GetUmaBasedClusteringOk() (*string, bool)`

GetUmaBasedClusteringOk returns a tuple with the UmaBasedClustering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmaBasedClustering

`func (o *BiosPolicyAllOf) SetUmaBasedClustering(v string)`

SetUmaBasedClustering sets UmaBasedClustering field to given value.

### HasUmaBasedClustering

`func (o *BiosPolicyAllOf) HasUmaBasedClustering() bool`

HasUmaBasedClustering returns a boolean if a field has been set.

### GetUsbEmul6064

`func (o *BiosPolicyAllOf) GetUsbEmul6064() string`

GetUsbEmul6064 returns the UsbEmul6064 field if non-nil, zero value otherwise.

### GetUsbEmul6064Ok

`func (o *BiosPolicyAllOf) GetUsbEmul6064Ok() (*string, bool)`

GetUsbEmul6064Ok returns a tuple with the UsbEmul6064 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbEmul6064

`func (o *BiosPolicyAllOf) SetUsbEmul6064(v string)`

SetUsbEmul6064 sets UsbEmul6064 field to given value.

### HasUsbEmul6064

`func (o *BiosPolicyAllOf) HasUsbEmul6064() bool`

HasUsbEmul6064 returns a boolean if a field has been set.

### GetUsbPortFront

`func (o *BiosPolicyAllOf) GetUsbPortFront() string`

GetUsbPortFront returns the UsbPortFront field if non-nil, zero value otherwise.

### GetUsbPortFrontOk

`func (o *BiosPolicyAllOf) GetUsbPortFrontOk() (*string, bool)`

GetUsbPortFrontOk returns a tuple with the UsbPortFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortFront

`func (o *BiosPolicyAllOf) SetUsbPortFront(v string)`

SetUsbPortFront sets UsbPortFront field to given value.

### HasUsbPortFront

`func (o *BiosPolicyAllOf) HasUsbPortFront() bool`

HasUsbPortFront returns a boolean if a field has been set.

### GetUsbPortInternal

`func (o *BiosPolicyAllOf) GetUsbPortInternal() string`

GetUsbPortInternal returns the UsbPortInternal field if non-nil, zero value otherwise.

### GetUsbPortInternalOk

`func (o *BiosPolicyAllOf) GetUsbPortInternalOk() (*string, bool)`

GetUsbPortInternalOk returns a tuple with the UsbPortInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortInternal

`func (o *BiosPolicyAllOf) SetUsbPortInternal(v string)`

SetUsbPortInternal sets UsbPortInternal field to given value.

### HasUsbPortInternal

`func (o *BiosPolicyAllOf) HasUsbPortInternal() bool`

HasUsbPortInternal returns a boolean if a field has been set.

### GetUsbPortKvm

`func (o *BiosPolicyAllOf) GetUsbPortKvm() string`

GetUsbPortKvm returns the UsbPortKvm field if non-nil, zero value otherwise.

### GetUsbPortKvmOk

`func (o *BiosPolicyAllOf) GetUsbPortKvmOk() (*string, bool)`

GetUsbPortKvmOk returns a tuple with the UsbPortKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortKvm

`func (o *BiosPolicyAllOf) SetUsbPortKvm(v string)`

SetUsbPortKvm sets UsbPortKvm field to given value.

### HasUsbPortKvm

`func (o *BiosPolicyAllOf) HasUsbPortKvm() bool`

HasUsbPortKvm returns a boolean if a field has been set.

### GetUsbPortRear

`func (o *BiosPolicyAllOf) GetUsbPortRear() string`

GetUsbPortRear returns the UsbPortRear field if non-nil, zero value otherwise.

### GetUsbPortRearOk

`func (o *BiosPolicyAllOf) GetUsbPortRearOk() (*string, bool)`

GetUsbPortRearOk returns a tuple with the UsbPortRear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortRear

`func (o *BiosPolicyAllOf) SetUsbPortRear(v string)`

SetUsbPortRear sets UsbPortRear field to given value.

### HasUsbPortRear

`func (o *BiosPolicyAllOf) HasUsbPortRear() bool`

HasUsbPortRear returns a boolean if a field has been set.

### GetUsbPortSdCard

`func (o *BiosPolicyAllOf) GetUsbPortSdCard() string`

GetUsbPortSdCard returns the UsbPortSdCard field if non-nil, zero value otherwise.

### GetUsbPortSdCardOk

`func (o *BiosPolicyAllOf) GetUsbPortSdCardOk() (*string, bool)`

GetUsbPortSdCardOk returns a tuple with the UsbPortSdCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortSdCard

`func (o *BiosPolicyAllOf) SetUsbPortSdCard(v string)`

SetUsbPortSdCard sets UsbPortSdCard field to given value.

### HasUsbPortSdCard

`func (o *BiosPolicyAllOf) HasUsbPortSdCard() bool`

HasUsbPortSdCard returns a boolean if a field has been set.

### GetUsbPortVmedia

`func (o *BiosPolicyAllOf) GetUsbPortVmedia() string`

GetUsbPortVmedia returns the UsbPortVmedia field if non-nil, zero value otherwise.

### GetUsbPortVmediaOk

`func (o *BiosPolicyAllOf) GetUsbPortVmediaOk() (*string, bool)`

GetUsbPortVmediaOk returns a tuple with the UsbPortVmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortVmedia

`func (o *BiosPolicyAllOf) SetUsbPortVmedia(v string)`

SetUsbPortVmedia sets UsbPortVmedia field to given value.

### HasUsbPortVmedia

`func (o *BiosPolicyAllOf) HasUsbPortVmedia() bool`

HasUsbPortVmedia returns a boolean if a field has been set.

### GetUsbXhciSupport

`func (o *BiosPolicyAllOf) GetUsbXhciSupport() string`

GetUsbXhciSupport returns the UsbXhciSupport field if non-nil, zero value otherwise.

### GetUsbXhciSupportOk

`func (o *BiosPolicyAllOf) GetUsbXhciSupportOk() (*string, bool)`

GetUsbXhciSupportOk returns a tuple with the UsbXhciSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbXhciSupport

`func (o *BiosPolicyAllOf) SetUsbXhciSupport(v string)`

SetUsbXhciSupport sets UsbXhciSupport field to given value.

### HasUsbXhciSupport

`func (o *BiosPolicyAllOf) HasUsbXhciSupport() bool`

HasUsbXhciSupport returns a boolean if a field has been set.

### GetVgaPriority

`func (o *BiosPolicyAllOf) GetVgaPriority() string`

GetVgaPriority returns the VgaPriority field if non-nil, zero value otherwise.

### GetVgaPriorityOk

`func (o *BiosPolicyAllOf) GetVgaPriorityOk() (*string, bool)`

GetVgaPriorityOk returns a tuple with the VgaPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgaPriority

`func (o *BiosPolicyAllOf) SetVgaPriority(v string)`

SetVgaPriority sets VgaPriority field to given value.

### HasVgaPriority

`func (o *BiosPolicyAllOf) HasVgaPriority() bool`

HasVgaPriority returns a boolean if a field has been set.

### GetVmdEnable

`func (o *BiosPolicyAllOf) GetVmdEnable() string`

GetVmdEnable returns the VmdEnable field if non-nil, zero value otherwise.

### GetVmdEnableOk

`func (o *BiosPolicyAllOf) GetVmdEnableOk() (*string, bool)`

GetVmdEnableOk returns a tuple with the VmdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdEnable

`func (o *BiosPolicyAllOf) SetVmdEnable(v string)`

SetVmdEnable sets VmdEnable field to given value.

### HasVmdEnable

`func (o *BiosPolicyAllOf) HasVmdEnable() bool`

HasVmdEnable returns a boolean if a field has been set.

### GetVolMemoryMode

`func (o *BiosPolicyAllOf) GetVolMemoryMode() string`

GetVolMemoryMode returns the VolMemoryMode field if non-nil, zero value otherwise.

### GetVolMemoryModeOk

`func (o *BiosPolicyAllOf) GetVolMemoryModeOk() (*string, bool)`

GetVolMemoryModeOk returns a tuple with the VolMemoryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolMemoryMode

`func (o *BiosPolicyAllOf) SetVolMemoryMode(v string)`

SetVolMemoryMode sets VolMemoryMode field to given value.

### HasVolMemoryMode

`func (o *BiosPolicyAllOf) HasVolMemoryMode() bool`

HasVolMemoryMode returns a boolean if a field has been set.

### GetWorkLoadConfig

`func (o *BiosPolicyAllOf) GetWorkLoadConfig() string`

GetWorkLoadConfig returns the WorkLoadConfig field if non-nil, zero value otherwise.

### GetWorkLoadConfigOk

`func (o *BiosPolicyAllOf) GetWorkLoadConfigOk() (*string, bool)`

GetWorkLoadConfigOk returns a tuple with the WorkLoadConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLoadConfig

`func (o *BiosPolicyAllOf) SetWorkLoadConfig(v string)`

SetWorkLoadConfig sets WorkLoadConfig field to given value.

### HasWorkLoadConfig

`func (o *BiosPolicyAllOf) HasWorkLoadConfig() bool`

HasWorkLoadConfig returns a boolean if a field has been set.

### GetXptPrefetch

`func (o *BiosPolicyAllOf) GetXptPrefetch() string`

GetXptPrefetch returns the XptPrefetch field if non-nil, zero value otherwise.

### GetXptPrefetchOk

`func (o *BiosPolicyAllOf) GetXptPrefetchOk() (*string, bool)`

GetXptPrefetchOk returns a tuple with the XptPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXptPrefetch

`func (o *BiosPolicyAllOf) SetXptPrefetch(v string)`

SetXptPrefetch sets XptPrefetch field to given value.

### HasXptPrefetch

`func (o *BiosPolicyAllOf) HasXptPrefetch() bool`

HasXptPrefetch returns a boolean if a field has been set.

### GetOrganization

`func (o *BiosPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BiosPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BiosPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BiosPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *BiosPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BiosPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BiosPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BiosPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *BiosPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *BiosPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


