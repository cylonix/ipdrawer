package serverpb

const (
	SWAGGER = `{
  "swagger": "2.0",
  "info": {
    "title": "serverpb/server.proto",
    "version": "version not set"
  },
  "tags": [
    {
      "name": "NetworkServiceV0"
    },
    {
      "name": "IPServiceV0"
    },
    {
      "name": "PoolServiceV0"
    }
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/{namespace}/drawip": {
      "get": {
        "operationId": "NetworkServiceV0_DrawIPEstimatingNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDrawIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "poolTag.key",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "poolTag.value",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "temporaryReserved",
            "in": "query",
            "required": false,
            "type": "boolean"
          },
          {
            "name": "sequential",
            "description": "Default false i.e. randomized",
            "in": "query",
            "required": false,
            "type": "boolean"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/list": {
      "get": {
        "operationId": "IPServiceV0_ListIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbListIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/temporary_reserved/list": {
      "get": {
        "operationId": "IPServiceV0_ListTemporaryReservedIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbListTemporaryReservedIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/{ip}/activate": {
      "post": {
        "operationId": "IPServiceV0_ActivateIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbCreateIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IPServiceV0ActivateIPBody"
            }
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/{ip}/create": {
      "post": {
        "operationId": "IPServiceV0_CreateIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbCreateIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IPServiceV0CreateIPBody"
            }
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/{ip}/deactivate": {
      "post": {
        "operationId": "IPServiceV0_DeactivateIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDeactivateIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/{ip}/network": {
      "get": {
        "operationId": "IPServiceV0_GetNetworkIncludingIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/ip/{ip}/update": {
      "post": {
        "operationId": "IPServiceV0_UpdateIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbUpdateIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IPServiceV0UpdateIPBody"
            }
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network": {
      "get": {
        "operationId": "NetworkServiceV0_GetEstimatedNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/update": {
      "post": {
        "operationId": "NetworkServiceV0_UpdateNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbUpdateNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0UpdateNetworkBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}": {
      "get": {
        "operationId": "NetworkServiceV0_GetNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "name",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}/create": {
      "post": {
        "operationId": "NetworkServiceV0_CreateNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbCreateNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0CreateNetworkBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}/delete": {
      "post": {
        "operationId": "NetworkServiceV0_DeleteNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDeleteNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}/drawip": {
      "post": {
        "operationId": "NetworkServiceV0_DrawIP",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDrawIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "description": "The ip address wanted or the network to draw ip from. Optional.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0DrawIPBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}/pool/create": {
      "post": {
        "operationId": "NetworkServiceV0_CreatePool",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbCreatePoolResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0CreatePoolBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{ip}/{mask}/pools": {
      "get": {
        "operationId": "NetworkServiceV0_GetPoolsInNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetPoolsInNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{name}": {
      "get": {
        "operationId": "NetworkServiceV0_GetNetwork2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "ip",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/network/{name}/drawip": {
      "post": {
        "operationId": "NetworkServiceV0_DrawIP2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDrawIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0DrawIPBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/networks": {
      "get": {
        "operationId": "NetworkServiceV0_ListNetwork",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbListNetworkResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/pool/list": {
      "get": {
        "operationId": "PoolServiceV0_ListPool",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbListPoolResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "PoolServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/delete": {
      "post": {
        "operationId": "PoolServiceV0_DeletePool",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDeletePoolResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeStart",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeEnd",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "PoolServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/drawip": {
      "post": {
        "operationId": "NetworkServiceV0_DrawIP3",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbDrawIPResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeStart",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeEnd",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NetworkServiceV0DrawIPBody"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/ip": {
      "get": {
        "operationId": "PoolServiceV0_GetIPInPool",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbGetIPInPoolResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeStart",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "rangeEnd",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "PoolServiceV0"
        ]
      }
    },
    "/api/v0/{namespace}/pool/{start}/{end}/update": {
      "post": {
        "operationId": "PoolServiceV0_UpdatePool",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/serverpbUpdatePoolResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "start",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "end",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PoolServiceV0UpdatePoolBody"
            }
          }
        ],
        "tags": [
          "PoolServiceV0"
        ]
      }
    }
  },
  "definitions": {
    "IPServiceV0ActivateIPBody": {
      "type": "object",
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "uuid": {
          "type": "string"
        }
      }
    },
    "IPServiceV0CreateIPBody": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/modelIPAddrStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "uuid": {
          "type": "string"
        }
      }
    },
    "IPServiceV0UpdateIPBody": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/modelIPAddrStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "uuid": {
          "type": "string"
        }
      }
    },
    "NetworkServiceV0CreateNetworkBody": {
      "type": "object",
      "properties": {
        "defaultGateways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "status": {
          "$ref": "#/definitions/modelNetworkStatus"
        }
      }
    },
    "NetworkServiceV0CreatePoolBody": {
      "type": "object",
      "properties": {
        "pool": {
          "$ref": "#/definitions/modelPool"
        }
      }
    },
    "NetworkServiceV0DrawIPBody": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string",
          "description": "The ip address wanted or the network to draw ip from. Optional."
        },
        "mask": {
          "type": "integer",
          "format": "int32"
        },
        "poolTag": {
          "$ref": "#/definitions/modelTag"
        },
        "name": {
          "type": "string"
        },
        "temporaryReserved": {
          "type": "boolean"
        },
        "uuid": {
          "type": "string"
        },
        "mustHaveWantIp": {
          "type": "boolean"
        },
        "sequential": {
          "type": "boolean",
          "title": "Default false i.e. randomized"
        }
      }
    },
    "NetworkServiceV0UpdateNetworkBody": {
      "type": "object",
      "properties": {
        "prefix": {
          "type": "string"
        },
        "gateways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "broadcast": {
          "type": "string"
        },
        "netmask": {
          "type": "string"
        },
        "status": {
          "$ref": "#/definitions/modelNetworkStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "PoolServiceV0UpdatePoolBody": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/modelPoolStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "googlerpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    },
    "modelIPAddr": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string"
        },
        "status": {
          "$ref": "#/definitions/modelIPAddrStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "uuid": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        }
      }
    },
    "modelIPAddrStatus": {
      "type": "integer",
      "format": "int32",
      "enum": [
        0,
        1,
        2,
        3
      ],
      "default": 0
    },
    "modelNetwork": {
      "type": "object",
      "properties": {
        "prefix": {
          "type": "string"
        },
        "gateways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "broadcast": {
          "type": "string"
        },
        "netmask": {
          "type": "string"
        },
        "status": {
          "$ref": "#/definitions/modelNetworkStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "namespace": {
          "type": "string"
        }
      }
    },
    "modelNetworkStatus": {
      "type": "integer",
      "format": "int32",
      "enum": [
        0,
        1,
        2
      ],
      "default": 0
    },
    "modelPool": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string"
        },
        "end": {
          "type": "string"
        },
        "status": {
          "$ref": "#/definitions/modelPoolStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "lastModifiedAt": {
          "type": "string",
          "format": "date-time"
        },
        "namespace": {
          "type": "string"
        }
      }
    },
    "modelPoolStatus": {
      "type": "integer",
      "format": "int32",
      "enum": [
        0,
        1,
        2
      ],
      "default": 0
    },
    "modelTag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        }
      },
      "additionalProperties": {}
    },
    "serverpbCreateIPResponse": {
      "type": "object"
    },
    "serverpbCreateNetworkResponse": {
      "type": "object"
    },
    "serverpbCreatePoolResponse": {
      "type": "object"
    },
    "serverpbDeactivateIPResponse": {
      "type": "object"
    },
    "serverpbDeleteNetworkResponse": {
      "type": "object"
    },
    "serverpbDeletePoolResponse": {
      "type": "object"
    },
    "serverpbDrawIPResponse": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "serverpbGetIPInPoolResponse": {
      "type": "object",
      "properties": {
        "pool": {
          "$ref": "#/definitions/modelPool"
        },
        "ips": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelIPAddr"
          }
        }
      }
    },
    "serverpbGetNetworkResponse": {
      "type": "object",
      "properties": {
        "network": {
          "type": "string"
        },
        "defaultGateways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "broadcast": {
          "type": "string"
        },
        "netmask": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelTag"
          }
        }
      }
    },
    "serverpbGetPoolsInNetworkResponse": {
      "type": "object",
      "properties": {
        "pools": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelPool"
          }
        }
      }
    },
    "serverpbListIPResponse": {
      "type": "object",
      "properties": {
        "ips": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelIPAddr"
          }
        }
      }
    },
    "serverpbListNetworkResponse": {
      "type": "object",
      "properties": {
        "networks": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelNetwork"
          }
        }
      }
    },
    "serverpbListPoolResponse": {
      "type": "object",
      "properties": {
        "pools": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelPool"
          }
        }
      }
    },
    "serverpbListTemporaryReservedIPResponse": {
      "type": "object",
      "properties": {
        "ips": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/modelIPAddr"
          }
        }
      }
    },
    "serverpbUpdateIPResponse": {
      "type": "object"
    },
    "serverpbUpdateNetworkResponse": {
      "type": "object"
    },
    "serverpbUpdatePoolResponse": {
      "type": "object"
    }
  }
}
`
)
