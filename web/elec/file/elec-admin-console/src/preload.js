const { spawn } = require("node:child_process");

window.addEventListener("load", async () => {
	const versions = {
		app: "0.0.1",
		node: process.versions.node,
		chrome: process.versions.chrome,
		electron: process.versions.electron,
	};
	console.log(versions);

	const cp = spawn("uname", ["-a"]);
	console.log(cp);
	const kernelInfo = await loadStream(cp.stdout);

	document.getElementById("app-version").textContent = versions.app;
	document.getElementById("node-version").textContent = versions.node;
	document.getElementById("chrome-version").textContent = versions.chrome;
	document.getElementById("electron-version").textContent = versions.electron;
	document.getElementById("kernel-info").textContent = kernelInfo.toString();
	document.getElementById("admin-footer").classList.remove("d-none");
});

const loadStream = (s) =>
	new Promise((resolve, reject) => {
		const chunks = [];
		s.on("data", (chunk) => chunks.push(chunk));
		s.on("error", reject);
		s.on("end", () => resolve(Buffer.concat(chunks)));
	});
