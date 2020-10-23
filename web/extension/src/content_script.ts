import { TabIPC } from "@/ipc/type";

if (typeof chrome.runtime.id !== "undefined") {
    const port = chrome.runtime.connect(chrome.runtime.id);

    // Process plug-in messages
    port.onMessage.addListener(msg => {
        if ({}.toString.call(msg) !== "[object Object]") {
            return;
        }

        if (msg.error) {
            console.error(`[rebirth plugin]: ${JSON.stringify(msg.error)}`);
            return;
        }
    });

    // The message of the webpage message is forwarded to the plug-in for processing.
    window.addEventListener(
        "message",
        (event: { data: TabIPC }) => {
            if (Object.keys(event.data).length === 0) {
                return;
            }

            if (event.data.action) {
                port.postMessage(event.data);
            }
        },
        false,
    );

    const injectedScript = document.createElement("script");
    injectedScript.src = chrome.extension.getURL("injected.js");
    (document.head || document.documentElement).prepend(injectedScript);
}
