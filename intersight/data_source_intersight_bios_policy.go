package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getBiosPolicySchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"acpi_srat_sp_flag_en": {
			Description: "BIOS Token for setting ACPI SRAT Special Purpose Memory Flag configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu1state": {
			Description: "BIOS Token for setting ACS Control GPU 1 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu2state": {
			Description: "BIOS Token for setting ACS Control GPU 2 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu3state": {
			Description: "BIOS Token for setting ACS Control GPU 3 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu4state": {
			Description: "BIOS Token for setting ACS Control GPU 4 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu5state": {
			Description: "BIOS Token for setting ACS Control GPU 5 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu6state": {
			Description: "BIOS Token for setting ACS Control GPU 6 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu7state": {
			Description: "BIOS Token for setting ACS Control GPU 7 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"acs_control_gpu8state": {
			Description: "BIOS Token for setting ACS Control GPU 8 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
		"adaptive_refresh_mgmt_level": {
			Description: "BIOS Token for setting Adaptive Refresh Management Level configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Default` - Value - Default for configuring AdaptiveRefreshMgmtLevel token.\n* `Level A` - Value - Level A for configuring AdaptiveRefreshMgmtLevel token.\n* `Level B` - Value - Level B for configuring AdaptiveRefreshMgmtLevel token.\n* `Level C` - Value - Level C for configuring AdaptiveRefreshMgmtLevel token.",
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
		"advanced_mem_test": {
			Description: "BIOS Token for setting Enhanced Memory Test configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring AdvancedMemTest token.\n* `disabled` - Value - disabled for configuring AdvancedMemTest token.\n* `enabled` - Value - enabled for configuring AdvancedMemTest token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"all_usb_devices": {
			Description: "BIOS Token for setting All USB Devices configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"altitude": {
			Description: "BIOS Token for setting Altitude configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `300-m` - Value - 300-m for configuring Altitude token.\n* `900-m` - Value - 900-m for configuring Altitude token.\n* `1500-m` - Value - 1500-m for configuring Altitude token.\n* `3000-m` - Value - 3000-m for configuring Altitude token.\n* `auto` - Value - auto for configuring Altitude token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
			Description: "BIOS Token for setting Autonomous Core C State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"autonumous_cstate_enable": {
			Description: "BIOS Token for setting CPU Autonomous C State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"baud_rate": {
			Description: "BIOS Token for setting Baud Rate configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `9600` - Value - 9600 for configuring BaudRate token.\n* `19200` - Value - 19200 for configuring BaudRate token.\n* `38400` - Value - 38400 for configuring BaudRate token.\n* `57600` - Value - 57600 for configuring BaudRate token.\n* `115200` - Value - 115200 for configuring BaudRate token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"bme_dma_mitigation": {
			Description: "BIOS Token for setting BME DMA Mitigation configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"boot_option_num_retry": {
			Description: "BIOS Token for setting Number of Retries configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `5` - Value - 5 for configuring BootOptionNumRetry token.\n* `13` - Value - 13 for configuring BootOptionNumRetry token.\n* `Infinite` - Value - Infinite for configuring BootOptionNumRetry token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"boot_option_re_cool_down": {
			Description: "BIOS Token for setting Cool Down Time  (sec) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `15` - Value - 15 for configuring BootOptionReCoolDown token.\n* `45` - Value - 45 for configuring BootOptionReCoolDown token.\n* `90` - Value - 90 for configuring BootOptionReCoolDown token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"boot_option_retry": {
			Description: "BIOS Token for setting Boot Option Retry configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"boot_performance_mode": {
			Description: "BIOS Token for setting Boot Performance Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Max Efficient` - Value - Max Efficient for configuring BootPerformanceMode token.\n* `Max Performance` - Value - Max Performance for configuring BootPerformanceMode token.\n* `Set by Intel NM` - Value - Set by Intel NM for configuring BootPerformanceMode token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"burst_and_postponed_refresh": {
			Description: "BIOS Token for setting Burst and Postponed Refresh configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"c1auto_demotion": {
			Description: "BIOS Token for setting C1 Auto Demotion configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring C1autoDemotion token.\n* `disabled` - Value - disabled for configuring C1autoDemotion token.\n* `enabled` - Value - enabled for configuring C1autoDemotion token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"c1auto_un_demotion": {
			Description: "BIOS Token for setting C1 Auto UnDemotion configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring C1autoUnDemotion token.\n* `disabled` - Value - disabled for configuring C1autoUnDemotion token.\n* `enabled` - Value - enabled for configuring C1autoUnDemotion token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_apbdis": {
			Description: "BIOS Token for setting APBDIS configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `0` - Value - 0 for configuring CbsCmnApbdis token.\n* `1` - Value - 1 for configuring CbsCmnApbdis token.\n* `Auto` - Value - Auto for configuring CbsCmnApbdis token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_apbdis_df_pstate_rs": {
			Description: "BIOS Token for setting Fixed SOC P-State SP5 F19h configuration (0 - 2 P State).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_avx512": {
			Description: "BIOS Token for setting AVX512 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuAvx512 token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuAvx512 token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuAvx512 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_cpb": {
			Description: "BIOS Token for setting Core Performance Boost configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuCpb token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuCpb token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_gen_downcore_ctrl": {
			Description: "BIOS Token for setting Downcore Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `TWO (1 + 1)` - Value - TWO (1 + 1) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `FOUR (2 + 2)` - Value - FOUR (2 + 2) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `TWO (2 + 0)` - Value - TWO (2 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `SIX (3 + 3)` - Value - SIX (3 + 3) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `THREE (3 + 0)` - Value - THREE (3 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.\n* `FOUR (4 + 0)` - Value - FOUR (4 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_global_cstate_ctrl": {
			Description: "BIOS Token for setting Global C State Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuGlobalCstateCtrl token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuGlobalCstateCtrl token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuGlobalCstateCtrl token.",
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
		"cbs_cmn_cpu_sev_asid_space_limit": {
			Description: "BIOS Token for setting SEV-ES ASID Space Limit configuration (1 - 1007 ASIDs).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_smee": {
			Description: "BIOS Token for setting CPU SMEE configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuSmee token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuSmee token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuSmee token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_cpu_streaming_stores_ctrl": {
			Description: "BIOS Token for setting Streaming Stores Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnCpuStreamingStoresCtrl token.\n* `disabled` - Value - disabled for configuring CbsCmnCpuStreamingStoresCtrl token.\n* `enabled` - Value - enabled for configuring CbsCmnCpuStreamingStoresCtrl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_determinism_slider": {
			Description: "BIOS Token for setting Determinism Slider configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnDeterminismSlider token.\n* `Performance` - Value - Performance for configuring CbsCmnDeterminismSlider token.\n* `Power` - Value - Power for configuring CbsCmnDeterminismSlider token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_edc_control_throttle": {
			Description: "BIOS Token for setting EDC Control Throttle configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnEdcControlThrottle token.\n* `disabled` - Value - disabled for configuring CbsCmnEdcControlThrottle token.\n* `enabled` - Value - enabled for configuring CbsCmnEdcControlThrottle token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_efficiency_mode_en": {
			Description: "BIOS Token for setting Efficiency Mode Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnEfficiencyModeEn token.\n* `Enabled` - Value - Enabled for configuring CbsCmnEfficiencyModeEn token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_efficiency_mode_en_rs": {
			Description: "BIOS Token for setting Power Profile Selection F19h configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnEfficiencyModeEnRs token.\n* `Balanced Core Memory Performance Mode` - Value - Balanced Core Memory Performance Mode for configuring CbsCmnEfficiencyModeEnRs token.\n* `Balanced Core Performance Mode` - Value - Balanced Core Performance Mode for configuring CbsCmnEfficiencyModeEnRs token.\n* `Balanced Memory Performance Mode` - Value - Balanced Memory Performance Mode for configuring CbsCmnEfficiencyModeEnRs token.\n* `Efficiency Mode` - Value - Efficiency Mode for configuring CbsCmnEfficiencyModeEnRs token.\n* `High Performance Mode` - Value - High Performance Mode for configuring CbsCmnEfficiencyModeEnRs token.\n* `Maximum IO Performance Mode` - Value - Maximum IO Performance Mode for configuring CbsCmnEfficiencyModeEnRs token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_fixed_soc_pstate": {
			Description: "BIOS Token for setting Fixed SOC P-State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnFixedSocPstate token.\n* `P0` - Value - P0 for configuring CbsCmnFixedSocPstate token.\n* `P1` - Value - P1 for configuring CbsCmnFixedSocPstate token.\n* `P2` - Value - P2 for configuring CbsCmnFixedSocPstate token.\n* `P3` - Value - P3 for configuring CbsCmnFixedSocPstate token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_gnb_nb_iommu": {
			Description: "BIOS Token for setting IOMMU configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbNbIommu token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbNbIommu token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbNbIommu token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_gnb_smu_df_cstates": {
			Description: "BIOS Token for setting DF C-States configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbSmuDfCstates token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbSmuDfCstates token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbSmuDfCstates token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_gnb_smu_dffo_rs": {
			Description: "BIOS Token for setting DF PState Frequency Optimizer configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbSmuDffoRs token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbSmuDffoRs token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbSmuDffoRs token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_gnb_smu_dlwm_support": {
			Description: "BIOS Token for setting DLWM Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbSmuDlwmSupport token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbSmuDlwmSupport token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbSmuDlwmSupport token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_gnb_smucppc": {
			Description: "BIOS Token for setting CPPC configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnGnbSmucppc token.\n* `disabled` - Value - disabled for configuring CbsCmnGnbSmucppc token.\n* `enabled` - Value - enabled for configuring CbsCmnGnbSmucppc token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_mem_ctrl_bank_group_swap_ddr4": {
			Description: "BIOS Token for setting Bank Group Swap configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.\n* `disabled` - Value - disabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.\n* `enabled` - Value - enabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_mem_ctrller_pwr_dn_en_ddr": {
			Description: "BIOS Token for setting Power Down Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnMemCtrllerPwrDnEnDdr token.\n* `disabled` - Value - disabled for configuring CbsCmnMemCtrllerPwrDnEnDdr token.\n* `enabled` - Value - enabled for configuring CbsCmnMemCtrllerPwrDnEnDdr token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_mem_map_bank_interleave_ddr4": {
			Description: "BIOS Token for setting Chipset Interleave configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnMemMapBankInterleaveDdr4 token.\n* `disabled` - Value - disabled for configuring CbsCmnMemMapBankInterleaveDdr4 token.\n* `Enabled` - Value - Enabled for configuring CbsCmnMemMapBankInterleaveDdr4 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_mem_speed_ddr47xx2": {
			Description: "BIOS Token for setting Memory Clock Speed 7xx2 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `667MHz` - Value - 667MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `800MHz` - Value - 800MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `933MHz` - Value - 933MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `1067MHz` - Value - 1067MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `1200MHz` - Value - 1200MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `1333MHz` - Value - 1333MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `1467MHz` - Value - 1467MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `1600MHz` - Value - 1600MHz for configuring CbsCmnMemSpeedDdr47xx2 token.\n* `Auto` - Value - Auto for configuring CbsCmnMemSpeedDdr47xx2 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_mem_speed_ddr47xx3": {
			Description: "BIOS Token for setting Memory Clock Speed 7xx3 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `400MHz` - Value - 400MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `800MHz` - Value - 800MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `933MHz` - Value - 933MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1067MHz` - Value - 1067MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1200MHz` - Value - 1200MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1333MHz` - Value - 1333MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1467MHz` - Value - 1467MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1600MHz` - Value - 1600MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1633MHz` - Value - 1633MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1667MHz` - Value - 1667MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1700MHz` - Value - 1700MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1733MHz` - Value - 1733MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1767MHz` - Value - 1767MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `1800MHz` - Value - 1800MHz for configuring CbsCmnMemSpeedDdr47xx3 token.\n* `Auto` - Value - Auto for configuring CbsCmnMemSpeedDdr47xx3 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_preferred_io7xx2": {
			Description: "BIOS Token for setting Preferred IO 7xx2 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnPreferredIo7xx2 token.\n* `Manual` - Value - Manual for configuring CbsCmnPreferredIo7xx2 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmn_preferred_io7xx3": {
			Description: "BIOS Token for setting Preferred IO 7xx3 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmnPreferredIo7xx3 token.\n* `Bus` - Value - Bus for configuring CbsCmnPreferredIo7xx3 token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmnc_tdp_ctl": {
			Description: "BIOS Token for setting cTDP Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCmncTdpCtl token.\n* `Manual` - Value - Manual for configuring CbsCmncTdpCtl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cmnx_gmi_force_link_width_rs": {
			Description: "BIOS Token for setting xGMI Force Link Width configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `0` - Value - 0 for configuring CbsCmnxGmiForceLinkWidthRs token.\n* `1` - Value - 1 for configuring CbsCmnxGmiForceLinkWidthRs token.\n* `2` - Value - 2 for configuring CbsCmnxGmiForceLinkWidthRs token.\n* `Auto` - Value - Auto for configuring CbsCmnxGmiForceLinkWidthRs token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cpu_ccd_ctrl_ssp": {
			Description: "BIOS Token for setting CCD Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `2 CCDs` - Value - 2 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `3 CCDs` - Value - 3 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `4 CCDs` - Value - 4 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `6 CCDs` - Value - 6 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `8 CCDs` - Value - 8 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `10 CCDs` - Value - 10 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `12 CCDs` - Value - 12 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `14 CCDs` - Value - 14 CCDs for configuring CbsCpuCcdCtrlSsp token.\n* `Auto` - Value - Auto for configuring CbsCpuCcdCtrlSsp token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cpu_core_ctrl": {
			Description: "BIOS Token for setting CPU Downcore control 7xx3 configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCpuCoreCtrl token.\n* `ONE (1 + 0)` - Value - ONE (1 + 0) for configuring CbsCpuCoreCtrl token.\n* `TWO (2 + 0)` - Value - TWO (2 + 0) for configuring CbsCpuCoreCtrl token.\n* `THREE (3 + 0)` - Value - THREE (3 + 0) for configuring CbsCpuCoreCtrl token.\n* `FOUR (4 + 0)` - Value - FOUR (4 + 0) for configuring CbsCpuCoreCtrl token.\n* `FIVE (5 + 0)` - Value - FIVE (5 + 0) for configuring CbsCpuCoreCtrl token.\n* `SIX (6 + 0)` - Value - SIX (6 + 0) for configuring CbsCpuCoreCtrl token.\n* `SEVEN (7 + 0)` - Value - SEVEN (7 + 0) for configuring CbsCpuCoreCtrl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cpu_down_core_ctrl_bergamo": {
			Description: "BIOS Token for setting Downcore control F19 MA0h-AFh configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCpuDownCoreCtrlBergamo token.\n* `TWO (1 + 1)` - Value - TWO (1 + 1) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `FOUR (2 + 2)` - Value - FOUR (2 + 2) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `SIX (3 + 3)` - Value - SIX (3 + 3) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `EIGHT (4 + 4)` - Value - EIGHT (4 + 4) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `TEN (5 + 5)` - Value - TEN (5 + 5) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `TWELVE (6 + 6)` - Value - TWELVE (6 + 6) for configuring CbsCpuDownCoreCtrlBergamo token.\n* `FOURTEEN (7 + 7)` - Value - FOURTEEN (7 + 7) for configuring CbsCpuDownCoreCtrlBergamo token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cpu_down_core_ctrl_genoa": {
			Description: "BIOS Token for setting CPU Downcore control F19 M10h-1Fh configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCpuDownCoreCtrlGenoa token.\n* `ONE (1 + 0)` - Value - ONE (1 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `TWO (2 + 0)` - Value - TWO (2 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `THREE (3 + 0)` - Value - THREE (3 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `FOUR (4 + 0)` - Value - FOUR (4 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `FIVE (5 + 0)` - Value - FIVE (5 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `SIX (6 + 0)` - Value - SIX (6 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `SEVEN (7 + 0)` - Value - SEVEN (7 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `EIGHT (8 + 0)` - Value - EIGHT (8 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `NINE (9 + 0)` - Value - NINE (9 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `TEN (10 + 0)` - Value - TEN (10 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `ELEVEN (11 + 0)` - Value - ELEVEN (11 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `TWELVE (12 + 0)` - Value - TWELVE (12 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `THIRTEEN (13 + 0)` - Value - THIRTEEN (13 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `FOURTEEN (14 + 0)` - Value - FOURTEEN (14 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.\n* `FIFTEEN (15 + 0)` - Value - FIFTEEN (15 + 0) for configuring CbsCpuDownCoreCtrlGenoa token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_cpu_smt_ctrl": {
			Description: "BIOS Token for setting CPU SMT Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsCpuSmtCtrl token.\n* `disabled` - Value - disabled for configuring CbsCpuSmtCtrl token.\n* `enabled` - Value - enabled for configuring CbsCpuSmtCtrl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_dbg_cpu_gen_cpu_wdt": {
			Description: "BIOS Token for setting Core Watchdog Timer Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDbgCpuGenCpuWdt token.\n* `disabled` - Value - disabled for configuring CbsDbgCpuGenCpuWdt token.\n* `enabled` - Value - enabled for configuring CbsDbgCpuGenCpuWdt token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_dbg_cpu_lapic_mode": {
			Description: "BIOS Token for setting Local APIC Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDbgCpuLapicMode token.\n* `Compatibility` - Value - Compatibility for configuring CbsDbgCpuLapicMode token.\n* `X2APIC` - Value - X2APIC for configuring CbsDbgCpuLapicMode token.\n* `XAPIC` - Value - XAPIC for configuring CbsDbgCpuLapicMode token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_dbg_cpu_snp_mem_cover": {
			Description: "BIOS Token for setting SNP Memory Coverage configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDbgCpuSnpMemCover token.\n* `Custom` - Value - Custom for configuring CbsDbgCpuSnpMemCover token.\n* `disabled` - Value - disabled for configuring CbsDbgCpuSnpMemCover token.\n* `enabled` - Value - enabled for configuring CbsDbgCpuSnpMemCover token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_dbg_cpu_snp_mem_size_cover": {
			Description: "BIOS Token for setting SNP Memory Size to Cover in MiB configuration (0 - 1048576 MiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn4link_max_xgmi_speed": {
			Description: "BIOS Token for setting 4-link xGMI max speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `20Gbps` - Value - 20Gbps for configuring CbsDfCmn4linkMaxXgmiSpeed token.\n* `25Gbps` - Value - 25Gbps for configuring CbsDfCmn4linkMaxXgmiSpeed token.\n* `32Gbps` - Value - 32Gbps for configuring CbsDfCmn4linkMaxXgmiSpeed token.\n* `Auto` - Value - Auto for configuring CbsDfCmn4linkMaxXgmiSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_acpi_srat_l3numa": {
			Description: "BIOS Token for setting ACPI SRAT L3 Cache As NUMA Domain configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDfCmnAcpiSratL3numa token.\n* `disabled` - Value - disabled for configuring CbsDfCmnAcpiSratL3numa token.\n* `enabled` - Value - enabled for configuring CbsDfCmnAcpiSratL3numa token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_dram_nps": {
			Description: "BIOS Token for setting NUMA Nodes per Socket configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDfCmnDramNps token.\n* `NPS0` - Value - NPS0 for configuring CbsDfCmnDramNps token.\n* `NPS1` - Value - NPS1 for configuring CbsDfCmnDramNps token.\n* `NPS2` - Value - NPS2 for configuring CbsDfCmnDramNps token.\n* `NPS4` - Value - NPS4 for configuring CbsDfCmnDramNps token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_dram_scrub_time": {
			Description: "BIOS Token for setting DRAM Scrub Time configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1 hour` - Value - 1 hour for configuring CbsDfCmnDramScrubTime token.\n* `4 hours` - Value - 4 hours for configuring CbsDfCmnDramScrubTime token.\n* `6 hours` - Value - 6 hours for configuring CbsDfCmnDramScrubTime token.\n* `8 hours` - Value - 8 hours for configuring CbsDfCmnDramScrubTime token.\n* `12 hours` - Value - 12 hours for configuring CbsDfCmnDramScrubTime token.\n* `16 hours` - Value - 16 hours for configuring CbsDfCmnDramScrubTime token.\n* `24 hours` - Value - 24 hours for configuring CbsDfCmnDramScrubTime token.\n* `48 hours` - Value - 48 hours for configuring CbsDfCmnDramScrubTime token.\n* `Auto` - Value - Auto for configuring CbsDfCmnDramScrubTime token.\n* `Disabled` - Value - Disabled for configuring CbsDfCmnDramScrubTime token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_mem_intlv": {
			Description: "BIOS Token for setting AMD Memory Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlv token.\n* `Channel` - Value - Channel for configuring CbsDfCmnMemIntlv token.\n* `Die` - Value - Die for configuring CbsDfCmnMemIntlv token.\n* `None` - Value - None for configuring CbsDfCmnMemIntlv token.\n* `Socket` - Value - Socket for configuring CbsDfCmnMemIntlv token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_mem_intlv_control": {
			Description: "BIOS Token for setting Memory Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlvControl token.\n* `disabled` - Value - disabled for configuring CbsDfCmnMemIntlvControl token.\n* `enabled` - Value - enabled for configuring CbsDfCmnMemIntlvControl token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_cmn_mem_intlv_size": {
			Description: "BIOS Token for setting AMD Memory Interleaving Size configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `256 Bytes` - Value - 256 Bytes for configuring CbsDfCmnMemIntlvSize token.\n* `512 Bytes` - Value - 512 Bytes for configuring CbsDfCmnMemIntlvSize token.\n* `1 KB` - Value - 1 KiB for configuring CbsDfCmnMemIntlvSize token.\n* `2 KB` - Value - 2 KiB for configuring CbsDfCmnMemIntlvSize token.\n* `4 KB` - Value - 4 KiB for configuring CbsDfCmnMemIntlvSize token.\n* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlvSize token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_df_dbg_xgmi_link_cfg": {
			Description: "BIOS Token for setting xGMI Link Configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `2 xGMI Links` - Value - 2 xGMI Links for configuring CbsDfDbgXgmiLinkCfg token.\n* `3 xGMI Links` - Value - 3 xGMI Links for configuring CbsDfDbgXgmiLinkCfg token.\n* `4 xGMI Links` - Value - 4 xGMI Links for configuring CbsDfDbgXgmiLinkCfg token.\n* `Auto` - Value - Auto for configuring CbsDfDbgXgmiLinkCfg token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_gnb_dbg_pcie_tbt_support": {
			Description: "BIOS Token for setting PCIe Ten Bit Tag Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsGnbDbgPcieTbtSupport token.\n* `disabled` - Value - disabled for configuring CbsGnbDbgPcieTbtSupport token.\n* `enabled` - Value - enabled for configuring CbsGnbDbgPcieTbtSupport token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cbs_sev_snp_support": {
			Description: "BIOS Token for setting SEV-SNP Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CbsSevSnpSupport token.\n* `disabled` - Value - disabled for configuring CbsSevSnpSupport token.\n* `enabled` - Value - enabled for configuring CbsSevSnpSupport token.",
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
		"cisco_xgmi_max_speed": {
			Description: "BIOS Token for setting Cisco xGMI Max Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cke_low_policy": {
			Description: "BIOS Token for setting CKE Low Policy configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring CkeLowPolicy token.\n* `disabled` - Value - disabled for configuring CkeLowPolicy token.\n* `fast` - Value - fast for configuring CkeLowPolicy token.\n* `slow` - Value - slow for configuring CkeLowPolicy token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"closed_loop_therm_throtl": {
			Description: "BIOS Token for setting Closed Loop Thermal Throttling configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
			Description: "BIOS Token for setting Console Redirection configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `com-0` - Value - com-0 for configuring ConsoleRedirection token.\n* `com-1` - Value - com-1 for configuring ConsoleRedirection token.\n* `disabled` - Value - disabled for configuring ConsoleRedirection token.\n* `enabled` - Value - enabled for configuring ConsoleRedirection token.\n* `serial-port-a` - Value - serial-port-a for configuring ConsoleRedirection token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"core_multi_processing": {
			Description: "BIOS Token for setting Core Multi Processing configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1` - Value - 1 for configuring CoreMultiProcessing token.\n* `2` - Value - 2 for configuring CoreMultiProcessing token.\n* `3` - Value - 3 for configuring CoreMultiProcessing token.\n* `4` - Value - 4 for configuring CoreMultiProcessing token.\n* `5` - Value - 5 for configuring CoreMultiProcessing token.\n* `6` - Value - 6 for configuring CoreMultiProcessing token.\n* `7` - Value - 7 for configuring CoreMultiProcessing token.\n* `8` - Value - 8 for configuring CoreMultiProcessing token.\n* `9` - Value - 9 for configuring CoreMultiProcessing token.\n* `10` - Value - 10 for configuring CoreMultiProcessing token.\n* `11` - Value - 11 for configuring CoreMultiProcessing token.\n* `12` - Value - 12 for configuring CoreMultiProcessing token.\n* `13` - Value - 13 for configuring CoreMultiProcessing token.\n* `14` - Value - 14 for configuring CoreMultiProcessing token.\n* `15` - Value - 15 for configuring CoreMultiProcessing token.\n* `16` - Value - 16 for configuring CoreMultiProcessing token.\n* `17` - Value - 17 for configuring CoreMultiProcessing token.\n* `18` - Value - 18 for configuring CoreMultiProcessing token.\n* `19` - Value - 19 for configuring CoreMultiProcessing token.\n* `20` - Value - 20 for configuring CoreMultiProcessing token.\n* `21` - Value - 21 for configuring CoreMultiProcessing token.\n* `22` - Value - 22 for configuring CoreMultiProcessing token.\n* `23` - Value - 23 for configuring CoreMultiProcessing token.\n* `24` - Value - 24 for configuring CoreMultiProcessing token.\n* `25` - Value - 25 for configuring CoreMultiProcessing token.\n* `26` - Value - 26 for configuring CoreMultiProcessing token.\n* `27` - Value - 27 for configuring CoreMultiProcessing token.\n* `28` - Value - 28 for configuring CoreMultiProcessing token.\n* `29` - Value - 29 for configuring CoreMultiProcessing token.\n* `30` - Value - 30 for configuring CoreMultiProcessing token.\n* `31` - Value - 31 for configuring CoreMultiProcessing token.\n* `32` - Value - 32 for configuring CoreMultiProcessing token.\n* `33` - Value - 33 for configuring CoreMultiProcessing token.\n* `34` - Value - 34 for configuring CoreMultiProcessing token.\n* `35` - Value - 35 for configuring CoreMultiProcessing token.\n* `36` - Value - 36 for configuring CoreMultiProcessing token.\n* `37` - Value - 37 for configuring CoreMultiProcessing token.\n* `38` - Value - 38 for configuring CoreMultiProcessing token.\n* `39` - Value - 39 for configuring CoreMultiProcessing token.\n* `40` - Value - 40 for configuring CoreMultiProcessing token.\n* `41` - Value - 41 for configuring CoreMultiProcessing token.\n* `42` - Value - 42 for configuring CoreMultiProcessing token.\n* `43` - Value - 43 for configuring CoreMultiProcessing token.\n* `44` - Value - 44 for configuring CoreMultiProcessing token.\n* `45` - Value - 45 for configuring CoreMultiProcessing token.\n* `46` - Value - 46 for configuring CoreMultiProcessing token.\n* `47` - Value - 47 for configuring CoreMultiProcessing token.\n* `48` - Value - 48 for configuring CoreMultiProcessing token.\n* `49` - Value - 49 for configuring CoreMultiProcessing token.\n* `50` - Value - 50 for configuring CoreMultiProcessing token.\n* `51` - Value - 51 for configuring CoreMultiProcessing token.\n* `52` - Value - 52 for configuring CoreMultiProcessing token.\n* `53` - Value - 53 for configuring CoreMultiProcessing token.\n* `54` - Value - 54 for configuring CoreMultiProcessing token.\n* `55` - Value - 55 for configuring CoreMultiProcessing token.\n* `56` - Value - 56 for configuring CoreMultiProcessing token.\n* `57` - Value - 57 for configuring CoreMultiProcessing token.\n* `58` - Value - 58 for configuring CoreMultiProcessing token.\n* `59` - Value - 59 for configuring CoreMultiProcessing token.\n* `60` - Value - 60 for configuring CoreMultiProcessing token.\n* `61` - Value - 61 for configuring CoreMultiProcessing token.\n* `62` - Value - 62 for configuring CoreMultiProcessing token.\n* `63` - Value - 63 for configuring CoreMultiProcessing token.\n* `64` - Value - 64 for configuring CoreMultiProcessing token.\n* `65` - Value - 65 for configuring CoreMultiProcessing token.\n* `66` - Value - 66 for configuring CoreMultiProcessing token.\n* `67` - Value - 67 for configuring CoreMultiProcessing token.\n* `68` - Value - 68 for configuring CoreMultiProcessing token.\n* `69` - Value - 69 for configuring CoreMultiProcessing token.\n* `70` - Value - 70 for configuring CoreMultiProcessing token.\n* `71` - Value - 71 for configuring CoreMultiProcessing token.\n* `72` - Value - 72 for configuring CoreMultiProcessing token.\n* `73` - Value - 73 for configuring CoreMultiProcessing token.\n* `74` - Value - 74 for configuring CoreMultiProcessing token.\n* `75` - Value - 75 for configuring CoreMultiProcessing token.\n* `76` - Value - 76 for configuring CoreMultiProcessing token.\n* `77` - Value - 77 for configuring CoreMultiProcessing token.\n* `78` - Value - 78 for configuring CoreMultiProcessing token.\n* `79` - Value - 79 for configuring CoreMultiProcessing token.\n* `80` - Value - 80 for configuring CoreMultiProcessing token.\n* `81` - Value - 81 for configuring CoreMultiProcessing token.\n* `82` - Value - 82 for configuring CoreMultiProcessing token.\n* `83` - Value - 83 for configuring CoreMultiProcessing token.\n* `84` - Value - 84 for configuring CoreMultiProcessing token.\n* `85` - Value - 85 for configuring CoreMultiProcessing token.\n* `86` - Value - 86 for configuring CoreMultiProcessing token.\n* `all` - Value - all for configuring CoreMultiProcessing token.",
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
		"cpu_pa_limit": {
			Description: "BIOS Token for setting Limit CPU PA to 46 Bits configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu_perf_enhancement": {
			Description: "BIOS Token for setting Enhanced CPU Performance configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CpuPerfEnhancement token.\n* `Disabled` - Value - Disabled for configuring CpuPerfEnhancement token.",
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
			Description: "BIOS Token for setting CR QoS configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring CrQos token.\n* `Mode 0 - Disable the PMem QoS Feature` - Value - Mode 0 - Disable the PMem QoS Feature for configuring CrQos token.\n* `Mode 1 - M2M QoS Enable and CHA QoS Disable` - Value - Mode 1 - M2M QoS Enable and CHA QoS Disable for configuring CrQos token.\n* `Mode 2 - M2M QoS Enable and CHA QoS Enable` - Value - Mode 2 - M2M QoS Enable and CHA QoS Enable for configuring CrQos token.\n* `Profile 1` - Value - Profile 1 for configuring CrQos token.\n* `Recipe 1` - Value - Recipe 1 for configuring CrQos token.\n* `Recipe 2` - Value - Recipe 2 for configuring CrQos token.\n* `Recipe 3` - Value - Recipe 3 for configuring CrQos token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"crfastgo_config": {
			Description: "BIOS Token for setting CR FastGo Config configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring CrfastgoConfig token.\n* `Default` - Value - Default for configuring CrfastgoConfig token.\n* `Disable optimization` - Value - Disable optimization for configuring CrfastgoConfig token.\n* `Enable optimization` - Value - Enable optimization for configuring CrfastgoConfig token.\n* `Option 1` - Value - Option 1 for configuring CrfastgoConfig token.\n* `Option 2` - Value - Option 2 for configuring CrfastgoConfig token.\n* `Option 3` - Value - Option 3 for configuring CrfastgoConfig token.\n* `Option 4` - Value - Option 4 for configuring CrfastgoConfig token.\n* `Option 5` - Value - Option 5 for configuring CrfastgoConfig token.",
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
		"dfx_osb_en": {
			Description: "BIOS Token for setting DFX OSB configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring DfxOsbEn token.\n* `disabled` - Value - disabled for configuring DfxOsbEn token.\n* `enabled` - Value - enabled for configuring DfxOsbEn token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"direct_cache_access": {
			Description: "BIOS Token for setting Direct Cache Access Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring DirectCacheAccess token.\n* `disabled` - Value - disabled for configuring DirectCacheAccess token.\n* `enabled` - Value - enabled for configuring DirectCacheAccess token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dma_ctrl_opt_in": {
			Description: "BIOS Token for setting DMA Control Opt-In Flag configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
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
		"dram_sw_thermal_throttling": {
			Description: "BIOS Token for setting DRAM SW Thermal Throttling configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"eadr_support": {
			Description: "BIOS Token for setting eADR Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring EadrSupport token.\n* `disabled` - Value - disabled for configuring EadrSupport token.\n* `enabled` - Value - enabled for configuring EadrSupport token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"edpc_en": {
			Description: "BIOS Token for setting IIO eDPC Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring EdpcEn token.\n* `On Fatal Error` - Value - On Fatal Error for configuring EdpcEn token.\n* `On Fatal and Non-Fatal Errors` - Value - On Fatal and Non-Fatal Errors for configuring EdpcEn token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_clock_spread_spec": {
			Description: "BIOS Token for setting External SSC Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `0P3_Percent` - Value - 0P3_Percent for configuring EnableClockSpreadSpec token.\n* `0P5_Percent` - Value - 0P5_Percent for configuring EnableClockSpreadSpec token.\n* `disabled` - Value - disabled for configuring EnableClockSpreadSpec token.\n* `enabled` - Value - enabled for configuring EnableClockSpreadSpec token.\n* `Hardware` - Value - Hardware for configuring EnableClockSpreadSpec token.\n* `Off` - Value - Off for configuring EnableClockSpreadSpec token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_mktme": {
			Description: "BIOS Token for setting Multikey Total Memory Encryption  (MK-TME) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_rmt": {
			Description: "BIOS Token for setting Rank Margin Tool configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_sgx": {
			Description: "BIOS Token for setting Software Guard Extensions  (SGX) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_tdx": {
			Description: "BIOS Token for setting Trust Domain Extension  (TDX) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_tdx_seamldr": {
			Description: "BIOS Token for setting TDX Secure Arbitration Mode  (SEAM) Loader configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"enable_tme": {
			Description: "BIOS Token for setting Total Memory Encryption  (TME) configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
		"epoch_update": {
			Description: "BIOS Token for setting Select Owner EPOCH Input Type configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Change to New Random Owner EPOCHs` - Value - Change to New Random Owner EPOCHs for configuring EpochUpdate token.\n* `Manual User Defined Owner EPOCHs` - Value - Manual User Defined Owner EPOCHs for configuring EpochUpdate token.\n* `SGX Owner EPOCH activated` - Value - SGX Owner EPOCH activated for configuring EpochUpdate token.\n* `SGX Owner EPOCH deactivated` - Value - SGX Owner EPOCH deactivated for configuring EpochUpdate token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"epp_enable": {
			Description: "BIOS Token for setting Processor EPP Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"epp_profile": {
			Description: "BIOS Token for setting EPP Profile configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Balanced Performance` - Value - Balanced Performance for configuring EppProfile token.\n* `Balanced Power` - Value - Balanced Power for configuring EppProfile token.\n* `Performance` - Value - Performance for configuring EppProfile token.\n* `Power` - Value - Power for configuring EppProfile token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"error_check_scrub": {
			Description: "BIOS Token for setting Error Check Scrub configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring ErrorCheckScrub token.\n* `Enabled with Result Collection` - Value - Enabled with Result Collection for configuring ErrorCheckScrub token.\n* `Enabled without Result Collection` - Value - Enabled without Result Collection for configuring ErrorCheckScrub token.",
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
		"intel_dynamic_speed_select": {
			Description: "BIOS Token for setting Intel Dynamic Speed Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"intel_hyper_threading_tech": {
			Description: "BIOS Token for setting Intel HyperThreading Tech configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"intel_speed_select": {
			Description: "BIOS Token for setting Intel Speed Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring IntelSpeedSelect token.\n* `Base` - Value - Base for configuring IntelSpeedSelect token.\n* `Config 1` - Value - Config 1 for configuring IntelSpeedSelect token.\n* `Config 2` - Value - Config 2 for configuring IntelSpeedSelect token.\n* `Config 3` - Value - Config 3 for configuring IntelSpeedSelect token.\n* `Config 4` - Value - Config 4 for configuring IntelSpeedSelect token.\n* `Level 0` - Value - Level 0 for configuring IntelSpeedSelect token.\n* `Level 1` - Value - Level 1 for configuring IntelSpeedSelect token.\n* `Level 2` - Value - Level 2 for configuring IntelSpeedSelect token.\n* `Level 3` - Value - Level 3 for configuring IntelSpeedSelect token.\n* `Level 4` - Value - Level 4 for configuring IntelSpeedSelect token.",
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
			Description: "BIOS Token for setting Intel VT for Directed IO configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
			Description: "BIOS Token for setting Intel (R) VT-d PassThrough DMA Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"intel_vtdats_support": {
			Description: "BIOS Token for setting Intel VTD ATS Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ioat_config_cpm": {
			Description: "BIOS Token for setting IOAT Configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
		"ipv4http": {
			Description: "BIOS Token for setting IPV4 HTTP Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ipv4pxe": {
			Description: "BIOS Token for setting IPv4 PXE Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ipv6http": {
			Description: "BIOS Token for setting IPV6 HTTP Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ipv6pxe": {
			Description: "BIOS Token for setting IPV6 PXE Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"kti_prefetch": {
			Description: "BIOS Token for setting KTI Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring KtiPrefetch token.\n* `disabled` - Value - disabled for configuring KtiPrefetch token.\n* `enabled` - Value - enabled for configuring KtiPrefetch token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"latency_optimized_mode": {
			Description: "BIOS Token for setting Latency Optimized Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"legacy_os_redirection": {
			Description: "BIOS Token for setting Legacy OS Redirection configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"legacy_usb_support": {
			Description: "BIOS Token for setting Legacy USB Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `auto` - Value - auto for configuring LegacyUsbSupport token.\n* `disabled` - Value - disabled for configuring LegacyUsbSupport token.\n* `enabled` - Value - enabled for configuring LegacyUsbSupport token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"llc_alloc": {
			Description: "BIOS Token for setting LLC Dead Line configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring LlcAlloc token.\n* `disabled` - Value - disabled for configuring LlcAlloc token.\n* `enabled` - Value - enabled for configuring LlcAlloc token.",
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
		"memory_bandwidth_boost": {
			Description: "BIOS Token for setting Memory Bandwidth Boost configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_inter_leave": {
			Description: "BIOS Token for setting Intel Memory Interleaving configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1 Way Node Interleave` - Value - 1 Way Node Interleave for configuring MemoryInterLeave token.\n* `2 Way Node Interleave` - Value - 2 Way Node Interleave for configuring MemoryInterLeave token.\n* `4 Way Node Interleave` - Value - 4 Way Node Interleave for configuring MemoryInterLeave token.\n* `8 Way Node Interleave` - Value - 8 Way Node Interleave for configuring MemoryInterLeave token.\n* `disabled` - Value - disabled for configuring MemoryInterLeave token.\n* `enabled` - Value - enabled for configuring MemoryInterLeave token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_mapped_io_above4gb": {
			Description: "BIOS Token for setting Memory Mapped IO above 4GiB configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_refresh_rate": {
			Description: "BIOS Token for setting Memory Refresh Rate configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1x Refresh` - Value - 1x Refresh for configuring MemoryRefreshRate token.\n* `2x Refresh` - Value - 2x Refresh for configuring MemoryRefreshRate token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_size_limit": {
			Description: "BIOS Token for setting Memory Size Limit in GiB configuration (0 - 65535 GiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_thermal_throttling": {
			Description: "BIOS Token for setting Memory Thermal Throttling Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `CLTT with PECI` - Value - CLTT with PECI for configuring MemoryThermalThrottling token.\n* `Disabled` - Value - Disabled for configuring MemoryThermalThrottling token.",
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
		"mmioh_base": {
			Description: "BIOS Token for setting MMIO High Base configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `512G` - Value - 512G for configuring MmiohBase token.\n* `1T` - Value - 1T for configuring MmiohBase token.\n* `2T` - Value - 2T for configuring MmiohBase token.\n* `4T` - Value - 4T for configuring MmiohBase token.\n* `16T` - Value - 16T for configuring MmiohBase token.\n* `24T` - Value - 24T for configuring MmiohBase token.\n* `32T` - Value - 32T for configuring MmiohBase token.\n* `40T` - Value - 40T for configuring MmiohBase token.\n* `56T` - Value - 56T for configuring MmiohBase token.\n* `Auto` - Value - Auto for configuring MmiohBase token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mmioh_size": {
			Description: "BIOS Token for setting MMIO High Granularity Size configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1G` - Value - 1G for configuring MmiohSize token.\n* `4G` - Value - 4G for configuring MmiohSize token.\n* `16G` - Value - 16G for configuring MmiohSize token.\n* `32G` - Value - 32G for configuring MmiohSize token.\n* `64G` - Value - 64G for configuring MmiohSize token.\n* `256G` - Value - 256G for configuring MmiohSize token.\n* `1024G` - Value - 1024G for configuring MmiohSize token.\n* `Auto` - Value - Auto for configuring MmiohSize token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
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
			Description: "BIOS Token for setting NUMA Optimized configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nvmdimm_perform_config": {
			Description: "BIOS Token for setting NVM Performance Setting configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `BW Optimized` - Value - BW Optimized for configuring NvmdimmPerformConfig token.\n* `Balanced Profile` - Value - Balanced Profile for configuring NvmdimmPerformConfig token.\n* `Latency Optimized` - Value - Latency Optimized for configuring NvmdimmPerformConfig token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"operation_mode": {
			Description: "BIOS Token for setting Operation Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Test Only` - Value - Test Only for configuring OperationMode token.\n* `Test and Repair` - Value - Test and Repair for configuring OperationMode token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"optimized_power_mode": {
			Description: "BIOS Token for setting Optimized Power Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
			Description: "BIOS Token for setting OS Boot Watchdog Timer Timeout configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `5-minutes` - Value - 5-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `10-minutes` - Value - 10-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `15-minutes` - Value - 15-minutes for configuring OsBootWatchdogTimerTimeout token.\n* `20-minutes` - Value - 20-minutes for configuring OsBootWatchdogTimerTimeout token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"out_of_band_mgmt_port": {
			Description: "BIOS Token for setting Out-of-Band Mgmt Port configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"package_cstate_limit": {
			Description: "BIOS Token for setting Package C State Limit configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PackageCstateLimit token.\n* `C0 C1 State` - Value - C0 C1 State for configuring PackageCstateLimit token.\n* `C0/C1` - Value - C0/C1 for configuring PackageCstateLimit token.\n* `C2` - Value - C2 for configuring PackageCstateLimit token.\n* `C6 Non Retention` - Value - C6 Non Retention for configuring PackageCstateLimit token.\n* `C6 Retention` - Value - C6 Retention for configuring PackageCstateLimit token.\n* `No Limit` - Value - No Limit for configuring PackageCstateLimit token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"panic_high_watermark": {
			Description: "BIOS Token for setting Panic and High Watermark configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `High` - Value - High for configuring PanicHighWatermark token.\n* `Low` - Value - Low for configuring PanicHighWatermark token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"partial_cache_line_sparing": {
			Description: "BIOS Token for setting Partial Cache Line Sparing configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_mode_config": {
			Description: "BIOS Token for setting Partial Memory Mirror Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring PartialMirrorModeConfig token.\n* `Percentage` - Value - Percentage for configuring PartialMirrorModeConfig token.\n* `Value in GB` - Value - Value in GiB for configuring PartialMirrorModeConfig token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_percent": {
			Description: "BIOS Token for setting Partial Mirror Percentage configuration (0.00 - 50.00 Percentage).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_value1": {
			Description: "BIOS Token for setting Partial Mirror1 Size in GiB configuration (0 - 65535 GiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_value2": {
			Description: "BIOS Token for setting Partial Mirror2 Size in GiB configuration (0 - 65535 GiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_value3": {
			Description: "BIOS Token for setting Partial Mirror3 Size in GiB configuration (0 - 65535 GiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"partial_mirror_value4": {
			Description: "BIOS Token for setting Partial Mirror4 Size in GiB configuration (0 - 65535 GiB).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"patrol_scrub": {
			Description: "BIOS Token for setting Patrol Scrub configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring PatrolScrub token.\n* `Enable at End of POST` - Value - Enable at End of POST for configuring PatrolScrub token.\n* `enabled` - Value - enabled for configuring PatrolScrub token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"patrol_scrub_duration": {
			Description: "BIOS Token for setting Patrol Scrub Interval configuration (5 - 23 Hour).",
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
		"pch_pcie_pll_ssc": {
			Description: "BIOS Token for setting PCIe PLL SSC Percent configuration (0 - 255 (n/10)%).",
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
		"pcie_slot_mraid1link_speed": {
			Description: "BIOS Token for setting MRAID1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotMraid1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotMraid1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotMraid1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotMraid1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotMraid1linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring PcieSlotMraid1linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring PcieSlotMraid1linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_mraid1option_rom": {
			Description: "BIOS Token for setting MRAID1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_mraid2link_speed": {
			Description: "BIOS Token for setting MRAID2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotMraid2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotMraid2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotMraid2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotMraid2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotMraid2linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring PcieSlotMraid2linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring PcieSlotMraid2linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_mraid2option_rom": {
			Description: "BIOS Token for setting MRAID2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_mstorraid_link_speed": {
			Description: "BIOS Token for setting PCIe Slot MSTOR Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotMstorraidLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotMstorraidLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotMstorraidLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotMstorraidLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotMstorraidLinkSpeed token.\n* `GEN4` - Value - GEN4 for configuring PcieSlotMstorraidLinkSpeed token.\n* `GEN5` - Value - GEN5 for configuring PcieSlotMstorraidLinkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_mstorraid_option_rom": {
			Description: "BIOS Token for setting PCIe Slot MSTOR RAID OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme1link_speed": {
			Description: "BIOS Token for setting NVME 1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme1linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme1option_rom": {
			Description: "BIOS Token for setting NVME 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme2link_speed": {
			Description: "BIOS Token for setting NVME 2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme2linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme2option_rom": {
			Description: "BIOS Token for setting NVME 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme3link_speed": {
			Description: "BIOS Token for setting NVME 3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme3linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme3option_rom": {
			Description: "BIOS Token for setting NVME 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme4link_speed": {
			Description: "BIOS Token for setting NVME 4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme4linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme4option_rom": {
			Description: "BIOS Token for setting NVME 4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme5link_speed": {
			Description: "BIOS Token for setting NVME 5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme5linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme5option_rom": {
			Description: "BIOS Token for setting NVME 5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme6link_speed": {
			Description: "BIOS Token for setting NVME 6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring PcieSlotNvme6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring PcieSlotNvme6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring PcieSlotNvme6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring PcieSlotNvme6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring PcieSlotNvme6linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slot_nvme6option_rom": {
			Description: "BIOS Token for setting NVME 6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pcie_slots_cdn_enable": {
			Description: "BIOS Token for setting PCIe Slots CDN Control configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
		"post_package_repair": {
			Description: "BIOS Token for setting Post Package Repair configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disabled` - Value - Disabled for configuring PostPackageRepair token.\n* `Hard PPR` - Value - Hard PPR for configuring PostPackageRepair token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"pre_boot_dma_protection": {
			Description: "BIOS Token for setting PreBoot DMA Protection configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"prmrr_size": {
			Description: "BIOS Token for setting PRMRR Size configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1G` - Value - 1G for configuring PrmrrSize token.\n* `2G` - Value - 2G for configuring PrmrrSize token.\n* `4G` - Value - 4G for configuring PrmrrSize token.\n* `8G` - Value - 8G for configuring PrmrrSize token.\n* `16G` - Value - 16G for configuring PrmrrSize token.\n* `32G` - Value - 32G for configuring PrmrrSize token.\n* `64G` - Value - 64G for configuring PrmrrSize token.\n* `128G` - Value - 128G for configuring PrmrrSize token.\n* `256G` - Value - 256G for configuring PrmrrSize token.\n* `512G` - Value - 512G for configuring PrmrrSize token.\n* `128M` - Value - 128M for configuring PrmrrSize token.\n* `256M` - Value - 256M for configuring PrmrrSize token.\n* `512M` - Value - 512M for configuring PrmrrSize token.\n* `Auto` - Value - Auto for configuring PrmrrSize token.\n* `Invalid Config.` - Value - Invalid Config for configuring PrmrrSize token.",
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
			Description: "BIOS Token for setting Processor C6 Report configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring ProcessorC6report token.\n* `disabled` - Value - disabled for configuring ProcessorC6report token.\n* `enabled` - Value - enabled for configuring ProcessorC6report token.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"psata": {
			Description: "BIOS Token for setting P-SATA Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `AHCI` - Value - AHCI for configuring Psata token.\n* `Disabled` - Value - Disabled for configuring Psata token.\n* `LSI SW RAID` - Value - LSI SW RAID for configuring Psata token.",
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
			Description: "BIOS Token for setting Power Performance Tuning configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `bios` - Value - BIOS for configuring PwrPerfTuning token.\n* `os` - Value - os for configuring PwrPerfTuning token.\n* `peci` - Value - peci for configuring PwrPerfTuning token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"qpi_link_frequency": {
			Description: "BIOS Token for setting QPI Link Frequency Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `6.4-gt/s` - Value - 6.4-gt/s for configuring QpiLinkFrequency token.\n* `7.2-gt/s` - Value - 7.2-gt/s for configuring QpiLinkFrequency token.\n* `8.0-gt/s` - Value - 8.0-gt/s for configuring QpiLinkFrequency token.\n* `9.6-gt/s` - Value - 9.6-gt/s for configuring QpiLinkFrequency token.\n* `auto` - Value - auto for configuring QpiLinkFrequency token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"qpi_link_speed": {
			Description: "BIOS Token for setting UPI Link Frequency Select configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `10.4GT/s` - Value - 10.4GT/s for configuring QpiLinkSpeed token.\n* `11.2GT/s` - Value - 11.2GT/s for configuring QpiLinkSpeed token.\n* `12.8GT/s` - Value - 12.8GT/s for configuring QpiLinkSpeed token.\n* `14.4GT/s` - Value - 14.4GT/s for configuring QpiLinkSpeed token.\n* `16.0GT/s` - Value - 16.0GT/s for configuring QpiLinkSpeed token.\n* `20.0GT/s` - Value - 20.0GT/s for configuring QpiLinkSpeed token.\n* `24.0GT/s` - Value - 24.0GT/s for configuring QpiLinkSpeed token.\n* `9.6GT/s` - Value - 9.6GT/s for configuring QpiLinkSpeed token.\n* `Auto` - Value - Auto for configuring QpiLinkSpeed token.\n* `Use Per Link Setting` - Value - Use Per Link Setting for configuring QpiLinkSpeed token.",
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
		"resize_bar_support": {
			Description: "BIOS Token for setting Re-Size BAR Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"runtime_post_package_repair": {
			Description: "BIOS Token for setting Runtime Post Package Repair configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sata_mode_select": {
			Description: "BIOS Token for setting SATA Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `AHCI` - Value - AHCI for configuring SataModeSelect token.\n* `Disabled` - Value - Disabled for configuring SataModeSelect token.\n* `LSI SW RAID` - Value - LSI SW RAID for configuring SataModeSelect token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"select_memory_ras_configuration": {
			Description: "BIOS Token for setting Memory RAS Configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `adddc-sparing` - Value - adddc-sparing for configuring SelectMemoryRasConfiguration token.\n* `lockstep` - Value - lockstep for configuring SelectMemoryRasConfiguration token.\n* `maximum-performance` - Value - maximum-performance for configuring SelectMemoryRasConfiguration token.\n* `mirror-mode-1lm` - Value - mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.\n* `mirroring` - Value - mirroring for configuring SelectMemoryRasConfiguration token.\n* `partial-mirror-mode-1lm` - Value - partial-mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.\n* `sparing` - Value - sparing for configuring SelectMemoryRasConfiguration token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"select_ppr_type": {
			Description: "BIOS Token for setting PPR Type configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SelectPprType token.\n* `Hard PPR` - Value - Hard PPR for configuring SelectPprType token.\n* `Soft PPR` - Value - Soft PPR for configuring SelectPprType token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial_mux": {
			Description: "BIOS Token for setting Serial Mux configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial_port_aenable": {
			Description: "BIOS Token for setting Serial A Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sev": {
			Description: "BIOS Token for setting Secured Encrypted Virtualization configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `253 ASIDs` - Value - 253 ASIDs for configuring Sev token.\n* `509 ASIDs` - Value - 509 ASIDs for configuring Sev token.\n* `Auto` - Value - Auto for configuring Sev token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_auto_registration_agent": {
			Description: "BIOS Token for setting SGX Auto MP Registration Agent configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_epoch0": {
			Description: "BIOS Token for setting SGX Epoch 0 configuration (0 - ffffffffffffffff Hash byte 7-0).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_epoch1": {
			Description: "BIOS Token for setting SGX Epoch 1 configuration (0 - ffffffffffffffff Hash byte 7-0).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_factory_reset": {
			Description: "BIOS Token for setting SGX Factory Reset configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_le_pub_key_hash0": {
			Description: "BIOS Token for setting SGX PubKey Hash0 configuration (0 - ffffffffffffffff Hash byte 7-0).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_le_pub_key_hash1": {
			Description: "BIOS Token for setting SGX PubKey Hash1 configuration (0 - ffffffffffffffff Hash byte 15-8).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_le_pub_key_hash2": {
			Description: "BIOS Token for setting SGX PubKey Hash2 configuration (0 - ffffffffffffffff Hash byte 23-16).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_le_pub_key_hash3": {
			Description: "BIOS Token for setting SGX PubKey Hash3 configuration (0 - ffffffffffffffff Hash byte 31-24).",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_le_wr": {
			Description: "BIOS Token for setting SGX Write Enable configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_package_info_in_band_access": {
			Description: "BIOS Token for setting SGX Package Information In-Band Access configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sgx_qos": {
			Description: "BIOS Token for setting SGX QoS configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sha1pcr_bank": {
			Description: "BIOS Token for setting SHA-1 PCR Bank configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sha256pcr_bank": {
			Description: "BIOS Token for setting SHA256 PCR Bank configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sha384pcr_bank": {
			Description: "BIOS Token for setting SHA384 PCR Bank configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
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
			Description: "BIOS Token for setting Slot 13 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot14state": {
			Description: "BIOS Token for setting Slot 14 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot1link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot1linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot1linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot1linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot1state": {
			Description: "BIOS Token for setting Slot 1 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot1state token.\n* `enabled` - Value - enabled for configuring Slot1state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot1state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot1state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot2link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot2linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot2linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot2linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot2state": {
			Description: "BIOS Token for setting Slot 2 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot2state token.\n* `enabled` - Value - enabled for configuring Slot2state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot2state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot2state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot3link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot3linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot3linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot3linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot3state": {
			Description: "BIOS Token for setting Slot 3 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot3state token.\n* `enabled` - Value - enabled for configuring Slot3state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot3state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot3state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot4link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot4linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot4linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot4linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot4state": {
			Description: "BIOS Token for setting Slot 4 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot4state token.\n* `enabled` - Value - enabled for configuring Slot4state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot4state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot4state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot5link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot5linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot5linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot5linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot5state": {
			Description: "BIOS Token for setting Slot 5 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot5state token.\n* `enabled` - Value - enabled for configuring Slot5state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot5state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot5state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot6link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot6linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot6linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot6linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot6state": {
			Description: "BIOS Token for setting Slot 6 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot6state token.\n* `enabled` - Value - enabled for configuring Slot6state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot6state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot6state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot7link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 7 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot7linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot7linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot7linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot7linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot7linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot7linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot7linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot7state": {
			Description: "BIOS Token for setting Slot 7 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot7state token.\n* `enabled` - Value - enabled for configuring Slot7state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot7state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot7state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot8link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 8 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot8linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot8linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot8linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot8linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot8linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot8linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring Slot8linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot8state": {
			Description: "BIOS Token for setting Slot 8 State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring Slot8state token.\n* `enabled` - Value - enabled for configuring Slot8state token.\n* `Legacy Only` - Value - Legacy Only for configuring Slot8state token.\n* `UEFI Only` - Value - UEFI Only for configuring Slot8state token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot9link_speed": {
			Description: "BIOS Token for setting PCIe Slot: 9 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Slot9linkSpeed token.\n* `Disabled` - Value - Disabled for configuring Slot9linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring Slot9linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring Slot9linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring Slot9linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring Slot9linkSpeed token.",
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
		"slot_front_nvme10link_speed": {
			Description: "BIOS Token for setting Front NVME 10 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme10linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme10linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme10linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme10linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme10linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme10linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme10linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme10option_rom": {
			Description: "BIOS Token for setting Front NVME 10 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme11link_speed": {
			Description: "BIOS Token for setting Front NVME 11 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme11linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme11linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme11linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme11linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme11linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme11linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme11linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme11option_rom": {
			Description: "BIOS Token for setting Front NVME 11 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme12link_speed": {
			Description: "BIOS Token for setting Front NVME 12 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme12linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme12linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme12linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme12linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme12linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme12linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme12linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme12option_rom": {
			Description: "BIOS Token for setting Front NVME 12 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme13link_speed": {
			Description: "BIOS Token for setting Front NVME 13 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme13linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme13linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme13linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme13linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme13linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme13linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme13linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme13option_rom": {
			Description: "BIOS Token for setting Front NVME 13 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme14link_speed": {
			Description: "BIOS Token for setting Front NVME 14 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme14linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme14linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme14linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme14linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme14linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme14linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme14linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme14option_rom": {
			Description: "BIOS Token for setting Front NVME 14 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme15link_speed": {
			Description: "BIOS Token for setting Front NVME 15 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme15linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme15linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme15linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme15linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme15linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme15linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme15linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme15option_rom": {
			Description: "BIOS Token for setting Front NVME 15 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme16link_speed": {
			Description: "BIOS Token for setting Front NVME 16 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme16linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme16linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme16linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme16linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme16linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme16linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme16linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme16option_rom": {
			Description: "BIOS Token for setting Front NVME 16 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme17link_speed": {
			Description: "BIOS Token for setting Front NVME 17 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme17linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme17linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme17linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme17linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme17linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme17linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme17linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme17option_rom": {
			Description: "BIOS Token for setting Front NVME 17 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme18link_speed": {
			Description: "BIOS Token for setting Front NVME 18 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme18linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme18linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme18linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme18linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme18linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme18linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme18linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme18option_rom": {
			Description: "BIOS Token for setting Front NVME 18 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme19link_speed": {
			Description: "BIOS Token for setting Front NVME 19 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme19linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme19linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme19linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme19linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme19linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme19linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme19linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme19option_rom": {
			Description: "BIOS Token for setting Front NVME 19 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme1link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme1linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme1linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme1option_rom": {
			Description: "BIOS Token for setting Front NVME 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme20link_speed": {
			Description: "BIOS Token for setting Front NVME 20 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme20linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme20linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme20linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme20linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme20linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme20linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme20linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme20option_rom": {
			Description: "BIOS Token for setting Front NVME 20 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme21link_speed": {
			Description: "BIOS Token for setting Front NVME 21 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme21linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme21linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme21linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme21linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme21linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme21linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme21linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme21option_rom": {
			Description: "BIOS Token for setting Front NVME 21 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme22link_speed": {
			Description: "BIOS Token for setting Front NVME 22 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme22linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme22linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme22linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme22linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme22linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme22linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme22linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme22option_rom": {
			Description: "BIOS Token for setting Front NVME 22 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme23link_speed": {
			Description: "BIOS Token for setting Front NVME 23 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme23linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme23linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme23linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme23linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme23linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme23linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme23linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme23option_rom": {
			Description: "BIOS Token for setting Front NVME 23 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme24link_speed": {
			Description: "BIOS Token for setting Front NVME 24 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme24linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme24linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme24linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme24linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme24linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme24linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme24linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme24option_rom": {
			Description: "BIOS Token for setting Front NVME 24 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme25link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 25 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme25linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme25linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme25linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme25linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme25linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme25linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme25linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme25option_rom": {
			Description: "BIOS Token for setting Front NVME 25 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme26link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 26 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme26linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme26linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme26linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme26linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme26linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme26linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme26linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme26option_rom": {
			Description: "BIOS Token for setting Front NVME 26 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme27link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 27 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme27linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme27linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme27linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme27linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme27linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme27linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme27linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme27option_rom": {
			Description: "BIOS Token for setting Front NVME 27 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme28link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 28 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme28linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme28linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme28linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme28linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme28linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme28linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme28linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme28option_rom": {
			Description: "BIOS Token for setting Front NVME 28 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme29link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 29 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme29linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme29linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme29linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme29linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme29linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme29linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme29linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme29option_rom": {
			Description: "BIOS Token for setting Front NVME 29 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme2link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme2linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme2linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme2option_rom": {
			Description: "BIOS Token for setting Front NVME 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme30link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 30 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme30linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme30linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme30linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme30linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme30linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme30linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme30linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme30option_rom": {
			Description: "BIOS Token for setting Front NVME 30 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme31link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 31 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme31linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme31linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme31linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme31linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme31linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme31linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme31linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme31option_rom": {
			Description: "BIOS Token for setting Front NVME 31 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme32link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Front NVME 32 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme32linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme32linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme32linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme32linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme32linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme32linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme32linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme32option_rom": {
			Description: "BIOS Token for setting Front NVME 32 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme3link_speed": {
			Description: "BIOS Token for setting Front NVME 3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme3linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme3linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme3linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme3option_rom": {
			Description: "BIOS Token for setting Front NVME 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme4link_speed": {
			Description: "BIOS Token for setting Front NVME 4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme4linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme4linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme4linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme4option_rom": {
			Description: "BIOS Token for setting Front NVME 4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme5link_speed": {
			Description: "BIOS Token for setting Front NVME 5 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme5linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme5linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme5linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme5linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme5linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme5linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme5linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme5option_rom": {
			Description: "BIOS Token for setting Front NVME 5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme6link_speed": {
			Description: "BIOS Token for setting Front NVME 6 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme6linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme6linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme6linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme6linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme6linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme6linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme6linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme6option_rom": {
			Description: "BIOS Token for setting Front NVME 6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme7link_speed": {
			Description: "BIOS Token for setting Front NVME 7 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme7linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme7linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme7linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme7linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme7linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme7linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme7linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme7option_rom": {
			Description: "BIOS Token for setting Front NVME 7 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme8link_speed": {
			Description: "BIOS Token for setting Front NVME 8 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme8linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme8linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme8linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme8linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme8linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme8linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme8linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme8option_rom": {
			Description: "BIOS Token for setting Front NVME 8 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme9link_speed": {
			Description: "BIOS Token for setting Front NVME 9 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotFrontNvme9linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotFrontNvme9linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotFrontNvme9linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotFrontNvme9linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotFrontNvme9linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotFrontNvme9linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotFrontNvme9linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_front_nvme9option_rom": {
			Description: "BIOS Token for setting Front NVME 9 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
			Description: "BIOS Token for setting GPU 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu2state": {
			Description: "BIOS Token for setting GPU 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu3state": {
			Description: "BIOS Token for setting GPU 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu4state": {
			Description: "BIOS Token for setting GPU 4 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu5state": {
			Description: "BIOS Token for setting GPU 5 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu6state": {
			Description: "BIOS Token for setting GPU 6 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu7state": {
			Description: "BIOS Token for setting GPU 7 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_gpu8state": {
			Description: "BIOS Token for setting GPU 8 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
			Description: "BIOS Token for setting PCIe Slot:MLOM Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotMlomLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotMlomLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotMlomLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotMlomLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotMlomLinkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotMlomLinkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotMlomLinkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_mlom_state": {
			Description: "BIOS Token for setting PCIe Slot MLOM OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring SlotMlomState token.\n* `enabled` - Value - enabled for configuring SlotMlomState token.\n* `Legacy Only` - Value - Legacy Only for configuring SlotMlomState token.\n* `UEFI Only` - Value - UEFI Only for configuring SlotMlomState token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_mraid_link_speed": {
			Description: "BIOS Token for setting MRAID Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotMraidLinkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotMraidLinkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotMraidLinkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotMraidLinkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotMraidLinkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotMraidLinkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotMraidLinkSpeed token.",
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
			Description: "BIOS Token for setting PCIe Slot:Rear NVME 1 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme1linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme1linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme1linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme1linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme1linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotRearNvme1linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotRearNvme1linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme1state": {
			Description: "BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme2link_speed": {
			Description: "BIOS Token for setting PCIe Slot:Rear NVME 2 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme2linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme2linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme2linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme2linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme2linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotRearNvme2linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotRearNvme2linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme2state": {
			Description: "BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme3link_speed": {
			Description: "BIOS Token for setting Rear NVME 3 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme3linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme3linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme3linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme3linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme3linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotRearNvme3linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotRearNvme3linkSpeed token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme3state": {
			Description: "BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_rear_nvme4link_speed": {
			Description: "BIOS Token for setting Rear NVME 4 Link Speed configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring SlotRearNvme4linkSpeed token.\n* `Disabled` - Value - Disabled for configuring SlotRearNvme4linkSpeed token.\n* `GEN1` - Value - GEN1 for configuring SlotRearNvme4linkSpeed token.\n* `GEN2` - Value - GEN2 for configuring SlotRearNvme4linkSpeed token.\n* `GEN3` - Value - GEN3 for configuring SlotRearNvme4linkSpeed token.\n* `GEN4` - Value - GEN4 for configuring SlotRearNvme4linkSpeed token.\n* `GEN5` - Value - GEN5 for configuring SlotRearNvme4linkSpeed token.",
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
			Description: "BIOS Token for setting Sub Numa Clustering configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Snc token.\n* `disabled` - Value - disabled for configuring Snc token.\n* `enabled` - Value - enabled for configuring Snc token.\n* `SNC2` - Value - SNC2 for configuring Snc token.\n* `SNC4` - Value - SNC4 for configuring Snc token.",
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
			Description: "BIOS Token for setting DCU Streamer Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring StreamerPrefetch token.\n* `disabled` - Value - disabled for configuring StreamerPrefetch token.\n* `enabled` - Value - enabled for configuring StreamerPrefetch token.",
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
					"ancestor_definitions": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"definition": {
						Description: "The definition is a reference to the tag definition object.\nThe tag definition object contains the properties of the tag such as name, type, and description.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"propagated": {
						Description: "Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"type": {
						Description: "An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.\n* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.\n* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".",
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
			Description: "BIOS Token for setting Terminal Type configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `pc-ansi` - Value - pc-ansi for configuring TerminalType token.\n* `vt100` - Value - vt100 for configuring TerminalType token.\n* `vt100-plus` - Value - vt100-plus for configuring TerminalType token.\n* `vt-utf8` - Value - vt-utf8 for configuring TerminalType token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tpm_control": {
			Description: "BIOS Token for setting Trusted Platform Module State configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tpm_pending_operation": {
			Description: "BIOS Token for setting TPM Pending Operation configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `None` - Value - None for configuring TpmPendingOperation token.\n* `TpmClear` - Value - TpmClear for configuring TpmPendingOperation token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tpm_ppi_required": {
			Description: "BIOS Token for setting TPM Minimal Physical Presence configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tpm_support": {
			Description: "BIOS Token for setting Security Device Support configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tsme": {
			Description: "BIOS Token for setting Transparent Secure Memory Encryption configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring Tsme token.\n* `disabled` - Value - disabled for configuring Tsme token.\n* `enabled` - Value - enabled for configuring Tsme token.",
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
		"uefi_mem_map_sp_flag_en": {
			Description: "BIOS Token for setting UEFI Memory Map Special Purpose Memory Flag configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ufs_disable": {
			Description: "BIOS Token for setting Uncore Frequency Scaling configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `disabled` - Value - disabled for configuring UfsDisable token.\n* `enabled` - Value - enabled for configuring UfsDisable token.\n* `Mode 0` - Value - Mode 0 for configuring UfsDisable token.\n* `Mode 1` - Value - Mode 1 for configuring UfsDisable token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"uma_based_clustering": {
			Description: "BIOS Token for setting UMA Based Clustering configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Disable (All2All)` - Value - Disable (All2All) for configuring UmaBasedClustering token.\n* `Hemisphere (2-clusters)` - Value - Hemisphere (2-clusters) for configuring UmaBasedClustering token.\n* `Quadrant (4-clusters)` - Value - Quadrant (4-clusters) for configuring UmaBasedClustering token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upi_link_enablement": {
			Description: "BIOS Token for setting UPI Link Enablement configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1` - Value - 1 for configuring UpiLinkEnablement token.\n* `2` - Value - 2 for configuring UpiLinkEnablement token.\n* `3` - Value - 3 for configuring UpiLinkEnablement token.\n* `Auto` - Value - Auto for configuring UpiLinkEnablement token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upi_power_management": {
			Description: "BIOS Token for setting UPI Power Manangement configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
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
		"version_context": {
			Description: "The versioning info for this managed object.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"marked_for_deletion": {
						Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vga_priority": {
			Description: "BIOS Token for setting VGA Priority configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Offboard` - Value - Offboard for configuring VgaPriority token.\n* `Onboard` - Value - Onboard for configuring VgaPriority token.\n* `Onboard VGA Disabled` - Value - Onboard VGA Disabled for configuring VgaPriority token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"virtual_numa": {
			Description: "BIOS Token for setting Virtual NUMA configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vmd_enable": {
			Description: "BIOS Token for setting VMD Enablement configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vol_memory_mode": {
			Description: "BIOS Token for setting Volatile Memory Mode configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `1LM` - Value - 1LM for configuring VolMemoryMode token.\n* `2LM` - Value - 2LM for configuring VolMemoryMode token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"work_load_config": {
			Description: "BIOS Token for setting Workload Configuration configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Balanced` - Value - Balanced for configuring WorkLoadConfig token.\n* `I/O Sensitive` - Value - I/O Sensitive for configuring WorkLoadConfig token.\n* `NUMA` - Value - NUMA for configuring WorkLoadConfig token.\n* `UMA` - Value - UMA for configuring WorkLoadConfig token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"x2apic_opt_out": {
			Description: "BIOS Token for setting X2APIC Opt-Out Flag configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `enabled` - Enables the BIOS setting.\n* `disabled` - Disables the BIOS setting.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"xpt_prefetch": {
			Description: "BIOS Token for setting XPT Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring XptPrefetch token.\n* `disabled` - Value - disabled for configuring XptPrefetch token.\n* `enabled` - Value - enabled for configuring XptPrefetch token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"xpt_remote_prefetch": {
			Description: "BIOS Token for setting XPT Remote Prefetch configuration.\n* `platform-default` - Default value used by the platform for the BIOS setting.\n* `Auto` - Value - Auto for configuring XptRemotePrefetch token.\n* `disabled` - Value - disabled for configuring XptRemotePrefetch token.\n* `enabled` - Value - enabled for configuring XptRemotePrefetch token.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
	return schemaMap
}

func dataSourceBiosPolicy() *schema.Resource {
	var subSchema = getBiosPolicySchema()
	var model = getBiosPolicySchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceBiosPolicyRead,
		Schema:      model}
}

func dataSourceBiosPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.BiosPolicy{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("acpi_srat_sp_flag_en"); ok {
		x := (v.(string))
		o.SetAcpiSratSpFlagEn(x)
	}

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

	if v, ok := d.GetOk("adaptive_refresh_mgmt_level"); ok {
		x := (v.(string))
		o.SetAdaptiveRefreshMgmtLevel(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("adjacent_cache_line_prefetch"); ok {
		x := (v.(string))
		o.SetAdjacentCacheLinePrefetch(x)
	}

	if v, ok := d.GetOk("advanced_mem_test"); ok {
		x := (v.(string))
		o.SetAdvancedMemTest(x)
	}

	if v, ok := d.GetOk("all_usb_devices"); ok {
		x := (v.(string))
		o.SetAllUsbDevices(x)
	}

	if v, ok := d.GetOk("altitude"); ok {
		x := (v.(string))
		o.SetAltitude(x)
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
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

	if v, ok := d.GetOk("burst_and_postponed_refresh"); ok {
		x := (v.(string))
		o.SetBurstAndPostponedRefresh(x)
	}

	if v, ok := d.GetOk("c1auto_demotion"); ok {
		x := (v.(string))
		o.SetC1autoDemotion(x)
	}

	if v, ok := d.GetOk("c1auto_un_demotion"); ok {
		x := (v.(string))
		o.SetC1autoUnDemotion(x)
	}

	if v, ok := d.GetOk("cbs_cmn_apbdis"); ok {
		x := (v.(string))
		o.SetCbsCmnApbdis(x)
	}

	if v, ok := d.GetOk("cbs_cmn_apbdis_df_pstate_rs"); ok {
		x := (v.(string))
		o.SetCbsCmnApbdisDfPstateRs(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_avx512"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuAvx512(x)
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

	if v, ok := d.GetOk("cbs_cmn_cpu_sev_asid_space_limit"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuSevAsidSpaceLimit(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_smee"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuSmee(x)
	}

	if v, ok := d.GetOk("cbs_cmn_cpu_streaming_stores_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuStreamingStoresCtrl(x)
	}

	if v, ok := d.GetOk("cbs_cmn_determinism_slider"); ok {
		x := (v.(string))
		o.SetCbsCmnDeterminismSlider(x)
	}

	if v, ok := d.GetOk("cbs_cmn_edc_control_throttle"); ok {
		x := (v.(string))
		o.SetCbsCmnEdcControlThrottle(x)
	}

	if v, ok := d.GetOk("cbs_cmn_efficiency_mode_en"); ok {
		x := (v.(string))
		o.SetCbsCmnEfficiencyModeEn(x)
	}

	if v, ok := d.GetOk("cbs_cmn_efficiency_mode_en_rs"); ok {
		x := (v.(string))
		o.SetCbsCmnEfficiencyModeEnRs(x)
	}

	if v, ok := d.GetOk("cbs_cmn_fixed_soc_pstate"); ok {
		x := (v.(string))
		o.SetCbsCmnFixedSocPstate(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_nb_iommu"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbNbIommu(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_smu_df_cstates"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbSmuDfCstates(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_smu_dffo_rs"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbSmuDffoRs(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_smu_dlwm_support"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbSmuDlwmSupport(x)
	}

	if v, ok := d.GetOk("cbs_cmn_gnb_smucppc"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbSmucppc(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_ctrl_bank_group_swap_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemCtrlBankGroupSwapDdr4(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_ctrller_pwr_dn_en_ddr"); ok {
		x := (v.(string))
		o.SetCbsCmnMemCtrllerPwrDnEnDdr(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_map_bank_interleave_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemMapBankInterleaveDdr4(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_speed_ddr47xx2"); ok {
		x := (v.(string))
		o.SetCbsCmnMemSpeedDdr47xx2(x)
	}

	if v, ok := d.GetOk("cbs_cmn_mem_speed_ddr47xx3"); ok {
		x := (v.(string))
		o.SetCbsCmnMemSpeedDdr47xx3(x)
	}

	if v, ok := d.GetOk("cbs_cmn_preferred_io7xx2"); ok {
		x := (v.(string))
		o.SetCbsCmnPreferredIo7xx2(x)
	}

	if v, ok := d.GetOk("cbs_cmn_preferred_io7xx3"); ok {
		x := (v.(string))
		o.SetCbsCmnPreferredIo7xx3(x)
	}

	if v, ok := d.GetOk("cbs_cmnc_tdp_ctl"); ok {
		x := (v.(string))
		o.SetCbsCmncTdpCtl(x)
	}

	if v, ok := d.GetOk("cbs_cmnx_gmi_force_link_width_rs"); ok {
		x := (v.(string))
		o.SetCbsCmnxGmiForceLinkWidthRs(x)
	}

	if v, ok := d.GetOk("cbs_cpu_ccd_ctrl_ssp"); ok {
		x := (v.(string))
		o.SetCbsCpuCcdCtrlSsp(x)
	}

	if v, ok := d.GetOk("cbs_cpu_core_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCpuCoreCtrl(x)
	}

	if v, ok := d.GetOk("cbs_cpu_down_core_ctrl_bergamo"); ok {
		x := (v.(string))
		o.SetCbsCpuDownCoreCtrlBergamo(x)
	}

	if v, ok := d.GetOk("cbs_cpu_down_core_ctrl_genoa"); ok {
		x := (v.(string))
		o.SetCbsCpuDownCoreCtrlGenoa(x)
	}

	if v, ok := d.GetOk("cbs_cpu_smt_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCpuSmtCtrl(x)
	}

	if v, ok := d.GetOk("cbs_dbg_cpu_gen_cpu_wdt"); ok {
		x := (v.(string))
		o.SetCbsDbgCpuGenCpuWdt(x)
	}

	if v, ok := d.GetOk("cbs_dbg_cpu_lapic_mode"); ok {
		x := (v.(string))
		o.SetCbsDbgCpuLapicMode(x)
	}

	if v, ok := d.GetOk("cbs_dbg_cpu_snp_mem_cover"); ok {
		x := (v.(string))
		o.SetCbsDbgCpuSnpMemCover(x)
	}

	if v, ok := d.GetOk("cbs_dbg_cpu_snp_mem_size_cover"); ok {
		x := (v.(string))
		o.SetCbsDbgCpuSnpMemSizeCover(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn4link_max_xgmi_speed"); ok {
		x := (v.(string))
		o.SetCbsDfCmn4linkMaxXgmiSpeed(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_acpi_srat_l3numa"); ok {
		x := (v.(string))
		o.SetCbsDfCmnAcpiSratL3numa(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_dram_nps"); ok {
		x := (v.(string))
		o.SetCbsDfCmnDramNps(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_dram_scrub_time"); ok {
		x := (v.(string))
		o.SetCbsDfCmnDramScrubTime(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlv(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv_control"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvControl(x)
	}

	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv_size"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvSize(x)
	}

	if v, ok := d.GetOk("cbs_df_dbg_xgmi_link_cfg"); ok {
		x := (v.(string))
		o.SetCbsDfDbgXgmiLinkCfg(x)
	}

	if v, ok := d.GetOk("cbs_gnb_dbg_pcie_tbt_support"); ok {
		x := (v.(string))
		o.SetCbsGnbDbgPcieTbtSupport(x)
	}

	if v, ok := d.GetOk("cbs_sev_snp_support"); ok {
		x := (v.(string))
		o.SetCbsSevSnpSupport(x)
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

	if v, ok := d.GetOk("cisco_xgmi_max_speed"); ok {
		x := (v.(string))
		o.SetCiscoXgmiMaxSpeed(x)
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

	if v, ok := d.GetOk("cpu_pa_limit"); ok {
		x := (v.(string))
		o.SetCpuPaLimit(x)
	}

	if v, ok := d.GetOk("cpu_perf_enhancement"); ok {
		x := (v.(string))
		o.SetCpuPerfEnhancement(x)
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

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
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

	if v, ok := d.GetOk("dfx_osb_en"); ok {
		x := (v.(string))
		o.SetDfxOsbEn(x)
	}

	if v, ok := d.GetOk("direct_cache_access"); ok {
		x := (v.(string))
		o.SetDirectCacheAccess(x)
	}

	if v, ok := d.GetOk("dma_ctrl_opt_in"); ok {
		x := (v.(string))
		o.SetDmaCtrlOptIn(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("dram_clock_throttling"); ok {
		x := (v.(string))
		o.SetDramClockThrottling(x)
	}

	if v, ok := d.GetOk("dram_refresh_rate"); ok {
		x := (v.(string))
		o.SetDramRefreshRate(x)
	}

	if v, ok := d.GetOk("dram_sw_thermal_throttling"); ok {
		x := (v.(string))
		o.SetDramSwThermalThrottling(x)
	}

	if v, ok := d.GetOk("eadr_support"); ok {
		x := (v.(string))
		o.SetEadrSupport(x)
	}

	if v, ok := d.GetOk("edpc_en"); ok {
		x := (v.(string))
		o.SetEdpcEn(x)
	}

	if v, ok := d.GetOk("enable_clock_spread_spec"); ok {
		x := (v.(string))
		o.SetEnableClockSpreadSpec(x)
	}

	if v, ok := d.GetOk("enable_mktme"); ok {
		x := (v.(string))
		o.SetEnableMktme(x)
	}

	if v, ok := d.GetOk("enable_rmt"); ok {
		x := (v.(string))
		o.SetEnableRmt(x)
	}

	if v, ok := d.GetOk("enable_sgx"); ok {
		x := (v.(string))
		o.SetEnableSgx(x)
	}

	if v, ok := d.GetOk("enable_tdx"); ok {
		x := (v.(string))
		o.SetEnableTdx(x)
	}

	if v, ok := d.GetOk("enable_tdx_seamldr"); ok {
		x := (v.(string))
		o.SetEnableTdxSeamldr(x)
	}

	if v, ok := d.GetOk("enable_tme"); ok {
		x := (v.(string))
		o.SetEnableTme(x)
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

	if v, ok := d.GetOk("epoch_update"); ok {
		x := (v.(string))
		o.SetEpochUpdate(x)
	}

	if v, ok := d.GetOk("epp_enable"); ok {
		x := (v.(string))
		o.SetEppEnable(x)
	}

	if v, ok := d.GetOk("epp_profile"); ok {
		x := (v.(string))
		o.SetEppProfile(x)
	}

	if v, ok := d.GetOk("error_check_scrub"); ok {
		x := (v.(string))
		o.SetErrorCheckScrub(x)
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

	if v, ok := d.GetOk("intel_dynamic_speed_select"); ok {
		x := (v.(string))
		o.SetIntelDynamicSpeedSelect(x)
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

	if v, ok := d.GetOk("ioat_config_cpm"); ok {
		x := (v.(string))
		o.SetIoatConfigCpm(x)
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

	if v, ok := d.GetOk("ipv4http"); ok {
		x := (v.(string))
		o.SetIpv4http(x)
	}

	if v, ok := d.GetOk("ipv4pxe"); ok {
		x := (v.(string))
		o.SetIpv4pxe(x)
	}

	if v, ok := d.GetOk("ipv6http"); ok {
		x := (v.(string))
		o.SetIpv6http(x)
	}

	if v, ok := d.GetOk("ipv6pxe"); ok {
		x := (v.(string))
		o.SetIpv6pxe(x)
	}

	if v, ok := d.GetOk("kti_prefetch"); ok {
		x := (v.(string))
		o.SetKtiPrefetch(x)
	}

	if v, ok := d.GetOk("latency_optimized_mode"); ok {
		x := (v.(string))
		o.SetLatencyOptimizedMode(x)
	}

	if v, ok := d.GetOk("legacy_os_redirection"); ok {
		x := (v.(string))
		o.SetLegacyOsRedirection(x)
	}

	if v, ok := d.GetOk("legacy_usb_support"); ok {
		x := (v.(string))
		o.SetLegacyUsbSupport(x)
	}

	if v, ok := d.GetOk("llc_alloc"); ok {
		x := (v.(string))
		o.SetLlcAlloc(x)
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

	if v, ok := d.GetOk("memory_bandwidth_boost"); ok {
		x := (v.(string))
		o.SetMemoryBandwidthBoost(x)
	}

	if v, ok := d.GetOk("memory_inter_leave"); ok {
		x := (v.(string))
		o.SetMemoryInterLeave(x)
	}

	if v, ok := d.GetOk("memory_mapped_io_above4gb"); ok {
		x := (v.(string))
		o.SetMemoryMappedIoAbove4gb(x)
	}

	if v, ok := d.GetOk("memory_refresh_rate"); ok {
		x := (v.(string))
		o.SetMemoryRefreshRate(x)
	}

	if v, ok := d.GetOk("memory_size_limit"); ok {
		x := (v.(string))
		o.SetMemorySizeLimit(x)
	}

	if v, ok := d.GetOk("memory_thermal_throttling"); ok {
		x := (v.(string))
		o.SetMemoryThermalThrottling(x)
	}

	if v, ok := d.GetOk("mirroring_mode"); ok {
		x := (v.(string))
		o.SetMirroringMode(x)
	}

	if v, ok := d.GetOk("mmcfg_base"); ok {
		x := (v.(string))
		o.SetMmcfgBase(x)
	}

	if v, ok := d.GetOk("mmioh_base"); ok {
		x := (v.(string))
		o.SetMmiohBase(x)
	}

	if v, ok := d.GetOk("mmioh_size"); ok {
		x := (v.(string))
		o.SetMmiohSize(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
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

	if v, ok := d.GetOk("operation_mode"); ok {
		x := (v.(string))
		o.SetOperationMode(x)
	}

	if v, ok := d.GetOk("optimized_power_mode"); ok {
		x := (v.(string))
		o.SetOptimizedPowerMode(x)
	}

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
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

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("package_cstate_limit"); ok {
		x := (v.(string))
		o.SetPackageCstateLimit(x)
	}

	if v, ok := d.GetOk("panic_high_watermark"); ok {
		x := (v.(string))
		o.SetPanicHighWatermark(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("partial_cache_line_sparing"); ok {
		x := (v.(string))
		o.SetPartialCacheLineSparing(x)
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

	if v, ok := d.GetOk("pch_pcie_pll_ssc"); ok {
		x := (v.(string))
		o.SetPchPciePllSsc(x)
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

	if v, ok := d.GetOk("pcie_slot_mraid1link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotMraid1linkSpeed(x)
	}

	if v, ok := d.GetOk("pcie_slot_mraid1option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotMraid1optionRom(x)
	}

	if v, ok := d.GetOk("pcie_slot_mraid2link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotMraid2linkSpeed(x)
	}

	if v, ok := d.GetOk("pcie_slot_mraid2option_rom"); ok {
		x := (v.(string))
		o.SetPcieSlotMraid2optionRom(x)
	}

	if v, ok := d.GetOk("pcie_slot_mstorraid_link_speed"); ok {
		x := (v.(string))
		o.SetPcieSlotMstorraidLinkSpeed(x)
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

	if v, ok := d.GetOk("pcie_slots_cdn_enable"); ok {
		x := (v.(string))
		o.SetPcieSlotsCdnEnable(x)
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("pop_support"); ok {
		x := (v.(string))
		o.SetPopSupport(x)
	}

	if v, ok := d.GetOk("post_error_pause"); ok {
		x := (v.(string))
		o.SetPostErrorPause(x)
	}

	if v, ok := d.GetOk("post_package_repair"); ok {
		x := (v.(string))
		o.SetPostPackageRepair(x)
	}

	if v, ok := d.GetOk("pre_boot_dma_protection"); ok {
		x := (v.(string))
		o.SetPreBootDmaProtection(x)
	}

	if v, ok := d.GetOk("prmrr_size"); ok {
		x := (v.(string))
		o.SetPrmrrSize(x)
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

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		o.SetProfiles(x)
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

	if v, ok := d.GetOk("resize_bar_support"); ok {
		x := (v.(string))
		o.SetResizeBarSupport(x)
	}

	if v, ok := d.GetOk("runtime_post_package_repair"); ok {
		x := (v.(string))
		o.SetRuntimePostPackageRepair(x)
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

	if v, ok := d.GetOk("serial_mux"); ok {
		x := (v.(string))
		o.SetSerialMux(x)
	}

	if v, ok := d.GetOk("serial_port_aenable"); ok {
		x := (v.(string))
		o.SetSerialPortAenable(x)
	}

	if v, ok := d.GetOk("sev"); ok {
		x := (v.(string))
		o.SetSev(x)
	}

	if v, ok := d.GetOk("sgx_auto_registration_agent"); ok {
		x := (v.(string))
		o.SetSgxAutoRegistrationAgent(x)
	}

	if v, ok := d.GetOk("sgx_epoch0"); ok {
		x := (v.(string))
		o.SetSgxEpoch0(x)
	}

	if v, ok := d.GetOk("sgx_epoch1"); ok {
		x := (v.(string))
		o.SetSgxEpoch1(x)
	}

	if v, ok := d.GetOk("sgx_factory_reset"); ok {
		x := (v.(string))
		o.SetSgxFactoryReset(x)
	}

	if v, ok := d.GetOk("sgx_le_pub_key_hash0"); ok {
		x := (v.(string))
		o.SetSgxLePubKeyHash0(x)
	}

	if v, ok := d.GetOk("sgx_le_pub_key_hash1"); ok {
		x := (v.(string))
		o.SetSgxLePubKeyHash1(x)
	}

	if v, ok := d.GetOk("sgx_le_pub_key_hash2"); ok {
		x := (v.(string))
		o.SetSgxLePubKeyHash2(x)
	}

	if v, ok := d.GetOk("sgx_le_pub_key_hash3"); ok {
		x := (v.(string))
		o.SetSgxLePubKeyHash3(x)
	}

	if v, ok := d.GetOk("sgx_le_wr"); ok {
		x := (v.(string))
		o.SetSgxLeWr(x)
	}

	if v, ok := d.GetOk("sgx_package_info_in_band_access"); ok {
		x := (v.(string))
		o.SetSgxPackageInfoInBandAccess(x)
	}

	if v, ok := d.GetOk("sgx_qos"); ok {
		x := (v.(string))
		o.SetSgxQos(x)
	}

	if v, ok := d.GetOk("sha1pcr_bank"); ok {
		x := (v.(string))
		o.SetSha1pcrBank(x)
	}

	if v, ok := d.GetOk("sha256pcr_bank"); ok {
		x := (v.(string))
		o.SetSha256pcrBank(x)
	}

	if v, ok := d.GetOk("sha384pcr_bank"); ok {
		x := (v.(string))
		o.SetSha384pcrBank(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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

	if v, ok := d.GetOk("slot_front_nvme10link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme10linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme10option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme10optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme11link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme11linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme11option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme11optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme12link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme12linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme12option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme12optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme13link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme13linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme13option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme13optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme14link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme14linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme14option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme14optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme15link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme15linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme15option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme15optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme16link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme16linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme16option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme16optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme17link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme17linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme17option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme17optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme18link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme18linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme18option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme18optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme19link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme19linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme19option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme19optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme1linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme1option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme1optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme20link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme20linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme20option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme20optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme21link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme21linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme21option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme21optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme22link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme22linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme22option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme22optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme23link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme23linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme23option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme23optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme24link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme24linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme24option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme24optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme25link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme25linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme25option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme25optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme26link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme26linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme26option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme26optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme27link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme27linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme27option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme27optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme28link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme28linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme28option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme28optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme29link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme29linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme29option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme29optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme2linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme2option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme2optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme30link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme30linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme30option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme30optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme31link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme31linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme31option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme31optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme32link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme32linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme32option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme32optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme3link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme3linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme3option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme3optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme4link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme4linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme4option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme4optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme5link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme5linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme5option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme5optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme6link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme6linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme6option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme6optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme7link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme7linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme7option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme7optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme8link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme8linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme8option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme8optionRom(x)
	}

	if v, ok := d.GetOk("slot_front_nvme9link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme9linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_front_nvme9option_rom"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme9optionRom(x)
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

	if v, ok := d.GetOk("slot_rear_nvme3link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme3linkSpeed(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme3state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme3state(x)
	}

	if v, ok := d.GetOk("slot_rear_nvme4link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme4linkSpeed(x)
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

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["ancestor_definitions"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetAncestorDefinitions(x)
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("terminal_type"); ok {
		x := (v.(string))
		o.SetTerminalType(x)
	}

	if v, ok := d.GetOk("tpm_control"); ok {
		x := (v.(string))
		o.SetTpmControl(x)
	}

	if v, ok := d.GetOk("tpm_pending_operation"); ok {
		x := (v.(string))
		o.SetTpmPendingOperation(x)
	}

	if v, ok := d.GetOk("tpm_ppi_required"); ok {
		x := (v.(string))
		o.SetTpmPpiRequired(x)
	}

	if v, ok := d.GetOk("tpm_support"); ok {
		x := (v.(string))
		o.SetTpmSupport(x)
	}

	if v, ok := d.GetOk("tsme"); ok {
		x := (v.(string))
		o.SetTsme(x)
	}

	if v, ok := d.GetOk("txt_support"); ok {
		x := (v.(string))
		o.SetTxtSupport(x)
	}

	if v, ok := d.GetOk("ucsm_boot_order_rule"); ok {
		x := (v.(string))
		o.SetUcsmBootOrderRule(x)
	}

	if v, ok := d.GetOk("uefi_mem_map_sp_flag_en"); ok {
		x := (v.(string))
		o.SetUefiMemMapSpFlagEn(x)
	}

	if v, ok := d.GetOk("ufs_disable"); ok {
		x := (v.(string))
		o.SetUfsDisable(x)
	}

	if v, ok := d.GetOk("uma_based_clustering"); ok {
		x := (v.(string))
		o.SetUmaBasedClustering(x)
	}

	if v, ok := d.GetOk("upi_link_enablement"); ok {
		x := (v.(string))
		o.SetUpiLinkEnablement(x)
	}

	if v, ok := d.GetOk("upi_power_management"); ok {
		x := (v.(string))
		o.SetUpiPowerManagement(x)
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

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if v, ok := d.GetOk("vga_priority"); ok {
		x := (v.(string))
		o.SetVgaPriority(x)
	}

	if v, ok := d.GetOk("virtual_numa"); ok {
		x := (v.(string))
		o.SetVirtualNuma(x)
	}

	if v, ok := d.GetOk("vmd_enable"); ok {
		x := (v.(string))
		o.SetVmdEnable(x)
	}

	if v, ok := d.GetOk("vol_memory_mode"); ok {
		x := (v.(string))
		o.SetVolMemoryMode(x)
	}

	if v, ok := d.GetOk("work_load_config"); ok {
		x := (v.(string))
		o.SetWorkLoadConfig(x)
	}

	if v, ok := d.GetOk("x2apic_opt_out"); ok {
		x := (v.(string))
		o.SetX2apicOptOut(x)
	}

	if v, ok := d.GetOk("xpt_prefetch"); ok {
		x := (v.(string))
		o.SetXptPrefetch(x)
	}

	if v, ok := d.GetOk("xpt_remote_prefetch"); ok {
		x := (v.(string))
		o.SetXptRemotePrefetch(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of BiosPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.BiosApi.GetBiosPolicyList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of BiosPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of BiosPolicy: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for BiosPolicy data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var biosPolicyResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.BiosApi.GetBiosPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching BiosPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching BiosPolicy: %s", responseErr.Error())
		}
		results := resMo.BiosPolicyList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["acpi_srat_sp_flag_en"] = (s.GetAcpiSratSpFlagEn())
				temp["acs_control_gpu1state"] = (s.GetAcsControlGpu1state())
				temp["acs_control_gpu2state"] = (s.GetAcsControlGpu2state())
				temp["acs_control_gpu3state"] = (s.GetAcsControlGpu3state())
				temp["acs_control_gpu4state"] = (s.GetAcsControlGpu4state())
				temp["acs_control_gpu5state"] = (s.GetAcsControlGpu5state())
				temp["acs_control_gpu6state"] = (s.GetAcsControlGpu6state())
				temp["acs_control_gpu7state"] = (s.GetAcsControlGpu7state())
				temp["acs_control_gpu8state"] = (s.GetAcsControlGpu8state())
				temp["acs_control_slot11state"] = (s.GetAcsControlSlot11state())
				temp["acs_control_slot12state"] = (s.GetAcsControlSlot12state())
				temp["acs_control_slot13state"] = (s.GetAcsControlSlot13state())
				temp["acs_control_slot14state"] = (s.GetAcsControlSlot14state())
				temp["adaptive_refresh_mgmt_level"] = (s.GetAdaptiveRefreshMgmtLevel())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["adjacent_cache_line_prefetch"] = (s.GetAdjacentCacheLinePrefetch())
				temp["advanced_mem_test"] = (s.GetAdvancedMemTest())
				temp["all_usb_devices"] = (s.GetAllUsbDevices())
				temp["altitude"] = (s.GetAltitude())

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["aspm_support"] = (s.GetAspmSupport())
				temp["assert_nmi_on_perr"] = (s.GetAssertNmiOnPerr())
				temp["assert_nmi_on_serr"] = (s.GetAssertNmiOnSerr())
				temp["auto_cc_state"] = (s.GetAutoCcState())
				temp["autonumous_cstate_enable"] = (s.GetAutonumousCstateEnable())
				temp["baud_rate"] = (s.GetBaudRate())
				temp["bme_dma_mitigation"] = (s.GetBmeDmaMitigation())
				temp["boot_option_num_retry"] = (s.GetBootOptionNumRetry())
				temp["boot_option_re_cool_down"] = (s.GetBootOptionReCoolDown())
				temp["boot_option_retry"] = (s.GetBootOptionRetry())
				temp["boot_performance_mode"] = (s.GetBootPerformanceMode())
				temp["burst_and_postponed_refresh"] = (s.GetBurstAndPostponedRefresh())
				temp["c1auto_demotion"] = (s.GetC1autoDemotion())
				temp["c1auto_un_demotion"] = (s.GetC1autoUnDemotion())
				temp["cbs_cmn_apbdis"] = (s.GetCbsCmnApbdis())
				temp["cbs_cmn_apbdis_df_pstate_rs"] = (s.GetCbsCmnApbdisDfPstateRs())
				temp["cbs_cmn_cpu_avx512"] = (s.GetCbsCmnCpuAvx512())
				temp["cbs_cmn_cpu_cpb"] = (s.GetCbsCmnCpuCpb())
				temp["cbs_cmn_cpu_gen_downcore_ctrl"] = (s.GetCbsCmnCpuGenDowncoreCtrl())
				temp["cbs_cmn_cpu_global_cstate_ctrl"] = (s.GetCbsCmnCpuGlobalCstateCtrl())
				temp["cbs_cmn_cpu_l1stream_hw_prefetcher"] = (s.GetCbsCmnCpuL1streamHwPrefetcher())
				temp["cbs_cmn_cpu_l2stream_hw_prefetcher"] = (s.GetCbsCmnCpuL2streamHwPrefetcher())
				temp["cbs_cmn_cpu_sev_asid_space_limit"] = (s.GetCbsCmnCpuSevAsidSpaceLimit())
				temp["cbs_cmn_cpu_smee"] = (s.GetCbsCmnCpuSmee())
				temp["cbs_cmn_cpu_streaming_stores_ctrl"] = (s.GetCbsCmnCpuStreamingStoresCtrl())
				temp["cbs_cmn_determinism_slider"] = (s.GetCbsCmnDeterminismSlider())
				temp["cbs_cmn_edc_control_throttle"] = (s.GetCbsCmnEdcControlThrottle())
				temp["cbs_cmn_efficiency_mode_en"] = (s.GetCbsCmnEfficiencyModeEn())
				temp["cbs_cmn_efficiency_mode_en_rs"] = (s.GetCbsCmnEfficiencyModeEnRs())
				temp["cbs_cmn_fixed_soc_pstate"] = (s.GetCbsCmnFixedSocPstate())
				temp["cbs_cmn_gnb_nb_iommu"] = (s.GetCbsCmnGnbNbIommu())
				temp["cbs_cmn_gnb_smu_df_cstates"] = (s.GetCbsCmnGnbSmuDfCstates())
				temp["cbs_cmn_gnb_smu_dffo_rs"] = (s.GetCbsCmnGnbSmuDffoRs())
				temp["cbs_cmn_gnb_smu_dlwm_support"] = (s.GetCbsCmnGnbSmuDlwmSupport())
				temp["cbs_cmn_gnb_smucppc"] = (s.GetCbsCmnGnbSmucppc())
				temp["cbs_cmn_mem_ctrl_bank_group_swap_ddr4"] = (s.GetCbsCmnMemCtrlBankGroupSwapDdr4())
				temp["cbs_cmn_mem_ctrller_pwr_dn_en_ddr"] = (s.GetCbsCmnMemCtrllerPwrDnEnDdr())
				temp["cbs_cmn_mem_map_bank_interleave_ddr4"] = (s.GetCbsCmnMemMapBankInterleaveDdr4())
				temp["cbs_cmn_mem_speed_ddr47xx2"] = (s.GetCbsCmnMemSpeedDdr47xx2())
				temp["cbs_cmn_mem_speed_ddr47xx3"] = (s.GetCbsCmnMemSpeedDdr47xx3())
				temp["cbs_cmn_preferred_io7xx2"] = (s.GetCbsCmnPreferredIo7xx2())
				temp["cbs_cmn_preferred_io7xx3"] = (s.GetCbsCmnPreferredIo7xx3())
				temp["cbs_cmnc_tdp_ctl"] = (s.GetCbsCmncTdpCtl())
				temp["cbs_cmnx_gmi_force_link_width_rs"] = (s.GetCbsCmnxGmiForceLinkWidthRs())
				temp["cbs_cpu_ccd_ctrl_ssp"] = (s.GetCbsCpuCcdCtrlSsp())
				temp["cbs_cpu_core_ctrl"] = (s.GetCbsCpuCoreCtrl())
				temp["cbs_cpu_down_core_ctrl_bergamo"] = (s.GetCbsCpuDownCoreCtrlBergamo())
				temp["cbs_cpu_down_core_ctrl_genoa"] = (s.GetCbsCpuDownCoreCtrlGenoa())
				temp["cbs_cpu_smt_ctrl"] = (s.GetCbsCpuSmtCtrl())
				temp["cbs_dbg_cpu_gen_cpu_wdt"] = (s.GetCbsDbgCpuGenCpuWdt())
				temp["cbs_dbg_cpu_lapic_mode"] = (s.GetCbsDbgCpuLapicMode())
				temp["cbs_dbg_cpu_snp_mem_cover"] = (s.GetCbsDbgCpuSnpMemCover())
				temp["cbs_dbg_cpu_snp_mem_size_cover"] = (s.GetCbsDbgCpuSnpMemSizeCover())
				temp["cbs_df_cmn4link_max_xgmi_speed"] = (s.GetCbsDfCmn4linkMaxXgmiSpeed())
				temp["cbs_df_cmn_acpi_srat_l3numa"] = (s.GetCbsDfCmnAcpiSratL3numa())
				temp["cbs_df_cmn_dram_nps"] = (s.GetCbsDfCmnDramNps())
				temp["cbs_df_cmn_dram_scrub_time"] = (s.GetCbsDfCmnDramScrubTime())
				temp["cbs_df_cmn_mem_intlv"] = (s.GetCbsDfCmnMemIntlv())
				temp["cbs_df_cmn_mem_intlv_control"] = (s.GetCbsDfCmnMemIntlvControl())
				temp["cbs_df_cmn_mem_intlv_size"] = (s.GetCbsDfCmnMemIntlvSize())
				temp["cbs_df_dbg_xgmi_link_cfg"] = (s.GetCbsDfDbgXgmiLinkCfg())
				temp["cbs_gnb_dbg_pcie_tbt_support"] = (s.GetCbsGnbDbgPcieTbtSupport())
				temp["cbs_sev_snp_support"] = (s.GetCbsSevSnpSupport())
				temp["cdn_enable"] = (s.GetCdnEnable())
				temp["cdn_support"] = (s.GetCdnSupport())
				temp["channel_inter_leave"] = (s.GetChannelInterLeave())
				temp["cisco_adaptive_mem_training"] = (s.GetCiscoAdaptiveMemTraining())
				temp["cisco_debug_level"] = (s.GetCiscoDebugLevel())
				temp["cisco_oprom_launch_optimization"] = (s.GetCiscoOpromLaunchOptimization())
				temp["cisco_xgmi_max_speed"] = (s.GetCiscoXgmiMaxSpeed())
				temp["cke_low_policy"] = (s.GetCkeLowPolicy())
				temp["class_id"] = (s.GetClassId())
				temp["closed_loop_therm_throtl"] = (s.GetClosedLoopThermThrotl())
				temp["cmci_enable"] = (s.GetCmciEnable())
				temp["config_tdp"] = (s.GetConfigTdp())
				temp["config_tdp_level"] = (s.GetConfigTdpLevel())
				temp["console_redirection"] = (s.GetConsoleRedirection())
				temp["core_multi_processing"] = (s.GetCoreMultiProcessing())
				temp["cpu_energy_performance"] = (s.GetCpuEnergyPerformance())
				temp["cpu_frequency_floor"] = (s.GetCpuFrequencyFloor())
				temp["cpu_pa_limit"] = (s.GetCpuPaLimit())
				temp["cpu_perf_enhancement"] = (s.GetCpuPerfEnhancement())
				temp["cpu_performance"] = (s.GetCpuPerformance())
				temp["cpu_power_management"] = (s.GetCpuPowerManagement())
				temp["cr_qos"] = (s.GetCrQos())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["crfastgo_config"] = (s.GetCrfastgoConfig())
				temp["dcpmm_firmware_downgrade"] = (s.GetDcpmmFirmwareDowngrade())
				temp["demand_scrub"] = (s.GetDemandScrub())
				temp["description"] = (s.GetDescription())
				temp["dfx_osb_en"] = (s.GetDfxOsbEn())
				temp["direct_cache_access"] = (s.GetDirectCacheAccess())
				temp["dma_ctrl_opt_in"] = (s.GetDmaCtrlOptIn())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["dram_clock_throttling"] = (s.GetDramClockThrottling())
				temp["dram_refresh_rate"] = (s.GetDramRefreshRate())
				temp["dram_sw_thermal_throttling"] = (s.GetDramSwThermalThrottling())
				temp["eadr_support"] = (s.GetEadrSupport())
				temp["edpc_en"] = (s.GetEdpcEn())
				temp["enable_clock_spread_spec"] = (s.GetEnableClockSpreadSpec())
				temp["enable_mktme"] = (s.GetEnableMktme())
				temp["enable_rmt"] = (s.GetEnableRmt())
				temp["enable_sgx"] = (s.GetEnableSgx())
				temp["enable_tdx"] = (s.GetEnableTdx())
				temp["enable_tdx_seamldr"] = (s.GetEnableTdxSeamldr())
				temp["enable_tme"] = (s.GetEnableTme())
				temp["energy_efficient_turbo"] = (s.GetEnergyEfficientTurbo())
				temp["eng_perf_tuning"] = (s.GetEngPerfTuning())
				temp["enhanced_intel_speed_step_tech"] = (s.GetEnhancedIntelSpeedStepTech())
				temp["epoch_update"] = (s.GetEpochUpdate())
				temp["epp_enable"] = (s.GetEppEnable())
				temp["epp_profile"] = (s.GetEppProfile())
				temp["error_check_scrub"] = (s.GetErrorCheckScrub())
				temp["execute_disable_bit"] = (s.GetExecuteDisableBit())
				temp["extended_apic"] = (s.GetExtendedApic())
				temp["flow_control"] = (s.GetFlowControl())
				temp["frb2enable"] = (s.GetFrb2enable())
				temp["hardware_prefetch"] = (s.GetHardwarePrefetch())
				temp["hwpm_enable"] = (s.GetHwpmEnable())
				temp["imc_interleave"] = (s.GetImcInterleave())
				temp["intel_dynamic_speed_select"] = (s.GetIntelDynamicSpeedSelect())
				temp["intel_hyper_threading_tech"] = (s.GetIntelHyperThreadingTech())
				temp["intel_speed_select"] = (s.GetIntelSpeedSelect())
				temp["intel_turbo_boost_tech"] = (s.GetIntelTurboBoostTech())
				temp["intel_virtualization_technology"] = (s.GetIntelVirtualizationTechnology())
				temp["intel_vt_for_directed_io"] = (s.GetIntelVtForDirectedIo())
				temp["intel_vtd_coherency_support"] = (s.GetIntelVtdCoherencySupport())
				temp["intel_vtd_interrupt_remapping"] = (s.GetIntelVtdInterruptRemapping())
				temp["intel_vtd_pass_through_dma_support"] = (s.GetIntelVtdPassThroughDmaSupport())
				temp["intel_vtdats_support"] = (s.GetIntelVtdatsSupport())
				temp["ioat_config_cpm"] = (s.GetIoatConfigCpm())
				temp["ioh_error_enable"] = (s.GetIohErrorEnable())
				temp["ioh_resource"] = (s.GetIohResource())
				temp["ip_prefetch"] = (s.GetIpPrefetch())
				temp["ipv4http"] = (s.GetIpv4http())
				temp["ipv4pxe"] = (s.GetIpv4pxe())
				temp["ipv6http"] = (s.GetIpv6http())
				temp["ipv6pxe"] = (s.GetIpv6pxe())
				temp["kti_prefetch"] = (s.GetKtiPrefetch())
				temp["latency_optimized_mode"] = (s.GetLatencyOptimizedMode())
				temp["legacy_os_redirection"] = (s.GetLegacyOsRedirection())
				temp["legacy_usb_support"] = (s.GetLegacyUsbSupport())
				temp["llc_alloc"] = (s.GetLlcAlloc())
				temp["llc_prefetch"] = (s.GetLlcPrefetch())
				temp["lom_port0state"] = (s.GetLomPort0state())
				temp["lom_port1state"] = (s.GetLomPort1state())
				temp["lom_port2state"] = (s.GetLomPort2state())
				temp["lom_port3state"] = (s.GetLomPort3state())
				temp["lom_ports_all_state"] = (s.GetLomPortsAllState())
				temp["lv_ddr_mode"] = (s.GetLvDdrMode())
				temp["make_device_non_bootable"] = (s.GetMakeDeviceNonBootable())
				temp["memory_bandwidth_boost"] = (s.GetMemoryBandwidthBoost())
				temp["memory_inter_leave"] = (s.GetMemoryInterLeave())
				temp["memory_mapped_io_above4gb"] = (s.GetMemoryMappedIoAbove4gb())
				temp["memory_refresh_rate"] = (s.GetMemoryRefreshRate())
				temp["memory_size_limit"] = (s.GetMemorySizeLimit())
				temp["memory_thermal_throttling"] = (s.GetMemoryThermalThrottling())
				temp["mirroring_mode"] = (s.GetMirroringMode())
				temp["mmcfg_base"] = (s.GetMmcfgBase())
				temp["mmioh_base"] = (s.GetMmiohBase())
				temp["mmioh_size"] = (s.GetMmiohSize())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["network_stack"] = (s.GetNetworkStack())
				temp["numa_optimized"] = (s.GetNumaOptimized())
				temp["nvmdimm_perform_config"] = (s.GetNvmdimmPerformConfig())
				temp["object_type"] = (s.GetObjectType())
				temp["onboard10gbit_lom"] = (s.GetOnboard10gbitLom())
				temp["onboard_gbit_lom"] = (s.GetOnboardGbitLom())
				temp["onboard_scu_storage_support"] = (s.GetOnboardScuStorageSupport())
				temp["onboard_scu_storage_sw_stack"] = (s.GetOnboardScuStorageSwStack())
				temp["operation_mode"] = (s.GetOperationMode())
				temp["optimized_power_mode"] = (s.GetOptimizedPowerMode())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["os_boot_watchdog_timer"] = (s.GetOsBootWatchdogTimer())
				temp["os_boot_watchdog_timer_policy"] = (s.GetOsBootWatchdogTimerPolicy())
				temp["os_boot_watchdog_timer_timeout"] = (s.GetOsBootWatchdogTimerTimeout())
				temp["out_of_band_mgmt_port"] = (s.GetOutOfBandMgmtPort())
				temp["owners"] = (s.GetOwners())
				temp["package_cstate_limit"] = (s.GetPackageCstateLimit())
				temp["panic_high_watermark"] = (s.GetPanicHighWatermark())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)
				temp["partial_cache_line_sparing"] = (s.GetPartialCacheLineSparing())
				temp["partial_mirror_mode_config"] = (s.GetPartialMirrorModeConfig())
				temp["partial_mirror_percent"] = (s.GetPartialMirrorPercent())
				temp["partial_mirror_value1"] = (s.GetPartialMirrorValue1())
				temp["partial_mirror_value2"] = (s.GetPartialMirrorValue2())
				temp["partial_mirror_value3"] = (s.GetPartialMirrorValue3())
				temp["partial_mirror_value4"] = (s.GetPartialMirrorValue4())
				temp["patrol_scrub"] = (s.GetPatrolScrub())
				temp["patrol_scrub_duration"] = (s.GetPatrolScrubDuration())
				temp["pc_ie_ras_support"] = (s.GetPcIeRasSupport())
				temp["pc_ie_ssd_hot_plug_support"] = (s.GetPcIeSsdHotPlugSupport())
				temp["pch_pcie_pll_ssc"] = (s.GetPchPciePllSsc())
				temp["pch_usb30mode"] = (s.GetPchUsb30mode())
				temp["pci_option_ro_ms"] = (s.GetPciOptionRoMs())
				temp["pci_rom_clp"] = (s.GetPciRomClp())
				temp["pcie_ari_support"] = (s.GetPcieAriSupport())
				temp["pcie_pll_ssc"] = (s.GetPciePllSsc())
				temp["pcie_slot_mraid1link_speed"] = (s.GetPcieSlotMraid1linkSpeed())
				temp["pcie_slot_mraid1option_rom"] = (s.GetPcieSlotMraid1optionRom())
				temp["pcie_slot_mraid2link_speed"] = (s.GetPcieSlotMraid2linkSpeed())
				temp["pcie_slot_mraid2option_rom"] = (s.GetPcieSlotMraid2optionRom())
				temp["pcie_slot_mstorraid_link_speed"] = (s.GetPcieSlotMstorraidLinkSpeed())
				temp["pcie_slot_mstorraid_option_rom"] = (s.GetPcieSlotMstorraidOptionRom())
				temp["pcie_slot_nvme1link_speed"] = (s.GetPcieSlotNvme1linkSpeed())
				temp["pcie_slot_nvme1option_rom"] = (s.GetPcieSlotNvme1optionRom())
				temp["pcie_slot_nvme2link_speed"] = (s.GetPcieSlotNvme2linkSpeed())
				temp["pcie_slot_nvme2option_rom"] = (s.GetPcieSlotNvme2optionRom())
				temp["pcie_slot_nvme3link_speed"] = (s.GetPcieSlotNvme3linkSpeed())
				temp["pcie_slot_nvme3option_rom"] = (s.GetPcieSlotNvme3optionRom())
				temp["pcie_slot_nvme4link_speed"] = (s.GetPcieSlotNvme4linkSpeed())
				temp["pcie_slot_nvme4option_rom"] = (s.GetPcieSlotNvme4optionRom())
				temp["pcie_slot_nvme5link_speed"] = (s.GetPcieSlotNvme5linkSpeed())
				temp["pcie_slot_nvme5option_rom"] = (s.GetPcieSlotNvme5optionRom())
				temp["pcie_slot_nvme6link_speed"] = (s.GetPcieSlotNvme6linkSpeed())
				temp["pcie_slot_nvme6option_rom"] = (s.GetPcieSlotNvme6optionRom())
				temp["pcie_slots_cdn_enable"] = (s.GetPcieSlotsCdnEnable())

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["pop_support"] = (s.GetPopSupport())
				temp["post_error_pause"] = (s.GetPostErrorPause())
				temp["post_package_repair"] = (s.GetPostPackageRepair())
				temp["pre_boot_dma_protection"] = (s.GetPreBootDmaProtection())
				temp["prmrr_size"] = (s.GetPrmrrSize())
				temp["processor_c1e"] = (s.GetProcessorC1e())
				temp["processor_c3report"] = (s.GetProcessorC3report())
				temp["processor_c6report"] = (s.GetProcessorC6report())
				temp["processor_cstate"] = (s.GetProcessorCstate())

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["psata"] = (s.GetPsata())
				temp["pstate_coord_type"] = (s.GetPstateCoordType())
				temp["putty_key_pad"] = (s.GetPuttyKeyPad())
				temp["pwr_perf_tuning"] = (s.GetPwrPerfTuning())
				temp["qpi_link_frequency"] = (s.GetQpiLinkFrequency())
				temp["qpi_link_speed"] = (s.GetQpiLinkSpeed())
				temp["qpi_snoop_mode"] = (s.GetQpiSnoopMode())
				temp["rank_inter_leave"] = (s.GetRankInterLeave())
				temp["redirection_after_post"] = (s.GetRedirectionAfterPost())
				temp["resize_bar_support"] = (s.GetResizeBarSupport())
				temp["runtime_post_package_repair"] = (s.GetRuntimePostPackageRepair())
				temp["sata_mode_select"] = (s.GetSataModeSelect())
				temp["select_memory_ras_configuration"] = (s.GetSelectMemoryRasConfiguration())
				temp["select_ppr_type"] = (s.GetSelectPprType())
				temp["serial_mux"] = (s.GetSerialMux())
				temp["serial_port_aenable"] = (s.GetSerialPortAenable())
				temp["sev"] = (s.GetSev())
				temp["sgx_auto_registration_agent"] = (s.GetSgxAutoRegistrationAgent())
				temp["sgx_epoch0"] = (s.GetSgxEpoch0())
				temp["sgx_epoch1"] = (s.GetSgxEpoch1())
				temp["sgx_factory_reset"] = (s.GetSgxFactoryReset())
				temp["sgx_le_pub_key_hash0"] = (s.GetSgxLePubKeyHash0())
				temp["sgx_le_pub_key_hash1"] = (s.GetSgxLePubKeyHash1())
				temp["sgx_le_pub_key_hash2"] = (s.GetSgxLePubKeyHash2())
				temp["sgx_le_pub_key_hash3"] = (s.GetSgxLePubKeyHash3())
				temp["sgx_le_wr"] = (s.GetSgxLeWr())
				temp["sgx_package_info_in_band_access"] = (s.GetSgxPackageInfoInBandAccess())
				temp["sgx_qos"] = (s.GetSgxQos())
				temp["sha1pcr_bank"] = (s.GetSha1pcrBank())
				temp["sha256pcr_bank"] = (s.GetSha256pcrBank())
				temp["sha384pcr_bank"] = (s.GetSha384pcrBank())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["single_pctl_enable"] = (s.GetSinglePctlEnable())
				temp["slot10link_speed"] = (s.GetSlot10linkSpeed())
				temp["slot10state"] = (s.GetSlot10state())
				temp["slot11link_speed"] = (s.GetSlot11linkSpeed())
				temp["slot11state"] = (s.GetSlot11state())
				temp["slot12link_speed"] = (s.GetSlot12linkSpeed())
				temp["slot12state"] = (s.GetSlot12state())
				temp["slot13state"] = (s.GetSlot13state())
				temp["slot14state"] = (s.GetSlot14state())
				temp["slot1link_speed"] = (s.GetSlot1linkSpeed())
				temp["slot1state"] = (s.GetSlot1state())
				temp["slot2link_speed"] = (s.GetSlot2linkSpeed())
				temp["slot2state"] = (s.GetSlot2state())
				temp["slot3link_speed"] = (s.GetSlot3linkSpeed())
				temp["slot3state"] = (s.GetSlot3state())
				temp["slot4link_speed"] = (s.GetSlot4linkSpeed())
				temp["slot4state"] = (s.GetSlot4state())
				temp["slot5link_speed"] = (s.GetSlot5linkSpeed())
				temp["slot5state"] = (s.GetSlot5state())
				temp["slot6link_speed"] = (s.GetSlot6linkSpeed())
				temp["slot6state"] = (s.GetSlot6state())
				temp["slot7link_speed"] = (s.GetSlot7linkSpeed())
				temp["slot7state"] = (s.GetSlot7state())
				temp["slot8link_speed"] = (s.GetSlot8linkSpeed())
				temp["slot8state"] = (s.GetSlot8state())
				temp["slot9link_speed"] = (s.GetSlot9linkSpeed())
				temp["slot9state"] = (s.GetSlot9state())
				temp["slot_flom_link_speed"] = (s.GetSlotFlomLinkSpeed())
				temp["slot_front_nvme10link_speed"] = (s.GetSlotFrontNvme10linkSpeed())
				temp["slot_front_nvme10option_rom"] = (s.GetSlotFrontNvme10optionRom())
				temp["slot_front_nvme11link_speed"] = (s.GetSlotFrontNvme11linkSpeed())
				temp["slot_front_nvme11option_rom"] = (s.GetSlotFrontNvme11optionRom())
				temp["slot_front_nvme12link_speed"] = (s.GetSlotFrontNvme12linkSpeed())
				temp["slot_front_nvme12option_rom"] = (s.GetSlotFrontNvme12optionRom())
				temp["slot_front_nvme13link_speed"] = (s.GetSlotFrontNvme13linkSpeed())
				temp["slot_front_nvme13option_rom"] = (s.GetSlotFrontNvme13optionRom())
				temp["slot_front_nvme14link_speed"] = (s.GetSlotFrontNvme14linkSpeed())
				temp["slot_front_nvme14option_rom"] = (s.GetSlotFrontNvme14optionRom())
				temp["slot_front_nvme15link_speed"] = (s.GetSlotFrontNvme15linkSpeed())
				temp["slot_front_nvme15option_rom"] = (s.GetSlotFrontNvme15optionRom())
				temp["slot_front_nvme16link_speed"] = (s.GetSlotFrontNvme16linkSpeed())
				temp["slot_front_nvme16option_rom"] = (s.GetSlotFrontNvme16optionRom())
				temp["slot_front_nvme17link_speed"] = (s.GetSlotFrontNvme17linkSpeed())
				temp["slot_front_nvme17option_rom"] = (s.GetSlotFrontNvme17optionRom())
				temp["slot_front_nvme18link_speed"] = (s.GetSlotFrontNvme18linkSpeed())
				temp["slot_front_nvme18option_rom"] = (s.GetSlotFrontNvme18optionRom())
				temp["slot_front_nvme19link_speed"] = (s.GetSlotFrontNvme19linkSpeed())
				temp["slot_front_nvme19option_rom"] = (s.GetSlotFrontNvme19optionRom())
				temp["slot_front_nvme1link_speed"] = (s.GetSlotFrontNvme1linkSpeed())
				temp["slot_front_nvme1option_rom"] = (s.GetSlotFrontNvme1optionRom())
				temp["slot_front_nvme20link_speed"] = (s.GetSlotFrontNvme20linkSpeed())
				temp["slot_front_nvme20option_rom"] = (s.GetSlotFrontNvme20optionRom())
				temp["slot_front_nvme21link_speed"] = (s.GetSlotFrontNvme21linkSpeed())
				temp["slot_front_nvme21option_rom"] = (s.GetSlotFrontNvme21optionRom())
				temp["slot_front_nvme22link_speed"] = (s.GetSlotFrontNvme22linkSpeed())
				temp["slot_front_nvme22option_rom"] = (s.GetSlotFrontNvme22optionRom())
				temp["slot_front_nvme23link_speed"] = (s.GetSlotFrontNvme23linkSpeed())
				temp["slot_front_nvme23option_rom"] = (s.GetSlotFrontNvme23optionRom())
				temp["slot_front_nvme24link_speed"] = (s.GetSlotFrontNvme24linkSpeed())
				temp["slot_front_nvme24option_rom"] = (s.GetSlotFrontNvme24optionRom())
				temp["slot_front_nvme25link_speed"] = (s.GetSlotFrontNvme25linkSpeed())
				temp["slot_front_nvme25option_rom"] = (s.GetSlotFrontNvme25optionRom())
				temp["slot_front_nvme26link_speed"] = (s.GetSlotFrontNvme26linkSpeed())
				temp["slot_front_nvme26option_rom"] = (s.GetSlotFrontNvme26optionRom())
				temp["slot_front_nvme27link_speed"] = (s.GetSlotFrontNvme27linkSpeed())
				temp["slot_front_nvme27option_rom"] = (s.GetSlotFrontNvme27optionRom())
				temp["slot_front_nvme28link_speed"] = (s.GetSlotFrontNvme28linkSpeed())
				temp["slot_front_nvme28option_rom"] = (s.GetSlotFrontNvme28optionRom())
				temp["slot_front_nvme29link_speed"] = (s.GetSlotFrontNvme29linkSpeed())
				temp["slot_front_nvme29option_rom"] = (s.GetSlotFrontNvme29optionRom())
				temp["slot_front_nvme2link_speed"] = (s.GetSlotFrontNvme2linkSpeed())
				temp["slot_front_nvme2option_rom"] = (s.GetSlotFrontNvme2optionRom())
				temp["slot_front_nvme30link_speed"] = (s.GetSlotFrontNvme30linkSpeed())
				temp["slot_front_nvme30option_rom"] = (s.GetSlotFrontNvme30optionRom())
				temp["slot_front_nvme31link_speed"] = (s.GetSlotFrontNvme31linkSpeed())
				temp["slot_front_nvme31option_rom"] = (s.GetSlotFrontNvme31optionRom())
				temp["slot_front_nvme32link_speed"] = (s.GetSlotFrontNvme32linkSpeed())
				temp["slot_front_nvme32option_rom"] = (s.GetSlotFrontNvme32optionRom())
				temp["slot_front_nvme3link_speed"] = (s.GetSlotFrontNvme3linkSpeed())
				temp["slot_front_nvme3option_rom"] = (s.GetSlotFrontNvme3optionRom())
				temp["slot_front_nvme4link_speed"] = (s.GetSlotFrontNvme4linkSpeed())
				temp["slot_front_nvme4option_rom"] = (s.GetSlotFrontNvme4optionRom())
				temp["slot_front_nvme5link_speed"] = (s.GetSlotFrontNvme5linkSpeed())
				temp["slot_front_nvme5option_rom"] = (s.GetSlotFrontNvme5optionRom())
				temp["slot_front_nvme6link_speed"] = (s.GetSlotFrontNvme6linkSpeed())
				temp["slot_front_nvme6option_rom"] = (s.GetSlotFrontNvme6optionRom())
				temp["slot_front_nvme7link_speed"] = (s.GetSlotFrontNvme7linkSpeed())
				temp["slot_front_nvme7option_rom"] = (s.GetSlotFrontNvme7optionRom())
				temp["slot_front_nvme8link_speed"] = (s.GetSlotFrontNvme8linkSpeed())
				temp["slot_front_nvme8option_rom"] = (s.GetSlotFrontNvme8optionRom())
				temp["slot_front_nvme9link_speed"] = (s.GetSlotFrontNvme9linkSpeed())
				temp["slot_front_nvme9option_rom"] = (s.GetSlotFrontNvme9optionRom())
				temp["slot_front_slot5link_speed"] = (s.GetSlotFrontSlot5linkSpeed())
				temp["slot_front_slot6link_speed"] = (s.GetSlotFrontSlot6linkSpeed())
				temp["slot_gpu1state"] = (s.GetSlotGpu1state())
				temp["slot_gpu2state"] = (s.GetSlotGpu2state())
				temp["slot_gpu3state"] = (s.GetSlotGpu3state())
				temp["slot_gpu4state"] = (s.GetSlotGpu4state())
				temp["slot_gpu5state"] = (s.GetSlotGpu5state())
				temp["slot_gpu6state"] = (s.GetSlotGpu6state())
				temp["slot_gpu7state"] = (s.GetSlotGpu7state())
				temp["slot_gpu8state"] = (s.GetSlotGpu8state())
				temp["slot_hba_link_speed"] = (s.GetSlotHbaLinkSpeed())
				temp["slot_hba_state"] = (s.GetSlotHbaState())
				temp["slot_lom1link"] = (s.GetSlotLom1link())
				temp["slot_lom2link"] = (s.GetSlotLom2link())
				temp["slot_mezz_state"] = (s.GetSlotMezzState())
				temp["slot_mlom_link_speed"] = (s.GetSlotMlomLinkSpeed())
				temp["slot_mlom_state"] = (s.GetSlotMlomState())
				temp["slot_mraid_link_speed"] = (s.GetSlotMraidLinkSpeed())
				temp["slot_mraid_state"] = (s.GetSlotMraidState())
				temp["slot_n10state"] = (s.GetSlotN10state())
				temp["slot_n11state"] = (s.GetSlotN11state())
				temp["slot_n12state"] = (s.GetSlotN12state())
				temp["slot_n13state"] = (s.GetSlotN13state())
				temp["slot_n14state"] = (s.GetSlotN14state())
				temp["slot_n15state"] = (s.GetSlotN15state())
				temp["slot_n16state"] = (s.GetSlotN16state())
				temp["slot_n17state"] = (s.GetSlotN17state())
				temp["slot_n18state"] = (s.GetSlotN18state())
				temp["slot_n19state"] = (s.GetSlotN19state())
				temp["slot_n1state"] = (s.GetSlotN1state())
				temp["slot_n20state"] = (s.GetSlotN20state())
				temp["slot_n21state"] = (s.GetSlotN21state())
				temp["slot_n22state"] = (s.GetSlotN22state())
				temp["slot_n23state"] = (s.GetSlotN23state())
				temp["slot_n24state"] = (s.GetSlotN24state())
				temp["slot_n2state"] = (s.GetSlotN2state())
				temp["slot_n3state"] = (s.GetSlotN3state())
				temp["slot_n4state"] = (s.GetSlotN4state())
				temp["slot_n5state"] = (s.GetSlotN5state())
				temp["slot_n6state"] = (s.GetSlotN6state())
				temp["slot_n7state"] = (s.GetSlotN7state())
				temp["slot_n8state"] = (s.GetSlotN8state())
				temp["slot_n9state"] = (s.GetSlotN9state())
				temp["slot_raid_link_speed"] = (s.GetSlotRaidLinkSpeed())
				temp["slot_raid_state"] = (s.GetSlotRaidState())
				temp["slot_rear_nvme1link_speed"] = (s.GetSlotRearNvme1linkSpeed())
				temp["slot_rear_nvme1state"] = (s.GetSlotRearNvme1state())
				temp["slot_rear_nvme2link_speed"] = (s.GetSlotRearNvme2linkSpeed())
				temp["slot_rear_nvme2state"] = (s.GetSlotRearNvme2state())
				temp["slot_rear_nvme3link_speed"] = (s.GetSlotRearNvme3linkSpeed())
				temp["slot_rear_nvme3state"] = (s.GetSlotRearNvme3state())
				temp["slot_rear_nvme4link_speed"] = (s.GetSlotRearNvme4linkSpeed())
				temp["slot_rear_nvme4state"] = (s.GetSlotRearNvme4state())
				temp["slot_rear_nvme5state"] = (s.GetSlotRearNvme5state())
				temp["slot_rear_nvme6state"] = (s.GetSlotRearNvme6state())
				temp["slot_rear_nvme7state"] = (s.GetSlotRearNvme7state())
				temp["slot_rear_nvme8state"] = (s.GetSlotRearNvme8state())
				temp["slot_riser1link_speed"] = (s.GetSlotRiser1linkSpeed())
				temp["slot_riser1slot1link_speed"] = (s.GetSlotRiser1slot1linkSpeed())
				temp["slot_riser1slot2link_speed"] = (s.GetSlotRiser1slot2linkSpeed())
				temp["slot_riser1slot3link_speed"] = (s.GetSlotRiser1slot3linkSpeed())
				temp["slot_riser2link_speed"] = (s.GetSlotRiser2linkSpeed())
				temp["slot_riser2slot4link_speed"] = (s.GetSlotRiser2slot4linkSpeed())
				temp["slot_riser2slot5link_speed"] = (s.GetSlotRiser2slot5linkSpeed())
				temp["slot_riser2slot6link_speed"] = (s.GetSlotRiser2slot6linkSpeed())
				temp["slot_sas_state"] = (s.GetSlotSasState())
				temp["slot_ssd_slot1link_speed"] = (s.GetSlotSsdSlot1linkSpeed())
				temp["slot_ssd_slot2link_speed"] = (s.GetSlotSsdSlot2linkSpeed())
				temp["smee"] = (s.GetSmee())
				temp["smt_mode"] = (s.GetSmtMode())
				temp["snc"] = (s.GetSnc())
				temp["snoopy_mode_for2lm"] = (s.GetSnoopyModeFor2lm())
				temp["snoopy_mode_for_ad"] = (s.GetSnoopyModeForAd())
				temp["sparing_mode"] = (s.GetSparingMode())
				temp["sr_iov"] = (s.GetSrIov())
				temp["streamer_prefetch"] = (s.GetStreamerPrefetch())
				temp["svm_mode"] = (s.GetSvmMode())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["terminal_type"] = (s.GetTerminalType())
				temp["tpm_control"] = (s.GetTpmControl())
				temp["tpm_pending_operation"] = (s.GetTpmPendingOperation())
				temp["tpm_ppi_required"] = (s.GetTpmPpiRequired())
				temp["tpm_support"] = (s.GetTpmSupport())
				temp["tsme"] = (s.GetTsme())
				temp["txt_support"] = (s.GetTxtSupport())
				temp["ucsm_boot_order_rule"] = (s.GetUcsmBootOrderRule())
				temp["uefi_mem_map_sp_flag_en"] = (s.GetUefiMemMapSpFlagEn())
				temp["ufs_disable"] = (s.GetUfsDisable())
				temp["uma_based_clustering"] = (s.GetUmaBasedClustering())
				temp["upi_link_enablement"] = (s.GetUpiLinkEnablement())
				temp["upi_power_management"] = (s.GetUpiPowerManagement())
				temp["usb_emul6064"] = (s.GetUsbEmul6064())
				temp["usb_port_front"] = (s.GetUsbPortFront())
				temp["usb_port_internal"] = (s.GetUsbPortInternal())
				temp["usb_port_kvm"] = (s.GetUsbPortKvm())
				temp["usb_port_rear"] = (s.GetUsbPortRear())
				temp["usb_port_sd_card"] = (s.GetUsbPortSdCard())
				temp["usb_port_vmedia"] = (s.GetUsbPortVmedia())
				temp["usb_xhci_support"] = (s.GetUsbXhciSupport())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vga_priority"] = (s.GetVgaPriority())
				temp["virtual_numa"] = (s.GetVirtualNuma())
				temp["vmd_enable"] = (s.GetVmdEnable())
				temp["vol_memory_mode"] = (s.GetVolMemoryMode())
				temp["work_load_config"] = (s.GetWorkLoadConfig())
				temp["x2apic_opt_out"] = (s.GetX2apicOptOut())
				temp["xpt_prefetch"] = (s.GetXptPrefetch())
				temp["xpt_remote_prefetch"] = (s.GetXptRemotePrefetch())
				biosPolicyResults = append(biosPolicyResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(biosPolicyResults))
	if err := d.Set("results", biosPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(biosPolicyResults[0]["moid"].(string))
	return de
}
