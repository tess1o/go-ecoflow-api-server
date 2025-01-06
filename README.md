# Ecoflow Rest API Server

## Table of Contents

1. [Description](#description)
2. [Ecoflow API implemented so far](#ecoflow-api-implemented-so-far)
3. [Deployment](#deployment)
    - [Get docker container from registry](#get-docker-container-from-registry)
    - [Build the Server from source](#build-the-server-from-source)
    - [Build a Docker Image from source](#build-a-docker-image-from-source)
4. [Requests / Responses](#requests--responses)
    - [Get all linked devices](#get-all-linked-devices)
    - [Get all parameters for given device](#get-all-parameters-for-given-device)
    - [Get specified parameters for specified device](#get-specified-parameters-for-specified-device)
    - [Enable/Disable AC/X-Boost](#enabledisable-acx-boost)
    - [Enable/Disable DC](#enabledisable-dc)
    - [Enable/Disable Car Output](#enabledisable-car-output)
    - [Change charging speed](#change-charging-speed)
    - [Change car input](#change-car-input)
    - [Change StandBy parameters](#change-standby-parameters)

## Description

This project implements a RESTful API server designed to manage and interact with Ecoflow devices.
The server provides endpoints to query and control Ecoflow-compatible hardware, offering seamless
integration for monitoring and management purposes.

The server does not store any customer data, to access your Ecoflow devices you need to send Access and Secret Tokens as
headers.These tokens are not logged and not recorded.

The access token is sent in `Authorization: Bearer XXX` header, the secret token as `X-Secret-Token` header.

## Ecoflow API implemented so far:

1. Get all linked devices
2. Get all parameters for given device
3. Get specified parameters for given device
4. Enable/disable AC
5. Enable/disable DC
6. Enable/disable Car input
7. Change charging speed
8. Change Car Input
9. Change stand by settings for device, AC, DC, LCD screen

## Deployment

### Get docker container from registry

Todo...

### Build the Server from source

1. Make sure you have Go installed (version 1.23 or later).
2. Clone the repository:

   ```shell
   git clone https://github.com/tess1o/go-ecoflow-rest-api
   cd go-ecoflow-rest-api
   ```

3. Build the project:

   ```shell
   go build -o go-ecoflow-rest-api .
   ```

4. Run the server locally:

   ```shell
   ./go-ecoflow-rest-api
   ```

Now the server should be accessible on `http://localhost:8080`.

### Build a Docker Image from source

1. Build the Docker image:

   ```shell
   docker build -t go-ecoflow-rest-api:latest .
   ```

2. Run the Docker container:

   ```shell
   docker run -p 8080:8080 go-ecoflow-rest-api:latest
   ```

Now the server should be accessible on `http://localhost:8080`.

## Requests / Responses

- ### Get all linked devices

**Request**

```shell
curl -XGET http://localhost:8080/api/devices \
-H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
-H "X-Secret-Token: YOUR_SECRET_TOKEN"
```

**Response**:

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success",
    "data": [
      {
        "sn": "R331ZCB4ZEXXXXXX",
        "online": 0
      },
      {
        "sn": "R351ZCB5HXXXXXX",
        "online": 1
      },
      {
        "sn": "R601ZCB5HXXXXXX",
        "online": 1
      }
    ],
    "eagleEyeTraceId": "ea1a2a58291736177544541XXXXXX",
    "tid": ""
  }
}
```

- ### Get all parameters for given device

**Request**

```shell
curl -XGET http://localhost:8080/api/devices/R351ZCB5HG8XXXXXX/parameters \
-H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
-H "X-Secret-Token: YOUR_SECRET_TOKEN"
```

<details>
  <summary>Expand big response</summary>

```json
{
  "success": true,
  "data": {
    "bms_bmsInfo.accuChgCap": 361937,
    "bms_bmsInfo.accuChgEnergy": 19776,
    "bms_bmsInfo.accuDsgCap": 215726,
    "bms_bmsInfo.accuDsgEnergy": 14344,
    "bms_bmsInfo.bmsLanchDate": 79,
    "bms_bmsInfo.bsmCycles": 9,
    "bms_bmsInfo.deepDsgCnt": 0,
    "bms_bmsInfo.highTempChgTime": 0,
    "bms_bmsInfo.highTempTime": 0,
    "bms_bmsInfo.lauchDateFlag": 3,
    "bms_bmsInfo.lowTempChgTime": 0,
    "bms_bmsInfo.lowTempTime": 0,
    "bms_bmsInfo.num": 0,
    "bms_bmsInfo.ohmRes": 20,
    "bms_bmsInfo.powerCapability": 1.5,
    "bms_bmsInfo.resetFlag": 0,
    "bms_bmsInfo.roundTrip": 70,
    "bms_bmsInfo.selfDsgRate": 5,
    "bms_bmsInfo.sn": "0000000000000000",
    "bms_bmsInfo.soh": 100,
    "bms_bmsStatus.actSoc": 98,
    "bms_bmsStatus.allBmsFault": 0,
    "bms_bmsStatus.allErrCode": 8388608,
    "bms_bmsStatus.amp": 5,
    "bms_bmsStatus.balanceState": 0,
    "bms_bmsStatus.bmsFault": 0,
    "bms_bmsStatus.bmsHeartVer": 259,
    "bms_bmsStatus.bqSysStatReg": 0,
    "bms_bmsStatus.caleSoh": 0,
    "bms_bmsStatus.cellId": 2,
    "bms_bmsStatus.cellTemp": [],
    "bms_bmsStatus.cellVol": [],
    "bms_bmsStatus.chgCap": 4294967295,
    "bms_bmsStatus.chgState": 0,
    "bms_bmsStatus.cycSoh": 99,
    "bms_bmsStatus.cycles": 9,
    "bms_bmsStatus.designCap": 40000,
    "bms_bmsStatus.diffSoc": 0,
    "bms_bmsStatus.dsgCap": 4294967295,
    "bms_bmsStatus.ecloudOcv": 65535,
    "bms_bmsStatus.errCode": 23,
    "bms_bmsStatus.f32ShowSoc": 98,
    "bms_bmsStatus.fullCap": 39278,
    "bms_bmsStatus.hwVersion": [
      86,
      48,
      46,
      49,
      46,
      50
    ],
    "bms_bmsStatus.inputWatts": 0,
    "bms_bmsStatus.loaderVer": 33619974,
    "bms_bmsStatus.maxCellTemp": 33,
    "bms_bmsStatus.maxCellVol": 3327,
    "bms_bmsStatus.maxMosTemp": 36,
    "bms_bmsStatus.maxVolDiff": 4,
    "bms_bmsStatus.minCellTemp": 31,
    "bms_bmsStatus.minCellVol": 3323,
    "bms_bmsStatus.minMosTemp": 25,
    "bms_bmsStatus.mosState": 2,
    "bms_bmsStatus.num": 0,
    "bms_bmsStatus.openBmsIdx": 1,
    "bms_bmsStatus.outputWatts": 0,
    "bms_bmsStatus.packSn": "0000000000000000",
    "bms_bmsStatus.productDetail": 2,
    "bms_bmsStatus.productType": 81,
    "bms_bmsStatus.realSoh": 0,
    "bms_bmsStatus.remainCap": 38889,
    "bms_bmsStatus.remainTime": 0,
    "bms_bmsStatus.soc": 98,
    "bms_bmsStatus.soh": 100,
    "bms_bmsStatus.sysState": 2,
    "bms_bmsStatus.sysVer": 33620258,
    "bms_bmsStatus.tagChgAmp": 40000,
    "bms_bmsStatus.targetSoc": 98,
    "bms_bmsStatus.temp": 33,
    "bms_bmsStatus.type": 1,
    "bms_bmsStatus.vol": 53355,
    "bms_emsStatus.bmsIsConnt": [
      3,
      0,
      0
    ],
    "bms_emsStatus.bmsModel": 1,
    "bms_emsStatus.bmsWarState": 0,
    "bms_emsStatus.chgAmp": 15761,
    "bms_emsStatus.chgCmd": 1,
    "bms_emsStatus.chgDisCond": 2,
    "bms_emsStatus.chgLinePlug": 34,
    "bms_emsStatus.chgRemainTime": 5999,
    "bms_emsStatus.chgState": 2,
    "bms_emsStatus.chgVol": 58928,
    "bms_emsStatus.dsgCmd": 1,
    "bms_emsStatus.dsgDisCond": 0,
    "bms_emsStatus.dsgRemainTime": 5999,
    "bms_emsStatus.emsIsNormalFlag": 1,
    "bms_emsStatus.emsVer": 259,
    "bms_emsStatus.f32LcdShowSoc": 98.1,
    "bms_emsStatus.fanLevel": 0,
    "bms_emsStatus.lcdShowSoc": 98,
    "bms_emsStatus.maxAvailNum": 1,
    "bms_emsStatus.maxChargeSoc": 100,
    "bms_emsStatus.maxCloseOilEb": 100,
    "bms_emsStatus.minDsgSoc": 1,
    "bms_emsStatus.minOpenOilEb": 0,
    "bms_emsStatus.openBmsIdx": 1,
    "bms_emsStatus.openUpsFlag": 1,
    "bms_emsStatus.paraVolMax": 54361,
    "bms_emsStatus.paraVolMin": 52287,
    "bms_emsStatus.sysChgDsgState": 0,
    "bms_kitInfo.aviDataLen": 83,
    "bms_kitInfo.kitNum": 2,
    "bms_kitInfo.version": 1,
    "bms_kitInfo.watts": [
      {
        "appState": 0,
        "appVer": 0,
        "avaFlag": 0,
        "curPower": 0,
        "detail": 0,
        "f32Soc": 0,
        "loadVer": 0,
        "sn": "",
        "soc": 0,
        "type": 0
      },
      {
        "appState": 0,
        "appVer": 0,
        "avaFlag": 0,
        "curPower": 0,
        "detail": 0,
        "f32Soc": 0,
        "loadVer": 0,
        "sn": "",
        "soc": 0,
        "type": 0
      }
    ],
    "inv.FastChgWatts": 2400,
    "inv.SlowChgWatts": 800,
    "inv.acChgRatedPower": 2400,
    "inv.acDipSwitch": 2,
    "inv.acInAmp": 729,
    "inv.acInFreq": 50,
    "inv.acInVol": 230579,
    "inv.acPassbyAutoEn": 0,
    "inv.cfgAcEnabled": 1,
    "inv.cfgAcOutFreq": 1,
    "inv.cfgAcOutVol": 220000,
    "inv.cfgAcWorkMode": 0,
    "inv.cfgAcXboost": 1,
    "inv.chargerType": 255,
    "inv.chgPauseFlag": 0,
    "inv.dcInAmp": 0,
    "inv.dcInTemp": 39,
    "inv.dcInVol": 0,
    "inv.dischargeType": 1,
    "inv.errCode": 0,
    "inv.fanState": 0,
    "inv.inputWatts": 170,
    "inv.invOutAmp": 729,
    "inv.invOutFreq": 50,
    "inv.invOutVol": 230304,
    "inv.invType": 10,
    "inv.outTemp": 36,
    "inv.outputWatts": 170,
    "inv.prBalanceMode": 0,
    "inv.reserved": [
      0,
      0,
      0,
      0,
      0,
      0
    ],
    "inv.standbyMin": 0,
    "inv.sysVer": 33554509,
    "mppt.carOutAmp": 46,
    "mppt.carOutVol": 0,
    "mppt.carOutWatts": 0,
    "mppt.carStandbyMin": 0,
    "mppt.carState": 0,
    "mppt.carTemp": 37,
    "mppt.cfgChgType": 0,
    "mppt.chgPauseFlag": 0,
    "mppt.chgState": 0,
    "mppt.chgType": 0,
    "mppt.dc24vState": 0,
    "mppt.dc24vTemp": 36,
    "mppt.dcChgCurrent": 8000,
    "mppt.dcdc12vAmp": 0,
    "mppt.dcdc12vVol": 0,
    "mppt.dcdc12vWatts": 0,
    "mppt.faultCode": 0,
    "mppt.inAmp": 0,
    "mppt.inVol": 5,
    "mppt.inWatts": 0,
    "mppt.mpptTemp": 36,
    "mppt.outAmp": 130,
    "mppt.outVol": 51635,
    "mppt.outWatts": 6,
    "mppt.pv2CfgChgType": 0,
    "mppt.pv2ChgPauseFlag": 0,
    "mppt.pv2ChgState": 0,
    "mppt.pv2ChgType": 0,
    "mppt.pv2DcChgCurrent": 8000,
    "mppt.pv2InAmp": 0,
    "mppt.pv2InVol": 2,
    "mppt.pv2InWatts": 0,
    "mppt.pv2MpptTemp": 38,
    "mppt.pv2Xt60ChgType": 0,
    "mppt.res": [
      0,
      0,
      0,
      0
    ],
    "mppt.swVer": 83886164,
    "mppt.x60ChgType": 0,
    "pd.XT150Watts1": 0,
    "pd.XT150Watts2": 0,
    "pd.acAutoOnCfg": 0,
    "pd.acAutoPause": 0,
    "pd.beepMode": 1,
    "pd.bmsInfoFull": 900000,
    "pd.bmsInfoIncre": 30000,
    "pd.bmsKitState": [
      0,
      0
    ],
    "pd.bmsRunIncre": 30000,
    "pd.bpPowerSoc": 80,
    "pd.brightLevel": 100,
    "pd.carState": 0,
    "pd.carTemp": 34,
    "pd.carUsedTime": 2272,
    "pd.carWatts": 0,
    "pd.chgDsgState": 1,
    "pd.chgPowerAC": 213052,
    "pd.chgPowerDC": 475,
    "pd.chgSunPower": 0,
    "pd.dcInUsedTime": 2215,
    "pd.dcOutState": 0,
    "pd.dsgPowerAC": 201414,
    "pd.dsgPowerDC": 265,
    "pd.errCode": 0,
    "pd.hysteresisAdd": 5,
    "pd.icoBytes": [
      0,
      0,
      132,
      0,
      128,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0
    ],
    "pd.invInWatts": 170,
    "pd.invOutWatts": 170,
    "pd.invUsedTime": 6690590,
    "pd.lcdOffSec": 60,
    "pd.minAcSoc": 6,
    "pd.model": 1,
    "pd.mpptUsedTime": 0,
    "pd.newAcAutoOnCfg": 1,
    "pd.otherKitState": 0,
    "pd.pdInfoFull": 900000,
    "pd.pdInfoIncre": 30000,
    "pd.pdRunIncre": 120000,
    "pd.pv1ChargeType": 0,
    "pd.pv1ChargeWatts": 0,
    "pd.pv2ChargeType": 0,
    "pd.pv2ChargeWatts": 0,
    "pd.pvChargePrioSet": 255,
    "pd.qcUsb1Watts": 0,
    "pd.qcUsb2Watts": 0,
    "pd.relaySwitchCnt": 0,
    "pd.remainTime": 5999,
    "pd.reserved": [
      0,
      0
    ],
    "pd.soc": 98,
    "pd.standbyMin": 0,
    "pd.sysVer": 16975450,
    "pd.typec1Temp": 30,
    "pd.typec1Watts": 0,
    "pd.typec2Temp": 30,
    "pd.typec2Watts": 0,
    "pd.typecUsedTime": 4503,
    "pd.usb1Watts": 0,
    "pd.usb2Watts": 0,
    "pd.usbUsedTime": 84035,
    "pd.usbqcUsedTime": 1793,
    "pd.watchIsConfig": 0,
    "pd.wattsInSum": 170,
    "pd.wattsOutSum": 170,
    "pd.wifiAutoRcvy": 0,
    "pd.wifiRssi": 0,
    "pd.wifiVer": 0,
    "pd.wireWatts": 0
  }
}
```

</details>

- ### Get specified parameters for specified device

```shell
curl -XPOST http://localhost:8080/api/devices/R351ZCB5HGXXXXXX/parameters/query \
-H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
-H "X-Secret-Token: YOUR_SECRET_TOKEN" \
-d '{"parameters": ["bms_bmsStatus.cycles", "bms_bmsStatus.maxMosTemp"]}'
```

**Parameters Explanation:**

- **parameters** - list of Ecoflow parameters that should be returned by the API. Must not be empty.

**Response**:

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success",
    "data": {
      "bms_bmsStatus.cycles": 9,
      "bms_bmsStatus.maxMosTemp": 36
    },
    "eagleEyeTraceId": "ea1a2a5c3217361965046798094d0007",
    "tid": ""
  }
}
```

- ### Enable/Disable AC/X-Boost

**Request**:

```shell
curl -XPUT http://localhost:8080/api/power_station/R601ZCB5HXXXXX/out/ac -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -H "X-Secret-Token: YOUR_SECRET_TOKEN" -d '{"ac_state": "on", "xboost_state": "on", "out_freq": 50, "out_voltage" : 220}'
```

**Parameters Explanation:**

- **`ac_state`**: A string that specifies the state of the AC. Acceptable values are `"on"` to
  enable or `"off"` to disable the AC functionality.

- **`xboost_state`**: A string that determines the state of the X-Boost. Acceptable values are `"on"` to enable
  or `"off"` to disable X-Boost mode, which optimizes load handling and power management.

- **`out_freq`**: An integer that specifies the output frequency in Hertz. Typical values are `50` (common in many
  regions) or `60` (used in others).

- **`out_voltage`**: An integer representing the output voltage in volts. For example, `220` volts is common in some
  regions, while others may use `110` volts.

**Response**:

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```

- ### Enable/Disable DC

**Request**:

```shell
curl -XPUT http://localhost:8080/api/power_station/R351ZCB5HGXXXXX/out/dc -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -H "X-Secret-Token: YOUR_SECRET_TOKEN" -d '{"state": "on"}'
```

**Explanation of Parameters**

- **`state`**: This parameter specifies the desired state of the DC (Direct Current) output.
    - Acceptable values:
        - `"on"`: Enables the DC functionality.
        - `"off"`: Disables the DC functionality.

**Response**:

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```

- ### Enable/Disable Car Output

**Request**:

```shell
curl -XPUT http://localhost:8080/api/power_station/R351ZCB5HGXXXXX/out/car -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -H "X-Secret-Token: YOUR_SECRET_TOKEN" -d '{"state": "on"}'
```

**Explanation of Parameters**

- **`state`**: This parameter specifies the desired state of the Car output.
    - Acceptable values:
        - `"on"`: Enables the Car Output
        - `"off"`: Disables the Car Output

**Response**:

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```

- ### Change charging speed

**Request**

```shell
 5280  curl -XPUT http://localhost:8080/api/power_station/R601ZCB5HEAXXXXX/input/speed \
 -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
 -H "X-Secret-Token: YOUR_SECRET_TOKENS" \
 -d '{"watts":150}'
```

**Explanation of Parameters**

- **`watts`**: This parameter specifies the charging speed in watts.

**Response**

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```

- ### Change car input

**Request**

```shell
 5280  curl -XPUT http://localhost:8080/api/power_station/R601ZCB5HEAXXXXX/input/car \
 -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
 -H "X-Secret-Token: YOUR_SECRET_TOKENS" \
 -d '{"amps":8}'
```

**Explanation of Parameters**

- **`amps`**: Set 12 V DC (car charger) charging current(Maximum DC charging current (mA)). From 4 to 10 amps (according
  to Ecoflow documents)

**Response**

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```

- ### Change StandBy parameters

**Request**

```shell
 5280  curl -XPOST http://localhost:8080/api/power_station/R601ZCB5HEAXXXXX/standby \
 -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
 -H "X-Secret-Token: YOUR_SECRET_TOKENS" \
 -d '{"type":"lcd", "stand_by":60}'
```

### Parameters Explanation

- **`type`**: Specifies the type of system for which the standby value is being set.  
  Possible values:
    - `device`
    - `ac`
    - `car`
    - `lcd`

- **`stand_by`**: The standby duration to set.
    - For `lcd`, specify the value in **seconds**.
    - For all other types (`device`, `ac`, `car`), specify the value in **minutes**.

**Response**

```json
{
  "success": true,
  "data": {
    "code": "0",
    "message": "Success"
  }
}
```