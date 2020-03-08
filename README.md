gopacparser
===========

gopacparser - это обертка для библиотеки `gopac`, которая обрабатывает данные 
возращаемые функцией FindProxyForURL (функция объявленная внутри PAC файла) и
приводит их в вид удобный для дальнейшего использования.

Пример обработанных данных:

```go
map[string]string{
    "http": "http://proxy-nossl.antizapret.prostovpn.org:29976",
    "https": "https:http://proxy-nossl.antizapret.prostovpn.org:29976",
}
```

```go
map[string]string{
    "http": "socks5://proxy-nossl.antizapret.prostovpn.org:29976",
    "https": "socks5:http://proxy-nossl.antizapret.prostovpn.org:29976",
}
```

## Использование

Для того, чтобы использовать данную библиотеку необходимо просто вызвать 
функцию `FindProxy`.

Пример:

```go
proxy, _ := FindProxy("https://antizapret.prostovpn.org/proxy.pac", "http://filmix.me")
```

```go
proxy, _ := FindProxy("/some/path/file.pac", "http://filmix.me")
```
