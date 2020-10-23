import { getRecordTasks, sendLog, taskFailed } from "@/lib/ajax";
import { tab } from "@/data/tab";

const openTab = () => {
    getRecordTasks().then(resp => {
        chrome.tabs.create(
            {
                url: resp.url,
            },
            () => {
                tab.status = "open";
                tab.filename = resp.filename;
                tab.fps = resp.fps;
                tab.initTimeoutId = window.setTimeout(() => {
                    sendLog(
                        "error",
                        "Five minutes after the browser was opened, the rebirth.init function was not called, resulting in a timeout. Please make sure that the website loads normally or view the webpage code normally calls the rebirth.init function",
                    );
                    taskFailed();
                }, 1000 * 60 * 5);

                sendLog("info", "chrome open url");
            },
        );
    });
};

export default () => {
    // waiting 3m, make sure everything is ok
    setTimeout(openTab, 1000 * 3);
};
