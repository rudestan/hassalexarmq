ARG BUILD_FROM
FROM $BUILD_FROM

ENV LANG C.UTF-8

RUN curl -sL https://github.com/rudestan/hassalexarmq/releases/download/v0.1-alpha/arm_alexa_rmq_consumer.tar.gz --output arm_alexa_rmq_consumer.tar.gz \
    && tar -xzvf arm_alexa_rmq_consumer.tar.gz

COPY run.sh /

RUN chmod a+x /run.sh

CMD [ "/run.sh" ]