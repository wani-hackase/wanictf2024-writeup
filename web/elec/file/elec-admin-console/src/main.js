const { app, BrowserWindow } = require("electron");
const path = require("node:path");

const pageUrl = process.argv.at(-1);

// this is required to run in the container
app.disableHardwareAcceleration();

const createWindow = () => {
	const win = new BrowserWindow({
		width: 800,
		height: 600,
		webPreferences: {
			preload: path.join(__dirname, "preload.js"),
			contextIsolation: false,
			sandbox: false,
		},
	});

	win.webContents.once("did-finish-load", () => {
		console.log("Page loaded!");
	});

	win.loadURL(pageUrl);
};

app.whenReady().then(() => {
	createWindow();

	app.on("activate", () => {
		if (BrowserWindow.getAllWindows().length === 0) {
			createWindow();
		}
	});
});

app.on("window-all-closed", () => {
	if (process.platform !== "darwin") {
		app.quit();
	}
});
