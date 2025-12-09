project(
    'myapp-gnome-meson',
    'c',
    version: '0.1.0',
    meson_version: '>= 1.0.0',
    default_options: [],
)
add_project_arguments(['-I' + meson.project_build_root()], language: 'c')

i18n = import('i18n')
gnome = import('gnome')

subdir('src')
subdir('po')
