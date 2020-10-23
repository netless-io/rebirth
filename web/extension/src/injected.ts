type NoParamsFuncName = "init" | "start" | "pause" | "resume" | "stop" | "fail";

type rebirth = {
    [key in NoParamsFuncName]: () => void;
} & {
    filename: (f: string) => void;
} & {
    extraInfo: (
        d: {
            [key in string]: string;
        },
    ) => void;
};

interface MyWindow extends Window {
    rebirth: rebirth;
}

(window as MyWindow & typeof globalThis).rebirth = {} as rebirth;

["init", "start", "pause", "resume", "stop", "fail"].forEach(m => {
    (window as MyWindow & typeof globalThis).rebirth[m as NoParamsFuncName] = () => {
        const msg = {
            action: m,
        };
        window.postMessage(msg, "*");
    };
});
