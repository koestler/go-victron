// Package veproduct provides a list of Victron Energy products, their device id, and their names as string.
//
// The source of the list is: https://www.victronenergy.com/upload/documents/VE.Direct-Protocol-3.33.pdf
package veproduct

// Product defines the exact model of a Victron Energy product.
// It is different for different product lines,, different power / voltage options and hardware revisions.
type Product uint16

const (
	BMV700                                 Product = 0x203
	BMV702                                 Product = 0x204
	BMV700H                                Product = 0x205
	BlueSolarMPPT70_15                     Product = 0x0300
	BlueSolarMPPT75_50                     Product = 0xA040
	BlueSolarMPPT150_35                    Product = 0xA041
	BlueSolarMPPT75_15                     Product = 0xA042
	BlueSolarMPPT100_15                    Product = 0xA043
	BlueSolarMPPT100_30                    Product = 0xA044
	BlueSolarMPPT100_50                    Product = 0xA045
	BlueSolarMPPT150_70                    Product = 0xA046
	BlueSolarMPPT150_100                   Product = 0xA047
	BlueSolarMPPT100_50rev2                Product = 0xA049
	BlueSolarMPPT100_30rev2                Product = 0xA04A
	BlueSolarMPPT150_35rev2                Product = 0xA04B
	BlueSolarMPPT75_10                     Product = 0xA04C
	BlueSolarMPPT150_45                    Product = 0xA04D
	BlueSolarMPPT150_60                    Product = 0xA04E
	BlueSolarMPPT150_85                    Product = 0xA04F
	SmartSolarMPPT250_100                  Product = 0xA050
	SmartSolarMPPT150_100                  Product = 0xA051
	SmartSolarMPPT150_85                   Product = 0xA052
	SmartSolarMPPT75_15                    Product = 0xA053
	SmartSolarMPPT75_10                    Product = 0xA054
	SmartSolarMPPT100_15                   Product = 0xA055
	SmartSolarMPPT100_30                   Product = 0xA056
	SmartSolarMPPT100_50                   Product = 0xA057
	SmartSolarMPPT150_35                   Product = 0xA058
	SmartSolarMPPT150_100rev2              Product = 0xA059
	SmartSolarMPPT150_85rev2               Product = 0xA05A
	SmartSolarMPPT250_70                   Product = 0xA05B
	SmartSolarMPPT250_85                   Product = 0xA05C
	SmartSolarMPPT250_60                   Product = 0xA05D
	SmartSolarMPPT250_45                   Product = 0xA05E
	SmartSolarMPPT100_20                   Product = 0xA05F
	SmartSolarMPPT100_2048V                Product = 0xA060
	SmartSolarMPPT150_45                   Product = 0xA061
	SmartSolarMPPT150_60                   Product = 0xA062
	SmartSolarMPPT150_70                   Product = 0xA063
	SmartSolarMPPT250_85rev2               Product = 0xA064
	SmartSolarMPPT250_100rev2              Product = 0xA065
	BlueSolarMPPT100_20                    Product = 0xA066
	BlueSolarMPPT100_2048V                 Product = 0xA067
	SmartSolarMPPT250_60rev2               Product = 0xA068
	SmartSolarMPPT250_70rev2               Product = 0xA069
	SmartSolarMPPT150_45rev2               Product = 0xA06A
	SmartSolarMPPT150_60rev2               Product = 0xA06B
	SmartSolarMPPT150_70rev2               Product = 0xA06C
	SmartSolarMPPT150_85rev3               Product = 0xA06D
	SmartSolarMPPT150_100rev3              Product = 0xA06E
	BlueSolarMPPT150_45rev2                Product = 0xA06F
	BlueSolarMPPT150_60rev2                Product = 0xA070
	BlueSolarMPPT150_70rev2                Product = 0xA071
	BlueSolarMPPT150_45rev3                Product = 0xA072
	SmatrtSolarMPPT150_45rev3              Product = 0xA073
	SmartSolarMPPT70_10rev2                Product = 0xA074
	SmartSolarMPPT75_15rev2                Product = 0xA075
	BlueSolarMPPT100_30rev3                Product = 0xA076
	BlueSolarMPPT100_50rev3                Product = 0xA077
	BlueSolarMPPT150_35rev3                Product = 0xA078
	BlueSolarMPPT75_10rev2                 Product = 0xA079
	BlueSolarMPPT75_15rev2                 Product = 0xA07A
	BlueSolarMPPT100_15rev2                Product = 0xA07B
	SmartSolarMPPTVECan150_70              Product = 0xA102
	SmartSolarMPPTVECan150_45              Product = 0xA103
	SmartSolarMPPTVECan150_60              Product = 0xA104
	SmartSolarMPPTVECan150_85              Product = 0xA105
	SmartSolarMPPTVECan150_100             Product = 0xA106
	SmartSolarMPPTVECan250_45              Product = 0xA107
	SmartSolarMPPTVECan250_60              Product = 0xA108
	SmartSolarMPPTVECan250_70              Product = 0xA109
	SmartSolarMPPTVECan250_85              Product = 0xA10A
	SmartSolarMPPTVECan250_100             Product = 0xA10B
	SmartSolarMPPTVECan150_70rev2          Product = 0xA10C
	SmartSolarMPPTVECan150_85rev2          Product = 0xA10D
	SmartSolarMPPTVECan150_100rev2         Product = 0xA10E
	BlueSolarMPPTVECan150_100              Product = 0xA10F
	BlueSolarMPPTVECan250_70               Product = 0xA112
	BlueSolarMPPTVECan250_100              Product = 0xA113
	SmartSolarMPPTVECan250_70rev2          Product = 0xA114
	SmartSolarMPPTVECan250_100rev2         Product = 0xA115
	SmartSolarMPPTVECan250_85rev2          Product = 0xA116
	BlueSolarMPPTVECan150_100rev2          Product = 0xA117
	PhoenixInverter12V250VA230V            Product = 0xA231
	PhoenixInverter24V250VA230V            Product = 0xA232
	PhoenixInverter48V250VA230V            Product = 0xA234
	PhoenixInverter12V250VA120V            Product = 0xA239
	PhoenixInverter24V250VA120V            Product = 0xA23A
	PhoenixInverter48V250VA120V            Product = 0xA23C
	PhoenixInverter12V375VA230V            Product = 0xA241
	PhoenixInverter24V375VA230V            Product = 0xA242
	PhoenixInverter48V375VA230V            Product = 0xA244
	PhoenixInverter12V375VA120V            Product = 0xA249
	PhoenixInverter24V375VA120V            Product = 0xA24A
	PhoenixInverter48V375VA120V            Product = 0xA24C
	PhoenixInverter12V500VA230V            Product = 0xA251
	PhoenixInverter24V500VA230V            Product = 0xA252
	PhoenixInverter48V500VA230V            Product = 0xA254
	PhoenixInverter12V500VA120V            Product = 0xA259
	PhoenixInverter24V500VA120V            Product = 0xA25A
	PhoenixInverter48V500VA120V            Product = 0xA25C
	PhoenixInverter12V800VA230V            Product = 0xA261
	PhoenixInverter24V800VA230V            Product = 0xA262
	PhoenixInverter48V800VA230V            Product = 0xA264
	PhoenixInverter12V800VA120V            Product = 0xA269
	PhoenixInverter24V800VA120V            Product = 0xA26A
	PhoenixInverter48V800VA120V            Product = 0xA26C
	PhoenixInverter12V1200VA230V           Product = 0xA271
	PhoenixInverter24V1200VA230V           Product = 0xA272
	PhoenixInverter48V1200VA230V           Product = 0xA274
	PhoenixInverter12V1200VA120V           Product = 0xA279
	PhoenixInverter24V1200VA120V           Product = 0xA27A
	PhoenixInverter48V1200VA120V           Product = 0xA27C
	PhoenixInverter12V1600VA230V           Product = 0xA281
	PhoenixInverter24V1600VA230V           Product = 0xA282
	PhoenixInverter48V1600VA230V           Product = 0xA284
	PhoenixInverter12V2000VA230V           Product = 0xA291
	PhoenixInverter24V2000VA230V           Product = 0xA292
	PhoenixInverter48V2000VA230V           Product = 0xA294
	PhoenixInverter12V3000VA230V           Product = 0xA2A1
	PhoenixInverter24V3000VA230V           Product = 0xA2A2
	PhoenixInverter48V3000VA230V           Product = 0xA2A4
	PhoenixInverterSmart12V5000VA230Vac64k Product = 0xA2B1
	PhoenixInverterSmart24V5000VA230Vac64k Product = 0xA2B2
	PhoenixInverterSmart48V5000VA230Vac64k Product = 0xA2B4
	PhoenixInverter12V800VA230Vac64kHS     Product = 0xA2E1
	PhoenixInverter24V800VA230Vac64kHS     Product = 0xA2E2
	PhoenixInverter48V800VA230Vac64kHS     Product = 0xA2E4
	PhoenixInverter12V800VA120Vac64kHS     Product = 0xA2E9
	PhoenixInverter24V800VA120Vac64kHS     Product = 0xA2EA
	PhoenixInverter48V800VA120Vac64kHS     Product = 0xA2EC
	PhoenixInverter12V1200VA230Vac64kHS    Product = 0xA2F1
	PhoenixInverter24V1200VA230Vac64kHS    Product = 0xA2F2
	PhoenixInverter48V1200VA230Vac64kHS    Product = 0xA2F4
	PhoenixInverter12V1200VA120Vac64kHS    Product = 0xA2F9
	PhoenixInverter24V1200VA120Vac64kHS    Product = 0xA2FA
	PhoenixInverter48V1200VA120Vac64kHS    Product = 0xA2FC
	PhoenixSmartIP43Charger12_50_1p1       Product = 0xA340
	PhoenixSmartIP43Charger12_50_3         Product = 0xA341
	PhoenixSmartIP43Charger24_25_1p1       Product = 0xA342
	PhoenixSmartIP43Charger24_25_3         Product = 0xA343
	PhoenixSmartIP43Charger12_30_1p1       Product = 0xA344
	PhoenixSmartIP43Charger12_30_3         Product = 0xA345
	PhoenixSmartIP43Charger24_16_1p1       Product = 0xA346
	PhoenixSmartIP43Charger24_16_3         Product = 0xA347
	BMV712Smart                            Product = 0xA381
	BMV710HSmart                           Product = 0xA382
	BMV712SmartRev2                        Product = 0xA383
	SmartShunt500A_50mV                    Product = 0xA389
	SmartShunt1000A_50mV                   Product = 0xA38A
	SmartShunt2000A_50mV                   Product = 0xA38B
)

// idStringMap is const and therefore only copies are exposed.
var idStringMap = map[Product]struct {
	Model           string
	Type            Type
	MaxPanelVoltage int
	MaxPanelCurrent int
}{
	BMV700:                                 {"700", TypeBMV, -1, -1},
	BMV702:                                 {"702", TypeBMV, -1, -1},
	BMV700H:                                {"700H", TypeBMV, -1, -1},
	BlueSolarMPPT70_15:                     {"70|15", TypeBlueSolarMPPT, 70, 15},
	BlueSolarMPPT75_50:                     {"75|50", TypeBlueSolarMPPT, 75, 50},
	BlueSolarMPPT150_35:                    {"150|35", TypeBlueSolarMPPT, 150, 35},
	BlueSolarMPPT75_15:                     {"75|15", TypeBlueSolarMPPT, 75, 15},
	BlueSolarMPPT100_15:                    {"100|15", TypeBlueSolarMPPT, 100, 15},
	BlueSolarMPPT100_30:                    {"100|30", TypeBlueSolarMPPT, 100, 30},
	BlueSolarMPPT100_50:                    {"100|50", TypeBlueSolarMPPT, 100, 50},
	BlueSolarMPPT150_70:                    {"150|70", TypeBlueSolarMPPT, 150, 70},
	BlueSolarMPPT150_100:                   {"150|100", TypeBlueSolarMPPT, 150, 100},
	BlueSolarMPPT100_50rev2:                {"100|50 rev2", TypeBlueSolarMPPT, 100, 50},
	BlueSolarMPPT100_30rev2:                {"100|30 rev2", TypeBlueSolarMPPT, 100, 30},
	BlueSolarMPPT150_35rev2:                {"150|35 rev2", TypeBlueSolarMPPT, 150, 35},
	BlueSolarMPPT75_10:                     {"75|10", TypeBlueSolarMPPT, 75, 10},
	BlueSolarMPPT150_45:                    {"150|45", TypeBlueSolarMPPT, 150, 45},
	BlueSolarMPPT150_60:                    {"150|60", TypeBlueSolarMPPT, 150, 60},
	BlueSolarMPPT150_85:                    {"150|85", TypeBlueSolarMPPT, 150, 85},
	SmartSolarMPPT250_100:                  {"250|100", TypeSmartSolarMPPT, 250, 100},
	SmartSolarMPPT150_100:                  {"150|100", TypeSmartSolarMPPT, 150, 100},
	SmartSolarMPPT150_85:                   {"150|85", TypeSmartSolarMPPT, 150, 85},
	SmartSolarMPPT75_15:                    {"75|15", TypeSmartSolarMPPT, 75, 15},
	SmartSolarMPPT75_10:                    {"75|10", TypeSmartSolarMPPT, 75, 10},
	SmartSolarMPPT100_15:                   {"100|15", TypeSmartSolarMPPT, 100, 15},
	SmartSolarMPPT100_30:                   {"100|30", TypeSmartSolarMPPT, 100, 30},
	SmartSolarMPPT100_50:                   {"100|50", TypeSmartSolarMPPT, 100, 50},
	SmartSolarMPPT150_35:                   {"150|35", TypeSmartSolarMPPT, 150, 35},
	SmartSolarMPPT150_100rev2:              {"150|100 rev2", TypeSmartSolarMPPT, 150, 100},
	SmartSolarMPPT150_85rev2:               {"150|85 rev2", TypeSmartSolarMPPT, 150, 85},
	SmartSolarMPPT250_70:                   {"250|70", TypeSmartSolarMPPT, 250, 70},
	SmartSolarMPPT250_85:                   {"250|85", TypeSmartSolarMPPT, 250, 85},
	SmartSolarMPPT250_60:                   {"250|60", TypeSmartSolarMPPT, 250, 60},
	SmartSolarMPPT250_45:                   {"250|45", TypeSmartSolarMPPT, 250, 45},
	SmartSolarMPPT100_20:                   {"100|20", TypeSmartSolarMPPT, 100, 20},
	SmartSolarMPPT100_2048V:                {"100|20 48V", TypeSmartSolarMPPT, 100, 20},
	SmartSolarMPPT150_45:                   {"150|45", TypeSmartSolarMPPT, 150, 45},
	SmartSolarMPPT150_60:                   {"150|60", TypeSmartSolarMPPT, 150, 60},
	SmartSolarMPPT150_70:                   {"150|70", TypeSmartSolarMPPT, 150, 70},
	SmartSolarMPPT250_85rev2:               {"250|85 rev2", TypeSmartSolarMPPT, 250, 85},
	SmartSolarMPPT250_100rev2:              {"250|100 rev2", TypeSmartSolarMPPT, 250, 100},
	BlueSolarMPPT100_20:                    {"100|20", TypeBlueSolarMPPT, 100, 20},
	BlueSolarMPPT100_2048V:                 {"100|20 48V", TypeBlueSolarMPPT, 100, 20},
	SmartSolarMPPT250_60rev2:               {"250|60 rev2", TypeSmartSolarMPPT, 250, 60},
	SmartSolarMPPT250_70rev2:               {"250|70 rev2", TypeSmartSolarMPPT, 250, 70},
	SmartSolarMPPT150_45rev2:               {"150|45 rev2", TypeSmartSolarMPPT, 150, 45},
	SmartSolarMPPT150_60rev2:               {"150|60 rev2", TypeSmartSolarMPPT, 150, 60},
	SmartSolarMPPT150_70rev2:               {"150|70 rev2", TypeSmartSolarMPPT, 150, 70},
	SmartSolarMPPT150_85rev3:               {"150|85 rev3", TypeSmartSolarMPPT, 150, 85},
	SmartSolarMPPT150_100rev3:              {"150|100 rev3", TypeSmartSolarMPPT, 150, 100},
	BlueSolarMPPT150_45rev2:                {"150|45 rev2", TypeBlueSolarMPPT, 150, 45},
	BlueSolarMPPT150_60rev2:                {"150|60 rev2", TypeBlueSolarMPPT, 150, 60},
	BlueSolarMPPT150_70rev2:                {"150|70 rev2", TypeBlueSolarMPPT, 150, 70},
	BlueSolarMPPT150_45rev3:                {"150|45 rev3", TypeBlueSolarMPPT, 150, 45},
	SmatrtSolarMPPT150_45rev3:              {"150|45 rev3", TypeSmartSolarMPPT, 150, 45},
	SmartSolarMPPT70_10rev2:                {"70|10 rev2", TypeSmartSolarMPPT, 70, 10},
	SmartSolarMPPT75_15rev2:                {"75|15 rev2", TypeSmartSolarMPPT, 75, 15},
	BlueSolarMPPT100_30rev3:                {"100|30 rev3", TypeBlueSolarMPPT, 100, 30},
	BlueSolarMPPT100_50rev3:                {"100|50 rev3", TypeBlueSolarMPPT, 100, 50},
	BlueSolarMPPT150_35rev3:                {"150|35 rev3", TypeBlueSolarMPPT, 150, 35},
	BlueSolarMPPT75_10rev2:                 {"75|10 rev2", TypeBlueSolarMPPT, 75, 10},
	BlueSolarMPPT75_15rev2:                 {"75|15 rev2", TypeBlueSolarMPPT, 75, 15},
	BlueSolarMPPT100_15rev2:                {"100|15 rev2", TypeBlueSolarMPPT, 100, 15},
	SmartSolarMPPTVECan150_70:              {"150/70", TypeSmartSolarMPPTVECan, 150, 70},
	SmartSolarMPPTVECan150_45:              {"150/45", TypeSmartSolarMPPTVECan, 150, 45},
	SmartSolarMPPTVECan150_60:              {"150/60", TypeSmartSolarMPPTVECan, 150, 60},
	SmartSolarMPPTVECan150_85:              {"150/85", TypeSmartSolarMPPTVECan, 150, 85},
	SmartSolarMPPTVECan150_100:             {"150/100", TypeSmartSolarMPPTVECan, 150, 100},
	SmartSolarMPPTVECan250_45:              {"250/45", TypeSmartSolarMPPTVECan, 250, 45},
	SmartSolarMPPTVECan250_60:              {"250/60", TypeSmartSolarMPPTVECan, 250, 60},
	SmartSolarMPPTVECan250_70:              {"250/70", TypeSmartSolarMPPTVECan, 250, 70},
	SmartSolarMPPTVECan250_85:              {"250/85", TypeSmartSolarMPPTVECan, 250, 85},
	SmartSolarMPPTVECan250_100:             {"250/100", TypeSmartSolarMPPTVECan, 250, 100},
	SmartSolarMPPTVECan150_70rev2:          {"150/70 rev2", TypeSmartSolarMPPTVECan, 150, 70},
	SmartSolarMPPTVECan150_85rev2:          {"150/85 rev2", TypeSmartSolarMPPTVECan, 150, 85},
	SmartSolarMPPTVECan150_100rev2:         {"150/100 rev2", TypeSmartSolarMPPTVECan, 150, 100},
	BlueSolarMPPTVECan150_100:              {"150/100", TypeBlueSolarMPPTVECan, 150, 100},
	BlueSolarMPPTVECan250_70:               {"250/70", TypeBlueSolarMPPTVECan, 250, 70},
	BlueSolarMPPTVECan250_100:              {"250/100", TypeBlueSolarMPPTVECan, 250, 100},
	SmartSolarMPPTVECan250_70rev2:          {"250/70 rev2", TypeSmartSolarMPPTVECan, 250, 70},
	SmartSolarMPPTVECan250_100rev2:         {"250/100 rev2", TypeSmartSolarMPPTVECan, 250, 100},
	SmartSolarMPPTVECan250_85rev2:          {"250/85 rev2", TypeSmartSolarMPPTVECan, 250, 85},
	BlueSolarMPPTVECan150_100rev2:          {"150/100 rev2", TypeBlueSolarMPPTVECan, 150, 100},
	PhoenixInverter12V250VA230V:            {"12V 250VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V250VA230V:            {"24V 250VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V250VA230V:            {"48V 250VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V250VA120V:            {"12V 250VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V250VA120V:            {"24V 250VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V250VA120V:            {"48V 250VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V375VA230V:            {"12V 375VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V375VA230V:            {"24V 375VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V375VA230V:            {"48V 375VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V375VA120V:            {"12V 375VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V375VA120V:            {"24V 375VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V375VA120V:            {"48V 375VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V500VA230V:            {"12V 500VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V500VA230V:            {"24V 500VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V500VA230V:            {"48V 500VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V500VA120V:            {"12V 500VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V500VA120V:            {"24V 500VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V500VA120V:            {"48V 500VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V800VA230V:            {"12V 800VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V800VA230V:            {"24V 800VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V800VA230V:            {"48V 800VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V800VA120V:            {"12V 800VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V800VA120V:            {"24V 800VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V800VA120V:            {"48V 800VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V1200VA230V:           {"12V 1200VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V1200VA230V:           {"24V 1200VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V1200VA230V:           {"48V 1200VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V1200VA120V:           {"12V 1200VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V1200VA120V:           {"24V 1200VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V1200VA120V:           {"48V 1200VA 120V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V1600VA230V:           {"12V 1600VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V1600VA230V:           {"24V 1600VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V1600VA230V:           {"48V 1600VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V2000VA230V:           {"12V 2000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V2000VA230V:           {"24V 2000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V2000VA230V:           {"48V 2000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V3000VA230V:           {"12V 3000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V3000VA230V:           {"24V 3000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V3000VA230V:           {"48V 3000VA 230V", TypePhoenixInverter, -1, -1},
	PhoenixInverterSmart12V5000VA230Vac64k: {"12V 5000VA 230Vac 64k", TypePhoenixInverterSmart, -1, -1},
	PhoenixInverterSmart24V5000VA230Vac64k: {"24V 5000VA 230Vac 64k", TypePhoenixInverterSmart, -1, -1},
	PhoenixInverterSmart48V5000VA230Vac64k: {"48V 5000VA 230Vac 64k", TypePhoenixInverterSmart, -1, -1},
	PhoenixInverter12V800VA230Vac64kHS:     {"12V 800VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V800VA230Vac64kHS:     {"24V 800VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V800VA230Vac64kHS:     {"48V 800VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V800VA120Vac64kHS:     {"12V 800VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V800VA120Vac64kHS:     {"24V 800VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V800VA120Vac64kHS:     {"48V 800VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V1200VA230Vac64kHS:    {"12V 1200VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V1200VA230Vac64kHS:    {"24V 1200VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V1200VA230Vac64kHS:    {"48V 1200VA 230Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter12V1200VA120Vac64kHS:    {"12V 1200VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter24V1200VA120Vac64kHS:    {"24V 1200VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixInverter48V1200VA120Vac64kHS:    {"48V 1200VA 120Vac 64k HS", TypePhoenixInverter, -1, -1},
	PhoenixSmartIP43Charger12_50_1p1:       {"12|50 (1+1)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger12_50_3:         {"12|50 (3)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger24_25_1p1:       {"24|25 (1+1)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger24_25_3:         {"24|25 (3)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger12_30_1p1:       {"12|30 (1+1)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger12_30_3:         {"12|30 (3)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger24_16_1p1:       {"24|16 (1+1)", TypePhoenixSmartIP43Charger, -1, -1},
	PhoenixSmartIP43Charger24_16_3:         {"24|16 (3)", TypePhoenixSmartIP43Charger, -1, -1},
	BMV712Smart:                            {"712", TypeBMVSmart, -1, -1},
	BMV710HSmart:                           {"710H", TypeBMVSmart, -1, -1},
	BMV712SmartRev2:                        {"712 Rev2", TypeBMVSmart, -1, -1},
	SmartShunt500A_50mV:                    {"500A/50mV", TypeSmartShunt, -1, -1},
	SmartShunt1000A_50mV:                   {"1000A/50mV", TypeSmartShunt, -1, -1},
	SmartShunt2000A_50mV:                   {"2000A/50mV", TypeSmartShunt, -1, -1},
}

// Exists returns true if the product is known
func (product Product) Exists() bool {
	_, ok := idStringMap[product]
	return ok
}

// Model returns the product model (e.g. "702", "250|70 rev2") if known, otherwise an empty string
func (product Product) Model() string {
	if v, ok := idStringMap[product]; ok {
		return v.Model
	}
	return ""
}

// Type returns type (e.g. "BMV", "SmartSolar", "BlueSolar", "SmartShunt" etc.) of the product
func (product Product) Type() Type {
	if v, ok := idStringMap[product]; ok {
		return v.Type
	}
	return TypeUnknown
}

// String returns the product type and model (e.g. "BMV 702", "BlueSolar MPPT 250|70 rev2"), otherwise an empty string
func (product Product) String() string {
	if v, ok := idStringMap[product]; ok {
		return v.Type.String() + " " + v.Model
	}
	return ""
}

// MaxPanelVoltage returns the specified maximal voltage of solar chargers in V.
// Returns -1 for products that are not known solar chargers.
func (product Product) MaxPanelVoltage() int {
	if v, ok := idStringMap[product]; ok {
		if v.Type.IsSolar() {
			return v.MaxPanelVoltage
		}
	}
	return -1
}

// MaxPanelCurrent returns the specified maximal current of solar chargers in A.
// Returns -1 for products that are not known solar chargers.
func (product Product) MaxPanelCurrent() int {
	if v, ok := idStringMap[product]; ok {
		if v.Type.IsSolar() {
			return v.MaxPanelCurrent
		}
	}
	return -1
}

// GetStringMap returns map of id to device name (= type + model)
func GetStringMap() map[Product]string {
	ret := make(map[Product]string, len(idStringMap))
	for k := range idStringMap {
		ret[k] = k.String()
	}
	return ret
}
