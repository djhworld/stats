# stats

`stats` - Outputs statistical information about line delimited numbers from stdin

## Warning

During computation, `stats` will store an entry for each line delimited item encountered, so memory usage will need be taken into consideration.

## Example usage

    $ echo "10\n3\n3\n2\n6\n8\n4\n10\n11" | stats
    count   sum     p99     p97     p95     min     max     avg     median  stddev
    9       57.0000 11.0000 11.0000 11.0000 2.0000  11.0000 6.3333  6.0000  3.2998

See `--help` for further information
