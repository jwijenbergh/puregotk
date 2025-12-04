#pragma once

#include <adwaita.h>

G_BEGIN_DECLS

#define MYLIB_MESON_INSIDE

/**
 * MyLibMesonMainApplicationWindow:
 *
 * Example application window with a test button and toast notifications.
 */
#define MYLIB_MESON_MAIN_APPLICATION_WINDOW                                    \
  (mylib_meson_main_application_window_get_type())

typedef struct _MyLibMesonMainApplicationWindow MyLibMesonMainApplicationWindow;

/**
 * mylib_meson_main_application_window_get_type:
 *
 * Gets the GType for MyLibMesonMainApplicationWindow.
 *
 * Returns: the GType for MyLibMesonMainApplicationWindow
 */
GType mylib_meson_main_application_window_get_type(void);

/**
 * mylib_meson_main_application_window_show_toast:
 * @window: a MyLibMesonMainApplicationWindow
 * @message: the message to display in the toast
 *
 * Shows a toast notification with the given message inside the window.
 */
void mylib_meson_main_application_window_show_toast(
    MyLibMesonMainApplicationWindow *window, const char *message);

#undef MYLIB_MESON_INSIDE

G_END_DECLS
