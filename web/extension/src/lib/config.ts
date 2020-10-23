export const captureConfig = (fps: number) => {
    return {
        video: true,
        audio: true,
        videoConstraints: {
            mandatory: {
                maxFrameRate: fps,
                minFrameRate: fps,
            },
        },
    };
};

export const mediaRecorderOptions = {
    audioBitsPerSecond: 128000,
    videoBitsPerSecond: 250000,
    mimeType: "video/webm;codecs=vp9",
};
