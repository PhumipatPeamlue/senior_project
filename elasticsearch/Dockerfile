FROM docker.elastic.co/elasticsearch/elasticsearch:8.12.0

ENV ES_JAVA_OPTS="-Xms2g -Xmx2g"
ENV bootstrap.memory_lock=true
ENV discovery.type=single-node
ENV xpack.security.enabled=false
ENV xpack.security.enrollment.enabled=false

EXPOSE 9200

RUN /usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-icu

ARG CPU_LIMIT=2.0
ARG CPU_RESERVATION=1.0

CMD ["elasticsearch"]