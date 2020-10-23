import { tab } from "@/data/tab";
import { TabIPC } from "@/ipc/type";
import action from "@/action";
import { sendLog } from "@/lib/ajax";

export default () => {
    // Listening to webpage messages
    chrome.runtime.onConnect.addListener(port => {
        port.onMessage.addListener((data: TabIPC) => {
            if (["start", "pause", "resume", "stop", "fail"].includes(data.action)) {
                sendLog("info", `current action is: ${data.action}`);
                return action[data.action as "start" | "pause" | "resume" | "stop" | "fail"]();
            }

            if (data.action === "extraInfo") {
                sendLog("info", `set extra info: ${JSON.stringify(data.extraInfo)}`);
                tab.extraInfo = {
                    ...tab.extraInfo,
                    ...data.extraInfo,
                };
                return;
            }

            if (data.action === "filename") {
                sendLog("info", `set output filename: ${data.filename}`);
                tab.filename = data.filename as string;
            }

            if (data.action === "init") {
                sendLog("debug", "website call init function");
                clearTimeout(tab.initTimeoutId);
            }
        });
    });
};
