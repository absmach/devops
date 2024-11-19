#!/bin/bash

helm install elasticsearch stable/elasticsearch \
    --set data.resources.requests.memory=512Mi \
    --set client.replicas=1 \
    --set master.replicas=1 \
    --set cluster.env.MINIMUM_MASTER_NODES=1 \
    --set cluster.env.RECOVER_AFTER_MASTER_NODES=1 \
    --set cluster.env.EXPECTED_MASTER_NODES=1 \
    --set data.replicas=1 \
    --set data.heapSize=300m \
    --set master.persistence.size=10Gi \
    --set data.persistence.size=10Gi \
    --wait

helm install fluent-bit stable/fluent-bit \
    --set backend.type=es \
    --set backend.es.host=elasticsearch-client \
    --set filter.mergeJSONLog=false

helm install kibana stable/kibana \
    --set env.ELASTICSEARCH_HOSTS=http://elasticsearch-client:9200
