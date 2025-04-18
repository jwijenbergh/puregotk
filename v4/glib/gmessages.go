// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the prototype of log handler functions.
//
// The default log handler, g_log_default_handler(), automatically appends a
// new-line character to @message when printing it. It is advised that any
// custom log handler functions behave similarly, so that logging calls in user
// code do not need modifying to add a new-line character to the message if the
// log handler is changed.
//
// This is not used if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
type LogFunc func(string, LogLevelFlags, string, uintptr)

// Writer function for log entries. A log entry is a collection of one or more
// #GLogFields, using the standard [field names from journal
// specification](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html).
// See g_log_structured() for more information.
//
// Writer functions must ignore fields which they do not recognise, unless they
// can write arbitrary binary output, as field values may be arbitrary binary.
//
// @log_level is guaranteed to be included in @fields as the `PRIORITY` field,
// but is provided separately for convenience of deciding whether or where to
// output the log entry.
//
// Writer functions should return %G_LOG_WRITER_HANDLED if they handled the log
// message successfully or if they deliberately ignored it. If there was an
// error handling the message (for example, if the writer function is meant to
// send messages to a remote logging server and there is a network error), it
// should return %G_LOG_WRITER_UNHANDLED. This allows writer functions to be
// chained and fall back to simpler handlers in case of failure.
type LogWriterFunc func(LogLevelFlags, []LogField, uint, uintptr) LogWriterOutput

// Specifies the type of the print handler functions.
// These are called with the complete formatted string to output.
type PrintFunc func(string)

// Structure representing a single field in a structured log entry. See
// g_log_structured() for details.
//
// Log fields may contain arbitrary values, including binary with embedded nul
// bytes. If the field contains a string, the string must be UTF-8 encoded and
// have a trailing nul byte. Otherwise, @length must be set to a non-negative
// value.
type LogField struct {
	_ structs.HostLayout

	Key uintptr

	Value uintptr

	Length int
}

func (x *LogField) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

const (
	// Defines the log domain. See [Log Domains](#log-domains).
	//
	// Libraries should define this so that any messages
	// which they log can be differentiated from messages from other
	// libraries and application code. But be careful not to define
	// it in any public header files.
	//
	// Log domains must be unique, and it is recommended that they are the
	// application or library name, optionally followed by a hyphen and a sub-domain
	// name. For example, `bloatpad` or `bloatpad-io`.
	//
	// If undefined, it defaults to the default %NULL (or `""`) log domain; this is
	// not advisable, as it cannot be filtered against using the `G_MESSAGES_DEBUG`
	// environment variable.
	//
	// For example, GTK+ uses this in its `Makefile.am`:
	// |[
	// AM_CPPFLAGS = -DG_LOG_DOMAIN=\"Gtk\"
	// ]|
	//
	// Applications can choose to leave it as the default %NULL (or `""`)
	// domain. However, defining the domain offers the same advantages as
	// above.
	LOG_DOMAIN byte = 0
	// GLib log levels that are considered fatal by default.
	//
	// This is not used if structured logging is enabled; see
	// [Using Structured Logging][using-structured-logging].
	LOG_FATAL_MASK int = 5
	// Log levels below 1&lt;&lt;G_LOG_LEVEL_USER_SHIFT are used by GLib.
	// Higher bits can be used for user-defined log levels.
	LOG_LEVEL_USER_SHIFT int = 8
)

// Flags specifying the level of log messages.
//
// It is possible to change how GLib treats messages of the various
// levels using g_log_set_handler() and g_log_set_fatal_mask().
type LogLevelFlags int

const (

	// internal flag
	GLogFlagRecursionValue LogLevelFlags = 1
	// internal flag
	GLogFlagFatalValue LogLevelFlags = 2
	// log level for errors, see g_error().
	//     This level is also used for messages produced by g_assert().
	GLogLevelErrorValue LogLevelFlags = 4
	// log level for critical warning messages, see
	//     g_critical().
	//     This level is also used for messages produced by g_return_if_fail()
	//     and g_return_val_if_fail().
	GLogLevelCriticalValue LogLevelFlags = 8
	// log level for warnings, see g_warning()
	GLogLevelWarningValue LogLevelFlags = 16
	// log level for messages, see g_message()
	GLogLevelMessageValue LogLevelFlags = 32
	// log level for informational messages, see g_info()
	GLogLevelInfoValue LogLevelFlags = 64
	// log level for debug messages, see g_debug()
	GLogLevelDebugValue LogLevelFlags = 128
	// a mask including all log levels
	GLogLevelMaskValue LogLevelFlags = -4
)

// Return values from #GLogWriterFuncs to indicate whether the given log entry
// was successfully handled by the writer, or whether there was an error in
// handling it (and hence a fallback writer should be used).
//
// If a #GLogWriterFunc ignores a log entry, it should return
// %G_LOG_WRITER_HANDLED.
type LogWriterOutput int

const (

	// Log writer has handled the log entry.
	GLogWriterHandledValue LogWriterOutput = 1
	// Log writer could not handle the log entry.
	GLogWriterUnhandledValue LogWriterOutput = 0
)

var xAssertWarning func(string, string, int, string, string)

func AssertWarning(LogDomainVar string, FileVar string, LineVar int, PrettyFunctionVar string, ExpressionVar string) {

	xAssertWarning(LogDomainVar, FileVar, LineVar, PrettyFunctionVar, ExpressionVar)

}

var xLog func(string, LogLevelFlags, string, ...interface{})

// Logs an error or debugging message.
//
// If the log level has been set as fatal, G_BREAKPOINT() is called
// to terminate the program. See the documentation for G_BREAKPOINT() for
// details of the debugging options this provides.
//
// If g_log_default_handler() is used as the log handler function, a new-line
// character will automatically be appended to @..., and need not be entered
// manually.
//
// If [structured logging is enabled][using-structured-logging] this will
// output via the structured log writer function (see g_log_set_writer_func()).
func Log(LogDomainVar string, LogLevelVar LogLevelFlags, FormatVar string, varArgs ...interface{}) {

	xLog(LogDomainVar, LogLevelVar, FormatVar, varArgs...)

}

var xLogDefaultHandler func(string, LogLevelFlags, string, uintptr)

// The default log handler set up by GLib; g_log_set_default_handler()
// allows to install an alternate default log handler.
// This is used if no log handler has been set for the particular log
// domain and log level combination. It outputs the message to stderr
// or stdout and if the log level is fatal it calls G_BREAKPOINT(). It automatically
// prints a new-line character after the message, so one does not need to be
// manually included in @message.
//
// The behavior of this log handler can be influenced by a number of
// environment variables:
//
//   - `G_MESSAGES_PREFIXED`: A :-separated list of log levels for which
//     messages should be prefixed by the program name and PID of the
//     application.
//
//   - `G_MESSAGES_DEBUG`: A space-separated list of log domains for
//     which debug and informational messages are printed. By default
//     these messages are not printed.
//
// stderr is used for levels %G_LOG_LEVEL_ERROR, %G_LOG_LEVEL_CRITICAL,
// %G_LOG_LEVEL_WARNING and %G_LOG_LEVEL_MESSAGE. stdout is used for
// the rest, unless stderr was requested by
// g_log_writer_default_set_use_stderr().
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
func LogDefaultHandler(LogDomainVar string, LogLevelVar LogLevelFlags, MessageVar string, UnusedDataVar uintptr) {

	xLogDefaultHandler(LogDomainVar, LogLevelVar, MessageVar, UnusedDataVar)

}

var xLogGetDebugEnabled func() bool

// Return whether debug output from the GLib logging system is enabled.
//
// Note that this should not be used to conditionalise calls to g_debug() or
// other logging functions; it should only be used from %GLogWriterFunc
// implementations.
//
// Note also that the value of this does not depend on `G_MESSAGES_DEBUG`; see
// the docs for g_log_set_debug_enabled().
func LogGetDebugEnabled() bool {

	cret := xLogGetDebugEnabled()
	return cret
}

var xLogRemoveHandler func(string, uint)

// Removes the log handler.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
func LogRemoveHandler(LogDomainVar string, HandlerIdVar uint) {

	xLogRemoveHandler(LogDomainVar, HandlerIdVar)

}

var xLogSetAlwaysFatal func(LogLevelFlags) LogLevelFlags

// Sets the message levels which are always fatal, in any log domain.
// When a message with any of these levels is logged the program terminates.
// You can only set the levels defined by GLib to be fatal.
// %G_LOG_LEVEL_ERROR is always fatal.
//
// You can also make some message levels fatal at runtime by setting
// the `G_DEBUG` environment variable (see
// [Running GLib Applications](glib-running.html)).
//
// Libraries should not call this function, as it affects all messages logged
// by a process, including those from other libraries.
//
// Structured log messages (using g_log_structured() and
// g_log_structured_array()) are fatal only if the default log writer is used;
// otherwise it is up to the writer function to determine which log messages
// are fatal. See [Using Structured Logging][using-structured-logging].
func LogSetAlwaysFatal(FatalMaskVar LogLevelFlags) LogLevelFlags {

	cret := xLogSetAlwaysFatal(FatalMaskVar)
	return cret
}

var xLogSetDebugEnabled func(bool)

// Enable or disable debug output from the GLib logging system for all domains.
// This value interacts disjunctively with `G_MESSAGES_DEBUG` — if either of
// them would allow a debug message to be outputted, it will be.
//
// Note that this should not be used from within library code to enable debug
// output — it is intended for external use.
func LogSetDebugEnabled(EnabledVar bool) {

	xLogSetDebugEnabled(EnabledVar)

}

var xLogSetDefaultHandler func(uintptr, uintptr) uintptr

// Installs a default log handler which is used if no
// log handler has been set for the particular log domain
// and log level combination. By default, GLib uses
// g_log_default_handler() as default log handler.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
func LogSetDefaultHandler(LogFuncVar *LogFunc, UserDataVar uintptr) uintptr {

	cret := xLogSetDefaultHandler(NewCallback(LogFuncVar), UserDataVar)
	return cret
}

var xLogSetFatalMask func(string, LogLevelFlags) LogLevelFlags

// Sets the log levels which are fatal in the given domain.
// %G_LOG_LEVEL_ERROR is always fatal.
//
// This has no effect on structured log messages (using g_log_structured() or
// g_log_structured_array()). To change the fatal behaviour for specific log
// messages, programs must install a custom log writer function using
// g_log_set_writer_func(). See
// [Using Structured Logging][using-structured-logging].
//
// This function is mostly intended to be used with
// %G_LOG_LEVEL_CRITICAL.  You should typically not set
// %G_LOG_LEVEL_WARNING, %G_LOG_LEVEL_MESSAGE, %G_LOG_LEVEL_INFO or
// %G_LOG_LEVEL_DEBUG as fatal except inside of test programs.
func LogSetFatalMask(LogDomainVar string, FatalMaskVar LogLevelFlags) LogLevelFlags {

	cret := xLogSetFatalMask(LogDomainVar, FatalMaskVar)
	return cret
}

var xLogSetHandler func(string, LogLevelFlags, uintptr, uintptr) uint

// Sets the log handler for a domain and a set of log levels.
//
// To handle fatal and recursive messages the @log_levels parameter
// must be combined with the %G_LOG_FLAG_FATAL and %G_LOG_FLAG_RECURSION
// bit flags.
//
// Note that since the %G_LOG_LEVEL_ERROR log level is always fatal, if
// you want to set a handler for this log level you must combine it with
// %G_LOG_FLAG_FATAL.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
//
// Here is an example for adding a log handler for all warning messages
// in the default domain:
//
// |[&lt;!-- language="C" --&gt;
// g_log_set_handler (NULL, G_LOG_LEVEL_WARNING | G_LOG_FLAG_FATAL
//
//	| G_LOG_FLAG_RECURSION, my_log_handler, NULL);
//
// ]|
//
// This example adds a log handler for all critical messages from GTK+:
//
// |[&lt;!-- language="C" --&gt;
// g_log_set_handler ("Gtk", G_LOG_LEVEL_CRITICAL | G_LOG_FLAG_FATAL
//
//	| G_LOG_FLAG_RECURSION, my_log_handler, NULL);
//
// ]|
//
// This example adds a log handler for all messages from GLib:
//
// |[&lt;!-- language="C" --&gt;
// g_log_set_handler ("GLib", G_LOG_LEVEL_MASK | G_LOG_FLAG_FATAL
//
//	| G_LOG_FLAG_RECURSION, my_log_handler, NULL);
//
// ]|
func LogSetHandler(LogDomainVar string, LogLevelsVar LogLevelFlags, LogFuncVar *LogFunc, UserDataVar uintptr) uint {

	cret := xLogSetHandler(LogDomainVar, LogLevelsVar, NewCallback(LogFuncVar), UserDataVar)
	return cret
}

var xLogSetHandlerFull func(string, LogLevelFlags, uintptr, uintptr, uintptr) uint

// Like g_log_set_handler(), but takes a destroy notify for the @user_data.
//
// This has no effect if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
func LogSetHandlerFull(LogDomainVar string, LogLevelsVar LogLevelFlags, LogFuncVar *LogFunc, UserDataVar uintptr, DestroyVar *DestroyNotify) uint {

	cret := xLogSetHandlerFull(LogDomainVar, LogLevelsVar, NewCallback(LogFuncVar), UserDataVar, NewCallback(DestroyVar))
	return cret
}

var xLogSetWriterFunc func(uintptr, uintptr, uintptr)

// Set a writer function which will be called to format and write out each log
// message. Each program should set a writer function, or the default writer
// (g_log_writer_default()) will be used.
//
// Libraries **must not** call this function — only programs are allowed to
// install a writer function, as there must be a single, central point where
// log messages are formatted and outputted.
//
// There can only be one writer function. It is an error to set more than one.
func LogSetWriterFunc(FuncVar *LogWriterFunc, UserDataVar uintptr, UserDataFreeVar *DestroyNotify) {

	xLogSetWriterFunc(NewCallback(FuncVar), UserDataVar, NewCallback(UserDataFreeVar))

}

var xLogStructured func(string, LogLevelFlags, ...interface{})

// Log a message with structured data.
//
// The message will be passed through to the log writer set by the application
// using g_log_set_writer_func(). If the message is fatal (i.e. its log level
// is %G_LOG_LEVEL_ERROR), the program will be aborted by calling
// G_BREAKPOINT() at the end of this function. If the log writer returns
// %G_LOG_WRITER_UNHANDLED (failure), no other fallback writers will be tried.
// See the documentation for #GLogWriterFunc for information on chaining
// writers.
//
// The structured data is provided as key–value pairs, where keys are UTF-8
// strings, and values are arbitrary pointers — typically pointing to UTF-8
// strings, but that is not a requirement. To pass binary (non-nul-terminated)
// structured data, use g_log_structured_array(). The keys for structured data
// should follow the [systemd journal
// fields](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html)
// specification. It is suggested that custom keys are namespaced according to
// the code which sets them. For example, custom keys from GLib all have a
// `GLIB_` prefix.
//
// The @log_domain will be converted into a `GLIB_DOMAIN` field. @log_level will
// be converted into a
// [`PRIORITY`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#PRIORITY=)
// field. The format string will have its placeholders substituted for the provided
// values and be converted into a
// [`MESSAGE`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#MESSAGE=)
// field.
//
// Other fields you may commonly want to pass into this function:
//
//   - [`MESSAGE_ID`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#MESSAGE_ID=)
//   - [`CODE_FILE`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#CODE_FILE=)
//   - [`CODE_LINE`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#CODE_LINE=)
//   - [`CODE_FUNC`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#CODE_FUNC=)
//   - [`ERRNO`](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#ERRNO=)
//
// Note that `CODE_FILE`, `CODE_LINE` and `CODE_FUNC` are automatically set by
// the logging macros, G_DEBUG_HERE(), g_message(), g_warning(), g_critical(),
// g_error(), etc, if the symbols `G_LOG_USE_STRUCTURED` is defined before including
// `glib.h`.
//
// For example:
//
// |[&lt;!-- language="C" --&gt;
// g_log_structured (G_LOG_DOMAIN, G_LOG_LEVEL_DEBUG,
//
//	"MESSAGE_ID", "06d4df59e6c24647bfe69d2c27ef0b4e",
//	"MY_APPLICATION_CUSTOM_FIELD", "some debug string",
//	"MESSAGE", "This is a debug message about pointer %p and integer %u.",
//	some_pointer, some_integer);
//
// ]|
//
// Note that each `MESSAGE_ID` must be [uniquely and randomly
// generated](https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#MESSAGE_ID=).
// If adding a `MESSAGE_ID`, consider shipping a [message
// catalog](https://www.freedesktop.org/wiki/Software/systemd/catalog/) with
// your software.
//
// To pass a user data pointer to the log writer function which is specific to
// this logging call, you must use g_log_structured_array() and pass the pointer
// as a field with #GLogField.length set to zero, otherwise it will be
// interpreted as a string.
//
// For example:
//
// |[&lt;!-- language="C" --&gt;
//
//	const GLogField fields[] = {
//	  { "MESSAGE", "This is a debug message.", -1 },
//	  { "MESSAGE_ID", "fcfb2e1e65c3494386b74878f1abf893", -1 },
//	  { "MY_APPLICATION_CUSTOM_FIELD", "some debug string", -1 },
//	  { "MY_APPLICATION_STATE", state_object, 0 },
//	};
//
// g_log_structured_array (G_LOG_LEVEL_DEBUG, fields, G_N_ELEMENTS (fields));
// ]|
//
// Note also that, even if no other structured fields are specified, there
// must always be a `MESSAGE` key before the format string. The `MESSAGE`-format
// pair has to be the last of the key-value pairs, and `MESSAGE` is the only
// field for which printf()-style formatting is supported.
//
// The default writer function for `stdout` and `stderr` will automatically
// append a new-line character after the message, so you should not add one
// manually to the format string.
func LogStructured(LogDomainVar string, LogLevelVar LogLevelFlags, varArgs ...interface{}) {

	xLogStructured(LogDomainVar, LogLevelVar, varArgs...)

}

var xLogStructuredArray func(LogLevelFlags, []LogField, uint)

// Log a message with structured data. The message will be passed through to the
// log writer set by the application using g_log_set_writer_func(). If the
// message is fatal (i.e. its log level is %G_LOG_LEVEL_ERROR), the program will
// be aborted at the end of this function.
//
// See g_log_structured() for more documentation.
//
// This assumes that @log_level is already present in @fields (typically as the
// `PRIORITY` field).
func LogStructuredArray(LogLevelVar LogLevelFlags, FieldsVar []LogField, NFieldsVar uint) {

	xLogStructuredArray(LogLevelVar, FieldsVar, NFieldsVar)

}

var xLogStructuredStandard func(string, LogLevelFlags, string, string, string, string, ...interface{})

func LogStructuredStandard(LogDomainVar string, LogLevelVar LogLevelFlags, FileVar string, LineVar string, FuncVar string, MessageFormatVar string, varArgs ...interface{}) {

	xLogStructuredStandard(LogDomainVar, LogLevelVar, FileVar, LineVar, FuncVar, MessageFormatVar, varArgs...)

}

var xLogVariant func(string, LogLevelFlags, *Variant)

// Log a message with structured data, accepting the data within a #GVariant. This
// version is especially useful for use in other languages, via introspection.
//
// The only mandatory item in the @fields dictionary is the "MESSAGE" which must
// contain the text shown to the user.
//
// The values in the @fields dictionary are likely to be of type String
// (%G_VARIANT_TYPE_STRING). Array of bytes (%G_VARIANT_TYPE_BYTESTRING) is also
// supported. In this case the message is handled as binary and will be forwarded
// to the log writer as such. The size of the array should not be higher than
// %G_MAXSSIZE. Otherwise it will be truncated to this size. For other types
// g_variant_print() will be used to convert the value into a string.
//
// For more details on its usage and about the parameters, see g_log_structured().
func LogVariant(LogDomainVar string, LogLevelVar LogLevelFlags, FieldsVar *Variant) {

	xLogVariant(LogDomainVar, LogLevelVar, FieldsVar)

}

var xLogWriterDefault func(LogLevelFlags, []LogField, uint, uintptr) LogWriterOutput

// Format a structured log message and output it to the default log destination
// for the platform. On Linux, this is typically the systemd journal, falling
// back to `stdout` or `stderr` if running from the terminal or if output is
// being redirected to a file.
//
// Support for other platform-specific logging mechanisms may be added in
// future. Distributors of GLib may modify this function to impose their own
// (documented) platform-specific log writing policies.
//
// This is suitable for use as a #GLogWriterFunc, and is the default writer used
// if no other is set using g_log_set_writer_func().
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or `all`) is listed in the space-separated
// `G_MESSAGES_DEBUG` environment variable.
//
// g_log_writer_default() uses the mask set by g_log_set_always_fatal() to
// determine which messages are fatal. When using a custom writer func instead it is
// up to the writer function to determine which log messages are fatal.
func LogWriterDefault(LogLevelVar LogLevelFlags, FieldsVar []LogField, NFieldsVar uint, UserDataVar uintptr) LogWriterOutput {

	cret := xLogWriterDefault(LogLevelVar, FieldsVar, NFieldsVar, UserDataVar)
	return cret
}

var xLogWriterDefaultSetUseStderr func(bool)

// Configure whether the built-in log functions
// (g_log_default_handler() for the old-style API, and both
// g_log_writer_default() and g_log_writer_standard_streams() for the
// structured API) will output all log messages to `stderr`.
//
// By default, log messages of levels %G_LOG_LEVEL_INFO and
// %G_LOG_LEVEL_DEBUG are sent to `stdout`, and other log messages are
// sent to `stderr`. This is problematic for applications that intend
// to reserve `stdout` for structured output such as JSON or XML.
//
// This function sets global state. It is not thread-aware, and should be
// called at the very start of a program, before creating any other threads
// or creating objects that could create worker threads of their own.
func LogWriterDefaultSetUseStderr(UseStderrVar bool) {

	xLogWriterDefaultSetUseStderr(UseStderrVar)

}

var xLogWriterDefaultWouldDrop func(LogLevelFlags, string) bool

// Check whether g_log_writer_default() and g_log_default_handler() would
// ignore a message with the given domain and level.
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or `all`) is listed in the space-separated
// `G_MESSAGES_DEBUG` environment variable.
//
// This can be used when implementing log writers with the same filtering
// behaviour as the default, but a different destination or output format:
//
// |[&lt;!-- language="C" --&gt;
//
//	if (g_log_writer_default_would_drop (log_level, log_domain))
//	  return G_LOG_WRITER_HANDLED;
//
// ]|
//
// or to skip an expensive computation if it is only needed for a debugging
// message, and `G_MESSAGES_DEBUG` is not set:
//
// |[&lt;!-- language="C" --&gt;
//
//	if (!g_log_writer_default_would_drop (G_LOG_LEVEL_DEBUG, G_LOG_DOMAIN))
//	  {
//	    gchar *result = expensive_computation (my_object);
//
//	    g_debug ("my_object result: %s", result);
//	    g_free (result);
//	  }
//
// ]|
func LogWriterDefaultWouldDrop(LogLevelVar LogLevelFlags, LogDomainVar string) bool {

	cret := xLogWriterDefaultWouldDrop(LogLevelVar, LogDomainVar)
	return cret
}

var xLogWriterFormatFields func(LogLevelFlags, []LogField, uint, bool) string

// Format a structured log message as a string suitable for outputting to the
// terminal (or elsewhere). This will include the values of all fields it knows
// how to interpret, which includes `MESSAGE` and `GLIB_DOMAIN` (see the
// documentation for g_log_structured()). It does not include values from
// unknown fields.
//
// The returned string does **not** have a trailing new-line character. It is
// encoded in the character set of the current locale, which is not necessarily
// UTF-8.
func LogWriterFormatFields(LogLevelVar LogLevelFlags, FieldsVar []LogField, NFieldsVar uint, UseColorVar bool) string {

	cret := xLogWriterFormatFields(LogLevelVar, FieldsVar, NFieldsVar, UseColorVar)
	return cret
}

var xLogWriterIsJournald func(int) bool

// Check whether the given @output_fd file descriptor is a connection to the
// systemd journal, or something else (like a log file or `stdout` or
// `stderr`).
//
// Invalid file descriptors are accepted and return %FALSE, which allows for
// the following construct without needing any additional error handling:
// |[&lt;!-- language="C" --&gt;
//
//	is_journald = g_log_writer_is_journald (fileno (stderr));
//
// ]|
func LogWriterIsJournald(OutputFdVar int) bool {

	cret := xLogWriterIsJournald(OutputFdVar)
	return cret
}

var xLogWriterJournald func(LogLevelFlags, []LogField, uint, uintptr) LogWriterOutput

// Format a structured log message and send it to the systemd journal as a set
// of key–value pairs. All fields are sent to the journal, but if a field has
// length zero (indicating program-specific data) then only its key will be
// sent.
//
// This is suitable for use as a #GLogWriterFunc.
//
// If GLib has been compiled without systemd support, this function is still
// defined, but will always return %G_LOG_WRITER_UNHANDLED.
func LogWriterJournald(LogLevelVar LogLevelFlags, FieldsVar []LogField, NFieldsVar uint, UserDataVar uintptr) LogWriterOutput {

	cret := xLogWriterJournald(LogLevelVar, FieldsVar, NFieldsVar, UserDataVar)
	return cret
}

var xLogWriterStandardStreams func(LogLevelFlags, []LogField, uint, uintptr) LogWriterOutput

// Format a structured log message and print it to either `stdout` or `stderr`,
// depending on its log level. %G_LOG_LEVEL_INFO and %G_LOG_LEVEL_DEBUG messages
// are sent to `stdout`, or to `stderr` if requested by
// g_log_writer_default_set_use_stderr();
// all other log levels are sent to `stderr`. Only fields
// which are understood by this function are included in the formatted string
// which is printed.
//
// If the output stream supports ANSI color escape sequences, they will be used
// in the output.
//
// A trailing new-line character is added to the log message when it is printed.
//
// This is suitable for use as a #GLogWriterFunc.
func LogWriterStandardStreams(LogLevelVar LogLevelFlags, FieldsVar []LogField, NFieldsVar uint, UserDataVar uintptr) LogWriterOutput {

	cret := xLogWriterStandardStreams(LogLevelVar, FieldsVar, NFieldsVar, UserDataVar)
	return cret
}

var xLogWriterSupportsColor func(int) bool

// Check whether the given @output_fd file descriptor supports ANSI color
// escape sequences. If so, they can safely be used when formatting log
// messages.
func LogWriterSupportsColor(OutputFdVar int) bool {

	cret := xLogWriterSupportsColor(OutputFdVar)
	return cret
}

var xLogv func(string, LogLevelFlags, string, []interface{})

// Logs an error or debugging message.
//
// If the log level has been set as fatal, G_BREAKPOINT() is called
// to terminate the program. See the documentation for G_BREAKPOINT() for
// details of the debugging options this provides.
//
// If g_log_default_handler() is used as the log handler function, a new-line
// character will automatically be appended to @..., and need not be entered
// manually.
//
// If [structured logging is enabled][using-structured-logging] this will
// output via the structured log writer function (see g_log_set_writer_func()).
func Logv(LogDomainVar string, LogLevelVar LogLevelFlags, FormatVar string, ArgsVar []interface{}) {

	xLogv(LogDomainVar, LogLevelVar, FormatVar, ArgsVar)

}

var xPrint func(string, ...interface{})

// Outputs a formatted message via the print handler.
// The default print handler simply outputs the message to stdout, without
// appending a trailing new-line character. Typically, @format should end with
// its own new-line character.
//
// g_print() should not be used from within libraries for debugging
// messages, since it may be redirected by applications to special
// purpose message windows or even files. Instead, libraries should
// use g_log(), g_log_structured(), or the convenience macros g_message(),
// g_warning() and g_error().
func Print(FormatVar string, varArgs ...interface{}) {

	xPrint(FormatVar, varArgs...)

}

var xPrinterr func(string, ...interface{})

// Outputs a formatted message via the error message handler.
// The default handler simply outputs the message to stderr, without appending
// a trailing new-line character. Typically, @format should end with its own
// new-line character.
//
// g_printerr() should not be used from within libraries.
// Instead g_log() or g_log_structured() should be used, or the convenience
// macros g_message(), g_warning() and g_error().
func Printerr(FormatVar string, varArgs ...interface{}) {

	xPrinterr(FormatVar, varArgs...)

}

var xPrintfStringUpperBound func(string, []interface{}) uint

// Calculates the maximum space needed to store the output
// of the sprintf() function.
func PrintfStringUpperBound(FormatVar string, ArgsVar []interface{}) uint {

	cret := xPrintfStringUpperBound(FormatVar, ArgsVar)
	return cret
}

var xReturnIfFailWarning func(string, string, string)

// Internal function used to print messages from the public g_return_if_fail()
// and g_return_val_if_fail() macros.
func ReturnIfFailWarning(LogDomainVar string, PrettyFunctionVar string, ExpressionVar string) {

	xReturnIfFailWarning(LogDomainVar, PrettyFunctionVar, ExpressionVar)

}

var xSetPrintHandler func(uintptr) uintptr

// Sets the print handler.
//
// Any messages passed to g_print() will be output via
// the new handler. The default handler simply outputs
// the message to stdout. By providing your own handler
// you can redirect the output, to a GTK+ widget or a
// log file for example.
func SetPrintHandler(FuncVar *PrintFunc) uintptr {

	cret := xSetPrintHandler(NewCallback(FuncVar))
	return cret
}

var xSetPrinterrHandler func(uintptr) uintptr

// Sets the handler for printing error messages.
//
// Any messages passed to g_printerr() will be output via
// the new handler. The default handler simply outputs the
// message to stderr. By providing your own handler you can
// redirect the output, to a GTK+ widget or a log file for
// example.
func SetPrinterrHandler(FuncVar *PrintFunc) uintptr {

	cret := xSetPrinterrHandler(NewCallback(FuncVar))
	return cret
}

var xWarnMessage func(string, string, int, string, string)

// Internal function used to print messages from the public g_warn_if_reached()
// and g_warn_if_fail() macros.
func WarnMessage(DomainVar string, FileVar string, LineVar int, FuncVar string, WarnexprVar string) {

	xWarnMessage(DomainVar, FileVar, LineVar, FuncVar, WarnexprVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAssertWarning, lib, "g_assert_warning")
	core.PuregoSafeRegister(&xLog, lib, "g_log")
	core.PuregoSafeRegister(&xLogDefaultHandler, lib, "g_log_default_handler")
	core.PuregoSafeRegister(&xLogGetDebugEnabled, lib, "g_log_get_debug_enabled")
	core.PuregoSafeRegister(&xLogRemoveHandler, lib, "g_log_remove_handler")
	core.PuregoSafeRegister(&xLogSetAlwaysFatal, lib, "g_log_set_always_fatal")
	core.PuregoSafeRegister(&xLogSetDebugEnabled, lib, "g_log_set_debug_enabled")
	core.PuregoSafeRegister(&xLogSetDefaultHandler, lib, "g_log_set_default_handler")
	core.PuregoSafeRegister(&xLogSetFatalMask, lib, "g_log_set_fatal_mask")
	core.PuregoSafeRegister(&xLogSetHandler, lib, "g_log_set_handler")
	core.PuregoSafeRegister(&xLogSetHandlerFull, lib, "g_log_set_handler_full")
	core.PuregoSafeRegister(&xLogSetWriterFunc, lib, "g_log_set_writer_func")
	core.PuregoSafeRegister(&xLogStructured, lib, "g_log_structured")
	core.PuregoSafeRegister(&xLogStructuredArray, lib, "g_log_structured_array")
	core.PuregoSafeRegister(&xLogStructuredStandard, lib, "g_log_structured_standard")
	core.PuregoSafeRegister(&xLogVariant, lib, "g_log_variant")
	core.PuregoSafeRegister(&xLogWriterDefault, lib, "g_log_writer_default")
	core.PuregoSafeRegister(&xLogWriterDefaultSetUseStderr, lib, "g_log_writer_default_set_use_stderr")
	core.PuregoSafeRegister(&xLogWriterDefaultWouldDrop, lib, "g_log_writer_default_would_drop")
	core.PuregoSafeRegister(&xLogWriterFormatFields, lib, "g_log_writer_format_fields")
	core.PuregoSafeRegister(&xLogWriterIsJournald, lib, "g_log_writer_is_journald")
	core.PuregoSafeRegister(&xLogWriterJournald, lib, "g_log_writer_journald")
	core.PuregoSafeRegister(&xLogWriterStandardStreams, lib, "g_log_writer_standard_streams")
	core.PuregoSafeRegister(&xLogWriterSupportsColor, lib, "g_log_writer_supports_color")
	core.PuregoSafeRegister(&xLogv, lib, "g_logv")
	core.PuregoSafeRegister(&xPrint, lib, "g_print")
	core.PuregoSafeRegister(&xPrinterr, lib, "g_printerr")
	core.PuregoSafeRegister(&xPrintfStringUpperBound, lib, "g_printf_string_upper_bound")
	core.PuregoSafeRegister(&xReturnIfFailWarning, lib, "g_return_if_fail_warning")
	core.PuregoSafeRegister(&xSetPrintHandler, lib, "g_set_print_handler")
	core.PuregoSafeRegister(&xSetPrinterrHandler, lib, "g_set_printerr_handler")
	core.PuregoSafeRegister(&xWarnMessage, lib, "g_warn_message")

}
