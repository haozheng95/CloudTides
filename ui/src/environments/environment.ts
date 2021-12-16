import { base } from './base';

// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
const url = window.location.href.split('/')
console.log(url)
const http = url[0]+ ':'
const ip = url[1]
export const environment = {
  production: false,
  apiPrefix: http+ip+'/api/v1',
  webssh: http+ip+':8888',
  apiIp: ip
  // apiPrefix: 'http://127.0.0.1:8033/api/v1',
  // webssh: 'http://127.0.0.1:8888',
  // apiIp: '127.0.0.1:8033'
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
