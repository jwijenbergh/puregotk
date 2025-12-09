#pragma once

#include <adwaita.h>

G_BEGIN_DECLS

#define MYLIB_GTK_MESON_INSIDE

/**
 * MyLibGtkMesonMainApplicationWindow:
 *
 * Example application window with a test button and toast notifications.
 */
#define MYLIB_GTK_MESON_MAIN_APPLICATION_WINDOW                                    \
  (mylib_gtk_meson_main_application_window_get_type())

typedef struct _MyLibGtkMesonMainApplicationWindow MyLibGtkMesonMainApplicationWindow;

/**
 * mylib_gtk_meson_main_application_window_get_type:
 *
 * Gets the GType for MyLibGtkMesonMainApplicationWindow.
 *
 * Returns: the GType for MyLibGtkMesonMainApplicationWindow
 */
GType mylib_gtk_meson_main_application_window_get_type(void);

/**
 * mylib_gtk_meson_main_application_window_show_toast:
 * @window: a MyLibGtkMesonMainApplicationWindow
 * @message: the message to display in the toast
 *
 * Shows a toast notification with the given message inside the window.
 */
void mylib_gtk_meson_main_application_window_show_toast(
    MyLibGtkMesonMainApplicationWindow *window, const char *message);

#undef MYLIB_GTK_MESON_INSIDE

G_END_DECLS
