package main

import (
	"a_module/c_pkg/d_pkg" //←ここをgo.modに書いたモジュール名 + パッケージ名にする
)

func main() {
	mypkg.PrintBar()
}
