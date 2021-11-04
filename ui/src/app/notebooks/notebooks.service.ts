import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { environment } from '@tide-environments/environment';
import { Router } from '@angular/router'
import { tap, map } from 'rxjs/operators';
import { FormBuilder, Validators } from '@angular/forms'
import { LOCAL_STORAGE_KEY } from '@tide-config/const';
import { WebSocketService } from '@tide-shared/service/web-socket.service'
@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor(private readonly http: HttpClient,private readonly router: Router, private fb: FormBuilder,public ws:WebSocketService) { }
  createInstanceFlag:boolean = false
  modifiable = true
  appList:AppModel[] = []
  appLogs: LogModel[] = []
  currentToken = ''
  loading = true
  wsUrl='ws://localhost:8033/api/v1/ws/application/instance/'
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
  getAppList () {
    return this.http.get(environment.apiPrefix + '/application/instance', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).
    pipe(
      tap(data => {})
    )
  }
  createNewApp (form) {
    return this.http.post(environment.apiPrefix + '/application/instance', form, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {})
    )
  }
  modifyApp (form) {
    return this.http.put(environment.apiPrefix + '/application/instance', form, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {})
    )
  }
  deleteApp (token) {
    return this.http.delete(environment.apiPrefix + `/application/instance/${token}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {})
    )
  }
  getApplictionList () {
    this.appList = [] 
    this.getAppList().subscribe((data:AppModel[]) => {
      data.forEach(el => {
        const port = el.link.split(':')[1]
        el.port = port.split('/')[0]
        el.logo = 'assets/img/jupyter.svg',
        // el.link = 'http://' + el.link.split('?')[0]
        console.log('link', el);
        
        this.appList.push(el)
      })
    })
  }
  getAppLogs (token: string) {
    this.currentToken = token
    return this.http.get(environment.apiPrefix + '/application/instance/'+ token,{
      headers: {
        Authorization: `Bearer ${localStorage.getItem(LOCAL_STORAGE_KEY.TOKEN)}`
      }
    }).pipe(
      tap(data => {})
    ).subscribe(
      (data: any[]) => {        
        this.appLogs = data
        this.loading = false
        this.buildWS(token)
      }
    )
  }
  unsub: any
  buildWS (token) {
    this.ws.connect(this.wsUrl+token)
    this.unsub = this.ws.messageSubject.subscribe(
      data => {
        // 剔除第一条，加入最新一条
        this.appLogs.shift()
        this.appLogs.push(data)
      }
    )
  }
  closeWs () {
    this.ws.onClose(false)
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

}
interface LogModel {
  content: string
  date: string
  leve: string
  source: string
}
