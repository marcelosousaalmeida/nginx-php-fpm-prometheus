# php-fpm-prometheus

Simple [PHP-FPM](http://php.net/manual/en/install.fpm.php) and [NGINX](https://www.nginx.com/) status exporter for [Prometheus](https://prometheus.io/).

## Installation

If you are using Go 1.6+ (or 1.5 with the `GO15VENDOREXPERIMENT=1` environment variable), you can install `php-fpm-prometheus` with the following command:

```bash
$ go get -u github.com/marceloalmeida/nginx-php-fpm-prometheus
```

## Usage

```bash
$ ./nginx-php-fpm-prometheus --help
Usage of ./nginx-php-fpm-prometheus:
  -addr="0.0.0.0:8080": IP/port for the HTTP server
  -fpm-status-url="": PHP-FPM status URL
  -nginx-status-url="": Nginx status URL

$ ./nginx-php-fpm-prometheus -status-url "http://example.com/status" -addr "127.0.0.1:8080"
```

Finally, point Prometheus to `http://127.0.0.1:8080/metrics`.

## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.
