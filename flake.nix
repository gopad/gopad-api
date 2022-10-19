{
  description = "Nix configuration for Gopad API";

  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixpkgs-unstable";
    };

    utils = {
      url = "github:numtide/flake-utils";
    };
  };

  outputs = { self, nixpkgs, utils, ... }@inputs:
    {

    } // utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            buf
            gnumake
            grpcurl
            protobuf
            protoc-gen-connect-go
            protoc-gen-go
          ];

          shellHook = ''
            export GOPAD_API_SERVER_PPROF=true
          '';
        };
      }
    );
}
