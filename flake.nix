{
  description = "Nix flake for development";

  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixpkgs-unstable";
    };

    devenv = {
      url = "github:cachix/devenv";
    };

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
    };

    git-hooks = {
      url = "github:cachix/git-hooks.nix";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        inputs.devenv.flakeModule
      ];

      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          system,
          ...
        }:
        {
          imports = [
            {
              _module.args.pkgs = import inputs.nixpkgs {
                inherit system;
                config.allowUnfree = true;
              };
            }
          ];

          devenv = {
            shells = {
              default = {
                name = "gopad-api";

                git-hooks = {
                  hooks = {
                    nixfmt-rfc-style = {
                      enable = true;
                    };

                    gofmt = {
                      enable = true;
                    };

                    golangci-lint = {
                      enable = true;
                      entry = "go tool github.com/golangci/golangci-lint/cmd/golangci-lint run ./...";
                      pass_filenames = false;
                    };
                  };
                };

                languages = {
                  go = {
                    enable = true;
                    package = pkgs.go_1_24;
                  };
                };

                packages = with pkgs; [
                  cosign
                  gnumake
                  goreleaser
                  httpie
                  nixfmt-rfc-style
                  posting
                  sqlite
                  yq-go
                ];

                env = {
                  CGO_ENABLED = "0";

                  GOPAD_API_LOG_LEVEL = "debug";
                  GOPAD_API_LOG_PRETTY = "true";
                  GOPAD_API_LOG_COLOR = "true";

                  GOPAD_API_TOKEN_SECRET = "TxHrYxMAg01rBeEWrHn1BjOP";
                  GOPAD_API_TOKEN_EXPIRE = "1h";

                  # GOPAD_API_SERVER_CERT = ".devenv/state/mkcert/localhost+1.pem";
                  # GOPAD_API_SERVER_KEY = ".devenv/state/mkcert/localhost+1-key.pem";

                  GOPAD_API_DATABASE_DRIVER = "sqlite3";
                  GOPAD_API_DATABASE_NAME = "storage/gopad.sqlite3";

                  # GOPAD_API_DATABASE_DRIVER = "mysql";
                  # GOPAD_API_DATABASE_ADDRESS = "127.0.0.1";
                  # GOPAD_API_DATABASE_PORT = "3306";
                  # GOPAD_API_DATABASE_USERNAME = "gopad";
                  # GOPAD_API_DATABASE_PASSWORD = "p455w0rd";
                  # GOPAD_API_DATABASE_NAME = "gopad";

                  # GOPAD_API_DATABASE_DRIVER = "postgres";
                  # GOPAD_API_DATABASE_ADDRESS = "127.0.0.1";
                  # GOPAD_API_DATABASE_PORT = "5432";
                  # GOPAD_API_DATABASE_USERNAME = "gopad";
                  # GOPAD_API_DATABASE_PASSWORD = "p455w0rd";
                  # GOPAD_API_DATABASE_NAME = "gopad";

                  GOPAD_API_UPLOAD_DRIVER = "file";
                  GOPAD_API_UPLOAD_PATH = "storage/uploads/";

                  # GOPAD_API_UPLOAD_DRIVER = "s3";
                  # GOPAD_API_UPLOAD_ENDPOINT = "127.0.0.1:9000";
                  # GOPAD_API_UPLOAD_BUCKET = "gopad";
                  # GOPAD_API_UPLOAD_REGION = "us-east-1";
                  # GOPAD_API_UPLOAD_ACCESS = "9VLV3OI55N1077Y9IALV";
                  # GOPAD_API_UPLOAD_SECRET = "bwcRkw5w6uF6BWBqOtsnMbwZSIDKQopy9DSo90ab";

                  GOPAD_API_ADMIN_USERNAME = "admin";
                  GOPAD_API_ADMIN_PASSWORD = "p455w0rd";
                  GOPAD_API_ADMIN_EMAIL = "gopad@webhippie.de";
                };

                services = {
                  minio = {
                    enable = true;
                    accessKey = "9VLV3OI55N1077Y9IALV";
                    secretKey = "bwcRkw5w6uF6BWBqOtsnMbwZSIDKQopy9DSo90ab";
                    buckets = [
                      "gopad"
                    ];
                  };
                  postgres = {
                    enable = true;
                    listen_addresses = "127.0.0.1";
                    initialScript = ''
                      CREATE USER gopad WITH ENCRYPTED PASSWORD 'p455w0rd';
                      GRANT ALL PRIVILEGES ON DATABASE gopad TO gopad;
                    '';
                    initialDatabases = [
                      {
                        name = "gopad";
                      }
                    ];
                  };
                };

                processes = {
                  gopad-server = {
                    exec = "make watch";

                    process-compose = {
                      readiness_probe = {
                        exec.command = "${pkgs.curl}/bin/curl -sSf http://localhost:8000/readyz";
                        initial_delay_seconds = 2;
                        period_seconds = 10;
                        timeout_seconds = 4;
                        success_threshold = 1;
                        failure_threshold = 5;
                      };

                      availability = {
                        restart = "on_failure";
                      };
                    };
                  };

                  minio = {
                    process-compose = {
                      readiness_probe = {
                        exec.command = "${pkgs.curl}/bin/curl -sSf http://localhost:9000/minio/health/live";
                        initial_delay_seconds = 2;
                        period_seconds = 10;
                        timeout_seconds = 4;
                        success_threshold = 1;
                        failure_threshold = 5;
                      };

                      availability = {
                        restart = "on_failure";
                      };
                    };
                  };
                };
              };
            };
          };
        };
    };
}
