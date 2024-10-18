[![Banner](https://github.com/clouddrove/terraform-module-template/assets/119565952/67a8a1af-2eb7-40b7-ae07-c94cde9ce062)][website]
<h1 align="center">
    Enigma
</h1>

<p align="center">
	<a href="https://goreportcard.com/report/github.com/clouddrove/enigma">
		<img alt="Go Report Status" src="https://goreportcard.com/badge/github.com/clouddrove/enigma">
	</a>
	<a href="https://github.com/clouddrove/enigma/">
		<img alt="Build Status" src="https://img.shields.io/badge/test-passing-green">
	</a>
	<a href="https://join.slack.com/t/devops-talks/shared_invite/zt-2s2rnal1e-bRStDKSyRC~dpXA~PaJ7vQ">
		<img alt="Slack Chat" src="https://img.shields.io/badge/join%20slack-click%20here-blue">
	</a>
	<a href="https://medium.com/devops-talks/announcing-devopstalks-spectacular-hacktoberfest-2024-363a09223c45">
		<img alt="Blog" src="https://img.shields.io/badge/hacktoberfest2024%20blog-8A2BE2">
	</a>
	<a href="https://choosealicense.com/licenses/mit/">
		<img alt="Apache-2.0 License" src="http://img.shields.io/badge/license-MIT-brightgreen.svg">
	</a>
</p>

<p align="center">
<a href='https://facebook.com/sharer/sharer.php?u=https://github.com/clouddrove/enigma'>
  <img title="Share on Facebook" src="https://user-images.githubusercontent.com/50652676/62817743-4f64cb80-bb59-11e9-90c7-b057252ded50.png" />
</a>
<a href='https://www.linkedin.com/shareArticle?mini=true&title=enigma&url=https://github.com/clouddrove/enigma'>
  <img title="Share on LinkedIn" src="https://user-images.githubusercontent.com/50652676/62817742-4e339e80-bb59-11e9-87b9-a1f68cae1049.png" />
</a>
<a href='https://twitter.com/intent/tweet/?text=enigma&url=https://github.com/clouddrove/enigma'>
  <img title="Share on Twitter" src="https://user-images.githubusercontent.com/50652676/62817740-4c69db00-bb59-11e9-8a79-3580fbbf6d5c.png" />
</a>
</p>

# Enigma Workflow

Enigma Workflow is a GitHub Action designed to simplify and streamline your workflows. It provides a unified interface for building, tagging, and publishing Docker images across multiple cloud platforms including AWS, GCP, Azure, and Docker Hub.

## Overview

The workflow automates the process of building and publishing Docker images with support for:
- Multi-architecture builds
- Custom build arguments
- Cache control
- Multiple registry destinations
- Post-push cleanup

## Usage

To use this action in your workflow, create a `.github/workflows/docker.yml` file with the following structure:

```yaml
name: Docker Build and Publish

on:
  push:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Build and Push Docker Image
        uses: clouddrove/enigma@v0.0.17
        with:
          command: 'publish'  # or 'bake' for building only
          enigma_file: './path/to/your/.enigma'
          DOCKER_IMAGE: 'your-image-name'
          DOCKER_TAG: 'latest'
          # Additional parameters as needed
```

## Input Parameters

### Required Parameters

| Parameter | Description |
|-----------|-------------|
| `command` | Workflow command to execute. Options: `bake` (build only) or `publish` (build and publish) |
| `enigma_file` | Path to the Enigma configuration file |
| `DOCKER_IMAGE` | Name of the Docker image to build |
| `DOCKER_TAG` | Tag to apply to the Docker image |

### Optional Parameters

#### Docker Build Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `DOCKERFILE_PATH` | Path to the Dockerfile | `./Dockerfile` |
| `DOCKER_BUILD_ARCHITECTURE` | Target build architecture | `amd64` |
| `DOCKER_BUILD_ARGS` | Additional Docker build arguments | |
| `DOCKER_NO_CACHE` | Disable Docker build cache | `false` |
| `DOCKER_MULTI_ARCH_BUILD` | Enable multi-architecture builds | `false` |
| `DOCKER_CLEANUP` | Remove local images after push | `false` |

#### Registry Credentials

##### AWS ECR
| Parameter | Description |
|-----------|-------------|
| `AWS_REGION` | AWS region for ECR |
| `AWS_ACCOUNT_ID` | AWS account ID |

##### Google Cloud Platform
| Parameter | Description |
|-----------|-------------|
| `GCP_PROJECT_ID` | GCP project identifier |
| `GCP_REGION` | GCP region for Container Registry |

##### Azure Container Registry
| Parameter | Description |
|-----------|-------------|
| `AZURE_REGISTRY_NAME` | Name of Azure Container Registry |
| `AZURE_CLIENT_ID` | Azure service principal client ID |
| `AZURE_TENANT_ID` | Azure tenant ID |
| `AZURE_CLIENT_SECRET` | Azure service principal secret |

##### Docker Hub
| Parameter | Description |
|-----------|-------------|
| `DOCKER_USERNAME` | Docker Hub username |
| `DOCKER_PASSWORD` | Docker Hub password/token |

##### GitHub Container Registry
| Parameter | Description |
|-----------|-------------|
| `TOKEN` | GitHub personal access token |
| `GITHUB_USERNAME` | GitHub username |

## Examples

### Basic Usage
```yaml
- uses: clouddrove/enigma@v0.0.17
  with:
    command: 'publish'
    enigma_file: './.enigma'
    DOCKER_IMAGE: 'myapp'
    DOCKER_TAG: 'v1.0.0'
```

### Multi-Architecture Build
```yaml
- uses: clouddrove/enigma@v0.0.17
  with:
    command: 'publish'
    enigma_file: './.enigma'
    DOCKER_IMAGE: 'myapp'
    DOCKER_TAG: 'v1.0.0'
    DOCKER_MULTI_ARCH_BUILD: 'true'
    DOCKER_BUILD_ARCHITECTURE: 'linux/amd64,linux/arm64'
```

### Publishing to AWS ECR
```yaml
- uses: clouddrove/enigma@v0.0.17
  with:
    command: 'publish'
    enigma_file: './.enigma'
    DOCKER_IMAGE: 'myapp'
    DOCKER_TAG: 'v1.0.0'
    AWS_REGION: 'us-west-2'
    AWS_ACCOUNT_ID: '${{ secrets.AWS_ACCOUNT_ID }}'
```

## Enigma Configuration File

The `enigma_file` should contain your workflow-specific configurations. Example structure:

```yaml
DOCKERFILE_PATH=
DOCKER_IMAGE=
DOCKER_TAG=
DOCKER_BUILD_ARCHITECTURE=
DOCKER_SCAN=
DOCKER_BUILD_ARGS=
DOCKER_NOCACHE=
DOCKER_CLEANUP=
DOCKER_MULTI_ARCH_BUILD=
```

## Notes

1. For security reasons, always use GitHub Secrets for sensitive information like credentials and tokens.
2. When using multiple registries, ensure all required credentials are provided.
3. The `bake` command will only build the image without publishing.
4. Enable `DOCKER_CLEANUP` to automatically remove local images after successful push.

## Support

For issues and feature requests, please create an issue in the [CloudDrove/Enigma](https://github.com/clouddrove/enigma) repository.


## Installation

### Source (Go >= 1.18)

* Clone and Build the Repo
```bash
# Clone the repository
git clone https://github.com/clouddrove/enigma

# Navigate into the project directory
cd enigma

# Install any dependencies (though this is typically handled automatically by go build)
go get .

# Build the project and output the binary as 'enigma'
go build -o bin/enigma
```

* Give permission to execute the binary
```
cd bin/

chmod +x enigma
```

* Create symlink to `/usr/local/bin`
```shell
sudo ln -s enigma /usr/local/bin/enigma

# Verify the symlink
ls -l /usr/local/bin/enigma
```

## Quick start

### Overview

Enigma is a tool designed to simplify the DevOps lifecycle, offering a seamless way to manage Docker environments, build, scan, and publish Docker images. Below is a quick guide to getting started with Enigma and using its core commands.

---

### 1. **Initialize the Enigmafile**

This command initializes the Enigmafile where environment variables are stored.

- **Command:** enigma

- **Arguments:**

    - --enigmafile [file_to_generate_enigma_vars] (Optional, Default: .enigma)

- **Subcommand:** init

- **Options:**

    - --d: Initialize for Docker

**Example:**

```bash

enigma --enigmafile .test init --d

Output:

Using enigma file: .test

Environment variables successfully written to .test

```

---

### 2. **Bake Command**

The bake command is used to set up and build Docker images.

- **Command:** enigma

- **Subcommand:** bake

- **Aliases:** bake

- **Description:** Builds and scans a Docker image.

- **Arguments:** No additional arguments required.

- **Options:**

    - --d: Uses Docker when set.

**Docker-specific steps:**

- Loads the Docker environment.

- Builds the Docker image.

- Scans the Docker image for vulnerabilities.

**Example:**

```bash

enigma bake --d

```

---

### 3. **Publish Command**

The publish command is used to tag and push Docker images to a remote repository.

- **Command:** enigma

- **Subcommand:** publish

- **Aliases:** publish

- **Description:** Publishes Docker images by tagging and pushing them.

- **Arguments:** No additional arguments required.

- **Options:**

    - --d: Uses Docker when set.

**Docker-specific steps:**

- Loads the Docker environment.

- Tags the Docker image.

- Pushes the Docker image to a registry.

**Example:**

```bash

enigma publish --d

```

---

### 4. **Build-Publish Command**

The build-publish command combines the steps of building, scanning, tagging, and publishing a Docker image in one command.

- **Command:** enigma

- **Subcommand:** build-publish

- **Description:** Builds, scans, and optionally publishes Docker images.

- **Arguments:** No additional arguments required.

- **Options:**

    - --d: Uses Docker when set.

**Docker-specific steps:**

- Loads the environment variables from the Enigmafile.

- Builds and scans the Docker image.

- If the PUBLISH environment variable is set to true, tags and pushes the image.

**Example:**

```bash

enigma bake-publish --d

```


## Options

```shell
Enigma is a tool designed to simplify the DevOps lifecycle, offering a seamless way to manage tools environments, 
        build, scan, and publish. Below is a quick guide to getting started with Enigma and using its core commands.

Usage:
  enigma [flags]
  enigma [command]

Available Commands:
  bake          To Bake the command
  bake-publish  To bake and publish
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  init          To init the command
  publish       To publish

Flags:
      --enigmafile string   Path to the .enigma file (default ".enigma")
  -h, --help                help for enigma

Use "enigma [command] --help" for more information about a command.
```





## ✨ Contributors

Big thanks to our contributors for elevating our project with their dedication and expertise! But, we do not wish to stop there, would like to invite contributions from the community in improving these projects and making them more versatile for better reach. Remember, every bit of contribution is immensely valuable, as, together, we are moving in only 1 direction, i.e. forward.

<a href="https://github.com/clouddrove/enigma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=clouddrove/enigma&max" />
</a>
<br>
<br> 

If you're considering contributing to our project, here are a few quick guidelines that we have been following (Got a suggestion? We are all ears!):

- **Fork the Repository:** Create a new branch for your feature or bug fix.
- **Coding Standards:** You know the drill.
- **Clear Commit Messages:** Write clear and concise commit messages to facilitate understanding.
- **Thorough Testing:** Test your changes thoroughly before submitting a pull request.
- **Documentation Updates:** Include relevant documentation updates if your changes impact it.


## Feedback
Spot a bug or have thoughts to share with us? Let's squash it together! Log it in our [issue tracker](https://github.com/clouddrove/enigma/issues), feel free to drop us an email at [hello@clouddrove.com](mailto:hello@clouddrove.com).

Show some love with a ★ on [our GitHub](https://github.com/clouddrove/enigma)!  if our work has brightened your day! – your feedback fuels our journey!

## Join Our Slack Community

Join our vibrant open-source slack community and embark on an ever-evolving journey with CloudDrove; helping you in moving upwards in your career path.
Join our vibrant Open Source Slack Community and embark on a learning journey with CloudDrove. Grow with us in the world of DevOps and set your career on a path of consistency.

🌐💬What you'll get after joining this Slack community:

- 🚀 Encouragement to upgrade your best version.
- 🌈 Learning companionship with our DevOps squad.
- 🌱 Relentless growth with daily updates on new advancements in technologies.

Join our tech elites [Join Now][slack] 🚀

## Explore Our Blogs

Click [here][blog] :books: :star2:

## Tap into our capabilities
We provide a platform for organizations to engage with experienced top-tier DevOps & Cloud services. Tap into our pool of certified engineers and architects to elevate your DevOps and Cloud Solutions.

At [CloudDrove][website], has extensive experience in designing, building & migrating environments, securing, consulting, monitoring, optimizing, automating, and maintaining complex and large modern systems. With remarkable client footprints in American & European corridors, our certified architects & engineers are ready to serve you as per your requirements & schedule. Write to us at [business@clouddrove.com](mailto:business@clouddrove.com).

<p align="center">We are <b> The Cloud Experts!</b></p>
<hr />
<p align="center">We ❤️  <a href="https://github.com/clouddrove">Open Source</a> and you can check out <a href="https://registry.terraform.io/namespaces/clouddrove">our other modules</a> to get help with your new Cloud ideas.</p>

[website]: https://clouddrove.com
[blog]: https://blog.clouddrove.com
[slack]: https://www.launchpass.com/devops-talks
[github]: https://github.com/clouddrove
[linkedin]: https://cpco.io/linkedin
[twitter]: https://twitter.com/clouddrove/
[email]: https://clouddrove.com/contact-us.html
[terraform_modules]: https://github.com/clouddrove?utf8=%E2%9C%93&q=terraform-&type=&language=
