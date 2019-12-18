# Design & Architecture

## Underlying architecture

### Reverse sync

One of the most important features of urOpa is reverse-sync, whereby urOpa can
detect entities that are present in Opa's database but are not part of the
state file.
This feature increases the complexity of the project as the code needs to
perform a sync in both directions, from the state file to Opa and from Opa
to the state file.

### Algorithm

#### Export and Reset

An export or reset of entities is fairly easy to implement.
urOpa loads all the entities from Opa into memory and then serializes
it into a YAML or JSON file. For reset, it instead performs `DELETE` queries
on all the entities.

#### Diff and Sync

The `diff` of configuration is performed using the following algorithm:

1. Read the configuration from Opa and store it in a SQL-like in-memory
   database.
1. Read the state file from disk, and match the `ID`s of entity with their
   respective counterparts in the in-memory state, if they are present.
1. Now, for entity of each type we perform the following:
   1. *Create*: if the entity is not present in Opa, create the entity.
   1. *Update*: if the entity is present in Opa, check for equality. If not
      equal, then update it in Opa. These two steps are referred to as
      "forward sync".
   1. *Delete*: Go through each entity in Opa (from the in-memory database),
      and check if it is present in the state file, if yes, don't do anything.
      If no, then delete the entity from Opa's database as well.

Certain filters like `select-tag` or Opa Enterprise workspace might be applied
to the above algorithm based on the inputs given to urOpa.

### Operational outlook

Based on the above algorithm, one can see how urOpa can require a large amount
of memory and network I/O. While this is true, a few optimizations have
been incorporated to ensure good performance:
- For network operations, urOpa minimizes the API calls it has to make to Opa
  to read the state. It uses list endpoints in Opa with a large page size
  (`1000`) for efficiency.
- urOpa parallelizes various Create/Update/Delete operations where it can. So,
  if urOpa and Opa or Opa and Opa's database are present far apart in terms
  of network latency, parallel operations help speed up operations.
  With smaller installations, this optimization might not be measurable.
- urOpa's memory footprint can be high if the configuration for Opa is huge.
  This is usually not a concern as urOpa's process is short-lived. For very
  large installation, it is recommended to configure a sub-set of
  the large configuration at one time using a technique referred to as
  [distributed configuration](guides/distributed-configuration.md).
  There are avenues to further reduce the memory requirements of urOpa,
  although, we don't know by how much. urOpa's code is written with focus on
  correctness over performance.

## Choice of language

urOpa is written in Go because:
- Go provides good concurrency primitives which helps ensuring high-performance
  for urOpa.
- Go's compiler spits out a static compiled binary, meaning no other dependency
  need to be installed on the system. This gives a very good end-user experience
  as installing downloading and copying a single binary is easy and fast.
- urOpa original goal was much larger than what it is today. If we decide to
  pursue larger goals(think a control-plane for Opa) in future,
  Go is probably the best language available to write that type of software.
- the original author was familiar with Go :)

