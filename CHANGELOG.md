# Changelog

## 0.4.0
- Veproduct: fix SmartSolarMPPT70_10rev2 (0xA074), renamed to SmartSolarMPPT75_10rev2.
- Veproduct: fix PhoenixInverter{12V,24V,48V}{1600,2000,3000}VA230V, renamed to
  PhoenixInverterSmart{12V,24V,48V}{1600,2000,3000}VA230Vac64k and changed their type from
  TypePhoenixInverter to TypePhoenixInverterSmart.
- Veproduct: add ids BlueSolarMPPT75_10rev3 (0xA07C), BlueSolarMPPT75_15rev3 (0xA07D),
  SmartSolarMPPT100_30_12V (0xA07E), AllInOneSmartSolarMPPT75_15_12V (0xA07F),
  SmartSolarMPPT250_60rev3 (0xA080), SmartSolarMPPT250_70rev3 (0xA081).
- Veproduct: add new TypeOrionXS and ids OrionXS12V_12V_50A (0xA3F0) and OrionXS1400 (0xA3F1)
  (register support is not implemented yet).
- Veregister: fix wrong signedness of BMV AuxVoltage, MainVoltageMinimum, MainVoltageMaximum
  and wrong unit casing of solar LoadSwitchHighLevel/LoadSwitchLowLevel.
- Require Go 1.22 (previously 1.21).
- Rename inconsistent names (DecodeSolarChargerRecord, LynxSmartBmsRecord, InverterWarningReason*, LiFePO4) and fix typos.
- vedirect/vedirectapi/veconst/ble bugs: Improve error handling, fix possible panics, make ordering deterministic, check for enum overflow

## 0.3.0
- fix exposed names:
  - Corrected the exposed name to `SmartSolarMPPT150_45rev3`.
- fix various typos in comments

## 0.2.1
- veregister: filter unavailable ProductRevision and Description registers for BMV700,BMV702,BMV700H

## 0.2.0
- veregister: add solar load registers for 10/15/20A solar chargers.

## 0.1.0
- veproduct.Type: add IsBMV, IsSolar, and IsInverter methods.
- veproduct.Product: add MaxPanelVoltage and MaxPanelCurrent methods.
- veproduct: Replace TestGetRegisterListByProductType by TestGetRegisterListByProduct.
  This is done because some solar chargers  do not have the Panel current available.
