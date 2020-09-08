package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceBiosPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBiosPolicyRead,
		Schema: map[string]*schema.Schema{
			"acs_control_gpu1state": {
				Description: "BIOS Token for setting ACS Control GPU-1 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu2state": {
				Description: "BIOS Token for setting ACS Control GPU-2 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu3state": {
				Description: "BIOS Token for setting ACS Control GPU-3 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu4state": {
				Description: "BIOS Token for setting ACS Control GPU-4 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu5state": {
				Description: "BIOS Token for setting ACS Control GPU-5 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu6state": {
				Description: "BIOS Token for setting ACS Control GPU-6 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu7state": {
				Description: "BIOS Token for setting ACS Control GPU-7 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu8state": {
				Description: "BIOS Token for setting ACS Control GPU-8 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot11state": {
				Description: "BIOS Token for setting ACS Control Slot 11 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot12state": {
				Description: "BIOS Token for setting ACS Control Slot 12 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot13state": {
				Description: "BIOS Token for setting ACS Control Slot 13 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot14state": {
				Description: "BIOS Token for setting ACS Control Slot 14 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"adjacent_cache_line_prefetch": {
				Description: "BIOS Token for setting Adjacent Cache Line Prefetcher configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"all_usb_devices": {
				Description: "BIOS Token for setting All USB Devices configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"altitude": {
				Description: "BIOS Token for setting Altitude configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1500-m` - Value - 1500-m for configuring Altitude token.\n* `300-m` - Value - 300-m for configuring Altitude token.\n* `3000-m` - Value - 3000-m for configuring Altitude token.\n* `900-m` - Value - 900-m for configuring Altitude token.\n* `auto` - Value - auto for configuring Altitude token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"aspm_support": {
				Description: "BIOS Token for setting ASPM Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring AspmSupport token.\n* `Disabled` - Value - Disabled for configuring AspmSupport token.\n* `Force L0s` - Value - Force L0s for configuring AspmSupport token.\n* `L1 Only` - Value - L1 Only for configuring AspmSupport token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"assert_nmi_on_perr": {
				Description: "BIOS Token for setting Assert NMI on PERR configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"assert_nmi_on_serr": {
				Description: "BIOS Token for setting Assert NMI on SERR configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auto_cc_state": {
				Description: "BIOS Token for setting Autonomous Core C-state configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"autonumous_cstate_enable": {
				Description: "BIOS Token for setting CPU Autonomous Cstate configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"baud_rate": {
				Description: "BIOS Token for setting Baud rate configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `115200` - Value - 115200 for configuring BaudRate token.\n* `19200` - Value - 19200 for configuring BaudRate token.\n* `38400` - Value - 38400 for configuring BaudRate token.\n* `57600` - Value - 57600 for configuring BaudRate token.\n* `9600` - Value - 9600 for configuring BaudRate token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"bme_dma_mitigation": {
				Description: "BIOS Token for setting BME DMA Mitigation configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_num_retry": {
				Description: "BIOS Token for setting Number of Retries configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `13` - Value - 13 for configuring BootOptionNumRetry token.\n* `5` - Value - 5 for configuring BootOptionNumRetry token.\n* `Infinite` - Value - Infinite for configuring BootOptionNumRetry token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_re_cool_down": {
				Description: "BIOS Token for setting Cool Down Time  (sec) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `15` - Value - 15 for configuring BootOptionReCoolDown token.\n* `45` - Value - 45 for configuring BootOptionReCoolDown token.\n* `90` - Value - 90 for configuring BootOptionReCoolDown token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_retry": {
				Description: "BIOS Token for setting Boot option retry configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_performance_mode": {
				Description: "BIOS Token for setting Boot Performance Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Max Efficient` - Value - Max Efficient for configuring BootPerformanceMode token.\n* `Max Performance` - Value - Max Performance for configuring BootPerformanceMode token.\n* `Set by Intel NM` - Value - Set by Intel NM for configuring BootPerformanceMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_cpb": {
				Description: "BIOS Token for setting Core Performance Boost configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuCpb token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuCpb token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_gen_downcore_ctrl": {
				Description: "BIOS Token for setting Downcore control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `FOUR (2 + 2)` - Value - FOUR (2 + 2) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `FOUR (4 + 0)` - Value - FOUR (4 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `SIX (3 + 3)` - Value - SIX (3 + 3) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `THREE (3 + 0)` - Value - THREE (3 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `TWO (1 + 1)` - Value - TWO (1 + 1) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `TWO (2 + 0)` - Value - TWO (2 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_global_cstate_ctrl": {
				Description: "BIOS Token for setting Global C-state Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuGlobalCstateCtrl token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuGlobalCstateCtrl token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuGlobalCstateCtrl token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_l1stream_hw_prefetcher": {
				Description: "BIOS Token for setting L1 Stream HW Prefetcher configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuL1streamHwPrefetcher token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuL1streamHwPrefetcher token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuL1streamHwPrefetcher token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_l2stream_hw_prefetcher": {
				Description: "BIOS Token for setting L2 Stream HW Prefetcher configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuL2streamHwPrefetcher token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuL2streamHwPrefetcher token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuL2streamHwPrefetcher token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_determinism_slider": {
				Description: "BIOS Token for setting Determinism Slider configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnDeterminismSlider token.\n* `Performance` - Value - Performance for configuring CbsCmnDeterminismSlider token.\n* `Power` - Value - Power for configuring CbsCmnDeterminismSlider token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_gnb_nb_iommu": {
				Description: "BIOS Token for setting IOMMU configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbNbIommu token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbNbIommu token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbNbIommu token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_mem_ctrl_bank_group_swap_ddr4": {
				Description: "BIOS Token for setting Bank Group Swap configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.\n* `disabled` - Value - disabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.\n* `enabled` - Value - enabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_mem_map_bank_interleave_ddr4": {
				Description: "BIOS Token for setting Chipselect Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnMemMapBankInterleaveDdr4 token.\n* `disabled` - Value - disabled for configuring CbsCmnMemMapBankInterleaveDdr4 token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmnc_tdp_ctl": {
				Description: "BIOS Token for setting cTDP Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmncTdpCtl token.\n* `Manual` - Value - Manual for configuring CbsCmncTdpCtl token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_df_cmn_mem_intlv": {
				Description: "BIOS Token for setting Memory interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlv token.\n* `Channel` - Value - Channel for configuring CbsDfCmnMemIntlv token.\n* `Channel NPS2` - Value - Channel NPS2 for configuring CbsDfCmnMemIntlv token.\n* `Channel NPS4` - Value - Channel NPS4 for configuring CbsDfCmnMemIntlv token.\n* `Die` - Value - Die for configuring CbsDfCmnMemIntlv token.\n* `Die NPS1` - Value - Die NPS1 for configuring CbsDfCmnMemIntlv token.\n* `None` - Value - None for configuring CbsDfCmnMemIntlv token.\n* `Socket` - Value - Socket for configuring CbsDfCmnMemIntlv token.\n* `Socket NPS0` - Value - Socket NPS0 for configuring CbsDfCmnMemIntlv token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_df_cmn_mem_intlv_size": {
				Description: "BIOS Token for setting Memory interleaving size configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1 KB` - Value - 1 KiB for configuring CbsDfCmnMemIntlvSize token.\n* `2 KB` - Value - 2 KiB for configuring CbsDfCmnMemIntlvSize token.\n* `256 Bytes` - Value - 256 Bytes for configuring CbsDfCmnMemIntlvSize token.\n* `512 Bytes` - Value - 512 Bytes for configuring CbsDfCmnMemIntlvSize token.\n* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlvSize token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cdn_enable": {
				Description: "BIOS Token for setting Consistent Device Naming configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cdn_support": {
				Description: "BIOS Token for setting CDN Support for LOM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring CdnSupport token.\n* `enabled` - Value - enabled for configuring CdnSupport token.\n* `LOMs Only` - Value - LOMs Only for configuring CdnSupport token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"channel_inter_leave": {
				Description: "BIOS Token for setting Channel Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1-way` - Value - 1-way for configuring ChannelInterLeave token.\n* `2-way` - Value - 2-way for configuring ChannelInterLeave token.\n* `3-way` - Value - 3-way for configuring ChannelInterLeave token.\n* `4-way` - Value - 4-way for configuring ChannelInterLeave token.\n* `auto` - Value - auto for configuring ChannelInterLeave token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_adaptive_mem_training": {
				Description: "BIOS Token for setting Adaptive Memory Training configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_debug_level": {
				Description: "BIOS Token for setting BIOS Techlog Level configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Maximum` - Value - Maximum for configuring CiscoDebugLevel token.\n* `Minimum` - Value - Minimum for configuring CiscoDebugLevel token.\n* `Normal` - Value - Normal for configuring CiscoDebugLevel token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_oprom_launch_optimization": {
				Description: "BIOS Token for setting OptionROM Launch Optimization configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cke_low_policy": {
				Description: "BIOS Token for setting CKE Low Policy configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring CkeLowPolicy token.\n* `disabled` - Value - disabled for configuring CkeLowPolicy token.\n* `fast` - Value - fast for configuring CkeLowPolicy token.\n* `slow` - Value - slow for configuring CkeLowPolicy token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"closed_loop_therm_throtl": {
				Description: "BIOS Token for setting Closed Loop Therm Throt configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cmci_enable": {
				Description: "BIOS Token for setting Processor CMCI configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_tdp": {
				Description: "BIOS Token for setting Config TDP configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_tdp_level": {
				Description: "BIOS Token for setting Configurable TDP Level configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Level 1` - Value - Level 1 for configuring ConfigTdpLevel token.\n* `Level 2` - Value - Level 2 for configuring ConfigTdpLevel token.\n* `Normal` - Value - Normal for configuring ConfigTdpLevel token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"console_redirection": {
				Description: "BIOS Token for setting Console redirection configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `com-0` - Value - com-0 for configuring ConsoleRedirection token.\n* `com-1` - Value - com-1 for configuring ConsoleRedirection token.\n* `disabled` - Value - disabled for configuring ConsoleRedirection token.\n* `enabled` - Value - enabled for configuring ConsoleRedirection token.\n* `serial-port-a` - Value - serial-port-a for configuring ConsoleRedirection token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"core_multi_processing": {
				Description: "BIOS Token for setting Core MultiProcessing configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1` - Value - 1 for configuring CoreMultiProcessing token.\n* `10` - Value - 10 for configuring CoreMultiProcessing token.\n* `11` - Value - 11 for configuring CoreMultiProcessing token.\n* `12` - Value - 12 for configuring CoreMultiProcessing token.\n* `13` - Value - 13 for configuring CoreMultiProcessing token.\n* `14` - Value - 14 for configuring CoreMultiProcessing token.\n* `15` - Value - 15 for configuring CoreMultiProcessing token.\n* `16` - Value - 16 for configuring CoreMultiProcessing token.\n* `17` - Value - 17 for configuring CoreMultiProcessing token.\n* `18` - Value - 18 for configuring CoreMultiProcessing token.\n* `19` - Value - 19 for configuring CoreMultiProcessing token.\n* `2` - Value - 2 for configuring CoreMultiProcessing token.\n* `20` - Value - 20 for configuring CoreMultiProcessing token.\n* `21` - Value - 21 for configuring CoreMultiProcessing token.\n* `22` - Value - 22 for configuring CoreMultiProcessing token.\n* `23` - Value - 23 for configuring CoreMultiProcessing token.\n* `24` - Value - 24 for configuring CoreMultiProcessing token.\n* `25` - Value - 25 for configuring CoreMultiProcessing token.\n* `26` - Value - 26 for configuring CoreMultiProcessing token.\n* `27` - Value - 27 for configuring CoreMultiProcessing token.\n* `28` - Value - 28 for configuring CoreMultiProcessing token.\n* `3` - Value - 3 for configuring CoreMultiProcessing token.\n* `4` - Value - 4 for configuring CoreMultiProcessing token.\n* `5` - Value - 5 for configuring CoreMultiProcessing token.\n* `6` - Value - 6 for configuring CoreMultiProcessing token.\n* `7` - Value - 7 for configuring CoreMultiProcessing token.\n* `8` - Value - 8 for configuring CoreMultiProcessing token.\n* `9` - Value - 9 for configuring CoreMultiProcessing token.\n* `all` - Value - all for configuring CoreMultiProcessing token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_energy_performance": {
				Description: "BIOS Token for setting Energy Performance configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `balanced-energy` - Value - balanced-energy for configuring CpuEnergyPerformance token.\n* `balanced-performance` - Value - balanced-performance for configuring CpuEnergyPerformance token.\n* `balanced-power` - Value - balanced-power for configuring CpuEnergyPerformance token.\n* `energy-efficient` - Value - energy-efficient for configuring CpuEnergyPerformance token.\n* `performance` - Value - performance for configuring CpuEnergyPerformance token.\n* `power` - Value - power for configuring CpuEnergyPerformance token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_frequency_floor": {
				Description: "BIOS Token for setting Frequency Floor Override configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_performance": {
				Description: "BIOS Token for setting CPU Performance configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `custom` - Value - custom for configuring CpuPerformance token.\n* `enterprise` - Value - enterprise for configuring CpuPerformance token.\n* `high-throughput` - Value - high-throughput for configuring CpuPerformance token.\n* `hpc` - Value - hpc for configuring CpuPerformance token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_power_management": {
				Description: "BIOS Token for setting Power Technology configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `custom` - Value - custom for configuring CpuPowerManagement token.\n* `disabled` - Value - disabled for configuring CpuPowerManagement token.\n* `energy-efficient` - Value - energy-efficient for configuring CpuPowerManagement token.\n* `performance` - Value - performance for configuring CpuPowerManagement token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cr_qos": {
				Description: "BIOS Token for setting CR QoS configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring CrQos token.\n* `Recipe 1` - Value - Recipe 1 for configuring CrQos token.\n* `Recipe 2` - Value - Recipe 2 for configuring CrQos token.\n* `Recipe 3` - Value - Recipe 3 for configuring CrQos token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"crfastgo_config": {
				Description: "BIOS Token for setting CR FastGo Config configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CrfastgoConfig token.\n* `Default` - Value - Default for configuring CrfastgoConfig token.\n* `Option 1` - Value - Option 1 for configuring CrfastgoConfig token.\n* `Option 2` - Value - Option 2 for configuring CrfastgoConfig token.\n* `Option 3` - Value - Option 3 for configuring CrfastgoConfig token.\n* `Option 4` - Value - Option 4 for configuring CrfastgoConfig token.\n* `Option 5` - Value - Option 5 for configuring CrfastgoConfig token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dcpmm_firmware_downgrade": {
				Description: "BIOS Token for setting DCPMM Firmware Downgrade configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"demand_scrub": {
				Description: "BIOS Token for setting Demand Scrub configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"direct_cache_access": {
				Description: "BIOS Token for setting Direct Cache Access Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring DirectCacheAccess token.\n* `disabled` - Value - disabled for configuring DirectCacheAccess token.\n* `enabled` - Value - enabled for configuring DirectCacheAccess token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dram_clock_throttling": {
				Description: "BIOS Token for setting DRAM Clock Throttling configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring DramClockThrottling token.\n* `Balanced` - Value - Balanced for configuring DramClockThrottling token.\n* `Energy Efficient` - Value - Energy Efficient for configuring DramClockThrottling token.\n* `Performance` - Value - Performance for configuring DramClockThrottling token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dram_refresh_rate": {
				Description: "BIOS Token for setting DRAM Refresh Rate configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1x` - Value - 1x for configuring DramRefreshRate token.\n* `2x` - Value - 2x for configuring DramRefreshRate token.\n* `3x` - Value - 3x for configuring DramRefreshRate token.\n* `4x` - Value - 4x for configuring DramRefreshRate token.\n* `Auto` - Value - Auto for configuring DramRefreshRate token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_clock_spread_spec": {
				Description: "BIOS Token for setting External SSC Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"energy_efficient_turbo": {
				Description: "BIOS Token for setting Energy Efficient Turbo configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"eng_perf_tuning": {
				Description: "BIOS Token for setting Energy Performance Tuning configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `BIOS` - Value - BIOS for configuring EngPerfTuning token.\n* `OS` - Value - OS for configuring EngPerfTuning token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enhanced_intel_speed_step_tech": {
				Description: "BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"epp_profile": {
				Description: "BIOS Token for setting EPP Profile configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Balanced Performance` - Value - Balanced Performance for configuring EppProfile token.\n* `Balanced Power` - Value - Balanced Power for configuring EppProfile token.\n* `Performance` - Value - Performance for configuring EppProfile token.\n* `Power` - Value - Power for configuring EppProfile token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"execute_disable_bit": {
				Description: "BIOS Token for setting Execute Disable Bit configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"extended_apic": {
				Description: "BIOS Token for setting Local X2 Apic configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring ExtendedApic token.\n* `enabled` - Value - enabled for configuring ExtendedApic token.\n* `X2APIC` - Value - X2APIC for configuring ExtendedApic token.\n* `XAPIC` - Value - XAPIC for configuring ExtendedApic token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"flow_control": {
				Description: "BIOS Token for setting Flow Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `none` - Value - none for configuring FlowControl token.\n* `rts-cts` - Value - rts-cts for configuring FlowControl token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"frb2enable": {
				Description: "BIOS Token for setting FRB-2 Timer configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hardware_prefetch": {
				Description: "BIOS Token for setting Hardware Prefetcher configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hwpm_enable": {
				Description: "BIOS Token for setting CPU Hardware Power Management configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring HwpmEnable token.\n* `HWPM Native Mode` - Value - HWPM Native Mode for configuring HwpmEnable token.\n* `HWPM OOB Mode` - Value - HWPM OOB Mode for configuring HwpmEnable token.\n* `NATIVE MODE` - Value - NATIVE MODE for configuring HwpmEnable token.\n* `Native Mode with no Legacy` - Value - Native Mode with no Legacy for configuring HwpmEnable token.\n* `OOB MODE` - Value - OOB MODE for configuring HwpmEnable token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"imc_interleave": {
				Description: "BIOS Token for setting IMC Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1-way Interleave` - Value - 1-way Interleave for configuring ImcInterleave token.\n* `2-way Interleave` - Value - 2-way Interleave for configuring ImcInterleave token.\n* `Auto` - Value - Auto for configuring ImcInterleave token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_hyper_threading_tech": {
				Description: "BIOS Token for setting Intel HyperThreading Tech configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_speed_select": {
				Description: "BIOS Token for setting Intel Speed Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Base` - Value - Base for configuring IntelSpeedSelect token.\n* `Config 1` - Value - Config 1 for configuring IntelSpeedSelect token.\n* `Config 2` - Value - Config 2 for configuring IntelSpeedSelect token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_turbo_boost_tech": {
				Description: "BIOS Token for setting Intel Turbo Boost Tech configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_virtualization_technology": {
				Description: "BIOS Token for setting Intel (R) VT configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vt_for_directed_io": {
				Description: "BIOS Token for setting Intel VT for directed IO configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_coherency_support": {
				Description: "BIOS Token for setting Intel (R) VT-d Coherency Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_interrupt_remapping": {
				Description: "BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_pass_through_dma_support": {
				Description: "BIOS Token for setting Intel (R) VT-d PassThrough DMA support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtdats_support": {
				Description: "BIOS Token for setting Intel VTD ATS support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ioh_error_enable": {
				Description: "BIOS Token for setting IIO Error Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `No` - Value - No for configuring IohErrorEnable token.\n* `Yes` - Value - Yes for configuring IohErrorEnable token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ioh_resource": {
				Description: "BIOS Token for setting IOH Resource Allocation configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `IOH0 24k IOH1 40k` - Value - IOH0 24k IOH1 40k for configuring IohResource token.\n* `IOH0 32k IOH1 32k` - Value - IOH0 32k IOH1 32k for configuring IohResource token.\n* `IOH0 40k IOH1 24k` - Value - IOH0 40k IOH1 24k for configuring IohResource token.\n* `IOH0 48k IOH1 16k` - Value - IOH0 48k IOH1 16k for configuring IohResource token.\n* `IOH0 56k IOH1 8k` - Value - IOH0 56k IOH1 8k for configuring IohResource token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_prefetch": {
				Description: "BIOS Token for setting DCU IP Prefetcher configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv4pxe": {
				Description: "BIOS Token for setting IPv4 PXE Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6pxe": {
				Description: "BIOS Token for setting IPV6 PXE Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"kti_prefetch": {
				Description: "BIOS Token for setting KTI Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"legacy_os_redirection": {
				Description: "BIOS Token for setting Legacy OS redirection configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"legacy_usb_support": {
				Description: "BIOS Token for setting Legacy USB Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring LegacyUsbSupport token.\n* `disabled` - Value - disabled for configuring LegacyUsbSupport token.\n* `enabled` - Value - enabled for configuring LegacyUsbSupport token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"llc_prefetch": {
				Description: "BIOS Token for setting LLC Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port0state": {
				Description: "BIOS Token for setting LOM Port 0 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring LomPort0state token.\n* `enabled` - Value - enabled for configuring LomPort0state token.\n* `Legacy Only` - Value - Legacy Only for configuring LomPort0state token.\n* `UEFI Only` - Value - UEFI Only for configuring LomPort0state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port1state": {
				Description: "BIOS Token for setting LOM Port 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring LomPort1state token.\n* `enabled` - Value - enabled for configuring LomPort1state token.\n* `Legacy Only` - Value - Legacy Only for configuring LomPort1state token.\n* `UEFI Only` - Value - UEFI Only for configuring LomPort1state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port2state": {
				Description: "BIOS Token for setting LOM Port 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring LomPort2state token.\n* `enabled` - Value - enabled for configuring LomPort2state token.\n* `Legacy Only` - Value - Legacy Only for configuring LomPort2state token.\n* `UEFI Only` - Value - UEFI Only for configuring LomPort2state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port3state": {
				Description: "BIOS Token for setting LOM Port 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring LomPort3state token.\n* `enabled` - Value - enabled for configuring LomPort3state token.\n* `Legacy Only` - Value - Legacy Only for configuring LomPort3state token.\n* `UEFI Only` - Value - UEFI Only for configuring LomPort3state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_ports_all_state": {
				Description: "BIOS Token for setting All Onboard LOM Ports configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lv_ddr_mode": {
				Description: "BIOS Token for setting Low Voltage DDR Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring LvDdrMode token.\n* `performance-mode` - Value - performance-mode for configuring LvDdrMode token.\n* `power-saving-mode` - Value - power-saving-mode for configuring LvDdrMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"make_device_non_bootable": {
				Description: "BIOS Token for setting Make Device Non Bootable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_inter_leave": {
				Description: "BIOS Token for setting Memory Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1 Way Node Interleave` - Value - 1 Way Node Interleave for configuring MemoryInterLeave token.\n* `2 Way Node Interleave` - Value - 2 Way Node Interleave for configuring MemoryInterLeave token.\n* `4 Way Node Interleave` - Value - 4 Way Node Interleave for configuring MemoryInterLeave token.\n* `8 Way Node Interleave` - Value - 8 Way Node Interleave for configuring MemoryInterLeave token.\n* `disabled` - Value - disabled for configuring MemoryInterLeave token.\n* `enabled` - Value - enabled for configuring MemoryInterLeave token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_mapped_io_above4gb": {
				Description: "BIOS Token for setting Memory mapped IO above 4GiB configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_size_limit": {
				Description: "BIOS Token for setting Memory Size Limit in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mirroring_mode": {
				Description: "BIOS Token for setting Mirroring Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `inter-socket` - Value - inter-socket for configuring MirroringMode token.\n* `intra-socket` - Value - intra-socket for configuring MirroringMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mmcfg_base": {
				Description: "BIOS Token for setting MMCFG BASE configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1 GB` - Value - 1 GiB for configuring MmcfgBase token.\n* `2 GB` - Value - 2 GiB for configuring MmcfgBase token.\n* `2.5 GB` - Value - 2.5 GiB for configuring MmcfgBase token.\n* `3 GB` - Value - 3 GiB for configuring MmcfgBase token.\n* `Auto` - Value - Auto for configuring MmcfgBase token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_stack": {
				Description: "BIOS Token for setting Network Stack configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"numa_optimized": {
				Description: "BIOS Token for setting NUMA optimized configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nvmdimm_perform_config": {
				Description: "BIOS Token for setting NVM Performance Setting configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `BW Optimized` - Value - BW Optimized for configuring NvmdimmPerformConfig token.\n* `Balanced Profile` - Value - Balanced Profile for configuring NvmdimmPerformConfig token.\n* `Latency Optimized` - Value - Latency Optimized for configuring NvmdimmPerformConfig token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"onboard10gbit_lom": {
				Description: "BIOS Token for setting Onboard 10Gbit LOM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_gbit_lom": {
				Description: "BIOS Token for setting Onboard Gbit LOM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_scu_storage_support": {
				Description: "BIOS Token for setting Onboard SCU Storage Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_scu_storage_sw_stack": {
				Description: "BIOS Token for setting Onboard SCU Storage SW Stack configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Intel RSTe` - Value - Intel RSTe for configuring OnboardScuStorageSwStack token.\n* `LSI SW RAID` - Value - LSI SW RAID for configuring OnboardScuStorageSwStack token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"os_boot_watchdog_timer": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_boot_watchdog_timer_policy": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Policy configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `do-nothing` - Value - do-nothing for configuring OsBootWatchdogTimerPolicy token.\n* `power-off` - Value - power-off for configuring OsBootWatchdogTimerPolicy token.\n* `reset` - Value - reset for configuring OsBootWatchdogTimerPolicy token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_boot_watchdog_timer_timeout": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Timeout configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `10-minutes` - Value - 10-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `15-minutes` - Value - 15-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `20-minutes` - Value - 20-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `5-minutes` - Value - 5-minutes for configuring OsBootWatchdogTimerTimeout token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_mgmt_port": {
				Description: "BIOS Token for setting Out-of-Band Mgmt Port configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"package_cstate_limit": {
				Description: "BIOS Token for setting Package C State Limit configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PackageCstateLimit token.\n* `C0 C1 State` - Value - C0 C1 State for configuring PackageCstateLimit token.\n* `C0/C1` - Value - C0/C1 for configuring PackageCstateLimit token.\n* `C2` - Value - C2 for configuring PackageCstateLimit token.\n* `C6 Non Retention` - Value - C6 Non Retention for configuring PackageCstateLimit token.\n* `C6 Retention` - Value - C6 Retention for configuring PackageCstateLimit token.\n* `No Limit` - Value - No Limit for configuring PackageCstateLimit token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_mode_config": {
				Description: "BIOS Token for setting Partial Memory Mirror Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring PartialMirrorModeConfig token.\n* `Percentage` - Value - Percentage for configuring PartialMirrorModeConfig token.\n* `Value in GB` - Value - Value in GiB for configuring PartialMirrorModeConfig token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_percent": {
				Description: "BIOS Token for setting Partial Mirror percentage configuration (0.00-50.00).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value1": {
				Description: "BIOS Token for setting Partial Mirror1 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value2": {
				Description: "BIOS Token for setting Partial Mirror2 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value3": {
				Description: "BIOS Token for setting Partial Mirror3 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value4": {
				Description: "BIOS Token for setting Partial Mirror4 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"patrol_scrub": {
				Description: "BIOS Token for setting Patrol Scrub configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"patrol_scrub_duration": {
				Description: "BIOS Token for setting Patrol Scrub Interval configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pc_ie_ras_support": {
				Description: "BIOS Token for setting PCIe RAS Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pc_ie_ssd_hot_plug_support": {
				Description: "BIOS Token for setting NVMe SSD Hot-Plug Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pch_usb30mode": {
				Description: "BIOS Token for setting xHCI Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pci_option_ro_ms": {
				Description: "BIOS Token for setting All PCIe Slots OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring PciOptionRoMs token.\n* `enabled` - Value - enabled for configuring PciOptionRoMs token.\n* `Legacy Only` - Value - Legacy Only for configuring PciOptionRoMs token.\n* `UEFI Only` - Value - UEFI Only for configuring PciOptionRoMs token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pci_rom_clp": {
				Description: "BIOS Token for setting PCI ROM CLP configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_ari_support": {
				Description: "BIOS Token for setting PCIe ARI Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieAriSupport token.\n* `disabled` - Value - disabled for configuring PcieAriSupport token.\n* `enabled` - Value - enabled for configuring PcieAriSupport token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_pll_ssc": {
				Description: "BIOS Token for setting PCIe PLL SSC configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PciePllSsc token.\n* `Disabled` - Value - Disabled for configuring PciePllSsc token.\n* `ZeroPointFive` - Value - ZeroPointFive for configuring PciePllSsc token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_mstorraid_option_rom": {
				Description: "BIOS Token for setting PCIe Slot MSTOR RAID OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme1link_speed": {
				Description: "BIOS Token for setting NVME-1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme1option_rom": {
				Description: "BIOS Token for setting NVME-1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme2link_speed": {
				Description: "BIOS Token for setting NVME-2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme2option_rom": {
				Description: "BIOS Token for setting NVME-2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme3link_speed": {
				Description: "BIOS Token for setting NVME-3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme3linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme3option_rom": {
				Description: "BIOS Token for setting NVME-3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme4link_speed": {
				Description: "BIOS Token for setting NVME-4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme4linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme4option_rom": {
				Description: "BIOS Token for setting NVME-4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme5link_speed": {
				Description: "BIOS Token for setting NVME-5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme5linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme5option_rom": {
				Description: "BIOS Token for setting NVME-5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme6link_speed": {
				Description: "BIOS Token for setting NVME-6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme6linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pcie_slot_nvme6option_rom": {
				Description: "BIOS Token for setting NVME-6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pop_support": {
				Description: "BIOS Token for setting Power ON Password configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"post_error_pause": {
				Description: "BIOS Token for setting POST Error Pause configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c1e": {
				Description: "BIOS Token for setting Processor C1E configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c3report": {
				Description: "BIOS Token for setting Processor C3 Report configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c6report": {
				Description: "BIOS Token for setting Processor C6 Report configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_cstate": {
				Description: "BIOS Token for setting CPU C State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"psata": {
				Description: "BIOS Token for setting P-SATA mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `AHCI` - Value - AHCI for configuring Psata token.\n* `Disabled` - Value - Disabled for configuring Psata token.\n* `LSI SW RAID` - Value - LSI SW RAID for configuring Psata token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pstate_coord_type": {
				Description: "BIOS Token for setting P-STATE Coordination configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `HW ALL` - Value - HW ALL for configuring PstateCoordType token.\n* `SW ALL` - Value - SW ALL for configuring PstateCoordType token.\n* `SW ANY` - Value - SW ANY for configuring PstateCoordType token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"putty_key_pad": {
				Description: "BIOS Token for setting Putty KeyPad configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `ESCN` - Value - ESCN for configuring PuttyKeyPad token.\n* `LINUX` - Value - LINUX for configuring PuttyKeyPad token.\n* `SCO` - Value - SCO for configuring PuttyKeyPad token.\n* `VT100` - Value - VT100 for configuring PuttyKeyPad token.\n* `VT400` - Value - VT400 for configuring PuttyKeyPad token.\n* `XTERMR6` - Value - XTERMR6 for configuring PuttyKeyPad token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pwr_perf_tuning": {
				Description: "BIOS Token for setting Power Performance Tuning configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `bios` - Value - BIOS for configuring PwrPerfTuning token.\n* `os` - Value - os for configuring PwrPerfTuning token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"qpi_link_frequency": {
				Description: "BIOS Token for setting QPI Link Frequency Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `6.4-gt/s` - Value - 6.4-gt/s for configuring QpiLinkFrequency token.\n* `7.2-gt/s` - Value - 7.2-gt/s for configuring QpiLinkFrequency token.\n* `8.0-gt/s` - Value - 8.0-gt/s for configuring QpiLinkFrequency token.\n* `9.6-gt/s` - Value - 9.6-gt/s for configuring QpiLinkFrequency token.\n* `auto` - Value - auto for configuring QpiLinkFrequency token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"qpi_link_speed": {
				Description: "BIOS Token for setting UPI Link Frequency Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `10.4GT/s` - Value - 10.4GT/s for configuring QpiLinkSpeed token.\n* `9.6GT/s` - Value - 9.6GT/s for configuring QpiLinkSpeed token.\n* `Auto` - Value - Auto for configuring QpiLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"qpi_snoop_mode": {
				Description: "BIOS Token for setting QPI Snoop Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring QpiSnoopMode token.\n* `cluster-on-die` - Value - cluster-on-die for configuring QpiSnoopMode token.\n* `early-snoop` - Value - early-snoop for configuring QpiSnoopMode token.\n* `home-directory-snoop` - Value - home-directory-snoop for configuring QpiSnoopMode token.\n* `home-directory-snoop-with-osb` - Value - home-directory-snoop-with-osb for configuring QpiSnoopMode token.\n* `home-snoop` - Value - home-snoop for configuring QpiSnoopMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"rank_inter_leave": {
				Description: "BIOS Token for setting Rank Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1-way` - Value - 1-way for configuring RankInterLeave token.\n* `2-way` - Value - 2-way for configuring RankInterLeave token.\n* `4-way` - Value - 4-way for configuring RankInterLeave token.\n* `8-way` - Value - 8-way for configuring RankInterLeave token.\n* `auto` - Value - auto for configuring RankInterLeave token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"redirection_after_post": {
				Description: "BIOS Token for setting Redirection After BIOS POST configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Always Enable` - Value - Always Enable for configuring RedirectionAfterPost token.\n* `Bootloader` - Value - Bootloader for configuring RedirectionAfterPost token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sata_mode_select": {
				Description: "BIOS Token for setting SATA mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `AHCI` - Value - AHCI for configuring SataModeSelect token.\n* `Disabled` - Value - Disabled for configuring SataModeSelect token.\n* `LSI SW RAID` - Value - LSI SW RAID for configuring SataModeSelect token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"select_memory_ras_configuration": {
				Description: "BIOS Token for setting SelectMemory RAS configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `adddc-sparing` - Value - adddc-sparing for configuring SelectMemoryRasConfiguration token.\n* `lockstep` - Value - lockstep for configuring SelectMemoryRasConfiguration token.\n* `maximum-performance` - Value - maximum-performance for configuring SelectMemoryRasConfiguration token.\n* `mirror-mode-1lm` - Value - mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.\n* `mirroring` - Value - mirroring for configuring SelectMemoryRasConfiguration token.\n* `partial-mirror-mode-1lm` - Value - partial-mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.\n* `sparing` - Value - sparing for configuring SelectMemoryRasConfiguration token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"select_ppr_type": {
				Description: "BIOS Token for setting Select PPR Type configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SelectPprType token.\n* `Hard PPR` - Value - Hard PPR for configuring SelectPprType token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial_port_aenable": {
				Description: "BIOS Token for setting Serial A Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"single_pctl_enable": {
				Description: "BIOS Token for setting Single PCTL configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `No` - Value - No for configuring SinglePctlEnable token.\n* `Yes` - Value - Yes for configuring SinglePctlEnable token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot10link_speed": {
				Description: "BIOS Token for setting PCIe Slot:10 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot10linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot10linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot10linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot10linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot10linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot10state": {
				Description: "BIOS Token for setting Slot 10 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot10state token.\n* `enabled` - Value - enabled for configuring Slot10state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot10state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot10state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot11link_speed": {
				Description: "BIOS Token for setting PCIe Slot:11 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot11linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot11linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot11linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot11linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot11linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot11state": {
				Description: "BIOS Token for setting Slot 11 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot12link_speed": {
				Description: "BIOS Token for setting PCIe Slot:12 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot12linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot12linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot12linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot12linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot12linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot12state": {
				Description: "BIOS Token for setting Slot 12 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot13state": {
				Description: "BIOS Token for setting PCIe Slot 13 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot14state": {
				Description: "BIOS Token for setting PCIe Slot 14 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot1state": {
				Description: "BIOS Token for setting Slot 1 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot1state token.\n* `enabled` - Value - enabled for configuring Slot1state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot1state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot1state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot2state": {
				Description: "BIOS Token for setting Slot 2 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot2state token.\n* `enabled` - Value - enabled for configuring Slot2state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot2state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot2state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot3linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot3state": {
				Description: "BIOS Token for setting Slot 3 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot3state token.\n* `enabled` - Value - enabled for configuring Slot3state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot3state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot3state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot4linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot4state": {
				Description: "BIOS Token for setting Slot 4 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot4state token.\n* `enabled` - Value - enabled for configuring Slot4state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot4state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot4state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot5linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot5state": {
				Description: "BIOS Token for setting Slot 5 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot5state token.\n* `enabled` - Value - enabled for configuring Slot5state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot5state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot5state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot6linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot6state": {
				Description: "BIOS Token for setting Slot 6 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot6state token.\n* `enabled` - Value - enabled for configuring Slot6state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot6state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot6state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot7link_speed": {
				Description: "BIOS Token for setting PCIe Slot:7 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot7linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot7linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot7linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot7linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot7linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot7state": {
				Description: "BIOS Token for setting Slot 7 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot7state token.\n* `enabled` - Value - enabled for configuring Slot7state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot7state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot7state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot8link_speed": {
				Description: "BIOS Token for setting PCIe Slot:8 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot8linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot8linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot8linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot8linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot8linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot8state": {
				Description: "BIOS Token for setting Slot 8 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot8state token.\n* `enabled` - Value - enabled for configuring Slot8state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot8state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot8state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot9link_speed": {
				Description: "BIOS Token for setting PCIe Slot:9 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot9linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot9linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot9linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot9linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot9linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot9state": {
				Description: "BIOS Token for setting Slot 9 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot9state token.\n* `enabled` - Value - enabled for configuring Slot9state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot9state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot9state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_flom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FLOM Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFlomLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFlomLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFlomLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFlomLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFlomLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontSlot5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontSlot5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontSlot5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontSlot5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontSlot5linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontSlot6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontSlot6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontSlot6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontSlot6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontSlot6linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu1state": {
				Description: "BIOS Token for setting GPU1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu2state": {
				Description: "BIOS Token for setting GPU2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu3state": {
				Description: "BIOS Token for setting GPU3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu4state": {
				Description: "BIOS Token for setting GPU4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu5state": {
				Description: "BIOS Token for setting GPU5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu6state": {
				Description: "BIOS Token for setting GPU6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu7state": {
				Description: "BIOS Token for setting GPU7 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu8state": {
				Description: "BIOS Token for setting GPU8 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_hba_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:HBA Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotHbaLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotHbaLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotHbaLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotHbaLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotHbaLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_hba_state": {
				Description: "BIOS Token for setting PCIe Slot:HBA OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotHbaState token.\n* `enabled` - Value - enabled for configuring SlotHbaState token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotHbaState token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotHbaState token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_lom1link": {
				Description: "BIOS Token for setting PCIe LOM:1 Link configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_lom2link": {
				Description: "BIOS Token for setting PCIe LOM:2 Link configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mezz_state": {
				Description: "BIOS Token for setting Slot Mezz State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotMezzState token.\n* `enabled` - Value - enabled for configuring SlotMezzState token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotMezzState token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotMezzState token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mlom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:MLOM Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotMlomLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotMlomLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotMlomLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotMlomLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotMlomLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mlom_state": {
				Description: "BIOS Token for setting PCIe Slot MLOM OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotMlomState token.\n* `enabled` - Value - enabled for configuring SlotMlomState token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotMlomState token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotMlomState token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mraid_link_speed": {
				Description: "BIOS Token for setting MRAID Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotMraidLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotMraidLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotMraidLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotMraidLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotMraidLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mraid_state": {
				Description: "BIOS Token for setting PCIe Slot MRAID OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n10state": {
				Description: "BIOS Token for setting PCIe Slot N10 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n11state": {
				Description: "BIOS Token for setting PCIe Slot N11 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n12state": {
				Description: "BIOS Token for setting PCIe Slot N12 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n13state": {
				Description: "BIOS Token for setting PCIe Slot N13 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n14state": {
				Description: "BIOS Token for setting PCIe Slot N14 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n15state": {
				Description: "BIOS Token for setting PCIe Slot N15 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n16state": {
				Description: "BIOS Token for setting PCIe Slot N16 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n17state": {
				Description: "BIOS Token for setting PCIe Slot N17 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n18state": {
				Description: "BIOS Token for setting PCIe Slot N18 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n19state": {
				Description: "BIOS Token for setting PCIe Slot N19 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n1state": {
				Description: "BIOS Token for setting PCIe Slot N1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotN1state token.\n* `enabled` - Value - enabled for configuring SlotN1state token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotN1state token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotN1state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n20state": {
				Description: "BIOS Token for setting PCIe Slot N20 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n21state": {
				Description: "BIOS Token for setting PCIe Slot N21 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n22state": {
				Description: "BIOS Token for setting PCIe Slot N22 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n23state": {
				Description: "BIOS Token for setting PCIe Slot N23 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n24state": {
				Description: "BIOS Token for setting PCIe Slot N24 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n2state": {
				Description: "BIOS Token for setting PCIe Slot N2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotN2state token.\n* `enabled` - Value - enabled for configuring SlotN2state token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotN2state token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotN2state token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n3state": {
				Description: "BIOS Token for setting PCIe Slot N3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n4state": {
				Description: "BIOS Token for setting PCIe Slot N4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n5state": {
				Description: "BIOS Token for setting PCIe Slot N5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n6state": {
				Description: "BIOS Token for setting PCIe Slot N6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n7state": {
				Description: "BIOS Token for setting PCIe Slot N7 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n8state": {
				Description: "BIOS Token for setting PCIe Slot N8 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n9state": {
				Description: "BIOS Token for setting PCIe Slot N9 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_raid_link_speed": {
				Description: "BIOS Token for setting RAID Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRaidLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRaidLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRaidLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRaidLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRaidLinkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_raid_state": {
				Description: "BIOS Token for setting PCIe Slot RAID OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme1state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme2state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme3state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme4state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme5state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme6state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme7state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 7 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme8state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 8 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser1slot1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser1slot1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser1slot1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser1slot1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser1slot1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser1slot2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser1slot2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser1slot2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser1slot2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser1slot2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser1slot3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser1slot3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser1slot3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser1slot3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser1slot3linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser2slot4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser2slot4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser2slot4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser2slot4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser2slot4linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser2slot5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser2slot5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser2slot5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser2slot5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser2slot5linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRiser2slot6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRiser2slot6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRiser2slot6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRiser2slot6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRiser2slot6linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_sas_state": {
				Description: "BIOS Token for setting PCIe Slot:SAS OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotSasState token.\n* `enabled` - Value - enabled for configuring SlotSasState token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotSasState token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotSasState token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_ssd_slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotSsdSlot1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotSsdSlot1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotSsdSlot1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotSsdSlot1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotSsdSlot1linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_ssd_slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotSsdSlot2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotSsdSlot2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotSsdSlot2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotSsdSlot2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotSsdSlot2linkSpeed token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"smee": {
				Description: "BIOS Token for setting SMEE configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"smt_mode": {
				Description: "BIOS Token for setting SMT Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SmtMode token.\n* `Off` - Value - Off for configuring SmtMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snc": {
				Description: "BIOS Token for setting Sub Numa Clustering configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Snc token.\n* `disabled` - Value - disabled for configuring Snc token.\n* `enabled` - Value - enabled for configuring Snc token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snoopy_mode_for2lm": {
				Description: "BIOS Token for setting Snoopy Mode for 2LM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snoopy_mode_for_ad": {
				Description: "BIOS Token for setting Snoopy Mode for AD configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sparing_mode": {
				Description: "BIOS Token for setting Sparing Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `dimm-sparing` - Value - dimm-sparing for configuring SparingMode token.\n* `rank-sparing` - Value - rank-sparing for configuring SparingMode token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sr_iov": {
				Description: "BIOS Token for setting SR-IOV Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"streamer_prefetch": {
				Description: "BIOS Token for setting DCU Streamer Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"svm_mode": {
				Description: "BIOS Token for setting SVM Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"terminal_type": {
				Description: "BIOS Token for setting Terminal Type configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `pc-ansi` - Value - pc-ansi for configuring TerminalType token.\n* `vt-utf8` - Value - vt-utf8 for configuring TerminalType token.\n* `vt100` - Value - vt100 for configuring TerminalType token.\n* `vt100-plus` - Value - vt100-plus for configuring TerminalType token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tpm_control": {
				Description: "BIOS Token for setting Trusted Platform Module State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tpm_support": {
				Description: "BIOS Token for setting TPM Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"txt_support": {
				Description: "BIOS Token for setting Intel Trusted Execution Technology Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ucsm_boot_order_rule": {
				Description: "BIOS Token for setting Boot Order Rules configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Loose` - Value - Loose for configuring UcsmBootOrderRule token.\n* `Strict` - Value - Strict for configuring UcsmBootOrderRule token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ufs_disable": {
				Description: "BIOS Token for setting Uncore Frequency Scaling configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_emul6064": {
				Description: "BIOS Token for setting Port 60/64 Emulation configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_front": {
				Description: "BIOS Token for setting USB Port Front configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_internal": {
				Description: "BIOS Token for setting USB Port Internal configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_kvm": {
				Description: "BIOS Token for setting USB Port KVM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_rear": {
				Description: "BIOS Token for setting USB Port Rear configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_sd_card": {
				Description: "BIOS Token for setting USB Port SD Card configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_vmedia": {
				Description: "BIOS Token for setting USB Port VMedia configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_xhci_support": {
				Description: "BIOS Token for setting XHCI Legacy Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vga_priority": {
				Description: "BIOS Token for setting VGA Priority configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Offboard` - Value - Offboard for configuring VgaPriority token.\n* `Onboard` - Value - Onboard for configuring VgaPriority token.\n* `Onboard VGA Disabled` - Value - Onboard VGA Disabled for configuring VgaPriority token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vmd_enable": {
				Description: "BIOS Token for setting VMD Enablement configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"work_load_config": {
				Description: "BIOS Token for setting Workload Configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Balanced` - Value - Balanced for configuring WorkLoadConfig token.\n* `I/O Sensitive` - Value - I/O Sensitive for configuring WorkLoadConfig token.\n* `NUMA` - Value - NUMA for configuring WorkLoadConfig token.\n* `UMA` - Value - UMA for configuring WorkLoadConfig token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"xpt_prefetch": {
				Description: "BIOS Token for setting XPT Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring XptPrefetch token.\n* `disabled` - Value - disabled for configuring XptPrefetch token.\n* `enabled` - Value - enabled for configuring XptPrefetch token.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceBiosPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewBiosPolicyWithDefaults()
	if v, ok := d.GetOk("acs_control_gpu1state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu1state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu2state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu2state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu3state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu3state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu4state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu4state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu5state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu5state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu6state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu6state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu7state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu7state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu8state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu8state(x)
	}
	if v, ok := d.GetOk("acs_control_slot11state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot11state(x)
	}
	if v, ok := d.GetOk("acs_control_slot12state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot12state(x)
	}
	if v, ok := d.GetOk("acs_control_slot13state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot13state(x)
	}
	if v, ok := d.GetOk("acs_control_slot14state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot14state(x)
	}
	if v, ok := d.GetOk("adjacent_cache_line_prefetch"); ok {
		x := (v.(string))
		o.SetAdjacentCacheLinePrefetch(x)
	}
	if v, ok := d.GetOk("all_usb_devices"); ok {
		x := (v.(string))
		o.SetAllUsbDevices(x)
	}
	if v, ok := d.GetOk("altitude"); ok {
		x := (v.(string))
		o.SetAltitude(x)
	}
	if v, ok := d.GetOk("aspm_support"); ok {
		x := (v.(string))
		o.SetAspmSupport(x)
	}
	if v, ok := d.GetOk("assert_nmi_on_perr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnPerr(x)
	}
	if v, ok := d.GetOk("assert_nmi_on_serr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnSerr(x)
	}
	if v, ok := d.GetOk("auto_cc_state"); ok {
		x := (v.(string))
		o.SetAutoCcState(x)
	}
	if v, ok := d.GetOk("autonumous_cstate_enable"); ok {
		x := (v.(string))
		o.SetAutonumousCstateEnable(x)
	}
	if v, ok := d.GetOk("baud_rate"); ok {
		x := (v.(string))
		o.SetBaudRate(x)
	}
	if v, ok := d.GetOk("bme_dma_mitigation"); ok {
		x := (v.(string))
		o.SetBmeDmaMitigation(x)
	}
	if v, ok := d.GetOk("boot_option_num_retry"); ok {
		x := (v.(string))
		o.SetBootOptionNumRetry(x)
	}
	if v, ok := d.GetOk("boot_option_re_cool_down"); ok {
		x := (v.(string))
		o.SetBootOptionReCoolDown(x)
	}
	if v, ok := d.GetOk("boot_option_retry"); ok {
		x := (v.(string))
		o.SetBootOptionRetry(x)
	}
	if v, ok := d.GetOk("boot_performance_mode"); ok {
		x := (v.(string))
		o.SetBootPerformanceMode(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_cpb"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuCpb(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_gen_downcore_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGenDowncoreCtrl(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_global_cstate_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGlobalCstateCtrl(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_l1stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL1streamHwPrefetcher(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_l2stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL2streamHwPrefetcher(x)
	}
	if v, ok := d.GetOk("cbs_cmn_determinism_slider"); ok {
		x := (v.(string))
		o.SetCbsCmnDeterminismSlider(x)
	}
	if v, ok := d.GetOk("cbs_cmn_gnb_nb_iommu"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbNbIommu(x)
	}
	if v, ok := d.GetOk("cbs_cmn_mem_ctrl_bank_group_swap_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemCtrlBankGroupSwapDdr4(x)
	}
	if v, ok := d.GetOk("cbs_cmn_mem_map_bank_interleave_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemMapBankInterleaveDdr4(x)
	}
	if v, ok := d.GetOk("cbs_cmnc_tdp_ctl"); ok {
		x := (v.(string))
		o.SetCbsCmncTdpCtl(x)
	}
	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlv(x)
	}
	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv_size"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvSize(x)
	}
	if v, ok := d.GetOk("cdn_enable"); ok {
		x := (v.(string))
		o.SetCdnEnable(x)
	}
	if v, ok := d.GetOk("cdn_support"); ok {
		x := (v.(string))
		o.SetCdnSupport(x)
	}
	if v, ok := d.GetOk("channel_inter_leave"); ok {
		x := (v.(string))
		o.SetChannelInterLeave(x)
	}
	if v, ok := d.GetOk("cisco_adaptive_mem_training"); ok {
		x := (v.(string))
		o.SetCiscoAdaptiveMemTraining(x)
	}
	if v, ok := d.GetOk("cisco_debug_level"); ok {
		x := (v.(string))
		o.SetCiscoDebugLevel(x)
	}
	if v, ok := d.GetOk("cisco_oprom_launch_optimization"); ok {
		x := (v.(string))
		o.SetCiscoOpromLaunchOptimization(x)
	}
	if v, ok := d.GetOk("cke_low_policy"); ok {
		x := (v.(string))
		o.SetCkeLowPolicy(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("closed_loop_therm_throtl"); ok {
		x := (v.(string))
		o.SetClosedLoopThermThrotl(x)
	}
	if v, ok := d.GetOk("cmci_enable"); ok {
		x := (v.(string))
		o.SetCmciEnable(x)
	}
	if v, ok := d.GetOk("config_tdp"); ok {
		x := (v.(string))
		o.SetConfigTdp(x)
	}
	if v, ok := d.GetOk("config_tdp_level"); ok {
		x := (v.(string))
		o.SetConfigTdpLevel(x)
	}
	if v, ok := d.GetOk("console_redirection"); ok {
		x := (v.(string))
		o.SetConsoleRedirection(x)
	}
	if v, ok := d.GetOk("core_multi_processing"); ok {
		x := (v.(string))
		o.SetCoreMultiProcessing(x)
	}
	if v, ok := d.GetOk("cpu_energy_performance"); ok {
		x := (v.(string))
		o.SetCpuEnergyPerformance(x)
	}
	if v, ok := d.GetOk("cpu_frequency_floor"); ok {
		x := (v.(string))
		o.SetCpuFrequencyFloor(x)
	}
	if v, ok := d.GetOk("cpu_performance"); ok {
		x := (v.(string))
		o.SetCpuPerformance(x)
	}
	if v, ok := d.GetOk("cpu_power_management"); ok {
		x := (v.(string))
		o.SetCpuPowerManagement(x)
	}
	if v, ok := d.GetOk("cr_qos"); ok {
		x := (v.(string))
		o.SetCrQos(x)
	}
	if v, ok := d.GetOk("crfastgo_config"); ok {
		x := (v.(string))
		o.SetCrfastgoConfig(x)
	}
	if v, ok := d.GetOk("dcpmm_firmware_downgrade"); ok {
		x := (v.(string))
		o.SetDcpmmFirmwareDowngrade(x)
	}
	if v, ok := d.GetOk("demand_scrub"); ok {
		x := (v.(string))
		o.SetDemandScrub(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("direct_cache_access"); ok {
		x := (v.(string))
		o.SetDirectCacheAccess(x)
	}
	if v, ok := d.GetOk("dram_clock_throttling"); ok {
		x := (v.(string))
		o.SetDramClockThrottling(x)
	}
	if v, ok := d.GetOk("dram_refresh_rate"); ok {
		x := (v.(string))
		o.SetDramRefreshRate(x)
	}
	if v, ok := d.GetOk("enable_clock_spread_spec"); ok {
		x := (v.(string))
		o.SetEnableClockSpreadSpec(x)
	}
	if v, ok := d.GetOk("energy_efficient_turbo"); ok {
		x := (v.(string))
		o.SetEnergyEfficientTurbo(x)
	}
	if v, ok := d.GetOk("eng_perf_tuning"); ok {
		x := (v.(string))
		o.SetEngPerfTuning(x)
	}
	if v, ok := d.GetOk("enhanced_intel_speed_step_tech"); ok {
		x := (v.(string))
		o.SetEnhancedIntelSpeedStepTech(x)
	}
	if v, ok := d.GetOk("epp_profile"); ok {
		x := (v.(string))
		o.SetEppProfile(x)
	}
	if v, ok := d.GetOk("execute_disable_bit"); ok {
		x := (v.(string))
		o.SetExecuteDisableBit(x)
	}
	if v, ok := d.GetOk("extended_apic"); ok {
		x := (v.(string))
		o.SetExtendedApic(x)
	}
	if v, ok := d.GetOk("flow_control"); ok {
		x := (v.(string))
		o.SetFlowControl(x)
	}
	if v, ok := d.GetOk("frb2enable"); ok {
		x := (v.(string))
		o.SetFrb2enable(x)
	}
	if v, ok := d.GetOk("hardware_prefetch"); ok {
		x := (v.(string))
		o.SetHardwarePrefetch(x)
	}
	if v, ok := d.GetOk("hwpm_enable"); ok {
		x := (v.(string))
		o.SetHwpmEnable(x)
	}
	if v, ok := d.GetOk("imc_interleave"); ok {
		x := (v.(string))
		o.SetImcInterleave(x)
	}
	if v, ok := d.GetOk("intel_hyper_threading_tech"); ok {
		x := (v.(string))
		o.SetIntelHyperThreadingTech(x)
	}
	if v, ok := d.GetOk("intel_speed_select"); ok {
		x := (v.(string))
		o.SetIntelSpeedSelect(x)
	}
	if v, ok := d.GetOk("intel_turbo_boost_tech"); ok {
		x := (v.(string))
		o.SetIntelTurboBoostTech(x)
	}
	if v, ok := d.GetOk("intel_virtualization_technology"); ok {
		x := (v.(string))
		o.SetIntelVirtualizationTechnology(x)
	}
	if v, ok := d.GetOk("intel_vt_for_directed_io"); ok {
		x := (v.(string))
		o.SetIntelVtForDirectedIo(x)
	}
	if v, ok := d.GetOk("intel_vtd_coherency_support"); ok {
		x := (v.(string))
		o.SetIntelVtdCoherencySupport(x)
	}
	if v, ok := d.GetOk("intel_vtd_interrupt_remapping"); ok {
		x := (v.(string))
		o.SetIntelVtdInterruptRemapping(x)
	}
	if v, ok := d.GetOk("intel_vtd_pass_through_dma_support"); ok {
		x := (v.(string))
		o.SetIntelVtdPassThroughDmaSupport(x)
	}
	if v, ok := d.GetOk("intel_vtdats_support"); ok {
		x := (v.(string))
		o.SetIntelVtdatsSupport(x)
	}
	if v, ok := d.GetOk("ioh_error_enable"); ok {
		x := (v.(string))
		o.SetIohErrorEnable(x)
	}
	if v, ok := d.GetOk("ioh_resource"); ok {
		x := (v.(string))
		o.SetIohResource(x)
	}
	if v, ok := d.GetOk("ip_prefetch"); ok {
		x := (v.(string))
		o.SetIpPrefetch(x)
	}
	if v, ok := d.GetOk("ipv4pxe"); ok {
		x := (v.(string))
		o.SetIpv4pxe(x)
	}
	if v, ok := d.GetOk("ipv6pxe"); ok {
		x := (v.(string))
		o.SetIpv6pxe(x)
	}
	if v, ok := d.GetOk("kti_prefetch"); ok {
		x := (v.(string))
		o.SetKtiPrefetch(x)
	}
	if v, ok := d.GetOk("legacy_os_redirection"); ok {
		x := (v.(string))
		o.SetLegacyOsRedirection(x)
	}
	if v, ok := d.GetOk("legacy_usb_support"); ok {
		x := (v.(string))
		o.SetLegacyUsbSupport(x)
	}
	if v, ok := d.GetOk("llc_prefetch"); ok {
		x := (v.(string))
		o.SetLlcPrefetch(x)
	}
	if v, ok := d.GetOk("lom_port0state"); ok {
		x := (v.(string))
		o.SetLomPort0state(x)
	}
	if v, ok := d.GetOk("lom_port1state"); ok {
		x := (v.(string))
		o.SetLomPort1state(x)
	}
	if v, ok := d.GetOk("lom_port2state"); ok {
		x := (v.(string))
		o.SetLomPort2state(x)
	}
	if v, ok := d.GetOk("lom_port3state"); ok {
		x := (v.(string))
		o.SetLomPort3state(x)
	}
	if v, ok := d.GetOk("lom_ports_all_state"); ok {
		x := (v.(string))
		o.SetLomPortsAllState(x)
	}
	if v, ok := d.GetOk("lv_ddr_mode"); ok {
		x := (v.(string))
		o.SetLvDdrMode(x)
	}
	if v, ok := d.GetOk("make_device_non_bootable"); ok {
		x := (v.(string))
		o.SetMakeDeviceNonBootable(x)
	}
	if v, ok := d.GetOk("memory_inter_leave"); ok {
		x := (v.(string))
		o.SetMemoryInterLeave(x)
	}
	if v, ok := d.GetOk("memory_mapped_io_above4gb"); ok {
		x := (v.(string))
		o.SetMemoryMappedIoAbove4gb(x)
	}
	if v, ok := d.GetOk("memory_size_limit"); ok {
		x := (v.(string))
		o.SetMemorySizeLimit(x)
	}
	if v, ok := d.GetOk("mirroring_mode"); ok {
		x := (v.(string))
		o.SetMirroringMode(x)
	}
	if v, ok := d.GetOk("mmcfg_base"); ok {
		x := (v.(string))
		o.SetMmcfgBase(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("network_stack"); ok {
		x := (v.(string))
		o.SetNetworkStack(x)
	}
	if v, ok := d.GetOk("numa_optimized"); ok {
		x := (v.(string))
		o.SetNumaOptimized(x)
	}
	if v, ok := d.GetOk("nvmdimm_perform_config"); ok {
		x := (v.(string))
		o.SetNvmdimmPerformConfig(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("onboard10gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboard10gbitLom(x)
	}
	if v, ok := d.GetOk("onboard_gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboardGbitLom(x)
	}
	if v, ok := d.GetOk("onboard_scu_storage_support"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSupport(x)
	}
	if v, ok := d.GetOk("onboard_scu_storage_sw_stack"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSwStack(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimer(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer_policy"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerPolicy(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer_timeout"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerTimeout(x)
	}
	if v, ok := d.GetOk("out_of_band_mgmt_port"); ok {
		x := (v.(string))
		o.SetOutOfBandMgmtPort(x)
	}
	if v, ok := d.GetOk("package_cstate_limit"); ok {
		x := (v.(string))
		o.SetPackageCstateLimit(x)
	}
	if v, ok := d.GetOk("partial_mirror_mode_config"); ok {
		x := (v.(string))
		o.SetPartialMirrorModeConfig(x)
	}
	if v, ok := d.GetOk("partial_mirror_percent"); ok {
		x := (v.(string))
		o.SetPartialMirrorPercent(x)
	}
	if v, ok := d.GetOk("partial_mirror_value1"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue1(x)
	}
	if v, ok := d.GetOk("partial_mirror_value2"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue2(x)
	}
	if v, ok := d.GetOk("partial_mirror_value3"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue3(x)
	}
	if v, ok := d.GetOk("partial_mirror_value4"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue4(x)
	}
	if v, ok := d.GetOk("patrol_scrub"); ok {
		x := (v.(string))
		o.SetPatrolScrub(x)
	}
	if v, ok := d.GetOk("patrol_scrub_duration"); ok {
		x := (v.(string))
		o.SetPatrolScrubDuration(x)
	}
	if v, ok := d.GetOk("pc_ie_ras_support"); ok {
		x := (v.(string))
		o.SetPcIeRasSupport(x)
	}
	if v, ok := d.GetOk("pc_ie_ssd_hot_plug_support"); ok {
		x := (v.(string))
		o.SetPcIeSsdHotPlugSupport(x)
	}
	if v, ok := d.GetOk("pch_usb30mode"); ok {
		x := (v.(string))
		o.SetPchUsb30mode(x)
	}
	if v, ok := d.GetOk("pci_option_ro_ms"); ok {
		x := (v.(string))
		o.SetPciOptionRoMs(x)
	}
	if v, ok := d.GetOk("pci_rom_clp"); ok {
		x := (v.(string))
		o.SetPciRomClp(x)
	}
	if v, ok := d.GetOk("pcie_ari_support"); ok {
		x := (v.(string))
		o.SetPcieAriSupport(x)
	}
	if v, ok := d.GetOk("pcie_pll_ssc"); ok {
		x := (v.(string))
		o.SetPciePllSsc(x)
	}
	if v, ok := d.GetOk("pcie_slot_mstorraid_option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotMstorraidOptionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme1linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme1option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme1optionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme2linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme2option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme2optionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme3link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme3linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme3option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme3optionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme4link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme4linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme4option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme4optionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme5link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme5linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme5option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme5optionRom(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme6link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme6linkSpeed(x)
	}
	if v, ok := d.GetOk("pcie_slot_nvme6option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotNvme6optionRom(x)
	}
	if v, ok := d.GetOk("pop_support"); ok {
		x := (v.(string))
		o.SetPopSupport(x)
	}
	if v, ok := d.GetOk("post_error_pause"); ok {
		x := (v.(string))
		o.SetPostErrorPause(x)
	}
	if v, ok := d.GetOk("processor_c1e"); ok {
		x := (v.(string))
		o.SetProcessorC1e(x)
	}
	if v, ok := d.GetOk("processor_c3report"); ok {
		x := (v.(string))
		o.SetProcessorC3report(x)
	}
	if v, ok := d.GetOk("processor_c6report"); ok {
		x := (v.(string))
		o.SetProcessorC6report(x)
	}
	if v, ok := d.GetOk("processor_cstate"); ok {
		x := (v.(string))
		o.SetProcessorCstate(x)
	}
	if v, ok := d.GetOk("psata"); ok {
		x := (v.(string))
		o.SetPsata(x)
	}
	if v, ok := d.GetOk("pstate_coord_type"); ok {
		x := (v.(string))
		o.SetPstateCoordType(x)
	}
	if v, ok := d.GetOk("putty_key_pad"); ok {
		x := (v.(string))
		o.SetPuttyKeyPad(x)
	}
	if v, ok := d.GetOk("pwr_perf_tuning"); ok {
		x := (v.(string))
		o.SetPwrPerfTuning(x)
	}
	if v, ok := d.GetOk("qpi_link_frequency"); ok {
		x := (v.(string))
		o.SetQpiLinkFrequency(x)
	}
	if v, ok := d.GetOk("qpi_link_speed"); ok {
		x := (v.(string))
		o.SetQpiLinkSpeed(x)
	}
	if v, ok := d.GetOk("qpi_snoop_mode"); ok {
		x := (v.(string))
		o.SetQpiSnoopMode(x)
	}
	if v, ok := d.GetOk("rank_inter_leave"); ok {
		x := (v.(string))
		o.SetRankInterLeave(x)
	}
	if v, ok := d.GetOk("redirection_after_post"); ok {
		x := (v.(string))
		o.SetRedirectionAfterPost(x)
	}
	if v, ok := d.GetOk("sata_mode_select"); ok {
		x := (v.(string))
		o.SetSataModeSelect(x)
	}
	if v, ok := d.GetOk("select_memory_ras_configuration"); ok {
		x := (v.(string))
		o.SetSelectMemoryRasConfiguration(x)
	}
	if v, ok := d.GetOk("select_ppr_type"); ok {
		x := (v.(string))
		o.SetSelectPprType(x)
	}
	if v, ok := d.GetOk("serial_port_aenable"); ok {
		x := (v.(string))
		o.SetSerialPortAenable(x)
	}
	if v, ok := d.GetOk("single_pctl_enable"); ok {
		x := (v.(string))
		o.SetSinglePctlEnable(x)
	}
	if v, ok := d.GetOk("slot10link_speed"); ok {
		x := (v.(string))
		o.SetSlot10linkSpeed(x)
	}
	if v, ok := d.GetOk("slot10state"); ok {
		x := (v.(string))
		o.SetSlot10state(x)
	}
	if v, ok := d.GetOk("slot11link_speed"); ok {
		x := (v.(string))
		o.SetSlot11linkSpeed(x)
	}
	if v, ok := d.GetOk("slot11state"); ok {
		x := (v.(string))
		o.SetSlot11state(x)
	}
	if v, ok := d.GetOk("slot12link_speed"); ok {
		x := (v.(string))
		o.SetSlot12linkSpeed(x)
	}
	if v, ok := d.GetOk("slot12state"); ok {
		x := (v.(string))
		o.SetSlot12state(x)
	}
	if v, ok := d.GetOk("slot13state"); ok {
		x := (v.(string))
		o.SetSlot13state(x)
	}
	if v, ok := d.GetOk("slot14state"); ok {
		x := (v.(string))
		o.SetSlot14state(x)
	}
	if v, ok := d.GetOk("slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot1state"); ok {
		x := (v.(string))
		o.SetSlot1state(x)
	}
	if v, ok := d.GetOk("slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlot2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot2state"); ok {
		x := (v.(string))
		o.SetSlot2state(x)
	}
	if v, ok := d.GetOk("slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlot3linkSpeed(x)
	}
	if v, ok := d.GetOk("slot3state"); ok {
		x := (v.(string))
		o.SetSlot3state(x)
	}
	if v, ok := d.GetOk("slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlot4linkSpeed(x)
	}
	if v, ok := d.GetOk("slot4state"); ok {
		x := (v.(string))
		o.SetSlot4state(x)
	}
	if v, ok := d.GetOk("slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot5state"); ok {
		x := (v.(string))
		o.SetSlot5state(x)
	}
	if v, ok := d.GetOk("slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot6state"); ok {
		x := (v.(string))
		o.SetSlot6state(x)
	}
	if v, ok := d.GetOk("slot7link_speed"); ok {
		x := (v.(string))
		o.SetSlot7linkSpeed(x)
	}
	if v, ok := d.GetOk("slot7state"); ok {
		x := (v.(string))
		o.SetSlot7state(x)
	}
	if v, ok := d.GetOk("slot8link_speed"); ok {
		x := (v.(string))
		o.SetSlot8linkSpeed(x)
	}
	if v, ok := d.GetOk("slot8state"); ok {
		x := (v.(string))
		o.SetSlot8state(x)
	}
	if v, ok := d.GetOk("slot9link_speed"); ok {
		x := (v.(string))
		o.SetSlot9linkSpeed(x)
	}
	if v, ok := d.GetOk("slot9state"); ok {
		x := (v.(string))
		o.SetSlot9state(x)
	}
	if v, ok := d.GetOk("slot_flom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotFlomLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_gpu1state"); ok {
		x := (v.(string))
		o.SetSlotGpu1state(x)
	}
	if v, ok := d.GetOk("slot_gpu2state"); ok {
		x := (v.(string))
		o.SetSlotGpu2state(x)
	}
	if v, ok := d.GetOk("slot_gpu3state"); ok {
		x := (v.(string))
		o.SetSlotGpu3state(x)
	}
	if v, ok := d.GetOk("slot_gpu4state"); ok {
		x := (v.(string))
		o.SetSlotGpu4state(x)
	}
	if v, ok := d.GetOk("slot_gpu5state"); ok {
		x := (v.(string))
		o.SetSlotGpu5state(x)
	}
	if v, ok := d.GetOk("slot_gpu6state"); ok {
		x := (v.(string))
		o.SetSlotGpu6state(x)
	}
	if v, ok := d.GetOk("slot_gpu7state"); ok {
		x := (v.(string))
		o.SetSlotGpu7state(x)
	}
	if v, ok := d.GetOk("slot_gpu8state"); ok {
		x := (v.(string))
		o.SetSlotGpu8state(x)
	}
	if v, ok := d.GetOk("slot_hba_link_speed"); ok {
		x := (v.(string))
		o.SetSlotHbaLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_hba_state"); ok {
		x := (v.(string))
		o.SetSlotHbaState(x)
	}
	if v, ok := d.GetOk("slot_lom1link"); ok {
		x := (v.(string))
		o.SetSlotLom1link(x)
	}
	if v, ok := d.GetOk("slot_lom2link"); ok {
		x := (v.(string))
		o.SetSlotLom2link(x)
	}
	if v, ok := d.GetOk("slot_mezz_state"); ok {
		x := (v.(string))
		o.SetSlotMezzState(x)
	}
	if v, ok := d.GetOk("slot_mlom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMlomLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_mlom_state"); ok {
		x := (v.(string))
		o.SetSlotMlomState(x)
	}
	if v, ok := d.GetOk("slot_mraid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMraidLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_mraid_state"); ok {
		x := (v.(string))
		o.SetSlotMraidState(x)
	}
	if v, ok := d.GetOk("slot_n10state"); ok {
		x := (v.(string))
		o.SetSlotN10state(x)
	}
	if v, ok := d.GetOk("slot_n11state"); ok {
		x := (v.(string))
		o.SetSlotN11state(x)
	}
	if v, ok := d.GetOk("slot_n12state"); ok {
		x := (v.(string))
		o.SetSlotN12state(x)
	}
	if v, ok := d.GetOk("slot_n13state"); ok {
		x := (v.(string))
		o.SetSlotN13state(x)
	}
	if v, ok := d.GetOk("slot_n14state"); ok {
		x := (v.(string))
		o.SetSlotN14state(x)
	}
	if v, ok := d.GetOk("slot_n15state"); ok {
		x := (v.(string))
		o.SetSlotN15state(x)
	}
	if v, ok := d.GetOk("slot_n16state"); ok {
		x := (v.(string))
		o.SetSlotN16state(x)
	}
	if v, ok := d.GetOk("slot_n17state"); ok {
		x := (v.(string))
		o.SetSlotN17state(x)
	}
	if v, ok := d.GetOk("slot_n18state"); ok {
		x := (v.(string))
		o.SetSlotN18state(x)
	}
	if v, ok := d.GetOk("slot_n19state"); ok {
		x := (v.(string))
		o.SetSlotN19state(x)
	}
	if v, ok := d.GetOk("slot_n1state"); ok {
		x := (v.(string))
		o.SetSlotN1state(x)
	}
	if v, ok := d.GetOk("slot_n20state"); ok {
		x := (v.(string))
		o.SetSlotN20state(x)
	}
	if v, ok := d.GetOk("slot_n21state"); ok {
		x := (v.(string))
		o.SetSlotN21state(x)
	}
	if v, ok := d.GetOk("slot_n22state"); ok {
		x := (v.(string))
		o.SetSlotN22state(x)
	}
	if v, ok := d.GetOk("slot_n23state"); ok {
		x := (v.(string))
		o.SetSlotN23state(x)
	}
	if v, ok := d.GetOk("slot_n24state"); ok {
		x := (v.(string))
		o.SetSlotN24state(x)
	}
	if v, ok := d.GetOk("slot_n2state"); ok {
		x := (v.(string))
		o.SetSlotN2state(x)
	}
	if v, ok := d.GetOk("slot_n3state"); ok {
		x := (v.(string))
		o.SetSlotN3state(x)
	}
	if v, ok := d.GetOk("slot_n4state"); ok {
		x := (v.(string))
		o.SetSlotN4state(x)
	}
	if v, ok := d.GetOk("slot_n5state"); ok {
		x := (v.(string))
		o.SetSlotN5state(x)
	}
	if v, ok := d.GetOk("slot_n6state"); ok {
		x := (v.(string))
		o.SetSlotN6state(x)
	}
	if v, ok := d.GetOk("slot_n7state"); ok {
		x := (v.(string))
		o.SetSlotN7state(x)
	}
	if v, ok := d.GetOk("slot_n8state"); ok {
		x := (v.(string))
		o.SetSlotN8state(x)
	}
	if v, ok := d.GetOk("slot_n9state"); ok {
		x := (v.(string))
		o.SetSlotN9state(x)
	}
	if v, ok := d.GetOk("slot_raid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotRaidLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_raid_state"); ok {
		x := (v.(string))
		o.SetSlotRaidState(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme1state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme2state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme3state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme3state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme4state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme4state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme5state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme5state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme6state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme6state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme7state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme7state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme8state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme8state(x)
	}
	if v, ok := d.GetOk("slot_riser1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot3linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot4linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_sas_state"); ok {
		x := (v.(string))
		o.SetSlotSasState(x)
	}
	if v, ok := d.GetOk("slot_ssd_slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_ssd_slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot2linkSpeed(x)
	}
	if v, ok := d.GetOk("smee"); ok {
		x := (v.(string))
		o.SetSmee(x)
	}
	if v, ok := d.GetOk("smt_mode"); ok {
		x := (v.(string))
		o.SetSmtMode(x)
	}
	if v, ok := d.GetOk("snc"); ok {
		x := (v.(string))
		o.SetSnc(x)
	}
	if v, ok := d.GetOk("snoopy_mode_for2lm"); ok {
		x := (v.(string))
		o.SetSnoopyModeFor2lm(x)
	}
	if v, ok := d.GetOk("snoopy_mode_for_ad"); ok {
		x := (v.(string))
		o.SetSnoopyModeForAd(x)
	}
	if v, ok := d.GetOk("sparing_mode"); ok {
		x := (v.(string))
		o.SetSparingMode(x)
	}
	if v, ok := d.GetOk("sr_iov"); ok {
		x := (v.(string))
		o.SetSrIov(x)
	}
	if v, ok := d.GetOk("streamer_prefetch"); ok {
		x := (v.(string))
		o.SetStreamerPrefetch(x)
	}
	if v, ok := d.GetOk("svm_mode"); ok {
		x := (v.(string))
		o.SetSvmMode(x)
	}
	if v, ok := d.GetOk("terminal_type"); ok {
		x := (v.(string))
		o.SetTerminalType(x)
	}
	if v, ok := d.GetOk("tpm_control"); ok {
		x := (v.(string))
		o.SetTpmControl(x)
	}
	if v, ok := d.GetOk("tpm_support"); ok {
		x := (v.(string))
		o.SetTpmSupport(x)
	}
	if v, ok := d.GetOk("txt_support"); ok {
		x := (v.(string))
		o.SetTxtSupport(x)
	}
	if v, ok := d.GetOk("ucsm_boot_order_rule"); ok {
		x := (v.(string))
		o.SetUcsmBootOrderRule(x)
	}
	if v, ok := d.GetOk("ufs_disable"); ok {
		x := (v.(string))
		o.SetUfsDisable(x)
	}
	if v, ok := d.GetOk("usb_emul6064"); ok {
		x := (v.(string))
		o.SetUsbEmul6064(x)
	}
	if v, ok := d.GetOk("usb_port_front"); ok {
		x := (v.(string))
		o.SetUsbPortFront(x)
	}
	if v, ok := d.GetOk("usb_port_internal"); ok {
		x := (v.(string))
		o.SetUsbPortInternal(x)
	}
	if v, ok := d.GetOk("usb_port_kvm"); ok {
		x := (v.(string))
		o.SetUsbPortKvm(x)
	}
	if v, ok := d.GetOk("usb_port_rear"); ok {
		x := (v.(string))
		o.SetUsbPortRear(x)
	}
	if v, ok := d.GetOk("usb_port_sd_card"); ok {
		x := (v.(string))
		o.SetUsbPortSdCard(x)
	}
	if v, ok := d.GetOk("usb_port_vmedia"); ok {
		x := (v.(string))
		o.SetUsbPortVmedia(x)
	}
	if v, ok := d.GetOk("usb_xhci_support"); ok {
		x := (v.(string))
		o.SetUsbXhciSupport(x)
	}
	if v, ok := d.GetOk("vga_priority"); ok {
		x := (v.(string))
		o.SetVgaPriority(x)
	}
	if v, ok := d.GetOk("vmd_enable"); ok {
		x := (v.(string))
		o.SetVmdEnable(x)
	}
	if v, ok := d.GetOk("work_load_config"); ok {
		x := (v.(string))
		o.SetWorkLoadConfig(x)
	}
	if v, ok := d.GetOk("xpt_prefetch"); ok {
		x := (v.(string))
		o.SetXptPrefetch(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.BiosApi.GetBiosPolicyList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.BiosPolicyList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to BiosPolicy: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewBiosPolicyWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("acs_control_gpu1state", (s.GetAcsControlGpu1state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu1state: %+v", err)
			}
			if err := d.Set("acs_control_gpu2state", (s.GetAcsControlGpu2state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu2state: %+v", err)
			}
			if err := d.Set("acs_control_gpu3state", (s.GetAcsControlGpu3state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu3state: %+v", err)
			}
			if err := d.Set("acs_control_gpu4state", (s.GetAcsControlGpu4state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu4state: %+v", err)
			}
			if err := d.Set("acs_control_gpu5state", (s.GetAcsControlGpu5state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu5state: %+v", err)
			}
			if err := d.Set("acs_control_gpu6state", (s.GetAcsControlGpu6state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu6state: %+v", err)
			}
			if err := d.Set("acs_control_gpu7state", (s.GetAcsControlGpu7state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu7state: %+v", err)
			}
			if err := d.Set("acs_control_gpu8state", (s.GetAcsControlGpu8state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlGpu8state: %+v", err)
			}
			if err := d.Set("acs_control_slot11state", (s.GetAcsControlSlot11state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlSlot11state: %+v", err)
			}
			if err := d.Set("acs_control_slot12state", (s.GetAcsControlSlot12state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlSlot12state: %+v", err)
			}
			if err := d.Set("acs_control_slot13state", (s.GetAcsControlSlot13state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlSlot13state: %+v", err)
			}
			if err := d.Set("acs_control_slot14state", (s.GetAcsControlSlot14state())); err != nil {
				return fmt.Errorf("error occurred while setting property AcsControlSlot14state: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("adjacent_cache_line_prefetch", (s.GetAdjacentCacheLinePrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property AdjacentCacheLinePrefetch: %+v", err)
			}
			if err := d.Set("all_usb_devices", (s.GetAllUsbDevices())); err != nil {
				return fmt.Errorf("error occurred while setting property AllUsbDevices: %+v", err)
			}
			if err := d.Set("altitude", (s.GetAltitude())); err != nil {
				return fmt.Errorf("error occurred while setting property Altitude: %+v", err)
			}
			if err := d.Set("aspm_support", (s.GetAspmSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property AspmSupport: %+v", err)
			}
			if err := d.Set("assert_nmi_on_perr", (s.GetAssertNmiOnPerr())); err != nil {
				return fmt.Errorf("error occurred while setting property AssertNmiOnPerr: %+v", err)
			}
			if err := d.Set("assert_nmi_on_serr", (s.GetAssertNmiOnSerr())); err != nil {
				return fmt.Errorf("error occurred while setting property AssertNmiOnSerr: %+v", err)
			}
			if err := d.Set("auto_cc_state", (s.GetAutoCcState())); err != nil {
				return fmt.Errorf("error occurred while setting property AutoCcState: %+v", err)
			}
			if err := d.Set("autonumous_cstate_enable", (s.GetAutonumousCstateEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property AutonumousCstateEnable: %+v", err)
			}
			if err := d.Set("baud_rate", (s.GetBaudRate())); err != nil {
				return fmt.Errorf("error occurred while setting property BaudRate: %+v", err)
			}
			if err := d.Set("bme_dma_mitigation", (s.GetBmeDmaMitigation())); err != nil {
				return fmt.Errorf("error occurred while setting property BmeDmaMitigation: %+v", err)
			}
			if err := d.Set("boot_option_num_retry", (s.GetBootOptionNumRetry())); err != nil {
				return fmt.Errorf("error occurred while setting property BootOptionNumRetry: %+v", err)
			}
			if err := d.Set("boot_option_re_cool_down", (s.GetBootOptionReCoolDown())); err != nil {
				return fmt.Errorf("error occurred while setting property BootOptionReCoolDown: %+v", err)
			}
			if err := d.Set("boot_option_retry", (s.GetBootOptionRetry())); err != nil {
				return fmt.Errorf("error occurred while setting property BootOptionRetry: %+v", err)
			}
			if err := d.Set("boot_performance_mode", (s.GetBootPerformanceMode())); err != nil {
				return fmt.Errorf("error occurred while setting property BootPerformanceMode: %+v", err)
			}
			if err := d.Set("cbs_cmn_cpu_cpb", (s.GetCbsCmnCpuCpb())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnCpuCpb: %+v", err)
			}
			if err := d.Set("cbs_cmn_cpu_gen_downcore_ctrl", (s.GetCbsCmnCpuGenDowncoreCtrl())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnCpuGenDowncoreCtrl: %+v", err)
			}
			if err := d.Set("cbs_cmn_cpu_global_cstate_ctrl", (s.GetCbsCmnCpuGlobalCstateCtrl())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnCpuGlobalCstateCtrl: %+v", err)
			}
			if err := d.Set("cbs_cmn_cpu_l1stream_hw_prefetcher", (s.GetCbsCmnCpuL1streamHwPrefetcher())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnCpuL1streamHwPrefetcher: %+v", err)
			}
			if err := d.Set("cbs_cmn_cpu_l2stream_hw_prefetcher", (s.GetCbsCmnCpuL2streamHwPrefetcher())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnCpuL2streamHwPrefetcher: %+v", err)
			}
			if err := d.Set("cbs_cmn_determinism_slider", (s.GetCbsCmnDeterminismSlider())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnDeterminismSlider: %+v", err)
			}
			if err := d.Set("cbs_cmn_gnb_nb_iommu", (s.GetCbsCmnGnbNbIommu())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnGnbNbIommu: %+v", err)
			}
			if err := d.Set("cbs_cmn_mem_ctrl_bank_group_swap_ddr4", (s.GetCbsCmnMemCtrlBankGroupSwapDdr4())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnMemCtrlBankGroupSwapDdr4: %+v", err)
			}
			if err := d.Set("cbs_cmn_mem_map_bank_interleave_ddr4", (s.GetCbsCmnMemMapBankInterleaveDdr4())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmnMemMapBankInterleaveDdr4: %+v", err)
			}
			if err := d.Set("cbs_cmnc_tdp_ctl", (s.GetCbsCmncTdpCtl())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsCmncTdpCtl: %+v", err)
			}
			if err := d.Set("cbs_df_cmn_mem_intlv", (s.GetCbsDfCmnMemIntlv())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsDfCmnMemIntlv: %+v", err)
			}
			if err := d.Set("cbs_df_cmn_mem_intlv_size", (s.GetCbsDfCmnMemIntlvSize())); err != nil {
				return fmt.Errorf("error occurred while setting property CbsDfCmnMemIntlvSize: %+v", err)
			}
			if err := d.Set("cdn_enable", (s.GetCdnEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property CdnEnable: %+v", err)
			}
			if err := d.Set("cdn_support", (s.GetCdnSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property CdnSupport: %+v", err)
			}
			if err := d.Set("channel_inter_leave", (s.GetChannelInterLeave())); err != nil {
				return fmt.Errorf("error occurred while setting property ChannelInterLeave: %+v", err)
			}
			if err := d.Set("cisco_adaptive_mem_training", (s.GetCiscoAdaptiveMemTraining())); err != nil {
				return fmt.Errorf("error occurred while setting property CiscoAdaptiveMemTraining: %+v", err)
			}
			if err := d.Set("cisco_debug_level", (s.GetCiscoDebugLevel())); err != nil {
				return fmt.Errorf("error occurred while setting property CiscoDebugLevel: %+v", err)
			}
			if err := d.Set("cisco_oprom_launch_optimization", (s.GetCiscoOpromLaunchOptimization())); err != nil {
				return fmt.Errorf("error occurred while setting property CiscoOpromLaunchOptimization: %+v", err)
			}
			if err := d.Set("cke_low_policy", (s.GetCkeLowPolicy())); err != nil {
				return fmt.Errorf("error occurred while setting property CkeLowPolicy: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("closed_loop_therm_throtl", (s.GetClosedLoopThermThrotl())); err != nil {
				return fmt.Errorf("error occurred while setting property ClosedLoopThermThrotl: %+v", err)
			}
			if err := d.Set("cmci_enable", (s.GetCmciEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property CmciEnable: %+v", err)
			}
			if err := d.Set("config_tdp", (s.GetConfigTdp())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigTdp: %+v", err)
			}
			if err := d.Set("config_tdp_level", (s.GetConfigTdpLevel())); err != nil {
				return fmt.Errorf("error occurred while setting property ConfigTdpLevel: %+v", err)
			}
			if err := d.Set("console_redirection", (s.GetConsoleRedirection())); err != nil {
				return fmt.Errorf("error occurred while setting property ConsoleRedirection: %+v", err)
			}
			if err := d.Set("core_multi_processing", (s.GetCoreMultiProcessing())); err != nil {
				return fmt.Errorf("error occurred while setting property CoreMultiProcessing: %+v", err)
			}
			if err := d.Set("cpu_energy_performance", (s.GetCpuEnergyPerformance())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuEnergyPerformance: %+v", err)
			}
			if err := d.Set("cpu_frequency_floor", (s.GetCpuFrequencyFloor())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuFrequencyFloor: %+v", err)
			}
			if err := d.Set("cpu_performance", (s.GetCpuPerformance())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuPerformance: %+v", err)
			}
			if err := d.Set("cpu_power_management", (s.GetCpuPowerManagement())); err != nil {
				return fmt.Errorf("error occurred while setting property CpuPowerManagement: %+v", err)
			}
			if err := d.Set("cr_qos", (s.GetCrQos())); err != nil {
				return fmt.Errorf("error occurred while setting property CrQos: %+v", err)
			}
			if err := d.Set("crfastgo_config", (s.GetCrfastgoConfig())); err != nil {
				return fmt.Errorf("error occurred while setting property CrfastgoConfig: %+v", err)
			}
			if err := d.Set("dcpmm_firmware_downgrade", (s.GetDcpmmFirmwareDowngrade())); err != nil {
				return fmt.Errorf("error occurred while setting property DcpmmFirmwareDowngrade: %+v", err)
			}
			if err := d.Set("demand_scrub", (s.GetDemandScrub())); err != nil {
				return fmt.Errorf("error occurred while setting property DemandScrub: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}
			if err := d.Set("direct_cache_access", (s.GetDirectCacheAccess())); err != nil {
				return fmt.Errorf("error occurred while setting property DirectCacheAccess: %+v", err)
			}
			if err := d.Set("dram_clock_throttling", (s.GetDramClockThrottling())); err != nil {
				return fmt.Errorf("error occurred while setting property DramClockThrottling: %+v", err)
			}
			if err := d.Set("dram_refresh_rate", (s.GetDramRefreshRate())); err != nil {
				return fmt.Errorf("error occurred while setting property DramRefreshRate: %+v", err)
			}
			if err := d.Set("enable_clock_spread_spec", (s.GetEnableClockSpreadSpec())); err != nil {
				return fmt.Errorf("error occurred while setting property EnableClockSpreadSpec: %+v", err)
			}
			if err := d.Set("energy_efficient_turbo", (s.GetEnergyEfficientTurbo())); err != nil {
				return fmt.Errorf("error occurred while setting property EnergyEfficientTurbo: %+v", err)
			}
			if err := d.Set("eng_perf_tuning", (s.GetEngPerfTuning())); err != nil {
				return fmt.Errorf("error occurred while setting property EngPerfTuning: %+v", err)
			}
			if err := d.Set("enhanced_intel_speed_step_tech", (s.GetEnhancedIntelSpeedStepTech())); err != nil {
				return fmt.Errorf("error occurred while setting property EnhancedIntelSpeedStepTech: %+v", err)
			}
			if err := d.Set("epp_profile", (s.GetEppProfile())); err != nil {
				return fmt.Errorf("error occurred while setting property EppProfile: %+v", err)
			}
			if err := d.Set("execute_disable_bit", (s.GetExecuteDisableBit())); err != nil {
				return fmt.Errorf("error occurred while setting property ExecuteDisableBit: %+v", err)
			}
			if err := d.Set("extended_apic", (s.GetExtendedApic())); err != nil {
				return fmt.Errorf("error occurred while setting property ExtendedApic: %+v", err)
			}
			if err := d.Set("flow_control", (s.GetFlowControl())); err != nil {
				return fmt.Errorf("error occurred while setting property FlowControl: %+v", err)
			}
			if err := d.Set("frb2enable", (s.GetFrb2enable())); err != nil {
				return fmt.Errorf("error occurred while setting property Frb2enable: %+v", err)
			}
			if err := d.Set("hardware_prefetch", (s.GetHardwarePrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property HardwarePrefetch: %+v", err)
			}
			if err := d.Set("hwpm_enable", (s.GetHwpmEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property HwpmEnable: %+v", err)
			}
			if err := d.Set("imc_interleave", (s.GetImcInterleave())); err != nil {
				return fmt.Errorf("error occurred while setting property ImcInterleave: %+v", err)
			}
			if err := d.Set("intel_hyper_threading_tech", (s.GetIntelHyperThreadingTech())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelHyperThreadingTech: %+v", err)
			}
			if err := d.Set("intel_speed_select", (s.GetIntelSpeedSelect())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelSpeedSelect: %+v", err)
			}
			if err := d.Set("intel_turbo_boost_tech", (s.GetIntelTurboBoostTech())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelTurboBoostTech: %+v", err)
			}
			if err := d.Set("intel_virtualization_technology", (s.GetIntelVirtualizationTechnology())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVirtualizationTechnology: %+v", err)
			}
			if err := d.Set("intel_vt_for_directed_io", (s.GetIntelVtForDirectedIo())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVtForDirectedIo: %+v", err)
			}
			if err := d.Set("intel_vtd_coherency_support", (s.GetIntelVtdCoherencySupport())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVtdCoherencySupport: %+v", err)
			}
			if err := d.Set("intel_vtd_interrupt_remapping", (s.GetIntelVtdInterruptRemapping())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVtdInterruptRemapping: %+v", err)
			}
			if err := d.Set("intel_vtd_pass_through_dma_support", (s.GetIntelVtdPassThroughDmaSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVtdPassThroughDmaSupport: %+v", err)
			}
			if err := d.Set("intel_vtdats_support", (s.GetIntelVtdatsSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property IntelVtdatsSupport: %+v", err)
			}
			if err := d.Set("ioh_error_enable", (s.GetIohErrorEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property IohErrorEnable: %+v", err)
			}
			if err := d.Set("ioh_resource", (s.GetIohResource())); err != nil {
				return fmt.Errorf("error occurred while setting property IohResource: %+v", err)
			}
			if err := d.Set("ip_prefetch", (s.GetIpPrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property IpPrefetch: %+v", err)
			}
			if err := d.Set("ipv4pxe", (s.GetIpv4pxe())); err != nil {
				return fmt.Errorf("error occurred while setting property Ipv4pxe: %+v", err)
			}
			if err := d.Set("ipv6pxe", (s.GetIpv6pxe())); err != nil {
				return fmt.Errorf("error occurred while setting property Ipv6pxe: %+v", err)
			}
			if err := d.Set("kti_prefetch", (s.GetKtiPrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property KtiPrefetch: %+v", err)
			}
			if err := d.Set("legacy_os_redirection", (s.GetLegacyOsRedirection())); err != nil {
				return fmt.Errorf("error occurred while setting property LegacyOsRedirection: %+v", err)
			}
			if err := d.Set("legacy_usb_support", (s.GetLegacyUsbSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property LegacyUsbSupport: %+v", err)
			}
			if err := d.Set("llc_prefetch", (s.GetLlcPrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property LlcPrefetch: %+v", err)
			}
			if err := d.Set("lom_port0state", (s.GetLomPort0state())); err != nil {
				return fmt.Errorf("error occurred while setting property LomPort0state: %+v", err)
			}
			if err := d.Set("lom_port1state", (s.GetLomPort1state())); err != nil {
				return fmt.Errorf("error occurred while setting property LomPort1state: %+v", err)
			}
			if err := d.Set("lom_port2state", (s.GetLomPort2state())); err != nil {
				return fmt.Errorf("error occurred while setting property LomPort2state: %+v", err)
			}
			if err := d.Set("lom_port3state", (s.GetLomPort3state())); err != nil {
				return fmt.Errorf("error occurred while setting property LomPort3state: %+v", err)
			}
			if err := d.Set("lom_ports_all_state", (s.GetLomPortsAllState())); err != nil {
				return fmt.Errorf("error occurred while setting property LomPortsAllState: %+v", err)
			}
			if err := d.Set("lv_ddr_mode", (s.GetLvDdrMode())); err != nil {
				return fmt.Errorf("error occurred while setting property LvDdrMode: %+v", err)
			}
			if err := d.Set("make_device_non_bootable", (s.GetMakeDeviceNonBootable())); err != nil {
				return fmt.Errorf("error occurred while setting property MakeDeviceNonBootable: %+v", err)
			}
			if err := d.Set("memory_inter_leave", (s.GetMemoryInterLeave())); err != nil {
				return fmt.Errorf("error occurred while setting property MemoryInterLeave: %+v", err)
			}
			if err := d.Set("memory_mapped_io_above4gb", (s.GetMemoryMappedIoAbove4gb())); err != nil {
				return fmt.Errorf("error occurred while setting property MemoryMappedIoAbove4gb: %+v", err)
			}
			if err := d.Set("memory_size_limit", (s.GetMemorySizeLimit())); err != nil {
				return fmt.Errorf("error occurred while setting property MemorySizeLimit: %+v", err)
			}
			if err := d.Set("mirroring_mode", (s.GetMirroringMode())); err != nil {
				return fmt.Errorf("error occurred while setting property MirroringMode: %+v", err)
			}
			if err := d.Set("mmcfg_base", (s.GetMmcfgBase())); err != nil {
				return fmt.Errorf("error occurred while setting property MmcfgBase: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("network_stack", (s.GetNetworkStack())); err != nil {
				return fmt.Errorf("error occurred while setting property NetworkStack: %+v", err)
			}
			if err := d.Set("numa_optimized", (s.GetNumaOptimized())); err != nil {
				return fmt.Errorf("error occurred while setting property NumaOptimized: %+v", err)
			}
			if err := d.Set("nvmdimm_perform_config", (s.GetNvmdimmPerformConfig())); err != nil {
				return fmt.Errorf("error occurred while setting property NvmdimmPerformConfig: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("onboard10gbit_lom", (s.GetOnboard10gbitLom())); err != nil {
				return fmt.Errorf("error occurred while setting property Onboard10gbitLom: %+v", err)
			}
			if err := d.Set("onboard_gbit_lom", (s.GetOnboardGbitLom())); err != nil {
				return fmt.Errorf("error occurred while setting property OnboardGbitLom: %+v", err)
			}
			if err := d.Set("onboard_scu_storage_support", (s.GetOnboardScuStorageSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property OnboardScuStorageSupport: %+v", err)
			}
			if err := d.Set("onboard_scu_storage_sw_stack", (s.GetOnboardScuStorageSwStack())); err != nil {
				return fmt.Errorf("error occurred while setting property OnboardScuStorageSwStack: %+v", err)
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Organization: %+v", err)
			}
			if err := d.Set("os_boot_watchdog_timer", (s.GetOsBootWatchdogTimer())); err != nil {
				return fmt.Errorf("error occurred while setting property OsBootWatchdogTimer: %+v", err)
			}
			if err := d.Set("os_boot_watchdog_timer_policy", (s.GetOsBootWatchdogTimerPolicy())); err != nil {
				return fmt.Errorf("error occurred while setting property OsBootWatchdogTimerPolicy: %+v", err)
			}
			if err := d.Set("os_boot_watchdog_timer_timeout", (s.GetOsBootWatchdogTimerTimeout())); err != nil {
				return fmt.Errorf("error occurred while setting property OsBootWatchdogTimerTimeout: %+v", err)
			}
			if err := d.Set("out_of_band_mgmt_port", (s.GetOutOfBandMgmtPort())); err != nil {
				return fmt.Errorf("error occurred while setting property OutOfBandMgmtPort: %+v", err)
			}
			if err := d.Set("package_cstate_limit", (s.GetPackageCstateLimit())); err != nil {
				return fmt.Errorf("error occurred while setting property PackageCstateLimit: %+v", err)
			}
			if err := d.Set("partial_mirror_mode_config", (s.GetPartialMirrorModeConfig())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorModeConfig: %+v", err)
			}
			if err := d.Set("partial_mirror_percent", (s.GetPartialMirrorPercent())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorPercent: %+v", err)
			}
			if err := d.Set("partial_mirror_value1", (s.GetPartialMirrorValue1())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorValue1: %+v", err)
			}
			if err := d.Set("partial_mirror_value2", (s.GetPartialMirrorValue2())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorValue2: %+v", err)
			}
			if err := d.Set("partial_mirror_value3", (s.GetPartialMirrorValue3())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorValue3: %+v", err)
			}
			if err := d.Set("partial_mirror_value4", (s.GetPartialMirrorValue4())); err != nil {
				return fmt.Errorf("error occurred while setting property PartialMirrorValue4: %+v", err)
			}
			if err := d.Set("patrol_scrub", (s.GetPatrolScrub())); err != nil {
				return fmt.Errorf("error occurred while setting property PatrolScrub: %+v", err)
			}
			if err := d.Set("patrol_scrub_duration", (s.GetPatrolScrubDuration())); err != nil {
				return fmt.Errorf("error occurred while setting property PatrolScrubDuration: %+v", err)
			}
			if err := d.Set("pc_ie_ras_support", (s.GetPcIeRasSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property PcIeRasSupport: %+v", err)
			}
			if err := d.Set("pc_ie_ssd_hot_plug_support", (s.GetPcIeSsdHotPlugSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property PcIeSsdHotPlugSupport: %+v", err)
			}
			if err := d.Set("pch_usb30mode", (s.GetPchUsb30mode())); err != nil {
				return fmt.Errorf("error occurred while setting property PchUsb30mode: %+v", err)
			}
			if err := d.Set("pci_option_ro_ms", (s.GetPciOptionRoMs())); err != nil {
				return fmt.Errorf("error occurred while setting property PciOptionRoMs: %+v", err)
			}
			if err := d.Set("pci_rom_clp", (s.GetPciRomClp())); err != nil {
				return fmt.Errorf("error occurred while setting property PciRomClp: %+v", err)
			}
			if err := d.Set("pcie_ari_support", (s.GetPcieAriSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieAriSupport: %+v", err)
			}
			if err := d.Set("pcie_pll_ssc", (s.GetPciePllSsc())); err != nil {
				return fmt.Errorf("error occurred while setting property PciePllSsc: %+v", err)
			}
			if err := d.Set("pcie_slot_mstorraid_option_rom", (s.GetPcieSlotMstorraidOptionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotMstorraidOptionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme1link_speed", (s.GetPcieSlotNvme1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme1linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme1option_rom", (s.GetPcieSlotNvme1optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme1optionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme2link_speed", (s.GetPcieSlotNvme2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme2linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme2option_rom", (s.GetPcieSlotNvme2optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme2optionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme3link_speed", (s.GetPcieSlotNvme3linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme3linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme3option_rom", (s.GetPcieSlotNvme3optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme3optionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme4link_speed", (s.GetPcieSlotNvme4linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme4linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme4option_rom", (s.GetPcieSlotNvme4optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme4optionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme5link_speed", (s.GetPcieSlotNvme5linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme5linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme5option_rom", (s.GetPcieSlotNvme5optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme5optionRom: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme6link_speed", (s.GetPcieSlotNvme6linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme6linkSpeed: %+v", err)
			}
			if err := d.Set("pcie_slot_nvme6option_rom", (s.GetPcieSlotNvme6optionRom())); err != nil {
				return fmt.Errorf("error occurred while setting property PcieSlotNvme6optionRom: %+v", err)
			}
			if err := d.Set("pop_support", (s.GetPopSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property PopSupport: %+v", err)
			}
			if err := d.Set("post_error_pause", (s.GetPostErrorPause())); err != nil {
				return fmt.Errorf("error occurred while setting property PostErrorPause: %+v", err)
			}
			if err := d.Set("processor_c1e", (s.GetProcessorC1e())); err != nil {
				return fmt.Errorf("error occurred while setting property ProcessorC1e: %+v", err)
			}
			if err := d.Set("processor_c3report", (s.GetProcessorC3report())); err != nil {
				return fmt.Errorf("error occurred while setting property ProcessorC3report: %+v", err)
			}
			if err := d.Set("processor_c6report", (s.GetProcessorC6report())); err != nil {
				return fmt.Errorf("error occurred while setting property ProcessorC6report: %+v", err)
			}
			if err := d.Set("processor_cstate", (s.GetProcessorCstate())); err != nil {
				return fmt.Errorf("error occurred while setting property ProcessorCstate: %+v", err)
			}

			if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Profiles: %+v", err)
			}
			if err := d.Set("psata", (s.GetPsata())); err != nil {
				return fmt.Errorf("error occurred while setting property Psata: %+v", err)
			}
			if err := d.Set("pstate_coord_type", (s.GetPstateCoordType())); err != nil {
				return fmt.Errorf("error occurred while setting property PstateCoordType: %+v", err)
			}
			if err := d.Set("putty_key_pad", (s.GetPuttyKeyPad())); err != nil {
				return fmt.Errorf("error occurred while setting property PuttyKeyPad: %+v", err)
			}
			if err := d.Set("pwr_perf_tuning", (s.GetPwrPerfTuning())); err != nil {
				return fmt.Errorf("error occurred while setting property PwrPerfTuning: %+v", err)
			}
			if err := d.Set("qpi_link_frequency", (s.GetQpiLinkFrequency())); err != nil {
				return fmt.Errorf("error occurred while setting property QpiLinkFrequency: %+v", err)
			}
			if err := d.Set("qpi_link_speed", (s.GetQpiLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property QpiLinkSpeed: %+v", err)
			}
			if err := d.Set("qpi_snoop_mode", (s.GetQpiSnoopMode())); err != nil {
				return fmt.Errorf("error occurred while setting property QpiSnoopMode: %+v", err)
			}
			if err := d.Set("rank_inter_leave", (s.GetRankInterLeave())); err != nil {
				return fmt.Errorf("error occurred while setting property RankInterLeave: %+v", err)
			}
			if err := d.Set("redirection_after_post", (s.GetRedirectionAfterPost())); err != nil {
				return fmt.Errorf("error occurred while setting property RedirectionAfterPost: %+v", err)
			}
			if err := d.Set("sata_mode_select", (s.GetSataModeSelect())); err != nil {
				return fmt.Errorf("error occurred while setting property SataModeSelect: %+v", err)
			}
			if err := d.Set("select_memory_ras_configuration", (s.GetSelectMemoryRasConfiguration())); err != nil {
				return fmt.Errorf("error occurred while setting property SelectMemoryRasConfiguration: %+v", err)
			}
			if err := d.Set("select_ppr_type", (s.GetSelectPprType())); err != nil {
				return fmt.Errorf("error occurred while setting property SelectPprType: %+v", err)
			}
			if err := d.Set("serial_port_aenable", (s.GetSerialPortAenable())); err != nil {
				return fmt.Errorf("error occurred while setting property SerialPortAenable: %+v", err)
			}
			if err := d.Set("single_pctl_enable", (s.GetSinglePctlEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property SinglePctlEnable: %+v", err)
			}
			if err := d.Set("slot10link_speed", (s.GetSlot10linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot10linkSpeed: %+v", err)
			}
			if err := d.Set("slot10state", (s.GetSlot10state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot10state: %+v", err)
			}
			if err := d.Set("slot11link_speed", (s.GetSlot11linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot11linkSpeed: %+v", err)
			}
			if err := d.Set("slot11state", (s.GetSlot11state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot11state: %+v", err)
			}
			if err := d.Set("slot12link_speed", (s.GetSlot12linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot12linkSpeed: %+v", err)
			}
			if err := d.Set("slot12state", (s.GetSlot12state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot12state: %+v", err)
			}
			if err := d.Set("slot13state", (s.GetSlot13state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot13state: %+v", err)
			}
			if err := d.Set("slot14state", (s.GetSlot14state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot14state: %+v", err)
			}
			if err := d.Set("slot1link_speed", (s.GetSlot1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot1linkSpeed: %+v", err)
			}
			if err := d.Set("slot1state", (s.GetSlot1state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot1state: %+v", err)
			}
			if err := d.Set("slot2link_speed", (s.GetSlot2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot2linkSpeed: %+v", err)
			}
			if err := d.Set("slot2state", (s.GetSlot2state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot2state: %+v", err)
			}
			if err := d.Set("slot3link_speed", (s.GetSlot3linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot3linkSpeed: %+v", err)
			}
			if err := d.Set("slot3state", (s.GetSlot3state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot3state: %+v", err)
			}
			if err := d.Set("slot4link_speed", (s.GetSlot4linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot4linkSpeed: %+v", err)
			}
			if err := d.Set("slot4state", (s.GetSlot4state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot4state: %+v", err)
			}
			if err := d.Set("slot5link_speed", (s.GetSlot5linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot5linkSpeed: %+v", err)
			}
			if err := d.Set("slot5state", (s.GetSlot5state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot5state: %+v", err)
			}
			if err := d.Set("slot6link_speed", (s.GetSlot6linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot6linkSpeed: %+v", err)
			}
			if err := d.Set("slot6state", (s.GetSlot6state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot6state: %+v", err)
			}
			if err := d.Set("slot7link_speed", (s.GetSlot7linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot7linkSpeed: %+v", err)
			}
			if err := d.Set("slot7state", (s.GetSlot7state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot7state: %+v", err)
			}
			if err := d.Set("slot8link_speed", (s.GetSlot8linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot8linkSpeed: %+v", err)
			}
			if err := d.Set("slot8state", (s.GetSlot8state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot8state: %+v", err)
			}
			if err := d.Set("slot9link_speed", (s.GetSlot9linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot9linkSpeed: %+v", err)
			}
			if err := d.Set("slot9state", (s.GetSlot9state())); err != nil {
				return fmt.Errorf("error occurred while setting property Slot9state: %+v", err)
			}
			if err := d.Set("slot_flom_link_speed", (s.GetSlotFlomLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotFlomLinkSpeed: %+v", err)
			}
			if err := d.Set("slot_front_nvme1link_speed", (s.GetSlotFrontNvme1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotFrontNvme1linkSpeed: %+v", err)
			}
			if err := d.Set("slot_front_nvme2link_speed", (s.GetSlotFrontNvme2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotFrontNvme2linkSpeed: %+v", err)
			}
			if err := d.Set("slot_front_slot5link_speed", (s.GetSlotFrontSlot5linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotFrontSlot5linkSpeed: %+v", err)
			}
			if err := d.Set("slot_front_slot6link_speed", (s.GetSlotFrontSlot6linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotFrontSlot6linkSpeed: %+v", err)
			}
			if err := d.Set("slot_gpu1state", (s.GetSlotGpu1state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu1state: %+v", err)
			}
			if err := d.Set("slot_gpu2state", (s.GetSlotGpu2state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu2state: %+v", err)
			}
			if err := d.Set("slot_gpu3state", (s.GetSlotGpu3state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu3state: %+v", err)
			}
			if err := d.Set("slot_gpu4state", (s.GetSlotGpu4state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu4state: %+v", err)
			}
			if err := d.Set("slot_gpu5state", (s.GetSlotGpu5state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu5state: %+v", err)
			}
			if err := d.Set("slot_gpu6state", (s.GetSlotGpu6state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu6state: %+v", err)
			}
			if err := d.Set("slot_gpu7state", (s.GetSlotGpu7state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu7state: %+v", err)
			}
			if err := d.Set("slot_gpu8state", (s.GetSlotGpu8state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotGpu8state: %+v", err)
			}
			if err := d.Set("slot_hba_link_speed", (s.GetSlotHbaLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotHbaLinkSpeed: %+v", err)
			}
			if err := d.Set("slot_hba_state", (s.GetSlotHbaState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotHbaState: %+v", err)
			}
			if err := d.Set("slot_lom1link", (s.GetSlotLom1link())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotLom1link: %+v", err)
			}
			if err := d.Set("slot_lom2link", (s.GetSlotLom2link())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotLom2link: %+v", err)
			}
			if err := d.Set("slot_mezz_state", (s.GetSlotMezzState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotMezzState: %+v", err)
			}
			if err := d.Set("slot_mlom_link_speed", (s.GetSlotMlomLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotMlomLinkSpeed: %+v", err)
			}
			if err := d.Set("slot_mlom_state", (s.GetSlotMlomState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotMlomState: %+v", err)
			}
			if err := d.Set("slot_mraid_link_speed", (s.GetSlotMraidLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotMraidLinkSpeed: %+v", err)
			}
			if err := d.Set("slot_mraid_state", (s.GetSlotMraidState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotMraidState: %+v", err)
			}
			if err := d.Set("slot_n10state", (s.GetSlotN10state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN10state: %+v", err)
			}
			if err := d.Set("slot_n11state", (s.GetSlotN11state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN11state: %+v", err)
			}
			if err := d.Set("slot_n12state", (s.GetSlotN12state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN12state: %+v", err)
			}
			if err := d.Set("slot_n13state", (s.GetSlotN13state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN13state: %+v", err)
			}
			if err := d.Set("slot_n14state", (s.GetSlotN14state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN14state: %+v", err)
			}
			if err := d.Set("slot_n15state", (s.GetSlotN15state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN15state: %+v", err)
			}
			if err := d.Set("slot_n16state", (s.GetSlotN16state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN16state: %+v", err)
			}
			if err := d.Set("slot_n17state", (s.GetSlotN17state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN17state: %+v", err)
			}
			if err := d.Set("slot_n18state", (s.GetSlotN18state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN18state: %+v", err)
			}
			if err := d.Set("slot_n19state", (s.GetSlotN19state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN19state: %+v", err)
			}
			if err := d.Set("slot_n1state", (s.GetSlotN1state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN1state: %+v", err)
			}
			if err := d.Set("slot_n20state", (s.GetSlotN20state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN20state: %+v", err)
			}
			if err := d.Set("slot_n21state", (s.GetSlotN21state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN21state: %+v", err)
			}
			if err := d.Set("slot_n22state", (s.GetSlotN22state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN22state: %+v", err)
			}
			if err := d.Set("slot_n23state", (s.GetSlotN23state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN23state: %+v", err)
			}
			if err := d.Set("slot_n24state", (s.GetSlotN24state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN24state: %+v", err)
			}
			if err := d.Set("slot_n2state", (s.GetSlotN2state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN2state: %+v", err)
			}
			if err := d.Set("slot_n3state", (s.GetSlotN3state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN3state: %+v", err)
			}
			if err := d.Set("slot_n4state", (s.GetSlotN4state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN4state: %+v", err)
			}
			if err := d.Set("slot_n5state", (s.GetSlotN5state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN5state: %+v", err)
			}
			if err := d.Set("slot_n6state", (s.GetSlotN6state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN6state: %+v", err)
			}
			if err := d.Set("slot_n7state", (s.GetSlotN7state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN7state: %+v", err)
			}
			if err := d.Set("slot_n8state", (s.GetSlotN8state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN8state: %+v", err)
			}
			if err := d.Set("slot_n9state", (s.GetSlotN9state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotN9state: %+v", err)
			}
			if err := d.Set("slot_raid_link_speed", (s.GetSlotRaidLinkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRaidLinkSpeed: %+v", err)
			}
			if err := d.Set("slot_raid_state", (s.GetSlotRaidState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRaidState: %+v", err)
			}
			if err := d.Set("slot_rear_nvme1link_speed", (s.GetSlotRearNvme1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme1linkSpeed: %+v", err)
			}
			if err := d.Set("slot_rear_nvme1state", (s.GetSlotRearNvme1state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme1state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme2link_speed", (s.GetSlotRearNvme2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme2linkSpeed: %+v", err)
			}
			if err := d.Set("slot_rear_nvme2state", (s.GetSlotRearNvme2state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme2state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme3state", (s.GetSlotRearNvme3state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme3state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme4state", (s.GetSlotRearNvme4state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme4state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme5state", (s.GetSlotRearNvme5state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme5state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme6state", (s.GetSlotRearNvme6state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme6state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme7state", (s.GetSlotRearNvme7state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme7state: %+v", err)
			}
			if err := d.Set("slot_rear_nvme8state", (s.GetSlotRearNvme8state())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRearNvme8state: %+v", err)
			}
			if err := d.Set("slot_riser1link_speed", (s.GetSlotRiser1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser1linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser1slot1link_speed", (s.GetSlotRiser1slot1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser1slot1linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser1slot2link_speed", (s.GetSlotRiser1slot2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser1slot2linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser1slot3link_speed", (s.GetSlotRiser1slot3linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser1slot3linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser2link_speed", (s.GetSlotRiser2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser2linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser2slot4link_speed", (s.GetSlotRiser2slot4linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser2slot4linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser2slot5link_speed", (s.GetSlotRiser2slot5linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser2slot5linkSpeed: %+v", err)
			}
			if err := d.Set("slot_riser2slot6link_speed", (s.GetSlotRiser2slot6linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotRiser2slot6linkSpeed: %+v", err)
			}
			if err := d.Set("slot_sas_state", (s.GetSlotSasState())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotSasState: %+v", err)
			}
			if err := d.Set("slot_ssd_slot1link_speed", (s.GetSlotSsdSlot1linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotSsdSlot1linkSpeed: %+v", err)
			}
			if err := d.Set("slot_ssd_slot2link_speed", (s.GetSlotSsdSlot2linkSpeed())); err != nil {
				return fmt.Errorf("error occurred while setting property SlotSsdSlot2linkSpeed: %+v", err)
			}
			if err := d.Set("smee", (s.GetSmee())); err != nil {
				return fmt.Errorf("error occurred while setting property Smee: %+v", err)
			}
			if err := d.Set("smt_mode", (s.GetSmtMode())); err != nil {
				return fmt.Errorf("error occurred while setting property SmtMode: %+v", err)
			}
			if err := d.Set("snc", (s.GetSnc())); err != nil {
				return fmt.Errorf("error occurred while setting property Snc: %+v", err)
			}
			if err := d.Set("snoopy_mode_for2lm", (s.GetSnoopyModeFor2lm())); err != nil {
				return fmt.Errorf("error occurred while setting property SnoopyModeFor2lm: %+v", err)
			}
			if err := d.Set("snoopy_mode_for_ad", (s.GetSnoopyModeForAd())); err != nil {
				return fmt.Errorf("error occurred while setting property SnoopyModeForAd: %+v", err)
			}
			if err := d.Set("sparing_mode", (s.GetSparingMode())); err != nil {
				return fmt.Errorf("error occurred while setting property SparingMode: %+v", err)
			}
			if err := d.Set("sr_iov", (s.GetSrIov())); err != nil {
				return fmt.Errorf("error occurred while setting property SrIov: %+v", err)
			}
			if err := d.Set("streamer_prefetch", (s.GetStreamerPrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property StreamerPrefetch: %+v", err)
			}
			if err := d.Set("svm_mode", (s.GetSvmMode())); err != nil {
				return fmt.Errorf("error occurred while setting property SvmMode: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("terminal_type", (s.GetTerminalType())); err != nil {
				return fmt.Errorf("error occurred while setting property TerminalType: %+v", err)
			}
			if err := d.Set("tpm_control", (s.GetTpmControl())); err != nil {
				return fmt.Errorf("error occurred while setting property TpmControl: %+v", err)
			}
			if err := d.Set("tpm_support", (s.GetTpmSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property TpmSupport: %+v", err)
			}
			if err := d.Set("txt_support", (s.GetTxtSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property TxtSupport: %+v", err)
			}
			if err := d.Set("ucsm_boot_order_rule", (s.GetUcsmBootOrderRule())); err != nil {
				return fmt.Errorf("error occurred while setting property UcsmBootOrderRule: %+v", err)
			}
			if err := d.Set("ufs_disable", (s.GetUfsDisable())); err != nil {
				return fmt.Errorf("error occurred while setting property UfsDisable: %+v", err)
			}
			if err := d.Set("usb_emul6064", (s.GetUsbEmul6064())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbEmul6064: %+v", err)
			}
			if err := d.Set("usb_port_front", (s.GetUsbPortFront())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortFront: %+v", err)
			}
			if err := d.Set("usb_port_internal", (s.GetUsbPortInternal())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortInternal: %+v", err)
			}
			if err := d.Set("usb_port_kvm", (s.GetUsbPortKvm())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortKvm: %+v", err)
			}
			if err := d.Set("usb_port_rear", (s.GetUsbPortRear())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortRear: %+v", err)
			}
			if err := d.Set("usb_port_sd_card", (s.GetUsbPortSdCard())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortSdCard: %+v", err)
			}
			if err := d.Set("usb_port_vmedia", (s.GetUsbPortVmedia())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbPortVmedia: %+v", err)
			}
			if err := d.Set("usb_xhci_support", (s.GetUsbXhciSupport())); err != nil {
				return fmt.Errorf("error occurred while setting property UsbXhciSupport: %+v", err)
			}
			if err := d.Set("vga_priority", (s.GetVgaPriority())); err != nil {
				return fmt.Errorf("error occurred while setting property VgaPriority: %+v", err)
			}
			if err := d.Set("vmd_enable", (s.GetVmdEnable())); err != nil {
				return fmt.Errorf("error occurred while setting property VmdEnable: %+v", err)
			}
			if err := d.Set("work_load_config", (s.GetWorkLoadConfig())); err != nil {
				return fmt.Errorf("error occurred while setting property WorkLoadConfig: %+v", err)
			}
			if err := d.Set("xpt_prefetch", (s.GetXptPrefetch())); err != nil {
				return fmt.Errorf("error occurred while setting property XptPrefetch: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
