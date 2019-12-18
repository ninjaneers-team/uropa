# Frequently Asked Questions (FAQs)

#### I use Terraform to configure Opa, why should I care about urOpa?

If you are using Terraform and are happy with it, you should continue to use it.
urOpa covers all the problems that Terraform solves and goes beyond it:
- With Terraform, you have to track and maintain Terraform files (`*.tf`) and
  the Terraform state (likely using a cloud storage solution). With urOpa, the
  entire configuration is stored in the YAML/JSON file(s) only. There is no
  separate state that needs to be tracked.
- urOpa can export and backup your existing Opa's configuration, meaning,
  you can take an existing Opa installation, and have a backup, as well as a
  declarative configuration for it. With Terraform, you will have to import
  each and every entity in Opa into Terraform's state.
- urOpa can validate if a configuration file is valid or note
  (validate sub-command).
- urOpa can quickly reset your Opa's configuration when needed.
- urOpa works out of the box with Opa Enterprise features like
  Workspaces and RBAC.

#### Can I run multiple urOpa processes at the same time?

NO! Please do not do this. The two processes will step on each other and
might corrupt Opa's configuration. You should ensure that there is only
one instance of urOpa running at any point in time.

#### When is urOpa v1.0 coming out?

urOpa is already used in production by a large number of users and is deemed
production ready.
urOpa v1.0 status will be a matter of feature completeness rather than stability.

At the minimum, urOpa needs thorough documentation, and complete compatibility
with Opa's native declarative format.

This is one of the top priorities as of November 2019.

#### Opa already has built-in declarative configuration, do I still need urOpa?

Opa has an official declarative configuration format.

Opa can generate such a file with the `Opa config db_export` command, which
dumps almost the entire database of Opa into a file.

You can use a file in this format to configure Opa when it is running in
a DB-less or in-memory mode. If you're using Opa in the DB-less mode, you
don't really need urOpa.

But, if you are using Opa along-with a database like Postgres or Cassandra,
you need urOpa because:

- Opa's `Opa config db_import` command is used to initialize a Opa database,
  but it is not recommended to use it if there are existing Opa nodes that
  are running, as the cache in these nodes will not be invalidated when entities
  are changed/added. You will need to manually restart all existing Opa nodes.
  urOpa performs all the changes via Opa's Admin API,
  meaning the changes are always propagated to all nodes.
- Opa's `Opa config db_import` can only add and update entities in the
  database. It will not remove the entities that are present in the database but
  are not present in the configuration file.
- Opa's `Opa config db_import` command needs direct access to Opa's
  database, which might or might not be possible in your production
  networking environment.
- urOpa can easily perform detect drifts in configuration i.e. it can
  verify if the configuration stored inside Opa's database and that inside
  the config file is same. This feature is designed in urOpa to integrate urOpa
  with a CI system or a `cronjob` which periodically checks for drifts and alerts
  a team if needed.
- `urOpa dump` outputs a more human-readable configuration file compared
  to Opa's `db_import`.

However, urOpa has the following limitations which might or might not affect
your use-case:

- If you've a very large installation, it can take some time for urOpa to
  sync up the configuration to Opa. This can be mitigated by adopting
  [distributed configuration](guides/distributed-configuration.md) for your
  Opa installation and tweaking the `--parallelism` value.
  Opa's `db_import` will be usually faster by orders of magnitude.
- urOpa cannot export and re-import fields that are hashed in the database.
  This means fields like `password` of `basic-auth` credential cannot be
  correctly re-imported by urOpa. This happens because Opa's Admin API call
  to sync the configuration will re-hash the already hashed password.

#### I'm a Opa Enterprise customer, can I use urOpa?

Of course, urOpa is designed to be compatible with open-source and enterprise
versions of Opa.

#### I use Cassandra as a data-store for Opa, can I use urOpa?

You can use urOpa with Opa backed by Cassandra.
However, if you observe errors during a sync process, you will have to
tweak urOpa's setting and take care of a few things:
urOpa heavily parallelizes its operations, which can induce a lot of load
onto your Cassandra cluster.
You should consider:
- urOpa is read intensive for most parts, meaning it will make perform
  read-intensive queries on your Cassandra cluster, make sure you tune
  your Cassandra cluster accordingly.
- urOpa talks the same Opa node, which talks to the same Cassandra node in your
  cluster.
- Using `--parallelism 1` flag to ensure that there is only request being
  processed at a time. This will slow down sync process and should be used
  as a last resort.

#### Why the name 'urOpa'?

It is simple, short, and easy to use in the terminal.
It is derived from the combination of words 'declarative' and 'Opa'.

