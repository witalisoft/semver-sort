# semver-sort

Simple tool to sort of a provided list of semantic versions from stdin.

## Usage

```console
Usage: ./semver-sort [options]
  -help
        Show help message
  -reverse
        Reverse the order of the sorted versions

Reads versions from stdin and prints them to stdout in sorted order.
```

## Example

```console
$ echo "1.2.3\n0.1.2\n6.7.8\n0.1.3\n1.2.3-rc.1"  | ./semver-sort
0.1.2
0.1.3
1.2.3-rc.1
1.2.3
6.7.8
```
