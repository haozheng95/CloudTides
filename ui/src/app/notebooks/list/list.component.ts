import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
import { environment } from '@tide-environments/environment';
@Component({
  selector: 'tide-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

  constructor(public nd: NotebooksService, private router:Router) {
  }

  ngOnInit(): void {
    this.nd.getApplictionList()
  }
  selected= 'selected'
  sureDeleteFalg = false
  filterSearchValue: string = ''
  token = ''
  toJupyter (app: AppModel) {
    const link = app.link.split('?')[0]    
    let form = document.createElement('form')
    form.action='http://'+ link
    form.method='get'
    form.target = '_blank'
    const body = document.documentElement
    body.appendChild(form)
    const input = document.createElement('input')
    input.type = 'hidden'
    input.name = 'token'
    input.value = app.token
    form.appendChild(input)
    setTimeout(() => {
      form.submit()
      form = null
    }, 1000)
  }
  toGromacs (app: AppModel) {
    const link = environment.webssh   
    let form = document.createElement('form')
    form.action= link
    form.method='get'
    form.target = '_blank'
    const body = document.documentElement
    body.appendChild(form)
    const opts = {
      hostname: app.sshHost,
      port: app.sshPort,
      username: app.sshUser,
      password: app.sshPassword,
      privatekey: '',
      // passphrase: '',
      totp: ''
    }
    for (const key in opts) {
      const input = document.createElement('input')
      input.type = 'hidden'
      input.name = key
      input.value = opts[key]
      form.appendChild(input)
    }

    setTimeout(() => {
      form.submit()
      form = null
    }, 1000)

  }
  modifyApp (app: AppModel) {
    this.nd.createInstanceFlag = true
    this.nd.createInstanceTitle = 'HOME.APPLICATION.Modify'
    this.nd.modifiable = false
    console.log('app', app);
    
    this.nd.instanceForm.setValue({
      instanceName: app.instanceName,
      port: app.port,
      appType: app.appType,
      sshHost: app.sshHost,
      sshPassword: app.sshPassword,
      sshPort: app.sshPort,
      sshUser: app.sshUser,
      token: app.token
    })
    this.nd.instanceForm.get('appType').disable({onlySelf: false, emitEvent: false})
  }
  deleteApp (app: AppModel) {
    this.sureDeleteFalg = true
    this.token = app.token
  }
  getCurrentAppLogs (app: AppModel) {
    this.nd.getAppLogs(app.token)
    this.router.navigate(['/cloudtides/notebooks/log'])
  }
  sure () {
    this.nd.deleteApp(this.token).subscribe(data => {
      console.log(data);
      this.sureDeleteFalg = false
      this.nd.getApplictionList()
    })
  }
  cancel () {
    this.sureDeleteFalg = false
    this.token = ''
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