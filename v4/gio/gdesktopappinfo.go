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

// During invocation, g_desktop_app_info_launch_uris_as_manager() may
// create one or more child processes.  This callback is invoked once
// for each, providing the process ID.
type DesktopAppLaunchCallback func(uintptr, glib.Pid, uintptr)

type DesktopAppInfoClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *DesktopAppInfoClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Interface that is used by backends to associate default
// handlers with URI schemes.
type DesktopAppInfoLookupIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *DesktopAppInfoLookupIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GDesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
type DesktopAppInfoLookup interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetDefaultForUriScheme(UriSchemeVar string) *AppInfoBase
}

var xDesktopAppInfoLookupGLibType func() types.GType

func DesktopAppInfoLookupGLibType() types.GType {
	return xDesktopAppInfoLookupGLibType()
}

type DesktopAppInfoLookupBase struct {
	Ptr uintptr
}

func (x *DesktopAppInfoLookupBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DesktopAppInfoLookupBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the default application for launching applications
// using this URI scheme for a particular #GDesktopAppInfoLookup
// implementation.
//
// The #GDesktopAppInfoLookup interface and this function is used
// to implement g_app_info_get_default_for_uri_scheme() backends
// in a GIO module. There is no reason for applications to use it
// directly. Applications should use g_app_info_get_default_for_uri_scheme().
func (x *DesktopAppInfoLookupBase) GetDefaultForUriScheme(UriSchemeVar string) *AppInfoBase {
	var cls *AppInfoBase

	cret := XGDesktopAppInfoLookupGetDefaultForUriScheme(x.GoPointer(), UriSchemeVar)

	if cret == 0 {
		return nil
	}
	cls = &AppInfoBase{}
	cls.Ptr = cret
	return cls
}

var XGDesktopAppInfoLookupGetDefaultForUriScheme func(uintptr, string) uintptr

const (
	// Extension point for default handler to URI association. See
	// [Extending GIO][extending-gio].
	DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME string = "gio-desktop-app-info-lookup"
)

// #GDesktopAppInfo is an implementation of #GAppInfo based on
// desktop files.
//
// Note that `&lt;gio/gdesktopappinfo.h&gt;` belongs to the UNIX-specific
// GIO interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
type DesktopAppInfo struct {
	gobject.Object
}

var xDesktopAppInfoGLibType func() types.GType

func DesktopAppInfoGLibType() types.GType {
	return xDesktopAppInfoGLibType()
}

func DesktopAppInfoNewFromInternalPtr(ptr uintptr) *DesktopAppInfo {
	cls := &DesktopAppInfo{}
	cls.Ptr = ptr
	return cls
}

var xNewDesktopAppInfo func(string) uintptr

// Creates a new #GDesktopAppInfo based on a desktop file id.
//
// A desktop file id is the basename of the desktop file, including the
// .desktop extension. GIO is looking for a desktop file with this name
// in the `applications` subdirectories of the XDG
// data directories (i.e. the directories specified in the `XDG_DATA_HOME`
// and `XDG_DATA_DIRS` environment variables). GIO also supports the
// prefix-to-subdirectory mapping that is described in the
// [Menu Spec](http://standards.freedesktop.org/menu-spec/latest/)
// (i.e. a desktop id of kde-foo.desktop will match
// `/usr/share/applications/kde/foo.desktop`).
func NewDesktopAppInfo(DesktopIdVar string) *DesktopAppInfo {
	var cls *DesktopAppInfo

	cret := xNewDesktopAppInfo(DesktopIdVar)

	if cret == 0 {
		return nil
	}
	cls = &DesktopAppInfo{}
	cls.Ptr = cret
	return cls
}

var xNewDesktopAppInfoFromFilename func(string) uintptr

// Creates a new #GDesktopAppInfo.
func NewDesktopAppInfoFromFilename(FilenameVar string) *DesktopAppInfo {
	var cls *DesktopAppInfo

	cret := xNewDesktopAppInfoFromFilename(FilenameVar)

	if cret == 0 {
		return nil
	}
	cls = &DesktopAppInfo{}
	cls.Ptr = cret
	return cls
}

var xNewDesktopAppInfoFromKeyfile func(*glib.KeyFile) uintptr

// Creates a new #GDesktopAppInfo.
func NewDesktopAppInfoFromKeyfile(KeyFileVar *glib.KeyFile) *DesktopAppInfo {
	var cls *DesktopAppInfo

	cret := xNewDesktopAppInfoFromKeyfile(KeyFileVar)

	if cret == 0 {
		return nil
	}
	cls = &DesktopAppInfo{}
	cls.Ptr = cret
	return cls
}

var xDesktopAppInfoGetActionName func(uintptr, string) string

// Gets the user-visible display name of the "additional application
// action" specified by @action_name.
//
// This corresponds to the "Name" key within the keyfile group for the
// action.
func (x *DesktopAppInfo) GetActionName(ActionNameVar string) string {

	cret := xDesktopAppInfoGetActionName(x.GoPointer(), ActionNameVar)
	return cret
}

var xDesktopAppInfoGetBoolean func(uintptr, string) bool

// Looks up a boolean value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (x *DesktopAppInfo) GetBoolean(KeyVar string) bool {

	cret := xDesktopAppInfoGetBoolean(x.GoPointer(), KeyVar)
	return cret
}

var xDesktopAppInfoGetCategories func(uintptr) string

// Gets the categories from the desktop file.
func (x *DesktopAppInfo) GetCategories() string {

	cret := xDesktopAppInfoGetCategories(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetFilename func(uintptr) string

// When @info was created from a known filename, return it.  In some
// situations such as the #GDesktopAppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return %NULL.
func (x *DesktopAppInfo) GetFilename() string {

	cret := xDesktopAppInfoGetFilename(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetGenericName func(uintptr) string

// Gets the generic name from the desktop file.
func (x *DesktopAppInfo) GetGenericName() string {

	cret := xDesktopAppInfoGetGenericName(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetIsHidden func(uintptr) bool

// A desktop file is hidden if the Hidden key in it is
// set to True.
func (x *DesktopAppInfo) GetIsHidden() bool {

	cret := xDesktopAppInfoGetIsHidden(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetKeywords func(uintptr) []string

// Gets the keywords from the desktop file.
func (x *DesktopAppInfo) GetKeywords() []string {

	cret := xDesktopAppInfoGetKeywords(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetLocaleString func(uintptr, string) string

// Looks up a localized string value in the keyfile backing @info
// translated to the current locale.
//
// The @key is looked up in the "Desktop Entry" group.
func (x *DesktopAppInfo) GetLocaleString(KeyVar string) string {

	cret := xDesktopAppInfoGetLocaleString(x.GoPointer(), KeyVar)
	return cret
}

var xDesktopAppInfoGetNodisplay func(uintptr) bool

// Gets the value of the NoDisplay key, which helps determine if the
// application info should be shown in menus. See
// %G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
func (x *DesktopAppInfo) GetNodisplay() bool {

	cret := xDesktopAppInfoGetNodisplay(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetShowIn func(uintptr, string) bool

// Checks if the application info should be shown in menus that list available
// applications for a specific name of the desktop, based on the
// `OnlyShowIn` and `NotShowIn` keys.
//
// @desktop_env should typically be given as %NULL, in which case the
// `XDG_CURRENT_DESKTOP` environment variable is consulted.  If you want
// to override the default mechanism then you may specify @desktop_env,
// but this is not recommended.
//
// Note that g_app_info_should_show() for @info will include this check (with
// %NULL for @desktop_env) as well as additional checks.
func (x *DesktopAppInfo) GetShowIn(DesktopEnvVar string) bool {

	cret := xDesktopAppInfoGetShowIn(x.GoPointer(), DesktopEnvVar)
	return cret
}

var xDesktopAppInfoGetStartupWmClass func(uintptr) string

// Retrieves the StartupWMClass field from @info. This represents the
// WM_CLASS property of the main window of the application, if launched
// through @info.
func (x *DesktopAppInfo) GetStartupWmClass() string {

	cret := xDesktopAppInfoGetStartupWmClass(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetString func(uintptr, string) string

// Looks up a string value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (x *DesktopAppInfo) GetString(KeyVar string) string {

	cret := xDesktopAppInfoGetString(x.GoPointer(), KeyVar)
	return cret
}

var xDesktopAppInfoGetStringList func(uintptr, string, uint) []string

// Looks up a string list value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (x *DesktopAppInfo) GetStringList(KeyVar string, LengthVar uint) []string {

	cret := xDesktopAppInfoGetStringList(x.GoPointer(), KeyVar, LengthVar)
	return cret
}

var xDesktopAppInfoHasKey func(uintptr, string) bool

// Returns whether @key exists in the "Desktop Entry" group
// of the keyfile backing @info.
func (x *DesktopAppInfo) HasKey(KeyVar string) bool {

	cret := xDesktopAppInfoHasKey(x.GoPointer(), KeyVar)
	return cret
}

var xDesktopAppInfoLaunchAction func(uintptr, string, uintptr)

// Activates the named application action.
//
// You may only call this function on action names that were
// returned from g_desktop_app_info_list_actions().
//
// Note that if the main entry of the desktop file indicates that the
// application supports startup notification, and @launch_context is
// non-%NULL, then startup notification will be used when activating the
// action (and as such, invocation of the action on the receiving side
// must signal the end of startup notification when it is completed).
// This is the expected behaviour of applications declaring additional
// actions, as per the desktop file specification.
//
// As with g_app_info_launch() there is no way to detect failures that
// occur while using this function.
func (x *DesktopAppInfo) LaunchAction(ActionNameVar string, LaunchContextVar *AppLaunchContext) {

	xDesktopAppInfoLaunchAction(x.GoPointer(), ActionNameVar, LaunchContextVar.GoPointer())

}

var xDesktopAppInfoLaunchUrisAsManager func(uintptr, *glib.List, uintptr, glib.SpawnFlags, uintptr, uintptr, uintptr, uintptr, **glib.Error) bool

// This function performs the equivalent of g_app_info_launch_uris(),
// but is intended primarily for operating system components that
// launch applications.  Ordinary applications should use
// g_app_info_launch_uris().
//
// If the application is launched via GSpawn, then @spawn_flags, @user_setup
// and @user_setup_data are used for the call to g_spawn_async().
// Additionally, @pid_callback (with @pid_callback_data) will be called to
// inform about the PID of the created process. See g_spawn_async_with_pipes()
// for information on certain parameter conditions that can enable an
// optimized posix_spawn() codepath to be used.
//
// If application launching occurs via some other mechanism (eg: D-Bus
// activation) then @spawn_flags, @user_setup, @user_setup_data,
// @pid_callback and @pid_callback_data are ignored.
func (x *DesktopAppInfo) LaunchUrisAsManager(UrisVar *glib.List, LaunchContextVar *AppLaunchContext, SpawnFlagsVar glib.SpawnFlags, UserSetupVar *glib.SpawnChildSetupFunc, UserSetupDataVar uintptr, PidCallbackVar *DesktopAppLaunchCallback, PidCallbackDataVar uintptr) (bool, error) {
	var cerr *glib.Error

	cret := xDesktopAppInfoLaunchUrisAsManager(x.GoPointer(), UrisVar, LaunchContextVar.GoPointer(), SpawnFlagsVar, glib.NewCallback(UserSetupVar), UserSetupDataVar, glib.NewCallback(PidCallbackVar), PidCallbackDataVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDesktopAppInfoLaunchUrisAsManagerWithFds func(uintptr, *glib.List, uintptr, glib.SpawnFlags, uintptr, uintptr, uintptr, uintptr, int, int, int, **glib.Error) bool

// Equivalent to g_desktop_app_info_launch_uris_as_manager() but allows
// you to pass in file descriptors for the stdin, stdout and stderr streams
// of the launched process.
//
// If application launching occurs via some non-spawn mechanism (e.g. D-Bus
// activation) then @stdin_fd, @stdout_fd and @stderr_fd are ignored.
func (x *DesktopAppInfo) LaunchUrisAsManagerWithFds(UrisVar *glib.List, LaunchContextVar *AppLaunchContext, SpawnFlagsVar glib.SpawnFlags, UserSetupVar *glib.SpawnChildSetupFunc, UserSetupDataVar uintptr, PidCallbackVar *DesktopAppLaunchCallback, PidCallbackDataVar uintptr, StdinFdVar int, StdoutFdVar int, StderrFdVar int) (bool, error) {
	var cerr *glib.Error

	cret := xDesktopAppInfoLaunchUrisAsManagerWithFds(x.GoPointer(), UrisVar, LaunchContextVar.GoPointer(), SpawnFlagsVar, glib.NewCallback(UserSetupVar), UserSetupDataVar, glib.NewCallback(PidCallbackVar), PidCallbackDataVar, StdinFdVar, StdoutFdVar, StderrFdVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDesktopAppInfoListActions func(uintptr) []string

// Returns the list of "additional application actions" supported on the
// desktop file, as per the desktop file specification.
//
// As per the specification, this is the list of actions that are
// explicitly listed in the "Actions" key of the [Desktop Entry] group.
func (x *DesktopAppInfo) ListActions() []string {

	cret := xDesktopAppInfoListActions(x.GoPointer())
	return cret
}

func (c *DesktopAppInfo) GoPointer() uintptr {
	return c.Ptr
}

func (c *DesktopAppInfo) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Adds a content type to the application information to indicate the
// application is capable of opening files with the given content type.
func (x *DesktopAppInfo) AddSupportsType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoAddSupportsType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Obtains the information whether the #GAppInfo can be deleted.
// See g_app_info_delete().
func (x *DesktopAppInfo) CanDelete() bool {

	cret := XGAppInfoCanDelete(x.GoPointer())
	return cret
}

// Checks if a supported content type can be removed from an application.
func (x *DesktopAppInfo) CanRemoveSupportsType() bool {

	cret := XGAppInfoCanRemoveSupportsType(x.GoPointer())
	return cret
}

// Tries to delete a #GAppInfo.
//
// On some platforms, there may be a difference between user-defined
// #GAppInfos which can be deleted, and system-wide ones which cannot.
// See g_app_info_can_delete().
func (x *DesktopAppInfo) Delete() bool {

	cret := XGAppInfoDelete(x.GoPointer())
	return cret
}

// Creates a duplicate of a #GAppInfo.
func (x *DesktopAppInfo) Dup() *AppInfoBase {
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
func (x *DesktopAppInfo) Equal(Appinfo2Var AppInfo) bool {

	cret := XGAppInfoEqual(x.GoPointer(), Appinfo2Var.GoPointer())
	return cret
}

// Gets the commandline with which the application will be
// started.
func (x *DesktopAppInfo) GetCommandline() string {

	cret := XGAppInfoGetCommandline(x.GoPointer())
	return cret
}

// Gets a human-readable description of an installed application.
func (x *DesktopAppInfo) GetDescription() string {

	cret := XGAppInfoGetDescription(x.GoPointer())
	return cret
}

// Gets the display name of the application. The display name is often more
// descriptive to the user than the name itself.
func (x *DesktopAppInfo) GetDisplayName() string {

	cret := XGAppInfoGetDisplayName(x.GoPointer())
	return cret
}

// Gets the executable's name for the installed application.
func (x *DesktopAppInfo) GetExecutable() string {

	cret := XGAppInfoGetExecutable(x.GoPointer())
	return cret
}

// Gets the icon for the application.
func (x *DesktopAppInfo) GetIcon() *IconBase {
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
func (x *DesktopAppInfo) GetId() string {

	cret := XGAppInfoGetId(x.GoPointer())
	return cret
}

// Gets the installed name of the application.
func (x *DesktopAppInfo) GetName() string {

	cret := XGAppInfoGetName(x.GoPointer())
	return cret
}

// Retrieves the list of content types that @app_info claims to support.
// If this information is not provided by the environment, this function
// will return %NULL.
// This function does not take in consideration associations added with
// g_app_info_add_supports_type(), but only those exported directly by
// the application.
func (x *DesktopAppInfo) GetSupportedTypes() []string {

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
func (x *DesktopAppInfo) Launch(FilesVar *glib.List, ContextVar *AppLaunchContext) (bool, error) {
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
func (x *DesktopAppInfo) LaunchUris(UrisVar *glib.List, ContextVar *AppLaunchContext) (bool, error) {
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
func (x *DesktopAppInfo) LaunchUrisAsync(UrisVar *glib.List, ContextVar *AppLaunchContext, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGAppInfoLaunchUrisAsync(x.GoPointer(), UrisVar, ContextVar.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes a g_app_info_launch_uris_async() operation.
func (x *DesktopAppInfo) LaunchUrisFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoLaunchUrisFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Removes a supported type from an application, if possible.
func (x *DesktopAppInfo) RemoveSupportsType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoRemoveSupportsType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the application as the default handler for the given file extension.
func (x *DesktopAppInfo) SetAsDefaultForExtension(ExtensionVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoSetAsDefaultForExtension(x.GoPointer(), ExtensionVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the application as the default handler for a given type.
func (x *DesktopAppInfo) SetAsDefaultForType(ContentTypeVar string) (bool, error) {
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
func (x *DesktopAppInfo) SetAsLastUsedForType(ContentTypeVar string) (bool, error) {
	var cerr *glib.Error

	cret := XGAppInfoSetAsLastUsedForType(x.GoPointer(), ContentTypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Checks if the application info should be shown in menus that
// list available applications.
func (x *DesktopAppInfo) ShouldShow() bool {

	cret := XGAppInfoShouldShow(x.GoPointer())
	return cret
}

// Checks if the application accepts files as arguments.
func (x *DesktopAppInfo) SupportsFiles() bool {

	cret := XGAppInfoSupportsFiles(x.GoPointer())
	return cret
}

// Checks if the application supports reading files and directories from URIs.
func (x *DesktopAppInfo) SupportsUris() bool {

	cret := XGAppInfoSupportsUris(x.GoPointer())
	return cret
}

var xDesktopAppInfoGetImplementations func(string) *glib.List

// Gets all applications that implement @interface.
//
// An application implements an interface if that interface is listed in
// the Implements= line of the desktop file of the application.
func DesktopAppInfoGetImplementations(InterfaceVar string) *glib.List {

	cret := xDesktopAppInfoGetImplementations(InterfaceVar)
	return cret
}

var xDesktopAppInfoSearch func(string) uintptr

// Searches desktop files for ones that match @search_string.
//
// The return value is an array of strvs.  Each strv contains a list of
// applications that matched @search_string with an equal score.  The
// outer list is sorted by score so that the first strv contains the
// best-matching applications, and so on.
// The algorithm for determining matches is undefined and may change at
// any time.
//
// None of the search results are subjected to the normal validation
// checks performed by g_desktop_app_info_new() (for example, checking that
// the executable referenced by a result exists), and so it is possible for
// g_desktop_app_info_new() to return %NULL when passed an app ID returned by
// this function. It is expected that calling code will do this when
// subsequently creating a #GDesktopAppInfo for each result.
func DesktopAppInfoSearch(SearchStringVar string) uintptr {

	cret := xDesktopAppInfoSearch(SearchStringVar)
	return cret
}

var xDesktopAppInfoSetDesktopEnv func(string)

// Sets the name of the desktop that the application is running in.
// This is used by g_app_info_should_show() and
// g_desktop_app_info_get_show_in() to evaluate the
// `OnlyShowIn` and `NotShowIn`
// desktop entry fields.
//
// Should be called only once; subsequent calls are ignored.
func DesktopAppInfoSetDesktopEnv(DesktopEnvVar string) {

	xDesktopAppInfoSetDesktopEnv(DesktopEnvVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDesktopAppInfoGLibType, lib, "g_desktop_app_info_get_type")

	core.PuregoSafeRegister(&xNewDesktopAppInfo, lib, "g_desktop_app_info_new")
	core.PuregoSafeRegister(&xNewDesktopAppInfoFromFilename, lib, "g_desktop_app_info_new_from_filename")
	core.PuregoSafeRegister(&xNewDesktopAppInfoFromKeyfile, lib, "g_desktop_app_info_new_from_keyfile")

	core.PuregoSafeRegister(&xDesktopAppInfoGetActionName, lib, "g_desktop_app_info_get_action_name")
	core.PuregoSafeRegister(&xDesktopAppInfoGetBoolean, lib, "g_desktop_app_info_get_boolean")
	core.PuregoSafeRegister(&xDesktopAppInfoGetCategories, lib, "g_desktop_app_info_get_categories")
	core.PuregoSafeRegister(&xDesktopAppInfoGetFilename, lib, "g_desktop_app_info_get_filename")
	core.PuregoSafeRegister(&xDesktopAppInfoGetGenericName, lib, "g_desktop_app_info_get_generic_name")
	core.PuregoSafeRegister(&xDesktopAppInfoGetIsHidden, lib, "g_desktop_app_info_get_is_hidden")
	core.PuregoSafeRegister(&xDesktopAppInfoGetKeywords, lib, "g_desktop_app_info_get_keywords")
	core.PuregoSafeRegister(&xDesktopAppInfoGetLocaleString, lib, "g_desktop_app_info_get_locale_string")
	core.PuregoSafeRegister(&xDesktopAppInfoGetNodisplay, lib, "g_desktop_app_info_get_nodisplay")
	core.PuregoSafeRegister(&xDesktopAppInfoGetShowIn, lib, "g_desktop_app_info_get_show_in")
	core.PuregoSafeRegister(&xDesktopAppInfoGetStartupWmClass, lib, "g_desktop_app_info_get_startup_wm_class")
	core.PuregoSafeRegister(&xDesktopAppInfoGetString, lib, "g_desktop_app_info_get_string")
	core.PuregoSafeRegister(&xDesktopAppInfoGetStringList, lib, "g_desktop_app_info_get_string_list")
	core.PuregoSafeRegister(&xDesktopAppInfoHasKey, lib, "g_desktop_app_info_has_key")
	core.PuregoSafeRegister(&xDesktopAppInfoLaunchAction, lib, "g_desktop_app_info_launch_action")
	core.PuregoSafeRegister(&xDesktopAppInfoLaunchUrisAsManager, lib, "g_desktop_app_info_launch_uris_as_manager")
	core.PuregoSafeRegister(&xDesktopAppInfoLaunchUrisAsManagerWithFds, lib, "g_desktop_app_info_launch_uris_as_manager_with_fds")
	core.PuregoSafeRegister(&xDesktopAppInfoListActions, lib, "g_desktop_app_info_list_actions")

	core.PuregoSafeRegister(&xDesktopAppInfoGetImplementations, lib, "g_desktop_app_info_get_implementations")
	core.PuregoSafeRegister(&xDesktopAppInfoSearch, lib, "g_desktop_app_info_search")
	core.PuregoSafeRegister(&xDesktopAppInfoSetDesktopEnv, lib, "g_desktop_app_info_set_desktop_env")

	core.PuregoSafeRegister(&xDesktopAppInfoLookupGLibType, lib, "g_desktop_app_info_lookup_get_type")

	core.PuregoSafeRegister(&XGDesktopAppInfoLookupGetDefaultForUriScheme, lib, "g_desktop_app_info_lookup_get_default_for_uri_scheme")

}
