import { fetchPost } from "@/lib/util";

/**
 * Request the server to obtain the recording task
 * @return {Promise<RecordInfo>}
 */
export const getRecordTasks = (): Promise<RecordInfo> => {
    return new Promise((resolve, reject) => {
        // link: /internal/app/server/recordInfo.go
        fetchPost("/recordInfo", {}).then(async resp => {
            const data: RecordInfo = await resp.json();
            return resp.ok
                ? resolve(data)
                : reject(new Error("recordInfo response status not is normal"));
        });
    });
};

export const taskComplete = (): void => {
    void fetchPost("/complete", {});
};

export const taskFailed = (): void => {
    void fetchPost("/failed", {});
};

export const sendLog = (level: "debug" | "info" | "warn" | "error", message: string): void => {
    void fetchPost("/log", {
        level,
        message,
    });
};

export type RecordInfo = {
    url: string;
    startTime: string;
    endTime: string;
    filename: string;
    fps: number;
};
