import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders, HttpRequest} from '@angular/common/http'
import {environment} from '@tide-environments/environment';
import {Router} from '@angular/router'
import {tap} from 'rxjs/operators';
import {FormBuilder, Validators} from '@angular/forms'
import {LOCAL_STORAGE_KEY} from '@tide-config/const';
import {WebSocketService} from '@tide-shared/service/web-socket.service'
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor(private readonly http: HttpClient, private readonly router: Router, private fb: FormBuilder, public ws: WebSocketService) {
  }

  createInstanceFlag: boolean = false
  modifiable = true
  appList: AppModel[] = []
  appLogs: LogModel[] = []
  currentToken = ''
  currentModel = ''
  loading = true
  url = environment.apiIp ? environment.apiIp : ''
  wsUrl = 'ws://' + this.url + '/api/v1/ws/application/instance/'
  instanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: [''],
    sshHost: [''],
    sshPassword: [''],
    sshPort: [''],
    sshUser: [''],
    token: ['']
  })
  createInstanceTitle = 'HOME.NOTEBOOKS.Create'

  getAppList() {
    return this.http.get(environment.apiPrefix + '/application/instance', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }

  createNewApp(form) {
    return this.http.post(environment.apiPrefix + '/application/instance', form, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }

  modifyApp(form) {
    return this.http.put(environment.apiPrefix + '/application/instance', form, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }

  deleteApp(token) {
    return this.http.delete(environment.apiPrefix + `/application/instance/${token}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }

  getApplictionList() {
    this.appList = []
    this.getAppList().subscribe((data: AppModel[]) => {
      data.forEach(el => {
        // Store hash
        localStorage.setItem(el.token, el.token)
        this.appList.push(el)
      })
    })
  }

  getAppLogs(token: string) {
    this.currentToken = token
    return this.http.get(environment.apiPrefix + '/application/instance/' + token, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    ).subscribe(
      (data: any[]) => {
        this.appLogs = data
        this.loading = false
        this.buildWS(token)
      }
    )
  }

  unsub: any

  buildWS(token) {
    this.ws.connect(this.wsUrl + token)
    this.ws.flag = true
    this.unsub = this.ws.messageSubject.subscribe(
      data => {
        // 剔除第一条，加入最新一条
        if (this.appLogs.length > 200) {
          this.appLogs.shift()
        }

        this.appLogs.push(data)
      }
    )
  }

  closeWs() {
    this.ws.onClose(false)
  }

  uploadData(file: any, token: string): Observable<any> {
    // httpOptionsMultipart.headers = httpOptionsMultipart.headers.delete('Content-Type');
    // return this.http.post('/data', 
    //   formData, httpOptionsMultipart);
    const formData: any = new FormData();
    formData.append('name', file.name);
    formData.append('file', file);
    const req = new HttpRequest('POST', environment.apiPrefix + '/application/instance/file/' + token, formData, {
      reportProgress: true,
    })
    req.headers.set('Authorization', `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`)
    const newReq = req.clone({
      headers: new HttpHeaders({
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      })
    })
    console.log('newReq', newReq);

    return this.http.request(newReq)
  }

  getFileList(token: string) {
    return this.http.get(environment.apiPrefix + `/application/instance/file/${token}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }
  getHostNameList () {
    return this.http.get<Data>(environment.apiPrefix + `/application/instance/file`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {
      })
    )
  }
}

interface AppModel {
  instanceName: string
  link: string
  token: string
  logo: string
  port: string,
  appType: string,
  sshHost: string,
  sshPassword: string
  sshPort: string
  sshUser: string
  extra: null | ExtraModel
}

interface ExtraModel {
  appType: string,
  sshHost: string,
  sshPassword: string
  sshPort: string
  sshUser: string
  cmd: string
}

interface LogModel {
  content: string
  date: string
  leve: string
  source: string
}
interface Data {
  data: any
  status: number
  messgae: string
}
