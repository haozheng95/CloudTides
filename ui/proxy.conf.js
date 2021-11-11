// @ts-check


const {
  npm_package_config_port: PORT = 4200,
  start_mode: MODE = 'local',
} = process.env;

const endpoint = {
  local: 'http://0.0.0.0:3000',
  e2e: 'http://0.0.0.0:8081',
  webssh: 'http://127.0.0.1:8888'
};

const target = endpoint[MODE];

console.log(`
==========================================
Starting Mode: [${MODE}]

  /api/       -> ${target}
  /api-local/ -> ${endpoint.local}
==========================================
`);

const PROXY_CONFIG = {

  '/api-local/': {
    target: endpoint.local,
    changeOrigin: true,
    secure: false,
    pathRewrite: {
      '^/api-local': '/api',
    },
  },

  '/api/': {
    target: target,
    changeOrigin: true,
    secure: false,
    pathRewrite: {
      '^/api': '/api',
    },
    onProxyRes(proxyResponse) {
      const cookies = proxyResponse.headers['set-cookie'];
      const prune = (cookie = '') => cookie.replace(/;\W*secure/gi, '');

      if (cookies) {
        proxyResponse.headers['set-cookie'] = cookies.map(prune);
      }
    }
  },
  '/api-webshh/': {
    target: endpoint.webssh,
    changeOrigin: true,
    secure: false,
    ws: true,
    pathRewrite: {
      '^/api-webshh/': '',
    },

  }

}

module.exports = PROXY_CONFIG;
