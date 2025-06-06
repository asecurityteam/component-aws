<a id="markdown-component-aws---settings-component-for-aws-clients" name="component-aws---settings-component-for-aws-clients"></a>
# component-aws - Settings component for AWS clients
[![GoDoc](https://godoc.org/github.com/asecurityteam/component-aws?status.svg)](https://godoc.org/github.com/asecurityteam/component-aws)

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=bugs)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=code_smells)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=coverage)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=ncloc)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=alert_status)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=security_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=sqale_index)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-aws&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=asecurityteam_component-aws)


<!-- TOC -->

- [component-aws - Settings component for AWS clients](#component-aws---settings-component-for-aws-clients)
    - [Overview](#overview)
    - [Quick Start](#quick-start)
    - [Status](#status)
    - [Contributing](#contributing)
        - [Building And Testing](#building-and-testing)
        - [License](#license)
        - [Contributing Agreement](#contributing-agreement)

<!-- /TOC -->

<a id="markdown-overview" name="overview"></a>
## Overview

This is a [`settings`](https://github.com/asecurityteam/settings) component that
enables constructing AWS credentials/sessions and some of the clients we
regularly use when building services.

<a id="markdown-quick-start" name="quick-start"></a>
## Quick Start

```golang
package main

import (
    "context"
    "net/http"

    aws "github.com/asecurityteam/component-aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/asecurityteam/settings"
)

func main() {
    ctx := context.Background()
    envSource := settings.NewEnvSource(os.Environ())

    sess, _ := aws.NewSession(ctx, envSource)
    d := dynamodb.New(sess)
    // or
    db, _ := aws.NewDynamoDB(ctx, envSource)
    // or
    s3 := aws.NewS3(ctx, envSource)
}
```

<a id="markdown-status" name="status"></a>
## Status

This project is in incubation which means we are not yet operating this tool in
production and the interfaces are subject to change.

<a id="markdown-contributing" name="contributing"></a>
## Contributing

<a id="markdown-building-and-testing" name="building-and-testing"></a>
### Building And Testing

We publish a docker image called [SDCLI](https://github.com/asecurityteam/sdcli) that
bundles all of our build dependencies. It is used by the included Makefile to help
make building and testing a bit easier. The following actions are available through
the Makefile:

-   make dep

    Install the project dependencies into a vendor directory

-   make lint

    Run our static analysis suite

-   make test

    Run unit tests and generate a coverage artifact

-   make integration

    Run integration tests and generate a coverage artifact

-   make coverage

    Report the combined coverage for unit and integration tests

<a id="markdown-license" name="license"></a>
### License

This project is licensed under Apache 2.0. See LICENSE.txt for details.

<a id="markdown-contributing-agreement" name="contributing-agreement"></a>
### Contributing Agreement

Atlassian requires signing a contributor's agreement before we can accept a patch. If
you are an individual you can fill out the [individual
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d).
If you are contributing on behalf of your company then please fill out the [corporate
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b).
