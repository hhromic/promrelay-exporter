FROM python:3.10-alpine AS base

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

FROM base AS depbuilder

RUN apk add --no-cache build-base

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

FROM base AS final

LABEL org.opencontainers.image.source https://github.com/hhromic/prometheus-relay-exporter

ENTRYPOINT ["python", "-m", "relay_exporter"]

EXPOSE 9878

COPY --from=depbuilder /usr/local /usr/local

ARG APP_HOME=/app
ARG APP_USER=app
ARG APP_GROUP=app
RUN addgroup ${APP_GROUP} \
    && adduser -D -h ${APP_HOME} -G ${APP_GROUP} ${APP_USER}
WORKDIR ${APP_HOME}

COPY relay_exporter relay_exporter

USER ${APP_USER}
