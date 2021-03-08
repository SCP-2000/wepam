{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable-small";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachSystem [ "aarch64-linux" "x86_64-linux" ]
      (system:
        let pkgs = import nixpkgs { inherit system; }; in
        {
          devShell = pkgs.mkShell {
            buildInputs = with pkgs;[ gcc pam ];
          };
        });
}
