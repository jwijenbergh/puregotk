{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    buildInputs = [ 
      pkgs.go
      pkgs.godef
      pkgs.go-tools
      pkgs.vim
      pkgs.gofumpt
      pkgs.gopls
      pkgs.python3
      pkgs.libadwaita.dev
      pkgs.python3.pkgs.pip
      pkgs.gotools
    ];
  shellHook = ''
    # Tells pip to put packages into $PIP_PREFIX instead of the usual locations.
    # See https://pip.pypa.io/en/stable/user_guide/#environment-variables.
    export PIP_PREFIX=$(pwd)/_build/pip_packages
    export PYTHONPATH="$PIP_PREFIX/${pkgs.python3.sitePackages}:$PYTHONPATH"
    export PATH="$PIP_PREFIX/bin:$PATH"
    export GOPATH="$HOME/go"
    unset SOURCE_DATE_EPOCH
  '';
}

