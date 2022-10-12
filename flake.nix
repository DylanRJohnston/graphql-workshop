{
  inputs.nixpkgs.url = "github:nixos/nixpkgs/release-22.05";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }: flake-utils.lib.eachDefaultSystem (system:
    let pkgs = import nixpkgs { inherit system; }; in
    {
      devShell = pkgs.mkShell {
        packages = with pkgs; [
          go_1_18
          golangci-lint
          air
          delve
          go-outline
          gopkgs
          gopls
          gotests
          impl
        ];
      };
    }
  );
}
