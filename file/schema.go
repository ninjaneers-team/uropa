// Code generated by go generate; DO NOT EDIT.
package file

const contentSchema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "_format_version": {
      "type": "string"
    },
    "policies": {
      "items": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "$ref": "#/definitions/FPolicy"
      },
      "type": "array"
    }
  },
  "additionalProperties": false,
  "type": "object",
  "definitions": {
    "FPolicy": {
      "properties": {
        "id": {
          "type": "string"
        },
        "raw": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object"
    }
  }
}`
