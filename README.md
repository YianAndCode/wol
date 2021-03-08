# wol
WOL （网络唤醒）服务

## 编译和运行

### Linux/macOS

```bash
make
```

### Windows

```cmd
go build -o ./bin/wol.exe .
```

### 运行

```bash
./bin/wol
# 或者在后台运行（Linux）
./bin/wol &
```

## 唤醒机器

假设运行本程序的机器内网 IP 地址为 `10.0.0.10`，则访问 `http://10.0.0.10:9394/wake?mac={你的MAC地址}`

## iOS 捷径

通过添加 iOS 捷径，你可以使用 Siri 来帮你唤醒局域网中的机器，捷径的设置非常简单：

![捷径设置](screenshots/iOS_Shortcut.jpg)

## 鸣谢

`magic_packet.go` 由 [https://github.com/sabhiram/go-wol/](https://github.com/sabhiram/go-wol/) 提供！