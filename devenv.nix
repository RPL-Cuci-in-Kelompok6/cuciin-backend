{ pkgs, ... }:

{
  # https://devenv.sh/basics/
  env.GREET = "devenv";

  # https://devenv.sh/packages/
  packages = [ pkgs.git pkgs.gcc ];

  # https://devenv.sh/scripts/
  scripts.hello.exec = "echo hello from $GREET";

  enterShell = ''
    hello
    git --version
    go version
    gcc --version | head -n1
  '';

  # https://devenv.sh/languages/
  languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";

  # See full reference at https://devenv.sh/reference/options/
}
