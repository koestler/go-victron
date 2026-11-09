# Changelog

## 0.4.0
- Veproduct: renamed SmartSolarMPPT70_10rev2 (0xA074), PhoenixInverter{12V,24V,48V}{1600,2000,3000}VA230V
- Veproduct: Add various new BlueSolarMPPT devices, add TypeOrionXS 
- Veregister: fix wrong signedness of BMV AuxVoltage, MainVoltageMinimum, MainVoltageMaximum
  and wrong unit casing of solar LoadSwitchHighLevel/LoadSwitchLowLevel.
- Require Go 1.22 (previously 1.21).
- Rename inconsistent names (DecodeSolarChargerRecord, LynxSmartBmsRecord, InverterWarningReason*, LiFePO4) and fix typos.
- vedirect/vedirectapi/veconst/ble bugs: Improve error handling, fix possible panics, make ordering deterministic, check for enum overflow
- bleparser: fix various decoder byte offset errors
- vedirect: fix command checksum being sent with a single hex digit when below 0x10.
- veregister: fix duplicate sort value
- ble: remove PKCS7 padding, verify first key byte, gate debug logs, decode all record types.

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
