#!/bin/bash
#
# To create an index in elasticsearch
#

# shards & replicas are default values.
#

source config.sh

TODAY=$(date '+%Y%m%d')
INDEX="publiccodes"

generate_index_settings() {
  cat <<EOF
{
  "settings" : {
    "index" : {
      "number_of_shards" : 5,
      "number_of_replicas" : 1
    }
  },
  "mappings": {
    "software": {
      "dynamic_templates": [
        {
          "description": {
            "path_match": "description.*",
            "mapping": {
              "type": "object",
              "properties": {
                "localisedName": {
                  "type": "text"
                },
                "genericName": {
                  "type": "text",
                  "fields": {
                    "keyword": { "type": "keyword", "ignore_above": 256 }
                  }
                },
                "shortDescription": {
                  "type": "text"
                },
                "longDescription": {
                  "type": "text"
                },
                "documentation": {
                  "type": "text",
                  "index": false
                },
                "apiDocumentation": {
                  "type": "text",
                  "index": false
                },
                "featureList": {
                  "type": "keyword"
                },
                "freeTags": {
                  "type": "keyword"
                },
                "screenshots": {
                  "type": "keyword",
                  "index": false
                },
                "videos": {
                  "type": "keyword",
                  "index": false
                },
                "awards": {
                  "type": "keyword"
                }
              }
            }
          }
        }
      ],
      "properties": {
        "publiccode-yaml-version": {
          "type": "text",
          "index": false
        },
        "name": {
          "type": "text"
        },
        "id": {
          "type": "text"
        },
        "crawltime": {
          "type": "text"
        },
        "applicationSuite": {
          "type": "text",
          "fields": { "keyword": { "type": "keyword", "ignore_above": 256 } }
        },
        "url": {
          "type": "text",
          "index": false,
          "fields": { "keyword": { "type": "keyword", "ignore_above": 256 } }
        },
        "landingURL": {
          "type": "text",
          "index": false
        },
        "isBasedOn": {
          "type": "text",
          "index": false
        },
        "softwareVersion": {
          "type": "keyword"
        },
        "releaseDate": {
          "type": "date",
          "format": "strict_date"
        },
        "logo": {
          "type": "text",
          "index": false
        },
        "monochromeLogo": {
          "type": "text",
          "index": false
        },
        "inputTypes": {
          "type": "keyword"
        },
        "outputTypes": {
          "type": "keyword"
        },
        "platforms": {
          "type": "keyword"
        },
        "tags": {
          "type": "keyword"
        },
        "usedBy": {
          "type": "text",
          "fields": { "keyword": { "type": "keyword", "ignore_above": 256 } }
        },
        "roadmap": {
          "type": "text",
          "index": false
        },
        "developmentStatus": {
          "type": "keyword"
        },
        "softwareType": {
          "type": "keyword"
        },
        "intendedAudience-onlyFor": {
          "type": "keyword"
        },
        "intendedAudience-countries": {
          "type": "keyword"
        },
        "intendedAudience-unsupportedCountries": {
          "type": "keyword"
        },
        "legal-license": {
          "type": "text",
          "fields": { "keyword": { "type": "keyword", "ignore_above": 256 } }
        },
        "legal-mainCopyrightOwner": {
          "type": "text"
        },
        "legal-repoOwner": {
          "type": "text"
        },
        "legal-authorsFile": {
          "type": "text",
          "index": false
        },
        "maintenance-type": {
          "type": "keyword"
        },
        "maintenance-contractors": {
          "type": "nested",
          "properties": {
            "name": {
              "type": "text"
            },
            "until": {
              "type": "date",
              "format": "strict_date"
            },
            "website": {
              "type": "text",
              "index": false
            }
          }
        },
        "maintenance-contacts": {
          "type": "nested",
          "properties": {
            "name": {
              "type": "text"
            },
            "email": {
              "type": "text"
            },
            "phone": {
              "type": "text",
              "index": false
            },
            "affiliation": {
              "type": "text"
            }
          }
        },
        "localisation-localisationReady": {
          "type": "boolean"
        },
        "localisation-availableLanguages": {
          "type": "keyword"
        },
        "dependsOn-open": {
          "type": "nested",
          "properties": {
            "name": {
              "type": "text"
            },
            "version-min": {
              "type": "text",
              "index": false
            },
            "version-max": {
              "type": "text",
              "index": false
            },
            "version": {
              "type": "text",
              "index": false
            },
            "optional": {
              "type": "boolean"
            }
          }
        },
        "dependsOn-proprietary": {
          "type": "nested",
          "properties": {
            "name": {
              "type": "text"
            },
            "version-min": {
              "type": "text",
              "index": false
            },
            "version-max": {
              "type": "text",
              "index": false
            },
            "version": {
              "type": "text",
              "index": false
            },
            "optional": {
              "type": "boolean"
            }
          }
        },
        "dependsOn-hardware": {
          "type": "nested",
          "properties": {
            "name": {
              "type": "text"
            },
            "version-min": {
              "type": "text",
              "index": false
            },
            "version-max": {
              "type": "text",
              "index": false
            },
            "version": {
              "type": "text",
              "index": false
            },
            "optional": {
              "type": "boolean"
            }
          }
        },
        "it-conforme-accessibile": {
          "type": "boolean"
        },
        "it-conforme-interoperabile": {
          "type": "boolean"
        },
        "it-conforme-sicuro": {
          "type": "boolean"
        },
        "it-conforme-privacy": {
          "type": "boolean"
        },
        "it-spid": {
          "type": "boolean"
        },
        "it-cie": {
          "type": "boolean"
        },
        "it-anpr": {
          "type": "boolean"
        },
        "it-pagopa": {
          "type": "boolean"
        },
        "it-riuso-codiceIPA": {
          "type": "keyword"
        },
        "it-ecosistemi": {
          "type": "keyword"
        },
        "it-designKit-seo": {
          "type": "boolean"
        },
        "it-designKit-ui": {
          "type": "boolean"
        },
        "it-designKit-web": {
          "type": "boolean"
        },
        "it-designKit-content": {
          "type": "boolean"
        },
        "suggest-name": {
          "type": "completion"
        },
        "vitality-score": {
          "type": "text",
          "index": false
        },
        "vitality-dataChart": {
          "type": "integer"
        },
        "related-software": {
          "properties": {
            "name": {
              "type": "text",
              "index": false
            },
            "image": {
              "type": "text",
              "index": false
            },
            "eng": {
              "properties": {
                "localised-name": {
                  "type": "text",
                  "index": false
                },
                "url": {
                  "type": "text",
                  "index": false
                }
              }
            },
            "ita": {
              "properties": {
                "localised-name": {
                  "type": "text",
                  "index": false
                },
                "url": {
                  "type": "text",
                  "index": false
                }
              }
            }
          }
        },
        "tags-related": {
          "type": "keyword"
        },
        "popular-tags": {
          "type": "keyword"
        },
        "share-tags": {
          "type": "keyword"
        },
        "old-variant": {
          "properties": {
            "name": {
              "type": "text",
              "index": false
            },
            "eng": {
              "properties": {
                "localised-name": {
                  "type": "text",
                  "index": false
                },
                "url": {
                  "type": "text",
                  "index": false
                },
                "feature-list": {
                  "type": "keyword",
                  "index": false
                },
                "vitality-score": {
                  "type": "integer",
                  "index": false
                },
                "legal-repo-owner": {
                  "type": "text",
                  "index": false
                }
              }
            },
            "ita": {
              "properties": {
                "localised-name": {
                  "type": "text",
                  "index": false
                },
                "url": {
                  "type": "text",
                  "index": false
                },
                "feature-list": {
                  "type": "keyword",
                  "index": false
                },
                "vitality-score": {
                  "type": "integer",
                  "index": false
                },
                "legal-repo-owner": {
                  "type": "text",
                  "index": false
                }
              }
            }
          }
        },
        "old-feature-list": {
          "properties": {
            "ita": {
              "type": "keyword",
              "index": false
            },
            "eng": {
              "type": "keyword",
              "index": false
            }
          }
        }
      }
    }
  }
}
EOF
}

curl -u "$BASICAUTH" -X PUT "$ELASTICSEARCH_URL/$INDEX" -H 'Content-Type: application/json' -d"$(generate_index_settings)"
