#!/bin/sh

set -e

for f in internal/gir/spec/*.gir; do flatpak run --filesystem="${PWD}" --command=sh org.gnome.Sdk -c "cp /usr/share/gir-1.0/$(basename ${f}) ${PWD}/${f}"; done
