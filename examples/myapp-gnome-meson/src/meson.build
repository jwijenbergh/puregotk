config_go = configuration_data()
config_go.set_quoted('PACKAGE_VERSION', meson.project_version())
config_go.set_quoted('GETTEXT_PACKAGE', 'myapp-gnome-meson')
config_go.set_quoted('LOCALEDIR', get_option('prefix') / get_option('localedir'))

config_go_file = custom_target(
    'config-go',
    input: configure_file(
        input: meson.current_source_dir() / 'config.go.in',
        output: 'config_rendered.go',
        configuration: config_go,
    ),
    output: 'config.go',
    command: ['cp', '@INPUT@', meson.current_source_dir() / 'config.go'],
)

myapp_gnome_meson_source_dir = meson.current_source_dir()

myapp_gnome_meson_source_go_files = files('main.go')

myapp_gnome_meson_deps = [
    dependency('gtk4'),
    dependency('libadwaita-1', version: '>= 1.4'),
    dependency('mylib-gtk-meson-0.1', version: '>= 0.1.0'),
]

myapp_gnome_meson_blueprints = custom_target(
    'myapp-gnome-meson-blueprints',
    input: files('window.blp'),
    output: '.',
    command: [
        find_program('blueprint-compiler'),
        'batch-compile',
        '@CURRENT_SOURCE_DIR@',
        '@CURRENT_SOURCE_DIR@',
        '@INPUT@',
    ],
)

myapp_gnome_meson_resources = custom_target(
    'myapp-gnome-meson-resources',
    input: gnome.compile_resources(
        'myapp-gnome-meson',
        'myapp-gnome-meson.gresource.xml',
        gresource_bundle: true,
        dependencies: [myapp_gnome_meson_blueprints],
    ),
    output: 'myapp-gnome-meson-resources.gresource',
    command: [
        'cp',
        '@INPUT@',
        meson.current_source_dir() / 'myapp-gnome-meson.gresource',
    ],
)

custom_target(
    'myapp-gnome-meson',
    build_always_stale: true,  # Go caches internally
    input: myapp_gnome_meson_source_go_files,
    output: 'myapp-gnome-meson',
    command: [
        find_program('go').full_path(),
        'build',
        '-x',
        '-o',
        '@OUTPUT@',
        myapp_gnome_meson_source_dir,
    ],
    depends: [myapp_gnome_meson_resources, config_go_file],
    install: true,
    install_dir: get_option('prefix') / get_option('bindir'),
)
