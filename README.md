# grogg-tdd
A TDD demo repository.

## Setup
Run `./setup.sh` to install the needed dependencies

## Using This Repository
Each commit of the master branch represents a step in time in the TDD process. Start from the first commit on master.

As you go forward in the commit history, this README will be updated with an explanation of the new code state added to the top of the Code State section below.

## JSON-Schema Registry Feature Acceptance Criteria
- A user should be able to create a new JSON Schema by providing a JSON string
- The schema is stored in our database only if it is a valid [JSON-Schema](https://json-schema.org) document

## Code State
- Impelment use of ORM abstraction, but ORM still not actually implemented, so the build app still fails.
- Add ORM abstraction to DataRepository, implemented first test for DataRepository#Save, but failing due to lack of implementation.
- Finished implementing the RegisterSchema, using a stub of a DataRepository. All tests passing, but the buit app encounters an error like `panic: runtime error: invalid memory address or nil pointer dereference` because the DataRepository is not actually implemented.
- Added a new test that introduces a mock of our data repository abstraction. At this point, the entire RegisterSchema behavior is now tested (but not yet implemented).
- Implemented the stub of a JSON Validator interface, and used it to get the test passing.
- Introduced the first test block related to JSON Schema validation. It is currently failing, but is ready for implementation.
- Basic repo setup with just the framework of tests in place, but no real functionality.