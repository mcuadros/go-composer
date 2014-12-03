go-composer [![Build Status](https://travis-ci.org/mcuadros/go-composer.png?branch=master)](https://travis-ci.org/mcuadros/go-composer)
===========

Basic replacement of [Composer](https://getcomposer.org/), the Dependency Manager for PHP.

This is not more than a proof of concept to try to build a faster dependency manager for PHP

The constrains version parsing is handled by [`go-version`](https://github.com/mcuadros/go-version) a port of the PHP [`version_compare`](http://php.net/manual/es/function.version-compare.php) function and Version comparsion classes from original Composer project

