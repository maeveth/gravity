# Gravity API Specification

## API Version

Currently all Gravity API is v1, versioning is added for future expansion and backwards compatability as Gravity API evolves

## URI Schema

Base Path: /api/{apiVersion}

## Plugins

### GET /plugins

Lists information plugin subsystem and brief status of loaded plugins

#### Example HTTP Response

````json
{
    "config": {
        "pluginPath": "/opt/gravity/plugins"
    },
    "plugins" : [
        {
            "name": "powerviewHub",
            "status": "active",
            "msg": null,
            "version": "1",
            "description": "Implements a Gravity interface for Hunter PowerView Hub REST API",
            "numDevices": 1
        }
    ]
}
````

### GET /plugins/{pluginName}

Displays full information, status, and associated devices for a single plugin. This does not trigger a update of status on all devices - returned device status is from cached data.

#### Example: GET /plugins/powerviewHub

````json
{
    "name": "powerviewHub",
    "status": "active",
    "msg": null,
    "version": "1",
    "description": "Implements a Gravity interface for Hunter PowerView Hub REST API",
    "numDevices": 2,
    "updateInterval": 60,
    "devices": [
        {
            "id": "A1890CF4D1D11A6C",
            "name": "My Hub",
            "typeId": "PowerviewHub Gen2",
            "position": 0,
            "status": "online",
            "msg": null,
            "updateOverride": 0,
            "pluginVer": 1,
            "lastUpdate": "2018-09-12T18:25:43.511Z",
            "commands": [
                "discover",
                "activateScene/{%s}"
            ],
            "outputs": [
                "name",
                "typeId",
                "status",
                "shades",
                "scenes"
            ]
        },
        {
            "id": "12346",
            "name": "Living Room Shade 1",
            "typeId": "Duette",
            "position": 0,
            "status": "online",
            "msg": null,
            "updateOverride": 0,
            "pluginVer": 1,
            "lastUpdate": "2018-09-12T18:25:43.511Z",
            "commands": [
                "open",
                "close",
                "up",
                "down",
                "stop",
                "incrUp/{%d}",
                "incrUp/{%d}",
                "set/{%d}",
                "update",
                "calibrate"
            ],
            "outputs": [
                "name",
                "typeId",
                "position",
                "status"
            ]
        }
    ]
}
````

### GET /device/{pluginName}/{deviceId}

Returns information on a single device associated with a specific plugin.  Gravity makes no assumptions on format for device IDs to allow for easier integration with downstream devices.  As a result devices must be addressed by plugin and its own internal device ID. Standard URL compliance applies however.

#### Example: GET /plugins/powerviewHub/A1890CF4D1D11A6C

````json
{
    "id": "A1890CF4D1D11A6C",
    "name": "My Hub",
    "typeId": "PowerviewHub Gen2",
    "position": 0,
    "status": "online",
    "msg": null,
    "updateOverride": 0,
    "pluginVer": 1,
    "lastUpdate": "2018-09-12T18:25:43.511Z",
    "commands": [
        "discover"
    ],
    "outputs": [
        "name",
        "typeId",
        "position",
        "status"
    ]
}
````

### POST /device/{pluginName}/{deviceId}/{command}

Execute a specific command on a device.  Supported commands are reflected in information for each devicea. In some cases no payload may be required for the commanad

#### Example: POST /device/powerviewHub/12346/up
