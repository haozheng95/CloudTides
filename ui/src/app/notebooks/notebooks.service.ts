import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { environment } from '@tide-environments/environment';
import { Router } from '@angular/router'
import { tap } from 'rxjs/operators';
import { FormBuilder, Validators } from '@angular/forms'
import { LOCAL_STORAGE_KEY } from '@tide-config/const';
@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor(private readonly http: HttpClient,private readonly router: Router, private fb: FormBuilder) { }
  createInstanceFlag:boolean = false
  modifiable = true
  instanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: [''],
    sshHost: [''],
    sshPassword: [''],
    sshPort: [''],
    sshUser: ['']
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
}
