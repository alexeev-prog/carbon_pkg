# Carbon PKG🚀
<a id="readme-top"></a> 

<div align="center">  
  <p align="center">
    Blazing fast and modern package manager written in Golang
    <br />
	<a href="./docs/en/index.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="#-key-features">Key Features</a>
    ·
    <a href="#-getting-started">Getting Started</a>
    ·
    <a href="#-usage-examples">Basic Usage</a>
    ·
    <a href="#-specifications">Specification</a>
    ·
    <a href="https://github.com/alexeev-prog/carbon_pkg/blob/main/LICENSE">License</a>
  </p>
</div>
<br>
<p align="center">
    <img src="https://img.shields.io/github/languages/top/alexeev-prog/carbon_pkg?style=for-the-badge">
    <img src="https://img.shields.io/github/languages/count/alexeev-prog/carbon_pkg?style=for-the-badge">
    <img src="https://img.shields.io/github/license/alexeev-prog/carbon_pkg?style=for-the-badge">
    <img src="https://img.shields.io/github/stars/alexeev-prog/carbon_pkg?style=for-the-badge">
    <img src="https://img.shields.io/github/issues/alexeev-prog/carbon_pkg?style=for-the-badge">
    <img src="https://img.shields.io/github/last-commit/alexeev-prog/carbon_pkg?style=for-the-badge">
</p>

## Check Other My Projects

 + [SQLSymphony](https://github.com/alexeev-prog/SQLSymphony) - simple and fast ORM in sqlite (and you can add other DBMS)
 + [Burn-Build](https://github.com/alexeev-prog/burn-build) - simple and fast build system written in python for C/C++ and other projects. With multiprocessing, project creation and caches!
 + [OptiArch](https://github.com/alexeev-prog/optiarch) - shell script for fast optimization of Arch Linux
 + [libnumerixpp](https://github.com/alexeev-prog/libnumerixpp) - a Powerful C++ Library for High-Performance Numerical Computing
 + [pycolor-palette](https://github.com/alexeev-prog/pycolor-palette) - display beautiful log messages, logging, debugging.
 + [shegang](https://github.com/alexeev-prog/shegang) - powerful command interpreter (shell) for linux written in C

## 🔧 Specifications

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
