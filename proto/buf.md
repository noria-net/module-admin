# Protobufs

This is the public protocol buffers API for [Module-Admin](https://github.com/noria-net/admin).

## Download

The `buf` CLI comes with an export command. Use `buf export -h` for details

#### Examples:

Download cosmwasm protos for a commit:
```bash
buf export buf.build/noria-net/admin:${commit} --output ./tmp
```

Download all project protos:
```bash
buf export . --output ./tmp
```