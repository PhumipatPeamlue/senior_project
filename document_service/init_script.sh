#!/bin/bash

/usr/local/bin/docker-entrypoint.sh &

until curl -s -XGET "http://localhost:9200/" > /dev/null; do
    echo "waiting elasticsearch..."
    sleep 1
done

if [ $? -eq 0 ]; then

  curl -X PUT "localhost:9200/video_doc" -H 'Content-Type: application/json' -d '{
     "settings": {
        "number_of_shards": 1
     },
     "mappings": {
        "properties": {
           "title": { "type": "text" },
           "video_url": { "type": "text" },
           "description": { "type": "text" },
           "created_at": { "type": "date" },
           "updated_at": { "type": "date" }
        }
     }
  }'

  curl -X PUT "localhost:9200/drug_doc" -H 'Content-Type: application/json' -d '{
     "settings": {
        "number_of_shards": 1
     },
     "mappings": {
        "properties": {
           "trade_name": { "type": "text" },
           "drug_name": { "type": "text" },
           "description": { "type": "text" },
           "preparation": { "type": "text" },
           "caution": { "type": "text" },
           "created_at": { "type": "date" },
           "updated_at": { "type": "date" }
        }
     }
  }'
fi

tail -f /dev/null