# Documentation Carbon PKG / Device project
Developing a package manager is not a common programming task.

carbonpkg is a package manager designed to simplify the process of installing software hosted on GitHub. It provides users with a convenient interface to install, update, and remove packages, as well as manage dependencies. In this documentation, we'll take a deep dive into key aspects of how carbonpkg works, including package specifications, directories, visibility, registry, versions, local files, source of truth, checksums, and dependencies.

This documentation covers in detail the key components of how carbon_pkg works, including the package specification, directories, visibility, registry, versioning, local files, checksums, and the process of installing and updating packages.

## Goals and objectives
The goals of our package manager:

+ **Simplified package management** - It may often be necessary to create a sandbox or virtual environment for testing packages. Virtual packages do not affect the overall list of system packages. The user can also change some parameters of an already installed project, if the developer allows this.
+ **Security** - checksums are used and verification of the integrity and originality of the package before installation. In the future, we also plan to add a search for package vulnerabilities during installation.
+ **Versioning support** - support the ability to have packages of different versions on the same system.

Tasks:

1. Provide integration with Git and GitHub.
2. Providing tools for managing dependencies.
3. Flexible configuration and integration into different systems.
4. Make packages and the package manager itself more independent of the system and platform.
5. Configure a system to protect against package damage (for example, the installation was completed for external reasons).
6. Make the package manager universal and convenient.

## Project architecture
At the time of planning, carbon_pkg consists of several key components:

1. Package specification parser (json, toml, yaml): loading and processing the package specification, as well as parsing the user configuration of the package manager.
2. RNP - trusted package registry (TPR). The registry contains two types of packages: Diamond-package and Monoxide-Package. The names are derived from the chemical compounds of carbon. Diamond-package is a reliable, tested package, while monoxide-package is a dangerous package that contains vulnerabilities or other problems. This marks packages until they are removed or the problem is resolved. Even ordinary packages in which a serious vulnerability has been discovered can become Monoxide packages.
3. Client and server parts: the server contains the RNP and all operations pass through it. The client is responsible for interaction with the user and his system.
4. Versioning system: a mechanism that allows you to work with packages. Especially useful if you have a package with multiple versions.
5. Security: responsible for checksums, vulnerabilities and protection against packet corruption.

## Structure of the package manager configuration file
By default it is stored in `~/.config/carbon_pkg/default/carbonpkg_settings.json`:

```json
{
"install_monoxide": false,
"always_skip_trp": false,
"timeout": 30,
"security" {
"verify_checksums": true,
},
"install": {
"install_location": "/usr/local/bin"
},
"backup": {
"enabled": true,
"backup_location": "~/.carbon_pkg_backups"
}
}
```

+ `install_monoxide` - allows potentially dangerous monoxide packages to be installed on the system.
+ `always_skip_trp` - the package skips the Trusted Package Registry check step and goes straight to installation.
+ `timeout` - timeout for waiting for a response from the registry server.
+ `security` - security settings.
- `verify_checksums` - whether to check checksums.
+ `install` - installation settings.
- `install_location` - package installation location.
+ `backup` - backup settings.
- `enabled` - enable backups.
- `backup_location` - location where backups are saved.

## Package specification structure
The package specification file in json format looks like this:

```json
{
"owner": "alexeev-pkg",
"name": "carbon_pkg",
"package_type": "binary",
"description": "Blazing fast and modern package manager written in Golang",
"maintainer": "Alexeev Bronislav <alexeev.dev@mail.ru>"
"tags": ["package manager", "package", "golang", "fast"]
"version": "latest",
"languages": ["go"],
"supported_architectures": ["amd64", "arm"],
"license": "GNU LGPL v2.1",
"repository": "https://github.com/alexeev-prog/carbon_pkg",
"homepage": "https://github.com/alexeev-prog/carbon_pkg",
"user-configurable": true,

"raw_source": {
"type": "git",
"url": "{repository}.git",
"branch": "main"
},

"assets": {
"default": "github-release",
"github-release": {
"path": "{repository}/releases/download/{version}",

"files": {
"linux-amd64": "carbonpkg-{version}-linux-amd64.tar.gz"
}

"checksums": "checksums.txt"
},
"building": {
"commands": [],
"output_files": [],
"building_instructions_file": "BUILDING.md"
}
},
}
```

+ `owner` - owner of the package.
+ `name` - package name.
+ `package_type` - package type.
+ `description` - a brief description of the project.
+ `maintainer` - information about the project maintainer.
+ `tags` - list of package tags.
+ `version` - version (latest or specific version) of the package.
+ `languages` - list of languages in which the project source code is written.
+ `supported_architectures` - supported processor architectures.
+ `license` - license.
+ `repository` - package repository.
+ `homepage` - the home page of the package.
+ `user-configurable` - allows the user to change some project specification parameters for themselves.
+ `raw_source` - raw sources:
- `type` - type.
- `url` - URL address
- `branch` - branch (required for type: git).
+ `assets` - project assets. Different types are supported (github-release, building):
- `default` - default type.
- `github-release` - installing a release via github api:
- `path` is the path (link).
- `files` - dictionary with files.
- `checksums` - checksums
- `building` - assembly (unstable):
- `commands` - list of commands to build.
- `output_files` - list of output files.
- `building_instructions_file` - building instructions file (the package manager will print it if an error occurs during automatic building).

---

[Contents](./index.md)
