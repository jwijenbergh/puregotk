// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Application Information interface, for operating system portability.
type AppInfoIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *AppInfoIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type AppLaunchContextClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *AppLaunchContextClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type AppLaunchContextPrivate struct {
	_ structs.HostLayout
}

func (x *AppLaunchContextPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GAppInfo and #GAppLaunchContext are used for describing and launching
// applications installed on the system.
//
// As of GLib 2.20, URIs will always be converted to POSIX paths
// (using g_file_get_path()) when using g_app_info_launch() even if
// the application requested an URI and not a POSIX path. For example
// for a desktop-file based application with Exec key `totem
// %U` and a single URI, `sftp://foo/file.avi`, then
// `/home/user/.gvfs/sftp on foo/file.avi` will be passed. This will
// only work if a set of suitable GIO extensions (such as gvfs 2.26
// compiled with FUSE support), is available and operational; if this
// is not the case, the URI will be passed unmodified to the application.
// Some URIs, such as `mailto:`, of course cannot be mapped to a POSIX
// path (in gvfs there's no FUSE mount for it); such URIs will be
// passed unmodified to the application.
//
// Specifically for gvfs 2.26 and later, the POSIX URI will be mapped
// back to the GIO URI in the #GFile constructors (since gvfs
// implements the #GVfs extension point). As such, if the application
// needs to examine the URI, it needs to use g_file_get_uri() or
// similar on #GFile. In other words, an application cannot assume
// that the URI passed to e.g. g_file_new_for_commandline_arg() is
// equal to the result of g_file_get_uri(). The following snippet
// illustrates this:
//
// |[
// GFile *f;
// char *uri;
//
// file = g_file_new_for_commandline_arg (uri_from_commandline);
//
// uri = g_file_get_uri (file);
// strcmp (uri, uri_from_commandline) == 0;
// g_free (uri);
//
// if (g_file_has_uri_scheme (file, "cdda"))
//
//	{
//	  // do something special with uri
//	}
//
// g_object_unref (file);
// ]|
//
// This code will work when both `cdda://sr0/Track 1.wav` and
// `/home/user/.gvfs/cdda on sr0/Track 1.wav` is passed to the
// application. It should be noted that it's generally not safe
// for applications to rely on the format of a particular URIs.
// Different launcher applications (e.g. file managers) may have
// different ideas of what a given URI means.
type AppInfo interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	AddSupportsType(ContentTypeVar string) bool
	CanDelete() bool
	CanRemoveSupportsType() bool
	Delete() bool
	Dup() *AppInfoBase
	Equal(Appinfo2Var AppInfo) bool
	GetCommandline() string
	GetDescription() string
	GetDisplayName() string
	GetExecutable() string
	GetIcon() *IconBase
	GetId() string
	GetName() string
	GetSupportedTypes() []string
	Launch(FilesVar *glib.List, ContextVar *AppLaunchContext) bool
	LaunchUris(UrisVar *glib.List, ContextVar *AppLaunchContext) bool
	LaunchUrisAsync(UrisVar *glib.List, ContextVar *AppLaunchContext, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	LaunchUrisFinish(ResultVar AsyncResult) bool
	RemoveSupportsType(ContentTypeVar string) bool
	SetAsDefaultForExtension(ExtensionVar string) bool
	SetAsDefaultForType(ContentTypeVar string) bool
	SetAsLastUsedForType(ContentTypeVar string) bool
	ShouldShow() bool
	SupportsFiles() bool
	SupportsUris() bool
}

var xAppInfoGLibType func() types.GType

func AppInfoGLibType() types.GType {
	return xAppInfoGLibType()
}

type AppInfoBase struct {
	Ptr uintptr
}

func (x *AppInfoBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *AppInfoBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Adds a content type to the application information to indicate the
// application is capable of opening files with the given content type.
func (x *AppInfoBase) AddSupportsType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoAddSupportsType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Obtains the information whether the #GAppInfo can be deleted.
// See g_app_info_delete().
func (x *AppInfoBase) CanDelete() bool {

	cret := XGAppInfoCanDelete(x.GoPointer())
	return cret
}

// Checks if a supported content type can be removed from an application.
func (x *AppInfoBase) CanRemoveSupportsType() bool {

	cret := XGAppInfoCanRemoveSupportsType(x.GoPointer())
	return cret
}

// Tries to delete a #GAppInfo.
//
// On some platforms, there may be a difference between user-defined
// #GAppInfos which can be deleted, and system-wide ones which cannot.
// See g_app_info_can_delete().
func (x *AppInfoBase) Delete() bool {

	cret := XGAppInfoDelete(x.GoPointer())
	return cret
}

// Creates a duplicate of a #GAppInfo.
func (x *AppInfoBase) Dup() *AppInfoBase {
	var cls *AppInfoBase

	cret := XGAppInfoDup(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &AppInfoBase{}
	cls.Ptr = cret
	return cls
}

// Checks if two #GAppInfos are equal.
//
// Note that the check *may not* compare each individual
// field, and only does an identity check. In case detecting changes in the
// contents is needed, program code must additionally compare relevant fields.
func (x *AppInfoBase) Equal(Appinfo2Var AppInfo) bool {

	cret := XGAppInfoEqual(x.GoPointer(), Appinfo2Var.GoPointer())
	return cret
}

// Gets the commandline with which the application will be
// started.
func (x *AppInfoBase) GetCommandline() string {

	cret := XGAppInfoGetCommandline(x.GoPointer())
	return cret
}

// Gets a human-readable description of an installed application.
func (x *AppInfoBase) GetDescription() string {

	cret := XGAppInfoGetDescription(x.GoPointer())
	return cret
}

// Gets the display name of the application. The display name is often more
// descriptive to the user than the name itself.
func (x *AppInfoBase) GetDisplayName() string {

	cret := XGAppInfoGetDisplayName(x.GoPointer())
	return cret
}

// Gets the executable's name for the installed application.
func (x *AppInfoBase) GetExecutable() string {

	cret := XGAppInfoGetExecutable(x.GoPointer())
	return cret
}

// Gets the icon for the application.
func (x *AppInfoBase) GetIcon() *IconBase {
	var cls *IconBase

	cret := XGAppInfoGetIcon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

// Gets the ID of an application. An id is a string that
// identifies the application. The exact format of the id is
// platform dependent. For instance, on Unix this is the
// desktop file id from the xdg menu specification.
//
// Note that the returned ID may be %NULL, depending on how
// the @appinfo has been constructed.
func (x *AppInfoBase) GetId() string {

	cret := XGAppInfoGetId(x.GoPointer())
	return cret
}

// Gets the installed name of the application.
func (x *AppInfoBase) GetName() string {

	cret := XGAppInfoGetName(x.GoPointer())
	return cret
}

// Retrieves the list of content types that @app_info claims to support.
// If this information is not provided by the environment, this function
// will return %NULL.
// This function does not take in consideration associations added with
// g_app_info_add_supports_type(), but only those exported directly by
// the application.
func (x *AppInfoBase) GetSupportedTypes() []string {

	cret := XGAppInfoGetSupportedTypes(x.GoPointer())
	return cret
}

// Launches the application. Passes @files to the launched application
// as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on).
// On error, @error will be set accordingly.
//
// To launch the application without arguments pass a %NULL @files list.
//
// Note that even if the launch is successful the application launched
// can fail to start if it runs into problems during startup. There is
// no way to detect this.
//
// Some URIs can be changed when passed through a GFile (for instance
// unsupported URIs with strange formats like mailto:), so if you have
// a textual URI you want to pass in as argument, consider using
// g_app_info_launch_uris() instead.
//
// The launched application inherits the environment of the launching
// process, but it can be modified with g_app_launch_context_setenv()
// and g_app_launch_context_unsetenv().
//
// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE`
// environment variable with the path of the launched desktop file and
// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`,
// should it be inherited by further processes. The `DISPLAY` and
// `DESKTOP_STARTUP_ID` environment variables are also set, based
// on information provided in @context.
func (x *AppInfoBase) Launch(FilesVar *glib.List, ContextVar *AppLaunchContext) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoLaunch(x.GoPointer(), FilesVar, ContextVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Launches the application. This passes the @uris to the launched application
// as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on).
// On error, @error will be set accordingly.
//
// To launch the application without arguments pass a %NULL @uris list.
//
// Note that even if the launch is successful the application launched
// can fail to start if it runs into problems during startup. There is
// no way to detect this.
func (x *AppInfoBase) LaunchUris(UrisVar *glib.List, ContextVar *AppLaunchContext) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoLaunchUris(x.GoPointer(), UrisVar, ContextVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Async version of g_app_info_launch_uris().
//
// The @callback is invoked immediately after the application launch, but it
// waits for activation in case of D-Bus–activated applications and also provides
// extended error information for sandboxed applications, see notes for
// g_app_info_launch_default_for_uri_async().
func (x *AppInfoBase) LaunchUrisAsync(UrisVar *glib.List, ContextVar *AppLaunchContext, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGAppInfoLaunchUrisAsync(x.GoPointer(), UrisVar, ContextVar.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes a g_app_info_launch_uris_async() operation.
func (x *AppInfoBase) LaunchUrisFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoLaunchUrisFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Removes a supported type from an application, if possible.
func (x *AppInfoBase) RemoveSupportsType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoRemoveSupportsType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the application as the default handler for the given file extension.
func (x *AppInfoBase) SetAsDefaultForExtension(ExtensionVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoSetAsDefaultForExtension(x.GoPointer(), ExtensionVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the application as the default handler for a given type.
func (x *AppInfoBase) SetAsDefaultForType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoSetAsDefaultForType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the application as the last used application for a given type.
// This will make the application appear as first in the list returned
// by g_app_info_get_recommended_for_type(), regardless of the default
// application for that content type.
func (x *AppInfoBase) SetAsLastUsedForType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoSetAsLastUsedForType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Checks if the application info should be shown in menus that
// list available applications.
func (x *AppInfoBase) ShouldShow() bool {

	cret := XGAppInfoShouldShow(x.GoPointer())
	return cret
}

// Checks if the application accepts files as arguments.
func (x *AppInfoBase) SupportsFiles() bool {

	cret := XGAppInfoSupportsFiles(x.GoPointer())
	return cret
}

// Checks if the application supports reading files and directories from URIs.
func (x *AppInfoBase) SupportsUris() bool {

	cret := XGAppInfoSupportsUris(x.GoPointer())
	return cret
}

var XGAppInfoAddSupportsType func(uintptr, string, **glib.Error) bool
var XGAppInfoCanDelete func(uintptr) bool
var XGAppInfoCanRemoveSupportsType func(uintptr) bool
var XGAppInfoDelete func(uintptr) bool
var XGAppInfoDup func(uintptr) uintptr
var XGAppInfoEqual func(uintptr, uintptr) bool
var XGAppInfoGetCommandline func(uintptr) string
var XGAppInfoGetDescription func(uintptr) string
var XGAppInfoGetDisplayName func(uintptr) string
var XGAppInfoGetExecutable func(uintptr) string
var XGAppInfoGetIcon func(uintptr) uintptr
var XGAppInfoGetId func(uintptr) string
var XGAppInfoGetName func(uintptr) string
var XGAppInfoGetSupportedTypes func(uintptr) []string
var XGAppInfoLaunch func(uintptr, *glib.List, uintptr, **glib.Error) bool
var XGAppInfoLaunchUris func(uintptr, *glib.List, uintptr, **glib.Error) bool
var XGAppInfoLaunchUrisAsync func(uintptr, *glib.List, uintptr, uintptr, uintptr, uintptr)
var XGAppInfoLaunchUrisFinish func(uintptr, uintptr, **glib.Error) bool
var XGAppInfoRemoveSupportsType func(uintptr, string, **glib.Error) bool
var XGAppInfoSetAsDefaultForExtension func(uintptr, string, **glib.Error) bool
var XGAppInfoSetAsDefaultForType func(uintptr, string, **glib.Error) bool
var XGAppInfoSetAsLastUsedForType func(uintptr, string, **glib.Error) bool
var XGAppInfoShouldShow func(uintptr) bool
var XGAppInfoSupportsFiles func(uintptr) bool
var XGAppInfoSupportsUris func(uintptr) bool

var xAppInfoCreateFromCommandline func(string, string, AppInfoCreateFlags, **glib.Error) uintptr

// Creates a new #GAppInfo from the given information.
//
// Note that for @commandline, the quoting rules of the Exec key of the
// [freedesktop.org Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec)
// are applied. For example, if the @commandline contains
// percent-encoded URIs, the percent-character must be doubled in order to prevent it from
// being swallowed by Exec key unquoting. See the specification for exact quoting rules.
func AppInfoCreateFromCommandline(CommandlineVar string, ApplicationNameVar string, FlagsVar AppInfoCreateFlags) (*AppInfoBase, error) {
	var cls *AppInfoBase
	var cerr *glib.Error

	cret := xAppInfoCreateFromCommandline(CommandlineVar, ApplicationNameVar, FlagsVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &AppInfoBase{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xAppInfoGetAll func() *glib.List

// Gets a list of all of the applications currently registered
// on this system.
//
// For desktop files, this includes applications that have
// `NoDisplay=true` set or are excluded from display by means
// of `OnlyShowIn` or `NotShowIn`. See g_app_info_should_show().
// The returned list does not include applications which have
// the `Hidden` key set.
func AppInfoGetAll() *glib.List {

	cret := xAppInfoGetAll()
	return cret
}

var xAppInfoGetAllForType func(string) *glib.List

// Gets a list of all #GAppInfos for a given content type,
// including the recommended and fallback #GAppInfos. See
// g_app_info_get_recommended_for_type() and
// g_app_info_get_fallback_for_type().
func AppInfoGetAllForType(ContentTypeVar string) *glib.List {

	cret := xAppInfoGetAllForType(ContentTypeVar)
	return cret
}

var xAppInfoGetDefaultForType func(string, bool) uintptr

// Gets the default #GAppInfo for a given content type.
func AppInfoGetDefaultForType(ContentTypeVar string, MustSupportUrisVar bool) *AppInfoBase {
	var cls *AppInfoBase

	cret := xAppInfoGetDefaultForType(ContentTypeVar, MustSupportUrisVar)

	if cret == 0 {
		return nil
	}
	cls = &AppInfoBase{}
	cls.Ptr = cret
	return cls
}

var xAppInfoGetDefaultForUriScheme func(string) uintptr

// Gets the default application for handling URIs with
// the given URI scheme. A URI scheme is the initial part
// of the URI, up to but not including the ':', e.g. "http",
// "ftp" or "sip".
func AppInfoGetDefaultForUriScheme(UriSchemeVar string) *AppInfoBase {
	var cls *AppInfoBase

	cret := xAppInfoGetDefaultForUriScheme(UriSchemeVar)

	if cret == 0 {
		return nil
	}
	cls = &AppInfoBase{}
	cls.Ptr = cret
	return cls
}

var xAppInfoGetFallbackForType func(string) *glib.List

// Gets a list of fallback #GAppInfos for a given content type, i.e.
// those applications which claim to support the given content type
// by MIME type subclassing and not directly.
func AppInfoGetFallbackForType(ContentTypeVar string) *glib.List {

	cret := xAppInfoGetFallbackForType(ContentTypeVar)
	return cret
}

var xAppInfoGetRecommendedForType func(string) *glib.List

// Gets a list of recommended #GAppInfos for a given content type, i.e.
// those applications which claim to support the given content type exactly,
// and not by MIME type subclassing.
// Note that the first application of the list is the last used one, i.e.
// the last one for which g_app_info_set_as_last_used_for_type() has been
// called.
func AppInfoGetRecommendedForType(ContentTypeVar string) *glib.List {

	cret := xAppInfoGetRecommendedForType(ContentTypeVar)
	return cret
}

var xAppInfoLaunchDefaultForUri func(string, uintptr, **glib.Error) bool

// Utility function that launches the default application
// registered to handle the specified uri. Synchronous I/O
// is done on the uri to detect the type of the file if
// required.
//
// The D-Bus–activated applications don't have to be started if your application
// terminates too soon after this function. To prevent this, use
// g_app_info_launch_default_for_uri_async() instead.
func AppInfoLaunchDefaultForUri(UriVar string, ContextVar *AppLaunchContext) (bool, error) {
	var cerr *glib.Error

	cret := xAppInfoLaunchDefaultForUri(UriVar, ContextVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xAppInfoLaunchDefaultForUriAsync func(string, uintptr, uintptr, uintptr, uintptr)

// Async version of g_app_info_launch_default_for_uri().
//
// This version is useful if you are interested in receiving
// error information in the case where the application is
// sandboxed and the portal may present an application chooser
// dialog to the user.
//
// This is also useful if you want to be sure that the D-Bus–activated
// applications are really started before termination and if you are interested
// in receiving error information from their activation.
func AppInfoLaunchDefaultForUriAsync(UriVar string, ContextVar *AppLaunchContext, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xAppInfoLaunchDefaultForUriAsync(UriVar, ContextVar.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xAppInfoLaunchDefaultForUriFinish func(uintptr, **glib.Error) bool

// Finishes an asynchronous launch-default-for-uri operation.
func AppInfoLaunchDefaultForUriFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := xAppInfoLaunchDefaultForUriFinish(ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xAppInfoResetTypeAssociations func(string)

// Removes all changes to the type associations done by
// g_app_info_set_as_default_for_type(),
// g_app_info_set_as_default_for_extension(),
// g_app_info_add_supports_type() or
// g_app_info_remove_supports_type().
func AppInfoResetTypeAssociations(ContentTypeVar string) {

	xAppInfoResetTypeAssociations(ContentTypeVar)

}

// Integrating the launch with the launching application. This is used to
// handle for instance startup notification and launching the new application
// on the same screen as the launching window.
type AppLaunchContext struct {
	gobject.Object
}

var xAppLaunchContextGLibType func() types.GType

func AppLaunchContextGLibType() types.GType {
	return xAppLaunchContextGLibType()
}

func AppLaunchContextNewFromInternalPtr(ptr uintptr) *AppLaunchContext {
	cls := &AppLaunchContext{}
	cls.Ptr = ptr
	return cls
}

var xNewAppLaunchContext func() uintptr

// Creates a new application launch context. This is not normally used,
// instead you instantiate a subclass of this, such as #GdkAppLaunchContext.
func NewAppLaunchContext() *AppLaunchContext {
	var cls *AppLaunchContext

	cret := xNewAppLaunchContext()

	if cret == 0 {
		return nil
	}
	cls = &AppLaunchContext{}
	cls.Ptr = cret
	return cls
}

var xAppLaunchContextGetDisplay func(uintptr, uintptr, *glib.List) string

// Gets the display string for the @context. This is used to ensure new
// applications are started on the same display as the launching
// application, by setting the `DISPLAY` environment variable.
func (x *AppLaunchContext) GetDisplay(InfoVar AppInfo, FilesVar *glib.List) string {

	cret := xAppLaunchContextGetDisplay(x.GoPointer(), InfoVar.GoPointer(), FilesVar)
	return cret
}

var xAppLaunchContextGetEnvironment func(uintptr) []string

// Gets the complete environment variable list to be passed to
// the child process when @context is used to launch an application.
// This is a %NULL-terminated array of strings, where each string has
// the form `KEY=VALUE`.
func (x *AppLaunchContext) GetEnvironment() []string {

	cret := xAppLaunchContextGetEnvironment(x.GoPointer())
	return cret
}

var xAppLaunchContextGetStartupNotifyId func(uintptr, uintptr, *glib.List) string

// Initiates startup notification for the application and returns the
// `DESKTOP_STARTUP_ID` for the launched operation, if supported.
//
// Startup notification IDs are defined in the
// [FreeDesktop.Org Startup Notifications standard](http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
func (x *AppLaunchContext) GetStartupNotifyId(InfoVar AppInfo, FilesVar *glib.List) string {

	cret := xAppLaunchContextGetStartupNotifyId(x.GoPointer(), InfoVar.GoPointer(), FilesVar)
	return cret
}

var xAppLaunchContextLaunchFailed func(uintptr, string)

// Called when an application has failed to launch, so that it can cancel
// the application startup notification started in g_app_launch_context_get_startup_notify_id().
func (x *AppLaunchContext) LaunchFailed(StartupNotifyIdVar string) {

	xAppLaunchContextLaunchFailed(x.GoPointer(), StartupNotifyIdVar)

}

var xAppLaunchContextSetenv func(uintptr, string, string)

// Arranges for @variable to be set to @value in the child's
// environment when @context is used to launch an application.
func (x *AppLaunchContext) Setenv(VariableVar string, ValueVar string) {

	xAppLaunchContextSetenv(x.GoPointer(), VariableVar, ValueVar)

}

var xAppLaunchContextUnsetenv func(uintptr, string)

// Arranges for @variable to be unset in the child's environment
// when @context is used to launch an application.
func (x *AppLaunchContext) Unsetenv(VariableVar string) {

	xAppLaunchContextUnsetenv(x.GoPointer(), VariableVar)

}

func (c *AppLaunchContext) GoPointer() uintptr {
	return c.Ptr
}

func (c *AppLaunchContext) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// The #GAppLaunchContext::launch-failed signal is emitted when a #GAppInfo launch
// fails. The startup notification id is provided, so that the launcher
// can cancel the startup notification.
func (x *AppLaunchContext) ConnectLaunchFailed(cb *func(AppLaunchContext, string)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "launch-failed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, StartupNotifyIdVarp string) {
		fa := AppLaunchContext{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, StartupNotifyIdVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "launch-failed", cbRefPtr)
}

// The #GAppLaunchContext::launch-started signal is emitted when a #GAppInfo is
// about to be launched. If non-null the @platform_data is an
// GVariant dictionary mapping strings to variants (ie `a{sv}`), which
// contains additional, platform-specific data about this launch. On
// UNIX, at least the `startup-notification-id` keys will be
// present.
//
// The value of the `startup-notification-id` key (type `s`) is a startup
// notification ID corresponding to the format from the [startup-notification
// specification](https://specifications.freedesktop.org/startup-notification-spec/startup-notification-0.1.txt).
// It allows tracking the progress of the launchee through startup.
//
// It is guaranteed that this signal is followed by either a #GAppLaunchContext::launched or
// #GAppLaunchContext::launch-failed signal.
func (x *AppLaunchContext) ConnectLaunchStarted(cb *func(AppLaunchContext, uintptr, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "launch-started", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, InfoVarp uintptr, PlatformDataVarp uintptr) {
		fa := AppLaunchContext{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, InfoVarp, PlatformDataVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "launch-started", cbRefPtr)
}

// The #GAppLaunchContext::launched signal is emitted when a #GAppInfo is successfully
// launched. The @platform_data is an GVariant dictionary mapping
// strings to variants (ie `a{sv}`), which contains additional,
// platform-specific data about this launch. On UNIX, at least the
// `pid` and `startup-notification-id` keys will be present.
//
// Since 2.72 the `pid` may be 0 if the process id wasn't known (for
// example if the process was launched via D-Bus). The `pid` may not be
// set at all in subsequent releases.
func (x *AppLaunchContext) ConnectLaunched(cb *func(AppLaunchContext, uintptr, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "launched", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, InfoVarp uintptr, PlatformDataVarp uintptr) {
		fa := AppLaunchContext{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, InfoVarp, PlatformDataVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "launched", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAppInfoCreateFromCommandline, lib, "g_app_info_create_from_commandline")
	core.PuregoSafeRegister(&xAppInfoGetAll, lib, "g_app_info_get_all")
	core.PuregoSafeRegister(&xAppInfoGetAllForType, lib, "g_app_info_get_all_for_type")
	core.PuregoSafeRegister(&xAppInfoGetDefaultForType, lib, "g_app_info_get_default_for_type")
	core.PuregoSafeRegister(&xAppInfoGetDefaultForUriScheme, lib, "g_app_info_get_default_for_uri_scheme")
	core.PuregoSafeRegister(&xAppInfoGetFallbackForType, lib, "g_app_info_get_fallback_for_type")
	core.PuregoSafeRegister(&xAppInfoGetRecommendedForType, lib, "g_app_info_get_recommended_for_type")
	core.PuregoSafeRegister(&xAppInfoLaunchDefaultForUri, lib, "g_app_info_launch_default_for_uri")
	core.PuregoSafeRegister(&xAppInfoLaunchDefaultForUriAsync, lib, "g_app_info_launch_default_for_uri_async")
	core.PuregoSafeRegister(&xAppInfoLaunchDefaultForUriFinish, lib, "g_app_info_launch_default_for_uri_finish")
	core.PuregoSafeRegister(&xAppInfoResetTypeAssociations, lib, "g_app_info_reset_type_associations")

	core.PuregoSafeRegister(&xAppLaunchContextGLibType, lib, "g_app_launch_context_get_type")

	core.PuregoSafeRegister(&xNewAppLaunchContext, lib, "g_app_launch_context_new")

	core.PuregoSafeRegister(&xAppLaunchContextGetDisplay, lib, "g_app_launch_context_get_display")
	core.PuregoSafeRegister(&xAppLaunchContextGetEnvironment, lib, "g_app_launch_context_get_environment")
	core.PuregoSafeRegister(&xAppLaunchContextGetStartupNotifyId, lib, "g_app_launch_context_get_startup_notify_id")
	core.PuregoSafeRegister(&xAppLaunchContextLaunchFailed, lib, "g_app_launch_context_launch_failed")
	core.PuregoSafeRegister(&xAppLaunchContextSetenv, lib, "g_app_launch_context_setenv")
	core.PuregoSafeRegister(&xAppLaunchContextUnsetenv, lib, "g_app_launch_context_unsetenv")

	core.PuregoSafeRegister(&xAppInfoGLibType, lib, "g_app_info_get_type")

	core.PuregoSafeRegister(&XGAppInfoAddSupportsType, lib, "g_app_info_add_supports_type")
	core.PuregoSafeRegister(&XGAppInfoCanDelete, lib, "g_app_info_can_delete")
	core.PuregoSafeRegister(&XGAppInfoCanRemoveSupportsType, lib, "g_app_info_can_remove_supports_type")
	core.PuregoSafeRegister(&XGAppInfoDelete, lib, "g_app_info_delete")
	core.PuregoSafeRegister(&XGAppInfoDup, lib, "g_app_info_dup")
	core.PuregoSafeRegister(&XGAppInfoEqual, lib, "g_app_info_equal")
	core.PuregoSafeRegister(&XGAppInfoGetCommandline, lib, "g_app_info_get_commandline")
	core.PuregoSafeRegister(&XGAppInfoGetDescription, lib, "g_app_info_get_description")
	core.PuregoSafeRegister(&XGAppInfoGetDisplayName, lib, "g_app_info_get_display_name")
	core.PuregoSafeRegister(&XGAppInfoGetExecutable, lib, "g_app_info_get_executable")
	core.PuregoSafeRegister(&XGAppInfoGetIcon, lib, "g_app_info_get_icon")
	core.PuregoSafeRegister(&XGAppInfoGetId, lib, "g_app_info_get_id")
	core.PuregoSafeRegister(&XGAppInfoGetName, lib, "g_app_info_get_name")
	core.PuregoSafeRegister(&XGAppInfoGetSupportedTypes, lib, "g_app_info_get_supported_types")
	core.PuregoSafeRegister(&XGAppInfoLaunch, lib, "g_app_info_launch")
	core.PuregoSafeRegister(&XGAppInfoLaunchUris, lib, "g_app_info_launch_uris")
	core.PuregoSafeRegister(&XGAppInfoLaunchUrisAsync, lib, "g_app_info_launch_uris_async")
	core.PuregoSafeRegister(&XGAppInfoLaunchUrisFinish, lib, "g_app_info_launch_uris_finish")
	core.PuregoSafeRegister(&XGAppInfoRemoveSupportsType, lib, "g_app_info_remove_supports_type")
	core.PuregoSafeRegister(&XGAppInfoSetAsDefaultForExtension, lib, "g_app_info_set_as_default_for_extension")
	core.PuregoSafeRegister(&XGAppInfoSetAsDefaultForType, lib, "g_app_info_set_as_default_for_type")
	core.PuregoSafeRegister(&XGAppInfoSetAsLastUsedForType, lib, "g_app_info_set_as_last_used_for_type")
	core.PuregoSafeRegister(&XGAppInfoShouldShow, lib, "g_app_info_should_show")
	core.PuregoSafeRegister(&XGAppInfoSupportsFiles, lib, "g_app_info_supports_files")
	core.PuregoSafeRegister(&XGAppInfoSupportsUris, lib, "g_app_info_supports_uris")

}
