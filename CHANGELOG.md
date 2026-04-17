# Changelog

## 0.3.0
- fix exposed names:
  - SmatrtSolarMPPT150_45rev3 -> SmartSolarMPPT150_45rev3
  - SmatrtSolarMPPT150_45rev3 -> SmartSolarMPPT150_45rev3
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
