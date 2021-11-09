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
    this.jupyterInstanceForm = this.nb.instanceForm
  }
  ngOnInit(): void {
    this.gromacsInstanceForm.get('appType').disable({
      onlySelf: true
    })
  }
  noteBook: NotebooksService
  // error tip flag
  loadingFlag:boolean = false
  errorMsg = ''
  // create disable for get
  get disabled () {
    const arr = []
    if (this.noteBook.currentModel === 'jupyter') {
      for (const key in this.jupyterInstanceForm.value) {
        if (key !== 'token') {
          arr.push(this.jupyterInstanceForm.value[key])
        }
      }
    } else if (this.noteBook.currentModel === 'gromacs') {
      for (const key in this.gromacsInstanceForm.value) {
        if (key !== 'token') {
          arr.push(this.gromacsInstanceForm.value[key])
        }
      }
    }
    const rsg = arr.every(el => el !== '')    
    return !rsg
  }
  jupyterInstanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    port: ['', Validators.required],
    appType: ['jupyter'],
    sshHost: [''],
    sshPassword: [''],
    sshPort: [''],
    sshUser: [''],
    token: ['']
  })
  gromacsInstanceForm = this.fb.group({
    instanceName: ['', Validators.required],
    cpu: ['', Validators.required],
    version: ['', Validators.required],
    appType: ['gromacs'],
    num: [1]
  })
  cancel () {    
    this.noteBook.createInstanceFlag = false
    this.jupyterInstanceForm.setValue({
      instanceName: '',
      port: '',
      appType: '',
      sshHost: '',
      sshPassword: '',
      sshPort: '',
      sshUser: '',
      token: ''
    })
    this.gromacsInstanceForm.setValue({
      instanceName: '',
      cpu: ['', Validators.required],
      version: ['', Validators.required],
      appType: ['gromacs'],
      num: [1]
      })
    this.errorMsg = ''
    this.loadingFlag = false

  }
  create (form) {
    this.errorMsg = ''
    this.loadingFlag = true
    const data = form.getRawValue()
    if (data.appType === 'jupyter') {
      const str = 'HOME.APPLICATION.Create'
      data.sshPort = +data.sshPort
      if (this.nb.createInstanceTitle === str) {
        delete data.token
        this.nb.createNewApp(data).subscribe((data:CreateData) => {
          this.loadingFlag = false
          this.noteBook.createInstanceFlag = false
          this.jupyterInstanceForm.setValue({
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
          this.loadingFlag = false
        })
      } else {
        this.nb.modifyApp(data).subscribe(data => {
          this.loadingFlag = false
          this.noteBook.createInstanceFlag = false
          this.nb.getApplictionList()
        },
        err => {
          this.errorMsg = err.msg
          this.loadingFlag = false
        })
      }
    } else if (data.appType === 'gromacs') {
      console.log('创建 gromace');
      this.loadingFlag = false
    }
  }
}
interface CreateData {
  link: string
  token: string
}