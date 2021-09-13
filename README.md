# webauthncf 
This repo is still in progress. As of right now, nothing works. Come back later!
A CICD-friendly WebAuthn command-line tool for testing WebAuthn conformance against FIDO Alliance standards.
## Installation
MacOS
```bash
brew install webauthncf
```
Linux
```
<list 500 ways to install>
```
Download the binary
```
wget https://some.cdn.maybe.github.release?
```
Build from source
```
git clone <url>
... build steps
```
## Usage
webauthncf offers 3 different modules for testing WebAuthn compliance.
```
USAGE:
    webauthncf [SUBCOMMAND] [OPTIONS]
FLAGS:
    -h, --help       Prints this message
SUBCOMMANDS:
    server    Runs the assertion server module
    files     Runs the files assertion module
    api       Runs the api assertion module
```
### Running as a server
This module providers a server and HTTP endponts to asert WebAuthn options and response.
- `POST http://localhost:9000/assert/attestation`
- `POST http://localhost:9000/assert/attestation`
The bodies of the request _may_ include an options key and _must_ include a response key. If the options key does not exist, the server will not assert that the response is conformant with the options that were provided.
**Example Request Body**:
```json
{
    // Optional: options used for attestation or asseretion response
    "options" { ... },
    // Required: attestation or assertion response created from the above options
    "response": { ... }
}
```
**CLI Reference**
```
USAGE:
    webauthncf server [OPTIONS]
FLAGS:
    -h, --help       Prints this message
    -v               Sets the level of verbosity
OPTIONS:
    -p, --port <PORT>    Set the port for the server
    -o, --output <PATH>  Output assertion results to a file
```
### Running against a rest API
**CLI Reference**
```
USAGE:
    webauthncf client [OPTIONS] <API BASE URL>
FLAGS:
    -h, --help       Prints this message
    -v               Sets the level of verbosity
OPTIONS:
    --assertion-options-url <>      Set URL to be used generating assertion options
    --attestation-options-url <>    Set URL to be used generating attestation options
    --assertion-url <>              Set URL to be used generating WebAuthn assertions
    --attestation-url <>            Set URL to be used generating WebAuthn attestations
    -o, --output <PATH>  Output assertion results to a file
```
### Running against files
The `file` module providers the ability to assert against static JSON files produced by a WebAuthn server. By default, the runner looks for JSON files in the following directories of the base directory provided:
- `./assertion`
- `./attestation`
The runner will assert against every JSON file found in the directory. The files _may_ include an options key and _must_ include a response key (see payload expectations for the server module).
**CLI Reference**
```
webauthncf files -d ./data/ 
Options <todo>
```
## Contributing
## License
[MIT](https://choosealicense.com/licenses/mit/)s
# webauthncf_cli
