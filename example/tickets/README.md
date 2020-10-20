# [goesgen](github.com/romshark/goesgen) example Tickets

This example demonstrates the use of [goesgen](github.com/romshark/goesgen). It models a simple JIRA-like ticket system in an event-oriented & event-sourced way. The following architectural paradigms are applied:
- [Command Query Responsibility Segregation (CQRS)](https://martinfowler.com/bliki/CQRS.html)
- [Event Sourcing](https://martinfowler.com/eaaDev/EventSourcing.html)
- [Microservices](https://en.wikipedia.org/wiki/Microservices)
- Transactions and [strong consistency guarantees](https://en.wikipedia.org/wiki/Strong_consistency) through [Optimistic Concurrency Control](https://en.wikipedia.org/wiki/Optimistic_concurrency_control)

Responsibilities are split up into 2 separate services:
- [**Users**](https://github.com/romshark/goesgen/tree/1.0.0/example/tickets/service/users) is responsible for creating new users. It uses an in-memory [SQLite 3](https://www.sqlite.org/index.html) database for integration SQL demonstration purposes.
- [**Tickets**](https://github.com/romshark/goesgen/tree/1.0.0/example/tickets/service/tickets) is responsible for managing tickets (creating tickets, assigning & unassigning users, closing tickets, updating tickets). It uses an in-memory data structure as a volatile store.

## Schema and code generation

The [schema YAML file](https://github.com/romshark/goesgen/blob/1.0.0/example/tickets/schema.yml) defines the schema of the entire system describing all event types, projections, services and service methods.

Package [`generated`](https://github.com/romshark/goesgen/tree/1.0.0/example/tickets/generated) is generated by executing the generator in the directory containing [schema.yml](https://github.com/romshark/goesgen/blob/1.0.0/example/tickets/schema.yml):
```bash
% goesgen
2020/10/20 17:30:33 package successfully generated: generated
```