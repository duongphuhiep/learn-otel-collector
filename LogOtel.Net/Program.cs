using Microsoft.Extensions.Logging;
using OpenTelemetry.Logs;

using var loggerFactory = LoggerFactory.Create(builder =>
{
	builder.AddOpenTelemetry(options =>
	{
		options.AddOtlpExporter();
	});
	builder.AddConsole();
});

var logger = loggerFactory.CreateLogger<Program>();
logger.LogInformation("Hello Otel Collector");
