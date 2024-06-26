const { chromium } = require("playwright");
const Redis = require("ioredis");
const connection = new Redis({
  host: process.env.REDIS_HOST,
  port: process.env.REDIS_PORT,
});

const APP_URL = process.env.APP_URL; // application URL
const HOST = process.env.HOST; // HOST
const FLAG = process.env.FLAG; // FLAG

const crawl = async (path) => {
  const browser = await chromium.launch();
  const page = await browser.newPage();
  const cookie = [
    {
      name: "flag",
      value: FLAG,
      domain: HOST,
      path: "/",
      expires: Date.now() / 1000 + 100000,
    },
  ];
  page.context().addCookies(cookie);
  try {
    await page.goto(APP_URL + path, {
      waitUntil: "domcontentloaded",
      timeout: 3000,
    });
    await page.waitForTimeout(1000);
    await page.close();
  } catch (err) {
    console.error("crawl", err.message);
  } finally {
    await browser.close();
    console.log("crawl", "browser closed");
  }
};

(async () => {
  while (true) {
    console.log(
      "[*] waiting new url",
      await connection.get("queued_count"),
      await connection.get("proceeded_count"),
    );
    await connection
      .blpop("url", 0)
      .then((v) => {
        const path = v[1];
        console.log("crawl", path);
        return crawl(path);
      })
      .then(() => {
        console.log("crawl", "finished");
        return connection.incr("proceeded_count");
      })
      .catch((e) => {
        console.log("crawl", e);
      });
  }
})();
