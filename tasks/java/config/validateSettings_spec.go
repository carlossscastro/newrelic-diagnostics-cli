package config

import (
	"strings"

	"github.com/newrelic/newrelic-diagnostics-cli/tasks/base/config"
)

func LoadSpec() map[string]interface{} {
	specBlob, _ := config.ParseYaml(strings.NewReader(spec))
	return specBlob.AsMap()
}

var spec = `
# This file configures the New Relic Agent.  New Relic monitors
# Java applications with deep visibility and low overhead.  For more details and additional
# configuration options visit https://docs.newrelic.com/docs/agents/java-agent/configuration/java-agent-configuration-config-file.
#
# <%= generated_for_user %>
#
# This section is for settings common to all environments.
# Do not add anything above this next line.
common: &default_settings

  # ============================== LICENSE KEY ===============================
  # You must specify the license key associated with your New Relic
  # account. For example, if your license key is 12345 use this:
  # license_key: '12345'
  # The key binds your Agent's data to your account in the New Relic service.
  license_key: String

  # Agent Enabled
  # Use this setting to disable the agent instead of removing it from the startup command.
  # Default is true.
  agent_enabled: Boolean

  # Set the name of your application as you'd like it show up in New Relic.
  # If enable_auto_app_naming is false, the agent reports all data to this application.
  # Otherwise, the agent reports only background tasks (transactions for non-web applications)
  # to this application. To report data to more than one application
  # (useful for rollup reporting), separate the application names with ";".
  # For example, to report data to "My Application" and "My Application 2" use this:
  # app_name: My Application;My Application 2
  # This setting is required. Up to 3 different application names can be specified.
  # The first application name must be unique.
  app_name: AppName

  # To enable high security, set this property to true. When in high
  # security mode, the agent will use SSL and obfuscated SQL. Additionally,
  # request parameters and message parameters will not be sent to New Relic.
  high_security: Boolean

  # Set to true to enable support for auto app naming.
  # The name of each web app is detected automatically
  # and the agent reports data separately for each one.
  # This provides a finer-grained performance breakdown for
  # web apps in New Relic.
  # Default is false.
  enable_auto_app_naming: Boolean

  # Set to true to enable component-based transaction naming.
  # Set to false to use the URI of a web request as the name of the transaction.
  # Default is true.
  enable_auto_transaction_naming: Boolean

  # The agent uses its own log file to keep its logging
  # separate from that of your application.  Specify the log level here.
  # This setting is dynamic, so changes do not require restarting your application.
  # The levels in increasing order of verboseness are:
  #   off, severe, warning, info, fine, finer, finest
  # Default is info.
  log_level: LogLevel

  # Log all data sent to and from New Relic in plain text.
  # This setting is dynamic, so changes do not require restarting your application.
  # Default is false.
  audit_mode: Boolean

  # The number of backup log files to save.
  # Default is 1.
  log_file_count: Integer

  # The maximum number of kbytes to write to any one log file.
  # The log_file_count must be set greater than 1.
  # Default is 0 (no limit).
  log_limit_in_kbytes: Integer

  # Override other log rolling configuration and roll the logs daily.
  # Default is false.
  log_daily: Boolean

  # The name of the log file.
  # Default is newrelic_agent.log.
  log_file_name: String

  # The log file directory.
  # Default is the logs directory in the newrelic.jar parent directory.
  log_file_path: String

  # Proxy settings for connecting to the New Relic server:
  # If a proxy is used, the host setting is required.  Other settings
  # are optional.  Default port is 8080.  The username and password
  # settings will be used to authenticate to Basic Auth challenges
  # from a proxy server. Proxy scheme will allow the agent to
  # connect through proxies using the HTTPS scheme.
  proxy_host: String
  proxy_port: Integer
  proxy_user: String
  proxy_password: String
  proxy_scheme: ProxyScheme

  # Limits the number of lines to capture for each stack trace.
  # Default is 30
  max_stack_trace_lines: Integer

  # Provides the ability to configure the attributes sent to New Relic. These
  # attributes can be found in transaction traces, traced errors, Insight's
  # transaction events, and Insight's page views.
  attributes:

    # When true, attributes will be sent to New Relic. The default is true.
    enabled: Boolean

    #A comma separated list of attribute keys whose values should
    # be sent to New Relic.
    include: CommaSeparatedStringList

    # A comma separated list of attribute keys whose values should
    # not be sent to New Relic.
    exclude: CommaSeparatedStringList


  # Transaction tracer captures deep information about slow
  # transactions and sends this to the New Relic service once a
  # minute. Included in the transaction is the exact call sequence of
  # the transactions including any SQL statements issued.
  transaction_tracer:

    # Transaction tracer is enabled by default. Set this to false to turn it off.
    # This feature is not available to Lite accounts and is automatically disabled.
    # Default is true.
    enabled: Boolean

    # Threshold in seconds for when to collect a transaction
    # trace. When the response time of a controller action exceeds
    # this threshold, a transaction trace will be recorded and sent to
    # New Relic. Valid values are any float value, or (default) "apdex_f",
    # which will use the threshold for the "Frustrated" Apdex level
    # (greater than four times the apdex_t value).
    # Default is apdex_f.
    transaction_threshold: TransactionThreshold

    # When transaction tracer is on, SQL statements can optionally be
    # recorded. The recorder has three modes, "off" which sends no
    # SQL, "raw" which sends the SQL statement in its original form,
    # and "obfuscated", which strips out numeric and string literals.
    # Default is obfuscated.
    record_sql: RecordSql

    # Set this to true to log SQL statements instead of recording them.
    # SQL is logged using the record_sql mode.
    # Default is false.
    log_sql: Boolean

    # Threshold in seconds for when to collect stack trace for a SQL
    # call. In other words, when SQL statements exceed this threshold,
    # then capture and send to New Relic the current stack trace. This is
    # helpful for pinpointing where long SQL calls originate from.
    # Default is 0.5 seconds.
    stack_trace_threshold: Float

    # Determines whether the agent will capture query plans for slow
    # SQL queries. Only supported for MySQL and PostgreSQL.
    # Default is true.
    explain_enabled: Boolean

    # Threshold for query execution time below which query plans will not
    # not be captured.  Relevant only when explain_enabled is true.
    # Default is 0.5 seconds.
    explain_threshold: Float

    # Use this setting to control the variety of transaction traces.
    # The higher the setting, the greater the variety.
    # Set this to 0 to always report the slowest transaction trace.
    # Default is 20.
    top_n: Integer

  # Error collector captures information about uncaught exceptions and
  # sends them to New Relic for viewing.
  error_collector:

    # This property enables the collection of errors. If the property is not
    # set or the property is set to false, then errors will not be collected.
    # Default is true.
    enabled: Boolean

    # Use this property to exclude specific exceptions from being reported as errors
    # by providing a comma separated list of full class names.
    # The default is to exclude akka.actor.ActorKilledException. If you want to override
    # this, you must provide any new value as an empty list is ignored.
    #
    # NOTE: this can be a list of strings
    ignore_errors: String # deprecated
    ignore_classes: String # new version

    # Use this property to exclude specific http status codes from being reported as errors
    # by providing a comma separated list of status codes.
    # The default is to exclude 404s. If you want to override
    # this, you must provide any new value as an empty list is ignored.
    ignore_status_codes: StatusCodeList

  # Transaction Events are used for Histograms and Percentiles. Unaggregated data is collected
  # for each web transaction and sent to the server on harvest. 
  transaction_events:

    # Set to false to disable transaction events.
    # Default is true.
    enabled: Boolean

    # Events are collected up to the configured amount. Afterwards, events are sampled to
    # maintain an even distribution across the harvest cycle.
    # Default is 2000.  Setting to 0 will disable.
    max_samples_stored: Integer

  # Distributed tracing lets you see the path that a request takes through your distributed system.
  # Enabling distributed tracing changes the behavior of some New Relic features, so carefully consult the transition
  # guide before you enable this feature: https://docs.newrelic.com/docs/apm/distributed-tracing/getting-started/transition-guide-distributed-tracing
  # Default is false.
  distributed_tracing:
    enabled: Boolean

  # Cross Application Tracing adds request and response headers to
  # external calls using supported HTTP libraries to provide better
  # performance data when calling applications monitored by other New Relic Agents.
  cross_application_tracer:

    # Set to false to disable cross application tracing.
    # Default is true.
    enabled: Boolean

  # Thread profiler measures wall clock time, CPU time, and method call counts
  # in your application's threads as they run.
  # This feature is not available to Lite accounts and is automatically disabled.
  thread_profiler:

    # Set to false to disable the thread profiler.
    # Default is true.
    enabled: Boolean

  # New Relic Real User Monitoring gives you insight into the performance real users are
  # experiencing with your website. This is accomplished by measuring the time it takes for
  # your users' browsers to download and render your web pages by injecting a small amount
  # of JavaScript code into the header and footer of each page. 
  browser_monitoring:

    # By default the agent automatically inserts API calls in compiled JSPs to
    # inject the monitoring JavaScript into web pages. Not all rendering engines are supported.
    # See https://docs.newrelic.com/docs/agents/java-agent/instrumentation/new-relic-browser-java-agent#manual_instrumentation
    # for instructions to add these manually to your pages.
    # Set this attribute to false to turn off this behavior.
    auto_instrument: Boolean

  class_transformer:
    # This instrumentation reports the name of the user principal returned from 
    # HttpServletRequest.getUserPrincipal() when servlets and filters are invoked.
    com.newrelic.instrumentation.servlet-user:
      enabled: Boolean

    com.newrelic.instrumentation.spring-aop-2:
      enabled: Boolean

    # This instrumentation reports metrics for resultset operations.
    com.newrelic.instrumentation.jdbc-resultset:
      enabled: Boolean

    # Classes loaded by classloaders in this list will not be instrumented.
    # This is a useful optimization for runtimes which use classloaders to
    # load dynamic classes which the agent would not instrument.
    classloader_excludes: String

  # User-configurable custom labels for this agent.  Labels are name-value pairs.
  # There is a maximum of 64 labels per agent.  Names and values are limited to 255 characters.
  # Names and values may not contain colons (:) or semicolons (;).
  labels: LabelList

    # An example label
    #label_name: label_value


# Application Environments
# ------------------------------------------
# Environment specific settings are in this section.
# You can use the environment to override the default settings.
# For example, to change the app_name setting.
# Use -Dnewrelic.environment=<environment> on the Java startup command line
# to set the environment.
# The default environment is production.

# NOTE if your application has other named environments, you should
# provide configuration settings for these environments here.

development:
  <<: *default_settings
  app_name: AppName

test:
  <<: *default_settings
  app_name: AppName

production:
  <<: *default_settings

staging:
  <<: *default_settings
  app_name: AppName
`
