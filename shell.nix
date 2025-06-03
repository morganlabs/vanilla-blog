{
  pkgs ? import <nixpkgs> { },
}:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    zsh
  ];

  shellHook = ''
    zsh
  '';
}
