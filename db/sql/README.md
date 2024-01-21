# Installation

## Requirements

- linux server - tested on Debian 12
- postgresql server - tested on 13

## Setup

```
    psql -U postgres -h localhost -a -f /route/to/archive/sql.sql -W
```