export const tab: Tab = {
    status: "none",
    initTimeoutId: 0,
    extraInfo: {},
    filename: "rebirth",
    fps: 30,
    mediaRecorder: null,
};

type Tab = {
    status: "none" | "open" | "ready" | "start";
    initTimeoutId: number;
    extraInfo: Record<string, string>;
    filename: string;
    fps: number;
    mediaRecorder: MediaRecorder | null;
};
