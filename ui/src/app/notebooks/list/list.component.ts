import { Component, OnInit } from '@angular/core';
import { NotebooksService } from '../notebooks.service'
@Component({
  selector: 'tide-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {

  constructor(private nd: NotebooksService) {
  }

  ngOnInit(): void {
    this.getApplictionList()
  }
  getApplictionList () {
    this.nd.getAppList().subscribe((data:AppModel[]) => {
      this.appList = data
    })
  }
  appList:AppModel[] = [
    {
      name: 'Jupyter',
      logo: 'assets/img/jupyter.svg',
      token: '1dc53b34f46aff0f91f8c65ec96f55eb3057d3770e2253b8',
      link: "http://120.133.15.12:8888/lab"
    }
  ]
  selected= 'selected'
  sureDeleteFalg = false
  filterSearchValue: string = ''
  token = ''
  toJupyter (app: AppModel) {
    let form = document.createElement('form')
    form.action= app.link
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
  modifyApp (app: AppModel) {
    this.nd.createInstanceFlag = true
    this.nd.createInstanceTitle = 'HOME.APPLICATION.Modify'
  }
  deleteApp (app: AppModel) {
    this.sureDeleteFalg = true
    this.token = app.token
  }
  sure () {
    this.nd.deleteApp(this.token).subscribe(data => {
      console.log(data);
      this.sureDeleteFalg = false
    })
  }
  cancel () {
    this.sureDeleteFalg = false
    this.token = ''
  }
}
interface AppModel {
  name: string
  link: string
  token: string
  logo: string
}