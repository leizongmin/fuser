const fs = require("node:fs");
const path = require("node:path");

async function readAllPids() {
  const pids = await fs.promises.readdir("/proc");
  return pids.map((pid) => parseInt(pid, 10)).filter((pid) => !isNaN(pid));
}

async function readAllOpenFilesFromPid(pid) {
  try {
    const fdPath = path.join("/proc", pid.toString(), "fd");
    const fds = await fs.promises.readdir(fdPath);
    const files = await Promise.all(
      fds.map(async (fd) => {
        try {
          const p = path.join(fdPath, fd);
          const link = await fs.promises.readlink(p);
          return link;
        } catch (e) {
          return null;
        }
      }),
    );
    return files.filter((p) => p !== null);
  } catch (e) {
    return [];
  }
}

async function buildMap(options = {}) {
  const map = {};
  const pids = await readAllPids();
  for (const pid of pids) {
    const files = await readAllOpenFilesFromPid(pid);
    for (const file of files) {
      if (options.filter && !options.filter(file)) {
        continue;
      }
      if (!map[file]) {
        map[file] = [];
      }
      map[file].push(pid);
    }
  }
  return map;
}

let cacheMap = {};

async function update(options = {}) {
  cacheMap = await buildMap(options);
}

function getPath(p) {
  if (!/^\w+:\[\d+\]$/g.test(p)) {
    p = path.resolve(p);
  }
  if (cacheMap[p]) {
    if (cacheMap[p].length > 0) {
      return cacheMap[p];
    }
  }
  return null;
}

exports.buildMap = buildMap;
exports.update = update;
exports.getPath = getPath;
