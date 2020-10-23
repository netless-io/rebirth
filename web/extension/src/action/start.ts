import { captureConfig, mediaRecorderOptions } from "@/lib/config";
import { tab } from "@/data/tab";
import { sendLog } from "@/lib/ajax";

export default () => {
    chrome.tabCapture.capture(captureConfig(tab.fps), stream => {
        if (stream === null) {
            sendLog("error", `chrome capture fail: ${JSON.stringify(chrome.runtime.lastError)}`);
            return;
        }

        const recordedBlobs: BlobPart[] = [];
        const mediaRecorder = new MediaRecorder(stream, mediaRecorderOptions);

        mediaRecorder.ondataavailable = event => {
            if (event.data && event.data.size > 0) {
                recordedBlobs.push(event.data);
            }
        };

        mediaRecorder.requestData()

        mediaRecorder.onstop = () => {
            const superBuffer = new Blob(recordedBlobs, {
                type: "video/webm",
            });

            // For the time being, we will not use the chrome.downloads.download API for downloading, because this API currently has bugs.
            // see: https://bugs.chromium.org/p/chromium/issues/detail?id=892133#makechanges
            const link = document.createElement("a");
            link.href = URL.createObjectURL(superBuffer);
            link.setAttribute("download", `${tab.filename}.webm`);
            link.click();

            // fileDownloadDone(fileName)
            //     .then(() => {
            //         completeRecordTask(fileName, id);
            //     })
            //     .catch((e: Error) => {
            //         sendLog("error", `download webm video failed: ${e.message}`);
            //     });
        };

        tab.mediaRecorder = mediaRecorder;
    });
};
