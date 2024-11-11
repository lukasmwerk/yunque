## Architecture
### DBs
- Thinking of using Postgres for the main stuff, since it's acid compliant, does relational well and can handle concurrency well, also its just generally the move. Use it for the OLTP and shop/catalog config stuff.

- Cassandra cuz it might be fun. Use it for real-time analytics bc of its high write throughput, just strean analytics relevant data from the main db and query offload for analytics (creating a near real-time analytics/reporting db that doesn't bog down the main db). Should be fun getting creative with this shiz.

- Prolly have to use kafka for the postgres - cassandra pipelining

- Maybe try some redis for shits and giggles. Message Queues, API & Rate limiting, interactions, dist locks?, fraud detection?, lower shopping cart abandonmnent with better snappiness/performance, just general caching, check out GAP case study.

- Other caching tricks? Data import/export, selling data >:-)

## Notes
- Context and Logging, need to pass the logger and ctx, going to implement basics without these injections but revisit them later and add logs everywhere.
- Waiting on mutex adding, implemented the basics of a logging mutex, but will wait to add the muteces until more of the structure is implemented