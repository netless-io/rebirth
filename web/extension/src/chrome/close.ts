/**
 * get chrome all window id
 */
const getAllWindowsId = (): Promise<number[]> => {
    return new Promise(resolve => {
        chrome.windows.getAll(
            {
                windowTypes: ["normal", "popup", "devtools"],
            },
            windows => {
                resolve(windows.map(w => w.id));
            },
        );
    });
};

/**
 * close chrome window by window id
 * @param ids
 */
const closeWindows = (ids: number[]) => {
    ids.forEach(id => {
        chrome.windows.remove(id);
    });
};

export default async () => {
    closeWindows(await getAllWindowsId());
};
