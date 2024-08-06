{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    devenv.url = "github:cachix/devenv";
    dagger.url = "github:dagger/nix";
    dagger.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        inputs.devenv.flakeModule
      ];

      systems = [ "x86_64-linux" "x86_64-darwin" "aarch64-darwin" ];

      perSystem = { config, self', inputs', pkgs, system, ... }: rec {
        _module.args.pkgs = import inputs.nixpkgs {
          inherit system;

          overlays = [
            (final: prev: {
              dagger = inputs'.dagger.packages.dagger;
              licensei = self'.packages.licensei;
            })
          ];
        };

        devenv.shells = {
          default = {
            languages = {
              go = {
                enable = true;
                package = pkgs.go_1_22;
              };

              python = {
                enable = true;
                package = pkgs.python39;
              };

              javascript = {
                enable = true;
                package = pkgs.nodejs_20;
              };
            };

            pre-commit.hooks = {
              nixpkgs-fmt.enable = true;
              commitizen.enable = true;

              commitizen-branch = {
                enable = true;
                name = "commitizen-branch check";
                description = ''
                  Check whether commit messages on the current HEAD follows committing rules.
                '';
                entry = "${pkgs.commitizen}/bin/cz check --allow-abort --rev-range origin/HEAD..HEAD";
                pass_filenames = false;
                stages = [ "manual" ];
              };
            };

            packages = with pkgs; [
              gnumake
              mage

              # Kafka build dependencies
              rdkafka # https://github.com/confluentinc/confluent-kafka-go#librdkafka
              cyrus_sasl
              pkg-config
              # confluent-platform

              # golangci-lint
              goreleaser
              air

              curl
              jq
              minikube
              kind
              kubectl
              helm-docs

              benthos

              # node
              nodePackages.pnpm

              # python
              poetry

              # FIXME: Building as a go module is not supported by maintainers:
              # - GH issue for atlas: https://github.com/ariga/atlas/issues/2582
              # - Module src we use: https://github.com/NixOS/nixpkgs/blob/6d54ef5bc208956deca392af75ad9568e4a58429/pkgs/development/tools/database/atlas/default.nix#L8
              # - building from src doesn't actually work: https://github.com/ariga/atlas/issues/2388
              # atlas
              self'.packages.atlas

              just
              semver-tool

              dagger
              licensei
            ];

            env = {
              KUBECONFIG = "${config.devenv.shells.default.env.DEVENV_STATE}/kube/config";
              KIND_CLUSTER_NAME = "openmeter";

              HELM_CACHE_HOME = "${config.devenv.shells.default.env.DEVENV_STATE}/helm/cache";
              HELM_CONFIG_HOME = "${config.devenv.shells.default.env.DEVENV_STATE}/helm/config";
              HELM_DATA_HOME = "${config.devenv.shells.default.env.DEVENV_STATE}/helm/data";

            };

            # https://github.com/cachix/devenv/issues/528#issuecomment-1556108767
            containers = pkgs.lib.mkForce { };
          };

          ci = devenv.shells.default;
        };

        packages = {
          licensei = pkgs.buildGoModule rec {
            pname = "licensei";
            version = "0.8.0";

            src = pkgs.fetchFromGitHub {
              owner = "goph";
              repo = "licensei";
              rev = "v${version}";
              sha256 = "sha256-Pvjmvfk0zkY2uSyLwAtzWNn5hqKImztkf8S6OhX8XoM=";
            };

            vendorHash = "sha256-ZIpZ2tPLHwfWiBywN00lPI1R7u7lseENIiybL3+9xG8=";

            subPackages = [ "cmd/licensei" ];

            ldflags = [
              "-w"
              "-s"
              "-X main.version=v${version}"
            ];
          };

          atlas = pkgs.stdenv.mkDerivation rec {
            pname = "atlas";
            version = "0.25.0";
            src = pkgs.fetchurl {
              url = "https://release.ariga.io/atlas/atlas-community-darwin-arm64-v${version}";
              hash = "sha256-jkKK1PAhS/jZiKux6ht7bciHuVhX+8CBfvw1Da9aY6k=";
            };

            unpackPhase = ''
              cp $src atlas
            '';

            installPhase = ''
              mkdir -p $out/bin
              cp atlas $out/bin/atlas
            '';

          };
        };
      };
    };
}
