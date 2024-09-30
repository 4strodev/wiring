> [!WARNING]
> This readme needs changes

# Wiring
A container for DI and autowiring written in Go. The main purpose is to provide an interface and a default implementation
for autowiring. You can use the default implementation or create your own if the requirements of your project are not
satisfied.

## Download

    go get github.com/4strodev/wiring@latest

## Usage
`Container` is the base interface, it provides all the methods necessary to start using DI in your project.

1. Instantiate a container. To do this wiring provides a default implementation defined in the `pkg/wiring` package.

## Thanks
This project was heavily inspired by [goloby/container](https://github.com/golobby/container).
