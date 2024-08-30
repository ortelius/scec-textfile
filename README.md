# Ortelius v11 Textfile Microservice

> Version 11.0.0

RestAPI for the Readme, License and Swagger objects.  Only for new objects.  Retrieval will be done directly against the db by other microservices.
![Release](https://img.shields.io/github/v/release/ortelius/scec-textfile?sort=semver)
![license](https://img.shields.io/github/license/ortelius/.github)

![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-textfile/build-push-chart.yml)
[![MegaLinter](https://github.com/ortelius/scec-textfile/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-textfile/actions?query=workflow%3AMegaLinter+branch%3Amain)
![CodeQL](https://github.com/ortelius/scec-textfile/workflows/CodeQL/badge.svg)
[![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-textfile)

![Discord](https://img.shields.io/discord/722468819091849316)

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/msapi/license](#postmsapilicense) | Create a License |
| POST | [/msapi/swagger](#postmsapiswagger) | Create a Swagger |
| POST | [/msapi/textfile](#postmsapitextfile) | Create a Textfile |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |

## Path Details

***

### [POST]/msapi/license

- Summary  
Create a License

- Description  
Create a new License and persist it

#### Responses

- 200 OK

***

### [POST]/msapi/swagger

- Summary  
Create a Swagger

- Description  
Create a new Swagger and persist it

#### Responses

- 200 OK

***

### [POST]/msapi/textfile

- Summary  
Create a Textfile

- Description  
Create a new Textfile and persist it

#### Responses

- 200 OK

## References
