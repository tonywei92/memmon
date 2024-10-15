# memmon

![build status](https://github.com/tonywei92/memmon/actions/workflows/go.yml/badge.svg)

![Memory Monitor Screenshot](screenshot.png)

`memmon` is a utility designed to monitor your Memory usage and provide notifications through beeps and/or by executing a specified command.

For example, it can be useful for sending an email by triggering an API with a `curl` command that corresponds to this action.

> Interested in monitoring your system CPU too? Check out [cpumon](https://github.com/tonywei92/cpumon)!

## Download

Head to [release](https://github.com/tonywei92/memmon/releases/latest) page, choose the appropriate binary for your operating system and architecture from the available options, and extract.

## Usage

Example:

```sh
# call `echo` command if memory is more than 80% usage
memmon -notify-on-memory-more-than=80 -no-beep -notify-with="echo memory usage high > output.log"
```

### Flags

- `-interval`: Interval in seconds (default: 1)
- `-notify-on-memory-more-than`: Notify when memory usage exceeds this percentage (default: 0)
- `-no-beep`: Disable beep notification (default: false)
- `-notify-with`: Command to run on notification, will not executed if not assigned

## Contributing

Thank you for considering contributing to `memmon`! We welcome contributions in the form of bug reports, feature requests, and pull requests. Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute to this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
