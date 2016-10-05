# wharf

![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)
[![build status](https://git.itch.ovh/itchio/wharf/badges/master/build.svg)](https://git.itch.ovh/itchio/wharf/commits/master)
[![coverage report](https://git.itch.ovh/itchio/wharf/badges/master/coverage.svg)](https://git.itch.ovh/itchio/wharf/commits/master)

wharf is a protocol for incrementally transferring software builds over
the network using minimal time/bandwidth.

It is used in production at <https://itch.io> to allow creators to
quickly iterate & players to keep their library always up-to-date.

This repository contains the reference [golang][] implementation of the wharf
protocol, along with the reference protobuf definition files.

[golang]: https://golang.org/

The complete spec is available online, as a book:

  * :memo: <http://docs.itch.ovh/wharf/master/>

And can be contributed to via its GitHub repository:

  * :evergreen_tree: <https://github.com/itchio/wharf-spec>

## See also

butler is the <https://itch.io> command-line tool and is a wharf client.
It's the easiest way to try out wharf without having to code anything yourself.

  * <https://github.com/itchio/butler>

## Hacking on wharf

wharf is a pretty typical golang project, all its dependencies are open-source,
it even has a few tests.

### Regenerating protobuf code

```bash
protoc --go_out=. pwr/*.proto
```

protobuf v3 is required, as we use the 'proto3' syntax.

## License

Licensed under MIT License, see `LICENSE` for details.

Contains modified code from the following projects:

  * [kardianos/rsync](https://bitbucket.org/kardianos/rsync/) (BSD) - golang implementation of the rsync algorithm

