// All parameters received from the web page
export type TabIPC = {
    action: "init" | "start" | "pause" | "resume" | "stop" | "fail" | "filename" | "extraInfo";
    filename?: string;
    extraInfo?: {
        [key in string]: string;
    }
}
