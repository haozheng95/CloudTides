import { Component, OnInit, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
import { NotebooksService } from './notebooks.service'
import { FormBuilder, Validators } from '@angular/forms'

@Component({
  selector: 'tide-notebooks',
  templateUrl: './notebooks.component.html',
  styleUrls: ['./notebooks.component.scss']
})
export class NotebooksComponent implements OnInit {

  constructor(public router: Router, private nb: NotebooksService, private fb: FormBuilder) {
    this.noteBook = this.nb
    this.instanceForm = this.nb.instanceForm
  }
  ngOnInit(): void {
  }
  noteBook: NotebooksService
  // error tip flag
  flag:boolean = false
  errorMsg = ''
  // create disable for get
  get disabled () {
    const arr = []
    for (const key in this.instanceForm.controls) {
      if (this.instanceForm.controls[key].value !== 'jupyter' && this.instanceForm.controls[key].value !== 'gromacs') {
        arr.push(this.instanceForm.controls[key].value)
      }
    }
    const rsg = arr.every(el => el !== '')
    return !rsg
  }
  instanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: ['jupyter'],
    sshHost: [''],
    sshPassword: [''],
    sshPort: [''],
    sshUser: [''],
    token: ['']
  })
  cancel () {    
    this.noteBook.createInstanceFlag = false
    this.instanceForm.setValue({
      instanceName: '',
      port: '',
      appType: '',
      sshHost: '',
      sshPassword: '',
      sshPort: '',
      sshUser: '',
      token: ''
    })
    this.errorMsg = ''
    this.flag = false

  }
  create (form) {
    this.errorMsg = ''
    this.flag = true
    const str = 'HOME.APPLICATION.Create'
    const data = form.getRawValue()
    data.sshPort = +data.sshPort
    if (this.nb.createInstanceTitle === str) {
      delete data.token
      this.nb.createNewApp(data).subscribe((data:CreateData) => {
        this.flag = false
        this.noteBook.createInstanceFlag = false
        this.instanceForm.setValue({
          instanceName: '',
          port: '',
          appType: '',
          sshHost: '',
          sshPassword: '',
          sshPort: '',
          sshUser: ''
        })
        window.open('http://' + data.link, "_blank")
      },
      err => {
        console.log('err', err);
        
        this.errorMsg = err.statusText
        this.flag = false
      })
    } else {
      this.nb.modifyApp(data).subscribe(data => {
        this.flag = false
        this.noteBook.createInstanceFlag = false
        this.nb.getApplictionList()
      },
      err => {
        this.errorMsg = err.msg
        this.flag = false
      })
    }
  }
}
interface CreateData {
  link: string
  token: string
}