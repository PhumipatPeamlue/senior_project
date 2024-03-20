#!/bin/bash

curl -X PUT "localhost:9200/video_doc" -H 'Content-Type: application/json' -d '{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "analyzer_shingle": {
            "tokenizer": "icu_tokenizer",
            "filter": [
              "filter_shingle"
            ]
          }
        },
        "filter": {
          "filter_shingle": {
            "type": "shingle",
            "max_shingle_size": 3,
            "min_shingle_size": 2,
            "output_unigrams": "true"
          }
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "title": {
        "analyzer": "analyzer_shingle",
        "type": "text"
      },
      "video_url": { "type": "text" },
      "description": {
        "analyzer": "analyzer_shingle",
        "type": "text"
      },
      "created_at": { "type": "date" },
      "updated_at": { "type": "date" }
    }
  }
}'

curl -X PUT "localhost:9200/drug_doc" -H 'Content-Type: application/json' -d '{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "analyzer_shingle": {
            "tokenizer": "icu_tokenizer",
            "filter": [
              "filter_shingle"
            ]
          }
        },
        "filter": {
          "filter_shingle": {
            "type": "shingle",
            "max_shingle_size": 3,
            "min_shingle_size": 2,
            "output_unigrams": "true"
          }
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "trade_name": {
        "analyzer": "analyzer_shingle",
        "type": "text"
      },
      "drug_name": {
        "analyzer": "analyzer_shingle",
        "type": "text"
      },
      "description": {
        "analyzer": "analyzer_shingle",
        "type": "text"
      },
      "preparation": { "type": "text" },
      "caution": { "type": "text" },
      "created_at": { "type": "date" },
      "updated_at": { "type": "date" }
    }
  }
}'