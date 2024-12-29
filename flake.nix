{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    zimeg.url = "github:zimeg/nur-packages";
  };
  outputs = { nixpkgs, flake-utils, ... } @ inputs:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        gon =
          if pkgs.stdenv.isDarwin then
            inputs.zimeg.packages.${pkgs.system}.gon
          else
            null;
      in
      {
        devShells.tom = pkgs.mkShell {
          packages = with pkgs; [
            time
          ];
        };
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            gnumake
            go
            gocyclo
            gofumpt
            golangci-lint
            gopls
            goreleaser
          ];
          shellHook = ''
            go mod tidy
          '';
        };
        devShells.gh = pkgs.mkShell {
          packages = with pkgs; [
            github-runner
          ];
        };
        devShells.gon =
          if pkgs.stdenv.isDarwin then
            pkgs.mkShell
              {
                packages = with pkgs; [
                  go
                  gon
                  goreleaser
                ];
                shellHook = ''
                  export PATH=/usr/bin:$PATH # https://github.com/zimeg/nur-packages/issues/4
                '';
              }
          else
            null;
        packages.default = pkgs.buildGoModule {
          pname = "emporia-time";
          version = "unversioned";
          src = ./.;
          ldflags = [
            "-s"
            "-w"
            "-X main.version=dev"
          ];
          doCheck = false;
          vendorHash = "sha256-G5sLF3awEQGkSqaXmhSw+IcBq+NoG3QoR+L8XymWfDU=";
          buildInputs = [
            pkgs.time
          ];
          meta = {
            description = "an energy aware time command";
            homepage = "https://github.com/zimeg/emporia-time";
            license = pkgs.lib.licenses.mit;
          };
        };
      });
}
