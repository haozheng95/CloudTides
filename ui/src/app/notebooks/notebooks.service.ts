import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { environment } from '@tide-environments/environment';
import { Router } from '@angular/router'
import { tap } from 'rxjs/operators';
@Injectable({
  providedIn: 'root'
})
export class NotebooksService {

  constructor(private readonly http: HttpClient,private readonly router: Router) { }
  createInstanceFlag:boolean = false
  instanceForm: any= {}
  createInstanceTitle = 'HOME.NOTEBOOKS.Create'
  getAppList () {
    return this.http.get(environment.apiPrefix + '/application/instance').pipe(
      tap(data => {})
    )
  }
  createNewApp () {
    return this.http.post(environment.apiPrefix + '/application/instance', this.instanceForm).pipe(
      tap(data => {})
    )
  }
  modifyApp () {
    return this.http.put(environment.apiPrefix + '/application/instance', this.instanceForm).pipe(
      tap(data => {})
    )
  }
  deleteApp (token) {
    return this.http.delete(environment.apiPrefix + `/application/instance/${token}`).pipe(
      tap(data => {})
    )
  }
}
