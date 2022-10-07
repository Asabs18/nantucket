{ pkgs }: {
	deps = [
  pkgs.gnuplot_aquaterm
  pkgs.clang_12
		pkgs.ccls
		pkgs.gdb
		pkgs.gnumake
	];
}