# `Packr2`テスト用
[packr/v2 at master · gobuffalo/packr · GitHub](https://github.com/gobuffalo/packr/tree/master/v2)を実際に利用して見るためのリポジトリです。

## ツールなどのインストール
```
$ go get -u github.com/gobuffalo/packr/v2/packr2
```

## 事前準備
`./templates/admin/index.html`を作成します:

```
kazu634@bastion% ll templates/admin/
total 4.0K
drwxr-xr-x 2 kazu634 kazu634  0 Aug  3 23:08 .
drwxr-xr-x 2 kazu634 kazu634  0 Aug  3 23:19 ..
-rwxr-xr-x 1 kazu634 kazu634 15 Aug  3 23:08 index.html
kazu634@bastion% cat templates/admin/index.html
foo

bar

buzz
```

## 使い方
こんな感じで使います:

```
package main

import (
  "fmt"
  "log"

  "github.com/gobuffalo/packr/v2"
)

func main() {
  box := packr.New("myBox", "./templates")

  s, err := box.FindString("admin/index.html")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(s)
}
```

## ビルドなどの仕方
何もしないと、作成したバイナリは`./templates/admin/index.html`を探してしまうため、バイナリにデータを埋め込みます。

実施する手順は以下になります:

1. `packr2`コマンドを実行して、データ埋め込み用ソースを作成する
2. `go build`でバイナリを作成する
3. `packr2 clean`で生成したデータ埋め込み用ソースを削除する

```
kazu634@bastion% packr2
kazu634@bastion% go build -o packr-test main.go
kazu634@bastion% ll
total 9.8M
drwxr-xr-x 2 kazu634 kazu634    0 Aug  4 08:42 .
drwxr-xr-x 2 kazu634 kazu634    0 Aug  3 22:56 ..
-rwxr-xr-x 1 kazu634 kazu634   59 Aug  4 08:32 .envrc
-rwxr-xr-x 1 kazu634 kazu634 2.1K Aug  4 08:16 go.mod
-rwxr-xr-x 1 kazu634 kazu634  29K Aug  4 08:16 go.sum
-rwxr-xr-x 1 kazu634 kazu634  242 Aug  3 23:09 main.go
-rwxr-xr-x 1 kazu634 kazu634  243 Aug  4 08:41 main-packr.go
drwxr-xr-x 2 kazu634 kazu634    0 Aug  4 08:41 packrd
-rwxr-xr-x 1 kazu634 kazu634 9.8M Aug  4 08:42 packr-test
-rwxr-xr-x 1 kazu634 kazu634  707 Aug  3 22:56 Rakefile
drwxr-xr-x 2 kazu634 kazu634    0 Aug  3 23:19 templates
drwxr-xr-x 2 kazu634 kazu634    0 Aug  4 08:11 vendor
kazu634@bastion% packr2 clean
kazu634@bastion% ./packr-test
foo

bar

buzz

```

## 参考
- [packr/v2 at master · gobuffalo/packr · GitHub](https://github.com/gobuffalo/packr/tree/master/v2)

