package victronDefinitions

type VeProduct uint16

const (
	VeProductBMV700                                 VeProduct = 0x203
	VeProductBMV702                                 VeProduct = 0x204
	VeProductBMV700H                                VeProduct = 0x205
	VeProductBlueSolarMPPT70_15                     VeProduct = 0x0300
	VeProductBlueSolarMPPT75_50                     VeProduct = 0xA040
	VeProductBlueSolarMPPT150_35                    VeProduct = 0xA041
	VeProductBlueSolarMPPT75_15                     VeProduct = 0xA042
	VeProductBlueSolarMPPT100_15                    VeProduct = 0xA043
	VeProductBlueSolarMPPT100_30                    VeProduct = 0xA044
	VeProductBlueSolarMPPT100_50                    VeProduct = 0xA045
	VeProductBlueSolarMPPT150_70                    VeProduct = 0xA046
	VeProductBlueSolarMPPT150_100                   VeProduct = 0xA047
	VeProductBlueSolarMPPT100_50rev2                VeProduct = 0xA049
	VeProductBlueSolarMPPT100_30rev2                VeProduct = 0xA04A
	VeProductBlueSolarMPPT150_35rev2                VeProduct = 0xA04B
	VeProductBlueSolarMPPT75_10                     VeProduct = 0xA04C
	VeProductBlueSolarMPPT150_45                    VeProduct = 0xA04D
	VeProductBlueSolarMPPT150_60                    VeProduct = 0xA04E
	VeProductBlueSolarMPPT150_85                    VeProduct = 0xA04F
	VeProductSmartSolarMPPT250_100                  VeProduct = 0xA050
	VeProductSmartSolarMPPT150_100                  VeProduct = 0xA051
	VeProductSmartSolarMPPT150_85                   VeProduct = 0xA052
	VeProductSmartSolarMPPT75_15                    VeProduct = 0xA053
	VeProductSmartSolarMPPT75_10                    VeProduct = 0xA054
	VeProductSmartSolarMPPT100_15                   VeProduct = 0xA055
	VeProductSmartSolarMPPT100_30                   VeProduct = 0xA056
	VeProductSmartSolarMPPT100_50                   VeProduct = 0xA057
	VeProductSmartSolarMPPT150_35                   VeProduct = 0xA058
	VeProductSmartSolarMPPT150_100rev2              VeProduct = 0xA059
	VeProductSmartSolarMPPT150_85rev2               VeProduct = 0xA05A
	VeProductSmartSolarMPPT250_70                   VeProduct = 0xA05B
	VeProductSmartSolarMPPT250_85                   VeProduct = 0xA05C
	VeProductSmartSolarMPPT250_60                   VeProduct = 0xA05D
	VeProductSmartSolarMPPT250_45                   VeProduct = 0xA05E
	VeProductSmartSolarMPPT100_20                   VeProduct = 0xA05F
	VeProductSmartSolarMPPT100_2048V                VeProduct = 0xA060
	VeProductSmartSolarMPPT150_45                   VeProduct = 0xA061
	VeProductSmartSolarMPPT150_60                   VeProduct = 0xA062
	VeProductSmartSolarMPPT150_70                   VeProduct = 0xA063
	VeProductSmartSolarMPPT250_85rev2               VeProduct = 0xA064
	VeProductSmartSolarMPPT250_100rev2              VeProduct = 0xA065
	VeProductBlueSolarMPPT100_20                    VeProduct = 0xA066
	VeProductBlueSolarMPPT100_2048V                 VeProduct = 0xA067
	VeProductSmartSolarMPPT250_60rev2               VeProduct = 0xA068
	VeProductSmartSolarMPPT250_70rev2               VeProduct = 0xA069
	VeProductSmartSolarMPPT150_45rev2               VeProduct = 0xA06A
	VeProductSmartSolarMPPT150_60rev2               VeProduct = 0xA06B
	VeProductSmartSolarMPPT150_70rev2               VeProduct = 0xA06C
	VeProductSmartSolarMPPT150_85rev3               VeProduct = 0xA06D
	VeProductSmartSolarMPPT150_100rev3              VeProduct = 0xA06E
	VeProductBlueSolarMPPT150_45rev2                VeProduct = 0xA06F
	VeProductBlueSolarMPPT150_60rev2                VeProduct = 0xA070
	VeProductBlueSolarMPPT150_70rev2                VeProduct = 0xA071
	VeProductBlueSolarMPPT150_45rev3                VeProduct = 0xA072
	VeProductSmatrtSolarMPPT150_45rev3              VeProduct = 0xA073
	VeProductSmartSolarMPPT70_10rev2                VeProduct = 0xA074
	VeProductSmartSolarMPPT75_15rev2                VeProduct = 0xA075
	VeProductBlueSolarMPPT100_30rev3                VeProduct = 0xA076
	VeProductBlueSolarMPPT100_50rev3                VeProduct = 0xA077
	VeProductBlueSolarMPPT150_35rev3                VeProduct = 0xA078
	VeProductBlueSolarMPPT75_10rev2                 VeProduct = 0xA079
	VeProductBlueSolarMPPT75_15rev2                 VeProduct = 0xA07A
	VeProductBlueSolarMPPT100_15rev2                VeProduct = 0xA07B
	VeProductSmartSolarMPPTVECan150_70              VeProduct = 0xA102
	VeProductSmartSolarMPPTVECan150_45              VeProduct = 0xA103
	VeProductSmartSolarMPPTVECan150_60              VeProduct = 0xA104
	VeProductSmartSolarMPPTVECan150_85              VeProduct = 0xA105
	VeProductSmartSolarMPPTVECan150_100             VeProduct = 0xA106
	VeProductSmartSolarMPPTVECan250_45              VeProduct = 0xA107
	VeProductSmartSolarMPPTVECan250_60              VeProduct = 0xA108
	VeProductSmartSolarMPPTVECan250_70              VeProduct = 0xA109
	VeProductSmartSolarMPPTVECan250_85              VeProduct = 0xA10A
	VeProductSmartSolarMPPTVECan250_100             VeProduct = 0xA10B
	VeProductSmartSolarMPPTVECan150_70rev2          VeProduct = 0xA10C
	VeProductSmartSolarMPPTVECan150_85rev2          VeProduct = 0xA10D
	VeProductSmartSolarMPPTVECan150_100rev2         VeProduct = 0xA10E
	VeProductBlueSolarMPPTVECan150_100              VeProduct = 0xA10F
	VeProductBlueSolarMPPTVECan250_70               VeProduct = 0xA112
	VeProductBlueSolarMPPTVECan250_100              VeProduct = 0xA113
	VeProductSmartSolarMPPTVECan250_70rev2          VeProduct = 0xA114
	VeProductSmartSolarMPPTVECan250_100rev2         VeProduct = 0xA115
	VeProductSmartSolarMPPTVECan250_85rev2          VeProduct = 0xA116
	VeProductBlueSolarMPPTVECan150_100rev2          VeProduct = 0xA117
	VeProductPhoenixInverter12V250VA230V            VeProduct = 0xA231
	VeProductPhoenixInverter24V250VA230V            VeProduct = 0xA232
	VeProductPhoenixInverter48V250VA230V            VeProduct = 0xA234
	VeProductPhoenixInverter12V250VA120V            VeProduct = 0xA239
	VeProductPhoenixInverter24V250VA120V            VeProduct = 0xA23A
	VeProductPhoenixInverter48V250VA120V            VeProduct = 0xA23C
	VeProductPhoenixInverter12V375VA230V            VeProduct = 0xA241
	VeProductPhoenixInverter24V375VA230V            VeProduct = 0xA242
	VeProductPhoenixInverter48V375VA230V            VeProduct = 0xA244
	VeProductPhoenixInverter12V375VA120V            VeProduct = 0xA249
	VeProductPhoenixInverter24V375VA120V            VeProduct = 0xA24A
	VeProductPhoenixInverter48V375VA120V            VeProduct = 0xA24C
	VeProductPhoenixInverter12V500VA230V            VeProduct = 0xA251
	VeProductPhoenixInverter24V500VA230V            VeProduct = 0xA252
	VeProductPhoenixInverter48V500VA230V            VeProduct = 0xA254
	VeProductPhoenixInverter12V500VA120V            VeProduct = 0xA259
	VeProductPhoenixInverter24V500VA120V            VeProduct = 0xA25A
	VeProductPhoenixInverter48V500VA120V            VeProduct = 0xA25C
	VeProductPhoenixInverter12V800VA230V            VeProduct = 0xA261
	VeProductPhoenixInverter24V800VA230V            VeProduct = 0xA262
	VeProductPhoenixInverter48V800VA230V            VeProduct = 0xA264
	VeProductPhoenixInverter12V800VA120V            VeProduct = 0xA269
	VeProductPhoenixInverter24V800VA120V            VeProduct = 0xA26A
	VeProductPhoenixInverter48V800VA120V            VeProduct = 0xA26C
	VeProductPhoenixInverter12V1200VA230V           VeProduct = 0xA271
	VeProductPhoenixInverter24V1200VA230V           VeProduct = 0xA272
	VeProductPhoenixInverter48V1200VA230V           VeProduct = 0xA274
	VeProductPhoenixInverter12V1200VA120V           VeProduct = 0xA279
	VeProductPhoenixInverter24V1200VA120V           VeProduct = 0xA27A
	VeProductPhoenixInverter48V1200VA120V           VeProduct = 0xA27C
	VeProductPhoenixInverter12V1600VA230V           VeProduct = 0xA281
	VeProductPhoenixInverter24V1600VA230V           VeProduct = 0xA282
	VeProductPhoenixInverter48V1600VA230V           VeProduct = 0xA284
	VeProductPhoenixInverter12V2000VA230V           VeProduct = 0xA291
	VeProductPhoenixInverter24V2000VA230V           VeProduct = 0xA292
	VeProductPhoenixInverter48V2000VA230V           VeProduct = 0xA294
	VeProductPhoenixInverter12V3000VA230V           VeProduct = 0xA2A1
	VeProductPhoenixInverter24V3000VA230V           VeProduct = 0xA2A2
	VeProductPhoenixInverter48V3000VA230V           VeProduct = 0xA2A4
	VeProductPhoenixInverterSmart12V5000VA230Vac64k VeProduct = 0xA2B1
	VeProductPhoenixInverterSmart24V5000VA230Vac64k VeProduct = 0xA2B2
	VeProductPhoenixInverterSmart48V5000VA230Vac64k VeProduct = 0xA2B4
	VeProductPhoenixInverter12V800VA230Vac64kHS     VeProduct = 0xA2E1
	VeProductPhoenixInverter24V800VA230Vac64kHS     VeProduct = 0xA2E2
	VeProductPhoenixInverter48V800VA230Vac64kHS     VeProduct = 0xA2E4
	VeProductPhoenixInverter12V800VA120Vac64kHS     VeProduct = 0xA2E9
	VeProductPhoenixInverter24V800VA120Vac64kHS     VeProduct = 0xA2EA
	VeProductPhoenixInverter48V800VA120Vac64kHS     VeProduct = 0xA2EC
	VeProductPhoenixInverter12V1200VA230Vac64kHS    VeProduct = 0xA2F1
	VeProductPhoenixInverter24V1200VA230Vac64kHS    VeProduct = 0xA2F2
	VeProductPhoenixInverter48V1200VA230Vac64kHS    VeProduct = 0xA2F4
	VeProductPhoenixInverter12V1200VA120Vac64kHS    VeProduct = 0xA2F9
	VeProductPhoenixInverter24V1200VA120Vac64kHS    VeProduct = 0xA2FA
	VeProductPhoenixInverter48V1200VA120Vac64kHS    VeProduct = 0xA2FC
	VeProductPhoenixSmartIP43Charger12_50_1p1       VeProduct = 0xA340
	VeProductPhoenixSmartIP43Charger12_50_3         VeProduct = 0xA341
	VeProductPhoenixSmartIP43Charger24_25_1p1       VeProduct = 0xA342
	VeProductPhoenixSmartIP43Charger24_25_3         VeProduct = 0xA343
	VeProductPhoenixSmartIP43Charger12_30_1p1       VeProduct = 0xA344
	VeProductPhoenixSmartIP43Charger12_30_3         VeProduct = 0xA345
	VeProductPhoenixSmartIP43Charger24_16_1p1       VeProduct = 0xA346
	VeProductPhoenixSmartIP43Charger24_16_3         VeProduct = 0xA347
	VeProductBMV712Smart                            VeProduct = 0xA381
	VeProductBMV710HSmart                           VeProduct = 0xA382
	VeProductBMV712SmartRev2                        VeProduct = 0xA383
	VeProductSmartShunt500A_50mV                    VeProduct = 0xA389
	VeProductSmartShunt1000A_50mV                   VeProduct = 0xA38A
	VeProductSmartShunt2000A_50mV                   VeProduct = 0xA38B
)

func GetVeProductMap() map[VeProduct]string {
	return map[VeProduct]string{
		VeProductBMV700:                                 "BMV-700",
		VeProductBMV702:                                 "BMV-702",
		VeProductBMV700H:                                "BMV-700H",
		VeProductBlueSolarMPPT70_15:                     "BlueSolar MPPT 70|15",
		VeProductBlueSolarMPPT75_50:                     "BlueSolar MPPT 75|50",
		VeProductBlueSolarMPPT150_35:                    "BlueSolar MPPT 150|35",
		VeProductBlueSolarMPPT75_15:                     "BlueSolar MPPT 75|15",
		VeProductBlueSolarMPPT100_15:                    "BlueSolar MPPT 100|15",
		VeProductBlueSolarMPPT100_30:                    "BlueSolar MPPT 100|30",
		VeProductBlueSolarMPPT100_50:                    "BlueSolar MPPT 100|50",
		VeProductBlueSolarMPPT150_70:                    "BlueSolar MPPT 150|70",
		VeProductBlueSolarMPPT150_100:                   "BlueSolar MPPT 150|100",
		VeProductBlueSolarMPPT100_50rev2:                "BlueSolar MPPT 100|50 rev2",
		VeProductBlueSolarMPPT100_30rev2:                "BlueSolar MPPT 100|30 rev2",
		VeProductBlueSolarMPPT150_35rev2:                "BlueSolar MPPT 150|35 rev2",
		VeProductBlueSolarMPPT75_10:                     "BlueSolar MPPT 75|10",
		VeProductBlueSolarMPPT150_45:                    "BlueSolar MPPT 150|45",
		VeProductBlueSolarMPPT150_60:                    "BlueSolar MPPT 150|60",
		VeProductBlueSolarMPPT150_85:                    "BlueSolar MPPT 150|85",
		VeProductSmartSolarMPPT250_100:                  "SmartSolar MPPT 250|100",
		VeProductSmartSolarMPPT150_100:                  "SmartSolar MPPT 150|100",
		VeProductSmartSolarMPPT150_85:                   "SmartSolar MPPT 150|85",
		VeProductSmartSolarMPPT75_15:                    "SmartSolar MPPT 75|15",
		VeProductSmartSolarMPPT75_10:                    "SmartSolar MPPT 75|10",
		VeProductSmartSolarMPPT100_15:                   "SmartSolar MPPT 100|15",
		VeProductSmartSolarMPPT100_30:                   "SmartSolar MPPT 100|30",
		VeProductSmartSolarMPPT100_50:                   "SmartSolar MPPT 100|50",
		VeProductSmartSolarMPPT150_35:                   "SmartSolar MPPT 150|35",
		VeProductSmartSolarMPPT150_100rev2:              "SmartSolar MPPT 150|100 rev2",
		VeProductSmartSolarMPPT150_85rev2:               "SmartSolar MPPT 150|85 rev2",
		VeProductSmartSolarMPPT250_70:                   "SmartSolar MPPT 250|70",
		VeProductSmartSolarMPPT250_85:                   "SmartSolar MPPT 250|85",
		VeProductSmartSolarMPPT250_60:                   "SmartSolar MPPT 250|60",
		VeProductSmartSolarMPPT250_45:                   "SmartSolar MPPT 250|45",
		VeProductSmartSolarMPPT100_20:                   "SmartSolar MPPT 100|20",
		VeProductSmartSolarMPPT100_2048V:                "SmartSolar MPPT 100|20 48V",
		VeProductSmartSolarMPPT150_45:                   "SmartSolar MPPT 150|45",
		VeProductSmartSolarMPPT150_60:                   "SmartSolar MPPT 150|60",
		VeProductSmartSolarMPPT150_70:                   "SmartSolar MPPT 150|70",
		VeProductSmartSolarMPPT250_85rev2:               "SmartSolar MPPT 250|85 rev2",
		VeProductSmartSolarMPPT250_100rev2:              "SmartSolar MPPT 250|100 rev2",
		VeProductBlueSolarMPPT100_20:                    "BlueSolar MPPT 100|20",
		VeProductBlueSolarMPPT100_2048V:                 "BlueSolar MPPT 100|20 48V",
		VeProductSmartSolarMPPT250_60rev2:               "SmartSolar MPPT 250|60 rev2",
		VeProductSmartSolarMPPT250_70rev2:               "SmartSolar MPPT 250|70 rev2",
		VeProductSmartSolarMPPT150_45rev2:               "SmartSolar MPPT 150|45 rev2",
		VeProductSmartSolarMPPT150_60rev2:               "SmartSolar MPPT 150|60 rev2",
		VeProductSmartSolarMPPT150_70rev2:               "SmartSolar MPPT 150|70 rev2",
		VeProductSmartSolarMPPT150_85rev3:               "SmartSolar MPPT 150|85 rev3",
		VeProductSmartSolarMPPT150_100rev3:              "SmartSolar MPPT 150|100 rev3",
		VeProductBlueSolarMPPT150_45rev2:                "BlueSolar MPPT 150|45 rev2",
		VeProductBlueSolarMPPT150_60rev2:                "BlueSolar MPPT 150|60 rev2",
		VeProductBlueSolarMPPT150_70rev2:                "BlueSolar MPPT 150|70 rev2",
		VeProductSmartSolarMPPTVECan150_70:              "SmartSolar MPPT VE.Can 150/70",
		VeProductSmartSolarMPPTVECan150_45:              "SmartSolar MPPT VE.Can 150/45",
		VeProductSmartSolarMPPTVECan150_60:              "SmartSolar MPPT VE.Can 150/60",
		VeProductSmartSolarMPPTVECan150_85:              "SmartSolar MPPT VE.Can 150/85",
		VeProductSmartSolarMPPTVECan150_100:             "SmartSolar MPPT VE.Can 150/100",
		VeProductSmartSolarMPPTVECan250_45:              "SmartSolar MPPT VE.Can 250/45",
		VeProductSmartSolarMPPTVECan250_60:              "SmartSolar MPPT VE.Can 250/60",
		VeProductSmartSolarMPPTVECan250_70:              "SmartSolar MPPT VE.Can 250/70",
		VeProductSmartSolarMPPTVECan250_85:              "SmartSolar MPPT VE.Can 250/85",
		VeProductSmartSolarMPPTVECan250_100:             "SmartSolar MPPT VE.Can 250/100",
		VeProductSmartSolarMPPTVECan150_70rev2:          "SmartSolar MPPT VE.Can 150/70 rev2",
		VeProductSmartSolarMPPTVECan150_85rev2:          "SmartSolar MPPT VE.Can 150/85 rev2",
		VeProductSmartSolarMPPTVECan150_100rev2:         "SmartSolar MPPT VE.Can 150/100 rev2",
		VeProductBlueSolarMPPTVECan150_100:              "BlueSolar MPPT VE.Can 150/100",
		VeProductBlueSolarMPPTVECan250_70:               "BlueSolar MPPT VE.Can 250/70",
		VeProductBlueSolarMPPTVECan250_100:              "BlueSolar MPPT VE.Can 250/100",
		VeProductSmartSolarMPPTVECan250_70rev2:          "SmartSolar MPPT VE.Can 250/70 rev2",
		VeProductSmartSolarMPPTVECan250_100rev2:         "SmartSolar MPPT VE.Can 250/100 rev2",
		VeProductSmartSolarMPPTVECan250_85rev2:          "SmartSolar MPPT VE.Can 250/85 rev2",
		VeProductPhoenixInverter12V250VA230V:            "Phoenix Inverter 12V 250VA 230V",
		VeProductPhoenixInverter24V250VA230V:            "Phoenix Inverter 24V 250VA 230V",
		VeProductPhoenixInverter48V250VA230V:            "Phoenix Inverter 48V 250VA 230V",
		VeProductPhoenixInverter12V250VA120V:            "Phoenix Inverter 12V 250VA 120V",
		VeProductPhoenixInverter24V250VA120V:            "Phoenix Inverter 24V 250VA 120V",
		VeProductPhoenixInverter48V250VA120V:            "Phoenix Inverter 48V 250VA 120V",
		VeProductPhoenixInverter12V375VA230V:            "Phoenix Inverter 12V 375VA 230V",
		VeProductPhoenixInverter24V375VA230V:            "Phoenix Inverter 24V 375VA 230V",
		VeProductPhoenixInverter48V375VA230V:            "Phoenix Inverter 48V 375VA 230V",
		VeProductPhoenixInverter12V375VA120V:            "Phoenix Inverter 12V 375VA 120V",
		VeProductPhoenixInverter24V375VA120V:            "Phoenix Inverter 24V 375VA 120V",
		VeProductPhoenixInverter48V375VA120V:            "Phoenix Inverter 48V 375VA 120V",
		VeProductPhoenixInverter12V500VA230V:            "Phoenix Inverter 12V 500VA 230V",
		VeProductPhoenixInverter24V500VA230V:            "Phoenix Inverter 24V 500VA 230V",
		VeProductPhoenixInverter48V500VA230V:            "Phoenix Inverter 48V 500VA 230V",
		VeProductPhoenixInverter12V500VA120V:            "Phoenix Inverter 12V 500VA 120V",
		VeProductPhoenixInverter24V500VA120V:            "Phoenix Inverter 24V 500VA 120V",
		VeProductPhoenixInverter48V500VA120V:            "Phoenix Inverter 48V 500VA 120V",
		VeProductPhoenixInverter12V800VA230V:            "Phoenix Inverter 12V 800VA 230V",
		VeProductPhoenixInverter24V800VA230V:            "Phoenix Inverter 24V 800VA 230V",
		VeProductPhoenixInverter48V800VA230V:            "Phoenix Inverter 48V 800VA 230V",
		VeProductPhoenixInverter12V800VA120V:            "Phoenix Inverter 12V 800VA 120V",
		VeProductPhoenixInverter24V800VA120V:            "Phoenix Inverter 24V 800VA 120V",
		VeProductPhoenixInverter48V800VA120V:            "Phoenix Inverter 48V 800VA 120V",
		VeProductPhoenixInverter12V1200VA230V:           "Phoenix Inverter 12V 1200VA 230V",
		VeProductPhoenixInverter24V1200VA230V:           "Phoenix Inverter 24V 1200VA 230V",
		VeProductPhoenixInverter48V1200VA230V:           "Phoenix Inverter 48V 1200VA 230V",
		VeProductPhoenixInverter12V1200VA120V:           "Phoenix Inverter 12V 1200VA 120V",
		VeProductPhoenixInverter24V1200VA120V:           "Phoenix Inverter 24V 1200VA 120V",
		VeProductPhoenixInverter48V1200VA120V:           "Phoenix Inverter 48V 1200VA 120V",
		VeProductPhoenixInverter12V1600VA230V:           "Phoenix Inverter 12V 1600VA 230V",
		VeProductPhoenixInverter24V1600VA230V:           "Phoenix Inverter 24V 1600VA 230V",
		VeProductPhoenixInverter48V1600VA230V:           "Phoenix Inverter 48V 1600VA 230V",
		VeProductPhoenixInverter12V2000VA230V:           "Phoenix Inverter 12V 2000VA 230V",
		VeProductPhoenixInverter24V2000VA230V:           "Phoenix Inverter 24V 2000VA 230V",
		VeProductPhoenixInverter48V2000VA230V:           "Phoenix Inverter 48V 2000VA 230V",
		VeProductPhoenixInverter12V3000VA230V:           "Phoenix Inverter 12V 3000VA 230V",
		VeProductPhoenixInverter24V3000VA230V:           "Phoenix Inverter 24V 3000VA 230V",
		VeProductPhoenixInverter48V3000VA230V:           "Phoenix Inverter 48V 3000VA 230V",
		VeProductPhoenixInverterSmart12V5000VA230Vac64k: "Phoenix Inverter Smart 12V 5000VA 230Vac 64k",
		VeProductPhoenixInverterSmart24V5000VA230Vac64k: "Phoenix Inverter Smart 24V 5000VA 230Vac 64k",
		VeProductPhoenixInverterSmart48V5000VA230Vac64k: "Phoenix Inverter Smart 48V 5000VA 230Vac 64k",
		VeProductPhoenixInverter12V800VA230Vac64kHS:     "Phoenix Inverter 12V 800VA 230Vac 64k HS",
		VeProductPhoenixInverter24V800VA230Vac64kHS:     "Phoenix Inverter 24V 800VA 230Vac 64k HS",
		VeProductPhoenixInverter48V800VA230Vac64kHS:     "Phoenix Inverter 48V 800VA 230Vac 64k HS",
		VeProductPhoenixInverter12V800VA120Vac64kHS:     "Phoenix Inverter 12V 800VA 120Vac 64k HS",
		VeProductPhoenixInverter24V800VA120Vac64kHS:     "Phoenix Inverter 24V 800VA 120Vac 64k HS",
		VeProductPhoenixInverter48V800VA120Vac64kHS:     "Phoenix Inverter 48V 800VA 120Vac 64k HS",
		VeProductPhoenixInverter12V1200VA230Vac64kHS:    "Phoenix Inverter 12V 1200VA 230Vac 64k HS",
		VeProductPhoenixInverter24V1200VA230Vac64kHS:    "Phoenix Inverter 24V 1200VA 230Vac 64k HS",
		VeProductPhoenixInverter48V1200VA230Vac64kHS:    "Phoenix Inverter 48V 1200VA 230Vac 64k HS",
		VeProductPhoenixInverter12V1200VA120Vac64kHS:    "Phoenix Inverter 12V 1200VA 120Vac 64k HS",
		VeProductPhoenixInverter24V1200VA120Vac64kHS:    "Phoenix Inverter 24V 1200VA 120Vac 64k HS",
		VeProductPhoenixInverter48V1200VA120Vac64kHS:    "Phoenix Inverter 48V 1200VA 120Vac 64k HS",
		VeProductPhoenixSmartIP43Charger12_50_1p1:       "Phoenix Smart IP43 Charger 12|50 (1+1)",
		VeProductPhoenixSmartIP43Charger12_50_3:         "Phoenix Smart IP43 Charger 12|50 (3)",
		VeProductPhoenixSmartIP43Charger24_25_1p1:       "Phoenix Smart IP43 Charger 24|25 (1+1)",
		VeProductPhoenixSmartIP43Charger24_25_3:         "Phoenix Smart IP43 Charger 24|25 (3)",
		VeProductPhoenixSmartIP43Charger12_30_1p1:       "Phoenix Smart IP43 Charger 12|30 (1+1)",
		VeProductPhoenixSmartIP43Charger12_30_3:         "Phoenix Smart IP43 Charger 12|30 (3)",
		VeProductPhoenixSmartIP43Charger24_16_1p1:       "Phoenix Smart IP43 Charger 24|16 (1+1)",
		VeProductPhoenixSmartIP43Charger24_16_3:         "Phoenix Smart IP43 Charger 24|16 (3)",
		VeProductBMV712Smart:                            "BMV-712 Smart",
		VeProductBMV710HSmart:                           "BMV-710H Smart",
		VeProductBMV712SmartRev2:                        "BMV-712 Smart Rev2",
		VeProductSmartShunt500A_50mV:                    "SmartShunt 500A/50mV",
		VeProductSmartShunt1000A_50mV:                   "SmartShunt 1000A/50mV",
		VeProductSmartShunt2000A_50mV:                   "SmartShunt 2000A/50mV",
	}
}

func (product VeProduct) String() string {
	m := GetVeProductMap()
	if v, ok := m[product]; ok {
		return v
	}
	return ""
}
