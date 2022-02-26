#!/usr/bin/env node

"use strict"

const request = require('request'),
    path = require('path'),
    tar = require('tar'),
    zlib = require('zlib'),
    mkdirp = require('mkdirp'),
    fs = require('fs'),
    exec = require('child_process').exec;

// Mapping from Node's `process.arch` to Golang's `$GOARCH`
const ARCH_MAPPING = {
    //"ia32": "386",
    "x64": "amd64",
    //"arm": "arm",
    "arm64": "arm64"
};

// Mapping between Node's `process.platform` to Golang's 
const PLATFORM_MAPPING = {
    //"darwin": "darwin",
    "linux": "linux",
    "win32": "windows",
    //"freebsd": "freebsd"
};

function validateConfiguration(packageJson) {

    if (!packageJson.version) {
        return "'version' property must be specified";
    }

    if (!packageJson.service || typeof(packageJson.service) !== "object") {
        return "'service' property must be defined and be an object";
    }

    if (!packageJson.service.version) {
        return "'version' property must be specified";
    }

    if (!packageJson.service.name) {
        return "'name' property is necessary";
    }

    if (!packageJson.service.path) {
        return "'path' property is necessary";
    }

    if (!packageJson.service.url) {
        return "'url' property is required";
    }

    // if (!packageJson.bin || typeof(packageJson.bin) !== "object") {
    //     return "'bin' property of package.json must be defined and be an object";
    // }
}

function parsePackageJson() {
    if (!(process.arch in ARCH_MAPPING)) {
        console.error("Installation is not supported for this architecture: " + process.arch);
        return;
    }

    if (!(process.platform in PLATFORM_MAPPING)) {
        console.error("Installation is not supported for this platform: " + process.platform);
        return
    }

    const packageJsonPath = path.join(".", "package.json");
    if (!fs.existsSync(packageJsonPath)) {
        console.error("Unable to find package.json. ");
        return
    }

    let packageJson = JSON.parse(fs.readFileSync(packageJsonPath));
    let error = validateConfiguration(packageJson);
    if (error && error.length > 0) {
        console.error("Invalid package.json: " + error);
        return
    }

    // We have validated the config. It exists in all its glory
    let binName = packageJson.service.name;
    let binPath = packageJson.service.path;
    let url = packageJson.service.url;
    let version = packageJson.service.version;
    if (version[0] === 'v') version = version.substr(1);  // strip the 'v' if necessary v0.0.1 => 0.0.1

    // Binary name on Windows has .exe suffix
    if (process.platform === "win32") {
        binName += ".exe"
    }

    // Interpolate variables in URL, if necessary
    url = url.replace(/{{arch}}/g, ARCH_MAPPING[process.arch]);
    url = url.replace(/{{platform}}/g, PLATFORM_MAPPING[process.platform]);
    url = url.replace(/{{version}}/g, version);
    url = url.replace(/{{bin_name}}/g, binName);

    return {
        binName: binName,
        binPath: path.join(process.cwd(),"bin"),
        url: url,
        version: version
    }
}

/**
 * Reads the configuration from application's package.json,
 * validates properties, downloads the binary, untars, and stores at
 * ./bin in the package's root. NPM already has support to install binary files
 * specific locations when invoked with "npm install -g"
 *
 *  See: https://docs.npmjs.com/files/package.json#bin
 */
const INVALID_INPUT = "Invalid inputs";
function install(callback) {
    
    let opts = parsePackageJson();
    if (!opts) return callback(INVALID_INPUT);

    mkdirp.sync(opts.binPath);
    let ungz = zlib.createGunzip();
    let untar = tar.x({cwd: opts.binPath});

    ungz.on('error', callback);
    untar.on('error', callback);

    console.log("Downloading from URL: " + opts.url);
    let req = request({uri: opts.url});
    req.on('error', callback.bind(null, "Error downloading from URL: " + opts.url));
    req.on('response', function(res) {
        if (res.statusCode !== 200) return callback("Error downloading binary. HTTP Status Code: " + res.statusCode);

        req.pipe(ungz).pipe(untar);
    });

}

function uninstall(callback) {
    let opts = parsePackageJson();
    try {
        fs.unlinkSync(path.join(opts.binPath, opts.binName));
    } catch(ex) {
        // Ignore errors when deleting the file.
    }

    return callback(null);
}


// Parse command line arguments and call the right method
let actions = {
    "install": install,
    "uninstall": uninstall
};

let argv = process.argv;
if (argv && argv.length > 2) {
    let cmd = process.argv[2];
    if (!actions[cmd]) {
        console.log("Invalid command to go-npm. `install` and `uninstall` are the only supported commands");
        process.exit(1);
    }

    actions[cmd](function(err) {
        if (err) {
            console.error(err);
            process.exit(1);
        } else {
            process.exit(0);
        }
    });
}



