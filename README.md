# grogg-tdd
A TDD demo repository.

## Setup
Run `./setup.sh` to install the needed dependencies. You can then run the tests via the `ginkgo -r` command.

## Using This Repository
Each commit of the master branch represents a step in time in the TDD process. Start from the first commit on master.

As you go forward in the commit history, this README will be updated with an explanation of the new code state added to the top of the Code State section below.

## JSON-Schema Registry Feature Acceptance Criteria
- A user should be able to create a new JSON Schema by providing a JSON string
- The schema is stored in our database only if it is a valid [JSON-Schema](https://json-schema.org) document

## Code State
- Implemented the stub of a JSON Validator interface, and used it to get the test passing.
- Introduced the first test block related to JSON Schema validation. It is currently failing, but is ready for implementation.
- Basic repo setup with just the framework of tests in place, but no real functionality.