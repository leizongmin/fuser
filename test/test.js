const fs = require("node:fs");
const os = require("node:os");
const path = require("node:path");
const { expect } = require("chai");
const fuser = require("../");

describe("fuser", function () {
  it("buildMap should return a map", async function () {
    const map = await fuser.buildMap();
    expect(map).to.be.an("object");
    expect(Object.keys(map).length).to.be.greaterThan(0);
  });

  it("update should update the cacheMap", async function () {
    await fuser.update();
  });

  it("getPath should return an array of pids", async function () {
    const tmpDir = fs.mkdtempSync(path.join(os.tmpdir(), "fuser-"));
    const tmpFile = path.join(tmpDir, "test.txt");
    {
      await fuser.update();
      const pids = await fuser.getPath(tmpFile);
      expect(pids).to.be.null;
    }
    {
      const fd = fs.openSync(tmpFile, "a+");
      const pids = await fuser.getPath(tmpFile);
      expect(pids).to.be.null;

      await fuser.update();
      const pids2 = await fuser.getPath(tmpFile);
      // console.log(await fuser.buildMap(), tmpFile, pids2);
      expect(pids2).to.be.deep.eq([process.pid]);

      fs.closeSync(fd);
      await fuser.update();
      const pids3 = await fuser.getPath(tmpFile);
      expect(pids3).to.be.null;
    }
  });
});