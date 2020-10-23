const path = require('path');

const resolvePath = (...relativePath) => path.resolve(__dirname, '..', ...relativePath);

module.exports = {
    dist: resolvePath('dist'),
    appSrc: resolvePath('src'),
    backgroundFile: resolvePath('src', 'background.ts'),
    contentFile: resolvePath('src', 'content_script.ts'),
    injectedFile: resolvePath('src', 'injected.ts'),
    tsConfig: resolvePath('tsconfig.json'),
    devEnv: resolvePath('config', '.env.dev'),
    devProd: resolvePath('config', '.env.prod'),
};
