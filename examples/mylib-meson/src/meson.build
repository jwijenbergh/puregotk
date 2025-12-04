api_version = '0.1'

pkg = import('pkgconfig')
gnome = import('gnome')
fs = import('fs')

config_go = configuration_data()
config_go.set_quoted('PACKAGE_VERSION', meson.project_version())
config_go.set_quoted('GETTEXT_PACKAGE', 'mylib-meson')
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

mylib_meson_source_dir = meson.current_source_dir()

mylib_meson_source_go_files = files('mylib-meson.go')

mylib_meson_headers = ['mylib-meson.h', 'window.h']

version_split = meson.project_version().split('.')
version_conf = configuration_data()
version_conf.set('VERSION', meson.project_version())
version_conf.set('MAJOR_VERSION', version_split[0])
version_conf.set('MINOR_VERSION', version_split[1])
version_conf.set('MICRO_VERSION', version_split[2])

configure_file(
    input: 'mylib-meson-version.h.in',
    output: 'mylib-meson-version.h',
    configuration: version_conf,
    install: true,
    install_dir: get_option('includedir') / 'mylib-meson',
)

mylib_meson_deps = [
    dependency('gtk4'),
    dependency('libadwaita-1', version: '>= 1.4'),
]

mylib_meson_blueprints = custom_target(
    'mylib-meson-blueprints',
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

mylib_meson_resources = custom_target(
    'mylib-meson-resources',
    input: gnome.compile_resources(
        'mylib-meson',
        'mylib-meson.gresource.xml',
        gresource_bundle: true,
        dependencies: [mylib_meson_blueprints],
    ),
    output: 'mylib-meson-resources.gresource',
    command: [
        'cp',
        '@INPUT@',
        meson.current_source_dir() / 'mylib-meson.gresource',
    ],
)

mylib_meson_lib_overwrite = custom_target(
    'mylib-meson-lib-overwrite',
    build_always_stale: true,  # Go caches internally
    input: mylib_meson_source_go_files,
    output: 'libmylib-meson-overwrite-' + api_version + '.so',
    command: [
        find_program('go').full_path(),
        'build',
        '-buildmode',
        'c-shared',
        '-x',
        '-o',
        '@OUTPUT@',
        mylib_meson_source_dir,
    ],
    depends: [mylib_meson_resources, config_go_file],
)

mylib_meson_archive = custom_target(
    'mylib-meson-archive',
    build_always_stale: true,  # Go caches internally
    input: mylib_meson_source_go_files,
    output: 'libmylib-meson-' + api_version + '.a',
    command: [
        find_program('go').full_path(),
        'build',
        '-buildmode',
        'c-archive',
        '-x',
        '-o',
        '@OUTPUT@',
        mylib_meson_source_dir,
    ],
    depends: [mylib_meson_resources, config_go_file, mylib_meson_lib_overwrite],
)

mylib_meson_lib = shared_library(
    'mylib-meson-' + api_version,
    [],
    link_whole: [mylib_meson_archive],
    dependencies: mylib_meson_deps,
    install: true,
)

# We overwrite the manually linked mylib_meson_lib with mylib_meson_lib_overwrite here since the manually linked library built with `c-archive` can not be used by another Go program without a `morestack` crash
meson.add_install_script(
    'install',
    mylib_meson_lib_overwrite,
    get_option('prefix') / get_option('libdir') / fs.name(
        mylib_meson_lib.full_path(),
    ),
)

install_headers(mylib_meson_headers, subdir: 'mylib-meson')

pkg.generate(
    description: 'Example GTK library built with Go',
    libraries: mylib_meson_lib,
    name: 'mylib-meson',
    filebase: 'mylib-meson-' + api_version,
    version: meson.project_version(),
    subdirs: 'mylib-meson',
    requires: ['gtk4 >= 4.17.5', 'libadwaita-1 >= 1.7.5'],
    install_dir: get_option('libdir') / 'pkgconfig',
)

mylib_meson_gir = gnome.generate_gir(
    mylib_meson_lib,
    sources: [mylib_meson_headers],
    header: 'mylib-meson.h',
    export_packages: 'mylib-meson-' + api_version,
    nsversion: api_version,
    namespace: 'MyLibMeson',
    symbol_prefix: 'mylib_meson',
    identifier_prefix: 'MyLibMeson',
    includes: ['Gtk-4.0', 'Adw-1'],
    install: true,
)

custom_target(
    'mylib-meson-docs',
    input: [
        configure_file(
            input: meson.current_source_dir() / 'docs.toml.in',
            output: 'docs.toml',
            configuration: config_go,
        ),
        mylib_meson_gir,
    ],
    output: 'mylib-meson-' + api_version,
    command: [
        find_program('gi-docgen'),
        'generate',
        '--no-namespace-dir',
        '--config=@INPUT0@',
        '--output-dir=@OUTPUT@',
        '@INPUT1@',
    ],
    install: true,
    install_dir: get_option('datadir') / 'doc',
)
