# stats

`stats` - Outputs statistical information about line delimited numbers from stdin

## Warning

During computation, `stats` will store an entry for each line delimited item encountered, so memory usage will need be taken into consideration.

## Example usage

    $ echo "10\n3\n3\n2\n6\n8\n4\n10\n11" | stats
    count   9
    sum 57
    p99 11.0000
    p97 11.0000
    p95 11.0000
    min 2.0000
    max 11.0000
    average 6.3333
    median  6.0000
    stddev  3.2998


    $ echo "10\n3\n3\n2\n6\n8\n4\n10\n11" | stats --output average --output max
    6.3333  
    11.0000

See `--help` for further information
