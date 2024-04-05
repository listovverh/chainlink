{ pkgs, scriptDir }:
with pkgs;
let
  go = go_1_21;
  postgresql = postgresql_14;
  nodejs = nodejs-18_x;
  nodePackages = pkgs.nodePackages.override { inherit nodejs; };
in
mkShell {
  nativeBuildInputs = [
    go
    goreleaser
    postgresql

    python3
    python3Packages.pip

    curl
    nodejs
    nodePackages.pnpm
    # TODO: compiler / gcc for secp compilation
    go-ethereum # geth
    # parity # openethereum
    go-mockery

    # tooling
    gotools
    gopls
    delve
    golangci-lint
    github-cli
    jq

    # deployment
    awscli2
    devspace
    kubectl
    kubernetes-helm

    # gofuzz
  ] ++ lib.optionals stdenv.isLinux [
    # some dependencies needed for node-gyp on pnpm install
    pkg-config
    libudev-zero
    libusb1
  ];
  LD_LIBRARY_PATH = "${stdenv.cc.cc.lib}/lib64:$LD_LIBRARY_PATH";
  GOROOT = "${go}/share/go";
  CGO_ENABLED = "0";
  HELM_REPOSITORY_CONFIG = "${scriptDir}/integration-tests/.helm-repositories.yaml";

  PGDATA = "db";
  CL_DATABASE_URL = "postgresql://chainlink:chainlink@localhost:5432/chainlink_test?sslmode=disable";

  shellHook = ''
    # Update helm repositories
    helm repo update > /dev/null
    # setup go bin for nix
    export GOBIN=$HOME/.nix-go/bin
    mkdir -p $GOBIN
    export PATH=$GOBIN:$PATH
    # install gotestloghelper
    go install github.com/smartcontractkit/chainlink-testing-framework/tools/gotestloghelper@latest
  '';
}
