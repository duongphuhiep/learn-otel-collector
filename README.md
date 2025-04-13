# Local Development Logging solution

As a .NET and Go developer, I want to use [VictoriaLogs]/[LogsQL] to troubleshoot my local development.

=> I will have to make my Go and .NET apps send logs to [VictoriaLogs]. However, to keep the standard with no vendor lock-in, my apps will send logs to an [Otel-collector] which will forward the logs to [VictoriaLogs].

## This repo includes

* A Docker Compose file to start an [Otel-collector] and [VictoriaLogs]. All logs received by the [Otel-collector] will be exported to the [VictoriaLogs] database.
* A Golang (console) app shows how to write logs to the [Otel-collector].
* A .NET (console) app shows how to write logs to the [Otel-collector].

## How It Works

1. The [Otel-collector] acts as a centralized logging agent that receives logs from the .NET and Go apps.
2. The [Otel-collector] forwards these logs to [VictoriaLogs], where they can be queried and analyse using [LogsQL].

Now, I can (ask AI to) add bunch of logs to troubleshoot my application, and comfortably visualize or query them in the VictoriaLogs UI: <http://localhost:9428/select/vmui>

## How to test the otel-collector

* In the [Otel-collector's config](./otel-collector-config.yaml), we configured the logs output to the debug (stdout)exporter
* Send a random log item and check if it is printed to the stdout of the otel collector container.

```sh
curl -X POST http://localhost:4318/v1/logs \
  -H "Content-Type: application/json" \
  -d '{
    "resourceLogs": [{
      "resource": {
        "attributes": [{
          "key": "service.name",
          "value": { "stringValue": "test-service" }
        }]
      },
      "scopeLogs": [{
        "logRecords": [{
          "timeUnixNano": '"$(date +%s000000000)"',
          "severityText": "INFO",
          "body": { "stringValue": "This is a test log message from curl" },
          "attributes": [{
            "key": "http.route",
            "value": { "stringValue": "/test" }
          }]
        }]
      }]
    }]
  }'
```

* We can also use [telemetrygen] to send a bunch of random telemetry data (logs, trace, metric) to the otel-collector:

```sh
go install github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen@latest
telemetrygen logs --duration 1s --otlp-insecure
```

[LogsQL]: https://docs.victoriametrics.com/victorialogs/logsql/
[VictoriaLogs]: https://docs.victoriametrics.com/victorialogs/
[Otel-collector]: https://www.youtube.com/watch?v=_CJrFW_yjRo
[telemetrygen]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/cmd/telemetrygen
