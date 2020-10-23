/**
 * Post request encapsulation
 * @param {string} path - request path
 * @param {object} data - Post body
 * @return {Promise<Response>>}
 */
export const fetchPost = (path: string, data: any) => {
    return fetch(`http://127.0.0.1:9182${path}`, {
        body: JSON.stringify(data),
        method: "POST",
        cache: "no-cache",
        headers: {
            "content-type": "application/json",
        },
        mode: "cors",
        credentials: "include",
    });
};
